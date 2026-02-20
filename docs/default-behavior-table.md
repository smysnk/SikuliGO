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
| `Pattern` | `similarity` | `0.70` | Set by `NewPattern` |
| `Pattern` | `resizeFactor` | `1.0` | Set by `NewPattern` |
| `Pattern` | `targetOffset` | `(0,0)` | Set by `NewPattern` |
| `Pattern` | `mask` | `nil` | No mask unless provided |
| `Match` | `Target` | center + offset | computed in `NewMatch` |
| `Finder` | matcher backend | `NCCMatcher` | set by `NewFinder` |
| `Finder` | `last` cache | `nil` | populated after find operations |
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
