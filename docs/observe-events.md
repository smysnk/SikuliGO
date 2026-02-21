# Observe Events

SikuliGO provides observe/event APIs through `ObserverController` with a concrete deterministic polling backend.

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

The default backend uses interval polling against the existing matcher to evaluate appear/vanish requests and pixel-delta checks for change requests.

Conformance tests validate deterministic timing for appear, vanish, and change event emission.
