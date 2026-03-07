#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${THIS_DIR}/paths.sh"

TMP_ROOT="$(mktemp -d /tmp/sikuli-go-node-e2e.XXXXXX)"
API_BINARY="${TMP_ROOT}/sikuli-go"
SPAWN_SQLITE="${TMP_ROOT}/node-spawn.db"
CONNECT_SQLITE="${TMP_ROOT}/node-connect.db"
KEEP_TMP="${KEEP_TMP:-0}"
GO_BUILD_TAGS="${GO_BUILD_TAGS:-gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d}"

cleanup() {
  if [[ "${KEEP_TMP}" == "1" ]]; then
    echo "[node-e2e] keeping temp dir: ${TMP_ROOT}"
    return
  fi
  rm -rf "${TMP_ROOT}"
}
trap cleanup EXIT

step() {
  echo "[node-e2e] $1"
}

step "1/3 Build local sikuli-go API binary"
(
  cd "${API_DIR}"
  go build -tags "${GO_BUILD_TAGS}" -o "${API_BINARY}" ./cmd/sikuligrpc
)

step "2/3 Build local Node client artifacts"
(
  cd "${CLIENT_NODE_DIR}"
  npm run build >/dev/null
)

step "3/3 Run Node client E2E startup/connect scenarios"
NODE_CLIENT_INDEX="${CLIENT_NODE_DIR}/dist/src/index.js" \
NODE_CLIENT_TRANSPORT="${CLIENT_NODE_DIR}/dist/src/client.js" \
SIKULIGO_BINARY="${API_BINARY}" \
SPAWN_SQLITE="${SPAWN_SQLITE}" \
CONNECT_SQLITE="${CONNECT_SQLITE}" \
node - <<'NODE'
const { spawn } = require("node:child_process");
const net = require("node:net");
const { once } = require("node:events");

const { Sikuli } = require(process.env.NODE_CLIENT_INDEX);
const { Sikuli: Transport } = require(process.env.NODE_CLIENT_TRANSPORT);

function hasContractRegression(message) {
  return /unknown gRPC method|serialization failure/i.test(String(message || ""));
}

function freePort() {
  return new Promise((resolve, reject) => {
    const server = net.createServer();
    server.once("error", reject);
    server.listen(0, "127.0.0.1", () => {
      const addr = server.address();
      if (!addr || typeof addr === "string") {
        server.close(() => reject(new Error("failed to allocate a free port")));
        return;
      }
      const port = addr.port;
      server.close((err) => {
        if (err) {
          reject(err);
          return;
        }
        resolve(port);
      });
    });
  });
}

async function waitForReady(address, authToken, timeoutMs = 5000) {
  const transport = new Transport({
    address,
    authToken,
    timeoutMs: 300
  });
  try {
    await transport.waitForReady(timeoutMs);
  } finally {
    transport.close();
  }
}

async function assertNoContractRegression(session, label) {
  try {
    await session.listWindows("NodeE2E");
  } catch (err) {
    if (hasContractRegression(err && err.message ? err.message : String(err))) {
      throw new Error(`${label} failed with contract regression: ${err.message || err}`);
    }
  }
}

async function stopProcess(child) {
  if (!child || child.exitCode !== null) {
    return;
  }
  child.kill("SIGTERM");
  await Promise.race([
    once(child, "exit").then(() => undefined),
    new Promise((resolve) => setTimeout(resolve, 1500))
  ]);
  if (child.exitCode === null) {
    child.kill("SIGKILL");
    await once(child, "exit").catch(() => undefined);
  }
}

async function main() {
  let external = null;
  try {
    const connectPort = await freePort();
    const connectAddress = `127.0.0.1:${connectPort}`;
    const connectToken = "node-e2e-token";
    external = spawn(
      process.env.SIKULIGO_BINARY,
      [
        "-listen",
        connectAddress,
        "-admin-listen",
        "",
        "-auth-token",
        connectToken,
        "-enable-reflection=false",
        "-sqlite-path",
        process.env.CONNECT_SQLITE
      ],
      {
        stdio: "ignore",
        env: {
          ...process.env,
          SIKULI_GRPC_AUTH_TOKEN: connectToken
        }
      }
    );
    await waitForReady(connectAddress, connectToken, 5000);

    const connected = await Sikuli({
      address: connectAddress,
      authToken: connectToken,
      binaryPath: process.env.SIKULIGO_BINARY,
      startupTimeoutMs: 1500,
      timeoutMs: 500
    });
    if (connected.meta.spawnedServer) {
      throw new Error("expected connect scenario to reuse existing server (spawnedServer=false)");
    }
    await assertNoContractRegression(connected, "connect scenario");
    await connected.close();

    await stopProcess(external);
    external = null;

    const spawnPort = await freePort();
    const spawnAddress = `127.0.0.1:${spawnPort}`;
    const spawned = await Sikuli({
      address: spawnAddress,
      binaryPath: process.env.SIKULIGO_BINARY,
      sqlitePath: process.env.SPAWN_SQLITE,
      startupTimeoutMs: 8000,
      timeoutMs: 500,
      stdio: "ignore"
    });
    if (!spawned.meta.spawnedServer) {
      throw new Error("expected spawn scenario to launch a new server (spawnedServer=true)");
    }
    await assertNoContractRegression(spawned, "spawn scenario");
    await spawned.close();

    console.log("[node-e2e] node client scenarios passed");
  } finally {
    await stopProcess(external);
  }
}

main().catch((err) => {
  console.error(err && err.stack ? err.stack : String(err));
  process.exit(1);
});
NODE

echo "[node-e2e] all checks passed"
