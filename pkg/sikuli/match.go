package sikuli

import "fmt"

type Match struct {
	Rect
	Score  float64
	Target Point
	Index  int
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

