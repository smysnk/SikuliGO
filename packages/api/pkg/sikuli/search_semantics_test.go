package sikuli

import (
	"errors"
	"testing"
	"time"
)

func TestSearchExistsAndWaitParityContract(t *testing.T) {
	t.Run("exists_single_probe_miss_returns_false_without_error", func(t *testing.T) {
		calls := 0
		match, ok, err := SearchExists(func() (Match, error) {
			calls++
			return Match{}, ErrFindFailed
		}, 0, time.Millisecond)
		if err != nil {
			t.Fatalf("exists miss should not error: %v", err)
		}
		if ok {
			t.Fatalf("expected ok=false for miss")
		}
		if match != (Match{}) {
			t.Fatalf("expected zero match on miss: %+v", match)
		}
		if calls != 1 {
			t.Fatalf("expected one probe, got %d", calls)
		}
	})

	t.Run("exists_retries_until_match", func(t *testing.T) {
		calls := 0
		match, ok, err := SearchExists(func() (Match, error) {
			calls++
			if calls < 3 {
				return Match{}, ErrFindFailed
			}
			return NewMatch(4, 5, 6, 7, 0.9, Point{}), nil
		}, 20*time.Millisecond, time.Millisecond)
		if err != nil {
			t.Fatalf("exists retry should not error: %v", err)
		}
		if !ok {
			t.Fatalf("expected ok=true after retry")
		}
		if match.X != 4 || match.Y != 5 {
			t.Fatalf("unexpected match after retry: %+v", match)
		}
		if calls != 3 {
			t.Fatalf("expected three probes, got %d", calls)
		}
	})

	t.Run("wait_promotes_miss_to_timeout", func(t *testing.T) {
		_, err := SearchWait(func() (Match, error) {
			return Match{}, ErrFindFailed
		}, 2*time.Millisecond, time.Millisecond)
		if !errors.Is(err, ErrTimeout) {
			t.Fatalf("expected ErrTimeout, got %v", err)
		}
	})

	t.Run("wait_vanish_times_out_without_error", func(t *testing.T) {
		vanished, err := SearchWaitVanish(func() (Match, error) {
			return NewMatch(1, 2, 3, 4, 1, Point{}), nil
		}, 2*time.Millisecond, time.Millisecond)
		if err != nil {
			t.Fatalf("wait vanish timeout should not error: %v", err)
		}
		if vanished {
			t.Fatalf("expected vanished=false on timeout")
		}
	})
}
