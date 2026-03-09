#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
DOCS_DIR="${ROOT_DIR}/docs"

python3 - "${DOCS_DIR}" <<'PY'
from __future__ import annotations

import re
import sys
from pathlib import Path
from urllib.parse import urljoin, urlsplit


docs_dir = Path(sys.argv[1]).resolve()

TEXT_SUFFIXES = {".md", ".html", ".yml", ".yaml"}
PAGE_SUFFIXES = {".md", ".html"}
EXTERNAL_PREFIXES = ("http://", "https://", "mailto:", "tel:", "data:", "javascript:")
LIQUID_RELATIVE_URL_RE = re.compile(r"""\{\{\s*['"]([^'"]+)['"]\s*\|\s*relative_url\s*\}\}""")
MARKDOWN_LINK_RE = re.compile(r"""!?\[[^\]]*\]\(([^)]+)\)""")
HTML_ATTR_RE = re.compile(r"""(?:href|src)=["']([^"']+)["']""")
YAML_URL_RE = re.compile(r"""^\s*url:\s*['"]?([^"'#\s]+)['"]?\s*(?:#.*)?$""")


def is_private_path(rel_path: Path) -> bool:
    return any(part.startswith("_") for part in rel_path.parts)


def to_posix_path(path: Path) -> str:
    raw = path.as_posix().strip("/")
    return f"/{raw}" if raw else "/"


def page_url_for(rel_path: Path) -> str | None:
    if rel_path.suffix not in PAGE_SUFFIXES:
        return None
    if is_private_path(rel_path):
        return None
    stem_path = rel_path.with_suffix("")
    if rel_path.name in {"index.md", "index.html"}:
        parent = stem_path.parent.as_posix().strip(".")
        parent = parent.strip("/")
        return f"/{parent}/" if parent else "/"
    return f"/{stem_path.as_posix().strip('/')}/"


published_urls: set[str] = set()
source_page_urls: dict[Path, str] = {}


def add_published_url(url: str) -> None:
    published_urls.add(url)
    if url != "/" and url.endswith("/"):
        published_urls.add(url[:-1])


for path in docs_dir.rglob("*"):
    if not path.is_file():
        continue
    rel_path = path.relative_to(docs_dir)
    page_url = page_url_for(rel_path)
    if page_url is not None:
        source_page_urls[path] = page_url
        add_published_url(page_url)
        continue
    if not is_private_path(rel_path):
        add_published_url(to_posix_path(rel_path))


def clean_markdown_target(raw: str) -> str:
    target = raw.strip()
    if target.startswith("<") and target.endswith(">"):
        target = target[1:-1].strip()
    if " " in target and not target.startswith("{{"):
        target = target.split()[0]
    return target


def normalize_target(target: str, base_url: str) -> str | None:
    target = target.strip()
    if not target or target.startswith("#") or target.startswith(EXTERNAL_PREFIXES):
        return None
    if "{{" in target or "{%" in target:
        return None
    parsed = urlsplit(target)
    raw_path = parsed.path.strip()
    if not raw_path:
        return None
    if raw_path.startswith("/"):
        resolved = raw_path
    else:
        resolved = urljoin(base_url, raw_path)
    resolved = re.sub(r"/{2,}", "/", resolved)
    if not resolved.startswith("/"):
        resolved = f"/{resolved}"
    return resolved


def published_candidate_for_source_path(path: str) -> str | None:
    if not path.endswith((".md", ".html")):
        return None
    rel_path = Path(path.lstrip("/"))
    page_url = page_url_for(rel_path)
    return page_url


def is_known_target(path: str) -> bool:
    if path in published_urls:
        return True
    if path != "/" and path.endswith("/") and path[:-1] in published_urls:
        return True
    if path != "/" and not path.endswith("/") and f"{path}/" in published_urls:
        return True
    source_candidate = published_candidate_for_source_path(path)
    if source_candidate and source_candidate in published_urls:
        return True
    return False


errors: list[str] = []

for path in docs_dir.rglob("*"):
    if not path.is_file() or path.suffix not in TEXT_SUFFIXES:
        continue

    rel_path = path.relative_to(docs_dir)
    base_url = source_page_urls.get(path, "/")
    try:
        text = path.read_text(encoding="utf-8")
    except UnicodeDecodeError:
        continue

    for lineno, line in enumerate(text.splitlines(), start=1):
        targets: list[str] = []
        targets.extend(match.group(1) for match in LIQUID_RELATIVE_URL_RE.finditer(line))
        targets.extend(clean_markdown_target(match.group(1)) for match in MARKDOWN_LINK_RE.finditer(line))
        targets.extend(match.group(1).strip() for match in HTML_ATTR_RE.finditer(line))
        yaml_match = YAML_URL_RE.match(line)
        if yaml_match:
            targets.append(yaml_match.group(1).strip())

        for target in targets:
            normalized = normalize_target(target, base_url)
            if normalized is None:
                continue
            if not is_known_target(normalized):
                errors.append(
                    f"{rel_path}:{lineno}: unresolved internal link: {target} -> {normalized}"
                )

if errors:
    print("Docs link check failed.", file=sys.stderr)
    for error in errors:
        print(error, file=sys.stderr)
    sys.exit(1)

print("Docs link check passed.")
PY
