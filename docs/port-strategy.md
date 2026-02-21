# GoLang Port Strategy

This document consolidates the project goals, locked architecture, implementation plan, and feature matrix for the SikuliGO port.

## Goals

- Build a feature-complete GoLang port of the core Sikuli API concepts.
- Preserve behavioral parity for matching, region semantics, and finder flows.
- Keep a stable public API while allowing backend upgrades.
- Make behavior measurable with deterministic tests and parity fixtures.

## Locked Architecture

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

| Type | Kind | Role | Status | Notes |
|---|---|---|---|---|
| `Point` | object | coordinate pair | ✅ | Implemented in current baseline |
| `Location` | object | parity-friendly coordinate object | ✅ | Implemented in current baseline |
| `Offset` | object | parity-friendly offset object | ✅ | Implemented in current baseline |
| `Rect` | object | geometry primitive | ✅ | Implemented in current baseline |
| `Region` | object | geometry + search defaults container | ✅ | Implemented in current baseline |
| `Screen` | object | screen identity/bounds abstraction | ✅ | Implemented in current baseline |
| `Image` | object | grayscale image holder | ✅ | Implemented in current baseline |
| `Pattern` | object | matching intent/configuration | ✅ | Implemented in current baseline |
| `Match` | object | match result payload | ✅ | Implemented in current baseline |
| `TextMatch` | object | OCR text match payload | ✅ | Implemented in current baseline |
| `OCRParams` | object | OCR request option payload | ✅ | Implemented in current baseline |
| `InputOptions` | object | input action option payload | ✅ | Implemented in current baseline |
| `InputController` | object | input automation orchestrator | ✅ | Implemented in current baseline |
| `ObserveOptions` | object | observe operation option payload | ✅ | Implemented in current baseline |
| `ObserveEventType` | object | observe event enum | ✅ | Implemented in current baseline |
| `ObserveEvent` | object | observe event payload | ✅ | Implemented in current baseline |
| `ObserverController` | object | observe orchestration controller | ✅ | Implemented in current baseline |
| `AppOptions` | object | app operation option payload | ✅ | Implemented in current baseline |
| `Window` | object | app/window payload | ✅ | Implemented in current baseline |
| `AppController` | object | app/window orchestration controller | ✅ | Implemented in current baseline |
| `Finder` | object | user-facing matching orchestrator | ✅ | Implemented in current baseline |
| `RuntimeSettings` | object | global runtime behavior values | ✅ | Implemented in current baseline |
| `Options` | object | typed string-map options wrapper | ✅ | Implemented in current baseline |

### `pkg/sikuli` interfaces

| Interface | Contract | Status | Notes |
|---|---|---|---|
| `ImageAPI` | stable image surface | ✅ | Signature and tests are in place |
| `PatternAPI` | stable pattern surface | ✅ | Signature and tests are in place |
| `FinderAPI` | stable finder surface | ✅ | Signature and tests are in place |
| `RegionAPI` | stable region surface | ✅ | Signature and tests are in place |
| `InputAPI` | stable input automation surface | ✅ | Signature and tests are in place |
| `ObserveAPI` | stable observe/event surface | ✅ | Signature and tests are in place |
| `AppAPI` | stable app/window surface | ✅ | Signature and tests are in place |

### `internal/core` protocol objects

| Type | Kind | Role | Status | Notes |
|---|---|---|---|---|
| `SearchRequest` | protocol object | backend-neutral match request | ✅ | Locked request contract |
| `MatchCandidate` | protocol object | backend-neutral match response item | ✅ | Locked response contract |
| `Matcher` | protocol interface | backend matcher boundary | ✅ | Used by finder protocol |
| `OCRRequest` | protocol object | backend-neutral OCR request | ✅ | Locked OCR request contract |
| `OCRWord` | protocol object | backend-neutral OCR word payload | ✅ | Locked OCR word contract |
| `OCRResult` | protocol object | backend-neutral OCR response payload | ✅ | Locked OCR response contract |
| `OCR` | protocol interface | backend OCR boundary | ✅ | Used by finder OCR protocol |
| `InputAction` | protocol object | backend-neutral input action enum | ✅ | Locked input action contract |
| `InputRequest` | protocol object | backend-neutral input request | ✅ | Locked input request contract |
| `Input` | protocol interface | backend input boundary | ✅ | Used by input controller |
| `ObserveEventType` | protocol object | backend-neutral observe event enum | ✅ | Locked observe event contract |
| `ObserveRequest` | protocol object | backend-neutral observe request | ✅ | Locked observe request contract |
| `ObserveEvent` | protocol object | backend-neutral observe event payload | ✅ | Locked observe payload contract |
| `Observer` | protocol interface | backend observe boundary | ✅ | Used by observer controller |
| `AppAction` | protocol object | backend-neutral app action enum | ✅ | Locked app action contract |
| `AppRequest` | protocol object | backend-neutral app request | ✅ | Locked app request contract |
| `WindowInfo` | protocol object | backend-neutral window payload | ✅ | Locked window payload contract |
| `AppResult` | protocol object | backend-neutral app response payload | ✅ | Locked app response contract |
| `App` | protocol interface | backend app boundary | ✅ | Used by app controller |

### `internal/cv` protocol implementation

| Type | Kind | Role | Status | Notes |
|---|---|---|---|---|
| `NCCMatcher` | protocol implementer | default matcher backend | ✅ | Primary backend in use |
| `SADMatcher` | protocol implementer | alternate matcher backend | ✅ | Conformance-tested alternate |

### `internal/ocr` protocol implementation

| Type | Kind | Role | Status | Notes |
|---|---|---|---|---|
| `unsupportedBackend` | protocol implementer | default OCR backend behavior | ✅ | returns unsupported unless gosseract tag is enabled |
| `gosseractBackend` | protocol implementer | OCR backend adapter | ✅ | enabled with `-tags gosseract` |

### `internal/input` protocol implementation

| Type | Kind | Role | Status | Notes |
|---|---|---|---|---|
| `darwinBackend` | protocol implementer | concrete input backend for macOS | ✅ | supports move/click/type/hotkey dispatch |
| `linuxBackend` | protocol implementer | concrete input backend for Linux | ✅ | command-driven move/click/type/hotkey via `xdotool` |
| `windowsBackend` | protocol implementer | concrete input backend for Windows | ✅ | PowerShell-driven move/click/type/hotkey |
| `unsupportedBackend` | protocol implementer | non-target fallback input behavior | ✅ | returns unsupported on `!darwin && !linux && !windows` builds |

### `internal/observe` protocol implementation

| Type | Kind | Role | Status | Notes |
|---|---|---|---|---|
| `pollingBackend` | protocol implementer | deterministic observe backend behavior | ✅ | matcher-driven interval polling for appear/vanish/change |

### `internal/app` protocol implementation

| Type | Kind | Role | Status | Notes |
|---|---|---|---|---|
| `darwinBackend` | protocol implementer | concrete app/window backend for macOS | ✅ | supports open/focus/close/is-running/list-windows |
| `linuxBackend` | protocol implementer | concrete app/window backend for Linux | ✅ | command-driven open/focus/close/is-running/list-windows |
| `windowsBackend` | protocol implementer | concrete app/window backend for Windows | ✅ | PowerShell-driven open/focus/close/is-running/list-windows |
| `unsupportedBackend` | protocol implementer | non-target fallback backend behavior | ✅ | returns unsupported for non-darwin/linux/windows builds |

### `internal/testharness` protocol objects

| Type | Kind | Role | Status | Notes |
|---|---|---|---|---|
| `GoldenCase` | protocol object | serialized test case schema | ✅ | Active fixture schema |
| `ExpectedMatch` | protocol object | expected match schema | ✅ | Active fixture schema |
| `CompareOptions` | protocol object | comparator tolerance schema | ✅ | Active comparator contract |

## Implementation Plan

### Workstream 1: Core API scaffolding

- Freeze signatures/defaults for:
  - `Image`, `Pattern`, `Match`, `Finder`, `Region`, `Screen`
- Define typed errors and runtime defaults.
- Enforce compatibility via signature freeze docs and interfaces.

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
2. Protocol completeness hardening

### Scaffold vs concrete backend status

| Workstream | Baseline scaffold | Concrete backend | Notes |
|---|---|---|---|
| Workstream 5: OCR and text-search parity | ✅ | ✅ | gosseract module version is pinned and enabled with `-tags gosseract` |
| Workstream 6: Input automation and hotkey parity | ✅ | ✅ | concrete `darwin`/`linux`/`windows` backends implemented |
| Workstream 7: Observe/event subsystem parity | ✅ | ✅ | deterministic polling backend implemented in `internal/observe` |
| Workstream 8: App/window/process control parity | ✅ | ✅ | concrete `darwin`/`linux`/`windows` backends implemented |

### Workstream 3: API parity surface expansion

- Expand `pkg/sikuli` to include additional parity objects and behaviors (location/offset aliases, broader region/finder helpers, options surfaces).
- Maintain non-breaking evolution under the API freeze protocol.

Status: 🟡 Planned

### Workstream 4: protocol completeness hardening

- Add alternate matcher backend(s) under the same `core.Matcher` protocol.
- Add conformance tests ensuring every backend obeys ordering/threshold/mask rules.

Status: 🟡 Planned

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

| Area | Scope | Priority | Status | Notes |
|---|---|---|---|---|
| Geometry primitives | `Point`, `Rect`, `Region` construction and transforms | P0 | ✅ | includes region union/intersection/containment and runtime setters |
| Location/offset parity types | `Location`, `Offset` value objects | P0 | ✅ | supports parity-friendly coordinate APIs |
| Screen abstraction | `Screen` id/bounds object | P1 | ✅ | add monitor discovery later |
| Image model | `Image` constructors, copy, dimensions | P0 | ✅ | add advanced image utilities later |
| Pattern semantics | similarity, exact, offset, resize, mask | P0 | ✅ | currently fully covered by default table |
| Match result model | score, target, index, geometry | P0 | ✅ | extend with comparator helpers if needed |
| Finder single target | `Find` + fail semantics | P0 | ✅ | includes `Exists` and `Has` helper semantics |
| Finder wait/vanish semantics | `Wait` and `WaitVanish` timeout polling | P0 | ✅ | global wait scan rate polling |
| Finder multi-target | `FindAll` ordering + indexing | P0 | ✅ | deterministic order locked |
| Finder sorted multi-target helpers | `FindAllByRow` / `FindAllByColumn` | P0 | ✅ | helper sorting + reindexing behavior |
| Region-scoped search | `Region.Find/Exists/Has/Wait` with timeout polling | P0 | ✅ | uses source crop + finder backend |
| Region sorted multi-target helpers | `FindAll` / `FindAllByRow` / `FindAllByColumn` | P0 | ✅ | region-scoped delegation |
| Image crop protocol | `Image.Crop(rect)` absolute-coordinate crop behavior | P0 | ✅ | enables region-scoped search protocol |
| Finder protocol swappability | `SetMatcher(core.Matcher)` | P0 | ✅ | enables backend evolution |
| Global settings | `RuntimeSettings` get/update/reset | P1 | ✅ | expand settings map as parity grows |
| Options/config object | typed get/set/delete/clone/merge | P1 | ✅ | string-map compatibility helper |
| Signature compatibility layer | `ImageAPI`, `PatternAPI`, `FinderAPI`, `RegionAPI`, `InputAPI`, `ObserveAPI`, `AppAPI` | P0 | ✅ | freeze enforced in docs |
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

1. It has frozen signature coverage in `docs/api-signature-freeze.md`.
2. It has default/behavior semantics in `docs/default-behavior-table.md`.
3. Its package boundary and role are defined in `docs/architecture-lock.md`.
4. It is covered by unit or parity tests where behavior is non-trivial.

## Related Documents

- `docs/architecture-lock.md`
- `docs/api-signature-freeze.md`
- `docs/default-behavior-table.md`
- `docs/backend-capability-matrix.md`
