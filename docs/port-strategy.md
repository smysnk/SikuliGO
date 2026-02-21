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

### `internal/core` protocol objects

| Type | Kind | Role | Status | Notes |
|---|---|---|---|---|
| `SearchRequest` | protocol object | backend-neutral match request | ✅ | Locked request contract |
| `MatchCandidate` | protocol object | backend-neutral match response item | ✅ | Locked response contract |
| `Matcher` | protocol interface | backend matcher boundary | ✅ | Used by finder protocol |

### `internal/cv` protocol implementation

| Type | Kind | Role | Status | Notes |
|---|---|---|---|---|
| `NCCMatcher` | protocol implementer | default matcher backend | ✅ | Primary backend in use |
| `SADMatcher` | protocol implementer | alternate matcher backend | ✅ | Conformance-tested alternate |

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
Current extension state additionally includes `Options` typed configuration helpers and sorted `FindAll` parity helpers.

### Workstream 2: Matching engine and parity harness

- Implement deterministic image matching (threshold + sort ordering + mask/resize support).
- Add golden matcher corpus and comparator assertions.
- Run `go test ./...` from repo root as the regression baseline.

Status: ✅ Completed (baseline implemented)

### Next planned workstreams

1. OCR and text-search parity
2. Input automation and hotkey parity
3. Observe/event subsystem parity
4. App/window/process control parity
5. Cross-platform backend hardening

### Workstream 3: API parity surface expansion

- Expand `pkg/sikuli` to include additional parity objects and behaviors (location/offset aliases, broader region/finder helpers, options surfaces).
- Maintain non-breaking evolution under the API freeze protocol.

Status: 🟡 Planned

### Workstream 4: protocol completeness hardening

- Add alternate matcher backend(s) under the same `core.Matcher` protocol.
- Add conformance tests ensuring every backend obeys ordering/threshold/mask rules.

Status: 🟡 Planned

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
| Signature compatibility layer | `ImageAPI`, `PatternAPI`, `FinderAPI` | P0 | ✅ | freeze enforced in docs |
| Core matcher protocol | `SearchRequest`, `MatchCandidate`, `Matcher` | P0 | ✅ | strict boundary maintained |
| Core image protocol util | `ResizeGrayNearest` | P1 | ✅ | may add interpolation variants later |
| CV backend implementation | `NCCMatcher` | P0 | ✅ | first backend |
| Alternate matcher backend | `SADMatcher` | P1 | ✅ | enables multi-backend protocol checks |
| Golden parity protocol | corpus loader + comparator + tests | P0 | ✅ | active in CI/local tests |
| Backend conformance protocol | ordering/threshold/mask/resize assertions | P0 | ✅ | active tests in `internal/cv` |
| CI test visibility | race tests + vet + tidy diff enforcement | P0 | ✅ | workflow publishes strict signal |
| OCR/text search | read text/find text parity | P1 | 🟡 | Not yet implemented |
| Input automation | mouse/keyboard parity | P1 | 🟡 | Not yet implemented |
| Observe/events | appear/vanish/change parity | P1 | 🟡 | Not yet implemented |
| App/window/process | focus/open/close/window parity | P2 | 🟡 | Not yet implemented |

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
