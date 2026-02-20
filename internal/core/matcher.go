package core

import (
	"fmt"
	"image"
)

type MatchCandidate struct {
	X     int
	Y     int
	W     int
	H     int
	Score float64
}

type SearchRequest struct {
	Haystack     *image.Gray
	Needle       *image.Gray
	Mask         *image.Gray
	Threshold    float64
	ResizeFactor float64
	MaxResults   int
}

func (r SearchRequest) Validate() error {
	if r.Haystack == nil {
		return fmt.Errorf("haystack cannot be nil")
	}
	if r.Needle == nil {
		return fmt.Errorf("needle cannot be nil")
	}
	if r.ResizeFactor <= 0 {
		return fmt.Errorf("resize factor must be > 0")
	}
	if r.Threshold < 0 || r.Threshold > 1 {
		return fmt.Errorf("threshold must be in [0,1]")
	}
	return nil
}

type Matcher interface {
	Find(req SearchRequest) ([]MatchCandidate, error)
}

