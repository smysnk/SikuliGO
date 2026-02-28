#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${THIS_DIR}/paths.sh"
source "${THIS_DIR}/npm-helpers.sh"

CLIENT_DIR="$CLIENT_NODE_DIR"
NPM_CACHE_DIR="${NPM_CONFIG_CACHE:-$ROOT_DIR/.test-results/npm-cache}"

mkdir -p "$NPM_CACHE_DIR"
export NPM_CONFIG_CACHE="$NPM_CACHE_DIR"

cd "$CLIENT_DIR"
if [[ "${SKIP_INSTALL:-0}" != "1" ]]; then
  NPM_CONFIG_OMIT= run_npm_no_workspace install --include=dev
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
  run_npm_no_workspace run build
fi

run_npm_no_workspace pack --dry-run --ignore-scripts

if [[ "${NPM_PUBLISH:-0}" == "1" ]]; then
  configure_npm_auth_token "${NPM_TOKEN:-}"
  verify_npm_auth
  check_npm_package_visibility "@sikuligo/sikuligo"
  run_npm_no_workspace publish --ignore-scripts --access public
else
  echo "Node package scaffold validated (publish skipped; set NPM_PUBLISH=1)"
fi
