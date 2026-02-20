package sikuli

import "fmt"

type Location struct {
	X int
	Y int
}

func NewLocation(x, y int) Location {
	return Location{X: x, Y: y}
}

func (l Location) ToPoint() Point {
	return NewPoint(l.X, l.Y)
}

func (l Location) Move(dx, dy int) Location {
	return NewLocation(l.X+dx, l.Y+dy)
}

func (l Location) String() string {
	return fmt.Sprintf("L[%d,%d]", l.X, l.Y)
}

type Offset struct {
	X int
	Y int
}

func NewOffset(x, y int) Offset {
	return Offset{X: x, Y: y}
}

func (o Offset) ToPoint() Point {
	return NewPoint(o.X, o.Y)
}

func (o Offset) String() string {
	return fmt.Sprintf("O[%d,%d]", o.X, o.Y)
}
