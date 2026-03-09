#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

cd "$ROOT_DIR/packages/api"
go test ./pkg/sikuli -run '^TestAPIParityContracts$' -count=1

cd "$ROOT_DIR"
./scripts/check-parity-docs.sh

echo "Parity gate passed."
