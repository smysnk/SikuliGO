package sikuli

import (
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

func (p Point) TargetPoint() Point {
	return p
}

// ToLocation converts a point to a parity-friendly Location value.
func (p Point) ToLocation() Location {
	return NewLocation(p.X, p.Y)
}

// ToOffset converts a point to a parity-friendly Offset value.
func (p Point) ToOffset() Offset {
	return NewOffset(p.X, p.Y)
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
	// ThrowException is retained as parity metadata for SikuliX-style ports.
	// The Go API uses explicit return values for misses and timeouts regardless of this flag.
	ThrowException  bool
	AutoWaitTimeout float64
	WaitScanRate    float64
	ObserveScanRate float64
}

// NewRegion constructs a rectangular search area with default timing settings.
func NewRegion(x, y, w, h int) Region {
	return Region{
		Rect:            NewRect(x, y, w, h),
		ThrowException:  true,
		AutoWaitTimeout: DefaultAutoWaitTimeout,
		WaitScanRate:    DefaultWaitScanRate,
		ObserveScanRate: DefaultObserveScanRate,
	}
}

// Center returns the midpoint of the region.
func (r Region) Center() Point {
	return Point{
		X: r.X + r.W/2,
		Y: r.Y + r.H/2,
	}
}

func (r Region) TargetPoint() Point {
	return r.Center()
}

// Grow expands the region outward in both directions.
func (r Region) Grow(dx, dy int) Region {
	return NewRegion(r.X-dx, r.Y-dy, r.W+dx*2, r.H+dy*2)
}

// Offset translates the region by dx and dy.
func (r Region) Offset(dx, dy int) Region {
	return NewRegion(r.X+dx, r.Y+dy, r.W, r.H)
}

// OffsetBy applies an Offset alias to this region position.
func (r Region) OffsetBy(off Offset) Region {
	return r.Offset(off.X, off.Y)
}

func (r Region) MoveTo(x, y int) Region {
	return NewRegion(x, y, r.W, r.H)
}

// MoveToLocation moves this region using a Location alias.
func (r Region) MoveToLocation(loc Location) Region {
	return r.MoveTo(loc.X, loc.Y)
}

// SetSize updates width and height while clamping negatives to zero.
func (r Region) SetSize(w, h int) Region {
	if w < 0 {
		w = 0
	}
	if h < 0 {
		h = 0
	}
	return NewRegion(r.X, r.Y, w, h)
}

// Contains reports whether a point lies within the region.
func (r Region) Contains(p Point) bool {
	return r.Rect.Contains(p)
}

// ContainsLocation reports whether this region contains the given location.
func (r Region) ContainsLocation(loc Location) bool {
	return r.Contains(loc.ToPoint())
}

func (r Region) ContainsRegion(other Region) bool {
	if r.Empty() || other.Empty() {
		return false
	}
	return r.Contains(NewPoint(other.X, other.Y)) &&
		r.Contains(NewPoint(other.X+other.W-1, other.Y+other.H-1))
}

// Union returns the smallest region containing both regions.
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

// Intersection returns the overlap between this region and another.
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

// Exists checks for pattern presence within timeout and returns first match when found.
func (r Region) Exists(source *Image, pattern *Pattern, timeout time.Duration) (Match, bool, error) {
	return SearchExists(func() (Match, error) {
		f, err := r.newFinder(source)
		if err != nil {
			return Match{}, err
		}
		return f.Find(pattern)
	}, timeout, r.waitInterval())
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
	return SearchWait(func() (Match, error) {
		f, err := r.newFinder(source)
		if err != nil {
			return Match{}, err
		}
		return f.Find(pattern)
	}, effectiveTimeout, r.waitInterval())
}

// WaitVanish waits until pattern disappears or timeout expires.
func (r Region) WaitVanish(source *Image, pattern *Pattern, timeout time.Duration) (bool, error) {
	return SearchWaitVanish(func() (Match, error) {
		f, err := r.newFinder(source)
		if err != nil {
			return Match{}, err
		}
		return f.Find(pattern)
	}, timeout, r.waitInterval())
}

// FindAll returns all matches in this region.
func (r Region) FindAll(source *Image, pattern *Pattern) ([]Match, error) {
	f, err := r.newFinder(source)
	if err != nil {
		return nil, err
	}
	return f.FindAll(pattern)
}

// Count returns the number of matches found in this region.
func (r Region) Count(source *Image, pattern *Pattern) (int, error) {
	matches, err := r.FindAll(source, pattern)
	if err != nil {
		return 0, err
	}
	return len(matches), nil
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

func (r Region) ReadText(source *Image, params OCRParams) (string, error) {
	f, err := r.newFinder(source)
	if err != nil {
		return "", err
	}
	return f.ReadText(params)
}

// CollectWords runs OCR in region and returns word-level results.
func (r Region) CollectWords(source *Image, params OCRParams) ([]OCRWord, error) {
	f, err := r.newFinder(source)
	if err != nil {
		return nil, err
	}
	return f.CollectWords(params)
}

// CollectLines runs OCR in region and returns line-level results.
func (r Region) CollectLines(source *Image, params OCRParams) ([]OCRLine, error) {
	f, err := r.newFinder(source)
	if err != nil {
		return nil, err
	}
	return f.CollectLines(params)
}

// FindText runs OCR in region and returns matches for the query.
func (r Region) FindText(source *Image, query string, params OCRParams) ([]TextMatch, error) {
	f, err := r.newFinder(source)
	if err != nil {
		return nil, err
	}
	return f.FindText(query, params)
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
	ID      int
	Name    string
	Bounds  Rect
	Primary bool
	runtime *Runtime
}

// NewScreen constructs a logical screen descriptor.
func NewScreen(id int, bounds Rect) Screen {
	return Screen{ID: id, Bounds: bounds}
}
