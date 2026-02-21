//go:build darwin

package input

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/sikulix/portgo/internal/core"
)

type commandRunner interface {
	Run(ctx context.Context, name string, args ...string) (string, error)
}

type execRunner struct{}

func (r execRunner) Run(ctx context.Context, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

type darwinBackend struct {
	runner commandRunner
	sleep  func(time.Duration)
}

func New() core.Input {
	return &darwinBackend{
		runner: execRunner{},
		sleep:  time.Sleep,
	}
}

func (b *darwinBackend) Execute(req core.InputRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	if b == nil || b.runner == nil {
		return fmt.Errorf("%w: backend not initialized", core.ErrInputUnsupported)
	}
	if req.Delay > 0 {
		if b.sleep != nil {
			b.sleep(req.Delay)
		} else {
			time.Sleep(req.Delay)
		}
	}

	ctx := context.Background()
	switch req.Action {
	case core.InputActionMouseMove:
		return b.runCliclick(ctx, fmt.Sprintf("m:%d,%d", req.X, req.Y))
	case core.InputActionClick:
		prefix := "c"
		switch strings.ToLower(strings.TrimSpace(req.Button)) {
		case "right":
			prefix = "rc"
		case "middle":
			prefix = "mc"
		}
		return b.runCliclick(ctx, fmt.Sprintf("%s:%d,%d", prefix, req.X, req.Y))
	case core.InputActionTypeText:
		return b.runOSA(ctx, fmt.Sprintf(`tell application "System Events" to keystroke %s`, asQuote(req.Text)))
	case core.InputActionHotkey:
		return b.runHotkey(ctx, req.Keys)
	default:
		return fmt.Errorf("unsupported input action %q", req.Action)
	}
}

func (b *darwinBackend) runHotkey(ctx context.Context, keys []string) error {
	if len(keys) == 0 {
		return fmt.Errorf("hotkey requires at least one key")
	}
	mods := make([]string, 0, len(keys))
	mainKey := ""
	for _, key := range keys {
		switch strings.ToLower(strings.TrimSpace(key)) {
		case "cmd", "command":
			mods = append(mods, "command down")
		case "ctrl", "control":
			mods = append(mods, "control down")
		case "shift":
			mods = append(mods, "shift down")
		case "alt", "option":
			mods = append(mods, "option down")
		default:
			mainKey = key
		}
	}
	if strings.TrimSpace(mainKey) == "" {
		return fmt.Errorf("hotkey requires one non-modifier key")
	}
	script := fmt.Sprintf(`tell application "System Events" to keystroke %s`, asQuote(mainKey))
	if len(mods) > 0 {
		script += " using {" + strings.Join(mods, ", ") + "}"
	}
	return b.runOSA(ctx, script)
}

func (b *darwinBackend) runCliclick(ctx context.Context, action string) error {
	out, err := b.runner.Run(ctx, "cliclick", action)
	if err != nil {
		return inputCommandError("cliclick", err, out)
	}
	return nil
}

func (b *darwinBackend) runOSA(ctx context.Context, script string) error {
	out, err := b.runner.Run(ctx, "osascript", "-e", script)
	if err != nil {
		return inputCommandError("osascript", err, out)
	}
	return nil
}

func asQuote(s string) string {
	return `"` + strings.ReplaceAll(strings.ReplaceAll(s, `\`, `\\`), `"`, `\"`) + `"`
}

func inputCommandError(tool string, err error, output string) error {
	trimmed := strings.TrimSpace(output)
	if trimmed == "" {
		return fmt.Errorf("%s input action failed: %w", tool, err)
	}
	return fmt.Errorf("%s input action failed: %w: %s", tool, err, trimmed)
}
