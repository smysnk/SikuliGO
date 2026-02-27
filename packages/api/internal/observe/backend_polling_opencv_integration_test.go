//go:build !opencv

package observe

import (
	"errors"
	"image"
	"image/color"
	"testing"

	"github.com/smysnk/sikuligo/internal/core"
	"github.com/smysnk/sikuligo/internal/cv"
)

func TestPollingBackendMapsOpenCVUnsupportedToObserveUnsupported(t *testing.T) {
	backend := newPollingBackend(cv.NewOpenCVMatcher())
	_, err := backend.Observe(core.ObserveRequest{
		Source:   grayObserveMat([][]uint8{{1, 1}, {1, 1}}),
		Region:   image.Rect(0, 0, 2, 2),
		Pattern:  grayObserveMat([][]uint8{{1}}),
		Event:    core.ObserveEventAppear,
		Interval: 0,
		Timeout:  0,
	})
	if err == nil {
		t.Fatalf("expected observe unsupported error")
	}
	if !errors.Is(err, core.ErrObserveUnsupported) {
		t.Fatalf("expected ErrObserveUnsupported, got=%v", err)
	}
}

func grayObserveMat(rows [][]uint8) *image.Gray {
	h := len(rows)
	w := len(rows[0])
	img := image.NewGray(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.SetGray(x, y, color.Gray{Y: rows[y][x]})
		}
	}
	return img
}
