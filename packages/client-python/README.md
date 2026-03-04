# SikuliGO Python Client

This directory contains the Python client for SikuliGO with Sikuli-style `Screen` + `Pattern` APIs.

## Links

- Main repository: [github.com/smysnk/SikuliGO](https://github.com/smysnk/SikuliGO)
- API reference: [smysnk.github.io/SikuliGO/reference/api](https://smysnk.github.io/SikuliGO/reference/api/)
- Client strategy: [smysnk.github.io/SikuliGO/strategy/client-strategy](https://smysnk.github.io/SikuliGO/strategy/client-strategy)
- Architecture docs: [Port Strategy](https://smysnk.github.io/SikuliGO/strategy/port-strategy), [gRPC Strategy](https://smysnk.github.io/SikuliGO/strategy/grpc-strategy), [Java Parity Map](https://smysnk.github.io/SikuliGO/reference/parity/java-to-go-mapping)

## Quickstart

`init:py-examples` prompts for the target directory, creates `requirements.txt`, installs into `.venv`, and copies examples.
Each example bootstraps `sikuligo` into `./.sikuligo/bin` and prepends it to PATH for the process.

```bash
pipx run sikuligo init:py-examples
cd sikuligo-demo
python3 examples/click.py
```

runs:
```python
from __future__ import annotations
from sikuligo import Pattern, Screen

screen = Screen()
try:
    match = screen.click(Pattern("assets/pattern.png").exact())
    print(f"clicked match target at ({match.target_x}, {match.target_y})")
finally:
    screen.close()
```

## Web Dashboard
```bash
pipx run sikuligo -listen 127.0.0.1:50051 -admin-listen :8080
```

Open:

- http://127.0.0.1:8080/dashboard

Additional endpoints:

- http://127.0.0.1:8080/healthz
- http://127.0.0.1:8080/metrics
- http://127.0.0.1:8080/snapshot

Install permanently on PATH:

```bash
pipx run sikuligo install-binary
source ~/.zshrc
# or
source ~/.bash_profile
```

<!-- BEGIN: FIND_ON_SCREEN_BENCH_AUTOGEN -->
## FindOnScreen Benchmark Test Results

Generated: `2026-03-04T08:47:04.585152+00:00`

### Reports

- [Markdown Summary](../../.test-results/bench/find-on-screen-e2e.md)
- [JSON Report](../../.test-results/bench/find-on-screen-e2e.json)
- [Raw go test Output](../../.test-results/bench/find-on-screen-e2e.txt)
- [Performance SVG](../../.test-results/bench/find-on-screen-performance.svg)
- [Accuracy SVG](../../.test-results/bench/find-on-screen-accuracy.svg)
- [Scenario Kind Match Time SVG](../../.test-results/bench/find-on-screen-kind-time.svg)
- [Scenario Kind Success SVG](../../.test-results/bench/find-on-screen-kind-success.svg)
- [Resolution Match Time SVG](../../.test-results/bench/find-on-screen-resolution-time.svg)
- [Resolution Matches SVG](../../.test-results/bench/find-on-screen-resolution-matches.svg)
- [Resolution Misses SVG](../../.test-results/bench/find-on-screen-resolution-misses.svg)
- [Resolution False Positives SVG](../../.test-results/bench/find-on-screen-resolution-false-positives.svg)

### Engine Summary

| Engine | Cases | OK | Partial | Not Found | Unsupported | Error | Overlap Miss | Avg ms/op | Median ms/op |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| akaze | 4 | 1 | 0 | 3 | 0 | 0 | 0 | 109.348 | 88.370 |
| brisk | 4 | 0 | 0 | 4 | 0 | 0 | 0 | 70.323 | 66.291 |
| hybrid | 4 | 4 | 0 | 0 | 0 | 0 | 0 | 85.527 | 69.467 |
| kaze | 4 | 2 | 0 | 0 | 0 | 0 | 2 | 464.642 | 374.102 |
| orb | 4 | 0 | 0 | 4 | 0 | 0 | 0 | 21.309 | 18.552 |
| sift | 4 | 4 | 0 | 0 | 0 | 0 | 0 | 165.972 | 138.164 |
| template | 4 | 4 | 0 | 0 | 0 | 0 | 0 | 86.545 | 68.910 |

### Run Mega Summary

![Run Mega Summary](../../.test-results/bench/visuals/summaries/summary-run-mega.jpg)

- [Open run mega summary image](../../.test-results/bench/visuals/summaries/summary-run-mega.jpg)

### Benchmark Graphs

![Performance Graph](../../.test-results/bench/find-on-screen-performance.svg)

- [Open performance graph](../../.test-results/bench/find-on-screen-performance.svg)

![Accuracy Graph](../../.test-results/bench/find-on-screen-accuracy.svg)

- [Open accuracy graph](../../.test-results/bench/find-on-screen-accuracy.svg)

### Scenario Kind Graphs

![Scenario Kind Match Time](../../.test-results/bench/find-on-screen-kind-time.svg)

- [Open scenario kind match time graph](../../.test-results/bench/find-on-screen-kind-time.svg)

![Scenario Kind Success](../../.test-results/bench/find-on-screen-kind-success.svg)

- [Open scenario kind success graph](../../.test-results/bench/find-on-screen-kind-success.svg)

### Resolution Group Graphs

![Resolution Match Time](../../.test-results/bench/find-on-screen-resolution-time.svg)

- [Open resolution match time graph](../../.test-results/bench/find-on-screen-resolution-time.svg)

![Resolution Matches](../../.test-results/bench/find-on-screen-resolution-matches.svg)

- [Open resolution matches graph](../../.test-results/bench/find-on-screen-resolution-matches.svg)

![Resolution Misses](../../.test-results/bench/find-on-screen-resolution-misses.svg)

- [Open resolution misses graph](../../.test-results/bench/find-on-screen-resolution-misses.svg)

![Resolution False Positives](../../.test-results/bench/find-on-screen-resolution-false-positives.svg)

- [Open resolution false positives graph](../../.test-results/bench/find-on-screen-resolution-false-positives.svg)

### Artifact Directories

- [Visual Root Directory](../../.test-results/bench/visuals)
- [Scenario Summaries Directory](../../.test-results/bench/visuals/summaries)
- [Attempt Images Directory](../../.test-results/bench/visuals/attempts)

### Scenario Summary Images (44)

#### `noise_stress_random_1024x576_i04_rotate_sc66006f4`

![noise_stress_random_1024x576_i04_rotate_sc66006f4](../../.test-results/bench/visuals/summaries/summary-noise_stress_random_1024x576_i04_rotate_sc66006f4.png)

- [Open `noise_stress_random_1024x576_i04_rotate_sc66006f4` image](../../.test-results/bench/visuals/summaries/summary-noise_stress_random_1024x576_i04_rotate_sc66006f4.png)

#### `noise_stress_random_1280x720_i04_rotate_sc37a59d9`

![noise_stress_random_1280x720_i04_rotate_sc37a59d9](../../.test-results/bench/visuals/summaries/summary-noise_stress_random_1280x720_i04_rotate_sc37a59d9.png)

- [Open `noise_stress_random_1280x720_i04_rotate_sc37a59d9` image](../../.test-results/bench/visuals/summaries/summary-noise_stress_random_1280x720_i04_rotate_sc37a59d9.png)

#### `noise_stress_random_1920x1080_i04_rotate_s17ff8896`

![noise_stress_random_1920x1080_i04_rotate_s17ff8896](../../.test-results/bench/visuals/summaries/summary-noise_stress_random_1920x1080_i04_rotate_s17ff8896.png)

- [Open `noise_stress_random_1920x1080_i04_rotate_s17ff8896` image](../../.test-results/bench/visuals/summaries/summary-noise_stress_random_1920x1080_i04_rotate_s17ff8896.png)

#### `noise_stress_random_480x270_i04_rotate_s67369b96`

![noise_stress_random_480x270_i04_rotate_s67369b96](../../.test-results/bench/visuals/summaries/summary-noise_stress_random_480x270_i04_rotate_s67369b96.png)

- [Open `noise_stress_random_480x270_i04_rotate_s67369b96` image](../../.test-results/bench/visuals/summaries/summary-noise_stress_random_480x270_i04_rotate_s67369b96.png)

#### `noise_stress_random_640x360_i04_rotate_s0f3b4e62`

![noise_stress_random_640x360_i04_rotate_s0f3b4e62](../../.test-results/bench/visuals/summaries/summary-noise_stress_random_640x360_i04_rotate_s0f3b4e62.png)

- [Open `noise_stress_random_640x360_i04_rotate_s0f3b4e62` image](../../.test-results/bench/visuals/summaries/summary-noise_stress_random_640x360_i04_rotate_s0f3b4e62.png)

#### `noise_stress_random_800x450_i04_rotate_s1e6f64ce`

![noise_stress_random_800x450_i04_rotate_s1e6f64ce](../../.test-results/bench/visuals/summaries/summary-noise_stress_random_800x450_i04_rotate_s1e6f64ce.png)

- [Open `noise_stress_random_800x450_i04_rotate_s1e6f64ce` image](../../.test-results/bench/visuals/summaries/summary-noise_stress_random_800x450_i04_rotate_s1e6f64ce.png)

- 38 additional scenario images available in the summaries directory.

<!-- END: FIND_ON_SCREEN_BENCH_AUTOGEN -->
