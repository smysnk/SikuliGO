# API Parity Phased Plan

This plan addresses items 1-9 from `docs/reference/parity/sikulix-api-docs-audit.md`.

Scope:
- Underlying protocol and Go API only.
- Target surfaces:
  - `packages/api/proto/sikuli/v1/sikuli.proto`
  - `packages/api/internal/grpcv1`
  - `packages/api/pkg/sikuli`

Guardrails:
- Do not change the runtime architecture described in the audit section `Areas Where the Go Port Exposes Different or Additional Behavior`.
- Keep the server/gRPC model.
- Keep explicit matcher-engine selection.
- Keep pattern masking support.
- Keep session/auth/logging/tracing behavior as architectural concerns of the API runtime.

## Item Coverage

| Audit Item | Topic | Planned In |
|---|---|---|
| 1 | `Region` live runtime model | Phases 1-2 |
| 2 | `Screen` surface parity | Phase 2 |
| 3 | `Match` as a `Region`-like result | Phase 3 |
| 4 | Direct action parity | Phase 4 |
| 5 | Finder iteration/lifecycle | Phase 5 |
| 6 | Exception/null semantics | Phases 1 and 5 |
| 7 | Multi-target search helpers | Phase 6 |
| 8 | OCR helper surface | Phase 7 |
| 9 | App/window model | Phase 8 |

## Phase 0: Contract Baseline

Goal:
- Freeze the parity target and prevent drift while implementation proceeds.

Deliverables:
- Convert the nine audit items into explicit parity tickets or checklist sections.
- Add a parity test manifest mapping each SikuliX-documented behavior to:
  - current Go behavior
  - target Go behavior
  - test owner
- Mark generated mapping docs where parity is currently overstated.

Exit criteria:
- Every item 1-9 has a testable acceptance definition.
- `java-to-go-mapping.md` no longer labels clearly incomplete areas as fully parity-ready.

## Phase 1: Search Semantics Foundation

Goal:
- Normalize miss, wait, and exception semantics before adding more surface area.

Items addressed:
- 1. `Region` runtime model groundwork
- 6. exception/null semantics

Deliverables:
- Define a canonical parity behavior matrix for:
  - `find`
  - `exists`
  - `wait`
  - `waitVanish`
  - `has`
- Wire `ThrowException` / `FindFailedThrows` into actual search behavior, or explicitly replace them with a different stable compatibility contract and document that replacement.
- Decide whether the stable Go parity contract is:
  - pure Go idiomatic returns, or
  - compatibility-mode behavior with configurable throw/null semantics
- Ensure image-backed and screen-backed search flows share one semantic implementation path.

Implementation notes:
- The semantic source of truth should live in `pkg/sikuli`, not separately in each client.
- gRPC handlers should map to the same behavior rules rather than inventing parallel wait/exists logic.

Exit criteria:
- One set of parity tests covers both image-backed and screen-backed search semantics.
- `exists`, `wait`, and miss behavior are deterministic and documented.

## Phase 2: Live `Region` and `Screen` Surface

Goal:
- Introduce a public live-screen API surface that matches SikuliX concepts without abandoning the current server architecture.

Items addressed:
- 1. `Region` live runtime model
- 2. `Screen` surface parity

Deliverables:
- Add monitor/screen protocol support for:
  - enumerate screens
  - resolve primary screen
  - address a specific screen by ID
  - capture from a specific screen or region
- Introduce a screen-backed `Region`/`Screen` facade in `pkg/sikuli` that routes to the API runtime instead of requiring a source `*Image` for live operations.
- Preserve current image-backed `Region` behavior for image-to-image workflows; do not remove it.
- Add public capture APIs that cover the SikuliX mental model:
  - whole screen capture
  - region capture
  - selected screen capture

Implementation notes:
- This phase should add public live-screen primitives, not merely client sugar.
- The current `Screen { ID, Bounds }` descriptor should evolve into a usable public surface rather than remain metadata-only.

Exit criteria:
- A Go caller can obtain a live `Screen`, derive a live `Region`, and perform documented screen-scoped search without manually calling gRPC stubs.
- Multi-monitor selection is testable and documented.

## Phase 3: `Match` as a First-Class Action Target

Goal:
- Make `Match` behave like a SikuliX action target instead of only returning geometry.

Items addressed:
- 3. `Match` semantics

Deliverables:
- Decide whether `Match` becomes:
  - a richer struct with action/search helpers, or
  - a lightweight wrapper around an internal region-like target object
- Add parity methods needed for direct action chaining from a match.
- Ensure `Match` retains:
  - rectangle
  - score
  - target point
  - stable region semantics

Implementation notes:
- Keep the existing value semantics where useful, but expose behavior expected by SikuliX-style flows.
- The added behavior should delegate to the same underlying live region/input layer created in Phase 2.

Exit criteria:
- A `Match` can be used directly in the same broad way SikuliX docs describe, without forcing callers to manually extract points and route through separate controllers.

## Phase 4: Direct Action API Parity

Goal:
- Close the major input/action gap on top of the live `Region`/`Match` surface.

Items addressed:
- 4. direct action parity

Deliverables:
- Add protocol and `pkg/sikuli` support for:
  - `hover`
  - `doubleClick`
  - `rightClick`
  - `dragDrop`
  - `paste`
  - wheel scrolling
  - stateful key down/up
  - stateful mouse down/up
- Decide which actions should be primitive RPCs versus composed helpers over existing RPCs.
- Keep low-level `InputController` available; add parity conveniences rather than replacing it.

Implementation notes:
- Prefer adding durable primitives for stateful behaviors (`keyDown`, `keyUp`, `mouseDown`, `mouseUp`) rather than emulating them unreliably in clients.
- Preserve current backends and dependency model.

Exit criteria:
- The documented SikuliX direct action vocabulary is represented at the API level, either as first-class methods or clearly documented composed operations with equivalent behavior.

## Phase 5: Finder Traversal and Lifecycle Semantics

Goal:
- Align the finder contract with SikuliX traversal behavior without regressing the current slice-based ergonomics.

Items addressed:
- 5. Finder iteration/lifecycle
- 6. remaining behavior semantics gaps

Deliverables:
- Add iterator-style traversal support on top of or alongside `[]Match` returns:
  - `HasNext`
  - `Next`
  - reset/close or `Destroy`
- Decide whether `FindAll` continues returning slices in Go while also exposing a compatibility iterator, or whether compatibility types are separated explicitly.
- Ensure `LastMatches()` and any new iterator state stay coherent.

Implementation notes:
- Do not remove slice returns if downstream code depends on them.
- Add a compatibility layer rather than forcing all Go callers into Java-like iteration.

Exit criteria:
- There is a stable finder traversal model that can support SikuliX-style ports and is covered by tests.

## Phase 6: Multi-Target Search Helpers

Goal:
- Add the multi-pattern helper family currently missing from the Go port.

Items addressed:
- 7. multi-target helpers

Deliverables:
- Add API-level support for helper families such as:
  - `findAnyList`
  - `findBestList`
  - equivalent list/batch matching helpers
- Decide whether batching is implemented as:
  - a multi-pattern RPC surface, or
  - a stable composed helper in `pkg/sikuli` over single-pattern RPCs
- Define deterministic tie-breaking and ordering rules.

Implementation notes:
- This phase should not merely expose repeated client loops; it should define stable API semantics.
- Explicitly document performance expectations when helpers are composed client-side versus executed server-side.

Exit criteria:
- SikuliX multi-target search flows can be expressed without client-specific ad hoc orchestration.

## Phase 7: OCR Surface Expansion

Goal:
- Expand OCR parity beyond the current `ReadText` and `FindText` pair.

Items addressed:
- 8. OCR helper surface

Deliverables:
- Add API support for richer OCR traversal helpers such as:
  - collect words
  - collect lines
  - char/word/line oriented results where justified
- Decide on stable result types for word and line segmentation.
- Align OCR configuration propagation across image-backed and screen-backed flows.

Implementation notes:
- Keep the current OCR backend model and build-tag expectations intact.
- Avoid redesigning OCR architecture; this phase is surface expansion only.

Exit criteria:
- OCR-heavy ports no longer need to reconstruct common SikuliX text workflows purely in clients.

## Phase 8: App and Window Parity Expansion

Goal:
- Close the most visible app/window parity gaps without trying to recreate the full Java runtime internals.

Items addressed:
- 9. App/window model

Deliverables:
- Expand protocol and Go app/window support for:
  - focused window lookup
  - richer window metadata where available
  - stable window selection/query helpers
- Decide whether app/window objects become richer first-class values or remain controller-returned records with helper methods.
- Document platform-specific variance where exact window metadata is not portable.

Implementation notes:
- Maintain the current app backend architecture.
- Focus on documented user-visible workflows rather than internal object parity for its own sake.

Exit criteria:
- Common SikuliX app/window flows have first-class API support in the Go port.

## Phase 9: Parity Hardening and Release Gates

Goal:
- Make parity regressions visible and block them.

Items addressed:
- all items 1-9 as stabilization work

Deliverables:
- Add parity contract tests for every newly implemented area.
- Add generated docs verification that reflects actual implementation maturity.
- Add migration examples demonstrating one image-backed flow and one live-screen flow for each newly closed parity area.

Exit criteria:
- Parity status is test-backed rather than narrative-only.
- New API parity work can ship without reopening already-closed gaps.

## Recommended Execution Order

1. Phase 0
2. Phase 1
3. Phase 2
4. Phase 3
5. Phase 4
6. Phase 5
7. Phase 6
8. Phase 7
9. Phase 8
10. Phase 9

Reasoning:
- Search semantics and live-screen surface decisions are foundational.
- `Match`, action methods, and client ergonomics should not move ahead until the live-screen contract is stable.
- OCR and app/window parity can come later without blocking the core automation model.
