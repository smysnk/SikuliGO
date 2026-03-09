# Live Screen Surface

Phase 2 introduces a public live-screen API in `packages/api/pkg/sikuli` without changing the existing API-runtime architecture.

## Public Entry Point

Use `NewRuntime` to connect to a running `sikuli-go` API process:

```go
runtime, err := sikuli.NewRuntime("127.0.0.1:50051")
```

The runtime exposes:

- `Screens()`
- `PrimaryScreen()`
- `Screen(id)`
- `Capture()`
- `CaptureRegion(region)`
- `Region(region)`

## `Screen` Model

`Screen` is now a usable live surface instead of a metadata-only descriptor.

It exposes:

- `FullRegion()`
- `Region(x, y, w, h)`
- `RegionRect(rect)`
- `Capture()`
- `Find(pattern)`
- `Exists(pattern, timeout)`
- `Has(pattern, timeout)`
- `Wait(pattern, timeout)`
- `WaitVanish(pattern, timeout)`

`Screen.Region(...)` coordinates are screen-local. The runtime translates them to the correct absolute desktop region before searching or capturing.

## `LiveRegion` Model

`LiveRegion` is the screen-backed region facade used for live screen operations.

It exposes:

- `Capture()`
- `Find(pattern)`
- `Exists(pattern, timeout)`
- `Has(pattern, timeout)`
- `Wait(pattern, timeout)`
- `WaitVanish(pattern, timeout)`
- `WithMatcherEngine(engine)`

The existing image-backed `Region` API remains intact and is still the correct surface for image-to-image workflows.

## Capture Scope Rules

- `Runtime.Capture()` captures the whole desktop image exposed by the runtime backend.
- `Screen.Capture()` captures only the selected screen.
- `Screen.Region(...).Capture()` captures only the selected local region on that screen.
- `Runtime.CaptureRegion(region)` captures a global desktop region.

## Multi-Monitor Selection

The runtime protocol now supports:

- enumerating available screens
- resolving the primary screen
- targeting a specific screen by `screen_id`
- capturing or searching within that screen

For screen-targeted operations, region coordinates are interpreted relative to the selected screen, not the full desktop.

## Semantics

Phase 1 search semantics remain in force:

- live `Find` miss -> `ErrFindFailed`
- live `Exists` miss -> `Match{}, false, nil`
- live `Wait` timeout -> `ErrTimeout`
- live `WaitVanish` timeout -> `false, nil`

The live surface reuses the same parity contract rather than inventing a separate runtime-specific search model.
