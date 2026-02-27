#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
API_DIR="$ROOT_DIR/packages/api"
DIST_DIR="${HOME_BREW_DIST_DIR:-$ROOT_DIR/.test-results/homebrew}"
OWNER_REPO="${GITHUB_REPOSITORY:-smysnk/SikuliGO}"
TAP_REPO="${HOMEBREW_TAP_REPO:-sikuligo/homebrew-tap}"
FORMULA_NAME="sikuligo"
FORMULA_PATH="Formula/${FORMULA_NAME}.rb"
ARM_ARCHIVE="$DIST_DIR/${FORMULA_NAME}-darwin-arm64.tar.gz"
AMD_ARCHIVE="$DIST_DIR/${FORMULA_NAME}-darwin-amd64.tar.gz"
PKG_JSON="$ROOT_DIR/packages/client-node/package.json"

mkdir -p "$DIST_DIR"
cd "$ROOT_DIR"

if [[ ! -f "$PKG_JSON" ]]; then
  echo "Missing package.json: $PKG_JSON" >&2
  exit 1
fi

if ! command -v go >/dev/null 2>&1; then
  echo "Missing go in PATH" >&2
  exit 1
fi

if ! command -v python3 >/dev/null 2>&1; then
  echo "Missing python3 in PATH" >&2
  exit 1
fi

VERSION="$(python3 - "$PKG_JSON" <<'PY'
import json
import sys

with open(sys.argv[1], "r", encoding="utf-8") as fh:
    pkg = json.load(fh)
print(pkg["version"])
PY
)"
TAG="v${VERSION}"
URL_BASE="https://github.com/${OWNER_REPO}/releases/download/${TAG}"

build_archive() {
  local arch="$1"
  local archive="$2"
  local stage_dir="$DIST_DIR/stage-${arch}"
  local files=("sikuligo")

  rm -rf "$stage_dir"
  mkdir -p "$stage_dir"

  echo "Building sikuligo darwin/${arch}"
  (
    cd "$API_DIR"
    CGO_ENABLED=0 GOOS=darwin GOARCH="$arch" \
      go build -trimpath -ldflags="-s -w" -o "$stage_dir/sikuligo" ./cmd/sikuligrpc
  )

  if [[ -d "$API_DIR/cmd/sikuligo-monitor" ]]; then
    echo "Building sikuligo-monitor darwin/${arch}"
    (
      cd "$API_DIR"
      CGO_ENABLED=0 GOOS=darwin GOARCH="$arch" \
        go build -trimpath -ldflags="-s -w" -o "$stage_dir/sikuligo-monitor" ./cmd/sikuligo-monitor
    )
    files+=("sikuligo-monitor")
  fi

  tar -C "$stage_dir" -czf "$archive" "${files[@]}"
}

sha256_file() {
  local file="$1"
  if command -v sha256sum >/dev/null 2>&1; then
    sha256sum "$file" | awk '{print $1}'
    return
  fi
  if command -v shasum >/dev/null 2>&1; then
    shasum -a 256 "$file" | awk '{print $1}'
    return
  fi
  echo "Missing sha256sum/shasum in PATH" >&2
  exit 1
}

build_archive "arm64" "$ARM_ARCHIVE"
build_archive "amd64" "$AMD_ARCHIVE"

ARM_SHA="$(sha256_file "$ARM_ARCHIVE")"
AMD_SHA="$(sha256_file "$AMD_ARCHIVE")"

FORMULA_FILE="$DIST_DIR/${FORMULA_NAME}.rb"
cat >"$FORMULA_FILE" <<EOF
class Sikuligo < Formula
  desc "SikuliGO desktop automation API server"
  homepage "https://github.com/${OWNER_REPO}"
  license "MIT"
  version "${VERSION}"

  on_macos do
    if Hardware::CPU.arm?
      url "${URL_BASE}/${FORMULA_NAME}-darwin-arm64.tar.gz"
      sha256 "${ARM_SHA}"
    else
      url "${URL_BASE}/${FORMULA_NAME}-darwin-amd64.tar.gz"
      sha256 "${AMD_SHA}"
    end
  end

  def install
    bin.install "sikuligo"
    bin.install "sikuligo-monitor" if (buildpath/"sikuligo-monitor").exist?
  end

  test do
    assert_predicate bin/"sikuligo", :exist?
  end
end
EOF

if [[ "${HOMEBREW_PUBLISH:-0}" != "1" ]]; then
  echo "Built Homebrew artifacts in $DIST_DIR (publish skipped; set HOMEBREW_PUBLISH=1)"
  echo "Formula preview: $FORMULA_FILE"
  exit 0
fi

GH_TOKEN="${GH_TOKEN:-${GITHUB_TOKEN:-}}"
if [[ -z "$GH_TOKEN" ]]; then
  echo "Missing GH_TOKEN (or GITHUB_TOKEN) for GitHub release operations" >&2
  exit 1
fi

if [[ -z "${HOMEBREW_TAP_TOKEN:-}" ]]; then
  echo "Missing HOMEBREW_TAP_TOKEN for tap repository updates" >&2
  exit 1
fi

if ! command -v gh >/dev/null 2>&1; then
  echo "Missing gh CLI in PATH" >&2
  exit 1
fi

echo "Publishing release assets to ${OWNER_REPO} tag ${TAG}"
if ! gh release view "$TAG" --repo "$OWNER_REPO" >/dev/null 2>&1; then
  gh release create "$TAG" \
    --repo "$OWNER_REPO" \
    --title "SikuliGO ${TAG}" \
    --notes "Automated release artifacts for Homebrew."
fi

gh release upload "$TAG" \
  "$ARM_ARCHIVE" \
  "$AMD_ARCHIVE" \
  --clobber \
  --repo "$OWNER_REPO"

TAP_DIR="$DIST_DIR/tap"
rm -rf "$TAP_DIR"
echo "Updating Homebrew tap ${TAP_REPO}"
git clone "https://x-access-token:${HOMEBREW_TAP_TOKEN}@github.com/${TAP_REPO}.git" "$TAP_DIR"
mkdir -p "$TAP_DIR/Formula"
cp "$FORMULA_FILE" "$TAP_DIR/$FORMULA_PATH"

pushd "$TAP_DIR" >/dev/null
git config user.name "github-actions[bot]"
git config user.email "41898282+github-actions[bot]@users.noreply.github.com"
git add "$FORMULA_PATH"
if git diff --cached --quiet; then
  echo "Homebrew formula unchanged"
else
  git commit -m "sikuligo ${VERSION}"
  git push origin HEAD
fi
popd >/dev/null

echo "Homebrew publish complete"
