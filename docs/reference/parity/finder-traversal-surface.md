# Finder Traversal Surface

Phase 5 adds a compatibility traversal model to `packages/api/pkg/sikuli/finder.go` without removing the existing Go value/slice API.

## What Changed

`Finder` now supports two additive traversal-preparation methods:

- `Iterate(pattern *Pattern) error`
- `IterateAll(pattern *Pattern) error`

These methods:

- populate the finder's internal traversal state
- do not return `ErrFindFailed` on a miss
- leave `HasNext()` as the presence check, matching the SikuliX mental model

The traversal methods operate over the same last-result cache used by:

- `HasNext() bool`
- `Next() (Match, bool)`
- `Reset()`
- `Destroy()`
- `LastMatches() []Match`

## Compatibility Contract

The Go port intentionally keeps both models:

### Existing Go-oriented search

- `Find(pattern)` returns `ErrFindFailed` on a miss
- `FindAll(pattern)` returns `[]Match`

### Additive SikuliX-style traversal

- `Iterate(pattern)` prepares a one-match traversal without raising `ErrFindFailed` on a miss
- `IterateAll(pattern)` prepares an all-match traversal without requiring callers to consume a slice immediately
- `HasNext()` and `Next()` advance through the prepared results

## LastMatches Coherence

`LastMatches()` continues returning the full most-recent match set even after `Next()` advances the iterator.

This is a deliberate Go-compatibility choice:

- SikuliX-style ports can use `HasNext()` and `Next()`
- existing Go callers keep stable slice access to the most recent results

`Destroy()` clears both the traversal cursor and the last-match cache. `Reset()` rewinds traversal over the current last-match cache.

## Scope

This phase does not change the runtime architecture or move traversal semantics into gRPC. The compatibility iterator is an additive `pkg/sikuli` layer on top of the existing search implementation.
