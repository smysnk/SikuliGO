# Client Strategy

This document defines the client delivery strategy for SikuliGO gRPC APIs across Python, Node.js, and Lua.

## Shared Client Rules

- Generate all clients from the same `proto/sikuli/v1/sikuli.proto`.
- Set per-call deadlines (do not use unbounded calls).
- Treat retries as opt-in and only for idempotent RPCs.
- Surface server error codes and details in client-friendly exceptions.

## Client Matrix

| Client | Recommended stack | Codegen approach | Maturity |
|---|---|---|---|
| Python | `grpcio`, `grpcio-tools`, `protobuf` | generate stubs directly from `.proto` | High |
| Node.js | `@grpc/grpc-js`, `@grpc/proto-loader`, generated JS/TS stubs | generate JS/TS stubs from `.proto` | High |
| Lua | `grpcurl` + generated descriptor (`protoset`) | direct gRPC method wrapper via CLI transport | Medium |

## Current Implementation Snapshot (February 21, 2026)

- Python client wrapper: `clients/python/sikuligo_client/client.py`
- Python examples: `clients/python/examples/`
- Python generator: `scripts/clients/generate-python-stubs.sh`
- Node client wrapper: `clients/node/src/client.ts`
- Node examples: `clients/node/examples/`
- Node generator: `scripts/clients/generate-node-stubs.sh`
- Lua client wrapper: `clients/lua/sikuligo_client.lua`
- Lua examples: `clients/lua/examples/`
- Lua descriptor generator: `scripts/clients/generate-lua-descriptor.sh`

## Python

Recommended for first external integration.

- Libraries: `grpcio`, `grpcio-tools`, `protobuf`.
- Generate stubs during build/release.
- Wrap generated stubs in a small typed client layer for retries, auth metadata, and deadline defaults.

Typical wrapper concerns:

- default timeout per RPC
- auth header attachment
- normalized exception mapping from gRPC status

## Node.js

Recommended for service-to-service integration and web tooling.

- Use `@grpc/grpc-js` transport.
- Use generated stubs (JS/TS) and keep call options centralized.
- Prefer async wrappers returning Promises for unary RPCs.

Typical wrapper concerns:

- centralized channel setup (TLS, creds, keepalive)
- per-RPC deadlines and retry policy
- consistent error-to-domain mapping

## Lua

Lua has more runtime variance; this implementation uses direct gRPC method invocation via `grpcurl`.

- generate a descriptor set (`protoset`) from `proto/sikuli/v1/sikuli.proto`.
- call `sikuli.v1.SikuliService/<Method>` from Lua through `grpcurl`.
- keep high-level Lua methods aligned with the same `v1` RPC names used by Python/Node.

## Implementation Phases

### Phase 1: Shared contract and tooling

Status: ✅ Implemented

- Keep one source contract: `proto/sikuli/v1/sikuli.proto`.
- Add language generation scripts under `scripts/clients/`.
- Pin generator/runtime versions in language-specific manifests.

### Phase 2: Python client

Status: ✅ Implemented (baseline wrapper and examples)

- Generate Python stubs into `clients/python/generated/`.
- Add wrapper in `clients/python/sikuligo_client/` for deadlines, metadata/auth, and error mapping.
- Add runnable examples in `clients/python/examples/` (`find.py`, `read_text.py`, `click_and_type.py`, `app_control.py`).
- Add CI smoke tests against local `cmd/sikuligrpc` (next step).

### Phase 3: Node.js client

Status: ✅ Implemented (baseline wrapper and examples)

- Generate JS/TS stubs into `clients/node/generated/`.
- Add Promise-based wrapper in `clients/node/src/`.
- Add runnable examples in `clients/node/examples/` (`find.ts`, `ocr.ts`, `input.ts`, `app.ts`).
- Add CI smoke tests against local `cmd/sikuligrpc` (next step).

### Phase 4: Lua client path

Status: ✅ Implemented (grpcurl method path)

- Use `grpcurl` transport with generated descriptor set for direct gRPC method calls.
- Implement a thin SDK in `clients/lua/` with the same high-level methods.
- Add runnable examples in `clients/lua/examples/`.
- Add runtime-appropriate CI smoke tests (next step).

### Phase 5: Documentation and distribution

Status: 🟡 In progress

- Publish language quickstarts and API usage docs.
- Document required env vars (`SIKULI_GRPC_ADDR`, auth values when enabled).
- Package and version each client with release notes.
  - Release scaffolding added:
  - Single-command version bump: `./scripts/clients/set-version.sh <X.Y.Z>`
  - Build-number versioning: `./scripts/clients/set-version-from-build.sh`
  - Python package metadata: `clients/python/pyproject.toml`
  - Node package metadata: `clients/node/package.json`
  - Node binary package metadata: `clients/node/packages/bin-*/package.json`
  - Manual release workflow: `.github/workflows/client-release.yml`
  - Release helpers: `scripts/clients/release-python-client.sh`, `scripts/clients/release-node-client.sh`, `scripts/clients/release-node-binaries.sh`
  - Protected-branch pushes auto-trigger publish flows through `.github/workflows/client-release.yml`

### Phase 6: Hardening and operations

Status: 🟡 Planned

- Standardize retries, timeout policies, and error mapping across clients.
- Add auth, tracing, and metrics guidance for production usage.
- Gate releases on cross-language integration checks.

## Client Deliverables

For each client, ship:

- generated stubs
- minimal high-level wrapper
- auth + timeout defaults
- one smoke test against staging/CI runtime
- short usage example

## Release Versioning Controls

`client-release.yml` computes client versions from CI build metadata before publishing.

- `BUILD_NUMBER`: defaults to GitHub `run_number`.
- `VERSION_MAJOR`: optional override for major (defaults to current major).
- `VERSION_MINOR`: optional override for minor (defaults to current minor).
- `PATCH_MODE`: `build` or `fixed-minus-build`.
- `PATCH_FIXED`: required only when `PATCH_MODE=fixed-minus-build`.

Examples:

- `PATCH_MODE=build` with build `412` and `VERSION_MAJOR=0`, `VERSION_MINOR=2` -> `0.2.412`
- `PATCH_MODE=fixed-minus-build`, `PATCH_FIXED=10000`, build `412`, `VERSION_MAJOR=1`, `VERSION_MINOR=0` -> `1.0.9588`
