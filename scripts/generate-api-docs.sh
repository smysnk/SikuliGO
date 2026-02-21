#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
OUT_ROOT="${API_DOCS_OUT_DIR:-$ROOT_DIR/docs/api}"
INDEX_FILE="$OUT_ROOT/index.md"
RENDERER="$ROOT_DIR/scripts/render-api-doc.awk"

if [[ ! -f "$RENDERER" ]]; then
  echo "Missing renderer: $RENDERER" >&2
  exit 1
fi

mkdir -p "$OUT_ROOT"
rm -f "$OUT_ROOT"/*.md

MODULE_PATH="$(cd "$ROOT_DIR" && go list -m)"

{
  echo "# API Reference"
  echo
  echo "This API reference is generated from package comments and exported symbols using \`go doc -all\`."
  echo
  cat <<'STYLE'
<style>
  .api-type { color: #0f766e; font-weight: 700; }
  .api-func { color: #1d4ed8; font-weight: 700; }
  .api-method { color: #7c3aed; font-weight: 700; }
  .api-signature { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace; }
</style>
STYLE
  echo
  echo "Legend: <span class=\"api-type\">Type</span>, <span class=\"api-func\">Function</span>, <span class=\"api-method\">Method</span>"
  echo
  echo "## Packages"
  echo
} > "$INDEX_FILE"

for pkg in $(cd "$ROOT_DIR" && go list ./pkg/... ./internal/... | sort); do
  rel="${pkg#${MODULE_PATH}/}"
  slug="${rel//\//-}"
  out_file="$OUT_ROOT/${slug}.md"
  tmp_doc="$(mktemp)"

  (cd "$ROOT_DIR" && go doc -all "$pkg") > "$tmp_doc"
  awk -v rel="$rel" -f "$RENDERER" "$tmp_doc" > "$out_file"
  rm -f "$tmp_doc"

  echo "- [\`$rel\`](./$slug)" >> "$INDEX_FILE"
done
