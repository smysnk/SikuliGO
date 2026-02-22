# Node Package User Flow

This document defines the target user story for a Node.js-first integration where automation can run after `npm install` with minimal code.

## User Story

As a Node.js user, I want to install SikuliGO from npm and run desktop automation with a few lines of code, without manually managing gRPC server startup.

## Target Developer Experience

Install:

```bash
npm install sikuligo
```

Use:

```ts
import { Sikuli } from "sikuligo";

const bot = await Sikuli.launch();
await bot.click({ x: 300, y: 220 });
await bot.typeText("hello");
await bot.hotkey(["cmd", "enter"]);
await bot.close();
```

## Required Components

1. `sikuligo` npm package (SDK/meta package):
- high-level API (`launch`, `find`, `click`, `typeText`, `hotkey`, app control methods).
- process manager that starts/stops `sikuligrpc`.
- gRPC client wrapper with deadlines, auth metadata, and error mapping.

2. `sikuligrpc` binary:
- packaged per OS/arch.
- spawned as a child process by SDK `launch()`.
- bound to localhost with ephemeral port and startup auth token.

3. client/server contract:
- SDK and binary both pinned to `proto/sikuli/v1/sikuli.proto` compatibility.
- clear version compatibility policy between npm SDK and binary builds.

## Binary Packaging Strategy

Recommended packaging model:

1. Publish one JS meta package:
- `sikuligo`

2. Publish per-platform binary packages as optional dependencies:
- `@sikuligo/bin-darwin-arm64`
- `@sikuligo/bin-darwin-x64`
- `@sikuligo/bin-linux-x64`
- `@sikuligo/bin-win32-x64`

3. Each binary package:
- includes one `sikuligrpc` executable.
- uses `os` and `cpu` constraints in `package.json`.
- installs to predictable path resolved by `sikuligo` at runtime.

4. Runtime resolution in `sikuligo`:
- detect active platform/arch.
- resolve installed binary package.
- return explicit install error if package missing/incompatible.

## Release and Build Requirements

- build binaries in CI for each supported OS/arch.
- produce and store checksums for release artifacts.
- sign/notarize macOS binaries before publish.
- preserve executable bit in packed artifact.
- version binaries alongside SDK compatibility matrix.

## Runtime Requirements

- desktop session available (not true headless automation).
- OS input/accessibility permissions granted (especially macOS).
- OCR runtime dependencies installed when OCR APIs are used.

## Operational Requirements

- startup health check after `launch()` before returning control.
- structured startup errors for:
  - missing binary
  - permission denial
  - unsupported platform
  - startup timeout
- trace/auth propagation into gRPC calls by default.
- `doctor` command (`npx sikuligo doctor`) for environment checks.

## Implementation Milestones

1. package split:
- create `sikuligo` meta package and `@sikuligo/bin-*` packages.

2. launch manager:
- implement child-process lifecycle management and cleanup.

3. compatibility and release:
- add SDK↔binary compatibility checks and release workflow gates.

4. onboarding:
- publish quickstart examples and troubleshooting docs.
