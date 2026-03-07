#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${THIS_DIR}/paths.sh"

TMP_ROOT="$(mktemp -d /tmp/sikuligo-local-verify.XXXXXX)"
PROJECT_DIR="${TMP_ROOT}/project"
LOG_FILE="${TMP_ROOT}/smoke.log"
KEEP_TMP="${KEEP_TMP:-0}"
VERIFY_PACKED_INSTALL="${VERIFY_PACKED_INSTALL:-0}"
GO_BUILD_TAGS="${GO_BUILD_TAGS:-gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d}"

cleanup() {
  if [[ "${KEEP_TMP}" == "1" ]]; then
    echo "[local-verify] keeping temp dir: ${TMP_ROOT}"
    return
  fi
  rm -rf "${TMP_ROOT}"
}
trap cleanup EXIT

step() {
  echo "[local-verify] $1"
}

fail() {
  echo "[local-verify] ERROR: $1" >&2
  exit 1
}

assert_file() {
  local file="$1"
  [[ -f "${file}" ]] || fail "missing file: ${file}"
}

assert_dir() {
  local dir="$1"
  [[ -d "${dir}" ]] || fail "missing directory: ${dir}"
}

step "1/8 Build local sikuli-go binary"
(
  cd "${API_DIR}"
  go build -tags "${GO_BUILD_TAGS}" -o "${API_DIR}/sikuli-go" ./cmd/sikuligrpc
)
assert_file "${API_DIR}/sikuli-go"

step "2/8 Build local Node client package"
(
  cd "${CLIENT_NODE_DIR}"
  npm run build >/dev/null
)

step "3/8 Prepare invocation mode"
CLI_RUNNER=()
if [[ "${VERIFY_PACKED_INSTALL}" == "1" ]]; then
  TARBALL_NAME="$(
    cd "${CLIENT_NODE_DIR}" &&
      npm pack --silent | tail -n 1
  )"
  TARBALL_PATH="${CLIENT_NODE_DIR}/${TARBALL_NAME}"
  assert_file "${TARBALL_PATH}"
  echo "[local-verify] tarball=${TARBALL_PATH}"
  step "4/8 Create temp Yarn project and install local tarball"
  mkdir -p "${PROJECT_DIR}"
  (
    cd "${PROJECT_DIR}"
    yarn init -2 -y >/dev/null
    YARN_ENABLE_GLOBAL_CACHE=0 YARN_CACHE_FOLDER="${PROJECT_DIR}/.yarn/cache" yarn add "${TARBALL_PATH}" >/dev/null
    YARN_ENABLE_GLOBAL_CACHE=0 YARN_CACHE_FOLDER="${PROJECT_DIR}/.yarn/cache" yarn install >/dev/null
  )
  CLI_RUNNER=(yarn exec sikuli-go)
else
  step "4/8 Create temp project (source-mode; no registry install)"
  mkdir -p "${PROJECT_DIR}"
  printf '{}\n' > "${PROJECT_DIR}/package.json"
  CLI_RUNNER=(node "${CLIENT_NODE_DIR}/dist/src/cli.js")
fi

run_cli() {
  (
    cd "${PROJECT_DIR}"
    SIKULIGO_BINARY_PATH="${API_DIR}/sikuli-go" "${CLI_RUNNER[@]}" "$@"
  )
}

if [[ "${VERIFY_PACKED_INSTALL}" == "1" ]]; then
  (
    cd "${PROJECT_DIR}"
    yarn bin sikuli-go >/dev/null
  ) || fail "packed install did not expose sikuli-go binary (check Yarn project setup and tarball bin field)"
fi

step "5/8 Verify CLI passthrough help from local binary"
HELP_OUT="$(run_cli help)"
echo "${HELP_OUT}" | rg -q "init:js-examples" || fail "help output missing init:js-examples"
echo "${HELP_OUT}" | rg -q "init:py-examples" || fail "help output missing init:py-examples"

step "6/8 Verify init:js-examples output (.mjs only)"
(
  cd "${PROJECT_DIR}"
  printf '\n' | SIKULIGO_BINARY_PATH="${API_DIR}/sikuli-go" "${CLI_RUNNER[@]}" init:js-examples --skip-install >/dev/null
)
assert_dir "${PROJECT_DIR}/sikuli-go-demo/examples"
assert_file "${PROJECT_DIR}/sikuli-go-demo/examples/click.mjs"
if find "${PROJECT_DIR}/sikuli-go-demo/examples" -maxdepth 1 -name '*.js' | rg -q .; then
  fail "init:js-examples produced .js files; expected .mjs only"
fi

step "7/8 Verify init:py-examples output and requirements"
(
  cd "${PROJECT_DIR}"
  SIKULIGO_BINARY_PATH="${API_DIR}/sikuli-go" "${CLI_RUNNER[@]}" init:py-examples --dir py-demo --skip-install >/dev/null
)
assert_dir "${PROJECT_DIR}/py-demo/examples"
assert_file "${PROJECT_DIR}/py-demo/examples/click.py"
assert_file "${PROJECT_DIR}/py-demo/requirements.txt"
rg -q '^sikuli-go' "${PROJECT_DIR}/py-demo/requirements.txt" || fail "requirements.txt missing sikuli-go dependency"

step "8/8 Smoke-run click example and assert no known transport regressions"
if [[ "${VERIFY_PACKED_INSTALL}" == "1" ]]; then
  set +e
  (
    cd "${PROJECT_DIR}/sikuli-go-demo"
    SIKULI_DEBUG=1 SIKULIGO_BINARY_PATH="${API_DIR}/sikuli-go" yarn -s node examples/click.mjs
  ) >"${LOG_FILE}" 2>&1
  SMOKE_RC=$?
  set -e
else
  set +e
  (
    cd "${PROJECT_DIR}"
    SIKULI_DEBUG=1 SIKULIGO_BINARY_PATH="${API_DIR}/sikuli-go" CLIENT_NODE_DIR="${CLIENT_NODE_DIR}" node - <<'NODE'
const { Sikuli } = require(`${process.env.CLIENT_NODE_DIR}/dist/src/index.js`);
async function main() {
  const client = await Sikuli.connect({ address: "127.0.0.1:50051", startupTimeoutMs: 50, timeoutMs: 200 });
  await client.close();
}
main().catch((err) => {
  console.error(err && err.stack ? err.stack : String(err));
  process.exit(1);
});
NODE
  ) >"${LOG_FILE}" 2>&1
  SMOKE_RC=$?
  set -e
fi
if rg -q "unknown gRPC method: ClickOnScreen|Request message serialization failure" "${LOG_FILE}"; then
  cat "${LOG_FILE}" >&2
  fail "detected transport regression in click example"
fi
echo "[local-verify] smoke exit code=${SMOKE_RC} (non-zero is allowed in this smoke check)"

step "Done"
echo "[local-verify] all local verification checks passed"
