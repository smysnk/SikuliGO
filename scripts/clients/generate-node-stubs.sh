#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
CLIENT_DIR="$ROOT_DIR/clients/node"
OUT_DIR="$CLIENT_DIR/generated"
PROTO_FILE="$ROOT_DIR/proto/sikuli/v1/sikuli.proto"
NODE_BIN="$CLIENT_DIR/node_modules/.bin"

if [[ ! -x "$NODE_BIN/grpc_tools_node_protoc" ]]; then
  echo "Missing grpc_tools_node_protoc. Run: (cd $CLIENT_DIR && npm install)" >&2
  exit 1
fi

if [[ ! -x "$NODE_BIN/grpc_tools_node_protoc_plugin" ]]; then
  echo "Missing grpc_tools_node_protoc_plugin. Run: (cd $CLIENT_DIR && npm install)" >&2
  exit 1
fi

if [[ ! -x "$NODE_BIN/protoc-gen-ts" ]]; then
  echo "Missing protoc-gen-ts. Run: (cd $CLIENT_DIR && npm install)" >&2
  exit 1
fi

mkdir -p "$OUT_DIR"

"$NODE_BIN/grpc_tools_node_protoc" \
  --proto_path="$ROOT_DIR/proto" \
  --js_out=import_style=commonjs,binary:"$OUT_DIR" \
  --grpc_out=grpc_js:"$OUT_DIR" \
  --plugin=protoc-gen-grpc="$NODE_BIN/grpc_tools_node_protoc_plugin" \
  "$PROTO_FILE"

"$NODE_BIN/grpc_tools_node_protoc" \
  --proto_path="$ROOT_DIR/proto" \
  --plugin=protoc-gen-ts="$NODE_BIN/protoc-gen-ts" \
  --ts_out=grpc_js:"$OUT_DIR" \
  "$PROTO_FILE"

echo "Node artifacts generated in: $OUT_DIR"
