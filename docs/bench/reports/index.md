---
layout: guide
title: Benchmark Reports
nav_key: benchmarks
kicker: Benchmarks
lead: Current benchmark artifacts, detailed reports, and raw outputs for the latest published run.
---

<div class="guide-grid">
<a class="guide-card" href="{{ '/bench' | relative_url }}">
  <span class="guide-card__eyebrow">Overview</span>
  <span class="guide-card__title">Benchmarks</span>
  <span class="guide-card__body">Return to the benchmark overview at the section root.</span>
</a>
<a class="guide-card" href="{{ '/bench/reports/find-on-screen-e2e' | relative_url }}">
  <span class="guide-card__eyebrow">Detailed Report</span>
  <span class="guide-card__title">FindOnScreen E2E</span>
  <span class="guide-card__body">Open the full benchmark breakdown page.</span>
</a>
<a class="guide-card" href="{{ '/bench/reports/find-on-screen-scenario-strategy' | relative_url }}">
  <span class="guide-card__eyebrow">Strategy</span>
  <span class="guide-card__title">Scenario Strategy</span>
  <span class="guide-card__body">Review the scenario set and visual examples.</span>
</a>
<a class="guide-card" href="{{ '/bench/reports/visuals' | relative_url }}">
  <span class="guide-card__eyebrow">Artifacts</span>
  <span class="guide-card__title">Visual Gallery</span>
  <span class="guide-card__body">Open the generated image gallery and summaries.</span>
</a>
</div>

## Run Summary

<div class="guide-meta">
  <div class="guide-meta__item">
    <span class="guide-meta__label">Generated</span>
    `2026-03-07T23:32:15.506029+00:00`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Target</span>
    `BenchmarkFindOnScreenE2E`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Engines</span>
    `akaze`, `brisk`, `hybrid`, `kaze`, `orb`, `sift`, `template`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Manifest</span>
    `docs/bench/find-on-screen-scenarios.example.json`
  </div>
</div>

## Artifacts

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
