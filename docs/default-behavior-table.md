# Default Behavior Table (Workstream 1)

| Setting | Default | Notes |
|---|---:|---|
| `Pattern.similarity` | `0.70` | Parity baseline for image find thresholds |
| `Pattern.exact()` similarity | `0.99` | Exact-like behavior |
| `Pattern.resizeFactor` | `1.0` | No resize unless explicitly set |
| `Pattern.targetOffset` | `(0,0)` | Match center by default |
| `Region.ThrowException` | `true` | Missing find may return `ErrFindFailed` |
| `Region.AutoWaitTimeout` | `3.0` sec | Baseline wait timeout |
| `Region.WaitScanRate` | `3.0` Hz | Wait poll cadence |
| `Region.ObserveScanRate` | `3.0` Hz | Observe poll cadence |
| `RuntimeSettings.ImageCache` | `64` | Initial image cache size |
| `RuntimeSettings.ShowActions` | `false` | Debug action drawing disabled |
| `RuntimeSettings.MinSimilarity` | `0.70` | Mirrors finder baseline threshold |

## Error defaults

| Condition | Error |
|---|---|
| Missing source image / pattern | `ErrInvalidTarget` |
| No match in `Finder.Find` | `ErrFindFailed` |
| Unsupported backend feature | `ErrBackendUnsupported` |

