# Client Parity Phased Plan

This plan follows `docs/reference/parity/api-parity-phased-plan.md` and covers how Node, Python, and Lua clients should adopt the API parity work for items 1-9 from `docs/reference/parity/sikulix-api-docs-audit.md`.

Scope:
- Client wrappers and examples only.
- Target surfaces:
  - `packages/client-node`
  - `packages/client-python`
  - `packages/client-lua`

Guardrails:
- Do not redesign the runtime architecture.
- Clients should consume the API parity work; they should not independently invent behavior that the underlying API does not guarantee.
- Node and Python remain the primary reference clients. Lua should follow once the core client contract is stable.

## Item Coverage

| Audit Item | Client Concern | Planned In |
|---|---|---|
| 1 | Live `Region` behavior exposed idiomatically in clients | Phases 1-2 |
| 2 | `Screen` monitor/capture model exposed in clients | Phase 2 |
| 3 | `Match` convenience behavior in client classes | Phase 3 |
| 4 | Direct action methods on region/match/screen classes | Phase 4 |
| 5 | Finder traversal wrappers | Phase 5 |
| 6 | Null/throw/timeout semantics mapped consistently | Phases 1 and 5 |
| 7 | Multi-target search helpers | Phase 6 |
| 8 | OCR helper families and richer result types | Phase 7 |
| 9 | App/window convenience surface | Phase 8 |

## Phase 0: Client Contract Freeze

Goal:
- Prevent the client wrappers from drifting while the API parity work lands.

Deliverables:
- Document a compatibility policy for Node, Python, and Lua that distinguishes:
  - stable parity wrappers
  - transitional wrappers
  - API features that are still client-composed
- Identify all client-only behaviors that currently simulate missing API parity.
- Mark which wrappers must be replaced once the underlying API phases land.

Exit criteria:
- The client layer has an explicit inventory of temporary shims versus permanent contract.

## Phase 1: Search Semantic Alignment

Goal:
- Make clients reflect the final API semantics for miss, exists, wait, and timeout behavior.

Depends on:
- API Phase 1

Items addressed:
- 1. `Region` runtime model groundwork
- 6. null/throw/timeout semantics

Deliverables:
- Align Node and Python `Region.find`, `exists`, `wait`, and `waitVanish` semantics with the API decisions from Phase 1.
- Remove any client-specific semantic drift where Node, Python, and Lua currently behave differently around misses or timeouts.
- Add compatibility tests that verify one behavior matrix across clients.

Implementation notes:
- The clients should not continue carrying parallel interpretation logic once the API contract is stable.
- Preserve ergonomic method names, but make behavior consistent.

Exit criteria:
- The same scenario produces the same semantic result across Node and Python.

## Phase 2: Live `Screen` and `Region` Client Surface

Goal:
- Expose the new live-screen API capabilities cleanly in the client object model.

Depends on:
- API Phase 2

Items addressed:
- 1. live `Region` model
- 2. `Screen` surface parity

Deliverables:
- Add client support for:
  - screen enumeration
  - primary/default screen
  - screen-by-ID selection
  - screen/region capture helpers
  - region derivation from a live screen
- Ensure `Screen` and `Region` wrappers mirror the same mental model across Node and Python.
- Update examples so live-screen workflows no longer need hidden client composition for capabilities that now exist in the API.

Implementation notes:
- Avoid introducing client-only screen concepts that are not represented in the API.
- Lua adoption can lag one phase behind Node/Python if necessary, but the shape should match.

Exit criteria:
- Node and Python can express documented SikuliX-like `Screen` / live `Region` workflows using the same API-backed object model.

## Phase 3: Client `Match` Behavior

Goal:
- Make client `Match` values reflect the richer API match contract without inventing extra client-only semantics.

Depends on:
- API Phase 3

Items addressed:
- 3. `Match` semantics

Deliverables:
- Add direct `Match` convenience methods once they exist at the API level.
- Keep `Match` geometry/score/target fields stable while layering behavior on top.
- Update client docs to show match-chaining flows rather than forcing coordinate extraction.

Exit criteria:
- Client `Match` objects support the same broad action model exposed by the API.

## Phase 4: Direct Action Vocabulary in Clients

Goal:
- Surface the full direct action vocabulary in the client wrappers.

Depends on:
- API Phase 4

Items addressed:
- 4. direct action parity

Deliverables:
- Add Node and Python wrappers for:
  - `doubleClick`
  - `rightClick`
  - `hover`
  - `dragDrop`
  - `paste`
  - wheel scrolling
  - key up/down
  - mouse up/down
- Keep method names aligned with SikuliX naming where reasonable.
- Update Lua only after Node and Python wrappers are stable and documented.

Implementation notes:
- Prefer thin wrappers over stable API primitives.
- Do not keep re-implementing these as client-specific composites once protocol support exists.

Exit criteria:
- Client users can write SikuliX-style action flows without dropping into low-level transport calls.

## Phase 5: Finder Traversal and Compatibility Wrappers

Goal:
- Support SikuliX-style traversal behavior in clients once the underlying API exposes it.

Depends on:
- API Phase 5

Items addressed:
- 5. Finder iteration/lifecycle
- 6. remaining semantic alignment

Deliverables:
- Add iterator-compatible client wrappers where justified.
- Decide whether client APIs should expose:
  - idiomatic language-native iteration only, or
  - an explicit compatibility iterator layer in addition to idiomatic iteration
- Ensure any compatibility wrapper does not drift from the underlying finder state model.

Exit criteria:
- SikuliX iterator-style ports can be expressed in Node and Python without custom client-side state machines.

## Phase 6: Multi-Target Helper Adoption

Goal:
- Surface multi-target search helpers once the API contract exists.

Depends on:
- API Phase 6

Items addressed:
- 7. multi-target helpers

Deliverables:
- Add client methods for batch/multi-target search helpers.
- Keep return ordering and tie-breaking behavior consistent with API guarantees.
- Add tests covering mixed-hit and no-hit cases across clients.

Exit criteria:
- Node and Python can express SikuliX multi-pattern search flows with stable semantics.

## Phase 7: OCR Workflow Expansion in Clients

Goal:
- Expose richer OCR helper families as thin client wrappers.

Depends on:
- API Phase 7

Items addressed:
- 8. OCR helper surface

Deliverables:
- Add client wrappers for word/line collection helpers.
- Standardize OCR result objects across Node and Python.
- Update OCR examples to demonstrate the richer API surface.

Exit criteria:
- OCR client usage no longer stops at `readText` and `findText`.

## Phase 8: App and Window Client Expansion

Goal:
- Expose the richer app/window API in the client wrappers.

Depends on:
- API Phase 8

Items addressed:
- 9. App/window model

Deliverables:
- Add Node and Python wrappers for focused-window and richer window metadata flows.
- Decide how much of the window object becomes behavior-bearing versus record-only in clients.
- Update docs and examples for common app/window workflows.

Exit criteria:
- Client app/window ergonomics match the API parity target and no longer undershoot the documented SikuliX workflows materially.

## Phase 9: Client Hardening, Docs, and Migration Guides

Goal:
- Make the new parity surface usable and durable for consumers.

Depends on:
- API Phase 9

Deliverables:
- Update client readmes and examples to prefer the parity-shaped workflows.
- Add migration guides from SikuliX docs/examples into Node and Python examples.
- Add CI parity smoke tests covering:
  - one image-backed flow
  - one live-screen flow
  - one OCR flow
  - one app/window flow
- Decide whether Lua remains a first-class parity target or a descriptor-level transport client.

Exit criteria:
- Clients document and test the parity surface rather than relying on scattered examples and one-off shims.

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
- Clients should lag the API by one stable milestone, not lead it.
- Search semantics and live-screen object model must be stable before client ergonomic expansion.
- Docs and migration examples should be refreshed only after the parity surface settles.
