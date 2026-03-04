#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SEED_FILE="$ROOT_DIR/docs/reference/parity/java-to-go-seed.tsv"
OUT_DIR="${PARITY_DOCS_OUT_DIR:-$ROOT_DIR/docs/reference/parity}"
OUT_FILE="$OUT_DIR/java-to-go-mapping.md"
SIG_FILE="$ROOT_DIR/packages/api/pkg/sikuli/signatures.go"
PROTO_FILE="$ROOT_DIR/packages/api/proto/sikuli/v1/sikuli.proto"

if [[ ! -f "$SEED_FILE" ]]; then
  echo "Missing seed mapping: $SEED_FILE" >&2
  exit 1
fi

mkdir -p "$OUT_DIR"

{
  echo "# Java to Go API Mapping"
  echo
  echo "This document is generated from \`docs/reference/parity/java-to-go-seed.tsv\` and source surfaces in \`packages/api/pkg/sikuli/signatures.go\` and \`packages/api/proto/sikuli/v1/sikuli.proto\`."
  echo
  echo "## Symbol Mapping"
  echo
  echo "| Java/SikuliX Symbol | Go Surface | gRPC RPC | Node API | Python API | Status | Notes |"
  echo "|---|---|---|---|---|---|---|"
  awk -F '\t' '!/^#/ && NF >= 7 {
    java=$1; go=$2; grpc=$3; node=$4; py=$5; status=$6; notes=$7;
    gsub(/\|/, "\\|", java);
    gsub(/\|/, "\\|", go);
    gsub(/\|/, "\\|", grpc);
    gsub(/\|/, "\\|", node);
    gsub(/\|/, "\\|", py);
    gsub(/\|/, "\\|", status);
    gsub(/\|/, "\\|", notes);
    printf("| `%s` | `%s` | `%s` | `%s` | `%s` | `%s` | %s |\n", java, go, grpc, node, py, status, notes);
  }' "$SEED_FILE"
  echo
  echo "## Status Summary"
  echo
  awk -F '\t' '!/^#/ && NF >= 7 { c[$6]++ } END {
    order[1]="parity-ready"; order[2]="partial"; order[3]="gap";
    for (i=1; i<=3; i++) {
      k=order[i];
      printf("- `%s`: %d\n", k, c[k]+0);
    }
  }' "$SEED_FILE"
  echo
  echo "## Go API Interface Surface"
  echo
  echo "Extracted from \`packages/api/pkg/sikuli/signatures.go\`:"
  echo
  awk '
    /^type [A-Za-z0-9_]+ interface \{/ {
      iface=$2;
      print "### `" iface "`";
      print "";
      in_iface=1;
      next;
    }
    in_iface && /^}/ {
      in_iface=0;
      print "";
      next;
    }
    in_iface {
      line=$0;
      sub(/^[[:space:]]+/, "", line);
      sub(/[[:space:]]+$/, "", line);
      if (line != "") {
        print "- `" line "`";
      }
    }
  ' "$SIG_FILE"
  echo
  echo "## gRPC Surface"
  echo
  echo "Extracted from \`packages/api/proto/sikuli/v1/sikuli.proto\`:"
  echo
  awk '
    /^[[:space:]]*rpc[[:space:]]+[A-Za-z0-9_]+[[:space:]]*\(/ {
      line=$0;
      sub(/^[[:space:]]*/, "", line);
      print "- `" line "`";
    }
  ' "$PROTO_FILE"
  echo
  echo "## Maintenance"
  echo
  echo "- Update the seed file when parity mappings change."
  echo "- Run \`./scripts/generate-parity-docs.sh\` after updates."
  echo "- CI verifies this file is up to date."
} > "$OUT_FILE"

echo "Generated parity mapping: $OUT_FILE"
