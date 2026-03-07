import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { spawnSync } from "node:child_process";
import { createHash } from "node:crypto";

const DEFAULT_BINARY_NAME = process.platform === "win32" ? "sikuli-go.exe" : "sikuli-go";
const DEFAULT_MONITOR_BINARY_NAME =
  process.platform === "win32" ? "sikuli-go-monitor.exe" : "sikuli-go-monitor";
const LEGACY_BINARY_NAMES =
  process.platform === "win32"
    ? ["sikuligo.exe", "sikuligrpc.exe"]
    : ["sikuligo", "sikuligrpc"];
const LEGACY_MONITOR_BINARY_NAMES =
  process.platform === "win32"
    ? ["sikuligo-monitor.exe", "sikuligrpc-monitor.exe"]
    : ["sikuligo-monitor", "sikuligrpc-monitor"];
const RUNTIME_NAMES = [
  DEFAULT_BINARY_NAME,
  DEFAULT_MONITOR_BINARY_NAME,
  ...LEGACY_BINARY_NAMES,
  ...LEGACY_MONITOR_BINARY_NAMES
];

const PLATFORM_BINARY_PACKAGES: Record<string, string[]> = {
  "darwin-arm64": ["@sikuligo/bin-darwin-arm64"],
  "darwin-x64": ["@sikuligo/bin-darwin-x64"],
  "linux-x64": ["@sikuligo/bin-linux-x64"],
  "win32-x64": ["@sikuligo/bin-win32-x64"]
};

function isExecutable(candidatePath: string): boolean {
  try {
    if (!candidatePath) {
      return false;
    }
    const stat = fs.statSync(candidatePath);
    if (!stat.isFile()) {
      return false;
    }
    if (process.platform === "win32") {
      fs.accessSync(candidatePath, fs.constants.F_OK);
    } else {
      fs.accessSync(candidatePath, fs.constants.F_OK | fs.constants.X_OK);
    }
    return true;
  } catch {
    return false;
  }
}

function candidateBinaryPaths(rootDir: string): string[] {
  const names = [DEFAULT_BINARY_NAME, ...LEGACY_BINARY_NAMES];
  return [
    ...names.map((name) => path.join(rootDir, name)),
    ...names.map((name) => path.join(rootDir, "bin", name)),
    ...names.map((name) => path.join(rootDir, "dist", name))
  ];
}

function candidateRuntimePaths(rootDir: string): string[] {
  return [
    ...RUNTIME_NAMES.map((name) => path.join(rootDir, name)),
    ...RUNTIME_NAMES.map((name) => path.join(rootDir, "bin", name)),
    ...RUNTIME_NAMES.map((name) => path.join(rootDir, "dist", name))
  ];
}

function isVirtualZipPath(candidatePath: string): boolean {
  if (!candidatePath) {
    return false;
  }
  return candidatePath.includes(".zip/") || candidatePath.includes(".zip\\") || candidatePath.startsWith("zip:");
}

function isRuntimeFile(candidatePath: string): boolean {
  try {
    return fs.statSync(candidatePath).isFile();
  } catch {
    return false;
  }
}

function canonicalRuntimeName(name: string): string | undefined {
  const ext = path.extname(name);
  const stem = name.slice(0, name.length - ext.length).toLowerCase();
  if (/^sikuli(?:-go|go|grpc)(?:-[0-9a-f]{8,})?$/.test(stem)) {
    return DEFAULT_BINARY_NAME;
  }
  if (/^sikuli(?:-go|go|grpc)-monitor(?:-[0-9a-f]{8,})?$/.test(stem)) {
    return DEFAULT_MONITOR_BINARY_NAME;
  }
  return undefined;
}

function runtimeSourceRank(name: string): number {
  const ext = path.extname(name);
  const stem = name.slice(0, name.length - ext.length).toLowerCase();
  switch (stem) {
    case "sikuli-go":
    case "sikuli-go-monitor":
      return 0;
    case "sikuligo":
    case "sikuligo-monitor":
      return 1;
    case "sikuligrpc":
    case "sikuligrpc-monitor":
      return 2;
    default:
      if (stem.startsWith("sikuli-go-monitor-") || stem.startsWith("sikuli-go-")) {
        return 3;
      }
      if (stem.startsWith("sikuligo-monitor-") || stem.startsWith("sikuligo-")) {
        return 4;
      }
      if (stem.startsWith("sikuligrpc-monitor-") || stem.startsWith("sikuligrpc-")) {
        return 5;
      }
      return 6;
  }
}

function resolveMaterializedRuntimeSources(candidatePath: string): Map<string, string> {
  const sources = new Map<string, { rank: number; source: string }>();
  for (const sibling of [candidatePath, ...candidateRuntimePaths(path.dirname(candidatePath))]) {
    if (!isRuntimeFile(sibling)) {
      continue;
    }
    const canonicalName = canonicalRuntimeName(path.basename(sibling));
    if (!canonicalName) {
      continue;
    }
    const rank = runtimeSourceRank(path.basename(sibling));
    const current = sources.get(canonicalName);
    if (!current || rank < current.rank || (rank === current.rank && sibling < current.source)) {
      sources.set(canonicalName, { rank, source: sibling });
    }
  }
  return new Map(Array.from(sources.entries(), ([name, value]) => [name, value.source]));
}

function materializeSpawnableBinary(candidatePath: string): string {
  if (!isVirtualZipPath(candidatePath)) {
    return candidatePath;
  }

  const canonicalBinaryName = canonicalRuntimeName(path.basename(candidatePath)) ?? DEFAULT_BINARY_NAME;
  const key = createHash("sha256").update(candidatePath).digest("hex").slice(0, 16);
  const cacheDir = path.join(os.tmpdir(), "sikuli-go-node-binaries", key);
  fs.mkdirSync(cacheDir, { recursive: true });

  for (const [runtimeName, sourcePath] of resolveMaterializedRuntimeSources(candidatePath)) {
    const outputPath = path.join(cacheDir, runtimeName);
    if (isExecutable(outputPath)) {
      continue;
    }
    const tmpPath = `${outputPath}.tmp-${process.pid}-${Date.now()}`;
    const binaryData = fs.readFileSync(sourcePath);
    try {
      fs.writeFileSync(tmpPath, binaryData, {
        mode: process.platform === "win32" ? 0o666 : 0o755
      });
      if (process.platform !== "win32") {
        fs.chmodSync(tmpPath, 0o755);
      }
      fs.renameSync(tmpPath, outputPath);
    } catch (err) {
      try {
        fs.rmSync(tmpPath, { force: true });
      } catch {
        // Ignore cleanup errors.
      }
      throw err;
    }
  }

  return path.join(cacheDir, canonicalBinaryName);
}

function resolvePackagedBinary(): string | undefined {
  const key = `${process.platform}-${process.arch}`;
  const packages = PLATFORM_BINARY_PACKAGES[key] ?? [];
  for (const pkgName of packages) {
    try {
      const packageJsonPath = require.resolve(`${pkgName}/package.json`);
      const packageRoot = path.dirname(packageJsonPath);
      for (const candidate of candidateBinaryPaths(packageRoot)) {
        if (isExecutable(candidate)) {
          return candidate;
        }
      }
    } catch {
      // Package not installed/resolvable for this platform.
    }
  }
  return undefined;
}

function resolveFromPath(): string | undefined {
  const cmd = process.platform === "win32" ? "where" : "which";
  const result = spawnSync(cmd, [DEFAULT_BINARY_NAME], { encoding: "utf8" });
  if (result.status !== 0) {
    return undefined;
  }
  const first = result.stdout
    .split(/\r?\n/)
    .map((line) => line.trim())
    .find((line) => line.length > 0);
  if (!first) {
    return undefined;
  }
  return isExecutable(first) ? first : undefined;
}

function resolveLocalRepoFallback(): string | undefined {
  const candidates = [
    path.resolve(process.cwd(), DEFAULT_BINARY_NAME),
    path.resolve(process.cwd(), "bin", DEFAULT_BINARY_NAME),
    path.resolve(__dirname, "..", "..", "..", DEFAULT_BINARY_NAME),
    path.resolve(__dirname, "..", "..", "..", "bin", DEFAULT_BINARY_NAME)
  ];
  for (const legacyName of LEGACY_BINARY_NAMES) {
    candidates.push(path.resolve(__dirname, "..", "..", "..", legacyName));
    candidates.push(path.resolve(process.cwd(), legacyName));
  }
  return candidates.find((candidate) => isExecutable(candidate));
}

function errorWithResolutionHelp(detail: string): Error {
  return new Error(
    `${detail}\n` +
      "Install @sikuligo/sikuli-go to auto-resolve the packaged platform binary, " +
      "or set SIKULIGO_BINARY_PATH, or place sikuli-go in PATH."
  );
}

export function resolveSikuliBinary(explicitPath?: string): string {
  const manual = explicitPath || process.env.SIKULIGO_BINARY_PATH || "";
  if (manual) {
    if (!isExecutable(manual)) {
      throw errorWithResolutionHelp(`Configured binary path is not executable: ${manual}`);
    }
    return materializeSpawnableBinary(manual);
  }

  const localFallback = resolveLocalRepoFallback();
  if (localFallback) {
    return materializeSpawnableBinary(localFallback);
  }

  const packagedBinary = resolvePackagedBinary();
  if (packagedBinary) {
    return materializeSpawnableBinary(packagedBinary);
  }

  const pathBinary = resolveFromPath();
  if (pathBinary) {
    return materializeSpawnableBinary(pathBinary);
  }

  throw errorWithResolutionHelp(
    `Unable to resolve sikuli-go binary for platform ${process.platform}/${process.arch}`
  );
}
