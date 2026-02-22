//go:build !opencv

package cv

import (
	"errors"
	"image"
	"image/color"
	"testing"

	"github.com/sikulix/portgo/internal/core"
)

func TestOpenCVMatcherReturnsUnsupportedWithoutTag(t *testing.T) {
	m := NewOpenCVMatcher()
	_, err := m.Find(core.SearchRequest{
		Haystack:     grayMat([][]uint8{{1, 2}, {3, 4}}),
		Needle:       grayMat([][]uint8{{1}}),
		Threshold:    0.5,
		ResizeFactor: 1.0,
	})
	if err == nil {
		t.Fatalf("expected unsupported error")
	}
	if !errors.Is(err, core.ErrMatcherUnsupported) {
		t.Fatalf("expected ErrMatcherUnsupported, got=%v", err)
	}
}

func grayMat(rows [][]uint8) *image.Gray {
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
