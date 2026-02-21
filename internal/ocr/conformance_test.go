//go:build !gosseract

package ocr

import (
	"errors"
	"image"
	"testing"

	"github.com/sikulix/portgo/internal/core"
)

func TestUnsupportedBackendConformance(t *testing.T) {
	backend := New()
	req := core.OCRRequest{
		Image:         image.NewGray(image.Rect(0, 0, 2, 2)),
		Language:      "eng",
		MinConfidence: 0.2,
	}
	_, err := backend.Read(req)
	if !errors.Is(err, core.ErrOCRUnsupported) {
		t.Fatalf("expected ErrOCRUnsupported, got=%v", err)
	}
}

func TestUnsupportedBackendValidation(t *testing.T) {
	backend := New()
	_, err := backend.Read(core.OCRRequest{
		Image:         nil,
		Language:      "eng",
		MinConfidence: 0,
	})
	if err == nil {
		t.Fatalf("expected validation error for nil image")
	}
}
