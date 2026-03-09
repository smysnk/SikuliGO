# Node Package User Flow

> Start with [Node.js Client]({{ '/nodejs-client/' | relative_url }}) if you want the current package guide. This page remains the deeper design and packaging note behind that guide.

This document defines the target user story for a Node.js-first integration where automation can run after `npm install` with minimal code.

## User Story

As a Node.js user, I want to install sikuli-go from npm and run desktop automation with a few lines of code, without manually managing gRPC server startup.

## Target Developer Experience

Install:

```bash
npm install @sikuligo/sikuli-go
```

Use:

```ts
import { Sikuli } from "@sikuligo/sikuli-go";

const bot = await Sikuli.launch();
await bot.click({ x: 300, y: 220 });
await bot.typeText("hello");
await bot.hotkey(["cmd", "enter"]);
await bot.close();
```

Implementation status (baseline in repo):

- SDK entrypoint exported from `packages/client-node/src/index.ts`.
- `Sikuli.launch()` and `Sikuli.connect()` implemented in `packages/client-node/src/sikuli.ts`.
- local process launcher implemented in `packages/client-node/src/launcher.ts`.
- binary resolution implemented in `packages/client-node/src/binary.ts`.
- diagnostics command available via `sikuli-go doctor` (`packages/client-node/src/doctor.ts`).

## Required Components

1. `@sikuligo/sikuli-go` npm package (SDK/meta package):
- high-level API (`launch`, `find`, `click`, `typeText`, `hotkey`, app control methods).
- process manager that starts/stops `sikuli-go`.
- gRPC client wrapper with deadlines, auth metadata, and error mapping.

2. `sikuli-go` binary:
- packaged per OS/arch.
- spawned as a child process by SDK `launch()`.
- bound to localhost with ephemeral port and startup auth token.

3. client/server contract:
- SDK and binary both pinned to `proto/sikuli/v1/sikuli.proto` compatibility.
- clear version compatibility policy between npm SDK and binary builds.

## Binary Packaging Strategy

Recommended packaging model:

1. Publish one JS meta package:
- `@sikuligo/sikuli-go`

2. Publish per-platform binary packages as required dependencies:
- `@sikuligo/bin-darwin-arm64`
- `@sikuligo/bin-darwin-x64`
- `@sikuligo/bin-linux-x64`
- `@sikuligo/bin-win32-x64`

Repository scaffolding:
- package manifests under `packages/client-node/packages/bin-*/package.json`
- build script: `scripts/clients/build-node-binaries.sh`
- release script: `scripts/clients/release-node-binaries.sh`

3. Each binary package:
- includes `sikuli-go` and `sikuli-go-monitor`.
- installs to predictable path resolved by `@sikuligo/sikuli-go` at runtime.

4. Runtime resolution in `@sikuligo/sikuli-go`:
- see Binary Resolution details in `docs/strategy/client-strategy.md` (Node.js section).

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
- `doctor` command (`npx @sikuligo/sikuli-go doctor`) for environment checks.

## Implementation Milestones

1. package split:
- create `@sikuligo/sikuli-go` meta package and `@sikuligo/bin-*` packages.

2. launch manager:
- implement child-process lifecycle management and cleanup.

3. compatibility and release:
- add SDK↔binary compatibility checks and release workflow gates.

4. onboarding:
- publish quickstart examples and troubleshooting docs.
