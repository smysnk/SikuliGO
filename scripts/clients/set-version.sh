#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"

if [[ $# -ne 1 ]]; then
  echo "Usage: $0 <semver>" >&2
  echo "Example: $0 0.1.5" >&2
  exit 1
fi

NEW_VERSION="$1"
if [[ ! "$NEW_VERSION" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "Invalid version: $NEW_VERSION (expected format: X.Y.Z)" >&2
  exit 1
fi

NODE_PKG="$ROOT_DIR/clients/node/package.json"
NODE_LOCK="$ROOT_DIR/clients/node/package-lock.json"
PYPROJECT="$ROOT_DIR/clients/python/pyproject.toml"
BIN_PKG_DIR="$ROOT_DIR/clients/node/packages"

if ! command -v node >/dev/null 2>&1; then
  echo "Missing node in PATH" >&2
  exit 1
fi

OLD_VERSION="$(node -e "console.log(require(process.argv[1]).version)" "$NODE_PKG")"

node - <<'JS' "$NODE_PKG" "$NODE_LOCK" "$BIN_PKG_DIR" "$OLD_VERSION" "$NEW_VERSION"
const fs = require("node:fs");
const path = require("node:path");

const [nodePkgPath, nodeLockPath, binPkgDir, oldVersion, newVersion] = process.argv.slice(2);
const binNames = [
  "@sikuligo/bin-darwin-arm64",
  "@sikuligo/bin-darwin-x64",
  "@sikuligo/bin-linux-x64",
  "@sikuligo/bin-win32-x64"
];

function writeJson(filePath, obj) {
  fs.writeFileSync(filePath, JSON.stringify(obj, null, 2) + "\n");
}

const nodePkg = JSON.parse(fs.readFileSync(nodePkgPath, "utf8"));
nodePkg.version = newVersion;
for (const name of binNames) {
  if (nodePkg.dependencies && name in nodePkg.dependencies) {
    nodePkg.dependencies[name] = newVersion;
  }
  if (nodePkg.optionalDependencies && name in nodePkg.optionalDependencies) {
    nodePkg.optionalDependencies[name] = newVersion;
  }
}
writeJson(nodePkgPath, nodePkg);

const lock = JSON.parse(fs.readFileSync(nodeLockPath, "utf8"));
lock.version = newVersion;
if (lock.packages && lock.packages[""]) {
  lock.packages[""].version = newVersion;
  for (const name of binNames) {
    if (lock.packages[""].dependencies && name in lock.packages[""].dependencies) {
      lock.packages[""].dependencies[name] = newVersion;
    }
    if (lock.packages[""].optionalDependencies && name in lock.packages[""].optionalDependencies) {
      lock.packages[""].optionalDependencies[name] = newVersion;
    }
  }
}
for (const name of binNames) {
  const key = `node_modules/${name}`;
  if (lock.packages && lock.packages[key]) {
    lock.packages[key].version = newVersion;
    if (typeof lock.packages[key].resolved === "string") {
      lock.packages[key].resolved = lock.packages[key].resolved.replace(oldVersion, newVersion);
    }
  }
}
writeJson(nodeLockPath, lock);

const binPkgFiles = fs
  .readdirSync(binPkgDir, { withFileTypes: true })
  .filter((entry) => entry.isDirectory())
  .map((entry) => path.join(binPkgDir, entry.name, "package.json"))
  .filter((filePath) => fs.existsSync(filePath));

for (const filePath of binPkgFiles) {
  const pkg = JSON.parse(fs.readFileSync(filePath, "utf8"));
  pkg.version = newVersion;
  writeJson(filePath, pkg);
}
JS

python3 - <<'PY' "$PYPROJECT" "$NEW_VERSION"
from pathlib import Path
import re
import sys

pyproject = Path(sys.argv[1])
new_version = sys.argv[2]
text = pyproject.read_text()

pattern = re.compile(r'(\[project\][\s\S]*?\nversion = ")([^"]+)(")')
updated, count = pattern.subn(rf'\g<1>{new_version}\3', text, count=1)
if count != 1:
    raise SystemExit("Failed to update [project].version in pyproject.toml")

pyproject.write_text(updated)
PY

echo "Updated project version: $OLD_VERSION -> $NEW_VERSION"
echo "Files updated:"
echo "  - clients/python/pyproject.toml"
echo "  - clients/node/package.json"
echo "  - clients/node/package-lock.json"
echo "  - clients/node/packages/bin-*/package.json"
