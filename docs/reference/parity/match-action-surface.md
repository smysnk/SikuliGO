# Match Action Surface

Phase 3 adds a first-class `Match` surface in `packages/api/pkg/sikuli` without changing the existing runtime architecture.

Delivered behavior:

- `Match` still retains its value semantics and exported data:
  - `Rect`
  - `Score`
  - `Target`
  - `Index`
- Live matches returned from the Phase 2 runtime surface now carry runtime binding internally.
- Plain image-backed matches remain plain values and do not pretend to be live desktop targets.

New `Match` methods:

- region-like geometry:
  - `Bounds() Region`
  - `Region() Region`
  - `Center() Point`
  - `TargetPoint() Point`
  - `Live() bool`
- live region delegation:
  - `Capture()`
  - `Find(...)`
  - `Exists(...)`
  - `Has(...)`
  - `Wait(...)`
  - `WaitVanish(...)`
- direct action chaining on live matches:
  - `MoveMouse(...)`
  - `Hover(...)`
  - `Click(...)`
  - `RightClick(...)`
  - `DoubleClick(...)`
  - `TypeText(...)`

Runtime rules:

- Matches returned from `Screen.Find`, `Screen.Exists`, `Screen.Wait`, and `LiveRegion` equivalents are live-bound.
- Matches returned from image-backed `Finder` / `Region` flows are not live-bound.
- Live-only operations on a plain match return `ErrRuntimeUnavailable`.

Scope boundary:

- This phase does not attempt to close the full SikuliX direct-action vocabulary gap.
- Broader verb parity for region/screen/action APIs remains Phase 4 work.
