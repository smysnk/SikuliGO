#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
CLIENT_DIR="$ROOT_DIR/clients/python"

cd "$ROOT_DIR"
./scripts/clients/generate-python-stubs.sh

cd "$CLIENT_DIR"
if [[ "${SKIP_INSTALL:-0}" != "1" ]]; then
  python3 -m pip install --upgrade pip >/dev/null
  python3 -m pip install --upgrade build twine >/dev/null
fi

python3 - <<'PY'
import importlib.util
missing = [m for m in ("build", "twine") if importlib.util.find_spec(m) is None]
if missing:
    raise SystemExit("Missing Python modules: " + ", ".join(missing) + ". Install with: python3 -m pip install build twine")
PY

rm -rf dist build ./*.egg-info
python3 -m build
python3 -m twine check dist/*

if [[ "${PYPI_PUBLISH:-0}" == "1" ]]; then
  if [[ -z "${PYPI_TOKEN:-}" ]]; then
    echo "Missing PYPI_TOKEN for publish" >&2
    exit 1
  fi
  python3 -m twine upload --non-interactive -u __token__ -p "$PYPI_TOKEN" dist/*
else
  echo "Built Python distributions in $CLIENT_DIR/dist (publish skipped; set PYPI_PUBLISH=1)"
fi
