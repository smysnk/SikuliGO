const fs = require("fs");
const os = require("os");
const path = require("path");
const { spawnSync } = require("child_process");
const test = require("node:test");
const assert = require("node:assert/strict");

const { instrumentNodeSource } = require("../src/execution/node-transform");

test("Node instrumentation only wraps supported top-level statements", async () => {
  const source = [
    'import { Screen } from "@sikuligo/sikuli-go";',
    "const value = 1;",
    "function helper() {",
    '  console.log("inside helper");',
    "}",
    "await Screen();",
    "",
  ].join("\n");

  const transformed = instrumentNodeSource(source).source;
  assert.match(transformed, /await globalThis\.__sikuliStep\(\{"line":2/);
  assert.match(transformed, /await globalThis\.__sikuliStep\(\{"line":6/);
  assert.doesNotMatch(transformed, /line":1/);
  assert.doesNotMatch(transformed, /function helper\(\) \{\n\s+await globalThis\.__sikuliStep/);
});

test("Python instrumentation keeps future imports and injects step hooks into function bodies", async () => {
  const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "sikuli-python-transform-"));
  const inputPath = path.join(tempDir, "input.py");
  const outputPath = path.join(tempDir, "output.py");
  fs.writeFileSync(
    inputPath,
    [
      "from __future__ import annotations",
      "",
      "VALUE = 1",
      "",
      "def main() -> int:",
      "    print(VALUE)",
      "    return 0",
      "",
    ].join("\n"),
    "utf8",
  );

  const result = spawnSync("python3", [
    path.join(__dirname, "..", "src", "execution", "python-transform.py"),
    inputPath,
    outputPath,
  ], {
    cwd: path.join(__dirname, ".."),
    encoding: "utf8",
  });

  assert.equal(result.status, 0, result.stderr || result.stdout);
  const transformed = fs.readFileSync(outputPath, "utf8");
  assert.match(transformed, /from __future__ import annotations/);
  assert.match(transformed, /__sikuli_step\(3, 1, 'python-3-1'\)/);
  assert.match(transformed, /__sikuli_step\(6, 5, 'python-6-2'\)/);
});
