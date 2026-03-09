---
layout: guide
title: Benchmarks
nav_key: benchmarks
kicker: Secondary
lead: Latest find-on-screen benchmark results, scenario strategy, and published artifacts rendered into the guide shell.
---

<div class="guide-grid">
<a class="guide-card" href="{{ '/bench/reports/find-on-screen-e2e' | relative_url }}">
  <span class="guide-card__eyebrow">Detailed Report</span>
  <span class="guide-card__title">Engine Breakdown</span>
  <span class="guide-card__body">Inspect engine latency, success, and resolution breakdowns.</span>
</a>
<a class="guide-card" href="{{ '/bench/reports/find-on-screen-scenario-strategy' | relative_url }}">
  <span class="guide-card__eyebrow">Scenario Strategy</span>
  <span class="guide-card__title">Corpus Design</span>
  <span class="guide-card__body">Review the scenario matrix, visual examples, and engine-selection rationale.</span>
</a>
<a class="guide-card" href="{{ '/bench/reports/visuals' | relative_url }}">
  <span class="guide-card__eyebrow">Artifacts</span>
  <span class="guide-card__title">Visual Gallery</span>
  <span class="guide-card__body">Open the generated screenshot gallery and summary images.</span>
</a>
<a class="guide-card guide-card--subtle" href="{{ '/bench/FIND_ON_SCREEN_SCENARIO_INTENT' | relative_url }}">
  <span class="guide-card__eyebrow">Scenario Docs</span>
  <span class="guide-card__title">Scenario Intent</span>
  <span class="guide-card__body">See what each benchmark scenario is trying to prove.</span>
</a>
</div>

## Latest Run

<div class="guide-meta">
  <div class="guide-meta__item">
    <span class="guide-meta__label">Generated</span>
    `2026-03-07T23:32:15.506029+00:00`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Platform</span>
    `darwin/arm64` on `Apple M4 Pro`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Target</span>
    `BenchmarkFindOnScreenE2E` with benchtime `200ms` and count `1`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Scenario Corpus</span>
    `10` scenario types across `4` resolution groups
  </div>
</div>

<div class="guide-callout">
  <strong>Key Findings</strong>
  <span>Most accurate: `hybrid` at `57.5%` success. Fastest: `orb` at `56.443` ms/op. Lowest false-positive rate: `template` at `0.0%`.</span>
</div>

## Engine Snapshot

| Engine | Cases | Avg ms/op | Median ms/op | Success % | False Positive % | No Match % |
| --- | --- | --- | --- | --- | --- | --- |
| `akaze` | 120 | 172.121 | 147.695 | 32.5 | 2.5 | 65.0 |
| `brisk` | 120 | 388.483 | 123.118 | 39.2 | 8.3 | 52.5 |
| `hybrid` | 120 | 171.017 | 134.411 | 57.5 | 5.0 | 37.5 |
| `kaze` | 120 | 824.898 | 640.512 | 52.5 | 5.8 | 41.7 |
| `orb` | 120 | 56.443 | 44.794 | 10.8 | 9.2 | 80.0 |
| `sift` | 120 | 256.756 | 198.264 | 46.7 | 7.5 | 45.8 |
| `template` | 120 | 154.257 | 114.466 | 53.3 | 0.0 | 46.7 |

## Charts

![Performance chart]({{ '/bench/reports/find-on-screen-performance.svg' | relative_url }})

![Accuracy chart]({{ '/bench/reports/find-on-screen-accuracy.svg' | relative_url }})

![Resolution time chart]({{ '/bench/reports/find-on-screen-resolution-time.svg' | relative_url }})

![Resolution matches chart]({{ '/bench/reports/find-on-screen-resolution-matches.svg' | relative_url }})

## Artifact Map

| Artifact | Purpose | Link |
| --- | --- | --- |
| Overview | Current benchmark summary at the section root. | [Open]({{ '/bench' | relative_url }}) |
| Reports Hub | Artifact map for the current benchmark run. | [Open]({{ '/bench/reports' | relative_url }}) |
| Detailed E2E | Full engine, resolution, and scenario breakdown. | [Open]({{ '/bench/reports/find-on-screen-e2e' | relative_url }}) |
| Benchmark JSON | Machine-readable benchmark summary. | [Open]({{ '/bench/reports/find-on-screen-e2e.json' | relative_url }}) |
| Benchmark Text | Raw `go test` benchmark output. | [Open]({{ '/bench/reports/find-on-screen-e2e.txt' | relative_url }}) |
| Scenario Strategy | Scenario corpus and engine-selection rationale. | [Open]({{ '/bench/reports/find-on-screen-scenario-strategy' | relative_url }}) |
| Strategy JSON | Machine-readable strategy summary. | [Open]({{ '/bench/reports/find-on-screen-scenario-strategy.json' | relative_url }}) |
| Visual Gallery | Generated benchmark screenshots and summaries. | [Open]({{ '/bench/reports/visuals' | relative_url }}) |
| Scenario Intent | What each scenario is intended to prove. | [Open]({{ '/bench/FIND_ON_SCREEN_SCENARIO_INTENT' | relative_url }}) |
| Scenario Schema | Manifest schema and region workflow. | [Open]({{ '/bench/FIND_ON_SCREEN_SCENARIO_SCHEMA' | relative_url }}) |

## Scenario Corpus

| Document | Purpose | Link |
| --- | --- | --- |
| Intent | Why each scenario exists and what it should stress. | [Open]({{ '/bench/FIND_ON_SCREEN_SCENARIO_INTENT' | relative_url }}) |
| Schema | Manifest structure, region-selection flow, and validation inputs. | [Open]({{ '/bench/FIND_ON_SCREEN_SCENARIO_SCHEMA' | relative_url }}) |
