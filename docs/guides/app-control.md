# App Control

SikuliGO provides app/window/process APIs through `AppController`.

## Public API

- `NewAppController()`
- `Open(name, args, opts)`
- `Focus(name, opts)`
- `Close(name, opts)`
- `IsRunning(name, opts)`
- `ListWindows(name, opts)`

## Request protocol

App actions flow through `core.AppRequest` with strict validation:

- action type is required
- app name is required
- timeout must be non-negative
- open action may include argument lists

## Backend behavior

- `darwin` builds use a concrete backend for open/focus/close/is-running/list-windows.
- `linux` builds use a concrete command-driven backend for open/focus/close/is-running/list-windows.
- `windows` builds use a concrete PowerShell-driven backend for open/focus/close/is-running/list-windows.
- non-target builds use an unsupported fallback backend and return `ErrBackendUnsupported` through the public API.

This keeps app/window contracts stable while enabling incremental cross-platform backend expansion.
