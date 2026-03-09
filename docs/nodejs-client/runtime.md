---
layout: guide
title: Node.js Client Runtime
nav_key: nodejs-client
kicker: Node.js Client
lead: Understand how the Node.js client connects to, starts, and scopes the runtime before you move into more complex workflows.
---

## Runtime Modes

The top-level exports support three ways to attach to the runtime:

- `Screen()` or `Screen.start()` uses auto mode.
- `Screen.connect()` attaches to an already-running runtime.
- `Screen.spawn()` forces a new runtime process.

The same pattern exists on the lower-level `Sikuli` constructor:

- `Sikuli.launch()` / `Sikuli()`
- `Sikuli.connect()`
- `Sikuli.spawn()`

## Auto Mode

Auto mode tries to connect first and spawns a runtime only when that probe fails. Use it for the default script flow when you do not want to manage server lifecycle manually.

## Screen And Region Surface

The current screen-facing wrapper supports:

- `find`
- `exists`
- `wait`
- `waitVanish`
- `click`
- `hover`
- `region(x, y, w, h)` for bounded searches

## Dashboard And Monitor

Start the runtime from the package path:

```bash
yarn dlx @sikuligo/sikuli-go -listen
```

Use the standalone monitor:

```bash
sikuli-go-monitor
```

Default admin URLs:

- `http://127.0.0.1:8080/dashboard`
- `http://127.0.0.1:8080/healthz`
- `http://127.0.0.1:8080/metrics`
- `http://127.0.0.1:8080/snapshot`

## Diagnostics

The current Node package flow also includes a diagnostics command:

```bash
npx @sikuligo/sikuli-go doctor
```

Use it when startup or binary resolution is failing and you want a quick environment check.

## Next Pages

- [Node.js Client: First Script]({{ '/nodejs-client/first-script' | relative_url }})
- [Node.js Client: Troubleshooting]({{ '/nodejs-client/troubleshooting' | relative_url }})
- [Node Package User Flow]({{ '/guides/node-package-user-flow' | relative_url }})
