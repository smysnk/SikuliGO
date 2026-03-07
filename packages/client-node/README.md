# sikuligo (Node.js)

SikuliGO is a GoLang implementation of Sikuli visual automation. This package provides the Node.js SDK for launching `sikuligo` locally and executing automation with a small API surface.

## Links

- Main repository: [github.com/smysnk/SikuliGO](https://github.com/smysnk/SikuliGO)
- API reference: [smysnk.github.io/SikuliGO/reference/api](https://smysnk.github.io/SikuliGO/reference/api/)
- Node user flow: [smysnk.github.io/SikuliGO/guides/node-package-user-flow](https://smysnk.github.io/SikuliGO/guides/node-package-user-flow)
- Client strategy: [smysnk.github.io/SikuliGO/strategy/client-strategy](https://smysnk.github.io/SikuliGO/strategy/client-strategy)
- Architecture docs: [Port Strategy](https://smysnk.github.io/SikuliGO/strategy/port-strategy), [gRPC Strategy](https://smysnk.github.io/SikuliGO/strategy/grpc-strategy), [Java Parity Map](https://smysnk.github.io/SikuliGO/reference/parity/java-to-go-mapping)

## Quickstart

`init:js-examples` prompts for a target directory, scaffolds a `package.json` with the latest `@sikuligo/sikuligo` dependency, runs `yarn install`, and copies `.mjs` examples into `examples/`.

```bash
yarn dlx @sikuligo/sikuligo init:js-examples
cd sikuligo-demo
yarn node examples/click.mjs
```

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

## Web Dashboard
```bash
yarn dlx @sikuligo/sikuligo -listen 127.0.0.1:50051 -admin-listen :8080
```

Open:

- http://127.0.0.1:8080/dashboard

Additional endpoints:

- http://127.0.0.1:8080/healthz
- http://127.0.0.1:8080/metrics
- http://127.0.0.1:8080/snapshot

Install permanently on PATH:

```bash
yarn dlx @sikuligo/sikuligo install-binary
source ~/.zshrc
# or
source ~/.bash_profile
```

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
