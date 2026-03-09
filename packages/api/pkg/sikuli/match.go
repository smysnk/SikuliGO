package sikuli

import (
	"fmt"
	"time"
)

type Match struct {
	Rect
	Score  float64
	Target Point
	Index  int

	runtime       *Runtime
	screenID      int
	hasScreenID   bool
	screenBounds  Rect
	matcherEngine MatcherEngine
	waitScanRate  float64
}

func NewMatch(x, y, w, h int, score float64, off Point) Match {
	target := Point{
		X: x + w/2 + off.X,
		Y: y + h/2 + off.Y,
	}
	return Match{
		Rect:   NewRect(x, y, w, h),
		Score:  score,
		Target: target,
	}
}

func (m Match) String() string {
	return fmt.Sprintf("M[%d,%d %dx%d score=%.4f]", m.X, m.Y, m.W, m.H, m.Score)
}

// Bounds returns the match as a region-like value so callers can treat a live
// match the same way SikuliX documents region-capable match results.
func (m Match) Bounds() Region {
	region := NewRegion(m.X, m.Y, m.W, m.H)
	if m.waitScanRate > 0 {
		region.WaitScanRate = m.waitScanRate
	}
	return region
}

// Region is an alias for Bounds for parity-friendly call sites.
func (m Match) Region() Region {
	return m.Bounds()
}

// Center returns the geometric center of the matched rectangle.
func (m Match) Center() Point {
	return m.Bounds().Center()
}

// TargetPoint returns the resolved click target point for this match.
func (m Match) TargetPoint() Point {
	return m.Target
}

// Live reports whether this match is bound to a running sikuli-go runtime.
func (m Match) Live() bool {
	return m.runtime != nil
}

// Capture captures the currently matched live region.
func (m Match) Capture() (*Image, error) {
	return m.liveRegion().Capture()
}

// Find searches again within the matched live region.
func (m Match) Find(pattern *Pattern) (Match, error) {
	return m.liveRegion().Find(pattern)
}

// Exists probes within the matched live region.
func (m Match) Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error) {
	return m.liveRegion().Exists(pattern, timeout)
}

// Has reports whether the target exists within the matched live region.
func (m Match) Has(pattern *Pattern, timeout time.Duration) (bool, error) {
	return m.liveRegion().Has(pattern, timeout)
}

// Wait waits for the target within the matched live region.
func (m Match) Wait(pattern *Pattern, timeout time.Duration) (Match, error) {
	return m.liveRegion().Wait(pattern, timeout)
}

// WaitVanish waits for the target to disappear from the matched live region.
func (m Match) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error) {
	return m.liveRegion().WaitVanish(pattern, timeout)
}

// ReadText runs OCR inside the matched live region.
func (m Match) ReadText(params OCRParams) (string, error) {
	return m.liveRegion().ReadText(params)
}

// FindText searches OCR text inside the matched live region.
func (m Match) FindText(query string, params OCRParams) ([]TextMatch, error) {
	return m.liveRegion().FindText(query, params)
}

// CollectWords returns OCR word-level results inside the matched live region.
func (m Match) CollectWords(params OCRParams) ([]OCRWord, error) {
	return m.liveRegion().CollectWords(params)
}

// CollectLines returns OCR line-level results inside the matched live region.
func (m Match) CollectLines(params OCRParams) ([]OCRLine, error) {
	return m.liveRegion().CollectLines(params)
}

// MoveMouse moves the pointer to the match target point.
func (m Match) MoveMouse(opts InputOptions) error {
	if !m.Live() {
		return ErrRuntimeUnavailable
	}
	return m.runtime.moveMouse(m.TargetPoint(), opts)
}

// Hover is a parity-friendly alias for MoveMouse.
func (m Match) Hover(opts InputOptions) error {
	return m.MoveMouse(opts)
}

// Click clicks the match target point.
func (m Match) Click(opts InputOptions) error {
	if !m.Live() {
		return ErrRuntimeUnavailable
	}
	return m.runtime.click(m.TargetPoint(), opts)
}

// RightClick clicks the match target point with the right mouse button.
func (m Match) RightClick(opts InputOptions) error {
	if !m.Live() {
		return ErrRuntimeUnavailable
	}
	return m.runtime.rightClick(m.TargetPoint(), opts)
}

// DoubleClick performs two click actions against the match target point.
func (m Match) DoubleClick(opts InputOptions) error {
	if !m.Live() {
		return ErrRuntimeUnavailable
	}
	return m.runtime.doubleClick(m.TargetPoint(), opts)
}

// MouseDown presses and holds the mouse button at the match target point.
func (m Match) MouseDown(opts InputOptions) error {
	if !m.Live() {
		return ErrRuntimeUnavailable
	}
	return m.runtime.mouseDown(m.TargetPoint(), opts)
}

// MouseUp releases the mouse button at the match target point.
func (m Match) MouseUp(opts InputOptions) error {
	if !m.Live() {
		return ErrRuntimeUnavailable
	}
	return m.runtime.mouseUp(m.TargetPoint(), opts)
}

// TypeText focuses the match target point and types text into it.
func (m Match) TypeText(text string, opts InputOptions) error {
	if !m.Live() {
		return ErrRuntimeUnavailable
	}
	if err := m.runtime.click(m.TargetPoint(), opts); err != nil {
		return err
	}
	return m.runtime.typeText(text, opts)
}

// Paste focuses the match target point and pastes text into it.
func (m Match) Paste(text string, opts InputOptions) error {
	if !m.Live() {
		return ErrRuntimeUnavailable
	}
	if err := m.runtime.click(m.TargetPoint(), opts); err != nil {
		return err
	}
	return m.runtime.pasteText(text, opts)
}

// DragDrop drags from the match target point to the target point.
func (m Match) DragDrop(target TargetPointProvider, opts InputOptions) error {
	if !m.Live() {
		return ErrRuntimeUnavailable
	}
	return m.runtime.dragDrop(m.TargetPoint(), target, opts)
}

// Wheel scrolls at the match target point.
func (m Match) Wheel(direction WheelDirection, steps int, opts InputOptions) error {
	if !m.Live() {
		return ErrRuntimeUnavailable
	}
	return m.runtime.wheel(m.TargetPoint(), direction, steps, opts)
}

// KeyDown holds the provided keys.
func (m Match) KeyDown(keys ...string) error {
	if !m.Live() {
		return ErrRuntimeUnavailable
	}
	return m.runtime.keyDown(keys...)
}

// KeyUp releases the provided keys.
func (m Match) KeyUp(keys ...string) error {
	if !m.Live() {
		return ErrRuntimeUnavailable
	}
	return m.runtime.keyUp(keys...)
}

func (m Match) liveRegion() LiveRegion {
	return LiveRegion{
		runtime:       m.runtime,
		region:        m.Bounds(),
		screenID:      m.screenID,
		hasScreenID:   m.hasScreenID,
		screenBounds:  m.screenBounds,
		matcherEngine: m.matcherEngine,
		waitScanRate:  m.waitScanRate,
	}
}
