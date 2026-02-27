import { ChildProcess } from "node:child_process";
import { launchSikuli, LaunchOptions, stopSpawnedProcess } from "./launcher";
import { RpcMessage, Sikuli as SikuliTransport, UnaryCallOptions } from "./client";

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

  static async spawn(opts: LaunchOptions = {}): Promise<Sikuli> {
    const result = await launchSikuli({ ...opts, spawnServer: opts.spawnServer ?? true });
    return new Sikuli(result.client, result.child, {
      address: result.address,
      authToken: result.authToken,
      spawnedServer: result.spawnedServer
    });
  }

  static async launch(opts: LaunchOptions = {}): Promise<Sikuli> {
    return await Sikuli.auto(opts);
  }

  static async connect(opts: LaunchOptions = {}): Promise<Sikuli> {
    const result = await launchSikuli({ ...opts, spawnServer: false });
    return new Sikuli(result.client, undefined, {
      address: result.address,
      authToken: result.authToken,
      spawnedServer: false
    });
  }

  static async auto(opts: LaunchOptions = {}): Promise<Sikuli> {
    const probeAddress = opts.address ?? process.env.SIKULI_GRPC_ADDR ?? "127.0.0.1:50051";
    try {
      return await Sikuli.connect({
        ...opts,
        address: probeAddress,
        startupTimeoutMs: 1_000
      });
    } catch {
      return await Sikuli.spawn(opts);
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

  async findOnScreen(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return await this.transport.findOnScreen(request, opts);
  }

  async existsOnScreen(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return await this.transport.existsOnScreen(request, opts);
  }

  async waitOnScreen(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return await this.transport.waitOnScreen(request, opts);
  }

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
