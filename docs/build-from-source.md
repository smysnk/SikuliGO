# Build From Source

This is the single source of truth for building SikuliGO from source.

## Prerequisites

- Go `1.24+`
- Node.js `20+`
- Yarn `4+`
- Python `3.10+`
- `protoc`

## Install Workspace Dependencies

```bash
cd /path/to/SikuliX1
yarn install
```

## Build Go API Binaries

Build `sikuligo`:

```bash
cd packages/api
go build -tags "gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d" -trimpath -ldflags="-s -w" -o ../../sikuligo ./cmd/sikuligrpc
```

Build `sikuligo-monitor`:

```bash
cd packages/api
go build -tags "gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d" -trimpath -ldflags="-s -w" -o ../../sikuligo-monitor ./cmd/sikuligo-monitor
```

## Build Protocol Artifacts

```bash
./scripts/generate-grpc-stubs.sh
./scripts/clients/generate-node-stubs.sh
./scripts/clients/generate-python-stubs.sh
./scripts/clients/generate-lua-descriptor.sh
```

## Build Node Client

```bash
yarn workspace @sikuligo/sikuligo build
```

## Build Python Distributions

```bash
./scripts/clients/release-python-client.sh
```

Skip installer steps:

```bash
SKIP_INSTALL=1 ./scripts/clients/release-python-client.sh
```

## Build Everything (Convenience)

```bash
make
```

## Optional OCR-Tagged Tests

```bash
cd packages/api
go test -tags gosseract ./...
```
