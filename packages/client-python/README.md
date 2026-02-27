# SikuliGO Python Client

This directory contains the Python client for SikuliGO with Sikuli-style `Screen` + `Pattern` APIs.

## Prerequisites
- Python 3.10+
- `protoc`

## Setup

Install `sikuligo` on macOS (optional, if not using local repo binary):

```bash
brew tap sikuligo/tap
brew install sikuligo/tap/sikuligo
```

```bash
cd packages/client-python
python3 -m venv .venv
source .venv/bin/activate
python3 -m pip install -r requirements.txt
```

For source build details, see:

- [Build From Source](../../docs/build-from-source.md)

Install from PyPI:

```bash
pip install sikuligo
```

## Quickstart

Run:

```bash
cd packages/client-python
python3 examples/workflow_connect.py
```

`python3 examples/workflow_connect.py` runs:

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

`python3 examples/workflow_auto_launch.py` uses the same primary constructor pattern (`connect -> spawn` fallback handled by `Screen()`):

```bash
cd packages/client-python
python3 examples/workflow_auto_launch.py
```
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

## Environment

- `SIKULI_GRPC_ADDR` (optional address used by `Screen()` probe/connect; default probe `127.0.0.1:50051`)
- `SIKULI_GRPC_AUTH_TOKEN` (optional; sent as `x-api-key`)
- `SIKULI_MATCHER_ENGINE` (optional default matcher engine: `template`, `orb`, `hybrid`)
- `SIKULIGO_SQLITE_PATH` (optional sqlite path for spawned server sessions; default `sikuligo.db`)

Primary constructors:
- `Screen()` = connect to default address first (1s), else spawn
- `Screen.connect()` = connect only
- `Screen.spawn()` = spawn only

## Matcher Engine Selection

Set matcher engine per session:

```python
from sikuligo import Pattern, Screen

screen = Screen(matcher_engine="hybrid")
try:
    match = screen.click(Pattern("assets/pattern.png").exact())
finally:
    screen.close()
```

Override matcher engine per call (ad-hoc):

```python
from sikuligo import Pattern, Screen

screen = Screen(matcher_engine="template")
try:
    match = screen.click(Pattern("assets/pattern.png").exact(), engine="orb")
finally:
    screen.close()
```

## Run Additional Examples

```bash
cd packages/client-python
python3 examples/find.py
python3 examples/read_text.py
python3 examples/click_and_type.py
python3 examples/app_control.py
```
