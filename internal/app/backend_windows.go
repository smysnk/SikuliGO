//go:build windows

package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/sikulix/portgo/internal/core"
)

type windowsBackend struct {
	runner commandRunner
}

func New() core.App {
	return &windowsBackend{
		runner: execRunner{},
	}
}

func (b *windowsBackend) Execute(req core.AppRequest) (core.AppResult, error) {
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
		return b.open(ctx, req.Name, req.Args)
	case core.AppActionFocus:
		return b.focus(ctx, req.Name)
	case core.AppActionClose:
		return b.close(ctx, req.Name)
	case core.AppActionIsRunning:
		return b.isRunning(ctx, req.Name)
	case core.AppActionListWindow:
		return b.listWindows(ctx, req.Name)
	default:
		return core.AppResult{}, fmt.Errorf("unsupported app action %q", req.Action)
	}
}

func (b *windowsBackend) open(ctx context.Context, name string, args []string) (core.AppResult, error) {
	script := fmt.Sprintf(`Start-Process -FilePath %s -ArgumentList %s`, psQuote(name), psArray(args))
	out, err := b.runner.Run(ctx, "powershell", "-NoProfile", "-Command", script)
	if err != nil {
		return core.AppResult{}, commandError("open", err, out)
	}
	return core.AppResult{}, nil
}

func (b *windowsBackend) focus(ctx context.Context, name string) (core.AppResult, error) {
	// Focusing is best-effort via process main window activation.
	script := fmt.Sprintf(`$p = Get-Process | Where-Object { $_.ProcessName -like %s } | Select-Object -First 1; if ($null -eq $p) { exit 0 }; [void]$p.MainWindowHandle`, psQuote(name+"*"))
	out, err := b.runner.Run(ctx, "powershell", "-NoProfile", "-Command", script)
	if err != nil {
		return core.AppResult{}, commandError("focus", err, out)
	}
	return core.AppResult{}, nil
}

func (b *windowsBackend) close(ctx context.Context, name string) (core.AppResult, error) {
	script := fmt.Sprintf(`Get-Process | Where-Object { $_.ProcessName -like %s } | Stop-Process -Force`, psQuote(name+"*"))
	out, err := b.runner.Run(ctx, "powershell", "-NoProfile", "-Command", script)
	if err != nil {
		return core.AppResult{}, commandError("close", err, out)
	}
	return core.AppResult{}, nil
}

func (b *windowsBackend) isRunning(ctx context.Context, name string) (core.AppResult, error) {
	script := fmt.Sprintf(`$n = (Get-Process | Where-Object { $_.ProcessName -like %s }).Count; if ($n -gt 0) { "true" } else { "false" }`, psQuote(name+"*"))
	out, err := b.runner.Run(ctx, "powershell", "-NoProfile", "-Command", script)
	if err != nil {
		return core.AppResult{}, commandError("is-running", err, out)
	}
	running, parseErr := parseBoolString(out)
	if parseErr != nil {
		return core.AppResult{}, parseErr
	}
	return core.AppResult{Running: running}, nil
}

func (b *windowsBackend) listWindows(ctx context.Context, name string) (core.AppResult, error) {
	script := fmt.Sprintf(`Get-Process | Where-Object { $_.ProcessName -like %s } | ForEach-Object { $title = $_.MainWindowTitle; if ([string]::IsNullOrWhiteSpace($title)) { $title = $_.ProcessName }; "$title||0||0||0||0||false" }`, psQuote(name+"*"))
	out, err := b.runner.Run(ctx, "powershell", "-NoProfile", "-Command", script)
	if err != nil {
		return core.AppResult{}, commandError("list-windows", err, out)
	}
	windows, parseErr := parseWindowsOutput(out)
	if parseErr != nil {
		return core.AppResult{}, parseErr
	}
	return core.AppResult{
		Running: len(windows) > 0,
		Windows: windows,
	}, nil
}

func psQuote(v string) string {
	return "'" + strings.ReplaceAll(v, "'", "''") + "'"
}

func psArray(items []string) string {
	if len(items) == 0 {
		return "@()"
	}
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, psQuote(item))
	}
	return "@(" + strings.Join(out, ",") + ")"
}
