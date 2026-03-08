import fs from "node:fs";
import path from "node:path";
import { spawnSync } from "node:child_process";
import { createHash } from "node:crypto";
import { resolveSikuliBinary } from "@sikuligo/sikuli-go";

function resolveOnPath(binaryName) {
  const cmd = process.platform === "win32" ? "where" : "which";
  const out = spawnSync(cmd, [binaryName], { encoding: "utf8" });
  if (out.status !== 0) {
    return "";
  }
  const first = String(out.stdout || "")
    .split(/\r?\n/)
    .map((line) => line.trim())
    .find((line) => line.length > 0);
  return first || "";
}

function ensureLocalInstallDirOnPath(installDir) {
  const current = process.env.PATH || "";
  const entries = current.split(path.delimiter).filter(Boolean);
  if (!entries.includes(installDir)) {
    process.env.PATH = `${installDir}${path.delimiter}${current}`;
  }
}

export function ensureSikuliGoOnPath() {
  const binaryName = process.platform === "win32" ? "sikuli-go.exe" : "sikuli-go";
  const existing = resolveOnPath(binaryName);
  if (existing) {
    process.env.SIKULI_GO_BINARY_PATH = existing;
    return existing;
  }

  const sourceBinary = resolveSikuliBinary();
  const installDir = path.resolve(process.cwd(), ".sikuli-go", "bin");
  const targetBinary = path.join(installDir, binaryName);
  fs.mkdirSync(installDir, { recursive: true });

  let shouldCopy = true;
  if (fs.existsSync(targetBinary)) {
    shouldCopy = fileDigest(sourceBinary) !== fileDigest(targetBinary);
  }
  if (shouldCopy) {
    fs.copyFileSync(sourceBinary, targetBinary);
    if (process.platform !== "win32") {
      fs.chmodSync(targetBinary, 0o755);
    }
  }

  ensureLocalInstallDirOnPath(installDir);
  process.env.SIKULI_GO_BINARY_PATH = targetBinary;
  return targetBinary;
}

function fileDigest(filePath) {
  const hash = createHash("sha256");
  hash.update(fs.readFileSync(filePath));
  return hash.digest("hex");
}
