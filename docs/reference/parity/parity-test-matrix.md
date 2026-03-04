# Parity Test Matrix

This matrix links parity expectations to concrete test coverage.

| Area | Java Concept | SikuliGO Surface | Test Location | Status |
|---|---|---|---|---|
| Pattern similarity/exact | `Pattern.similar/exact` | `pkg/sikuli.Pattern` + client Pattern | `packages/api/pkg/sikuli/*_test.go` | ✅ |
| Finder search/wait | `find/exists/wait` | `Finder`, `Region`, gRPC screen RPCs | `packages/api/internal/grpcv1/*_test.go` | ✅ |
| OCR read/find text | OCR APIs | `ReadText`, `FindText` | `packages/api/internal/ocr/*_test.go`, `packages/api/pkg/sikuli/*_test.go` | ✅ |
| Input automation | click/type/hotkey | `InputController`, gRPC input RPCs | `packages/api/internal/input/*_test.go`, `packages/api/internal/grpcv1/server_test.go` | ✅ |
| Observe events | appear/vanish/change | `ObserverController`, observe RPCs | `packages/api/internal/observe/*_test.go`, `packages/api/internal/grpcv1/*_test.go` | ✅ |
| App control | open/focus/close/windows | `AppController`, app RPCs | `packages/api/internal/app/*_test.go`, `packages/api/internal/grpcv1/*_test.go` | ✅ |
| KeyDown/KeyUp split | modifier lifecycle | not first-class yet | pending | 🟡 |

## Gate

Parity is considered protected when:

- mapping exists in `java-to-go-mapping.md`
- behavior is covered by tests in this matrix
- CI checks pass for API docs + parity docs freshness
