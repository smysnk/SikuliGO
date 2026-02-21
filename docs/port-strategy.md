# GoLang Port Strategy

This document consolidates the project goals, architecture overview, implementation plan, and feature matrix for the SikuliGO port.

## Goals

- Build a feature-complete GoLang port of the core Sikuli API concepts.
- Preserve behavioral parity for matching, region semantics, and finder flows.
- Keep a stable public API while allowing backend upgrades.
- Make behavior measurable with deterministic tests and parity fixtures.

## Architecture Overview

### Module and package layout

- `go.mod`: root GoLang module
- `pkg/sikuli`: public API surface and compatibility-facing types
- `internal/core`: shared contracts and primitives (`SearchRequest`, `Matcher`, resize helpers)
- `internal/cv`: concrete matching engine implementation
- `internal/ocr`: OCR backend adapters and hOCR parsing helpers
- `internal/input`: input automation backend adapters
- `internal/observe`: observe/event backend adapters
- `internal/app`: app/window backend adapters
- `internal/testharness`: golden corpus loading and parity comparators

### Backend boundaries

The matcher boundary is fixed behind `core.Matcher`:

```go
type Matcher interface {
  Find(req SearchRequest) ([]MatchCandidate, error)
}
```

This keeps `pkg/sikuli` stable while allowing alternate implementations (e.g., `gocv`) later.

## Complete Current Object, Interface, and Protocol Inventory

### `pkg/sikuli` objects

| Type | Kind | Role |  |  |
|---|---|---|---|---|
| `Point` | object | coordinate pair | ✅ |  |
| `Location` | object | parity-friendly coordinate object | ✅ |  |
| `Offset` | object | parity-friendly offset object | ✅ |  |
| `Rect` | object | geometry primitive | ✅ |  |
| `Region` | object | geometry + search defaults container | ✅ |  |
| `Screen` | object | screen identity/bounds abstraction | ✅ |  |
| `Image` | object | grayscale image holder | ✅ |  |
| `Pattern` | object | matching intent/configuration | ✅ |  |
| `Match` | object | match result payload | ✅ |  |
| `TextMatch` | object | OCR text match payload | ✅ |  |
| `OCRParams` | object | OCR request option payload | ✅ |  |
| `InputOptions` | object | input action option payload | ✅ |  |
| `InputController` | object | input automation orchestrator | ✅ |  |
| `ObserveOptions` | object | observe operation option payload | ✅ |  |
| `ObserveEventType` | object | observe event enum | ✅ |  |
| `ObserveEvent` | object | observe event payload | ✅ |  |
| `ObserverController` | object | observe orchestration controller | ✅ |  |
| `AppOptions` | object | app operation option payload | ✅ |  |
| `Window` | object | app/window payload | ✅ |  |
| `AppController` | object | app/window orchestration controller | ✅ |  |
| `Finder` | object | user-facing matching orchestrator | ✅ |  |
| `RuntimeSettings` | object | global runtime behavior values | ✅ |  |
| `Options` | object | typed string-map options wrapper | ✅ |  |

### `pkg/sikuli` interfaces

| Interface | Contract |  |  |
|---|---|---|---|
| `ImageAPI` | image surface | ✅ | Signature and tests are in place |
| `PatternAPI` | pattern surface | ✅ | Signature and tests are in place |
| `FinderAPI` | finder surface | ✅ | Signature and tests are in place |
| `RegionAPI` | region surface | ✅ | Signature and tests are in place |
| `InputAPI` | input automation surface | ✅ | Signature and tests are in place |
| `ObserveAPI` | observe/event surface | ✅ | Signature and tests are in place |
| `AppAPI` | app/window surface | ✅ | Signature and tests are in place |

### `internal/core` protocol objects

| Type | Kind | Role |  |  |
|---|---|---|---|---|
| `SearchRequest` | protocol object | match request | ✅ | Stable request contract |
| `MatchCandidate` | protocol object | match response item | ✅ | Stable response contract |
| `Matcher` | protocol interface | matcher boundary | ✅ | Used by finder protocol |
| `OCRRequest` | protocol object | OCR request | ✅ | Stable OCR request contract |
| `OCRWord` | protocol object | OCR word payload | ✅ | Stable OCR word contract |
| `OCRResult` | protocol object | OCR response payload | ✅ | Stable OCR response contract |
| `OCR` | protocol interface | OCR boundary | ✅ | Used by finder OCR protocol |
| `InputAction` | protocol object | input action enum | ✅ | Stable input action contract |
| `InputRequest` | protocol object | input request | ✅ | Stable input request contract |
| `Input` | protocol interface | input boundary | ✅ | Used by input controller |
| `ObserveEventType` | protocol object | observe event enum | ✅ | Stable observe event contract |
| `ObserveRequest` | protocol object | observe request | ✅ | Stable observe request contract |
| `ObserveEvent` | protocol object | observe event payload | ✅ | Stable observe payload contract |
| `Observer` | protocol interface | observe boundary | ✅ | Used by observer controller |
| `AppAction` | protocol object | app action enum | ✅ | Stable app action contract |
| `AppRequest` | protocol object | app request | ✅ | Stable app request contract |
| `WindowInfo` | protocol object | window payload | ✅ | Stable window payload contract |
| `AppResult` | protocol object | app response payload | ✅ | Stable app response contract |
| `App` | protocol interface | app boundary | ✅ | Used by app controller |

### `internal/cv` protocol implementation

| Type | Kind | Role |  |  |
|---|---|---|---|---|
| `NCCMatcher` | protocol implementer | default matcher | ✅ | Primary backend in use |
| `SADMatcher` | protocol implementer | alternate matcher | ✅ | Conformance-tested alternate |

### `internal/ocr` protocol implementation

| Type | Kind | Role |  |  |
|---|---|---|---|---|
| `unsupportedBackend` | protocol implementer | default OCR behavior | ✅ | returns unsupported unless gosseract tag is enabled |
| `gosseractBackend` | protocol implementer | OCR adapter | ✅ | enabled with `-tags gosseract` |

### `internal/input` protocol implementation

| Type | Kind | Role |  |  |
|---|---|---|---|---|
| `darwinBackend` | protocol implementer | concrete input for macOS | ✅ | supports move/click/type/hotkey dispatch |
| `linuxBackend` | protocol implementer | concrete input for Linux | ✅ | command-driven move/click/type/hotkey via `xdotool` |
| `windowsBackend` | protocol implementer | concrete input for Windows | ✅ | PowerShell-driven move/click/type/hotkey |
| `unsupportedBackend` | protocol implementer | non-target fallback input behavior | ✅ | returns unsupported on `!darwin && !linux && !windows` builds |

### `internal/observe` protocol implementation

| Type | Kind | Role |  |  |
|---|---|---|---|---|
| `pollingBackend` | protocol implementer | deterministic observe behavior | ✅ | matcher-driven interval polling for appear/vanish/change |

### `internal/app` protocol implementation

| Type | Kind | Role |  |  |
|---|---|---|---|---|
| `darwinBackend` | protocol implementer | concrete app/window for macOS | ✅ | supports open/focus/close/is-running/list-windows |
| `linuxBackend` | protocol implementer | concrete app/window for Linux | ✅ | command-driven open/focus/close/is-running/list-windows |
| `windowsBackend` | protocol implementer | concrete app/window for Windows | ✅ | PowerShell-driven open/focus/close/is-running/list-windows |
| `unsupportedBackend` | protocol implementer | non-target fallback behavior | ✅ | returns unsupported for non-darwin/linux/windows builds |

### `internal/testharness` protocol objects

| Type | Kind | Role |  |  |
|---|---|---|---|---|
| `GoldenCase` | protocol object | serialized test case schema | ✅ | Active fixture schema |
| `ExpectedMatch` | protocol object | expected match schema | ✅ | Active fixture schema |
| `CompareOptions` | protocol object | comparator tolerance schema | ✅ | Active comparator contract |

## Implementation Plan

### Workstream 1: Core API scaffolding

- Define signatures/defaults for:
  - `Image`, `Pattern`, `Match`, `Finder`, `Region`, `Screen`
- Define typed errors and runtime defaults.
- Enforce compatibility via documented API interfaces.

Status: ✅ Completed (baseline implemented)

Current extension state: Region geometry/runtime helper surface, Finder wait/vanish helpers, Region-scoped search/wait parity scaffolding, and Location/Offset parity objects are implemented and covered by unit tests.
Current extension state additionally includes `Options` typed configuration helpers, sorted `FindAll` parity helpers, OCR text-search APIs (`ReadText`/`FindText`) with optional `gosseract` backend integration, input automation scaffolding, observe/event scaffolding, and app/window scaffolding.

### Workstream 2: Matching engine and parity harness

- Implement deterministic image matching (threshold + sort ordering + mask/resize support).
- Add golden matcher corpus and comparator assertions.
- Run `go test ./...` from repo root as the regression baseline.

Status: ✅ Completed (baseline implemented)

### Next planned workstreams

1. Cross-platform backend hardening

### Scaffold vs concrete backend status

| Workstream | Baseline scaffold | Concrete backend |  |
|---|---|---|---|
| Workstream 5: OCR and text-search parity | ✅ | ✅ | gosseract module version is pinned and enabled with `-tags gosseract` |
| Workstream 6: Input automation and hotkey parity | ✅ | ✅ | concrete `darwin`/`linux`/`windows` backends implemented |
| Workstream 7: Observe/event subsystem parity | ✅ | ✅ | deterministic polling backend implemented in `internal/observe` |
| Workstream 8: App/window/process control parity | ✅ | ✅ | concrete `darwin`/`linux`/`windows` backends implemented |

### Workstream 3: API parity surface expansion

- Expand `pkg/sikuli` to include additional parity objects and behaviors (location/offset aliases, broader region/finder helpers, options surfaces).
- Maintain non-breaking evolution under the API compatibility protocol.

Status: ✅ Completed

Completed scope:
- Added location/offset alias conversions and parity-friendly wrappers on region helpers.
- Expanded finder/region helper surface with match-count convenience APIs.
- Preserved backward compatibility with additive-only API changes and coverage tests.

### Workstream 4: protocol completeness hardening

- Add alternate matcher backend(s) under the same `core.Matcher` protocol.
- Add conformance tests ensuring every backend obeys ordering/threshold/mask rules.

Status: ✅ Completed

Completed scope:
- Added and maintained alternate backend coverage via `NCCMatcher` and `SADMatcher`.
- Enforced shared behavior with protocol conformance tests for ordering, threshold, mask, and resize rules.

### Workstream 5: OCR and text-search parity

- Add OCR protocol contract in `internal/core`.
- Expose `Finder.ReadText/FindText` and region-scoped text operations.
- Integrate optional backend support through the pinned `gosseract` module version.

Status (Baseline scaffold): ✅ Completed
Status (Concrete backend): ✅ Completed (pinned `gosseract` backend with tagged tests)

### Workstream 6: Input automation and hotkey parity

- Add input protocol contract in `internal/core`.
- Expose `InputController` with move/click/type/hotkey APIs.
- Maintain deterministic request/validation tests while expanding concrete platform backends.

Status (Baseline scaffold): ✅ Completed
Status (Concrete backend): ✅ Completed (`darwin` + `linux` + `windows` backends implemented)

### Workstream 7: Observe/event subsystem parity

- Add observe protocol contract in `internal/core`.
- Expose `ObserverController` with appear/vanish/change APIs.
- Implement deterministic matcher-driven polling backend and conformance timing tests.

Status (Baseline scaffold): ✅ Completed
Status (Concrete backend): ✅ Completed (deterministic polling backend implemented)

### Workstream 8: App/window/process control parity

- Add app/window protocol contract in `internal/core`.
- Expose `AppController` with open/focus/close/is-running/list-window APIs.
- Implement concrete platform backends behind the protocol boundary (`darwin`, `linux`, `windows`).

Status (Baseline scaffold): ✅ Completed
Status (Concrete backend): ✅ Completed (`darwin` + `linux` + `windows` backends implemented)

## Feature Matrix (Current and Planned)

| Area | Scope | Priority |  |  |
|---|---|---|---|---|
| Geometry primitives | `Point`, `Rect`, `Region` construction and transforms | P0 | ✅ | includes region union/intersection/containment and runtime setters |
| Location/offset parity types | `Location`, `Offset` value objects | P0 | ✅ | supports parity-friendly coordinate APIs |
| Screen abstraction | `Screen` id/bounds object | P1 | ✅ | add monitor discovery later |
| Image model | `Image` constructors, copy, dimensions | P0 | ✅ | add advanced image utilities later |
| Pattern semantics | similarity, exact, offset, resize, mask | P0 | ✅ | currently fully covered by default table |
| Match result model | score, target, index, geometry | P0 | ✅ | extend with comparator helpers if needed |
| Finder single target | `Find` + fail semantics | P0 | ✅ | includes `Exists` and `Has` helper semantics |
| Finder wait/vanish semantics | `Wait` and `WaitVanish` timeout polling | P0 | ✅ | global wait scan rate polling |
| Finder multi-target | `FindAll` ordering + indexing | P0 | ✅ | deterministic order defined |
| Finder sorted multi-target helpers | `FindAllByRow` / `FindAllByColumn` | P0 | ✅ | helper sorting + reindexing behavior |
| Region-scoped search | `Region.Find/Exists/Has/Wait` with timeout polling | P0 | ✅ | uses source crop + finder backend |
| Region sorted multi-target helpers | `FindAll` / `FindAllByRow` / `FindAllByColumn` | P0 | ✅ | region-scoped delegation |
| Image crop protocol | `Image.Crop(rect)` absolute-coordinate crop behavior | P0 | ✅ | enables region-scoped search protocol |
| Finder protocol swappability | `SetMatcher(core.Matcher)` | P0 | ✅ | enables backend evolution |
| Global settings | `RuntimeSettings` get/update/reset | P1 | ✅ | expand settings map as parity grows |
| Options/config object | typed get/set/delete/clone/merge | P1 | ✅ | string-map compatibility helper |
| Signature compatibility layer | `ImageAPI`, `PatternAPI`, `FinderAPI`, `RegionAPI`, `InputAPI`, `ObserveAPI`, `AppAPI` | P0 | ✅ | compatibility documented in API docs |
| Core matcher protocol | `SearchRequest`, `MatchCandidate`, `Matcher` | P0 | ✅ | strict boundary maintained |
| Core image protocol util | `ResizeGrayNearest` | P1 | ✅ | may add interpolation variants later |
| CV backend implementation | `NCCMatcher` | P0 | ✅ | first backend |
| Alternate matcher backend | `SADMatcher` | P1 | ✅ | enables multi-backend protocol checks |
| Golden parity protocol | corpus loader + comparator + tests | P0 | ✅ | active in CI/local tests |
| Backend conformance protocol | ordering/threshold/mask/resize assertions | P0 | ✅ | active tests in `internal/cv` |
| CI test visibility | race tests + vet + tidy diff enforcement | P0 | ✅ | workflow publishes strict signal |
| End-to-end parity flows | app + input + observe + OCR chained behavior | P1 | ✅ | dedicated parity e2e tests for default and `-tags gosseract` builds |
| OCR/text search | read text/find text parity | P1 | ✅ | finder/region OCR APIs with optional `gosseract` backend |
| OCR backend swappability | `core.OCR` protocol + backend selection | P1 | ✅ | unsupported default + pinned `gosseract` build-tag backend |
| OCR conformance tests | confidence filtering + ordering + backend behavior | P1 | ✅ | includes unsupported backend and tagged hOCR parser conformance tests |
| Input automation | mouse/keyboard parity | P1 | ✅ | `InputController` scaffolding with protocol boundary and tests |
| Input backend swappability | `core.Input` protocol + backend selection | P1 | ✅ | concrete `darwin`/`linux`/`windows` backends + non-target fallback |
| Observe/events | appear/vanish/change parity | P1 | ✅ | `ObserverController` + concrete deterministic polling backend |
| Observe backend swappability | `core.Observer` protocol + backend selection | P1 | ✅ | concrete default polling backend via `internal/observe` |
| App/window/process | focus/open/close/window parity | P2 | ✅ | `AppController` protocol with concrete `darwin`/`linux`/`windows` backends |
| App backend swappability | `core.App` protocol + backend selection | P2 | ✅ | concrete backends for major desktop OS targets |

## Protocol Completion Criteria

Each existing object/interface/protocol is considered feature-complete when:

1. It has signature coverage in the generated API docs and compatibility interfaces.
2. It has default/behavior semantics in `docs/default-behavior-table.md`.
3. Its package boundary and role are defined in this strategy document.
4. It is covered by unit or parity tests where behavior is non-trivial.

## Related Documents

- `docs/default-behavior-table.md`
- `docs/backend-capability-matrix.md`
