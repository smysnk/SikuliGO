# sikuligo (Node.js)

SikuliGO is a GoLang implementation of Sikuli visual automation. This package provides the Node.js SDK for launching `sikuligrpc` locally and executing automation with a small API surface.

## Setup

```bash
cd clients/node
npm install
npm run build
```

## Quickstart User Flow

```ts
import { Screen, Pattern } from "@sikuligo/sikuligo";

const screen = await Screen.start();
try {
  await screen.click(new Pattern("assets/pattern.png").exact());
} finally {
  await screen.close();
}
```

For OpenCV-backed matching, run `sikuligrpc` with `-tags opencv`.

## Run Examples

```bash
cd clients/node
npm run example:find
npm run example:click
npm run example:ocr
npm run example:input
npm run example:app
npm run example:user-flow
npm run doctor
```

## Environment
- `SIKULIGO_BINARY_PATH` (optional explicit path to `sikuligrpc`)
- `SIKULI_GRPC_ADDR` (used by `Sikuli.connect`; default `127.0.0.1:50051`)
- `SIKULI_GRPC_AUTH_TOKEN` (optional; sent as `x-api-key` for spawned/connected sessions)
- `SIKULI_APP_NAME` (optional; used by `examples/app.ts`)
