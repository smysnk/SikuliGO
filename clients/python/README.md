# SikuliGO Python Client

This directory contains the Python client for SikuliGO with Sikuli-style `Screen` + `Pattern` APIs.

## Prerequisites
- Python 3.10+
- `protoc`

## Setup

```bash
cd clients/python
python3 -m venv .venv
source .venv/bin/activate
python3 -m pip install -r requirements.txt
../../scripts/clients/generate-python-stubs.sh
```

Install from PyPI:

```bash
pip install sikuligo
```

## Quickstart

Run:

```bash
cd clients/python
python3 examples/workflow_connect.py
```

`python3 examples/workflow_connect.py` runs:

```python
from __future__ import annotations
from sikuligo import Pattern, Screen

screen = Screen.auto()
try:
    match = screen.click(Pattern("assets/pattern.png").exact())
    print(f"clicked match target at ({match.target_x}, {match.target_y})")
finally:
    screen.close()
```

`python3 examples/workflow_auto_launch.py` uses the same constructor pattern (`connect -> launch`):

```bash
cd clients/python
python3 examples/workflow_auto_launch.py
```
```python
from __future__ import annotations
from sikuligo import Pattern, Screen

screen = Screen.auto()
try:
    match = screen.click(Pattern("assets/pattern.png").exact())
    print(f"clicked match target at ({match.target_x}, {match.target_y})")
finally:
    screen.close()
```

## Environment

- `SIKULI_GRPC_ADDR` (optional address used by `auto` probe/connect; default probe `127.0.0.1:50051`)
- `SIKULI_GRPC_AUTH_TOKEN` (optional; sent as `x-api-key`)
- `SIKULIGO_SQLITE_PATH` (optional sqlite path for spawned server sessions; default `sikuligo.db`)

## Run Additional Examples

```bash
cd clients/python
python3 examples/find.py
python3 examples/read_text.py
python3 examples/click_and_type.py
python3 examples/app_control.py
```

## Build/Release Scaffold

Build distributions and validate metadata:

```bash
./scripts/clients/release-python-client.sh
```

If build tools are already installed, skip installer steps:

```bash
SKIP_INSTALL=1 ./scripts/clients/release-python-client.sh
```

Publish to PyPI (requires `PYPI_TOKEN`):

```bash
PYPI_PUBLISH=1 PYPI_TOKEN=... ./scripts/clients/release-python-client.sh
```
