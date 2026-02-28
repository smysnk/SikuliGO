#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${THIS_DIR}/paths.sh"

OUT_DIR="$CLIENT_PYTHON_DIR/generated"
PROTO_FILE="sikuli/v1/sikuli.proto"

if ! command -v protoc >/dev/null 2>&1; then
  echo "Missing protoc in PATH" >&2
  exit 1
fi

mkdir -p "$OUT_DIR"

cd "$API_DIR"
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
  echo "Install with: python3 -m pip install -r packages/client-python/requirements.txt" >&2
fi

# Ensure importable Python package structure.
mkdir -p "$OUT_DIR/sikuli" "$OUT_DIR/sikuli/v1"
touch "$OUT_DIR/__init__.py" "$OUT_DIR/sikuli/__init__.py" "$OUT_DIR/sikuli/v1/__init__.py"

# grpcio-tools generates absolute imports using proto package names (e.g. `from sikuli.v1 ...`).
# Rewrite to the packaged module path so examples/imports work without PYTHONPATH tweaks.
if [[ -f "$OUT_DIR/sikuli/v1/sikuli_pb2_grpc.py" ]]; then
  perl -i -pe 's/^from sikuli\.v1 import sikuli_pb2 as /from generated.sikuli.v1 import sikuli_pb2 as /' \
    "$OUT_DIR/sikuli/v1/sikuli_pb2_grpc.py"
fi

echo "Python artifacts generated in: $OUT_DIR"
