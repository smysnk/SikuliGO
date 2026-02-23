# sikuligo (Node.js)

SikuliGO is a GoLang implementation of Sikuli visual automation. This package provides the Node.js SDK for launching `sikuligo` locally and executing automation with a small API surface.

## Setup

```bash
cd clients/node
npm install
npm run build
```

## Quickstart

### 1) Start API manually, then run the client script

Start `sikuligo` yourself first:

```bash
# from repo root
./sikuligo -listen 127.0.0.1:50051
```

Then run:

```bash
cd clients/node
npm run example:workflow:connect
```

`npm run example:workflow:connect` runs:

```js
import { Screen, Pattern } from "../src";

async function main() {
  const screen = await Screen.connect({
    address: process.env.SIKULI_GRPC_ADDR ?? "127.0.0.1:50051"
  });
  try {
    const match = await screen.click(new Pattern("assets/pattern.png").exact());
    console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
  } finally {
    await screen.close();
  }
}
```

### 2) Run script only (auto-launch API)

```bash
cd clients/node
npm run example:workflow:auto
```

`npm run example:workflow:auto` runs:

```js
import { Screen, Pattern } from "../src";

async function main() {
  const screen = await Screen.start();
  try {
    const match = await screen.click(new Pattern("assets/pattern.png").exact());
    console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
  } finally {
    await screen.close();
  }
}
```

## Run Examples

```bash
cd clients/node
npm run example:workflow:auto
npm run example:workflow:connect
npm run example:find
npm run example:click
npm run example:ocr
npm run example:input
npm run example:app
npm run example:user-flow
npm run doctor
```

## Environment
- `SIKULIGO_BINARY_PATH` (optional explicit path to `sikuligo`)
- `SIKULI_GRPC_ADDR` (used by `Sikuli.connect`; default `127.0.0.1:50051`)
- `SIKULI_GRPC_AUTH_TOKEN` (optional; sent as `x-api-key` for spawned/connected sessions)
- `SIKULI_APP_NAME` (optional; used by `examples/app.js`)
