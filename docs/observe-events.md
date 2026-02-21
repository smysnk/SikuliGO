# Observe Events

SikuliGO includes a baseline observe/event scaffold through `ObserverController`.

## Public API

- `NewObserverController()`
- `ObserveAppear(source, region, pattern, opts)`
- `ObserveVanish(source, region, pattern, opts)`
- `ObserveChange(source, region, opts)`

## Request protocol

Observe actions flow through `core.ObserveRequest` with strict validation:

- source image is required
- region must be non-empty and intersect source bounds
- event type must be `appear`, `vanish`, or `change`
- pattern is required for `appear` and `vanish`
- interval and timeout must be non-negative

## Backend behavior

The default backend is an unsupported stub and returns `ErrBackendUnsupported` through the public API.

This locks the observe protocol and deterministic tests now, while platform-specific observe backends can be added later without breaking the public surface.
