//go:build opencv

package cv

import "testing"

func TestNewDefaultMatcherUsesOpenCVWithTag(t *testing.T) {
	m := NewDefaultMatcher()
	if m == nil {
		t.Fatalf("expected matcher, got nil")
	}
	if _, ok := m.(*OpenCVMatcher); !ok {
		t.Fatalf("expected *OpenCVMatcher default with opencv tag, got %T", m)
	}
}
