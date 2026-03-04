# Behavioral Differences (Java SikuliX vs SikuliGO)

This page documents intentional, user-visible differences between Java SikuliX behavior and the GoLang port.

## Runtime and Process Model

- SikuliGO uses an API/server runtime model with gRPC transport; Java SikuliX is traditionally in-process scripting.
- Screen capture and input actions are executed by the API process, not by language clients.

## Matching and Engine Selection

- SikuliGO exposes explicit matcher engine selection (`template`, `orb`, `akaze`, `brisk`, `kaze`, `sift`, `hybrid`).
- Java SikuliX engine behavior is less explicitly surfaced as a transport-level option.

## OCR Backend Enablement

- OCR in SikuliGO requires builds with OCR backend tags and native dependencies.
- Java SikuliX OCR packaging/runtime defaults differ by distribution.

## Input Backend Dependencies

- Platform input backends in SikuliGO depend on OS tools/libraries (`cliclick`, `xdotool`, PowerShell path/runtime).
- SikuliGO now performs startup/runtime dependency checks and surfaces actionable warnings.

## Constructor and Session Model

- SikuliGO clients use explicit `auto/connect/spawn` constructor patterns with API session tracking.
- Java SikuliX often assumes direct scripting session context.

## Documentation Contract

- Any parity-sensitive behavior delta must be documented here and cross-linked in `parity-gaps.md` or mapping notes.
