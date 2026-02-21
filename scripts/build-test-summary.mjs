#!/usr/bin/env node

import fs from "node:fs";
import path from "node:path";
import readline from "node:readline";

const ROOT = process.cwd();
const RESULTS_DIR = path.join(ROOT, ".test-results");
const INPUT_PATH = path.join(RESULTS_DIR, "go-test.json");
const OUTPUT_PATH = path.join(RESULTS_DIR, "summary.json");

fs.mkdirSync(RESULTS_DIR, { recursive: true });

if (!fs.existsSync(INPUT_PATH)) {
  const emptySummary = {
    summary: {
      packageCount: 0,
      failedPackages: 0,
      tests: 0,
      passed: 0,
      failed: 0,
      skipped: 0,
      durationMs: 0,
    },
    packages: [],
  };
  fs.writeFileSync(OUTPUT_PATH, JSON.stringify(emptySummary, null, 2) + "\n", "utf8");
  console.log(`No go-test JSON found at ${path.relative(ROOT, INPUT_PATH)}. Wrote empty summary.`);
  process.exit(0);
}

const packages = new Map();

function ensurePackage(name) {
  if (!packages.has(name)) {
    packages.set(name, {
      package: name,
      tests: 0,
      passed: 0,
      failed: 0,
      skipped: 0,
      durationMs: 0,
      status: "passed",
      seenTests: new Set(),
      maxElapsedSec: 0,
    });
  }
  return packages.get(name);
}

const rl = readline.createInterface({
  input: fs.createReadStream(INPUT_PATH, { encoding: "utf8" }),
  crlfDelay: Infinity,
});

for await (const line of rl) {
  const trimmed = line.trim();
  if (!trimmed) {
    continue;
  }
  let evt;
  try {
    evt = JSON.parse(trimmed);
  } catch {
    continue;
  }

  const pkgName = evt.Package || "(unknown)";
  const pkg = ensurePackage(pkgName);

  const elapsedSec = Number(evt.Elapsed || 0);
  if (elapsedSec > pkg.maxElapsedSec) {
    pkg.maxElapsedSec = elapsedSec;
  }

  if (evt.Test && evt.Action === "run") {
    pkg.seenTests.add(evt.Test);
  }
  if (evt.Test && evt.Action === "pass") {
    pkg.passed += 1;
  }
  if (evt.Test && evt.Action === "fail") {
    pkg.failed += 1;
    pkg.status = "failed";
  }
  if (evt.Test && evt.Action === "skip") {
    pkg.skipped += 1;
  }

  if (!evt.Test && evt.Action === "fail") {
    pkg.status = "failed";
  }
}

const packageRows = [...packages.values()]
  .map((pkg) => {
    pkg.tests = pkg.seenTests.size;
    pkg.durationMs = Math.round(pkg.maxElapsedSec * 1000);
    if (pkg.failed > 0) {
      pkg.status = "failed";
    }
    return {
      package: pkg.package,
      tests: pkg.tests,
      passed: pkg.passed,
      failed: pkg.failed,
      skipped: pkg.skipped,
      durationMs: pkg.durationMs,
      status: pkg.status,
    };
  })
  .sort((a, b) => a.package.localeCompare(b.package));

const summary = {
  packageCount: packageRows.length,
  failedPackages: packageRows.filter((pkg) => pkg.status === "failed").length,
  tests: packageRows.reduce((n, pkg) => n + pkg.tests, 0),
  passed: packageRows.reduce((n, pkg) => n + pkg.passed, 0),
  failed: packageRows.reduce((n, pkg) => n + pkg.failed, 0),
  skipped: packageRows.reduce((n, pkg) => n + pkg.skipped, 0),
  durationMs: packageRows.reduce((n, pkg) => n + pkg.durationMs, 0),
};

const report = {
  summary,
  packages: packageRows,
};

fs.writeFileSync(OUTPUT_PATH, JSON.stringify(report, null, 2) + "\n", "utf8");
console.log(`Wrote test summary to ${path.relative(ROOT, OUTPUT_PATH)}`);
