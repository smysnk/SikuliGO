package sikuli

import (
	"fmt"
	"sort"

	"github.com/sikulix/portgo/internal/core"
	"github.com/sikulix/portgo/internal/cv"
)

type Finder struct {
	source  *Image
	matcher core.Matcher
	last    []Match
}

func NewFinder(source *Image) (*Finder, error) {
	if source == nil || source.Gray() == nil {
		return nil, fmt.Errorf("%w: source image is nil", ErrInvalidTarget)
	}
	return &Finder{
		source:  source,
		matcher: cv.NewNCCMatcher(),
		last:    nil,
	}, nil
}

func (f *Finder) SetMatcher(m core.Matcher) {
	if m == nil {
		return
	}
	f.matcher = m
}

func (f *Finder) Find(pattern *Pattern) (Match, error) {
	req, err := f.buildRequest(pattern, 1)
	if err != nil {
		return Match{}, err
	}
	rawMatches, err := f.matcher.Find(req)
	if err != nil {
		return Match{}, err
	}
	if len(rawMatches) == 0 {
		f.last = nil
		return Match{}, ErrFindFailed
	}
	match := toMatch(rawMatches[0], pattern.Offset())
	match.Index = 0
	f.last = []Match{match}
	return match, nil
}

func (f *Finder) FindAll(pattern *Pattern) ([]Match, error) {
	req, err := f.buildRequest(pattern, 0)
	if err != nil {
		return nil, err
	}
	rawMatches, err := f.matcher.Find(req)
	if err != nil {
		return nil, err
	}
	matches := make([]Match, 0, len(rawMatches))
	for i, m := range rawMatches {
		out := toMatch(m, pattern.Offset())
		out.Index = i
		matches = append(matches, out)
	}
	f.last = matches
	return matches, nil
}

// Exists returns the first match when present. Missing targets return (Match{}, false, nil).
func (f *Finder) Exists(pattern *Pattern) (Match, bool, error) {
	match, err := f.Find(pattern)
	if err != nil {
		if err == ErrFindFailed {
			return Match{}, false, nil
		}
		return Match{}, false, err
	}
	return match, true, nil
}

// Has reports whether the target exists and bubbles non-find errors.
func (f *Finder) Has(pattern *Pattern) (bool, error) {
	_, ok, err := f.Exists(pattern)
	return ok, err
}

func (f *Finder) LastMatches() []Match {
	if len(f.last) == 0 {
		return nil
	}
	out := make([]Match, len(f.last))
	copy(out, f.last)
	return out
}

func (f *Finder) buildRequest(pattern *Pattern, maxResults int) (core.SearchRequest, error) {
	if f == nil || f.source == nil || f.source.Gray() == nil {
		return core.SearchRequest{}, fmt.Errorf("%w: source image is nil", ErrInvalidTarget)
	}
	if pattern == nil || pattern.Image() == nil || pattern.Image().Gray() == nil {
		return core.SearchRequest{}, fmt.Errorf("%w: pattern image is nil", ErrInvalidTarget)
	}
	req := core.SearchRequest{
		Haystack:     f.source.Gray(),
		Needle:       pattern.Image().Gray(),
		Mask:         pattern.Mask(),
		Threshold:    pattern.Similarity(),
		ResizeFactor: pattern.ResizeFactor(),
		MaxResults:   maxResults,
	}
	return req, nil
}

func toMatch(candidate core.MatchCandidate, off Point) Match {
	return NewMatch(candidate.X, candidate.Y, candidate.W, candidate.H, candidate.Score, off)
}

// SortMatchesByRowColumn keeps parity with Java helper behavior for "by row".
func SortMatchesByRowColumn(matches []Match) {
	sort.Slice(matches, func(i, j int) bool {
		if matches[i].Y == matches[j].Y {
			return matches[i].X < matches[j].X
		}
		return matches[i].Y < matches[j].Y
	})
}

// SortMatchesByColumnRow keeps parity with Java helper behavior for "by column".
func SortMatchesByColumnRow(matches []Match) {
	sort.Slice(matches, func(i, j int) bool {
		if matches[i].X == matches[j].X {
			return matches[i].Y < matches[j].Y
		}
		return matches[i].X < matches[j].X
	})
}
