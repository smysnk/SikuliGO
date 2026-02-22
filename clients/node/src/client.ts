import fs from "node:fs";
import path from "node:path";
import grpc from "@grpc/grpc-js";
import protoLoader from "@grpc/proto-loader";

const DEFAULT_ADDR = "127.0.0.1:50051";
const DEFAULT_TIMEOUT_MS = 5000;
const TRACE_HEADER = "x-trace-id";

export type RpcMessage = Record<string, unknown>;

export interface SikuliClientOptions {
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

export class SikuliGrpcError extends Error {
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

export class SikuliGrpcClient {
  private readonly client: grpc.Client & Record<string, unknown>;
  private readonly authToken: string;
  private readonly traceId: string;
  private readonly defaultTimeoutMs: number;

  constructor(opts: SikuliClientOptions = {}) {
    const address = opts.address ?? process.env.SIKULI_GRPC_ADDR ?? DEFAULT_ADDR;
    this.authToken = opts.authToken ?? process.env.SIKULI_GRPC_AUTH_TOKEN ?? "";
    this.traceId = opts.traceId ?? "";
    this.defaultTimeoutMs = opts.timeoutMs ?? DEFAULT_TIMEOUT_MS;

    const protoPath = opts.protoPath ?? resolveDefaultProtoPath();
    const ctor = serviceConstructorFromProto(protoPath);
    const credentials = opts.credentials ?? grpc.credentials.createInsecure();

    this.client = new ctor(address, credentials) as grpc.Client & Record<string, unknown>;
  }

  close(): void {
    this.client.close();
  }

  waitForReady(timeoutMs = DEFAULT_TIMEOUT_MS): Promise<void> {
    const deadline = new Date(Date.now() + timeoutMs);
    return new Promise((resolve, reject) => {
      this.client.waitForReady(deadline, (err?: Error | null) => {
        if (err) {
          reject(err);
          return;
        }
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

  private grpcError(err: grpc.ServiceError): SikuliGrpcError {
    const traceValues = err.metadata?.get(TRACE_HEADER) ?? [];
    const traceId = traceValues.length > 0 ? String(traceValues[0]) : undefined;
    return new SikuliGrpcError(err.code ?? grpc.status.UNKNOWN, err.details || err.message, traceId);
  }

  private unary(methodName: string, request: RpcMessage, opts: UnaryCallOptions = {}): Promise<RpcMessage> {
    const callFn = this.client[methodName] as Function | undefined;
    if (typeof callFn !== "function") {
      return Promise.reject(new Error(`unknown gRPC method: ${methodName}`));
    }

    const timeoutMs = opts.timeoutMs ?? this.defaultTimeoutMs;
    const deadline = new Date(Date.now() + timeoutMs);
    const metadata = this.buildMetadata(opts.metadata);

    return new Promise((resolve, reject) => {
      callFn.call(
        this.client,
        request,
        metadata,
        { deadline },
        (err: grpc.ServiceError | null, response: RpcMessage) => {
          if (err) {
            reject(this.grpcError(err));
            return;
          }
          resolve(response);
        }
      );
    });
  }

  find(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("Find", request, opts);
  }

  findAll(request: RpcMessage, opts?: UnaryCallOptions): Promise<RpcMessage> {
    return this.unary("FindAll", request, opts);
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
