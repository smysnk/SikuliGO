# API: `internal/testharness`

[Back to API Index](./)

<style>
  .api-type { color: #0f766e; font-weight: 700; }
  .api-func { color: #1d4ed8; font-weight: 700; }
  .api-method { color: #7c3aed; font-weight: 700; }
  .api-signature { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace; }
</style>

Legend: <span class="api-type">Type</span>, <span class="api-func">Function</span>, <span class="api-method">Method</span>

Package: `package testharness // import "github.com/sikulix/portgo/internal/testharness"`

## Symbol Index

### Types

- <span class="api-type">[`CompareOptions`](#type-compareoptions)</span>
- <span class="api-type">[`ExpectedMatch`](#type-expectedmatch)</span>
- <span class="api-type">[`GoldenCase`](#type-goldencase)</span>

### Functions

- <span class="api-func">[`AlmostEqual`](#func-almostequal)</span>
- <span class="api-func">[`CompareMatches`](#func-comparematches)</span>
- <span class="api-func">[`MatrixToGray`](#func-matrixtogray)</span>
- <span class="api-func">[`LoadCorpus`](#func-loadcorpus)</span>

### Methods

- none

## Declarations

### Types

#### <a id="type-compareoptions"></a><span class="api-type">Type</span> `CompareOptions`

- Signature: <span class="api-signature">`type CompareOptions struct {`</span>

#### <a id="type-expectedmatch"></a><span class="api-type">Type</span> `ExpectedMatch`

- Signature: <span class="api-signature">`type ExpectedMatch struct {`</span>

#### <a id="type-goldencase"></a><span class="api-type">Type</span> `GoldenCase`

- Signature: <span class="api-signature">`type GoldenCase struct {`</span>

### Functions

#### <a id="func-almostequal"></a><span class="api-func">Function</span> `AlmostEqual`

- Signature: <span class="api-signature">`func AlmostEqual(a, b, tol float64) bool`</span>

#### <a id="func-comparematches"></a><span class="api-func">Function</span> `CompareMatches`

- Signature: <span class="api-signature">`func CompareMatches(got []core.MatchCandidate, want []ExpectedMatch, opts CompareOptions) error`</span>
- Uses: [`CompareOptions`](#type-compareoptions), [`ExpectedMatch`](#type-expectedmatch)

#### <a id="func-matrixtogray"></a><span class="api-func">Function</span> `MatrixToGray`

- Signature: <span class="api-signature">`func MatrixToGray(rows [][]int) (*image.Gray, error)`</span>

#### <a id="func-loadcorpus"></a><span class="api-func">Function</span> `LoadCorpus`

- Signature: <span class="api-signature">`func LoadCorpus() ([]GoldenCase, error)`</span>
- Uses: [`GoldenCase`](#type-goldencase)

## Raw Package Doc

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
