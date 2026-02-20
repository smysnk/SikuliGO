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
| `Pattern` | `similarity` | `0.70` | Set by `NewPattern` |
| `Pattern` | `resizeFactor` | `1.0` | Set by `NewPattern` |
| `Pattern` | `targetOffset` | `(0,0)` | Set by `NewPattern` |
| `Pattern` | `mask` | `nil` | No mask unless provided |
| `Image` | `Crop(rect)` | preserves absolute coordinate bounds in crop | errors if rect is empty or fully outside source |
| `Match` | `Target` | center + offset | computed in `NewMatch` |
| `Finder` | matcher backend | `NCCMatcher` | set by `NewFinder` |
| `Finder` | `last` cache | `nil` | populated after find operations |
| `Finder` | `Exists(pattern)` | `(Match{}, false, nil)` on missing targets | does not return `ErrFindFailed` for misses |
| `Finder` | `Has(pattern)` | `false` on missing targets | forwards non-find errors |
| `Finder` | `Wait(pattern, timeout)` | timeout `<= 0` performs one-shot then `ErrTimeout` if missing | timeout `> 0` polls using global wait scan rate |
| `Finder` | `WaitVanish(pattern, timeout)` | timeout `<= 0` performs one-shot vanish check | timeout `> 0` polls until vanished or timeout |
| `Region` | `Find(source, pattern)` | one-shot match within region crop | returns `ErrFindFailed` if not found |
| `Region` | `Exists(source, pattern, timeout)` | one-shot when timeout `<= 0` | polls using `WaitScanRate` when timeout `> 0` |
| `Region` | `Has(source, pattern, timeout)` | bool wrapper over `Exists` | forwards non-find errors |
| `Region` | `Wait(source, pattern, timeout)` | uses `AutoWaitTimeout` when timeout `<= 0` | returns `ErrTimeout` on miss |
| `Region` | `WaitVanish(source, pattern, timeout)` | one-shot when timeout `<= 0` | polls using `WaitScanRate` when timeout `> 0` |
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

## Matcher request protocol defaults

| Request field | Behavior |
|---|---|
| `SearchRequest.Haystack` | required, must be non-nil |
| `SearchRequest.Needle` | required, must be non-nil |
| `SearchRequest.Threshold` | must be in `[0,1]` |
| `SearchRequest.ResizeFactor` | must be `> 0` |
| `SearchRequest.MaxResults` | `0` means unlimited |
| `SearchRequest.Mask` | optional; if set, must match effective needle dimensions |

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
| no results in `Finder.Find` | `ErrFindFailed` |
| unsupported backend path | `ErrBackendUnsupported` |
