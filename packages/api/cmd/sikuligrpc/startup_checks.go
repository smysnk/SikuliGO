package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type startupState struct {
	SuppressCliclickPrompt bool `json:"suppress_cliclick_prompt"`
}

func runStartupChecks(stderr io.Writer) {
	if runtime.GOOS != "darwin" {
		return
	}
	if _, err := exec.LookPath("cliclick"); err == nil {
		return
	}

	statePath := startupStatePath()
	state, err := loadStartupState(statePath)
	if err != nil {
		fmt.Fprintf(stderr, "sikuli-go: warning: could not read startup state %q: %v\n", statePath, err)
	}

	fmt.Fprintln(
		stderr,
		`sikuli-go: warning: missing "cliclick" in PATH (required for mouse automation on macOS).`,
	)

	if state.SuppressCliclickPrompt {
		return
	}
	if !isInteractiveTerminal() {
		return
	}

	installNow, err := promptInstallCliclick(os.Stdin, stderr)
	if err != nil {
		fmt.Fprintf(stderr, "sikuli-go: warning: prompt failed: %v\n", err)
		return
	}
	if !installNow {
		state.SuppressCliclickPrompt = true
		if err := saveStartupState(statePath, state); err != nil {
			fmt.Fprintf(stderr, "sikuli-go: warning: could not persist startup state: %v\n", err)
		} else {
			fmt.Fprintf(stderr, "sikuli-go: warning: not prompting again (remove %s to re-enable prompt)\n", statePath)
		}
		return
	}

	if err := installCliclick(stderr); err != nil {
		fmt.Fprintf(stderr, "sikuli-go: warning: %v\n", err)
		return
	}
	fmt.Fprintln(stderr, `sikuli-go: cliclick install complete.`)
}

func promptInstallCliclick(stdin io.Reader, stderr io.Writer) (bool, error) {
	fmt.Fprint(stderr, `Install "cliclick" now with Homebrew? [Y/n]: `)
	reader := bufio.NewReader(stdin)
	line, err := reader.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return false, err
	}
	answer := strings.TrimSpace(strings.ToLower(line))
	switch answer {
	case "", "y", "yes":
		return true, nil
	case "n", "no":
		return false, nil
	default:
		// Keep startup non-blocking: unknown input means no install and no persistence.
		fmt.Fprintln(stderr, `sikuli-go: warning: unrecognized answer, skipping install prompt for this run`)
		return false, nil
	}
}

func installCliclick(stderr io.Writer) error {
	if _, err := exec.LookPath("brew"); err != nil {
		return errors.New(`Homebrew not found in PATH; install Homebrew or run "brew install cliclick" manually`)
	}
	cmd := exec.Command("brew", "install", "cliclick")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf(`failed to install cliclick using brew: %w`, err)
	}
	if _, err := exec.LookPath("cliclick"); err != nil {
		return errors.New(`cliclick still not found in PATH after install; ensure /opt/homebrew/bin or /usr/local/bin is on PATH`)
	}
	return nil
}

func startupStatePath() string {
	dir := os.Getenv("XDG_CONFIG_HOME")
	if strings.TrimSpace(dir) == "" {
		home, err := os.UserHomeDir()
		if err != nil || strings.TrimSpace(home) == "" {
			return ".sikuligo-startup-state.json"
		}
		dir = filepath.Join(home, ".config")
	}
	return filepath.Join(dir, "sikuligo", "startup-state.json")
}

func loadStartupState(path string) (startupState, error) {
	var state startupState
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return state, nil
		}
		return state, err
	}
	if len(data) == 0 {
		return state, nil
	}
	if err := json.Unmarshal(data, &state); err != nil {
		return startupState{}, err
	}
	return state, nil
}

func saveStartupState(path string, state startupState) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	payload, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	payload = append(payload, '\n')
	return os.WriteFile(path, payload, 0o644)
}

func isInteractiveTerminal() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (info.Mode() & os.ModeCharDevice) != 0
}
