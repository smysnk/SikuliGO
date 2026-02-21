#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
GEN_SCRIPT="$ROOT_DIR/scripts/generate-grpc-stubs.sh"

if [[ ! -x "$GEN_SCRIPT" ]]; then
  echo "Missing generator script: $GEN_SCRIPT" >&2
  exit 1
fi

cd "$ROOT_DIR"
"$GEN_SCRIPT"

if ! git diff --quiet -- internal/grpcv1/pb/sikuli.pb.go internal/grpcv1/pb/sikuli_grpc.pb.go; then
  echo "gRPC stubs are out of date. Run ./scripts/generate-grpc-stubs.sh and commit changes." >&2
  exit 1
fi
