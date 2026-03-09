//go:build windows

package input

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/smysnk/sikuligo/internal/core"
)

type windowsFakeCall struct {
	name string
	args []string
}

type windowsFakeResp struct {
	out string
	err error
}

type windowsFakeRunner struct {
	calls []windowsFakeCall
	resp  []windowsFakeResp
}

func (f *windowsFakeRunner) Run(_ context.Context, name string, args ...string) (string, error) {
	f.calls = append(f.calls, windowsFakeCall{name: name, args: append([]string(nil), args...)})
	idx := len(f.calls) - 1
	if idx < len(f.resp) {
		return f.resp[idx].out, f.resp[idx].err
	}
	return "", nil
}

func TestWindowsInputBackendDispatch(t *testing.T) {
	runner := &windowsFakeRunner{}
	slept := time.Duration(0)
	backend := &windowsBackend{
		runner: runner,
		sleep: func(d time.Duration) {
			slept += d
		},
	}

	if err := backend.Execute(core.InputRequest{Action: core.InputActionMouseMove, X: 10, Y: 20}); err != nil {
		t.Fatalf("mouse move failed: %v", err)
	}
	if err := backend.Execute(core.InputRequest{Action: core.InputActionClick, X: 30, Y: 40, Button: "left"}); err != nil {
		t.Fatalf("click failed: %v", err)
	}
	if err := backend.Execute(core.InputRequest{Action: core.InputActionMouseDown, X: 31, Y: 41, Button: "left"}); err != nil {
		t.Fatalf("mouse down failed: %v", err)
	}
	if err := backend.Execute(core.InputRequest{Action: core.InputActionMouseUp, X: 31, Y: 41, Button: "left"}); err != nil {
		t.Fatalf("mouse up failed: %v", err)
	}
	if err := backend.Execute(core.InputRequest{Action: core.InputActionTypeText, Text: "hello", Delay: 5 * time.Millisecond}); err != nil {
		t.Fatalf("type failed: %v", err)
	}
	if err := backend.Execute(core.InputRequest{Action: core.InputActionPasteText, Text: "world"}); err != nil {
		t.Fatalf("paste failed: %v", err)
	}
	if err := backend.Execute(core.InputRequest{Action: core.InputActionHotkey, Keys: []string{"ctrl", "shift", "p"}}); err != nil {
		t.Fatalf("hotkey failed: %v", err)
	}
	if err := backend.Execute(core.InputRequest{Action: core.InputActionKeyDown, Keys: []string{"ctrl", "shift"}}); err != nil {
		t.Fatalf("key down failed: %v", err)
	}
	if err := backend.Execute(core.InputRequest{Action: core.InputActionKeyUp, Keys: []string{"ctrl", "shift"}}); err != nil {
		t.Fatalf("key up failed: %v", err)
	}
	if err := backend.Execute(core.InputRequest{Action: core.InputActionWheel, X: 32, Y: 42, ScrollDirection: "down", ScrollSteps: 2}); err != nil {
		t.Fatalf("wheel failed: %v", err)
	}

	if len(runner.calls) != 10 {
		t.Fatalf("expected 10 commands, got=%d", len(runner.calls))
	}
	for i, call := range runner.calls {
		if call.name != "powershell" {
			t.Fatalf("call[%d] name mismatch: %+v", i, call)
		}
		if len(call.args) < 3 || call.args[0] != "-NoProfile" || call.args[1] != "-Command" {
			t.Fatalf("call[%d] args mismatch: %+v", i, call)
		}
	}
	if !strings.Contains(runner.calls[0].args[2], "SetCursorPos(10, 20)") {
		t.Fatalf("move script mismatch: %s", runner.calls[0].args[2])
	}
	if !strings.Contains(runner.calls[1].args[2], "SetCursorPos(30, 40)") || !strings.Contains(runner.calls[1].args[2], "mouse_event") {
		t.Fatalf("click script mismatch: %s", runner.calls[1].args[2])
	}
	if !strings.Contains(runner.calls[2].args[2], "SetCursorPos(31, 41)") || !strings.Contains(runner.calls[2].args[2], "0x0002") {
		t.Fatalf("mouse down script mismatch: %s", runner.calls[2].args[2])
	}
	if !strings.Contains(runner.calls[3].args[2], "SetCursorPos(31, 41)") || !strings.Contains(runner.calls[3].args[2], "0x0004") {
		t.Fatalf("mouse up script mismatch: %s", runner.calls[3].args[2])
	}
	if !strings.Contains(runner.calls[4].args[2], "SendWait") || !strings.Contains(runner.calls[4].args[2], "'hello'") {
		t.Fatalf("type script mismatch: %s", runner.calls[4].args[2])
	}
	if !strings.Contains(runner.calls[5].args[2], "Set-Clipboard -Value 'world'") || !strings.Contains(runner.calls[5].args[2], "SendWait('^v')") {
		t.Fatalf("paste script mismatch: %s", runner.calls[5].args[2])
	}
	if !strings.Contains(runner.calls[6].args[2], "SendWait") || !strings.Contains(runner.calls[6].args[2], "'^+p'") {
		t.Fatalf("hotkey script mismatch: %s", runner.calls[6].args[2])
	}
	if !strings.Contains(runner.calls[7].args[2], "keybd_event") || !strings.Contains(runner.calls[7].args[2], "0x11") || !strings.Contains(runner.calls[7].args[2], "0x10") {
		t.Fatalf("key down script mismatch: %s", runner.calls[7].args[2])
	}
	if !strings.Contains(runner.calls[8].args[2], "keybd_event") || !strings.Contains(runner.calls[8].args[2], "0x0002") {
		t.Fatalf("key up script mismatch: %s", runner.calls[8].args[2])
	}
	if !strings.Contains(runner.calls[9].args[2], "mouse_event") || !strings.Contains(runner.calls[9].args[2], "0x0800") || !strings.Contains(runner.calls[9].args[2], "-240") {
		t.Fatalf("wheel script mismatch: %s", runner.calls[9].args[2])
	}
	if slept != 5*time.Millisecond {
		t.Fatalf("delay sleep mismatch: got=%v", slept)
	}
}

func TestWindowsInputBackendKeyUp(t *testing.T) {
	runner := &windowsFakeRunner{}
	backend := &windowsBackend{runner: runner}

	if err := backend.Execute(core.InputRequest{Action: core.InputActionKeyUp, Keys: []string{"ctrl", "shift"}}); err != nil {
		t.Fatalf("key up failed: %v", err)
	}
	if !strings.Contains(runner.calls[0].args[2], "keybd_event") || !strings.Contains(runner.calls[0].args[2], "0x0002") {
		t.Fatalf("key up script mismatch: %s", runner.calls[0].args[2])
	}
}

func TestWindowsInputBackendErrors(t *testing.T) {
	backend := &windowsBackend{runner: &windowsFakeRunner{}}
	err := backend.Execute(core.InputRequest{
		Action: core.InputActionClick,
		X:      1,
		Y:      2,
		Button: "invalid",
	})
	if err == nil {
		t.Fatalf("expected invalid button error")
	}

	runner := &windowsFakeRunner{
		resp: []windowsFakeResp{
			{out: "Access denied", err: errors.New("exit status 1")},
		},
	}
	backend = &windowsBackend{runner: runner}
	err = backend.Execute(core.InputRequest{
		Action: core.InputActionTypeText,
		Text:   "hello",
	})
	if err == nil {
		t.Fatalf("expected command error")
	}
	if !strings.Contains(err.Error(), "powershell input action failed") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestWindowsSendKeysEscaping(t *testing.T) {
	got := windowsSendKeysEscapeText("a+b^{x}\n")
	want := "a{+}b{^}{{}x{}}{ENTER}"
	if got != want {
		t.Fatalf("escape mismatch: got=%q want=%q", got, want)
	}

	chord, err := windowsHotkeyChord([]string{"ctrl", "shift", "enter"})
	if err != nil {
		t.Fatalf("hotkey chord failed: %v", err)
	}
	if chord != "^+{ENTER}" {
		t.Fatalf("hotkey chord mismatch: got=%q", chord)
	}
}
