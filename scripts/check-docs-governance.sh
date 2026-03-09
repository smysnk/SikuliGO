#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SOURCE_MAP="$ROOT_DIR/docs/strategy/documentation-source-map.md"
WORKFLOW_DOC="$ROOT_DIR/docs/contribution/docs-workflow.md"

fail() {
  echo "$1" >&2
  exit 1
}

require_file() {
  local path="$1"
  [[ -f "$path" ]] || fail "Missing required phase-0 documentation file: $path"
}

require_line() {
  local file="$1"
  local text="$2"
  grep -Fq "$text" "$file" || fail "$file: missing required text: $text"
}

require_file "$SOURCE_MAP"
require_file "$WORKFLOW_DOC"

sections=(
  "## Downloads"
  "## Getting Started"
  "## Node.js Client"
  "## Python Client"
  "## Golang API"
  "## Getting Help"
  "## Contribution"
  "## License"
)

for section in "${sections[@]}"; do
  require_line "$SOURCE_MAP" "$section"
done

targets=(
  "$ROOT_DIR/docs/downloads/index.md"
  "$ROOT_DIR/docs/getting-started/index.md"
  "$ROOT_DIR/docs/nodejs-client/index.md"
  "$ROOT_DIR/docs/python-client/index.md"
  "$ROOT_DIR/docs/golang-api/index.md"
  "$ROOT_DIR/docs/getting-help/index.md"
  "$ROOT_DIR/docs/contribution/index.md"
  "$ROOT_DIR/docs/license/index.md"
)

for target in "${targets[@]}"; do
  require_file "$target"
done

require_line "$WORKFLOW_DOC" "Do not land end-user documentation changes only in a package README."
require_line "$WORKFLOW_DOC" '`docs/strategy/documentation-source-map.md`'
require_line "$WORKFLOW_DOC" "./scripts/check-docs-governance.sh"

require_line "$ROOT_DIR/README.md" "<!-- DOCS_CANONICAL_TARGET: docs/downloads/index.md -->"
require_line "$ROOT_DIR/README.md" "<!-- DOCS_CANONICAL_TARGET: docs/getting-started/index.md -->"
require_line "$ROOT_DIR/README.md" "<!-- DOCS_SOURCE_MAP: docs/strategy/documentation-source-map.md -->"
require_line "$ROOT_DIR/README.md" "<!-- DOCS_WORKFLOW: docs/contribution/docs-workflow.md -->"

require_line "$ROOT_DIR/packages/client-node/README.md" "<!-- DOCS_CANONICAL_TARGET: docs/nodejs-client/index.md -->"
require_line "$ROOT_DIR/packages/client-node/README.md" "<!-- DOCS_SOURCE_MAP: docs/strategy/documentation-source-map.md -->"
require_line "$ROOT_DIR/packages/client-node/README.md" "<!-- DOCS_WORKFLOW: docs/contribution/docs-workflow.md -->"

require_line "$ROOT_DIR/packages/client-python/README.md" "<!-- DOCS_CANONICAL_TARGET: docs/python-client/index.md -->"
require_line "$ROOT_DIR/packages/client-python/README.md" "<!-- DOCS_SOURCE_MAP: docs/strategy/documentation-source-map.md -->"
require_line "$ROOT_DIR/packages/client-python/README.md" "<!-- DOCS_WORKFLOW: docs/contribution/docs-workflow.md -->"

echo "Documentation governance check passed."
