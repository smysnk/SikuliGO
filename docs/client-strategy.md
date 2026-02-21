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
| Node.js | `@grpc/grpc-js` + generated stubs | generate JS/TS stubs from `.proto` | High |
| Lua | direct gRPC runtime where available, or HTTP gateway fallback | direct bindings or gateway client wrapper | Medium |

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

Lua has more runtime variance; choose one of two paths:

1. Direct gRPC client runtime:
- use when your target runtime has stable HTTP/2 + protobuf gRPC support.
- keep generated/request models aligned with `v1` schema.

2. Gateway wrapper path:
- expose HTTP/JSON in front of gRPC and call from Lua with a standard HTTP client.
- recommended when direct gRPC support is weak in your Lua runtime.

For initial delivery, gateway wrapper is usually lower risk unless direct gRPC is already proven in your runtime.

## Implementation Phases

### Phase 1: Shared contract and tooling

- Keep one source contract: `proto/sikuli/v1/sikuli.proto`.
- Add language generation scripts under `scripts/clients/`.
- Pin generator/runtime versions and verify generated drift in CI.

### Phase 2: Python client

- Generate Python stubs into `clients/python/generated/`.
- Add wrapper in `clients/python/sikuligo_client/` for deadlines, metadata/auth, and error mapping.
- Add runnable examples in `clients/python/examples/` (`find.py`, `read_text.py`, `click_and_type.py`, `app_control.py`).
- Add CI smoke tests against local `cmd/sikuligrpc`.

### Phase 3: Node.js client

- Generate JS/TS stubs into `clients/node/generated/`.
- Add Promise-based wrapper in `clients/node/src/`.
- Add runnable examples in `clients/node/examples/` (`find.ts`, `ocr.ts`, `input.ts`, `app.ts`).
- Add CI smoke tests against local `cmd/sikuligrpc`.

### Phase 4: Lua client path

- Choose direct gRPC runtime vs HTTP/JSON gateway fallback per target runtime constraints.
- Implement a thin SDK in `clients/lua/` with the same high-level methods.
- Add runnable examples in `clients/lua/examples/`.
- Add runtime-appropriate CI smoke tests.

### Phase 5: Documentation and distribution

- Publish language quickstarts and API usage docs.
- Document required env vars (`SIKULI_GRPC_ADDR`, auth values when enabled).
- Package and version each client with release notes.

### Phase 6: Hardening and operations

- Standardize retries, timeout policies, and error mapping across clients.
- Add auth, tracing, and metrics guidance for production usage.
- Gate releases on cross-language integration checks.

## Client Deliverables

For each client, ship:

- generated stubs
- minimal high-level wrapper
- auth + timeout defaults
- one smoke test against staging
- short usage example
