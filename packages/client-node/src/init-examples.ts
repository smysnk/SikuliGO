#!/usr/bin/env node

import fs from "fs";
import path from "path";

function usageAndExit(message?: string): never {
  if (message) {
    console.error(message);
  }
  console.error("Usage: sikuligo-init-examples [--dir <targetDir>] [--force]");
  process.exit(1);
}

function parseArgs(argv: string[]): { targetDir: string; force: boolean } {
  let targetDir = process.cwd();
  let force = false;

  for (let i = 0; i < argv.length; i += 1) {
    const arg = argv[i];
    if (arg === "--force") {
      force = true;
      continue;
    }
    if (arg === "--dir") {
      const value = argv[i + 1];
      if (!value) {
        usageAndExit("Missing value for --dir");
      }
      targetDir = path.resolve(value);
      i += 1;
      continue;
    }
    usageAndExit(`Unknown argument: ${arg}`);
  }

  return { targetDir, force };
}

function copyDirRecursive(sourceDir: string, targetDir: string): void {
  fs.mkdirSync(targetDir, { recursive: true });
  for (const entry of fs.readdirSync(sourceDir, { withFileTypes: true })) {
    const src = path.join(sourceDir, entry.name);
    const dst = path.join(targetDir, entry.name);
    if (entry.isDirectory()) {
      copyDirRecursive(src, dst);
      continue;
    }
    fs.copyFileSync(src, dst);
  }
}

function main(): void {
  const { targetDir, force } = parseArgs(process.argv.slice(2));
  const packageRoot = path.resolve(__dirname, "..", "..");
  const packagedExamplesDir = path.join(packageRoot, "examples");

  if (!fs.existsSync(packagedExamplesDir)) {
    throw new Error(`Packaged examples directory not found: ${packagedExamplesDir}`);
  }

  const outputDir = path.join(targetDir, "examples");
  if (fs.existsSync(outputDir)) {
    if (!force) {
      throw new Error(`Target already exists: ${outputDir} (use --force to overwrite)`);
    }
    fs.rmSync(outputDir, { recursive: true, force: true });
  }

  copyDirRecursive(packagedExamplesDir, outputDir);
  console.log(`Initialized SikuliGO examples in: ${outputDir}`);
}

try {
  main();
} catch (err) {
  const msg = err instanceof Error ? err.message : String(err);
  console.error(msg);
  process.exit(1);
}
