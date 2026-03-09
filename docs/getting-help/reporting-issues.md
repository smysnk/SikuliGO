---
layout: guide
title: Getting Help Reporting Issues
nav_key: getting-help
kicker: Getting Help
lead: Collect the right runtime, platform, and reproduction details before opening a bug report.
---

## Where To Report

- [GitHub Issues](https://github.com/smysnk/SikuliGO/issues)

## Include These Details

- OS and architecture
- whether you used Node.js, Python, or a direct Go/runtime path
- the package version or commit SHA
- the exact command used to start the runtime
- the exact script or minimal reproduction
- full error output and logs
- screenshots when the issue depends on desktop state

## Issue Report Skeleton

```text
Environment:
- OS:
- Architecture:
- Path used: Node.js / Python / Go runtime
- Package version or commit:

Runtime start command:

Minimal reproduction:

Expected behavior:

Actual behavior:

Logs / screenshots:
```

## Before You File

- Check [Getting Help: FAQ]({{ '/getting-help/faq' | relative_url }}) first.
- If the problem is package-path specific, check the matching troubleshooting page.
- If the problem is about local contributor setup, check [Contribution: Development Setup]({{ '/contribution/development-setup' | relative_url }}).

## Next Pages

- [Getting Help: FAQ]({{ '/getting-help/faq' | relative_url }})
- [Contribution]({{ '/contribution/' | relative_url }})
