# API: `internal/core`

[Back to API Index](./)

## Full Package Doc

```text
package core // import "github.com/sikulix/portgo/internal/core"


FUNCTIONS

func ResizeGrayNearest(src *image.Gray, factor float64) *image.Gray

TYPES

type MatchCandidate struct {
	X     int
	Y     int
	W     int
	H     int
	Score float64
}

type Matcher interface {
	Find(req SearchRequest) ([]MatchCandidate, error)
}

type SearchRequest struct {
	Haystack     *image.Gray
	Needle       *image.Gray
	Mask         *image.Gray
	Threshold    float64
	ResizeFactor float64
	MaxResults   int
}

func (r SearchRequest) Validate() error

```
