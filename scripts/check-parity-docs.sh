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
  if ! diff -u "$PARITY_DIR/java-to-go-mapping.md" "$tmp_out/java-to-go-mapping.md" >/dev/null; then
    echo "Parity docs are out of date. Run ./scripts/generate-parity-docs.sh and commit updates." >&2
    diff -u "$PARITY_DIR/java-to-go-mapping.md" "$tmp_out/java-to-go-mapping.md" || true
    exit 1
  fi
fi

required=(
  "$PARITY_DIR/java-to-go-mapping.md"
  "$PARITY_DIR/behavioral-differences.md"
  "$PARITY_DIR/parity-gaps.md"
  "$PARITY_DIR/parity-test-matrix.md"
)

for file in "${required[@]}"; do
  if [[ ! -f "$file" ]]; then
    echo "Missing parity document: $file" >&2
    exit 1
  fi
done

echo "Parity docs validation passed."
