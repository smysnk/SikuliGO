//go:build opencv

package cv

import (
	"fmt"
	"math"

	"gocv.io/x/gocv"

	"github.com/smysnk/sikuligo/internal/core"
)

const (
	orbRatioTestThreshold   = 0.75
	orbMinGoodMatches       = 8
	orbHomographyThreshold  = 3.0
	orbHomographyMaxIters   = 2000
	orbHomographyConfidence = 0.995
)

type ORBMatcher struct{}

func NewORBMatcher() *ORBMatcher {
	return &ORBMatcher{}
}

func (m *ORBMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error) {
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

	emptyMask := gocv.NewMat()
	defer emptyMask.Close()

	needleMask := emptyMask
	hasNeedleMask := false
	if mask != nil {
		needleMask, err = grayToMat(mask)
		if err != nil {
			return nil, err
		}
		hasNeedleMask = true
	}
	if hasNeedleMask {
		defer needleMask.Close()
	}

	orb := gocv.NewORB()
	defer orb.Close()

	hayKP, hayDesc := orb.DetectAndCompute(hayMat, emptyMask)
	defer hayDesc.Close()

	needleKP, needleDesc := orb.DetectAndCompute(needleMat, needleMask)
	defer needleDesc.Close()

	if len(hayKP) == 0 || len(needleKP) == 0 || hayDesc.Empty() || needleDesc.Empty() {
		return nil, nil
	}

	bf := gocv.NewBFMatcherWithParams(gocv.NormHamming, false)
	defer bf.Close()

	knn := bf.KnnMatch(needleDesc, hayDesc, 2)
	good := make([]gocv.DMatch, 0, len(knn))
	for _, pair := range knn {
		if len(pair) < 2 {
			continue
		}
		if pair[0].Distance < orbRatioTestThreshold*pair[1].Distance {
			good = append(good, pair[0])
		}
	}
	if len(good) < orbMinGoodMatches {
		return nil, nil
	}

	srcPoints := make([]gocv.Point2f, 0, len(good))
	dstPoints := make([]gocv.Point2f, 0, len(good))
	var distanceSum float64
	for _, match := range good {
		if match.QueryIdx < 0 || match.QueryIdx >= len(needleKP) || match.TrainIdx < 0 || match.TrainIdx >= len(hayKP) {
			continue
		}
		n := needleKP[match.QueryIdx]
		h := hayKP[match.TrainIdx]
		srcPoints = append(srcPoints, gocv.NewPoint2f(float32(n.X), float32(n.Y)))
		dstPoints = append(dstPoints, gocv.NewPoint2f(float32(h.X), float32(h.Y)))
		distanceSum += match.Distance
	}
	if len(srcPoints) < 4 || len(dstPoints) < 4 {
		return nil, nil
	}

	srcVec := gocv.NewPoint2fVectorFromPoints(srcPoints)
	defer srcVec.Close()
	dstVec := gocv.NewPoint2fVectorFromPoints(dstPoints)
	defer dstVec.Close()

	srcMat := gocv.NewMatFromPoint2fVector(srcVec, true)
	defer srcMat.Close()
	dstMat := gocv.NewMatFromPoint2fVector(dstVec, true)
	defer dstMat.Close()

	inlierMask := gocv.NewMat()
	defer inlierMask.Close()

	homography := gocv.FindHomography(
		srcMat,
		dstMat,
		gocv.HomographyMethodRANSAC,
		orbHomographyThreshold,
		&inlierMask,
		orbHomographyMaxIters,
		orbHomographyConfidence,
	)
	defer homography.Close()
	if homography.Empty() {
		return nil, nil
	}

	corners := []gocv.Point2f{
		gocv.NewPoint2f(0, 0),
		gocv.NewPoint2f(float32(nw), 0),
		gocv.NewPoint2f(float32(nw), float32(nh)),
		gocv.NewPoint2f(0, float32(nh)),
	}
	cornerVec := gocv.NewPoint2fVectorFromPoints(corners)
	defer cornerVec.Close()
	cornerMat := gocv.NewMatFromPoint2fVector(cornerVec, true)
	defer cornerMat.Close()

	dstCornerMat := gocv.NewMat()
	defer dstCornerMat.Close()
	if err := gocv.PerspectiveTransform(cornerMat, &dstCornerMat, homography); err != nil {
		return nil, err
	}
	dstCornerVec := gocv.NewPoint2fVectorFromMat(dstCornerMat)
	defer dstCornerVec.Close()
	projected := dstCornerVec.ToPoints()
	if len(projected) < 4 {
		return nil, nil
	}

	minX := float64(projected[0].X)
	maxX := minX
	minY := float64(projected[0].Y)
	maxY := minY
	for i := 1; i < len(projected); i++ {
		minX = math.Min(minX, float64(projected[i].X))
		maxX = math.Max(maxX, float64(projected[i].X))
		minY = math.Min(minY, float64(projected[i].Y))
		maxY = math.Max(maxY, float64(projected[i].Y))
	}

	x := clampInt(int(math.Round(minX))+hb.Min.X, hb.Min.X, hb.Max.X-1)
	y := clampInt(int(math.Round(minY))+hb.Min.Y, hb.Min.Y, hb.Max.Y-1)
	w := clampInt(int(math.Round(maxX-minX)), 1, hb.Max.X-x)
	h := clampInt(int(math.Round(maxY-minY)), 1, hb.Max.Y-y)
	if w <= 0 || h <= 0 {
		return nil, nil
	}

	inlierCount := 0
	if !inlierMask.Empty() {
		rows := inlierMask.Rows()
		for i := 0; i < rows; i++ {
			if inlierMask.GetUCharAt(i, 0) > 0 {
				inlierCount++
			}
		}
	}
	if inlierCount == 0 {
		inlierCount = len(srcPoints)
	}

	inlierRatio := float64(inlierCount) / float64(len(srcPoints))
	avgDistance := distanceSum / float64(len(srcPoints))
	distanceScore := clamp01(1.0 - (avgDistance / 256.0))
	score := clamp01((inlierRatio + distanceScore) / 2.0)
	if score < req.Threshold {
		return nil, nil
	}

	results := []core.MatchCandidate{
		{
			X:     x,
			Y:     y,
			W:     w,
			H:     h,
			Score: score,
		},
	}
	if req.MaxResults > 0 && len(results) > req.MaxResults {
		results = results[:req.MaxResults]
	}
	return results, nil
}

func clampInt(v, minV, maxV int) int {
	if maxV < minV {
		return minV
	}
	if v < minV {
		return minV
	}
	if v > maxV {
		return maxV
	}
	return v
}

func clamp01(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
