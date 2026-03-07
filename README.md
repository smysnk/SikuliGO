# SikuliGO [![GoLang Tests](https://github.com/smysnk/SikuliGO/actions/workflows/go-test.yml/badge.svg)](https://github.com/smysnk/SikuliGO/actions/workflows/go-test.yml) [![Testing Gate](https://img.shields.io/github/actions/workflow/status/smysnk/SikuliGO/go-test.yml?branch=master&label=testing%20gate)](https://github.com/smysnk/SikuliGO/actions/workflows/go-test.yml) [![PyPI version](https://img.shields.io/pypi/v/sikuligo)](https://pypi.org/project/sikuligo/) [![npm version](https://img.shields.io/npm/v/%40sikuligo%2Fsikuligo)](https://www.npmjs.com/package/@sikuligo/sikuligo)

![SikuliX Logo](docs/images/logo.png)

Sikuli is an open-source tool for automating anything visible on a computer screen using image recognition. Instead of relying on internal source code or object IDs, it identifies and interacts with graphical user interface (GUI) components (buttons, text boxes, etc.) by using screenshots. **This repo houses a GoLang port of the original concept.**

## Project Intent

- Build a feature-complete GoLang port of the core [Sikuli](https://sikulix.github.io/) concepts.
- Preserve behavioral parity (image matching, regions, patterns, finder semantics).
- Provide a modern, testable architecture with explicit contracts and deterministic matching behavior.
- Establish a maintainable foundation for cross-platform automation features.

## Examples

### Node.js

Scaffold and run Node.js examples:

```bash
yarn dlx @sikuligo/sikuligo init:js-examples
cd sikuligo-demo
yarn node examples/click.mjs
```
`init:js-examples` prompts for a target directory, scaffolds a `package.json` with the latest `@sikuligo/sikuligo` dependency, runs `yarn install`, and copies `.mjs` examples into `examples/`.

```js
import { Screen, Pattern } from "@sikuligo/sikuligo";

const screen = await Screen();
try {
  const match = await screen.click(Pattern("assets/pattern.png").exact());
  console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
} finally {
  await screen.close();
}
```

### Python

Scaffold and run Python examples:

```bash
pipx run sikuligo init:py-examples
cd sikuligo-demo
python3 examples/click.py
```

`init:py-examples` prompts for a target directory, writes `requirements.txt`, creates `.venv`, installs dependencies, and copies Python examples into `examples/`.

```python
from sikuligo import Pattern, Screen

screen = Screen()
try:
    match = screen.click(Pattern("assets/pattern.png").exact())
    print(f"clicked match target at ({match.target_x}, {match.target_y})")
finally:
    screen.close()
```

## API Dashboard

![SikuliGO Dashboard Demo](docs/images/dashboard.png)

## Install API Binary On PATH

Install via Yarn (Node ecosystem):

```bash
yarn dlx @sikuligo/sikuligo install-binary
```

Install via Python:

```bash
pipx run sikuligo install-binary
```

Both commands create `~/.local/bin` if needed, copy all available `sikuli*` runtimes, and can prompt to add PATH to `~/.zshrc` or `~/.bash_profile`.
If PATH is updated, reload with `source ~/.zshrc` or `source ~/.bash_profile`.

## Available Clients

| Client |  | Notes |
| :---  | --- | :---  |
| [Python](https://pypi.org/project/sikuligo/)  | ✅ | Implemented |
| [Node](https://www.npmjs.com/package/@sikuligo/sikuligo)  | ✅ | Implemented |
| Lua  | ✅ | Implemented |
| Robot Framework | 🟡 | Planned |
| Web IDE | 🟡 | Planned |


## Current Focus

| Roadmap Item | Scope |  |
| :---  | :---  |---|
| Core API scaffolding | Public SikuliGo API surface and parity-facing core objects | ✅ |
| Matching engine and parity harness | Deterministic matcher behavior, golden corpus, backend conformance tests | ✅ |
| API parity surface expansion | Additional parity helpers and compatibility APIs | ✅ |
| Protocol completeness hardening | Alternate matcher backend + cross-backend conformance rules | ✅ |
| OCR and text-search parity | OCR contracts, finder/region text flows, optional backend integration | ✅ |
| Input automation and hotkey parity | Input controller contracts, request validation, backend protocol scaffold | 🟡 |
| Observe/event subsystem parity | Observer contracts, request validation, backend protocol scaffold | ✅ |
| App/window/process control parity | App/window contracts, request validation, backend protocol scaffold | ✅ |
| Cross-platform backend hardening | Platform integration hardening and backend portability | 🟡 |

# Docs
- [Docs Home](https://smysnk.github.io/SikuliGO/)

## Strategy
- [Mission](https://smysnk.github.io/SikuliGO/strategy/mission)
- [Port](https://smysnk.github.io/SikuliGO/strategy/port-strategy)
- [gRPC](https://smysnk.github.io/SikuliGO/strategy/grpc-strategy)
- [Client](https://smysnk.github.io/SikuliGO/strategy/client-strategy)
- [ORB Search Delivery](https://smysnk.github.io/SikuliGO/strategy/orb-search-delivery)

## Integration & Implementation
- [API Reference](https://smysnk.github.io/SikuliGO/reference/api/)
- [OCR](https://smysnk.github.io/SikuliGO/guides/ocr-integration)
- [OpenCV](https://smysnk.github.io/SikuliGO/guides/opencv-integration)
- [Input Automation](https://smysnk.github.io/SikuliGO/guides/input-automation)
- [Observe Events](https://smysnk.github.io/SikuliGO/guides/observe-events)
- [App Control](https://smysnk.github.io/SikuliGO/guides/app-control)
- [Defaults Table](https://smysnk.github.io/SikuliGO/guides/default-behavior-table)
- [Backend Capability Matrix](https://smysnk.github.io/SikuliGO/guides/backend-capability-matrix)
- [Node Package User Flow](https://smysnk.github.io/SikuliGO/guides/node-package-user-flow)
- [Java to Go API Mapping](https://smysnk.github.io/SikuliGO/reference/parity/java-to-go-mapping)
- [Parity Test Matrix](https://smysnk.github.io/SikuliGO/reference/parity/parity-test-matrix)

## Repository Layout

- [`pkg`](pkg) : public GoLang API packages
- [`packages/api`](packages/api) : GoLang API module (`cmd`, `internal`, `pkg`, `proto`)
- [`packages/editor`](packages/editor) : Next.js editor app
- [`packages/api-electron`](packages/api-electron) : macOS Electron API + dashboard/session viewer app
- [`clients`](clients) : language client SDKs and packaging artifacts
- [`docs`](docs) : documentation and assets
- [`legacy`](legacy) : previous Java-era project directories retained for reference

## Install and Build

Build from source:

- [Build From Source](docs/guides/build-from-source.md)

## Project History and Credits

Sikuli started in 2009 as an open-source research effort at the MIT User Interface Design Group, led by **Tsung-Hsiang Chang** and **Tom Yeh**, with early development connected to **Prof. Rob Miller**'s work at **MIT CSAIL**. The project introduced a practical idea that was unusual at the time: instead of relying on internal application APIs, users could automate **Graphical User Interfaces (GUI)** by teaching scripts what to click through screenshots of buttons, icons, and other visual elements. Even the name reflected that vision, drawing from the Huichol concept of the "**God's Eye**," a symbol of seeing and understanding what is otherwise hidden.

In 2012, after the original creators moved on, the project's active development continued under **RaiMan** and evolved into **SikuliX**. That branch carried the platform forward for real-world desktop and web automation, using scripting ecosystems such as **Jython/Python**, **Java**, and **Ruby**, and refining image-based interaction workflows over time. Because this style of automation simulates real **mouse** and **keyboard** behavior, it has always worked best in environments with an active graphical session rather than truly **headless** execution.

The GoLang port in this repository began in **2026**. It stands on the work of the original Sikuli authors, **RaiMan**, and the broader contributor community that kept visual automation practical and accessible over the years.

## Sikuli References

- [SikuliX Official Site](https://sikulix.github.io/)
- [Wikipedia](https://de.wikipedia.org/wiki/Sikuli_(Software))
- [Original Sikuli Github](https://github.com/sikuli/sikuli)
- [Sikuli Framework](https://github.com/smysnk/sikuli-framework) = Sikuli + Robot Framework

<!-- BEGIN: FIND_ON_SCREEN_BENCH_AUTOGEN -->
## FindOnScreen Benchmark Test Results

Generated: `2026-03-07T08:22:49.092218+00:00`

### Reports

- [Markdown Summary](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-e2e.md)
- [JSON Report](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-e2e.json)
- [Raw go test Output](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-e2e.txt)
- [Performance SVG](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-performance.svg)
- [Accuracy SVG](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-accuracy.svg)
- [Scenario Kind Match Time SVG](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-kind-time.svg)
- [Scenario Kind Success SVG](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-kind-success.svg)
- [Resolution Match Time SVG](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-resolution-time.svg)
- [Resolution Matches SVG](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-resolution-matches.svg)
- [Resolution Misses SVG](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-resolution-misses.svg)
- [Resolution False Positives SVG](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-resolution-false-positives.svg)

### Engine Summary

_Cases/OK metrics are query-level counts (regions x scenarios x resolutions), not just benchmark row count._

| Engine | Cases | OK | Partial | Not Found | Unsupported | Error | Overlap Miss | Avg ms/op | Median ms/op |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| akaze | 120 | 39 | 0 | 78 | 0 | 0 | 3 | 171.032 | 147.060 |
| brisk | 120 | 47 | 0 | 63 | 0 | 0 | 10 | 331.625 | 123.717 |
| hybrid | 120 | 69 | 0 | 45 | 0 | 0 | 6 | 171.317 | 135.641 |
| kaze | 120 | 63 | 0 | 50 | 0 | 0 | 7 | 787.238 | 637.542 |
| orb | 120 | 13 | 0 | 96 | 0 | 0 | 11 | 55.807 | 43.110 |
| sift | 120 | 56 | 0 | 55 | 0 | 0 | 9 | 251.719 | 195.118 |
| template | 120 | 64 | 0 | 56 | 0 | 0 | 0 | 155.890 | 117.008 |

### Run Mega Summary

![Run Mega Summary](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-run-mega.jpg)

- [Open run mega summary image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-run-mega.jpg)

### Benchmark Graphs

![Performance Graph](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-performance.svg)

- [Open performance graph](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-performance.svg)

![Accuracy Graph](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-accuracy.svg)

- [Open accuracy graph](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-accuracy.svg)

### Scenario Kind Graphs

![Scenario Kind Match Time](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-kind-time.svg)

- [Open scenario kind match time graph](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-kind-time.svg)

![Scenario Kind Success](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-kind-success.svg)

- [Open scenario kind success graph](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-kind-success.svg)

### Resolution Group Graphs

![Resolution Match Time](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-resolution-time.svg)

- [Open resolution match time graph](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-resolution-time.svg)

![Resolution Matches](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-resolution-matches.svg)

- [Open resolution matches graph](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-resolution-matches.svg)

![Resolution Misses](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-resolution-misses.svg)

- [Open resolution misses graph](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-resolution-misses.svg)

![Resolution False Positives](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-resolution-false-positives.svg)

- [Open resolution false positives graph](https://smysnk.github.io/SikuliGO/bench/reports/find-on-screen-resolution-false-positives.svg)

### Artifact Directories

- [Visual Root](https://smysnk.github.io/SikuliGO/bench/reports/visuals/index.html)
- [Scenario Summaries](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/index.html)
- [Attempt Images](https://smysnk.github.io/SikuliGO/bench/reports/visuals/attempts/index.html)

### Scenario Summary Images (10)

#### `hybrid_gate_conflicts_1920x1080_i09`

![hybrid_gate_conflicts_1920x1080_i09](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-hybrid_gate_conflicts_1920x1080_i09.png)

- [Open `hybrid_gate_conflicts_1920x1080_i09` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-hybrid_gate_conflicts_1920x1080_i09.png)

#### `multi_monitor_dpi_shift_1920x1080_i10`

![multi_monitor_dpi_shift_1920x1080_i10](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-multi_monitor_dpi_shift_1920x1080_i10.png)

- [Open `multi_monitor_dpi_shift_1920x1080_i10` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-multi_monitor_dpi_shift_1920x1080_i10.png)

#### `noise_stress_random_1920x1080_i04`

![noise_stress_random_1920x1080_i04](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-noise_stress_random_1920x1080_i04.png)

- [Open `noise_stress_random_1920x1080_i04` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-noise_stress_random_1920x1080_i04.png)

#### `orb_feature_rich_1920x1080_i07`

![orb_feature_rich_1920x1080_i07](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-orb_feature_rich_1920x1080_i07.png)

- [Open `orb_feature_rich_1920x1080_i07` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-orb_feature_rich_1920x1080_i07.png)

#### `perspective_skew_sweep_1920x1080_i06`

![perspective_skew_sweep_1920x1080_i06](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-perspective_skew_sweep_1920x1080_i06.png)

- [Open `perspective_skew_sweep_1920x1080_i06` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-perspective_skew_sweep_1920x1080_i06.png)

#### `photo_clutter_1920x1080_i02`

![photo_clutter_1920x1080_i02](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-photo_clutter_1920x1080_i02.png)

- [Open `photo_clutter_1920x1080_i02` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-photo_clutter_1920x1080_i02.png)

#### `repetitive_grid_camouflage_1920x1080_i03`

![repetitive_grid_camouflage_1920x1080_i03](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-repetitive_grid_camouflage_1920x1080_i03.png)

- [Open `repetitive_grid_camouflage_1920x1080_i03` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-repetitive_grid_camouflage_1920x1080_i03.png)

#### `scale_rotate_sweep_1920x1080_i05`

![scale_rotate_sweep_1920x1080_i05](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-scale_rotate_sweep_1920x1080_i05.png)

- [Open `scale_rotate_sweep_1920x1080_i05` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-scale_rotate_sweep_1920x1080_i05.png)

#### `template_control_exact_1920x1080_i08`

![template_control_exact_1920x1080_i08](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-template_control_exact_1920x1080_i08.png)

- [Open `template_control_exact_1920x1080_i08` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-template_control_exact_1920x1080_i08.png)

#### `vector_ui_baseline_1920x1080_i01`

![vector_ui_baseline_1920x1080_i01](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-vector_ui_baseline_1920x1080_i01.png)

- [Open `vector_ui_baseline_1920x1080_i01` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-vector_ui_baseline_1920x1080_i01.png)

<!-- END: FIND_ON_SCREEN_BENCH_AUTOGEN -->
