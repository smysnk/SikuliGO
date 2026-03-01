#!/usr/bin/env node

import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { pathToFileURL } from "node:url";
import { spawnSync } from "node:child_process";
import readline from "node:readline/promises";
import { resolveSikuliBinary } from "./binary";
import { runInitExamples } from "./init-examples";

const EXAMPLE_NAMES = new Set([
  "workflow-auto-launch",
  "workflow-connect",
  "find",
  "click",
  "ocr",
  "input",
  "app",
  "user-flow"
]);

function usage(): string {
  return [
    "Usage: sikuligo <command> [options]",
    "",
    "Commands:",
    "  init-examples [--dir <targetDir>]            Scaffold a project and copy examples into ./examples",
    "  init:js-examples [--dir <targetDir>]         Scaffold a JS project and copy .js examples into ./examples",
    "  example <name>                               Run a packaged example by name",
    "  doctor                                       Print binary/platform diagnostics",
    "  install-binary [--dir <binDir>]              Copy sikuli runtimes to a PATH-ready directory",
    "",
    `Example names: ${Array.from(EXAMPLE_NAMES).join(", ")}`
  ].join("\n");
}

function normalizeExampleName(raw: string): string {
  const normalized = raw.trim().toLowerCase().replace(/\.mjs$/, "");
  if (!EXAMPLE_NAMES.has(normalized)) {
    throw new Error(`Unknown example "${raw}". Valid names: ${Array.from(EXAMPLE_NAMES).join(", ")}`);
  }
  return normalized;
}

async function runPackagedExample(nameArg: string): Promise<void> {
  const name = normalizeExampleName(nameArg);
  const packageRoot = path.resolve(__dirname, "..", "..");
  const examplePath = path.join(packageRoot, "examples", `${name}.mjs`);
  if (!fs.existsSync(examplePath)) {
    throw new Error(`Packaged example not found: ${examplePath}`);
  }
  await import(pathToFileURL(examplePath).href);
}

type InitExamplesArgs = {
  targetDir?: string;
  skipInstall: boolean;
};

function parseInitExamplesArgs(argv: string[]): InitExamplesArgs {
  let targetDir: string | undefined;
  let skipInstall = false;
  for (let i = 0; i < argv.length; i += 1) {
    const arg = argv[i];
    if (arg === "--dir") {
      const value = argv[i + 1];
      if (!value) {
        throw new Error("Missing value for --dir");
      }
      targetDir = path.resolve(value);
      i += 1;
      continue;
    }
    if (arg === "--skip-install") {
      skipInstall = true;
      continue;
    }
    throw new Error(`Unknown argument: ${arg}`);
  }
  return { targetDir, skipInstall };
}

async function promptProjectDir(): Promise<string> {
  const defaultDir = "sikuligo-demo";
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
  });
  try {
    const answer = (await rl.question(`Project directory name [${defaultDir}]: `)).trim();
    if (!answer) {
      return path.resolve(defaultDir);
    }
    return path.resolve(answer);
  } finally {
    rl.close();
  }
}

function packageVersion(): string {
  const packageRoot = path.resolve(__dirname, "..", "..");
  const packageJsonPath = path.join(packageRoot, "package.json");
  const raw = JSON.parse(fs.readFileSync(packageJsonPath, "utf8")) as { version?: string };
  const version = String(raw.version || "").trim();
  if (!version) {
    throw new Error(`Unable to determine package version from ${packageJsonPath}`);
  }
  return version;
}

function ensureProjectPackageJson(projectDir: string, opts: { typeModule?: boolean } = {}): void {
  const pkgPath = path.join(projectDir, "package.json");
  const version = packageVersion();
  const dependencyVersion = `^${version}`;
  let pkg: Record<string, unknown>;

  if (fs.existsSync(pkgPath)) {
    pkg = JSON.parse(fs.readFileSync(pkgPath, "utf8")) as Record<string, unknown>;
  } else {
    pkg = {
      name: path.basename(projectDir) || "sikuligo-project",
      private: true
    };
  }

  const dependencies = (pkg.dependencies as Record<string, string> | undefined) ?? {};
  dependencies["@sikuligo/sikuligo"] = dependencyVersion;
  pkg.dependencies = dependencies;
  if (opts.typeModule) {
    pkg.type = "module";
  }

  fs.writeFileSync(pkgPath, `${JSON.stringify(pkg, null, 2)}\n`);
}

function runYarnInstall(projectDir: string): void {
  const out = spawnSync("yarn", ["install"], {
    cwd: projectDir,
    stdio: "inherit",
    env: process.env
  });
  if (out.status !== 0) {
    throw new Error("yarn install failed");
  }
}

function createJsExampleVariants(projectDir: string): void {
  const examplesDir = path.join(projectDir, "examples");
  if (!fs.existsSync(examplesDir)) {
    return;
  }
  const entries = fs.readdirSync(examplesDir, { withFileTypes: true });
  for (const entry of entries) {
    if (!entry.isFile() || !entry.name.endsWith(".mjs")) {
      continue;
    }
    const source = path.join(examplesDir, entry.name);
    const target = path.join(examplesDir, `${entry.name.slice(0, -4)}.js`);
    fs.copyFileSync(source, target);
  }
}

async function runInitExamplesScaffold(
  argv: string[],
  opts: { jsMode?: boolean; defaultSkipInstall?: boolean } = {}
): Promise<void> {
  const args = parseInitExamplesArgs(argv);
  const skipInstall = args.skipInstall || opts.defaultSkipInstall === true;
  const projectDir = args.targetDir ?? (await promptProjectDir());

  fs.mkdirSync(projectDir, { recursive: true });
  ensureProjectPackageJson(projectDir, { typeModule: opts.jsMode === true });
  if (!skipInstall) {
    runYarnInstall(projectDir);
  }
  runInitExamples(["--dir", projectDir, "--force"]);
  if (opts.jsMode === true) {
    createJsExampleVariants(projectDir);
  }
  console.log(`Initialized SikuliGO project in: ${projectDir}`);
  console.log(`Examples copied to: ${path.join(projectDir, "examples")}`);
}

function parseInstallBinaryArgs(argv: string[]): { targetDir: string; yes: boolean; noShellUpdate: boolean } {
  let targetDir = path.join(os.homedir(), ".local", "bin");
  let yes = false;
  let noShellUpdate = false;
  for (let i = 0; i < argv.length; i += 1) {
    const arg = argv[i];
    if (arg === "--dir") {
      const value = argv[i + 1];
      if (!value) {
        throw new Error("Missing value for --dir");
      }
      targetDir = path.resolve(value);
      i += 1;
      continue;
    }
    if (arg === "--yes") {
      yes = true;
      continue;
    }
    if (arg === "--no-shell-update") {
      noShellUpdate = true;
      continue;
    }
    throw new Error(`Unknown argument: ${arg}`);
  }
  return { targetDir, yes, noShellUpdate };
}

function discoverRuntimeSources(primary: string): string[] {
  const out = new Set<string>();
  out.add(primary);
  const dir = path.dirname(primary);
  try {
    for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
      if (!entry.isFile()) {
        continue;
      }
      if (!/^sikuli.*(\.exe)?$/i.test(entry.name)) {
        continue;
      }
      out.add(path.join(dir, entry.name));
    }
  } catch {
    // Ignore sibling scan errors.
  }
  return Array.from(out).filter((candidate) => {
    try {
      fs.accessSync(candidate, fs.constants.F_OK);
      return true;
    } catch {
      return false;
    }
  });
}

function shellProfilePath(): { profile: string; sourceCmd: string } | undefined {
  const shell = process.env.SHELL ?? "";
  if (shell.includes("zsh")) {
    return { profile: path.join(os.homedir(), ".zshrc"), sourceCmd: "source ~/.zshrc" };
  }
  if (shell.includes("bash")) {
    return { profile: path.join(os.homedir(), ".bash_profile"), sourceCmd: "source ~/.bash_profile" };
  }
  return undefined;
}

async function promptYesNo(question: string): Promise<boolean> {
  if (!process.stdin.isTTY) {
    return false;
  }
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
  });
  try {
    const answer = (await rl.question(`${question} [y/N]: `)).trim().toLowerCase();
    return answer === "y" || answer === "yes";
  } finally {
    rl.close();
  }
}

function ensurePathExport(profilePath: string, binDir: string): boolean {
  const marker = "# Added by sikuligo install-binary";
  const exportLine = `export PATH="${binDir}:$PATH"`;
  const snippet = `${marker}\n${exportLine}\n`;
  const existing = fs.existsSync(profilePath) ? fs.readFileSync(profilePath, "utf8") : "";
  if (existing.includes(exportLine)) {
    return false;
  }
  const prefix = existing.length > 0 && !existing.endsWith("\n") ? "\n" : "";
  fs.writeFileSync(profilePath, `${existing}${prefix}${snippet}`);
  return true;
}

async function installBinary(argv: string[]): Promise<{ copied: string[]; targetDir: string; sourceCmd?: string }> {
  const { targetDir, yes, noShellUpdate } = parseInstallBinaryArgs(argv);
  const source = resolveSikuliBinary();
  fs.mkdirSync(targetDir, { recursive: true });
  const copied: string[] = [];
  for (const runtime of discoverRuntimeSources(source)) {
    const runtimeBase = path.basename(runtime);
    const targets = new Set<string>([runtimeBase]);
    if (/^sikuligrpc(\.exe)?$/i.test(runtimeBase)) {
      targets.add(runtimeBase.replace(/sikuligrpc/i, "sikuligo"));
    }
    for (const targetBase of targets) {
      const target = path.join(targetDir, targetBase);
      fs.copyFileSync(runtime, target);
      if (process.platform !== "win32") {
        fs.chmodSync(target, 0o755);
      }
      copied.push(target);
    }
  }

  let sourceCmd: string | undefined;
  if (!noShellUpdate) {
    const shell = shellProfilePath();
    if (shell) {
      const shouldUpdate = yes ? true : await promptYesNo(`Add ${targetDir} to PATH in ${shell.profile}?`);
      if (shouldUpdate) {
        ensurePathExport(shell.profile, targetDir);
        sourceCmd = shell.sourceCmd;
      }
    }
  }
  return { copied, targetDir, sourceCmd };
}

function runDoctor(): number {
  try {
    const binary = resolveSikuliBinary();
    console.log("sikuligo doctor: ok");
    console.log(`binary: ${binary}`);
    console.log(`platform: ${process.platform}/${process.arch}`);
    return 0;
  } catch (err) {
    const message = err instanceof Error ? err.message : String(err);
    console.error("sikuligo doctor: failed");
    console.error(message);
    return 1;
  }
}

async function main(): Promise<number> {
  const [command = "", ...rest] = process.argv.slice(2);
  switch (command) {
    case "init-examples": {
      await runInitExamplesScaffold(rest);
      return 0;
    }
    case "init:js-examples": {
      await runInitExamplesScaffold(rest, {
        jsMode: true
      });
      return 0;
    }
    case "example": {
      const [name = ""] = rest;
      if (!name) {
        throw new Error(`Missing example name\n${usage()}`);
      }
      await runPackagedExample(name);
      return 0;
    }
    case "doctor": {
      return runDoctor();
    }
    case "install-binary": {
      const result = await installBinary(rest);
      for (const copied of result.copied) {
        console.log(copied);
      }
      if (result.sourceCmd) {
        console.log(`Run ${result.sourceCmd} to reload PATH in this shell.`);
      } else {
        console.log(`Ensure ${result.targetDir} is on PATH for new shells.`);
      }
      return 0;
    }
    case "help":
    case "--help":
    case "-h":
    case "": {
      console.log(usage());
      return command === "" ? 1 : 0;
    }
    default: {
      throw new Error(`Unknown command: ${command}\n${usage()}`);
    }
  }
}

main()
  .then((code) => {
    process.exitCode = code;
  })
  .catch((err) => {
    const msg = err instanceof Error ? err.message : String(err);
    console.error(msg);
    process.exit(1);
  });
