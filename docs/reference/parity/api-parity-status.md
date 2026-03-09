# API Parity Status

This document is generated from `docs/reference/parity/api-parity-status.tsv`. It tracks API-level implementation maturity independently from client wrapper maturity.

| Area | API Status | Test Location | Primary Contract Test | Migration Examples | Notes |
|---|---|---|---|---|---|
| Search semantics foundation | `closed` | `packages/api/pkg/sikuli/parity_contract_test.go` | `TestAPIParityContracts/SearchSemantics` | [Examples](/reference/parity/api-migration-examples#search-semantics-and-exceptionnull-behavior) | Image-backed and live-screen miss/wait semantics share one stable contract. |
| Exception/null semantics | `closed` | `packages/api/pkg/sikuli/parity_contract_test.go` | `TestAPIParityContracts/SearchSemantics` | [Examples](/reference/parity/api-migration-examples#search-semantics-and-exceptionnull-behavior) | The Go contract keeps explicit miss/timeout returns while preserving documented compatibility metadata. |
| Live screen and region surface | `closed` | `packages/api/pkg/sikuli/parity_contract_test.go` | `TestAPIParityContracts/LiveScreenAndRegionSurface` | [Examples](/reference/parity/api-migration-examples#live-screen-and-region-surface) | Screen discovery, region derivation, and capture are public API workflows. |
| Match as a first-class action target | `closed` | `packages/api/pkg/sikuli/parity_contract_test.go` | `TestAPIParityContracts/MatchAsActionTarget` | [Examples](/reference/parity/api-migration-examples#match-as-a-first-class-action-target) | Live matches retain stable region semantics and can drive follow-up work without manual point extraction. |
| Direct action API parity | `closed` | `packages/api/pkg/sikuli/parity_contract_test.go` | `TestAPIParityContracts/DirectActionSurface` | [Examples](/reference/parity/api-migration-examples#direct-action-surface) | Hover, click variants, drag-drop, paste, wheel, and key state are stabilized at the API layer. |
| Finder traversal and lifecycle | `closed` | `packages/api/pkg/sikuli/parity_contract_test.go` | `TestAPIParityContracts/FinderTraversalAndLifecycle` | [Examples](/reference/parity/api-migration-examples#finder-traversal-and-lifecycle) | Slice-oriented Go search remains, with additive iterator compatibility for legacy ports. |
| Multi-target search helpers | `closed` | `packages/api/pkg/sikuli/parity_contract_test.go` | `TestAPIParityContracts/MultiTargetSearchHelpers` | [Examples](/reference/parity/api-migration-examples#multi-target-search-helpers) | Best/any/wait helper families now have deterministic API-level semantics. |
| OCR collection surface | `closed` | `packages/api/pkg/sikuli/parity_contract_test.go` | `TestAPIParityContracts/OCRCollectionSurface` | [Examples](/reference/parity/api-migration-examples#ocr-collection-surface) | Word and line collection are available across image-backed and live-screen flows. |
| App and window surface | `closed` | `packages/api/pkg/sikuli/parity_contract_test.go` | `TestAPIParityContracts/AppWindowSurface` | [Examples](/reference/parity/api-migration-examples#app-and-window-surface) | Focused window lookup and stable window selection are API-backed, with platform variance documented. |

## Status Summary

- `closed`: 9

## Maintenance

- Update the status seed when API parity maturity changes.
- Run `./scripts/generate-parity-docs.sh` after updates.
- CI verifies this file is up to date.
