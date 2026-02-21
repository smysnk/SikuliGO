#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
API_DIR="$ROOT_DIR/docs/api"
AWK_RENDERER="$ROOT_DIR/scripts/render-api-doc.awk"

if [[ ! -f "$AWK_RENDERER" ]]; then
  echo "Missing renderer: $AWK_RENDERER" >&2
  exit 1
fi

for file in "$API_DIR"/*.md; do
  base="$(basename "$file")"
  if [[ "$base" == "index.md" ]]; then
    continue
  fi

  rel="$(sed -n '1s/^# API: `\(.*\)`$/\1/p' "$file")"
  if [[ -z "$rel" ]]; then
    echo "Skipping $file (missing API header)" >&2
    continue
  fi

  raw_tmp="$(mktemp)"
  out_tmp="$(mktemp)"

  awk '/^```text$/{inside=1;next} /^```$/{if (inside) {inside=0; exit}} inside {print}' "$file" > "$raw_tmp"

  if [[ ! -s "$raw_tmp" ]]; then
    echo "Skipping $file (missing raw doc block)" >&2
    rm -f "$raw_tmp" "$out_tmp"
    continue
  fi

  awk -v rel="$rel" -f "$AWK_RENDERER" "$raw_tmp" > "$out_tmp"
  mv "$out_tmp" "$file"

  rm -f "$raw_tmp"
done
