# SikuliGO Python Client

This directory contains a minimal Python gRPC client wrapper for `sikuli.v1.SikuliService`.

## Prerequisites

- Python 3.10+
- `protoc`
- SikuliGO gRPC server running (default `127.0.0.1:50051`)

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

## Environment

- `SIKULI_GRPC_ADDR` (default: `127.0.0.1:50051`)
- `SIKULI_GRPC_AUTH_TOKEN` (optional; sent as `x-api-key`)

## Quick Example

```python
from generated.sikuli.v1 import sikuli_pb2 as pb
from sikuligo_client.client import SikuliGrpcClient

client = SikuliGrpcClient(address="127.0.0.1:50051")
try:
    client.click(pb.ClickRequest(x=300, y=220))
    client.type_text(pb.TypeTextRequest(text="hello from sikuligo"))
    client.hotkey(pb.HotkeyRequest(keys=["cmd", "enter"]))
finally:
    client.close()
```

## Run Examples

```bash
cd clients/python
PYTHONPATH=. python3 examples/find.py
PYTHONPATH=. python3 examples/read_text.py
PYTHONPATH=. python3 examples/click_and_type.py
PYTHONPATH=. python3 examples/app_control.py
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
