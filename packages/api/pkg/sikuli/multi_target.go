package sikuli

import (
	"errors"
	"fmt"
	"time"

	"github.com/smysnk/sikuligo/internal/core"
	"github.com/smysnk/sikuligo/internal/cv"
)

type multiSearchProbe func() ([]Match, error)

func validatePatternList(patterns []*Pattern) error {
	if len(patterns) == 0 {
		return fmt.Errorf("%w: pattern list is empty", ErrInvalidTarget)
	}
	for i, pattern := range patterns {
		if pattern == nil || pattern.Image() == nil || pattern.Image().Gray() == nil {
			return fmt.Errorf("%w: pattern[%d] image is nil", ErrInvalidTarget, i)
		}
	}
	return nil
}

func multiSearchExists(probe multiSearchProbe) ([]Match, bool, error) {
	if probe == nil {
		return nil, false, fmt.Errorf("%w: search probe is nil", ErrInvalidTarget)
	}
	matches, err := probe()
	if err != nil {
		if errors.Is(err, ErrFindFailed) {
			return nil, false, nil
		}
		return nil, false, err
	}
	if len(matches) == 0 {
		return nil, false, nil
	}
	return matches, true, nil
}

func multiSearchWait(probe multiSearchProbe, timeout, interval time.Duration) ([]Match, error) {
	if timeout <= 0 {
		matches, ok, err := multiSearchExists(probe)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, ErrTimeout
		}
		return matches, nil
	}
	deadline := time.Now().Add(timeout)
	for {
		matches, ok, err := multiSearchExists(probe)
		if err != nil {
			return nil, err
		}
		if ok {
			return matches, nil
		}
		if !time.Now().Before(deadline) {
			return nil, ErrTimeout
		}
		time.Sleep(searchSleepInterval(interval, deadline))
	}
}

func bestMultiTargetMatch(matches []Match) (Match, error) {
	if len(matches) == 0 {
		return Match{}, ErrFindFailed
	}
	best := matches[0]
	for _, candidate := range matches[1:] {
		if multiTargetMatchLess(candidate, best) {
			best = candidate
		}
	}
	return best, nil
}

func multiTargetMatchLess(a, b Match) bool {
	if a.Score != b.Score {
		return a.Score > b.Score
	}
	if a.Index != b.Index {
		return a.Index < b.Index
	}
	if a.Y != b.Y {
		return a.Y < b.Y
	}
	if a.X != b.X {
		return a.X < b.X
	}
	if a.W != b.W {
		return a.W < b.W
	}
	return a.H < b.H
}

func (f *Finder) findAnyListMatches(patterns []*Pattern) ([]Match, error) {
	if err := validatePatternList(patterns); err != nil {
		return nil, err
	}
	matches := make([]Match, 0, len(patterns))
	for idx, pattern := range patterns {
		found, err := f.searchMatches(pattern, 1)
		if err != nil {
			return nil, err
		}
		if len(found) == 0 {
			continue
		}
		match := found[0]
		match.Index = idx
		matches = append(matches, match)
	}
	if len(matches) == 0 {
		return nil, ErrFindFailed
	}
	return matches, nil
}

func (f *Finder) FindAnyList(patterns []*Pattern) ([]Match, error) {
	matches, err := f.findAnyListMatches(patterns)
	if err != nil {
		if errors.Is(err, ErrFindFailed) {
			f.clearTraversal()
		}
		return nil, err
	}
	f.setTraversal(matches)
	return matches, nil
}

func (f *Finder) FindBestList(patterns []*Pattern) (Match, error) {
	matches, err := f.findAnyListMatches(patterns)
	if err != nil {
		if errors.Is(err, ErrFindFailed) {
			f.clearTraversal()
		}
		return Match{}, err
	}
	best, err := bestMultiTargetMatch(matches)
	if err != nil {
		f.clearTraversal()
		return Match{}, err
	}
	f.setTraversal([]Match{best})
	return best, nil
}

func (f *Finder) WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error) {
	matches, err := multiSearchWait(func() ([]Match, error) {
		return f.findAnyListMatches(patterns)
	}, timeout, finderWaitInterval())
	if err != nil {
		f.clearTraversal()
		return nil, err
	}
	f.setTraversal(matches)
	return matches, nil
}

func (f *Finder) WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error) {
	matches, err := multiSearchWait(func() ([]Match, error) {
		return f.findAnyListMatches(patterns)
	}, timeout, finderWaitInterval())
	if err != nil {
		f.clearTraversal()
		return Match{}, err
	}
	best, err := bestMultiTargetMatch(matches)
	if err != nil {
		f.clearTraversal()
		return Match{}, err
	}
	f.setTraversal([]Match{best})
	return best, nil
}

func (r Region) FindAnyList(source *Image, patterns []*Pattern) ([]Match, error) {
	finder, err := r.newFinder(source)
	if err != nil {
		return nil, err
	}
	return finder.FindAnyList(patterns)
}

func (r Region) FindBestList(source *Image, patterns []*Pattern) (Match, error) {
	finder, err := r.newFinder(source)
	if err != nil {
		return Match{}, err
	}
	return finder.FindBestList(patterns)
}

func (r Region) WaitAnyList(source *Image, patterns []*Pattern, timeout time.Duration) ([]Match, error) {
	effectiveTimeout := timeout
	if effectiveTimeout <= 0 {
		effectiveTimeout = time.Duration(r.AutoWaitTimeout * float64(time.Second))
	}
	return multiSearchWait(func() ([]Match, error) {
		finder, err := r.newFinder(source)
		if err != nil {
			return nil, err
		}
		return finder.findAnyListMatches(patterns)
	}, effectiveTimeout, r.waitInterval())
}

func (r Region) WaitBestList(source *Image, patterns []*Pattern, timeout time.Duration) (Match, error) {
	matches, err := r.WaitAnyList(source, patterns, timeout)
	if err != nil {
		return Match{}, err
	}
	return bestMultiTargetMatch(matches)
}

func matcherForEngine(engine MatcherEngine) (core.Matcher, error) {
	parsed, err := cv.ParseMatcherEngine(string(engine))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrBackendUnsupported, err)
	}
	matcher, err := cv.NewMatcherForEngine(parsed)
	if err != nil {
		if errors.Is(err, core.ErrMatcherUnsupported) {
			return nil, ErrBackendUnsupported
		}
		return nil, fmt.Errorf("%w: %v", ErrBackendUnsupported, err)
	}
	return matcher, nil
}

func (r LiveRegion) effectiveMatcherEngine() MatcherEngine {
	if r.matcherEngine != "" {
		return r.matcherEngine
	}
	if r.runtime != nil {
		return r.runtime.defaultMatcherEngine()
	}
	return MatcherEngineHybrid
}

func (r LiveRegion) captureFinder() (*Finder, error) {
	capture, err := r.Capture()
	if err != nil {
		return nil, err
	}
	finder, err := NewFinder(capture)
	if err != nil {
		return nil, err
	}
	matcher, err := matcherForEngine(r.effectiveMatcherEngine())
	if err != nil {
		return nil, err
	}
	finder.SetMatcher(matcher)
	return finder, nil
}

func (r LiveRegion) bindCapturedMatch(match Match) Match {
	if match == (Match{}) {
		return Match{}
	}
	match.Rect = NewRect(match.X+r.region.X, match.Y+r.region.Y, match.W, match.H)
	match.Target = NewPoint(match.Target.X+r.region.X, match.Target.Y+r.region.Y)
	return r.bindMatch(match)
}

func (r LiveRegion) findAnyListCapture(patterns []*Pattern) ([]Match, error) {
	if err := r.ensure(); err != nil {
		return nil, err
	}
	finder, err := r.captureFinder()
	if err != nil {
		return nil, err
	}
	matches, err := finder.findAnyListMatches(patterns)
	if err != nil {
		return nil, err
	}
	bound := make([]Match, len(matches))
	for i, match := range matches {
		bound[i] = r.bindCapturedMatch(match)
	}
	return bound, nil
}

func (r LiveRegion) FindAnyList(patterns []*Pattern) ([]Match, error) {
	return r.findAnyListCapture(patterns)
}

func (r LiveRegion) FindBestList(patterns []*Pattern) (Match, error) {
	matches, err := r.findAnyListCapture(patterns)
	if err != nil {
		return Match{}, err
	}
	return bestMultiTargetMatch(matches)
}

func (r LiveRegion) WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error) {
	return multiSearchWait(func() ([]Match, error) {
		return r.findAnyListCapture(patterns)
	}, timeout, r.waitInterval())
}

func (r LiveRegion) WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error) {
	matches, err := r.WaitAnyList(patterns, timeout)
	if err != nil {
		return Match{}, err
	}
	return bestMultiTargetMatch(matches)
}

func (s Screen) FindAnyList(patterns []*Pattern) ([]Match, error) {
	return s.FullRegion().FindAnyList(patterns)
}

func (s Screen) FindBestList(patterns []*Pattern) (Match, error) {
	return s.FullRegion().FindBestList(patterns)
}

func (s Screen) WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error) {
	return s.FullRegion().WaitAnyList(patterns, timeout)
}

func (s Screen) WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error) {
	return s.FullRegion().WaitBestList(patterns, timeout)
}

func (m Match) FindAnyList(patterns []*Pattern) ([]Match, error) {
	return m.liveRegion().FindAnyList(patterns)
}

func (m Match) FindBestList(patterns []*Pattern) (Match, error) {
	return m.liveRegion().FindBestList(patterns)
}

func (m Match) WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error) {
	return m.liveRegion().WaitAnyList(patterns, timeout)
}

func (m Match) WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error) {
	return m.liveRegion().WaitBestList(patterns, timeout)
}
