//go:build darwin

package app

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/sikulix/portgo/internal/core"
)

type fakeRunnerCall struct {
	name string
	args []string
}

type fakeRunnerResponse struct {
	out string
	err error
}

type fakeRunner struct {
	calls     []fakeRunnerCall
	responses []fakeRunnerResponse
}

func (f *fakeRunner) Run(_ context.Context, name string, args ...string) (string, error) {
	f.calls = append(f.calls, fakeRunnerCall{name: name, args: append([]string(nil), args...)})
	idx := len(f.calls) - 1
	if idx < len(f.responses) {
		return f.responses[idx].out, f.responses[idx].err
	}
	return "", nil
}

func TestDarwinBackendOpenCommand(t *testing.T) {
	runner := &fakeRunner{}
	backend := &darwinBackend{runner: runner}

	_, err := backend.Execute(core.AppRequest{
		Action: core.AppActionOpen,
		Name:   "TextEdit",
		Args:   []string{"--foo", "bar"},
	})
	if err != nil {
		t.Fatalf("open action failed: %v", err)
	}
	if len(runner.calls) != 1 {
		t.Fatalf("expected 1 command call, got=%d", len(runner.calls))
	}
	call := runner.calls[0]
	if call.name != "open" {
		t.Fatalf("expected open command, got=%q", call.name)
	}
	want := []string{"-a", "TextEdit", "--args", "--foo", "bar"}
	if len(call.args) != len(want) {
		t.Fatalf("open args length mismatch: got=%v want=%v", call.args, want)
	}
	for i := range want {
		if call.args[i] != want[i] {
			t.Fatalf("open arg[%d] mismatch: got=%q want=%q", i, call.args[i], want[i])
		}
	}
}

func TestDarwinBackendFocusAndCloseScripts(t *testing.T) {
	runner := &fakeRunner{}
	backend := &darwinBackend{runner: runner}

	_, err := backend.Execute(core.AppRequest{
		Action: core.AppActionFocus,
		Name:   "Preview",
	})
	if err != nil {
		t.Fatalf("focus action failed: %v", err)
	}
	_, err = backend.Execute(core.AppRequest{
		Action: core.AppActionClose,
		Name:   "Preview",
	})
	if err != nil {
		t.Fatalf("close action failed: %v", err)
	}
	if len(runner.calls) != 2 {
		t.Fatalf("expected 2 command calls, got=%d", len(runner.calls))
	}
	if runner.calls[0].name != "osascript" || !containsArg(runner.calls[0].args, `tell application "Preview" to activate`) {
		t.Fatalf("focus command mismatch: %+v", runner.calls[0])
	}
	if runner.calls[1].name != "osascript" || !containsArg(runner.calls[1].args, `tell application "Preview" to quit`) {
		t.Fatalf("close command mismatch: %+v", runner.calls[1])
	}
}

func TestDarwinBackendIsRunningAndListWindows(t *testing.T) {
	runner := &fakeRunner{
		responses: []fakeRunnerResponse{
			{out: "true\n"},
			{out: "true\n"},
			{out: "Main||10||20||800||600||true\nTools||30||40||300||200||false\n"},
		},
	}
	backend := &darwinBackend{runner: runner}

	isRunning, err := backend.Execute(core.AppRequest{
		Action: core.AppActionIsRunning,
		Name:   "Demo",
	})
	if err != nil {
		t.Fatalf("is-running action failed: %v", err)
	}
	if !isRunning.Running {
		t.Fatalf("expected running=true")
	}

	list, err := backend.Execute(core.AppRequest{
		Action: core.AppActionListWindow,
		Name:   "Demo",
	})
	if err != nil {
		t.Fatalf("list-windows action failed: %v", err)
	}
	if !list.Running {
		t.Fatalf("expected list-windows running=true")
	}
	if len(list.Windows) != 2 {
		t.Fatalf("expected 2 windows, got=%d", len(list.Windows))
	}
	if list.Windows[0].Title != "Main" || list.Windows[0].X != 10 || list.Windows[0].Y != 20 || !list.Windows[0].Focused {
		t.Fatalf("first window mismatch: %+v", list.Windows[0])
	}
	if list.Windows[1].Title != "Tools" || list.Windows[1].W != 300 || list.Windows[1].H != 200 || list.Windows[1].Focused {
		t.Fatalf("second window mismatch: %+v", list.Windows[1])
	}
	if len(runner.calls) != 3 {
		t.Fatalf("expected 3 command calls, got=%d", len(runner.calls))
	}
}

func TestDarwinBackendCommandErrors(t *testing.T) {
	runner := &fakeRunner{
		responses: []fakeRunnerResponse{
			{
				out: "permission denied",
				err: errors.New("exit status 1"),
			},
		},
	}
	backend := &darwinBackend{runner: runner}

	_, err := backend.Execute(core.AppRequest{
		Action: core.AppActionFocus,
		Name:   "Blocked",
	})
	if err == nil {
		t.Fatalf("expected focus error")
	}
	if !strings.Contains(err.Error(), "permission denied") {
		t.Fatalf("expected command output in error, got=%v", err)
	}
}

func containsArg(args []string, expected string) bool {
	for _, arg := range args {
		if strings.Contains(arg, expected) {
			return true
		}
	}
	return false
}
