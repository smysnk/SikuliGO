# Search Semantics Matrix

Phase 1 establishes the stable search contract shared by:
- `packages/api/pkg/sikuli`
- `packages/api/internal/grpcv1`

This phase does not change the underlying architecture described in `Areas Where the Go Port Exposes Different or Additional Behavior`.

## Compatibility Decision

The stable parity contract for the Go port is explicit Go return semantics, not runtime-configurable throw/null mode.

Retained compatibility flags:
- `Region.ThrowException`
- `RuntimeSettings.FindFailedThrows`

Phase 1 replacement contract:
- these flags remain available as compatibility metadata for ports coming from SikuliX
- they do not alter Go method signatures
- they do not suppress `ErrFindFailed`
- they do not suppress `ErrTimeout`
- screen-backed gRPC calls map to the same contract rather than maintaining a separate miss/wait policy

## Canonical Behavior Matrix

| Operation | Image-backed `pkg/sikuli` behavior | Screen-backed gRPC behavior |
|---|---|---|
| `find` | returns `Match, nil` on hit; returns `ErrFindFailed` on miss | returns `FindResponse` on hit; maps miss to `codes.NotFound` |
| `exists` | returns `(Match{}, false, nil)` on miss; returns `(Match, true, nil)` on hit | returns `ExistsOnScreenResponse{exists:false}` on miss; returns `exists:true` with `match` on hit |
| `has` | returns `bool, error`; miss is `false, nil` | composed from `exists` semantics |
| `wait` | polls until hit; returns `ErrTimeout` when the wait budget is exhausted | polls until hit; maps timeout to `codes.DeadlineExceeded` |
| `waitVanish` | returns `true, nil` once absent; returns `false, nil` on timeout | not yet exposed as a dedicated RPC in Phase 1 |

## Shared Semantic Source

Phase 1 moves the search source of truth into:
- `packages/api/pkg/sikuli/search_semantics.go`

Shared helpers:
- `SearchExists`
- `SearchWait`
- `SearchWaitVanish`

These helpers define:
- `ErrFindFailed` means "miss"
- timeout `<= 0` means "probe once" for `exists` and `waitVanish`
- timeout `<= 0` for `wait` means one probe and `ErrTimeout` on miss
- non-find errors abort immediately

## Phase 1 Test Coverage

Image-backed coverage:
- `packages/api/pkg/sikuli/search_semantics_test.go`
- `packages/api/pkg/sikuli/scaffolding_test.go`

Screen-backed coverage:
- `packages/api/internal/grpcv1/server_test.go`

## Deferred Work

Phase 1 intentionally does not do the following:
- add live `Screen` capture/search APIs to `pkg/sikuli`
- make `Region` a live server-backed object
- introduce configurable null/exception mode into public Go signatures
- add `waitVanish` as a screen RPC

Those changes belong to later phases in:
- `docs/reference/parity/api-parity-phased-plan.md`
