# Build From Source

This is the single source of truth for building SikuliGO from source.

## Prerequisites

- Go `1.24+`
- Node.js `20+`
- Yarn `4+`
- Python `3.10+`
- `protoc`

## Install Workspace Dependencies

```bash
cd /path/to/SikuliX1
yarn install
```

## Build Go API Binaries

Build `sikuligo`:

```bash
cd packages/api
go build -tags "gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d" -trimpath -ldflags="-s -w" -o ../../sikuligo ./cmd/sikuligrpc
```

Build `sikuligo-monitor`:

```bash
cd packages/api
go build -tags "gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d" -trimpath -ldflags="-s -w" -o ../../sikuligo-monitor ./cmd/sikuligo-monitor
```

## Build Protocol Artifacts

```bash
./scripts/generate-grpc-stubs.sh
./scripts/clients/generate-node-stubs.sh
./scripts/clients/generate-python-stubs.sh
./scripts/clients/generate-lua-descriptor.sh
```

## Build Node Client

```bash
yarn workspace @sikuligo/sikuligo build
```

## Build Python Distributions

```bash
./scripts/clients/release-python-client.sh
```

Skip installer steps:

```bash
SKIP_INSTALL=1 ./scripts/clients/release-python-client.sh
```

## Build Everything (Convenience)

```bash
make
```

## Publish GitHub Docs

Build docs locally and publish the generated site to the `gh-pages` branch:

```bash
make gh-publish
```

Options:

- `GH_PAGES_BRANCH=gh-pages` target branch
- `GH_PAGES_REMOTE=origin` target remote
- `GH_PAGES_INCLUDE_BENCH=1` include local benchmark screenshots from `.test-results/bench`
- `GH_PAGES_FORCE_PUSH=1` force-push the publish branch
- `GH_PAGES_CNAME=docs.example.com` include a CNAME file in published output
- `GH_PAGES_CONFIGURE_SOURCE=1` configure repository Pages source to `gh-pages` using `gh api`
- `GH_PAGES_REMOTE_URL=https://<token>@github.com/<owner>/<repo>.git` explicit authenticated remote URL for non-interactive CI pushes

CI automation:

- `.github/workflows/docs-pages.yml` runs benchmark + publish on pushes to `main` (and `master` compatibility) and on manual dispatch.

## Local End-to-End Verification

Run the local verifier:

```bash
make test-publish
```

What it checks:

- Builds local `sikuligo` binary and Node client
- Verifies CLI passthrough help output
- Verifies `init:js-examples` scaffolds `.mjs`-only examples
- Verifies `init:py-examples` scaffolds `requirements.txt` + Python examples
- Runs a transport smoke check and fails on known regressions:
  - `unknown gRPC method: ClickOnScreen`
  - `Request message serialization failure`

Optional: verify using packed tarball install flow (closest to published package install):

```bash
VERIFY_PACKED_INSTALL=1 make test-publish
```

Run the full local integration suite:

```bash
make test-integration
```

What `test-integration` adds on top of `test-publish`:

- gRPC RPC-surface integration coverage for all service methods
- gRPC image/OCR E2E verification (`FindOnScreen`, `ReadText`, `FindText`)
- API package cross-protocol integration flow tests
- Node client E2E startup/connect scenarios (`auto -> connect` and `auto -> spawn`)
- Python client E2E startup/connect scenarios (`auto -> connect` and `auto -> spawn`)
- Optional real-desktop E2E when `RUN_REAL_DESKTOP_E2E=1` (manual desktop fixture + live OCR)

Run optional real-desktop E2E directly:

```bash
REAL_DESKTOP_E2E=1 make test-e2e
```

Select a specific monitor/display for capture (useful on multi-monitor setups):

```bash
REAL_DESKTOP_E2E=1 REAL_DESKTOP_E2E_DISPLAY=2 make test-e2e
```

Notes:

- Intended for local/manual execution on a real desktop session (not CI/headless).
- Builds `sikuligo` with OCR/OpenCV tags, opens a fixture page, then validates:
  - `FindOnScreen` against a visible image
  - OCR (`ReadText` + `FindText`) from a live `/snapshot`
- `REAL_DESKTOP_E2E_DISPLAY` maps to `SIKULI_CAPTURE_DISPLAY`/`SIKULIGO_CAPTURE_DISPLAY` for `screencapture -D <display>`.

## Find Benchmark E2E

Benchmark `FindOnScreen` across matcher implementations (`template`, `orb`, `hybrid`) and multiple fixture scenarios:

- different image families (`grid`, `glyph`)
- ORB-friendly image families (`noise`, `orbtex`)
- different pattern sizes (`small`, `medium`, `large`)
- different orientations (`0`, `90`, `180`, `270`)
- different synthetic screen sizes (`480x270`, `640x360`, `800x450`, `960x540`, `1024x576`)

Run:

```bash
make benchmark
```

Output artifacts (default):

- `.test-results/bench/find-on-screen-e2e.txt` (raw `go test` benchmark output)
- `.test-results/bench/find-on-screen-e2e.json` (machine-readable report)
- `.test-results/bench/find-on-screen-e2e.md` (human-readable summary)
- `.test-results/bench/find-on-screen-performance.svg` (engine-level latency chart)
- `.test-results/bench/find-on-screen-accuracy.svg` (engine-level success vs false-positive chart)
- `.test-results/bench/find-on-screen-resolution-time.svg` (resolution-grouped latency chart by engine)
- `.test-results/bench/find-on-screen-resolution-matches.svg` (resolution-grouped match counts by engine)
- `.test-results/bench/find-on-screen-resolution-misses.svg` (resolution-grouped miss counts by engine)
- `.test-results/bench/find-on-screen-resolution-false-positives.svg` (resolution-grouped false-positive counts by engine)
- `.test-results/bench/visuals/attempts/...` (annotated per-attempt screenshots)
- `.test-results/bench/visuals/summaries/...` (per-scenario combined summary screenshots across engines/attempts)

Useful options:

```bash
FIND_BENCH_TIME=500ms FIND_BENCH_COUNT=3 make benchmark
FIND_BENCH_TAGS="opencv gocv_specific_modules gocv_features2d gocv_calib3d" make benchmark
FIND_BENCH_TAGS="" make benchmark
FIND_BENCH_REPORT_DIR=.test-results/custom make benchmark
FIND_BENCH_VISUAL=1 FIND_BENCH_VISUAL_MAX_ATTEMPTS=2 make benchmark
FIND_BENCH_VISUAL_DIR=.test-results/bench/custom-visuals FIND_BENCH_VISUAL_TIMEOUT=8s make benchmark
make benchmark
FIND_BENCH_PATCH_READMES=1 FIND_BENCH_README_PATHS="$PWD/README.md,$PWD/packages/client-node/README.md,$PWD/packages/client-python/README.md" make benchmark
FIND_BENCH_PATCH_READMES=1 FIND_BENCH_README_INLINE_IMAGES=4 FIND_BENCH_README_SECTION_TITLE="Latest Benchmark Evidence" make benchmark
FIND_BENCH_PATCH_READMES=0 make benchmark
```

By default, the benchmark runs with OpenCV-related tags enabled so `orb` and `hybrid` implementations can be compared.

README patching is enabled by default. On each benchmark run, the script appends/updates an autogenerated section at the bottom of each target README with:

- links to `.txt/.json/.md` benchmark artifacts
- engine summary table
- embedded benchmark screenshots (run mega summary + scenario summaries)
- direct links to generated images

README patch controls:

- `FIND_BENCH_README_PATHS` comma-separated paths (absolute or repo-root relative)
- `FIND_BENCH_README_SECTION_TITLE` section heading text
- `FIND_BENCH_README_INLINE_IMAGES` number of scenario images embedded inline per README

## Optional OCR-Tagged Tests

```bash
cd packages/api
go test -tags gosseract ./...
```
