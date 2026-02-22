# OpenCV Integration

SikuliGO now includes an OpenCV matcher backend for image search.

## Backend Selection

- default build (`go build ./...`): uses pure Go matcher backend.
- OpenCV build (`-tags opencv`): uses OpenCV matcher backend by default.

The OpenCV backend is wired through:

- `internal/cv/opencv_matcher_opencv.go`
- `internal/cv/default_matcher_opencv.go`
- `pkg/sikuli/finder.go`
- `internal/observe/backend_polling.go`

## Build with OpenCV

```bash
go test -tags opencv ./...
go build -tags opencv ./...
```

## Native Requirements

- OpenCV libraries and headers installed on the host
- CGO enabled

On macOS (Homebrew), a typical setup is:

```bash
brew install opencv pkg-config
export HOMEBREW_PREFIX="$(brew --prefix)"
export PKG_CONFIG_PATH="$HOMEBREW_PREFIX/lib/pkgconfig:$PKG_CONFIG_PATH"
```

## Notes

- When `-tags opencv` is enabled, `Finder` and observe polling automatically use the OpenCV matcher backend.
- Existing pure-Go matchers (`NCC`, `SAD`) remain available.
