package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestDiscoverRuntimeSourcesPrefersCanonicalNames(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	files := []string{
		"sikuli-go",
		"sikuli-go-1234abcd1234abcd",
		"sikuligo",
		"sikuligrpc",
		"sikuligo-abcdef0123456789",
		"sikuli-go-monitor",
		"sikuligo-monitor-deadbeefcafebabe",
	}
	for _, name := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(name), 0o755); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}

	got := discoverRuntimeSources(filepath.Join(dir, "sikuligo-abcdef0123456789"))
	want := map[string]string{
		"sikuli-go":         filepath.Join(dir, "sikuli-go"),
		"sikuli-go-monitor": filepath.Join(dir, "sikuli-go-monitor"),
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("discoverRuntimeSources() = %#v, want %#v", got, want)
	}
}

func TestCleanupInstalledRuntimeAliasesRemovesNonCanonicalNames(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	keep := []string{"sikuli-go", "sikuli-go-monitor", "notes.txt"}
	remove := []string{"sikuli-go-deadbeefcafebabe", "sikuligo-abcdef0123456789", "sikuligrpc", "sikuligo-monitor-deadbeefcafebabe"}
	for _, name := range append(append([]string{}, keep...), remove...) {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(name), 0o644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}

	if err := cleanupInstalledRuntimeAliases(dir, nil); err != nil {
		t.Fatalf("cleanupInstalledRuntimeAliases() error = %v", err)
	}

	for _, name := range keep {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Fatalf("expected %s to remain: %v", name, err)
		}
	}
	for _, name := range remove {
		if _, err := os.Stat(filepath.Join(dir, name)); !os.IsNotExist(err) {
			t.Fatalf("expected %s to be removed, stat err=%v", name, err)
		}
	}
}
