#!/usr/bin/env node

import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import { spawnSync } from "node:child_process";
import readline from "node:readline/promises";
import { resolveSikuliBinary } from "./binary";
import { runInitExamples } from "./init-examples";
import { launchSikuli, stopSpawnedProcess } from "./launcher";

function usage(): string {
  return [
    "Usage: sikuligo <command> [options]",
    "",
    "Wrapper Commands:",
    "  init:js-examples [--dir <targetDir>]         Scaffold a JS project and copy .mjs examples into ./examples",
    "  install-binary [--dir <binDir>]              Copy sikuli runtimes to a PATH-ready directory",
    "",
    "All other commands/flags are forwarded to the native sikuligo binary."
  ].join("\n");
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

function runBinaryProbe(binary: string): { ok: boolean; detail: string } {
  const out = spawnSync(binary, ["-h"], {
    env: process.env,
    encoding: "utf8"
  });
  if (out.error) {
    return { ok: false, detail: out.error.message };
  }
  if (typeof out.status === "number" && out.status !== 0) {
    return { ok: false, detail: `exit=${out.status}` };
  }
  return { ok: true, detail: "ok" };
}

async function runDoctor(): Promise<number> {
  const report: Array<{ check: string; ok: boolean; detail: string }> = [];
  let tmpDir = "";
  try {
    const binary = resolveSikuliBinary();
    report.push({ check: "binary.resolve", ok: true, detail: binary });

    const probe = runBinaryProbe(binary);
    report.push({ check: "binary.execute", ok: probe.ok, detail: probe.detail });
    if (!probe.ok) {
      throw new Error(`binary execution probe failed: ${probe.detail}`);
    }

    tmpDir = fs.mkdtempSync(path.join(os.tmpdir(), "sikuligo-doctor-"));
    const sqlitePath = path.join(tmpDir, "sikuligo-doctor.db");
    const launch = await launchSikuli({
      binaryPath: binary,
      startupTimeoutMs: 3_000,
      adminListen: "",
      sqlitePath,
      spawnServer: true,
      timeoutMs: 3_000
    });
    report.push({ check: "grpc.launch", ok: true, detail: `address=${launch.address}` });
    launch.client.close();
    await stopSpawnedProcess(launch.child);
    report.push({ check: "grpc.shutdown", ok: true, detail: "ok" });

    console.log("sikuligo doctor: ok");
    console.log(`platform: ${process.platform}/${process.arch}`);
    for (const row of report) {
      console.log(`- ${row.check}: ok (${row.detail})`);
    }
    return 0;
  } catch (err) {
    const message = err instanceof Error ? err.message : String(err);
    console.error("sikuligo doctor: failed");
    for (const row of report) {
      const status = row.ok ? "ok" : "failed";
      console.error(`- ${row.check}: ${status} (${row.detail})`);
    }
    console.error(`- error: ${message}`);
    return 1;
  } finally {
    if (tmpDir) {
      try {
        fs.rmSync(tmpDir, { recursive: true, force: true });
      } catch {
        // Ignore cleanup errors.
      }
    }
  }
}

function runBinaryPassthrough(argv: string[]): number {
  const binary = resolveSikuliBinary();
  const out = spawnSync(binary, argv, {
    stdio: "inherit",
    env: process.env
  });
  if (out.error) {
    throw out.error;
  }
  if (typeof out.status === "number") {
    return out.status;
  }
  return 1;
}

async function main(): Promise<number> {
  const [command = "", ...rest] = process.argv.slice(2);
  switch (command) {
    case "init:js-examples": {
      await runInitExamplesScaffold(rest, {
        jsMode: true
      });
      return 0;
    }
    case "doctor": {
      return await runDoctor();
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
    case "--help":
    case "-h":
    case "help": {
      console.log(usage());
      return 0;
    }
    default: {
      return runBinaryPassthrough(process.argv.slice(2));
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
