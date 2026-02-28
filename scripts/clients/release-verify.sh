#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${THIS_DIR}/paths.sh"
source "${THIS_DIR}/npm-helpers.sh"

echo "Client release verification preflight..."
for required in "$NODE_PACKAGE_JSON" "$NODE_PACKAGE_LOCK" "$PYTHON_PROJECT_TOML"; do
  if [[ ! -f "$required" ]]; then
    echo "Missing required file: $required" >&2
    exit 1
  fi
done

if [[ -n "${BUILD_NUMBER:-${GITHUB_RUN_NUMBER:-}}" ]]; then
  echo "Setting versions from build metadata..."
  "${THIS_DIR}/set-version-from-build.sh"
else
  echo "BUILD_NUMBER not set; skipping set-version-from-build."
fi

if [[ "${RUN_NODE_BINARIES_VERIFY:-1}" == "1" ]]; then
  echo "Verifying Node binary package flow (no publish)..."
  NPM_PUBLISH=0 "${THIS_DIR}/release-node-binaries.sh"
fi

if [[ "${RUN_NODE_CLIENT_VERIFY:-1}" == "1" ]]; then
  echo "Verifying Node client package flow (no publish)..."
  NPM_PUBLISH=0 "${THIS_DIR}/release-node-client.sh"
fi

if [[ "${RUN_PYTHON_VERIFY:-1}" == "1" ]]; then
  echo "Verifying Python client package flow (no publish)..."
  PYPI_PUBLISH=0 "${THIS_DIR}/release-python-client.sh"
fi

if [[ "${RUN_HOMEBREW_VERIFY:-0}" == "1" ]]; then
  echo "Verifying Homebrew packaging flow (no publish)..."
  HOMEBREW_PUBLISH=0 "${THIS_DIR}/release-homebrew.sh"
fi

echo "Client release verification complete."
