package cv

import "github.com/sikulix/portgo/internal/core"

// NewDefaultMatcher returns the matcher backend used by default in Sikuli flows.
func NewDefaultMatcher() core.Matcher {
	return newDefaultMatcher()
}
