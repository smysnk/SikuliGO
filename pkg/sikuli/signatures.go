package sikuli

import "image"

// This file intentionally defines the frozen workstream-1 public signatures.
// If these interfaces are changed, update docs/api-signature-freeze.md.

type ImageAPI interface {
	Name() string
	Width() int
	Height() int
	Gray() *image.Gray
	Clone() *Image
}

type PatternAPI interface {
	Image() *Image
	Similar(sim float64) *Pattern
	Similarity() float64
	Exact() *Pattern
	TargetOffset(dx, dy int) *Pattern
	Offset() Point
	Resize(factor float64) *Pattern
	ResizeFactor() float64
	Mask() *image.Gray
}

type FinderAPI interface {
	Find(pattern *Pattern) (Match, error)
	FindAll(pattern *Pattern) ([]Match, error)
	LastMatches() []Match
}

var (
	_ ImageAPI   = (*Image)(nil)
	_ PatternAPI = (*Pattern)(nil)
	_ FinderAPI  = (*Finder)(nil)
)
