#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
API_DIR="$ROOT_DIR/packages/api"
PACKAGES_DIR="$ROOT_DIR/packages/client-node/packages"
NODE_CLIENT_PKG="$ROOT_DIR/packages/client-node/package.json"
TARGETS=(
  "darwin arm64 bin-darwin-arm64"
  "darwin amd64 bin-darwin-x64"
  "linux amd64 bin-linux-x64"
  "windows amd64 bin-win32-x64"
)

if ! command -v go >/dev/null 2>&1; then
  echo "Missing go in PATH" >&2
  exit 1
fi

if ! command -v node >/dev/null 2>&1; then
  echo "Missing node in PATH" >&2
  exit 1
fi

NODE_VERSION="$(node -e "console.log(require(process.argv[1]).version)" "$NODE_CLIENT_PKG")"

ensure_pkg_scaffold() {
  local pkg="$1"
  local goos="$2"
  local goarch="$3"
  local pkg_dir="$PACKAGES_DIR/$pkg"
  local readme="$pkg_dir/README.md"
  local manifest="$pkg_dir/package.json"
  local bin_name="sikuligo"
  if [[ "$goos" == "windows" ]]; then
    bin_name="sikuligo.exe"
  fi

  mkdir -p "$pkg_dir/bin"

  if [[ ! -f "$readme" ]]; then
    cat >"$readme" <<EOF
# @sikuligo/$pkg

Platform binary package for SikuliGO ($goos/$goarch).
EOF
  fi

  if [[ ! -f "$manifest" ]]; then
    cat >"$manifest" <<EOF
{
  "name": "@sikuligo/$pkg",
  "version": "$NODE_VERSION",
  "description": "sikuligo binary for $goos $goarch",
  "license": "MIT",
  "files": [
    "bin/$bin_name",
    "README.md"
  ],
  "publishConfig": {
    "access": "public"
  }
}
EOF
  fi
}

for target in "${TARGETS[@]}"; do
  IFS=' ' read -r goos goarch pkg <<<"$target"
  pkg_dir="$PACKAGES_DIR/$pkg"
  bin_dir="$pkg_dir/bin"
  ensure_pkg_scaffold "$pkg" "$goos" "$goarch"
  mkdir -p "$bin_dir"

  if [[ "$goos" == "windows" ]]; then
    out="$bin_dir/sikuligo.exe"
    rm -f "$bin_dir/sikuligrpc" "$bin_dir/sikuligrpc.exe" "$bin_dir/sikuligo"
  else
    out="$bin_dir/sikuligo"
    rm -f "$bin_dir/sikuligrpc" "$bin_dir/sikuligrpc.exe" "$bin_dir/sikuligo.exe"
  fi

  echo "Building $pkg ($goos/$goarch)"
  (
    cd "$API_DIR"
    CGO_ENABLED=0 GOOS="$goos" GOARCH="$goarch" \
      go build -trimpath -ldflags="-s -w" -o "$out" ./cmd/sikuligrpc
  )

  if [[ "$goos" != "windows" ]]; then
    chmod +x "$out"
  fi
done

checksum_file="$PACKAGES_DIR/checksums.txt"
rm -f "$checksum_file"
if command -v sha256sum >/dev/null 2>&1; then
  (
    cd "$PACKAGES_DIR"
    sha256sum bin-*/bin/sikuligo* > checksums.txt
  )
elif command -v shasum >/dev/null 2>&1; then
  (
    cd "$PACKAGES_DIR"
    shasum -a 256 bin-*/bin/sikuligo* > checksums.txt
  )
fi

echo "Built Node binary package payloads in: $PACKAGES_DIR"
