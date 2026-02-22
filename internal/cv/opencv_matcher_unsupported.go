//go:build !opencv

package cv

import (
	"fmt"

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
	return nil, fmt.Errorf("%w: build with -tags opencv", core.ErrMatcherUnsupported)
}
