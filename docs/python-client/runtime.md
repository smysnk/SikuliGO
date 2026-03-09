---
layout: guide
title: Python Client Runtime
nav_key: python-client
kicker: Python Client
lead: Understand how the Python client connects to, starts, and scopes the runtime before you move into more complex workflows.
---

## Runtime Modes

The screen-facing wrapper supports three explicit ways to attach to the runtime:

- `Screen()` or `Screen.start()` uses auto mode
- `Screen.connect()` attaches to an already-running runtime
- `Screen.spawn()` forces a new runtime process

There is also an explicit `Screen.auto()` classmethod when you want to say that mode directly.

## Auto Mode

Auto mode probes for an existing runtime first and only spawns a new process if that connection attempt fails.

## Screen And Region Surface

The current Python wrapper supports:

- `find`
- `exists`
- `wait`
- `wait_vanish`
- `click`
- `hover`
- `region(x, y, w, h)` for bounded searches

## Environment Inputs

The runtime path and connection flow can be shaped by environment variables already used in the client source:

- `SIKULI_GRPC_ADDR`
- `SIKULI_GRPC_AUTH_TOKEN`
- `SIKULI_GO_BINARY_PATH`
- `SIKULI_GO_SQLITE_PATH`

## Dashboard And Monitor

Start the runtime from the package path:

```bash
pipx run sikuli-go -listen
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

## Next Pages

- [Python Client: First Script]({{ '/python-client/first-script' | relative_url }})
- [Python Client: Troubleshooting]({{ '/python-client/troubleshooting' | relative_url }})
- [Getting Started: Dashboard]({{ '/getting-started/dashboard' | relative_url }})
