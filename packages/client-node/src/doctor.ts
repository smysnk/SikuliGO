#!/usr/bin/env node
import { resolveSikuliBinary } from "./binary";

function run(): number {
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

process.exitCode = run();
