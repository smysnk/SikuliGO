//go:build linux

package input

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/smysnk/sikuligo/internal/core"
)

type linuxBackend struct {
	runner commandRunner
	sleep  func(time.Duration)
}

func New() core.Input {
	return &linuxBackend{
		runner: execRunner{},
		sleep:  time.Sleep,
	}
}

func (b *linuxBackend) Execute(req core.InputRequest) error {
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
		return b.runXdotool(ctx, "mousemove", strconv.Itoa(req.X), strconv.Itoa(req.Y))
	case core.InputActionClick:
		code, err := linuxMouseButtonCode(req.Button)
		if err != nil {
			return err
		}
		return b.runXdotool(ctx, "mousemove", strconv.Itoa(req.X), strconv.Itoa(req.Y), "click", code)
	case core.InputActionMouseDown:
		code, err := linuxMouseButtonCode(req.Button)
		if err != nil {
			return err
		}
		return b.runXdotool(ctx, "mousemove", strconv.Itoa(req.X), strconv.Itoa(req.Y), "mousedown", code)
	case core.InputActionMouseUp:
		code, err := linuxMouseButtonCode(req.Button)
		if err != nil {
			return err
		}
		return b.runXdotool(ctx, "mousemove", strconv.Itoa(req.X), strconv.Itoa(req.Y), "mouseup", code)
	case core.InputActionTypeText:
		return b.runXdotool(ctx, "type", "--", req.Text)
	case core.InputActionPasteText:
		return b.runXdotool(ctx, "type", "--", req.Text)
	case core.InputActionHotkey:
		chord, err := linuxHotkeyChord(req.Keys)
		if err != nil {
			return err
		}
		return b.runXdotool(ctx, "key", chord)
	case core.InputActionKeyDown:
		for _, key := range req.Keys {
			norm := normalizeLinuxKey(key)
			if norm == "" {
				return fmt.Errorf("key down requires non-empty keys")
			}
			if err := b.runXdotool(ctx, "keydown", norm); err != nil {
				return err
			}
		}
		return nil
	case core.InputActionKeyUp:
		for _, key := range req.Keys {
			norm := normalizeLinuxKey(key)
			if norm == "" {
				return fmt.Errorf("key up requires non-empty keys")
			}
			if err := b.runXdotool(ctx, "keyup", norm); err != nil {
				return err
			}
		}
		return nil
	case core.InputActionWheel:
		code, err := linuxWheelButtonCode(req.ScrollDirection)
		if err != nil {
			return err
		}
		return b.runXdotool(ctx, "mousemove", strconv.Itoa(req.X), strconv.Itoa(req.Y), "click", "--repeat", strconv.Itoa(req.ScrollSteps), code)
	default:
		return fmt.Errorf("unsupported input action %q", req.Action)
	}
}

func (b *linuxBackend) runXdotool(ctx context.Context, args ...string) error {
	out, err := b.runner.Run(ctx, "xdotool", args...)
	if err != nil {
		return inputCommandError("xdotool", err, out)
	}
	return nil
}

func linuxMouseButtonCode(button string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(button)) {
	case "left":
		return "1", nil
	case "middle":
		return "2", nil
	case "right":
		return "3", nil
	default:
		return "", fmt.Errorf("click requires supported button, got %q", button)
	}
}

func linuxHotkeyChord(keys []string) (string, error) {
	if len(keys) == 0 {
		return "", fmt.Errorf("hotkey requires at least one key")
	}
	mods := make([]string, 0, len(keys))
	modSeen := map[string]struct{}{}
	mainKey := ""
	for _, key := range keys {
		norm := normalizeLinuxKey(key)
		switch norm {
		case "ctrl", "shift", "alt", "super":
			if _, ok := modSeen[norm]; ok {
				continue
			}
			modSeen[norm] = struct{}{}
			mods = append(mods, norm)
		default:
			if norm != "" {
				mainKey = norm
			}
		}
	}
	if mainKey == "" {
		return "", fmt.Errorf("hotkey requires one non-modifier key")
	}
	parts := append(mods, mainKey)
	return strings.Join(parts, "+"), nil
}

func linuxWheelButtonCode(direction string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(direction)) {
	case "up":
		return "4", nil
	case "down":
		return "5", nil
	case "left":
		return "6", nil
	case "right":
		return "7", nil
	default:
		return "", fmt.Errorf("wheel requires supported direction, got %q", direction)
	}
}

func normalizeLinuxKey(key string) string {
	raw := strings.TrimSpace(key)
	if raw == "" {
		return ""
	}
	lower := strings.ToLower(raw)
	switch lower {
	case "cmd", "command", "win", "windows", "super":
		return "super"
	case "ctrl", "control":
		return "ctrl"
	case "alt", "option":
		return "alt"
	case "shift":
		return "shift"
	case "enter", "return":
		return "Return"
	case "tab":
		return "Tab"
	case "esc", "escape":
		return "Escape"
	case "backspace":
		return "BackSpace"
	case "delete", "del":
		return "Delete"
	case "space":
		return "space"
	case "up", "arrowup", "up arrow":
		return "Up"
	case "down", "arrowdown", "down arrow":
		return "Down"
	case "left", "arrowleft", "left arrow":
		return "Left"
	case "right", "arrowright", "right arrow":
		return "Right"
	case "home":
		return "Home"
	case "end":
		return "End"
	case "pageup", "page_up", "page up":
		return "Page_Up"
	case "pagedown", "page_down", "page down":
		return "Page_Down"
	}
	if len(raw) == 1 {
		return strings.ToLower(raw)
	}
	if len(lower) > 1 && strings.HasPrefix(lower, "f") {
		isFn := true
		for _, r := range lower[1:] {
			if r < '0' || r > '9' {
				isFn = false
				break
			}
		}
		if isFn {
			return strings.ToUpper(lower)
		}
	}
	lower = strings.ReplaceAll(lower, " ", "_")
	lower = strings.ReplaceAll(lower, "-", "_")
	return lower
}
