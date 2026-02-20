package testharness

import (
	"fmt"
	"math"

	"github.com/sikulix/portgo/internal/core"
)

type CompareOptions struct {
	ScoreTolerance float64
}

func CompareMatches(got []core.MatchCandidate, want []ExpectedMatch, opts CompareOptions) error {
	if len(got) != len(want) {
		return fmt.Errorf("match count mismatch: got=%d want=%d", len(got), len(want))
	}
	for i := 0; i < len(want); i++ {
		g := got[i]
		w := want[i]
		if g.X != w.X || g.Y != w.Y || g.W != w.W || g.H != w.H {
			return fmt.Errorf(
				"match[%d] geometry mismatch: got=(%d,%d %dx%d) want=(%d,%d %dx%d)",
				i, g.X, g.Y, g.W, g.H, w.X, w.Y, w.W, w.H,
			)
		}
		minAllowed := w.ScoreMin - opts.ScoreTolerance
		if g.Score < minAllowed {
			return fmt.Errorf(
				"match[%d] score below minimum: got=%.6f min=%.6f",
				i, g.Score, minAllowed,
			)
		}
		if w.ScoreMax > 0 {
			maxAllowed := w.ScoreMax + opts.ScoreTolerance
			if g.Score > maxAllowed {
				return fmt.Errorf(
					"match[%d] score above maximum: got=%.6f max=%.6f",
					i, g.Score, maxAllowed,
				)
			}
		}
	}
	return nil
}

func AlmostEqual(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}

