import { ChildProcess } from "node:child_process";
import { launchSikuli, LaunchOptions, stopSpawnedProcess } from "./launcher";
import { RpcMessage, Sikuli as SikuliTransport, UnaryCallOptions } from "./client";

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

export interface InputOptions {
  delayMillis?: number;
  button?: string;
}

export interface MoveMouseRequest {
  x: number;
  y: number;
  opts?: InputOptions;
}

export interface ClickRequest {
  x: number;
  y: number;
  button?: string;
  delayMillis?: number;
}

export interface TypeTextRequest {
  text: string;
  delayMillis?: number;
}

export interface LaunchResultMeta {
  address: string;
  authToken: string;
  spawnedServer: boolean;
}

/**
 * Sikuli is the transport-level API client.
 * It maps high-level SikuliX-style actions to gRPC calls against the SikuliGO API.
 */
export class Sikuli {
  private readonly transport: SikuliTransport;
  private readonly child?: ChildProcess;
  readonly meta: LaunchResultMeta;
  private closed = false;

  private constructor(client: SikuliTransport, child: ChildProcess | undefined, meta: LaunchResultMeta) {
    this.transport = client;
    this.child = child;
    this.meta = meta;
  }

  /** Force-launch a new API process and connect to it. */
  static async spawn(opts: LaunchOptions = {}): Promise<Sikuli> {
    const result = await launchSikuli({ ...opts, spawnServer: opts.spawnServer ?? true });
    return new Sikuli(result.client, result.child, {
      address: result.address,
      authToken: result.authToken,
      spawnedServer: result.spawnedServer
    });
  }

  /** Auto mode entry point (`connect` first, then spawn fallback). */
  static async launch(opts: LaunchOptions = {}): Promise<Sikuli> {
    return await Sikuli.auto(opts);
  }

  /** Connect only. Does not spawn a new API process. */
  static async connect(opts: LaunchOptions = {}): Promise<Sikuli> {
    const result = await launchSikuli({ ...opts, spawnServer: false });
    return new Sikuli(result.client, undefined, {
      address: result.address,
      authToken: result.authToken,
      spawnedServer: false
    });
  }

  /** Connect first, and spawn only when connect probe fails. */
  static async auto(opts: LaunchOptions = {}): Promise<Sikuli> {
    const probeAddress = opts.address ?? process.env.SIKULI_GRPC_ADDR ?? "127.0.0.1:50051";
    const probeTimeoutMs = 1_000;
    debugLog("launcher.auto.start", {
      probe_timeout_ms: probeTimeoutMs,
      explicit_address: opts.address ? "yes" : "no",
      env_address: process.env.SIKULI_GRPC_ADDR ? "yes" : "no"
    });
    try {
      debugLog("launcher.auto.probe.connect", {
        probe_timeout_ms: probeTimeoutMs
      });
      const connected = await Sikuli.connect({
        ...opts,
        address: probeAddress,
        startupTimeoutMs: probeTimeoutMs
      });
      debugLog("launcher.auto.probe.connected_existing", {
        spawn_attempted: "no",
        spawned_server: connected.meta.spawnedServer ? "yes" : "no"
      });
      return connected;
    } catch (err) {
      debugLog("launcher.auto.probe.failed", {
        probe_timeout_ms: probeTimeoutMs,
        reason: (err as Error)?.message ?? "connect probe failed"
      });
      debugLog("launcher.auto.spawn.start", {
        spawn_attempted: "yes",
        startup_timeout_ms: opts.startupTimeoutMs ?? undefined
      });
      const spawned = await Sikuli.spawn(opts);
      debugLog("launcher.auto.spawn.ready", {
        spawned_server: spawned.meta.spawnedServer ? "yes" : "no"
      });
      return spawned;
    }
  }

  client(): SikuliTransport {
    return this.transport;
  }

  async close(): Promise<void> {
    if (this.closed) {
      return;
    }
    this.closed = true;
    this.transport.close();
    await stopSpawnedProcess(this.child);
  }

  /** Unary passthrough: FindOnScreen RPC. */
  async findOnScreen(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return await this.transport.findOnScreen(request, opts);
  }

  /** Unary passthrough: ExistsOnScreen RPC. */
  async existsOnScreen(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return await this.transport.existsOnScreen(request, opts);
  }

  /** Unary passthrough: WaitOnScreen RPC. */
  async waitOnScreen(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return await this.transport.waitOnScreen(request, opts);
  }

  /** Unary passthrough: ClickOnScreen RPC. */
  async clickOnScreen(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return await this.transport.clickOnScreen(request, opts);
  }

  async readText(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return await this.transport.readText(request, opts);
  }

  async findText(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return await this.transport.findText(request, opts);
  }

  async moveMouse(request: MoveMouseRequest, opts?: UnaryCallOptions): Promise<void> {
    await this.transport.moveMouse(
      {
        x: request.x,
        y: request.y,
        opts: {
          delay_millis: request.opts?.delayMillis
        }
      },
      opts
    );
  }

  async click(request: ClickRequest, opts?: UnaryCallOptions): Promise<void> {
    await this.transport.click(
      {
        x: request.x,
        y: request.y,
        opts: {
          button: request.button,
          delay_millis: request.delayMillis
        }
      },
      opts
    );
  }

  async typeText(request: TypeTextRequest | string, opts?: UnaryCallOptions): Promise<void> {
    const input = typeof request === "string" ? { text: request } : request;
    await this.transport.typeText(
      {
        text: input.text,
        opts: {
          delay_millis: input.delayMillis
        }
      },
      opts
    );
  }

  async hotkey(keys: string[], opts?: UnaryCallOptions): Promise<void> {
    await this.transport.hotkey({ keys }, opts);
  }

  async openApp(request: { name: string; args?: string[] }, opts?: UnaryCallOptions): Promise<void> {
    await this.transport.openApp(
      {
        name: request.name,
        args: request.args ?? []
      },
      opts
    );
  }

  async focusApp(name: string, opts?: UnaryCallOptions): Promise<void> {
    await this.transport.focusApp({ name }, opts);
  }

  async closeApp(name: string, opts?: UnaryCallOptions): Promise<void> {
    await this.transport.closeApp({ name }, opts);
  }

  async isAppRunning(name: string, opts?: UnaryCallOptions): Promise<boolean> {
    const out = await this.transport.isAppRunning({ name }, opts);
    return Boolean((out as { running?: boolean }).running);
  }

  async listWindows(name: string, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return await this.transport.listWindows({ name }, opts);
  }
}
