const { contextBridge, ipcRenderer } = require("electron");

contextBridge.exposeInMainWorld("sikuliIde", {
  listExamples: async () => await ipcRenderer.invoke("ide:list-examples"),
  readExample: async (exampleId) => await ipcRenderer.invoke("ide:read-example", exampleId),
  saveExample: async (exampleId, source) => await ipcRenderer.invoke("ide:save-example", exampleId, source),
  createExample: async (language, fileName, source) => await ipcRenderer.invoke("ide:create-example", language, fileName, source),
  cloneExample: async (exampleId, fileName, source) => await ipcRenderer.invoke("ide:clone-example", exampleId, fileName, source),
  renameExample: async (exampleId, fileName, source) => await ipcRenderer.invoke("ide:rename-example", exampleId, fileName, source),
  listWorkspaceFiles: async (exampleId) => await ipcRenderer.invoke("ide:list-workspace-files", exampleId),
  readWorkspaceFile: async (exampleId, relativeWorkspacePath) => await ipcRenderer.invoke("ide:read-workspace-file", exampleId, relativeWorkspacePath),
  runExample: async (exampleId, options) => await ipcRenderer.invoke("ide:run-example", exampleId, options || {}),
  stopExecution: async () => await ipcRenderer.invoke("ide:stop-execution"),
  pauseExecution: async () => await ipcRenderer.invoke("ide:pause-execution"),
  resumeExecution: async () => await ipcRenderer.invoke("ide:resume-execution"),
  stepExecution: async () => await ipcRenderer.invoke("ide:step-execution"),
  runToLine: async (line) => await ipcRenderer.invoke("ide:run-to-line", line),
  continueToRuntimeCall: async () => await ipcRenderer.invoke("ide:continue-to-runtime-call"),
  getExecutionState: async () => await ipcRenderer.invoke("ide:get-execution-state"),
  listSessions: async () => await ipcRenderer.invoke("ide:list-sessions"),
  getSession: async (sessionId) => await ipcRenderer.invoke("ide:get-session", sessionId),
  getRuntimeStatus: async () => await ipcRenderer.invoke("ide:get-runtime-status"),
  startRuntime: async () => await ipcRenderer.invoke("ide:start-runtime"),
  restartRuntime: async () => await ipcRenderer.invoke("ide:restart-runtime"),
  stopRuntime: async () => await ipcRenderer.invoke("ide:stop-runtime"),
  onExecutionEvent: (handler) => {
    const wrapped = (_event, payload) => handler(payload);
    ipcRenderer.on("ide:execution-event", wrapped);
    return () => {
      ipcRenderer.removeListener("ide:execution-event", wrapped);
    };
  },
});
