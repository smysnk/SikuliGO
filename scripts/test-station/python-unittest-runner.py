#!/usr/bin/env python3
from __future__ import annotations

import argparse
import contextlib
import inspect
import io
import json
import os
import sys
import time
import traceback
import unittest
from pathlib import Path


class RecordingResult(unittest.TextTestResult):
    def __init__(self, stream, descriptions, verbosity):
        super().__init__(stream, descriptions, verbosity)
        self.records: list[dict[str, object]] = []
        self._started_at: dict[int, float] = {}

    def startTest(self, test):  # noqa: N802
        self._started_at[id(test)] = time.perf_counter()
        super().startTest(test)

    def addSuccess(self, test):  # noqa: N802
        super().addSuccess(test)
        self.records.append(build_test_record(test, "passed", self._duration_ms(test)))

    def addFailure(self, test, err):  # noqa: N802
        super().addFailure(test, err)
        self.records.append(
            build_test_record(
                test,
                "failed",
                self._duration_ms(test),
                [self._exc_info_to_string(err, test)],
            )
        )

    def addError(self, test, err):  # noqa: N802
        super().addError(test, err)
        self.records.append(
            build_test_record(
                test,
                "failed",
                self._duration_ms(test),
                [self._exc_info_to_string(err, test)],
            )
        )

    def addSkip(self, test, reason):  # noqa: N802
        super().addSkip(test, reason)
        self.records.append(build_test_record(test, "skipped", self._duration_ms(test), [str(reason)]))

    def _duration_ms(self, test) -> int:
        started_at = self._started_at.pop(id(test), None)
        if started_at is None:
            return 0
        return int((time.perf_counter() - started_at) * 1000)


def build_test_record(test, status: str, duration_ms: int, failure_messages: list[str] | None = None) -> dict[str, object]:
    test_id = test.id()
    file_path, line_number = resolve_test_location(test)
    method_name = getattr(test, "_testMethodName", None)
    return {
        "name": method_name or test_id,
        "fullName": test_id,
        "status": status,
        "durationMs": duration_ms,
        "file": file_path,
        "line": line_number,
        "column": None,
        "failureMessages": failure_messages or [],
        "assertions": [],
        "setup": [],
        "mocks": [],
        "rawDetails": {},
        "classificationSource": "adapter",
    }


def resolve_test_location(test) -> tuple[str | None, int | None]:
    method_name = getattr(test, "_testMethodName", None)
    method = getattr(test, method_name, None) if method_name else None
    line_number = None
    if method is not None:
        try:
            _, line_number = inspect.getsourcelines(method)
        except (OSError, TypeError):
            line_number = None

    module = sys.modules.get(test.__class__.__module__)
    file_path = None
    for candidate in (method, test.__class__, module):
        if candidate is None:
            continue
        try:
            file_path = inspect.getsourcefile(candidate) or inspect.getfile(candidate)
        except (OSError, TypeError):
            continue
        if file_path:
            break

    if file_path and not os.path.isabs(file_path):
        file_path = str(Path(file_path).resolve())

    return file_path, line_number


def create_summary(records: list[dict[str, object]]) -> dict[str, int]:
    return {
        "total": len(records),
        "passed": sum(1 for record in records if record["status"] == "passed"),
        "failed": sum(1 for record in records if record["status"] == "failed"),
        "skipped": sum(1 for record in records if record["status"] == "skipped"),
    }


def create_coverage_metric(covered: int, total: int) -> dict[str, object]:
    safe_total = max(0, int(total))
    safe_covered = max(0, min(safe_total, int(covered)))
    pct = 100.0 if safe_total == 0 else round((safe_covered / safe_total) * 100.0, 2)
    return {
        "covered": safe_covered,
        "total": safe_total,
        "pct": pct,
    }


def collect_coverage(package_root: Path) -> tuple[dict[str, object] | None, dict[str, object] | None, list[str]]:
    warnings: list[str] = []
    try:
        import coverage  # type: ignore
    except ImportError:
        warnings.append("Coverage module is not installed; Python coverage was skipped.")
        return None, None, warnings

    cov = coverage.Coverage(source=[str(package_root)])
    cov.start()
    return cov, None, warnings


def finalize_coverage(cov, package_root: Path) -> tuple[dict[str, object] | None, dict[str, object] | None]:
    if cov is None:
        return None, None

    cov.stop()
    cov.save()

    files: list[dict[str, object]] = []
    total_lines = 0
    total_lines_covered = 0

    for measured_file in sorted(cov.get_data().measured_files()):
        analysis = cov.analysis2(measured_file)
        statements = analysis[1]
        missing = set(analysis[3])
        covered_lines = [line for line in statements if line not in missing]
        line_metric = create_coverage_metric(len(covered_lines), len(statements))
        files.append(
            {
                "path": str(Path(measured_file).resolve()),
                "lines": line_metric,
                "statements": line_metric,
                "functions": None,
                "branches": None,
            }
        )
        total_lines += len(statements)
        total_lines_covered += len(covered_lines)

    coverage_summary = {
        "lines": create_coverage_metric(total_lines_covered, total_lines),
        "statements": create_coverage_metric(total_lines_covered, total_lines),
        "functions": None,
        "branches": None,
        "files": files,
    }
    return coverage_summary, {"files": files}


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--package-root", required=True)
    parser.add_argument("--test-dir", required=True)
    parser.add_argument("--pattern", default="test_*.py")
    parser.add_argument("--coverage", action="store_true")
    parser.add_argument("--python-executable")
    args = parser.parse_args()

    package_root = Path(args.package_root).resolve()
    test_dir = Path(args.test_dir).resolve()

    for candidate in (package_root, test_dir):
        if str(candidate) not in sys.path:
            sys.path.insert(0, str(candidate))

    loader = unittest.defaultTestLoader
    suite = loader.discover(start_dir=str(test_dir), pattern=args.pattern)

    coverage_controller = None
    coverage_warnings: list[str] = []
    if args.coverage:
      coverage_controller, _, coverage_warnings = collect_coverage(package_root)

    runner_stream = io.StringIO()
    capture_stdout = io.StringIO()
    capture_stderr = io.StringIO()
    result: RecordingResult

    with contextlib.redirect_stdout(capture_stdout), contextlib.redirect_stderr(capture_stderr):
        runner = unittest.TextTestRunner(
            stream=runner_stream,
            verbosity=2,
            resultclass=RecordingResult,
        )
        result = runner.run(suite)

    coverage_summary, coverage_raw = finalize_coverage(coverage_controller, package_root)
    records = result.records
    summary = create_summary(records)
    status = "failed" if summary["failed"] > 0 or not result.wasSuccessful() else "passed"
    if summary["total"] == 0:
        status = "skipped"

    stdout_text = "\n".join(filter(None, [capture_stdout.getvalue().strip(), runner_stream.getvalue().strip()])).strip()
    stderr_text = capture_stderr.getvalue().strip()

    payload = {
        "status": status,
        "durationMs": sum(int(record["durationMs"]) for record in records),
        "summary": summary,
        "coverage": coverage_summary,
        "tests": records,
        "warnings": coverage_warnings,
        "output": {
            "stdout": stdout_text,
            "stderr": stderr_text,
        },
        "rawArtifacts": [
            {
                "relativePath": "client-python-unittest.log",
                "content": "\n".join(filter(None, [stdout_text, stderr_text])).strip(),
            },
            *(
                [
                    {
                        "relativePath": "client-python-coverage.json",
                        "content": json.dumps(coverage_raw, indent=2),
                    }
                ]
                if coverage_raw is not None
                else []
            ),
        ],
    }

    print(json.dumps(payload))
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except SystemExit:
        raise
    except Exception:  # pragma: no cover - helper failure path
        traceback.print_exc(file=sys.stderr)
        raise SystemExit(1)
