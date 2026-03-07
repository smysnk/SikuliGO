#!/usr/bin/env node
import { spawnSync } from "node:child_process";
import { resolveSikuliBinary } from "./binary";

function findOnPath(cmd: string): string | undefined {
  const probe = process.platform === "win32" ? "where" : "which";
  const result = spawnSync(probe, [cmd], { encoding: "utf8" });
  if (result.status !== 0) {
    return undefined;
  }
  return result.stdout
    .split(/\r?\n/)
    .map((line) => line.trim())
    .find((line) => line.length > 0);
}

function run(): number {
  try {
    const binary = resolveSikuliBinary();
    console.log("sikuli-go doctor: ok");
    console.log(`binary: ${binary}`);
    console.log(`platform: ${process.platform}/${process.arch}`);
    if (process.platform === "darwin") {
      const cliclick = findOnPath("cliclick");
      if (!cliclick) {
        console.error("runtime deps: missing cliclick (required for click/move on macOS)");
        console.error('Install with: brew install cliclick');
        return 1;
      }
      console.log(`runtime deps: cliclick=${cliclick}`);
    }
    return 0;
  } catch (err) {
    const message = err instanceof Error ? err.message : String(err);
    console.error("sikuli-go doctor: failed");
    console.error(message);
    return 1;
  }
}

process.exitCode = run();
