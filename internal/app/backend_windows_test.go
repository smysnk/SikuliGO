//go:build windows

package app

import (
	"context"
	"errors"
	"testing"

	"github.com/sikulix/portgo/internal/core"
)

type windowsFakeRunnerCall struct {
	name string
	args []string
}

type windowsFakeRunnerResponse struct {
	out string
	err error
}

type windowsFakeRunner struct {
	calls     []windowsFakeRunnerCall
	responses []windowsFakeRunnerResponse
}

func (f *windowsFakeRunner) Run(_ context.Context, name string, args ...string) (string, error) {
	f.calls = append(f.calls, windowsFakeRunnerCall{name: name, args: append([]string(nil), args...)})
	idx := len(f.calls) - 1
	if idx < len(f.responses) {
		return f.responses[idx].out, f.responses[idx].err
	}
	return "", nil
}

func TestWindowsBackendCommandDispatch(t *testing.T) {
	runner := &windowsFakeRunner{
		responses: []windowsFakeRunnerResponse{
			{},              // open
			{},              // focus
			{},              // close
			{out: "true\n"}, // is-running
			{out: "Demo||0||0||0||0||false\n"},
		},
	}
	backend := &windowsBackend{runner: runner}

	_, err := backend.Execute(core.AppRequest{Action: core.AppActionOpen, Name: "demo.exe", Args: []string{"--debug"}})
	if err != nil {
		t.Fatalf("open failed: %v", err)
	}
	_, err = backend.Execute(core.AppRequest{Action: core.AppActionFocus, Name: "Demo"})
	if err != nil {
		t.Fatalf("focus failed: %v", err)
	}
	_, err = backend.Execute(core.AppRequest{Action: core.AppActionClose, Name: "Demo"})
	if err != nil {
		t.Fatalf("close failed: %v", err)
	}
	isRunning, err := backend.Execute(core.AppRequest{Action: core.AppActionIsRunning, Name: "Demo"})
	if err != nil {
		t.Fatalf("is-running failed: %v", err)
	}
	if !isRunning.Running {
		t.Fatalf("expected running=true")
	}
	list, err := backend.Execute(core.AppRequest{Action: core.AppActionListWindow, Name: "Demo"})
	if err != nil {
		t.Fatalf("list-windows failed: %v", err)
	}
	if len(list.Windows) != 1 || list.Windows[0].Title != "Demo" {
		t.Fatalf("windows mismatch: %+v", list.Windows)
	}
	if len(runner.calls) != 5 {
		t.Fatalf("expected 5 command calls, got=%d", len(runner.calls))
	}
	for _, c := range runner.calls {
		if c.name != "powershell" {
			t.Fatalf("expected powershell command, got=%q", c.name)
		}
	}
}

func TestWindowsBackendCommandErrors(t *testing.T) {
	runner := &windowsFakeRunner{
		responses: []windowsFakeRunnerResponse{
			{out: "access denied", err: errors.New("exit status 1")},
		},
	}
	backend := &windowsBackend{runner: runner}
	_, err := backend.Execute(core.AppRequest{Action: core.AppActionClose, Name: "Denied"})
	if err == nil {
		t.Fatalf("expected close error")
	}
}
