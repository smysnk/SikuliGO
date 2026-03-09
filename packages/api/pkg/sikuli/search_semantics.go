package sikuli

import (
	"errors"
	"fmt"
	"time"
)

// SearchProbe returns the next match candidate for a parity search operation.
// Returning ErrFindFailed indicates a miss; any other error aborts the search.
type SearchProbe func() (Match, error)

// SearchExists applies the canonical sikuli-go parity contract to a search probe.
// Misses are reported as (Match{}, false, nil); timeout <= 0 performs one probe.
func SearchExists(probe SearchProbe, timeout, interval time.Duration) (Match, bool, error) {
	if timeout <= 0 {
		return searchExistsOnce(probe)
	}

	deadline := time.Now().Add(timeout)
	for {
		match, ok, err := searchExistsOnce(probe)
		if err != nil {
			return Match{}, false, err
		}
		if ok {
			return match, true, nil
		}
		if !time.Now().Before(deadline) {
			return Match{}, false, nil
		}
		time.Sleep(searchSleepInterval(interval, deadline))
	}
}

// SearchWait applies the canonical sikuli-go wait contract to a search probe.
// Misses are promoted to ErrTimeout once the wait budget is exhausted.
func SearchWait(probe SearchProbe, timeout, interval time.Duration) (Match, error) {
	match, ok, err := SearchExists(probe, timeout, interval)
	if err != nil {
		return Match{}, err
	}
	if !ok {
		return Match{}, ErrTimeout
	}
	return match, nil
}

// SearchWaitVanish applies the canonical vanish contract to a search probe.
// It returns true when the target is absent and false on timeout.
func SearchWaitVanish(probe SearchProbe, timeout, interval time.Duration) (bool, error) {
	checkOnce := func() (bool, error) {
		_, ok, err := searchExistsOnce(probe)
		if err != nil {
			return false, err
		}
		return !ok, nil
	}

	if timeout <= 0 {
		return checkOnce()
	}

	deadline := time.Now().Add(timeout)
	for {
		vanished, err := checkOnce()
		if err != nil {
			return false, err
		}
		if vanished {
			return true, nil
		}
		if !time.Now().Before(deadline) {
			return false, nil
		}
		time.Sleep(searchSleepInterval(interval, deadline))
	}
}

func searchExistsOnce(probe SearchProbe) (Match, bool, error) {
	if probe == nil {
		return Match{}, false, fmt.Errorf("%w: search probe is nil", ErrInvalidTarget)
	}
	match, err := probe()
	if err != nil {
		if errors.Is(err, ErrFindFailed) {
			return Match{}, false, nil
		}
		return Match{}, false, err
	}
	return match, true, nil
}

func searchSleepInterval(interval time.Duration, deadline time.Time) time.Duration {
	remaining := time.Until(deadline)
	if remaining <= 0 {
		return 0
	}
	if interval <= 0 {
		interval = 100 * time.Millisecond
	}
	if interval > remaining {
		return remaining
	}
	return interval
}
