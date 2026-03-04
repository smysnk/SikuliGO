package main

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestPromptInstallCliclick(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  string
		expect bool
	}{
		{name: "default yes", input: "\n", expect: true},
		{name: "explicit yes", input: "yes\n", expect: true},
		{name: "explicit no", input: "n\n", expect: false},
		{name: "unknown", input: "maybe\n", expect: false},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var out bytes.Buffer
			got, err := promptInstallCliclick(strings.NewReader(tc.input), &out)
			if err != nil {
				t.Fatalf("promptInstallCliclick error: %v", err)
			}
			if got != tc.expect {
				t.Fatalf("promptInstallCliclick=%v want=%v", got, tc.expect)
			}
			if !strings.Contains(out.String(), "Install \"cliclick\" now with Homebrew?") {
				t.Fatalf("expected prompt output, got=%q", out.String())
			}
		})
	}
}

func TestStartupStateRoundTrip(t *testing.T) {
	t.Parallel()
	tmp := t.TempDir()
	statePath := filepath.Join(tmp, "state.json")
	want := startupState{SuppressCliclickPrompt: true}
	if err := saveStartupState(statePath, want); err != nil {
		t.Fatalf("saveStartupState error: %v", err)
	}
	got, err := loadStartupState(statePath)
	if err != nil {
		t.Fatalf("loadStartupState error: %v", err)
	}
	if got.SuppressCliclickPrompt != want.SuppressCliclickPrompt {
		t.Fatalf("state mismatch got=%+v want=%+v", got, want)
	}
}

func TestStartupStatePathUsesXDGConfigHome(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", tmp)
	got := startupStatePath()
	want := filepath.Join(tmp, "sikuligo", "startup-state.json")
	if got != want {
		t.Fatalf("startupStatePath=%q want=%q", got, want)
	}
}

func TestLoadStartupStateMissingFile(t *testing.T) {
	t.Parallel()
	got, err := loadStartupState(filepath.Join(t.TempDir(), "missing.json"))
	if err != nil {
		t.Fatalf("loadStartupState missing file error: %v", err)
	}
	if got.SuppressCliclickPrompt {
		t.Fatalf("expected default state, got=%+v", got)
	}
}

func TestStartupStatePathFallbackWithoutHome(t *testing.T) {
	t.Setenv("HOME", "")
	t.Setenv("XDG_CONFIG_HOME", "")
	got := startupStatePath()
	if got == "" {
		t.Fatalf("startupStatePath should not be empty")
	}
}
