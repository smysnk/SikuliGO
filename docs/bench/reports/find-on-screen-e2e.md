---
layout: guide
title: FindOnScreen Benchmark Report
nav_key: benchmarks
kicker: Benchmarks
lead: Full engine, resolution, scenario-kind, and artifact breakdown for the latest benchmark run.
---

<div class="guide-grid">
<a class="guide-card" href="{{ '/bench' | relative_url }}">
  <span class="guide-card__eyebrow">Overview</span>
  <span class="guide-card__title">Benchmark Overview</span>
  <span class="guide-card__body">Return to the section summary and latest charts.</span>
</a>
<a class="guide-card" href="{{ '/bench/reports' | relative_url }}">
  <span class="guide-card__eyebrow">Reports Hub</span>
  <span class="guide-card__title">Artifact Index</span>
  <span class="guide-card__body">Browse the raw outputs and related benchmark pages.</span>
</a>
<a class="guide-card" href="{{ '/bench/reports/find-on-screen-scenario-strategy' | relative_url }}">
  <span class="guide-card__eyebrow">Strategy</span>
  <span class="guide-card__title">Scenario Strategy</span>
  <span class="guide-card__body">Inspect the scenario corpus and visual examples.</span>
</a>
<a class="guide-card guide-card--subtle" href="{{ '/bench/reports/visuals' | relative_url }}">
  <span class="guide-card__eyebrow">Artifacts</span>
  <span class="guide-card__title">Visual Gallery</span>
  <span class="guide-card__body">Open generated images and summary boards.</span>
</a>
</div>

## Run Metadata

<div class="guide-meta">
  <div class="guide-meta__item">
    <span class="guide-meta__label">Generated</span>
    `2026-03-07T23:32:15.506029+00:00`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Package</span>
    `github.com/smysnk/sikuligo/internal/grpcv1`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Target</span>
    `BenchmarkFindOnScreenE2E`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Benchtime</span>
    `200ms`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Count</span>
    `1`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Tags</span>
    `opencv gocv_specific_modules gocv_features2d gocv_calib3d`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Platform</span>
    `darwin/arm64` on `Apple M4 Pro`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Manifest</span>
    `docs/bench/find-on-screen-scenarios.example.json`
  </div>
</div>

## Engine Summary

| Engine | Cases | OK | Partial | Not Found | Unsupported | Error | Overlap Miss | Avg ms/op | Median ms/op | Best Scenario | Worst Scenario |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `akaze` | 120 | 39 | 0 | 78 | 0 | 0 | 3 | 172.121 | 147.695 | `photo_clutter_800x600_i02_rotate_s89976fd6` | `perspective_skew_sweep_1920x1080_i06_perspective_s6cc01a57` |
| `brisk` | 120 | 47 | 0 | 63 | 0 | 0 | 10 | 388.483 | 123.118 | `template_control_exact_800x600_i08_sb012a65c` | `noise_stress_random_1920x1080_i04_rotate_s17ff8896` |
| `hybrid` | 120 | 69 | 0 | 45 | 0 | 0 | 6 | 171.017 | 134.411 | `template_control_exact_800x600_i08_sb012a65c` | `perspective_skew_sweep_1920x1080_i06_perspective_s6cc01a57` |
| `kaze` | 120 | 63 | 0 | 50 | 0 | 0 | 7 | 824.898 | 640.512 | `vector_ui_baseline_800x600_i01_scale_s57c08455` | `perspective_skew_sweep_1920x1080_i06_perspective_s6cc01a57` |
| `orb` | 120 | 13 | 0 | 96 | 0 | 0 | 11 | 56.443 | 44.794 | `vector_ui_baseline_800x600_i01_scale_s57c08455` | `photo_clutter_1920x1080_i02_rotate_sa31cde71` |
| `sift` | 120 | 56 | 0 | 55 | 0 | 0 | 9 | 256.756 | 198.264 | `noise_stress_random_800x600_i04_rotate_s2892b109` | `perspective_skew_sweep_1920x1080_i06_perspective_s6cc01a57` |
| `template` | 120 | 64 | 0 | 56 | 0 | 0 | 0 | 154.257 | 114.466 | `template_control_exact_800x600_i08_sb012a65c` | `scale_rotate_sweep_1920x1080_i05_scale_s93bb9915` |

## Summary Metrics

| Engine | Cases | Rows | Avg ms/op | Median ms/op | Success % | False Positive % | No Match % | Unsupported % | Error % |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `akaze` | 120 | 40 | 172.121 | 147.695 | 32.5 | 2.5 | 65.0 | 0.0 | 0.0 |
| `brisk` | 120 | 40 | 388.483 | 123.118 | 39.2 | 8.3 | 52.5 | 0.0 | 0.0 |
| `hybrid` | 120 | 40 | 171.017 | 134.411 | 57.5 | 5.0 | 37.5 | 0.0 | 0.0 |
| `kaze` | 120 | 40 | 824.898 | 640.512 | 52.5 | 5.8 | 41.7 | 0.0 | 0.0 |
| `orb` | 120 | 40 | 56.443 | 44.794 | 10.8 | 9.2 | 80.0 | 0.0 | 0.0 |
| `sift` | 120 | 40 | 256.756 | 198.264 | 46.7 | 7.5 | 45.8 | 0.0 | 0.0 |
| `template` | 120 | 40 | 154.257 | 114.466 | 53.3 | 0.0 | 46.7 | 0.0 | 0.0 |

## Charts

![Performance chart]({{ '/bench/reports/find-on-screen-performance.svg' | relative_url }})

![Accuracy chart]({{ '/bench/reports/find-on-screen-accuracy.svg' | relative_url }})

![Scenario kind time chart]({{ '/bench/reports/find-on-screen-kind-time.svg' | relative_url }})

![Scenario kind success chart]({{ '/bench/reports/find-on-screen-kind-success.svg' | relative_url }})

![Resolution time chart]({{ '/bench/reports/find-on-screen-resolution-time.svg' | relative_url }})

![Resolution matches chart]({{ '/bench/reports/find-on-screen-resolution-matches.svg' | relative_url }})

![Resolution misses chart]({{ '/bench/reports/find-on-screen-resolution-misses.svg' | relative_url }})

![Resolution false positives chart]({{ '/bench/reports/find-on-screen-resolution-false-positives.svg' | relative_url }})

## Resolution Breakdown

| Resolution | Engine | Cases | Avg ms/op | Matches | Misses | False Positives |
| --- | --- | --- | --- | --- | --- | --- |
| `800x600` | `akaze` | 10 | 102.854 | 6 | 24 | 0 |
| `800x600` | `brisk` | 10 | 225.001 | 10 | 18 | 2 |
| `800x600` | `hybrid` | 10 | 77.853 | 17 | 12 | 1 |
| `800x600` | `kaze` | 10 | 353.025 | 14 | 16 | 0 |
| `800x600` | `orb` | 10 | 32.972 | 2 | 26 | 2 |
| `800x600` | `sift` | 10 | 127.999 | 13 | 16 | 1 |
| `800x600` | `template` | 10 | 66.950 | 16 | 14 | 0 |
| `1024x768` | `akaze` | 10 | 137.759 | 9 | 20 | 1 |
| `1024x768` | `brisk` | 10 | 270.581 | 10 | 16 | 4 |
| `1024x768` | `hybrid` | 10 | 127.208 | 17 | 11 | 2 |
| `1024x768` | `kaze` | 10 | 627.845 | 15 | 13 | 2 |
| `1024x768` | `orb` | 10 | 49.779 | 3 | 25 | 2 |
| `1024x768` | `sift` | 10 | 199.202 | 14 | 13 | 3 |
| `1024x768` | `template` | 10 | 110.610 | 16 | 14 | 0 |
| `1280x720` | `akaze` | 10 | 157.623 | 9 | 20 | 1 |
| `1280x720` | `brisk` | 10 | 314.513 | 12 | 15 | 3 |
| `1280x720` | `hybrid` | 10 | 151.516 | 16 | 12 | 2 |
| `1280x720` | `kaze` | 10 | 682.533 | 14 | 12 | 4 |
| `1280x720` | `orb` | 10 | 54.096 | 1 | 26 | 3 |
| `1280x720` | `sift` | 10 | 218.127 | 13 | 13 | 4 |
| `1280x720` | `template` | 10 | 132.035 | 16 | 14 | 0 |
| `1920x1080` | `akaze` | 10 | 290.249 | 15 | 14 | 1 |
| `1920x1080` | `brisk` | 10 | 743.838 | 15 | 14 | 1 |
| `1920x1080` | `hybrid` | 10 | 327.491 | 19 | 10 | 1 |
| `1920x1080` | `kaze` | 10 | 1636.191 | 20 | 9 | 1 |
| `1920x1080` | `orb` | 10 | 88.928 | 7 | 19 | 4 |
| `1920x1080` | `sift` | 10 | 481.695 | 16 | 13 | 1 |
| `1920x1080` | `template` | 10 | 307.431 | 16 | 14 | 0 |

## Scenario Kind Breakdown

| Kind | Engine | Cases | Avg ms/op | Success % | Not Found % | False Positive % |
| --- | --- | --- | --- | --- | --- | --- |
| `multi_monitor_dpi` | `akaze` | 4 | 204.595 | 0.0 | 100.0 | 0.0 |
| `multi_monitor_dpi` | `brisk` | 4 | 54.409 | 8.3 | 83.3 | 8.3 |
| `multi_monitor_dpi` | `hybrid` | 4 | 198.244 | 0.0 | 83.3 | 16.7 |
| `multi_monitor_dpi` | `kaze` | 4 | 976.377 | 8.3 | 83.3 | 8.3 |
| `multi_monitor_dpi` | `orb` | 4 | 28.499 | 0.0 | 83.3 | 16.7 |
| `multi_monitor_dpi` | `sift` | 4 | 274.505 | 0.0 | 91.7 | 8.3 |
| `multi_monitor_dpi` | `template` | 4 | 193.188 | 0.0 | 100.0 | 0.0 |
| `noise_stress` | `akaze` | 4 | 138.820 | 66.7 | 33.3 | 0.0 |
| `noise_stress` | `brisk` | 4 | 1345.607 | 83.3 | 16.7 | 0.0 |
| `noise_stress` | `hybrid` | 4 | 230.338 | 25.0 | 58.3 | 16.7 |
| `noise_stress` | `kaze` | 4 | 600.132 | 100.0 | 0.0 | 0.0 |
| `noise_stress` | `orb` | 4 | 78.654 | 25.0 | 58.3 | 16.7 |
| `noise_stress` | `sift` | 4 | 179.982 | 100.0 | 0.0 | 0.0 |
| `noise_stress` | `template` | 4 | 202.530 | 0.0 | 100.0 | 0.0 |
| `orb_feature_rich` | `akaze` | 4 | 195.933 | 0.0 | 100.0 | 0.0 |
| `orb_feature_rich` | `brisk` | 4 | 74.856 | 0.0 | 100.0 | 0.0 |
| `orb_feature_rich` | `hybrid` | 4 | 206.233 | 33.3 | 66.7 | 0.0 |
| `orb_feature_rich` | `kaze` | 4 | 977.411 | 8.3 | 91.7 | 0.0 |
| `orb_feature_rich` | `orb` | 4 | 39.425 | 0.0 | 100.0 | 0.0 |
| `orb_feature_rich` | `sift` | 4 | 263.661 | 16.7 | 83.3 | 0.0 |
| `orb_feature_rich` | `template` | 4 | 179.018 | 33.3 | 66.7 | 0.0 |
| `perspective_skew` | `akaze` | 8 | 207.635 | 4.2 | 95.8 | 0.0 |
| `perspective_skew` | `brisk` | 8 | 469.901 | 0.0 | 100.0 | 0.0 |
| `perspective_skew` | `hybrid` | 8 | 205.818 | 50.0 | 50.0 | 0.0 |
| `perspective_skew` | `kaze` | 8 | 960.689 | 37.5 | 62.5 | 0.0 |
| `perspective_skew` | `orb` | 8 | 63.492 | 4.2 | 95.8 | 0.0 |
| `perspective_skew` | `sift` | 8 | 365.885 | 4.2 | 95.8 | 0.0 |
| `perspective_skew` | `template` | 8 | 161.562 | 50.0 | 50.0 | 0.0 |
| `photographic` | `akaze` | 4 | 147.158 | 100.0 | 0.0 | 0.0 |
| `photographic` | `brisk` | 4 | 996.318 | 100.0 | 0.0 | 0.0 |
| `photographic` | `hybrid` | 4 | 101.079 | 100.0 | 0.0 | 0.0 |
| `photographic` | `kaze` | 4 | 1114.282 | 100.0 | 0.0 | 0.0 |
| `photographic` | `orb` | 4 | 130.924 | 0.0 | 83.3 | 16.7 |
| `photographic` | `sift` | 4 | 338.407 | 100.0 | 0.0 | 0.0 |
| `photographic` | `template` | 4 | 100.317 | 100.0 | 0.0 | 0.0 |
| `repetitive_grid` | `akaze` | 4 | 147.038 | 41.7 | 58.3 | 0.0 |
| `repetitive_grid` | `brisk` | 4 | 130.125 | 83.3 | 8.3 | 8.3 |
| `repetitive_grid` | `hybrid` | 4 | 114.883 | 100.0 | 0.0 | 0.0 |
| `repetitive_grid` | `kaze` | 4 | 613.537 | 75.0 | 25.0 | 0.0 |
| `repetitive_grid` | `orb` | 4 | 52.358 | 0.0 | 91.7 | 8.3 |
| `repetitive_grid` | `sift` | 4 | 188.392 | 91.7 | 8.3 | 0.0 |
| `repetitive_grid` | `template` | 4 | 113.758 | 100.0 | 0.0 | 0.0 |
| `scale_rotate` | `akaze` | 4 | 163.147 | 16.7 | 58.3 | 25.0 |
| `scale_rotate` | `brisk` | 4 | 202.417 | 25.0 | 16.7 | 58.3 |
| `scale_rotate` | `hybrid` | 4 | 238.930 | 16.7 | 66.7 | 16.7 |
| `scale_rotate` | `kaze` | 4 | 582.282 | 25.0 | 25.0 | 50.0 |
| `scale_rotate` | `orb` | 4 | 50.519 | 16.7 | 66.7 | 16.7 |
| `scale_rotate` | `sift` | 4 | 181.435 | 25.0 | 8.3 | 66.7 |
| `scale_rotate` | `template` | 4 | 222.800 | 0.0 | 100.0 | 0.0 |
| `template_control` | `akaze` | 4 | 175.418 | 25.0 | 75.0 | 0.0 |
| `template_control` | `brisk` | 4 | 50.894 | 33.3 | 66.7 | 0.0 |
| `template_control` | `hybrid` | 4 | 100.134 | 100.0 | 0.0 | 0.0 |
| `template_control` | `kaze` | 4 | 920.839 | 33.3 | 66.7 | 0.0 |
| `template_control` | `orb` | 4 | 28.586 | 8.3 | 91.7 | 0.0 |
| `template_control` | `sift` | 4 | 238.808 | 33.3 | 66.7 | 0.0 |
| `template_control` | `template` | 4 | 100.624 | 100.0 | 0.0 | 0.0 |
| `vector_ui` | `akaze` | 4 | 133.833 | 66.7 | 33.3 | 0.0 |
| `vector_ui` | `brisk` | 4 | 90.407 | 58.3 | 33.3 | 8.3 |
| `vector_ui` | `hybrid` | 4 | 108.691 | 100.0 | 0.0 | 0.0 |
| `vector_ui` | `kaze` | 4 | 542.745 | 100.0 | 0.0 | 0.0 |
| `vector_ui` | `orb` | 4 | 28.485 | 50.0 | 33.3 | 16.7 |
| `vector_ui` | `sift` | 4 | 170.600 | 91.7 | 8.3 | 0.0 |
| `vector_ui` | `template` | 4 | 107.204 | 100.0 | 0.0 | 0.0 |

## Resolution + Kind Breakdown

| Resolution | Kind | Engine | Cases | Avg ms/op | Success % | Not Found % | False Positive % |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `800x600` | `multi_monitor_dpi` | `akaze` | 1 | 104.000 | 0.0 | 100.0 | 0.0 |
| `800x600` | `multi_monitor_dpi` | `brisk` | 1 | 43.261 | 0.0 | 100.0 | 0.0 |
| `800x600` | `multi_monitor_dpi` | `hybrid` | 1 | 93.550 | 0.0 | 100.0 | 0.0 |
| `800x600` | `multi_monitor_dpi` | `kaze` | 1 | 424.733 | 0.0 | 100.0 | 0.0 |
| `800x600` | `multi_monitor_dpi` | `orb` | 1 | 17.768 | 0.0 | 100.0 | 0.0 |
| `800x600` | `multi_monitor_dpi` | `sift` | 1 | 129.833 | 0.0 | 100.0 | 0.0 |
| `800x600` | `multi_monitor_dpi` | `template` | 1 | 78.801 | 0.0 | 100.0 | 0.0 |
| `800x600` | `noise_stress` | `akaze` | 1 | 123.367 | 33.3 | 66.7 | 0.0 |
| `800x600` | `noise_stress` | `brisk` | 1 | 638.611 | 33.3 | 66.7 | 0.0 |
| `800x600` | `noise_stress` | `hybrid` | 1 | 117.685 | 0.0 | 66.7 | 33.3 |
| `800x600` | `noise_stress` | `kaze` | 1 | 238.426 | 100.0 | 0.0 | 0.0 |
| `800x600` | `noise_stress` | `orb` | 1 | 45.050 | 0.0 | 66.7 | 33.3 |
| `800x600` | `noise_stress` | `sift` | 1 | 84.590 | 100.0 | 0.0 | 0.0 |
| `800x600` | `noise_stress` | `template` | 1 | 87.716 | 0.0 | 100.0 | 0.0 |
| `800x600` | `orb_feature_rich` | `akaze` | 1 | 97.409 | 0.0 | 100.0 | 0.0 |
| `800x600` | `orb_feature_rich` | `brisk` | 1 | 52.525 | 0.0 | 100.0 | 0.0 |
| `800x600` | `orb_feature_rich` | `hybrid` | 1 | 91.094 | 33.3 | 66.7 | 0.0 |
| `800x600` | `orb_feature_rich` | `kaze` | 1 | 425.584 | 0.0 | 100.0 | 0.0 |
| `800x600` | `orb_feature_rich` | `orb` | 1 | 21.164 | 0.0 | 100.0 | 0.0 |
| `800x600` | `orb_feature_rich` | `sift` | 1 | 133.134 | 0.0 | 100.0 | 0.0 |
| `800x600` | `orb_feature_rich` | `template` | 1 | 77.782 | 33.3 | 66.7 | 0.0 |
| `800x600` | `perspective_skew` | `akaze` | 2 | 105.089 | 0.0 | 100.0 | 0.0 |
| `800x600` | `perspective_skew` | `brisk` | 2 | 266.538 | 0.0 | 100.0 | 0.0 |
| `800x600` | `perspective_skew` | `hybrid` | 2 | 92.351 | 50.0 | 50.0 | 0.0 |
| `800x600` | `perspective_skew` | `kaze` | 2 | 424.219 | 16.7 | 83.3 | 0.0 |
| `800x600` | `perspective_skew` | `orb` | 2 | 34.178 | 0.0 | 100.0 | 0.0 |
| `800x600` | `perspective_skew` | `sift` | 2 | 151.949 | 16.7 | 83.3 | 0.0 |
| `800x600` | `perspective_skew` | `template` | 2 | 70.480 | 50.0 | 50.0 | 0.0 |
| `800x600` | `photographic` | `akaze` | 1 | 71.644 | 100.0 | 0.0 | 0.0 |
| `800x600` | `photographic` | `brisk` | 1 | 640.692 | 100.0 | 0.0 | 0.0 |
| `800x600` | `photographic` | `hybrid` | 1 | 44.352 | 100.0 | 0.0 | 0.0 |
| `800x600` | `photographic` | `kaze` | 1 | 367.949 | 100.0 | 0.0 | 0.0 |
| `800x600` | `photographic` | `orb` | 1 | 82.962 | 0.0 | 100.0 | 0.0 |
| `800x600` | `photographic` | `sift` | 1 | 183.902 | 100.0 | 0.0 | 0.0 |
| `800x600` | `photographic` | `template` | 1 | 45.318 | 100.0 | 0.0 | 0.0 |
| `800x600` | `repetitive_grid` | `akaze` | 1 | 109.760 | 0.0 | 100.0 | 0.0 |
| `800x600` | `repetitive_grid` | `brisk` | 1 | 91.307 | 66.7 | 0.0 | 33.3 |
| `800x600` | `repetitive_grid` | `hybrid` | 1 | 47.043 | 100.0 | 0.0 | 0.0 |
| `800x600` | `repetitive_grid` | `kaze` | 1 | 294.739 | 66.7 | 33.3 | 0.0 |
| `800x600` | `repetitive_grid` | `orb` | 1 | 33.426 | 0.0 | 100.0 | 0.0 |
| `800x600` | `repetitive_grid` | `sift` | 1 | 113.159 | 66.7 | 33.3 | 0.0 |
| `800x600` | `repetitive_grid` | `template` | 1 | 48.591 | 100.0 | 0.0 | 0.0 |
| `800x600` | `scale_rotate` | `akaze` | 1 | 115.360 | 33.3 | 66.7 | 0.0 |
| `800x600` | `scale_rotate` | `brisk` | 1 | 138.919 | 33.3 | 33.3 | 33.3 |
| `800x600` | `scale_rotate` | `hybrid` | 1 | 109.393 | 33.3 | 66.7 | 0.0 |
| `800x600` | `scale_rotate` | `kaze` | 1 | 353.617 | 33.3 | 66.7 | 0.0 |
| `800x600` | `scale_rotate` | `orb` | 1 | 26.904 | 33.3 | 66.7 | 0.0 |
| `800x600` | `scale_rotate` | `sift` | 1 | 102.557 | 33.3 | 33.3 | 33.3 |
| `800x600` | `scale_rotate` | `template` | 1 | 101.483 | 0.0 | 100.0 | 0.0 |
| `800x600` | `template_control` | `akaze` | 1 | 103.059 | 0.0 | 100.0 | 0.0 |
| `800x600` | `template_control` | `brisk` | 1 | 38.168 | 33.3 | 66.7 | 0.0 |
| `800x600` | `template_control` | `hybrid` | 1 | 41.339 | 100.0 | 0.0 | 0.0 |
| `800x600` | `template_control` | `kaze` | 1 | 348.177 | 33.3 | 66.7 | 0.0 |
| `800x600` | `template_control` | `orb` | 1 | 17.223 | 0.0 | 100.0 | 0.0 |
| `800x600` | `template_control` | `sift` | 1 | 107.185 | 33.3 | 66.7 | 0.0 |
| `800x600` | `template_control` | `template` | 1 | 42.962 | 100.0 | 0.0 | 0.0 |
| `800x600` | `vector_ui` | `akaze` | 1 | 93.762 | 33.3 | 66.7 | 0.0 |
| `800x600` | `vector_ui` | `brisk` | 1 | 73.452 | 66.7 | 33.3 | 0.0 |
| `800x600` | `vector_ui` | `hybrid` | 1 | 49.370 | 100.0 | 0.0 | 0.0 |
| `800x600` | `vector_ui` | `kaze` | 1 | 228.583 | 100.0 | 0.0 | 0.0 |
| `800x600` | `vector_ui` | `orb` | 1 | 16.863 | 33.3 | 33.3 | 33.3 |
| `800x600` | `vector_ui` | `sift` | 1 | 121.732 | 66.7 | 33.3 | 0.0 |
| `800x600` | `vector_ui` | `template` | 1 | 45.888 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `multi_monitor_dpi` | `akaze` | 1 | 160.789 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `multi_monitor_dpi` | `brisk` | 1 | 56.359 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `multi_monitor_dpi` | `hybrid` | 1 | 133.637 | 0.0 | 66.7 | 33.3 |
| `1024x768` | `multi_monitor_dpi` | `kaze` | 1 | 807.870 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `multi_monitor_dpi` | `orb` | 1 | 22.277 | 0.0 | 66.7 | 33.3 |
| `1024x768` | `multi_monitor_dpi` | `sift` | 1 | 227.169 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `multi_monitor_dpi` | `template` | 1 | 134.919 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `noise_stress` | `akaze` | 1 | 110.782 | 66.7 | 33.3 | 0.0 |
| `1024x768` | `noise_stress` | `brisk` | 1 | 584.451 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `noise_stress` | `hybrid` | 1 | 187.895 | 33.3 | 66.7 | 0.0 |
| `1024x768` | `noise_stress` | `kaze` | 1 | 437.172 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `noise_stress` | `orb` | 1 | 69.752 | 33.3 | 66.7 | 0.0 |
| `1024x768` | `noise_stress` | `sift` | 1 | 145.630 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `noise_stress` | `template` | 1 | 142.395 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `orb_feature_rich` | `akaze` | 1 | 148.733 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `orb_feature_rich` | `brisk` | 1 | 66.857 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `orb_feature_rich` | `hybrid` | 1 | 150.871 | 33.3 | 66.7 | 0.0 |
| `1024x768` | `orb_feature_rich` | `kaze` | 1 | 777.856 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `orb_feature_rich` | `orb` | 1 | 31.582 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `orb_feature_rich` | `sift` | 1 | 194.033 | 33.3 | 66.7 | 0.0 |
| `1024x768` | `orb_feature_rich` | `template` | 1 | 127.450 | 33.3 | 66.7 | 0.0 |
| `1024x768` | `perspective_skew` | `akaze` | 2 | 161.179 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `perspective_skew` | `brisk` | 2 | 357.684 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `perspective_skew` | `hybrid` | 2 | 151.894 | 50.0 | 50.0 | 0.0 |
| `1024x768` | `perspective_skew` | `kaze` | 2 | 661.903 | 50.0 | 50.0 | 0.0 |
| `1024x768` | `perspective_skew` | `orb` | 2 | 51.840 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `perspective_skew` | `sift` | 2 | 291.130 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `perspective_skew` | `template` | 2 | 115.161 | 50.0 | 50.0 | 0.0 |
| `1024x768` | `photographic` | `akaze` | 1 | 118.957 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `photographic` | `brisk` | 1 | 877.598 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `photographic` | `hybrid` | 1 | 75.246 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `photographic` | `kaze` | 1 | 744.107 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `photographic` | `orb` | 1 | 133.762 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `photographic` | `sift` | 1 | 250.309 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `photographic` | `template` | 1 | 73.443 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `repetitive_grid` | `akaze` | 1 | 133.518 | 33.3 | 66.7 | 0.0 |
| `1024x768` | `repetitive_grid` | `brisk` | 1 | 131.239 | 66.7 | 33.3 | 0.0 |
| `1024x768` | `repetitive_grid` | `hybrid` | 1 | 88.787 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `repetitive_grid` | `kaze` | 1 | 554.872 | 66.7 | 33.3 | 0.0 |
| `1024x768` | `repetitive_grid` | `orb` | 1 | 45.839 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `repetitive_grid` | `sift` | 1 | 139.990 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `repetitive_grid` | `template` | 1 | 88.904 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `scale_rotate` | `akaze` | 1 | 130.809 | 0.0 | 66.7 | 33.3 |
| `1024x768` | `scale_rotate` | `brisk` | 1 | 144.116 | 0.0 | 0.0 | 100.0 |
| `1024x768` | `scale_rotate` | `hybrid` | 1 | 178.792 | 0.0 | 66.7 | 33.3 |
| `1024x768` | `scale_rotate` | `kaze` | 1 | 546.474 | 0.0 | 33.3 | 66.7 |
| `1024x768` | `scale_rotate` | `orb` | 1 | 41.364 | 0.0 | 66.7 | 33.3 |
| `1024x768` | `scale_rotate` | `sift` | 1 | 134.335 | 0.0 | 0.0 | 100.0 |
| `1024x768` | `scale_rotate` | `template` | 1 | 161.292 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `template_control` | `akaze` | 1 | 139.871 | 33.3 | 66.7 | 0.0 |
| `1024x768` | `template_control` | `brisk` | 1 | 45.542 | 33.3 | 66.7 | 0.0 |
| `1024x768` | `template_control` | `hybrid` | 1 | 72.154 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `template_control` | `kaze` | 1 | 655.391 | 33.3 | 66.7 | 0.0 |
| `1024x768` | `template_control` | `orb` | 1 | 25.096 | 0.0 | 100.0 | 0.0 |
| `1024x768` | `template_control` | `sift` | 1 | 188.121 | 33.3 | 66.7 | 0.0 |
| `1024x768` | `template_control` | `template` | 1 | 71.231 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `vector_ui` | `akaze` | 1 | 111.770 | 66.7 | 33.3 | 0.0 |
| `1024x768` | `vector_ui` | `brisk` | 1 | 84.281 | 33.3 | 33.3 | 33.3 |
| `1024x768` | `vector_ui` | `hybrid` | 1 | 80.904 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `vector_ui` | `kaze` | 1 | 430.900 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `vector_ui` | `orb` | 1 | 24.437 | 66.7 | 33.3 | 0.0 |
| `1024x768` | `vector_ui` | `sift` | 1 | 130.174 | 100.0 | 0.0 | 0.0 |
| `1024x768` | `vector_ui` | `template` | 1 | 76.143 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `multi_monitor_dpi` | `akaze` | 1 | 177.019 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `multi_monitor_dpi` | `brisk` | 1 | 48.794 | 0.0 | 66.7 | 33.3 |
| `1280x720` | `multi_monitor_dpi` | `hybrid` | 1 | 194.301 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `multi_monitor_dpi` | `kaze` | 1 | 718.669 | 0.0 | 66.7 | 33.3 |
| `1280x720` | `multi_monitor_dpi` | `orb` | 1 | 27.472 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `multi_monitor_dpi` | `sift` | 1 | 205.642 | 0.0 | 66.7 | 33.3 |
| `1280x720` | `multi_monitor_dpi` | `template` | 1 | 170.140 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `noise_stress` | `akaze` | 1 | 125.038 | 66.7 | 33.3 | 0.0 |
| `1280x720` | `noise_stress` | `brisk` | 1 | 619.543 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `noise_stress` | `hybrid` | 1 | 209.925 | 0.0 | 66.7 | 33.3 |
| `1280x720` | `noise_stress` | `kaze` | 1 | 494.669 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `noise_stress` | `orb` | 1 | 73.740 | 0.0 | 66.7 | 33.3 |
| `1280x720` | `noise_stress` | `sift` | 1 | 154.148 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `noise_stress` | `template` | 1 | 167.037 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `orb_feature_rich` | `akaze` | 1 | 183.120 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `orb_feature_rich` | `brisk` | 1 | 68.857 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `orb_feature_rich` | `hybrid` | 1 | 180.077 | 33.3 | 66.7 | 0.0 |
| `1280x720` | `orb_feature_rich` | `kaze` | 1 | 858.105 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `orb_feature_rich` | `orb` | 1 | 34.248 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `orb_feature_rich` | `sift` | 1 | 246.567 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `orb_feature_rich` | `template` | 1 | 156.729 | 33.3 | 66.7 | 0.0 |
| `1280x720` | `perspective_skew` | `akaze` | 2 | 191.217 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `perspective_skew` | `brisk` | 2 | 458.157 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `perspective_skew` | `hybrid` | 2 | 186.480 | 50.0 | 50.0 | 0.0 |
| `1280x720` | `perspective_skew` | `kaze` | 2 | 816.047 | 33.3 | 66.7 | 0.0 |
| `1280x720` | `perspective_skew` | `orb` | 2 | 57.282 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `perspective_skew` | `sift` | 2 | 317.121 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `perspective_skew` | `template` | 2 | 140.050 | 50.0 | 50.0 | 0.0 |
| `1280x720` | `photographic` | `akaze` | 1 | 124.443 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `photographic` | `brisk` | 1 | 996.510 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `photographic` | `hybrid` | 1 | 84.439 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `photographic` | `kaze` | 1 | 857.750 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `photographic` | `orb` | 1 | 136.591 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `photographic` | `sift` | 1 | 274.028 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `photographic` | `template` | 1 | 85.594 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `repetitive_grid` | `akaze` | 1 | 152.873 | 33.3 | 66.7 | 0.0 |
| `1280x720` | `repetitive_grid` | `brisk` | 1 | 134.760 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `repetitive_grid` | `hybrid` | 1 | 95.281 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `repetitive_grid` | `kaze` | 1 | 625.634 | 66.7 | 33.3 | 0.0 |
| `1280x720` | `repetitive_grid` | `orb` | 1 | 55.275 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `repetitive_grid` | `sift` | 1 | 177.677 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `repetitive_grid` | `template` | 1 | 95.223 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `scale_rotate` | `akaze` | 1 | 155.637 | 0.0 | 66.7 | 33.3 |
| `1280x720` | `scale_rotate` | `brisk` | 1 | 219.295 | 0.0 | 33.3 | 66.7 |
| `1280x720` | `scale_rotate` | `hybrid` | 1 | 204.396 | 0.0 | 66.7 | 33.3 |
| `1280x720` | `scale_rotate` | `kaze` | 1 | 457.675 | 0.0 | 0.0 | 100.0 |
| `1280x720` | `scale_rotate` | `orb` | 1 | 45.238 | 0.0 | 66.7 | 33.3 |
| `1280x720` | `scale_rotate` | `sift` | 1 | 145.077 | 0.0 | 0.0 | 100.0 |
| `1280x720` | `scale_rotate` | `template` | 1 | 186.018 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `template_control` | `akaze` | 1 | 146.657 | 33.3 | 66.7 | 0.0 |
| `1280x720` | `template_control` | `brisk` | 1 | 52.158 | 33.3 | 66.7 | 0.0 |
| `1280x720` | `template_control` | `hybrid` | 1 | 87.506 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `template_control` | `kaze` | 1 | 713.100 | 33.3 | 66.7 | 0.0 |
| `1280x720` | `template_control` | `orb` | 1 | 27.485 | 0.0 | 100.0 | 0.0 |
| `1280x720` | `template_control` | `sift` | 1 | 202.494 | 33.3 | 66.7 | 0.0 |
| `1280x720` | `template_control` | `template` | 1 | 90.819 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `vector_ui` | `akaze` | 1 | 129.010 | 66.7 | 33.3 | 0.0 |
| `1280x720` | `vector_ui` | `brisk` | 1 | 88.900 | 66.7 | 33.3 | 0.0 |
| `1280x720` | `vector_ui` | `hybrid` | 1 | 86.270 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `vector_ui` | `kaze` | 1 | 467.631 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `vector_ui` | `orb` | 1 | 26.342 | 33.3 | 33.3 | 33.3 |
| `1280x720` | `vector_ui` | `sift` | 1 | 141.396 | 100.0 | 0.0 | 0.0 |
| `1280x720` | `vector_ui` | `template` | 1 | 88.689 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `multi_monitor_dpi` | `akaze` | 1 | 376.572 | 0.0 | 100.0 | 0.0 |
| `1920x1080` | `multi_monitor_dpi` | `brisk` | 1 | 69.223 | 33.3 | 66.7 | 0.0 |
| `1920x1080` | `multi_monitor_dpi` | `hybrid` | 1 | 371.488 | 0.0 | 66.7 | 33.3 |
| `1920x1080` | `multi_monitor_dpi` | `kaze` | 1 | 1954.237 | 33.3 | 66.7 | 0.0 |
| `1920x1080` | `multi_monitor_dpi` | `orb` | 1 | 46.478 | 0.0 | 66.7 | 33.3 |
| `1920x1080` | `multi_monitor_dpi` | `sift` | 1 | 535.374 | 0.0 | 100.0 | 0.0 |
| `1920x1080` | `multi_monitor_dpi` | `template` | 1 | 388.894 | 0.0 | 100.0 | 0.0 |
| `1920x1080` | `noise_stress` | `akaze` | 1 | 196.092 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `noise_stress` | `brisk` | 1 | 3539.823 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `noise_stress` | `hybrid` | 1 | 405.846 | 66.7 | 33.3 | 0.0 |
| `1920x1080` | `noise_stress` | `kaze` | 1 | 1230.262 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `noise_stress` | `orb` | 1 | 126.076 | 66.7 | 33.3 | 0.0 |
| `1920x1080` | `noise_stress` | `sift` | 1 | 335.558 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `noise_stress` | `template` | 1 | 412.973 | 0.0 | 100.0 | 0.0 |
| `1920x1080` | `orb_feature_rich` | `akaze` | 1 | 354.472 | 0.0 | 100.0 | 0.0 |
| `1920x1080` | `orb_feature_rich` | `brisk` | 1 | 111.184 | 0.0 | 100.0 | 0.0 |
| `1920x1080` | `orb_feature_rich` | `hybrid` | 1 | 402.890 | 33.3 | 66.7 | 0.0 |
| `1920x1080` | `orb_feature_rich` | `kaze` | 1 | 1848.098 | 33.3 | 66.7 | 0.0 |
| `1920x1080` | `orb_feature_rich` | `orb` | 1 | 70.705 | 0.0 | 100.0 | 0.0 |
| `1920x1080` | `orb_feature_rich` | `sift` | 1 | 480.909 | 33.3 | 66.7 | 0.0 |
| `1920x1080` | `orb_feature_rich` | `template` | 1 | 354.110 | 33.3 | 66.7 | 0.0 |
| `1920x1080` | `perspective_skew` | `akaze` | 2 | 373.054 | 16.7 | 83.3 | 0.0 |
| `1920x1080` | `perspective_skew` | `brisk` | 2 | 797.224 | 0.0 | 100.0 | 0.0 |
| `1920x1080` | `perspective_skew` | `hybrid` | 2 | 392.547 | 50.0 | 50.0 | 0.0 |
| `1920x1080` | `perspective_skew` | `kaze` | 2 | 1940.587 | 50.0 | 50.0 | 0.0 |
| `1920x1080` | `perspective_skew` | `orb` | 2 | 110.667 | 16.7 | 83.3 | 0.0 |
| `1920x1080` | `perspective_skew` | `sift` | 2 | 703.339 | 0.0 | 100.0 | 0.0 |
| `1920x1080` | `perspective_skew` | `template` | 2 | 320.558 | 50.0 | 50.0 | 0.0 |
| `1920x1080` | `photographic` | `akaze` | 1 | 273.588 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `photographic` | `brisk` | 1 | 1470.471 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `photographic` | `hybrid` | 1 | 200.278 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `photographic` | `kaze` | 1 | 2487.322 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `photographic` | `orb` | 1 | 170.380 | 0.0 | 33.3 | 66.7 |
| `1920x1080` | `photographic` | `sift` | 1 | 645.391 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `photographic` | `template` | 1 | 196.911 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `repetitive_grid` | `akaze` | 1 | 192.004 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `repetitive_grid` | `brisk` | 1 | 163.192 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `repetitive_grid` | `hybrid` | 1 | 228.421 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `repetitive_grid` | `kaze` | 1 | 978.902 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `repetitive_grid` | `orb` | 1 | 74.892 | 0.0 | 66.7 | 33.3 |
| `1920x1080` | `repetitive_grid` | `sift` | 1 | 322.742 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `repetitive_grid` | `template` | 1 | 222.315 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `scale_rotate` | `akaze` | 1 | 250.782 | 33.3 | 33.3 | 33.3 |
| `1920x1080` | `scale_rotate` | `brisk` | 1 | 307.339 | 66.7 | 0.0 | 33.3 |
| `1920x1080` | `scale_rotate` | `hybrid` | 1 | 463.139 | 33.3 | 66.7 | 0.0 |
| `1920x1080` | `scale_rotate` | `kaze` | 1 | 971.362 | 66.7 | 0.0 | 33.3 |
| `1920x1080` | `scale_rotate` | `orb` | 1 | 88.571 | 33.3 | 66.7 | 0.0 |
| `1920x1080` | `scale_rotate` | `sift` | 1 | 343.772 | 66.7 | 0.0 | 33.3 |
| `1920x1080` | `scale_rotate` | `template` | 1 | 442.409 | 0.0 | 100.0 | 0.0 |
| `1920x1080` | `template_control` | `akaze` | 1 | 312.083 | 33.3 | 66.7 | 0.0 |
| `1920x1080` | `template_control` | `brisk` | 1 | 67.705 | 33.3 | 66.7 | 0.0 |
| `1920x1080` | `template_control` | `hybrid` | 1 | 199.535 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `template_control` | `kaze` | 1 | 1966.688 | 33.3 | 66.7 | 0.0 |
| `1920x1080` | `template_control` | `orb` | 1 | 44.539 | 33.3 | 66.7 | 0.0 |
| `1920x1080` | `template_control` | `sift` | 1 | 457.429 | 33.3 | 66.7 | 0.0 |
| `1920x1080` | `template_control` | `template` | 1 | 197.485 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `vector_ui` | `akaze` | 1 | 200.790 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `vector_ui` | `brisk` | 1 | 114.997 | 66.7 | 33.3 | 0.0 |
| `1920x1080` | `vector_ui` | `hybrid` | 1 | 218.219 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `vector_ui` | `kaze` | 1 | 1043.864 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `vector_ui` | `orb` | 1 | 46.300 | 66.7 | 33.3 | 0.0 |
| `1920x1080` | `vector_ui` | `sift` | 1 | 289.096 | 100.0 | 0.0 | 0.0 |
| `1920x1080` | `vector_ui` | `template` | 1 | 218.099 | 100.0 | 0.0 | 0.0 |

## Slowest Benchmark Rows

| Scenario | Engine | Status | Iterations | ms/op | Bytes/op | Allocs/op |
| --- | --- | --- | --- | --- | --- | --- |
| `noise_stress_random_1920x1080_i04_rotate_s17ff8896` | `brisk` | `ok` | 1 | 3539.823 | 20629608 | 1317 |
| `perspective_skew_sweep_1920x1080_i06_perspective_s6cc01a57` | `kaze` | `not_found` | 1 | 2669.433 | 19990440 | 3568 |
| `photo_clutter_1920x1080_i02_rotate_sa31cde71` | `kaze` | `ok` | 1 | 2487.322 | 16496176 | 2731 |
| `template_control_exact_1920x1080_i08_s60d1e717` | `kaze` | `not_found` | 1 | 1966.688 | 13674256 | 1316 |
| `multi_monitor_dpi_shift_1920x1080_i10_scale_s6a949f54` | `kaze` | `not_found` | 1 | 1954.237 | 11704360 | 1230 |
| `orb_feature_rich_1920x1080_i07_rotate_s2be28017` | `kaze` | `not_found` | 1 | 1848.098 | 11063840 | 1167 |
| `photo_clutter_1920x1080_i02_rotate_sa31cde71` | `brisk` | `ok` | 1 | 1470.471 | 17409256 | 2807 |
| `perspective_skew_sweep_1920x1080_i06_perspective_s6cc01a57` | `brisk` | `not_found` | 1 | 1448.909 | 24326360 | 6230 |
| `noise_stress_random_1920x1080_i04_rotate_s17ff8896` | `kaze` | `ok` | 1 | 1230.262 | 10172784 | 1310 |
| `hybrid_gate_conflicts_1920x1080_i09_perspective_s5a3c19a0` | `kaze` | `ok` | 1 | 1211.741 | 7892032 | 857 |
| `perspective_skew_sweep_1280x720_i06_perspective_s033dc796` | `kaze` | `not_found` | 1 | 1056.743 | 8557808 | 2238 |
| `vector_ui_baseline_1920x1080_i01_rotate_s2332e8ba` | `kaze` | `ok` | 1 | 1043.864 | 8528040 | 1196 |

## Raw Artifacts

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
