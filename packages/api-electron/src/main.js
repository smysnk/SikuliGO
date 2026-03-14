const { app, BrowserWindow, Menu, ipcMain } = require('electron');
const { spawn } = require('child_process');
const http = require('http');
const path = require('path');
const { ExecutionManager } = require('./execution');
const {
  cloneExample,
  createExample,
  listExamples,
  listWorkspaceFiles,
  readExample,
  readWorkspaceFile,
  renameExample,
  saveExample,
} = require('./examples');
const { buildAdminUrls, resolveRuntimeListenConfig } = require('./runtime-addresses');

const REQUESTED_API_LISTEN = process.env.SIKULI_GO_API_LISTEN || '127.0.0.1:50051';
const REQUESTED_ADMIN_LISTEN = process.env.SIKULI_GO_ADMIN_LISTEN || '127.0.0.1:8080';
const API_LISTEN_EXPLICIT = Boolean(process.env.SIKULI_GO_API_LISTEN);
const ADMIN_LISTEN_EXPLICIT = Boolean(process.env.SIKULI_GO_ADMIN_LISTEN);
const EDITOR_URL = process.env.SIKULI_GO_EDITOR_URL || '';
const INITIAL_VIEW = process.env.SIKULI_GO_INITIAL_VIEW || (EDITOR_URL ? 'editor' : 'dashboard');
const API_BINARY_PATH =
  process.env.SIKULI_GO_BINARY_PATH || path.resolve(__dirname, '../../../sikuli-go');
const API_AUTO_START = process.env.SIKULI_GO_API_AUTO_START !== '0';
const API_STARTUP_TIMEOUT_MS = Number(process.env.SIKULI_GO_API_STARTUP_TIMEOUT_MS || '8000');

let mainWindow = null;
let managedApiProcess = null;
const runtimeState = {
  requestedApiListen: REQUESTED_API_LISTEN,
  requestedAdminListen: REQUESTED_ADMIN_LISTEN,
  apiListen: REQUESTED_API_LISTEN,
  adminListen: REQUESTED_ADMIN_LISTEN,
  portFallbackActive: false,
  runtimeWarning: '',
};

function currentAdminUrls() {
  return buildAdminUrls(runtimeState.adminListen);
}

function requestedAdminUrls() {
  return buildAdminUrls(runtimeState.requestedAdminListen);
}

const executionManager = new ExecutionManager({
  startRuntime: async () => await startApi(),
  runtimeEnv: () => ({
    SIKULI_GO_BINARY_PATH: API_BINARY_PATH,
    SIKULI_GRPC_ADDR: runtimeState.apiListen,
  }),
  onEvent: (event) => {
    if (mainWindow && !mainWindow.isDestroyed()) {
      mainWindow.webContents.send('ide:execution-event', event);
    }
  }
});

function isHealthy(url) {
  return new Promise((resolve) => {
    const req = http.get(url, (res) => {
      res.resume();
      resolve(res.statusCode === 200);
    });
    req.on('error', () => resolve(false));
    req.setTimeout(800, () => {
      req.destroy();
      resolve(false);
    });
  });
}

async function waitUntilHealthy(url, timeoutMs) {
  const startedAt = Date.now();
  while (Date.now() - startedAt < timeoutMs) {
    if (await isHealthy(url)) {
      return true;
    }
    await new Promise((r) => setTimeout(r, 200));
  }
  return false;
}

async function startApi() {
  if (await isHealthy(requestedAdminUrls().healthUrl)) {
    runtimeState.apiListen = runtimeState.requestedApiListen;
    runtimeState.adminListen = runtimeState.requestedAdminListen;
    runtimeState.portFallbackActive = false;
    runtimeState.runtimeWarning = '';
    return { ok: true, external: true };
  }
  if (managedApiProcess && !managedApiProcess.killed) {
    return {
      ok: await waitUntilHealthy(currentAdminUrls().healthUrl, API_STARTUP_TIMEOUT_MS),
      external: false,
      fallbackActive: runtimeState.portFallbackActive,
      warning: runtimeState.runtimeWarning,
    };
  }

  const resolvedRuntime = await resolveRuntimeListenConfig({
    apiListen: runtimeState.requestedApiListen,
    adminListen: runtimeState.requestedAdminListen,
    allowApiFallback: !API_LISTEN_EXPLICIT,
    allowAdminFallback: !ADMIN_LISTEN_EXPLICIT,
  });

  runtimeState.apiListen = resolvedRuntime.apiListen;
  runtimeState.adminListen = resolvedRuntime.adminListen;
  runtimeState.portFallbackActive = resolvedRuntime.fallbackActive;
  runtimeState.runtimeWarning = resolvedRuntime.warning;

  if (resolvedRuntime.warning && !resolvedRuntime.fallbackActive) {
    return {
      ok: false,
      error: resolvedRuntime.warning,
      external: false,
    };
  }

  const args = ['-listen', runtimeState.apiListen, '-admin-listen', runtimeState.adminListen];
  managedApiProcess = spawn(API_BINARY_PATH, args, {
    stdio: 'inherit'
  });

  managedApiProcess.once('exit', (code, signal) => {
    console.error(`sikuli-go exited code=${code} signal=${signal}`);
    managedApiProcess = null;
  });

  const ok = await waitUntilHealthy(currentAdminUrls().healthUrl, API_STARTUP_TIMEOUT_MS);
  return {
    ok,
    external: false,
    fallbackActive: runtimeState.portFallbackActive,
    warning: runtimeState.runtimeWarning,
    error: ok ? '' : `Unable to reach sikuli-go admin health endpoint: ${currentAdminUrls().healthUrl}`,
  };
}

function stopApi() {
  if (!managedApiProcess) {
    return false;
  }
  const child = managedApiProcess;
  managedApiProcess = null;
  if (process.platform === 'win32') {
    spawn('taskkill', ['/pid', String(child.pid), '/f', '/t'], { stdio: 'ignore' });
  } else {
    child.kill('SIGTERM');
  }
  return true;
}

async function restartApi() {
  stopApi();
  return startApi();
}

async function runtimeStatus() {
  const adminUrls = currentAdminUrls();
  return {
    apiListen: runtimeState.apiListen,
    adminListen: runtimeState.adminListen,
    requestedApiListen: runtimeState.requestedApiListen,
    requestedAdminListen: runtimeState.requestedAdminListen,
    portFallbackActive: runtimeState.portFallbackActive,
    runtimeWarning: runtimeState.runtimeWarning,
    healthy: await isHealthy(adminUrls.healthUrl),
    dashboardUrl: adminUrls.dashboardUrl,
    sessionUrl: adminUrls.sessionUrl,
    healthUrl: adminUrls.healthUrl,
    metricsUrl: adminUrls.metricsUrl,
    snapshotUrl: adminUrls.snapshotUrl,
    binaryPath: API_BINARY_PATH,
    managedPid: managedApiProcess ? managedApiProcess.pid : null,
  };
}

function registerIpcHandlers() {
  ipcMain.handle('ide:list-examples', async () => listExamples());
  ipcMain.handle('ide:read-example', async (_event, exampleId) => readExample(exampleId));
  ipcMain.handle('ide:save-example', async (_event, exampleId, source) => saveExample(exampleId, source));
  ipcMain.handle('ide:create-example', async (_event, language, fileName, source) => createExample(language, fileName, source));
  ipcMain.handle('ide:clone-example', async (_event, exampleId, fileName, source) => cloneExample(exampleId, fileName, source));
  ipcMain.handle('ide:rename-example', async (_event, exampleId, fileName, source) => renameExample(exampleId, fileName, source));
  ipcMain.handle('ide:list-workspace-files', async (_event, exampleId) => listWorkspaceFiles(exampleId));
  ipcMain.handle('ide:read-workspace-file', async (_event, exampleId, relativeWorkspacePath) => readWorkspaceFile(exampleId, relativeWorkspacePath));
  ipcMain.handle('ide:run-example', async (_event, exampleId, options) => await executionManager.runExample(exampleId, options || {}));
  ipcMain.handle('ide:stop-execution', async () => await executionManager.stopCurrentRun());
  ipcMain.handle('ide:pause-execution', async () => await executionManager.pauseCurrentRun());
  ipcMain.handle('ide:resume-execution', async () => await executionManager.resumeCurrentRun());
  ipcMain.handle('ide:step-execution', async () => await executionManager.stepCurrentRun());
  ipcMain.handle('ide:run-to-line', async (_event, line) => await executionManager.runCurrentRunToLine(line));
  ipcMain.handle('ide:continue-to-runtime-call', async () => await executionManager.continueCurrentRunToRuntimeCall());
  ipcMain.handle('ide:get-execution-state', async () => executionManager.getCurrentRun());
  ipcMain.handle('ide:list-sessions', async () => executionManager.listSessions());
  ipcMain.handle('ide:get-session', async (_event, sessionId) => executionManager.getSession(sessionId));
  ipcMain.handle('ide:get-runtime-status', async () => await runtimeStatus());
  ipcMain.handle('ide:start-runtime', async () => {
    const result = await startApi();
    if (!result.ok) {
      throw new Error(result.error || 'Runtime failed to start');
    }
    return result;
  });
  ipcMain.handle('ide:restart-runtime', async () => {
    const result = await restartApi();
    if (!result.ok) {
      throw new Error(result.error || 'Runtime failed to restart');
    }
    return result;
  });
  ipcMain.handle('ide:stop-runtime', async () => stopApi());
}

async function createWindow() {
  const win = new BrowserWindow({
    width: 1440,
    height: 900,
    title: 'sikuli-go IDE',
    webPreferences: {
      preload: path.join(__dirname, 'preload.js'),
      contextIsolation: true,
      nodeIntegration: false
    }
  });
  mainWindow = win;

  const loadEditor = () => {
    if (!EDITOR_URL) {
      return loadDashboard();
    }
    return win.loadURL(EDITOR_URL);
  };
  const loadDashboard = () => win.loadURL(currentAdminUrls().dashboardUrl);
  const loadSessions = () => win.loadURL(currentAdminUrls().sessionUrl);

  const template = [
    {
      label: 'Workspace',
      submenu: [
        ...(EDITOR_URL ? [{ label: 'Editor', click: loadEditor }] : []),
        { label: 'Live Dashboard', click: loadDashboard },
        { label: 'Session Viewer', click: loadSessions }
      ]
    },
    {
      label: 'API',
      submenu: [
        {
          label: 'Start API',
          click: async () => {
            const started = await startApi();
            if (started.ok && mainWindow) {
              loadDashboard();
            }
          }
        },
        {
          label: 'Restart API',
          click: async () => {
            const restarted = await restartApi();
            if (restarted.ok && mainWindow) {
              loadDashboard();
            }
          }
        },
        {
          label: 'Stop API',
          click: () => {
            stopApi();
          }
        }
      ]
    },
    {
      label: 'View',
      submenu: [
        ...(EDITOR_URL ? [{ label: 'Editor', click: loadEditor }] : []),
        { label: 'Live Dashboard', click: loadDashboard },
        { label: 'Session Viewer', click: loadSessions },
        { type: 'separator' },
        { role: 'reload' },
        { role: 'toggleDevTools' }
      ]
    }
  ];

  Menu.setApplicationMenu(Menu.buildFromTemplate(template));
  if (API_AUTO_START) {
    const started = await startApi();
    if (!started.ok) {
      console.error(
        `${started.error || `Unable to reach sikuli-go admin health endpoint: ${currentAdminUrls().healthUrl}`}. ` +
          `Set SIKULI_GO_BINARY_PATH or start sikuli-go manually.`
      );
    }
  }
  if (INITIAL_VIEW === 'editor' && EDITOR_URL) {
    await loadEditor();
  } else {
    await loadDashboard();
  }
}

registerIpcHandlers();

app.whenReady().then(() => {
  createWindow();
  app.on('activate', () => {
    if (BrowserWindow.getAllWindows().length === 0) {
      createWindow();
    }
  });
});

app.on('window-all-closed', () => {
  stopApi();
  if (process.platform !== 'darwin') {
    app.quit();
  }
});

app.on('before-quit', () => {
  executionManager.stopCurrentRun();
  stopApi();
});
