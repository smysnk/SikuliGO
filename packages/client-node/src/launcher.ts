import { ChildProcess, spawn } from "node:child_process";
import { randomBytes } from "node:crypto";
import net from "node:net";
import { Sikuli as SikuliTransport, SikuliOptions } from "./client";
import { resolveSikuliBinary } from "./binary";

export interface LaunchOptions extends SikuliOptions {
  spawnServer?: boolean;
  startupTimeoutMs?: number;
  binaryPath?: string;
  adminListen?: string;
  sqlitePath?: string;
  serverArgs?: string[];
  stdio?: "ignore" | "pipe" | "inherit";
}

export interface LaunchResult {
  address: string;
  authToken: string;
  client: SikuliTransport;
  child?: ChildProcess;
  spawnedServer: boolean;
}

const DEFAULT_STARTUP_TIMEOUT_MS = 10_000;
const DEBUG_ENABLED = /^(1|true|yes|on)$/i.test(process.env.SIKULI_DEBUG ?? "");

function debugLog(message: string, fields: Record<string, unknown> = {}): void {
  if (!DEBUG_ENABLED) {
    return;
  }
  const parts = Object.entries(fields)
    .filter(([k, v]) => k !== "address" && v !== undefined && v !== null && v !== "")
    .map(([k, v]) => `${k}=${String(v)}`);
  const suffix = parts.length > 0 ? ` ${parts.join(" ")}` : "";
  // eslint-disable-next-line no-console
  console.error(`[sikuligo-debug] ${message}${suffix}`);
}

async function findOpenPort(): Promise<number> {
  return await new Promise((resolve, reject) => {
    const srv = net.createServer();
    srv.unref();
    srv.once("error", reject);
    srv.listen(0, "127.0.0.1", () => {
      const addr = srv.address();
      if (!addr || typeof addr === "string") {
        srv.close(() => reject(new Error("failed to allocate local port")));
        return;
      }
      const port = addr.port;
      srv.close((err) => {
        if (err) {
          reject(err);
          return;
        }
        resolve(port);
      });
    });
  });
}

function wireShutdown(child: ChildProcess): Array<() => void> {
  const onExit = () => {
    if (child.exitCode === null && !child.killed) {
      child.kill("SIGTERM");
    }
  };
  process.on("exit", onExit);
  process.on("SIGINT", onExit);
  process.on("SIGTERM", onExit);
  return [
    () => process.off("exit", onExit),
    () => process.off("SIGINT", onExit),
    () => process.off("SIGTERM", onExit)
  ];
}

async function waitForStartup(
  client: SikuliTransport,
  child: ChildProcess,
  timeoutMs: number,
  exitDetail?: () => string
): Promise<void> {
  let rejected = false;
  await Promise.race([
    client.waitForReady(timeoutMs),
    new Promise<never>((_, reject) => {
      child.once("exit", (code, signal) => {
        rejected = true;
        const detail = exitDetail ? exitDetail() : "";
        const suffix = detail ? ` stderr_tail=${JSON.stringify(detail)}` : "";
        reject(
          new Error(
            `sikuligo exited before startup completed (code=${code ?? "nil"} signal=${signal ?? "nil"})${suffix}`
          )
        );
      });
    })
  ]);
  if (rejected) {
    throw new Error("sikuligo exited before ready");
  }
}

export async function stopSpawnedProcess(child?: ChildProcess, timeoutMs = 3_000): Promise<void> {
  if (!child || child.exitCode !== null) {
    return;
  }
  child.kill("SIGTERM");
  await Promise.race([
    new Promise<void>((resolve) => {
      child.once("exit", () => resolve());
    }),
    new Promise<void>((resolve) => {
      setTimeout(resolve, timeoutMs);
    })
  ]);
  if (child.exitCode === null) {
    child.kill("SIGKILL");
  }
}

export async function launchSikuli(opts: LaunchOptions = {}): Promise<LaunchResult> {
  const spawnServer = opts.spawnServer !== false;
  const startupTimeoutMs = opts.startupTimeoutMs ?? DEFAULT_STARTUP_TIMEOUT_MS;
  const explicitAddress = opts.address || process.env.SIKULI_GRPC_ADDR || "";
  const address = explicitAddress || (spawnServer ? `127.0.0.1:${await findOpenPort()}` : "127.0.0.1:50051");
  const authToken = opts.authToken || process.env.SIKULI_GRPC_AUTH_TOKEN || "";

  if (!spawnServer) {
    debugLog("launcher.connect.start", {
      address,
      startup_timeout_ms: startupTimeoutMs
    });
    const client = new SikuliTransport({
      address,
      authToken,
      traceId: opts.traceId,
      timeoutMs: opts.timeoutMs,
      credentials: opts.credentials
    });
    await client.waitForReady(startupTimeoutMs);
    debugLog("launcher.connect.ready", { address });
    return {
      address,
      authToken,
      client,
      spawnedServer: false
    };
  }

  const binaryPath = resolveSikuliBinary(opts.binaryPath);
  const token = authToken || randomBytes(24).toString("hex");
  const sqlitePath = opts.sqlitePath || process.env.SIKULIGO_SQLITE_PATH || "sikuligo.db";
  const stdioMode: "ignore" | "pipe" | "inherit" = opts.stdio ?? (DEBUG_ENABLED ? "inherit" : "ignore");
  const serverArgs = [
    "-listen",
    address,
    "-admin-listen",
    opts.adminListen ?? "",
    "-auth-token",
    token,
    "-enable-reflection=false",
    "-sqlite-path",
    sqlitePath,
    ...(opts.serverArgs ?? [])
  ];
  debugLog("launcher.spawn.start", {
    binary: binaryPath,
    address,
    admin_listen: opts.adminListen ?? "",
    sqlite_path: sqlitePath,
    stdio: stdioMode,
    startup_timeout_ms: startupTimeoutMs
  });

  const child = spawn(binaryPath, serverArgs, {
    stdio: stdioMode,
    env: {
      ...process.env,
      SIKULI_GRPC_AUTH_TOKEN: token
    }
  });
  let stderrTail = "";
  const appendStderr = (chunk: Buffer | string) => {
    const text = typeof chunk === "string" ? chunk : chunk.toString("utf8");
    stderrTail = `${stderrTail}${text}`;
    const max = 2000;
    if (stderrTail.length > max) {
      stderrTail = stderrTail.slice(stderrTail.length - max);
    }
  };
  if (child.stderr) {
    child.stderr.on("data", appendStderr);
  }
  const unwire = wireShutdown(child);

  const client = new SikuliTransport({
    address,
    authToken: token,
    traceId: opts.traceId,
    timeoutMs: opts.timeoutMs,
    credentials: opts.credentials
  });

  try {
    await waitForStartup(client, child, startupTimeoutMs, () => stderrTail.trim());
    debugLog("launcher.spawn.ready", {
      address,
      pid: child.pid ?? "unknown"
    });
  } catch (err) {
    const canFallbackToConnect = explicitAddress !== "";
    if (canFallbackToConnect) {
      try {
        await client.waitForReady(Math.max(250, Math.min(startupTimeoutMs, 1_500)));
        debugLog("launcher.spawn.fallback_connect", {
          address,
          reason: (err as Error)?.message ?? "spawn failed"
        });
        unwire.forEach((fn) => fn());
        return {
          address,
          authToken: opts.authToken || process.env.SIKULI_GRPC_AUTH_TOKEN || "",
          client,
          spawnedServer: false
        };
      } catch {
        // Fall through to original failure handling.
      }
    }
    await stopSpawnedProcess(child);
    client.close();
    unwire.forEach((fn) => fn());
    throw err;
  }

  child.once("exit", () => {
    debugLog("launcher.spawn.exit", {
      address,
      pid: child.pid ?? "unknown",
      code: child.exitCode ?? "nil"
    });
    unwire.forEach((fn) => fn());
  });

  return {
    address,
    authToken: token,
    client,
    child,
    spawnedServer: true
  };
}
