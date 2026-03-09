---
layout: guide
title: Getting Started Installation
nav_key: getting-started
kicker: Getting Started
lead: Install the fastest working path for Node.js, Python, or a direct runtime workflow before running the first script.
---

## Who This Page Is For

Use this page when you have not installed anything yet and want the shortest route to a working setup.

## Prerequisites

- Node.js `20+`
- Yarn `4+`
- Python `3.10+`
- Go `1.24+` and `protoc` only if you are building from source

## Fast Install Options

### Node.js Quickstart Project

```bash
yarn dlx @sikuligo/sikuli-go init:js-examples
```

Use this when you want a JavaScript project scaffold with examples and the current package dependency already wired in.

### Python Quickstart Project

```bash
pipx run sikuli-go init:py-examples
```

Use this when you want a Python project scaffold with `requirements.txt`, a local `.venv`, and copied examples.

### Install Runtime Binaries On PATH

```bash
yarn dlx @sikuligo/sikuli-go install-binary
# or
pipx run sikuli-go install-binary
```

This puts `sikuli-go` and `sikuli-go-monitor` in `~/.local/bin` and can update your shell profile.

### Build From Source

```bash
make
```

Use this path if you are working in the repository, need generated artifacts, or want to verify local changes.

## Verify The Install

If the runtime is on your PATH, start it explicitly:

```bash
sikuli-go -listen 127.0.0.1:50051 -admin-listen :8080
```

Then open `http://127.0.0.1:8080/dashboard`.

## Common Questions

- If the binary is not found after `install-binary`, reload `~/.zshrc` or `~/.bash_profile`.
- If you only need the package workflow, you do not need to install the runtime on PATH first.
- If you are already inside the repo, use [Build From Source]({{ '/guides/build-from-source' | relative_url }}) instead of repeating package-level setup.

## Next Pages

- [Getting Started: First Script]({{ '/getting-started/first-script' | relative_url }})
- [Getting Started: Dashboard]({{ '/getting-started/dashboard' | relative_url }})
- [Downloads]({{ '/downloads/' | relative_url }})
