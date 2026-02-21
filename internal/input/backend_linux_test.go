//go:build linux

package input

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/sikulix/portgo/internal/core"
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
	if err := backend.Execute(core.InputRequest{Action: core.InputActionTypeText, Text: "hello", Delay: 5 * time.Millisecond}); err != nil {
		t.Fatalf("type failed: %v", err)
	}
	if err := backend.Execute(core.InputRequest{Action: core.InputActionHotkey, Keys: []string{"ctrl", "shift", "p"}}); err != nil {
		t.Fatalf("hotkey failed: %v", err)
	}

	if len(runner.calls) != 4 {
		t.Fatalf("expected 4 commands, got=%d", len(runner.calls))
	}
	if runner.calls[0].name != "xdotool" || strings.Join(runner.calls[0].args, " ") != "mousemove 10 20" {
		t.Fatalf("move command mismatch: %+v", runner.calls[0])
	}
	if runner.calls[1].name != "xdotool" || strings.Join(runner.calls[1].args, " ") != "mousemove 30 40 click 3" {
		t.Fatalf("click command mismatch: %+v", runner.calls[1])
	}
	if runner.calls[2].name != "xdotool" || strings.Join(runner.calls[2].args, " ") != "type -- hello" {
		t.Fatalf("type command mismatch: %+v", runner.calls[2])
	}
	if runner.calls[3].name != "xdotool" || strings.Join(runner.calls[3].args, " ") != "key ctrl+shift+p" {
		t.Fatalf("hotkey command mismatch: %+v", runner.calls[3])
	}
	if slept != 5*time.Millisecond {
		t.Fatalf("delay sleep mismatch: got=%v", slept)
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
