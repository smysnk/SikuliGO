# API: `internal/testharness`

[Back to API Index](./)

## Full Package Doc

```text
package testharness // import "github.com/sikulix/portgo/internal/testharness"


FUNCTIONS

func AlmostEqual(a, b, tol float64) bool
func CompareMatches(got []core.MatchCandidate, want []ExpectedMatch, opts CompareOptions) error
func MatrixToGray(rows [][]int) (*image.Gray, error)

TYPES

type CompareOptions struct {
	ScoreTolerance float64
}

type ExpectedMatch struct {
	X        int     `json:"x"`
	Y        int     `json:"y"`
	W        int     `json:"w"`
	H        int     `json:"h"`
	ScoreMin float64 `json:"score_min"`
	ScoreMax float64 `json:"score_max"`
}

type GoldenCase struct {
	Name         string          `json:"name"`
	Haystack     [][]int         `json:"haystack"`
	Needle       [][]int         `json:"needle"`
	Mask         [][]int         `json:"mask,omitempty"`
	Threshold    float64         `json:"threshold"`
	ResizeFactor float64         `json:"resize_factor"`
	MaxResults   int             `json:"max_results"`
	Expected     []ExpectedMatch `json:"expected"`
}

func LoadCorpus() ([]GoldenCase, error)

```
