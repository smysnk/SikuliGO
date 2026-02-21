# Input Automation

SikuliGO includes a baseline input automation scaffold through `InputController`.

## Public API

- `NewInputController()`
- `MoveMouse(x, y, opts)`
- `Click(x, y, opts)`
- `TypeText(text, opts)`
- `Hotkey(keys...)`

## Request protocol

Input actions flow through `core.InputRequest` with strict validation:

- action type is required
- delays must be non-negative
- click requires a button
- type requires non-empty text
- hotkey requires at least one key

## Backend behavior

The default backend is an unsupported stub and returns `ErrBackendUnsupported` through the public API.

This locks the protocol and tests now, while platform-specific backends can be added later without breaking the public surface.
