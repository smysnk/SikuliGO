const fs = require("fs");
const net = require("net");
const path = require("path");
const { fork, spawn, spawnSync } = require("child_process");

const { getExampleById: defaultGetExampleById } = require("./examples");
const { createInstrumentedNodeExample } = require("./execution/node-transform");

const MAX_SESSION_HISTORY = 30;
const NODE_RUNNER_PATH = path.join(__dirname, "execution", "node-runner.js");
const PYTHON_RUNNER_PATH = path.join(__dirname, "execution", "python-runner.py");
const PYTHON_TRANSFORM_PATH = path.join(__dirname, "execution", "python-transform.py");

function nowIso() {
  return new Date().toISOString();
}

function emitToHandler(handler, payload) {
  if (typeof handler === "function") {
    handler(payload);
  }
}

function cloneRuntimeCall(runtimeCall) {
  if (!runtimeCall) {
    return null;
  }
  return { ...runtimeCall };
}

function updateRuntimeCallRecord(existing, payload, state) {
  return {
    callId: payload.callId || existing?.callId || `runtime-${Date.now()}`,
    method: payload.method || existing?.method || "Runtime Call",
    state,
    line: payload.line ?? existing?.line ?? null,
    column: payload.column ?? existing?.column ?? null,
    statementId: payload.statementId ?? existing?.statementId ?? null,
    startedAt: existing?.startedAt || payload.emittedAt || nowIso(),
    endedAt: state === "running" ? null : payload.emittedAt || nowIso(),
    durationMs:
      typeof payload.durationMs === "number"
        ? payload.durationMs
        : typeof existing?.durationMs === "number"
          ? existing.durationMs
          : null,
    requestSummary: payload.requestSummary || existing?.requestSummary || "",
    responseSummary: payload.responseSummary || existing?.responseSummary || "",
    error: payload.error || existing?.error || null,
    traceId: payload.traceId || existing?.traceId || "",
    matchRect: payload.matchRect || existing?.matchRect || "",
    targetPoint: payload.targetPoint || existing?.targetPoint || "",
    score:
      typeof payload.score === "number"
        ? payload.score
        : typeof existing?.score === "number"
          ? existing.score
          : null,
    exists:
      typeof payload.exists === "boolean"
        ? payload.exists
        : typeof existing?.exists === "boolean"
          ? existing.exists
          : null,
  };
}

function findPythonCommand(packageRoot) {
  const candidates = process.platform === "win32"
    ? [
        path.join(packageRoot, ".venv", "Scripts", "python.exe"),
        process.env.PYTHON_BIN,
        "py",
        "python",
      ]
    : [
        path.join(packageRoot, ".venv", "bin", "python"),
        process.env.PYTHON_BIN,
        "python3",
        "python",
      ];
  for (const candidate of candidates) {
    if (!candidate) {
      continue;
    }
    if (!candidate.includes(path.sep)) {
      return { command: candidate, argsPrefix: candidate === "py" ? ["-3"] : [] };
    }
    if (fs.existsSync(candidate)) {
      return { command: candidate, argsPrefix: [] };
    }
  }
  return { command: process.platform === "win32" ? "py" : "python3", argsPrefix: process.platform === "win32" ? ["-3"] : [] };
}

function safeUnlink(filePath) {
  if (!filePath) {
    return;
  }
  try {
    fs.unlinkSync(filePath);
  } catch (error) {
    if (error && error.code !== "ENOENT") {
      console.warn(`Unable to remove temporary file: ${filePath}`, error);
    }
  }
}

function createPythonControlChannel() {
  return new Promise((resolve, reject) => {
    const server = net.createServer();
    let socket = null;
    let pendingLines = [];
    let buffer = "";
    let closed = false;
    let onMessage = null;

    function closeSocket() {
      if (socket && !socket.destroyed) {
        socket.destroy();
      }
      socket = null;
    }

    function closeServer() {
      if (closed) {
        return;
      }
      closed = true;
      closeSocket();
      server.close();
    }

    function handleLine(line) {
      if (!line.trim()) {
        return;
      }
      try {
        const payload = JSON.parse(line);
        if (typeof onMessage === "function") {
          onMessage(payload);
        }
      } catch (error) {
        console.warn(`Unable to parse Python step payload: ${line}`, error);
      }
    }

    function attachSocket(nextSocket) {
      if (socket && socket !== nextSocket) {
        nextSocket.destroy();
        return;
      }

      socket = nextSocket;
      socket.setEncoding("utf8");
      socket.on("data", (chunk) => {
        buffer += chunk;
        while (buffer.includes("\n")) {
          const newlineIndex = buffer.indexOf("\n");
          const line = buffer.slice(0, newlineIndex);
          buffer = buffer.slice(newlineIndex + 1);
          handleLine(line);
        }
      });
      socket.on("error", () => {});
      socket.on("close", () => {
        if (socket === nextSocket) {
          socket = null;
        }
      });

      for (const line of pendingLines) {
        nextSocket.write(line);
      }
      pendingLines = [];
    }

    server.on("connection", attachSocket);
    server.once("error", (error) => {
      closeServer();
      reject(error);
    });

    server.listen(0, "127.0.0.1", () => {
      const address = server.address();
      if (!address || typeof address === "string") {
        closeServer();
        reject(new Error("Unable to allocate Python control channel"));
        return;
      }

      resolve({
        host: "127.0.0.1",
        port: address.port,
        setMessageHandler(handler) {
          onMessage = handler;
        },
        sendCommand(command) {
          const payload = command && typeof command === "object" ? command : { command };
          const line = `${JSON.stringify(payload)}\n`;
          if (socket && !socket.destroyed) {
            socket.write(line);
            return;
          }
          pendingLines.push(line);
        },
        close() {
          closeServer();
        },
      });
    });
  });
}

function createInstrumentedPythonExample(example, sessionId) {
  const extension = path.extname(example.absolutePath) || ".py";
  const baseName = path.basename(example.absolutePath, extension);
  const instrumentedPath = path.join(
    path.dirname(example.absolutePath),
    `.${baseName}.sikuli-ide-${sessionId}${extension}`,
  );

  const python = findPythonCommand(example.packageRoot);
  const result = spawnSync(python.command, [...python.argsPrefix, PYTHON_TRANSFORM_PATH, example.absolutePath, instrumentedPath], {
    cwd: example.packageRoot,
    encoding: "utf8",
  });

  if (result.status !== 0) {
    safeUnlink(instrumentedPath);
    throw new Error((result.stderr || result.stdout || "Python transform failed").trim());
  }

  return {
    instrumentedPath,
  };
}

async function buildNodeRunSpec(example, sessionId, options) {
  const instrumented = createInstrumentedNodeExample(example, sessionId);
  const nodeExecPath = process.env.SIKULI_GO_NODE_BIN || process.execPath;

  return {
    command: process.env.SIKULI_GO_NODE_BIN || "node",
    args: [example.relativePath],
    env: {
      ...process.env,
      SIKULI_GO_STEP_START_PAUSED: options.startPaused ? "1" : "0",
    },
    supportsStepping: true,
    startPaused: Boolean(options.startPaused),
    launch(cwd, env) {
      return {
        child: fork(NODE_RUNNER_PATH, [instrumented.instrumentedPath], {
          cwd,
          env,
          execPath: nodeExecPath,
          silent: true,
        }),
        cleanup() {
          safeUnlink(instrumented.instrumentedPath);
        },
        sendCommand(command) {
          if (typeof this.child?.send === "function" && command && typeof command === "object") {
            this.child.send(command);
            return true;
          }
          return false;
        },
      };
    },
  };
}

async function buildPythonRunSpec(example, sessionId, options) {
  const python = findPythonCommand(example.packageRoot);
  const instrumented = createInstrumentedPythonExample(example, sessionId);
  const controlChannel = await createPythonControlChannel();
  return {
    command: python.command,
    args: [...python.argsPrefix, example.relativePath],
    env: {
      ...process.env,
      PYTHONUNBUFFERED: "1",
      SIKULI_GO_STEP_HOST: controlChannel.host,
      SIKULI_GO_STEP_PORT: String(controlChannel.port),
      SIKULI_GO_STEP_START_PAUSED: options.startPaused ? "1" : "0",
    },
    supportsStepping: true,
    startPaused: Boolean(options.startPaused),
    setMessageHandler(handler) {
      controlChannel.setMessageHandler(handler);
    },
    launch(cwd, env) {
      return {
        child: spawn(python.command, [...python.argsPrefix, PYTHON_RUNNER_PATH, instrumented.instrumentedPath], {
          cwd,
          env,
          stdio: ["ignore", "pipe", "pipe"],
        }),
        cleanup() {
          controlChannel.close();
          safeUnlink(instrumented.instrumentedPath);
        },
        sendCommand(command) {
          controlChannel.sendCommand(command);
          return true;
        },
      };
    },
  };
}

async function resolveRunSpec(example, sessionId, options = {}) {
  if (example.language === "nodejs") {
    return await buildNodeRunSpec(example, sessionId, options);
  }
  return await buildPythonRunSpec(example, sessionId, options);
}

function initialSessionState(example, runSpec, sessionId) {
  return {
    sessionId,
    exampleId: example.id,
    exampleName: example.name,
    language: example.language,
    languageLabel: example.languageLabel,
    relativePath: example.relativePath,
    packageRoot: example.packageRoot,
    command: runSpec.command,
    args: [...runSpec.args],
    cwd: example.packageRoot,
    state: "starting",
    supportsStepping: Boolean(runSpec.supportsStepping),
    pid: null,
    startedAt: nowIso(),
    endedAt: null,
    currentLine: null,
    currentColumn: null,
    currentStatementId: null,
    pauseReason: null,
    runToLineTarget: null,
    awaitingRuntimeCall: false,
    currentRuntimeCall: null,
    lastRuntimeCall: null,
    runtimeCallCount: 0,
    runtimeErrorCount: 0,
    stdout: "",
    stderr: "",
    output: "",
    exitCode: null,
    signal: null,
    error: null,
    timeline: [],
  };
}

function sessionSnapshot(session) {
  return {
    sessionId: session.sessionId,
    exampleId: session.exampleId,
    exampleName: session.exampleName,
    language: session.language,
    languageLabel: session.languageLabel,
    relativePath: session.relativePath,
    packageRoot: session.packageRoot,
    command: session.command,
    args: [...session.args],
    cwd: session.cwd,
    state: session.state,
    supportsStepping: session.supportsStepping,
    pid: session.pid,
    startedAt: session.startedAt,
    endedAt: session.endedAt,
    currentLine: session.currentLine,
    currentColumn: session.currentColumn,
    currentStatementId: session.currentStatementId,
    pauseReason: session.pauseReason,
    runToLineTarget: session.runToLineTarget,
    awaitingRuntimeCall: session.awaitingRuntimeCall,
    currentRuntimeCall: cloneRuntimeCall(session.currentRuntimeCall),
    lastRuntimeCall: cloneRuntimeCall(session.lastRuntimeCall),
    runtimeCallCount: session.runtimeCallCount,
    runtimeErrorCount: session.runtimeErrorCount,
    exitCode: session.exitCode,
    signal: session.signal,
    error: session.error,
  };
}

function sessionRecord(session) {
  return {
    ...sessionSnapshot(session),
    stdout: session.stdout,
    stderr: session.stderr,
    output: session.output,
    currentRuntimeCall: cloneRuntimeCall(session.currentRuntimeCall),
    lastRuntimeCall: cloneRuntimeCall(session.lastRuntimeCall),
    timeline: session.timeline.map((entry) => ({ ...entry })),
  };
}

function outputTextForEvent(payload) {
  switch (payload.type) {
    case "stdout":
    case "stderr":
      return payload.chunk || "";
    case "session:starting":
      return `$ ${(payload.command || "").trim()} ${payload.args?.join(" ") || ""}\n[cwd] ${payload.cwd || ""}\n`;
    case "session:started":
      return `[started] pid=${payload.pid ?? "unknown"}\n`;
    case "session:stopping":
      return `[stopping] pid=${payload.pid ?? "unknown"}\n`;
    case "session:stopped":
      return `[stopped] signal=${payload.signal || "SIGTERM"}\n`;
    case "session:completed":
      return `[completed] exit=${payload.exitCode ?? 0}\n`;
    case "session:failed":
      if (payload.error) {
        return `[failed] ${payload.error}\n`;
      }
      return `[failed] exit=${payload.exitCode ?? "unknown"} signal=${payload.signal || "none"}\n`;
    case "runtime:call:start":
      return `[runtime] ${payload.method || "call"} start line=${payload.line ?? "?"}${payload.requestSummary ? ` ${payload.requestSummary}` : ""}\n`;
    case "runtime:call:end":
      return `[runtime] ${payload.method || "call"} ok${typeof payload.durationMs === "number" ? ` ${payload.durationMs}ms` : ""}${payload.responseSummary ? ` ${payload.responseSummary}` : ""}\n`;
    case "runtime:call:error":
      return `[runtime] ${payload.method || "call"} failed${typeof payload.durationMs === "number" ? ` ${payload.durationMs}ms` : ""}${payload.error ? ` ${payload.error}` : ""}\n`;
    default:
      return "";
  }
}

function previewChunk(chunk) {
  if (!chunk) {
    return "";
  }
  const singleLine = String(chunk).replace(/\s+/g, " ").trim();
  if (!singleLine) {
    return "(whitespace)";
  }
  return singleLine.length > 96 ? `${singleLine.slice(0, 93)}...` : singleLine;
}

function timelineSummary(payload) {
  switch (payload.type) {
    case "stdout":
      return `stdout: ${previewChunk(payload.chunk)}`;
    case "stderr":
      return `stderr: ${previewChunk(payload.chunk)}`;
    case "session:starting":
      return `Starting ${payload.command || ""} ${(payload.args || []).join(" ")}`.trim();
    case "session:started":
      return `Process started (pid ${payload.pid ?? "unknown"})`;
    case "session:stopping":
      return `Stopping process (pid ${payload.pid ?? "unknown"})`;
    case "session:stopped":
      return `Process stopped${payload.signal ? ` (${payload.signal})` : ""}`;
    case "session:completed":
      return `Process completed with exit ${payload.exitCode ?? 0}`;
    case "session:failed":
      if (payload.error) {
        return `Process failed: ${payload.error}`;
      }
      return `Process failed with exit ${payload.exitCode ?? "unknown"}`;
    case "step:start":
      return `Step start at line ${payload.line ?? "unknown"}`;
    case "step:end":
      return `Step complete at line ${payload.line ?? "unknown"}`;
    case "step:pause":
      if (payload.reason === "runtime-call") {
        return `Paused before ${payload.runtimeMethod || "runtime call"} at line ${payload.line ?? "unknown"}`;
      }
      if (payload.reason === "run-to-line") {
        return `Paused at target line ${payload.targetLine ?? payload.line ?? "unknown"}`;
      }
      return `Paused at line ${payload.line ?? "unknown"}`;
    case "step:resume":
      if (payload.mode === "runtime-call") {
        return `Continuing to next runtime call from line ${payload.line ?? "unknown"}`;
      }
      if (payload.mode === "run-to-line") {
        return `Running to line ${payload.targetLine ?? "unknown"} from line ${payload.line ?? "unknown"}`;
      }
      return `Resumed at line ${payload.line ?? "unknown"}${payload.mode ? ` (${payload.mode})` : ""}`;
    case "step:error":
      return `Step error at line ${payload.line ?? "unknown"}: ${payload.error || "unknown error"}`;
    case "runtime:call:start":
      return `${payload.method || "Runtime call"} started at line ${payload.line ?? "unknown"}${payload.requestSummary ? ` (${payload.requestSummary})` : ""}`;
    case "runtime:call:end":
      return `${payload.method || "Runtime call"} completed${typeof payload.durationMs === "number" ? ` in ${payload.durationMs}ms` : ""}${payload.responseSummary ? ` (${payload.responseSummary})` : ""}`;
    case "runtime:call:error":
      return `${payload.method || "Runtime call"} failed${typeof payload.durationMs === "number" ? ` in ${payload.durationMs}ms` : ""}: ${payload.error || "unknown error"}`;
    default:
      return payload.type;
  }
}

function timelineEntryForSession(session, payload) {
  return {
    id: `${session.sessionId}:${session.timeline.length + 1}`,
    type: payload.type,
    emittedAt: payload.emittedAt,
    summary: timelineSummary(payload),
  };
}

function updateSessionForEvent(session, payload) {
  if (payload.type === "step:start" || payload.type === "step:end" || payload.type === "step:pause" || payload.type === "step:resume" || payload.type === "step:error") {
    session.currentLine = payload.line ?? session.currentLine;
    session.currentColumn = payload.column ?? session.currentColumn;
    session.currentStatementId = payload.statementId ?? session.currentStatementId;
  }

  if (payload.type === "step:pause") {
    session.state = "paused";
    session.pauseReason = payload.reason || "step";
    if (payload.reason === "runtime-call") {
      session.awaitingRuntimeCall = false;
    } else if (payload.reason === "run-to-line") {
      session.runToLineTarget = payload.targetLine ?? session.runToLineTarget;
    }
  } else if (payload.type === "step:resume") {
    session.state = "running";
    session.pauseReason = null;
    if (payload.mode === "runtime-call") {
      session.awaitingRuntimeCall = true;
      session.runToLineTarget = null;
    } else if (payload.mode === "run-to-line") {
      session.awaitingRuntimeCall = false;
      session.runToLineTarget = payload.targetLine ?? session.runToLineTarget;
    } else {
      session.awaitingRuntimeCall = false;
      session.runToLineTarget = null;
    }
  } else if (payload.type === "step:error" && payload.error) {
    session.error = payload.error;
    session.pauseReason = "error";
  }

  if (payload.type === "runtime:call:start") {
    const runtimeCall = updateRuntimeCallRecord(null, payload, "running");
    session.currentRuntimeCall = runtimeCall;
    session.lastRuntimeCall = runtimeCall;
  } else if (payload.type === "runtime:call:end") {
    const existing =
      session.currentRuntimeCall && (!payload.callId || session.currentRuntimeCall.callId === payload.callId)
        ? session.currentRuntimeCall
        : session.lastRuntimeCall;
    const runtimeCall = updateRuntimeCallRecord(existing, payload, "completed");
    session.currentRuntimeCall = null;
    session.lastRuntimeCall = runtimeCall;
    session.runtimeCallCount += 1;
  } else if (payload.type === "runtime:call:error") {
    const existing =
      session.currentRuntimeCall && (!payload.callId || session.currentRuntimeCall.callId === payload.callId)
        ? session.currentRuntimeCall
        : session.lastRuntimeCall;
    const runtimeCall = updateRuntimeCallRecord(existing, payload, "failed");
    session.currentRuntimeCall = null;
    session.lastRuntimeCall = runtimeCall;
    session.runtimeCallCount += 1;
    session.runtimeErrorCount += 1;
    if (payload.error) {
      session.error = payload.error;
    }
  }
}

class ExecutionManager {
  constructor(options) {
    this.startRuntime = options.startRuntime;
    this.onEvent = options.onEvent;
    this.runtimeEnv = options.runtimeEnv || (() => ({}));
    this.getExampleById = options.getExampleById || defaultGetExampleById;
    this.maxSessions = options.maxSessions || MAX_SESSION_HISTORY;
    this.currentRun = null;
    this.sessions = [];
    this.sessionsById = new Map();
    this.sessionCounter = 0;
  }

  getCurrentRun() {
    if (!this.currentRun) {
      return null;
    }
    return sessionSnapshot(this.currentRun.session);
  }

  listSessions() {
    return this.sessions.map((session) => sessionRecord(session));
  }

  getSession(sessionId) {
    const session = this.sessionsById.get(sessionId);
    return session ? sessionRecord(session) : null;
  }

  registerSession(session) {
    this.sessions.unshift(session);
    this.sessionsById.set(session.sessionId, session);
    if (this.sessions.length <= this.maxSessions) {
      return;
    }
    const removed = this.sessions.pop();
    if (!removed) {
      return;
    }
    this.sessionsById.delete(removed.sessionId);
  }

  emitSessionEvent(session, payload) {
    const emittedAt = payload.emittedAt || nowIso();
    const eventPayload = {
      ...payload,
      emittedAt,
    };
    updateSessionForEvent(session, eventPayload);
    const outputText = outputTextForEvent(eventPayload);
    const timelineEntry = timelineEntryForSession(session, eventPayload);

    if (eventPayload.type === "stdout") {
      session.stdout += eventPayload.chunk || "";
    } else if (eventPayload.type === "stderr") {
      session.stderr += eventPayload.chunk || "";
    }

    if (outputText) {
      session.output += outputText;
    }
    session.timeline.push(timelineEntry);

    emitToHandler(this.onEvent, {
      ...eventPayload,
      outputText,
      timelineEntry,
      session: sessionSnapshot(session),
    });
  }

  cleanupRun(run) {
    if (!run || typeof run.cleanup !== "function") {
      return;
    }
    const cleanup = run.cleanup;
    run.cleanup = null;
    cleanup();
  }

  async runExample(exampleId, options = {}) {
    if (this.currentRun) {
      throw new Error("An example is already running");
    }

    const example = this.getExampleById(exampleId);
    if (!example) {
      throw new Error(`Unknown example: ${exampleId}`);
    }

    const runtime = await this.startRuntime();
    if (!runtime || runtime.ok !== true) {
      throw new Error("Runtime failed to start");
    }

    const sessionId = `run-${Date.now()}-${++this.sessionCounter}`;
    const runSpec = await resolveRunSpec(example, sessionId, options);
    const session = initialSessionState(example, runSpec, sessionId);
    this.registerSession(session);
    this.currentRun = {
      child: null,
      session,
      cleanup: null,
      sendCommand: null,
    };

    if (typeof runSpec.setMessageHandler === "function") {
      runSpec.setMessageHandler((message) => {
        if (
          !message ||
          typeof message !== "object" ||
          typeof message.type !== "string" ||
          (!message.type.startsWith("step:") && !message.type.startsWith("runtime:"))
        ) {
          return;
        }
        this.emitSessionEvent(session, {
          ...message,
          sessionId,
          exampleId,
        });
      });
    }

    this.emitSessionEvent(session, {
      type: "session:starting",
      sessionId,
      exampleId,
      language: example.language,
      command: runSpec.command,
      args: runSpec.args,
      cwd: example.packageRoot,
    });

    let child;
    try {
      const launched = runSpec.launch(example.packageRoot, {
        ...runSpec.env,
        ...this.runtimeEnv(),
      });
      child = launched.child;
      this.currentRun.cleanup = launched.cleanup;
      this.currentRun.sendCommand = typeof launched.sendCommand === "function" ? launched.sendCommand.bind(launched) : null;
      this.currentRun.child = child;
    } catch (error) {
      session.state = "failed";
      session.error = error instanceof Error ? error.message : String(error);
      session.endedAt = nowIso();
      this.cleanupRun(this.currentRun);
      this.currentRun = null;
      this.emitSessionEvent(session, {
        type: "session:failed",
        sessionId,
        exampleId,
        error: session.error,
      });
      throw error;
    }

    child.once("spawn", () => {
      if (!this.currentRun || this.currentRun.session.sessionId !== sessionId) {
        return;
      }
      session.state = "running";
      session.pid = child.pid || null;
      this.emitSessionEvent(session, {
        type: "session:started",
        sessionId,
        exampleId,
        pid: session.pid,
      });
    });

    child.on("message", (message) => {
      if (
        !message ||
        typeof message !== "object" ||
        typeof message.type !== "string" ||
        (!message.type.startsWith("step:") && !message.type.startsWith("runtime:"))
      ) {
        return;
      }
      this.emitSessionEvent(session, {
        ...message,
        sessionId,
        exampleId,
      });
    });

    child.stdout.on("data", (chunk) => {
      this.emitSessionEvent(session, {
        type: "stdout",
        sessionId,
        exampleId,
        chunk: String(chunk),
      });
    });

    child.stderr.on("data", (chunk) => {
      this.emitSessionEvent(session, {
        type: "stderr",
        sessionId,
        exampleId,
        chunk: String(chunk),
      });
    });

    child.on("error", (error) => {
      if (session.endedAt) {
        return;
      }
      session.state = "failed";
      session.error = error.message;
      session.endedAt = nowIso();
      if (this.currentRun && this.currentRun.session.sessionId === sessionId) {
        this.cleanupRun(this.currentRun);
        this.currentRun = null;
      }
      this.emitSessionEvent(session, {
        type: "session:failed",
        sessionId,
        exampleId,
        error: error.message,
      });
    });

    child.on("close", (code, signal) => {
      if (session.endedAt) {
        return;
      }
      const stoppedByUser = session.state === "stopping";
      session.exitCode = typeof code === "number" ? code : null;
      session.signal = signal || null;
      session.endedAt = nowIso();
      if (stoppedByUser) {
        session.state = "stopped";
      } else if (code === 0) {
        session.state = "completed";
      } else {
        session.state = "failed";
        if (!session.error) {
          session.error = signal
            ? `Runner exited unexpectedly with signal ${signal}`
            : `Runner exited unexpectedly with code ${session.exitCode ?? "unknown"}`;
        }
      }
      if (this.currentRun && this.currentRun.session.sessionId === sessionId) {
        this.cleanupRun(this.currentRun);
        this.currentRun = null;
      }
      this.emitSessionEvent(session, {
        type: stoppedByUser ? "session:stopped" : code === 0 ? "session:completed" : "session:failed",
        sessionId,
        exampleId,
        exitCode: session.exitCode,
        signal: session.signal,
        error: session.state === "failed" ? session.error : null,
      });
    });

    return {
      sessionId,
      exampleId,
    };
  }

  sendCurrentRunCommand(command) {
    if (!this.currentRun || !this.currentRun.session.supportsStepping || typeof this.currentRun.sendCommand !== "function") {
      return false;
    }
    return this.currentRun.sendCommand(command) === true;
  }

  async pauseCurrentRun() {
    return this.sendCurrentRunCommand({ command: "pause" });
  }

  async resumeCurrentRun() {
    return this.sendCurrentRunCommand({ command: "resume" });
  }

  async stepCurrentRun() {
    if (!this.currentRun || this.currentRun.session.state !== "paused") {
      return false;
    }
    return this.sendCurrentRunCommand({ command: "step" });
  }

  async runCurrentRunToLine(line) {
    if (!this.currentRun || this.currentRun.session.state !== "paused") {
      return false;
    }
    const targetLine = Number(line);
    if (!Number.isFinite(targetLine) || targetLine < 1) {
      return false;
    }
    return this.sendCurrentRunCommand({
      command: "run-to-line",
      line: Math.floor(targetLine),
    });
  }

  async continueCurrentRunToRuntimeCall() {
    if (!this.currentRun || this.currentRun.session.state !== "paused") {
      return false;
    }
    return this.sendCurrentRunCommand({ command: "continue-to-runtime-call" });
  }

  async stopCurrentRun() {
    if (!this.currentRun) {
      return false;
    }

    const activeRun = this.currentRun;
    activeRun.session.state = "stopping";
    this.emitSessionEvent(activeRun.session, {
      type: "session:stopping",
      sessionId: activeRun.session.sessionId,
      exampleId: activeRun.session.exampleId,
      pid: activeRun.child ? activeRun.child.pid || null : null,
    });

    if (!activeRun.child) {
      return true;
    }

    if (process.platform === "win32") {
      spawn("taskkill", ["/pid", String(activeRun.child.pid), "/f", "/t"], { stdio: "ignore" });
    } else {
      activeRun.child.kill("SIGTERM");
      setTimeout(() => {
        if (this.currentRun && this.currentRun.session.sessionId === activeRun.session.sessionId) {
          this.currentRun.child.kill("SIGKILL");
        }
      }, 3000).unref();
    }
    return true;
  }
}

module.exports = {
  ExecutionManager,
};
