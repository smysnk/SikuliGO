#!/usr/bin/env node

import { spawnSync } from "node:child_process";
import { resolveSikuliBinary } from "./binary";

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

function main(): number {
  const args = process.argv.slice(2);
  if (args.length === 0) {
    return runBinaryPassthrough(["-h"]);
  }
  return runBinaryPassthrough(args);
}

try {
  process.exitCode = main();
} catch (err) {
  const msg = err instanceof Error ? err.message : String(err);
  console.error(msg);
  process.exit(1);
}
