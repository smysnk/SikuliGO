#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${THIS_DIR}/paths.sh"
source "${THIS_DIR}/macos-ocr-env.sh"

TMP_ROOT="$(mktemp -d /tmp/sikuli-go-python-e2e.XXXXXX)"
API_BINARY="${TMP_ROOT}/sikuli-go"
SPAWN_SQLITE="${TMP_ROOT}/python-spawn.db"
CONNECT_SQLITE="${TMP_ROOT}/python-connect.db"
KEEP_TMP="${KEEP_TMP:-0}"
GO_BUILD_TAGS="${GO_BUILD_TAGS:-gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d}"

cleanup() {
  if [[ "${KEEP_TMP}" == "1" ]]; then
    echo "[python-e2e] keeping temp dir: ${TMP_ROOT}"
    return
  fi
  rm -rf "${TMP_ROOT}"
}
trap cleanup EXIT

step() {
  echo "[python-e2e] $1"
}

step "1/2 Build local sikuli-go API binary"
(
  cd "${API_DIR}"
  configure_macos_ocr_env
  go build -tags "${GO_BUILD_TAGS}" -o "${API_BINARY}" ./cmd/sikuli-go
)

step "2/2 Run Python client E2E startup/connect scenarios"
PY_E2E_CLIENT_PYTHON_DIR="${CLIENT_PYTHON_DIR}" \
SIKULI_GO_BINARY="${API_BINARY}" \
SPAWN_SQLITE="${SPAWN_SQLITE}" \
CONNECT_SQLITE="${CONNECT_SQLITE}" \
python3 - <<'PY'
from __future__ import annotations

import os
import signal
import socket
import subprocess
import time

import grpc

import sys
sys.path.insert(0, os.environ["PY_E2E_CLIENT_PYTHON_DIR"])

from generated.sikuli.v1 import sikuli_pb2 as pb
from sikuligo import Screen


def has_contract_regression(message: str) -> bool:
    lowered = message.lower()
    return "unknown method" in lowered or "serialization failure" in lowered


def free_port() -> int:
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
        sock.bind(("127.0.0.1", 0))
        return int(sock.getsockname()[1])


def stop_process(child: subprocess.Popen | None, timeout: float = 1.5) -> None:
    if child is None:
        return
    if child.poll() is not None:
        return
    child.terminate()
    try:
        child.wait(timeout=timeout)
    except subprocess.TimeoutExpired:
        child.kill()
        child.wait(timeout=timeout)


def wait_for_ready(address: str, auth_token: str, timeout_seconds: float = 5.0) -> None:
    channel = grpc.insecure_channel(address)
    try:
        grpc.channel_ready_future(channel).result(timeout=timeout_seconds)
    finally:
        channel.close()


def assert_no_contract_regression(screen, label: str) -> None:
    try:
        screen.client.list_windows(pb.AppActionRequest(name="PythonE2E"), timeout_seconds=0.5)
    except Exception as exc:
        if has_contract_regression(str(exc)):
            raise AssertionError(f"{label} failed with contract regression: {exc}") from exc


external: subprocess.Popen | None = None
try:
    connect_port = free_port()
    connect_address = f"127.0.0.1:{connect_port}"
    connect_token = "python-e2e-token"
    external = subprocess.Popen(
        [
            os.environ["SIKULI_GO_BINARY"],
            "-listen",
            connect_address,
            "-admin-listen",
            "",
            "-auth-token",
            connect_token,
            "-enable-reflection=false",
            "-sqlite-path",
            os.environ["CONNECT_SQLITE"],
        ],
        env={**os.environ, "SIKULI_GRPC_AUTH_TOKEN": connect_token},
        stdout=subprocess.DEVNULL,
        stderr=subprocess.DEVNULL,
    )
    wait_for_ready(connect_address, connect_token, timeout_seconds=5.0)

    connected = Screen(
        address=connect_address,
        auth_token=connect_token,
        binary_path=os.environ["SIKULI_GO_BINARY"],
        startup_timeout_seconds=1.5,
        timeout_seconds=0.5,
        stdio="ignore",
    )
    if connected.meta.spawned_server:
        raise AssertionError("expected connect scenario to reuse existing server (spawned_server=False)")
    assert_no_contract_regression(connected, "connect scenario")
    connected.close()

    stop_process(external)
    external = None

    spawn_port = free_port()
    spawn_address = f"127.0.0.1:{spawn_port}"
    spawned = Screen(
        address=spawn_address,
        binary_path=os.environ["SIKULI_GO_BINARY"],
        sqlite_path=os.environ["SPAWN_SQLITE"],
        startup_timeout_seconds=8.0,
        timeout_seconds=0.5,
        stdio="ignore",
    )
    if not spawned.meta.spawned_server:
        raise AssertionError("expected spawn scenario to launch a new server (spawned_server=True)")
    assert_no_contract_regression(spawned, "spawn scenario")
    spawned.close()

    print("[python-e2e] python client scenarios passed")
finally:
    stop_process(external)
PY

echo "[python-e2e] all checks passed"
