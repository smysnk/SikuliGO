#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
EDITOR_URL="${SIKULI_GO_EDITOR_URL:-http://127.0.0.1:3000}"
EDITOR_DEV_STARTED=0
EDITOR_PID=""

cleanup() {
  if [[ "${EDITOR_DEV_STARTED}" == "1" && -n "${EDITOR_PID}" ]]; then
    kill "${EDITOR_PID}" >/dev/null 2>&1 || true
    wait "${EDITOR_PID}" >/dev/null 2>&1 || true
  fi
}

trap cleanup EXIT INT TERM

cd "${ROOT_DIR}"

if ! curl -fsS "${EDITOR_URL}" >/dev/null 2>&1; then
  yarn workspace @sikuligo/editor dev >/dev/null 2>&1 &
  EDITOR_PID="$!"
  EDITOR_DEV_STARTED=1

  for _ in $(seq 1 60); do
    if curl -fsS "${EDITOR_URL}" >/dev/null 2>&1; then
      break
    fi
    sleep 1
  done
fi

if ! curl -fsS "${EDITOR_URL}" >/dev/null 2>&1; then
  echo "Editor dev server did not become ready at ${EDITOR_URL}" >&2
  exit 1
fi

SIKULI_GO_EDITOR_URL="${EDITOR_URL}" \
SIKULI_GO_INITIAL_VIEW="editor" \
yarn workspace @sikuligo/api-electron dev
