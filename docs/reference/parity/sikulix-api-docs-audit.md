# SikuliX API Docs Audit

This audit compares the documented behavior in the official SikuliX API docs with the current Go port implementation in `packages/api/pkg/sikuli` and `packages/api/internal/grpcv1`.

Scope:
- Official docs reviewed: [API index](https://sikulix.github.io/docs/api/), [Region](https://sikulix.github.io/docs/api/region/), [Screen](https://sikulix.github.io/docs/api/screen/), [Finder](https://sikulix.github.io/docs/api/finder/), [Pattern](https://sikulix.github.io/docs/api/pattern/), [Match](https://sikulix.github.io/docs/api/match/), [App](https://sikulix.github.io/docs/api/appclass/), [Text and OCR](https://sikulix.github.io/docs/api/textandocr/), [Keys](https://sikulix.github.io/docs/api/keys/).
- Go surfaces reviewed: `packages/api/pkg/sikuli/*.go`, `packages/api/proto/sikuli/v1/sikuli.proto`, `packages/api/internal/grpcv1/server.go`.
- This document compares SikuliX docs to the Go port itself. Node and Python client wrappers intentionally are not treated as the source of truth when they add convenience methods on top of the API.

## Executive Summary

The largest parity gap is architectural: SikuliX documents `Screen`, `Region`, and `Match` as live, action-capable desktop objects. The current Go port exposes image-scoped search primitives in `pkg/sikuli` and moves live desktop work into a separate gRPC server plus controller types.

That means the Go port currently differs from the SikuliX docs in four material ways:

1. The public Go `Region` API is image-driven, not live-screen-driven.
2. The public Go `Screen` type is only a descriptor, not a live monitor/search/input surface.
3. `Match` is a value object, not a `Region`-like object with direct action methods.
4. Several documented SikuliX convenience APIs are either absent or split across separate controllers/RPCs.

## Findings

| Area | SikuliX Docs Describe | Current Go Port | Impact |
|---|---|---|---|
| `Region` runtime model | `Region` is a live search and action surface on the desktop | `Region` methods require a source `*Image`; live-screen behavior exists only behind gRPC screen RPCs | High |
| `Screen` surface | `Screen` extends region behavior, represents monitors, supports capture and monitor selection | `Screen` is only `{ID, Bounds}` in `pkg/sikuli`; no public capture/search/action methods | High |
| `Match` semantics | `Match` behaves like a `Region` and can be clicked/hovered/used directly | `Match` is a plain struct (`Rect`, `Score`, `Target`, `Index`) | High |
| Action methods | Direct `click`, `doubleClick`, `rightClick`, `hover`, `dragDrop`, `type`, `paste`, wheel/mouse/key state operations | Input is split into `InputController` plus gRPC RPCs; many direct convenience methods are absent | High |
| Iteration model | `Finder.findAll()` uses `hasNext()/next()` and `destroy()` lifecycle | Go returns `[]Match`; no iterator state or `destroy()` | Medium |
| Exception/null semantics | `find()` throws `FindFailed`; `exists()` returns `null`; `setThrowException` changes behavior | Go uses `(Match, bool, error)` and sentinel errors; `ThrowException` / `FindFailedThrows` exist but are not wired into search behavior | High |
| Multi-target helpers | `findAnyList`, `findBestList`, `getAll`, related helper families are documented | No equivalent helpers in `pkg/sikuli` or gRPC surface | Medium |
| OCR helper surface | Docs describe `text()`, `findText()`, `collectLines()`, `collectWords()` and related text workflows | Go exposes `ReadText` and `FindText` only | Medium |
| App/window model | Docs describe richer `App` / `Window` workflows (`window()`, `focusedWindow()`, `allWindows()`) | Go exposes `Open`, `Focus`, `Close`, `IsRunning`, `ListWindows` only | Medium |

## Detailed Differences

### 1. `Region` is not a live desktop object in the public Go API

SikuliX documentation presents `Region` as the primary live desktop abstraction: search, wait, click, hover, type, paste, drag/drop, OCR, and observe all hang directly off the region or screen.

Current Go behavior:
- `Region.Find`, `Region.Exists`, `Region.Wait`, `Region.FindAll`, `Region.ReadText`, and `Region.FindText` all require a source `*Image`.
- `Region` has no direct desktop input methods.
- Live-screen searching is implemented by gRPC methods such as `FindOnScreen`, `ExistsOnScreen`, `WaitOnScreen`, and `ClickOnScreen` in `packages/api/internal/grpcv1/server.go`.

Why this matters:
- Existing SikuliX users reading the official docs expect `Region` to be the live desktop handle.
- In the Go port, `Region` is closer to an image ROI helper, while live desktop interaction is a server concern.

Relevant implementation:
- `packages/api/pkg/sikuli/types.go`
- `packages/api/internal/grpcv1/server.go`
- `packages/api/proto/sikuli/v1/sikuli.proto`

### 2. `Screen` is only a descriptor, not a live action surface

SikuliX documentation describes `Screen` as a specialized region representing monitors, with monitor enumeration, monitor IDs, primary screen semantics, and capture helpers.

Current Go behavior:
- `pkg/sikuli` defines:
  - `type Screen struct { ID int; Bounds Rect }`
  - `func NewScreen(id int, bounds Rect) Screen`
- There are no public `Screen` methods for capture, find, click, monitor enumeration, or selection.
- Screen capture exists only internally on the server side via `capture(...)` and platform capture helpers.

Why this matters:
- The public Go package does not currently provide a SikuliX-like `Screen` API.
- Multi-monitor behavior exists operationally in benchmarks and platform capture internals, but not as a first-class public contract.

Relevant implementation:
- `packages/api/pkg/sikuli/types.go`
- `packages/api/internal/grpcv1/server.go`

### 3. `Match` is not `Region`-like

SikuliX documentation treats `Match` as a region-like result object that can be acted on directly.

Current Go behavior:
- `Match` is a plain value type with geometry and score only.
- It has no methods for click, hover, type, drag/drop, or further search.

Why this matters:
- SikuliX flows often chain directly from a match to an action.
- In the Go port, callers must extract the match target and route the action through `InputController` or the corresponding RPC.

Relevant implementation:
- `packages/api/pkg/sikuli/match.go`

### 4. Direct action parity is incomplete

The Region docs describe direct action methods such as:
- `click`
- `doubleClick`
- `rightClick`
- `hover`
- `dragDrop`
- `type`
- `paste`
- `wheel`
- `mouseDown` / `mouseUp`
- `keyDown` / `keyUp`

Current Go behavior:
- Public input surface is limited to:
  - `MoveMouse`
  - `Click`
  - `TypeText`
  - `Hotkey`
- There is no public Go equivalent for `doubleClick`, `rightClick` as a convenience verb, `hover` as a verb, `dragDrop`, `paste`, wheel scrolling, or stateful key/mouse down/up operations.
- `ClickOnScreen` exists server-side, but only for image-find then click.

Why this matters:
- SikuliX scripts often rely on these direct ergonomics.
- The current Go port requires lower-level composition or client-wrapper sugar for behavior that SikuliX documents as a first-class API.

Relevant implementation:
- `packages/api/pkg/sikuli/input.go`
- `packages/api/internal/grpcv1/server.go`
- `packages/api/proto/sikuli/v1/sikuli.proto`

### 5. Finder iteration and lifecycle differ

SikuliX Finder docs describe iterator-style behavior with `findAll()`, `hasNext()`, `next()`, and `destroy()`.

Current Go behavior:
- `FindAll` returns a full `[]Match`.
- `FindAllByRow` and `FindAllByColumn` return reordered slices.
- `LastMatches()` exposes the last slice copy.
- There is no iterator API and no `destroy()` lifecycle method.

Why this matters:
- Script logic written around iterator behavior does not translate directly.
- Lifetime and resource-management semantics differ from what the docs describe.

Relevant implementation:
- `packages/api/pkg/sikuli/finder.go`

### 6. Miss and timeout semantics differ from the documented exception model

SikuliX docs describe:
- `find()` throws `FindFailed` on miss.
- `exists()` returns `null` on miss.
- `wait()` throws on timeout.
- `waitVanish()` returns a boolean.
- `setThrowException(false)` changes miss behavior.

Current Go behavior:
- `Find()` returns `ErrFindFailed`.
- `Exists()` returns `(Match{}, false, nil)` on miss.
- `Wait()` returns `ErrTimeout`.
- `WaitVanish()` returns `(bool, error)`.
- `Region.ThrowException` and global `FindFailedThrows` settings exist but are not consulted by the search methods.

Why this matters:
- The Go API is explicit and idiomatic, but not behaviorally identical to the documented SikuliX exception/null model.
- The presence of `ThrowException`-style fields suggests parity, but the implementation does not currently honor them.

Relevant implementation:
- `packages/api/pkg/sikuli/errors.go`
- `packages/api/pkg/sikuli/types.go`
- `packages/api/pkg/sikuli/settings.go`

### 7. Multi-target search helpers documented in SikuliX are absent

The Region docs describe helper families such as `findAnyList`, `findBestList`, and `getAll`-style multi-target workflows.

Current Go behavior:
- Public search surface is limited to single-pattern `Find`, `FindAll`, `Exists`, `Wait`, and ordering helpers.
- No equivalent multi-pattern search helpers exist in `pkg/sikuli` or the gRPC surface.

Why this matters:
- Porting higher-level search flows from SikuliX requires custom orchestration.

Relevant implementation:
- `packages/api/pkg/sikuli/finder.go`
- `packages/api/pkg/sikuli/types.go`
- `packages/api/proto/sikuli/v1/sikuli.proto`

### 8. OCR helper surface is narrower than the docs

The Text and OCR docs describe a broader helper vocabulary, including plain text extraction, search, and collection helpers such as `collectLines()` and `collectWords()`.

Current Go behavior:
- OCR surface is limited to:
  - `ReadText(params)`
  - `FindText(query, params)`
- No collection helpers for words/lines are exposed.

Why this matters:
- OCR-heavy SikuliX scripts expect richer result traversal than the current Go API exposes.

Relevant implementation:
- `packages/api/pkg/sikuli/ocr.go`
- `packages/api/pkg/sikuli/finder.go`
- `packages/api/internal/grpcv1/server.go`

### 9. App and window APIs are narrower than the documented SikuliX model

The App docs describe a richer app/window model around app instances, focused windows, and window-specific selection.

Current Go behavior:
- Public API supports:
  - `Open`
  - `Focus`
  - `Close`
  - `IsRunning`
  - `ListWindows`
- `Window` only contains `Title`, `Bounds`, and `Focused`.
- There is no richer window object API, no `focusedWindow()` equivalent, and no documented per-window control layer.

Why this matters:
- Users porting window-oriented SikuliX code will still need an adapter layer.

Relevant implementation:
- `packages/api/pkg/sikuli/app.go`
- `packages/api/proto/sikuli/v1/sikuli.proto`

## Areas Where the Go Port Exposes Different or Additional Behavior

These are not gaps in the strict sense, but they are still behavior differences relative to the SikuliX docs:

- The gRPC surface exposes explicit matcher-engine selection (`template`, `orb`, `akaze`, `brisk`, `kaze`, `sift`, `hybrid`), which is more explicit than the SikuliX docs.
- Pattern masking is exposed directly in the Go API (`WithMask`, `WithMaskMatrix`).
- The runtime is server-based, with transport/session/auth/logging concerns that do not exist in the same way in the SikuliX docs.

Those differences are already covered at a higher level in `behavioral-differences.md`; this audit is focused on the public API behavior mismatches.

## Recommended Follow-up Work

1. Decide whether `pkg/sikuli` is intended to be a true SikuliX-shaped API or only a compatibility-oriented core for server/client wrappers.
2. If true public parity is desired, prioritize:
   - direct `Screen` and `Region` live-action methods
   - `Match` as an action-capable region-like value
   - iterator or equivalent traversal semantics for Finder
   - wiring `ThrowException` / `FindFailedThrows` into actual search behavior
3. Keep `java-to-go-mapping.md` and this audit aligned. The generated mapping currently overstates parity in a few places where the core Go surface still differs materially from the docs.
