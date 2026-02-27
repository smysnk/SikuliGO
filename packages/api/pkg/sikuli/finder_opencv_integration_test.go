//go:build !opencv

package sikuli

import (
	"errors"
	"testing"

	"github.com/smysnk/sikuligo/internal/cv"
)

func TestFinderMapsOpenCVUnsupportedToBackendUnsupported(t *testing.T) {
	haystack, err := NewImageFromMatrix("haystack", [][]uint8{
		{10, 10, 10},
		{10, 10, 10},
		{10, 10, 10},
	})
	if err != nil {
		t.Fatalf("new haystack: %v", err)
	}
	needle, err := NewImageFromMatrix("needle", [][]uint8{
		{10},
	})
	if err != nil {
		t.Fatalf("new needle: %v", err)
	}
	pattern, err := NewPattern(needle)
	if err != nil {
		t.Fatalf("new pattern: %v", err)
	}

	finder, err := NewFinder(haystack)
	if err != nil {
		t.Fatalf("new finder: %v", err)
	}
	finder.SetMatcher(cv.NewOpenCVMatcher())

	_, err = finder.Find(pattern)
	if !errors.Is(err, ErrBackendUnsupported) {
		t.Fatalf("expected ErrBackendUnsupported for Find, got=%v", err)
	}

	_, err = finder.FindAll(pattern)
	if !errors.Is(err, ErrBackendUnsupported) {
		t.Fatalf("expected ErrBackendUnsupported for FindAll, got=%v", err)
	}
}
