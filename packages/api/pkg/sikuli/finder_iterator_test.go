package sikuli

import (
	"errors"
	"testing"
)

func testFinderWithTwoMatches(t *testing.T) (*Finder, *Pattern) {
	t.Helper()

	hay, err := NewImageFromMatrix("hay", [][]uint8{
		{10, 10, 10, 10, 10, 10, 10, 10},
		{10, 0, 255, 10, 10, 10, 10, 10},
		{10, 255, 0, 10, 0, 255, 10, 10},
		{10, 10, 10, 10, 255, 0, 10, 10},
		{10, 10, 10, 10, 10, 10, 10, 10},
	})
	if err != nil {
		t.Fatalf("new hay image: %v", err)
	}
	needle, err := NewImageFromMatrix("needle", [][]uint8{
		{0, 255},
		{255, 0},
	})
	if err != nil {
		t.Fatalf("new needle image: %v", err)
	}
	p, err := NewPattern(needle)
	if err != nil {
		t.Fatalf("new pattern: %v", err)
	}
	p.Exact()
	f, err := NewFinder(hay)
	if err != nil {
		t.Fatalf("new finder: %v", err)
	}
	return f, p
}

func TestFinderCompatibilityIteratorLifecycle(t *testing.T) {
	f, p := testFinderWithTwoMatches(t)

	if err := f.IterateAll(p); err != nil {
		t.Fatalf("iterate all failed: %v", err)
	}
	if !f.HasNext() {
		t.Fatalf("expected iterator to have next")
	}

	first, ok := f.Next()
	if !ok {
		t.Fatalf("expected first iterator match")
	}
	second, ok := f.Next()
	if !ok {
		t.Fatalf("expected second iterator match")
	}
	if first.Index != 0 || second.Index != 1 {
		t.Fatalf("iterator order mismatch: first=%+v second=%+v", first, second)
	}
	if _, ok := f.Next(); ok {
		t.Fatalf("expected iterator exhaustion")
	}
	if f.HasNext() {
		t.Fatalf("hasNext should be false after exhaustion")
	}

	last := f.LastMatches()
	if len(last) != 2 || last[0].Index != 0 || last[1].Index != 1 {
		t.Fatalf("last matches should remain stable after iteration: %+v", last)
	}

	f.Reset()
	if !f.HasNext() {
		t.Fatalf("expected iterator to rewind after reset")
	}
	again, ok := f.Next()
	if !ok || again.Index != 0 {
		t.Fatalf("reset iteration mismatch: ok=%v match=%+v", ok, again)
	}

	f.Destroy()
	if f.HasNext() {
		t.Fatalf("destroy should clear iterator state")
	}
	if last := f.LastMatches(); last != nil {
		t.Fatalf("destroy should clear last matches: %+v", last)
	}

	if err := f.IterateAll(p); err != nil {
		t.Fatalf("finder should remain reusable after destroy: %v", err)
	}
	if !f.HasNext() {
		t.Fatalf("expected iterator to work after destroy+reuse")
	}
}

func TestFinderIterateMissDoesNotRaiseFindFailed(t *testing.T) {
	f, p := testFinderWithTwoMatches(t)

	missingNeedle, err := NewImageFromMatrix("missing", [][]uint8{
		{255, 255},
		{255, 255},
	})
	if err != nil {
		t.Fatalf("new missing needle: %v", err)
	}
	missingPattern, err := NewPattern(missingNeedle)
	if err != nil {
		t.Fatalf("new missing pattern: %v", err)
	}
	missingPattern.Exact()

	if _, err := f.Find(missingPattern); !errors.Is(err, ErrFindFailed) {
		t.Fatalf("find miss should still return ErrFindFailed, got=%v", err)
	}
	if err := f.Iterate(missingPattern); err != nil {
		t.Fatalf("iterate miss should not fail, got=%v", err)
	}
	if f.HasNext() {
		t.Fatalf("iterator miss should have no next element")
	}
	if last := f.LastMatches(); last != nil {
		t.Fatalf("iterator miss should clear last matches: %+v", last)
	}

	if err := f.Iterate(p); err != nil {
		t.Fatalf("iterate best-match failed: %v", err)
	}
	if !f.HasNext() {
		t.Fatalf("expected one best-match iterator result")
	}
	match, ok := f.Next()
	if !ok {
		t.Fatalf("expected best match after iterate")
	}
	if match.Index != 0 {
		t.Fatalf("best-match iterator index mismatch: %+v", match)
	}
	if _, ok := f.Next(); ok {
		t.Fatalf("single-match iterator should be exhausted after one result")
	}
}
