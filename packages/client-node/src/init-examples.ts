#!/usr/bin/env node

import fs from "fs";
import path from "path";

function usage(): string {
  return "Usage: sikuli-go init-examples [--dir <targetDir>]";
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
        throw new Error(`Missing value for --dir\n${usage()}`);
      }
      targetDir = path.resolve(value);
      i += 1;
      continue;
    }
    throw new Error(`Unknown argument: ${arg}\n${usage()}`);
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

export function runInitExamples(argv: string[] = process.argv.slice(2)): string {
  const { targetDir, force } = parseArgs(argv);
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
  return outputDir;
}

if (require.main === module) {
  try {
    const outputDir = runInitExamples();
    console.log(`Initialized SikuliGO examples in: ${outputDir}`);
  } catch (err) {
    const msg = err instanceof Error ? err.message : String(err);
    console.error(msg);
    process.exit(1);
  }
}
