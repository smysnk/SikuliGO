#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${THIS_DIR}/paths.sh"
source "${THIS_DIR}/npm-helpers.sh"

CLIENT_DIR="$CLIENT_NODE_DIR"
NPM_CACHE_DIR="${NPM_CONFIG_CACHE:-$ROOT_DIR/.test-results/npm-cache}"
NODE_PUBLISH_MODE="${NPM_PUBLISH:-0}"
NODE_CLIENT_BUILD_MODE="${NODE_CLIENT_BUILD:-0}"
NODE_SKIP_INSTALL="${SKIP_INSTALL:-0}"

step() {
  echo "[node-client-release] $1"
}

step "1/10 Preflight required tools and files"
if ! command -v npm >/dev/null 2>&1; then
  echo "Missing npm in PATH" >&2
  exit 1
fi
if ! command -v node >/dev/null 2>&1; then
  echo "Missing node in PATH" >&2
  exit 1
fi
if [[ ! -f "$NODE_PACKAGE_JSON" ]]; then
  echo "Missing package.json: $NODE_PACKAGE_JSON" >&2
  exit 1
fi

step "2/10 Prepare npm cache and working directory"
mkdir -p "$NPM_CACHE_DIR"
export NPM_CONFIG_CACHE="$NPM_CACHE_DIR"
cd "$CLIENT_DIR"

step "3/10 Verify package identity"
PKG_NAME="$(node -e "console.log(require(process.argv[1]).name)" "$NODE_PACKAGE_JSON")"
PKG_VERSION="$(node -e "console.log(require(process.argv[1]).version)" "$NODE_PACKAGE_JSON")"
if [[ "$PKG_NAME" != "@sikuligo/sikuligo" ]]; then
  echo "Unexpected package name: $PKG_NAME" >&2
  exit 1
fi
echo "Target package: ${PKG_NAME}@${PKG_VERSION}"

step "4/10 Install dependencies (dev included)"
if [[ "$NODE_SKIP_INSTALL" != "1" ]]; then
  NPM_CONFIG_OMIT= run_npm_no_workspace install --include=dev
else
  echo "Skipping install (SKIP_INSTALL=1)"
fi

step "5/10 Validate required artifacts"
required_files=(
  "dist/src/index.js"
  "dist/src/index.d.ts"
  "dist/src/binary.js"
  "dist/src/binary.d.ts"
  "dist/src/client.js"
  "dist/src/client.d.ts"
  "dist/src/doctor.js"
  "dist/src/doctor.d.ts"
  "dist/src/image.js"
  "dist/src/image.d.ts"
  "dist/src/launcher.js"
  "dist/src/launcher.d.ts"
  "dist/src/sikuli.js"
  "dist/src/sikuli.d.ts"
  "dist/src/sikulix.js"
  "dist/src/sikulix.d.ts"
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

step "6/10 Build artifacts when required"
if [[ "$NODE_CLIENT_BUILD_MODE" == "1" || ${#missing[@]} -gt 0 ]]; then
  if [[ ${#missing[@]} -gt 0 ]]; then
    echo "Missing Node client artifacts: ${missing[*]}"
    echo "Attempting npm run build to regenerate dist/generated artifacts..."
  else
    echo "Forcing build (NODE_CLIENT_BUILD=1)"
  fi
  run_npm_no_workspace run build
fi

step "7/10 Re-validate required artifacts after build"
missing_post=()
for f in "${required_files[@]}"; do
  if [[ ! -f "$f" ]]; then
    missing_post+=("$f")
  fi
done
if [[ ${#missing_post[@]} -gt 0 ]]; then
  echo "Required artifacts still missing after build: ${missing_post[*]}" >&2
  exit 1
fi

step "8/10 Validate package tarball"
run_npm_no_workspace pack --dry-run --ignore-scripts "$CLIENT_DIR"

if [[ "$NODE_PUBLISH_MODE" == "1" ]]; then
  step "9/10 Authenticate and preflight npm registry checks"
  (
    cd "$ROOT_DIR"
    configure_npm_auth_token "${NPM_TOKEN:-}"
    verify_npm_auth
    check_npm_package_visibility "@sikuligo/sikuligo"
  )

  step "10/10 Publish @sikuligo/sikuligo"
  (
    cd "$ROOT_DIR"
    run_npm_no_workspace publish --ignore-scripts --access public "$CLIENT_DIR"
  )
else
  step "9/10 Publish preflight skipped (NPM_PUBLISH!=1)"
  step "10/10 Publish skipped; release validation complete"
fi
