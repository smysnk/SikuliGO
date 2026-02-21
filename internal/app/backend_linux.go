//go:build linux

package app

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/sikulix/portgo/internal/core"
)

type linuxBackend struct {
	runner commandRunner
}

func New() core.App {
	return &linuxBackend{
		runner: execRunner{},
	}
}

func (b *linuxBackend) Execute(req core.AppRequest) (core.AppResult, error) {
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
		out, err := b.runner.Run(ctx, req.Name, req.Args...)
		if err != nil {
			return core.AppResult{}, commandError("open", err, out)
		}
		return core.AppResult{}, nil
	case core.AppActionFocus:
		out, err := b.runner.Run(ctx, "wmctrl", "-a", req.Name)
		if err != nil {
			return core.AppResult{}, commandError("focus", err, out)
		}
		return core.AppResult{}, nil
	case core.AppActionClose:
		out, err := b.runner.Run(ctx, "pkill", "-f", req.Name)
		if err != nil {
			return core.AppResult{}, commandError("close", err, out)
		}
		return core.AppResult{}, nil
	case core.AppActionIsRunning:
		out, err := b.runner.Run(ctx, "pgrep", "-f", req.Name)
		if err != nil {
			// pgrep returns non-zero when no match.
			if strings.TrimSpace(out) == "" {
				return core.AppResult{Running: false}, nil
			}
			return core.AppResult{}, commandError("is-running", err, out)
		}
		return core.AppResult{Running: strings.TrimSpace(out) != ""}, nil
	case core.AppActionListWindow:
		return b.listWindows(ctx, req.Name)
	default:
		return core.AppResult{}, fmt.Errorf("unsupported app action %q", req.Action)
	}
}

func (b *linuxBackend) listWindows(ctx context.Context, name string) (core.AppResult, error) {
	runningResult, err := b.Execute(core.AppRequest{
		Action: core.AppActionIsRunning,
		Name:   name,
	})
	if err != nil {
		return core.AppResult{}, err
	}
	if !runningResult.Running {
		return core.AppResult{Running: false, Windows: nil}, nil
	}

	out, err := b.runner.Run(ctx, "wmctrl", "-lxG")
	if err != nil {
		return core.AppResult{}, commandError("list-windows", err, out)
	}
	windows, parseErr := parseLinuxWindowRows(out, name)
	if parseErr != nil {
		return core.AppResult{}, parseErr
	}
	return core.AppResult{
		Running: true,
		Windows: windows,
	}, nil
}

func parseLinuxWindowRows(out, needle string) ([]core.WindowInfo, error) {
	lines := strings.Split(strings.TrimSpace(out), "\n")
	windows := make([]core.WindowInfo, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 9 {
			continue
		}
		x, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, fmt.Errorf("linux list-windows parse x failed: %w", err)
		}
		y, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil, fmt.Errorf("linux list-windows parse y failed: %w", err)
		}
		w, err := strconv.Atoi(parts[4])
		if err != nil {
			return nil, fmt.Errorf("linux list-windows parse w failed: %w", err)
		}
		h, err := strconv.Atoi(parts[5])
		if err != nil {
			return nil, fmt.Errorf("linux list-windows parse h failed: %w", err)
		}
		title := strings.Join(parts[8:], " ")
		if !strings.Contains(strings.ToLower(title), strings.ToLower(needle)) {
			continue
		}
		windows = append(windows, core.WindowInfo{
			Title:   title,
			X:       x,
			Y:       y,
			W:       w,
			H:       h,
			Focused: false,
		})
	}
	return windows, nil
}
