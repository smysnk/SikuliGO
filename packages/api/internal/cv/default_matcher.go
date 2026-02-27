package cv

import "github.com/smysnk/sikuligo/internal/core"

// NewDefaultMatcher returns the matcher backend used by default in Sikuli flows.
func NewDefaultMatcher() core.Matcher {
	return newDefaultMatcher()
}
