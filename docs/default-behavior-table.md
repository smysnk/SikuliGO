# Default Behavior Table

This table captures current default and protocol behavior for all existing exported GoLang objects/interfaces plus active internal matching protocols.

## Constants and global defaults

| Key | Value | Scope |
|---|---:|---|
| `DefaultSimilarity` | `0.70` | `Pattern`, `RuntimeSettings` |
| `ExactSimilarity` | `0.99` | `Pattern.Exact()` |
| `DefaultAutoWaitTimeout` | `3.0` | `Region`, `RuntimeSettings` |
| `DefaultWaitScanRate` | `3.0` | `Region`, `RuntimeSettings` |
| `DefaultObserveScanRate` | `3.0` | `Region`, `RuntimeSettings` |
| `DefaultOCRLanguage` | `"eng"` | `OCRParams`, `Finder.ReadText/FindText` |

## Object defaults

| Object | Field or behavior | Default | Notes |
|---|---|---:|---|
| `Region` | `ThrowException` | `true` | Set by `NewRegion` |
| `Region` | `AutoWaitTimeout` | `3.0` | Set by `NewRegion` |
| `Region` | `WaitScanRate` | `3.0` | Set by `NewRegion` |
| `Region` | `ObserveScanRate` | `3.0` | Set by `NewRegion` |
| `Region` | `SetThrowException(flag)` | sets exact flag value | mutates receiver |
| `Region` | `ResetThrowException()` | `true` | mutates receiver |
| `Region` | `SetAutoWaitTimeout(sec)` | negative values clamp to `0` | mutates receiver |
| `Region` | `SetWaitScanRate(rate)` | non-positive values fallback to `DefaultWaitScanRate` | mutates receiver |
| `Region` | `SetObserveScanRate(rate)` | non-positive values fallback to `DefaultObserveScanRate` | mutates receiver |
| `Location` | `NewLocation(x,y)` | exact coordinates | value type |
| `Offset` | `NewOffset(x,y)` | exact coordinates | value type |
| `Point` | `ToLocation()/ToOffset()` | exact coordinate alias conversion | non-mutating value conversion |
| `Pattern` | `similarity` | `0.70` | Set by `NewPattern` |
| `Pattern` | `resizeFactor` | `1.0` | Set by `NewPattern` |
| `Pattern` | `targetOffset` | `(0,0)` | Set by `NewPattern` |
| `Pattern` | `mask` | `nil` | No mask unless provided |
| `Image` | `Crop(rect)` | preserves absolute coordinate bounds in crop | errors if rect is empty or fully outside source |
| `Match` | `Target` | center + offset | computed in `NewMatch` |
| `Finder` | matcher backend | `NCCMatcher` | set by `NewFinder` |
| `Finder` | OCR backend | unsupported backend unless built with `-tags gosseract` | gosseract module version is pinned in `go.mod` for tagged builds |
| `Finder` | `last` cache | `nil` | populated after find operations |
| `Finder` | `Exists(pattern)` | `(Match{}, false, nil)` on missing targets | does not return `ErrFindFailed` for misses |
| `Finder` | `Has(pattern)` | `false` on missing targets | forwards non-find errors |
| `Finder` | `Wait(pattern, timeout)` | timeout `<= 0` performs one-shot then `ErrTimeout` if missing | timeout `> 0` polls using global wait scan rate |
| `Finder` | `WaitVanish(pattern, timeout)` | timeout `<= 0` performs one-shot vanish check | timeout `> 0` polls until vanished or timeout |
| `Finder` | `FindAllByRow(pattern)` | sorts by row/column then reindexes | updates `last` cache |
| `Finder` | `FindAllByColumn(pattern)` | sorts by column/row then reindexes | updates `last` cache |
| `Finder` | `Count(pattern)` | returns `len(FindAll(pattern))` | updates `last` cache via `FindAll` |
| `Finder` | `ReadText(params)` | returns trimmed OCR text | wraps `core.ErrOCRUnsupported` as `ErrBackendUnsupported` |
| `Finder` | `FindText(query, params)` | case-insensitive by default | returns `ErrFindFailed` when no matching text is detected |
| `InputController` | input backend | concrete `darwin`/`linux`/`windows` backends with non-target fallback unsupported | set by `NewInputController` |
| `InputController` | `MoveMouse(x, y, opts)` | normalizes delay to non-negative duration | delegates to `core.Input` |
| `InputController` | `Click(x, y, opts)` | default button is `left` | delegates to `core.Input` |
| `InputController` | `TypeText(text, opts)` | trims text and rejects empty values | returns `ErrInvalidTarget` on empty text |
| `InputController` | `Hotkey(keys...)` | requires at least one non-empty key | returns `ErrInvalidTarget` on invalid keys |
| `ObserverController` | observe backend | deterministic polling backend | set by `NewObserverController` |
| `ObserverController` | `ObserveAppear/ObserveVanish` | requires non-empty region and non-nil pattern | delegates to `core.Observer` |
| `ObserverController` | `ObserveChange` | pattern not required | delegates to `core.Observer` |
| `AppController` | app backend | concrete `darwin`/`linux`/`windows` backends with non-target fallback unsupported | set by `NewAppController` |
| `AppController` | `Open/Focus/Close` | requires non-empty app name | delegates to `core.App` |
| `AppController` | `IsRunning` | returns backend running state | delegates to `core.App` |
| `AppController` | `ListWindows` | maps backend window payloads to `Window` values | delegates to `core.App` |
| `Region` | `Find(source, pattern)` | one-shot match within region crop | returns `ErrFindFailed` if not found |
| `Region` | `Exists(source, pattern, timeout)` | one-shot when timeout `<= 0` | polls using `WaitScanRate` when timeout `> 0` |
| `Region` | `Has(source, pattern, timeout)` | bool wrapper over `Exists` | forwards non-find errors |
| `Region` | `Wait(source, pattern, timeout)` | uses `AutoWaitTimeout` when timeout `<= 0` | returns `ErrTimeout` on miss |
| `Region` | `WaitVanish(source, pattern, timeout)` | one-shot when timeout `<= 0` | polls using `WaitScanRate` when timeout `> 0` |
| `Region` | `FindAll/FindAllByRow/FindAllByColumn` | region-scoped via source crop | delegates to finder helper semantics |
| `Region` | `Count(source, pattern)` | returns `len(FindAll(source, pattern))` | region-scoped via source crop |
| `Region` | `OffsetBy(offset)` | delegates to `Offset(offset.X, offset.Y)` | value-style alias helper |
| `Region` | `MoveToLocation(location)` | delegates to `MoveTo(location.X, location.Y)` | value-style alias helper |
| `Region` | `ContainsLocation(location)` | delegates to `Contains(location.ToPoint())` | alias helper for parity ergonomics |
| `Region` | `ReadText(source, params)` | region-scoped OCR over source crop | delegates to finder OCR backend |
| `Region` | `FindText(source, query, params)` | region-scoped text search over OCR words | delegates to finder OCR backend |
| `Options` | typed getters | parse from string map with default fallback | invalid parse returns provided default |
| `Options` | typed setters | store string representation in map | canonical serialization via `strconv` |
| `RuntimeSettings` | `ImageCache` | `64` | initial global value |
| `RuntimeSettings` | `ShowActions` | `false` | initial global value |
| `RuntimeSettings` | `WaitScanRate` | `3.0` | initial global value |
| `RuntimeSettings` | `ObserveScanRate` | `3.0` | initial global value |
| `RuntimeSettings` | `AutoWaitTimeout` | `3.0` | initial global value |
| `RuntimeSettings` | `MinSimilarity` | `0.70` | initial global value |
| `RuntimeSettings` | `FindFailedThrows` | `true` | initial global value |

## Setter normalization behaviors

| API | Normalization rule |
|---|---|
| `Pattern.Similar(sim)` | clamps to `[0,1]` |
| `Pattern.Exact()` | sets similarity to `ExactSimilarity` |
| `Pattern.Resize(factor)` | values `<= 0` become `1.0` |
| `Pattern.WithMask(mask)` | `nil` clears mask |
| `Pattern.WithMaskMatrix(rows)` | empty rows clear mask |
| `Region.SetSize(w,h)` | negative dimensions clamp to `0` |
| `Region.SetAutoWaitTimeout(sec)` | negative values clamp to `0` |
| `Region.SetWaitScanRate(rate)` | non-positive values fallback to `DefaultWaitScanRate` |
| `Region.SetObserveScanRate(rate)` | non-positive values fallback to `DefaultObserveScanRate` |
| `OCRParams` | empty `Language` becomes `DefaultOCRLanguage` |
| `OCRParams` | `MinConfidence` clamps to `[0,1]` |
| `OCRParams` | negative `Timeout` becomes `0` |
| `InputOptions` | invalid/empty `Button` becomes `MouseButtonLeft` |
| `InputOptions` | negative `Delay` becomes `0` |
| `ObserveOptions` | non-positive `Interval` becomes derived observe interval from settings |
| `ObserveOptions` | negative `Timeout` becomes `0` |
| `AppOptions` | negative `Timeout` becomes `0` |

## Matcher request protocol defaults

| Request field | Behavior |
|---|---|
| `SearchRequest.Haystack` | required, must be non-nil |
| `SearchRequest.Needle` | required, must be non-nil |
| `SearchRequest.Threshold` | must be in `[0,1]` |
| `SearchRequest.ResizeFactor` | must be `> 0` |
| `SearchRequest.MaxResults` | `0` means unlimited |
| `SearchRequest.Mask` | optional; if set, must match effective needle dimensions |

## OCR request protocol defaults

| Request field | Behavior |
|---|---|
| `OCRRequest.Image` | required, must be non-nil |
| `OCRRequest.Language` | required, defaults to `DefaultOCRLanguage` through `OCRParams` normalization |
| `OCRRequest.TrainingDataPath` | optional |
| `OCRRequest.MinConfidence` | must be in `[0,1]` |
| `OCRRequest.Timeout` | must be `>= 0` |

## Input request protocol defaults

| Request field | Behavior |
|---|---|
| `InputRequest.Action` | required, must be one of move/click/type/hotkey |
| `InputRequest.Delay` | must be `>= 0` |
| `InputRequest.Button` | required for click actions |
| `InputRequest.Text` | required for type actions |
| `InputRequest.Keys` | at least one key required for hotkey actions |

## Observe request protocol defaults

| Request field | Behavior |
|---|---|
| `ObserveRequest.Source` | required, must be non-nil |
| `ObserveRequest.Region` | required, must be non-empty |
| `ObserveRequest.Event` | required, must be appear/vanish/change |
| `ObserveRequest.Pattern` | required for appear/vanish; optional for change |
| `ObserveRequest.Interval` | must be `>= 0` |
| `ObserveRequest.Timeout` | must be `>= 0` |

## App request protocol defaults

| Request field | Behavior |
|---|---|
| `AppRequest.Action` | required, must be open/focus/close/is_running/list_windows |
| `AppRequest.Name` | required, must be non-empty |
| `AppRequest.Timeout` | must be `>= 0` |
| `AppRequest.Args` | optional; typically used for open action |

## Matching and ordering protocol defaults

| Behavior | Rule |
|---|---|
| Resize interpolation | nearest-neighbor (`ResizeGrayNearest`) |
| Correlation model | normalized cross-correlation |
| Score normalization | mapped to `[0,1]` |
| Match filtering | include score `>= threshold` |
| Match ordering | score desc, y asc, x asc |

## Error behavior defaults

| Condition | Error |
|---|---|
| nil/invalid source image | `ErrInvalidTarget` |
| nil/invalid pattern image | `ErrInvalidTarget` |
| empty OCR query | `ErrInvalidTarget` |
| empty input text/hotkey args | `ErrInvalidTarget` |
| empty app name / invalid observe request | `ErrInvalidTarget` |
| no results in `Finder.Find` | `ErrFindFailed` |
| no results in `Finder.FindText` | `ErrFindFailed` |
| unsupported backend path | `ErrBackendUnsupported` |
