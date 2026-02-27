# sikuligo (Node.js)

SikuliGO is a GoLang implementation of Sikuli visual automation. This package provides the Node.js SDK for launching `sikuligo` locally and executing automation with a small API surface.

## Setup

Install `sikuligo` on macOS (optional, if not using local repo binary):

```bash
brew tap sikuligo/tap
brew install sikuligo/tap/sikuligo
```

```bash
yarn install
```

For source build details, see:

- [Build From Source](../../docs/build-from-source.md)

## Quickstart

Run:

```bash
yarn workspace @sikuligo/sikuligo example:workflow:connect
```

`yarn workspace @sikuligo/sikuligo example:workflow:connect` runs:

```js
import { Screen, Pattern } from "../src";

async function main() {
  const screen = await Screen();
  try {
    const match = await screen.click(Pattern("assets/pattern.png").exact());
    console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
  } finally {
    await screen.close();
  }
}
```

`yarn workspace @sikuligo/sikuligo example:workflow:auto` uses the same primary constructor pattern (`connect -> spawn` fallback handled by `Screen()`):

```bash
yarn workspace @sikuligo/sikuligo example:workflow:auto
```

`yarn workspace @sikuligo/sikuligo example:workflow:auto` runs:

```js
import { Screen, Pattern } from "../src";

async function main() {
  const screen = await Screen();
  try {
    const match = await screen.click(Pattern("assets/pattern.png").exact());
    console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
  } finally {
    await screen.close();
  }
}
```

## Run Examples

```bash
yarn workspace @sikuligo/sikuligo example:workflow:auto
yarn workspace @sikuligo/sikuligo example:workflow:connect
yarn workspace @sikuligo/sikuligo example:find
yarn workspace @sikuligo/sikuligo example:click
yarn workspace @sikuligo/sikuligo example:ocr
yarn workspace @sikuligo/sikuligo example:input
yarn workspace @sikuligo/sikuligo example:app
yarn workspace @sikuligo/sikuligo example:user-flow
yarn workspace @sikuligo/sikuligo doctor
```

## Environment
- `SIKULIGO_BINARY_PATH` (optional explicit path to `sikuligo`)
- `SIKULI_GRPC_ADDR` (optional address used by `start/launch` probe/connect; default probe `127.0.0.1:50051`)
- `SIKULI_GRPC_AUTH_TOKEN` (optional; sent as `x-api-key` for spawned/connected sessions)
- `SIKULI_MATCHER_ENGINE` (optional default matcher engine: `template`, `orb`, `hybrid`)
- `SIKULI_DEBUG` (optional; set to `1` to log launcher and per-RPC timing details; spawned `sikuligo` logs are shown too)
- `SIKULIGO_SQLITE_PATH` (optional sqlite path for spawned server sessions; default `sikuligo.db`)
- `SIKULI_APP_NAME` (optional; used by `examples/app.js`)

Primary constructors:
- `Screen()` / `Sikuli()` = connect to default address first (1s), else spawn
- `Screen.connect()` / `Sikuli.connect()` = connect only
- `Screen.spawn()` / `Sikuli.spawn()` = spawn only

## Matcher Engine Selection

Set matcher engine per session:

```js
import { Screen, Pattern } from "../src";

const screen = await Screen({ matcherEngine: "hybrid" });
const match = await screen.click(Pattern("assets/pattern.png").exact());
await screen.close();
```

Override matcher engine per call (ad-hoc):

```js
import { Screen, Pattern } from "../src";

const screen = await Screen({ matcherEngine: "template" });
const match = await screen.click(Pattern("assets/pattern.png").exact(), "orb");
await screen.close();
```
