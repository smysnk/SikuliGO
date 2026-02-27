package cv

import (
	"errors"
	"fmt"

	"github.com/smysnk/sikuligo/internal/core"
)

type HybridMatcher struct {
	primary  core.Matcher
	fallback core.Matcher
}

func NewHybridMatcher(primary, fallback core.Matcher) *HybridMatcher {
	return &HybridMatcher{
		primary:  primary,
		fallback: fallback,
	}
}

func (m *HybridMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error) {
	if m == nil {
		return nil, fmt.Errorf("matcher is nil")
	}
	if m.primary == nil && m.fallback == nil {
		return nil, fmt.Errorf("no matcher backends configured")
	}

	results, err := runMatcher(m.primary, req)
	if err == nil && len(results) > 0 {
		return results, nil
	}
	if err == nil && len(results) == 0 {
		return runMatcher(m.fallback, req)
	}
	if errors.Is(err, core.ErrMatcherUnsupported) {
		return runMatcher(m.fallback, req)
	}
	return nil, err
}

func runMatcher(m core.Matcher, req core.SearchRequest) ([]core.MatchCandidate, error) {
	if m == nil {
		return nil, nil
	}
	return m.Find(req)
}
