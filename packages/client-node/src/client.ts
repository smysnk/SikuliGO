import fs from "node:fs";
import path from "node:path";
import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";

const DEFAULT_ADDR = "127.0.0.1:50051";
const DEFAULT_TIMEOUT_MS = 5000;
const TRACE_HEADER = "x-trace-id";

export type MatcherEngine = "template" | "orb" | "akaze" | "brisk" | "kaze" | "sift" | "hybrid";
const SCREEN_SEARCH_METHODS = new Set(["FindOnScreen", "ExistsOnScreen", "WaitOnScreen", "ClickOnScreen"]);

export type RpcMessage = Record<string, unknown>;

export interface SikuliOptions {
  address?: string;
  authToken?: string;
  traceId?: string;
  timeoutMs?: number;
  matcherEngine?: MatcherEngine;
  protoPath?: string;
  credentials?: grpc.ChannelCredentials;
}

export interface UnaryCallOptions {
  timeoutMs?: number;
  matcherEngine?: MatcherEngine;
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

function normalizeMatcherEngine(raw: string | undefined): MatcherEngine {
  const normalized = String(raw ?? "").trim().toLowerCase();
  if (
    normalized === "orb" ||
    normalized === "akaze" ||
    normalized === "brisk" ||
    normalized === "kaze" ||
    normalized === "sift" ||
    normalized === "hybrid"
  ) {
    return normalized;
  }
  return "template";
}

function resolveDefaultProtoPath(): string {
  const candidates = [
    path.resolve(__dirname, "../../proto/sikuli/v1/sikuli.proto"),
    path.resolve(__dirname, "../proto/sikuli/v1/sikuli.proto"),
    path.resolve(process.cwd(), "proto/sikuli/v1/sikuli.proto"),
    path.resolve(process.cwd(), "packages/client-node/proto/sikuli/v1/sikuli.proto"),
    path.resolve(process.cwd(), "packages/api/proto/sikuli/v1/sikuli.proto")
  ];
  for (const candidate of candidates) {
    if (fs.existsSync(candidate)) {
      return candidate;
    }
  }
  throw new Error(
    `Unable to resolve sikuli.proto. Tried: ${candidates.join(", ")}. ` +
      "Rebuild package artifacts before running."
  );
}

function serviceConstructorFromProto(protoPath: string): grpc.ServiceClientConstructor {
  const packageDefinition = protoLoader.loadSync(protoPath, {
    includeDirs: [path.dirname(protoPath)],
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

function matcherEngineToProtoValue(engine: MatcherEngine): number {
  switch (engine) {
    case "orb":
      return 2;
    case "akaze":
      return 4;
    case "brisk":
      return 5;
    case "kaze":
      return 6;
    case "sift":
      return 7;
    case "hybrid":
      return 3;
    default:
      return 1;
  }
}

function withMatcherEngine(methodName: string, request: RpcMessage, engine: MatcherEngine): RpcMessage {
  const req: RpcMessage = { ...request };
  const protoValue = matcherEngineToProtoValue(engine);
  if (methodName === "Find" || methodName === "FindAll") {
    const current = req.matcher_engine;
    if (current === undefined || current === null || current === 0 || current === "MATCHER_ENGINE_UNSPECIFIED") {
      req.matcher_engine = protoValue;
    }
    return req;
  }
  if (SCREEN_SEARCH_METHODS.has(methodName)) {
    const opts =
      req.opts && typeof req.opts === "object"
        ? ({ ...(req.opts as RpcMessage) } as RpcMessage)
        : ({} as RpcMessage);
    const current = opts.matcher_engine;
    if (current === undefined || current === null || current === 0 || current === "MATCHER_ENGINE_UNSPECIFIED") {
      opts.matcher_engine = protoValue;
    }
    req.opts = opts;
  }
  return req;
}

function screenRequestDebugFields(methodName: string, request: RpcMessage): Record<string, unknown> {
  if (!SCREEN_SEARCH_METHODS.has(methodName)) {
    return {};
  }
  const req = request as Record<string, unknown>;
  const pattern = (req.pattern ?? {}) as Record<string, unknown>;
  const image = (pattern.image ?? {}) as Record<string, unknown>;
  const opts = (req.opts ?? {}) as Record<string, unknown>;
  const region = (opts.region ?? {}) as Record<string, unknown>;

  const width = Number(image.width ?? 0);
  const height = Number(image.height ?? 0);
  const rx = Number(region.x ?? 0);
  const ry = Number(region.y ?? 0);
  const rw = Number(region.w ?? 0);
  const rh = Number(region.h ?? 0);

  return {
    screen_capture: "server",
    pattern_w: Number.isFinite(width) && width > 0 ? width : undefined,
    pattern_h: Number.isFinite(height) && height > 0 ? height : undefined,
    region: Number.isFinite(rw) && Number.isFinite(rh) && rw > 0 && rh > 0 ? `${rx},${ry},${rw},${rh}` : "full_screen"
  };
}

function screenResponseDebugFields(methodName: string, response: RpcMessage): Record<string, unknown> {
  if (!SCREEN_SEARCH_METHODS.has(methodName)) {
    return {};
  }
  const resp = response as Record<string, unknown>;
  const match = (resp.match ?? {}) as Record<string, unknown>;
  const rect = (match.rect ?? {}) as Record<string, unknown>;
  const x = Number(rect.x ?? 0);
  const y = Number(rect.y ?? 0);
  const w = Number(rect.w ?? 0);
  const h = Number(rect.h ?? 0);
  if (!(Number.isFinite(w) && Number.isFinite(h) && w > 0 && h > 0)) {
    return { match_found: "no" };
  }
  return {
    match_found: "yes",
    match_rect: `${x},${y},${w},${h}`
  };
}

function isLikelyCaptureFailure(methodName: string, code: number, details: string): boolean {
  if (!SCREEN_SEARCH_METHODS.has(methodName)) {
    return false;
  }
  const lower = String(details || "").toLowerCase();
  if (lower.includes("screencapture")) {
    return true;
  }
  if (lower.includes("screen capture")) {
    return true;
  }
  if (lower.includes("capture failed")) {
    return true;
  }
  return false;
}

function missingInputDependency(details: string): string {
  const lower = String(details || "").toLowerCase();
  if (lower.includes("exec: \"cliclick\": executable file not found") || (lower.includes("cliclick") && lower.includes("executable file not found"))) {
    return "cliclick";
  }
  if (lower.includes("exec: \"xdotool\": executable file not found") || (lower.includes("xdotool") && lower.includes("executable file not found"))) {
    return "xdotool";
  }
  if (lower.includes("exec: \"powershell\": executable file not found") || (lower.includes("powershell") && lower.includes("executable file not found"))) {
    return "powershell";
  }
  return "";
}

export class Sikuli {
  private readonly client: grpc.Client & Record<string, unknown>;
  private readonly address: string;
  private readonly authToken: string;
  private readonly traceId: string;
  private readonly defaultTimeoutMs: number;
  private readonly matcherEngine: MatcherEngine;
  private readonly debugEnabled: boolean;

  constructor(opts: SikuliOptions = {}) {
    const address = opts.address ?? process.env.SIKULI_GRPC_ADDR ?? DEFAULT_ADDR;
    this.address = address;
    this.authToken = opts.authToken ?? process.env.SIKULI_GRPC_AUTH_TOKEN ?? "";
    this.traceId = opts.traceId ?? "";
    this.defaultTimeoutMs = opts.timeoutMs ?? DEFAULT_TIMEOUT_MS;
    this.matcherEngine = normalizeMatcherEngine(opts.matcherEngine ?? process.env.SIKULI_MATCHER_ENGINE);
    this.debugEnabled = /^(1|true|yes|on)$/i.test(process.env.SIKULI_DEBUG ?? "");

    const protoPath = opts.protoPath ?? resolveDefaultProtoPath();
    const serviceClientCtor = serviceConstructorFromProto(protoPath);
    const credentials = opts.credentials ?? grpc.credentials.createInsecure();
    this.client = new serviceClientCtor(address, credentials) as unknown as grpc.Client &
      Record<string, unknown>;
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
    const captureFailure = isLikelyCaptureFailure(methodName, code, details);
    const isScreenMethod = SCREEN_SEARCH_METHODS.has(methodName);
    const missingDep = missingInputDependency(details);
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
        "Set SIKULI_DEBUG=1 to log RPC and launcher details.";
      if (isScreenMethod) {
        if (captureFailure) {
          details += " Screen capture failed; verify Screen Recording permission for your terminal/IDE on macOS.";
        } else {
          details +=
            " Capture appears server-side; check server logs for find_on_screen.match.start/progress to confirm matcher timing.";
        }
      }
    } else if (isScreenMethod && captureFailure) {
      details += " Screen capture failed; verify Screen Recording permission for your terminal/IDE on macOS.";
    }
    if (missingDep === "cliclick") {
      details += ' Missing macOS input dependency "cliclick". Install with: brew install cliclick.';
    } else if (missingDep === "xdotool") {
      details += ' Missing Linux input dependency "xdotool". Install with your distro package manager.';
    } else if (missingDep === "powershell") {
      details += ' Missing Windows input dependency "powershell".';
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
    const matcherEngine = normalizeMatcherEngine(opts.matcherEngine ?? this.matcherEngine);
    const metadata = this.buildMetadata(opts.metadata);
    const wireRequest = withMatcherEngine(methodName, request, matcherEngine);
    const screenFields = screenRequestDebugFields(methodName, wireRequest);
    this.debugLog("rpc.start", {
      method: methodName,
      address: this.address,
      timeout_ms: timeoutMs,
      matcher_engine: matcherEngine,
      ...screenFields
    });

    return new Promise((resolve, reject) => {
      callFn.call(
        this.client,
        wireRequest,
        metadata,
        { deadline },
        (err: grpc.ServiceError | null, response: RpcMessage) => {
          if (err) {
            const traceValues = err.metadata?.get(TRACE_HEADER) ?? [];
            const traceId = traceValues.length > 0 ? String(traceValues[0]) : "";
            const details = err.details || err.message;
            const captureFailure = isLikelyCaptureFailure(methodName, err.code ?? grpc.status.UNKNOWN, details);
            this.debugLog("rpc.error", {
              method: methodName,
              address: this.address,
              timeout_ms: timeoutMs,
              duration_ms: Date.now() - startedAt,
              matcher_engine: matcherEngine,
              grpc_code: err.code,
              trace_id: traceId || undefined,
              capture_stage: captureFailure ? "failed" : "unknown_or_ok",
              details
            });
            if (SCREEN_SEARCH_METHODS.has(methodName)) {
              if (captureFailure) {
                this.debugLog("rpc.screen_capture.note", {
                  method: methodName,
                  hint: "capture likely failed (check find_on_screen.capture.error / capture_screen.* logs)"
                });
              } else {
                this.debugLog("rpc.matcher.note", {
                  method: methodName,
                  hint: "capture likely succeeded; inspect find_on_screen.match.start/progress logs for matcher slowdown"
                });
              }
            }
            reject(this.clientError(methodName, err, timeoutMs));
            return;
          }
          const screenResponseFields = screenResponseDebugFields(methodName, response);
          this.debugLog("rpc.ok", {
            method: methodName,
            address: this.address,
            duration_ms: Date.now() - startedAt,
            matcher_engine: matcherEngine,
            ...screenResponseFields
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
