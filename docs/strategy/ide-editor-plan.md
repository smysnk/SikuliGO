---
layout: guide
title: IDE Editor Implementation Plan
nav_key: strategy
kicker: Strategy
lead: The phased implementation plan for a local IDE that can run Node.js and Python examples, step through execution line by line, and embed the existing runtime dashboards.
---

This document turns the earlier editor concept into an implementation plan. The goal is not just to describe the desired IDE, but to define the build order, package ownership, repo touchpoints, validation gates, and exit criteria needed to ship it safely.

## Outcome

Ship a local desktop IDE for sikuli-go that can:

- run the existing Node.js examples in `packages/client-node/examples/`
- run the existing Python examples in `packages/client-python/examples/`
- pause and step line by line through those examples
- show stdout, stderr, runtime state, and step events in one workspace
- embed the same runtime dashboards already served by `sikuli-go`

## Current Implementation Status (March 13, 2026)

- `packages/editor/` exists as a Next.js renderer shell with a workflow sidebar and editor-area placeholder.
- `packages/api-electron/` already manages the `sikuli-go` binary lifecycle and opens the runtime dashboard/session viewer in Electron.
- Node.js and Python examples already exist and are runnable from their package directories.
- The runtime already serves the operational surfaces we want to reuse:
  - `/dashboard`
  - `/healthz`
  - `/metrics`
  - `/snapshot`
  - session viewer via the existing dashboard query flow

What is missing:

- a shared desktop shell that hosts the editor renderer and owns local process control
- a shared execution layer for Node.js and Python examples
- step instrumentation for both languages
- UI wiring between editor state, execution state, and runtime dashboards

## Product Decision

Recommendation:
- implement this as an Electron desktop IDE

Why:

- local file access, Python process execution, Node.js process execution, and runtime supervision all belong in a desktop host
- the existing `packages/api-electron/` package already solves the hardest runtime-hosting problem
- a browser-only app would still require a separate local agent, which adds more moving pieces than it removes

Implementation consequence:

- Electron becomes the process owner
- `packages/editor/` becomes the renderer UI
- `packages/api-electron/` becomes the desktop host and IPC boundary

## Scope

In scope:

- example discovery and execution for Node.js and Python
- line-by-line stepping for example files
- runtime start/connect/restart/stop controls
- embedded dashboard and session views
- execution timeline, output panel, and current-line highlighting

Out of scope for the first useful release:

- arbitrary-language debugging beyond curated example flows
- full IDE features like search across workspace, refactors, or extensions
- rebuilding the runtime dashboard UI in React
- cloud or remote execution

## Non-Goals

- replacing the existing runtime dashboard with a new editor-native implementation
- shipping a general-purpose debugger before example stepping works
- supporting Lua in the first IDE release
- turning `packages/editor/` into a browser-hosted SaaS app

## Delivery Principles

- Reuse repo pieces that already exist before adding new packages.
- Build one shared event model for both languages.
- Treat example execution as the first supported contract.
- Keep the runtime dashboard binary-owned and embed it directly.
- Add explicit validation gates before moving from one phase to the next.

## Package Ownership Plan

### `packages/api-electron/`

Primary owner for:

- runtime lifecycle management
- Electron window and webview management
- IPC surface
- Node.js/Python child-process supervision
- local workspace access

### `packages/editor/`

Primary owner for:

- editor renderer
- file tree / example picker
- code pane
- execution controls
- output panels
- step timeline
- dashboard embedding layout

### `packages/client-node/` and `packages/client-python/`

Remain owners for:

- example corpus
- client runtime behavior
- language-specific expectations the IDE must preserve

## Planned Repository Touchpoints

Likely existing files to extend:

- `packages/api-electron/src/main.js`
- `packages/api-electron/src/preload.js`
- `packages/editor/app/page.tsx`
- `packages/editor/app/page.module.css`

Likely new desktop-host files:

- `packages/api-electron/src/ipc/runtime.js`
- `packages/api-electron/src/ipc/examples.js`
- `packages/api-electron/src/ipc/execution.js`
- `packages/api-electron/src/execution/session.js`
- `packages/api-electron/src/execution/node-runner.js`
- `packages/api-electron/src/execution/python-runner.js`

Likely new renderer files:

- `packages/editor/app/lib/ipc.ts`
- `packages/editor/app/lib/types.ts`
- `packages/editor/app/components/example-tree.tsx`
- `packages/editor/app/components/code-pane.tsx`
- `packages/editor/app/components/output-panel.tsx`
- `packages/editor/app/components/step-timeline.tsx`
- `packages/editor/app/components/dashboard-pane.tsx`

Likely new language instrumentation files:

- `packages/api-electron/src/execution/instrumentation/node-transform.js`
- `packages/api-electron/src/execution/instrumentation/python-transform.py`

This list is intentional guidance, not a locked file map. The constraint is simpler: do not create a second desktop host or a second editor app.

## Shared Runtime Model

The IDE should support three runtime modes:

1. Auto-start
2. Connect to existing runtime
3. Manual restart/stop from the IDE

Shared runtime session state:

- runtime mode
- gRPC listen address
- admin listen address
- health status
- dashboard URL
- session viewer URL
- managed process id when auto-started

The embedded dashboard must point to the exact runtime instance the example runner is using.

## Shared Execution Model

Each run should create an `ExecutionSession` with:

- session id
- language
- source file path
- working directory
- runtime mode
- process id
- run state (`idle`, `starting`, `running`, `paused`, `completed`, `failed`)
- current line
- current statement id
- stdout buffer
- stderr buffer
- step events
- runtime call events

Shared event types:

- `session:starting`
- `session:started`
- `session:completed`
- `session:failed`
- `stdout`
- `stderr`
- `step:start`
- `step:end`
- `step:pause`
- `step:resume`
- `step:error`
- `runtime:call`
- `runtime:result`

The editor renderer should consume only this shared event model. It should not contain Node.js-specific or Python-specific stepping logic.

## Instrumentation Strategy

Recommendation:
- use source instrumentation for phase 1 stepping, not language-native debugger integrations

Why:

- we only need to support the example corpus first
- we want consistent stepping semantics across both languages
- debugger integration in both languages is a separate project with more platform drift

### Node.js

Implementation approach:

- parse `.mjs` examples
- wrap executable statements with a step hook
- suspend before each statement until the IDE sends continue/step

### Python

Implementation approach:

- parse `.py` examples with `ast`
- inject the same step hook before executable statements
- suspend on the same shared pause/step contract

Constraint:

- phase 1 instrumentation only needs to support the current example patterns
- broadening to arbitrary user-authored scripts happens after the event model and runner architecture are stable

## Milestones

### Milestone 0: Desktop Shell Convergence

Objective:
- turn the existing Electron shell and editor shell into one app boundary

Status:
- Implemented (baseline Electron host + editor renderer wiring)

Implementation tasks:

- load the `packages/editor/` renderer inside Electron
- keep runtime management in the Electron main process
- expose a minimal preload bridge for:
  - runtime status
  - example discovery
  - run
  - stop
  - pause
  - resume
  - step

Dependencies:

- current Electron host in `packages/api-electron/`
- current Next.js editor in `packages/editor/`

Validation:

- the editor renderer opens inside Electron
- the app can still open the live dashboard
- runtime health can be requested from the renderer

Exit criteria:

- one desktop shell runs the editor renderer and can still manage the runtime

### Milestone 1: Example Discovery And Execution

Objective:
- run the existing Node.js and Python examples from the editor

Status:
- Implemented (baseline discovery, source viewing, run, stop, and output streaming)

Implementation tasks:

- discover examples from:
  - `packages/client-node/examples/`
  - `packages/client-python/examples/`
- add an example tree or picker in the renderer
- start Node.js or Python child processes from the Electron host
- stage working directories and assets correctly
- stream stdout/stderr and exit status to the renderer

Validation:

- run `click` and `find` examples from both language sets
- confirm output appears in the IDE
- confirm runtime auto-start and connect modes both work

Exit criteria:

- both language example sets are runnable from the IDE without stepping

### Milestone 2: Shared Execution Session And UI

Objective:
- make execution state visible and consistent before adding stepping

Status:
- Implemented (shared execution sessions, history, metadata, and timeline UI)

Implementation tasks:

- add `ExecutionSession` model
- add renderer output panel
- add run-state indicator
- add session history/timeline container
- show current file and run metadata

Validation:

- session state remains coherent across multiple runs
- output panel resets correctly per run
- failing examples are visible as failed sessions, not silent exits

Exit criteria:

- the IDE can show execution state reliably for both languages

### Milestone 3: Node.js Step Instrumentation

Objective:
- make Node.js examples pausable and step-able line by line

Status:
- Implemented (Node.js line stepping, pause or resume controls, and current-line highlighting)

Implementation tasks:

- add Node.js AST transform
- inject step hooks before statements
- add pause, resume, and step-over controls
- highlight the current line in the code pane
- append step events to the timeline

Validation:

- `click.mjs`, `find.mjs`, and one workflow example can be stepped end to end
- pause and step do not deadlock the process
- current-line highlighting tracks the emitted step location

Exit criteria:

- Node.js examples can be stepped line by line from the IDE

### Milestone 4: Python Step Instrumentation

Objective:
- match Node.js stepping behavior for Python examples

Status:
- Implemented (Python line stepping on the shared session and control model)

Implementation tasks:

- add Python `ast` transform
- inject step hooks
- normalize Python step events to the same execution model
- reuse the same renderer controls and timeline

Validation:

- `click.py`, `find.py`, and one workflow example can be stepped end to end
- stepping semantics match the Node.js UX closely enough that the UI does not branch

Exit criteria:

- Python examples can be stepped line by line from the IDE

### Milestone 5: Embedded Dashboard Workspace

Objective:
- show the same runtime dashboards inside the IDE

Status:
- Implemented (embedded live dashboard, session viewer, and admin-surface panes in the editor workspace)

Implementation tasks:

- add split-pane or tabbed dashboard area in the renderer layout
- load:
  - live dashboard
  - session viewer
  - optional health/metrics panes
- bind the dashboard views to the same runtime session used by the current run

Validation:

- dashboard view reflects the managed runtime
- session viewer is reachable from the same IDE workspace
- restarting runtime updates the embedded views correctly

Exit criteria:

- the IDE shows the existing runtime dashboard and session viewer without rebuilding them

### Milestone 6: Runtime-Aware Stepping UX

Objective:
- make stepping useful for visual automation debugging

Status:
- Implemented (runtime-call tracing, pause-on-call / run-to-line controls, and call-aware session diagnostics in the editor)

Implementation tasks:

- correlate step events with runtime calls
- show runtime call boundaries in the timeline
- attach screenshots, match metadata, or error context when available
- add `run to line` and `continue to next runtime call`

Validation:

- users can identify which source line triggered the runtime action
- a failed runtime call is visible in both the output panel and the step timeline

Exit criteria:

- the IDE is useful for debugging visual automation flows, not just pausing generic code

### Milestone 7: Authoring And Hardening

Objective:
- move from example runner to usable local IDE

Status:
- Implemented (workspace authoring flows, asset awareness, runtime hardening, and host-side regression coverage)

Implementation tasks:

- support save, clone, rename, and new-file flows
- support assets next to scripts
- add crash recovery for runner processes
- add port-conflict handling for the runtime
- persist recent workspaces and panel layout
- add smoke coverage for Node.js and Python execution paths
- add regression tests for both instrumentation paths

Validation:

- edited files survive restart
- runner crashes surface clearly and recover cleanly
- repeated runs across both languages remain stable

Exit criteria:

- the IDE is stable enough for regular local use, not just demos

## Cross-Cutting Workstreams

### Workstream A: IPC Contract

Needed early because every milestone depends on it.

Minimum IPC surface:

- `runtime.getStatus`
- `runtime.start`
- `runtime.stop`
- `runtime.restart`
- `examples.list`
- `execution.run`
- `execution.stop`
- `execution.pause`
- `execution.resume`
- `execution.step`
- `execution.subscribe`

### Workstream B: Editor UX

Needed in parallel with execution milestones.

Minimum UX pieces:

- example tree
- code viewer/editor
- run controls
- output panel
- timeline
- runtime badge
- dashboard tabs or split pane

### Workstream C: Verification

Needed before hardening.

Minimum test layers:

- renderer smoke tests
- Electron host smoke tests
- example execution tests for Node.js
- example execution tests for Python
- instrumentation regression fixtures for both languages

## Validation Gates

Do not advance beyond each gate until the previous one is green.

Gate 1:

- Electron hosts the editor renderer
- runtime status is visible from the renderer

Gate 2:

- Node.js and Python examples both run from the IDE

Gate 3:

- Node.js stepping works on the current example corpus

Gate 4:

- Python stepping works on the current example corpus

Gate 5:

- embedded dashboard and session viewer reflect the same runtime instance as the current run

Gate 6:

- repeated run/step/restart flows remain stable across both languages

## Risks And Mitigations

- Risk: source instrumentation may break on edge-case syntax.
  - Mitigation: support the current example corpus first and add fixture coverage before broadening scope.
- Risk: the Electron host and Next renderer can drift into duplicated state ownership.
  - Mitigation: keep all process and filesystem ownership in Electron; keep the renderer event-driven.
- Risk: dashboard embedding can fail silently when runtime ports change or startup races occur.
  - Mitigation: make runtime status explicit and block dashboard tabs on health readiness.
- Risk: Python environment differences can make runner behavior flaky.
  - Mitigation: execute against the known package example flows first and normalize environment setup in the host.

## Definition Of Done

The IDE work is done for the initial release when:

- Node.js examples run and step in the IDE
- Python examples run and step in the IDE
- the current line is highlighted during pause/step
- stdout, stderr, and runtime status are visible in one workspace
- the live dashboard and session viewer are embedded from the existing runtime
- runtime restart/stop is controlled from the IDE
- execution and instrumentation have automated smoke coverage

## Recommended Build Order

If we want the shortest path to a useful result, build in this order:

1. Electron + editor convergence
2. Example discovery and execution
3. Shared session model and output UI
4. Node.js stepping
5. Python stepping
6. Embedded dashboard workspace
7. Runtime-aware debugging polish
8. Authoring and hardening

## Repository References

- editor renderer: `packages/editor/`
- current editor shell: `packages/editor/app/page.tsx`
- desktop host: `packages/api-electron/`
- runtime host logic: `packages/api-electron/src/main.js`
- Node.js examples: `packages/client-node/examples/`
- Python examples: `packages/client-python/examples/`
- dashboard guide: `docs/getting-started/dashboard.md`
- client delivery context: `docs/strategy/client-strategy.md`
