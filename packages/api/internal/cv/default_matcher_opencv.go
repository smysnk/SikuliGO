//go:build opencv

package cv

import "github.com/smysnk/sikuligo/internal/core"

func newDefaultMatcher() core.Matcher {
	return NewOpenCVMatcher()
}
