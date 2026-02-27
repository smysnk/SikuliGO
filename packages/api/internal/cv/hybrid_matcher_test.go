package cv

import (
	"errors"
	"image"
	"testing"

	"github.com/smysnk/sikuligo/internal/core"
)

type stubMatcher struct {
	results []core.MatchCandidate
	err     error
}

func (s stubMatcher) Find(_ core.SearchRequest) ([]core.MatchCandidate, error) {
	return s.results, s.err
}

func TestHybridMatcherPrimaryWins(t *testing.T) {
	m := NewHybridMatcher(
		stubMatcher{results: []core.MatchCandidate{{X: 1, Y: 2, W: 3, H: 4, Score: 0.9}}},
		stubMatcher{results: []core.MatchCandidate{{X: 5, Y: 6, W: 3, H: 4, Score: 0.8}}},
	)
	got, err := m.Find(core.SearchRequest{
		Haystack:     image.NewGray(image.Rect(0, 0, 2, 2)),
		Needle:       image.NewGray(image.Rect(0, 0, 1, 1)),
		Threshold:    0,
		ResizeFactor: 1,
	})
	if err != nil {
		t.Fatalf("find failed: %v", err)
	}
	if len(got) != 1 || got[0].X != 1 || got[0].Y != 2 {
		t.Fatalf("unexpected primary result: %+v", got)
	}
}

func TestHybridMatcherFallbackOnNoResults(t *testing.T) {
	m := NewHybridMatcher(
		stubMatcher{results: nil},
		stubMatcher{results: []core.MatchCandidate{{X: 9, Y: 9, W: 1, H: 1, Score: 0.7}}},
	)
	got, err := m.Find(core.SearchRequest{
		Haystack:     image.NewGray(image.Rect(0, 0, 2, 2)),
		Needle:       image.NewGray(image.Rect(0, 0, 1, 1)),
		Threshold:    0,
		ResizeFactor: 1,
	})
	if err != nil {
		t.Fatalf("find failed: %v", err)
	}
	if len(got) != 1 || got[0].X != 9 || got[0].Y != 9 {
		t.Fatalf("unexpected fallback result: %+v", got)
	}
}

func TestHybridMatcherFallbackOnUnsupported(t *testing.T) {
	m := NewHybridMatcher(
		stubMatcher{err: core.ErrMatcherUnsupported},
		stubMatcher{results: []core.MatchCandidate{{X: 3, Y: 4, W: 1, H: 1, Score: 0.7}}},
	)
	got, err := m.Find(core.SearchRequest{
		Haystack:     image.NewGray(image.Rect(0, 0, 2, 2)),
		Needle:       image.NewGray(image.Rect(0, 0, 1, 1)),
		Threshold:    0,
		ResizeFactor: 1,
	})
	if err != nil {
		t.Fatalf("find failed: %v", err)
	}
	if len(got) != 1 || got[0].X != 3 || got[0].Y != 4 {
		t.Fatalf("unexpected fallback result: %+v", got)
	}
}

func TestHybridMatcherReturnsPrimaryError(t *testing.T) {
	wantErr := errors.New("boom")
	m := NewHybridMatcher(
		stubMatcher{err: wantErr},
		stubMatcher{results: []core.MatchCandidate{{X: 1, Y: 1, W: 1, H: 1, Score: 0.8}}},
	)
	_, err := m.Find(core.SearchRequest{
		Haystack:     image.NewGray(image.Rect(0, 0, 2, 2)),
		Needle:       image.NewGray(image.Rect(0, 0, 1, 1)),
		Threshold:    0,
		ResizeFactor: 1,
	})
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected primary error, got %v", err)
	}
}
