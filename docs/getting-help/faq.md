---
layout: guide
title: Getting Help FAQ
nav_key: getting-help
kicker: Getting Help
lead: The common install, runtime, and workflow questions that come up before you need a full issue report.
---

## Why Is `sikuli-go` Not Found After `install-binary`?

The installer can add `~/.local/bin` to your shell profile, but your current shell still needs to reload that profile. Reload `~/.zshrc` or `~/.bash_profile`, then try again.

## Which Python Name Do I Use?

- install: `sikuli-go`
- import: `sikuligo`

The package and import names are intentionally different.

## What Is The Difference Between The Runtime And The Monitor?

- `sikuli-go` runs the automation API and can expose the dashboard.
- `sikuli-go-monitor` serves the monitor UI against an existing `sikuli-go.db` session store without starting another automation server.

## Where Does The Session Store Live?

By default, `sikuli-go-monitor` reads `sikuli-go.db` from the current working directory.

## Can This Run Truly Headless?

Desktop automation depends on an active graphical session and real desktop state. It is not a true headless browser-style automation flow.

## The Runtime Starts But Automation Still Fails

Check for:

- OS accessibility/input permissions
- a visible target image on the actual desktop
- a live graphical session
- the correct connect vs spawn mode for your workflow

## Next Pages

- [Getting Help: Reporting Issues]({{ '/getting-help/reporting-issues' | relative_url }})
- [Node.js Client: Troubleshooting]({{ '/nodejs-client/troubleshooting' | relative_url }})
- [Python Client: Troubleshooting]({{ '/python-client/troubleshooting' | relative_url }})
