---
layout: guide
title: Documentation Canonical Source Map
nav_key: strategy
kicker: Strategy
lead: The current ownership map for the public docs sections and their guide landing pages.
---

This document started as the phase 0 ownership freeze and now tracks the current canonical source for each public docs section as the guide-format rewrite lands.

Use this file before editing user-facing documentation:

1. Find the section you are changing.
2. Update the current canonical source listed here.
3. Update the matching target page in `docs/` if it exists.
4. Do not land end-user documentation changes only in a package README.

## Rules

- `docs/` is the long-term home for end-user documentation.
- Package `README.md` files remain valid quickstarts, but they are not allowed to drift without a matching docs target.
- Generated content under `docs/reference/api/` stays script-owned.
- Strategy and parity pages can support end-user pages, but they are not the primary onboarding surface.

## Section Inventory

## Downloads

Current canonical source:
- `docs/downloads/index.md`

Target page:
- `docs/downloads/index.md`

Supporting references:
- `docs/guides/api-publish-install.md`
- `packages/client-node/README.md`
- `packages/client-python/README.md`
- `.github/workflows/client-release.yml`
- `scripts/clients/release-node-client.sh`
- `scripts/clients/release-python-client.sh`
- `scripts/clients/release-homebrew.sh`

## Getting Started

Current canonical source:
- `docs/getting-started/index.md`
- `docs/getting-started/installation.md`
- `docs/getting-started/first-script.md`
- `docs/getting-started/dashboard.md`

Target page:
- `docs/getting-started/index.md`

Supporting references:
- `docs/index.md`
- `packages/client-node/README.md`
- `packages/client-python/README.md`

## Node.js Client

Current canonical source:
- `docs/nodejs-client/index.md`
- `docs/nodejs-client/installation.md`
- `docs/nodejs-client/first-script.md`
- `docs/nodejs-client/runtime.md`
- `docs/nodejs-client/troubleshooting.md`

Target page:
- `docs/nodejs-client/index.md`

Supporting references:
- `packages/client-node/README.md`
- `docs/guides/node-package-user-flow.md`
- `docs/strategy/client-strategy.md`
- `packages/client-node/src/`

## Python Client

Current canonical source:
- `docs/python-client/index.md`
- `docs/python-client/installation.md`
- `docs/python-client/first-script.md`
- `docs/python-client/runtime.md`
- `docs/python-client/troubleshooting.md`

Target page:
- `docs/python-client/index.md`

Supporting references:
- `packages/client-python/README.md`
- `docs/strategy/client-strategy.md`
- `packages/client-python/sikuligo/`

## Golang API

Current canonical source:
- `docs/golang-api/index.md`
- `docs/golang-api/installation.md`
- `docs/golang-api/first-program.md`
- `docs/golang-api/core-types.md`
- `docs/golang-api/runtime-and-reference.md`

Target page:
- `docs/golang-api/index.md`

Supporting references:
- `docs/reference/api/index.md`
- `docs/reference/api/pkg-sikuli.md`
- `docs/guides/build-from-source.md`
- `docs/guides/ocr-integration.md`
- `docs/guides/input-automation.md`
- `docs/guides/app-control.md`
- `packages/api/pkg/sikuli/`

## Getting Help

Current canonical source:
- `docs/getting-help/index.md`
- `docs/getting-help/faq.md`
- `docs/getting-help/reporting-issues.md`

Target page:
- `docs/getting-help/index.md`

Supporting references:
- `packages/client-node/README.md`
- `packages/client-python/README.md`
- `docs/guides/build-from-source.md`

## Contribution

Current canonical source:
- `docs/contribution/index.md`
- `docs/contribution/development-setup.md`
- `docs/contribution/docs-workflow.md`

Target page:
- `docs/contribution/index.md`

Supporting references:
- `README.md`
- `docs/guides/build-from-source.md`
- `.github/workflows/go-test.yml`
- `.github/workflows/docs-pages.yml`
- `scripts/`

## License

Current canonical source:
- `LICENSE`

Target page:
- `docs/license/index.md`

Supporting references:
- `README.md`

## Secondary Sections

These sections are not part of the primary guide flow, but they still have fixed ownership:

- `Reference`: current canonical sources are `docs/reference/index.md` and generated pages under `docs/reference/api/`.
- `Guides`: current canonical source is `docs/guides/index.md` plus the guide pages under `docs/guides/`.
- `Strategy`: current canonical source is `docs/strategy/index.md` plus strategy pages under `docs/strategy/`.
- `Benchmarks`: current canonical source is the bench pages and reports under `docs/bench/`.
