#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${THIS_DIR}/paths.sh"
source "${THIS_DIR}/npm-helpers.sh"

if [[ $# -lt 1 || $# -gt 2 ]]; then
  echo "Usage: $0 <bin-package> [source-version]" >&2
  echo "Example: $0 bin-darwin-x64 0.1.54" >&2
  exit 1
fi

pkg="$1"
source_version="${2:-}"
pkg_name="@sikuligo/${pkg}"
pkg_dir="${NODE_BIN_PACKAGES_DIR}/${pkg}"
target_version="$(node -e "console.log(require(process.argv[1]).version)" "$NODE_PACKAGE_JSON")"
tmp_dir="$(mktemp -d)"
tarball_path="${tmp_dir}/${pkg}.tgz"

cleanup() {
  rm -rf "$tmp_dir"
}
trap cleanup EXIT

case "$pkg" in
  bin-darwin-arm64|bin-darwin-x64|bin-linux-x64|bin-win32-x64)
    ;;
  *)
    echo "Unsupported node binary package: $pkg" >&2
    exit 1
    ;;
esac

mkdir -p "${NODE_BIN_PACKAGES_DIR}"

if [[ -z "$source_version" ]]; then
  source_version="$(run_npm_no_workspace view "$pkg_name" versions --json | node -e '
const versions = JSON.parse(require("node:fs").readFileSync(0, "utf8"));
if (!Array.isArray(versions) || versions.length === 0) process.exit(1);
console.log(versions[versions.length - 1]);
')"
fi

if [[ -z "${source_version}" ]]; then
  echo "Unable to determine source version for ${pkg_name}" >&2
  exit 1
fi

if run_npm_no_workspace view "${pkg_name}@${target_version}" version >/dev/null 2>&1; then
  echo "Backfill target already published: ${pkg_name}@${target_version}"
  exit 0
fi

tarball_url="$(run_npm_no_workspace view "${pkg_name}@${source_version}" dist.tarball)"
if [[ -z "$tarball_url" ]]; then
  echo "Unable to resolve tarball URL for ${pkg_name}@${source_version}" >&2
  exit 1
fi

curl -fsSL "$tarball_url" -o "$tarball_path"
tar -xzf "$tarball_path" -C "$tmp_dir"

rm -rf "$pkg_dir"
mkdir -p "$pkg_dir"
cp -R "${tmp_dir}/package/." "$pkg_dir/"

node - <<'JS' "$pkg_dir/package.json" "$target_version"
const fs = require("node:fs");
const filePath = process.argv[2];
const newVersion = process.argv[3];
const pkg = JSON.parse(fs.readFileSync(filePath, "utf8"));
pkg.version = newVersion;
fs.writeFileSync(filePath, JSON.stringify(pkg, null, 2) + "\n");
JS

if [[ "$pkg" == "bin-win32-x64" ]]; then
  test -f "$pkg_dir/bin/sikuli-go.exe" || { echo "Missing binary for ${pkg_name} backfill" >&2; exit 1; }
else
  test -f "$pkg_dir/bin/sikuli-go" || { echo "Missing binary for ${pkg_name} backfill" >&2; exit 1; }
fi

run_npm_no_workspace pack --dry-run --ignore-scripts "$pkg_dir" >/dev/null

if [[ "${NPM_PUBLISH:-0}" == "1" ]]; then
  configure_npm_auth_token "${NPM_TOKEN:-}"
  verify_npm_auth
  run_npm_no_workspace publish --ignore-scripts --access public "$pkg_dir"
  echo "Backfilled ${pkg_name}@${target_version} from ${source_version}"
else
  echo "Validated backfill ${pkg_name}@${target_version} from ${source_version} (publish skipped)"
fi
