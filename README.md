# SikuliGO [![GoLang Tests](https://github.com/smysnk/SikuliGO/actions/workflows/go-test.yml/badge.svg)](https://github.com/smysnk/SikuliGO/actions/workflows/go-test.yml) [![Testing Gate](https://img.shields.io/github/actions/workflow/status/smysnk/SikuliGO/go-test.yml?branch=master&label=testing%20gate)](https://github.com/smysnk/SikuliGO/actions/workflows/go-test.yml) [![PyPI version](https://img.shields.io/pypi/v/sikuligo)](https://pypi.org/project/sikuligo/) [![npm version](https://img.shields.io/npm/v/%40sikuligo%2Fsikuligo)](https://www.npmjs.com/package/@sikuligo/sikuligo)

![SikuliX Logo](docs/images/logo.png)

Sikuli is an open-source tool for automating anything visible on a computer screen using image recognition. Instead of relying on internal source code or object IDs, it identifies and interacts with graphical user interface (GUI) components (buttons, text boxes, etc.) by using screenshots. **This repo houses a GoLang port of the original concept.**

## Project Intent

- Build a feature-complete GoLang port of the core [Sikuli](https://sikulix.github.io/) concepts.
- Preserve behavioral parity (image matching, regions, patterns, finder semantics).
- Provide a modern, testable architecture with explicit contracts and deterministic matching behavior.
- Establish a maintainable foundation for cross-platform automation features.

## Examples

### Node.js

Install the Node client package:

```bash
yarn dlx @sikuligo/sikuligo init:js-examples
cd sikuligo-demo
yarn node examples/click.mjs
```

Runs:
```js
import { Screen, Pattern } from "@sikuligo/sikuligo";

const screen = await Screen();
try {
  const match = await screen.click(Pattern("assets/pattern.png").exact());
  console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
} finally {
  await screen.close();
}
```

### Python

Scaffold and run Python examples:

```bash
pipx run sikuligo init:py-examples
cd sikuligo-demo
python3 examples/click.py
```

Runs:
```python
from sikuligo import Pattern, Screen

screen = Screen()
try:
    match = screen.click(Pattern("assets/pattern.png").exact())
    print(f"clicked match target at ({match.target_x}, {match.target_y})")
finally:
    screen.close()
```

## API Dashboard

![SikuliGO Dashboard Demo](docs/images/dashboard.png)

## Install API Binary On PATH

Install via Yarn (Node ecosystem):

```bash
yarn dlx @sikuligo/sikuligo install-binary
```

Install via Python:

```bash
pipx run sikuligo install-binary
```

Both commands create `~/.local/bin` if needed, copy all available `sikuli*` runtimes, and can prompt to add PATH to `~/.zshrc` or `~/.bash_profile`.
If PATH is updated, reload with `source ~/.zshrc` or `source ~/.bash_profile`.

## Available Clients

| Client |  | Notes |
| :---  | --- | :---  |
| [Python](https://pypi.org/project/sikuligo/)  | ✅ | Implemented |
| [Node](https://www.npmjs.com/package/@sikuligo/sikuligo)  | ✅ | Implemented |
| Lua  | ✅ | Implemented |
| Robot Framework | 🟡 | Planned |
| Web IDE | 🟡 | Planned |


## Current Focus

| Roadmap Item | Scope |  |
| :---  | :---  |---|
| Core API scaffolding | Public SikuliGo API surface and parity-facing core objects | ✅ |
| Matching engine and parity harness | Deterministic matcher behavior, golden corpus, backend conformance tests | ✅ |
| API parity surface expansion | Additional parity helpers and compatibility APIs | ✅ |
| Protocol completeness hardening | Alternate matcher backend + cross-backend conformance rules | ✅ |
| OCR and text-search parity | OCR contracts, finder/region text flows, optional backend integration | ✅ |
| Input automation and hotkey parity | Input controller contracts, request validation, backend protocol scaffold | 🟡 |
| Observe/event subsystem parity | Observer contracts, request validation, backend protocol scaffold | ✅ |
| App/window/process control parity | App/window contracts, request validation, backend protocol scaffold | ✅ |
| Cross-platform backend hardening | Platform integration hardening and backend portability | 🟡 |

# Docs
- [Docs Home](https://smysnk.github.io/SikuliGO/)

## Strategy
- [Port](https://smysnk.github.io/SikuliGO/port-strategy)
- [gRPC](https://smysnk.github.io/SikuliGO/grpc-strategy)
- [Client](https://smysnk.github.io/SikuliGO/client-strategy)
- [ORB Search Delivery](https://smysnk.github.io/SikuliGO/orb-search-delivery)

## Integration & Implementation
- [API Reference](https://smysnk.github.io/SikuliGO/api/)
- [OCR](https://smysnk.github.io/SikuliGO/ocr-integration)
- [OpenCV](https://smysnk.github.io/SikuliGO/opencv-integration)
- [Input Automation](https://smysnk.github.io/SikuliGO/input-automation)
- [Observe Events](https://smysnk.github.io/SikuliGO/observe-events)
- [App Control](https://smysnk.github.io/SikuliGO/app-control)
- [Defaults Table](https://smysnk.github.io/SikuliGO/default-behavior-table)
- [Backend Capability Matrix](https://smysnk.github.io/SikuliGO/backend-capability-matrix)
- [Node Package User Flow](https://smysnk.github.io/SikuliGO/node-package-user-flow)

## Repository Layout

- [`pkg`](pkg) : public GoLang API packages
- [`packages/api`](packages/api) : GoLang API module (`cmd`, `internal`, `pkg`, `proto`)
- [`packages/editor`](packages/editor) : Next.js editor app
- [`packages/api-electron`](packages/api-electron) : macOS Electron API + dashboard/session viewer app
- [`clients`](clients) : language client SDKs and packaging artifacts
- [`docs`](docs) : documentation and assets
- [`legacy`](legacy) : previous Java-era project directories retained for reference

## Install and Build

Build from source:

- [Build From Source](docs/build-from-source.md)

## Project History and Credits

Sikuli started in 2009 as an open-source research effort at the MIT User Interface Design Group, led by **Tsung-Hsiang Chang** and **Tom Yeh**, with early development connected to **Prof. Rob Miller**'s work at **MIT CSAIL**. The project introduced a practical idea that was unusual at the time: instead of relying on internal application APIs, users could automate **Graphical User Interfaces (GUI)** by teaching scripts what to click through screenshots of buttons, icons, and other visual elements. Even the name reflected that vision, drawing from the Huichol concept of the "**God's Eye**," a symbol of seeing and understanding what is otherwise hidden.

In 2012, after the original creators moved on, the project's active development continued under **RaiMan** and evolved into **SikuliX**. That branch carried the platform forward for real-world desktop and web automation, using scripting ecosystems such as **Jython/Python**, **Java**, and **Ruby**, and refining image-based interaction workflows over time. Because this style of automation simulates real **mouse** and **keyboard** behavior, it has always worked best in environments with an active graphical session rather than truly **headless** execution.

The GoLang port in this repository began in **2026**. It stands on the work of the original Sikuli authors, **RaiMan**, and the broader contributor community that kept visual automation practical and accessible over the years.

## Sikuli References

- [SikuliX Official Site](https://sikulix.github.io/)
- [Wikipedia](https://de.wikipedia.org/wiki/Sikuli_(Software))
- [Original Sikuli Github](https://github.com/sikuli/sikuli)
- [Sikuli Framework](https://github.com/smysnk/sikuli-framework) = Sikuli + Robot Framework
