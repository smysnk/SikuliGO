#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
CLIENT_DIR="$ROOT_DIR/clients/node"

cd "$CLIENT_DIR"
if [[ "${SKIP_INSTALL:-0}" != "1" ]]; then
  npm ci
fi

if [[ ! -x "node_modules/.bin/tsc" ]]; then
  echo "Missing node_modules tooling. Run: (cd $CLIENT_DIR && npm ci) or set SKIP_INSTALL=0" >&2
  exit 1
fi

npm run generate
npm run build
npm pack --dry-run

if [[ "${NPM_PUBLISH:-0}" == "1" ]]; then
  if [[ -z "${NPM_TOKEN:-}" ]]; then
    echo "Missing NPM_TOKEN for publish" >&2
    exit 1
  fi
  echo "//registry.npmjs.org/:_authToken=${NPM_TOKEN}" > "$HOME/.npmrc"
  npm publish --access public
else
  echo "Node package scaffold validated (publish skipped; set NPM_PUBLISH=1)"
fi
