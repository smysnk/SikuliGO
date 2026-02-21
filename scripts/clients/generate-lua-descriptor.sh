#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
OUT_DIR="$ROOT_DIR/clients/lua/generated"
OUT_FILE="$OUT_DIR/sikuli.protoset"
PROTO_FILE="sikuli/v1/sikuli.proto"

if ! command -v protoc >/dev/null 2>&1; then
  echo "Missing protoc in PATH" >&2
  exit 1
fi

mkdir -p "$OUT_DIR"

cd "$ROOT_DIR"
protoc \
  --proto_path=proto \
  --descriptor_set_out="$OUT_FILE" \
  --include_imports \
  "proto/$PROTO_FILE"

echo "Lua descriptor generated in: $OUT_FILE"
