import fs from "node:fs";
import path from "node:path";
import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";

const DEFAULT_ADDR = "127.0.0.1:50051";
const DEFAULT_TIMEOUT_MS = 5000;
const TRACE_HEADER = "x-trace-id";

export type RpcMessage = Record<string, unknown>;

export interface SikuliOptions {
  address?: string;
  authToken?: string;
  traceId?: string;
  timeoutMs?: number;
  protoPath?: string;
  credentials?: grpc.ChannelCredentials;
}

export interface UnaryCallOptions {
  timeoutMs?: number;
  metadata?: Record<string, string>;
}

export class SikuliError extends Error {
  readonly code: number;
  readonly details: string;
  readonly traceId?: string;

  constructor(code: number, details: string, traceId?: string) {
    const suffix = traceId ? ` trace_id=${traceId}` : "";
    super(`grpc_code=${code} details=${details}${suffix}`);
    this.code = code;
    this.details = details;
    this.traceId = traceId;
  }
}

function resolveDefaultProtoPath(): string {
  const candidates = [
    path.resolve(__dirname, "../../../proto/sikuli/v1/sikuli.proto"),
    path.resolve(__dirname, "../../../../proto/sikuli/v1/sikuli.proto"),
    path.resolve(process.cwd(), "proto/sikuli/v1/sikuli.proto"),
    path.resolve(process.cwd(), "../../proto/sikuli/v1/sikuli.proto")
  ];
  for (const candidate of candidates) {
    if (fs.existsSync(candidate)) {
      return candidate;
    }
  }
  return candidates[0];
}

function serviceConstructorFromProto(protoPath: string): grpc.ServiceClientConstructor {
  const packageDefinition = protoLoader.loadSync(protoPath, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
  });
  const root = grpc.loadPackageDefinition(packageDefinition) as any;
  const serviceCtor = root?.sikuli?.v1?.SikuliService;
  if (!serviceCtor) {
    throw new Error(`SikuliService not found in proto: ${protoPath}`);
  }
  return serviceCtor as grpc.ServiceClientConstructor;
}

export class Sikuli {
  private readonly client: grpc.Client & Record<string, unknown>;
  private readonly address: string;
  private readonly authToken: string;
  private readonly traceId: string;
  private readonly defaultTimeoutMs: number;
  private readonly debugEnabled: boolean;

  constructor(opts: SikuliOptions = {}) {
    const address = opts.address ?? process.env.SIKULI_GRPC_ADDR ?? DEFAULT_ADDR;
    this.address = address;
    this.authToken = opts.authToken ?? process.env.SIKULI_GRPC_AUTH_TOKEN ?? "";
    this.traceId = opts.traceId ?? "";
    this.defaultTimeoutMs = opts.timeoutMs ?? DEFAULT_TIMEOUT_MS;
    this.debugEnabled = /^(1|true|yes|on)$/i.test(process.env.SIKULI_DEBUG ?? "");

    const protoPath = opts.protoPath ?? resolveDefaultProtoPath();
    const ctor = serviceConstructorFromProto(protoPath);
    const credentials = opts.credentials ?? grpc.credentials.createInsecure();

    this.client = new ctor(address, credentials) as grpc.Client & Record<string, unknown>;
  }

  private debugLog(message: string, fields: Record<string, unknown> = {}): void {
    if (!this.debugEnabled) {
      return;
    }
    const parts = Object.entries(fields)
      .filter(([k, v]) => k !== "address" && v !== undefined && v !== null && v !== "")
      .map(([k, v]) => `${k}=${String(v)}`);
    const suffix = parts.length > 0 ? ` ${parts.join(" ")}` : "";
    // eslint-disable-next-line no-console
    console.error(`[sikuligo-debug] ${message}${suffix}`);
  }

  close(): void {
    this.debugLog("grpc.close", { address: this.address });
    this.client.close();
  }

  waitForReady(timeoutMs = DEFAULT_TIMEOUT_MS): Promise<void> {
    const startedAt = Date.now();
    this.debugLog("grpc.wait_for_ready.start", { address: this.address, timeout_ms: timeoutMs });
    const deadline = new Date(Date.now() + timeoutMs);
    return new Promise((resolve, reject) => {
      this.client.waitForReady(deadline, (err?: Error | null) => {
        if (err) {
          this.debugLog("grpc.wait_for_ready.error", {
            address: this.address,
            timeout_ms: timeoutMs,
            duration_ms: Date.now() - startedAt,
            error: err.message
          });
          reject(err);
          return;
        }
        this.debugLog("grpc.wait_for_ready.ok", {
          address: this.address,
          duration_ms: Date.now() - startedAt
        });
        resolve();
      });
    });
  }

  private buildMetadata(extra: Record<string, string> = {}): grpc.Metadata {
    const md = new grpc.Metadata();
    if (this.authToken) {
      md.set("x-api-key", this.authToken);
    }
    if (this.traceId) {
      md.set(TRACE_HEADER, this.traceId);
    }
    for (const [k, v] of Object.entries(extra)) {
      if (v) {
        md.set(k, v);
      }
    }
    return md;
  }

  private clientError(methodName: string, err: grpc.ServiceError, timeoutMs: number): SikuliError {
    const traceValues = err.metadata?.get(TRACE_HEADER) ?? [];
    const traceId = traceValues.length > 0 ? String(traceValues[0]) : undefined;
    const code = err.code ?? grpc.status.UNKNOWN;
    let details = err.details || err.message;
    if (
      code === grpc.status.UNIMPLEMENTED &&
      typeof details === "string" &&
      details.toLowerCase().includes("unknown method")
    ) {
      details +=
        `; server does not implement ${methodName}. ` +
        "This usually means the sikuligo binary is older than this client. " +
        "Build/update sikuligo or set SIKULIGO_BINARY_PATH to a current binary.";
    }
    if (code === grpc.status.DEADLINE_EXCEEDED) {
      details +=
        `; client deadline=${timeoutMs}ms. ` +
        "Set SIKULI_DEBUG=1 to log RPC and launcher details. " +
        "If this is ClickOnScreen/FindOnScreen on macOS, verify Screen Recording permission for your terminal/IDE.";
    }
    return new SikuliError(code, details, traceId);
  }

  private unary(methodName: string, request: RpcMessage, opts: UnaryCallOptions = {}): Promise<RpcMessage> {
    const callFn = this.client[methodName] as Function | undefined;
    if (typeof callFn !== "function") {
      return Promise.reject(new Error(`unknown gRPC method: ${methodName}`));
    }

    const timeoutMs = opts.timeoutMs ?? this.defaultTimeoutMs;
    const startedAt = Date.now();
    const deadline = new Date(Date.now() + timeoutMs);
    const metadata = this.buildMetadata(opts.metadata);
    this.debugLog("rpc.start", {
      method: methodName,
      address: this.address,
      timeout_ms: timeoutMs
    });

    return new Promise((resolve, reject) => {
      callFn.call(
        this.client,
        request,
        metadata,
        { deadline },
        (err: grpc.ServiceError | null, response: RpcMessage) => {
          if (err) {
            this.debugLog("rpc.error", {
              method: methodName,
              address: this.address,
              timeout_ms: timeoutMs,
              duration_ms: Date.now() - startedAt,
              grpc_code: err.code,
              details: err.details || err.message
            });
            reject(this.clientError(methodName, err, timeoutMs));
            return;
          }
          this.debugLog("rpc.ok", {
            method: methodName,
            address: this.address,
            duration_ms: Date.now() - startedAt
          });
          resolve(response);
        }
      );
    });
  }

  findOnScreen(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("FindOnScreen", request, opts);
  }

  existsOnScreen(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("ExistsOnScreen", request, opts);
  }

  waitOnScreen(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("WaitOnScreen", request, opts);
  }

  clickOnScreen(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("ClickOnScreen", request, opts);
  }

  readText(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("ReadText", request, opts);
  }

  findText(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("FindText", request, opts);
  }

  moveMouse(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("MoveMouse", request, opts);
  }

  click(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("Click", request, opts);
  }

  typeText(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("TypeText", request, opts);
  }

  hotkey(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("Hotkey", request, opts);
  }

  openApp(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("OpenApp", request, opts);
  }

  focusApp(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("FocusApp", request, opts);
  }

  closeApp(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("CloseApp", request, opts);
  }

  isAppRunning(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("IsAppRunning", request, opts);
  }

  listWindows(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("ListWindows", request, opts);
  }
}
