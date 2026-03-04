# Input Automation

SikuliGO provides input automation APIs through `InputController`.

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

- `darwin` builds use a concrete backend for move/click/type/hotkey dispatch.
- `linux` builds use a concrete backend for move/click/type/hotkey dispatch via `xdotool`.
- `windows` builds use a concrete backend for move/click/type/hotkey dispatch via PowerShell.
- non-target builds (`!darwin && !linux && !windows`) use an unsupported fallback backend and return `ErrBackendUnsupported` through the public API.

This keeps the protocol stable while allowing platform-specific backend implementations behind `core.Input`.
