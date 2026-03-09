---
layout: guide
title: FindOnScreen Scenario Strategy
nav_key: benchmarks
kicker: Benchmarks
lead: Scenario-corpus design, engine comparison context, and generated visual examples for the current benchmark set.
---

<div class="guide-grid">
<a class="guide-card" href="{{ '/bench' | relative_url }}">
  <span class="guide-card__eyebrow">Overview</span>
  <span class="guide-card__title">Benchmark Overview</span>
  <span class="guide-card__body">Return to the benchmark section summary.</span>
</a>
<a class="guide-card" href="{{ '/bench/reports/find-on-screen-e2e' | relative_url }}">
  <span class="guide-card__eyebrow">Detailed Report</span>
  <span class="guide-card__title">E2E Results</span>
  <span class="guide-card__body">Compare strategy guidance against the latest benchmark outcome.</span>
</a>
<a class="guide-card" href="{{ '/bench/FIND_ON_SCREEN_SCENARIO_INTENT' | relative_url }}">
  <span class="guide-card__eyebrow">Scenario Docs</span>
  <span class="guide-card__title">Scenario Intent</span>
  <span class="guide-card__body">Review what each scenario is intended to prove.</span>
</a>
<a class="guide-card guide-card--subtle" href="{{ '/bench/FIND_ON_SCREEN_SCENARIO_SCHEMA' | relative_url }}">
  <span class="guide-card__eyebrow">Schema</span>
  <span class="guide-card__title">Scenario Schema</span>
  <span class="guide-card__body">Inspect manifest structure and region-selection workflow.</span>
</a>
</div>

## Strategy Metadata

<div class="guide-meta">
  <div class="guide-meta__item">
    <span class="guide-meta__label">Generated</span>
    `2026-03-07T23:32:16.066359+00:00`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Manifest</span>
    `docs/bench/find-on-screen-scenarios.example.json`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Schema Version</span>
    `1.0.0`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Engines</span>
    `template`, `orb`, `akaze`, `brisk`, `kaze`, `sift`, `hybrid`
  </div>
</div>

## Strategy Summary

| Metric | Value |
| --- | --- |
| Engine | `hybrid` |
| Engine Match Rate | `57.5%` |
| Other Engines Match Rate | `39.2%` |
| Delta vs Others | `18.3 pts` |
| Engine Rank | `1/7` |
| Benchmark Source | `docs/bench/reports/find-on-screen-e2e.json` |

## Visual Examples

<div class="guide-meta">
  <div class="guide-meta__item">
    <span class="guide-meta__label">Resolution</span>
    `1280x720`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Engine</span>
    `hybrid`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Log</span>
    `docs/bench/reports/strategy-visuals-1280x720/go-test.log`
  </div>
  <div class="guide-meta__item">
    <span class="guide-meta__label">Command</span>
    `go test ./internal/grpcv1 -run '^$' -bench 'BenchmarkFindOnScreenE2E/engine=hybrid/.*_1280x720_' -benchmem -benchtime 1x -count 1 -tags 'opencv gocv_specific_modules gocv_features2d gocv_calib3d'`
  </div>
</div>

| Scenario | Example |
| --- | --- |
| `vector_ui_baseline` | ![vector_ui_baseline]({{ '/bench/reports/strategy-visuals-1280x720/summaries/summary-vector_ui_baseline_1280x720_i01.png' | relative_url }}) |
| `photo_clutter` | ![photo_clutter]({{ '/bench/reports/strategy-visuals-1280x720/summaries/summary-photo_clutter_1280x720_i02.png' | relative_url }}) |
| `repetitive_grid_camouflage` | ![repetitive_grid_camouflage]({{ '/bench/reports/strategy-visuals-1280x720/summaries/summary-repetitive_grid_camouflage_1280x720_i03.png' | relative_url }}) |
| `noise_stress_random` | ![noise_stress_random]({{ '/bench/reports/strategy-visuals-1280x720/summaries/summary-noise_stress_random_1280x720_i04.png' | relative_url }}) |
| `scale_rotate_sweep` | ![scale_rotate_sweep]({{ '/bench/reports/strategy-visuals-1280x720/summaries/summary-scale_rotate_sweep_1280x720_i05.png' | relative_url }}) |
| `perspective_skew_sweep` | ![perspective_skew_sweep]({{ '/bench/reports/strategy-visuals-1280x720/summaries/summary-perspective_skew_sweep_1280x720_i06.png' | relative_url }}) |
| `orb_feature_rich` | ![orb_feature_rich]({{ '/bench/reports/strategy-visuals-1280x720/summaries/summary-orb_feature_rich_1280x720_i07.png' | relative_url }}) |
| `template_control_exact` | ![template_control_exact]({{ '/bench/reports/strategy-visuals-1280x720/summaries/summary-template_control_exact_1280x720_i08.png' | relative_url }}) |
| `hybrid_gate_conflicts` | ![hybrid_gate_conflicts]({{ '/bench/reports/strategy-visuals-1280x720/summaries/summary-hybrid_gate_conflicts_1280x720_i09.png' | relative_url }}) |
| `multi_monitor_dpi_shift` | ![multi_monitor_dpi_shift]({{ '/bench/reports/strategy-visuals-1280x720/summaries/summary-multi_monitor_dpi_shift_1280x720_i10.png' | relative_url }}) |

## Diversity Summary

| Metric | Value |
| --- | --- |
| Scenario Types | `10/10` |
| Resolution Groups | `4` |
| Scenarios Per Resolution | `10` |
| Expected Positive | `10` |
| Expected Negative | `0` |

### Kinds

| Metric | Value |
| --- | --- |
| `hybrid_gate` | `1` |
| `multi_monitor_dpi` | `1` |
| `noise_stress` | `1` |
| `orb_feature_rich` | `1` |
| `perspective_skew` | `1` |
| `photographic` | `1` |
| `repetitive_grid` | `1` |
| `scale_rotate` | `1` |
| `template_control` | `1` |
| `vector_ui` | `1` |

### Styles

| Metric | Value |
| --- | --- |
| `grid` | `1` |
| `mixed` | `4` |
| `noise` | `1` |
| `orbtex` | `1` |
| `photo` | `1` |
| `ui` | `1` |
| `vector` | `1` |

### Target Sources

| Metric | Value |
| --- | --- |
| `asset` | `1` |
| `mixed` | `4` |
| `synthetic` | `5` |

### Decoy Placements

| Metric | Value |
| --- | --- |
| `clustered` | `1` |
| `grid` | `3` |
| `mixed` | `4` |
| `random` | `2` |

### Noise Types

| Metric | Value |
| --- | --- |
| `banding` | `2` |
| `compression_blocks` | `5` |
| `gaussian` | `5` |
| `poisson` | `1` |
| `salt_pepper` | `1` |

### Transform Coverage

| Metric | Value |
| --- | --- |
| `perspective_enabled` | `3` |
| `rotate` | `10` |
| `scale` | `10` |
| `skew_x_nonzero` | `8` |
| `skew_y_nonzero` | `8` |

## Scenario Intent

| Scenario ID | Kind | Style | Looking For |
| --- | --- | --- | --- |
| `vector_ui_baseline` | `vector_ui` | `vector` | vector_ui_baseline: vector_ui (positive=True, iou_min=0.92, area_max=1.25, partial=False, decoys=grid, occlusion=False) |
| `photo_clutter` | `photographic` | `photo` | photo_clutter: photographic (positive=True, iou_min=0.75, area_max=1.6, partial=True, decoys=mixed, occlusion=True) |
| `repetitive_grid_camouflage` | `repetitive_grid` | `grid` | repetitive_grid_camouflage: repetitive_grid (positive=True, iou_min=0.85, area_max=1.35, partial=False, decoys=grid, occlusion=False) |
| `noise_stress_random` | `noise_stress` | `noise` | noise_stress_random: noise_stress (positive=True, iou_min=0.7, area_max=1.8, partial=True, decoys=random, occlusion=True) |
| `scale_rotate_sweep` | `scale_rotate` | `mixed` | scale_rotate_sweep: scale_rotate (positive=True, iou_min=0.68, area_max=2.2, partial=True, decoys=mixed, occlusion=True) |
| `perspective_skew_sweep` | `perspective_skew` | `mixed` | perspective_skew_sweep: perspective_skew (positive=True, iou_min=0.66, area_max=2.4, partial=True, decoys=random, occlusion=True) |
| `orb_feature_rich` | `orb_feature_rich` | `orbtex` | orb_feature_rich: orb_feature_rich (positive=True, iou_min=0.64, area_max=2.0, partial=True, decoys=clustered, occlusion=True) |
| `template_control_exact` | `template_control` | `ui` | template_control_exact: template_control (positive=True, iou_min=0.95, area_max=1.15, partial=False, decoys=grid, occlusion=False) |
| `hybrid_gate_conflicts` | `hybrid_gate` | `mixed` | hybrid_gate_conflicts: hybrid_gate (positive=True, iou_min=0.72, area_max=1.9, partial=True, decoys=mixed, occlusion=True) |
| `multi_monitor_dpi_shift` | `multi_monitor_dpi` | `mixed` | multi_monitor_dpi_shift: multi_monitor_dpi (positive=True, iou_min=0.78, area_max=1.5, partial=False, decoys=mixed, occlusion=False) |

## Scenario Configuration Details

### `vector_ui_baseline`

- Kind: `vector_ui`
- Style: `vector`
- Target: source=`synthetic` size=`{'min': 48, 'max': 160}` rotation=`{'min': 0, 'max': 180}` assets=`none`
- Background: palette=`ui_light` clutter=`0.72` continuous_canvas=`True`
- Transforms: scale=`{'min': 0.9, 'max': 1.2}` rotate=`{'min': -5, 'max': 5}` perspective_enabled=`False` skew_x=`{'min': 0.0, 'max': 0.0}` skew_y=`{'min': 0.0, 'max': 0.0}`
- Photometric: brightness=`{'min': -0.08, 'max': 0.08}` contrast=`{'min': 0.9, 'max': 1.12}` gamma=`{'min': 0.95, 'max': 1.06}` blur=`{'min': 0.0, 'max': 0.6}` jpeg_quality=`{'min': 88, 'max': 100}` noise_types=`compression_blocks`
- Decoys: enabled=`True` count=`{'min': 14, 'max': 34}` similarity=`{'min': 0.86, 'max': 0.97}` placement=`grid`
- Occlusion: enabled=`False` coverage=`{'min': 0.0, 'max': 0.0}`
- Expected: positive=`True` iou_min=`0.92` area_ratio_max=`1.25` allow_partial=`False`

### `photo_clutter`

- Kind: `photographic`
- Style: `photo`
- Target: source=`asset` size=`{'min': 56, 'max': 176}` rotation=`{'min': -25, 'max': 25}` assets=`docs/bench/assets/photo/4256_clutter_crop_zoom.jpg, docs/bench/assets/photo/4256_clutter_cool_grain.jpg, docs/bench/assets/photo/4256_clutter_warm_soft.jpg, docs/bench/assets/photo/4256_clutter_lowlight.jpg, docs/bench/assets/photo/4256_clutter_highiso.jpg`
- Background: palette=`mixed` clutter=`0.88` continuous_canvas=`True`
- Transforms: scale=`{'min': 0.75, 'max': 1.35}` rotate=`{'min': -20, 'max': 20}` perspective_enabled=`False` skew_x=`{'min': -1.0, 'max': 1.0}` skew_y=`{'min': -1.0, 'max': 1.0}`
- Photometric: brightness=`{'min': -0.2, 'max': 0.2}` contrast=`{'min': 0.7, 'max': 1.4}` gamma=`{'min': 0.8, 'max': 1.3}` blur=`{'min': 0.0, 'max': 2.0}` jpeg_quality=`{'min': 45, 'max': 100}` noise_types=`gaussian, poisson`
- Decoys: enabled=`True` count=`{'min': 20, 'max': 56}` similarity=`{'min': 0.8, 'max': 0.95}` placement=`mixed`
- Occlusion: enabled=`True` coverage=`{'min': 0.0, 'max': 0.25}`
- Expected: positive=`True` iou_min=`0.75` area_ratio_max=`1.6` allow_partial=`True`

### `repetitive_grid_camouflage`

- Kind: `repetitive_grid`
- Style: `grid`
- Target: source=`synthetic` size=`{'min': 48, 'max': 160}` rotation=`{'min': -15, 'max': 15}` assets=`none`
- Background: palette=`grayscale` clutter=`0.96` continuous_canvas=`True`
- Transforms: scale=`{'min': 0.85, 'max': 1.2}` rotate=`{'min': -15, 'max': 15}` perspective_enabled=`False` skew_x=`{'min': -2.0, 'max': 2.0}` skew_y=`{'min': -1.0, 'max': 1.0}`
- Photometric: brightness=`{'min': -0.12, 'max': 0.12}` contrast=`{'min': 0.85, 'max': 1.2}` gamma=`{'min': 0.9, 'max': 1.15}` blur=`{'min': 0.0, 'max': 1.0}` jpeg_quality=`{'min': 60, 'max': 100}` noise_types=`banding`
- Decoys: enabled=`True` count=`{'min': 36, 'max': 120}` similarity=`{'min': 0.9, 'max': 0.99}` placement=`grid`
- Occlusion: enabled=`False` coverage=`{'min': 0.0, 'max': 0.0}`
- Expected: positive=`True` iou_min=`0.85` area_ratio_max=`1.35` allow_partial=`False`

### `noise_stress_random`

- Kind: `noise_stress`
- Style: `noise`
- Target: source=`synthetic` size=`{'min': 64, 'max': 176}` rotation=`{'min': -45, 'max': 45}` assets=`none`
- Background: palette=`grayscale` clutter=`0.93` continuous_canvas=`True`
- Transforms: scale=`{'min': 0.7, 'max': 1.5}` rotate=`{'min': -30, 'max': 30}` perspective_enabled=`False` skew_x=`{'min': -2.0, 'max': 2.0}` skew_y=`{'min': -2.0, 'max': 2.0}`
- Photometric: brightness=`{'min': -0.25, 'max': 0.25}` contrast=`{'min': 0.65, 'max': 1.45}` gamma=`{'min': 0.75, 'max': 1.35}` blur=`{'min': 0.0, 'max': 2.2}` jpeg_quality=`{'min': 40, 'max': 95}` noise_types=`compression_blocks, gaussian, salt_pepper`
- Decoys: enabled=`True` count=`{'min': 24, 'max': 92}` similarity=`{'min': 0.75, 'max': 0.94}` placement=`random`
- Occlusion: enabled=`True` coverage=`{'min': 0.0, 'max': 0.35}`
- Expected: positive=`True` iou_min=`0.7` area_ratio_max=`1.8` allow_partial=`True`

### `scale_rotate_sweep`

- Kind: `scale_rotate`
- Style: `mixed`
- Target: source=`mixed` size=`{'min': 48, 'max': 180}` rotation=`{'min': -180, 'max': 180}` assets=`none`
- Background: palette=`mixed` clutter=`0.82` continuous_canvas=`True`
- Transforms: scale=`{'min': 0.4, 'max': 2.2}` rotate=`{'min': -180, 'max': 180}` perspective_enabled=`False` skew_x=`{'min': -4.0, 'max': 4.0}` skew_y=`{'min': -2.0, 'max': 2.0}`
- Photometric: brightness=`{'min': -0.16, 'max': 0.16}` contrast=`{'min': 0.78, 'max': 1.28}` gamma=`{'min': 0.85, 'max': 1.2}` blur=`{'min': 0.0, 'max': 1.6}` jpeg_quality=`{'min': 55, 'max': 100}` noise_types=`gaussian`
- Decoys: enabled=`True` count=`{'min': 16, 'max': 44}` similarity=`{'min': 0.8, 'max': 0.96}` placement=`mixed`
- Occlusion: enabled=`True` coverage=`{'min': 0.0, 'max': 0.2}`
- Expected: positive=`True` iou_min=`0.68` area_ratio_max=`2.2` allow_partial=`True`

### `perspective_skew_sweep`

- Kind: `perspective_skew`
- Style: `mixed`
- Target: source=`mixed` size=`{'min': 56, 'max': 176}` rotation=`{'min': -45, 'max': 45}` assets=`none`
- Background: palette=`mixed` clutter=`0.86` continuous_canvas=`True`
- Transforms: scale=`{'min': 0.75, 'max': 1.35}` rotate=`{'min': -45, 'max': 45}` perspective_enabled=`True` skew_x=`{'min': -10.0, 'max': 10.0}` skew_y=`{'min': -4.0, 'max': 4.0}`
- Photometric: brightness=`{'min': -0.14, 'max': 0.14}` contrast=`{'min': 0.8, 'max': 1.3}` gamma=`{'min': 0.84, 'max': 1.24}` blur=`{'min': 0.0, 'max': 1.5}` jpeg_quality=`{'min': 55, 'max': 100}` noise_types=`compression_blocks`
- Decoys: enabled=`True` count=`{'min': 14, 'max': 42}` similarity=`{'min': 0.78, 'max': 0.95}` placement=`random`
- Occlusion: enabled=`True` coverage=`{'min': 0.0, 'max': 0.28}`
- Expected: positive=`True` iou_min=`0.66` area_ratio_max=`2.4` allow_partial=`True`

### `orb_feature_rich`

- Kind: `orb_feature_rich`
- Style: `orbtex`
- Target: source=`synthetic` size=`{'min': 56, 'max': 176}` rotation=`{'min': -180, 'max': 180}` assets=`none`
- Background: palette=`grayscale` clutter=`0.9` continuous_canvas=`True`
- Transforms: scale=`{'min': 0.65, 'max': 1.7}` rotate=`{'min': -180, 'max': 180}` perspective_enabled=`True` skew_x=`{'min': -8.0, 'max': 8.0}` skew_y=`{'min': -4.0, 'max': 4.0}`
- Photometric: brightness=`{'min': -0.15, 'max': 0.15}` contrast=`{'min': 0.7, 'max': 1.5}` gamma=`{'min': 0.8, 'max': 1.25}` blur=`{'min': 0.0, 'max': 1.3}` jpeg_quality=`{'min': 45, 'max': 100}` noise_types=`banding, gaussian`
- Decoys: enabled=`True` count=`{'min': 22, 'max': 84}` similarity=`{'min': 0.72, 'max': 0.93}` placement=`clustered`
- Occlusion: enabled=`True` coverage=`{'min': 0.0, 'max': 0.18}`
- Expected: positive=`True` iou_min=`0.64` area_ratio_max=`2.0` allow_partial=`True`

### `template_control_exact`

- Kind: `template_control`
- Style: `ui`
- Target: source=`synthetic` size=`{'min': 48, 'max': 144}` rotation=`{'min': 0, 'max': 0}` assets=`none`
- Background: palette=`ui_dark` clutter=`0.55` continuous_canvas=`True`
- Transforms: scale=`{'min': 1.0, 'max': 1.0}` rotate=`{'min': 0.0, 'max': 0.0}` perspective_enabled=`False` skew_x=`{'min': 0.0, 'max': 0.0}` skew_y=`{'min': 0.0, 'max': 0.0}`
- Photometric: brightness=`{'min': -0.04, 'max': 0.04}` contrast=`{'min': 0.95, 'max': 1.08}` gamma=`{'min': 0.98, 'max': 1.04}` blur=`{'min': 0.0, 'max': 0.3}` jpeg_quality=`{'min': 95, 'max': 100}` noise_types=`none`
- Decoys: enabled=`True` count=`{'min': 8, 'max': 24}` similarity=`{'min': 0.85, 'max': 0.95}` placement=`grid`
- Occlusion: enabled=`False` coverage=`{'min': 0.0, 'max': 0.0}`
- Expected: positive=`True` iou_min=`0.95` area_ratio_max=`1.15` allow_partial=`False`

### `hybrid_gate_conflicts`

- Kind: `hybrid_gate`
- Style: `mixed`
- Target: source=`mixed` size=`{'min': 56, 'max': 176}` rotation=`{'min': -60, 'max': 60}` assets=`none`
- Background: palette=`mixed` clutter=`0.92` continuous_canvas=`True`
- Transforms: scale=`{'min': 0.7, 'max': 1.6}` rotate=`{'min': -80, 'max': 80}` perspective_enabled=`True` skew_x=`{'min': -7.0, 'max': 7.0}` skew_y=`{'min': -3.0, 'max': 3.0}`
- Photometric: brightness=`{'min': -0.2, 'max': 0.2}` contrast=`{'min': 0.75, 'max': 1.35}` gamma=`{'min': 0.82, 'max': 1.22}` blur=`{'min': 0.0, 'max': 1.8}` jpeg_quality=`{'min': 50, 'max': 100}` noise_types=`compression_blocks, gaussian`
- Decoys: enabled=`True` count=`{'min': 24, 'max': 96}` similarity=`{'min': 0.82, 'max': 0.99}` placement=`mixed`
- Occlusion: enabled=`True` coverage=`{'min': 0.0, 'max': 0.3}`
- Expected: positive=`True` iou_min=`0.72` area_ratio_max=`1.9` allow_partial=`True`
- Hybrid Policy: `{'must_consider_all_engines': True, 'select_by': 'score_then_iou', 'fallback_order': ['orb', 'akaze', 'brisk', 'kaze', 'sift', 'template']}`

### `multi_monitor_dpi_shift`

- Kind: `multi_monitor_dpi`
- Style: `mixed`
- Target: source=`mixed` size=`{'min': 56, 'max': 176}` rotation=`{'min': -30, 'max': 30}` assets=`none`
- Background: palette=`mixed` clutter=`0.8` continuous_canvas=`True`
- Transforms: scale=`{'min': 0.8, 'max': 1.3}` rotate=`{'min': -25, 'max': 25}` perspective_enabled=`False` skew_x=`{'min': -2.0, 'max': 2.0}` skew_y=`{'min': -1.0, 'max': 1.0}`
- Photometric: brightness=`{'min': -0.1, 'max': 0.1}` contrast=`{'min': 0.82, 'max': 1.18}` gamma=`{'min': 0.9, 'max': 1.18}` blur=`{'min': 0.0, 'max': 1.0}` jpeg_quality=`{'min': 70, 'max': 100}` noise_types=`compression_blocks`
- Decoys: enabled=`True` count=`{'min': 12, 'max': 36}` similarity=`{'min': 0.82, 'max': 0.95}` placement=`mixed`
- Occlusion: enabled=`False` coverage=`{'min': 0.0, 'max': 0.0}`
- Expected: positive=`True` iou_min=`0.78` area_ratio_max=`1.5` allow_partial=`False`
- Monitor Selector: `{'mode': 'round_robin', 'monitor_ids': ['mon-0', 'mon-1']}`

## Raw Artifacts

| Artifact | Purpose | Link |
| --- | --- | --- |
| Benchmark Overview | Latest benchmark section root. | [Open]({{ '/bench' | relative_url }}) |
| Detailed E2E Report | Compare strategy against measured results. | [Open]({{ '/bench/reports/find-on-screen-e2e' | relative_url }}) |
| Strategy JSON | Machine-readable strategy summary. | [Open]({{ '/bench/reports/find-on-screen-scenario-strategy.json' | relative_url }}) |
| Visual Gallery | Generated screenshots and summary boards. | [Open]({{ '/bench/reports/visuals' | relative_url }}) |
