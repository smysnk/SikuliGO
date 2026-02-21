#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
OUT_ROOT="$ROOT_DIR/docs/api"
INDEX_FILE="$OUT_ROOT/index.md"

mkdir -p "$OUT_ROOT"
rm -f "$OUT_ROOT"/*.md

MODULE_PATH="$(cd "$ROOT_DIR" && go list -m)"

{
  echo "# GoLang API Reference"
  echo
  echo "This API reference is generated from package comments and exported symbols using \`go doc -all\`."
  echo
  echo "## Packages"
  echo
} > "$INDEX_FILE"

for pkg in $(cd "$ROOT_DIR" && go list ./pkg/... ./internal/... | sort); do
  rel="${pkg#${MODULE_PATH}/}"
  slug="${rel//\//-}"
  out_file="$OUT_ROOT/${slug}.md"

  {
    echo "# API: \`$rel\`"
    echo
    echo "[Back to API Index](./)"
    echo
    echo "## Full Package Doc"
    echo
    echo '```text'
    (cd "$ROOT_DIR" && go doc -all "$pkg")
    echo '```'
  } > "$out_file"

  echo "- [\`$rel\`](./$slug)" >> "$INDEX_FILE"
done
