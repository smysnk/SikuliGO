# SikuliGO Node.js Client

This directory contains a Node.js gRPC client wrapper for `sikuli.v1.SikuliService`.

## Prerequisites

- Node.js 20+
- npm
- SikuliGO gRPC server running (default `127.0.0.1:50051`)

## Setup

```bash
cd clients/node
npm install
npm run generate
```

## Environment

- `SIKULI_GRPC_ADDR` (default: `127.0.0.1:50051`)
- `SIKULI_GRPC_AUTH_TOKEN` (optional; sent as `x-api-key`)
- `SIKULI_APP_NAME` (optional; used by `examples/app.ts`)

## Run Examples

```bash
cd clients/node
npm run example:find
npm run example:ocr
npm run example:input
npm run example:app
```

## Build/Release Scaffold

Validate package contents and build:

```bash
./scripts/clients/release-node-client.sh
```

If dependencies are already installed in `node_modules`, skip `npm ci`:

```bash
SKIP_INSTALL=1 ./scripts/clients/release-node-client.sh
```

Publish to npm (requires `NPM_TOKEN`):

```bash
NPM_PUBLISH=1 NPM_TOKEN=... ./scripts/clients/release-node-client.sh
```
