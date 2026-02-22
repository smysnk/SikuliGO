//go:build !opencv

package cv

import "testing"

func TestNewDefaultMatcherFallsBackToNCCWithoutOpenCVTag(t *testing.T) {
	m := NewDefaultMatcher()
	if m == nil {
		t.Fatalf("expected matcher, got nil")
	}
	if _, ok := m.(*NCCMatcher); !ok {
		t.Fatalf("expected *NCCMatcher default without opencv tag, got %T", m)
	}
}
