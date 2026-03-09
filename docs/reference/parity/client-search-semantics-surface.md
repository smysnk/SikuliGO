# Client Search Semantics Surface

Phase 1 aligns the search contract for the reference clients with the API parity decisions from `docs/reference/parity/search-semantics-matrix.md`.

## Contract

For live screen search in Node and Python:

- `find`
  - success: returns a `Match`
  - miss: preserves the transport/API miss error
- `exists`
  - success with hit: returns a `Match`
  - miss or timeout budget exhaustion: returns `null` / `None`
- `wait`
  - success: returns a `Match`
  - timeout: preserves the transport/API timeout error
- `waitVanish` / `wait_vanish`
  - returns `true` when the target is absent
  - returns `false` after timeout
  - does not convert timeout into an exception

## Node

Node now treats `find` and `wait` as thin wrappers over the API contract. It no longer synthesizes local `"match not found"` or `"wait timeout"` errors when the transport already returns a stable gRPC status.

`waitVanish` is client-composed over `exists`, which is acceptable for this phase because the underlying API parity contract is already frozen.

## Python

Python now follows the same model as Node:

- `find` propagates the API miss error
- `wait` propagates the API timeout error
- `wait_vanish` is client-composed over `exists_on_screen` and returns `bool`

## Lua

Lua remains a descriptor-level `grpcurl` transport client in this phase. It does not currently expose the same object wrapper surface as Node and Python, but it also does not add independent miss/timeout reinterpretation on top of the API.

That means:

- Lua continues to reflect raw transport success/error behavior
- higher-level parity wrappers for Lua are deferred until the core object model is stabilized in the later client phases
