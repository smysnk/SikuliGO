import fs from "node:fs";
import path from "node:path";
import { spawnSync } from "node:child_process";

const DEFAULT_BINARY_NAME = process.platform === "win32" ? "sikuligo.exe" : "sikuligo";

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
  const names = process.platform === "win32" ? [DEFAULT_BINARY_NAME, "sikuligrpc.exe"] : [DEFAULT_BINARY_NAME, "sikuligrpc"];
  return [
    ...names.map((name) => path.join(rootDir, name)),
    ...names.map((name) => path.join(rootDir, "bin", name)),
    ...names.map((name) => path.join(rootDir, "dist", name))
  ];
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
    path.resolve(__dirname, "..", "..", "..", process.platform === "win32" ? "sikuligo.exe" : "sikuligo"),
    path.resolve(__dirname, "..", "..", "..", "bin", DEFAULT_BINARY_NAME)
  ];
  if (process.platform === "win32") {
    candidates.push(path.resolve(__dirname, "..", "..", "..", "sikuligrpc.exe"));
    candidates.push(path.resolve(process.cwd(), "sikuligrpc.exe"));
  } else {
    candidates.push(path.resolve(__dirname, "..", "..", "..", "sikuligrpc"));
    candidates.push(path.resolve(process.cwd(), "sikuligrpc"));
  }
  return candidates.find((candidate) => isExecutable(candidate));
}

function errorWithResolutionHelp(detail: string): Error {
  return new Error(
    `${detail}\n` +
      "Install @sikuligo/sikuligo to auto-resolve the packaged platform binary, " +
      "or set SIKULIGO_BINARY_PATH, or place sikuligo in PATH."
  );
}

export function resolveSikuliBinary(explicitPath?: string): string {
  const manual = explicitPath || process.env.SIKULIGO_BINARY_PATH || "";
  if (manual) {
    if (!isExecutable(manual)) {
      throw errorWithResolutionHelp(`Configured binary path is not executable: ${manual}`);
    }
    return manual;
  }

  const localFallback = resolveLocalRepoFallback();
  if (localFallback) {
    return localFallback;
  }

  const packagedBinary = resolvePackagedBinary();
  if (packagedBinary) {
    return packagedBinary;
  }

  const pathBinary = resolveFromPath();
  if (pathBinary) {
    return pathBinary;
  }

  throw errorWithResolutionHelp(
    `Unable to resolve sikuligo binary for platform ${process.platform}/${process.arch}`
  );
}
