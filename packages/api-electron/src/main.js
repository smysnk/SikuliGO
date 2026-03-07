const { app, BrowserWindow, Menu } = require('electron');
const { spawn } = require('child_process');
const http = require('http');
const path = require('path');

const API_LISTEN = process.env.SIKULIGO_API_LISTEN || '127.0.0.1:50051';
const ADMIN_LISTEN = process.env.SIKULIGO_ADMIN_LISTEN || '127.0.0.1:8080';
const ADMIN_HEALTH_URL = process.env.SIKULIGO_ADMIN_HEALTH_URL || `http://${ADMIN_LISTEN}/healthz`;
const DASHBOARD_URL = process.env.SIKULIGO_DASHBOARD_URL || `http://${ADMIN_LISTEN}/dashboard`;
const SESSION_URL = process.env.SIKULIGO_SESSION_VIEW_URL || `${DASHBOARD_URL}?view=session-viewer`;
const API_BINARY_PATH =
  process.env.SIKULIGO_BINARY_PATH ||
  path.resolve(__dirname, '../../../sikuli-go');
const API_AUTO_START = process.env.SIKULIGO_API_AUTO_START !== '0';
const API_STARTUP_TIMEOUT_MS = Number(process.env.SIKULIGO_API_STARTUP_TIMEOUT_MS || '8000');

let mainWindow = null;
let managedApiProcess = null;

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
  if (await isHealthy(ADMIN_HEALTH_URL)) {
    return { ok: true, external: true };
  }
  if (managedApiProcess && !managedApiProcess.killed) {
    return { ok: await waitUntilHealthy(ADMIN_HEALTH_URL, API_STARTUP_TIMEOUT_MS) };
  }

  const args = ['-listen', API_LISTEN, '-admin-listen', ADMIN_LISTEN];
  managedApiProcess = spawn(API_BINARY_PATH, args, {
    stdio: 'inherit'
  });

  managedApiProcess.once('exit', (code, signal) => {
    console.error(`sikuli-go exited code=${code} signal=${signal}`);
    managedApiProcess = null;
  });

  const ok = await waitUntilHealthy(ADMIN_HEALTH_URL, API_STARTUP_TIMEOUT_MS);
  return { ok, external: false };
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

async function createWindow() {
  const win = new BrowserWindow({
    width: 1440,
    height: 900,
    title: 'SikuliGO Dashboard',
    webPreferences: {
      preload: path.join(__dirname, 'preload.js')
    }
  });
  mainWindow = win;

  const loadDashboard = () => win.loadURL(DASHBOARD_URL);
  const loadSessions = () => win.loadURL(SESSION_URL);

  const template = [
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
        `Unable to reach SikuliGO admin health endpoint: ${ADMIN_HEALTH_URL}. ` +
          `Set SIKULIGO_BINARY_PATH or start SikuliGO manually.`
      );
    }
  }
  loadDashboard();
}

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
  stopApi();
});
