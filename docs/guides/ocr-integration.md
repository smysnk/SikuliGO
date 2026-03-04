# OCR

SikuliGO provides OCR APIs in `Finder` and `Region`:

- `Finder.ReadText(params OCRParams)`
- `Finder.FindText(query, params OCRParams)`
- `Region.ReadText(source, params OCRParams)`
- `Region.FindText(source, query, params OCRParams)`

By default, OCR is disabled at build time and these APIs return `ErrBackendUnsupported`.

## Enable gosseract backend

SikuliGO includes an optional OCR backend adapter for:

- module path: `github.com/otiai10/gosseract/v2`
- pinned module version in `go.mod`

Native runtime requirements:

- Tesseract OCR and Leptonica shared libraries available on the host
- language training data installed for the selected OCR language (for example, `eng`)

Build or test with OCR enabled:

```bash
go test -tags gosseract ./...
go build -tags gosseract ./...
```

## macOS setup (Homebrew)

Install native dependencies:

```bash
brew install leptonica tesseract pkg-config
```

Export build/link flags (Apple Silicon/Homebrew):

```bash
export HOMEBREW_PREFIX="$(brew --prefix)"
export PKG_CONFIG_PATH="$HOMEBREW_PREFIX/lib/pkgconfig:$PKG_CONFIG_PATH"
export CGO_CFLAGS="-I$HOMEBREW_PREFIX/include"
export CGO_CPPFLAGS="-I$HOMEBREW_PREFIX/include"
export CGO_LDFLAGS="-L$HOMEBREW_PREFIX/lib -llept -ltesseract"
```

Validate package config and run tests:

```bash
pkg-config --cflags lept tesseract
pkg-config --libs lept tesseract
go clean -cache -testcache
go test -tags gosseract ./internal/ocr ./pkg/sikuli
```

If you see `ld: library 'lept' not found`, create compatibility symlinks:

```bash
sudo ln -sf "$(brew --prefix leptonica)/lib/libleptonica.dylib" /opt/homebrew/lib/liblept.dylib
sudo ln -sf "$(brew --prefix leptonica)/lib/libleptonica.dylib" /usr/local/lib/liblept.dylib
```

## macOS troubleshooting: `fatal error: 'leptonica/allheaders.h' file not found`

This means the compiler cannot find Leptonica headers while building the `gosseract` CGO path.

1. Verify dependencies are installed:

```bash
brew install leptonica tesseract pkg-config
```

2. Confirm where `allheaders.h` is located:

```bash
find /opt/homebrew /usr/local -name "allheaders.h" 2>/dev/null
```

3. Prefer `pkg-config` + Homebrew prefix exports (portable across Intel/Apple Silicon):

```bash
export HOMEBREW_PREFIX="$(brew --prefix)"
export PKG_CONFIG_PATH="$HOMEBREW_PREFIX/lib/pkgconfig:$PKG_CONFIG_PATH"
export CGO_CFLAGS="-I$HOMEBREW_PREFIX/include"
export CGO_CPPFLAGS="-I$HOMEBREW_PREFIX/include"
export CGO_CXXFLAGS="-I$HOMEBREW_PREFIX/include"
export CGO_LDFLAGS="-L$HOMEBREW_PREFIX/lib -llept -ltesseract"
```

4. If your headers/libs are versioned under Cellar and still not found, export explicit paths:

```bash
export LEPT_PREFIX="$(brew --prefix leptonica)"
export TESS_PREFIX="$(brew --prefix tesseract)"
export CGO_CFLAGS="-I$LEPT_PREFIX/include -I$TESS_PREFIX/include"
export CGO_CPPFLAGS="-I$LEPT_PREFIX/include -I$TESS_PREFIX/include"
export CGO_CXXFLAGS="-I$LEPT_PREFIX/include -I$TESS_PREFIX/include"
export CGO_LDFLAGS="-L$LEPT_PREFIX/lib -L$TESS_PREFIX/lib -llept -ltesseract"
```

5. Validate and retry:

```bash
pkg-config --cflags lept tesseract
pkg-config --libs lept tesseract
go clean -cache -testcache
go test -tags gosseract ./internal/ocr ./pkg/sikuli
```

## OCR parameters

`OCRParams` supports:

- `Language` (default: `"eng"`)
- `TrainingDataPath` (optional tessdata path)
- `MinConfidence` (clamped to `[0,1]`)
- `Timeout` (negative values become `0`)
- `CaseSensitive` (for `FindText`)

## Example

```go
txt, err := finder.ReadText(sikuli.OCRParams{
  Language: "eng",
})

matches, err := finder.FindText("Submit", sikuli.OCRParams{
  MinConfidence: 0.6,
})
```
