const fs = require("fs");
const os = require("os");
const path = require("path");
const test = require("node:test");
const assert = require("node:assert/strict");

const { createExampleStore } = require("../src/examples");

function makeWorkspaceFixture() {
  const rootDir = fs.mkdtempSync(path.join(os.tmpdir(), "sikuli-examples-"));
  const nodePackageRoot = path.join(rootDir, "client-node");
  const pythonPackageRoot = path.join(rootDir, "client-python");
  const nodeExamplesDir = path.join(nodePackageRoot, "examples");
  const pythonExamplesDir = path.join(pythonPackageRoot, "examples");
  fs.mkdirSync(path.join(nodeExamplesDir, "assets"), { recursive: true });
  fs.mkdirSync(path.join(pythonExamplesDir, "assets"), { recursive: true });
  fs.writeFileSync(path.join(nodeExamplesDir, "find.mjs"), "console.log('node find');\n", "utf8");
  fs.writeFileSync(path.join(nodeExamplesDir, "assets", "pattern.png"), "node-image", "utf8");
  fs.writeFileSync(path.join(pythonExamplesDir, "find.py"), "print('python find')\n", "utf8");
  fs.writeFileSync(path.join(pythonExamplesDir, "assets", "pattern.txt"), "python-asset", "utf8");
  return {
    rootDir,
    store: createExampleStore({
      sources: [
        {
          language: "nodejs",
          label: "Node.js",
          packageRoot: nodePackageRoot,
          examplesDir: nodeExamplesDir,
          extension: ".mjs",
          ignoredFiles: new Set(["bootstrap.mjs"]),
        },
        {
          language: "python",
          label: "Python",
          packageRoot: pythonPackageRoot,
          examplesDir: pythonExamplesDir,
          extension: ".py",
          ignoredFiles: new Set(["bootstrap.py"]),
        },
      ],
    }),
  };
}

test("example store supports create, save, clone, rename, and workspace asset reads", async () => {
  const fixture = makeWorkspaceFixture();
  const { store } = fixture;

  const initialExamples = store.listExamples();
  assert.equal(initialExamples.length, 2);

  const created = store.createExample("nodejs", "new workflow", "console.log('created');\n");
  assert.equal(created.fileName, "new-workflow.mjs");
  assert.match(created.source, /created/);

  const saved = store.saveExample(created.id, "console.log('saved');\n");
  assert.match(saved.source, /saved/);

  const cloned = store.cloneExample(saved.id, "saved-copy.mjs");
  assert.equal(cloned.fileName, "saved-copy.mjs");
  assert.equal(cloned.source, saved.source);

  const renamed = store.renameExample(cloned.id, "renamed-script.mjs", "console.log('renamed');\n");
  assert.equal(renamed.fileName, "renamed-script.mjs");
  assert.match(renamed.source, /renamed/);

  const workspaceFiles = store.listWorkspaceFiles("nodejs:find.mjs");
  assert.ok(workspaceFiles.some((file) => file.relativeWorkspacePath === "assets/pattern.png"));
  assert.ok(workspaceFiles.some((file) => file.relativeWorkspacePath === "renamed-script.mjs"));

  const asset = store.readWorkspaceFile("nodejs:find.mjs", "assets/pattern.png");
  assert.equal(asset.isImage, true);
  assert.match(asset.dataUrl, /^data:image\/png;base64,/);

  const textAsset = store.readWorkspaceFile("python:find.py", "assets/pattern.txt");
  assert.equal(textAsset.isText, true);
  assert.equal(textAsset.content, "python-asset");
});
