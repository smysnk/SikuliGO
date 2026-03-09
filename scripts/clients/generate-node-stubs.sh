#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${THIS_DIR}/paths.sh"

CLIENT_DIR="$CLIENT_NODE_DIR"
OUT_DIR="$CLIENT_DIR/generated"
PROTO_FILE="$API_DIR/proto/sikuli/v1/sikuli.proto"
NODE_BIN_CANDIDATES=(
  "$CLIENT_DIR/node_modules/.bin"
  "$ROOT_DIR/node_modules/.bin"
)
NODE_BIN=""
REQUIRED_GENERATED=(
  "$OUT_DIR/sikuli/v1/sikuli_pb.js"
  "$OUT_DIR/sikuli/v1/sikuli_pb.d.ts"
  "$OUT_DIR/sikuli/v1/sikuli_grpc_pb.js"
  "$OUT_DIR/sikuli/v1/sikuli_grpc_pb.d.ts"
)

has_generated_artifacts() {
  for f in "${REQUIRED_GENERATED[@]}"; do
    if [[ ! -f "$f" ]]; then
      return 1
    fi
  done
  return 0
}

missing_tools=()
for candidate in "${NODE_BIN_CANDIDATES[@]}"; do
  if [[ -x "$candidate/grpc_tools_node_protoc" && -x "$candidate/grpc_tools_node_protoc_plugin" && -x "$candidate/protoc-gen-ts" ]]; then
    NODE_BIN="$candidate"
    break
  fi
done

for tool in grpc_tools_node_protoc grpc_tools_node_protoc_plugin protoc-gen-ts; do
  if [[ -z "$NODE_BIN" || ! -x "$NODE_BIN/$tool" ]]; then
    missing_tools+=("$tool")
  fi
done

if [[ ${#missing_tools[@]} -gt 0 ]]; then
  if has_generated_artifacts; then
    echo "Missing Node protobuf tooling (${missing_tools[*]}), reusing committed generated artifacts in $OUT_DIR" >&2
    exit 0
  fi
  echo "Missing Node protobuf tooling: ${missing_tools[*]}. Run: (cd $CLIENT_DIR && npm install)" >&2
  echo "No generated artifacts found at $OUT_DIR, cannot continue." >&2
  exit 1
fi

mkdir -p "$OUT_DIR"

"$NODE_BIN/grpc_tools_node_protoc" \
  --proto_path="$API_DIR/proto" \
  --js_out=import_style=commonjs,binary:"$OUT_DIR" \
  --grpc_out=grpc_js:"$OUT_DIR" \
  --plugin=protoc-gen-grpc="$NODE_BIN/grpc_tools_node_protoc_plugin" \
  "$PROTO_FILE"

"$NODE_BIN/grpc_tools_node_protoc" \
  --proto_path="$API_DIR/proto" \
  --plugin=protoc-gen-ts="$NODE_BIN/protoc-gen-ts" \
  --ts_out=grpc_js:"$OUT_DIR" \
  "$PROTO_FILE"

echo "Node artifacts generated in: $OUT_DIR"
