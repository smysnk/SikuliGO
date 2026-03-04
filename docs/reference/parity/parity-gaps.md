# Parity Gaps

Open Java parity gaps tracked for SikuliGO.

## API-Level Gaps

- Stateful keyboard modifier API (`keyDown`/`keyUp`) as first-class operations.
- Full window metadata parity vs Java runtime window model.
- Additional convenience/legacy aliases that exist in Java scripting layers.

## Behavioral Gaps

- Engine-specific tuning defaults may not exactly mirror Java expectations.
- Some wait/vanish edge behaviors are implemented via explicit polling wrappers in clients.

## Priority

- P0: user-facing script portability blockers
- P1: high-frequency workflow ergonomics
- P2: compatibility conveniences

## Resolution Rules

- Add mapping entry in `java-to-go-seed.tsv` before implementation.
- Add/update tests in `parity-test-matrix.md` once behavior ships.
