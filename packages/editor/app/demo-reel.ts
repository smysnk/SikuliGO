type ExampleSummary = {
  id: string;
  language: string;
  languageLabel: string;
  name: string;
  fileName: string;
  relativePath: string;
  packageRoot: string;
  commandPreview: string;
};

type ExampleDocument = ExampleSummary & {
  absolutePath: string;
  source: string;
};

type WorkspaceFileSummary = {
  id: string;
  language: string;
  languageLabel: string;
  fileName: string;
  relativeWorkspacePath: string;
  relativePath: string;
  absolutePath: string;
  packageRoot: string;
  workspaceRoot: string;
  kind: "script" | "asset";
  extension: string;
  isImage: boolean;
  isText: boolean;
  size: number;
  updatedAt: string;
};

type WorkspaceFileDocument = WorkspaceFileSummary & {
  content: string;
  dataUrl: string;
};

type RuntimeStatus = {
  apiListen: string;
  adminListen: string;
  requestedApiListen: string;
  requestedAdminListen: string;
  portFallbackActive: boolean;
  runtimeWarning: string;
  healthy: boolean;
  dashboardUrl: string;
  sessionUrl: string;
  healthUrl: string;
  metricsUrl: string;
  snapshotUrl: string;
  binaryPath: string;
  managedPid: number | null;
};

type DashboardTab = "dashboard" | "session" | "health" | "metrics" | "snapshot";
type AppMode = "ide" | "dashboard";

type ExecutionRuntimeCall = {
  callId: string;
  method: string;
  state: string;
  line: number | null;
  column: number | null;
  statementId: string | null;
  startedAt: string;
  endedAt: string | null;
  durationMs: number | null;
  requestSummary: string;
  responseSummary: string;
  error: string | null;
  traceId: string;
  matchRect: string;
  targetPoint: string;
  score: number | null;
  exists: boolean | null;
};

type ExecutionTimelineEntry = {
  id: string;
  type: string;
  emittedAt: string;
  summary: string;
};

type ExecutionSessionSnapshot = {
  sessionId: string;
  exampleId: string;
  exampleName: string;
  language: string;
  languageLabel: string;
  relativePath: string;
  packageRoot: string;
  command: string;
  args: string[];
  cwd: string;
  state: string;
  supportsStepping: boolean;
  pid: number | null;
  startedAt: string;
  endedAt: string | null;
  currentLine: number | null;
  currentColumn: number | null;
  currentStatementId: string | null;
  pauseReason: string | null;
  runToLineTarget: number | null;
  awaitingRuntimeCall: boolean;
  currentRuntimeCall: ExecutionRuntimeCall | null;
  lastRuntimeCall: ExecutionRuntimeCall | null;
  runtimeCallCount: number;
  runtimeErrorCount: number;
  exitCode: number | null;
  signal: string | null;
  error: string | null;
};

type ExecutionSession = ExecutionSessionSnapshot & {
  stdout: string;
  stderr: string;
  output: string;
  timeline: ExecutionTimelineEntry[];
};

type ExecutionState = ExecutionSessionSnapshot | null;

type ExecutionEvent = {
  type: string;
  sessionId?: string;
  exampleId?: string;
  chunk?: string;
  pid?: number | null;
  exitCode?: number | null;
  signal?: string | null;
  error?: string;
  line?: number | null;
  column?: number | null;
  statementId?: string | null;
  mode?: string;
  reason?: string;
  targetLine?: number | null;
  runtimeMethod?: string | null;
  runtimeCallId?: string | null;
  command?: string;
  args?: string[];
  cwd?: string;
  emittedAt?: string;
  outputText?: string;
  timelineEntry?: ExecutionTimelineEntry;
  session?: ExecutionSessionSnapshot;
};

type IdeBridge = {
  listExamples(): Promise<ExampleSummary[]>;
  readExample(exampleId: string): Promise<ExampleDocument>;
  saveExample(exampleId: string, source: string): Promise<ExampleDocument>;
  createExample(language: string, fileName: string, source?: string): Promise<ExampleDocument>;
  cloneExample(exampleId: string, fileName: string, source?: string): Promise<ExampleDocument>;
  renameExample(exampleId: string, fileName: string, source?: string): Promise<ExampleDocument>;
  listWorkspaceFiles(exampleId: string): Promise<WorkspaceFileSummary[]>;
  readWorkspaceFile(exampleId: string, relativeWorkspacePath: string): Promise<WorkspaceFileDocument>;
  runExample(exampleId: string, options?: { startPaused?: boolean }): Promise<{ sessionId: string; exampleId: string }>;
  stopExecution(): Promise<boolean>;
  pauseExecution(): Promise<boolean>;
  resumeExecution(): Promise<boolean>;
  stepExecution(): Promise<boolean>;
  runToLine(line: number): Promise<boolean>;
  continueToRuntimeCall(): Promise<boolean>;
  getExecutionState(): Promise<ExecutionState>;
  listSessions(): Promise<ExecutionSession[]>;
  getSession(sessionId: string): Promise<ExecutionSession | null>;
  getRuntimeStatus(): Promise<RuntimeStatus>;
  startRuntime(): Promise<unknown>;
  restartRuntime(): Promise<unknown>;
  stopRuntime(): Promise<boolean>;
  onExecutionEvent(handler: (event: ExecutionEvent) => void): () => void;
};

export type WorkspaceTab = "authoring" | "sessions" | "runtime";

export type DemoScene = {
  durationMs: number;
  label: string;
  appMode: AppMode;
  selectedExampleId: string;
  selectedSessionId: string;
  selectedCodeLine: number;
  workspaceTab: WorkspaceTab;
  dashboardTab: DashboardTab;
  selectedWorkspaceFilePath: string;
  saveMessage: string;
};

const ROOT_DIR = "/Users/josh/play/sikuli-go";
const NODE_EXAMPLES_DIR = `${ROOT_DIR}/packages/client-node/examples`;
const PYTHON_EXAMPLES_DIR = `${ROOT_DIR}/packages/client-python/examples`;

const DEMO_NODE_CLICK_SOURCE = `import { ensureSikuliGoOnPath } from "./bootstrap.mjs";
import { Screen, Pattern } from "@sikuligo/sikuli-go";

ensureSikuliGoOnPath();

const screen = await Screen();
try {
  const pattern = Pattern("assets/pattern.png").similar(0.9);
  const match = await screen.click(pattern);
  console.log(\`clicked match target at (\${match.targetX}, \${match.targetY})\`);
} finally {
  await screen.close();
}
`;

const DEMO_NODE_CLICK_DRAFT = `import { ensureSikuliGoOnPath } from "./bootstrap.mjs";
import { Screen, Pattern } from "@sikuligo/sikuli-go";

ensureSikuliGoOnPath();

const screen = await Screen();
try {
  const pattern = Pattern("assets/pattern.png").similar(0.9);
  const match = await screen.click(pattern);
  console.log(\`clicked match target at (\${match.targetX}, \${match.targetY})\`);
  console.log("annotated from the IDE draft");
} finally {
  await screen.close();
}
`;

const DEMO_PYTHON_FIND_SOURCE = `from __future__ import annotations

from bootstrap import ensure_sikuli_go_on_path
from sikuligo import Pattern, Screen

ensure_sikuli_go_on_path()


def main() -> int:
    screen = Screen()
    try:
        match = screen.find(Pattern("assets/pattern.png").exact(), timeout_millis=3000)
        print(
            f"match rect=({match.x},{match.y},{match.w},{match.h}) "
            f"score={match.score:.3f} target=({match.target_x},{match.target_y})"
        )
    finally:
        screen.close()
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
`;

const DEMO_INVESTIGATION_SOURCE = `import { ensureSikuliGoOnPath } from "./bootstrap.mjs";
import { Screen, Pattern } from "@sikuligo/sikuli-go";

ensureSikuliGoOnPath();

const screen = await Screen();
try {
  const button = Pattern("assets/pattern.png").similar(0.92);
  const match = await screen.find(button, { timeoutMillis: 3000 });
  console.log("captured match", {
    score: match.score,
    rect: [match.x, match.y, match.w, match.h],
    target: [match.targetX, match.targetY],
  });
} finally {
  await screen.close();
}
`;

const DEMO_TEXT_ASSET = `notes:
- capture source and pattern pairs beside each script
- keep one Node.js and one Python session in history
- verify runtime call traces against the live dashboard pane
`;

const DEMO_EXAMPLES: ExampleDocument[] = [
  {
    id: "demo-node-click",
    language: "nodejs",
    languageLabel: "Node.js",
    name: "Click",
    fileName: "click.mjs",
    relativePath: "click.mjs",
    packageRoot: "packages/client-node",
    commandPreview: "node click.mjs",
    absolutePath: `${NODE_EXAMPLES_DIR}/click.mjs`,
    source: DEMO_NODE_CLICK_SOURCE,
  },
  {
    id: "demo-python-find",
    language: "python",
    languageLabel: "Python",
    name: "Find",
    fileName: "find.py",
    relativePath: "find.py",
    packageRoot: "packages/client-python",
    commandPreview: "python find.py",
    absolutePath: `${PYTHON_EXAMPLES_DIR}/find.py`,
    source: DEMO_PYTHON_FIND_SOURCE,
  },
  {
    id: "demo-node-investigation",
    language: "nodejs",
    languageLabel: "Node.js",
    name: "Investigation",
    fileName: "capture-investigation.mjs",
    relativePath: "capture-investigation.mjs",
    packageRoot: "packages/client-node",
    commandPreview: "node capture-investigation.mjs",
    absolutePath: `${NODE_EXAMPLES_DIR}/capture-investigation.mjs`,
    source: DEMO_INVESTIGATION_SOURCE,
  },
];

const DEMO_WORKSPACE_FILES_BY_EXAMPLE_ID: Record<string, WorkspaceFileDocument[]> = Object.fromEntries(
  DEMO_EXAMPLES.map((example) => {
    const workspaceRoot = example.absolutePath.slice(0, -example.fileName.length);
    const files: WorkspaceFileDocument[] = [
      {
        id: `${example.id}:asset:source`,
        language: example.language,
        languageLabel: example.languageLabel,
        fileName: "source.png",
        relativeWorkspacePath: "assets/source.png",
        relativePath: `${example.relativePath.replace(example.fileName, "")}assets/source.png`,
        absolutePath: `${workspaceRoot}assets/source.png`,
        packageRoot: example.packageRoot,
        workspaceRoot,
        kind: "asset",
        extension: ".png",
        isImage: true,
        isText: false,
        size: 184232,
        updatedAt: "2026-03-13T12:15:00.000Z",
        content: "",
        dataUrl: "/demo/assets/source.png",
      },
      {
        id: `${example.id}:asset:pattern`,
        language: example.language,
        languageLabel: example.languageLabel,
        fileName: "pattern.png",
        relativeWorkspacePath: "assets/pattern.png",
        relativePath: `${example.relativePath.replace(example.fileName, "")}assets/pattern.png`,
        absolutePath: `${workspaceRoot}assets/pattern.png`,
        packageRoot: example.packageRoot,
        workspaceRoot,
        kind: "asset",
        extension: ".png",
        isImage: true,
        isText: false,
        size: 12134,
        updatedAt: "2026-03-13T12:15:00.000Z",
        content: "",
        dataUrl: "/demo/assets/pattern.png",
      },
      {
        id: `${example.id}:asset:notes`,
        language: example.language,
        languageLabel: example.languageLabel,
        fileName: "notes.txt",
        relativeWorkspacePath: "notes.txt",
        relativePath: `${example.relativePath.replace(example.fileName, "")}notes.txt`,
        absolutePath: `${workspaceRoot}notes.txt`,
        packageRoot: example.packageRoot,
        workspaceRoot,
        kind: "asset",
        extension: ".txt",
        isImage: false,
        isText: true,
        size: DEMO_TEXT_ASSET.length,
        updatedAt: "2026-03-13T12:15:00.000Z",
        content: DEMO_TEXT_ASSET,
        dataUrl: "",
      },
    ];
    return [example.id, files];
  }),
);

const LIVE_RUNTIME_CALL: ExecutionRuntimeCall = {
  callId: "call-node-click-03",
  method: "Click",
  state: "completed",
  line: 8,
  column: 22,
  statementId: "stmt-click-08",
  startedAt: "2026-03-13T12:20:42.320Z",
  endedAt: "2026-03-13T12:20:42.364Z",
  durationMs: 44,
  requestSummary: 'Pattern("assets/pattern.png").similar(0.90)',
  responseSummary: "match rect=(1096,612,54,52) target=(1123,638)",
  error: null,
  traceId: "trace-7c61f2",
  matchRect: "1096,612,54,52",
  targetPoint: "1123,638",
  score: 0.994,
  exists: true,
};

const DEMO_SESSIONS: ExecutionSession[] = [
  {
    sessionId: "session-node-live",
    exampleId: "demo-node-click",
    exampleName: "Click",
    language: "nodejs",
    languageLabel: "Node.js",
    relativePath: "click.mjs",
    packageRoot: "packages/client-node",
    command: "node",
    args: ["click.mjs"],
    cwd: NODE_EXAMPLES_DIR,
    state: "paused",
    supportsStepping: true,
    pid: 48216,
    startedAt: "2026-03-13T12:20:41.900Z",
    endedAt: null,
    currentLine: 8,
    currentColumn: 22,
    currentStatementId: "stmt-click-08",
    pauseReason: "runtime-call",
    runToLineTarget: 10,
    awaitingRuntimeCall: false,
    currentRuntimeCall: LIVE_RUNTIME_CALL,
    lastRuntimeCall: LIVE_RUNTIME_CALL,
    runtimeCallCount: 3,
    runtimeErrorCount: 0,
    exitCode: null,
    signal: null,
    error: null,
    stdout: "clicked match target at (1123, 638)\nannotated from the IDE draft\n",
    stderr: "",
    output: [
      "[ide] session started",
      "[trace] Pattern prepared from assets/pattern.png",
      "[trace] Click request sent",
      "clicked match target at (1123, 638)",
      "annotated from the IDE draft",
      "[ide] paused before next runtime call",
    ].join("\n"),
    timeline: [
      {
        id: "session-node-live:01",
        type: "session:started",
        emittedAt: "2026-03-13T12:20:41.900Z",
        summary: "Node.js click.mjs launched in paused stepping mode.",
      },
      {
        id: "session-node-live:02",
        type: "step:pause",
        emittedAt: "2026-03-13T12:20:42.100Z",
        summary: "Paused at line 7 before screen.click(pattern).",
      },
      {
        id: "session-node-live:03",
        type: "runtime:call:end",
        emittedAt: "2026-03-13T12:20:42.364Z",
        summary: "Click completed with score 0.994 at target (1123, 638).",
      },
      {
        id: "session-node-live:04",
        type: "stdout",
        emittedAt: "2026-03-13T12:20:42.510Z",
        summary: "stdout: clicked match target at (1123, 638)",
      },
      {
        id: "session-node-live:05",
        type: "step:pause",
        emittedAt: "2026-03-13T12:20:42.610Z",
        summary: "Paused at line 10 with Continue To Runtime Call armed for the next step.",
      },
    ],
  },
  {
    sessionId: "session-python-find",
    exampleId: "demo-python-find",
    exampleName: "Find",
    language: "python",
    languageLabel: "Python",
    relativePath: "find.py",
    packageRoot: "packages/client-python",
    command: "python",
    args: ["find.py"],
    cwd: PYTHON_EXAMPLES_DIR,
    state: "completed",
    supportsStepping: true,
    pid: 48103,
    startedAt: "2026-03-13T12:18:09.110Z",
    endedAt: "2026-03-13T12:18:10.442Z",
    currentLine: 12,
    currentColumn: 15,
    currentStatementId: "stmt-find-12",
    pauseReason: null,
    runToLineTarget: null,
    awaitingRuntimeCall: false,
    currentRuntimeCall: null,
    lastRuntimeCall: {
      callId: "call-python-find-01",
      method: "Find",
      state: "completed",
      line: 12,
      column: 15,
      statementId: "stmt-find-12",
      startedAt: "2026-03-13T12:18:09.884Z",
      endedAt: "2026-03-13T12:18:10.141Z",
      durationMs: 257,
      requestSummary: 'Pattern("assets/pattern.png").exact() timeout=3000ms',
      responseSummary: "match rect=(1089,606,64,62) score=0.999",
      error: null,
      traceId: "trace-py-118e",
      matchRect: "1089,606,64,62",
      targetPoint: "1121,637",
      score: 0.999,
      exists: true,
    },
    runtimeCallCount: 1,
    runtimeErrorCount: 0,
    exitCode: 0,
    signal: null,
    error: null,
    stdout: "match rect=(1089,606,64,62) score=0.999 target=(1121,637)\n",
    stderr: "",
    output: [
      "[ide] Python session started",
      "match rect=(1089,606,64,62) score=0.999 target=(1121,637)",
      "[ide] session exited with code 0",
    ].join("\n"),
    timeline: [
      {
        id: "session-python-find:01",
        type: "session:started",
        emittedAt: "2026-03-13T12:18:09.110Z",
        summary: "Python find.py launched with stepping enabled.",
      },
      {
        id: "session-python-find:02",
        type: "runtime:call:end",
        emittedAt: "2026-03-13T12:18:10.141Z",
        summary: "Find returned a single match with score 0.999.",
      },
      {
        id: "session-python-find:03",
        type: "session:completed",
        emittedAt: "2026-03-13T12:18:10.442Z",
        summary: "Python run completed successfully.",
      },
    ],
  },
];

const DEMO_EXECUTION_SNAPSHOT: ExecutionSessionSnapshot = {
  sessionId: "session-node-live",
  exampleId: "demo-node-click",
  exampleName: "Click",
  language: "nodejs",
  languageLabel: "Node.js",
  relativePath: "click.mjs",
  packageRoot: "packages/client-node",
  command: "node",
  args: ["click.mjs"],
  cwd: NODE_EXAMPLES_DIR,
  state: "paused",
  supportsStepping: true,
  pid: 48216,
  startedAt: "2026-03-13T12:20:41.900Z",
  endedAt: null,
  currentLine: 8,
  currentColumn: 22,
  currentStatementId: "stmt-click-08",
  pauseReason: "runtime-call",
  runToLineTarget: 10,
  awaitingRuntimeCall: false,
  currentRuntimeCall: LIVE_RUNTIME_CALL,
  lastRuntimeCall: LIVE_RUNTIME_CALL,
  runtimeCallCount: 3,
  runtimeErrorCount: 0,
  exitCode: null,
  signal: null,
  error: null,
};

const DEMO_EXECUTION_STATE: ExecutionState = DEMO_EXECUTION_SNAPSHOT;

export const DEMO_RECENT_EXAMPLE_IDS = [
  "demo-node-click",
  "demo-python-find",
  "demo-node-investigation",
];

export const DEMO_DRAFTS_BY_EXAMPLE_ID: Record<string, string> = {
  "demo-node-click": DEMO_NODE_CLICK_DRAFT,
};

export const DEMO_SCENES: DemoScene[] = [
  {
    durationMs: 1400,
    label: "Authoring stays focused on the script, draft, and local assets instead of stacking every tool in one scroll view.",
    appMode: "ide",
    selectedExampleId: "demo-node-click",
    selectedSessionId: "session-node-live",
    selectedCodeLine: 8,
    workspaceTab: "authoring",
    dashboardTab: "dashboard",
    selectedWorkspaceFilePath: "assets/source.png",
    saveMessage: "",
  },
  {
    durationMs: 1400,
    label: "The Sessions tab keeps runtime diagnostics, buffered output, and timeline data attached to the selected run.",
    appMode: "ide",
    selectedExampleId: "demo-node-click",
    selectedSessionId: "session-node-live",
    selectedCodeLine: 10,
    workspaceTab: "sessions",
    dashboardTab: "session",
    selectedWorkspaceFilePath: "assets/pattern.png",
    saveMessage: "Saved click.mjs",
  },
  {
    durationMs: 1400,
    label: "Python examples use the same authoring workspace, with their own draft, assets, and selected line state.",
    appMode: "ide",
    selectedExampleId: "demo-python-find",
    selectedSessionId: "session-python-find",
    selectedCodeLine: 12,
    workspaceTab: "authoring",
    dashboardTab: "metrics",
    selectedWorkspaceFilePath: "notes.txt",
    saveMessage: "",
  },
  {
    durationMs: 1400,
    label: "New files and investigations stay inside the authoring tab so you can branch a workflow without leaving the editor.",
    appMode: "ide",
    selectedExampleId: "demo-node-investigation",
    selectedSessionId: "session-node-live",
    selectedCodeLine: 9,
    workspaceTab: "authoring",
    dashboardTab: "snapshot",
    selectedWorkspaceFilePath: "assets/source.png",
    saveMessage: "Cloned to capture-investigation.mjs",
  },
  {
    durationMs: 1400,
    label: "Dashboard mode promotes the runtime dashboard into the main content pane and moves the available views into the sidebar.",
    appMode: "dashboard",
    selectedExampleId: "demo-node-click",
    selectedSessionId: "session-node-live",
    selectedCodeLine: 8,
    workspaceTab: "runtime",
    dashboardTab: "dashboard",
    selectedWorkspaceFilePath: "assets/source.png",
    saveMessage: "",
  },
  {
    durationMs: 1800,
    label: "Once dashboard mode is active, the sidebar becomes a view navigator for the live dashboard, session, metrics, and snapshot panes.",
    appMode: "dashboard",
    selectedExampleId: "demo-node-click",
    selectedSessionId: "session-node-live",
    selectedCodeLine: 8,
    workspaceTab: "runtime",
    dashboardTab: "snapshot",
    selectedWorkspaceFilePath: "assets/source.png",
    saveMessage: "",
  },
];

function cloneDemoValue<T>(value: T): T {
  return JSON.parse(JSON.stringify(value)) as T;
}

function summarizeExample(example: ExampleDocument): ExampleSummary {
  return {
    id: example.id,
    language: example.language,
    languageLabel: example.languageLabel,
    name: example.name,
    fileName: example.fileName,
    relativePath: example.relativePath,
    packageRoot: example.packageRoot,
    commandPreview: example.commandPreview,
  };
}

function buildRuntimeStatus(origin: string): RuntimeStatus {
  return {
    apiListen: "127.0.0.1:7210",
    adminListen: "127.0.0.1:7310",
    requestedApiListen: "127.0.0.1:7210",
    requestedAdminListen: "127.0.0.1:7310",
    portFallbackActive: true,
    runtimeWarning: "Demo bridge is serving browser-safe runtime panes.",
    healthy: true,
    dashboardUrl: new URL("/demo/runtime/dashboard.html", origin).toString(),
    sessionUrl: new URL("/demo/runtime/sessions.html", origin).toString(),
    healthUrl: new URL("/demo/runtime/health.html", origin).toString(),
    metricsUrl: new URL("/demo/runtime/metrics.html", origin).toString(),
    snapshotUrl: new URL("/demo/runtime/snapshot.html", origin).toString(),
    binaryPath: `${ROOT_DIR}/bin/sikuli-go`,
    managedPid: 41782,
  };
}

export function createDemoBridge(origin: string): IdeBridge {
  const runtimeStatus = buildRuntimeStatus(origin);
  const examplesById = new Map(DEMO_EXAMPLES.map((example) => [example.id, cloneDemoValue(example)]));
  const workspaceFilesByExampleId = new Map(
    Object.entries(DEMO_WORKSPACE_FILES_BY_EXAMPLE_ID).map(([exampleId, files]) => [exampleId, cloneDemoValue(files)]),
  );
  const sessions = cloneDemoValue(DEMO_SESSIONS);
  let executionState = cloneDemoValue(DEMO_EXECUTION_STATE);
  const listeners = new Set<(event: ExecutionEvent) => void>();

  function orderedExamples(): ExampleDocument[] {
    return Array.from(examplesById.values()).sort((left, right) => left.fileName.localeCompare(right.fileName));
  }

  function exampleOrThrow(exampleId: string): ExampleDocument {
    const example = examplesById.get(exampleId);
    if (!example) {
      throw new Error(`Unknown demo example: ${exampleId}`);
    }
    return example;
  }

  function emit(event: ExecutionEvent) {
    for (const listener of listeners) {
      listener(cloneDemoValue(event));
    }
  }

  return {
    async listExamples() {
      return orderedExamples().map((example) => summarizeExample(cloneDemoValue(example)));
    },

    async readExample(exampleId: string) {
      return cloneDemoValue(exampleOrThrow(exampleId));
    },

    async saveExample(exampleId: string, source: string) {
      const example = exampleOrThrow(exampleId);
      example.source = source;
      return cloneDemoValue(example);
    },

    async createExample(language: string, fileName: string, source?: string) {
      const packageRoot = language === "python" ? "packages/client-python" : "packages/client-node";
      const examplesDir = language === "python" ? PYTHON_EXAMPLES_DIR : NODE_EXAMPLES_DIR;
      const normalizedId = `demo-created-${language}-${fileName.replace(/[^a-zA-Z0-9]+/g, "-").replace(/^-+|-+$/g, "").toLowerCase()}`;
      const nextExample: ExampleDocument = {
        id: normalizedId,
        language,
        languageLabel: language === "python" ? "Python" : "Node.js",
        name: fileName.replace(/\.[^.]+$/, "").replace(/[-_]/g, " ").replace(/\b\w/g, (value) => value.toUpperCase()),
        fileName,
        relativePath: fileName,
        packageRoot,
        commandPreview: `${language === "python" ? "python" : "node"} ${fileName}`,
        absolutePath: `${examplesDir}/${fileName}`,
        source:
          source ||
          (language === "python"
            ? "from sikuligo import Screen\n\nscreen = Screen()\n"
            : 'import { Screen } from "@sikuligo/sikuli-go";\n\nconst screen = await Screen();\n'),
      };
      examplesById.set(nextExample.id, nextExample);
      workspaceFilesByExampleId.set(nextExample.id, cloneDemoValue(DEMO_WORKSPACE_FILES_BY_EXAMPLE_ID["demo-node-click"]));
      return cloneDemoValue(nextExample);
    },

    async cloneExample(exampleId: string, fileName: string, source?: string) {
      const existing = exampleOrThrow(exampleId);
      const cloned = await this.createExample(existing.language, fileName, source || existing.source);
      return cloneDemoValue(cloned);
    },

    async renameExample(exampleId: string, fileName: string, source?: string) {
      const existing = exampleOrThrow(exampleId);
      examplesById.delete(exampleId);
      const renamedId = `demo-renamed-${existing.language}-${fileName.replace(/[^a-zA-Z0-9]+/g, "-").replace(/^-+|-+$/g, "").toLowerCase()}`;
      const renamed: ExampleDocument = {
        ...existing,
        id: renamedId,
        fileName,
        name: fileName.replace(/\.[^.]+$/, "").replace(/[-_]/g, " ").replace(/\b\w/g, (value) => value.toUpperCase()),
        relativePath: fileName,
        absolutePath: `${existing.absolutePath.slice(0, -existing.fileName.length)}${fileName}`,
        commandPreview: `${existing.language === "python" ? "python" : "node"} ${fileName}`,
        source: source || existing.source,
      };
      examplesById.set(renamed.id, renamed);
      const files = workspaceFilesByExampleId.get(exampleId) || [];
      workspaceFilesByExampleId.delete(exampleId);
      workspaceFilesByExampleId.set(renamed.id, files.map((file) => ({ ...file, id: file.id.replace(exampleId, renamed.id) })));
      return cloneDemoValue(renamed);
    },

    async listWorkspaceFiles(exampleId: string) {
      return cloneDemoValue((workspaceFilesByExampleId.get(exampleId) || []).map((file) => ({ ...file })));
    },

    async readWorkspaceFile(exampleId: string, relativeWorkspacePath: string) {
      const file = (workspaceFilesByExampleId.get(exampleId) || []).find(
        (candidate) => candidate.relativeWorkspacePath === relativeWorkspacePath,
      );
      if (!file) {
        throw new Error(`Unknown demo workspace file: ${relativeWorkspacePath}`);
      }
      return cloneDemoValue(file);
    },

    async runExample(exampleId: string) {
      const nextExecutionState = cloneDemoValue({
        ...DEMO_EXECUTION_SNAPSHOT,
        exampleId,
        exampleName: exampleOrThrow(exampleId).name,
      });
      executionState = nextExecutionState;
      emit({
        type: "session:started",
        sessionId: nextExecutionState.sessionId,
        exampleId,
        emittedAt: new Date().toISOString(),
        session: cloneDemoValue(nextExecutionState),
      });
      return { sessionId: nextExecutionState.sessionId, exampleId };
    },

    async stopExecution() {
      executionState = null;
      emit({
        type: "session:stopped",
        sessionId: "session-node-live",
        emittedAt: new Date().toISOString(),
      });
      return true;
    },

    async pauseExecution() {
      return true;
    },

    async resumeExecution() {
      return true;
    },

    async stepExecution() {
      return true;
    },

    async runToLine() {
      return true;
    },

    async continueToRuntimeCall() {
      return true;
    },

    async getExecutionState() {
      return cloneDemoValue(executionState);
    },

    async listSessions() {
      return cloneDemoValue(sessions);
    },

    async getSession(sessionId: string) {
      return cloneDemoValue(sessions.find((session) => session.sessionId === sessionId) || null);
    },

    async getRuntimeStatus() {
      return cloneDemoValue(runtimeStatus);
    },

    async startRuntime() {
      return { ok: true };
    },

    async restartRuntime() {
      return { ok: true };
    },

    async stopRuntime() {
      return true;
    },

    onExecutionEvent(handler: (event: ExecutionEvent) => void) {
      listeners.add(handler);
      return () => {
        listeners.delete(handler);
      };
    },
  };
}
