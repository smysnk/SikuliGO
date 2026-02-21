# Architecture Lock

This document defines the locked package boundaries, object responsibilities, and internal protocols for the current GoLang port baseline.

## Package boundaries

- `pkg/sikuli`:
  - Public types, defaults, and errors.
  - Public API orchestration (`Finder` delegates matching work to `core.Matcher` and `core.OCR`).
  - Input automation controller (`InputController`) delegates to `core.Input`.
  - Observe/event controller (`ObserverController`) delegates to `core.Observer`.
  - App/window controller (`AppController`) delegates to `core.App`.
  - Region/Finder helper semantics (`Region` geometry/runtime setters, `Finder.Exists/Has`).
  - Region-scoped search protocol (`Region.Find/Exists/Has/Wait`) over source image crops.
  - OCR and text-search protocol (`Finder.ReadText/FindText`, `Region.ReadText/FindText`).
  - Location/offset value objects for parity-friendly coordinate APIs.
  - Options map protocol (`Options`) for typed config compatibility helpers.
- `internal/core`:
  - Matching protocol contracts and transport objects.
  - OCR protocol contracts and transport objects.
  - Input automation protocol contracts and transport objects.
  - Shared image operations used by backends.
- `internal/cv`:
  - Primary matcher backend implementation (`NCCMatcher`).
- `internal/ocr`:
  - OCR backend implementation with optional `gogosseract` integration.
- `internal/input`:
  - Input backend implementation with unsupported default.
- `internal/observe`:
  - Observe backend implementation with unsupported default.
- `internal/app`:
  - App/window backend implementation with unsupported default.
- `internal/testharness`:
  - Corpus loading, comparator policy, and parity tests.

## Object inventory by package

### `pkg/sikuli`

- Value objects:
  - `Point`, `Location`, `Offset`, `Rect`, `Region`, `Screen`, `Match`, `TextMatch`
- OCR request helper objects:
  - `OCRParams`
- Input helper objects:
  - `InputOptions`, `MouseButton`
- Observe helper objects:
  - `ObserveOptions`, `ObserveEventType`, `ObserveEvent`
- App helper objects:
  - `AppOptions`, `Window`
- Stateful objects:
  - `Image`, `Pattern`, `Finder`
- Input automation objects:
  - `InputController`
- Observe automation objects:
  - `ObserverController`
- App/window automation objects:
  - `AppController`
- Global configuration:
  - `RuntimeSettings`, settings mutator/accessor functions
- Options/configuration helpers:
  - `Options` typed map wrapper
- Compatibility interfaces:
  - `ImageAPI`, `PatternAPI`, `FinderAPI`, `RegionAPI`, `InputAPI`, `ObserveAPI`, `AppAPI`
- Error protocol:
  - `ErrFindFailed`, `ErrTimeout`, `ErrInvalidTarget`, `ErrBackendUnsupported`

### `internal/core`

- `MatchCandidate`: backend-neutral match payload.
- `SearchRequest`: backend-neutral request payload.
- `Matcher`: backend protocol interface.
- `OCRRequest`: backend-neutral OCR request payload.
- `OCRWord`: backend-neutral OCR word payload.
- `OCRResult`: backend-neutral OCR response payload.
- `OCR`: backend OCR protocol interface.
- `ErrOCRUnsupported`: backend capability sentinel.
- `InputAction`: backend-neutral input action enum.
- `InputRequest`: backend-neutral input request payload.
- `Input`: backend input protocol interface.
- `ErrInputUnsupported`: backend capability sentinel.
- `ObserveEventType`: backend-neutral observe event enum.
- `ObserveRequest`: backend-neutral observe request payload.
- `ObserveEvent`: backend-neutral observe event payload.
- `Observer`: backend observe protocol interface.
- `ErrObserveUnsupported`: backend capability sentinel.
- `AppAction`: backend-neutral app action enum.
- `AppRequest`: backend-neutral app request payload.
- `AppResult`: backend-neutral app response payload.
- `WindowInfo`: backend-neutral window payload.
- `App`: backend app protocol interface.
- `ErrAppUnsupported`: backend capability sentinel.
- `ResizeGrayNearest`: canonical nearest-neighbor resize helper.

### `internal/cv`

- `NCCMatcher`: concrete implementation of `core.Matcher`.
- `SADMatcher`: alternate implementation of `core.Matcher`.
- Internal protocol helpers:
  - normalized cross-correlation scoring
  - mask inclusion policy
  - grayscale pixel accessor policy

### `internal/ocr`

- `unsupportedBackend`: default implementation returning `core.ErrOCRUnsupported`.
- `gogosseractBackend`: optional implementation (build tag `gogosseract`) using `github.com/smysnk/gogosseract`.
- Internal protocol helpers:
  - reflective method adaptation for fork compatibility
  - hOCR word parsing and confidence filtering

### `internal/input`

- `unsupportedBackend`: default implementation returning `core.ErrInputUnsupported`.

### `internal/observe`

- `unsupportedBackend`: default implementation returning `core.ErrObserveUnsupported`.

### `internal/app`

- `unsupportedBackend`: default implementation returning `core.ErrAppUnsupported`.

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

## Protocol lock: OCR boundary

The OCR backend boundary remains strictly behind this interface:

```go
type OCR interface {
  Read(req OCRRequest) (OCRResult, error)
}
```

`pkg/sikuli.Finder` text APIs (`ReadText` and `FindText`) must consume only this protocol and must not depend on backend-specific types.

Default builds use the unsupported backend and return `ErrBackendUnsupported` through the public API unless built with `-tags gogosseract`.

## Protocol lock: input boundary

The input backend boundary remains strictly behind this interface:

```go
type Input interface {
  Execute(req InputRequest) error
}
```

`pkg/sikuli.InputController` must consume only this protocol and must not depend on backend-specific types.

## Protocol lock: observe boundary

The observe backend boundary remains strictly behind this interface:

```go
type Observer interface {
  Observe(req ObserveRequest) ([]ObserveEvent, error)
}
```

`pkg/sikuli.ObserverController` must consume only this protocol and must not depend on backend-specific types.

## Protocol lock: app boundary

The app backend boundary remains strictly behind this interface:

```go
type App interface {
  Execute(req AppRequest) (AppResult, error)
}
```

`pkg/sikuli.AppController` must consume only this protocol and must not depend on backend-specific types.

## Protocol lock: region-scoped search

`Region` search methods are layered on top of the same matcher protocol and are locked to this behavior:

1. crop the given source image to the region rect via `Image.Crop`
2. create a `Finder` over the cropped image
3. run `Find`/`Exists`/`Has`/`Wait` semantics against that scope

Timeout semantics for `Region.Exists/Wait` are polling-based and driven by:

- explicit timeout argument when `> 0`
- `Region.AutoWaitTimeout` when `Wait` receives `<= 0`
- `Region.WaitScanRate` for polling interval

`Finder.Wait/WaitVanish` are polling-based and driven by global wait scan rate from `RuntimeSettings`.

`Region.ReadText/FindText` remain layered on the same region crop protocol (`Image.Crop` then `Finder` execution), preserving region scoping for OCR requests.

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
