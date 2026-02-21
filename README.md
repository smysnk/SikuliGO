# SikuliGO [![GoLang Tests](https://github.com/smysnk/SikuliGO/actions/workflows/go-test.yml/badge.svg)](https://github.com/smysnk/SikuliGO/actions/workflows/go-test.yml)

![SikuliX Logo](docs/images/logo.png)

This repository houses a GoLang implementation of Sikuli visual automation.

## Project Intent

- Build a feature-complete GoLang port of the core Sikuli API concepts.
- Preserve behavioral parity where it matters (image matching, regions, patterns, finder semantics).
- Provide a modern, testable architecture with explicit contracts and deterministic matching behavior.
- Establish a maintainable foundation for cross-platform automation features.

## Current Focus

- Core API scaffolding in `pkg/sikuli`
- Matching engine implementation in `internal/cv`
- Golden/parity test harness in `internal/testharness`

## Port Strategy

- [Docs Home](https://smysnk.github.io/SikuliGO/)
- [Consolidated Strategy](https://smysnk.github.io/SikuliGO/port-strategy)
- [Locked Architecture](https://smysnk.github.io/SikuliGO/architecture-lock)
- [API Freeze](https://smysnk.github.io/SikuliGO/api-signature-freeze)
- [Defaults Table](https://smysnk.github.io/SikuliGO/default-behavior-table)
- [CI Test Reporting](https://smysnk.github.io/SikuliGO/ci-test-reporting)

## Repository Layout

- [`pkg`](pkg) : public GoLang API packages
- [`internal`](internal) : internal GoLang implementation packages
- [`docs`](docs) : documentation and assets (including logo)
- [`legacy`](legacy) : previous Java-era project directories retained for reference

## Getting Started

```bash
go mod tidy
go test ./...
```

## Project History and Credits

Sikuli started in 2009 as an open-source research effort at the MIT User Interface Design Group, led by **Tsung-Hsiang Chang** and **Tom Yeh**, with early development connected to **Prof. Rob Miller**'s work at **MIT CSAIL**. The project introduced a practical idea that was unusual at the time: instead of relying on internal application APIs, users could automate **Graphical User Interfaces (GUI)** by teaching scripts what to click through screenshots of buttons, icons, and other visual elements. Even the name reflected that vision, drawing from the Huichol concept of the "**God's Eye**," a symbol of seeing and understanding what is otherwise hidden.

In 2012, after the original creators moved on, the project's active development continued under **RaiMan** and evolved into **SikuliX**. That branch carried the platform forward for real-world desktop and web automation, using scripting ecosystems such as **Jython/Python**, **Java**, and **Ruby**, and refining image-based interaction workflows over time. Because this style of automation simulates real **mouse** and **keyboard** behavior, it has always worked best in environments with an active graphical session rather than truly **headless** execution.

The GoLang port in this repository began in **2026**. It stands on the work of the original Sikuli authors, **RaiMan**, and the broader contributor community that kept visual automation practical and accessible over the years.

## Sikuli References

- [SikuliX Official Site](https://sikulix.github.io/)
- [Sikuli (Software) - Wikipedia](https://de.wikipedia.org/wiki/Sikuli_(Software))
