# OCR Collection Surface

Phase 7 expands OCR parity beyond the original `ReadText` and `FindText` pair without changing the current sikuli-go OCR backend architecture.

## Scope

The API now exposes richer OCR traversal helpers across the image-backed and live-screen surfaces:

- `Finder.CollectWords`
- `Finder.CollectLines`
- `Region.CollectWords`
- `Region.CollectLines`
- `LiveRegion.CollectWords`
- `LiveRegion.CollectLines`
- `Screen.CollectWords`
- `Screen.CollectLines`
- `Match.CollectWords`
- `Match.CollectLines`

Existing OCR helpers remain unchanged:

- `ReadText`
- `FindText`

## Stable Result Types

The Go API now uses two stable OCR result values:

- `OCRWord`
  - word text
  - confidence
  - geometry
  - stable word index
- `OCRLine`
  - line text
  - confidence
  - geometry
  - grouped `[]OCRWord`

This keeps the OCR surface predictable for client ports and parity wrappers.

## Geometry and Binding

Image-backed helpers return geometry in the source-image coordinate space.

Live helpers capture once per call and then rebind OCR geometry back into the absolute screen or live-region coordinate space. That means:

- `LiveRegion` OCR results are offset into the live region bounds
- `Screen` OCR results are offset into the screen bounds
- live `Match` OCR results stay bound to the matched live region

## OCR Parameter Propagation

`OCRParams` now flows through the same way for image-backed and live-screen helpers:

- language
- training data path
- minimum confidence
- timeout
- case sensitivity for `FindText`

The live surface does not add a separate OCR protocol. It captures through the runtime, then runs the same finder/OCR helper logic used by image-backed search.

## Char-Level Note

Phase 7 does not add char-level OCR result types.

That is deliberate:

- the current OCR backend surface exposes stable word geometry
- line grouping can be derived deterministically from those word results
- char-level geometry would require backend-specific segmentation rules that are not yet stable across the current OCR backends/builds

So this phase closes the main SikuliX OCR workflow gap while keeping the existing OCR architecture intact.

## Result

OCR-heavy ports no longer need to reconstruct common SikuliX text traversal workflows entirely in clients just to iterate words or lines.
