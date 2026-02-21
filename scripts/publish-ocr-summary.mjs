#!/usr/bin/env node

import fs from "node:fs";
import path from "node:path";

const ROOT = process.cwd();
const RESULTS_FILE = path.join(ROOT, ".test-results", "go-test.json");
const summaryFile = process.env.GITHUB_STEP_SUMMARY || null;

function writeLine(line = "") {
  if (summaryFile) {
    fs.appendFileSync(summaryFile, `${line}\n`, "utf8");
  } else {
    console.log(line);
  }
}

if (!fs.existsSync(RESULTS_FILE)) {
  writeLine("### OCR Test Summary");
  writeLine("");
  writeLine("- OCR test JSON not found.");
  process.exit(0);
}

const lines = fs.readFileSync(RESULTS_FILE, "utf8").split(/\r?\n/).filter(Boolean);

const tests = new Map();

for (const line of lines) {
  let evt;
  try {
    evt = JSON.parse(line);
  } catch {
    continue;
  }
  const testName = String(evt.Test || "");
  if (!testName || !testName.toLowerCase().includes("ocr")) {
    continue;
  }

  const key = `${evt.Package || "(unknown)"}::${testName}`;
  if (!tests.has(key)) {
    tests.set(key, {
      package: evt.Package || "(unknown)",
      test: testName,
      status: "run",
      elapsedMs: 0,
    });
  }
  const item = tests.get(key);
  if (["pass", "fail", "skip"].includes(evt.Action)) {
    item.status = evt.Action;
  }
  if (typeof evt.Elapsed === "number") {
    item.elapsedMs = Math.max(item.elapsedMs, Math.round(evt.Elapsed * 1000));
  }
}

const rows = [...tests.values()].sort((a, b) =>
  `${a.package}/${a.test}`.localeCompare(`${b.package}/${b.test}`),
);

const passed = rows.filter((r) => r.status === "pass").length;
const failed = rows.filter((r) => r.status === "fail").length;
const skipped = rows.filter((r) => r.status === "skip").length;

writeLine("### OCR Test Summary");
writeLine("");
writeLine(`- Total OCR tests: ${rows.length}`);
writeLine(`- Passed: ${passed}`);
writeLine(`- Failed: ${failed}`);
writeLine(`- Skipped: ${skipped}`);

if (rows.length > 0) {
  writeLine("");
  for (const row of rows) {
    const icon = row.status === "pass" ? "✅" : row.status === "fail" ? "❌" : "⚪";
    writeLine(`- ${icon} \`${row.package}\` :: \`${row.test}\` (${(row.elapsedMs / 1000).toFixed(2)}s)`);
  }
}
