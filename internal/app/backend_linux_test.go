//go:build linux

package app

import (
	"context"
	"errors"
	"testing"

	"github.com/sikulix/portgo/internal/core"
)

type linuxFakeRunnerCall struct {
	name string
	args []string
}

type linuxFakeRunnerResponse struct {
	out string
	err error
}

type linuxFakeRunner struct {
	calls     []linuxFakeRunnerCall
	responses []linuxFakeRunnerResponse
}

func (f *linuxFakeRunner) Run(_ context.Context, name string, args ...string) (string, error) {
	f.calls = append(f.calls, linuxFakeRunnerCall{name: name, args: append([]string(nil), args...)})
	idx := len(f.calls) - 1
	if idx < len(f.responses) {
		return f.responses[idx].out, f.responses[idx].err
	}
	return "", nil
}

func TestLinuxBackendCommandDispatch(t *testing.T) {
	runner := &linuxFakeRunner{
		responses: []linuxFakeRunnerResponse{
			{},             // open
			{},             // focus
			{},             // close
			{out: "123\n"}, // is-running
			{out: "123\n"}, // list-windows is-running precheck
			{out: "0x001 0 10 20 640 480 host app.Demo Demo Main Window\n"},
		},
	}
	backend := &linuxBackend{runner: runner}

	_, err := backend.Execute(core.AppRequest{Action: core.AppActionOpen, Name: "demo-app", Args: []string{"--flag"}})
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
	running, err := backend.Execute(core.AppRequest{Action: core.AppActionIsRunning, Name: "Demo"})
	if err != nil {
		t.Fatalf("is-running failed: %v", err)
	}
	if !running.Running {
		t.Fatalf("expected running=true")
	}
	windows, err := backend.Execute(core.AppRequest{Action: core.AppActionListWindow, Name: "Demo"})
	if err != nil {
		t.Fatalf("list-windows failed: %v", err)
	}
	if len(windows.Windows) != 1 {
		t.Fatalf("expected one window, got=%d", len(windows.Windows))
	}
	if windows.Windows[0].X != 10 || windows.Windows[0].Y != 20 || windows.Windows[0].W != 640 || windows.Windows[0].H != 480 {
		t.Fatalf("window geometry mismatch: %+v", windows.Windows[0])
	}
	if len(runner.calls) != 6 {
		t.Fatalf("expected 6 command calls, got=%d", len(runner.calls))
	}
	if runner.calls[1].name != "wmctrl" || runner.calls[2].name != "pkill" || runner.calls[3].name != "pgrep" {
		t.Fatalf("command dispatch mismatch: %+v", runner.calls)
	}
}

func TestLinuxBackendParsersAndErrors(t *testing.T) {
	rows := "0x001 0 10 20 640 480 host app.Demo Demo Main Window\n0x002 0 1 2 3 4 host app.Other Other"
	windows, err := parseLinuxWindowRows(rows, "demo")
	if err != nil {
		t.Fatalf("parse rows failed: %v", err)
	}
	if len(windows) != 1 || windows[0].Title != "Demo Main Window" {
		t.Fatalf("row filter mismatch: %+v", windows)
	}

	runner := &linuxFakeRunner{
		responses: []linuxFakeRunnerResponse{
			{out: "", err: errors.New("exit status 1")}, // pgrep no match
		},
	}
	backend := &linuxBackend{runner: runner}
	result, err := backend.Execute(core.AppRequest{Action: core.AppActionIsRunning, Name: "missing"})
	if err != nil {
		t.Fatalf("expected no hard error for missing process, got=%v", err)
	}
	if result.Running {
		t.Fatalf("expected running=false for missing process")
	}
}
