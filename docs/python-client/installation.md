---
layout: guide
title: Python Client Installation
nav_key: python-client
kicker: Python Client
lead: Install the Python package, scaffold a working project, and optionally install the runtime binaries on PATH.
---

## Who This Page Is For

Use this page when Python is the primary entry point and you need package install and bootstrap guidance before writing scripts.

## Fastest Path

```bash
pipx run sikuli-go init:py-examples
```

This creates a project directory, writes `requirements.txt`, creates `.venv`, installs dependencies, and copies Python examples into `examples/`.

## Package Naming

- Published package name: `sikuli-go`
- Import module: `sikuligo`

## Install Into Your Own Environment

If you manage your own Python environment instead of using the scaffold, install the published package from PyPI:

```bash
python -m pip install sikuli-go
```

## Install The Runtime On PATH

```bash
pipx run sikuli-go install-binary
```

This installs `sikuli-go` and `sikuli-go-monitor` into `~/.local/bin`.

## Next Pages

- [Python Client: First Script]({{ '/python-client/first-script' | relative_url }})
- [Python Client: Runtime]({{ '/python-client/runtime' | relative_url }})
- [Python Client: Troubleshooting]({{ '/python-client/troubleshooting' | relative_url }})
