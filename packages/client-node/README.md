# sikuli-go (Node.js)
<!-- DOCS_CANONICAL_TARGET: docs/nodejs-client/index.md -->
<!-- DOCS_SOURCE_MAP: docs/strategy/documentation-source-map.md -->
<!-- DOCS_WORKFLOW: docs/contribution/docs-workflow.md -->

sikuli-go is a Go implementation of Sikuli visual automation. This package provides the Node.js SDK for launching `sikuli-go` locally and executing automation with a small API surface.

## Canonical Documentation

Long-form Node.js docs now live in the published guide:

- [Node.js Client](https://smysnk.github.io/sikuli-go/nodejs-client/)
- [Installation](https://smysnk.github.io/sikuli-go/nodejs-client/installation)
- [First Script](https://smysnk.github.io/sikuli-go/nodejs-client/first-script)
- [Runtime](https://smysnk.github.io/sikuli-go/nodejs-client/runtime)
- [Troubleshooting](https://smysnk.github.io/sikuli-go/nodejs-client/troubleshooting)
- [Getting Help](https://smysnk.github.io/sikuli-go/getting-help/)

## Quickstart

`init:js-examples` prompts for a target directory, scaffolds a `package.json` with the latest `@sikuligo/sikuli-go` dependency, runs `yarn install`, and copies `.mjs` examples into `examples/`.

```bash
yarn dlx @sikuligo/sikuli-go init:js-examples
cd sikuli-go-demo
yarn node examples/click.mjs
```

```js
import { Screen, Pattern } from "@sikuligo/sikuli-go";

const screen = await Screen();
try {
  const match = await screen.click(Pattern("assets/pattern.png").exact());
  console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
} finally {
  await screen.close();
}
```

## Install In An Existing Project

```bash
npm install @sikuligo/sikuli-go
# or
yarn add @sikuligo/sikuli-go
```

## Runtime Helpers

Install the runtime binaries on PATH:

```bash
yarn dlx @sikuligo/sikuli-go install-binary
```

Start the runtime and dashboard:

```bash
yarn dlx @sikuligo/sikuli-go -listen
```

Open:

- http://127.0.0.1:8080/dashboard

Run the standalone monitor after installing the binaries:

```bash
sikuli-go-monitor
```

By default it serves the monitor UI on `:8080` and reads `sikuli-go.db` from the current working directory.

## Common Package Entry Points

- `Screen()` or `Screen.start()` uses auto mode
- `Screen.connect()` attaches to an existing runtime
- `Screen.spawn()` forces a new runtime process
- `screen.region(x, y, w, h)` scopes the search area

For the full runtime model, diagnostics flow, and troubleshooting notes, use the Node.js guide pages above.
