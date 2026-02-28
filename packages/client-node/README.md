# sikuligo (Node.js)

SikuliGO is a GoLang implementation of Sikuli visual automation. This package provides the Node.js SDK for launching `sikuligo` locally and executing automation with a small API surface.

## Quickstart

Use the published package examples in any project directory:

```bash
mkdir sikuligo-demo
cd sikuligo-demo
yarn init -y
yarn add @sikuligo/sikuligo
yarn init-examples
yarn node examples/click.mjs
```

```js
import { Screen, Pattern } from "@sikuligo/sikuligo";

const screen = await Screen();
try {
  const match = await screen.click(Pattern("assets/pattern.png").exact());
  console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
} finally {
  await screen.close();
}
```

## Run Examples

```bash
yarn node examples/workflow-auto-launch.mjs
yarn node examples/workflow-connect.mjs
yarn node examples/find.mjs
yarn node examples/click.mjs
yarn node examples/ocr.mjs
yarn node examples/input.mjs
yarn node examples/app.mjs
yarn node examples/user-flow.mjs
yarn sikuligo-doctor
```

## Environment
- `SIKULIGO_BINARY_PATH` (optional explicit path to `sikuligo`)
- `SIKULI_GRPC_ADDR` (optional address used by `start/launch` probe/connect; default probe `127.0.0.1:50051`)
- `SIKULI_GRPC_AUTH_TOKEN` (optional; sent as `x-api-key` for spawned/connected sessions)
- `SIKULI_MATCHER_ENGINE` (optional default matcher engine: `template`, `orb`, `hybrid`)
- `SIKULI_DEBUG` (optional; set to `1` to log launcher and per-RPC timing details; spawned `sikuligo` logs are shown too)
- `SIKULIGO_SQLITE_PATH` (optional sqlite path for spawned server sessions; default `sikuligo.db`)
- `SIKULI_APP_NAME` (optional; used by `examples/app.mjs`)

Primary constructors:
- `Screen()` / `Sikuli()` = connect to default address first (1s), else spawn
- `Screen.connect()` / `Sikuli.connect()` = connect only
- `Screen.spawn()` / `Sikuli.spawn()` = spawn only

## Matcher Engine Selection

Set matcher engine per session:

```js
import { Screen, Pattern } from "@sikuligo/sikuligo";

const screen = await Screen({ matcherEngine: "hybrid" });
const match = await screen.click(Pattern("assets/pattern.png").exact());
await screen.close();
```

Override matcher engine per call (ad-hoc):

```js
import { Screen, Pattern } from "@sikuligo/sikuligo";

const screen = await Screen({ matcherEngine: "template" });
const match = await screen.click(Pattern("assets/pattern.png").exact(), "orb");
await screen.close();
```
