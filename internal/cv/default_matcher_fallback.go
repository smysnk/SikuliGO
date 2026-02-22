//go:build !opencv

package cv

import "github.com/sikulix/portgo/internal/core"

func newDefaultMatcher() core.Matcher {
	return NewNCCMatcher()
}
