#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
PACKAGES_DIR="$ROOT_DIR/clients/node/packages"
NPM_CACHE_DIR="${NPM_CONFIG_CACHE:-$ROOT_DIR/.test-results/npm-cache}"
PACKAGES=(
  "bin-darwin-arm64"
  "bin-darwin-x64"
  "bin-linux-x64"
  "bin-win32-x64"
)

mkdir -p "$NPM_CACHE_DIR"
export NPM_CONFIG_CACHE="$NPM_CACHE_DIR"

cd "$ROOT_DIR"
./scripts/clients/build-node-binaries.sh

if [[ "${NPM_PUBLISH:-0}" == "1" ]]; then
  if [[ -z "${NPM_TOKEN:-}" ]]; then
    echo "Missing NPM_TOKEN for publish" >&2
    exit 1
  fi
  echo "//registry.npmjs.org/:_authToken=${NPM_TOKEN}" > "$HOME/.npmrc"
fi

for pkg in "${PACKAGES[@]}"; do
  pkg_dir="$PACKAGES_DIR/$pkg"
  cd "$pkg_dir"

  if [[ "$pkg" == "bin-win32-x64" ]]; then
    test -f "bin/sikuligrpc.exe" || { echo "Missing binary for $pkg" >&2; exit 1; }
  else
    test -f "bin/sikuligrpc" || { echo "Missing binary for $pkg" >&2; exit 1; }
  fi

  npm pack --dry-run >/dev/null

  if [[ "${NPM_PUBLISH:-0}" == "1" ]]; then
    npm publish --access public
  fi
done

if [[ "${NPM_PUBLISH:-0}" == "1" ]]; then
  echo "Node binary packages published"
else
  echo "Node binary packages validated (publish skipped; set NPM_PUBLISH=1)"
fi
