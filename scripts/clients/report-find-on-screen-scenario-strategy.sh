#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${THIS_DIR}/paths.sh"

MANIFEST_RAW="${FIND_BENCH_SCENARIO_MANIFEST:-example}"
if [[ "${MANIFEST_RAW}" == "example" || "${MANIFEST_RAW}" == "default" || -z "${MANIFEST_RAW}" ]]; then
  MANIFEST_RAW="docs/bench/find-on-screen-scenarios.example.json"
fi

REPORT_DIR="${FIND_BENCH_REPORT_DIR:-${ROOT_DIR}/.test-results/bench}"
STRATEGY_BASENAME="${FIND_BENCH_STRATEGY_BASENAME:-find-on-screen-scenario-strategy}"
STRATEGY_JSON_OUT="${FIND_BENCH_STRATEGY_JSON_OUT:-${REPORT_DIR}/${STRATEGY_BASENAME}.json}"
STRATEGY_MD_OUT="${FIND_BENCH_STRATEGY_MD_OUT:-${REPORT_DIR}/${STRATEGY_BASENAME}.md}"
STRATEGY_BENCH_JSON="${FIND_BENCH_STRATEGY_BENCH_JSON:-${REPORT_DIR}/find-on-screen-e2e.json}"
STRATEGY_VISUAL_EXAMPLES="${FIND_BENCH_STRATEGY_VISUAL_EXAMPLES:-1}"
STRATEGY_VISUAL_RES="${FIND_BENCH_STRATEGY_VISUAL_RES:-1280x720}"
STRATEGY_VISUAL_ENGINE="${FIND_BENCH_STRATEGY_VISUAL_ENGINE:-hybrid}"
STRATEGY_VISUAL_BENCHTIME="${FIND_BENCH_STRATEGY_VISUAL_BENCHTIME:-1x}"
STRATEGY_VISUAL_COUNT="${FIND_BENCH_STRATEGY_VISUAL_COUNT:-1}"
STRATEGY_VISUAL_TAGS="${FIND_BENCH_STRATEGY_VISUAL_TAGS:-opencv gocv_specific_modules gocv_features2d gocv_calib3d}"
STRATEGY_VISUAL_MAX_ATTEMPTS="${FIND_BENCH_STRATEGY_VISUAL_MAX_ATTEMPTS:-2}"
STRATEGY_VISUAL_TIMEOUT="${FIND_BENCH_STRATEGY_VISUAL_TIMEOUT:-5s}"
STRATEGY_VISUAL_DIR="${FIND_BENCH_STRATEGY_VISUAL_DIR:-${REPORT_DIR}/strategy-visuals-${STRATEGY_VISUAL_RES}}"

mkdir -p "${REPORT_DIR}"

echo "[find-bench-strategy] manifest=${MANIFEST_RAW}"
echo "[find-bench-strategy] json=${STRATEGY_JSON_OUT}"
echo "[find-bench-strategy] markdown=${STRATEGY_MD_OUT}"
echo "[find-bench-strategy] bench_json=${STRATEGY_BENCH_JSON}"
echo "[find-bench-strategy] visual_examples=${STRATEGY_VISUAL_EXAMPLES} resolution=${STRATEGY_VISUAL_RES} engine=${STRATEGY_VISUAL_ENGINE}"
echo "[find-bench-strategy] visual_dir=${STRATEGY_VISUAL_DIR} benchtime=${STRATEGY_VISUAL_BENCHTIME} count=${STRATEGY_VISUAL_COUNT}"

PROJECT_ROOT="${ROOT_DIR}" \
MANIFEST_RAW="${MANIFEST_RAW}" \
STRATEGY_JSON_OUT="${STRATEGY_JSON_OUT}" \
STRATEGY_MD_OUT="${STRATEGY_MD_OUT}" \
STRATEGY_BENCH_JSON="${STRATEGY_BENCH_JSON}" \
STRATEGY_VISUAL_EXAMPLES="${STRATEGY_VISUAL_EXAMPLES}" \
STRATEGY_VISUAL_RES="${STRATEGY_VISUAL_RES}" \
STRATEGY_VISUAL_ENGINE="${STRATEGY_VISUAL_ENGINE}" \
STRATEGY_VISUAL_BENCHTIME="${STRATEGY_VISUAL_BENCHTIME}" \
STRATEGY_VISUAL_COUNT="${STRATEGY_VISUAL_COUNT}" \
STRATEGY_VISUAL_TAGS="${STRATEGY_VISUAL_TAGS}" \
STRATEGY_VISUAL_DIR="${STRATEGY_VISUAL_DIR}" \
STRATEGY_VISUAL_MAX_ATTEMPTS="${STRATEGY_VISUAL_MAX_ATTEMPTS}" \
STRATEGY_VISUAL_TIMEOUT="${STRATEGY_VISUAL_TIMEOUT}" \
BENCH_SCHEMA="${FIND_BENCH_SCENARIO_SCHEMA:-}" \
BENCH_PHOTO_ASSET="${FIND_BENCH_PHOTO_ASSET:-}" \
python3 - <<'PY'
from __future__ import annotations

import copy
import json
import os
import re
import shlex
import subprocess
from collections import Counter
from datetime import datetime, timezone
from pathlib import Path
from typing import Any


def resolve_manifest(raw: str, root: Path) -> Path:
    raw = (raw or "").strip()
    candidates: list[Path] = []
    p = Path(raw)
    candidates.append(p)
    if not p.is_absolute():
        cwd = Path.cwd()
        candidates.append(cwd / p)
        candidates.append(cwd / ".." / ".." / p)
        candidates.append(root / p)
    for c in candidates:
        c = c.resolve()
        if c.exists() and c.is_file():
            return c
    joined = ", ".join(str(c.resolve()) for c in candidates)
    raise FileNotFoundError(f"manifest not found: {raw} candidates=[{joined}]")


def summarize_transforms(t: dict[str, Any]) -> dict[str, Any]:
    perspective = (t.get("perspective") or {})
    return {
        "scale": t.get("scale"),
        "rotate": t.get("rotate"),
        "perspective_enabled": bool(perspective.get("enabled", False)),
        "perspective_corner_shift_pct": perspective.get("corner_shift_pct"),
        "skew_x": t.get("skew_x"),
        "skew_y": t.get("skew_y"),
    }


def summarize_noise(photometric: dict[str, Any]) -> dict[str, Any]:
    noise_specs = photometric.get("noise") or []
    types: list[str] = []
    amounts: dict[str, Any] = {}
    for n in noise_specs:
        ntype = str(n.get("type", "")).strip()
        if not ntype:
            continue
        types.append(ntype)
        amounts[ntype] = n.get("amount")
    return {
        "types": sorted(set(types)),
        "amounts": amounts,
    }


def scenario_goal_line(s: dict[str, Any]) -> str:
    sid = s.get("id", "")
    kind = s.get("kind", "")
    expected = s.get("expected") or {}
    decoys = s.get("decoys") or {}
    occ = s.get("occlusion") or {}
    positive = bool(expected.get("positive", True))
    iou = expected.get("iou_min")
    area = expected.get("area_ratio_max")
    partial = bool(expected.get("allow_partial", False))
    decoy_place = decoys.get("placement", "n/a")
    occ_enabled = bool(occ.get("enabled", False))
    return (
        f"{sid}: {kind} "
        f"(positive={positive}, iou_min={iou}, area_max={area}, partial={partial}, "
        f"decoys={decoy_place}, occlusion={occ_enabled})"
    )


def _float_or_zero(v: Any) -> float:
    try:
        return float(v)
    except Exception:
        return 0.0


def resolve_optional_file(raw: str, root: Path, report_dir: Path) -> Path | None:
    raw = (raw or "").strip()
    candidates: list[Path] = []
    if raw:
        p = Path(raw)
        candidates.append(p)
        if not p.is_absolute():
            candidates.append(report_dir / p)
            candidates.append(root / p)
    else:
        candidates.append(report_dir / "find-on-screen-e2e.json")

    # fallback: newest seeded/default benchmark json in report dir
    if report_dir.exists():
        for p in sorted(report_dir.glob("find-on-screen-e2e*.json"), key=lambda x: x.stat().st_mtime, reverse=True):
            candidates.append(p)

    seen: set[Path] = set()
    for c in candidates:
        c = c.resolve()
        if c in seen:
            continue
        seen.add(c)
        if c.exists() and c.is_file():
            return c
    return None


def compute_engine_match_comparison(bench_report_path: Path | None, selected_engine: str) -> dict[str, Any]:
    out: dict[str, Any] = {
        "available": False,
        "engine": selected_engine,
        "bench_report_path": str(bench_report_path) if bench_report_path else "",
    }
    if bench_report_path is None:
        out["error"] = "benchmark report not found"
        return out

    try:
        payload = json.loads(bench_report_path.read_text(encoding="utf-8"))
    except Exception as exc:
        out["error"] = f"failed to parse benchmark report: {type(exc).__name__}: {exc}"
        return out

    summary = payload.get("summary") or {}
    rows = summary.get("by_engine") or payload.get("by_engine") or []
    if not isinstance(rows, list) or not rows:
        out["error"] = "benchmark report missing summary.by_engine rows"
        return out

    parsed: list[dict[str, Any]] = []
    for row in rows:
        if not isinstance(row, dict):
            continue
        engine = str(row.get("engine", "")).strip()
        if not engine:
            continue
        cases = _float_or_zero(row.get("cases", 0))
        ok = _float_or_zero(row.get("query_ok", 0))
        match_rate = (ok / cases * 100.0) if cases > 0 else 0.0
        parsed.append(
            {
                "engine": engine,
                "cases": cases,
                "query_ok": ok,
                "match_rate_pct": match_rate,
            }
        )

    if not parsed:
        out["error"] = "benchmark report contains no valid engine rows"
        return out

    selected_lc = (selected_engine or "").strip().lower()
    selected = next((r for r in parsed if r["engine"].lower() == selected_lc), None)
    if selected is None:
        out["error"] = f"selected engine '{selected_engine}' not present in benchmark report"
        out["engines_present"] = sorted(r["engine"] for r in parsed)
        return out

    others = [r for r in parsed if r["engine"].lower() != selected_lc]
    others_cases = sum(r["cases"] for r in others)
    others_ok = sum(r["query_ok"] for r in others)
    others_match_rate = (others_ok / others_cases * 100.0) if others_cases > 0 else 0.0

    ranked = sorted(parsed, key=lambda r: (r["match_rate_pct"], r["query_ok"]), reverse=True)
    rank = 1
    for i, row in enumerate(ranked, start=1):
        if row["engine"].lower() == selected_lc:
            rank = i
            break

    out.update(
        {
            "available": True,
            "bench_report_path": str(bench_report_path),
            "engine": selected["engine"],
            "engine_match_rate_pct": selected["match_rate_pct"],
            "engine_cases": selected["cases"],
            "engine_query_ok": selected["query_ok"],
            "others_match_rate_pct": others_match_rate,
            "others_cases": others_cases,
            "others_query_ok": others_ok,
            "delta_pct_points": selected["match_rate_pct"] - others_match_rate,
            "engine_rank": rank,
            "engine_count": len(parsed),
        }
    )
    return out


def truthy(raw: str) -> bool:
    return (raw or "").strip().lower() in {"1", "true", "yes", "on"}


def to_docs_site_link(root: Path, target: Path) -> str | None:
    docs_root = (root / "docs").resolve()
    try:
        rel = target.resolve().relative_to(docs_root)
    except Exception:
        return None
    return f"/{rel.as_posix()}"


def to_display_path(root: Path, target: Path | str) -> str:
    path = Path(target).resolve()
    try:
        return path.relative_to(root).as_posix()
    except Exception:
        return str(target)


def parse_resolution(raw: str) -> tuple[int, int]:
    m = re.match(r"^\s*(\d+)\s*[xX]\s*(\d+)\s*$", raw or "")
    if not m:
        raise ValueError(f"invalid resolution string: {raw!r}")
    w = int(m.group(1))
    h = int(m.group(2))
    if w < 64 or h < 64:
        raise ValueError(f"resolution too small: {w}x{h}")
    return w, h


def sanitize_token(raw: str) -> str:
    raw = (raw or "").strip().lower()
    if not raw:
        return "scenario"
    out: list[str] = []
    prev_us = False
    for ch in raw:
        if ch.isalnum():
            out.append(ch)
            prev_us = False
        else:
            if not prev_us:
                out.append("_")
            prev_us = True
    token = "".join(out).strip("_")
    return token or "scenario"


def pick_resolution_id(manifest: dict[str, Any], width: int, height: int) -> str | None:
    groups = manifest.get("resolution_groups") or []
    exact_enabled: list[str] = []
    exact_any: list[str] = []
    for g in groups:
        gid = str(g.get("id", "")).strip()
        if not gid:
            continue
        if int(g.get("width", 0)) == width and int(g.get("height", 0)) == height:
            exact_any.append(gid)
            if bool(g.get("enabled", True)):
                exact_enabled.append(gid)
    if exact_enabled:
        return exact_enabled[0]
    if exact_any:
        return exact_any[0]
    return None


def run_visual_examples(
    manifest: dict[str, Any],
    manifest_path: Path,
    api_dir: Path,
    report_dir: Path,
    visual_dir: Path,
    width: int,
    height: int,
    engine: str,
    benchtime: str,
    count: str,
    tags: str,
    visual_max_attempts: str,
    visual_timeout: str,
    schema_path: str,
    photo_asset: str,
) -> dict[str, Any]:
    result: dict[str, Any] = {
        "enabled": True,
        "resolution": f"{width}x{height}",
        "engine": engine,
        "manifest_for_visuals": "",
        "command": "",
        "log_path": "",
        "exit_code": None,
        "error": "",
        "summaries": [],
    }

    resolution_id = pick_resolution_id(manifest, width, height)
    if not resolution_id:
        result["error"] = f"resolution {width}x{height} not present in manifest resolution_groups"
        return result

    scenario_ids = [str(s.get("id", "")).strip() for s in (manifest.get("scenario_types") or []) if s.get("enabled", True) and str(s.get("id", "")).strip()]
    if not scenario_ids:
        result["error"] = "manifest has no enabled scenario types"
        return result

    vm = copy.deepcopy(manifest)
    vm_matrix = vm.setdefault("matrix", {})
    vm_matrix["resolution_group_ids"] = [resolution_id]
    vm_matrix["scenario_type_ids"] = scenario_ids
    vm_matrix["scenarios_per_resolution"] = len(scenario_ids)
    vm_matrix["enforce_equal_scenario_count"] = True
    if isinstance(vm_matrix.get("per_resolution_overrides"), dict):
        ovr = vm_matrix["per_resolution_overrides"]
        vm_matrix["per_resolution_overrides"] = {resolution_id: ovr[resolution_id]} if resolution_id in ovr else {}

    visual_manifest_path = report_dir / f".tmp-find-bench-strategy-visual-{width}x{height}.json"
    visual_manifest_path.write_text(json.dumps(vm, indent=2), encoding="utf-8")
    result["manifest_for_visuals"] = str(visual_manifest_path)

    visual_dir.mkdir(parents=True, exist_ok=True)
    bench_filter = f"BenchmarkFindOnScreenE2E/engine={engine}/.*_{width}x{height}_"
    cmd = [
        "go",
        "test",
        "./internal/grpcv1",
        "-run",
        "^$",
        "-bench",
        bench_filter,
        "-benchmem",
        "-benchtime",
        benchtime,
        "-count",
        str(count),
    ]
    if tags.strip():
        cmd.extend(["-tags", tags.strip()])
    result["command"] = " ".join(shlex.quote(x) for x in cmd)

    env = os.environ.copy()
    env.update(
        {
            "FIND_BENCH_VISUAL": "1",
            "FIND_BENCH_VISUAL_DIR": str(visual_dir),
            "FIND_BENCH_VISUAL_MAX_ATTEMPTS": str(visual_max_attempts),
            "FIND_BENCH_VISUAL_TIMEOUT": str(visual_timeout),
            "FIND_BENCH_VISUAL_ATTEMPT_OVERLAY": "0",
            "FIND_BENCH_VISUAL_SUMMARY_NATIVE": "1",
            "FIND_BENCH_VISUAL_SUMMARY_SHOW_PATTERN": "1",
            "FIND_BENCH_HIGH_RES": "1",
            "FIND_BENCH_ULTRA_RES": "1",
            "FIND_BENCH_SCENARIO_MANIFEST": str(visual_manifest_path),
        }
    )
    if schema_path.strip():
        env["FIND_BENCH_SCENARIO_SCHEMA"] = schema_path.strip()
    if photo_asset.strip():
        env["FIND_BENCH_PHOTO_ASSET"] = photo_asset.strip()

    proc = subprocess.run(cmd, cwd=api_dir, env=env, stdout=subprocess.PIPE, stderr=subprocess.STDOUT, text=True)
    result["exit_code"] = proc.returncode
    log_path = visual_dir / "go-test.log"
    log_path.write_text(proc.stdout or "", encoding="utf-8")
    result["log_path"] = str(log_path)
    if proc.returncode != 0:
        result["error"] = f"go test exited with code {proc.returncode}"
        return result

    summaries_dir = visual_dir / "summaries"
    files = sorted(summaries_dir.glob("summary-*.png"))
    entries: list[dict[str, str]] = []
    for p in files:
        name = p.stem.replace("summary-", "", 1)
        if f"_{width}x{height}_" not in name:
            continue
        entries.append({"scenario_name": name, "path": str(p)})

    # Prefer one image per scenario type id in declared order.
    picked: list[dict[str, str]] = []
    for sid in scenario_ids:
        token = sanitize_token(sid)
        for e in entries:
            if e["scenario_name"].startswith(f"{token}_{width}x{height}_"):
                picked.append({"scenario_id": sid, "scenario_name": e["scenario_name"], "path": e["path"]})
                break
    if not picked:
        picked = [{"scenario_id": "", "scenario_name": e["scenario_name"], "path": e["path"]} for e in entries]
    result["summaries"] = picked
    return result


root = Path(os.environ["PROJECT_ROOT"]).resolve()
api_dir = (root / "packages" / "api").resolve()
manifest_raw = os.environ["MANIFEST_RAW"]
json_out = Path(os.environ["STRATEGY_JSON_OUT"]).resolve()
md_out = Path(os.environ["STRATEGY_MD_OUT"]).resolve()
strategy_visual_examples = truthy(os.environ.get("STRATEGY_VISUAL_EXAMPLES", "1"))
strategy_visual_res = os.environ.get("STRATEGY_VISUAL_RES", "1280x720")
strategy_visual_engine = (os.environ.get("STRATEGY_VISUAL_ENGINE", "hybrid") or "hybrid").strip()
strategy_visual_benchtime = os.environ.get("STRATEGY_VISUAL_BENCHTIME", "1x")
strategy_visual_count = os.environ.get("STRATEGY_VISUAL_COUNT", "1")
strategy_visual_tags = os.environ.get("STRATEGY_VISUAL_TAGS", "")
strategy_visual_dir = Path(os.environ.get("STRATEGY_VISUAL_DIR", str(json_out.parent / f"strategy-visuals-{strategy_visual_res}"))).resolve()
strategy_visual_max_attempts = os.environ.get("STRATEGY_VISUAL_MAX_ATTEMPTS", "2")
strategy_visual_timeout = os.environ.get("STRATEGY_VISUAL_TIMEOUT", "5s")
strategy_bench_json = os.environ.get("STRATEGY_BENCH_JSON", "")
bench_schema = os.environ.get("BENCH_SCHEMA", "")
bench_photo_asset = os.environ.get("BENCH_PHOTO_ASSET", "")

manifest_path = resolve_manifest(manifest_raw, root)
manifest = json.loads(manifest_path.read_text(encoding="utf-8"))

scenario_types = manifest.get("scenario_types") or []
matrix = manifest.get("matrix") or {}
resolution_groups = manifest.get("resolution_groups") or []
engines = manifest.get("engines") or []

enabled_scenarios = [s for s in scenario_types if s.get("enabled", True)]

style_counts = Counter(str(s.get("style", "")).strip() for s in enabled_scenarios)
kind_counts = Counter(str(s.get("kind", "")).strip() for s in enabled_scenarios)
source_counts = Counter(str((s.get("target") or {}).get("source", "")).strip() for s in enabled_scenarios)
placement_counts = Counter(str((s.get("decoys") or {}).get("placement", "")).strip() for s in enabled_scenarios)

noise_type_counts: Counter[str] = Counter()
for s in enabled_scenarios:
    photometric = s.get("photometric") or {}
    for n in photometric.get("noise") or []:
        ntype = str(n.get("type", "")).strip()
        if ntype:
            noise_type_counts[ntype] += 1

transform_coverage = {
    "scale": 0,
    "rotate": 0,
    "perspective_enabled": 0,
    "skew_x_nonzero": 0,
    "skew_y_nonzero": 0,
}
for s in enabled_scenarios:
    tr = s.get("transforms") or {}
    if tr.get("scale"):
        transform_coverage["scale"] += 1
    if tr.get("rotate"):
        transform_coverage["rotate"] += 1
    persp = tr.get("perspective") or {}
    if bool(persp.get("enabled", False)):
        transform_coverage["perspective_enabled"] += 1
    skew_x = tr.get("skew_x") or {}
    skew_y = tr.get("skew_y") or {}
    if float(skew_x.get("min", 0.0)) != 0.0 or float(skew_x.get("max", 0.0)) != 0.0:
        transform_coverage["skew_x_nonzero"] += 1
    if float(skew_y.get("min", 0.0)) != 0.0 or float(skew_y.get("max", 0.0)) != 0.0:
        transform_coverage["skew_y_nonzero"] += 1

per_scenario: list[dict[str, Any]] = []
for s in enabled_scenarios:
    target = s.get("target") or {}
    background = s.get("background") or {}
    transforms = s.get("transforms") or {}
    photometric = s.get("photometric") or {}
    occlusion = s.get("occlusion") or {}
    decoys = s.get("decoys") or {}
    expected = s.get("expected") or {}
    monitor_selector = s.get("monitor_selector") or {}
    hybrid_policy = s.get("hybrid_policy") or {}

    per_scenario.append(
        {
            "id": s.get("id"),
            "kind": s.get("kind"),
            "style": s.get("style"),
            "goal": scenario_goal_line(s),
            "target": {
                "source": target.get("source"),
                "asset_pool": target.get("asset_pool", []),
                "size_px": target.get("size_px"),
                "rotation_degrees": target.get("rotation_degrees"),
                "aspect_jitter": target.get("aspect_jitter"),
                "subpixel_offset": target.get("subpixel_offset"),
            },
            "background": {
                "continuous_canvas": background.get("continuous_canvas"),
                "clutter_density": background.get("clutter_density"),
                "palette": background.get("palette"),
                "texture_seed": background.get("texture_seed"),
            },
            "transforms": summarize_transforms(transforms),
            "photometric": {
                "brightness": photometric.get("brightness"),
                "contrast": photometric.get("contrast"),
                "gamma": photometric.get("gamma"),
                "blur_sigma": photometric.get("blur_sigma"),
                "jpeg_quality": photometric.get("jpeg_quality"),
                "noise": summarize_noise(photometric),
            },
            "occlusion": {
                "enabled": occlusion.get("enabled"),
                "target_coverage_pct": occlusion.get("target_coverage_pct"),
            },
            "decoys": {
                "enabled": decoys.get("enabled"),
                "count": decoys.get("count"),
                "similarity": decoys.get("similarity"),
                "placement": decoys.get("placement"),
            },
            "expected": expected,
            "monitor_selector": monitor_selector,
            "hybrid_policy": hybrid_policy,
        }
    )

positive_count = sum(1 for s in enabled_scenarios if (s.get("expected") or {}).get("positive", True))
negative_count = len(enabled_scenarios) - positive_count

summary = {
    "generated_at_utc": datetime.now(timezone.utc).isoformat(),
    "manifest_path": str(manifest_path),
    "schema_version": manifest.get("schema_version"),
    "manifest_name": manifest.get("name"),
    "engines": engines,
    "resolution_group_count": len(resolution_groups),
    "resolution_group_ids": [r.get("id") for r in resolution_groups],
    "scenario_type_count": len(scenario_types),
    "enabled_scenario_type_count": len(enabled_scenarios),
    "scenarios_per_resolution": matrix.get("scenarios_per_resolution"),
    "matrix_resolution_ids": matrix.get("resolution_group_ids", []),
    "matrix_scenario_type_ids": matrix.get("scenario_type_ids", []),
    "distribution": {
        "kinds": dict(kind_counts),
        "styles": dict(style_counts),
        "target_sources": dict(source_counts),
        "decoy_placements": dict(placement_counts),
        "noise_types": dict(noise_type_counts),
        "expected_positive": positive_count,
        "expected_negative": negative_count,
    },
    "transform_coverage": transform_coverage,
}
bench_report_path = resolve_optional_file(strategy_bench_json, root, json_out.parent)
summary["engine_match_comparison"] = compute_engine_match_comparison(bench_report_path, strategy_visual_engine)

visual_examples: dict[str, Any] = {"enabled": strategy_visual_examples}
if strategy_visual_examples:
    try:
        vw, vh = parse_resolution(strategy_visual_res)
        visual_examples = run_visual_examples(
            manifest=manifest,
            manifest_path=manifest_path,
            api_dir=api_dir,
            report_dir=json_out.parent,
            visual_dir=strategy_visual_dir,
            width=vw,
            height=vh,
            engine=strategy_visual_engine,
            benchtime=strategy_visual_benchtime,
            count=strategy_visual_count,
            tags=strategy_visual_tags,
            visual_max_attempts=strategy_visual_max_attempts,
            visual_timeout=strategy_visual_timeout,
            schema_path=bench_schema,
            photo_asset=bench_photo_asset,
        )
    except Exception as exc:
        visual_examples = {
            "enabled": True,
            "resolution": strategy_visual_res,
            "engine": strategy_visual_engine,
            "error": f"{type(exc).__name__}: {exc}",
            "summaries": [],
        }

report = {
    "summary": summary,
    "scenarios": per_scenario,
    "visual_examples": visual_examples,
}

json_out.parent.mkdir(parents=True, exist_ok=True)
json_out.write_text(json.dumps(report, indent=2), encoding="utf-8")

def md_dict_rows(title: str, data: dict[str, Any]) -> list[str]:
    lines = [f"### {title}", "", "| Metric | Value |", "|---|---|"]
    for k in sorted(data.keys()):
        lines.append(f"| `{k}` | `{data[k]}` |")
    lines.append("")
    return lines

lines: list[str] = []
lines.append("# FindOnScreen Scenario Strategy Report")
lines.append("")
lines.append(f"- Manifest: `{to_display_path(root, manifest_path)}`")
lines.append(f"- Generated: `{summary['generated_at_utc']}`")
lines.append(f"- Schema version: `{summary['schema_version']}`")
lines.append(f"- Engines: `{', '.join(summary['engines'])}`")
lines.append("")

lines.append("## Strategy Summary Card")
lines.append("")
engine_cmp = summary.get("engine_match_comparison") or {}
if engine_cmp.get("available"):
    lines.append("| Metric | Value |")
    lines.append("|---|---|")
    lines.append(f"| Engine | `{engine_cmp.get('engine')}` |")
    lines.append(f"| Engine match rate | `{float(engine_cmp.get('engine_match_rate_pct', 0.0)):.1f}%` |")
    lines.append(f"| Other engines match rate (weighted) | `{float(engine_cmp.get('others_match_rate_pct', 0.0)):.1f}%` |")
    lines.append(f"| Delta vs others | `{float(engine_cmp.get('delta_pct_points', 0.0)):+.1f} pts` |")
    lines.append(f"| Engine rank by match rate | `{int(engine_cmp.get('engine_rank', 0))}/{int(engine_cmp.get('engine_count', 0))}` |")
    lines.append(f"| Benchmark source | `{to_display_path(root, engine_cmp.get('bench_report_path', ''))}` |")
else:
    lines.append(f"_Match-rate comparison unavailable: {engine_cmp.get('error', 'no benchmark summary found')}._")
    if engine_cmp.get("bench_report_path"):
        lines.append("")
        lines.append(f"- Expected benchmark report path: `{engine_cmp.get('bench_report_path')}`")
lines.append("")

if visual_examples.get("enabled"):
    lines.append("## Visual Examples")
    lines.append("")
    lines.append(f"- Resolution: `{visual_examples.get('resolution', strategy_visual_res)}`")
    lines.append(f"- Engine: `{visual_examples.get('engine', strategy_visual_engine)}`")
    if visual_examples.get("command"):
        lines.append(f"- Command: `{visual_examples.get('command')}`")
    if visual_examples.get("log_path"):
        lines.append(f"- Log: `{to_display_path(root, visual_examples.get('log_path'))}`")
    if visual_examples.get("error"):
        lines.append(f"- Error: `{visual_examples['error']}`")
    lines.append("")
    summaries = visual_examples.get("summaries") or []
    if summaries:
        lines.append("| Scenario | Example |")
        lines.append("|---|---|")
        for item in summaries:
            sid = item.get("scenario_id") or item.get("scenario_name") or "scenario"
            img_path = Path(item.get("path", ""))
            rel_img = to_docs_site_link(root, img_path) or (os.path.relpath(img_path, md_out.parent) if img_path else "")
            lines.append(f"| `{sid}` | ![{sid}]({rel_img}) |")
        lines.append("")
    else:
        lines.append("_No visual examples were generated._")
        lines.append("")

lines.append("## Diversity Summary")
lines.append("")
lines.append("| Metric | Value |")
lines.append("|---|---|")
lines.append(f"| Scenario types (enabled/total) | `{summary['enabled_scenario_type_count']}/{summary['scenario_type_count']}` |")
lines.append(f"| Resolution groups | `{summary['resolution_group_count']}` |")
lines.append(f"| Scenarios per resolution | `{summary['scenarios_per_resolution']}` |")
lines.append(f"| Expected positive | `{summary['distribution']['expected_positive']}` |")
lines.append(f"| Expected negative | `{summary['distribution']['expected_negative']}` |")
lines.append("")

lines.extend(md_dict_rows("Kinds", summary["distribution"]["kinds"]))
lines.extend(md_dict_rows("Styles", summary["distribution"]["styles"]))
lines.extend(md_dict_rows("Target Sources", summary["distribution"]["target_sources"]))
lines.extend(md_dict_rows("Decoy Placements", summary["distribution"]["decoy_placements"]))
lines.extend(md_dict_rows("Noise Types", summary["distribution"]["noise_types"]))
lines.extend(md_dict_rows("Transform Coverage", summary["transform_coverage"]))

lines.append("## Scenario Intent")
lines.append("")
lines.append("| Scenario ID | Kind | Style | Looking For |")
lines.append("|---|---|---|---|")
for s in per_scenario:
    lines.append(f"| `{s['id']}` | `{s['kind']}` | `{s['style']}` | {s['goal']} |")
lines.append("")

lines.append("## Scenario Configuration Details")
lines.append("")
for s in per_scenario:
    lines.append(f"### `{s['id']}`")
    lines.append("")
    lines.append(f"- Kind: `{s['kind']}`")
    lines.append(f"- Style: `{s['style']}`")
    t = s["target"]
    lines.append(
        f"- Target: source=`{t['source']}` size=`{t['size_px']}` rotation=`{t['rotation_degrees']}` "
        f"assets=`{', '.join(t['asset_pool']) if t['asset_pool'] else 'none'}`"
    )
    b = s["background"]
    lines.append(
        f"- Background: palette=`{b['palette']}` clutter=`{b['clutter_density']}` continuous_canvas=`{b['continuous_canvas']}`"
    )
    tr = s["transforms"]
    lines.append(
        f"- Transforms: scale=`{tr['scale']}` rotate=`{tr['rotate']}` "
        f"perspective_enabled=`{tr['perspective_enabled']}` skew_x=`{tr['skew_x']}` skew_y=`{tr['skew_y']}`"
    )
    p = s["photometric"]
    lines.append(
        f"- Photometric: brightness=`{p['brightness']}` contrast=`{p['contrast']}` gamma=`{p['gamma']}` "
        f"blur=`{p['blur_sigma']}` jpeg_quality=`{p['jpeg_quality']}` noise_types=`{', '.join(p['noise']['types']) if p['noise']['types'] else 'none'}`"
    )
    d = s["decoys"]
    lines.append(
        f"- Decoys: enabled=`{d['enabled']}` count=`{d['count']}` similarity=`{d['similarity']}` placement=`{d['placement']}`"
    )
    o = s["occlusion"]
    lines.append(
        f"- Occlusion: enabled=`{o['enabled']}` coverage=`{o['target_coverage_pct']}`"
    )
    e = s["expected"]
    lines.append(
        f"- Expected: positive=`{e.get('positive')}` iou_min=`{e.get('iou_min')}` "
        f"area_ratio_max=`{e.get('area_ratio_max')}` allow_partial=`{e.get('allow_partial')}`"
    )
    if s["monitor_selector"]:
        lines.append(f"- Monitor selector: `{s['monitor_selector']}`")
    if s["hybrid_policy"]:
        lines.append(f"- Hybrid policy: `{s['hybrid_policy']}`")
    lines.append("")

md_out.parent.mkdir(parents=True, exist_ok=True)
md_out.write_text("\n".join(lines), encoding="utf-8")

print(f"[find-bench-strategy] wrote json report: {json_out}")
print(f"[find-bench-strategy] wrote markdown report: {md_out}")
if visual_examples.get("enabled"):
    if visual_examples.get("error"):
        print(f"[find-bench-strategy] visual examples error: {visual_examples['error']}")
    else:
        print(f"[find-bench-strategy] visual examples generated: {len(visual_examples.get('summaries') or [])}")
PY
