import { spawn } from "node:child_process";

export function parseCommandSpec(command) {
  if (Array.isArray(command) && command.length > 0) {
    return {
      command: String(command[0]),
      args: command.slice(1).map((entry) => String(entry)),
    };
  }

  if (typeof command === "string" && command.trim().length > 0) {
    const tokens = tokenizeCommand(command);
    if (tokens.length === 0) {
      throw new Error("Command string was empty after tokenization.");
    }
    return {
      command: tokens[0],
      args: tokens.slice(1),
    };
  }

  throw new Error("Expected suite.command to be a non-empty string or array.");
}

export function createSummary(values = {}) {
  return {
    total: Number.isFinite(values.total) ? values.total : 0,
    passed: Number.isFinite(values.passed) ? values.passed : 0,
    failed: Number.isFinite(values.failed) ? values.failed : 0,
    skipped: Number.isFinite(values.skipped) ? values.skipped : 0,
  };
}

export function createCoverageMetric(covered, total) {
  if (!Number.isFinite(total)) {
    return null;
  }
  const safeTotal = Math.max(0, Number(total));
  const safeCovered = Number.isFinite(covered) ? Math.max(0, Math.min(safeTotal, Number(covered))) : 0;
  const pct = safeTotal === 0 ? 100 : Number(((safeCovered / safeTotal) * 100).toFixed(2));
  return {
    covered: safeCovered,
    total: safeTotal,
    pct,
  };
}

export function resolveSuiteEnv(overrides = {}) {
  const env = { ...process.env };
  for (const [key, value] of Object.entries(overrides || {})) {
    env[key] = String(value);
  }
  return env;
}

export function slugify(value) {
  return String(value || "")
    .trim()
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, "-")
    .replace(/^-+|-+$/g, "") || "item";
}

export function trimForReport(value, maxLength = 4000) {
  const text = String(value || "");
  if (text.length <= maxLength) {
    return text;
  }
  return `${text.slice(0, maxLength - 1)}…`;
}

export function applySuiteClassification(result, suite) {
  const moduleName = normalizeClassificationField(suite?.module);
  const themeName = normalizeClassificationField(suite?.theme);

  const tests = Array.isArray(result?.tests)
    ? result.tests.map((test) => ({
        ...test,
        module: normalizeClassificationField(test?.module) || moduleName || "uncategorized",
        theme: normalizeClassificationField(test?.theme) || themeName || "uncategorized",
        classificationSource: test?.classificationSource || (moduleName || themeName ? "adapter" : "default"),
      }))
    : [];

  const coverage = result?.coverage && typeof result.coverage === "object"
    ? {
        ...result.coverage,
        files: Array.isArray(result.coverage.files)
          ? result.coverage.files.map((file) => ({
              ...file,
              module: normalizeClassificationField(file?.module) || moduleName || null,
              theme: normalizeClassificationField(file?.theme) || themeName || null,
            }))
          : [],
      }
    : result?.coverage || null;

  return {
    ...result,
    tests,
    coverage,
  };
}

export function spawnCommand(command, args, options = {}) {
  return new Promise((resolve) => {
    const startedAt = Date.now();
    const child = spawn(command, args, {
      cwd: options.cwd,
      env: options.env,
      stdio: ["ignore", "pipe", "pipe"],
    });

    let stdout = "";
    let stderr = "";

    child.stdout.on("data", (chunk) => {
      stdout += chunk.toString();
    });

    child.stderr.on("data", (chunk) => {
      stderr += chunk.toString();
    });

    child.on("error", (error) => {
      stderr += `${error instanceof Error ? error.message : String(error)}\n`;
      resolve({
        exitCode: 1,
        stdout,
        stderr,
        durationMs: Date.now() - startedAt,
      });
    });

    child.on("close", (code) => {
      resolve({
        exitCode: Number.isInteger(code) ? code : 1,
        stdout,
        stderr,
        durationMs: Date.now() - startedAt,
      });
    });
  });
}

function tokenizeCommand(command) {
  const tokens = [];
  let current = "";
  let quote = null;

  for (let index = 0; index < command.length; index += 1) {
    const char = command[index];
    if (quote) {
      if (char === quote) {
        quote = null;
      } else if (char === "\\" && quote === '"' && index + 1 < command.length) {
        current += command[index + 1];
        index += 1;
      } else {
        current += char;
      }
      continue;
    }

    if (char === '"' || char === "'") {
      quote = char;
      continue;
    }

    if (/\s/.test(char)) {
      if (current.length > 0) {
        tokens.push(current);
        current = "";
      }
      continue;
    }

    current += char;
  }

  if (current.length > 0) {
    tokens.push(current);
  }

  return tokens;
}

function normalizeClassificationField(value) {
  return typeof value === "string" && value.trim() ? value.trim() : null;
}
