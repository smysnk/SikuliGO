package sikuli

import (
	"errors"
	"fmt"
	"time"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

type Rect struct {
	X int
	Y int
	W int
	H int
}

func NewRect(x, y, w, h int) Rect {
	return Rect{X: x, Y: y, W: w, H: h}
}

func (r Rect) Empty() bool {
	return r.W <= 0 || r.H <= 0
}

func (r Rect) Contains(p Point) bool {
	return p.X >= r.X && p.Y >= r.Y && p.X < r.X+r.W && p.Y < r.Y+r.H
}

func (r Rect) String() string {
	return fmt.Sprintf("R[%d,%d %dx%d]", r.X, r.Y, r.W, r.H)
}

type Region struct {
	Rect
	ThrowException  bool
	AutoWaitTimeout float64
	WaitScanRate    float64
	ObserveScanRate float64
}

func NewRegion(x, y, w, h int) Region {
	return Region{
		Rect:            NewRect(x, y, w, h),
		ThrowException:  true,
		AutoWaitTimeout: DefaultAutoWaitTimeout,
		WaitScanRate:    DefaultWaitScanRate,
		ObserveScanRate: DefaultObserveScanRate,
	}
}

func (r Region) Center() Point {
	return Point{
		X: r.X + r.W/2,
		Y: r.Y + r.H/2,
	}
}

func (r Region) Grow(dx, dy int) Region {
	return NewRegion(r.X-dx, r.Y-dy, r.W+dx*2, r.H+dy*2)
}

func (r Region) Offset(dx, dy int) Region {
	return NewRegion(r.X+dx, r.Y+dy, r.W, r.H)
}

func (r Region) MoveTo(x, y int) Region {
	return NewRegion(x, y, r.W, r.H)
}

func (r Region) SetSize(w, h int) Region {
	if w < 0 {
		w = 0
	}
	if h < 0 {
		h = 0
	}
	return NewRegion(r.X, r.Y, w, h)
}

func (r Region) Contains(p Point) bool {
	return r.Rect.Contains(p)
}

func (r Region) ContainsRegion(other Region) bool {
	if r.Empty() || other.Empty() {
		return false
	}
	return r.Contains(NewPoint(other.X, other.Y)) &&
		r.Contains(NewPoint(other.X+other.W-1, other.Y+other.H-1))
}

func (r Region) Union(other Region) Region {
	if r.Empty() {
		return other
	}
	if other.Empty() {
		return r
	}
	left := min(r.X, other.X)
	top := min(r.Y, other.Y)
	right := max(r.X+r.W, other.X+other.W)
	bottom := max(r.Y+r.H, other.Y+other.H)
	return NewRegion(left, top, right-left, bottom-top)
}

func (r Region) Intersection(other Region) Region {
	if r.Empty() || other.Empty() {
		return NewRegion(0, 0, 0, 0)
	}
	left := max(r.X, other.X)
	top := max(r.Y, other.Y)
	right := min(r.X+r.W, other.X+other.W)
	bottom := min(r.Y+r.H, other.Y+other.H)
	if right <= left || bottom <= top {
		return NewRegion(left, top, 0, 0)
	}
	return NewRegion(left, top, right-left, bottom-top)
}

func (r *Region) SetThrowException(flag bool) {
	r.ThrowException = flag
}

func (r *Region) ResetThrowException() {
	r.ThrowException = true
}

func (r *Region) SetAutoWaitTimeout(sec float64) {
	if sec < 0 {
		sec = 0
	}
	r.AutoWaitTimeout = sec
}

func (r *Region) SetWaitScanRate(rate float64) {
	if rate <= 0 {
		rate = DefaultWaitScanRate
	}
	r.WaitScanRate = rate
}

func (r *Region) SetObserveScanRate(rate float64) {
	if rate <= 0 {
		rate = DefaultObserveScanRate
	}
	r.ObserveScanRate = rate
}

func (r Region) Find(source *Image, pattern *Pattern) (Match, error) {
	f, err := r.newFinder(source)
	if err != nil {
		return Match{}, err
	}
	return f.Find(pattern)
}

func (r Region) Exists(source *Image, pattern *Pattern, timeout time.Duration) (Match, bool, error) {
	checkOnce := func() (Match, bool, error) {
		f, err := r.newFinder(source)
		if err != nil {
			return Match{}, false, err
		}
		m, err := f.Find(pattern)
		if err != nil {
			if errors.Is(err, ErrFindFailed) {
				return Match{}, false, nil
			}
			return Match{}, false, err
		}
		return m, true, nil
	}

	if timeout <= 0 {
		return checkOnce()
	}

	deadline := time.Now().Add(timeout)
	interval := r.waitInterval()
	for {
		m, ok, err := checkOnce()
		if err != nil {
			return Match{}, false, err
		}
		if ok {
			return m, true, nil
		}
		if time.Now().After(deadline) {
			return Match{}, false, nil
		}
		sleep := interval
		if remaining := time.Until(deadline); remaining < sleep {
			sleep = remaining
		}
		if sleep > 0 {
			time.Sleep(sleep)
		}
	}
}

func (r Region) Has(source *Image, pattern *Pattern, timeout time.Duration) (bool, error) {
	_, ok, err := r.Exists(source, pattern, timeout)
	return ok, err
}

func (r Region) Wait(source *Image, pattern *Pattern, timeout time.Duration) (Match, error) {
	effectiveTimeout := timeout
	if effectiveTimeout <= 0 {
		effectiveTimeout = time.Duration(r.AutoWaitTimeout * float64(time.Second))
	}
	m, ok, err := r.Exists(source, pattern, effectiveTimeout)
	if err != nil {
		return Match{}, err
	}
	if !ok {
		return Match{}, ErrTimeout
	}
	return m, nil
}

func (r Region) WaitVanish(source *Image, pattern *Pattern, timeout time.Duration) (bool, error) {
	checkOnce := func() (bool, error) {
		_, ok, err := r.Exists(source, pattern, 0)
		if err != nil {
			return false, err
		}
		return !ok, nil
	}

	if timeout <= 0 {
		return checkOnce()
	}

	deadline := time.Now().Add(timeout)
	interval := r.waitInterval()
	for {
		vanished, err := checkOnce()
		if err != nil {
			return false, err
		}
		if vanished {
			return true, nil
		}
		if time.Now().After(deadline) {
			return false, nil
		}
		sleep := interval
		if remaining := time.Until(deadline); remaining < sleep {
			sleep = remaining
		}
		if sleep > 0 {
			time.Sleep(sleep)
		}
	}
}

func (r Region) FindAll(source *Image, pattern *Pattern) ([]Match, error) {
	f, err := r.newFinder(source)
	if err != nil {
		return nil, err
	}
	return f.FindAll(pattern)
}

func (r Region) FindAllByRow(source *Image, pattern *Pattern) ([]Match, error) {
	f, err := r.newFinder(source)
	if err != nil {
		return nil, err
	}
	return f.FindAllByRow(pattern)
}

func (r Region) FindAllByColumn(source *Image, pattern *Pattern) ([]Match, error) {
	f, err := r.newFinder(source)
	if err != nil {
		return nil, err
	}
	return f.FindAllByColumn(pattern)
}

func (r Region) newFinder(source *Image) (*Finder, error) {
	if source == nil || source.Gray() == nil {
		return nil, fmt.Errorf("%w: source image is nil", ErrInvalidTarget)
	}
	if r.Empty() {
		return nil, fmt.Errorf("%w: region is empty", ErrInvalidTarget)
	}
	crop, err := source.Crop(r.Rect)
	if err != nil {
		return nil, err
	}
	return NewFinder(crop)
}

func (r Region) waitInterval() time.Duration {
	rate := r.WaitScanRate
	if rate <= 0 {
		rate = DefaultWaitScanRate
	}
	interval := time.Duration(float64(time.Second) / rate)
	if interval < time.Millisecond {
		return time.Millisecond
	}
	return interval
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Screen struct {
	ID     int
	Bounds Rect
}

func NewScreen(id int, bounds Rect) Screen {
	return Screen{ID: id, Bounds: bounds}
}
