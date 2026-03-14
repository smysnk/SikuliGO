const net = require("net");
const test = require("node:test");
const assert = require("node:assert/strict");

const { buildAdminUrls, parseListenAddress, resolveRuntimeListenConfig } = require("../src/runtime-addresses");

function listenOnEphemeralPort() {
  return new Promise((resolve, reject) => {
    const server = net.createServer();
    server.unref();
    server.once("error", reject);
    server.listen(0, "127.0.0.1", () => {
      const address = server.address();
      if (!address || typeof address === "string") {
        server.close(() => reject(new Error("Unable to allocate test port")));
        return;
      }
      resolve({
        server,
        listenAddress: `127.0.0.1:${address.port}`,
      });
    });
  });
}

test("resolveRuntimeListenConfig falls back to open ports when defaults are occupied", async () => {
  const apiServer = await listenOnEphemeralPort();
  const adminServer = await listenOnEphemeralPort();

  try {
    const resolved = await resolveRuntimeListenConfig({
      apiListen: apiServer.listenAddress,
      adminListen: adminServer.listenAddress,
      allowApiFallback: true,
      allowAdminFallback: true,
    });

    assert.equal(resolved.fallbackActive, true);
    assert.notEqual(resolved.apiListen, apiServer.listenAddress);
    assert.notEqual(resolved.adminListen, adminServer.listenAddress);
    assert.match(resolved.warning, /moved to/);

    const urls = buildAdminUrls(resolved.adminListen);
    assert.match(urls.dashboardUrl, /\/dashboard$/);
    assert.deepEqual(parseListenAddress(resolved.apiListen).host, "127.0.0.1");
  } finally {
    await new Promise((resolve) => apiServer.server.close(resolve));
    await new Promise((resolve) => adminServer.server.close(resolve));
  }
});

test("resolveRuntimeListenConfig reports a conflict when fallback is disabled", async () => {
  const apiServer = await listenOnEphemeralPort();
  const adminServer = await listenOnEphemeralPort();
  try {
    const resolved = await resolveRuntimeListenConfig({
      apiListen: apiServer.listenAddress,
      adminListen: adminServer.listenAddress,
      allowApiFallback: false,
      allowAdminFallback: true,
    });

    assert.equal(resolved.fallbackActive, false);
    assert.match(resolved.warning, /already in use/);
  } finally {
    await new Promise((resolve) => apiServer.server.close(resolve));
    await new Promise((resolve) => adminServer.server.close(resolve));
  }
});
