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

Generated: `2026-03-05T04:26:02.429969+00:00`

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
| akaze | 120 | 60 | 0 | 59 | 0 | 0 | 1 | 155.526 | 129.394 |
| brisk | 120 | 65 | 0 | 47 | 0 | 0 | 8 | 328.511 | 108.479 |
| hybrid | 120 | 90 | 0 | 29 | 0 | 0 | 1 | 130.107 | 88.616 |
| kaze | 120 | 75 | 0 | 38 | 0 | 0 | 7 | 726.687 | 557.288 |
| orb | 120 | 37 | 0 | 73 | 0 | 0 | 10 | 51.102 | 37.761 |
| sift | 120 | 75 | 0 | 43 | 0 | 0 | 2 | 233.382 | 191.003 |
| template | 120 | 72 | 0 | 48 | 0 | 0 | 0 | 121.025 | 86.137 |

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

- [Visual Root Directory](https://smysnk.github.io/SikuliGO/bench/reports/visuals)
- [Scenario Summaries Directory](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries)
- [Attempt Images Directory](https://smysnk.github.io/SikuliGO/bench/reports/visuals/attempts)

### Scenario Summary Images (40)

#### `hybrid_gate_conflicts_1024x768_i09`

![hybrid_gate_conflicts_1024x768_i09](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-hybrid_gate_conflicts_1024x768_i09.png)

- [Open `hybrid_gate_conflicts_1024x768_i09` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-hybrid_gate_conflicts_1024x768_i09.png)

#### `hybrid_gate_conflicts_1280x720_i09`

![hybrid_gate_conflicts_1280x720_i09](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-hybrid_gate_conflicts_1280x720_i09.png)

- [Open `hybrid_gate_conflicts_1280x720_i09` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-hybrid_gate_conflicts_1280x720_i09.png)

#### `hybrid_gate_conflicts_1920x1080_i09`

![hybrid_gate_conflicts_1920x1080_i09](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-hybrid_gate_conflicts_1920x1080_i09.png)

- [Open `hybrid_gate_conflicts_1920x1080_i09` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-hybrid_gate_conflicts_1920x1080_i09.png)

#### `hybrid_gate_conflicts_800x600_i09`

![hybrid_gate_conflicts_800x600_i09](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-hybrid_gate_conflicts_800x600_i09.png)

- [Open `hybrid_gate_conflicts_800x600_i09` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-hybrid_gate_conflicts_800x600_i09.png)

#### `multi_monitor_dpi_shift_1024x768_i10`

![multi_monitor_dpi_shift_1024x768_i10](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-multi_monitor_dpi_shift_1024x768_i10.png)

- [Open `multi_monitor_dpi_shift_1024x768_i10` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-multi_monitor_dpi_shift_1024x768_i10.png)

#### `multi_monitor_dpi_shift_1280x720_i10`

![multi_monitor_dpi_shift_1280x720_i10](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-multi_monitor_dpi_shift_1280x720_i10.png)

- [Open `multi_monitor_dpi_shift_1280x720_i10` image](https://smysnk.github.io/SikuliGO/bench/reports/visuals/summaries/summary-multi_monitor_dpi_shift_1280x720_i10.png)

- 34 additional scenario images available in the summaries directory.

<!-- END: FIND_ON_SCREEN_BENCH_AUTOGEN -->
