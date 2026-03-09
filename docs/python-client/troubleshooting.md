---
layout: guide
title: Python Client Troubleshooting
nav_key: python-client
kicker: Python Client
lead: Resolve the common package, binary-resolution, and desktop-environment issues that show up in the Python workflow.
---

## Package Name vs Import Name

The package name and import name are intentionally different:

- install: `sikuli-go`
- import: `sikuligo`

If your import fails after install, check that you installed the published package into the Python environment you are actually running.

## Runtime Binary Cannot Be Resolved

The client looks for the runtime in several places:

- an explicit `binary_path`
- `SIKULI_GO_BINARY_PATH`
- common repo-local locations
- executables on PATH

If resolution fails, build the runtime locally or set `SIKULI_GO_BINARY_PATH`.

## Runtime Startup Fails

Common causes:

- no active desktop session
- missing OS accessibility/input permissions
- startup timeout while spawning the runtime
- trying to run a desktop-automation flow in a truly headless environment

## Existing Runtime vs Spawned Runtime

If you already have a runtime listening, connect explicitly:

```python
from sikuligo import Screen

screen = Screen.connect(address="127.0.0.1:50051")
```

## More References

- [Python Client: Runtime]({{ '/python-client/runtime' | relative_url }})
- [Getting Help: FAQ]({{ '/getting-help/faq' | relative_url }})
- [Build From Source]({{ '/guides/build-from-source' | relative_url }})
