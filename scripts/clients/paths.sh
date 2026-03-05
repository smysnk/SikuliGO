#!/usr/bin/env bash

# Shared paths/constants for client release/build scripts.
# This file is sourced by other scripts and intentionally avoids `set -euo pipefail`.

if [[ "${SIKULIGO_CLIENT_PATHS_LOADED:-0}" == "1" && -n "${NODE_BIN_PACKAGES_DIR:-}" ]]; then
  return 0
fi
SIKULIGO_CLIENT_PATHS_LOADED=1

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "${SCRIPT_DIR}/../.." && pwd)"

API_DIR="${ROOT_DIR}/packages/api"
CLIENT_NODE_DIR="${ROOT_DIR}/packages/client-node"
CLIENT_PYTHON_DIR="${ROOT_DIR}/packages/client-python"
CLIENT_LUA_DIR="${ROOT_DIR}/packages/client-lua"

NODE_BIN_PACKAGES_DIR="${CLIENT_NODE_DIR}/packages"
NODE_PACKAGE_JSON="${CLIENT_NODE_DIR}/package.json"
NODE_PACKAGE_LOCK="${CLIENT_NODE_DIR}/package-lock.json"
PYTHON_PROJECT_TOML="${CLIENT_PYTHON_DIR}/pyproject.toml"
SIKULIGO_GO_BUILD_TAGS="${SIKULIGO_GO_BUILD_TAGS:-gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d}"

readonly ROOT_DIR API_DIR CLIENT_NODE_DIR CLIENT_PYTHON_DIR CLIENT_LUA_DIR NODE_BIN_PACKAGES_DIR NODE_PACKAGE_JSON NODE_PACKAGE_LOCK PYTHON_PROJECT_TOML SIKULIGO_GO_BUILD_TAGS

NODE_BIN_PACKAGES=(
  "bin-darwin-arm64"
  "bin-darwin-x64"
  "bin-linux-x64"
  "bin-win32-x64"
)
readonly NODE_BIN_PACKAGES
