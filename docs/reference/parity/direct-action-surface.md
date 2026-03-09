# Direct Action Surface

Phase 4 closes the main direct-action parity gap on top of the live runtime without changing the existing sikuli-go architecture.

## Scope

The live API now represents the SikuliX-style direct action vocabulary across:

- `LiveRegion`
- `Match`
- `Screen`
- `InputController`

Supported actions:

- `hover`
- `click`
- `rightClick`
- `doubleClick`
- `mouseDown`
- `mouseUp`
- `typeText`
- `paste`
- `dragDrop`
- wheel scrolling
- `keyDown`
- `keyUp`

## Primitive vs Composed

The protocol only adds primitives where stateful or platform-sensitive behavior needs durable server-side semantics.

Primitive RPCs:

- `MouseDown`
- `MouseUp`
- `PasteText`
- `KeyDown`
- `KeyUp`
- `ScrollWheel`

Composed helpers in `pkg/sikuli`:

- `Hover` -> `MoveMouse`
- `RightClick` -> `Click` with right-button input options
- `DoubleClick` -> two `Click` calls
- `DragDrop` -> `MoveMouse` -> `MouseDown` -> `MoveMouse` -> `MouseUp`

This keeps the runtime protocol small while still exposing the full direct-action vocabulary at the public API layer.

## Match Behavior

`Match` direct actions now route through the live runtime instead of the local `InputController`.

That change is important for parity:

- a live `Match` keeps using its resolved match target point
- actions execute on the same remote/API-backed desktop context as the screen search that produced the match
- image-only matches still return `ErrRuntimeUnavailable` for live-only actions

## Backend Notes

No backend architecture was changed.

Current platform notes:

- macOS:
  - `mouseDown`/`mouseUp` are implemented via `cliclick`
  - `keyDown`/`keyUp` currently support modifier-key holds
  - wheel scrolling remains backend-unsupported
- Linux:
  - primitives map to `xdotool`
  - wheel scrolling maps to button-repeat scroll events
- Windows:
  - primitives map to PowerShell/user32-based input helpers
  - wheel scrolling is supported

## Result

The SikuliX direct-action vocabulary is now represented at the API level without changing the current separation between:

- image-backed local search APIs
- live runtime/server-backed screen APIs
- low-level platform backends
