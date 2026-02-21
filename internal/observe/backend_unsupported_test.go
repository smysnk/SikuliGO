package observe

import (
	"errors"
	"image"
	"testing"

	"github.com/sikulix/portgo/internal/core"
)

func TestUnsupportedObserveBackend(t *testing.T) {
	backend := New()
	_, err := backend.Observe(core.ObserveRequest{
		Source:  image.NewGray(image.Rect(0, 0, 10, 10)),
		Region:  image.Rect(0, 0, 10, 10),
		Pattern: image.NewGray(image.Rect(0, 0, 2, 2)),
		Event:   core.ObserveEventAppear,
	})
	if !errors.Is(err, core.ErrObserveUnsupported) {
		t.Fatalf("expected ErrObserveUnsupported, got=%v", err)
	}
}
