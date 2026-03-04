# Backend Capability Matrix

This matrix tracks concrete backend implementation status by subsystem and platform.

| Subsystem | macOS (`darwin`) | Linux | Windows | Notes |
|---|---|---|---|---|
| Matcher (`core.Matcher`) | ✅ | ✅ | ✅ | Pure Go backend(s) via `internal/cv` |
| OCR (`core.OCR`, `-tags gosseract`) | ✅ | ✅ | ✅ | Pinned `gosseract` module version through `go.mod`; requires Tesseract runtime libs |
| Input (`core.Input`) | ✅ | ✅ | ✅ | Concrete backends in `internal/input/backend_darwin.go`, `internal/input/backend_linux.go`, `internal/input/backend_windows.go` |
| Observe (`core.Observer`) | ✅ | ✅ | ✅ | Deterministic polling backend in `internal/observe` |
| App (`core.App`) | ✅ | ✅ | ✅ | Concrete backends in `internal/app/backend_darwin.go`, `internal/app/backend_linux.go`, `internal/app/backend_windows.go` |

Legend:
- `✅`: concrete backend implemented
- `⚪`: protocol scaffold present, concrete backend pending
