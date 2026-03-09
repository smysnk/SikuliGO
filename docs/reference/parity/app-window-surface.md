# App and Window Surface

Phase 8 expands app/window parity without changing the current app-backend architecture in `sikuli-go`.

## Public Go Surface

`packages/api/pkg/sikuli/app.go` now exposes a richer record and deterministic query helpers:

- `AppController.ListWindows(name, opts)`
- `AppController.FindWindows(name, query, opts)`
- `AppController.GetWindow(name, query, opts)`
- `AppController.FocusedWindow(name, opts)`

The returned `Window` record now includes:

- `ID`
- `App`
- `PID`
- `Title`
- `Bounds`
- `Focused`

This keeps `Window` as a returned value type with helper-driven selection. It does not introduce a separate Java-style mutable window object hierarchy.

## Stable Query Model

`WindowQuery` supports:

- exact ID match
- exact title match
- title substring match
- focused-only filtering
- stable index selection over the filtered result set

Selection rules are deterministic:

1. the backend returns windows in platform order
2. `FindWindows` filters that list without reordering it
3. `GetWindow` applies `Index` to the filtered list
4. `FocusedWindow` is shorthand for the focused-only query

This gives ports a stable way to express SikuliX-style `window()`, `focusedWindow()`, and `allWindows()` flows without changing the app backend model.

## Protocol Surface

The gRPC API now exposes:

- `ListWindows`
- `FindWindows`
- `GetWindow`
- `GetFocusedWindow`

The RPC layer delegates to the same `AppController` query logic used by the Go package surface. There is no parallel window-selection implementation in the transport layer.

## Platform Variance

Window metadata is richer, but still platform-dependent.

### Linux

Linux uses `wmctrl -lxG` plus `_NET_ACTIVE_WINDOW` probing.

Available where the desktop environment supports those tools:

- window ID
- app/class string
- title
- geometry
- focused flag

Limitations:

- PID is not currently populated
- focused state depends on active-window support from the window manager

### macOS

macOS uses `System Events` / AppleScript window enumeration.

Available where accessibility permissions allow it:

- app name
- PID
- title
- geometry
- focused/main state when AX attributes are available

Limitations:

- stable per-window IDs are not exposed in this phase
- focus metadata depends on accessibility attributes exposed by the app

### Windows

Windows uses PowerShell with Win32 window probing.

Available where a main window handle exists:

- window handle ID
- process name
- PID
- title
- geometry
- focused flag

Limitations:

- processes without a usable main window handle fall back to partial metadata
- geometry remains best-effort for windows that do not report a standard top-level rect

## Result

Common SikuliX app/window workflows now have first-class API support in the Go port:

- list all windows
- find windows by title or ID
- select a deterministic window from multiple matches
- ask for the currently focused window

The remaining differences are about platform portability of metadata, not about missing API concepts.
