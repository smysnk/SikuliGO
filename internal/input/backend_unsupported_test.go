//go:build !darwin

package input

import (
	"errors"
	"testing"

	"github.com/sikulix/portgo/internal/core"
)

func TestUnsupportedInputBackend(t *testing.T) {
	backend := New()
	err := backend.Execute(core.InputRequest{
		Action: core.InputActionClick,
		X:      10,
		Y:      20,
		Button: "left",
	})
	if !errors.Is(err, core.ErrInputUnsupported) {
		t.Fatalf("expected ErrInputUnsupported, got=%v", err)
	}
}
