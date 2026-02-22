//go:build opencv

package cv

import (
	"fmt"
	"image"
	"sort"

	"gocv.io/x/gocv"

	"github.com/sikulix/portgo/internal/core"
)

type OpenCVMatcher struct{}

func NewOpenCVMatcher() *OpenCVMatcher {
	return &OpenCVMatcher{}
}

func (m *OpenCVMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error) {
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

	hayMat, err := grayToMat(req.Haystack)
	if err != nil {
		return nil, err
	}
	defer hayMat.Close()

	needleMat, err := grayToMat(needle)
	if err != nil {
		return nil, err
	}
	defer needleMat.Close()

	method := gocv.TmCcoeffNormed
	maskMat := gocv.NewMat()
	useMask := false
	if mask != nil {
		maskMat, err = grayToMat(mask)
		if err != nil {
			return nil, err
		}
		useMask = true
		method = gocv.TmCcorrNormed
	}
	if useMask {
		defer maskMat.Close()
	}

	rows := hh - nh + 1
	cols := hw - nw + 1
	result := gocv.NewMatWithSize(rows, cols, gocv.MatTypeCV32F)
	defer result.Close()

	if useMask {
		if err := gocv.MatchTemplate(hayMat, needleMat, &result, method, maskMat); err != nil {
			return nil, err
		}
	} else {
		emptyMask := gocv.NewMat()
		defer emptyMask.Close()
		if err := gocv.MatchTemplate(hayMat, needleMat, &result, method, emptyMask); err != nil {
			return nil, err
		}
	}

	results := make([]core.MatchCandidate, 0)
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			raw := float64(result.GetFloatAt(y, x))
			score := normalizeTemplateScore(raw, method)
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

func normalizeTemplateScore(raw float64, method gocv.TemplateMatchMode) float64 {
	score := raw
	if method == gocv.TmCcoeffNormed {
		score = (raw + 1.0) / 2.0
	}
	if score < 0 {
		return 0
	}
	if score > 1 {
		return 1
	}
	return score
}

func grayToMat(img *image.Gray) (gocv.Mat, error) {
	if img == nil {
		return gocv.NewMat(), fmt.Errorf("image is nil")
	}
	b := img.Bounds()
	w := b.Dx()
	h := b.Dy()
	if w <= 0 || h <= 0 {
		return gocv.NewMat(), fmt.Errorf("image dimensions must be > 0")
	}

	data := make([]byte, w*h)
	for y := 0; y < h; y++ {
		srcStart := img.PixOffset(b.Min.X, b.Min.Y+y)
		copy(data[y*w:(y+1)*w], img.Pix[srcStart:srcStart+w])
	}
	mat, err := gocv.NewMatFromBytes(h, w, gocv.MatTypeCV8UC1, data)
	if err != nil {
		return gocv.NewMat(), err
	}
	return mat, nil
}
