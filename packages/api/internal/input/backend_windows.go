//go:build windows

package input

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/smysnk/sikuligo/internal/core"
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
	case core.InputActionMouseDown:
		script, err := windowsMouseButtonScript(req.X, req.Y, req.Button, true)
		if err != nil {
			return err
		}
		return b.runPowerShell(ctx, script)
	case core.InputActionMouseUp:
		script, err := windowsMouseButtonScript(req.X, req.Y, req.Button, false)
		if err != nil {
			return err
		}
		return b.runPowerShell(ctx, script)
	case core.InputActionTypeText:
		return b.runPowerShell(ctx, windowsSendKeysScript(windowsSendKeysEscapeText(req.Text)))
	case core.InputActionPasteText:
		return b.runPowerShell(ctx, windowsPasteScript(req.Text))
	case core.InputActionHotkey:
		chord, err := windowsHotkeyChord(req.Keys)
		if err != nil {
			return err
		}
		return b.runPowerShell(ctx, windowsSendKeysScript(chord))
	case core.InputActionKeyDown:
		script, err := windowsKeyStateScript(req.Keys, true)
		if err != nil {
			return err
		}
		return b.runPowerShell(ctx, script)
	case core.InputActionKeyUp:
		script, err := windowsKeyStateScript(req.Keys, false)
		if err != nil {
			return err
		}
		return b.runPowerShell(ctx, script)
	case core.InputActionWheel:
		script, err := windowsWheelScript(req.X, req.Y, req.ScrollDirection, req.ScrollSteps)
		if err != nil {
			return err
		}
		return b.runPowerShell(ctx, script)
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
	down, up, err := windowsMouseFlags(button)
	if err != nil {
		return "", err
	}
	return windowsSetCursorScript(x, y) + fmt.Sprintf(`; [SikuliGoInput]::mouse_event(0x%04x, 0, 0, 0, [UIntPtr]::Zero); [SikuliGoInput]::mouse_event(0x%04x, 0, 0, 0, [UIntPtr]::Zero)`, down, up), nil
}

func windowsMouseButtonScript(x, y int, button string, downEvent bool) (string, error) {
	down, up, err := windowsMouseFlags(button)
	if err != nil {
		return "", err
	}
	flag := up
	if downEvent {
		flag = down
	}
	return windowsSetCursorScript(x, y) + fmt.Sprintf(`; [SikuliGoInput]::mouse_event(0x%04x, 0, 0, 0, [UIntPtr]::Zero)`, flag), nil
}

func windowsMouseFlags(button string) (uint32, uint32, error) {
	switch strings.ToLower(strings.TrimSpace(button)) {
	case "left":
		return 0x0002, 0x0004, nil
	case "right":
		return 0x0008, 0x0010, nil
	case "middle":
		return 0x0020, 0x0040, nil
	default:
		return 0, 0, fmt.Errorf("click requires supported button, got %q", button)
	}
}

func windowsSetCursorScript(x, y int) string {
	return windowsInputSetupScript() + fmt.Sprintf(`; [SikuliGoInput]::SetCursorPos(%d, %d) | Out-Null`, x, y)
}

func windowsInputSetupScript() string {
	return `$sig = @'
using System;
using System.Runtime.InteropServices;
public static class SikuliGoInput {
  [DllImport("user32.dll")]
  public static extern bool SetCursorPos(int X, int Y);
  [DllImport("user32.dll")]
  public static extern void mouse_event(uint dwFlags, uint dx, uint dy, int dwData, UIntPtr dwExtraInfo);
  [DllImport("user32.dll")]
  public static extern void keybd_event(byte bVk, byte bScan, uint dwFlags, UIntPtr dwExtraInfo);
}
'@; if (-not ("SikuliGoInput" -as [type])) { Add-Type -TypeDefinition $sig }`
}

func windowsSendKeysScript(chord string) string {
	return fmt.Sprintf(`Add-Type -AssemblyName System.Windows.Forms; [System.Windows.Forms.SendKeys]::SendWait(%s)`, psQuote(chord))
}

func windowsPasteScript(text string) string {
	return fmt.Sprintf(`Set-Clipboard -Value %s; %s`, psQuote(text), windowsSendKeysScript("^v"))
}

func windowsKeyStateScript(keys []string, down bool) (string, error) {
	script := windowsInputSetupScript()
	flags := uint32(0)
	if !down {
		flags = 0x0002
	}
	for _, key := range keys {
		vk, err := windowsVirtualKeyCode(key)
		if err != nil {
			return "", err
		}
		script += fmt.Sprintf(`; [SikuliGoInput]::keybd_event(0x%02X, 0, 0x%04X, [UIntPtr]::Zero)`, vk, flags)
	}
	return script, nil
}

func windowsWheelScript(x, y int, direction string, steps int) (string, error) {
	flag := uint32(0x0800)
	delta := 120 * steps
	switch strings.ToLower(strings.TrimSpace(direction)) {
	case "up":
	case "down":
		delta = -delta
	case "right":
		flag = 0x01000
	case "left":
		flag = 0x01000
		delta = -delta
	default:
		return "", fmt.Errorf("wheel requires supported direction, got %q", direction)
	}
	return windowsSetCursorScript(x, y) + fmt.Sprintf(`; [SikuliGoInput]::mouse_event(0x%04X, 0, 0, %d, [UIntPtr]::Zero)`, flag, delta), nil
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

func windowsVirtualKeyCode(key string) (byte, error) {
	raw := strings.TrimSpace(key)
	if raw == "" {
		return 0, fmt.Errorf("key cannot be empty")
	}
	lower := strings.ToLower(raw)
	switch lower {
	case "ctrl", "control":
		return 0x11, nil
	case "shift":
		return 0x10, nil
	case "alt", "option":
		return 0x12, nil
	case "cmd", "command", "win", "windows", "super":
		return 0x5B, nil
	case "enter", "return":
		return 0x0D, nil
	case "tab":
		return 0x09, nil
	case "esc", "escape":
		return 0x1B, nil
	case "delete", "del":
		return 0x2E, nil
	case "backspace":
		return 0x08, nil
	case "space":
		return 0x20, nil
	case "up", "arrowup":
		return 0x26, nil
	case "down", "arrowdown":
		return 0x28, nil
	case "left", "arrowleft":
		return 0x25, nil
	case "right", "arrowright":
		return 0x27, nil
	case "home":
		return 0x24, nil
	case "end":
		return 0x23, nil
	case "pageup", "page up", "page_up":
		return 0x21, nil
	case "pagedown", "page down", "page_down":
		return 0x22, nil
	}
	if len(raw) == 1 {
		r := raw[0]
		switch {
		case r >= 'a' && r <= 'z':
			return r - 'a' + 'A', nil
		case r >= 'A' && r <= 'Z':
			return r, nil
		case r >= '0' && r <= '9':
			return r, nil
		}
	}
	if strings.HasPrefix(lower, "f") && len(lower) > 1 {
		n := 0
		for _, r := range lower[1:] {
			if r < '0' || r > '9' {
				n = 0
				break
			}
			n = n*10 + int(r-'0')
		}
		if n >= 1 && n <= 24 {
			return byte(0x70 + n - 1), nil
		}
	}
	return 0, fmt.Errorf("unsupported key %q", key)
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
