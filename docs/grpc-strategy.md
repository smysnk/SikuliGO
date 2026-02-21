# gRPC Strategy

This document defines a practical plan to expose the SikuliGO API surface through gRPC and support generated clients across multiple languages.

## Goals

- Provide a versioned network API contract for SikuliGO capabilities.
- Keep protocol behavior aligned with existing `core` contracts and `pkg/sikuli` semantics.
- Support generated clients for Python, Node.js, and Lua-oriented integrations.
- Preserve additive evolution (`v1` stable, future changes via `v2`).

## Non-goals

- Replacing existing in-process Go APIs.
- Delivering every future feature in `v1` on day one.
- Introducing breaking changes to `v1` once published.

## Proposed API Surface

Start with a single `sikuli.v1.SikuliService` and expand by adding RPCs, not by changing existing fields.

Recommended first RPC set:

- Matching: `Find`, `FindAll`
- OCR: `ReadText`, `FindText`
- Input: `MoveMouse`, `Click`, `TypeText`, `Hotkey`
- Observe: `ObserveAppear`, `ObserveVanish`, `ObserveChange` (server-streaming where applicable)
- App control: `OpenApp`, `FocusApp`, `CloseApp`, `IsAppRunning`, `ListWindows`

## Contract Layout

Recommended repo layout:

```text
proto/
  sikuli/
    v1/
      sikuli.proto
```

Proto package conventions:

- `package sikuli.v1;`
- language package options for Go/Python/Node.
- shared request metadata fields for timeout, request id, and optional tags.

## Implementation Phases

### Phase 1: Contract and generation

- Define `proto/sikuli/v1/sikuli.proto`.
- Add codegen workflow (`buf` or `protoc`) with reproducible outputs.
- Add CI checks for proto formatting and generated-code drift.

### Phase 2: Go server transport

- Add gRPC server bootstrap and service registration.
- Map RPC handlers to existing `pkg/sikuli` controllers (`Finder`, `InputController`, `ObserverController`, `AppController`).
- Implement request validation and status-code mapping.

### Phase 3: Cross-cutting concerns

- Add unary/stream interceptors for auth, logging, and tracing.
- Enforce deadlines and cancellation propagation.
- Normalize error payloads (`code`, message, details).

### Phase 4: Client integration enablement

- Generate and publish client stubs for Python and Node.js.
- Provide Lua integration path (direct gRPC where runtime supports it, or JSON gateway path).
- Add language-specific quickstart examples.

### Phase 5: Verification and rollout

- Add end-to-end tests for success, validation errors, and timeout behavior.
- Run canary rollout in staging.
- Track latency, error-rate, and per-RPC volume before broad release.

## Testing Requirements

- Contract tests: server output matches proto schema and semantic defaults.
- Conformance tests: behavior parity against existing in-process APIs.
- Integration tests: Python and Node smoke calls in CI.
- Compatibility tests: reject breaking proto changes in `v1`.

## Definition of Done

- `v1` proto is versioned and documented.
- Go gRPC server is running with authenticated endpoints.
- Python and Node generated clients are validated in CI.
- Lua integration path is documented and tested in at least one target runtime.
- Operational metrics and tracing are available for every published RPC.
