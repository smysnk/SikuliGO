# App Control

SikuliGO includes a baseline app/window/process scaffold through `AppController`.

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

The default backend is an unsupported stub and returns `ErrBackendUnsupported` through the public API.

This locks app/window contracts now, while cross-platform app-control backends can be added later without breaking the public surface.
