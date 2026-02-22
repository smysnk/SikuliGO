# sikuligo (Node.js)

SikuliGO is a GoLang implementation of Sikuli visual automation. This package provides the Node.js SDK for launching `sikuligrpc` locally and executing automation with a small API surface.

## Prerequisites

- Node.js 20+

## Binary Resolution

`@sikuligo/sikuligo` declares all supported platform binaries (`@sikuligo/bin-*`) and resolves the correct binary automatically for the current OS/CPU at runtime.

Resolution order:

1. `SIKULIGO_BINARY_PATH` (explicit override)
2. auto-resolved packaged platform binary
3. `sikuligrpc` found in `PATH`

## Setup

```bash
cd clients/node
npm install
npm run generate
npm run build
```

## Environment

- `SIKULIGO_BINARY_PATH` (optional explicit path to `sikuligrpc`)
- `SIKULI_GRPC_ADDR` (used by `Sikuli.connect`; default `127.0.0.1:50051`)
- `SIKULI_GRPC_AUTH_TOKEN` (optional; sent as `x-api-key` for spawned/connected sessions)
- `SIKULI_APP_NAME` (optional; used by `examples/app.ts`)

## Quickstart User Flow

```ts
import { Sikuli } from "@sikuligo/sikuligo";

const bot = await Sikuli.launch();
await bot.click({ x: 300, y: 220 });
await bot.typeText("hello");
await bot.hotkey(["cmd", "enter"]);
await bot.close();
```

## Run Examples

```bash
cd clients/node
npm run example:find
npm run example:ocr
npm run example:input
npm run example:app
npm run example:user-flow
npm run doctor
```
