---
layout: guide
title: Golang API Runtime And Reference
nav_key: golang-api
kicker: Golang API
lead: Connect the live runtime concepts to the generated API reference and the deeper integration guides.
---

## Runtime Connection

The public runtime entry point is:

```go
runtime, err := sikuli.NewRuntime("127.0.0.1:50051")
```

If the address is empty, the runtime code falls back to `SIKULI_GRPC_ADDR` and then to the default `127.0.0.1:50051`.

Useful runtime options include:

- `WithRuntimeAuthToken`
- `WithRuntimeRPCTimeout`
- `WithRuntimeDialTimeout`
- `WithRuntimeMatcherEngine`
- `WithRuntimeConn`
- `WithRuntimeContextDialer`

## Runtime Surface

From `Runtime`, the main live entry points are:

- `Screens()`
- `PrimaryScreen()`
- `Screen(id)`
- `Capture()`
- `CaptureRegion(region)`
- `Region(region)`

From live `Screen` and `LiveRegion`, the main operations include:

- `Find`
- `Exists`
- `Wait`
- `WaitVanish`
- `ReadText`
- `FindText`
- `CollectWords`
- `CollectLines`
- `Click`, `Hover`, `RightClick`, `DoubleClick`

## Generated Reference

Use the generated reference when you need exported symbol detail rather than a guide summary:

- [API Reference]({{ '/reference/api/' | relative_url }})
- [pkg/sikuli]({{ '/reference/api/pkg-sikuli' | relative_url }})

The generated pages are script-owned and should not be hand-edited.

## Related Guides

- [Build From Source]({{ '/guides/build-from-source' | relative_url }})
- [OCR Integration]({{ '/guides/ocr-integration' | relative_url }})
- [Input Automation]({{ '/guides/input-automation' | relative_url }})
- [App Control]({{ '/guides/app-control' | relative_url }})

## Next Pages

- [Golang API: Core Types]({{ '/golang-api/core-types' | relative_url }})
- [Getting Help: FAQ]({{ '/getting-help/faq' | relative_url }})
