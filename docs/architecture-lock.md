# Architecture Lock

This document defines the locked package boundaries, object responsibilities, and internal protocols for the current GoLang port baseline.

## Package boundaries

- `pkg/sikuli`:
  - Public types, defaults, and errors.
  - Public API orchestration (`Finder` delegates matching work to `core.Matcher`).
  - Region/Finder helper semantics (`Region` geometry/runtime setters, `Finder.Exists/Has`).
  - Region-scoped search protocol (`Region.Find/Exists/Has/Wait`) over source image crops.
  - Location/offset value objects for parity-friendly coordinate APIs.
- `internal/core`:
  - Matching protocol contracts and transport objects.
  - Shared image operations used by backends.
- `internal/cv`:
  - Primary matcher backend implementation (`NCCMatcher`).
- `internal/testharness`:
  - Corpus loading, comparator policy, and parity tests.

## Object inventory by package

### `pkg/sikuli`

- Value objects:
  - `Point`, `Location`, `Offset`, `Rect`, `Region`, `Screen`, `Match`
- Stateful objects:
  - `Image`, `Pattern`, `Finder`
- Global configuration:
  - `RuntimeSettings`, settings mutator/accessor functions
- Compatibility interfaces:
  - `ImageAPI`, `PatternAPI`, `FinderAPI`
- Error protocol:
  - `ErrFindFailed`, `ErrTimeout`, `ErrInvalidTarget`, `ErrBackendUnsupported`

### `internal/core`

- `MatchCandidate`: backend-neutral match payload.
- `SearchRequest`: backend-neutral request payload.
- `Matcher`: backend protocol interface.
- `ResizeGrayNearest`: canonical nearest-neighbor resize helper.

### `internal/cv`

- `NCCMatcher`: concrete implementation of `core.Matcher`.
- `SADMatcher`: alternate implementation of `core.Matcher`.
- Internal protocol helpers:
  - normalized cross-correlation scoring
  - mask inclusion policy
  - grayscale pixel accessor policy

### `internal/testharness`

- Data protocol types:
  - `GoldenCase`
  - `ExpectedMatch`
  - `CompareOptions`
- Harness protocol functions:
  - `LoadCorpus`
  - `MatrixToGray`
  - `CompareMatches`
  - `AlmostEqual`

## Protocol lock: matcher boundary

The backend boundary remains strictly behind this interface:

```go
type Matcher interface {
  Find(req SearchRequest) ([]MatchCandidate, error)
}
```

`pkg/sikuli.Finder` must consume only this protocol and must not depend on backend-specific types.

## Protocol lock: region-scoped search

`Region` search methods are layered on top of the same matcher protocol and are locked to this behavior:

1. crop the given source image to the region rect via `Image.Crop`
2. create a `Finder` over the cropped image
3. run `Find`/`Exists`/`Has`/`Wait` semantics against that scope

Timeout semantics for `Region.Exists/Wait` are polling-based and driven by:

- explicit timeout argument when `> 0`
- `Region.AutoWaitTimeout` when `Wait` receives `<= 0`
- `Region.WaitScanRate` for polling interval

## Protocol lock: request and validation

`SearchRequest` fields and behavior are currently locked:

- Required:
  - `Haystack`
  - `Needle`
- Optional:
  - `Mask`
  - `MaxResults`
- Behavioral controls:
  - `Threshold` in `[0,1]`
  - `ResizeFactor > 0`

Validation failures are returned as `error`.

## Protocol lock: result ordering

Current sort policy in `internal/cv` is locked for deterministic behavior:

1. score descending
2. y ascending
3. x ascending

`pkg/sikuli` also exposes row-first and column-first secondary sort helpers for post-processing:

- `SortMatchesByRowColumn`
- `SortMatchesByColumnRow`

## Protocol lock: masking and resize

- Mask dimensions must match the effective needle dimensions.
- Zero mask values exclude pixels from scoring.
- Non-zero mask values include pixels.
- Needle and mask are both resized using nearest-neighbor when `ResizeFactor != 1.0`.

## Protocol lock: parity harness

`internal/testharness` defines the active parity protocol:

- JSON corpus fixture format (`golden_match_cases.json`)
- matrix-to-gray conversion
- geometry + score-window assertions through `CompareMatches`

This is the required baseline for backend refactors or backend replacements.

Backend protocol conformance is additionally validated by matcher conformance tests in `internal/cv/conformance_test.go`.
