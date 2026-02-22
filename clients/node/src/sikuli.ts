import { ChildProcess } from "node:child_process";
import { launchClient, LaunchOptions, stopSpawnedProcess } from "./launcher";
import { RpcMessage, SikuliGrpcClient, UnaryCallOptions } from "./client";

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
  private readonly client: SikuliGrpcClient;
  private readonly child?: ChildProcess;
  readonly meta: LaunchResultMeta;
  private closed = false;

  private constructor(client: SikuliGrpcClient, child: ChildProcess | undefined, meta: LaunchResultMeta) {
    this.client = client;
    this.child = child;
    this.meta = meta;
  }

  static async launch(opts: LaunchOptions = {}): Promise<Sikuli> {
    const result = await launchClient({ ...opts, spawnServer: opts.spawnServer ?? true });
    return new Sikuli(result.client, result.child, {
      address: result.address,
      authToken: result.authToken,
      spawnedServer: result.spawnedServer
    });
  }

  static async connect(opts: LaunchOptions = {}): Promise<Sikuli> {
    const result = await launchClient({ ...opts, spawnServer: false });
    return new Sikuli(result.client, undefined, {
      address: result.address,
      authToken: result.authToken,
      spawnedServer: false
    });
  }

  grpc(): SikuliGrpcClient {
    return this.client;
  }

  async close(): Promise<void> {
    if (this.closed) {
      return;
    }
    this.closed = true;
    this.client.close();
    await stopSpawnedProcess(this.child);
  }

  async find(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return await this.client.find(request, opts);
  }

  async findAll(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return await this.client.findAll(request, opts);
  }

  async readText(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return await this.client.readText(request, opts);
  }

  async findText(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return await this.client.findText(request, opts);
  }

  async moveMouse(request: MoveMouseRequest, opts?: UnaryCallOptions): Promise<void> {
    await this.client.moveMouse(
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
    await this.client.click(
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
    await this.client.typeText(
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
    await this.client.hotkey({ keys }, opts);
  }

  async openApp(request: { name: string; args?: string[] }, opts?: UnaryCallOptions): Promise<void> {
    await this.client.openApp(
      {
        name: request.name,
        args: request.args ?? []
      },
      opts
    );
  }

  async focusApp(name: string, opts?: UnaryCallOptions): Promise<void> {
    await this.client.focusApp({ name }, opts);
  }

  async closeApp(name: string, opts?: UnaryCallOptions): Promise<void> {
    await this.client.closeApp({ name }, opts);
  }

  async isAppRunning(name: string, opts?: UnaryCallOptions): Promise<boolean> {
    const out = await this.client.isAppRunning({ name }, opts);
    return Boolean((out as { running?: boolean }).running);
  }

  async listWindows(name: string, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return await this.client.listWindows({ name }, opts);
  }
}
