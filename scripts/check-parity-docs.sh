#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
PARITY_DIR="$ROOT_DIR/docs/reference/parity"
REGENERATE=1

if [[ "${1:-}" == "--skip-regenerate" ]] || [[ "${SKIP_REGENERATE:-0}" == "1" ]]; then
  REGENERATE=0
fi

if [[ ! -d "$PARITY_DIR" ]]; then
  echo "Missing parity docs directory: $PARITY_DIR" >&2
  exit 1
fi

if [[ "$REGENERATE" -eq 1 ]]; then
  tmp_out="$(mktemp -d)"
  trap 'rm -rf "$tmp_out"' EXIT

  PARITY_DOCS_OUT_DIR="$tmp_out" "$ROOT_DIR/scripts/generate-parity-docs.sh"
  for generated in java-to-go-mapping.md api-parity-status.md; do
    if ! diff -u "$PARITY_DIR/$generated" "$tmp_out/$generated" >/dev/null; then
      echo "Parity docs are out of date. Run ./scripts/generate-parity-docs.sh and commit updates." >&2
      diff -u "$PARITY_DIR/$generated" "$tmp_out/$generated" || true
      exit 1
    fi
  done
fi

required=(
  "$PARITY_DIR/java-to-go-mapping.md"
  "$PARITY_DIR/api-parity-status.md"
  "$PARITY_DIR/behavioral-differences.md"
  "$PARITY_DIR/parity-gaps.md"
  "$PARITY_DIR/parity-test-matrix.md"
  "$PARITY_DIR/api-migration-examples.md"
  "$PARITY_DIR/search-semantics-matrix.md"
  "$PARITY_DIR/live-screen-surface.md"
  "$PARITY_DIR/match-action-surface.md"
  "$PARITY_DIR/direct-action-surface.md"
  "$PARITY_DIR/finder-traversal-surface.md"
  "$PARITY_DIR/multi-target-search-surface.md"
  "$PARITY_DIR/ocr-collection-surface.md"
  "$PARITY_DIR/app-window-surface.md"
)

for file in "${required[@]}"; do
  if [[ ! -f "$file" ]]; then
    echo "Missing parity document: $file" >&2
    exit 1
  fi
done

for heading in \
  "## Search Semantics and Exception/Null Behavior" \
  "## Live Screen and Region Surface" \
  "## Match as a First-Class Action Target" \
  "## Direct Action Surface" \
  "## Finder Traversal and Lifecycle" \
  "## Multi-Target Search Helpers" \
  "## OCR Collection Surface" \
  "## App and Window Surface"; do
  if ! grep -F "$heading" "$PARITY_DIR/api-migration-examples.md" >/dev/null; then
    echo "Missing migration example section: $heading" >&2
    exit 1
  fi
done

echo "Parity docs validation passed."
