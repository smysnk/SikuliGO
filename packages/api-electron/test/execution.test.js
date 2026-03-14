const fs = require("fs");
const os = require("os");
const path = require("path");
const test = require("node:test");
const assert = require("node:assert/strict");

const { ExecutionManager } = require("../src/execution");

function makeExample(tempRoot, language, fileName, source) {
  const packageRoot = path.join(tempRoot, language === "nodejs" ? "client-node" : "client-python");
  const examplesDir = path.join(packageRoot, "examples");
  fs.mkdirSync(examplesDir, { recursive: true });
  const absolutePath = path.join(examplesDir, fileName);
  fs.writeFileSync(absolutePath, source, "utf8");
  return {
    id: `${language}:${fileName}`,
    language,
    languageLabel: language === "nodejs" ? "Node.js" : "Python",
    name: fileName,
    fileName,
    relativePath: path.posix.join("examples", fileName),
    absolutePath,
    packageRoot,
    commandPreview: `${language === "nodejs" ? "node" : "python"} examples/${fileName}`,
    source,
  };
}

async function waitForFinalSession(manager, sessionId) {
  const deadline = Date.now() + 8_000;
  while (Date.now() < deadline) {
    const session = manager.getSession(sessionId);
    if (session && ["completed", "failed", "stopped"].includes(session.state)) {
      return session;
    }
    await new Promise((resolve) => setTimeout(resolve, 50));
  }
  throw new Error(`Timed out waiting for session ${sessionId} to finish`);
}

test("ExecutionManager runs Node.js and Python examples end to end", async () => {
  const tempRoot = fs.mkdtempSync(path.join(os.tmpdir(), "sikuli-execution-"));
  const nodeExample = makeExample(tempRoot, "nodejs", "smoke.mjs", 'console.log("node smoke");\n');
  const pythonExample = makeExample(tempRoot, "python", "smoke.py", 'print("python smoke")\n');
  const examples = new Map([
    [nodeExample.id, nodeExample],
    [pythonExample.id, pythonExample],
  ]);

  const manager = new ExecutionManager({
    startRuntime: async () => ({ ok: true }),
    getExampleById: (exampleId) => examples.get(exampleId) || null,
    onEvent: () => {},
  });

  const firstRun = await manager.runExample(nodeExample.id);
  const firstSession = await waitForFinalSession(manager, firstRun.sessionId);
  assert.equal(firstSession.state, "completed");
  assert.match(firstSession.output, /node smoke/);

  const secondRun = await manager.runExample(pythonExample.id);
  const secondSession = await waitForFinalSession(manager, secondRun.sessionId);
  assert.equal(secondSession.state, "completed");
  assert.match(secondSession.output, /python smoke/);
});

test("ExecutionManager recovers cleanly after a runner crash and allows a later rerun", async () => {
  const tempRoot = fs.mkdtempSync(path.join(os.tmpdir(), "sikuli-crash-"));
  const crashExample = makeExample(tempRoot, "nodejs", "crash.mjs", 'throw new Error("boom");\n');
  const goodExample = makeExample(tempRoot, "nodejs", "good.mjs", 'console.log("recovered");\n');
  const examples = new Map([
    [crashExample.id, crashExample],
    [goodExample.id, goodExample],
  ]);
  const events = [];

  const manager = new ExecutionManager({
    startRuntime: async () => ({ ok: true }),
    getExampleById: (exampleId) => examples.get(exampleId) || null,
    onEvent: (event) => {
      events.push(event);
    },
  });

  const crashedRun = await manager.runExample(crashExample.id);
  const crashedSession = await waitForFinalSession(manager, crashedRun.sessionId);
  assert.equal(crashedSession.state, "failed");
  assert.match(crashedSession.error || "", /boom|Runner exited unexpectedly/);
  assert.equal(manager.getCurrentRun(), null);

  const recoveredRun = await manager.runExample(goodExample.id);
  const recoveredSession = await waitForFinalSession(manager, recoveredRun.sessionId);
  assert.equal(recoveredSession.state, "completed");
  assert.match(recoveredSession.output, /recovered/);
  assert.ok(events.some((event) => event.type === "session:failed"));
  assert.ok(events.some((event) => event.type === "session:completed"));
});
