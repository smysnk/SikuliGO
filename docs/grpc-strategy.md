# gRPC Strategy

This document defines and tracks the gRPC approach for exposing the SikuliGO API surface across multiple languages.

## Goals

- Provide a versioned network API contract for SikuliGO capabilities.
- Keep protocol behavior aligned with existing `core` contracts and `pkg/sikuli` semantics.
- Support generated clients for Python, Node.js, and Lua-oriented integrations.
- Preserve additive evolution (`v1` stable, future changes via `v2`).

## Non-goals

- Replacing existing in-process Go APIs.
- Delivering every future feature in `v1` on day one.
- Introducing breaking changes to `v1` once published.

## Current Implementation

- Versioned contract at `proto/sikuli/v1/sikuli.proto`.
- Generated Go protobuf/gRPC stubs at `internal/grpcv1/pb/`.
- gRPC service adapter at `internal/grpcv1/server.go`.
- Runnable gRPC server entrypoint at `cmd/sikuligrpc/main.go`.
- Unary/stream interceptors for auth, logging, and trace IDs in `internal/grpcv1/interceptors.go`.
- Admin operational endpoints (`/healthz`, `/snapshot`, `/metrics`, `/dashboard`) in `internal/grpcv1/ops.go`.
- Stub generation/check scripts:
  - `scripts/generate-grpc-stubs.sh`
  - `scripts/check-grpc-stubs.sh`

## Proposed API Surface

Start with a single `sikuli.v1.SikuliService` and expand by adding RPCs, not by changing existing fields.

Implemented `v1` RPC set:

- Matching: `Find`, `FindAll`
- OCR: `ReadText`, `FindText`
- Input: `MoveMouse`, `Click`, `TypeText`, `Hotkey`
- Observe: `ObserveAppear`, `ObserveVanish`, `ObserveChange`
- App control: `OpenApp`, `FocusApp`, `CloseApp`, `IsAppRunning`, `ListWindows`

## Contract Layout

Repository layout:

```text
proto/
  sikuli/
    v1/
      sikuli.proto
```

Proto conventions:

- `package sikuli.v1;`
- language package options for Go/Python/Node.
- request payloads include timeout/options fields where needed for controller behavior.

## Implementation Phases

### Phase 1: Contract and generation

Status: ✅ Implemented

- `proto/sikuli/v1/sikuli.proto` is defined.
- `protoc`-based generation script is added.
- Generated Go stubs are committed under `internal/grpcv1/pb`.

### Phase 2: Go server transport

Status: ✅ Implemented

- gRPC server bootstrap and registration are implemented.
- RPC handlers map to existing `pkg/sikuli` controllers (`Finder`, `InputController`, `ObserverController`, `AppController`).
- Request validation and gRPC status-code mapping are implemented.

### Phase 3: Cross-cutting concerns

Status: ✅ Implemented (auth/logging/tracing interceptors)

- Add unary/stream interceptors for auth, logging, and tracing.
- Enforce deadlines and cancellation propagation.
- Normalize structured error detail payloads.

### Phase 4: Client integration enablement

Status: 🟡 Planned

- Generate and publish client stubs for Python and Node.js.
- Provide Lua integration path (direct gRPC where runtime supports it, or JSON gateway path).
- Add language-specific quickstart examples.

### Phase 5: Verification and rollout

Status: 🟡 In progress

- Add end-to-end tests for success, validation errors, and timeout behavior.
- Run canary rollout in staging.
- Track latency, error-rate, and per-RPC volume before broad release.
- Provide operational dashboards and endpoint health visibility.

## Local Usage

Generate stubs:

```bash
./scripts/generate-grpc-stubs.sh
```

Run server:

```bash
go run ./cmd/sikuligrpc -listen :50051 -admin-listen :8080
```

Run server with auth:

```bash
go run ./cmd/sikuligrpc -listen :50051 -auth-token "$SIKULI_GRPC_AUTH_TOKEN"
```

## Testing Requirements

- Contract tests: server output matches proto schema and semantic defaults.
- Conformance tests: behavior parity against existing in-process APIs.
- Integration tests: Python and Node smoke calls in CI.
- Compatibility tests: reject breaking proto changes in `v1`.

## Definition of Done

- `v1` proto is versioned and documented.
- Go gRPC server runs with endpoint coverage for matching, OCR, input, observe, and app control.
- Generated client stubs are validated for Python/Node and integration path is documented for Lua.
- Operational concerns (auth, tracing, metrics) are implemented via interceptors.
