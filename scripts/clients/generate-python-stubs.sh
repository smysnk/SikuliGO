#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
OUT_DIR="$ROOT_DIR/clients/python/generated"
PROTO_FILE="sikuli/v1/sikuli.proto"

if ! command -v protoc >/dev/null 2>&1; then
  echo "Missing protoc in PATH" >&2
  exit 1
fi

mkdir -p "$OUT_DIR"

cd "$ROOT_DIR"
if command -v grpc_python_plugin >/dev/null 2>&1; then
  protoc \
    --proto_path=proto \
    --python_out="$OUT_DIR" \
    --grpc_out="$OUT_DIR" \
    --plugin=protoc-gen-grpc="$(command -v grpc_python_plugin)" \
    "$PROTO_FILE"
elif python3 - <<'PY' >/dev/null 2>&1
import importlib.util
raise SystemExit(0 if importlib.util.find_spec("grpc_tools.protoc") else 1)
PY
then
  python3 -m grpc_tools.protoc \
    -I proto \
    --python_out="$OUT_DIR" \
    --grpc_python_out="$OUT_DIR" \
    "proto/$PROTO_FILE"
else
  protoc \
    --proto_path=proto \
    --python_out="$OUT_DIR" \
    "$PROTO_FILE"
  echo "Missing grpc_python_plugin and grpcio-tools; generated protobuf messages only (no *_pb2_grpc.py)." >&2
  echo "Install with: python3 -m pip install -r clients/python/requirements.txt" >&2
fi

# Ensure importable Python package structure.
mkdir -p "$OUT_DIR/sikuli" "$OUT_DIR/sikuli/v1"
touch "$OUT_DIR/__init__.py" "$OUT_DIR/sikuli/__init__.py" "$OUT_DIR/sikuli/v1/__init__.py"

echo "Python artifacts generated in: $OUT_DIR"
