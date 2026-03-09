---
layout: guide
title: Guide-Format Documentation Plan
nav_key: strategy
kicker: Strategy
lead: The implementation plan for moving the docs site from link lists and deep technical pages to a guide-first structure.
---

This plan restructures the public docs into a shallow, task-first guide with consistent top-level sections for onboarding, client usage, API usage, support, and project metadata.

## Objectives

- Make the docs readable in a guide format instead of a loose collection of indexes and deep technical references.
- Put first-run instructions and language-specific workflows ahead of internal strategy and parity material.
- Reuse existing content from the repo instead of rewriting from scratch.
- Keep the current GitHub Pages + Jekyll publish flow intact in the first implementation pass.
- Avoid external-brand references in the new end-user guide pages.

## Current-State Assessment

The current docs already contain most of the raw material, but it is split across:

- `README.md` for the fastest Node.js and Python starts, binary install, dashboard, downloads, and project overview.
- `packages/client-node/README.md` for Node-specific setup and runtime behavior.
- `packages/client-python/README.md` for Python-specific setup and runtime behavior.
- `docs/guides/*.md` for build, install, OCR, input, app control, and Node package flow.
- `docs/reference/api/` for generated Go API reference.
- `docs/strategy/*.md` and `docs/reference/parity/*.md` for internal architecture and parity planning.

The gaps are structural, not informational:

- No clear top-level journey from download to first working script.
- Client docs live partly in package READMEs and partly in repo docs.
- Help, contribution, license, and downloads are not first-class sections.
- The docs home page is a link list, not a guide entry point.
- The generated Go API reference exists, but there is no curated Go-facing landing page for it.

## Constraints

- Docs are published from `docs/` through `.github/workflows/docs-pages.yml`.
- The site currently builds with Jekyll from raw Markdown under `docs/`.
- Generated API docs under `docs/reference/api/` should remain script-owned.
- Existing deep links should keep working during migration.
- The first pass should not require moving the site to a different docs framework.

## Target Information Architecture

Primary navigation order:

1. `Downloads`
2. `Getting Started`
3. `Node.js Client`
4. `Python Client`
5. `Golang API`
6. `Getting Help`
7. `Contribution`
8. `License`

Secondary navigation:

- `Reference`
- `Guides`
- `Strategy`
- `Benchmarks`

Recommended file layout:

- `docs/downloads/index.md`
- `docs/getting-started/index.md`
- `docs/getting-started/installation.md`
- `docs/getting-started/first-script.md`
- `docs/getting-started/dashboard.md`
- `docs/nodejs-client/index.md`
- `docs/nodejs-client/installation.md`
- `docs/nodejs-client/first-script.md`
- `docs/nodejs-client/runtime.md`
- `docs/nodejs-client/troubleshooting.md`
- `docs/python-client/index.md`
- `docs/python-client/installation.md`
- `docs/python-client/first-script.md`
- `docs/python-client/runtime.md`
- `docs/python-client/troubleshooting.md`
- `docs/golang-api/index.md`
- `docs/golang-api/installation.md`
- `docs/golang-api/first-program.md`
- `docs/golang-api/core-types.md`
- `docs/golang-api/runtime-and-reference.md`
- `docs/getting-help/index.md`
- `docs/getting-help/faq.md`
- `docs/getting-help/reporting-issues.md`
- `docs/contribution/index.md`
- `docs/contribution/development-setup.md`
- `docs/contribution/docs-workflow.md`
- `docs/license/index.md`

## Section Plan

### Downloads

Purpose:
- Give users one obvious place to choose a distribution channel.

Content:
- Go binary downloads from GitHub Releases.
- Node package install via npm/Yarn.
- Python package install via PyPI/pipx.
- Binary install helper commands.
- Platform matrix and what each download contains.

Primary source material:
- `README.md`
- `docs/guides/api-publish-install.md`
- `packages/client-node/README.md`
- `packages/client-python/README.md`
- `.github/workflows/client-release.yml`
- `scripts/clients/release-node-client.sh`
- `scripts/clients/release-python-client.sh`
- `scripts/clients/release-homebrew.sh`

### Getting Started

Purpose:
- Replace the current docs home page with the shortest successful first-run path.

Content:
- What sikuli-go is.
- Which path to choose: Node.js, Python, or Go.
- Install prerequisites.
- Run the first example.
- Start the API and dashboard.
- Verify the environment works.

Primary source material:
- `README.md`
- `docs/index.md`
- `packages/client-node/README.md`
- `packages/client-python/README.md`

### Node.js Client

Purpose:
- Mirror the same style as the API section, but centered on the Node client workflow.

Content shape:
- Overview and intended use.
- Install and bootstrap.
- First script.
- Runtime model: auto-started binary, monitor, dashboard, ports, session store.
- Client concepts: `Screen`, `Region`, `Pattern`, `Match`.
- Links to deeper reference and advanced guides.
- Troubleshooting and compatibility notes.

Primary source material:
- `packages/client-node/README.md`
- `docs/guides/node-package-user-flow.md`
- `docs/strategy/client-strategy.md`
- `packages/client-node/src/*.ts`

### Python Client

Purpose:
- Match the Node section pattern while staying idiomatic for Python packaging and usage.

Content shape:
- Overview and intended use.
- Install and bootstrap.
- First script.
- Runtime model: bundled binary resolution, monitor, dashboard, ports, session store.
- Client concepts: `Screen`, `Region`, `Pattern`, `Match`.
- Links to deeper reference and advanced guides.
- Troubleshooting and compatibility notes.

Primary source material:
- `packages/client-python/README.md`
- `packages/client-python/sikuligo/*.py`
- `docs/strategy/client-strategy.md`

### Golang API

Purpose:
- Provide a curated guide for the base implementation instead of dropping users directly into generated API output.

Content shape:
- Overview and intended use.
- Build or install the Go binary/API runtime.
- First program or first API-backed workflow.
- Core concepts: runtime, `Screen`, `Region`, `Pattern`, `Match`, search, OCR, input, app control.
- Link into generated API reference for details.
- Call out what is stable API surface versus deeper internal packages.

Primary source material:
- `docs/reference/api/index.md`
- `docs/reference/api/pkg-sikuli.md`
- `docs/guides/build-from-source.md`
- `docs/guides/ocr-integration.md`
- `docs/guides/input-automation.md`
- `docs/guides/app-control.md`
- `packages/api/pkg/sikuli/*.go`

### Getting Help

Purpose:
- Make support paths explicit and reduce support traffic caused by scattered information.

Content:
- FAQ.
- Common install/runtime failures.
- How to collect version, platform, and log details.
- Where to report bugs and ask for help.
- Pointers to API reference and advanced guides when the answer is already documented.

Primary source material:
- `README.md`
- `packages/client-node/README.md`
- `packages/client-python/README.md`
- existing troubleshooting notes embedded in guides and scripts

### Contribution

Purpose:
- Give contributors a clear path from local setup to a docs/code change.

Content:
- Repository layout.
- Local build/test flow.
- Docs generation and verification flow.
- Client release and API docs generation references.
- Pull request expectations.

Primary source material:
- `README.md`
- `docs/guides/build-from-source.md`
- `.github/workflows/go-test.yml`
- `.github/workflows/docs-pages.yml`
- scripts under `scripts/`

### License

Purpose:
- Make project licensing visible from the public guide.

Content:
- Short plain-language summary.
- Link to full license text.
- Copyright notice source.

Primary source material:
- `LICENSE`

## Page Template Standard

Each primary section should use the same page rhythm:

1. One-sentence summary.
2. Who this section is for.
3. Fast path or quickstart.
4. Core concepts.
5. Common tasks.
6. Troubleshooting or next steps.
7. Links to deeper reference.

This keeps the Node.js Client, Python Client, and Golang API sections visually and structurally parallel.

## Navigation and Presentation Plan

Phase 1 should add a reusable docs shell inside `docs/`:

- `docs/_config.yml` for site title, defaults, and collections if needed.
- `docs/_layouts/guide.html` for a persistent guide layout.
- `docs/_includes/sidebar.html` for top-level and section navigation.
- `docs/_data/navigation.yml` to define the ordered left-nav structure.

If the layout work is deferred, the fallback is still valid:

- keep plain Markdown pages
- create section landing pages with identical heading order
- update `docs/index.md` to act as the main guide gateway

The preferred path is the reusable shell, because it is what makes the guide format feel deliberate instead of just link-heavy.

## Migration Strategy

### Phase 0: Freeze inputs

- Identify the canonical source for each section.
- Stop adding new end-user guidance only to package READMEs without a matching docs-page target.

### Phase 1: Build the guide shell

- Add the shared layout, sidebar data, and top-level landing pages.
- Rework `docs/index.md` into the public entry page for the guide.

### Phase 2: Publish the top-level sections

- Ship `Downloads`, `Getting Started`, `Node.js Client`, `Python Client`, `Golang API`, `Getting Help`, `Contribution`, and `License` as landing pages first.
- Link existing detailed guides from these pages before rewriting all detail pages.

### Phase 3: Split dense pages into guide subpages

- Move install, first-run, runtime, and troubleshooting material into subpages.
- Keep generated API reference where it is and link it from `Golang API`.

### Phase 4: Reconcile README and docs ownership

- Shorten package READMEs so they point to the guide for canonical docs.
- Keep package READMEs useful for registry users, but stop duplicating long-form docs there.

### Phase 5: Clean up and harden

- Add redirects or stub pages for moved content.
- Remove stale wording and external-brand references from new user-facing guide pages.
- Verify all internal links in the published site.

## Content Ownership Rules

- `docs/` becomes the canonical home for long-form end-user documentation.
- package `README.md` files become package entry pages and quickstarts, not the full manual.
- generated API content stays under `docs/reference/api/`.
- strategy and parity pages remain valuable, but they move out of the primary user journey.

## Acceptance Criteria

- A new user can go from docs home to a working first script in no more than three clicks.
- Node.js, Python, and Go each have a dedicated landing page with the same structure.
- Downloads, help, contribution, and license are visible in the primary navigation.
- The Go API section explains the curated surface before linking into generated API reference.
- The docs site keeps building through the existing Pages workflow.
- New user-facing pages do not rely on external-brand references.

## Recommended Execution Order

1. Create the guide shell and top-level navigation.
2. Publish `Getting Started` and `Downloads`.
3. Publish `Node.js Client`, `Python Client`, and `Golang API`.
4. Publish `Getting Help`, `Contribution`, and `License`.
5. Reconcile README duplication and old page links.
