//go:build windows

package input

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sikulix/portgo/internal/core"
)

type windowsBackend struct {
	runner commandRunner
	sleep  func(time.Duration)
}

func New() core.Input {
	return &windowsBackend{
		runner: execRunner{},
		sleep:  time.Sleep,
	}
}

func (b *windowsBackend) Execute(req core.InputRequest) error {
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
		return b.runPowerShell(ctx, windowsSetCursorScript(req.X, req.Y))
	case core.InputActionClick:
		script, err := windowsClickScript(req.X, req.Y, req.Button)
		if err != nil {
			return err
		}
		return b.runPowerShell(ctx, script)
	case core.InputActionTypeText:
		return b.runPowerShell(ctx, windowsSendKeysScript(windowsSendKeysEscapeText(req.Text)))
	case core.InputActionHotkey:
		chord, err := windowsHotkeyChord(req.Keys)
		if err != nil {
			return err
		}
		return b.runPowerShell(ctx, windowsSendKeysScript(chord))
	default:
		return fmt.Errorf("unsupported input action %q", req.Action)
	}
}

func (b *windowsBackend) runPowerShell(ctx context.Context, script string) error {
	out, err := b.runner.Run(ctx, "powershell", "-NoProfile", "-Command", script)
	if err != nil {
		return inputCommandError("powershell", err, out)
	}
	return nil
}

func windowsClickScript(x, y int, button string) (string, error) {
	var down uint32
	var up uint32
	switch strings.ToLower(strings.TrimSpace(button)) {
	case "left":
		down = 0x0002
		up = 0x0004
	case "right":
		down = 0x0008
		up = 0x0010
	case "middle":
		down = 0x0020
		up = 0x0040
	default:
		return "", fmt.Errorf("click requires supported button, got %q", button)
	}
	return windowsSetCursorScript(x, y) + fmt.Sprintf(`; [SikuliGoInput]::mouse_event(0x%04x, 0, 0, 0, [UIntPtr]::Zero); [SikuliGoInput]::mouse_event(0x%04x, 0, 0, 0, [UIntPtr]::Zero)`, down, up), nil
}

func windowsSetCursorScript(x, y int) string {
	return fmt.Sprintf(`$sig = @'
using System;
using System.Runtime.InteropServices;
public static class SikuliGoInput {
  [DllImport("user32.dll")]
  public static extern bool SetCursorPos(int X, int Y);
  [DllImport("user32.dll")]
  public static extern void mouse_event(uint dwFlags, uint dx, uint dy, uint dwData, UIntPtr dwExtraInfo);
}
'@; if (-not ("SikuliGoInput" -as [type])) { Add-Type -TypeDefinition $sig }; [SikuliGoInput]::SetCursorPos(%d, %d) | Out-Null`, x, y)
}

func windowsSendKeysScript(chord string) string {
	return fmt.Sprintf(`Add-Type -AssemblyName System.Windows.Forms; [System.Windows.Forms.SendKeys]::SendWait(%s)`, psQuote(chord))
}

func windowsHotkeyChord(keys []string) (string, error) {
	if len(keys) == 0 {
		return "", fmt.Errorf("hotkey requires at least one key")
	}
	mods := make([]string, 0, len(keys))
	seenMods := map[string]struct{}{}
	main := ""
	for _, key := range keys {
		switch strings.ToLower(strings.TrimSpace(key)) {
		case "ctrl", "control", "cmd", "command":
			if _, ok := seenMods["^"]; !ok {
				seenMods["^"] = struct{}{}
				mods = append(mods, "^")
			}
		case "shift":
			if _, ok := seenMods["+"]; !ok {
				seenMods["+"] = struct{}{}
				mods = append(mods, "+")
			}
		case "alt", "option":
			if _, ok := seenMods["%"]; !ok {
				seenMods["%"] = struct{}{}
				mods = append(mods, "%")
			}
		default:
			token := windowsHotkeyToken(key)
			if token != "" {
				main = token
			}
		}
	}
	if main == "" {
		return "", fmt.Errorf("hotkey requires one non-modifier key")
	}
	return strings.Join(mods, "") + main, nil
}

func windowsHotkeyToken(key string) string {
	raw := strings.TrimSpace(key)
	if raw == "" {
		return ""
	}
	lower := strings.ToLower(raw)
	switch lower {
	case "enter", "return":
		return "{ENTER}"
	case "tab":
		return "{TAB}"
	case "esc", "escape":
		return "{ESC}"
	case "delete", "del":
		return "{DEL}"
	case "backspace":
		return "{BACKSPACE}"
	case "space":
		return " "
	case "up", "arrowup":
		return "{UP}"
	case "down", "arrowdown":
		return "{DOWN}"
	case "left", "arrowleft":
		return "{LEFT}"
	case "right", "arrowright":
		return "{RIGHT}"
	case "home":
		return "{HOME}"
	case "end":
		return "{END}"
	case "pageup", "page up", "page_up":
		return "{PGUP}"
	case "pagedown", "page down", "page_down":
		return "{PGDN}"
	}
	if len(raw) == 1 {
		return windowsSendKeysEscapeText(raw)
	}
	if strings.HasPrefix(lower, "f") && len(lower) > 1 {
		isFn := true
		for _, r := range lower[1:] {
			if r < '0' || r > '9' {
				isFn = false
				break
			}
		}
		if isFn {
			return "{" + strings.ToUpper(lower) + "}"
		}
	}
	return "{" + strings.ToUpper(raw) + "}"
}

func windowsSendKeysEscapeText(s string) string {
	if s == "" {
		return ""
	}
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch r {
		case '+':
			b.WriteString("{+}")
		case '^':
			b.WriteString("{^}")
		case '%':
			b.WriteString("{%}")
		case '~':
			b.WriteString("{~}")
		case '(':
			b.WriteString("{(}")
		case ')':
			b.WriteString("{)}")
		case '[':
			b.WriteString("{[}")
		case ']':
			b.WriteString("{]}")
		case '{':
			b.WriteString("{{}")
		case '}':
			b.WriteString("{}}")
		case '\n', '\r':
			b.WriteString("{ENTER}")
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}

func psQuote(v string) string {
	return "'" + strings.ReplaceAll(v, "'", "''") + "'"
}
