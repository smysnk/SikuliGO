package observe

import (
	"image"
	"image/color"
	"testing"
	"time"

	"github.com/sikulix/portgo/internal/core"
)

type fakeClock struct {
	current time.Time
}

func (c *fakeClock) now() time.Time {
	return c.current
}

func (c *fakeClock) sleep(d time.Duration) {
	if d > 0 {
		c.current = c.current.Add(d)
	}
}

func TestPollingBackendConformanceAppearTiming(t *testing.T) {
	source := grayFromRows([][]uint8{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	})
	pattern := grayFromRows([][]uint8{
		{10, 200},
		{220, 15},
	})

	start := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	clock := &fakeClock{current: start}
	backend := newPollingBackend(nil)
	backend.now = clock.now
	backend.sleep = clock.sleep
	backend.onPoll = func(iteration int) {
		if iteration != 2 {
			return
		}
		writePattern(source, pattern, 1, 2)
	}

	events, err := backend.Observe(core.ObserveRequest{
		Source:   source,
		Region:   source.Bounds(),
		Pattern:  pattern,
		Event:    core.ObserveEventAppear,
		Interval: 50 * time.Millisecond,
		Timeout:  500 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("observe appear failed: %v", err)
	}
	if len(events) != 1 {
		t.Fatalf("expected 1 appear event, got=%d", len(events))
	}
	ev := events[0]
	if ev.Event != core.ObserveEventAppear {
		t.Fatalf("expected appear event, got=%q", ev.Event)
	}
	if ev.X != 1 || ev.Y != 2 || ev.W != 2 || ev.H != 2 {
		t.Fatalf("appear match geometry mismatch: %+v", ev)
	}
	wantTS := start.Add(100 * time.Millisecond)
	if !ev.Timestamp.Equal(wantTS) {
		t.Fatalf("appear timestamp mismatch: got=%v want=%v", ev.Timestamp, wantTS)
	}
}

func TestPollingBackendConformanceVanishTiming(t *testing.T) {
	source := grayFromRows([][]uint8{
		{0, 0, 0, 0, 0},
		{0, 0, 9, 180, 0},
		{0, 0, 220, 11, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	})
	pattern := grayFromRows([][]uint8{
		{9, 180},
		{220, 11},
	})

	start := time.Date(2026, 1, 1, 11, 0, 0, 0, time.UTC)
	clock := &fakeClock{current: start}
	backend := newPollingBackend(nil)
	backend.now = clock.now
	backend.sleep = clock.sleep
	backend.onPoll = func(iteration int) {
		if iteration != 3 {
			return
		}
		clearPattern(source, 2, 1, 2, 2)
	}

	events, err := backend.Observe(core.ObserveRequest{
		Source:   source,
		Region:   source.Bounds(),
		Pattern:  pattern,
		Event:    core.ObserveEventVanish,
		Interval: 40 * time.Millisecond,
		Timeout:  500 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("observe vanish failed: %v", err)
	}
	if len(events) != 1 {
		t.Fatalf("expected 1 vanish event, got=%d", len(events))
	}
	ev := events[0]
	if ev.Event != core.ObserveEventVanish {
		t.Fatalf("expected vanish event, got=%q", ev.Event)
	}
	if ev.X != 2 || ev.Y != 1 || ev.W != 2 || ev.H != 2 {
		t.Fatalf("vanish geometry mismatch: %+v", ev)
	}
	wantTS := start.Add(120 * time.Millisecond)
	if !ev.Timestamp.Equal(wantTS) {
		t.Fatalf("vanish timestamp mismatch: got=%v want=%v", ev.Timestamp, wantTS)
	}
}

func TestPollingBackendConformanceChangeTiming(t *testing.T) {
	source := grayFromRows([][]uint8{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	})

	start := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	clock := &fakeClock{current: start}
	backend := newPollingBackend(nil)
	backend.now = clock.now
	backend.sleep = clock.sleep
	backend.onPoll = func(iteration int) {
		if iteration != 1 {
			return
		}
		source.SetGray(1, 1, color.Gray{Y: 250})
	}

	events, err := backend.Observe(core.ObserveRequest{
		Source:   source,
		Region:   source.Bounds(),
		Event:    core.ObserveEventChange,
		Interval: 75 * time.Millisecond,
		Timeout:  300 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("observe change failed: %v", err)
	}
	if len(events) != 1 {
		t.Fatalf("expected 1 change event, got=%d", len(events))
	}
	ev := events[0]
	if ev.Event != core.ObserveEventChange {
		t.Fatalf("expected change event, got=%q", ev.Event)
	}
	if ev.Score <= 0 {
		t.Fatalf("expected positive change score, got=%v", ev.Score)
	}
	wantTS := start.Add(75 * time.Millisecond)
	if !ev.Timestamp.Equal(wantTS) {
		t.Fatalf("change timestamp mismatch: got=%v want=%v", ev.Timestamp, wantTS)
	}
}

func TestPollingBackendConformanceTimeout(t *testing.T) {
	source := grayFromRows([][]uint8{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	})
	pattern := grayFromRows([][]uint8{
		{255, 255},
		{255, 255},
	})

	start := time.Date(2026, 1, 1, 13, 0, 0, 0, time.UTC)
	clock := &fakeClock{current: start}
	backend := newPollingBackend(nil)
	backend.now = clock.now
	backend.sleep = clock.sleep

	events, err := backend.Observe(core.ObserveRequest{
		Source:   source,
		Region:   source.Bounds(),
		Pattern:  pattern,
		Event:    core.ObserveEventAppear,
		Interval: 30 * time.Millisecond,
		Timeout:  90 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("observe timeout case failed: %v", err)
	}
	if len(events) != 0 {
		t.Fatalf("expected timeout with no events, got=%+v", events)
	}
	wantNow := start.Add(90 * time.Millisecond)
	if !clock.current.Equal(wantNow) {
		t.Fatalf("timeout clock mismatch: got=%v want=%v", clock.current, wantNow)
	}
}

func grayFromRows(rows [][]uint8) *image.Gray {
	h := len(rows)
	w := len(rows[0])
	g := image.NewGray(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			g.SetGray(x, y, color.Gray{Y: rows[y][x]})
		}
	}
	return g
}

func writePattern(dst, pattern *image.Gray, x0, y0 int) {
	pb := pattern.Bounds()
	for y := 0; y < pb.Dy(); y++ {
		for x := 0; x < pb.Dx(); x++ {
			dst.SetGray(x0+x, y0+y, pattern.GrayAt(pb.Min.X+x, pb.Min.Y+y))
		}
	}
}

func clearPattern(dst *image.Gray, x0, y0, w, h int) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dst.SetGray(x0+x, y0+y, color.Gray{Y: 0})
		}
	}
}
