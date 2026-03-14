const net = require("net");

function parseListenAddress(rawAddress) {
  const value = String(rawAddress || "").trim();
  if (!value) {
    throw new Error("Listen address is required");
  }

  if (/^\d+$/.test(value)) {
    return { host: "127.0.0.1", port: Number(value) };
  }

  const colonIndex = value.lastIndexOf(":");
  if (colonIndex < 0) {
    throw new Error(`Listen address must include a port: ${value}`);
  }

  const host = value.slice(0, colonIndex) || "127.0.0.1";
  const port = Number(value.slice(colonIndex + 1));
  if (!Number.isInteger(port) || port <= 0 || port > 65535) {
    throw new Error(`Invalid listen address port: ${value}`);
  }

  return {
    host,
    port,
  };
}

function formatListenAddress(address) {
  return `${address.host}:${address.port}`;
}

function buildAdminUrls(listenAddress) {
  const baseUrl = `http://${listenAddress}`;
  return {
    dashboardUrl: `${baseUrl}/dashboard`,
    healthUrl: `${baseUrl}/healthz`,
    metricsUrl: `${baseUrl}/metrics`,
    sessionUrl: `${baseUrl}/sessions`,
    snapshotUrl: `${baseUrl}/snapshot`,
  };
}

function isListenAvailable(rawAddress) {
  return new Promise((resolve, reject) => {
    let parsedAddress;
    try {
      parsedAddress = parseListenAddress(rawAddress);
    } catch (error) {
      reject(error);
      return;
    }

    const server = net.createServer();
    server.unref();
    server.once("error", (error) => {
      if (error && (error.code === "EADDRINUSE" || error.code === "EACCES")) {
        resolve(false);
        return;
      }
      reject(error);
    });
    server.listen(parsedAddress.port, parsedAddress.host, () => {
      server.close(() => resolve(true));
    });
  });
}

function findOpenListen(rawAddress) {
  return new Promise((resolve, reject) => {
    let parsedAddress;
    try {
      parsedAddress = parseListenAddress(rawAddress);
    } catch (error) {
      reject(error);
      return;
    }

    const server = net.createServer();
    server.unref();
    server.once("error", reject);
    server.listen(0, parsedAddress.host, () => {
      const nextAddress = server.address();
      if (!nextAddress || typeof nextAddress === "string") {
        server.close(() => reject(new Error("Unable to allocate an open port")));
        return;
      }
      const resolvedListen = formatListenAddress({
        host: parsedAddress.host,
        port: nextAddress.port,
      });
      server.close(() => resolve(resolvedListen));
    });
  });
}

async function resolveRuntimeListenConfig(options) {
  const {
    apiListen,
    adminListen,
    allowApiFallback = true,
    allowAdminFallback = true,
  } = options;

  const apiAvailable = await isListenAvailable(apiListen);
  const adminAvailable = await isListenAvailable(adminListen);

  if (apiAvailable && adminAvailable) {
    return {
      apiListen,
      adminListen,
      fallbackActive: false,
      warning: "",
    };
  }

  if ((!apiAvailable && !allowApiFallback) || (!adminAvailable && !allowAdminFallback)) {
    const conflicts = [];
    if (!apiAvailable) {
      conflicts.push(`gRPC listen address ${apiListen}`);
    }
    if (!adminAvailable) {
      conflicts.push(`admin listen address ${adminListen}`);
    }
    return {
      apiListen,
      adminListen,
      fallbackActive: false,
      warning: `${conflicts.join(" and ")} is already in use`,
    };
  }

  const nextApiListen = apiAvailable ? apiListen : await findOpenListen(apiListen);
  const nextAdminListen = adminAvailable ? adminListen : await findOpenListen(adminListen);
  const warnings = [];
  if (!apiAvailable) {
    warnings.push(`gRPC listen moved to ${nextApiListen}`);
  }
  if (!adminAvailable) {
    warnings.push(`admin listen moved to ${nextAdminListen}`);
  }

  return {
    apiListen: nextApiListen,
    adminListen: nextAdminListen,
    fallbackActive: nextApiListen !== apiListen || nextAdminListen !== adminListen,
    warning: warnings.join("; "),
  };
}

module.exports = {
  buildAdminUrls,
  findOpenListen,
  formatListenAddress,
  isListenAvailable,
  parseListenAddress,
  resolveRuntimeListenConfig,
};
