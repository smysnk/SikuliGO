#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${THIS_DIR}/paths.sh"

PACKAGES_DIR="$NODE_BIN_PACKAGES_DIR"
NODE_CLIENT_PKG="$NODE_PACKAGE_JSON"
GO_CACHE_DIR="${GOCACHE:-$ROOT_DIR/.test-results/go-build}"
GO_MOD_CACHE_DIR="${GOMODCACHE:-$ROOT_DIR/.test-results/go-mod}"
GO_BUILD_TAGS="${GO_BUILD_TAGS:-$SIKULIGO_GO_BUILD_TAGS}"
TARGETS=(
  "darwin arm64 bin-darwin-arm64"
  "darwin amd64 bin-darwin-x64"
  "linux amd64 bin-linux-x64"
  "windows amd64 bin-win32-x64"
)
built_manifest="$PACKAGES_DIR/.built-packages"

host_goos() {
  case "$(uname -s)" in
    Darwin) echo "darwin" ;;
    Linux) echo "linux" ;;
    MINGW*|MSYS*|CYGWIN*|Windows_NT) echo "windows" ;;
    *) go env GOOS ;;
  esac
}

host_goarch() {
  case "$(uname -m)" in
    arm64|aarch64) echo "arm64" ;;
    x86_64|amd64) echo "amd64" ;;
    *) go env GOARCH ;;
  esac
}

configure_macos_ocr_env() {
  if [[ "$(uname -s)" != "Darwin" ]]; then
    return 0
  fi
  if ! command -v brew >/dev/null 2>&1; then
    return 0
  fi

  local homebrew_prefix lept_prefix tess_prefix
  homebrew_prefix="$(brew --prefix)"
  lept_prefix="$(brew --prefix leptonica 2>/dev/null || true)"
  tess_prefix="$(brew --prefix tesseract 2>/dev/null || true)"

  if [[ -n "$homebrew_prefix" ]]; then
    export PKG_CONFIG_PATH="${homebrew_prefix}/lib/pkgconfig${PKG_CONFIG_PATH:+:${PKG_CONFIG_PATH}}"
    export CGO_CFLAGS="${CGO_CFLAGS:-} -I${homebrew_prefix}/include"
    export CGO_CPPFLAGS="${CGO_CPPFLAGS:-} -I${homebrew_prefix}/include"
    export CGO_CXXFLAGS="${CGO_CXXFLAGS:-} -I${homebrew_prefix}/include"
    export CGO_LDFLAGS="${CGO_LDFLAGS:-} -L${homebrew_prefix}/lib -llept -ltesseract"
  fi
  if [[ -n "$lept_prefix" && -n "$tess_prefix" ]]; then
    export CGO_CFLAGS="${CGO_CFLAGS:-} -I${lept_prefix}/include -I${tess_prefix}/include"
    export CGO_CPPFLAGS="${CGO_CPPFLAGS:-} -I${lept_prefix}/include -I${tess_prefix}/include"
    export CGO_CXXFLAGS="${CGO_CXXFLAGS:-} -I${lept_prefix}/include -I${tess_prefix}/include"
    export CGO_LDFLAGS="${CGO_LDFLAGS:-} -L${lept_prefix}/lib -L${tess_prefix}/lib -llept -ltesseract"
  fi
}

normalize_env_flags() {
  export CGO_CFLAGS="$(echo "${CGO_CFLAGS:-}" | xargs 2>/dev/null || true)"
  export CGO_CPPFLAGS="$(echo "${CGO_CPPFLAGS:-}" | xargs 2>/dev/null || true)"
  export CGO_CXXFLAGS="$(echo "${CGO_CXXFLAGS:-}" | xargs 2>/dev/null || true)"
  export CGO_LDFLAGS="$(echo "${CGO_LDFLAGS:-}" | xargs 2>/dev/null || true)"
}

should_build_target() {
  local goos="$1"
  local goarch="$2"
  local pkg="$3"
  local requested="${NODE_BIN_TARGETS:-}"
  if [[ -z "${requested//[[:space:]]/}" && "$(host_goos)" == "darwin" ]]; then
    [[ "$goos" == "$(host_goos)" && "$goarch" == "$(host_goarch)" ]]
    return
  fi
  if [[ -z "${requested//[[:space:]]/}" ]]; then
    return 0
  fi
  requested="${requested//,/ }"
  local token=""
  for token in $requested; do
    token="$(echo "$token" | tr '[:upper:]' '[:lower:]')"
    if [[ "$token" == "$pkg" || "$token" == "$goos/$goarch" || "$token" == "$goos-$goarch" ]]; then
      return 0
    fi
  done
  return 1
}

if ! command -v go >/dev/null 2>&1; then
  echo "Missing go in PATH" >&2
  exit 1
fi

if ! command -v node >/dev/null 2>&1; then
  echo "Missing node in PATH" >&2
  exit 1
fi

configure_macos_ocr_env
normalize_env_flags

mkdir -p "$GO_CACHE_DIR" "$GO_MOD_CACHE_DIR"
rm -f "$built_manifest"

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

built_pkgs=()
for target in "${TARGETS[@]}"; do
  IFS=' ' read -r goos goarch pkg <<<"$target"
  if ! should_build_target "$goos" "$goarch" "$pkg"; then
    continue
  fi
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
    export GOCACHE="$GO_CACHE_DIR"
    export GOMODCACHE="$GO_MOD_CACHE_DIR"
    GOOS="$goos" GOARCH="$goarch" \
      go build -tags "$GO_BUILD_TAGS" -trimpath -ldflags="-s -w" -o "$out" ./cmd/sikuligrpc
  )

  if [[ "$goos" != "windows" ]]; then
    chmod +x "$out"
  fi
  built_pkgs+=("$pkg")
done

if [[ "${#built_pkgs[@]}" -eq 0 ]]; then
  echo "No Node binary targets selected. Set NODE_BIN_TARGETS to one or more of: ${NODE_BIN_PACKAGES[*]}" >&2
  exit 1
fi

printf '%s\n' "${built_pkgs[@]}" > "$built_manifest"

checksum_file="$PACKAGES_DIR/checksums.txt"
rm -f "$checksum_file"
if command -v sha256sum >/dev/null 2>&1; then
  for pkg in "${built_pkgs[@]}"; do
    if [[ -f "$PACKAGES_DIR/$pkg/bin/sikuligo" ]]; then
      sha256sum "$PACKAGES_DIR/$pkg/bin/sikuligo" >> "$checksum_file"
    elif [[ -f "$PACKAGES_DIR/$pkg/bin/sikuligo.exe" ]]; then
      sha256sum "$PACKAGES_DIR/$pkg/bin/sikuligo.exe" >> "$checksum_file"
    fi
  done
elif command -v shasum >/dev/null 2>&1; then
  for pkg in "${built_pkgs[@]}"; do
    if [[ -f "$PACKAGES_DIR/$pkg/bin/sikuligo" ]]; then
      shasum -a 256 "$PACKAGES_DIR/$pkg/bin/sikuligo" >> "$checksum_file"
    elif [[ -f "$PACKAGES_DIR/$pkg/bin/sikuligo.exe" ]]; then
      shasum -a 256 "$PACKAGES_DIR/$pkg/bin/sikuligo.exe" >> "$checksum_file"
    fi
  done
fi

echo "Built Node binary packages: ${built_pkgs[*]}"
echo "Built Node binary package payloads in: $PACKAGES_DIR"
