package sikuli

import (
	"fmt"
	"image"
	"image/color"
)

type Pattern struct {
	image        *Image
	similarity   float64
	targetOffset Point
	resizeFactor float64
	mask         *image.Gray
}

// NewPattern creates a match pattern from an image with default similarity settings.
func NewPattern(img *Image) (*Pattern, error) {
	if img == nil || img.Gray() == nil {
		return nil, fmt.Errorf("%w: pattern image is nil", ErrInvalidTarget)
	}
	return &Pattern{
		image:        img,
		similarity:   DefaultSimilarity,
		targetOffset: Point{},
		resizeFactor: 1.0,
	}, nil
}

// Image returns the underlying pattern image.
func (p *Pattern) Image() *Image {
	return p.image
}

// Similar sets the acceptance threshold in the [0,1] range.
// Higher values require a closer match.
func (p *Pattern) Similar(sim float64) *Pattern {
	if sim < 0 {
		sim = 0
	}
	if sim > 1 {
		sim = 1
	}
	p.similarity = sim
	return p
}

// Similarity returns the current acceptance threshold.
func (p *Pattern) Similarity() float64 {
	return p.similarity
}

// Exact is a convenience for Similar(1.0).
func (p *Pattern) Exact() *Pattern {
	p.similarity = ExactSimilarity
	return p
}

// TargetOffset sets the click anchor relative to the matched rectangle.
func (p *Pattern) TargetOffset(dx, dy int) *Pattern {
	p.targetOffset = Point{X: dx, Y: dy}
	return p
}

// Offset returns the configured click anchor offset.
func (p *Pattern) Offset() Point {
	return p.targetOffset
}

// Resize scales the pattern before matching.
func (p *Pattern) Resize(factor float64) *Pattern {
	if factor <= 0 {
		factor = 1.0
	}
	p.resizeFactor = factor
	return p
}

// ResizeFactor returns the currently configured resize factor.
func (p *Pattern) ResizeFactor() float64 {
	return p.resizeFactor
}

// WithMask sets an optional per-pixel mask where 0 excludes and >0 includes pixels.
func (p *Pattern) WithMask(mask *image.Gray) (*Pattern, error) {
	if mask == nil {
		p.mask = nil
		return p, nil
	}
	if p.image == nil || p.image.Gray() == nil {
		return nil, fmt.Errorf("%w: pattern image is nil", ErrInvalidTarget)
	}
	if mask.Bounds().Dx() != p.image.Width() || mask.Bounds().Dy() != p.image.Height() {
		return nil, fmt.Errorf("%w: mask dimensions do not match pattern", ErrInvalidTarget)
	}
	p.mask = mask
	return p, nil
}

// WithMaskMatrix sets an optional binary mask from matrix rows.
func (p *Pattern) WithMaskMatrix(rows [][]uint8) (*Pattern, error) {
	if len(rows) == 0 {
		p.mask = nil
		return p, nil
	}
	if p.image == nil || p.image.Gray() == nil {
		return nil, fmt.Errorf("%w: pattern image is nil", ErrInvalidTarget)
	}
	h := len(rows)
	w := len(rows[0])
	if w != p.image.Width() || h != p.image.Height() {
		return nil, fmt.Errorf("%w: mask dimensions do not match pattern", ErrInvalidTarget)
	}
	m := image.NewGray(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		if len(rows[y]) != w {
			return nil, fmt.Errorf("%w: mask row width mismatch", ErrInvalidTarget)
		}
		for x := 0; x < w; x++ {
			if rows[y][x] == 0 {
				m.SetGray(x, y, color.Gray{Y: 0})
				continue
			}
			m.SetGray(x, y, color.Gray{Y: 255})
		}
	}
	p.mask = m
	return p, nil
}

// Mask returns the currently configured mask.
func (p *Pattern) Mask() *image.Gray {
	return p.mask
}
