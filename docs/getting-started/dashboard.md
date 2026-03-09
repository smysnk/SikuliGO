---
layout: guide
title: Getting Started Dashboard
nav_key: getting-started
kicker: Getting Started
lead: Start the runtime, open the dashboard, and use the standalone monitor without digging through the deeper runtime docs.
---

## Runtime And Monitor

- `sikuli-go` is the runtime process. It serves the API and can also expose the dashboard.
- `sikuli-go-monitor` is the monitor-only process. It reads the shared `sikuli-go.db` session store without starting another automation server.

## Start The Runtime

From the Node.js package path:

```bash
yarn dlx @sikuligo/sikuli-go -listen
```

From the Python package path:

```bash
pipx run sikuli-go -listen
```

From a binary already on PATH:

```bash
sikuli-go -listen 127.0.0.1:50051 -admin-listen :8080
```

## Open The Dashboard

Use the default admin URLs:

- `http://127.0.0.1:8080/dashboard`
- `http://127.0.0.1:8080/healthz`
- `http://127.0.0.1:8080/metrics`
- `http://127.0.0.1:8080/snapshot`

## Start The Standalone Monitor

```bash
sikuli-go-monitor
```

By default it serves the monitor UI on `:8080` and reads `sikuli-go.db` from the current working directory.

## Common Questions

- If you only want to inspect sessions, use `sikuli-go-monitor`.
- If the dashboard is unavailable, confirm the runtime is listening and that nothing else is bound to `:8080`.
- If you are using the package wrappers, the runtime can still be started in auto mode by the client even when you do not start it manually first.

## Next Pages

- [Node.js Client: Runtime]({{ '/nodejs-client/runtime' | relative_url }})
- [Python Client: Runtime]({{ '/python-client/runtime' | relative_url }})
- [Getting Help: FAQ]({{ '/getting-help/faq' | relative_url }})
