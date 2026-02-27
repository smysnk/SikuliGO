#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
API_DIR="$ROOT_DIR/packages/api"
OUT_DIR="$ROOT_DIR/packages/client-lua/generated"
OUT_FILE="$OUT_DIR/sikuli.protoset"
PROTO_FILE="sikuli/v1/sikuli.proto"

if ! command -v protoc >/dev/null 2>&1; then
  echo "Missing protoc in PATH" >&2
  exit 1
fi

mkdir -p "$OUT_DIR"

cd "$API_DIR"
protoc \
  --proto_path=proto \
  --descriptor_set_out="$OUT_FILE" \
  --include_imports \
  "proto/$PROTO_FILE"

echo "Lua descriptor generated in: $OUT_FILE"
