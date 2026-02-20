# Architecture Lock (for Workstreams 1 and 2)

The implementation is pinned to the following package structure:

- `pkg/sikuli`: public API layer and compatibility-facing types.
- `internal/core`: matcher contracts (`SearchRequest`, `Matcher`, candidate types) and shared image helpers.
- `internal/cv`: concrete matcher implementation (NCC template matching + resize + mask support).
- `internal/testharness`: golden corpus loader and parity comparator.

The matcher contract boundary is intentionally narrow:

```go
type Matcher interface {
  Find(req SearchRequest) ([]MatchCandidate, error)
}
```

This keeps `pkg/sikuli.Finder` stable while allowing backend swaps
(`gocv` implementation, GPU implementation, etc.) without API breakage.

This architecture is the baseline for the GoLang port.
