import { ChildProcess, spawn } from "node:child_process";
import { randomBytes } from "node:crypto";
import net from "node:net";
import { SikuliGrpcClient, SikuliClientOptions } from "./client";
import { resolveSikuliGrpcBinary } from "./binary";

export interface LaunchOptions extends SikuliClientOptions {
  spawnServer?: boolean;
  startupTimeoutMs?: number;
  binaryPath?: string;
  adminListen?: string;
  serverArgs?: string[];
  stdio?: "ignore" | "pipe" | "inherit";
}

export interface LaunchResult {
  address: string;
  authToken: string;
  client: SikuliGrpcClient;
  child?: ChildProcess;
  spawnedServer: boolean;
}

const DEFAULT_STARTUP_TIMEOUT_MS = 10_000;

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

async function waitForStartup(client: SikuliGrpcClient, child: ChildProcess, timeoutMs: number): Promise<void> {
  let rejected = false;
  await Promise.race([
    client.waitForReady(timeoutMs),
    new Promise<never>((_, reject) => {
      child.once("exit", (code, signal) => {
        rejected = true;
        reject(
          new Error(
            `sikuligrpc exited before startup completed (code=${code ?? "nil"} signal=${signal ?? "nil"})`
          )
        );
      });
    })
  ]);
  if (rejected) {
    throw new Error("sikuligrpc exited before ready");
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

export async function launchClient(opts: LaunchOptions = {}): Promise<LaunchResult> {
  const spawnServer = opts.spawnServer !== false;
  const startupTimeoutMs = opts.startupTimeoutMs ?? DEFAULT_STARTUP_TIMEOUT_MS;
  const address =
    opts.address ||
    process.env.SIKULI_GRPC_ADDR ||
    (spawnServer ? `127.0.0.1:${await findOpenPort()}` : "127.0.0.1:50051");
  const authToken = opts.authToken || process.env.SIKULI_GRPC_AUTH_TOKEN || "";

  if (!spawnServer) {
    const client = new SikuliGrpcClient({
      address,
      authToken,
      traceId: opts.traceId,
      timeoutMs: opts.timeoutMs,
      protoPath: opts.protoPath,
      credentials: opts.credentials
    });
    await client.waitForReady(startupTimeoutMs);
    return {
      address,
      authToken,
      client,
      spawnedServer: false
    };
  }

  const binaryPath = resolveSikuliGrpcBinary(opts.binaryPath);
  const token = authToken || randomBytes(24).toString("hex");
  const serverArgs = [
    "-listen",
    address,
    "-admin-listen",
    opts.adminListen ?? "",
    "-auth-token",
    token,
    "-enable-reflection=false",
    ...(opts.serverArgs ?? [])
  ];

  const child = spawn(binaryPath, serverArgs, {
    stdio: opts.stdio ?? "ignore",
    env: {
      ...process.env,
      SIKULI_GRPC_AUTH_TOKEN: token
    }
  });
  const unwire = wireShutdown(child);

  const client = new SikuliGrpcClient({
    address,
    authToken: token,
    traceId: opts.traceId,
    timeoutMs: opts.timeoutMs,
    protoPath: opts.protoPath,
    credentials: opts.credentials
  });

  try {
    await waitForStartup(client, child, startupTimeoutMs);
  } catch (err) {
    await stopSpawnedProcess(child);
    client.close();
    unwire.forEach((fn) => fn());
    throw err;
  }

  child.once("exit", () => {
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
