package cv

import (
	"fmt"
	"image"
	"math"
	"sort"

	"github.com/sikulix/portgo/internal/core"
)

type SADMatcher struct{}

func NewSADMatcher() *SADMatcher {
	return &SADMatcher{}
}

func (m *SADMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	needle := req.Needle
	mask := req.Mask
	if req.ResizeFactor != 1.0 {
		needle = core.ResizeGrayNearest(needle, req.ResizeFactor)
		if mask != nil {
			mask = core.ResizeGrayNearest(mask, req.ResizeFactor)
		}
	}
	if mask != nil {
		if mask.Bounds().Dx() != needle.Bounds().Dx() || mask.Bounds().Dy() != needle.Bounds().Dy() {
			return nil, fmt.Errorf("mask dimensions must match needle dimensions")
		}
	}

	hb := req.Haystack.Bounds()
	nb := needle.Bounds()
	hw := hb.Dx()
	hh := hb.Dy()
	nw := nb.Dx()
	nh := nb.Dy()
	if nw <= 0 || nh <= 0 || hw <= 0 || hh <= 0 {
		return nil, nil
	}
	if nw > hw || nh > hh {
		return nil, nil
	}

	results := make([]core.MatchCandidate, 0)
	for y := 0; y <= hh-nh; y++ {
		for x := 0; x <= hw-nw; x++ {
			score := sadScoreAt(req.Haystack, needle, mask, x, y)
			if score < req.Threshold {
				continue
			}
			results = append(results, core.MatchCandidate{
				X:     hb.Min.X + x,
				Y:     hb.Min.Y + y,
				W:     nw,
				H:     nh,
				Score: score,
			})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Score == results[j].Score {
			if results[i].Y == results[j].Y {
				return results[i].X < results[j].X
			}
			return results[i].Y < results[j].Y
		}
		return results[i].Score > results[j].Score
	})
	if req.MaxResults > 0 && len(results) > req.MaxResults {
		results = results[:req.MaxResults]
	}
	return results, nil
}

func sadScoreAt(haystack, needle, mask *image.Gray, x0, y0 int) float64 {
	hb := haystack.Bounds()
	nb := needle.Bounds()
	nw := nb.Dx()
	nh := nb.Dy()

	var count float64
	var totalDiff float64
	for y := 0; y < nh; y++ {
		for x := 0; x < nw; x++ {
			if !maskInclude(mask, x, y) {
				continue
			}
			nv := float64(grayAt(needle, nb.Min.X+x, nb.Min.Y+y))
			hv := float64(grayAt(haystack, hb.Min.X+x0+x, hb.Min.Y+y0+y))
			totalDiff += math.Abs(nv - hv)
			count++
		}
	}
	if count == 0 {
		return 0
	}
	score := 1.0 - (totalDiff / (count * 255.0))
	if score < 0 {
		return 0
	}
	if score > 1 {
		return 1
	}
	return score
}
