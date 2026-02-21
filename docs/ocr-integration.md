# OCR Integration

SikuliGO provides OCR APIs in `Finder` and `Region`:

- `Finder.ReadText(params OCRParams)`
- `Finder.FindText(query, params OCRParams)`
- `Region.ReadText(source, params OCRParams)`
- `Region.FindText(source, query, params OCRParams)`

By default, OCR is disabled at build time and these APIs return `ErrBackendUnsupported`.

## Enable gogosseract backend

SikuliGO includes an optional OCR backend adapter for:

- `github.com/smysnk/gogosseract`
- branch/tag: `wazero-1-11-lean-fork-compat`

Install and pin the backend dependency:

```bash
go get github.com/smysnk/gogosseract@wazero-1-11-lean-fork-compat
go mod tidy
```

Build or test with OCR enabled:

```bash
go test -tags gogosseract ./...
go build -tags gogosseract ./...
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
