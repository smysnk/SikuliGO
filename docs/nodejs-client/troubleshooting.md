---
layout: guide
title: Node.js Client Troubleshooting
nav_key: nodejs-client
kicker: Node.js Client
lead: Resolve the common runtime, binary, and desktop-environment issues that show up in the Node.js workflow.
---

## Binary Not Found

- Run `yarn dlx @sikuligo/sikuli-go install-binary`.
- Reload your shell if `sikuli-go` is still not on PATH.
- If you manage the runtime separately, point the client at that runtime with `Screen.connect()`.

## Runtime Startup Fails

Common causes:

- no active desktop session
- missing OS accessibility/input permissions
- unsupported platform packaging
- startup timeout while spawning the runtime

The Node package user flow also calls out these startup error classes explicitly:

- missing binary
- permission denial
- unsupported platform
- startup timeout

## Existing Runtime vs Spawned Runtime

If you already have a runtime listening, use the connect path instead of forcing a new process:

```js
import { Screen } from "@sikuligo/sikuli-go";

const screen = await Screen.connect({ address: "127.0.0.1:50051" });
```

## Diagnostics

```bash
npx @sikuligo/sikuli-go doctor
```

Use this to validate environment assumptions before digging into deeper runtime issues.

## More References

- [Node.js Client: Runtime]({{ '/nodejs-client/runtime' | relative_url }})
- [Getting Help: FAQ]({{ '/getting-help/faq' | relative_url }})
- [Node Package User Flow]({{ '/guides/node-package-user-flow' | relative_url }})
