# SikuliGO Mission

SikuliGO exists to provide a production-grade GoLang port of SikuliX visual automation semantics.

## Why This Port Exists

- Keep the core Sikuli model (Pattern, Region, Finder, Screen) usable in modern service/runtime architectures.
- Preserve behavioral parity where it matters most: image matching behavior, wait/exists semantics, targeting, and workflow ergonomics.
- Expose the same core capabilities through a stable gRPC surface and language clients.

## Core Outcomes

- Deterministic matching contracts with measurable behavior across engines.
- Cross-platform automation backends behind stable interfaces.
- Explicit parity mapping between Java/SikuliX concepts and Go/gRPC/client APIs.
- Test and documentation gates that prevent silent drift.

## Non-Goals

- Byte-for-byte API equivalence with Java internals.
- Hidden behavior differences without documentation.
- One-off client APIs that bypass shared contracts.

## Engineering Principles

- Stable public interfaces first, backend implementations behind protocol boundaries.
- Additive evolution and clear compatibility notes.
- Behavior-driven tests for parity-sensitive features.
- Documentation generated from source where possible; curated where needed.
