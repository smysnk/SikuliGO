#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
API_DIR="$ROOT_DIR/docs/api"
REGENERATE=1

if [[ "${1:-}" == "--skip-regenerate" ]] || [[ "${SKIP_REGENERATE:-0}" == "1" ]]; then
  REGENERATE=0
fi

if [[ "$REGENERATE" -eq 1 ]]; then
  tmp_out="$(mktemp -d)"
  trap 'rm -rf "$tmp_out"' EXIT

  API_DOCS_OUT_DIR="$tmp_out" "$ROOT_DIR/scripts/generate-api-docs.sh"
  if ! diff -ru "$API_DIR" "$tmp_out" >/dev/null; then
    echo "API docs are out of date. Run ./scripts/generate-api-docs.sh and commit the updates." >&2
    diff -ru "$API_DIR" "$tmp_out" || true
    exit 1
  fi
fi

if [[ ! -d "$API_DIR" ]]; then
  echo "Missing API docs directory: $API_DIR" >&2
  exit 1
fi

failures=0
for file in "$API_DIR"/*.md; do
  base="$(basename "$file")"
  if [[ "$base" == "index.md" ]]; then
    continue
  fi

  if ! grep -q "<style>" "$file"; then
    echo "$file: missing embedded style block" >&2
    failures=$((failures + 1))
  fi

  if ! grep -q "Legend: <span class=\"api-type\">Type</span>" "$file"; then
    echo "$file: missing symbol legend" >&2
    failures=$((failures + 1))
  fi

  if ! grep -q "^## Symbol Index$" "$file"; then
    echo "$file: missing symbol index section" >&2
    failures=$((failures + 1))
  fi

  if ! grep -Eq "<a id=\"(type|func|method)-" "$file"; then
    echo "$file: missing declaration anchors" >&2
    failures=$((failures + 1))
  fi

  if ! grep -Eq "\(#(type|func|method)-" "$file"; then
    echo "$file: missing symbol links" >&2
    failures=$((failures + 1))
  fi

  if ! grep -q "^## Raw Package Doc$" "$file"; then
    echo "$file: missing raw package doc section" >&2
    failures=$((failures + 1))
  fi
done

if [[ "$failures" -ne 0 ]]; then
  echo "API docs validation failed with $failures issue(s)." >&2
  exit 1
fi

echo "API docs validation passed."
