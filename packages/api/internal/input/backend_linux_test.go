//go:build linux

package input

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/smysnk/sikuligo/internal/core"
)

type linuxFakeCall struct {
	name string
	args []string
}

type linuxFakeResp struct {
	out string
	err error
}

type linuxFakeRunner struct {
	calls []linuxFakeCall
	resp  []linuxFakeResp
}

func (f *linuxFakeRunner) Run(_ context.Context, name string, args ...string) (string, error) {
	f.calls = append(f.calls, linuxFakeCall{name: name, args: append([]string(nil), args...)})
	idx := len(f.calls) - 1
	if idx < len(f.resp) {
		return f.resp[idx].out, f.resp[idx].err
	}
	return "", nil
}

func TestLinuxInputBackendDispatch(t *testing.T) {
	runner := &linuxFakeRunner{}
	slept := time.Duration(0)
	backend := &linuxBackend{
		runner: runner,
		sleep: func(d time.Duration) {
			slept += d
		},
	}

	if err := backend.Execute(core.InputRequest{Action: core.InputActionMouseMove, X: 10, Y: 20}); err != nil {
		t.Fatalf("mouse move failed: %v", err)
	}
	if err := backend.Execute(core.InputRequest{Action: core.InputActionClick, X: 30, Y: 40, Button: "right"}); err != nil {
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

	if len(runner.calls) != 12 {
		t.Fatalf("expected 12 commands, got=%d", len(runner.calls))
	}
	if runner.calls[0].name != "xdotool" || strings.Join(runner.calls[0].args, " ") != "mousemove 10 20" {
		t.Fatalf("move command mismatch: %+v", runner.calls[0])
	}
	if runner.calls[1].name != "xdotool" || strings.Join(runner.calls[1].args, " ") != "mousemove 30 40 click 3" {
		t.Fatalf("click command mismatch: %+v", runner.calls[1])
	}
	if runner.calls[2].name != "xdotool" || strings.Join(runner.calls[2].args, " ") != "mousemove 31 41 mousedown 1" {
		t.Fatalf("mouse down command mismatch: %+v", runner.calls[2])
	}
	if runner.calls[3].name != "xdotool" || strings.Join(runner.calls[3].args, " ") != "mousemove 31 41 mouseup 1" {
		t.Fatalf("mouse up command mismatch: %+v", runner.calls[3])
	}
	if runner.calls[4].name != "xdotool" || strings.Join(runner.calls[4].args, " ") != "type -- hello" {
		t.Fatalf("type command mismatch: %+v", runner.calls[4])
	}
	if runner.calls[5].name != "xdotool" || strings.Join(runner.calls[5].args, " ") != "type -- world" {
		t.Fatalf("paste command mismatch: %+v", runner.calls[5])
	}
	if runner.calls[6].name != "xdotool" || strings.Join(runner.calls[6].args, " ") != "key ctrl+shift+p" {
		t.Fatalf("hotkey command mismatch: %+v", runner.calls[6])
	}
	if runner.calls[7].name != "xdotool" || strings.Join(runner.calls[7].args, " ") != "keydown ctrl" {
		t.Fatalf("key down ctrl mismatch: %+v", runner.calls[7])
	}
	if runner.calls[8].name != "xdotool" || strings.Join(runner.calls[8].args, " ") != "keydown shift" {
		t.Fatalf("key down shift mismatch: %+v", runner.calls[8])
	}
	if runner.calls[9].name != "xdotool" || strings.Join(runner.calls[9].args, " ") != "keyup ctrl" {
		t.Fatalf("key up ctrl mismatch: %+v", runner.calls[9])
	}
	if runner.calls[10].name != "xdotool" || strings.Join(runner.calls[10].args, " ") != "keyup shift" {
		t.Fatalf("key up shift mismatch: %+v", runner.calls[10])
	}
	if runner.calls[11].name != "xdotool" || strings.Join(runner.calls[11].args, " ") != "mousemove 32 42 click --repeat 2 5" {
		t.Fatalf("wheel command mismatch: %+v", runner.calls[11])
	}
	if slept != 5*time.Millisecond {
		t.Fatalf("delay sleep mismatch: got=%v", slept)
	}
}

func TestLinuxInputBackendKeyUpAndWheel(t *testing.T) {
	runner := &linuxFakeRunner{}
	backend := &linuxBackend{runner: runner}

	if err := backend.Execute(core.InputRequest{Action: core.InputActionKeyUp, Keys: []string{"shift"}}); err != nil {
		t.Fatalf("key up failed: %v", err)
	}
	if got := strings.Join(runner.calls[0].args, " "); got != "keyup shift" {
		t.Fatalf("key up command mismatch: %q", got)
	}
	if err := backend.Execute(core.InputRequest{Action: core.InputActionWheel, X: 1, Y: 2, ScrollDirection: "right", ScrollSteps: 2}); err != nil {
		t.Fatalf("wheel failed: %v", err)
	}
	if got := strings.Join(runner.calls[1].args, " "); got != "mousemove 1 2 click --repeat 2 7" {
		t.Fatalf("wheel command mismatch: %q", got)
	}
}

func TestLinuxInputBackendErrors(t *testing.T) {
	backend := &linuxBackend{runner: &linuxFakeRunner{}}
	err := backend.Execute(core.InputRequest{
		Action: core.InputActionClick,
		X:      1,
		Y:      2,
		Button: "invalid",
	})
	if err == nil {
		t.Fatalf("expected invalid button error")
	}

	runner := &linuxFakeRunner{
		resp: []linuxFakeResp{
			{out: "xdotool: not found", err: errors.New("exit status 127")},
		},
	}
	backend = &linuxBackend{runner: runner}
	err = backend.Execute(core.InputRequest{
		Action: core.InputActionTypeText,
		Text:   "hello",
	})
	if err == nil {
		t.Fatalf("expected command error")
	}
	if !strings.Contains(err.Error(), "xdotool input action failed") {
		t.Fatalf("unexpected command error: %v", err)
	}
}
