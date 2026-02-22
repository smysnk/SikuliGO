# sikuligo (Node.js)

This package provides a Node.js SDK that can launch a local `sikuligrpc` process and execute automation with a small API surface.

## Prerequisites

- Node.js 20+
- npm
- `sikuligrpc` binary available from one of:
  - this package's `optionalDependencies` (`@sikuligo/bin-*`) after binary packages are published
  - `SIKULIGO_BINARY_PATH`
  - a manually installed platform package (for example `@sikuligo/bin-darwin-arm64`)
  - `sikuligrpc` in `PATH`

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
