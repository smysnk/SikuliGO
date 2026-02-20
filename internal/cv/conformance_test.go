package cv

import (
	"image"
	"image/color"
	"testing"

	"github.com/sikulix/portgo/internal/core"
)

func TestMatcherConformance(t *testing.T) {
	matchers := []struct {
		name string
		m    core.Matcher
	}{
		{name: "ncc", m: NewNCCMatcher()},
		{name: "sad", m: NewSADMatcher()},
	}

	for _, tc := range matchers {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Run("ordering_and_max_results", func(t *testing.T) {
				hay := mat([][]uint8{
					{10, 10, 10, 10, 10, 10, 10, 10},
					{10, 0, 255, 10, 10, 10, 10, 10},
					{10, 255, 0, 10, 0, 255, 10, 10},
					{10, 10, 10, 10, 255, 0, 10, 10},
					{10, 10, 10, 10, 10, 10, 10, 10},
				})
				needle := mat([][]uint8{
					{0, 255},
					{255, 0},
				})
				got, err := tc.m.Find(core.SearchRequest{
					Haystack:     hay,
					Needle:       needle,
					Threshold:    0.999,
					ResizeFactor: 1.0,
					MaxResults:   0,
				})
				if err != nil {
					t.Fatalf("find failed: %v", err)
				}
				if len(got) != 2 {
					t.Fatalf("expected 2 matches, got %d", len(got))
				}
				if got[0].X != 1 || got[0].Y != 1 {
					t.Fatalf("first match order mismatch: %+v", got[0])
				}
				if got[1].X != 4 || got[1].Y != 2 {
					t.Fatalf("second match order mismatch: %+v", got[1])
				}
				got, err = tc.m.Find(core.SearchRequest{
					Haystack:     hay,
					Needle:       needle,
					Threshold:    0.999,
					ResizeFactor: 1.0,
					MaxResults:   1,
				})
				if err != nil {
					t.Fatalf("find(max=1) failed: %v", err)
				}
				if len(got) != 1 {
					t.Fatalf("expected max result truncation to 1, got %d", len(got))
				}
			})

			t.Run("mask_behavior", func(t *testing.T) {
				hay := mat([][]uint8{
					{0, 0, 0, 0, 0},
					{0, 10, 20, 0, 0},
					{0, 30, 40, 0, 0},
					{0, 0, 0, 0, 0},
				})
				needle := mat([][]uint8{
					{10, 200},
					{30, 40},
				})
				mask := mat([][]uint8{
					{255, 0},
					{255, 255},
				})
				got, err := tc.m.Find(core.SearchRequest{
					Haystack:     hay,
					Needle:       needle,
					Mask:         mask,
					Threshold:    0.999,
					ResizeFactor: 1.0,
				})
				if err != nil {
					t.Fatalf("masked find failed: %v", err)
				}
				if len(got) != 1 || got[0].X != 1 || got[0].Y != 1 {
					t.Fatalf("masked match mismatch: %+v", got)
				}
			})

			t.Run("resize_behavior", func(t *testing.T) {
				hay := mat([][]uint8{
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 100, 100, 0, 0},
					{0, 0, 0, 100, 100, 0, 0},
					{0, 150, 150, 255, 255, 0, 0},
					{0, 150, 150, 255, 255, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0},
				})
				needle := mat([][]uint8{
					{0, 100},
					{150, 255},
				})
				got, err := tc.m.Find(core.SearchRequest{
					Haystack:     hay,
					Needle:       needle,
					Threshold:    0.999,
					ResizeFactor: 2.0,
					MaxResults:   1,
				})
				if err != nil {
					t.Fatalf("resized find failed: %v", err)
				}
				if len(got) != 1 {
					t.Fatalf("expected 1 match, got %d", len(got))
				}
				if got[0].X != 1 || got[0].Y != 1 || got[0].W != 4 || got[0].H != 4 {
					t.Fatalf("resized match geometry mismatch: %+v", got[0])
				}
			})
		})
	}
}

func mat(rows [][]uint8) *image.Gray {
	h := len(rows)
	w := len(rows[0])
	g := image.NewGray(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			g.SetGray(x, y, color.Gray{Y: rows[y][x]})
		}
	}
	return g
}
