# SikuliGO Lua Client

This directory provides a Lua gRPC method client using `grpcurl` as transport.

## Prerequisites

- Lua 5.3+
- `grpcurl`
- `protoc` (for descriptor generation)
- SikuliGO gRPC server running (default `127.0.0.1:50051`)

## Setup

For source build details, see:

- [Build From Source](../../docs/build-from-source.md)

## Environment

- `SIKULI_GRPC_ADDR` (default: `127.0.0.1:50051`)
- `SIKULI_GRPC_AUTH_TOKEN` (optional; sent as `x-api-key`)
- `SIKULI_APP_NAME` (optional; used by `examples/app.lua`)

## Run Examples

```bash
cd packages/client-lua/examples
lua find.lua
lua input.lua
lua app.lua
```
