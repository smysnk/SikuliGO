package testharness

import (
	"image"
	"testing"

	"github.com/sikulix/portgo/internal/core"
	"github.com/sikulix/portgo/internal/cv"
)

func TestGoldenMatcherCorpus(t *testing.T) {
	cases, err := LoadCorpus()
	if err != nil {
		t.Fatalf("load corpus: %v", err)
	}
	matcher := cv.NewNCCMatcher()
	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			hay, err := MatrixToGray(tc.Haystack)
			if err != nil {
				t.Fatalf("haystack: %v", err)
			}
			needle, err := MatrixToGray(tc.Needle)
			if err != nil {
				t.Fatalf("needle: %v", err)
			}
			var maskImg = (*image.Gray)(nil)
			if len(tc.Mask) > 0 {
				maskImg, err = MatrixToGray(tc.Mask)
				if err != nil {
					t.Fatalf("mask: %v", err)
				}
			}

			req := core.SearchRequest{
				Haystack:     hay,
				Needle:       needle,
				Mask:         maskImg,
				Threshold:    tc.Threshold,
				ResizeFactor: tc.ResizeFactor,
				MaxResults:   tc.MaxResults,
			}
			got, err := matcher.Find(req)
			if err != nil {
				t.Fatalf("find: %v", err)
			}
			if err := CompareMatches(got, tc.Expected, CompareOptions{ScoreTolerance: 0.001}); err != nil {
				t.Fatalf("compare: %v", err)
			}
		})
	}
}
