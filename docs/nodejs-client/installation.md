---
layout: guide
title: Node.js Client Installation
nav_key: nodejs-client
kicker: Node.js Client
lead: Install the Node package, scaffold a working project, and optionally install the runtime binaries on PATH.
---

## Who This Page Is For

Use this page when Node.js is the primary entry point and you need package install and bootstrap guidance before writing scripts.

## Fastest Path

```bash
yarn dlx @sikuligo/sikuli-go init:js-examples
```

This scaffolds a project directory, writes `package.json`, installs the current package dependency, and copies `.mjs` examples into `examples/`.

## Package Install

If you are wiring the package into an existing project, use the published package:

```bash
npm install @sikuligo/sikuli-go
# or
yarn add @sikuligo/sikuli-go
```

## Install The Runtime On PATH

```bash
yarn dlx @sikuligo/sikuli-go install-binary
```

This installs `sikuli-go` and `sikuli-go-monitor` into `~/.local/bin`.

## Verify The Node Path

```bash
yarn dlx @sikuligo/sikuli-go init:js-examples
cd sikuli-go-demo
yarn node examples/click.mjs
```

## Next Pages

- [Node.js Client: First Script]({{ '/nodejs-client/first-script' | relative_url }})
- [Node.js Client: Runtime]({{ '/nodejs-client/runtime' | relative_url }})
- [Node.js Client: Troubleshooting]({{ '/nodejs-client/troubleshooting' | relative_url }})
