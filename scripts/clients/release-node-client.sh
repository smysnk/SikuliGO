#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
CLIENT_DIR="$ROOT_DIR/packages/client-node"
NPM_CACHE_DIR="${NPM_CONFIG_CACHE:-$ROOT_DIR/.test-results/npm-cache}"

mkdir -p "$NPM_CACHE_DIR"
export NPM_CONFIG_CACHE="$NPM_CACHE_DIR"

cd "$CLIENT_DIR"
if [[ "${SKIP_INSTALL:-0}" != "1" ]]; then
  NPM_CONFIG_OMIT= npm install --include=dev
fi

required_files=(
  "dist/src/index.js"
  "dist/src/index.d.ts"
  "dist/src/client.js"
  "generated/sikuli/v1/sikuli_pb.js"
  "generated/sikuli/v1/sikuli_pb.d.ts"
  "generated/sikuli/v1/sikuli_grpc_pb.js"
  "generated/sikuli/v1/sikuli_grpc_pb.d.ts"
)

missing=()
for f in "${required_files[@]}"; do
  if [[ ! -f "$f" ]]; then
    missing+=("$f")
  fi
done

if [[ "${NODE_CLIENT_BUILD:-0}" == "1" || ${#missing[@]} -gt 0 ]]; then
  if [[ ${#missing[@]} -gt 0 ]]; then
    echo "Missing Node client artifacts: ${missing[*]}"
    echo "Attempting npm run build to regenerate dist/generated artifacts..."
  fi
  npm run build
fi

run_npm_without_workspace_flags() {
  env \
    -u NPM_CONFIG_WORKSPACE \
    -u npm_config_workspace \
    -u NPM_CONFIG_WORKSPACES \
    -u npm_config_workspaces \
    npm "$@"
}

run_npm_without_workspace_flags pack --dry-run --ignore-scripts

if [[ "${NPM_PUBLISH:-0}" == "1" ]]; then
  if [[ -z "${NPM_TOKEN:-}" ]]; then
    echo "Missing NPM_TOKEN for publish" >&2
    exit 1
  fi
  npm config set //registry.npmjs.org/:_authToken="${NPM_TOKEN}"
  run_npm_without_workspace_flags publish --ignore-scripts --access public
else
  echo "Node package scaffold validated (publish skipped; set NPM_PUBLISH=1)"
fi
