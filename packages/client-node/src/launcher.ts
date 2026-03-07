import { ChildProcess, spawn } from "node:child_process";
import { randomBytes } from "node:crypto";
import fs from "node:fs";
import net from "node:net";
import path from "node:path";
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
  console.error(`[sikuli-go-debug] ${message}${suffix}`);
}

function mergeRuntimePath(currentPath: string | undefined): { pathValue: string; added: string[] } {
  const delimiter = process.platform === "win32" ? ";" : ":";
  const existing = String(currentPath || "")
    .split(delimiter)
    .map((part) => part.trim())
    .filter((part) => part.length > 0);
  const keyFor = (part: string): string => (process.platform === "win32" ? part.toLowerCase() : part);
  const seen = new Set(existing.map(keyFor));
  const extra: string[] = [];
  const homeLocal = process.env.HOME ? `${process.env.HOME}/.local/bin` : "";
  const defaults =
    process.platform === "darwin"
      ? ["/opt/homebrew/bin", "/usr/local/bin", "/usr/bin", "/bin", homeLocal]
      : process.platform === "linux"
      ? ["/usr/local/bin", "/usr/bin", "/bin", homeLocal]
      : [];
  for (const candidate of defaults) {
    const trimmed = candidate.trim();
    if (!trimmed) {
      continue;
    }
    const key = keyFor(trimmed);
    if (seen.has(key)) {
      continue;
    }
    seen.add(key);
    existing.push(trimmed);
    extra.push(trimmed);
  }
  return {
    pathValue: existing.join(delimiter),
    added: extra
  };
}

function splitPathList(pathValue: string): string[] {
  const delimiter = process.platform === "win32" ? ";" : ":";
  return String(pathValue || "")
    .split(delimiter)
    .map((part) => part.trim())
    .filter((part) => part.length > 0);
}

function isExecutableFile(candidatePath: string): boolean {
  try {
    const stat = fs.statSync(candidatePath);
    if (!stat.isFile()) {
      return false;
    }
    if (process.platform === "win32") {
      fs.accessSync(candidatePath, fs.constants.F_OK);
    } else {
      fs.accessSync(candidatePath, fs.constants.F_OK | fs.constants.X_OK);
    }
    return true;
  } catch {
    return false;
  }
}

function resolveCommandFromPath(cmd: string, pathValue: string): string | undefined {
  if (!cmd) {
    return undefined;
  }
  if (cmd.includes(path.sep)) {
    return isExecutableFile(cmd) ? cmd : undefined;
  }
  for (const dir of splitPathList(pathValue)) {
    const candidate = path.join(dir, cmd);
    if (isExecutableFile(candidate)) {
      return candidate;
    }
  }
  return undefined;
}

function startupStatePath(): string {
  const xdg = process.env.XDG_CONFIG_HOME?.trim();
  if (xdg) {
    return path.join(xdg, "sikuligo", "startup-state.json");
  }
  const home = process.env.HOME?.trim();
  if (home) {
    return path.join(home, ".config", "sikuligo", "startup-state.json");
  }
  return path.resolve(process.cwd(), ".sikuligo-startup-state.json");
}

function suppressCliclickPromptEnabled(statePath: string): boolean {
  try {
    const raw = fs.readFileSync(statePath, "utf8");
    const parsed = JSON.parse(raw) as { suppress_cliclick_prompt?: unknown };
    return parsed.suppress_cliclick_prompt === true;
  } catch {
    return false;
  }
}

function sleep(ms: number): Promise<void> {
  return new Promise((resolve) => {
    setTimeout(resolve, ms);
  });
}

function splitHostPort(address: string): { host: string; port: number } | undefined {
  const value = String(address || "").trim();
  if (!value) {
    return undefined;
  }
  const idx = value.lastIndexOf(":");
  if (idx < 0) {
    return undefined;
  }
  const hostRaw = value.slice(0, idx).trim();
  const portRaw = value.slice(idx + 1).trim();
  const port = Number(portRaw);
  if (!Number.isInteger(port) || port <= 0 || port > 65535) {
    return undefined;
  }
  const host = hostRaw.length > 0 ? hostRaw : "127.0.0.1";
  return { host, port };
}

async function isAddressReachable(address: string, timeoutMs: number): Promise<boolean> {
  const parsed = splitHostPort(address);
  if (!parsed) {
    return false;
  }
  return await new Promise<boolean>((resolve) => {
    const socket = net.createConnection({ host: parsed.host, port: parsed.port });
    let settled = false;
    const finish = (value: boolean): void => {
      if (settled) {
        return;
      }
      settled = true;
      socket.destroy();
      resolve(value);
    };
    socket.setTimeout(timeoutMs);
    socket.once("connect", () => finish(true));
    socket.once("timeout", () => finish(false));
    socket.once("error", () => finish(false));
  });
}

async function waitForCliclickGate(child: ChildProcess, pathValue: string, address: string): Promise<void> {
  const statePath = startupStatePath();
  const pollMs = 500;
  let lastLog = 0;
  debugLog("launcher.cliclick.poll.start", {
    state_path: statePath,
    poll_ms: pollMs
  });
  for (;;) {
    if (child.exitCode !== null) {
      throw new Error(
        `sikuli-go exited while waiting for cliclick dependency gate (code=${child.exitCode ?? "nil"})`
      );
    }
    const binaryPath = resolveCommandFromPath("cliclick", pathValue);
    if (binaryPath) {
      debugLog("launcher.cliclick.poll.ready", { binary: binaryPath });
      return;
    }
    if (suppressCliclickPromptEnabled(statePath)) {
      debugLog("launcher.cliclick.poll.suppressed", {
        state_path: statePath
      });
      return;
    }
    if (await isAddressReachable(address, 150)) {
      debugLog("launcher.cliclick.poll.server_ready_without_cliclick", {
        reason: "server_ready"
      });
      return;
    }
    const now = Date.now();
    if (now - lastLog >= 5000) {
      lastLog = now;
      debugLog("launcher.cliclick.poll.wait", {
        state_path: statePath
      });
    }
    await sleep(pollMs);
  }
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
            `sikuli-go exited before startup completed (code=${code ?? "nil"} signal=${signal ?? "nil"})${suffix}`
          )
        );
      });
    })
  ]);
  if (rejected) {
    throw new Error("sikuli-go exited before ready");
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
  const startupTimeoutRequestedMs = opts.startupTimeoutMs ?? DEFAULT_STARTUP_TIMEOUT_MS;
  const stdioMode: "ignore" | "pipe" | "inherit" = opts.stdio ?? (DEBUG_ENABLED ? "inherit" : "ignore");
  const startupTimeoutMs = startupTimeoutRequestedMs;
  const explicitAddress = opts.address || process.env.SIKULI_GRPC_ADDR || "";
  const address = explicitAddress || (spawnServer ? `127.0.0.1:${await findOpenPort()}` : "127.0.0.1:50051");
  const authToken = opts.authToken || process.env.SIKULI_GRPC_AUTH_TOKEN || "";
  debugLog("launcher.start", {
    mode: spawnServer ? "spawn" : "connect",
    explicit_address: explicitAddress ? "yes" : "no",
    startup_timeout_requested_ms: startupTimeoutRequestedMs,
    startup_timeout_ms: startupTimeoutMs
  });

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
    startup_timeout_requested_ms: startupTimeoutRequestedMs,
    startup_timeout_ms: startupTimeoutMs
  });
  const mergedPath = mergeRuntimePath(process.env.PATH);
  debugLog("launcher.spawn.path", {
    path_augmented: mergedPath.added.length > 0 ? "yes" : "no",
    path_added: mergedPath.added.join(",")
  });

  const child = spawn(binaryPath, serverArgs, {
    stdio: stdioMode,
    env: {
      ...process.env,
      PATH: mergedPath.pathValue || process.env.PATH || "",
      SIKULI_GRPC_AUTH_TOKEN: token,
      ...(DEBUG_ENABLED ? { SIKULI_DEBUG: "1" } : {})
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
    if (process.platform === "darwin" && stdioMode === "inherit" && Boolean(process.stdin.isTTY)) {
      await waitForCliclickGate(child, mergedPath.pathValue, address);
    }
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
