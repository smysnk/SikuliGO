# SikuliGO Lua Client

This directory provides a Lua gRPC method client using `grpcurl` as transport.

## Prerequisites

- Lua 5.3+
- `grpcurl`
- `protoc` (for descriptor generation)
- SikuliGO gRPC server running (default `127.0.0.1:50051`)

## Setup

```bash
./scripts/clients/generate-lua-descriptor.sh
```

## Environment

- `SIKULI_GRPC_ADDR` (default: `127.0.0.1:50051`)
- `SIKULI_GRPC_AUTH_TOKEN` (optional; sent as `x-api-key`)
- `SIKULI_APP_NAME` (optional; used by `examples/app.lua`)

## Run Examples

```bash
cd clients/lua/examples
lua find.lua
lua input.lua
lua app.lua
```
