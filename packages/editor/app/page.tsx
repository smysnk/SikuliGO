"use client";

import { useEffect, useMemo, useRef, useState, type PointerEvent as ReactPointerEvent } from "react";
import styles from "./page.module.css";
import {
  DEMO_DRAFTS_BY_EXAMPLE_ID,
  DEMO_RECENT_EXAMPLE_IDS,
  DEMO_SCENES,
  type WorkspaceTab,
  createDemoBridge,
} from "./demo-reel";

const SIDEBAR_MIN = 240;
const SIDEBAR_MAX = 560;
const SIDEBAR_DEFAULT = 320;
const APP_MODE_KEY = "sikuli-go.editor.app.mode";
const EDITOR_THEME_KEY = "sikuli-go.editor.theme";
const SIDEBAR_WIDTH_KEY = "sikuli-go.editor.examples.sidebar.width";
const SELECTED_EXAMPLE_KEY = "sikuli-go.editor.examples.selected";
const DASHBOARD_TAB_KEY = "sikuli-go.editor.dashboard.tab";
const WORKSPACE_TAB_KEY = "sikuli-go.editor.workspace.tab";
const RECENT_EXAMPLES_KEY = "sikuli-go.editor.recent.examples";
const DRAFTS_KEY = "sikuli-go.editor.drafts";

type AppMode = "ide" | "dashboard";
type EditorTheme = "dark" | "dashboard";

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

declare global {
  interface Window {
    sikuliIde?: IdeBridge;
  }
}

function clampWidth(width: number): number {
  return Math.max(SIDEBAR_MIN, Math.min(SIDEBAR_MAX, width));
}

function packageLabel(packageRoot: string): string {
  return packageRoot.split("/").pop() || packageRoot;
}

function exampleWorkspacePath(example: ExampleSummary | ExampleDocument): string {
  return `${packageLabel(example.packageRoot)}/examples/${example.fileName}`;
}

function formatDateTime(value: string | null | undefined): string {
  if (!value) {
    return "not finished";
  }
  const parsed = new Date(value);
  if (Number.isNaN(parsed.getTime())) {
    return value;
  }
  return parsed.toLocaleString([], {
    month: "short",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
  });
}

function formatDuration(startedAt: string, endedAt: string | null): string {
  const started = new Date(startedAt).getTime();
  const ended = endedAt ? new Date(endedAt).getTime() : Date.now();
  if (!Number.isFinite(started) || !Number.isFinite(ended)) {
    return "unknown";
  }
  const totalSeconds = Math.max(0, Math.round((ended - started) / 1000));
  const minutes = Math.floor(totalSeconds / 60);
  const seconds = totalSeconds % 60;
  if (minutes === 0) {
    return `${seconds}s`;
  }
  return `${minutes}m ${seconds}s`;
}

function buildSessionFromSnapshot(snapshot: ExecutionSessionSnapshot): ExecutionSession {
  return {
    ...snapshot,
    args: [...snapshot.args],
    stdout: "",
    stderr: "",
    output: "",
    timeline: [],
  };
}

function fallbackSessionFromEvent(event: ExecutionEvent): ExecutionSession {
  const emittedAt = event.emittedAt || new Date().toISOString();
  return buildSessionFromSnapshot({
    sessionId: event.sessionId || `event-${emittedAt}`,
    exampleId: event.exampleId || "",
    exampleName: event.session?.exampleName || "Unknown Example",
    language: event.session?.language || "",
    languageLabel: event.session?.languageLabel || "Unknown",
    relativePath: event.session?.relativePath || "",
    packageRoot: event.session?.packageRoot || "",
    command: event.command || event.session?.command || "",
    args: event.args ? [...event.args] : event.session?.args ? [...event.session.args] : [],
    cwd: event.cwd || event.session?.cwd || "",
    state: event.session?.state || "running",
    supportsStepping: event.session?.supportsStepping || false,
    pid: event.pid ?? event.session?.pid ?? null,
    startedAt: event.session?.startedAt || emittedAt,
    endedAt: event.session?.endedAt || null,
    currentLine: event.line ?? event.session?.currentLine ?? null,
    currentColumn: event.column ?? event.session?.currentColumn ?? null,
    currentStatementId: event.statementId ?? event.session?.currentStatementId ?? null,
    pauseReason: event.session?.pauseReason || event.reason || null,
    runToLineTarget: event.session?.runToLineTarget ?? event.targetLine ?? null,
    awaitingRuntimeCall: event.session?.awaitingRuntimeCall || false,
    currentRuntimeCall: event.session?.currentRuntimeCall || null,
    lastRuntimeCall: event.session?.lastRuntimeCall || null,
    runtimeCallCount: event.session?.runtimeCallCount || 0,
    runtimeErrorCount: event.session?.runtimeErrorCount || 0,
    exitCode: event.exitCode ?? event.session?.exitCode ?? null,
    signal: event.signal ?? event.session?.signal ?? null,
    error: event.error || event.session?.error || null,
  });
}

function mergeSessionSnapshot(
  session: ExecutionSession,
  snapshot: ExecutionSessionSnapshot | undefined,
): ExecutionSession {
  if (!snapshot) {
    return session;
  }
  return {
    ...session,
    ...snapshot,
    args: [...snapshot.args],
  };
}

function applyExecutionEvent(
  previousSessions: ExecutionSession[],
  event: ExecutionEvent,
): ExecutionSession[] {
  if (!event.sessionId) {
    return previousSessions;
  }

  const nextSessions = [...previousSessions];
  const index = nextSessions.findIndex((session) => session.sessionId === event.sessionId);
  let session = index >= 0 ? nextSessions[index] : event.session ? buildSessionFromSnapshot(event.session) : fallbackSessionFromEvent(event);

  session = mergeSessionSnapshot(session, event.session);

  if (event.type === "stdout") {
    session = {
      ...session,
      stdout: `${session.stdout}${event.chunk || ""}`,
    };
  } else if (event.type === "stderr") {
    session = {
      ...session,
      stderr: `${session.stderr}${event.chunk || ""}`,
    };
  }

  if (event.outputText) {
    session = {
      ...session,
      output: `${session.output}${event.outputText}`,
    };
  }

  if (event.timelineEntry && !session.timeline.some((entry) => entry.id === event.timelineEntry?.id)) {
    session = {
      ...session,
      timeline: [...session.timeline, event.timelineEntry],
    };
  }

  if (index >= 0) {
    nextSessions[index] = session;
    return nextSessions;
  }

  nextSessions.unshift(session);
  return nextSessions;
}

function snapshotFromEvent(event: ExecutionEvent): ExecutionState {
  if (event.session) {
    return event.session;
  }
  if (!event.sessionId) {
    return null;
  }
  return {
    sessionId: event.sessionId,
    exampleId: event.exampleId || "",
    exampleName: "Unknown Example",
    language: "",
    languageLabel: "Unknown",
    relativePath: "",
    packageRoot: "",
    command: event.command || "",
    args: event.args ? [...event.args] : [],
    cwd: event.cwd || "",
    state: "running",
    supportsStepping: false,
    pid: event.pid ?? null,
    startedAt: event.emittedAt || new Date().toISOString(),
    endedAt: null,
    currentLine: event.line ?? null,
    currentColumn: event.column ?? null,
    currentStatementId: event.statementId ?? null,
    pauseReason: event.reason || null,
    runToLineTarget: event.targetLine ?? null,
    awaitingRuntimeCall: false,
    currentRuntimeCall: null,
    lastRuntimeCall: null,
    runtimeCallCount: 0,
    runtimeErrorCount: 0,
    exitCode: event.exitCode ?? null,
    signal: event.signal ?? null,
    error: event.error || null,
  };
}

function badgeTone(state: string): string {
  switch (state) {
    case "starting":
      return styles.badgeStarting;
    case "running":
      return styles.badgeRunning;
    case "paused":
      return styles.badgePaused;
    case "stopping":
      return styles.badgeStopping;
    case "completed":
      return styles.badgeCompleted;
    case "failed":
      return styles.badgeFailed;
    case "stopped":
      return styles.badgeStopped;
    default:
      return "";
  }
}

function isActiveExecutionState(state: ExecutionState): boolean {
  return Boolean(state && !["completed", "failed", "stopped"].includes(state.state));
}

function dashboardTabLabel(tab: DashboardTab): string {
  switch (tab) {
    case "dashboard":
      return "Live Dashboard";
    case "session":
      return "Session Viewer";
    case "health":
      return "Health";
    case "metrics":
      return "Metrics";
    case "snapshot":
      return "Snapshot";
    default:
      return tab;
  }
}

function workspaceTabLabel(tab: WorkspaceTab): string {
  switch (tab) {
    case "authoring":
      return "Authoring";
    case "sessions":
      return "Sessions";
    case "runtime":
      return "Runtime";
    default:
      return tab;
  }
}

function appModeLabel(mode: AppMode): string {
  switch (mode) {
    case "dashboard":
      return "Dashboard";
    case "ide":
    default:
      return "IDE";
  }
}

function editorThemeLabel(theme: EditorTheme): string {
  switch (theme) {
    case "dashboard":
      return "Dashboard Light";
    case "dark":
    default:
      return "Midnight IDE";
  }
}

function withEmbeddedDashboardTheme(url: string, theme: EditorTheme): string {
  try {
    const next = new URL(url);
    next.searchParams.set("theme", theme === "dashboard" ? "original" : "editor");
    next.searchParams.set("embed", "1");
    return next.toString();
  } catch {
    return url;
  }
}

function formatPauseReason(reason: string | null | undefined): string {
  switch (reason) {
    case "runtime-call":
      return "before runtime call";
    case "run-to-line":
      return "target line reached";
    case "step":
      return "step";
    case "error":
      return "runtime error";
    default:
      return reason || "none";
  }
}

function runtimeCallLabel(runtimeCall: ExecutionRuntimeCall | null): string {
  if (!runtimeCall) {
    return "none";
  }
  return `${runtimeCall.method} @ line ${runtimeCall.line ?? "?"}`;
}

function parseStoredStringArray(raw: string | null): string[] {
  if (!raw) {
    return [];
  }
  try {
    const parsed = JSON.parse(raw);
    if (!Array.isArray(parsed)) {
      return [];
    }
    return parsed.filter((value): value is string => typeof value === "string" && value.length > 0);
  } catch {
    return [];
  }
}

function parseStoredDraftMap(raw: string | null): Record<string, string> {
  if (!raw) {
    return {};
  }
  try {
    const parsed = JSON.parse(raw);
    if (!parsed || typeof parsed !== "object" || Array.isArray(parsed)) {
      return {};
    }
    return Object.fromEntries(
      Object.entries(parsed).filter(
        (entry): entry is [string, string] => typeof entry[0] === "string" && typeof entry[1] === "string",
      ),
    );
  } catch {
    return {};
  }
}

function nextRecentExampleIds(previous: string[], exampleId: string): string[] {
  return [exampleId, ...previous.filter((candidate) => candidate !== exampleId)].slice(0, 8);
}

export default function Home() {
  const demoBridgeRef = useRef<IdeBridge | null>(null);
  const menuBarRef = useRef<HTMLElement | null>(null);
  const [sidebarWidth, setSidebarWidth] = useState(SIDEBAR_DEFAULT);
  const [dragging, setDragging] = useState(false);
  const [demoMode, setDemoMode] = useState<boolean | null>(null);
  const [demoSceneIndex, setDemoSceneIndex] = useState(0);
  const [bridgeReady, setBridgeReady] = useState(false);
  const [loadingExamples, setLoadingExamples] = useState(true);
  const [examples, setExamples] = useState<ExampleSummary[]>([]);
  const [recentExampleIds, setRecentExampleIds] = useState<string[]>([]);
  const [selectedExampleId, setSelectedExampleId] = useState("");
  const [selectedExample, setSelectedExample] = useState<ExampleDocument | null>(null);
  const [selectedLoadError, setSelectedLoadError] = useState("");
  const [selectedCodeLine, setSelectedCodeLine] = useState<number | null>(null);
  const [draftsByExampleId, setDraftsByExampleId] = useState<Record<string, string>>({});
  const [draftSource, setDraftSource] = useState("");
  const [workspaceFiles, setWorkspaceFiles] = useState<WorkspaceFileSummary[]>([]);
  const [selectedWorkspaceFilePath, setSelectedWorkspaceFilePath] = useState("");
  const [selectedWorkspaceFile, setSelectedWorkspaceFile] = useState<WorkspaceFileDocument | null>(null);
  const [workspaceLoadError, setWorkspaceLoadError] = useState("");
  const [saveMessage, setSaveMessage] = useState("");
  const [runtimeStatus, setRuntimeStatus] = useState<RuntimeStatus | null>(null);
  const [executionState, setExecutionState] = useState<ExecutionState>(null);
  const [sessions, setSessions] = useState<ExecutionSession[]>([]);
  const [selectedSessionId, setSelectedSessionId] = useState("");
  const [appMode, setAppMode] = useState<AppMode>("ide");
  const [editorTheme, setEditorTheme] = useState<EditorTheme>("dark");
  const [workspaceTab, setWorkspaceTab] = useState<WorkspaceTab>("authoring");
  const [dashboardTab, setDashboardTab] = useState<DashboardTab>("dashboard");
  const [dashboardRefreshToken, setDashboardRefreshToken] = useState(0);
  const [collapsedTreeNodes, setCollapsedTreeNodes] = useState<Record<string, boolean>>({});
  const [openMenuId, setOpenMenuId] = useState<string | null>(null);
  const [actionError, setActionError] = useState("");
  const dragState = useRef<{ startX: number; startWidth: number } | null>(null);
  const outputRef = useRef<HTMLPreElement | null>(null);

  const bridge = typeof window !== "undefined" ? window.sikuliIde ?? (demoMode ? demoBridgeRef.current : undefined) : undefined;

  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const shouldUseDemo = params.get("demo") === "1";
    if (shouldUseDemo) {
      if (!demoBridgeRef.current) {
        demoBridgeRef.current = createDemoBridge(window.location.origin);
      }
      setDemoMode(true);
      setSidebarWidth(clampWidth(SIDEBAR_DEFAULT + 32));
      setSelectedExampleId(DEMO_SCENES[0]?.selectedExampleId || "");
      setSelectedSessionId(DEMO_SCENES[0]?.selectedSessionId || "");
      setSelectedCodeLine(DEMO_SCENES[0]?.selectedCodeLine || null);
      setAppMode(DEMO_SCENES[0]?.appMode || "ide");
      setEditorTheme("dark");
      setWorkspaceTab(DEMO_SCENES[0]?.workspaceTab || "authoring");
      setDashboardTab(DEMO_SCENES[0]?.dashboardTab || "dashboard");
      setRecentExampleIds([...DEMO_RECENT_EXAMPLE_IDS]);
      setDraftsByExampleId({ ...DEMO_DRAFTS_BY_EXAMPLE_ID });
      setSelectedWorkspaceFilePath(DEMO_SCENES[0]?.selectedWorkspaceFilePath || "");
      return;
    }

    setDemoMode(false);
    const parsed = Number(window.localStorage.getItem(SIDEBAR_WIDTH_KEY));
    if (Number.isFinite(parsed)) {
      setSidebarWidth(clampWidth(parsed));
    }
    const storedAppMode = window.localStorage.getItem(APP_MODE_KEY);
    if (storedAppMode === "ide" || storedAppMode === "dashboard") {
      setAppMode(storedAppMode);
    }
    const storedEditorTheme = window.localStorage.getItem(EDITOR_THEME_KEY);
    if (storedEditorTheme === "dark" || storedEditorTheme === "dashboard") {
      setEditorTheme(storedEditorTheme);
    }
    setSelectedExampleId(window.localStorage.getItem(SELECTED_EXAMPLE_KEY) || "");
    const workspaceTab = window.localStorage.getItem(WORKSPACE_TAB_KEY);
    if (workspaceTab === "authoring" || workspaceTab === "sessions" || workspaceTab === "runtime") {
      setWorkspaceTab(workspaceTab);
    }
    const dashboardTab = window.localStorage.getItem(DASHBOARD_TAB_KEY);
    if (dashboardTab === "dashboard" || dashboardTab === "session" || dashboardTab === "health" || dashboardTab === "metrics" || dashboardTab === "snapshot") {
      setDashboardTab(dashboardTab);
    }
    setRecentExampleIds(parseStoredStringArray(window.localStorage.getItem(RECENT_EXAMPLES_KEY)));
    setDraftsByExampleId(parseStoredDraftMap(window.localStorage.getItem(DRAFTS_KEY)));
  }, []);

  useEffect(() => {
    if (demoMode !== false) {
      return;
    }
    window.localStorage.setItem(SIDEBAR_WIDTH_KEY, String(sidebarWidth));
  }, [demoMode, sidebarWidth]);

  useEffect(() => {
    if (demoMode !== false) {
      return;
    }
    window.localStorage.setItem(APP_MODE_KEY, appMode);
  }, [appMode, demoMode]);

  useEffect(() => {
    if (demoMode !== false) {
      return;
    }
    window.localStorage.setItem(EDITOR_THEME_KEY, editorTheme);
  }, [demoMode, editorTheme]);

  useEffect(() => {
    if (demoMode !== false) {
      return;
    }
    if (!selectedExampleId) {
      return;
    }
    window.localStorage.setItem(SELECTED_EXAMPLE_KEY, selectedExampleId);
  }, [demoMode, selectedExampleId]);

  useEffect(() => {
    if (demoMode !== false) {
      return;
    }
    window.localStorage.setItem(WORKSPACE_TAB_KEY, workspaceTab);
  }, [demoMode, workspaceTab]);

  useEffect(() => {
    if (demoMode !== false) {
      return;
    }
    window.localStorage.setItem(DASHBOARD_TAB_KEY, dashboardTab);
  }, [dashboardTab, demoMode]);

  useEffect(() => {
    if (demoMode !== false) {
      return;
    }
    window.localStorage.setItem(RECENT_EXAMPLES_KEY, JSON.stringify(recentExampleIds));
  }, [demoMode, recentExampleIds]);

  useEffect(() => {
    if (demoMode !== false) {
      return;
    }
    window.localStorage.setItem(DRAFTS_KEY, JSON.stringify(draftsByExampleId));
  }, [demoMode, draftsByExampleId]);

  useEffect(() => {
    if (typeof window === "undefined") {
      return;
    }

    const handlePointerDown = (event: PointerEvent) => {
      if (!menuBarRef.current || menuBarRef.current.contains(event.target as Node)) {
        return;
      }
      setOpenMenuId(null);
    };

    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.key === "Escape") {
        setOpenMenuId(null);
      }
    };

    window.addEventListener("pointerdown", handlePointerDown);
    window.addEventListener("keydown", handleKeyDown);
    return () => {
      window.removeEventListener("pointerdown", handlePointerDown);
      window.removeEventListener("keydown", handleKeyDown);
    };
  }, []);

  useEffect(() => {
    const onMove = (event: globalThis.PointerEvent) => {
      if (!dragState.current) {
        return;
      }
      const delta = event.clientX - dragState.current.startX;
      setSidebarWidth(clampWidth(dragState.current.startWidth + delta));
    };

    const onUp = () => {
      dragState.current = null;
      setDragging(false);
    };

    window.addEventListener("pointermove", onMove);
    window.addEventListener("pointerup", onUp);
    return () => {
      window.removeEventListener("pointermove", onMove);
      window.removeEventListener("pointerup", onUp);
    };
  }, []);

  const selectedSession = useMemo(
    () => sessions.find((session) => session.sessionId === selectedSessionId) || null,
    [selectedSessionId, sessions],
  );
  const activeSession = useMemo(
    () => (executionState ? sessions.find((session) => session.sessionId === executionState.sessionId) || null : null),
    [executionState, sessions],
  );

  useEffect(() => {
    if (!outputRef.current) {
      return;
    }
    outputRef.current.scrollTop = outputRef.current.scrollHeight;
  }, [selectedSession?.output]);

  useEffect(() => {
    if (demoMode === null) {
      return;
    }
    if (!bridge) {
      setLoadingExamples(false);
      setBridgeReady(false);
      return;
    }
    const ide = bridge;

    let cancelled = false;
    setBridgeReady(true);

    const unsubscribe = ide.onExecutionEvent((event) => {
      if (cancelled) {
        return;
      }

      setSessions((previous) => applyExecutionEvent(previous, event));

      if (event.type === "session:starting" && event.sessionId) {
        setSelectedSessionId(event.sessionId);
        setWorkspaceTab("sessions");
        setDashboardTab("session");
        if (event.exampleId) {
          setSelectedExampleId(event.exampleId);
        }
      }

      if (
        event.type === "session:starting" ||
        event.type === "session:started" ||
        event.type === "session:stopping" ||
        event.type === "step:start" ||
        event.type === "step:pause" ||
        event.type === "step:resume" ||
        event.type === "runtime:call:start" ||
        event.type === "runtime:call:end" ||
        event.type === "runtime:call:error"
      ) {
        setExecutionState(snapshotFromEvent(event));
        return;
      }

      if (event.type === "session:completed" || event.type === "session:failed" || event.type === "session:stopped") {
        setExecutionState((current) => (current && current.sessionId === event.sessionId ? null : current));
        return;
      }

      if (
        (event.type === "stdout" ||
          event.type === "stderr" ||
          event.type === "step:end" ||
          event.type === "step:error") &&
        event.session
      ) {
        setExecutionState((current) => (current && current.sessionId === event.sessionId ? event.session || current : current));
      }
    });

    async function loadInitialData() {
      try {
        const [listedExamples, currentRuntime, currentExecution, sessionHistory] = await Promise.all([
          ide.listExamples(),
          ide.getRuntimeStatus(),
          ide.getExecutionState(),
          ide.listSessions(),
        ]);

        if (cancelled) {
          return;
        }

        setExamples(listedExamples);
        setRuntimeStatus(currentRuntime);
        setExecutionState(currentExecution);
        setSessions(sessionHistory);

        if (listedExamples.length > 0) {
          const storedRecentExamples =
            typeof window !== "undefined"
              ? parseStoredStringArray(window.localStorage.getItem(RECENT_EXAMPLES_KEY))
              : [];
          setSelectedExampleId((previous) => {
            if (previous && listedExamples.some((example) => example.id === previous)) {
              return previous;
            }
            const recentExampleId = storedRecentExamples.find((exampleId) => listedExamples.some((example) => example.id === exampleId));
            return currentExecution?.exampleId || sessionHistory[0]?.exampleId || recentExampleId || listedExamples[0].id;
          });
        }

        if (sessionHistory.length > 0) {
          const initialSessionId = currentExecution?.sessionId || sessionHistory[0].sessionId;
          setSelectedSessionId((previous) => previous || initialSessionId);
        }
      } catch (error) {
        if (cancelled) {
          return;
        }
        setActionError(error instanceof Error ? error.message : String(error));
      } finally {
        if (!cancelled) {
          setLoadingExamples(false);
        }
      }
    }

    loadInitialData();

    return () => {
      cancelled = true;
      unsubscribe();
    };
  }, [bridge]);

  useEffect(() => {
    if (!sessions.length) {
      if (selectedSessionId) {
        setSelectedSessionId("");
      }
      return;
    }

    if (selectedSessionId && sessions.some((session) => session.sessionId === selectedSessionId)) {
      return;
    }

    const fallbackSessionId =
      (executionState && sessions.some((session) => session.sessionId === executionState.sessionId)
        ? executionState.sessionId
        : sessions[0].sessionId);
    setSelectedSessionId(fallbackSessionId);
  }, [executionState, selectedSessionId, sessions]);

  useEffect(() => {
    if (!selectedExampleId) {
      return;
    }
    setRecentExampleIds((previous) => nextRecentExampleIds(previous, selectedExampleId));
  }, [selectedExampleId]);

  useEffect(() => {
    if (!bridge || !selectedExampleId) {
      setSelectedExample(null);
      setSelectedCodeLine(null);
      setDraftSource("");
      return;
    }
    const ide = bridge;

    let cancelled = false;
    setSelectedLoadError("");

    async function loadExample() {
      try {
        const loaded = await ide.readExample(selectedExampleId);
        if (!cancelled) {
          const storedDrafts =
            demoMode === true
              ? DEMO_DRAFTS_BY_EXAMPLE_ID
              : typeof window !== "undefined"
              ? parseStoredDraftMap(window.localStorage.getItem(DRAFTS_KEY))
              : {};
          setSelectedExample(loaded);
          setSelectedCodeLine((current) => (current && current > 0 ? current : 1));
          setDraftSource(draftsByExampleId[selectedExampleId] ?? storedDrafts[selectedExampleId] ?? loaded.source);
          setSaveMessage("");
        }
      } catch (error) {
        if (!cancelled) {
          setSelectedExample(null);
          setSelectedCodeLine(null);
          setDraftSource("");
          setSelectedLoadError(error instanceof Error ? error.message : String(error));
        }
      }
    }

    loadExample();
    return () => {
      cancelled = true;
    };
  }, [bridge, demoMode, draftsByExampleId, selectedExampleId]);

  useEffect(() => {
    if (demoMode !== true || examples.length === 0) {
      return;
    }

    const timers: number[] = [];
    const applyScene = (sceneIndex: number) => {
      const scene = DEMO_SCENES[sceneIndex];
      if (!scene) {
        return;
      }
      setDemoSceneIndex(sceneIndex);
      setSelectedExampleId(scene.selectedExampleId);
      setSelectedSessionId(scene.selectedSessionId);
      setSelectedCodeLine(scene.selectedCodeLine);
      setAppMode(scene.appMode);
      setWorkspaceTab(scene.workspaceTab);
      setDashboardTab(scene.dashboardTab);
      setSelectedWorkspaceFilePath(scene.selectedWorkspaceFilePath);
      timers.push(window.setTimeout(() => setSaveMessage(scene.saveMessage), 120));
    };

    applyScene(0);
    let elapsed = 0;
    for (let index = 1; index < DEMO_SCENES.length; index += 1) {
      elapsed += DEMO_SCENES[index - 1].durationMs;
      const sceneIndex = index;
      timers.push(window.setTimeout(() => applyScene(sceneIndex), elapsed));
    }

    return () => {
      for (const timer of timers) {
        window.clearTimeout(timer);
      }
    };
  }, [demoMode, examples.length]);

  useEffect(() => {
    if (!selectedExampleId || !selectedExample) {
      return;
    }
    setDraftsByExampleId((previous) => {
      const next = { ...previous };
      if (draftSource === selectedExample.source) {
        delete next[selectedExampleId];
      } else {
        next[selectedExampleId] = draftSource;
      }
      return next;
    });
  }, [draftSource, selectedExample, selectedExampleId]);

  useEffect(() => {
    if (!bridge || !selectedExample) {
      setWorkspaceFiles([]);
      setSelectedWorkspaceFilePath("");
      setSelectedWorkspaceFile(null);
      setWorkspaceLoadError("");
      return;
    }

    const ide = bridge;
    const exampleId = selectedExample.id;
    let cancelled = false;
    setWorkspaceLoadError("");

    async function loadWorkspaceFiles() {
      try {
        const files = await ide.listWorkspaceFiles(exampleId);
        if (cancelled) {
          return;
        }
        setWorkspaceFiles(files);
        const assetFiles = files.filter((file) => file.kind === "asset");
        setSelectedWorkspaceFilePath((current) => {
          if (current && assetFiles.some((file) => file.relativeWorkspacePath === current)) {
            return current;
          }
          return assetFiles[0]?.relativeWorkspacePath || "";
        });
      } catch (error) {
        if (!cancelled) {
          setWorkspaceFiles([]);
          setSelectedWorkspaceFilePath("");
          setSelectedWorkspaceFile(null);
          setWorkspaceLoadError(error instanceof Error ? error.message : String(error));
        }
      }
    }

    loadWorkspaceFiles();
    return () => {
      cancelled = true;
    };
  }, [bridge, selectedExample]);

  useEffect(() => {
    if (!bridge || !selectedExample || !selectedWorkspaceFilePath) {
      setSelectedWorkspaceFile(null);
      return;
    }

    const ide = bridge;
    const exampleId = selectedExample.id;
    let cancelled = false;

    async function loadWorkspaceFile() {
      try {
        const file = await ide.readWorkspaceFile(exampleId, selectedWorkspaceFilePath);
        if (!cancelled) {
          setSelectedWorkspaceFile(file);
        }
      } catch (error) {
        if (!cancelled) {
          setSelectedWorkspaceFile(null);
          setWorkspaceLoadError(error instanceof Error ? error.message : String(error));
        }
      }
    }

    loadWorkspaceFile();
    return () => {
      cancelled = true;
    };
  }, [bridge, selectedExample, selectedWorkspaceFilePath]);

  const recentExamples = useMemo(
    () =>
      recentExampleIds
        .map((exampleId) => examples.find((example) => example.id === exampleId) || null)
        .filter((example): example is ExampleSummary => Boolean(example)),
    [examples, recentExampleIds],
  );
  const examplesByPackage = useMemo(() => {
    const groups = new Map<string, ExampleSummary[]>();
    for (const example of examples) {
      const current = groups.get(example.packageRoot) || [];
      current.push(example);
      groups.set(example.packageRoot, current);
    }
    return Array.from(groups.entries())
      .map(([packageRoot, items]) => ({
        packageRoot,
        label: packageLabel(packageRoot),
        items: items.sort((left, right) => left.fileName.localeCompare(right.fileName)),
      }))
      .sort((left, right) => left.label.localeCompare(right.label));
  }, [examples]);
  const editorLineNumbers = useMemo(() => {
    const source = draftSource || selectedExample?.source || "";
    const lineCount = Math.max(1, source.split("\n").length);
    return Array.from({ length: lineCount }, (_, index) => index + 1);
  }, [draftSource, selectedExample?.source]);
  const viewItems: { id: WorkspaceTab; title: string; subtitle: string }[] = [
    {
      id: "authoring",
      title: "Editor",
      subtitle: selectedExample ? selectedExample.fileName : "No file selected",
    },
    {
      id: "sessions",
      title: "Sessions",
      subtitle: selectedSession ? selectedSession.exampleName : `${sessions.length} stored`,
    },
    {
      id: "runtime",
      title: "Runtime",
      subtitle: runtimeStatus?.healthy ? runtimeStatus.adminListen : "not connected",
    },
  ];
  const dashboardViewItems: { id: DashboardTab; title: string; subtitle: string }[] = [
    {
      id: "dashboard",
      title: "Main Dashboard",
      subtitle: "Requests, sessions, and runtime health",
    },
    {
      id: "session",
      title: "Session Viewer",
      subtitle: activeSession ? `${activeSession.exampleName} (${activeSession.sessionId})` : "Latest live session surface",
    },
    {
      id: "health",
      title: "Health",
      subtitle: "Liveness and readiness endpoints",
    },
    {
      id: "metrics",
      title: "Metrics",
      subtitle: "Prometheus metrics stream",
    },
    {
      id: "snapshot",
      title: "Snapshot",
      subtitle: "Rendered capture preview",
    },
  ];
  const isTreeNodeExpanded = (nodeId: string) => !collapsedTreeNodes[nodeId];
  const toggleTreeNode = (nodeId: string) => {
    setCollapsedTreeNodes((current) => ({
      ...current,
      [nodeId]: !current[nodeId],
    }));
  };
  const focusExample = (exampleId: string) => {
    setWorkspaceTab("authoring");
    setSaveMessage("");
    setSelectedCodeLine(null);
    setSelectedExampleId(exampleId);
  };
  const handleSelectAppMode = (mode: AppMode) => {
    setOpenMenuId(null);
    setAppMode(mode);
    if (mode === "dashboard") {
      setDashboardTab("dashboard");
    }
  };

  const onResizePointerDown = (event: ReactPointerEvent<HTMLDivElement>) => {
    dragState.current = { startX: event.clientX, startWidth: sidebarWidth };
    setDragging(true);
    event.currentTarget.setPointerCapture(event.pointerId);
  };

  const reloadExamples = async (preferredExampleId?: string) => {
    if (!bridge) {
      return;
    }
    const listedExamples = await bridge.listExamples();
    setExamples(listedExamples);
    if (listedExamples.length === 0) {
      setSelectedExampleId("");
      return;
    }
    setSelectedExampleId((current) => {
      if (preferredExampleId && listedExamples.some((example) => example.id === preferredExampleId)) {
        return preferredExampleId;
      }
      if (current && listedExamples.some((example) => example.id === current)) {
        return current;
      }
      return listedExamples[0].id;
    });
  };

  const refreshRuntimeStatus = async (options?: { reloadPanes?: boolean }) => {
    if (!bridge) {
      return;
    }
    try {
      setActionError("");
      setRuntimeStatus(await bridge.getRuntimeStatus());
      if (options?.reloadPanes !== false) {
        setDashboardRefreshToken((current) => current + 1);
      }
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const saveDirtyDraftIfNeeded = async () => {
    if (!bridge || !selectedExample || !draftDirty) {
      return selectedExample;
    }
    const saved = await bridge.saveExample(selectedExample.id, draftSource);
    setSelectedExample(saved);
    setDraftSource(saved.source);
    setSaveMessage(`Saved ${saved.fileName}`);
    await reloadExamples(saved.id);
    return saved;
  };

  const handleRun = async () => {
    if (!bridge || !selectedExampleId) {
      return;
    }
    try {
      setActionError("");
      await saveDirtyDraftIfNeeded();
      await bridge.runExample(selectedExampleId, { startPaused: false });
      await refreshRuntimeStatus({ reloadPanes: false });
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const handleRunPaused = async () => {
    if (!bridge || !selectedExampleId) {
      return;
    }
    try {
      setActionError("");
      await saveDirtyDraftIfNeeded();
      await bridge.runExample(selectedExampleId, { startPaused: true });
      await refreshRuntimeStatus({ reloadPanes: false });
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const handlePause = async () => {
    if (!bridge) {
      return;
    }
    try {
      setActionError("");
      await bridge.pauseExecution();
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const handleResume = async () => {
    if (!bridge) {
      return;
    }
    try {
      setActionError("");
      await bridge.resumeExecution();
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const handleStep = async () => {
    if (!bridge) {
      return;
    }
    try {
      setActionError("");
      await bridge.stepExecution();
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const handleContinueToRuntimeCall = async () => {
    if (!bridge) {
      return;
    }
    try {
      setActionError("");
      await bridge.continueToRuntimeCall();
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const handleRunToLine = async () => {
    if (!bridge || !selectedCodeLine) {
      return;
    }
    try {
      setActionError("");
      await bridge.runToLine(selectedCodeLine);
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const handleStop = async () => {
    if (!bridge) {
      return;
    }
    try {
      setActionError("");
      await bridge.stopExecution();
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const handleStartRuntime = async () => {
    if (!bridge) {
      return;
    }
    try {
      setActionError("");
      await bridge.startRuntime();
      await refreshRuntimeStatus();
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const handleRestartRuntime = async () => {
    if (!bridge) {
      return;
    }
    try {
      setActionError("");
      await bridge.restartRuntime();
      await refreshRuntimeStatus();
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const handleStopRuntime = async () => {
    if (!bridge) {
      return;
    }
    try {
      setActionError("");
      await bridge.stopRuntime();
      await refreshRuntimeStatus();
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const handleSaveExample = async () => {
    if (!bridge || !selectedExample) {
      return;
    }
    try {
      setActionError("");
      const saved = await bridge.saveExample(selectedExample.id, draftSource);
      setSelectedExample(saved);
      setDraftSource(saved.source);
      setSaveMessage(`Saved ${saved.fileName}`);
      await reloadExamples(saved.id);
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const handleRevertExample = () => {
    if (!selectedExample) {
      return;
    }
    setDraftSource(selectedExample.source);
    setSaveMessage(`Reverted ${selectedExample.fileName}`);
  };

  const handleCreateExample = async () => {
    if (!bridge) {
      return;
    }
    const language =
      selectedExample?.language ||
      examples.find((example) => example.id === selectedExampleId)?.language ||
      examples[0]?.language ||
      "nodejs";
    const suggestedName = language === "python" ? "new-example.py" : "new-example.mjs";
    const nextFileName = window.prompt(`Create a new ${language === "python" ? "Python" : "Node.js"} example`, suggestedName);
    if (!nextFileName) {
      return;
    }
    try {
      setActionError("");
      const created = await bridge.createExample(language, nextFileName);
      setDraftSource(created.source);
      setSaveMessage(`Created ${created.fileName}`);
      await reloadExamples(created.id);
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const handleCloneExample = async () => {
    if (!bridge || !selectedExample) {
      return;
    }
    const suggestedName = selectedExample.fileName.replace(/\.[^.]+$/, (extension) => `-copy${extension}`);
    const nextFileName = window.prompt("Clone the current example to", suggestedName);
    if (!nextFileName) {
      return;
    }
    try {
      setActionError("");
      const cloned = await bridge.cloneExample(selectedExample.id, nextFileName, draftSource);
      setDraftSource(cloned.source);
      setSaveMessage(`Cloned to ${cloned.fileName}`);
      await reloadExamples(cloned.id);
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const handleRenameExample = async () => {
    if (!bridge || !selectedExample) {
      return;
    }
    const nextFileName = window.prompt("Rename the current example", selectedExample.fileName);
    if (!nextFileName) {
      return;
    }
    try {
      setActionError("");
      const renamed = await bridge.renameExample(selectedExample.id, nextFileName, draftSource);
      setDraftSource(renamed.source);
      setSaveMessage(`Renamed to ${renamed.fileName}`);
      await reloadExamples(renamed.id);
    } catch (error) {
      setActionError(error instanceof Error ? error.message : String(error));
    }
  };

  const handleSelectSession = (session: ExecutionSession) => {
    setWorkspaceTab("sessions");
    setSelectedSessionId(session.sessionId);
    setSelectedExampleId(session.exampleId);
    setSelectedCodeLine(session.currentLine ?? 1);
  };

  const sidebarStyle = { width: `${sidebarWidth}px` };
  const activeRunState = executionState?.state || "idle";
  const hasActiveRun = isActiveExecutionState(executionState);
  const canStartPaused =
    selectedExample?.language === "nodejs" ||
    selectedExample?.language === "python" ||
    examples.find((example) => example.id === selectedExampleId)?.language === "nodejs" ||
    examples.find((example) => example.id === selectedExampleId)?.language === "python";
  const canPause = Boolean(activeSession?.supportsStepping && activeRunState === "running");
  const canResume = Boolean(activeSession?.supportsStepping && activeRunState === "paused");
  const canStep = Boolean(activeSession?.supportsStepping && activeRunState === "paused");
  const canContinueToRuntimeCall = Boolean(activeSession?.supportsStepping && activeRunState === "paused");
  const canRunToLine = Boolean(activeSession?.supportsStepping && activeRunState === "paused" && selectedCodeLine);
  const selectedSessionState = selectedSession?.state || activeRunState;
  const highlightedLine = selectedSession && selectedSession.exampleId === selectedExample?.id ? selectedSession.currentLine : null;
  const selectedRuntimeCall = selectedSession?.currentRuntimeCall || selectedSession?.lastRuntimeCall || null;
  const draftDirty = Boolean(selectedExample && draftSource !== selectedExample.source);
  const assetFiles = workspaceFiles.filter((file) => file.kind === "asset");
  const dashboardUrls = runtimeStatus
    ? {
        dashboard: withEmbeddedDashboardTheme(runtimeStatus.dashboardUrl, editorTheme),
        session: withEmbeddedDashboardTheme(runtimeStatus.sessionUrl, editorTheme),
        health: withEmbeddedDashboardTheme(runtimeStatus.healthUrl, editorTheme),
        metrics: withEmbeddedDashboardTheme(runtimeStatus.metricsUrl, editorTheme),
        snapshot: withEmbeddedDashboardTheme(runtimeStatus.snapshotUrl, editorTheme),
      }
    : null;
  const activeDashboardUrl = dashboardUrls ? dashboardUrls[dashboardTab] : "";
  const activeDashboardLabel = dashboardTabLabel(dashboardTab);
  const dashboardBindingKey = dashboardTab === "session" ? activeSession?.sessionId || "idle" : "runtime";
  const dashboardFrameKey = `${dashboardTab}:${dashboardRefreshToken}:${runtimeStatus?.healthy ? "healthy" : "unhealthy"}:${dashboardBindingKey}`;
  const demoScene = demoMode ? DEMO_SCENES[demoSceneIndex] || null : null;
  const statusLineNumber = highlightedLine || selectedCodeLine || null;
  const statusWorkspacePath = selectedExample ? exampleWorkspacePath(selectedExample) : "No file selected";
  const statusRuntimeTarget = runtimeStatus?.healthy ? runtimeStatus.adminListen : runtimeStatus?.requestedAdminListen || "runtime unavailable";
  const statusPrimaryItem = appMode === "dashboard" ? activeDashboardUrl || "No runtime url available" : statusWorkspacePath;
  const statusModeItem = appMode === "dashboard" ? "Dashboard mode" : workspaceTabLabel(workspaceTab);
  const statusDetailItem = appMode === "dashboard" ? activeDashboardLabel : statusLineNumber ? `Line ${statusLineNumber}` : "Line -";
  const toggleMenu = (menuId: string) => {
    setOpenMenuId((current) => (current === menuId ? null : menuId));
  };
  const closeMenu = () => setOpenMenuId(null);
  const menuDefinitions: {
    id: string;
    label: string;
    items: {
      label: string;
      description: string;
      action: () => void;
      disabled?: boolean;
      active?: boolean;
    }[];
  }[] = [
    {
      id: "file",
      label: "File",
      items: [
        {
          label: "New File",
          description: "Create a new Node.js or Python example",
          action: () => {
            closeMenu();
            void handleCreateExample();
          },
          disabled: !bridgeReady,
        },
        {
          label: "Save",
          description: selectedExample ? `Save ${selectedExample.fileName}` : "No file selected",
          action: () => {
            closeMenu();
            void handleSaveExample();
          },
          disabled: !bridgeReady || !draftDirty || !selectedExample,
        },
        {
          label: "Clone",
          description: selectedExample ? `Clone ${selectedExample.fileName}` : "No file selected",
          action: () => {
            closeMenu();
            void handleCloneExample();
          },
          disabled: !bridgeReady || !selectedExample,
        },
        {
          label: "Rename",
          description: selectedExample ? `Rename ${selectedExample.fileName}` : "No file selected",
          action: () => {
            closeMenu();
            void handleRenameExample();
          },
          disabled: !bridgeReady || !selectedExample,
        },
      ],
    },
    {
      id: "edit",
      label: "Edit",
      items: [
        {
          label: "Revert Draft",
          description: selectedExample ? `Restore saved source for ${selectedExample.fileName}` : "No file selected",
          action: () => {
            closeMenu();
            handleRevertExample();
          },
          disabled: !draftDirty || !selectedExample,
        },
        {
          label: "Clear Target Line",
          description: "Remove the current run-to-line target",
          action: () => {
            closeMenu();
            setSelectedCodeLine(null);
          },
          disabled: selectedCodeLine === null,
        },
      ],
    },
    {
      id: "selection",
      label: "Selection",
      items: [
        {
          label: "Authoring View",
          description: "Focus the code editor workspace",
          action: () => {
            closeMenu();
            setWorkspaceTab("authoring");
          },
          active: appMode === "ide" && workspaceTab === "authoring",
          disabled: appMode !== "ide",
        },
        {
          label: "Sessions View",
          description: "Inspect stored runs and diagnostics",
          action: () => {
            closeMenu();
            setWorkspaceTab("sessions");
          },
          active: appMode === "ide" && workspaceTab === "sessions",
          disabled: appMode !== "ide",
        },
        {
          label: "Runtime View",
          description: "Show the embedded runtime workspace",
          action: () => {
            closeMenu();
            setWorkspaceTab("runtime");
          },
          active: appMode === "ide" && workspaceTab === "runtime",
          disabled: appMode !== "ide",
        },
      ],
    },
    {
      id: "run",
      label: "Run",
      items: [
        {
          label: "Run",
          description: "Run the selected example immediately",
          action: () => {
            closeMenu();
            void handleRun();
          },
          disabled: !selectedExampleId || hasActiveRun || !bridgeReady,
        },
        {
          label: "Run Paused",
          description: "Start and pause on the first step",
          action: () => {
            closeMenu();
            void handleRunPaused();
          },
          disabled: !selectedExampleId || hasActiveRun || !bridgeReady || !canStartPaused,
        },
        {
          label: "Pause",
          description: "Pause the active run",
          action: () => {
            closeMenu();
            void handlePause();
          },
          disabled: !canPause || !bridgeReady,
        },
        {
          label: "Resume",
          description: "Resume the active paused run",
          action: () => {
            closeMenu();
            void handleResume();
          },
          disabled: !canResume || !bridgeReady,
        },
        {
          label: "Step",
          description: "Advance the paused run one statement",
          action: () => {
            closeMenu();
            void handleStep();
          },
          disabled: !canStep || !bridgeReady,
        },
        {
          label: "Stop",
          description: "Stop the active execution session",
          action: () => {
            closeMenu();
            void handleStop();
          },
          disabled: !hasActiveRun || !bridgeReady,
        },
      ],
    },
    {
      id: "view",
      label: "View",
      items: [
        {
          label: "Midnight IDE",
          description: "Use the dark editor shell",
          action: () => {
            closeMenu();
            setEditorTheme("dark");
          },
          active: editorTheme === "dark",
        },
        {
          label: "Dashboard Light",
          description: "Match the standalone white dashboard styling",
          action: () => {
            closeMenu();
            setEditorTheme("dashboard");
          },
          active: editorTheme === "dashboard",
        },
        {
          label: "IDE Mode",
          description: "Show the editor workspace",
          action: () => {
            closeMenu();
            handleSelectAppMode("ide");
          },
          active: appMode === "ide",
        },
        {
          label: "Dashboard Mode",
          description: "Promote runtime dashboards into the main pane",
          action: () => {
            closeMenu();
            handleSelectAppMode("dashboard");
          },
          active: appMode === "dashboard",
        },
      ],
    },
    {
      id: "tools",
      label: "Tools",
      items: [
        {
          label: "Refresh Runtime",
          description: "Reload runtime status and dashboard panes",
          action: () => {
            closeMenu();
            void refreshRuntimeStatus();
          },
          disabled: !bridgeReady,
        },
        {
          label: "Start Runtime",
          description: "Start the managed runtime process",
          action: () => {
            closeMenu();
            void handleStartRuntime();
          },
          disabled: !bridgeReady,
        },
        {
          label: "Restart Runtime",
          description: "Restart the managed runtime process",
          action: () => {
            closeMenu();
            void handleRestartRuntime();
          },
          disabled: !bridgeReady,
        },
        {
          label: "Reset Layout",
          description: "Restore the default sidebar width",
          action: () => {
            closeMenu();
            setSidebarWidth(SIDEBAR_DEFAULT);
          },
        },
      ],
    },
    {
      id: "help",
      label: "Help",
      items: [
        {
          label: "Current Style",
          description: editorThemeLabel(editorTheme),
          action: closeMenu,
          active: true,
        },
        {
          label: "Current Mode",
          description: appModeLabel(appMode),
          action: closeMenu,
          active: true,
        },
      ],
    },
  ];

  return (
    <div
      className={`${styles.root} ${editorTheme === "dashboard" ? styles.rootLightTheme : ""}`}
      style={dragging ? { userSelect: "none" } : undefined}
    >
      <header ref={menuBarRef} className={styles.menuBar}>
        <div className={styles.modeSwitch} role="group" aria-label="Application mode">
          <button
            type="button"
            className={`${styles.modeSwitchButton} ${appMode === "ide" ? styles.modeSwitchButtonActive : ""}`}
            onClick={() => handleSelectAppMode("ide")}
          >
            IDE
          </button>
          <button
            type="button"
            className={`${styles.modeSwitchButton} ${appMode === "dashboard" ? styles.modeSwitchButtonActive : ""}`}
            onClick={() => handleSelectAppMode("dashboard")}
          >
            Dashboard
          </button>
        </div>
        <div className={styles.menuList}>
          {menuDefinitions.map((menu) => (
            <div key={menu.id} className={styles.menuDropdown}>
              <button
                type="button"
                className={`${styles.menuItem} ${openMenuId === menu.id ? styles.menuItemActive : ""}`}
                aria-expanded={openMenuId === menu.id}
                onClick={() => toggleMenu(menu.id)}
              >
                {menu.label}
              </button>
              {openMenuId === menu.id ? (
                <div className={styles.menuPanel} role="menu" aria-label={`${menu.label} menu`}>
                  {menu.items.map((item) => (
                    <button
                      key={`${menu.id}:${item.label}`}
                      type="button"
                      role="menuitem"
                      className={`${styles.menuPanelItem} ${item.active ? styles.menuPanelItemActive : ""}`}
                      onClick={item.action}
                      disabled={item.disabled}
                    >
                      <span className={styles.menuPanelTitle}>{item.label}</span>
                      <span className={styles.menuPanelMeta}>{item.description}</span>
                    </button>
                  ))}
                </div>
              ) : null}
            </div>
          ))}
        </div>
      </header>

      <div className={styles.appBody}>
        <aside className={styles.sidebar} style={sidebarStyle}>
          <div className={styles.sidebarHeader}>
            <div>
              <h2>{appMode === "dashboard" ? "Dashboard" : "Explorer"}</h2>
              <div className={styles.sidebarMeta}>
                {appMode === "dashboard"
                  ? `${dashboardViewItems.length} views available`
                  : loadingExamples
                  ? "Loading workspace..."
                  : `${examples.length} files indexed`}
              </div>
            </div>
          </div>

          {!bridgeReady ? (
            <div className={styles.emptyState}>
              <strong>Electron bridge unavailable</strong>
              <p>Start the editor through the Electron shell to discover and run examples.</p>
            </div>
          ) : null}

          {appMode === "dashboard" ? (
            <>
              <section className={styles.sidebarPanel}>
                <div className={styles.sidebarSectionTitle}>Available Views</div>
                <div className={styles.dashboardTabList} role="tablist" aria-label="Dashboard views">
                  {dashboardViewItems.map((item) => (
                    <button
                      key={item.id}
                      type="button"
                      role="tab"
                      aria-selected={dashboardTab === item.id}
                      className={`${styles.dashboardTab} ${dashboardTab === item.id ? styles.dashboardTabActive : ""}`}
                      onClick={() => setDashboardTab(item.id)}
                    >
                      <span className={styles.dashboardTabTitle}>{item.title}</span>
                      <span className={styles.dashboardTabMeta}>{item.subtitle}</span>
                    </button>
                  ))}
                </div>
              </section>

              <section className={styles.sidebarPanel}>
                <div className={styles.sidebarSectionTitle}>Runtime</div>
                <dl className={styles.dashboardMeta}>
                  <div>
                    <dt>Pane</dt>
                    <dd>{activeDashboardLabel}</dd>
                  </div>
                  <div>
                    <dt>Runtime</dt>
                    <dd>{statusRuntimeTarget}</dd>
                  </div>
                  <div>
                    <dt>Health</dt>
                    <dd>{runtimeStatus?.healthy ? "healthy" : "unreachable"}</dd>
                  </div>
                  <div>
                    <dt>Active Run</dt>
                    <dd>{activeSession ? `${activeSession.exampleName} (${activeSession.sessionId})` : "no active run"}</dd>
                  </div>
                </dl>
                <div className={styles.dashboardHint}>
                  Switch views from the sidebar while keeping the selected dashboard surface in the main content pane.
                </div>
              </section>
            </>
          ) : (
            <div className={styles.treeRoot}>
              <button
                type="button"
                className={styles.treeFolderRow}
                onClick={() => toggleTreeNode("project")}
                aria-expanded={isTreeNodeExpanded("project")}
              >
                <span className={styles.treeCaret}>{isTreeNodeExpanded("project") ? "v" : ">"}</span>
                <span className={styles.treeFolderName}>sikuli-go</span>
                <span className={styles.treeFolderMeta}>{loadingExamples ? "indexing..." : `${examples.length} files`}</span>
              </button>

              {isTreeNodeExpanded("project") ? (
                <div className={styles.treeChildren}>
                  <div className={styles.treeSection}>
                    <button
                      type="button"
                      className={styles.treeFolderRow}
                      onClick={() => toggleTreeNode("views")}
                      aria-expanded={isTreeNodeExpanded("views")}
                    >
                      <span className={styles.treeCaret}>{isTreeNodeExpanded("views") ? "v" : ">"}</span>
                      <span className={styles.treeFolderName}>views</span>
                    </button>
                    {isTreeNodeExpanded("views") ? (
                      <div className={styles.treeChildren}>
                        {viewItems.map((item) => (
                          <button
                            key={item.id}
                            type="button"
                            className={`${styles.treeFileButton} ${workspaceTab === item.id ? styles.treeFileButtonActive : ""}`}
                            onClick={() => setWorkspaceTab(item.id)}
                          >
                            <span className={styles.treeSpacer} />
                            <span className={styles.treeFileStack}>
                              <span className={styles.treeFileName}>{item.title}</span>
                              <span className={styles.treeFileMeta}>{item.subtitle}</span>
                            </span>
                          </button>
                        ))}
                      </div>
                    ) : null}
                  </div>

                  {recentExamples.length > 0 ? (
                    <div className={styles.treeSection}>
                      <button
                        type="button"
                        className={styles.treeFolderRow}
                        onClick={() => toggleTreeNode("open-editors")}
                        aria-expanded={isTreeNodeExpanded("open-editors")}
                      >
                        <span className={styles.treeCaret}>{isTreeNodeExpanded("open-editors") ? "v" : ">"}</span>
                        <span className={styles.treeFolderName}>open editors</span>
                      </button>
                      {isTreeNodeExpanded("open-editors") ? (
                        <div className={styles.treeChildren}>
                          {recentExamples.map((example) => (
                            <button
                              key={`recent:${example.id}`}
                              type="button"
                              className={`${styles.treeFileButton} ${selectedExampleId === example.id ? styles.treeFileButtonActive : ""}`}
                              onClick={() => focusExample(example.id)}
                            >
                              <span className={styles.treeSpacer} />
                              <span className={styles.treeFileStack}>
                                <span className={styles.treeFileName}>{example.fileName}</span>
                                <span className={styles.treeFileMeta}>{exampleWorkspacePath(example)}</span>
                              </span>
                            </button>
                          ))}
                        </div>
                      ) : null}
                    </div>
                  ) : null}

                  <div className={styles.treeSection}>
                    <button
                      type="button"
                      className={styles.treeFolderRow}
                      onClick={() => toggleTreeNode("workspace")}
                      aria-expanded={isTreeNodeExpanded("workspace")}
                    >
                      <span className={styles.treeCaret}>{isTreeNodeExpanded("workspace") ? "v" : ">"}</span>
                      <span className={styles.treeFolderName}>workspace</span>
                    </button>
                    {isTreeNodeExpanded("workspace") ? (
                      <div className={styles.treeChildren}>
                        {examplesByPackage.map((group) => {
                          const packageNodeId = `package:${group.packageRoot}`;
                          const packageExamplesNodeId = `${packageNodeId}:examples`;
                          return (
                            <div key={group.packageRoot} className={styles.treeSection}>
                              <button
                                type="button"
                                className={styles.treeFolderRow}
                                onClick={() => toggleTreeNode(packageNodeId)}
                                aria-expanded={isTreeNodeExpanded(packageNodeId)}
                              >
                                <span className={styles.treeCaret}>{isTreeNodeExpanded(packageNodeId) ? "v" : ">"}</span>
                                <span className={styles.treeFolderName}>{group.label}</span>
                              </button>
                              {isTreeNodeExpanded(packageNodeId) ? (
                                <div className={styles.treeChildren}>
                                  <button
                                    type="button"
                                    className={styles.treeFolderRow}
                                    onClick={() => toggleTreeNode(packageExamplesNodeId)}
                                    aria-expanded={isTreeNodeExpanded(packageExamplesNodeId)}
                                  >
                                    <span className={styles.treeCaret}>{isTreeNodeExpanded(packageExamplesNodeId) ? "v" : ">"}</span>
                                    <span className={styles.treeFolderName}>examples</span>
                                  </button>
                                  {isTreeNodeExpanded(packageExamplesNodeId) ? (
                                    <div className={styles.treeChildren}>
                                      {group.items.map((example) => {
                                        const isSelected = selectedExampleId === example.id;
                                        return (
                                          <button
                                            key={example.id}
                                            type="button"
                                            className={`${styles.treeFileButton} ${isSelected ? styles.treeFileButtonActive : ""}`}
                                            onClick={() => focusExample(example.id)}
                                          >
                                            <span className={styles.treeSpacer} />
                                            <span className={styles.treeFileName}>{example.fileName}</span>
                                          </button>
                                        );
                                      })}
                                    </div>
                                  ) : null}
                                </div>
                              ) : null}
                            </div>
                          );
                        })}

                        {selectedExample ? (
                          <div className={styles.treeSection}>
                            <button
                              type="button"
                              className={styles.treeFolderRow}
                              onClick={() => toggleTreeNode("current-workspace")}
                              aria-expanded={isTreeNodeExpanded("current-workspace")}
                            >
                              <span className={styles.treeCaret}>{isTreeNodeExpanded("current-workspace") ? "v" : ">"}</span>
                              <span className={styles.treeFolderName}>current workspace</span>
                            </button>
                            {isTreeNodeExpanded("current-workspace") ? (
                              <div className={styles.treeChildren}>
                                <button
                                  type="button"
                                  className={`${styles.treeFileButton} ${styles.treeFileButtonActive}`}
                                  onClick={() => setWorkspaceTab("authoring")}
                                >
                                  <span className={styles.treeSpacer} />
                                  <span className={styles.treeFileName}>{selectedExample.fileName}</span>
                                </button>

                                {assetFiles.length > 0 ? (
                                  <div className={styles.treeSection}>
                                    <button
                                      type="button"
                                      className={styles.treeFolderRow}
                                      onClick={() => toggleTreeNode("current-workspace:assets")}
                                      aria-expanded={isTreeNodeExpanded("current-workspace:assets")}
                                    >
                                      <span className={styles.treeCaret}>{isTreeNodeExpanded("current-workspace:assets") ? "v" : ">"}</span>
                                      <span className={styles.treeFolderName}>assets</span>
                                    </button>
                                    {isTreeNodeExpanded("current-workspace:assets") ? (
                                      <div className={styles.treeChildren}>
                                        {assetFiles.map((file) => (
                                          <button
                                            key={file.id}
                                            type="button"
                                            className={`${styles.treeFileButton} ${selectedWorkspaceFilePath === file.relativeWorkspacePath ? styles.treeFileButtonActive : ""}`}
                                            onClick={() => {
                                              setWorkspaceTab("authoring");
                                              setSelectedWorkspaceFilePath(file.relativeWorkspacePath);
                                            }}
                                          >
                                            <span className={styles.treeSpacer} />
                                            <span className={styles.treeFileName}>{file.fileName}</span>
                                          </button>
                                        ))}
                                      </div>
                                    ) : null}
                                  </div>
                                ) : null}
                              </div>
                            ) : null}
                          </div>
                        ) : null}
                      </div>
                    ) : null}
                  </div>
                </div>
              ) : null}
            </div>
          )}
        </aside>

        <div
          className={styles.resizer}
          role="separator"
          aria-orientation="vertical"
          aria-label="Resize explorer sidebar"
          onPointerDown={onResizePointerDown}
          onDoubleClick={() => setSidebarWidth(SIDEBAR_DEFAULT)}
        />

        <main className={`${styles.main} ${appMode === "dashboard" ? styles.mainDashboardMode : ""}`}>
          {demoScene ? (
            <section className={styles.demoBanner}>
              <div>
                <div className={styles.eyebrow}>Demo Reel</div>
                <div className={styles.demoTitle}>Editor walkthrough for capture output</div>
                <p className={styles.demoCopy}>{demoScene.label}</p>
              </div>
              <div className={styles.demoProgress} aria-label="Demo progress">
                {DEMO_SCENES.map((scene, index) => (
                  <span
                    key={scene.label}
                    className={`${styles.demoProgressDot} ${index === demoSceneIndex ? styles.demoProgressDotActive : ""}`}
                    aria-hidden="true"
                  />
                ))}
              </div>
            </section>
          ) : null}

          <section className={styles.commandBar}>
            <div className={styles.commandContext}>
              <div className={styles.commandPath}>
                {appMode === "dashboard"
                  ? activeDashboardLabel
                  : selectedExample
                  ? exampleWorkspacePath(selectedExample)
                  : "Select a workspace file"}
              </div>
              <div className={styles.commandMeta}>
                {appMode === "dashboard"
                  ? `${activeDashboardUrl || "No runtime url available"} • ${appModeLabel(appMode)} mode`
                  : `${selectedExample ? selectedExample.commandPreview : "No command selected"} • ${workspaceTabLabel(workspaceTab)} view`}
                {selectedSession ? ` • Session ${selectedSession.sessionId}` : ""}
              </div>
            </div>

            <div className={styles.toolbar}>
              {appMode === "dashboard" ? (
                <>
                  <div className={`${styles.badge} ${runtimeStatus?.healthy ? styles.badgeRunning : styles.badgeFailed}`}>
                    {runtimeStatus?.healthy ? "runtime healthy" : "runtime unavailable"}
                  </div>
                  <button type="button" className={styles.primaryButton} onClick={handleStartRuntime} disabled={!bridgeReady}>
                    Start Runtime
                  </button>
                  <button type="button" className={styles.secondaryButton} onClick={handleRestartRuntime} disabled={!bridgeReady}>
                    Restart Runtime
                  </button>
                  <button type="button" className={styles.secondaryButton} onClick={handleStopRuntime} disabled={!bridgeReady}>
                    Stop Runtime
                  </button>
                  <button type="button" className={styles.secondaryButton} onClick={() => void refreshRuntimeStatus()} disabled={!bridgeReady}>
                    Reload Pane
                  </button>
                </>
              ) : (
                <>
                  <div className={`${styles.badge} ${badgeTone(activeRunState)}`}>Run {activeRunState}</div>
                  <button type="button" className={styles.primaryButton} onClick={handleRun} disabled={!selectedExampleId || hasActiveRun || !bridgeReady}>
                    {hasActiveRun ? "Run Active" : "Run"}
                  </button>
                  <button
                    type="button"
                    className={styles.secondaryButton}
                    onClick={handleRunPaused}
                    disabled={!selectedExampleId || hasActiveRun || !bridgeReady || !canStartPaused}
                  >
                    Run Paused
                  </button>
                  <button type="button" className={styles.secondaryButton} onClick={handlePause} disabled={!canPause || !bridgeReady}>
                    Pause
                  </button>
                  <button type="button" className={styles.secondaryButton} onClick={handleResume} disabled={!canResume || !bridgeReady}>
                    Resume
                  </button>
                  <button type="button" className={styles.secondaryButton} onClick={handleStep} disabled={!canStep || !bridgeReady}>
                    Step
                  </button>
                  <button
                    type="button"
                    className={styles.secondaryButton}
                    onClick={handleContinueToRuntimeCall}
                    disabled={!canContinueToRuntimeCall || !bridgeReady}
                  >
                    To Runtime Call
                  </button>
                  <button
                    type="button"
                    className={styles.secondaryButton}
                    onClick={handleRunToLine}
                    disabled={!canRunToLine || !bridgeReady}
                  >
                    {selectedCodeLine ? `To Line ${selectedCodeLine}` : "Run To Line"}
                  </button>
                  <button type="button" className={styles.secondaryButton} onClick={handleStop} disabled={!hasActiveRun || !bridgeReady}>
                    Stop
                  </button>
                </>
              )}
            </div>
          </section>

          {actionError ? <div className={styles.errorBanner}>{actionError}</div> : null}

          {appMode === "dashboard" ? (
            <section className={styles.dashboardViewport}>
              {runtimeStatus && activeDashboardUrl ? (
                <iframe
                  key={dashboardFrameKey}
                  title={activeDashboardLabel}
                  src={activeDashboardUrl}
                  className={`${styles.dashboardFrame} ${styles.dashboardFrameFull}`}
                />
              ) : (
                <div className={styles.emptyState}>
                  <strong>No runtime pane yet</strong>
                  <p>Start the runtime and reload the pane to show the live dashboard surface in the main content area.</p>
                </div>
              )}
            </section>
          ) : (
            <>
        {workspaceTab === "authoring" ? (
          <div className={styles.workspaceGrid}>
            <section className={`${styles.panel} ${styles.editorSurfacePanel}`}>
              <div className={styles.editorTabBar}>
                <button type="button" className={`${styles.editorTab} ${styles.editorTabActive}`}>
                  {selectedExample?.fileName || "untitled"}
                </button>
                {selectedWorkspaceFile ? (
                  <button type="button" className={styles.editorTab}>
                    {selectedWorkspaceFile.fileName}
                  </button>
                ) : null}
              </div>

              {selectedLoadError ? <div className={styles.inlineError}>{selectedLoadError}</div> : null}

              {selectedExample ? (
                <>
                  <div className={styles.editorHeader}>
                    <div>
                      <div className={styles.panelTitle}>{selectedExample.name}</div>
                      <div className={styles.panelSubtitle}>{exampleWorkspacePath(selectedExample)}</div>
                    </div>
                    <div className={styles.editorHeaderMeta}>
                      <span className={`${styles.badge} ${badgeTone(selectedSessionState)}`}>{selectedSessionState}</span>
                      <span className={`${styles.badge} ${draftDirty ? styles.badgePaused : styles.badgeCompleted}`}>
                        {draftDirty ? "unsaved" : "saved"}
                      </span>
                    </div>
                  </div>

                  <div className={styles.editorToolbar}>
                    <button type="button" className={styles.primaryButton} onClick={handleSaveExample} disabled={!draftDirty || !bridgeReady}>
                      Save
                    </button>
                    <button type="button" className={styles.secondaryButton} onClick={handleRevertExample} disabled={!draftDirty}>
                      Revert
                    </button>
                    <button type="button" className={styles.secondaryButton} onClick={handleRenameExample} disabled={!selectedExample || !bridgeReady}>
                      Rename
                    </button>
                    <button type="button" className={styles.secondaryButton} onClick={handleCloneExample} disabled={!selectedExample || !bridgeReady}>
                      Clone
                    </button>
                    <button type="button" className={styles.secondaryButton} onClick={handleCreateExample} disabled={!bridgeReady}>
                      New File
                    </button>
                  </div>

                  <div className={styles.editorCanvas}>
                    <div className={styles.editorGutter}>
                      {editorLineNumbers.map((lineNumber) => (
                        <button
                          key={`line:${lineNumber}`}
                          type="button"
                          className={`${styles.editorGutterLine} ${highlightedLine === lineNumber ? styles.editorGutterLineActive : ""} ${selectedCodeLine === lineNumber ? styles.editorGutterLineTarget : ""}`}
                          onClick={() => setSelectedCodeLine(lineNumber)}
                        >
                          {lineNumber}
                        </button>
                      ))}
                    </div>
                    <textarea
                      className={styles.editorCodeTextarea}
                      value={draftSource}
                      onChange={(event) => {
                        setSaveMessage("");
                        setDraftSource(event.target.value);
                      }}
                      spellCheck={false}
                    />
                  </div>

                  <div className={styles.editorStatusBar}>
                    <span>Command: {selectedExample.commandPreview}</span>
                    <span>Line: {selectedCodeLine ?? "-"}</span>
                    <span>Current: {highlightedLine ?? "-"}</span>
                    <span>{draftDirty ? "Unsaved draft" : saveMessage || "Synced to disk"}</span>
                  </div>
                </>
              ) : (
                <div className={styles.emptyState}>
                  <strong>Select an example</strong>
                  <p>Choose a Node.js or Python example from the explorer to inspect and edit its source.</p>
                </div>
              )}
            </section>

            <section className={styles.stack}>
              <section className={styles.panel}>
                <div className={styles.panelHeader}>
                  <div>
                    <div className={styles.panelTitle}>Overview</div>
                    <div className={styles.panelSubtitle}>File, runtime binding, and selected line context</div>
                  </div>
                  <div className={`${styles.badge} ${badgeTone(selectedSessionState)}`}>{selectedSessionState}</div>
                </div>

                {selectedExample ? (
                  <dl className={styles.metaGrid}>
                    <div>
                      <dt>File</dt>
                      <dd>{selectedExample.fileName}</dd>
                    </div>
                    <div>
                      <dt>Package</dt>
                      <dd>{packageLabel(selectedExample.packageRoot)}</dd>
                    </div>
                    <div>
                      <dt>Language</dt>
                      <dd>{selectedExample.languageLabel}</dd>
                    </div>
                    <div>
                      <dt>Stepping</dt>
                      <dd>{selectedExample.language === "nodejs" || selectedExample.language === "python" ? "available" : "not available"}</dd>
                    </div>
                    <div>
                      <dt>Current Line</dt>
                      <dd>{highlightedLine ?? "none"}</dd>
                    </div>
                    <div>
                      <dt>Target Line</dt>
                      <dd>{selectedCodeLine ?? "none"}</dd>
                    </div>
                    <div>
                      <dt>Runtime</dt>
                      <dd>{runtimeStatus?.adminListen || "not connected"}</dd>
                    </div>
                    <div>
                      <dt>Last Call</dt>
                      <dd>{runtimeCallLabel(selectedRuntimeCall)}</dd>
                    </div>
                  </dl>
                ) : (
                  <div className={styles.emptyState}>
                    <strong>No file selected</strong>
                    <p>Select a file to inspect its metadata and run context.</p>
                  </div>
                )}
              </section>

              <section className={styles.panel}>
                <div className={styles.panelHeader}>
                  <div>
                    <div className={styles.panelTitle}>Workspace Assets</div>
                    <div className={styles.panelSubtitle}>Preview assets that live beside the current script</div>
                  </div>
                  <div className={styles.badge}>{assetFiles.length} assets</div>
                </div>

                {workspaceLoadError ? <div className={styles.inlineError}>{workspaceLoadError}</div> : null}

                {assetFiles.length > 0 ? (
                  <div className={styles.assetGrid}>
                    <ul className={styles.assetList}>
                      {assetFiles.map((file) => (
                        <li key={file.id}>
                          <button
                            type="button"
                            className={`${styles.assetItem} ${selectedWorkspaceFilePath === file.relativeWorkspacePath ? styles.assetItemActive : ""}`}
                            onClick={() => setSelectedWorkspaceFilePath(file.relativeWorkspacePath)}
                          >
                            <span className={styles.sessionItemTitle}>{file.fileName}</span>
                            <span className={styles.sessionItemMeta}>{file.relativeWorkspacePath}</span>
                            <span className={styles.sessionItemMeta}>{file.isImage ? "image" : file.isText ? "text" : "binary"} · {file.size} bytes</span>
                          </button>
                        </li>
                      ))}
                    </ul>

                    <div className={styles.assetPreview}>
                      {selectedWorkspaceFile ? (
                        selectedWorkspaceFile.isImage && selectedWorkspaceFile.dataUrl ? (
                          <img
                            src={selectedWorkspaceFile.dataUrl}
                            alt={selectedWorkspaceFile.fileName}
                            className={styles.assetPreviewImage}
                          />
                        ) : selectedWorkspaceFile.isText ? (
                          <pre className={styles.assetPreviewText}>{selectedWorkspaceFile.content}</pre>
                        ) : (
                          <div className={styles.emptyState}>
                            <strong>Binary asset</strong>
                            <p>{selectedWorkspaceFile.relativeWorkspacePath} cannot be previewed inline yet.</p>
                          </div>
                        )
                      ) : (
                        <div className={styles.emptyState}>
                          <strong>No asset selected</strong>
                          <p>Select an asset to preview it in the inspector.</p>
                        </div>
                      )}
                    </div>
                  </div>
                ) : (
                  <div className={styles.emptyState}>
                    <strong>No assets found</strong>
                    <p>The current workspace has no previewable asset files yet.</p>
                  </div>
                )}
              </section>
            </section>
          </div>
        ) : null}

        {workspaceTab === "sessions" ? (
          <div className={styles.workspaceGrid}>
            <section className={styles.stack}>
              <section className={styles.panel}>
                <div className={styles.panelHeader}>
                  <div>
                    <div className={styles.panelTitle}>Session History</div>
                    <div className={styles.panelSubtitle}>Recent runs across both client languages</div>
                  </div>
                  <div className={styles.badge}>{sessions.length} stored</div>
                </div>

                {sessions.length > 0 ? (
                  <div className={styles.sessionListPane}>
                    <ul className={styles.sessionList}>
                      {sessions.map((session) => (
                        <li key={session.sessionId}>
                          <button
                            type="button"
                            className={`${styles.sessionItem} ${selectedSessionId === session.sessionId ? styles.sessionItemActive : ""}`}
                            onClick={() => handleSelectSession(session)}
                          >
                            <span className={styles.sessionItemHeader}>
                              <span className={styles.sessionItemTitle}>{session.exampleName}</span>
                              <span className={`${styles.badge} ${badgeTone(session.state)}`}>{session.state}</span>
                            </span>
                            <span className={styles.sessionItemMeta}>
                              {session.languageLabel} · {session.relativePath}
                            </span>
                            <span className={styles.sessionItemMeta}>
                              {formatDateTime(session.startedAt)} · {formatDuration(session.startedAt, session.endedAt)}
                            </span>
                          </button>
                        </li>
                      ))}
                    </ul>
                  </div>
                ) : (
                  <div className={styles.emptyState}>
                    <strong>No sessions yet</strong>
                    <p>Run an example to create the first execution session and timeline.</p>
                  </div>
                )}
              </section>

              <section className={styles.panel}>
                <div className={styles.panelHeader}>
                  <div>
                    <div className={styles.panelTitle}>Selected Session</div>
                    <div className={styles.panelSubtitle}>
                      {selectedSession ? `Session ${selectedSession.sessionId}` : "Choose a run from session history"}
                    </div>
                  </div>
                  <div className={`${styles.badge} ${badgeTone(selectedSessionState)}`}>{selectedSessionState}</div>
                </div>

                {selectedSession ? (
                  <dl className={styles.metaGrid}>
                    <div>
                      <dt>Example</dt>
                      <dd>{selectedSession.exampleName}</dd>
                    </div>
                    <div>
                      <dt>Language</dt>
                      <dd>{selectedSession.languageLabel}</dd>
                    </div>
                    <div>
                      <dt>File</dt>
                      <dd>{selectedSession.relativePath}</dd>
                    </div>
                    <div>
                      <dt>PID</dt>
                      <dd>{selectedSession.pid ?? "not running"}</dd>
                    </div>
                    <div>
                      <dt>Stepping</dt>
                      <dd>{selectedSession.supportsStepping ? "enabled" : "not available"}</dd>
                    </div>
                    <div>
                      <dt>Started</dt>
                      <dd>{formatDateTime(selectedSession.startedAt)}</dd>
                    </div>
                    <div>
                      <dt>Ended</dt>
                      <dd>{formatDateTime(selectedSession.endedAt)}</dd>
                    </div>
                    <div>
                      <dt>Duration</dt>
                      <dd>{formatDuration(selectedSession.startedAt, selectedSession.endedAt)}</dd>
                    </div>
                    <div>
                      <dt>Current Line</dt>
                      <dd>{selectedSession.currentLine ?? "none"}</dd>
                    </div>
                    <div>
                      <dt>Pause Reason</dt>
                      <dd>{formatPauseReason(selectedSession.pauseReason)}</dd>
                    </div>
                    <div>
                      <dt>Run To Line</dt>
                      <dd>{selectedSession.runToLineTarget ?? "not armed"}</dd>
                    </div>
                    <div>
                      <dt>Awaiting Runtime Call</dt>
                      <dd>{selectedSession.awaitingRuntimeCall ? "armed" : "no"}</dd>
                    </div>
                    <div>
                      <dt>Statement</dt>
                      <dd>{selectedSession.currentStatementId ?? "none"}</dd>
                    </div>
                    <div>
                      <dt>Runtime Calls</dt>
                      <dd>{selectedSession.runtimeCallCount}</dd>
                    </div>
                    <div>
                      <dt>Runtime Failures</dt>
                      <dd>{selectedSession.runtimeErrorCount}</dd>
                    </div>
                    <div>
                      <dt>Exit</dt>
                      <dd>{selectedSession.exitCode ?? selectedSession.signal ?? "in progress"}</dd>
                    </div>
                    <div>
                      <dt>Command</dt>
                      <dd>{`${selectedSession.command} ${selectedSession.args.join(" ")}`.trim()}</dd>
                    </div>
                    <div>
                      <dt>Working Directory</dt>
                      <dd>{selectedSession.cwd}</dd>
                    </div>
                  </dl>
                ) : (
                  <div className={styles.emptyState}>
                    <strong>No session selected</strong>
                    <p>Pick a run from session history to inspect its metadata, output, and timeline.</p>
                  </div>
                )}
              </section>
            </section>

            <section className={styles.stack}>
              <section className={styles.panel}>
                <div className={styles.panelHeader}>
                  <div>
                    <div className={styles.panelTitle}>Runtime Diagnostics</div>
                    <div className={styles.panelSubtitle}>Latest runtime call, match metadata, and failure context for the selected session</div>
                  </div>
                  <div className={`${styles.badge} ${badgeTone(selectedRuntimeCall?.state || "")}`}>
                    {selectedRuntimeCall?.state || "idle"}
                  </div>
                </div>

                {selectedSession ? (
                  <dl className={styles.metaGrid}>
                    <div>
                      <dt>Current Call</dt>
                      <dd>{runtimeCallLabel(selectedSession.currentRuntimeCall)}</dd>
                    </div>
                    <div>
                      <dt>Latest Call</dt>
                      <dd>{runtimeCallLabel(selectedSession.lastRuntimeCall)}</dd>
                    </div>
                    <div>
                      <dt>Request</dt>
                      <dd>{selectedRuntimeCall?.requestSummary || "none"}</dd>
                    </div>
                    <div>
                      <dt>Response</dt>
                      <dd>{selectedRuntimeCall?.responseSummary || "none"}</dd>
                    </div>
                    <div>
                      <dt>Match Rect</dt>
                      <dd>{selectedRuntimeCall?.matchRect || "none"}</dd>
                    </div>
                    <div>
                      <dt>Target</dt>
                      <dd>{selectedRuntimeCall?.targetPoint || "none"}</dd>
                    </div>
                    <div>
                      <dt>Duration</dt>
                      <dd>{typeof selectedRuntimeCall?.durationMs === "number" ? `${selectedRuntimeCall.durationMs}ms` : "none"}</dd>
                    </div>
                    <div>
                      <dt>Trace ID</dt>
                      <dd>{selectedRuntimeCall?.traceId || "none"}</dd>
                    </div>
                    <div>
                      <dt>Score</dt>
                      <dd>{typeof selectedRuntimeCall?.score === "number" ? selectedRuntimeCall.score.toFixed(3) : "none"}</dd>
                    </div>
                    <div>
                      <dt>Exists</dt>
                      <dd>{typeof selectedRuntimeCall?.exists === "boolean" ? (selectedRuntimeCall.exists ? "yes" : "no") : "n/a"}</dd>
                    </div>
                    <div>
                      <dt>Error</dt>
                      <dd>{selectedRuntimeCall?.error || "none"}</dd>
                    </div>
                  </dl>
                ) : (
                  <div className={styles.emptyState}>
                    <strong>No session selected</strong>
                    <p>Pick a run from session history to inspect the latest runtime call and any attached match or failure details.</p>
                  </div>
                )}
              </section>

              <section className={styles.panel}>
                <div className={styles.panelHeader}>
                  <div>
                    <div className={styles.panelTitle}>Execution Output</div>
                    <div className={styles.panelSubtitle}>
                      {selectedSession ? "Buffered output for the selected session" : "Run an example to stream stdout and stderr here"}
                    </div>
                  </div>
                  <div className={styles.badge}>
                    {selectedSession?.pid ? `pid ${selectedSession.pid}` : "no process"}
                  </div>
                </div>

                <pre ref={outputRef} className={styles.outputPane}>
                  {selectedSession?.output || "Run an example to stream its output here."}
                </pre>
              </section>

              <section className={styles.panel}>
                <div className={styles.panelHeader}>
                  <div>
                    <div className={styles.panelTitle}>Timeline</div>
                    <div className={styles.panelSubtitle}>Lifecycle and output events for the selected session</div>
                  </div>
                  <div className={styles.badge}>
                    {selectedSession ? `${selectedSession.timeline.length} events` : "no session"}
                  </div>
                </div>

                {selectedSession && selectedSession.timeline.length > 0 ? (
                  <div className={styles.timelinePane}>
                    <ol className={styles.timelineList}>
                      {selectedSession.timeline.map((entry) => (
                        <li key={entry.id} className={styles.timelineItem}>
                          <div className={styles.timelineTime}>{formatDateTime(entry.emittedAt)}</div>
                          <div className={styles.timelineSummary}>{entry.summary}</div>
                        </li>
                      ))}
                    </ol>
                  </div>
                ) : (
                  <div className={styles.emptyState}>
                    <strong>No timeline yet</strong>
                    <p>The timeline will populate as soon as the selected session emits lifecycle or output events.</p>
                  </div>
                )}
              </section>
            </section>
          </div>
        ) : null}

        {workspaceTab === "runtime" ? (
          <>
            <section className={styles.panel}>
              <div className={styles.panelHeader}>
                <div>
                  <div className={styles.panelTitle}>Runtime</div>
                  <div className={styles.panelSubtitle}>Current desktop-host runtime wiring</div>
                </div>
                <button type="button" className={styles.linkButton} onClick={() => void refreshRuntimeStatus()} disabled={!bridgeReady}>
                  Refresh
                </button>
              </div>

              {runtimeStatus ? (
                <dl className={styles.metaGrid}>
                  <div>
                    <dt>gRPC</dt>
                    <dd>{runtimeStatus.apiListen}</dd>
                  </div>
                  <div>
                    <dt>Admin</dt>
                    <dd>{runtimeStatus.adminListen}</dd>
                  </div>
                  <div>
                    <dt>Requested gRPC</dt>
                    <dd>{runtimeStatus.requestedApiListen}</dd>
                  </div>
                  <div>
                    <dt>Requested Admin</dt>
                    <dd>{runtimeStatus.requestedAdminListen}</dd>
                  </div>
                  <div>
                    <dt>Health</dt>
                    <dd>{runtimeStatus.healthy ? "healthy" : "unreachable"}</dd>
                  </div>
                  <div>
                    <dt>Port Fallback</dt>
                    <dd>{runtimeStatus.portFallbackActive ? "active" : "not needed"}</dd>
                  </div>
                  <div>
                    <dt>Dashboard</dt>
                    <dd>{runtimeStatus.dashboardUrl}</dd>
                  </div>
                  <div>
                    <dt>Session Viewer</dt>
                    <dd>{runtimeStatus.sessionUrl}</dd>
                  </div>
                  <div>
                    <dt>Metrics</dt>
                    <dd>{runtimeStatus.metricsUrl}</dd>
                  </div>
                  <div>
                    <dt>Snapshot</dt>
                    <dd>{runtimeStatus.snapshotUrl}</dd>
                  </div>
                  <div>
                    <dt>Binary</dt>
                    <dd>{runtimeStatus.binaryPath}</dd>
                  </div>
                  <div>
                    <dt>Managed PID</dt>
                    <dd>{runtimeStatus.managedPid ?? "external or not started"}</dd>
                  </div>
                  <div>
                    <dt>Warning</dt>
                    <dd>{runtimeStatus.runtimeWarning || "none"}</dd>
                  </div>
                </dl>
              ) : (
                <div className={styles.emptyState}>
                  <strong>No runtime status yet</strong>
                  <p>Start the Electron shell and refresh runtime status to see the active addresses.</p>
                </div>
              )}
            </section>

            <section className={`${styles.panel} ${styles.dashboardWorkspace}`}>
          <div className={styles.panelHeader}>
            <div>
              <div className={styles.panelTitle}>Embedded Runtime Workspace</div>
              <div className={styles.panelSubtitle}>
                The live admin surfaces from the same runtime instance used by the current example execution
              </div>
            </div>
            <div className={styles.dashboardActions}>
              <div className={`${styles.badge} ${runtimeStatus?.healthy ? styles.badgeRunning : styles.badgeFailed}`}>
                {runtimeStatus?.healthy ? "runtime healthy" : "runtime unavailable"}
              </div>
              <button type="button" className={styles.secondaryButton} onClick={() => void refreshRuntimeStatus()} disabled={!bridgeReady}>
                Reload Pane
              </button>
            </div>
          </div>

          <div className={styles.dashboardGrid}>
            <section className={styles.dashboardSidebar}>
              <div className={styles.dashboardTabList} role="tablist" aria-label="Runtime workspace panes">
                {(["dashboard", "session", "health", "metrics", "snapshot"] as DashboardTab[]).map((tab) => (
                  <button
                    key={tab}
                    type="button"
                    role="tab"
                    aria-selected={dashboardTab === tab}
                    className={`${styles.dashboardTab} ${dashboardTab === tab ? styles.dashboardTabActive : ""}`}
                    onClick={() => setDashboardTab(tab)}
                  >
                    {dashboardTabLabel(tab)}
                  </button>
                ))}
              </div>

              <dl className={styles.dashboardMeta}>
                <div>
                  <dt>Pane</dt>
                  <dd>{activeDashboardLabel}</dd>
                </div>
                <div>
                  <dt>Runtime</dt>
                  <dd>{runtimeStatus?.adminListen || "not connected"}</dd>
                </div>
                <div>
                  <dt>URL</dt>
                  <dd>{activeDashboardUrl || "no runtime url available"}</dd>
                </div>
                <div>
                  <dt>Bound Run</dt>
                  <dd>{activeSession ? `${activeSession.exampleName} (${activeSession.sessionId})` : "no active run"}</dd>
                </div>
                <div>
                  <dt>Managed PID</dt>
                  <dd>{runtimeStatus?.managedPid ?? "external or not started"}</dd>
                </div>
                <div>
                  <dt>Health</dt>
                  <dd>{runtimeStatus?.healthy ? "healthy" : "unreachable"}</dd>
                </div>
              </dl>

              <div className={styles.dashboardHint}>
                The embedded panes follow the same runtime addresses used by the editor host. Restarting the runtime or refreshing this panel reloads the current view in place.
              </div>
            </section>

            <section className={styles.dashboardPane}>
              {runtimeStatus && activeDashboardUrl ? (
                <iframe
                  key={dashboardFrameKey}
                  title={activeDashboardLabel}
                  src={activeDashboardUrl}
                  className={styles.dashboardFrame}
                />
              ) : (
                <div className={styles.emptyState}>
                  <strong>No runtime pane yet</strong>
                  <p>Start the runtime and refresh status to load the live dashboard surfaces inside the IDE.</p>
                </div>
              )}
            </section>
          </div>
        </section>
          </>
        ) : null}
            </>
          )}
      </main>
      </div>

      <footer className={styles.statusBar}>
        <div className={styles.statusBarSection}>
          <span className={styles.statusBarItem}>{statusPrimaryItem}</span>
          <span className={styles.statusBarItem}>{statusModeItem}</span>
          <span className={styles.statusBarItem}>{statusDetailItem}</span>
        </div>
        <div className={styles.statusBarSection}>
          <span className={styles.statusBarItem}>{statusRuntimeTarget}</span>
          <span className={`${styles.statusBarItem} ${styles.statusBarBadge} ${badgeTone(activeRunState)}`}>{activeRunState}</span>
        </div>
      </footer>
    </div>
  );
}
