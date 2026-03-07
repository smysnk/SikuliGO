#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${THIS_DIR}/paths.sh"
source "${THIS_DIR}/npm-helpers.sh"

PACKAGES_DIR="$NODE_BIN_PACKAGES_DIR"
NPM_CACHE_DIR="${NPM_CONFIG_CACHE:-$ROOT_DIR/.test-results/npm-cache}"
built_manifest="$PACKAGES_DIR/.built-packages"

mkdir -p "$NPM_CACHE_DIR"
export NPM_CONFIG_CACHE="$NPM_CACHE_DIR"

cd "$ROOT_DIR"
./scripts/clients/build-node-binaries.sh

target_packages=("${NODE_BIN_PACKAGES[@]}")
if [[ -f "$built_manifest" ]]; then
  target_packages=()
  while IFS= read -r line; do
    if [[ -n "${line//[[:space:]]/}" ]]; then
      target_packages+=("$line")
    fi
  done < "$built_manifest"
fi
if [[ "${#target_packages[@]}" -eq 0 ]]; then
  echo "No Node binary packages selected for release." >&2
  exit 1
fi
echo "Node binary release targets: ${target_packages[*]}"

if [[ "${NPM_PUBLISH:-0}" == "1" ]]; then
  configure_npm_auth_token "${NPM_TOKEN:-}"
  verify_npm_auth
fi

for pkg in "${target_packages[@]}"; do
  pkg_dir="$PACKAGES_DIR/$pkg"
  if [[ ! -f "$pkg_dir/package.json" ]]; then
    echo "Missing package.json for $pkg at $pkg_dir" >&2
    exit 1
  fi

  if [[ "$pkg" == "bin-win32-x64" ]]; then
    test -f "$pkg_dir/bin/sikuli-go.exe" || { echo "Missing binary for $pkg" >&2; exit 1; }
    test -f "$pkg_dir/bin/sikuli-go-monitor.exe" || { echo "Missing monitor binary for $pkg" >&2; exit 1; }
  else
    test -f "$pkg_dir/bin/sikuli-go" || { echo "Missing binary for $pkg" >&2; exit 1; }
    test -f "$pkg_dir/bin/sikuli-go-monitor" || { echo "Missing monitor binary for $pkg" >&2; exit 1; }
  fi

  # Pack from explicit directory and disable scripts to avoid workspace/root prepack hooks.
  run_npm_no_workspace pack --dry-run --ignore-scripts "$pkg_dir" >/dev/null

  if [[ "${NPM_PUBLISH:-0}" == "1" ]]; then
    check_npm_package_visibility "@sikuligo/${pkg}"
    run_npm_no_workspace publish --ignore-scripts --access public "$pkg_dir"
  fi
done

if [[ "${NPM_PUBLISH:-0}" == "1" ]]; then
  echo "Node binary packages published"
else
  echo "Node binary packages validated (publish skipped; set NPM_PUBLISH=1)"
fi
