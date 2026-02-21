# OCR Integration

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
