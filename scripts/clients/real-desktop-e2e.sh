#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${THIS_DIR}/paths.sh"

# Safety gate so this cannot accidentally run in CI/headless flows.
if [[ "${REAL_DESKTOP_E2E:-0}" != "1" ]]; then
  echo "[real-desktop-e2e] skipped (set REAL_DESKTOP_E2E=1 to run)"
  echo "[real-desktop-e2e] example: REAL_DESKTOP_E2E=1 make test-e2e"
  exit 0
fi

TMP_ROOT="$(mktemp -d /tmp/sikuligo-real-desktop-e2e.XXXXXX)"
API_BINARY="${TMP_ROOT}/sikuli-go"
SQLITE_PATH="${TMP_ROOT}/sikuligo-real-desktop-e2e.db"
SNAPSHOT_PATH="${TMP_ROOT}/snapshot.png"
FIXTURE_HTML="${TMP_ROOT}/fixture.html"
KEEP_TMP="${KEEP_TMP:-0}"
GO_TAGS="${REAL_DESKTOP_E2E_GO_TAGS:-gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d}"
PATTERN_PATH="${REAL_DESKTOP_E2E_PATTERN_PATH:-}"
FIXTURE_IMAGE_PATH="${REAL_DESKTOP_E2E_IMAGE_PATH:-}"
DISPLAY_SELECTION="${REAL_DESKTOP_E2E_DISPLAY:-}"
CONFIRM="${REAL_DESKTOP_E2E_CONFIRM:-1}"
AUTO_OPEN="${REAL_DESKTOP_E2E_AUTO_OPEN:-1}"
SETTLE_SECONDS="${REAL_DESKTOP_E2E_SETTLE_SECONDS:-3}"

API_PID=""
FIXTURE_OPENED=0

cleanup() {
  if [[ -n "${API_PID}" ]] && kill -0 "${API_PID}" >/dev/null 2>&1; then
    kill "${API_PID}" >/dev/null 2>&1 || true
    wait "${API_PID}" 2>/dev/null || true
  fi
  if [[ "${KEEP_TMP}" == "1" ]]; then
    echo "[real-desktop-e2e] keeping temp dir: ${TMP_ROOT}"
    return
  fi
  rm -rf "${TMP_ROOT}"
}
trap cleanup EXIT

step() {
  echo "[real-desktop-e2e] $1"
}

fail() {
  echo "[real-desktop-e2e] ERROR: $1" >&2
  exit 1
}

command -v node >/dev/null 2>&1 || fail "node is required"
command -v python3 >/dev/null 2>&1 || fail "python3 is required"

if [[ -z "${PATTERN_PATH}" ]]; then
  PATTERN_PATH="${TMP_ROOT}/fixture-target.png"
  export PATTERN_PATH
  python3 - <<'PY'
import base64
import os
from pathlib import Path

png_b64 = (
    "iVBORw0KGgoAAAANSUhEUgAAAGAAAABgCAIAAABt+uBvAAABUElEQVR4nO3cMWoCURRA0RjSTZHaytoii0iVVaXKqlK5iBTWVtYp7FMEgqDDFR0SieeUI8rn8hhm4OFst9vdMe7+rw9w7QQKAgWBgkBBoCBQeBj7YBiG3zzHNTj6SGiCgkBBoDB6D9r3j9/X8lZrgoJAQaAgUBAoCBQECic9B43ZLpdTnWO+Xh9eXLxtp/r9zev8vC+aoCBQECgIFAQKAgWBgkBBoCBQECgIFAQKAgWBgkBBoCBQECgIFAQKAgWBgkBhNrYctb9ZdCMLVJY4zyFQECgIFAQKAoWL9oMWq5epzrF5fj+8+Pkx2f7R49OR/aNTmKAgUBAoCBQECgIFgYJAQaAgUBAoCBQECgIFgYJAQaAgUBAoCBQECgIFgYJAwX6Q/aDLCBQECgIFgYJAQaBw0gLVDf7x7Q8TFAQKAoXRdzG+maAgUBAoCBQECgKFL+zsKmS9wrOYAAAAAElFTkSuQmCC"
)
Path(os.environ["PATTERN_PATH"]).write_bytes(base64.b64decode(png_b64))
PY
fi

if [[ -z "${FIXTURE_IMAGE_PATH}" ]]; then
  FIXTURE_IMAGE_PATH="${PATTERN_PATH}"
fi

[[ -f "${PATTERN_PATH}" ]] || fail "pattern image not found: ${PATTERN_PATH}"
[[ -f "${FIXTURE_IMAGE_PATH}" ]] || fail "fixture image not found: ${FIXTURE_IMAGE_PATH}"

step "1/7 Build sikuli-go with OCR/OpenCV tags"
(
  cd "${API_DIR}"
  go build -tags "${GO_TAGS}" -o "${API_BINARY}" ./cmd/sikuligrpc
)

step "2/7 Build Node client artifacts"
(
  cd "${CLIENT_NODE_DIR}"
  npm run build >/dev/null
)

read -r GRPC_PORT ADMIN_PORT < <(
  python3 - <<'PY'
import socket

def alloc_port():
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.bind(("127.0.0.1", 0))
    port = s.getsockname()[1]
    s.close()
    return port

print(alloc_port(), alloc_port())
PY
)

GRPC_ADDR="127.0.0.1:${GRPC_PORT}"
ADMIN_ADDR="127.0.0.1:${ADMIN_PORT}"
AUTH_TOKEN="real-desktop-e2e-token"

step "3/7 Launch sikuli-go grpc=${GRPC_ADDR} admin=${ADMIN_ADDR}"
if [[ -n "${DISPLAY_SELECTION}" ]]; then
  echo "[real-desktop-e2e] capture display selector=${DISPLAY_SELECTION}"
fi
SIKULI_CAPTURE_DISPLAY="${DISPLAY_SELECTION}" \
SIKULIGO_CAPTURE_DISPLAY="${DISPLAY_SELECTION}" \
"${API_BINARY}" \
  -listen "${GRPC_ADDR}" \
  -admin-listen "${ADMIN_ADDR}" \
  -auth-token "${AUTH_TOKEN}" \
  -enable-reflection=false \
  -sqlite-path "${SQLITE_PATH}" \
  >/dev/null 2>&1 &
API_PID=$!

# Wait for gRPC ready using the generated Node transport.
NODE_CLIENT_TRANSPORT="${CLIENT_NODE_DIR}/dist/src/client.js" \
GRPC_ADDR="${GRPC_ADDR}" \
AUTH_TOKEN="${AUTH_TOKEN}" \
node - <<'NODE'
const { Sikuli } = require(process.env.NODE_CLIENT_TRANSPORT);

async function main() {
  const client = new Sikuli({
    address: process.env.GRPC_ADDR,
    authToken: process.env.AUTH_TOKEN,
    timeoutMs: 500
  });
  try {
    await client.waitForReady(10000);
  } finally {
    client.close();
  }
}

main().catch((err) => {
  console.error(err && err.stack ? err.stack : String(err));
  process.exit(1);
});
NODE

step "4/7 Prepare desktop fixture page"
cat >"${FIXTURE_HTML}" <<HTML
<!doctype html>
<html>
<head>
<meta charset="utf-8" />
<title>SikuliGO Real Desktop E2E Fixture</title>
<style>
  body { font-family: -apple-system, BlinkMacSystemFont, Segoe UI, sans-serif; margin: 24px; }
  h1 { font-size: 48px; margin: 0 0 16px; }
  p { font-size: 20px; margin: 0 0 16px; }
  img { image-rendering: pixelated; border: 2px solid #333; display: block; }
</style>
</head>
<body>
  <h1>SIKULIGO OCR E2E</h1>
  <p>Keep this window visible while the real-desktop test runs.</p>
  <img alt="fixture" src="file://${FIXTURE_IMAGE_PATH}" />
</body>
</html>
HTML

step "5/7 Open fixture page and wait for visibility"
if [[ "${AUTO_OPEN}" == "1" ]]; then
  if command -v open >/dev/null 2>&1; then
    open "${FIXTURE_HTML}" >/dev/null 2>&1 || true
    FIXTURE_OPENED=1
  elif command -v xdg-open >/dev/null 2>&1; then
    xdg-open "${FIXTURE_HTML}" >/dev/null 2>&1 || true
    FIXTURE_OPENED=1
  fi
fi

if [[ "${CONFIRM}" == "1" && -t 0 ]]; then
  echo "[real-desktop-e2e] fixture: ${FIXTURE_HTML}"
  if [[ "${FIXTURE_OPENED}" == "0" ]]; then
    echo "[real-desktop-e2e] open the fixture manually in a visible window."
  fi
  read -r -p "[real-desktop-e2e] Press Enter when fixture is visible on your desktop... " _
else
  sleep "${SETTLE_SECONDS}"
fi

step "6/7 Run real FindOnScreen + OCR checks"
NODE_CLIENT_INDEX="${CLIENT_NODE_DIR}/dist/src/index.js" \
PATTERN_PATH="${PATTERN_PATH}" \
SNAPSHOT_PATH="${SNAPSHOT_PATH}" \
GRPC_ADDR="${GRPC_ADDR}" \
ADMIN_ADDR="${ADMIN_ADDR}" \
AUTH_TOKEN="${AUTH_TOKEN}" \
node - <<'NODE'
const fs = require("node:fs");
const { Screen, Pattern, Sikuli, loadGrayImage } = require(process.env.NODE_CLIENT_INDEX);

function assert(condition, message) {
  if (!condition) {
    throw new Error(message);
  }
}

async function main() {
  const screen = await Screen.connect({
    address: process.env.GRPC_ADDR,
    authToken: process.env.AUTH_TOKEN,
    startupTimeoutMs: 5000,
    timeoutMs: 15000
  });
  try {
    const match = await screen.find(Pattern(process.env.PATTERN_PATH).exact());
    assert(match && match.w > 0 && match.h > 0, "FindOnScreen returned invalid match");
    console.log(`[real-desktop-e2e] find_on_screen ok target=(${match.targetX},${match.targetY}) score=${match.score}`);
  } finally {
    await screen.close();
  }

  const snapshotRes = await fetch(`http://${process.env.ADMIN_ADDR}/snapshot`);
  assert(snapshotRes.ok, `snapshot endpoint failed: HTTP ${snapshotRes.status}`);
  const bytes = Buffer.from(await snapshotRes.arrayBuffer());
  fs.writeFileSync(process.env.SNAPSHOT_PATH, bytes);

  const client = await Sikuli.connect({
    address: process.env.GRPC_ADDR,
    authToken: process.env.AUTH_TOKEN,
    startupTimeoutMs: 5000,
    timeoutMs: 15000
  });
  try {
    const source = loadGrayImage(process.env.SNAPSHOT_PATH);
    const read = await client.readText({
      source,
      params: { language: "eng", min_confidence: 0.2 }
    });
    const text = String(read?.text ?? "");
    assert(/SIKULIGO/i.test(text), `ReadText did not detect expected text. got=${JSON.stringify(text)}`);

    const find = await client.findText({
      source,
      query: "SIKULIGO",
      params: { language: "eng", case_sensitive: false, min_confidence: 0.2 }
    });
    const matches = Array.isArray(find?.matches) ? find.matches : [];
    assert(matches.length > 0, "FindText returned no matches for 'SIKULIGO'");
    console.log(`[real-desktop-e2e] ocr ok read_text_len=${text.length} find_text_matches=${matches.length}`);
  } finally {
    await client.close();
  }
}

main().catch((err) => {
  console.error(err && err.stack ? err.stack : String(err));
  process.exit(1);
});
NODE

step "7/7 Completed real-desktop E2E"
echo "[real-desktop-e2e] all checks passed"
