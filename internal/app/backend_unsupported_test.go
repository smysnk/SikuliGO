package app

import (
	"errors"
	"testing"

	"github.com/sikulix/portgo/internal/core"
)

func TestUnsupportedAppBackend(t *testing.T) {
	backend := New()
	_, err := backend.Execute(core.AppRequest{
		Action: core.AppActionOpen,
		Name:   "demo",
	})
	if !errors.Is(err, core.ErrAppUnsupported) {
		t.Fatalf("expected ErrAppUnsupported, got=%v", err)
	}
}
