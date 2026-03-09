# Multi-Target Search Surface

Phase 6 adds the SikuliX-style multi-pattern helper family to the Go API without changing the underlying runtime architecture.

## Scope

Added public helpers:
- `(*Finder).FindAnyList(patterns)`
- `(*Finder).FindBestList(patterns)`
- `(*Finder).WaitAnyList(patterns, timeout)`
- `(*Finder).WaitBestList(patterns, timeout)`
- `(*Region).FindAnyList(source, patterns)`
- `(*Region).FindBestList(source, patterns)`
- `(*Region).WaitAnyList(source, patterns, timeout)`
- `(*Region).WaitBestList(source, patterns, timeout)`
- `LiveRegion.FindAnyList(patterns)`
- `LiveRegion.FindBestList(patterns)`
- `LiveRegion.WaitAnyList(patterns, timeout)`
- `LiveRegion.WaitBestList(patterns, timeout)`
- `Screen.FindAnyList(patterns)`
- `Screen.FindBestList(patterns)`
- `Screen.WaitAnyList(patterns, timeout)`
- `Screen.WaitBestList(patterns, timeout)`
- live `Match` equivalents through the existing region-like match surface

No new RPCs were added in this phase.

## Execution Model

Image-backed helpers:
- evaluate all requested patterns against the same source image
- reuse the caller's configured matcher backend on `Finder`

Live-screen helpers:
- capture one screen image per helper call
- run every requested pattern against that same captured image
- on `WaitAnyList` and `WaitBestList`, poll by taking one capture per poll cycle, then evaluate all patterns against that capture

This keeps semantics stable and avoids screenshot drift across patterns inside one helper call.

## Semantics

`FindAnyList(patterns)`:
- returns one best match per matched pattern
- returns `ErrFindFailed` if none of the patterns match
- returned matches preserve input-pattern identity in `Match.Index`
- result ordering follows input-pattern order

`FindBestList(patterns)`:
- returns the best match across the matched pattern list
- returns `ErrFindFailed` if none of the patterns match
- tie-breaking is deterministic:
  1. higher `Score`
  2. lower input pattern index
  3. smaller `Y`
  4. smaller `X`
  5. smaller `W`
  6. smaller `H`

`WaitAnyList(patterns, timeout)`:
- polls until at least one pattern matches
- returns one best match per matched pattern from the successful poll
- returns `ErrTimeout` if the wait budget is exhausted

`WaitBestList(patterns, timeout)`:
- polls until at least one pattern matches
- returns the deterministic best match from the successful poll
- returns `ErrTimeout` if the wait budget is exhausted

## Match Index Contract

For this helper family, `Match.Index` is the zero-based input pattern index, not the match-occurrence index inside a single-pattern result set.

That aligns the additive Go helper surface with the SikuliX expectation that callers can map the returned match back to the pattern list they supplied.

## Performance Note

This phase intentionally implements multi-target helpers as stable API composition in `pkg/sikuli`, not as a new batch RPC surface.

Implications:
- image-backed helpers incur local matcher work only
- live helpers incur one API capture round-trip per helper call or poll cycle
- live helpers do not incur one capture per pattern inside a single helper call

A future performance phase can still add batch RPCs without changing the public helper contract introduced here.
