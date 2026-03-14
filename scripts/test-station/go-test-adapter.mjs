import fs from "node:fs";
import os from "node:os";
import path from "node:path";

import { applySuiteClassification, createCoverageMetric, createSummary, parseCommandSpec, resolveSuiteEnv, slugify, spawnCommand, trimForReport } from "./common.mjs";

export default {
  id: "go-test",
  description: "Go test adapter",
  phase: 1,
  async run({ project, suite, execution }) {
    const tempDir = fs.mkdtempSync(path.join(os.tmpdir(), "sikuli-go-test-station-go-"));
    const coverageEnabled = Boolean(execution?.coverage) && suite?.coverage?.enabled !== false;
    const coveragePath = path.join(tempDir, "coverage.out");
    const modulePath = readGoModulePath(suite.cwd || project.rootDir);
    const commandSpec = buildGoTestCommand(parseCommandSpec(suite.command), coverageEnabled ? coveragePath : null);
    const result = await spawnCommand(commandSpec.command, commandSpec.args, {
      cwd: suite.cwd || project.rootDir,
      env: resolveSuiteEnv(suite.env),
    });

    const parsed = parseGoTestJson(result.stdout, {
      packageName: suite.packageName || "default",
      moduleName: suite.module || null,
      themeName: suite.theme || null,
    });
    const coverage = coverageEnabled && fs.existsSync(coveragePath)
      ? parseGoCoverageProfile(coveragePath, {
          modulePath,
          moduleRoot: suite.cwd || project.rootDir,
          packageName: suite.packageName || "default",
          moduleName: suite.module || null,
          themeName: suite.theme || null,
        })
      : null;

    const warnings = [];
    if (coverageEnabled && !coverage) {
      warnings.push("Coverage was requested, but no Go coverage profile was produced.");
    }

    return applySuiteClassification({
      status: result.exitCode === 0 && parsed.summary.failed === 0 ? "passed" : "failed",
      durationMs: result.durationMs,
      summary: parsed.summary.total > 0 ? parsed.summary : createSummary({ total: 1, failed: result.exitCode === 0 ? 0 : 1, passed: result.exitCode === 0 ? 1 : 0 }),
      coverage,
      tests: parsed.tests.length > 0 ? parsed.tests : [
        {
          name: suite.label || suite.id || "go test",
          fullName: suite.label || suite.id || "go test",
          status: result.exitCode === 0 ? "passed" : "failed",
          durationMs: result.durationMs,
          file: null,
          line: null,
          column: null,
          assertions: ["Go test command completed."],
          setup: [],
          mocks: [],
          failureMessages: result.exitCode === 0 ? [] : [trimForReport(result.stderr || result.stdout || "Go test failed.")],
          rawDetails: {},
          module: suite.module || "runtime",
          theme: suite.theme || "go",
          classificationSource: "adapter",
        },
      ],
      warnings,
      output: {
        stdout: result.stdout,
        stderr: result.stderr,
      },
      rawArtifacts: [
        {
          relativePath: `${slugify(suite.packageName || "default")}-${slugify(suite.id || "go-test")}.ndjson`,
          content: result.stdout,
        },
        ...(coverage && fs.existsSync(coveragePath)
          ? [{
              relativePath: `${slugify(suite.packageName || "default")}-${slugify(suite.id || "go-test")}-coverage.out`,
              sourcePath: coveragePath,
            }]
          : []),
      ],
    }, suite);
  },
};

function buildGoTestCommand(commandSpec, coveragePath) {
  const isGoTest = commandSpec.command === "go" && commandSpec.args[0] === "test";
  if (!isGoTest) {
    throw new Error("Go test adapter requires a direct `go test ...` command.");
  }

  const args = ["test", "-json"];
  if (coveragePath) {
    args.push("-covermode=atomic", `-coverprofile=${coveragePath}`);
  }
  args.push(...commandSpec.args.slice(1));

  return {
    command: commandSpec.command,
    args,
  };
}

function parseGoTestJson(output, options = {}) {
  const tests = new Map();
  const packageOutput = new Map();
  const packageFailures = new Set();

  for (const line of String(output || "").split(/\r?\n/)) {
    const trimmed = line.trim();
    if (!trimmed) {
      continue;
    }

    let event;
    try {
      event = JSON.parse(trimmed);
    } catch {
      continue;
    }

    const packageName = event.Package || options.packageName || "default";

    if (event.Action === "output") {
      appendPackageOutput(packageOutput, packageName, String(event.Output || ""));
    }

    if (!event.Test) {
      if (event.Action === "fail") {
        packageFailures.add(packageName);
      }
      continue;
    }

    const key = `${packageName}::${event.Test}`;
    const existing = tests.get(key) || {
      name: event.Test,
      fullName: `${packageName} ${event.Test}`,
      status: "passed",
      durationMs: 0,
      file: null,
      line: null,
      column: null,
      assertions: [],
      setup: [],
      mocks: [],
      failureMessages: [],
      rawDetails: {
        package: packageName,
      },
      module: options.moduleName || "runtime",
      theme: options.themeName || "go",
      classificationSource: "adapter",
      output: [],
    };

    if (event.Action === "output") {
      existing.output.push(String(event.Output || "").replace(/\r?\n$/, ""));
    } else if (event.Action === "skip") {
      existing.status = "skipped";
      existing.durationMs = toDurationMs(event.Elapsed);
    } else if (event.Action === "pass") {
      existing.status = "passed";
      existing.durationMs = toDurationMs(event.Elapsed);
    } else if (event.Action === "fail") {
      existing.status = "failed";
      existing.durationMs = toDurationMs(event.Elapsed);
      const failureText = existing.output.filter(Boolean).join("\n").trim();
      if (failureText) {
        existing.failureMessages = [trimForReport(failureText, 6000)];
      }
    }

    tests.set(key, existing);
  }

  for (const packageName of packageFailures) {
    const hasPackageTests = [...tests.values()].some((entry) => entry.rawDetails?.package === packageName);
    if (hasPackageTests) {
      continue;
    }
    const outputText = (packageOutput.get(packageName) || []).join("").trim();
    tests.set(`${packageName}::__package__`, {
      name: `${packageName} package checks`,
      fullName: `${packageName} package checks`,
      status: "failed",
      durationMs: 0,
      file: null,
      line: null,
      column: null,
      assertions: [],
      setup: [],
      mocks: [],
      failureMessages: outputText ? [trimForReport(outputText, 6000)] : ["Package-level Go test failure."],
      rawDetails: {
        package: packageName,
      },
      module: options.moduleName || "runtime",
      theme: options.themeName || "go",
      classificationSource: "adapter",
    });
  }

  const normalizedTests = [...tests.values()].map((entry) => {
    const { output: ignoredOutput, ...test } = entry;
    return test;
  });

  return {
    tests: normalizedTests,
    summary: createSummary({
      total: normalizedTests.length,
      passed: normalizedTests.filter((entry) => entry.status === "passed").length,
      failed: normalizedTests.filter((entry) => entry.status === "failed").length,
      skipped: normalizedTests.filter((entry) => entry.status === "skipped").length,
    }),
  };
}

function appendPackageOutput(map, packageName, output) {
  if (!map.has(packageName)) {
    map.set(packageName, []);
  }
  map.get(packageName).push(output);
}

function parseGoCoverageProfile(filePath, options = {}) {
  const lines = fs.readFileSync(filePath, "utf8").split(/\r?\n/).map((line) => line.trim()).filter(Boolean);
  const dataLines = lines.filter((line) => !line.startsWith("mode:"));
  if (dataLines.length === 0) {
    return null;
  }

  const fileMap = new Map();

  for (const line of dataLines) {
    const [rawFile, remainder] = splitCoverageLine(line);
    if (!rawFile || !remainder) {
      continue;
    }

    const match = remainder.match(/^(\d+)\.(\d+),(\d+)\.(\d+)\s+(\d+)\s+([0-9.]+)$/);
    if (!match) {
      continue;
    }

    const startLine = Number(match[1]);
    const endLine = Number(match[3]);
    const numStatements = Number(match[5]);
    const hitCount = Number(match[6]);
    const resolvedPath = resolveGoCoverageFilePath(rawFile, options);

    if (!fileMap.has(resolvedPath)) {
      fileMap.set(resolvedPath, {
        statementsTotal: 0,
        statementsCovered: 0,
        linesTotal: new Set(),
        linesCovered: new Set(),
      });
    }

    const entry = fileMap.get(resolvedPath);
    entry.statementsTotal += numStatements;
    if (hitCount > 0) {
      entry.statementsCovered += numStatements;
    }
    for (let lineNumber = startLine; lineNumber <= endLine; lineNumber += 1) {
      entry.linesTotal.add(lineNumber);
      if (hitCount > 0) {
        entry.linesCovered.add(lineNumber);
      }
    }
  }

  const files = [...fileMap.entries()]
    .map(([filePathValue, entry]) => ({
      path: filePathValue,
      lines: createCoverageMetric(entry.linesCovered.size, entry.linesTotal.size),
      statements: createCoverageMetric(entry.statementsCovered, entry.statementsTotal),
      functions: null,
      branches: null,
      packageName: options.packageName || null,
      module: options.moduleName || null,
      theme: options.themeName || null,
    }))
    .sort((left, right) => left.path.localeCompare(right.path));

  return {
    lines: aggregateCoverageMetric(files, "lines"),
    statements: aggregateCoverageMetric(files, "statements"),
    functions: null,
    branches: null,
    files,
  };
}

function splitCoverageLine(line) {
  const index = line.indexOf(":");
  if (index === -1) {
    return [null, null];
  }
  return [line.slice(0, index), line.slice(index + 1)];
}

function resolveGoCoverageFilePath(rawPath, options = {}) {
  const normalized = String(rawPath || "").replace(/\\/g, "/");
  const modulePath = String(options.modulePath || "");
  if (modulePath && normalized.startsWith(`${modulePath}/`)) {
    return path.join(options.moduleRoot || process.cwd(), normalized.slice(modulePath.length + 1));
  }
  if (path.isAbsolute(normalized)) {
    return normalized;
  }
  return path.join(options.moduleRoot || process.cwd(), normalized);
}

function readGoModulePath(moduleRoot) {
  const goModPath = path.join(moduleRoot, "go.mod");
  const content = fs.readFileSync(goModPath, "utf8");
  const match = content.match(/^module\s+(.+)$/m);
  return match ? match[1].trim() : "";
}

function aggregateCoverageMetric(files, key) {
  const entries = files.map((file) => file[key]).filter(Boolean);
  if (entries.length === 0) {
    return null;
  }
  const covered = entries.reduce((sum, metric) => sum + metric.covered, 0);
  const total = entries.reduce((sum, metric) => sum + metric.total, 0);
  return createCoverageMetric(covered, total);
}

function toDurationMs(value) {
  const seconds = Number(value || 0);
  return Number.isFinite(seconds) ? Math.round(seconds * 1000) : 0;
}
