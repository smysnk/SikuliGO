# ORB Search Delivery Phases

This document defines the phased delivery plan for adding ORB-based image search as an alternative matcher in SikuliGO.

## Goals

- Add ORB as a first-class matcher backend.
- Keep backward compatibility with current template matching behavior.
- Support engine selection both per script session and per call (ad-hoc override).
- Measure quality and performance before expanding default usage.

## Implementation Status (Current)

- Completed: matcher abstraction with engine routing (`template`, `orb`, `hybrid`).
- Completed: ORB matcher backend wired for OpenCV builds.
- Completed: hybrid matcher backend (`template` primary + ORB fallback).
- Completed: typed proto matcher-engine fields for engine selection.
- Completed: Node and Python client support for session-level defaults and per-call overrides.
- Pending: expanded benchmark corpus and default-engine rollout policy updates.

## Scope

In scope:

- matcher abstraction and ORB backend implementation
- gRPC/client API options for engine selection
- hybrid fallback logic (`template` + `orb`)
- telemetry, conformance tests, and rollout controls

Out of scope:

- deprecating template matching
- changing existing public method names
- forcing ORB as universal default before conformance data

## Engine Selection Model

Supported engine values:

- `template`
- `orb`
- `hybrid`

Selection precedence:

1. per-call override (ad-hoc)
2. client session default
3. server default

Enablement paths:

- Per script session: set default engine when creating the client (`Screen(engine="orb")`, `Sikuli(engine="hybrid")`).
- Ad-hoc per call: override on individual calls (`find(..., engine="orb")`, `click(..., engine="hybrid")`).
- Transport field path:
  - `FindRequest.matcher_engine`
  - `ScreenQueryOptions.matcher_engine`

## Phase 1: Core Matcher Abstraction

Objective:

- Introduce a matcher interface and keep current behavior as the default implementation.

Deliverables:

- `Matcher` interface in core search path
- `TemplateMatcher` wrapper for current behavior
- factory/registry for backend selection

Entry criteria:

- existing template matching tests green

Exit criteria:

- no functional regression in current template behavior
- all matcher calls routed through abstraction

## Phase 2: ORB Backend (Feature-Flagged)

Objective:

- Implement ORB matching without changing default production behavior.

Deliverables:

- ORB pipeline (keypoint detect, descriptor match, ratio test, RANSAC/homography)
- normalized `Match` output mapping (target coordinates + score normalization)
- server-side feature flag gate

Entry criteria:

- phase 1 abstraction complete

Exit criteria:

- ORB backend compiles and runs across supported platforms
- baseline ORB smoke tests pass
- flag-off keeps behavior unchanged

## Phase 3: Session + Ad-Hoc Engine Controls

Objective:

- Enable engine selection per client session and per call.

Deliverables:

- proto/API options for matcher engine
- server logic for precedence rules
- Node/Python client constructor defaults + per-method overrides

Entry criteria:

- ORB backend available behind flag

Exit criteria:

- client can set engine once for session
- client can override engine per call
- server records selected engine for each interaction

## Phase 4: Hybrid Mode and Reliability Guardrails

Objective:

- Improve reliability by combining deterministic template matching with ORB resilience.

Deliverables:

- `hybrid` engine mode
- configurable fallback order and thresholds
- descriptor/keypoint caching for repeated patterns
- failure reason tagging (insufficient keypoints, low inlier ratio, etc.)

Entry criteria:

- per-session and ad-hoc controls complete

Exit criteria:

- hybrid fallback produces lower false-negative rate on benchmark set
- no increase in critical false positives beyond agreed threshold

## Phase 5: Conformance, Performance, and Rollout

Objective:

- Validate quality and speed, then define safe rollout defaults.

Deliverables:

- benchmark corpus across real UI assets
- conformance matrix (`template` vs `orb` vs `hybrid`)
- docs/examples for each engine mode
- rollout policy (default engine per environment/profile)

Entry criteria:

- telemetry and hybrid mode available

Exit criteria:

- performance targets documented and met for target workloads
- rollout recommendation approved (`template` default or `hybrid` default)
- release notes and migration notes published

## Practical Considerations

- ORB generally improves robustness for scale/rotation variance.
- ORB can underperform on low-texture UI targets.
- Template matching remains important for deterministic pixel-stable UI.
- Hybrid mode is typically the safest first default for broad compatibility.
- Always capture per-engine metrics before changing defaults globally.

## Recommended Initial Rollout

- Server default: `hybrid`
- Opt-in pilots: `hybrid` for selected scripts/suites
- Advanced opt-in: `orb` per session or per call where template instability is known
- Revisit default after phase 5 benchmark evidence
