# SikuliGO [![GoLang Tests](https://github.com/smysnk/SikuliGO/actions/workflows/go-test.yml/badge.svg)](https://github.com/smysnk/SikuliGO/actions/workflows/go-test.yml)

![SikuliX Logo](docs/images/logo.png)

This repository houses a GoLang implementation of Sikuli visual automation.

## Project Intent

- Build a feature-complete GoLang port of the core Sikuli API concepts.
- Preserve behavioral parity where it matters (image matching, regions, patterns, finder semantics).
- Provide a modern, testable architecture with explicit contracts and deterministic matching behavior.
- Establish a maintainable foundation for cross-platform automation features.

## Current Focus

| Roadmap Item | Scope | Current Status |
|---|---|---|
| Core API scaffolding | Public SikuliGo API surface and parity-facing core objects | ✅ Completed (baseline + extensions) |
| Matching engine and parity harness | Deterministic matcher behavior, golden corpus, backend conformance tests | ✅ Completed (baseline + extensions) |
| API parity surface expansion | Additional parity helpers and compatibility APIs | 🟡 Planned / In progress |
| OCR and text-search parity | OCR contracts, finder/region text flows, optional backend integration | ✅ Completed (pinned gosseract module integration) |
| Input automation and hotkey parity | Input controller contracts, request validation, backend protocol scaffold | 🟡 In progress (concrete `darwin` backend; non-`darwin` fallback unsupported) |
| Observe/event subsystem parity | Observer contracts, request validation, backend protocol scaffold | ✅ Completed (concrete deterministic polling backend) |
| App/window/process control parity | App/window contracts, request validation, backend protocol scaffold | ✅ Completed (concrete `darwin`/`linux`/`windows` backends) |
| Cross-platform backend hardening | Platform integration hardening and backend portability | 🟡 Planned |

## Port Strategy

- [Docs Home](https://smysnk.github.io/SikuliGO/)
- [Consolidated Strategy](https://smysnk.github.io/SikuliGO/port-strategy)
- [API Reference](https://smysnk.github.io/SikuliGO/api/)
- [Defaults Table](https://smysnk.github.io/SikuliGO/default-behavior-table)

Integration & Implementation:
- [OCR](https://smysnk.github.io/SikuliGO/ocr-integration)
- [Input Automation](https://smysnk.github.io/SikuliGO/input-automation)
- [Observe Events](https://smysnk.github.io/SikuliGO/observe-events)
- [App Control](https://smysnk.github.io/SikuliGO/app-control)
- [Backend Capability Matrix](https://smysnk.github.io/SikuliGO/backend-capability-matrix)

API reference pages are generated from source with `./scripts/generate-api-docs.sh` and validated with `./scripts/check-api-docs.sh`.

## Repository Layout

- [`pkg`](pkg) : public GoLang API packages
- [`internal`](internal) : internal GoLang implementation packages
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
- [Sikuli (Software) - Wikipedia](https://de.wikipedia.org/wiki/Sikuli_(Software))
- [Original Sikuli Github](https://github.com/sikuli/sikuli)
