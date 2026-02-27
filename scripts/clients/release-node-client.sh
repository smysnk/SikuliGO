#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
CLIENT_DIR="$ROOT_DIR/packages/client-node"
NPM_CACHE_DIR="${NPM_CONFIG_CACHE:-$ROOT_DIR/.test-results/npm-cache}"

mkdir -p "$NPM_CACHE_DIR"
export NPM_CONFIG_CACHE="$NPM_CACHE_DIR"

cd "$CLIENT_DIR"
if [[ "${SKIP_INSTALL:-0}" != "1" ]]; then
  NPM_CONFIG_PRODUCTION=false NPM_CONFIG_OMIT= npm install --include=dev
fi

npm run build
npm pack --dry-run

if [[ "${NPM_PUBLISH:-0}" == "1" ]]; then
  if [[ -z "${NPM_TOKEN:-}" ]]; then
    echo "Missing NPM_TOKEN for publish" >&2
    exit 1
  fi
  npm config set //registry.npmjs.org/:_authToken="${NPM_TOKEN}"
  npm publish --access public
else
  echo "Node package scaffold validated (publish skipped; set NPM_PUBLISH=1)"
fi
