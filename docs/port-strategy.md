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

| Type | Kind | Role | Status |
|---|---|---|---|
| `Point` | object | coordinate pair | Implemented |
| `Rect` | object | geometry primitive | Implemented |
| `Region` | object | geometry + search defaults container | Implemented |
| `Screen` | object | screen identity/bounds abstraction | Implemented |
| `Image` | object | grayscale image holder | Implemented |
| `Pattern` | object | matching intent/configuration | Implemented |
| `Match` | object | match result payload | Implemented |
| `Finder` | object | user-facing matching orchestrator | Implemented |
| `RuntimeSettings` | object | global runtime behavior values | Implemented |

### `pkg/sikuli` interfaces

| Interface | Contract | Status |
|---|---|---|
| `ImageAPI` | stable image surface | Implemented |
| `PatternAPI` | stable pattern surface | Implemented |
| `FinderAPI` | stable finder surface | Implemented |

### `internal/core` protocol objects

| Type | Kind | Role | Status |
|---|---|---|---|
| `SearchRequest` | protocol object | backend-neutral match request | Implemented |
| `MatchCandidate` | protocol object | backend-neutral match response item | Implemented |
| `Matcher` | protocol interface | backend matcher boundary | Implemented |

### `internal/cv` protocol implementation

| Type | Kind | Role | Status |
|---|---|---|---|
| `NCCMatcher` | protocol implementer | default matcher backend | Implemented |

### `internal/testharness` protocol objects

| Type | Kind | Role | Status |
|---|---|---|---|
| `GoldenCase` | protocol object | serialized test case schema | Implemented |
| `ExpectedMatch` | protocol object | expected match schema | Implemented |
| `CompareOptions` | protocol object | comparator tolerance schema | Implemented |

## Implementation Plan

### Workstream 1: Core API scaffolding

- Freeze signatures/defaults for:
  - `Image`, `Pattern`, `Match`, `Finder`, `Region`, `Screen`
- Define typed errors and runtime defaults.
- Enforce compatibility via signature freeze docs and interfaces.

Status: complete baseline implemented.

Current extension state: Region geometry/runtime helper surface, Finder existence helpers, and Region-scoped search/wait parity scaffolding are implemented and covered by unit tests.

### Workstream 2: Matching engine and parity harness

- Implement deterministic image matching (threshold + sort ordering + mask/resize support).
- Add golden matcher corpus and comparator assertions.
- Run `go test ./...` from repo root as the regression baseline.

Status: complete baseline implemented.

### Next planned workstreams

1. OCR and text-search parity
2. Input automation and hotkey parity
3. Observe/event subsystem parity
4. App/window/process control parity
5. Cross-platform backend hardening

### Workstream 3: API parity surface expansion

- Expand `pkg/sikuli` to include additional parity objects and behaviors (location/offset aliases, broader region/finder helpers, options surfaces).
- Maintain non-breaking evolution under the API freeze protocol.

Status: planned.

### Workstream 4: protocol completeness hardening

- Add alternate matcher backend(s) under the same `core.Matcher` protocol.
- Add conformance tests ensuring every backend obeys ordering/threshold/mask rules.

Status: planned.

## Feature Matrix (Current and Planned)

| Area | Scope | Priority | Status | Notes |
|---|---|---|---|---|
| Geometry primitives | `Point`, `Rect`, `Region` construction and transforms | P0 | Implemented extended | includes region union/intersection/containment and runtime setters |
| Screen abstraction | `Screen` id/bounds object | P1 | Implemented baseline | add monitor discovery later |
| Image model | `Image` constructors, copy, dimensions | P0 | Implemented baseline | add advanced image utilities later |
| Pattern semantics | similarity, exact, offset, resize, mask | P0 | Implemented baseline | currently fully covered by default table |
| Match result model | score, target, index, geometry | P0 | Implemented baseline | extend with comparator helpers if needed |
| Finder single target | `Find` + fail semantics | P0 | Implemented extended | includes `Exists` and `Has` helper semantics |
| Finder multi-target | `FindAll` ordering + indexing | P0 | Implemented baseline | deterministic order locked |
| Region-scoped search | `Region.Find/Exists/Has/Wait` with timeout polling | P0 | Implemented extended | uses source crop + finder backend |
| Image crop protocol | `Image.Crop(rect)` absolute-coordinate crop behavior | P0 | Implemented extended | enables region-scoped search protocol |
| Finder protocol swappability | `SetMatcher(core.Matcher)` | P0 | Implemented baseline | enables backend evolution |
| Global settings | `RuntimeSettings` get/update/reset | P1 | Implemented baseline | expand settings map as parity grows |
| Signature compatibility layer | `ImageAPI`, `PatternAPI`, `FinderAPI` | P0 | Implemented baseline | freeze enforced in docs |
| Core matcher protocol | `SearchRequest`, `MatchCandidate`, `Matcher` | P0 | Implemented baseline | strict boundary maintained |
| Core image protocol util | `ResizeGrayNearest` | P1 | Implemented baseline | may add interpolation variants later |
| CV backend implementation | `NCCMatcher` | P0 | Implemented baseline | first backend |
| Golden parity protocol | corpus loader + comparator + tests | P0 | Implemented baseline | active in CI/local tests |
| OCR/text search | read text/find text parity | P1 | Planned | Not yet implemented |
| Input automation | mouse/keyboard parity | P1 | Planned | Not yet implemented |
| Observe/events | appear/vanish/change parity | P1 | Planned | Not yet implemented |
| App/window/process | focus/open/close/window parity | P2 | Planned | Not yet implemented |

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
