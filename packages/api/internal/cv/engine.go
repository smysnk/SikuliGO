package cv

import (
	"fmt"
	"strings"

	"github.com/smysnk/sikuligo/internal/core"
)

type MatcherEngine string

const (
	MatcherEngineTemplate MatcherEngine = "template"
	MatcherEngineORB      MatcherEngine = "orb"
	MatcherEngineHybrid   MatcherEngine = "hybrid"
)

func ParseMatcherEngine(raw string) (MatcherEngine, error) {
	engine := strings.ToLower(strings.TrimSpace(raw))
	switch engine {
	case "", string(MatcherEngineTemplate), "ncc", "opencv":
		return MatcherEngineTemplate, nil
	case string(MatcherEngineORB):
		return MatcherEngineORB, nil
	case string(MatcherEngineHybrid):
		return MatcherEngineHybrid, nil
	default:
		return "", fmt.Errorf("unsupported matcher engine %q", raw)
	}
}

func NewMatcherForEngine(engine MatcherEngine) (core.Matcher, error) {
	switch engine {
	case MatcherEngineTemplate:
		return NewDefaultMatcher(), nil
	case MatcherEngineORB:
		return NewORBMatcher(), nil
	case MatcherEngineHybrid:
		return NewHybridMatcher(NewDefaultMatcher(), NewORBMatcher()), nil
	default:
		return nil, fmt.Errorf("unsupported matcher engine %q", engine)
	}
}
