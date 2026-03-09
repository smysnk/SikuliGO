# Parity Test Matrix

This matrix links the closed API parity areas to explicit contract coverage.

| Area | SikuliX Concept | Primary Contract Test | Supporting Coverage | Status |
|---|---|---|---|---|
| Search semantics foundation | `find/exists/wait/waitVanish` | `packages/api/pkg/sikuli/parity_contract_test.go` -> `TestAPIParityContracts/SearchSemantics` | `packages/api/pkg/sikuli/search_semantics_test.go`, `packages/api/internal/grpcv1/server_test.go` | ✅ |
| Exception/null semantics | `setThrowException`, miss/null behavior | `packages/api/pkg/sikuli/parity_contract_test.go` -> `TestAPIParityContracts/SearchSemantics` | `packages/api/pkg/sikuli/search_semantics_test.go`, `packages/api/pkg/sikuli/scaffolding_test.go` | ✅ |
| Live screen and region surface | `Screen`, live `Region`, capture | `packages/api/pkg/sikuli/parity_contract_test.go` -> `TestAPIParityContracts/LiveScreenAndRegionSurface` | `packages/api/pkg/sikuli/live_runtime_test.go`, `packages/api/internal/grpcv1/server_screen_test.go` | ✅ |
| Match as action target | live `Match` reuse | `packages/api/pkg/sikuli/parity_contract_test.go` -> `TestAPIParityContracts/MatchAsActionTarget` | `packages/api/pkg/sikuli/match_live_runtime_test.go`, `packages/api/pkg/sikuli/match_action_test.go` | ✅ |
| Direct action API parity | hover/click/right/double/drag/paste/key/mouse lifecycle/wheel | `packages/api/pkg/sikuli/parity_contract_test.go` -> `TestAPIParityContracts/DirectActionSurface` | `packages/api/internal/input/*_test.go`, `packages/api/internal/grpcv1/rpc_surface_integration_test.go`, `packages/api/pkg/sikuli/match_action_test.go` | ✅ |
| Finder traversal and lifecycle | `hasNext/next/destroy` | `packages/api/pkg/sikuli/parity_contract_test.go` -> `TestAPIParityContracts/FinderTraversalAndLifecycle` | `packages/api/pkg/sikuli/finder_iterator_test.go` | ✅ |
| Multi-target search helpers | `findAnyList/findBestList/waitAnyList/waitBestList` | `packages/api/pkg/sikuli/parity_contract_test.go` -> `TestAPIParityContracts/MultiTargetSearchHelpers` | `packages/api/pkg/sikuli/multi_target_test.go` | ✅ |
| OCR collection surface | collect words/lines | `packages/api/pkg/sikuli/parity_contract_test.go` -> `TestAPIParityContracts/OCRCollectionSurface` | `packages/api/pkg/sikuli/live_runtime_ocr_test.go`, `packages/api/pkg/sikuli/scaffolding_test.go`, `packages/api/internal/ocr/*_test.go` | ✅ |
| App and window surface | focused window, window metadata, stable selection | `packages/api/pkg/sikuli/parity_contract_test.go` -> `TestAPIParityContracts/AppWindowSurface` | `packages/api/internal/app/*_test.go`, `packages/api/internal/grpcv1/rpc_surface_integration_test.go`, `packages/api/pkg/sikuli/scaffolding_test.go` | ✅ |

## Gate

API parity is considered protected when:

- `./scripts/check-parity-gates.sh` passes.
- `docs/reference/parity/api-parity-status.md` is regenerated and committed.
- `docs/reference/parity/api-migration-examples.md` remains aligned with the closed API parity areas.
- CI blocks merges when the parity gate fails.
