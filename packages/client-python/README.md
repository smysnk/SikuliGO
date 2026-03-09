# sikuli-go (Python)
<!-- DOCS_CANONICAL_TARGET: docs/python-client/index.md -->
<!-- DOCS_SOURCE_MAP: docs/strategy/documentation-source-map.md -->
<!-- DOCS_WORKFLOW: docs/contribution/docs-workflow.md -->

This directory contains the Python client for sikuli-go with Sikuli-style `Screen` + `Pattern` APIs.

## Canonical Documentation

Long-form Python docs now live in the published guide:

- [Python Client](https://smysnk.github.io/sikuli-go/python-client/)
- [Installation](https://smysnk.github.io/sikuli-go/python-client/installation)
- [First Script](https://smysnk.github.io/sikuli-go/python-client/first-script)
- [Runtime](https://smysnk.github.io/sikuli-go/python-client/runtime)
- [Troubleshooting](https://smysnk.github.io/sikuli-go/python-client/troubleshooting)
- [Getting Help](https://smysnk.github.io/sikuli-go/getting-help/)

## Quickstart

`init:py-examples` prompts for the target directory, creates `requirements.txt`, installs into `.venv`, and copies examples.
Each example bootstraps `sikuli-go` into `./.sikuli-go/bin` and prepends it to PATH for the process.
The published package name is `sikuli-go`; the import module remains `sikuligo`.

```bash
pipx run sikuli-go init:py-examples
cd sikuli-go-demo
python3 examples/click.py
```

runs:
```python
from __future__ import annotations
from sikuligo import Pattern, Screen

screen = Screen()
try:
    match = screen.click(Pattern("assets/pattern.png").exact())
    print(f"clicked match target at ({match.target_x}, {match.target_y})")
finally:
    screen.close()
```

## Install Into An Existing Environment

If you manage your own environment instead of using the scaffold:

```bash
python -m pip install sikuli-go
```

Package name vs import name:

- install: `sikuli-go`
- import: `sikuligo`

## Runtime Helpers

Install the runtime binaries on PATH:

```bash
pipx run sikuli-go install-binary
```

Start the runtime and dashboard:

```bash
pipx run sikuli-go -listen
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

For the full runtime model, environment inputs, and troubleshooting notes, use the Python guide pages above.
