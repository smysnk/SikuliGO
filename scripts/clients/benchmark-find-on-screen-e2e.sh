#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${THIS_DIR}/paths.sh"

BENCH_NAME="${FIND_BENCH_NAME:-BenchmarkFindOnScreenE2E}"
BENCH_TARGET="all"
BENCH_FILTER="${BENCH_NAME}"
BENCH_TIME="${FIND_BENCH_TIME:-200ms}"
BENCH_COUNT="${FIND_BENCH_COUNT:-1}"
BENCH_TAGS="${FIND_BENCH_TAGS:-opencv gocv_specific_modules gocv_features2d gocv_calib3d}"
RUN_MODE="full_report"
HIGH_RES="1"
ULTRA_RES="1"
BENCH_MANIFEST="${FIND_BENCH_SCENARIO_MANIFEST:-docs/bench/find-on-screen-scenarios.example.json}"
if [[ "${BENCH_MANIFEST}" == "example" || "${BENCH_MANIFEST}" == "default" ]]; then
  BENCH_MANIFEST="docs/bench/find-on-screen-scenarios.example.json"
fi
BENCH_SCHEMA="${FIND_BENCH_SCENARIO_SCHEMA:-}"
BENCH_REGION_SPEC="${FIND_BENCH_REGION_SPEC:-packages/api/internal/grpcv1/testdata/find-bench-assets/regions.json}"
BENCH_SEED="${FIND_BENCH_SEED:-}"
BENCH_PHOTO_ASSET="${FIND_BENCH_PHOTO_ASSET:-}"
DETERMINISTIC_REPORT="${FIND_BENCH_DETERMINISTIC_REPORT:-1}"
OUTPUT_BY_SEED="${FIND_BENCH_OUTPUT_BY_SEED:-0}"
GO_TEST_ARGS="${FIND_BENCH_GO_TEST_ARGS:-}"
GO_TEST_TIMEOUT="${FIND_BENCH_TEST_TIMEOUT:-90m}"
REPORT_DIR="${FIND_BENCH_REPORT_DIR:-${ROOT_DIR}/.test-results/bench}"
DOCS_REPORT_DIR="${FIND_BENCH_DOCS_REPORT_DIR:-${ROOT_DIR}/docs/bench/reports}"
DOCS_PUBLISH="${FIND_BENCH_PUBLISH_DOCS:-1}"
BENCH_BASENAME="find-on-screen-e2e"
if [[ -n "${BENCH_SEED}" && "${OUTPUT_BY_SEED}" =~ ^(1|true|yes|on)$ ]]; then
  BENCH_BASENAME="${BENCH_BASENAME}-seed${BENCH_SEED}"
fi
TEXT_OUT="${FIND_BENCH_TEXT_OUT:-${REPORT_DIR}/${BENCH_BASENAME}.txt}"
JSON_OUT="${FIND_BENCH_JSON_OUT:-${REPORT_DIR}/${BENCH_BASENAME}.json}"
MD_OUT="${FIND_BENCH_MD_OUT:-${REPORT_DIR}/${BENCH_BASENAME}.md}"
PERF_SVG_OUT="${FIND_BENCH_PERF_SVG_OUT:-${REPORT_DIR}/find-on-screen-performance.svg}"
ACCURACY_SVG_OUT="${FIND_BENCH_ACCURACY_SVG_OUT:-${REPORT_DIR}/find-on-screen-accuracy.svg}"
KIND_TIME_SVG_OUT="${FIND_BENCH_KIND_TIME_SVG_OUT:-${REPORT_DIR}/find-on-screen-kind-time.svg}"
KIND_SUCCESS_SVG_OUT="${FIND_BENCH_KIND_SUCCESS_SVG_OUT:-${REPORT_DIR}/find-on-screen-kind-success.svg}"
RESOLUTION_TIME_SVG_OUT="${FIND_BENCH_RESOLUTION_TIME_SVG_OUT:-${REPORT_DIR}/find-on-screen-resolution-time.svg}"
RESOLUTION_MATCHES_SVG_OUT="${FIND_BENCH_RESOLUTION_MATCHES_SVG_OUT:-${REPORT_DIR}/find-on-screen-resolution-matches.svg}"
RESOLUTION_MISSES_SVG_OUT="${FIND_BENCH_RESOLUTION_MISSES_SVG_OUT:-${REPORT_DIR}/find-on-screen-resolution-misses.svg}"
RESOLUTION_FALSE_POS_SVG_OUT="${FIND_BENCH_RESOLUTION_FALSE_POS_SVG_OUT:-${REPORT_DIR}/find-on-screen-resolution-false-positives.svg}"
VISUAL_ENABLE="1"
VISUAL_DIR="${FIND_BENCH_VISUAL_DIR:-${REPORT_DIR}/visuals}"
VISUAL_MAX_ATTEMPTS="${FIND_BENCH_VISUAL_MAX_ATTEMPTS:-2}"
VISUAL_TIMEOUT="${FIND_BENCH_VISUAL_TIMEOUT:-5s}"
PATCH_READMES="1"
README_PATHS="${FIND_BENCH_README_PATHS:-${ROOT_DIR}/README.md,${ROOT_DIR}/packages/client-node/README.md,${ROOT_DIR}/packages/client-python/README.md}"
README_SECTION_TITLE="${FIND_BENCH_README_SECTION_TITLE:-FindOnScreen Benchmark Test Results}"
README_INLINE_IMAGES="${FIND_BENCH_README_INLINE_IMAGES:-6}"
README_LINK_MODE="${FIND_BENCH_README_LINK_MODE:-pages}"
README_BASE_URL="${FIND_BENCH_README_BASE_URL:-https://smysnk.github.io/SikuliGO}"
STRATEGY_BASENAME="find-on-screen-scenario-strategy"
if [[ -n "${BENCH_SEED}" && "${OUTPUT_BY_SEED}" =~ ^(1|true|yes|on)$ ]]; then
  STRATEGY_BASENAME="${STRATEGY_BASENAME}-seed${BENCH_SEED}"
fi
STRATEGY_JSON_OUT="${FIND_BENCH_STRATEGY_JSON_OUT:-${REPORT_DIR}/${STRATEGY_BASENAME}.json}"
STRATEGY_MD_OUT="${FIND_BENCH_STRATEGY_MD_OUT:-${REPORT_DIR}/${STRATEGY_BASENAME}.md}"

mkdir -p "${REPORT_DIR}"

echo "[find-bench] package=${API_DIR}/internal/grpcv1"
echo "[find-bench] bench=${BENCH_NAME} benchtime=${BENCH_TIME} count=${BENCH_COUNT}"
echo "[find-bench] go_test_timeout=${GO_TEST_TIMEOUT}"
echo "[find-bench] target=${BENCH_TARGET} filter=${BENCH_FILTER}"
if [[ -n "${BENCH_TAGS}" ]]; then
  echo "[find-bench] tags=${BENCH_TAGS}"
fi
echo "[find-bench] run_mode=${RUN_MODE} high_res=${HIGH_RES} ultra_res=${ULTRA_RES}"
if [[ -n "${BENCH_MANIFEST}" ]]; then
  echo "[find-bench] manifest=${BENCH_MANIFEST}"
fi
if [[ -n "${BENCH_SCHEMA}" ]]; then
  echo "[find-bench] schema=${BENCH_SCHEMA}"
fi
if [[ -n "${BENCH_REGION_SPEC}" ]]; then
  echo "[find-bench] region_spec=${BENCH_REGION_SPEC}"
fi
if [[ -n "${BENCH_SEED}" ]]; then
  echo "[find-bench] seed=${BENCH_SEED}"
fi
if [[ -n "${BENCH_PHOTO_ASSET}" ]]; then
  echo "[find-bench] photo_asset=${BENCH_PHOTO_ASSET}"
fi
echo "[find-bench] deterministic_report=${DETERMINISTIC_REPORT}"

echo "[find-bench] text=${TEXT_OUT}"
echo "[find-bench] json=${JSON_OUT}"
echo "[find-bench] markdown=${MD_OUT}"
echo "[find-bench] perf_svg=${PERF_SVG_OUT}"
echo "[find-bench] accuracy_svg=${ACCURACY_SVG_OUT}"
echo "[find-bench] kind_time_svg=${KIND_TIME_SVG_OUT}"
echo "[find-bench] kind_success_svg=${KIND_SUCCESS_SVG_OUT}"
echo "[find-bench] resolution_time_svg=${RESOLUTION_TIME_SVG_OUT}"
echo "[find-bench] resolution_matches_svg=${RESOLUTION_MATCHES_SVG_OUT}"
echo "[find-bench] resolution_misses_svg=${RESOLUTION_MISSES_SVG_OUT}"
echo "[find-bench] resolution_false_pos_svg=${RESOLUTION_FALSE_POS_SVG_OUT}"
echo "[find-bench] visuals=${VISUAL_ENABLE} dir=${VISUAL_DIR} max_attempts=${VISUAL_MAX_ATTEMPTS} timeout=${VISUAL_TIMEOUT}"
echo "[find-bench] patch_readmes=${PATCH_READMES} readmes=${README_PATHS} inline_images=${README_INLINE_IMAGES}"
echo "[find-bench] readme_link_mode=${README_LINK_MODE} readme_base_url=${README_BASE_URL}"
echo "[find-bench] strategy_json=${STRATEGY_JSON_OUT}"
echo "[find-bench] strategy_md=${STRATEGY_MD_OUT}"
echo "[find-bench] publish_docs=${DOCS_PUBLISH} docs_dir=${DOCS_REPORT_DIR}"

cd "${API_DIR}"

FIND_BENCH_SCENARIO_MANIFEST="${BENCH_MANIFEST}" \
FIND_BENCH_REPORT_DIR="${REPORT_DIR}" \
FIND_BENCH_STRATEGY_JSON_OUT="${STRATEGY_JSON_OUT}" \
FIND_BENCH_STRATEGY_MD_OUT="${STRATEGY_MD_OUT}" \
"${THIS_DIR}/report-find-on-screen-scenario-strategy.sh"

FIND_BENCH_SCENARIO_MANIFEST="${BENCH_MANIFEST}" \
FIND_BENCH_SCENARIO_SCHEMA="${BENCH_SCHEMA}" \
FIND_BENCH_REGION_SPEC="${BENCH_REGION_SPEC}" \
FIND_BENCH_SEED="${BENCH_SEED}" \
FIND_BENCH_PHOTO_ASSET="${BENCH_PHOTO_ASSET}" \
go test ./internal/grpcv1 -run '^TestFindBenchManifestPreflightFromEnv$' -count=1 -v

cmd=(
  go test
  ./internal/grpcv1
  -run '^$'
  -bench "${BENCH_FILTER}"
  -benchmem
  -benchtime "${BENCH_TIME}"
  -count "${BENCH_COUNT}"
)

if [[ -n "${GO_TEST_TIMEOUT}" ]]; then
  cmd+=( -timeout "${GO_TEST_TIMEOUT}" )
fi

if [[ -n "${BENCH_TAGS}" ]]; then
  cmd+=( -tags "${BENCH_TAGS}" )
fi

if [[ -n "${GO_TEST_ARGS}" ]]; then
  # shellcheck disable=SC2206
  extra=( ${GO_TEST_ARGS} )
  cmd+=( "${extra[@]}" )
fi

FIND_BENCH_VISUAL="${VISUAL_ENABLE}" \
FIND_BENCH_VISUAL_DIR="${VISUAL_DIR}" \
FIND_BENCH_VISUAL_MAX_ATTEMPTS="${VISUAL_MAX_ATTEMPTS}" \
FIND_BENCH_VISUAL_TIMEOUT="${VISUAL_TIMEOUT}" \
FIND_BENCH_HIGH_RES="${HIGH_RES}" \
FIND_BENCH_ULTRA_RES="${ULTRA_RES}" \
FIND_BENCH_SCENARIO_MANIFEST="${BENCH_MANIFEST}" \
FIND_BENCH_SCENARIO_SCHEMA="${BENCH_SCHEMA}" \
FIND_BENCH_REGION_SPEC="${BENCH_REGION_SPEC}" \
FIND_BENCH_SEED="${BENCH_SEED}" \
FIND_BENCH_PHOTO_ASSET="${BENCH_PHOTO_ASSET}" \
"${cmd[@]}" | tee "${TEXT_OUT}"

BENCH_NAME="${BENCH_NAME}" \
BENCH_TARGET="${BENCH_TARGET}" \
BENCH_FILTER="${BENCH_FILTER}" \
BENCH_TIME="${BENCH_TIME}" \
BENCH_COUNT="${BENCH_COUNT}" \
BENCH_TAGS="${BENCH_TAGS}" \
RUN_MODE="${RUN_MODE}" \
HIGH_RES="${HIGH_RES}" \
ULTRA_RES="${ULTRA_RES}" \
BENCH_SEED="${BENCH_SEED}" \
DETERMINISTIC_REPORT="${DETERMINISTIC_REPORT}" \
BENCH_MANIFEST="${BENCH_MANIFEST}" \
BENCH_SCHEMA="${BENCH_SCHEMA}" \
BENCH_REGION_SPEC="${BENCH_REGION_SPEC}" \
REPORT_DIR="${REPORT_DIR}" \
DOCS_REPORT_DIR="${DOCS_REPORT_DIR}" \
DOCS_PUBLISH="${DOCS_PUBLISH}" \
PROJECT_ROOT="${ROOT_DIR}" \
VISUAL_ENABLE="${VISUAL_ENABLE}" \
VISUAL_DIR="${VISUAL_DIR}" \
VISUAL_MAX_ATTEMPTS="${VISUAL_MAX_ATTEMPTS}" \
VISUAL_TIMEOUT="${VISUAL_TIMEOUT}" \
PATCH_READMES="${PATCH_READMES}" \
README_PATHS="${README_PATHS}" \
README_SECTION_TITLE="${README_SECTION_TITLE}" \
README_INLINE_IMAGES="${README_INLINE_IMAGES}" \
README_LINK_MODE="${README_LINK_MODE}" \
README_BASE_URL="${README_BASE_URL}" \
TEXT_OUT="${TEXT_OUT}" \
JSON_OUT="${JSON_OUT}" \
MD_OUT="${MD_OUT}" \
PERF_SVG_OUT="${PERF_SVG_OUT}" \
ACCURACY_SVG_OUT="${ACCURACY_SVG_OUT}" \
KIND_TIME_SVG_OUT="${KIND_TIME_SVG_OUT}" \
KIND_SUCCESS_SVG_OUT="${KIND_SUCCESS_SVG_OUT}" \
RESOLUTION_TIME_SVG_OUT="${RESOLUTION_TIME_SVG_OUT}" \
RESOLUTION_MATCHES_SVG_OUT="${RESOLUTION_MATCHES_SVG_OUT}" \
RESOLUTION_MISSES_SVG_OUT="${RESOLUTION_MISSES_SVG_OUT}" \
RESOLUTION_FALSE_POS_SVG_OUT="${RESOLUTION_FALSE_POS_SVG_OUT}" \
python3 - <<'PY'
from __future__ import annotations

import json
import os
import re
import shutil
from collections import defaultdict
from datetime import datetime, timezone
from pathlib import Path
from statistics import median

text_path = Path(os.environ["TEXT_OUT"])
json_path = Path(os.environ["JSON_OUT"])
md_path = Path(os.environ["MD_OUT"])
perf_svg_path = Path(os.environ["PERF_SVG_OUT"])
accuracy_svg_path = Path(os.environ["ACCURACY_SVG_OUT"])
kind_time_svg_path = Path(os.environ["KIND_TIME_SVG_OUT"])
kind_success_svg_path = Path(os.environ["KIND_SUCCESS_SVG_OUT"])
resolution_time_svg_path = Path(os.environ["RESOLUTION_TIME_SVG_OUT"])
resolution_matches_svg_path = Path(os.environ["RESOLUTION_MATCHES_SVG_OUT"])
resolution_misses_svg_path = Path(os.environ["RESOLUTION_MISSES_SVG_OUT"])
resolution_false_pos_svg_path = Path(os.environ["RESOLUTION_FALSE_POS_SVG_OUT"])

bench_name = os.environ["BENCH_NAME"]
bench_target = os.environ.get("BENCH_TARGET", "all")
bench_filter = os.environ.get("BENCH_FILTER", bench_name)
bench_time = os.environ["BENCH_TIME"]
bench_count = os.environ["BENCH_COUNT"]
bench_tags = os.environ.get("BENCH_TAGS", "")
high_res = os.environ.get("HIGH_RES", "")
ultra_res = os.environ.get("ULTRA_RES", "")
run_mode = os.environ.get("RUN_MODE", "full_range")
bench_seed = os.environ.get("BENCH_SEED", "").strip()
deterministic_report = os.environ.get("DETERMINISTIC_REPORT", "").strip().lower() in {"1", "true", "yes", "on"}
bench_manifest = os.environ.get("BENCH_MANIFEST", "").strip()
bench_schema = os.environ.get("BENCH_SCHEMA", "").strip()
bench_region_spec = os.environ.get("BENCH_REGION_SPEC", "").strip()
visual_enable = os.environ.get("VISUAL_ENABLE", "")
visual_dir = os.environ.get("VISUAL_DIR", "")
visual_max_attempts = os.environ.get("VISUAL_MAX_ATTEMPTS", "")
visual_timeout = os.environ.get("VISUAL_TIMEOUT", "")
patch_readmes = os.environ.get("PATCH_READMES", "")
readme_paths = os.environ.get("README_PATHS", "")
readme_section_title = os.environ.get("README_SECTION_TITLE", "FindOnScreen Benchmark Test Results")
readme_inline_images = int(os.environ.get("README_INLINE_IMAGES", "6"))
readme_link_mode = os.environ.get("README_LINK_MODE", "pages").strip().lower()
readme_base_url = os.environ.get("README_BASE_URL", "https://smysnk.github.io/SikuliGO").strip().rstrip("/")
root_dir = Path(os.environ.get("PROJECT_ROOT", "")).resolve()
report_dir = Path(os.environ.get("REPORT_DIR", "")).resolve()
docs_report_dir_raw = os.environ.get("DOCS_REPORT_DIR", "").strip()
docs_publish = os.environ.get("DOCS_PUBLISH", "").strip().lower() in {"1", "true", "yes", "on"}
docs_root = (root_dir / "docs").resolve()

content = text_path.read_text(encoding="utf-8", errors="replace").splitlines()

meta = {
    "goos": "",
    "goarch": "",
    "package": "",
    "cpu": "",
    "timestamp_utc": f"seed:{bench_seed}" if deterministic_report and bench_seed else datetime.now(timezone.utc).isoformat(),
    "bench_name": bench_name,
    "bench_target": bench_target,
    "bench_filter": bench_filter,
    "benchtime": bench_time,
    "count": int(bench_count),
    "tags": bench_tags,
    "seed": bench_seed,
    "manifest": bench_manifest,
    "schema": bench_schema,
    "region_spec": bench_region_spec,
}

label_re = re.compile(r"^Benchmark[^/]+/engine=([^/]+)/(.+)-\d+$")

def as_float(value: str) -> float | None:
    try:
        return float(value)
    except ValueError:
        return None

def write_svg(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content, encoding="utf-8")

def svg_escape(s: str) -> str:
    return (
        s.replace("&", "&amp;")
        .replace("<", "&lt;")
        .replace(">", "&gt;")
        .replace('"', "&quot;")
        .replace("'", "&#39;")
    )

def render_single_series_bar_chart_svg(
    title: str,
    y_label: str,
    labels: list[str],
    values: list[float],
    fill: str,
) -> str:
    width, height = 980, 520
    left, right, top, bottom = 90, 40, 65, 90
    plot_w = width - left - right
    plot_h = height - top - bottom
    max_v = max(values) if values else 1.0
    ymax = max(1.0, max_v * 1.15)
    n = max(1, len(values))
    gap = 20
    bar_w = max(18, int((plot_w - gap * (n + 1)) / n))

    parts = [
        f'<svg xmlns="http://www.w3.org/2000/svg" width="{width}" height="{height}" viewBox="0 0 {width} {height}">',
        '<style>text{font-family:-apple-system,BlinkMacSystemFont,Segoe UI,Roboto,Arial,sans-serif;fill:#0f172a} .muted{fill:#475569}</style>',
        '<rect x="0" y="0" width="100%" height="100%" fill="#f8fafc"/>',
        f'<text x="{width/2:.0f}" y="34" font-size="22" text-anchor="middle">{svg_escape(title)}</text>',
        f'<text x="{left - 55}" y="{top + plot_h/2:.0f}" font-size="14" text-anchor="middle" transform="rotate(-90 {left - 55},{top + plot_h/2:.0f})" class="muted">{svg_escape(y_label)}</text>',
        f'<line x1="{left}" y1="{top}" x2="{left}" y2="{top + plot_h}" stroke="#334155" stroke-width="2"/>',
        f'<line x1="{left}" y1="{top + plot_h}" x2="{left + plot_w}" y2="{top + plot_h}" stroke="#334155" stroke-width="2"/>',
    ]

    ticks = 5
    for i in range(ticks + 1):
        t = i / ticks
        y = top + plot_h - (plot_h * t)
        v = ymax * t
        parts.append(f'<line x1="{left}" y1="{y:.1f}" x2="{left + plot_w}" y2="{y:.1f}" stroke="#e2e8f0" stroke-width="1"/>')
        parts.append(f'<text x="{left - 10}" y="{y + 5:.1f}" font-size="12" text-anchor="end" class="muted">{v:.1f}</text>')

    for i, (label, value) in enumerate(zip(labels, values)):
        x = left + gap + i * (bar_w + gap)
        h = 0 if ymax <= 0 else (value / ymax) * plot_h
        y = top + plot_h - h
        parts.append(f'<rect x="{x:.1f}" y="{y:.1f}" width="{bar_w}" height="{h:.1f}" fill="{fill}" rx="4"/>')
        parts.append(f'<text x="{x + bar_w/2:.1f}" y="{y - 6:.1f}" font-size="12" text-anchor="middle">{value:.2f}</text>')
        parts.append(f'<text x="{x + bar_w/2:.1f}" y="{top + plot_h + 20:.1f}" font-size="11" text-anchor="middle" class="muted">{svg_escape(label)}</text>')

    parts.append("</svg>")
    return "\n".join(parts)

def render_grouped_accuracy_chart_svg(
    title: str,
    labels: list[str],
    success_vals: list[float],
    false_pos_vals: list[float],
) -> str:
    width, height = 980, 560
    left, right, top, bottom = 90, 40, 65, 110
    plot_w = width - left - right
    plot_h = height - top - bottom
    n = max(1, len(labels))
    group_gap = 24
    bar_w = max(14, int((plot_w - group_gap * (n + 1)) / (n * 2)))

    parts = [
        f'<svg xmlns="http://www.w3.org/2000/svg" width="{width}" height="{height}" viewBox="0 0 {width} {height}">',
        '<style>text{font-family:-apple-system,BlinkMacSystemFont,Segoe UI,Roboto,Arial,sans-serif;fill:#0f172a} .muted{fill:#475569}</style>',
        '<rect x="0" y="0" width="100%" height="100%" fill="#f8fafc"/>',
        f'<text x="{width/2:.0f}" y="34" font-size="22" text-anchor="middle">{svg_escape(title)}</text>',
        f'<line x1="{left}" y1="{top}" x2="{left}" y2="{top + plot_h}" stroke="#334155" stroke-width="2"/>',
        f'<line x1="{left}" y1="{top + plot_h}" x2="{left + plot_w}" y2="{top + plot_h}" stroke="#334155" stroke-width="2"/>',
    ]

    ticks = 5
    for i in range(ticks + 1):
        t = i / ticks
        y = top + plot_h - (plot_h * t)
        v = 100 * t
        parts.append(f'<line x1="{left}" y1="{y:.1f}" x2="{left + plot_w}" y2="{y:.1f}" stroke="#e2e8f0" stroke-width="1"/>')
        parts.append(f'<text x="{left - 10}" y="{y + 5:.1f}" font-size="12" text-anchor="end" class="muted">{v:.0f}</text>')

    for i, label in enumerate(labels):
        x0 = left + group_gap + i * (2 * bar_w + group_gap)
        s = success_vals[i]
        f = false_pos_vals[i]
        sh = (s / 100.0) * plot_h
        fh = (f / 100.0) * plot_h
        sy = top + plot_h - sh
        fy = top + plot_h - fh
        parts.append(f'<rect x="{x0:.1f}" y="{sy:.1f}" width="{bar_w}" height="{sh:.1f}" fill="#16a34a" rx="3"/>')
        parts.append(f'<rect x="{x0 + bar_w + 6:.1f}" y="{fy:.1f}" width="{bar_w}" height="{fh:.1f}" fill="#dc2626" rx="3"/>')
        parts.append(f'<text x="{x0 + bar_w/2:.1f}" y="{sy - 6:.1f}" font-size="11" text-anchor="middle">{s:.1f}</text>')
        parts.append(f'<text x="{x0 + bar_w + 6 + bar_w/2:.1f}" y="{fy - 6:.1f}" font-size="11" text-anchor="middle">{f:.1f}</text>')
        parts.append(f'<text x="{x0 + bar_w + 3:.1f}" y="{top + plot_h + 20:.1f}" font-size="11" text-anchor="middle" class="muted">{svg_escape(label)}</text>')

    legend_y = height - 32
    parts.append(f'<rect x="{left}" y="{legend_y}" width="14" height="14" fill="#16a34a" rx="2"/>')
    parts.append(f'<text x="{left + 22}" y="{legend_y + 12}" font-size="12" class="muted">Success %</text>')
    parts.append(f'<rect x="{left + 130}" y="{legend_y}" width="14" height="14" fill="#dc2626" rx="2"/>')
    parts.append(f'<text x="{left + 152}" y="{legend_y + 12}" font-size="12" class="muted">False Positive %</text>')

    parts.append("</svg>")
    return "\n".join(parts)

def ordered_engines(raw_engines: list[str]) -> list[str]:
    preferred = ["orb", "akaze", "brisk", "kaze", "sift", "template", "hybrid"]
    known = [e for e in preferred if e in raw_engines]
    rest = sorted([e for e in raw_engines if e not in preferred])
    return known + rest

def scenario_resolution(scenario: str) -> str:
    m = re.search(r"(\d+x\d+)$", scenario)
    if m:
        return m.group(1)
    m2 = re.search(r"(\d+x\d+)_", scenario)
    if m2:
        return m2.group(1)
    return "unknown"

def scenario_kind(scenario: str) -> str:
    s = scenario.lower()
    if s.startswith("vector_") or "vector_ui" in s:
        return "vector_ui"
    if s.startswith("photo_") or "photo_clutter" in s:
        return "photographic"
    if s.startswith("ui_") or "template_control" in s:
        return "template_control"
    if s.startswith("grid_") or s.startswith("glyph_") or "repetitive_grid" in s:
        return "repetitive_grid"
    if s.startswith("noise_") or "noise_stress" in s:
        return "noise_stress"
    if "scale_rotate" in s or "mix_resize" in s:
        return "scale_rotate"
    if "perspective" in s or "skew" in s:
        return "perspective_skew"
    if "orb_feature_rich" in s or "orbtex" in s:
        return "orb_feature_rich"
    if "hybrid_gate" in s:
        return "hybrid_gate"
    if "multi_monitor" in s:
        return "multi_monitor_dpi"
    if "mix_rotate" in s:
        return "scale_rotate"
    return "unknown"

def resolution_sort_key(res: str) -> tuple[int, int, str]:
    m = re.match(r"^(\d+)x(\d+)$", res)
    if not m:
        return (10**9, 10**9, res)
    w = int(m.group(1))
    h = int(m.group(2))
    return (w * h, w, res)

def engine_color(name: str, idx: int) -> str:
    palette = {
        "orb": "#ea580c",
        "akaze": "#a21caf",
        "brisk": "#db2777",
        "kaze": "#7c3aed",
        "sift": "#9333ea",
        "template": "#2563eb",
        "hybrid": "#16a34a",
    }
    fallback = ["#0f766e", "#c2410c", "#4f46e5", "#0ea5e9", "#64748b"]
    if name in palette:
        return palette[name]
    return fallback[idx % len(fallback)]

def render_grouped_multi_series_bar_chart_svg(
    title: str,
    y_label: str,
    group_labels: list[str],
    series: list[tuple[str, list[float], str]],
) -> str:
    width, height = 1160, 580
    left, right, top, bottom = 90, 40, 70, 130
    plot_w = width - left - right
    plot_h = height - top - bottom

    n_groups = max(1, len(group_labels))
    n_series = max(1, len(series))
    group_gap = 24
    slot_w = (plot_w - group_gap * (n_groups + 1)) / n_groups
    inner_pad = 8
    bar_gap = 5
    usable = max(1.0, slot_w - inner_pad * 2 - bar_gap * (n_series - 1))
    bar_w = max(8.0, usable / n_series)

    max_v = 0.0
    for _, vals, _ in series:
        if vals:
            max_v = max(max_v, max(vals))
    ymax = max(1.0, max_v * 1.15)

    parts = [
        f'<svg xmlns="http://www.w3.org/2000/svg" width="{width}" height="{height}" viewBox="0 0 {width} {height}">',
        '<style>text{font-family:-apple-system,BlinkMacSystemFont,Segoe UI,Roboto,Arial,sans-serif;fill:#0f172a} .muted{fill:#475569}</style>',
        '<rect x="0" y="0" width="100%" height="100%" fill="#f8fafc"/>',
        f'<text x="{width/2:.0f}" y="36" font-size="22" text-anchor="middle">{svg_escape(title)}</text>',
        f'<text x="{left - 58}" y="{top + plot_h/2:.0f}" font-size="14" text-anchor="middle" transform="rotate(-90 {left - 58},{top + plot_h/2:.0f})" class="muted">{svg_escape(y_label)}</text>',
        f'<line x1="{left}" y1="{top}" x2="{left}" y2="{top + plot_h}" stroke="#334155" stroke-width="2"/>',
        f'<line x1="{left}" y1="{top + plot_h}" x2="{left + plot_w}" y2="{top + plot_h}" stroke="#334155" stroke-width="2"/>',
    ]

    ticks = 5
    for i in range(ticks + 1):
        t = i / ticks
        y = top + plot_h - (plot_h * t)
        v = ymax * t
        parts.append(f'<line x1="{left}" y1="{y:.1f}" x2="{left + plot_w}" y2="{y:.1f}" stroke="#e2e8f0" stroke-width="1"/>')
        parts.append(f'<text x="{left - 10}" y="{y + 5:.1f}" font-size="12" text-anchor="end" class="muted">{v:.1f}</text>')

    for gi, g_label in enumerate(group_labels):
        group_x = left + group_gap + gi * (slot_w + group_gap)
        gx_center = group_x + slot_w / 2.0
        parts.append(f'<text x="{gx_center:.1f}" y="{top + plot_h + 24:.1f}" font-size="11" text-anchor="middle" class="muted">{svg_escape(g_label)}</text>')
        for si, (_, vals, color) in enumerate(series):
            value = vals[gi] if gi < len(vals) else 0.0
            x = group_x + inner_pad + si * (bar_w + bar_gap)
            h = 0 if ymax <= 0 else (value / ymax) * plot_h
            y = top + plot_h - h
            parts.append(f'<rect x="{x:.1f}" y="{y:.1f}" width="{bar_w:.1f}" height="{h:.1f}" fill="{color}" rx="3"/>')
            if value > 0:
                parts.append(f'<text x="{x + bar_w/2:.1f}" y="{max(top + 12, y - 4):.1f}" font-size="10" text-anchor="middle">{value:.2f}</text>')

    legend_x = left
    legend_y = height - 44
    for si, (name, _, color) in enumerate(series):
        x = legend_x + si * 170
        parts.append(f'<rect x="{x}" y="{legend_y}" width="14" height="14" fill="{color}" rx="2"/>')
        parts.append(f'<text x="{x + 22}" y="{legend_y + 12}" font-size="12" class="muted">{svg_escape(name)}</text>')

    parts.append("</svg>")
    return "\n".join(parts)

def derive_status(metrics: dict[str, float]) -> str:
    success = metrics.get("success/op", 1.0)
    not_found = metrics.get("not_found/op", 0.0)
    unsupported = metrics.get("unsupported/op", 0.0)
    error = metrics.get("error/op", 0.0)
    overlap_miss = metrics.get("overlap_miss/op", 0.0)

    if success >= 0.999 and not_found == 0.0 and unsupported == 0.0 and error == 0.0 and overlap_miss == 0.0:
        return "ok"
    if unsupported > 0.0 and success == 0.0:
        return "unsupported"
    if error > 0.0 and success == 0.0:
        return "error"
    if overlap_miss > 0.0 and success == 0.0:
        return "overlap_miss"
    if not_found > 0.0 and success == 0.0:
        return "not_found"
    return "partial"

results: list[dict[str, object]] = []
for line in content:
    if line.startswith("goos:"):
        meta["goos"] = line.split(":", 1)[1].strip()
        continue
    if line.startswith("goarch:"):
        meta["goarch"] = line.split(":", 1)[1].strip()
        continue
    if line.startswith("pkg:"):
        meta["package"] = line.split(":", 1)[1].strip()
        continue
    if line.startswith("cpu:"):
        meta["cpu"] = line.split(":", 1)[1].strip()
        continue

    raw = line.strip()
    if not raw.startswith("Benchmark"):
        continue
    parts = raw.split()
    if len(parts) < 4:
        continue

    label = parts[0]
    iter_val = as_float(parts[1])
    if iter_val is None:
        continue
    iters = int(iter_val)

    metrics: dict[str, float] = {}
    idx = 2
    while idx + 1 < len(parts):
        value = as_float(parts[idx])
        metric = parts[idx + 1]
        if value is None:
            idx += 1
            continue
        metrics[metric] = value
        idx += 2

    ns_per_op = metrics.get("ns/op")
    if ns_per_op is None:
        continue
    bytes_per_op = metrics.get("B/op")
    allocs_per_op = metrics.get("allocs/op")

    engine = "unknown"
    scenario = label
    lm = label_re.match(label)
    if lm:
        engine = lm.group(1)
        scenario = lm.group(2)

    results.append(
        {
            "label": label,
            "engine": engine,
            "scenario": scenario,
            "iterations": iters,
            "ns_per_op": ns_per_op,
            "ms_per_op": ns_per_op / 1_000_000.0,
            "bytes_per_op": bytes_per_op,
            "allocs_per_op": allocs_per_op,
            "status": derive_status(metrics),
            "metrics": metrics,
        }
    )

if not results:
    raise SystemExit(f"No benchmark rows parsed from {text_path}")

engines: dict[str, list[dict[str, object]]] = defaultdict(list)
for r in results:
    engines[str(r["engine"])].append(r)
for rows in engines.values():
    rows.sort(key=lambda x: float(x["ns_per_op"]))

summary_rows: list[dict[str, object]] = []
for engine, rows in sorted(engines.items()):
    ms_values = [float(r["ms_per_op"]) for r in rows]
    best = min(rows, key=lambda r: float(r["ms_per_op"]))
    worst = max(rows, key=lambda r: float(r["ms_per_op"]))
    status_counts: dict[str, int] = defaultdict(int)
    query_cases = 0.0
    query_ok = 0.0
    query_not_found = 0.0
    query_unsupported = 0.0
    query_error = 0.0
    query_overlap_miss = 0.0
    for row in rows:
        status_counts[str(row["status"])] += 1
        metrics = row.get("metrics", {})
        q_total = float(metrics.get("query_total/op", 0.0))
        if q_total <= 0.0:
            q_total = 1.0
        query_cases += q_total
        if "query_ok/op" in metrics:
            query_ok += float(metrics.get("query_ok/op", 0.0))
            query_not_found += float(metrics.get("query_not_found/op", 0.0))
            query_unsupported += float(metrics.get("query_unsupported/op", 0.0))
            query_error += float(metrics.get("query_error/op", 0.0))
            query_overlap_miss += float(metrics.get("query_overlap_miss/op", 0.0))
            continue

        # Backward-compatible fallback for older benchmark outputs without query_* metrics.
        st = str(row.get("status", "error"))
        if st == "ok":
            query_ok += q_total
        elif st == "not_found":
            query_not_found += q_total
        elif st == "unsupported":
            query_unsupported += q_total
        elif st == "overlap_miss":
            query_overlap_miss += q_total
        elif st == "error":
            query_error += q_total

    query_partial = max(0.0, query_cases - query_ok - query_not_found - query_unsupported - query_error - query_overlap_miss)
    summary_rows.append(
        {
            "engine": engine,
            "cases": query_cases,
            "status_counts": dict(status_counts),
            "query_ok": query_ok,
            "query_partial": query_partial,
            "query_not_found": query_not_found,
            "query_unsupported": query_unsupported,
            "query_error": query_error,
            "query_overlap_miss": query_overlap_miss,
            "row_count": len(rows),
            "avg_ms_per_op": sum(ms_values) / len(ms_values),
            "median_ms_per_op": float(median(ms_values)),
            "best_scenario": str(best["scenario"]),
            "best_ms_per_op": float(best["ms_per_op"]),
            "worst_scenario": str(worst["scenario"]),
            "worst_ms_per_op": float(worst["ms_per_op"]),
        }
    )

metrics_chart_rows: list[dict[str, object]] = []
for engine, rows in sorted(engines.items()):
    success_vals: list[float] = []
    false_positive_vals: list[float] = []
    not_found_vals: list[float] = []
    unsupported_vals: list[float] = []
    error_vals: list[float] = []
    query_cases = 0.0
    query_rows = 0
    for r in rows:
        metrics = r.get("metrics", {})
        q_total = float(metrics.get("query_total/op", 0.0))
        if q_total > 0:
            query_rows += 1
            query_cases += q_total
            denom = max(1.0, q_total)
            success_vals.append((float(metrics.get("query_ok/op", 0.0)) / denom))
            false_positive_vals.append((float(metrics.get("query_overlap_miss/op", 0.0)) / denom))
            not_found_vals.append((float(metrics.get("query_not_found/op", 0.0)) / denom))
            unsupported_vals.append((float(metrics.get("query_unsupported/op", 0.0)) / denom))
            error_vals.append((float(metrics.get("query_error/op", 0.0)) / denom))
        else:
            success_vals.append(float(metrics.get("success/op", 1.0)))
            false_positive_vals.append(float(metrics.get("overlap_miss/op", 0.0)))
            not_found_vals.append(float(metrics.get("not_found/op", 0.0)))
            unsupported_vals.append(float(metrics.get("unsupported/op", 0.0)))
            error_vals.append(float(metrics.get("error/op", 0.0)))
            query_cases += 1.0
    ms_vals = [float(r["ms_per_op"]) for r in rows]
    metrics_chart_rows.append(
        {
            "engine": engine,
            "cases": query_cases,
            "rows": len(rows),
            "query_rows": query_rows,
            "avg_ms_per_op": sum(ms_vals) / len(ms_vals),
            "median_ms_per_op": float(median(ms_vals)),
            "success_rate_pct": (sum(success_vals) / len(success_vals)) * 100.0,
            "false_positive_rate_pct": (sum(false_positive_vals) / len(false_positive_vals)) * 100.0,
            "not_found_rate_pct": (sum(not_found_vals) / len(not_found_vals)) * 100.0,
            "unsupported_rate_pct": (sum(unsupported_vals) / len(unsupported_vals)) * 100.0,
            "error_rate_pct": (sum(error_vals) / len(error_vals)) * 100.0,
        }
    )

engine_names = ordered_engines(sorted(engines.keys()))
resolution_buckets: dict[str, dict[str, dict[str, float]]] = defaultdict(lambda: defaultdict(lambda: {
    "ms_sum": 0.0,
    "n": 0.0,
    "matches": 0.0,
    "misses": 0.0,
    "false_pos": 0.0,
}))
for row in results:
    scenario = str(row.get("scenario", ""))
    engine = str(row.get("engine", "unknown"))
    res_group = scenario_resolution(scenario)
    metrics = row.get("metrics", {})
    query_total = float(metrics.get("query_total/op", 0.0))
    if query_total > 0:
        success = float(metrics.get("query_ok/op", 0.0))
        miss = float(metrics.get("query_not_found/op", 0.0))
        false_pos = float(metrics.get("query_overlap_miss/op", 0.0))
    else:
        success = float(metrics.get("success/op", 1.0))
        miss = float(metrics.get("not_found/op", 0.0))
        false_pos = float(metrics.get("overlap_miss/op", 0.0))
    bucket = resolution_buckets[res_group][engine]
    bucket["ms_sum"] += float(row.get("ms_per_op", 0.0))
    bucket["n"] += 1.0
    bucket["matches"] += success
    bucket["misses"] += miss
    bucket["false_pos"] += false_pos

resolution_groups = sorted(resolution_buckets.keys(), key=resolution_sort_key)
resolution_rows: list[dict[str, object]] = []
for res in resolution_groups:
    engines_map: dict[str, dict[str, float]] = {}
    for eng in engine_names:
        b = resolution_buckets[res].get(eng)
        if b is None or b["n"] <= 0:
            engines_map[eng] = {
                "avg_ms_per_op": 0.0,
                "match_count": 0.0,
                "miss_count": 0.0,
                "false_positive_count": 0.0,
                "cases": 0.0,
            }
            continue
        n = b["n"]
        engines_map[eng] = {
            "avg_ms_per_op": b["ms_sum"] / n,
            "match_count": b["matches"],
            "miss_count": b["misses"],
            "false_positive_count": b["false_pos"],
            "cases": n,
        }
    resolution_rows.append({"resolution": res, "engines": engines_map})

kind_buckets: dict[str, dict[str, dict[str, float]]] = defaultdict(lambda: defaultdict(lambda: {
    "ms_sum": 0.0,
    "n": 0.0,
    "success_sum": 0.0,
    "not_found_sum": 0.0,
    "false_pos_sum": 0.0,
}))
resolution_kind_buckets: dict[tuple[str, str], dict[str, dict[str, float]]] = defaultdict(lambda: defaultdict(lambda: {
    "ms_sum": 0.0,
    "n": 0.0,
    "success_sum": 0.0,
    "not_found_sum": 0.0,
    "false_pos_sum": 0.0,
}))

for row in results:
    scenario = str(row.get("scenario", ""))
    engine = str(row.get("engine", "unknown"))
    kind = scenario_kind(scenario)
    res = scenario_resolution(scenario)
    metrics = row.get("metrics", {})
    query_total = float(metrics.get("query_total/op", 0.0))
    if query_total > 0:
        denom = max(1.0, query_total)
        success = float(metrics.get("query_ok/op", 0.0)) / denom
        not_found = float(metrics.get("query_not_found/op", 0.0)) / denom
        false_pos = float(metrics.get("query_overlap_miss/op", 0.0)) / denom
    else:
        success = float(metrics.get("success/op", 1.0))
        not_found = float(metrics.get("not_found/op", 0.0))
        false_pos = float(metrics.get("overlap_miss/op", 0.0))
    ms = float(row.get("ms_per_op", 0.0))

    kb = kind_buckets[kind][engine]
    kb["ms_sum"] += ms
    kb["n"] += 1.0
    kb["success_sum"] += success
    kb["not_found_sum"] += not_found
    kb["false_pos_sum"] += false_pos

    rb = resolution_kind_buckets[(res, kind)][engine]
    rb["ms_sum"] += ms
    rb["n"] += 1.0
    rb["success_sum"] += success
    rb["not_found_sum"] += not_found
    rb["false_pos_sum"] += false_pos

kind_rows: list[dict[str, object]] = []
for kind in sorted(kind_buckets.keys()):
    engines_map: dict[str, dict[str, float]] = {}
    for eng in engine_names:
        b = kind_buckets[kind].get(eng)
        if b is None or b["n"] <= 0:
            engines_map[eng] = {
                "avg_ms_per_op": 0.0,
                "success_pct": 0.0,
                "not_found_pct": 0.0,
                "false_positive_pct": 0.0,
                "cases": 0.0,
            }
            continue
        n = b["n"]
        engines_map[eng] = {
            "avg_ms_per_op": b["ms_sum"] / n,
            "success_pct": (b["success_sum"] / n) * 100.0,
            "not_found_pct": (b["not_found_sum"] / n) * 100.0,
            "false_positive_pct": (b["false_pos_sum"] / n) * 100.0,
            "cases": n,
        }
    kind_rows.append({"kind": kind, "engines": engines_map})

resolution_kind_rows: list[dict[str, object]] = []
for (res, kind) in sorted(resolution_kind_buckets.keys(), key=lambda x: (resolution_sort_key(x[0]), x[1])):
    row: dict[str, object] = {"resolution": res, "kind": kind, "engines": {}}
    for eng in engine_names:
        b = resolution_kind_buckets[(res, kind)].get(eng)
        if b is None or b["n"] <= 0:
            row["engines"][eng] = {
                "avg_ms_per_op": 0.0,
                "success_pct": 0.0,
                "not_found_pct": 0.0,
                "false_positive_pct": 0.0,
                "cases": 0.0,
            }
            continue
        n = b["n"]
        row["engines"][eng] = {
            "avg_ms_per_op": b["ms_sum"] / n,
            "success_pct": (b["success_sum"] / n) * 100.0,
            "not_found_pct": (b["not_found_sum"] / n) * 100.0,
            "false_positive_pct": (b["false_pos_sum"] / n) * 100.0,
            "cases": n,
        }
    resolution_kind_rows.append(row)

report = {
    "metadata": meta,
    "summary": {
        "total_rows": len(results),
        "engines": sorted(engines.keys()),
        "by_engine": summary_rows,
        "metrics_chart": metrics_chart_rows,
        "by_resolution": resolution_rows,
        "by_kind": kind_rows,
        "by_resolution_kind": resolution_kind_rows,
    },
    "results": results,
}
json_path.write_text(json.dumps(report, indent=2), encoding="utf-8")

lines: list[str] = []
lines.append("# FindOnScreen Benchmark Report")
lines.append("")
lines.append(f"Generated: `{meta['timestamp_utc']}`")
lines.append("")
lines.append("## Run Metadata")
lines.append("")
lines.append(f"- Package: `{meta['package']}`")
lines.append(f"- Target: `{meta['bench_name']}`")
lines.append(f"- Benchtime: `{meta['benchtime']}`")
lines.append(f"- Count: `{meta['count']}`")
lines.append(f"- Tags: `{meta['tags'] or '(none)'}`")
lines.append(f"- Seed: `{meta['seed'] or '(unset)'}`")
if meta["manifest"]:
    lines.append(f"- Scenario Manifest: `{meta['manifest']}`")
if meta["schema"]:
    lines.append(f"- Scenario Schema: `{meta['schema']}`")
lines.append(f"- Run Mode: `{run_mode}`")
lines.append(f"- High Resolution Scenarios: `{high_res}`")
lines.append(f"- Ultra Resolution Scenarios: `{ultra_res}`")
lines.append(f"- Platform: `{meta['goos']}/{meta['goarch']}`")
lines.append(f"- CPU: `{meta['cpu']}`")
lines.append(f"- Visuals Enabled: `{visual_enable}`")
lines.append(f"- Visual Output: `{visual_dir}`")
lines.append(f"- Visual Max Attempts: `{visual_max_attempts}`")
lines.append(f"- Visual Timeout: `{visual_timeout}`")
lines.append("")
if visual_enable.lower() in {"1", "true", "yes", "on"} and visual_dir:
    visual_path = Path(visual_dir)
    attempt_count = 0
    summary_count = 0
    mega_summary = visual_path / "summaries" / "summary-run-mega.jpg"
    if visual_path.exists():
        attempt_count = len(list((visual_path / "attempts").glob("**/*.png")))
        if not mega_summary.exists():
            mega_summary = visual_path / "summaries" / "summary-run-mega.png"
        summary_count = len([
            p for p in (visual_path / "summaries").glob("summary-*.png")
            if p.name != "summary-run-mega.png" and p.name != "summary-run-mega.jpg"
        ])
    lines.append("## Visual Artifacts")
    lines.append("")
    lines.append(f"- Directory: `{visual_dir}`")
    lines.append(f"- Per-attempt screenshots: `{attempt_count}`")
    lines.append(f"- Per-scenario summary screenshots: `{summary_count}`")
    lines.append(f"- Run-level mega summary: `{mega_summary}`")
    lines.append("")

lines.append("## Engine Summary")
lines.append("")
lines.append("_Cases/OK metrics are query-level counts (regions x scenarios x resolutions), not just benchmark row count._")
lines.append("")
lines.append("| Engine | Cases | OK | Partial | Not Found | Unsupported | Error | Overlap Miss | Avg ms/op | Median ms/op | Best Scenario | Best ms/op | Worst Scenario | Worst ms/op |")
lines.append("|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---|---:|")
for row in summary_rows:
    lines.append(
        "| {engine} | {cases:.0f} | {ok:.0f} | {partial:.0f} | {not_found:.0f} | {unsupported:.0f} | {error:.0f} | {overlap_miss:.0f} | {avg:.3f} | {med:.3f} | `{best_s}` | {best:.3f} | `{worst_s}` | {worst:.3f} |".format(
            engine=row["engine"],
            cases=row["cases"],
            ok=row.get("query_ok", 0.0),
            partial=row.get("query_partial", 0.0),
            not_found=row.get("query_not_found", 0.0),
            unsupported=row.get("query_unsupported", 0.0),
            error=row.get("query_error", 0.0),
            overlap_miss=row.get("query_overlap_miss", 0.0),
            avg=row["avg_ms_per_op"],
            med=row["median_ms_per_op"],
            best_s=row["best_scenario"],
            best=row["best_ms_per_op"],
            worst_s=row["worst_scenario"],
            worst=row["worst_ms_per_op"],
        )
    )

lines.append("")
lines.append("## Summary Metrics Chart")
lines.append("")
lines.append("| Engine | Cases | Benchmark Rows | Avg ms/op | Median ms/op | Success % | False Positive % | No Match % | Unsupported % | Error % |")
lines.append("|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|")
for row in metrics_chart_rows:
    lines.append(
        "| {engine} | {cases:.0f} | {rows} | {avg:.3f} | {med:.3f} | {success:.1f} | {fp:.1f} | {nf:.1f} | {unsup:.1f} | {err:.1f} |".format(
            engine=row["engine"],
            cases=row["cases"],
            rows=row.get("rows", 0),
            avg=row["avg_ms_per_op"],
            med=row["median_ms_per_op"],
            success=row["success_rate_pct"],
            fp=row["false_positive_rate_pct"],
            nf=row["not_found_rate_pct"],
            unsup=row["unsupported_rate_pct"],
            err=row["error_rate_pct"],
        )
    )

if metrics_chart_rows:
    engines_axis = [str(r["engine"]) for r in metrics_chart_rows]
    perf_vals = [round(float(r["avg_ms_per_op"]), 3) for r in metrics_chart_rows]
    success_vals = [round(float(r["success_rate_pct"]), 1) for r in metrics_chart_rows]
    false_positive_vals = [round(float(r["false_positive_rate_pct"]), 1) for r in metrics_chart_rows]
    axis_labels = ", ".join([f'"{e}"' for e in engines_axis])
    perf_max = max(perf_vals) if perf_vals else 1.0
    perf_ymax = max(1, int((perf_max + 1.0) // 1 + 1))

    lines.append("")
    lines.append("## Performance Graph")
    lines.append("")
    lines.append("```mermaid")
    lines.append("xychart-beta")
    lines.append('  title "Engine Average Latency (ms/op)"')
    lines.append(f"  x-axis [{axis_labels}]")
    lines.append(f'  y-axis "ms/op" 0 --> {perf_ymax}')
    lines.append(f'  bar [{", ".join([str(v) for v in perf_vals])}]')
    lines.append("```")

    lines.append("")
    lines.append("## Accuracy Graph")
    lines.append("")
    lines.append("```mermaid")
    lines.append("xychart-beta")
    lines.append('  title "Engine Accuracy and False Positive Rate (%)"')
    lines.append(f"  x-axis [{axis_labels}]")
    lines.append('  y-axis "percent" 0 --> 100')
    lines.append(f'  bar [{", ".join([str(v) for v in success_vals])}]')
    lines.append(f'  line [{", ".join([str(v) for v in false_positive_vals])}]')
    lines.append("```")

    perf_svg = render_single_series_bar_chart_svg(
        title="Engine Average Latency (ms/op)",
        y_label="ms/op",
        labels=engines_axis,
        values=perf_vals,
        fill="#2563eb",
    )
    acc_svg = render_grouped_accuracy_chart_svg(
        title="Engine Accuracy vs False Positive Rate (%)",
        labels=engines_axis,
        success_vals=success_vals,
        false_pos_vals=false_positive_vals,
    )
    write_svg(perf_svg_path, perf_svg)
    write_svg(accuracy_svg_path, acc_svg)
    rel_perf_md = os.path.relpath(perf_svg_path, md_path.parent).replace(os.sep, "/")
    rel_acc_md = os.path.relpath(accuracy_svg_path, md_path.parent).replace(os.sep, "/")

    lines.append("")
    lines.append("## Static Graphs (SVG)")
    lines.append("")
    lines.append(f"- [Performance SVG]({rel_perf_md})")
    lines.append(f"- [Accuracy SVG]({rel_acc_md})")
    lines.append("")
    lines.append(f"![Performance SVG]({rel_perf_md})")
    lines.append("")
    lines.append(f"![Accuracy SVG]({rel_acc_md})")

if resolution_rows and engine_names:
    res_labels = [str(r["resolution"]) for r in resolution_rows]
    time_series: list[tuple[str, list[float], str]] = []
    match_series: list[tuple[str, list[float], str]] = []
    miss_series: list[tuple[str, list[float], str]] = []
    false_pos_series: list[tuple[str, list[float], str]] = []

    for i, eng in enumerate(engine_names):
        color = engine_color(eng, i)
        t_vals: list[float] = []
        ok_vals: list[float] = []
        miss_vals: list[float] = []
        fp_vals: list[float] = []
        for res_row in resolution_rows:
            emap = res_row["engines"]
            stats = emap.get(eng, {})
            t_vals.append(round(float(stats.get("avg_ms_per_op", 0.0)), 3))
            ok_vals.append(round(float(stats.get("match_count", 0.0)), 3))
            miss_vals.append(round(float(stats.get("miss_count", 0.0)), 3))
            fp_vals.append(round(float(stats.get("false_positive_count", 0.0)), 3))
        time_series.append((eng, t_vals, color))
        match_series.append((eng, ok_vals, color))
        miss_series.append((eng, miss_vals, color))
        false_pos_series.append((eng, fp_vals, color))

    write_svg(
        resolution_time_svg_path,
        render_grouped_multi_series_bar_chart_svg(
            title="Resolution Groups: Match Time by Engine",
            y_label="ms/op",
            group_labels=res_labels,
            series=time_series,
        ),
    )
    write_svg(
        resolution_matches_svg_path,
        render_grouped_multi_series_bar_chart_svg(
            title="Resolution Groups: Matches by Engine",
            y_label="match count",
            group_labels=res_labels,
            series=match_series,
        ),
    )
    write_svg(
        resolution_misses_svg_path,
        render_grouped_multi_series_bar_chart_svg(
            title="Resolution Groups: Misses by Engine",
            y_label="miss count",
            group_labels=res_labels,
            series=miss_series,
        ),
    )
    write_svg(
        resolution_false_pos_svg_path,
        render_grouped_multi_series_bar_chart_svg(
            title="Resolution Groups: False Positives by Engine",
            y_label="false positive count",
            group_labels=res_labels,
            series=false_pos_series,
        ),
    )

    rel_time_md = os.path.relpath(resolution_time_svg_path, md_path.parent).replace(os.sep, "/")
    rel_match_md = os.path.relpath(resolution_matches_svg_path, md_path.parent).replace(os.sep, "/")
    rel_miss_md = os.path.relpath(resolution_misses_svg_path, md_path.parent).replace(os.sep, "/")
    rel_fp_md = os.path.relpath(resolution_false_pos_svg_path, md_path.parent).replace(os.sep, "/")

    lines.append("")
    lines.append("## Resolution Group Graphs (SVG)")
    lines.append("")
    lines.append(f"- [Resolution Match Time]({rel_time_md})")
    lines.append(f"- [Resolution Matches]({rel_match_md})")
    lines.append(f"- [Resolution Misses]({rel_miss_md})")
    lines.append(f"- [Resolution False Positives]({rel_fp_md})")
    lines.append("")
    lines.append(f"![Resolution Match Time]({rel_time_md})")
    lines.append("")
    lines.append(f"![Resolution Matches]({rel_match_md})")
    lines.append("")
    lines.append(f"![Resolution Misses]({rel_miss_md})")
    lines.append("")
    lines.append(f"![Resolution False Positives]({rel_fp_md})")

if kind_rows and engine_names:
    kind_labels = [str(r["kind"]) for r in kind_rows]
    kind_time_series: list[tuple[str, list[float], str]] = []
    kind_success_series: list[tuple[str, list[float], str]] = []
    for i, eng in enumerate(engine_names):
        color = engine_color(eng, i)
        t_vals: list[float] = []
        s_vals: list[float] = []
        for row in kind_rows:
            stats = row["engines"].get(eng, {})
            t_vals.append(round(float(stats.get("avg_ms_per_op", 0.0)), 3))
            s_vals.append(round(float(stats.get("success_pct", 0.0)), 2))
        kind_time_series.append((eng, t_vals, color))
        kind_success_series.append((eng, s_vals, color))

    write_svg(
        kind_time_svg_path,
        render_grouped_multi_series_bar_chart_svg(
            title="Scenario Kind: Match Time by Engine",
            y_label="ms/op",
            group_labels=kind_labels,
            series=kind_time_series,
        ),
    )
    write_svg(
        kind_success_svg_path,
        render_grouped_multi_series_bar_chart_svg(
            title="Scenario Kind: Success % by Engine",
            y_label="success percent",
            group_labels=kind_labels,
            series=kind_success_series,
        ),
    )

    rel_kind_time_md = os.path.relpath(kind_time_svg_path, md_path.parent).replace(os.sep, "/")
    rel_kind_success_md = os.path.relpath(kind_success_svg_path, md_path.parent).replace(os.sep, "/")

    lines.append("")
    lines.append("## Scenario Kind Graphs (SVG)")
    lines.append("")
    lines.append(f"- [Kind Match Time]({rel_kind_time_md})")
    lines.append(f"- [Kind Success Rate]({rel_kind_success_md})")
    lines.append("")
    lines.append(f"![Kind Match Time]({rel_kind_time_md})")
    lines.append("")
    lines.append(f"![Kind Success Rate]({rel_kind_success_md})")

if resolution_kind_rows:
    lines.append("")
    lines.append("## Resolution + Scenario Kind Breakdown")
    lines.append("")
    lines.append("| Resolution | Scenario Kind | Engine | Cases | Avg ms/op | Success % | Not Found % | False Positive % |")
    lines.append("|---|---|---|---:|---:|---:|---:|---:|")
    for rk in resolution_kind_rows:
        res = str(rk["resolution"])
        kind = str(rk["kind"])
        eng_map = rk["engines"]
        for eng in engine_names:
            stats = eng_map.get(eng, {})
            lines.append(
                "| {res} | `{kind}` | {eng} | {cases:.0f} | {avg:.3f} | {success:.2f} | {miss:.2f} | {fp:.2f} |".format(
                    res=res,
                    kind=kind,
                    eng=eng,
                    cases=float(stats.get("cases", 0.0)),
                    avg=float(stats.get("avg_ms_per_op", 0.0)),
                    success=float(stats.get("success_pct", 0.0)),
                    miss=float(stats.get("not_found_pct", 0.0)),
                    fp=float(stats.get("false_positive_pct", 0.0)),
                )
            )

for engine in sorted(engines.keys()):
    lines.append("")
    lines.append(f"## Engine: `{engine}`")
    lines.append("")
    lines.append("| Scenario | Status | Iterations | ms/op | ns/op | Success % | Not Found % | Unsupported % | Error % | Overlap Miss % | KB/op | allocs/op |")
    lines.append("|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|")
    for row in engines[engine]:
        bytes_per_op = row.get("bytes_per_op")
        allocs_per_op = row.get("allocs_per_op")
        metrics = row.get("metrics", {})
        query_total = float(metrics.get("query_total/op", 0.0))
        if query_total > 0:
            denom = max(1.0, query_total)
            success_rate = (float(metrics.get("query_ok/op", 0.0)) / denom) * 100.0
            not_found_rate = (float(metrics.get("query_not_found/op", 0.0)) / denom) * 100.0
            unsupported_rate = (float(metrics.get("query_unsupported/op", 0.0)) / denom) * 100.0
            error_rate = (float(metrics.get("query_error/op", 0.0)) / denom) * 100.0
            overlap_miss_rate = (float(metrics.get("query_overlap_miss/op", 0.0)) / denom) * 100.0
        else:
            success_rate = float(metrics.get("success/op", 1.0)) * 100.0
            not_found_rate = float(metrics.get("not_found/op", 0.0)) * 100.0
            unsupported_rate = float(metrics.get("unsupported/op", 0.0)) * 100.0
            error_rate = float(metrics.get("error/op", 0.0)) * 100.0
            overlap_miss_rate = float(metrics.get("overlap_miss/op", 0.0)) * 100.0
        kb = ""
        if isinstance(bytes_per_op, (int, float)):
            kb = f"{float(bytes_per_op)/1024.0:.2f}"
        allocs = ""
        if isinstance(allocs_per_op, (int, float)):
            allocs = f"{float(allocs_per_op):.2f}"
        lines.append(
            "| `{scenario}` | `{status}` | {iters} | {ms:.3f} | {ns:.0f} | {success:.1f} | {not_found:.1f} | {unsupported:.1f} | {error:.1f} | {overlap_miss:.1f} | {kb} | {allocs} |".format(
                scenario=row["scenario"],
                status=row["status"],
                iters=row["iterations"],
                ms=row["ms_per_op"],
                ns=row["ns_per_op"],
                success=success_rate,
                not_found=not_found_rate,
                unsupported=unsupported_rate,
                error=error_rate,
                overlap_miss=overlap_miss_rate,
                kb=kb,
                allocs=allocs,
            )
        )

md_path.write_text("\n".join(lines) + "\n", encoding="utf-8")

def env_true(raw: str) -> bool:
    return raw.strip().lower() in {"1", "true", "yes", "on"}

published_report_dir = report_dir
if docs_publish and docs_report_dir_raw and report_dir.exists():
    docs_report_dir = Path(docs_report_dir_raw)
    if not docs_report_dir.is_absolute():
        docs_report_dir = (root_dir / docs_report_dir).resolve()
    if docs_report_dir != report_dir:
        if docs_report_dir.exists():
            shutil.rmtree(docs_report_dir)
        docs_report_dir.parent.mkdir(parents=True, exist_ok=True)
        shutil.copytree(report_dir, docs_report_dir)
    published_report_dir = docs_report_dir
    print(f"[find-bench] published docs report dir: {published_report_dir}")

def published_path(path: Path) -> Path:
    try:
        rel = path.resolve().relative_to(report_dir)
    except Exception:
        return path
    return published_report_dir / rel

def to_rel_link(base_dir: Path, target: Path) -> str:
    try:
        rel = os.path.relpath(target, base_dir)
    except ValueError:
        rel = str(target)
    return rel.replace(os.sep, "/")

def to_readme_link(base_dir: Path, target: Path) -> str:
    if readme_link_mode == "pages" and readme_base_url:
        try:
            rel = target.resolve().relative_to(docs_root)
            return f"{readme_base_url}/{rel.as_posix()}"
        except Exception:
            pass
    return to_rel_link(base_dir, target)

IMAGE_EXTS = {".png", ".jpg", ".jpeg", ".webp", ".gif", ".svg", ".bmp", ".avif"}

def write_gallery_index_js(target_dir: Path, title: str, recursive: bool) -> None:
    target_dir.mkdir(parents=True, exist_ok=True)
    index_js = target_dir / "index.js"
    index_html = target_dir / "index.html"

    def is_gallery_support_file(p: Path) -> bool:
        return p.name in {"index.js", "index.html"}

    def to_href(base: Path, p: Path) -> str:
        return os.path.relpath(p, base).replace(os.sep, "/")

    files: list[Path] = []
    if recursive:
        files = sorted(p for p in target_dir.rglob("*") if p.is_file() and not is_gallery_support_file(p))
    else:
        files = sorted(p for p in target_dir.glob("*") if p.is_file() and not is_gallery_support_file(p))

    entries: list[dict[str, object]] = []
    for p in files:
        ext = p.suffix.lower()
        href = to_href(target_dir, p)
        name = href
        is_image = ext in IMAGE_EXTS
        entries.append({
            "name": name,
            "href": href,
            "isImage": is_image,
        })

    index_js.write_text(
        "\n".join(
            [
                f"const GALLERY_TITLE = {json.dumps(title)};",
                f"const GALLERY_ENTRIES = {json.dumps(entries, ensure_ascii=True)};",
                "",
                "const app = document.getElementById('app');",
                "const titleEl = document.getElementById('title');",
                "if (titleEl) titleEl.textContent = GALLERY_TITLE;",
                "",
                "if (!GALLERY_ENTRIES.length) {",
                "  const empty = document.createElement('p');",
                "  empty.className = 'empty';",
                "  empty.textContent = 'No files found in this directory.';",
                "  app.appendChild(empty);",
                "} else {",
                "  for (const entry of GALLERY_ENTRIES) {",
                "    const card = document.createElement('article');",
                "    card.className = 'card';",
                "    const label = document.createElement('div');",
                "    label.className = 'label';",
                "    label.textContent = entry.name;",
                "    const link = document.createElement('a');",
                "    link.href = entry.href;",
                "    link.target = '_blank';",
                "    link.rel = 'noopener noreferrer';",
                "    link.className = 'open';",
                "    link.textContent = 'Open file';",
                "    card.appendChild(label);",
                "    if (entry.isImage) {",
                "      const img = document.createElement('img');",
                "      img.src = entry.href;",
                "      img.alt = entry.name;",
                "      img.loading = 'lazy';",
                "      img.className = 'preview';",
                "      card.appendChild(img);",
                "    } else {",
                "      const meta = document.createElement('div');",
                "      meta.className = 'meta';",
                "      meta.textContent = 'Non-image file';",
                "      card.appendChild(meta);",
                "    }",
                "    card.appendChild(link);",
                "    app.appendChild(card);",
                "  }",
                "}",
                "",
            ]
        )
        + "\n",
        encoding="utf-8",
    )

    index_html.write_text(
        """<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width,initial-scale=1" />
  <title>Benchmark Visual Gallery</title>
  <style>
    :root { color-scheme: light dark; }
    body {
      margin: 0;
      padding: 20px;
      font-family: ui-sans-serif, system-ui, -apple-system, "Segoe UI", Roboto, sans-serif;
      background: #0b1220;
      color: #e5e7eb;
    }
    h1 {
      margin: 0 0 14px;
      font-size: 24px;
      font-weight: 700;
    }
    .hint {
      margin: 0 0 18px;
      color: #94a3b8;
      font-size: 14px;
    }
    #app {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(520px, 1fr));
      gap: 16px;
    }
    .card {
      background: #111827;
      border: 1px solid #334155;
      border-radius: 10px;
      padding: 12px;
      display: flex;
      flex-direction: column;
      gap: 10px;
      min-height: 140px;
    }
    .label {
      font-size: 13px;
      color: #93c5fd;
      word-break: break-all;
    }
    .preview {
      width: 100%;
      min-height: 240px;
      max-height: 700px;
      object-fit: contain;
      background: #0f172a;
      border-radius: 8px;
      border: 1px solid #334155;
    }
    .meta {
      color: #94a3b8;
      font-size: 14px;
      padding: 12px;
      border: 1px dashed #334155;
      border-radius: 8px;
      background: #0f172a;
    }
    .open {
      color: #22c55e;
      text-decoration: none;
      font-weight: 600;
      font-size: 14px;
    }
    .open:hover { text-decoration: underline; }
    .empty {
      color: #94a3b8;
      font-size: 14px;
    }
  </style>
</head>
<body>
  <h1 id="title">Benchmark Visual Gallery</h1>
  <p class="hint">Large previews for files in this directory.</p>
  <section id="app"></section>
  <script src="./index.js"></script>
</body>
</html>
""",
        encoding="utf-8",
    )

def select_top_resolution_summary_per_scenario(images: list[Path]) -> list[Path]:
    # Filename shape: summary-<scenario>_<width>x<height>_iNN[...].png
    summary_re = re.compile(r"^summary-(.+)_(\d+)x(\d+)_i\d+(?:_.+)?\.png$")
    best: dict[str, tuple[int, int, int, Path]] = {}
    for img in images:
        name = img.name
        m = summary_re.match(name)
        if m:
            scenario_key = m.group(1)
            width = int(m.group(2))
            height = int(m.group(3))
        else:
            scenario_key = img.stem.replace("summary-", "")
            width = 0
            height = 0
        area = width * height
        current = best.get(scenario_key)
        candidate = (area, width, height, img)
        if current is None or candidate[:3] > current[:3] or (
            candidate[:3] == current[:3] and img.name < current[3].name
        ):
            best[scenario_key] = candidate
    return [best[key][3] for key in sorted(best.keys())]

def patch_readme(path: Path, block_lines: list[str]) -> None:
    begin = "<!-- BEGIN: FIND_ON_SCREEN_BENCH_AUTOGEN -->"
    end = "<!-- END: FIND_ON_SCREEN_BENCH_AUTOGEN -->"
    block = "\n".join([begin, *block_lines, end]) + "\n"

    original = ""
    if path.exists():
        original = path.read_text(encoding="utf-8", errors="replace")
        original = re.sub(
            rf"\n?{re.escape(begin)}.*?{re.escape(end)}\n?",
            "\n",
            original,
            flags=re.S,
        ).rstrip()

    if original:
        updated = f"{original}\n\n{block}"
    else:
        updated = block
    path.write_text(updated, encoding="utf-8")

if env_true(patch_readmes):
    report_text_path = published_path(text_path)
    report_json_path = published_path(json_path)
    report_md_path = published_path(md_path)
    report_perf_svg_path = published_path(perf_svg_path)
    report_accuracy_svg_path = published_path(accuracy_svg_path)
    report_kind_time_svg_path = published_path(kind_time_svg_path)
    report_kind_success_svg_path = published_path(kind_success_svg_path)
    report_resolution_time_svg_path = published_path(resolution_time_svg_path)
    report_resolution_matches_svg_path = published_path(resolution_matches_svg_path)
    report_resolution_misses_svg_path = published_path(resolution_misses_svg_path)
    report_resolution_false_pos_svg_path = published_path(resolution_false_pos_svg_path)

    root_visual = Path(visual_dir) if visual_dir else None
    if root_visual is not None:
        root_visual = published_path(root_visual)
    mega_summary = None
    scenario_summaries: list[Path] = []
    attempt_images: list[Path] = []
    if root_visual and root_visual.exists():
        maybe_mega = root_visual / "summaries" / "summary-run-mega.jpg"
        if not maybe_mega.exists():
            maybe_mega = root_visual / "summaries" / "summary-run-mega.png"
        if maybe_mega.exists():
            mega_summary = maybe_mega
        scenario_summaries = sorted(
            p for p in (root_visual / "summaries").glob("summary-*.png")
            if p.name != "summary-run-mega.png" and p.name != "summary-run-mega.jpg"
        )
        scenario_summaries = select_top_resolution_summary_per_scenario(scenario_summaries)
        attempt_images = sorted((root_visual / "attempts").glob("**/*.png"))
        write_gallery_index_js(root_visual, "FindOnScreen Visual Root", recursive=False)
        write_gallery_index_js(root_visual / "summaries", "FindOnScreen Scenario Summaries", recursive=False)
        write_gallery_index_js(root_visual / "attempts", "FindOnScreen Attempt Images", recursive=True)

    readmes: list[Path] = []
    for raw_path in readme_paths.split(","):
        raw_path = raw_path.strip()
        if not raw_path:
            continue
        p = Path(raw_path)
        if not p.is_absolute():
            p = (root_dir / p).resolve()
        readmes.append(p)
    for readme in readmes:
        if not readme.exists():
            print(f"[find-bench] skip missing readme: {readme}")
            continue
        section: list[str] = []
        section.append(f"## {readme_section_title}")
        section.append("")
        section.append(f"Generated: `{meta['timestamp_utc']}`")
        section.append("")
        section.append("### Reports")
        section.append("")
        section.append(f"- [Markdown Summary]({to_readme_link(readme.parent, report_md_path)})")
        section.append(f"- [JSON Report]({to_readme_link(readme.parent, report_json_path)})")
        section.append(f"- [Raw go test Output]({to_readme_link(readme.parent, report_text_path)})")
        if report_perf_svg_path.exists():
            section.append(f"- [Performance SVG]({to_readme_link(readme.parent, report_perf_svg_path)})")
        if report_accuracy_svg_path.exists():
            section.append(f"- [Accuracy SVG]({to_readme_link(readme.parent, report_accuracy_svg_path)})")
        if report_kind_time_svg_path.exists():
            section.append(f"- [Scenario Kind Match Time SVG]({to_readme_link(readme.parent, report_kind_time_svg_path)})")
        if report_kind_success_svg_path.exists():
            section.append(f"- [Scenario Kind Success SVG]({to_readme_link(readme.parent, report_kind_success_svg_path)})")
        if report_resolution_time_svg_path.exists():
            section.append(f"- [Resolution Match Time SVG]({to_readme_link(readme.parent, report_resolution_time_svg_path)})")
        if report_resolution_matches_svg_path.exists():
            section.append(f"- [Resolution Matches SVG]({to_readme_link(readme.parent, report_resolution_matches_svg_path)})")
        if report_resolution_misses_svg_path.exists():
            section.append(f"- [Resolution Misses SVG]({to_readme_link(readme.parent, report_resolution_misses_svg_path)})")
        if report_resolution_false_pos_svg_path.exists():
            section.append(f"- [Resolution False Positives SVG]({to_readme_link(readme.parent, report_resolution_false_pos_svg_path)})")
        section.append("")
        section.append("### Engine Summary")
        section.append("")
        section.append("_Cases/OK metrics are query-level counts (regions x scenarios x resolutions), not just benchmark row count._")
        section.append("")
        section.append("| Engine | Cases | OK | Partial | Not Found | Unsupported | Error | Overlap Miss | Avg ms/op | Median ms/op |")
        section.append("|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|")
        for row in summary_rows:
            section.append(
                "| {engine} | {cases:.0f} | {ok:.0f} | {partial:.0f} | {not_found:.0f} | {unsupported:.0f} | {error:.0f} | {overlap_miss:.0f} | {avg:.3f} | {med:.3f} |".format(
                    engine=row["engine"],
                    cases=row["cases"],
                    ok=row.get("query_ok", 0.0),
                    partial=row.get("query_partial", 0.0),
                    not_found=row.get("query_not_found", 0.0),
                    unsupported=row.get("query_unsupported", 0.0),
                    error=row.get("query_error", 0.0),
                    overlap_miss=row.get("query_overlap_miss", 0.0),
                    avg=row["avg_ms_per_op"],
                    med=row["median_ms_per_op"],
                )
            )

        if mega_summary is not None:
            section.append("")
            section.append("### Run Mega Summary")
            section.append("")
            section.append(f"![Run Mega Summary]({to_readme_link(readme.parent, mega_summary)})")
            section.append("")
            section.append(f"- [Open run mega summary image]({to_readme_link(readme.parent, mega_summary)})")

        if report_perf_svg_path.exists() or report_accuracy_svg_path.exists():
            section.append("")
            section.append("### Benchmark Graphs")
            section.append("")
            if report_perf_svg_path.exists():
                rel_perf = to_readme_link(readme.parent, report_perf_svg_path)
                section.append(f"![Performance Graph]({rel_perf})")
                section.append("")
                section.append(f"- [Open performance graph]({rel_perf})")
            if report_accuracy_svg_path.exists():
                rel_acc = to_readme_link(readme.parent, report_accuracy_svg_path)
                section.append("")
                section.append(f"![Accuracy Graph]({rel_acc})")
                section.append("")
                section.append(f"- [Open accuracy graph]({rel_acc})")

        if report_kind_time_svg_path.exists() or report_kind_success_svg_path.exists():
            section.append("")
            section.append("### Scenario Kind Graphs")
            section.append("")
            if report_kind_time_svg_path.exists():
                rel = to_readme_link(readme.parent, report_kind_time_svg_path)
                section.append(f"![Scenario Kind Match Time]({rel})")
                section.append("")
                section.append(f"- [Open scenario kind match time graph]({rel})")
                section.append("")
            if report_kind_success_svg_path.exists():
                rel = to_readme_link(readme.parent, report_kind_success_svg_path)
                section.append(f"![Scenario Kind Success]({rel})")
                section.append("")
                section.append(f"- [Open scenario kind success graph]({rel})")

        if report_resolution_time_svg_path.exists() or report_resolution_matches_svg_path.exists() or report_resolution_misses_svg_path.exists() or report_resolution_false_pos_svg_path.exists():
            section.append("")
            section.append("### Resolution Group Graphs")
            section.append("")
            if report_resolution_time_svg_path.exists():
                rel = to_readme_link(readme.parent, report_resolution_time_svg_path)
                section.append(f"![Resolution Match Time]({rel})")
                section.append("")
                section.append(f"- [Open resolution match time graph]({rel})")
                section.append("")
            if report_resolution_matches_svg_path.exists():
                rel = to_readme_link(readme.parent, report_resolution_matches_svg_path)
                section.append(f"![Resolution Matches]({rel})")
                section.append("")
                section.append(f"- [Open resolution matches graph]({rel})")
                section.append("")
            if report_resolution_misses_svg_path.exists():
                rel = to_readme_link(readme.parent, report_resolution_misses_svg_path)
                section.append(f"![Resolution Misses]({rel})")
                section.append("")
                section.append(f"- [Open resolution misses graph]({rel})")
                section.append("")
            if report_resolution_false_pos_svg_path.exists():
                rel = to_readme_link(readme.parent, report_resolution_false_pos_svg_path)
                section.append(f"![Resolution False Positives]({rel})")
                section.append("")
                section.append(f"- [Open resolution false positives graph]({rel})")

        if root_visual is not None:
            section.append("")
            section.append("### Artifact Directories")
            section.append("")
            section.append(f"- [Visual Root]({to_readme_link(readme.parent, root_visual / 'index.html')})")
            section.append(f"- [Scenario Summaries]({to_readme_link(readme.parent, root_visual / 'summaries' / 'index.html')})")
            section.append(f"- [Attempt Images]({to_readme_link(readme.parent, root_visual / 'attempts' / 'index.html')})")

        if scenario_summaries:
            section.append("")
            section.append(f"### Scenario Summary Images ({len(scenario_summaries)})")
            section.append("")
            for img in scenario_summaries:
                name = img.stem.replace("summary-", "")
                rel = to_readme_link(readme.parent, img)
                section.append(f"#### `{name}`")
                section.append("")
                section.append(f"![{name}]({rel})")
                section.append("")
                section.append(f"- [Open `{name}` image]({rel})")
                section.append("")

        patch_readme(readme, section)
        print(f"[find-bench] patched readme: {readme}")

print(f"[find-bench] wrote json report: {json_path}")
print(f"[find-bench] wrote markdown report: {md_path}")
if perf_svg_path.exists():
    print(f"[find-bench] wrote performance graph: {perf_svg_path}")
if accuracy_svg_path.exists():
    print(f"[find-bench] wrote accuracy graph: {accuracy_svg_path}")
if kind_time_svg_path.exists():
    print(f"[find-bench] wrote scenario kind time graph: {kind_time_svg_path}")
if kind_success_svg_path.exists():
    print(f"[find-bench] wrote scenario kind success graph: {kind_success_svg_path}")
if resolution_time_svg_path.exists():
    print(f"[find-bench] wrote resolution time graph: {resolution_time_svg_path}")
if resolution_matches_svg_path.exists():
    print(f"[find-bench] wrote resolution matches graph: {resolution_matches_svg_path}")
if resolution_misses_svg_path.exists():
    print(f"[find-bench] wrote resolution misses graph: {resolution_misses_svg_path}")
if resolution_false_pos_svg_path.exists():
    print(f"[find-bench] wrote resolution false positive graph: {resolution_false_pos_svg_path}")
PY
