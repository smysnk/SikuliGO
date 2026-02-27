//go:build !opencv

package cv

import (
	"fmt"

	"github.com/smysnk/sikuligo/internal/core"
)

type ORBMatcher struct{}

func NewORBMatcher() *ORBMatcher {
	return &ORBMatcher{}
}

func (m *ORBMatcher) Find(req core.SearchRequest) ([]core.MatchCandidate, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("%w: ORB matcher requires build with -tags opencv", core.ErrMatcherUnsupported)
}
