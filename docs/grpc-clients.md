# gRPC Client Options

This document outlines practical client options for SikuliGO gRPC APIs, with emphasis on Python, Lua, and Node.js.

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

## Suggested Rollout Order

1. Python client (reference client + CI integration tests)
2. Node.js client (service integration and tooling)
3. Lua client path (direct or gateway based on runtime constraints)

## Client Deliverables

For each client, ship:

- generated stubs
- minimal high-level wrapper
- auth + timeout defaults
- one smoke test against staging
- short usage example
