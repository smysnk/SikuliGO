//go:build darwin

package input

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/sikulix/portgo/internal/core"
)

type darwinFakeCall struct {
	name string
	args []string
}

type darwinFakeResp struct {
	out string
	err error
}

type darwinFakeRunner struct {
	calls []darwinFakeCall
	resp  []darwinFakeResp
}

func (f *darwinFakeRunner) Run(_ context.Context, name string, args ...string) (string, error) {
	f.calls = append(f.calls, darwinFakeCall{name: name, args: append([]string(nil), args...)})
	idx := len(f.calls) - 1
	if idx < len(f.resp) {
		return f.resp[idx].out, f.resp[idx].err
	}
	return "", nil
}

func TestDarwinInputBackendDispatch(t *testing.T) {
	runner := &darwinFakeRunner{}
	slept := time.Duration(0)
	backend := &darwinBackend{
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
	if err := backend.Execute(core.InputRequest{Action: core.InputActionHotkey, Keys: []string{"cmd", "shift", "p"}}); err != nil {
		t.Fatalf("hotkey failed: %v", err)
	}

	if len(runner.calls) != 4 {
		t.Fatalf("expected 4 commands, got=%d", len(runner.calls))
	}
	if runner.calls[0].name != "cliclick" || runner.calls[0].args[0] != "m:10,20" {
		t.Fatalf("move command mismatch: %+v", runner.calls[0])
	}
	if runner.calls[1].name != "cliclick" || runner.calls[1].args[0] != "rc:30,40" {
		t.Fatalf("click command mismatch: %+v", runner.calls[1])
	}
	if runner.calls[2].name != "osascript" {
		t.Fatalf("type command mismatch: %+v", runner.calls[2])
	}
	if runner.calls[3].name != "osascript" {
		t.Fatalf("hotkey command mismatch: %+v", runner.calls[3])
	}
	if slept != 5*time.Millisecond {
		t.Fatalf("delay sleep mismatch: got=%v", slept)
	}
}

func TestDarwinInputBackendErrors(t *testing.T) {
	runner := &darwinFakeRunner{
		resp: []darwinFakeResp{
			{out: "missing cliclick", err: errors.New("exit status 127")},
		},
	}
	backend := &darwinBackend{runner: runner}
	err := backend.Execute(core.InputRequest{
		Action: core.InputActionClick,
		X:      1,
		Y:      2,
		Button: "left",
	})
	if err == nil {
		t.Fatalf("expected click error")
	}
}
