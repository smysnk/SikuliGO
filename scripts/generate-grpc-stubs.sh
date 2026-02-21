#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
PROTO_FILE="sikuli/v1/sikuli.proto"

if ! command -v protoc >/dev/null 2>&1; then
  echo "Missing protoc in PATH" >&2
  exit 1
fi

if ! command -v protoc-gen-go >/dev/null 2>&1; then
  echo "Missing protoc-gen-go in PATH" >&2
  exit 1
fi

if ! command -v protoc-gen-go-grpc >/dev/null 2>&1; then
  echo "Missing protoc-gen-go-grpc in PATH" >&2
  exit 1
fi

cd "$ROOT_DIR"
protoc \
  --proto_path=proto \
  --go_out=. \
  --go_opt=module=github.com/sikulix/portgo \
  --go-grpc_out=. \
  --go-grpc_opt=module=github.com/sikulix/portgo \
  "$PROTO_FILE"
