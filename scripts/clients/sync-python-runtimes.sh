#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${THIS_DIR}/paths.sh"

SRC_ROOT="${NODE_BIN_PACKAGES_DIR}"
DST_ROOT="${CLIENT_PYTHON_DIR}/sikuligo/runtime"

mkdir -p "${DST_ROOT}"
find "${DST_ROOT}" -mindepth 1 -maxdepth 1 -type d -exec rm -rf {} +

for pkg in "${NODE_BIN_PACKAGES[@]}"; do
  src_pkg_dir="${SRC_ROOT}/${pkg}/bin"
  if [[ ! -d "${src_pkg_dir}" ]]; then
    continue
  fi
  dst_pkg_dir="${DST_ROOT}/${pkg}/bin"
  mkdir -p "${dst_pkg_dir}"
  find "${src_pkg_dir}" -maxdepth 1 -type f ! -name '.gitkeep' -exec cp -f {} "${dst_pkg_dir}/" \;
done

echo "Synced Python packaged runtimes into: ${DST_ROOT}"
