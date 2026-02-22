# SikuliGO [![GoLang Tests](https://github.com/smysnk/SikuliGO/actions/workflows/go-test.yml/badge.svg)](https://github.com/smysnk/SikuliGO/actions/workflows/go-test.yml)

![SikuliX Logo](docs/images/logo.png)

Sikuli is an open-source tool for automating anything visible on a computer screen using image recognition. Instead of relying on internal source code or object IDs, it identifies and interacts with graphical user interface (GUI) components (buttons, text boxes, etc.) by using screenshots. **This repo houses a GoLang port of the original concept.**

## Project Intent

- Build a feature-complete GoLang port of the core [Sikuli](https://sikulix.github.io/) concepts.
- Preserve behavioral parity (image matching, regions, patterns, finder semantics).
- Provide a modern, testable architecture with explicit contracts and deterministic matching behavior.
- Establish a maintainable foundation for cross-platform automation features.

## Available Clients

| Client |  | Notes |
| :---  | --- | :---  |
| [Python](https://pypi.org/project/sikuligo/)  | ✅ | Implemented |
| [Node](https://www.npmjs.com/package/@sikuligo/sikuligo)  | ✅ | Implemented |
| Lua  | ✅ | Implemented |
| Robot Framework | 🟡 | Planned |
| Web IDE | 🟡 | Planned |

## Python
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

## Node
```ts
import { Sikuli } from "@sikuligo/sikuligo";

const bot = await Sikuli.launch();
await bot.click({ x: 300, y: 220 });
await bot.typeText("hello");
await bot.hotkey(["cmd", "enter"]);
await bot.close();
```

## Example: Dashboard

![SikuliGO Dashboard Demo](docs/images/dashboard.png)


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

## Integration & Implementation
- [API Reference](https://smysnk.github.io/SikuliGO/api/)
- [OCR](https://smysnk.github.io/SikuliGO/ocr-integration)
- [Input Automation](https://smysnk.github.io/SikuliGO/input-automation)
- [Observe Events](https://smysnk.github.io/SikuliGO/observe-events)
- [App Control](https://smysnk.github.io/SikuliGO/app-control)
- [Defaults Table](https://smysnk.github.io/SikuliGO/default-behavior-table)
- [Backend Capability Matrix](https://smysnk.github.io/SikuliGO/backend-capability-matrix)
- [Node Package User Flow](https://smysnk.github.io/SikuliGO/node-package-user-flow)

## Repository Layout

- [`pkg`](pkg) : public GoLang API packages
- [`internal`](internal) : internal GoLang implementation packages
- [`clients`](clients) : language client SDKs and packaging artifacts
- [`docs`](docs) : documentation and assets
- [`legacy`](legacy) : previous Java-era project directories retained for reference

## Getting Started

Requires GoLang `1.24+`.

```bash
go mod tidy
go test ./...
```

Optional OCR backend (gosseract):

```bash
go test -tags gosseract ./...
```

Tagged OCR builds require native Tesseract + Leptonica runtime libraries and installed language data.
See [OCR](https://smysnk.github.io/SikuliGO/ocr-integration) for full macOS/Homebrew setup and troubleshooting steps.

## Project History and Credits

Sikuli started in 2009 as an open-source research effort at the MIT User Interface Design Group, led by **Tsung-Hsiang Chang** and **Tom Yeh**, with early development connected to **Prof. Rob Miller**'s work at **MIT CSAIL**. The project introduced a practical idea that was unusual at the time: instead of relying on internal application APIs, users could automate **Graphical User Interfaces (GUI)** by teaching scripts what to click through screenshots of buttons, icons, and other visual elements. Even the name reflected that vision, drawing from the Huichol concept of the "**God's Eye**," a symbol of seeing and understanding what is otherwise hidden.

In 2012, after the original creators moved on, the project's active development continued under **RaiMan** and evolved into **SikuliX**. That branch carried the platform forward for real-world desktop and web automation, using scripting ecosystems such as **Jython/Python**, **Java**, and **Ruby**, and refining image-based interaction workflows over time. Because this style of automation simulates real **mouse** and **keyboard** behavior, it has always worked best in environments with an active graphical session rather than truly **headless** execution.

The GoLang port in this repository began in **2026**. It stands on the work of the original Sikuli authors, **RaiMan**, and the broader contributor community that kept visual automation practical and accessible over the years.

## Sikuli References

- [SikuliX Official Site](https://sikulix.github.io/)
- [Wikipedia](https://de.wikipedia.org/wiki/Sikuli_(Software))
- [Original Sikuli Github](https://github.com/sikuli/sikuli)
- [Sikuli Framework](https://github.com/smysnk/sikuli-framework) = Sikuli + Robot Framework
