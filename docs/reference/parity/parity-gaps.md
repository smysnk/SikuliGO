# Parity Gaps

Open Java parity gaps tracked for sikuli-go.

## API-Level Gaps

- Additional convenience and legacy aliases from Java scripting layers remain out of scope or are deferred to compatibility wrappers.
- Full Java runtime window object parity is still intentionally narrower in sikuli-go; portable window metadata varies by OS backend.
- Vision-engine defaults and tuning are not guaranteed to match Java/SikuliX heuristics exactly.

## Client-Level Gaps

- Node/Python/Lua wrappers still need to expose all newly closed API parity helpers consistently.
- Some SikuliX-style migration ergonomics remain in the client parity plan rather than the API parity plan.

## Priority

- P0: user-facing script portability blockers
- P1: high-frequency workflow ergonomics
- P2: compatibility conveniences

## Resolution Rules

- Add mapping entry in `java-to-go-seed.tsv` before implementation.
- Add/update tests in `parity-test-matrix.md` once behavior ships.
- Keep `api-parity-status.md` aligned with actual API maturity.
