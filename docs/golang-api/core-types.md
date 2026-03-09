---
layout: guide
title: Golang API Core Types
nav_key: golang-api
kicker: Golang API
lead: Use this page to understand the main value types and live runtime surfaces exposed by the public Go package.
---

## Runtime And Screen Types

- `Runtime` connects to a running runtime and exposes live screen operations.
- `Screen` represents a live screen returned from `Runtime.Screens()` or `Runtime.PrimaryScreen()`.
- `LiveRegion` is the runtime-backed region surface used for find, OCR, and input actions on a bounded part of the screen.

## Image And Search Types

- `Image` holds grayscale image data used for matching and image-scoped operations.
- `Pattern` wraps an image and carries search settings such as `Similar`, `Exact`, `TargetOffset`, `Resize`, and optional masking.
- `Match` carries geometry, score, target point, and runtime bindings when the result came from a live screen operation.

## Geometry And Settings Types

- `Point` and `Rect` are the base geometry values.
- `Region` is the rectangular value object used for image-scoped search helpers and runtime region definitions.
- `Options` is the typed settings map for string, int, float, and bool values.

## App And Window Types

- `AppController` exposes app lifecycle helpers such as `Open`, `Focus`, `Close`, `IsRunning`, and window queries.
- `Window` and `WindowQuery` are the basic window records and filters for app/window flows.

## How The Types Fit Together

- `Pattern` defines what to search for.
- `Runtime` and `Screen` decide where the live operation runs.
- `LiveRegion` narrows the search or action surface.
- `Match` carries the result and can become the action target for click, hover, OCR, and other live operations.

## Next Pages

- [Golang API: First Program]({{ '/golang-api/first-program' | relative_url }})
- [Golang API: Runtime And Reference]({{ '/golang-api/runtime-and-reference' | relative_url }})
- [API Reference]({{ '/reference/api/' | relative_url }})
