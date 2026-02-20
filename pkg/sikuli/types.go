package sikuli

import "fmt"

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

type Screen struct {
	ID     int
	Bounds Rect
}

func NewScreen(id int, bounds Rect) Screen {
	return Screen{ID: id, Bounds: bounds}
}

