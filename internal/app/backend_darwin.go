//go:build darwin

package app

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
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
}

func New() core.App {
	return &darwinBackend{
		runner: execRunner{},
	}
}

func (b *darwinBackend) Execute(req core.AppRequest) (core.AppResult, error) {
	if err := req.Validate(); err != nil {
		return core.AppResult{}, err
	}
	if b == nil || b.runner == nil {
		return core.AppResult{}, fmt.Errorf("%w: backend not initialized", core.ErrAppUnsupported)
	}

	ctx, cancel := contextForTimeout(req.Timeout)
	defer cancel()

	switch req.Action {
	case core.AppActionOpen:
		if err := b.open(ctx, req.Name, req.Args); err != nil {
			return core.AppResult{}, err
		}
		return core.AppResult{}, nil
	case core.AppActionFocus:
		if err := b.focus(ctx, req.Name); err != nil {
			return core.AppResult{}, err
		}
		return core.AppResult{}, nil
	case core.AppActionClose:
		if err := b.close(ctx, req.Name); err != nil {
			return core.AppResult{}, err
		}
		return core.AppResult{}, nil
	case core.AppActionIsRunning:
		running, err := b.isRunning(ctx, req.Name)
		if err != nil {
			return core.AppResult{}, err
		}
		return core.AppResult{Running: running}, nil
	case core.AppActionListWindow:
		return b.listWindows(ctx, req.Name)
	default:
		return core.AppResult{}, fmt.Errorf("unsupported app action %q", req.Action)
	}
}

func (b *darwinBackend) open(ctx context.Context, name string, args []string) error {
	cmdArgs := []string{"-a", name}
	if len(args) > 0 {
		cmdArgs = append(cmdArgs, "--args")
		cmdArgs = append(cmdArgs, args...)
	}
	out, err := b.runner.Run(ctx, "open", cmdArgs...)
	if err != nil {
		return commandError("open", err, out)
	}
	return nil
}

func (b *darwinBackend) focus(ctx context.Context, name string) error {
	script := fmt.Sprintf(`tell application %s to activate`, strconv.Quote(name))
	out, err := b.runner.Run(ctx, "osascript", "-e", script)
	if err != nil {
		return commandError("focus", err, out)
	}
	return nil
}

func (b *darwinBackend) close(ctx context.Context, name string) error {
	script := fmt.Sprintf(`tell application %s to quit`, strconv.Quote(name))
	out, err := b.runner.Run(ctx, "osascript", "-e", script)
	if err != nil {
		return commandError("close", err, out)
	}
	return nil
}

func (b *darwinBackend) isRunning(ctx context.Context, name string) (bool, error) {
	script := fmt.Sprintf(`tell application "System Events" to (name of application processes) contains %s`, strconv.Quote(name))
	out, err := b.runner.Run(ctx, "osascript", "-e", script)
	if err != nil {
		return false, commandError("is-running", err, out)
	}
	parsed, parseErr := parseBoolString(out)
	if parseErr != nil {
		return false, fmt.Errorf("is-running parse failed: %w", parseErr)
	}
	return parsed, nil
}

func (b *darwinBackend) listWindows(ctx context.Context, name string) (core.AppResult, error) {
	running, err := b.isRunning(ctx, name)
	if err != nil {
		return core.AppResult{}, err
	}
	if !running {
		return core.AppResult{Running: false, Windows: nil}, nil
	}

	script := fmt.Sprintf(`
tell application "System Events"
	set appName to %s
	set matched to application processes whose name is appName
	if (count of matched) is 0 then
		return ""
	end if
	set proc to item 1 of matched
	set focusedState to frontmost of proc as string
	set rows to {}
	repeat with w in windows of proc
		set winTitle to ""
		try
			set winTitle to name of w as string
		end try
		set xPos to 0
		set yPos to 0
		try
			set {xPos, yPos} to position of w
		end try
		set wSize to 0
		set hSize to 0
		try
			set {wSize, hSize} to size of w
		end try
		set row to winTitle & "||" & (xPos as string) & "||" & (yPos as string) & "||" & (wSize as string) & "||" & (hSize as string) & "||" & focusedState
		copy row to end of rows
	end repeat
	set AppleScript's text item delimiters to linefeed
	return rows as text
end tell`, strconv.Quote(name))

	out, err := b.runner.Run(ctx, "osascript", "-e", script)
	if err != nil {
		return core.AppResult{}, commandError("list-windows", err, out)
	}

	windows, parseErr := parseWindowsOutput(out)
	if parseErr != nil {
		return core.AppResult{}, parseErr
	}
	return core.AppResult{
		Running: true,
		Windows: windows,
	}, nil
}

func contextForTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	if timeout > 0 {
		return context.WithTimeout(context.Background(), timeout)
	}
	return context.WithCancel(context.Background())
}

func parseBoolString(s string) (bool, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "true":
		return true, nil
	case "false", "":
		return false, nil
	default:
		return false, fmt.Errorf("unexpected bool value %q", strings.TrimSpace(s))
	}
}

func parseWindowsOutput(s string) ([]core.WindowInfo, error) {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	windows := make([]core.WindowInfo, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, "||")
		if len(parts) != 6 {
			return nil, fmt.Errorf("window parse failed for line %q", line)
		}
		x, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, fmt.Errorf("window parse x failed: %w", err)
		}
		y, err := strconv.Atoi(strings.TrimSpace(parts[2]))
		if err != nil {
			return nil, fmt.Errorf("window parse y failed: %w", err)
		}
		w, err := strconv.Atoi(strings.TrimSpace(parts[3]))
		if err != nil {
			return nil, fmt.Errorf("window parse w failed: %w", err)
		}
		h, err := strconv.Atoi(strings.TrimSpace(parts[4]))
		if err != nil {
			return nil, fmt.Errorf("window parse h failed: %w", err)
		}
		focused, err := parseBoolString(parts[5])
		if err != nil {
			return nil, fmt.Errorf("window parse focused failed: %w", err)
		}
		windows = append(windows, core.WindowInfo{
			Title:   parts[0],
			X:       x,
			Y:       y,
			W:       w,
			H:       h,
			Focused: focused,
		})
	}
	return windows, nil
}

func commandError(action string, err error, output string) error {
	trimmed := strings.TrimSpace(output)
	if trimmed == "" {
		return fmt.Errorf("%s app action failed: %w", action, err)
	}
	return fmt.Errorf("%s app action failed: %w: %s", action, err, trimmed)
}
