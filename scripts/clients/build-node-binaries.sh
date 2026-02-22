#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
PACKAGES_DIR="$ROOT_DIR/clients/node/packages"
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

for target in "${TARGETS[@]}"; do
  IFS=' ' read -r goos goarch pkg <<<"$target"
  pkg_dir="$PACKAGES_DIR/$pkg"
  bin_dir="$pkg_dir/bin"
  mkdir -p "$bin_dir"

  if [[ "$goos" == "windows" ]]; then
    out="$bin_dir/sikuligrpc.exe"
    rm -f "$bin_dir/sikuligrpc"
  else
    out="$bin_dir/sikuligrpc"
    rm -f "$bin_dir/sikuligrpc.exe"
  fi

  echo "Building $pkg ($goos/$goarch)"
  CGO_ENABLED=0 GOOS="$goos" GOARCH="$goarch" \
    go build -trimpath -ldflags="-s -w" -o "$out" ./cmd/sikuligrpc

  if [[ "$goos" != "windows" ]]; then
    chmod +x "$out"
  fi
done

checksum_file="$PACKAGES_DIR/checksums.txt"
rm -f "$checksum_file"
if command -v sha256sum >/dev/null 2>&1; then
  (
    cd "$PACKAGES_DIR"
    sha256sum bin-*/bin/sikuligrpc* > checksums.txt
  )
elif command -v shasum >/dev/null 2>&1; then
  (
    cd "$PACKAGES_DIR"
    shasum -a 256 bin-*/bin/sikuligrpc* > checksums.txt
  )
fi

echo "Built Node binary package payloads in: $PACKAGES_DIR"
