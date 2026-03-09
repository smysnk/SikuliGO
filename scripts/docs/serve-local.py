#!/usr/bin/env python3
"""Serve staged docs with markdown-aware rendering and pretty URLs."""

from __future__ import annotations

import argparse
import html
import http.server
import mimetypes
import re
import socketserver
import sys
from pathlib import Path
from typing import Any
from urllib.parse import unquote, urlparse

import markdown
import yaml


FRONT_MATTER_RE = re.compile(r"\A---\s*\n(.*?)\n---\s*\n?", re.DOTALL)
RELATIVE_URL_RE = re.compile(r"""\{\{\s*(['"])(.*?)\1\s*\|\s*relative_url\s*\}\}""")
LINK_TEXT_RE = re.compile(r"\[([^\]]+)\]\([^)]+\)")
TAG_RE = re.compile(r"<[^>]+>")
WHITESPACE_RE = re.compile(r"\s+")
HEADING_RE = re.compile(r"^\s*#\s+(.+?)\s*$", re.MULTILINE)


def strip_tags(value: str) -> str:
    return WHITESPACE_RE.sub(" ", TAG_RE.sub("", value)).strip()


def strip_inline_markdown(value: str) -> str:
    value = LINK_TEXT_RE.sub(r"\1", value)
    value = value.replace("`", "")
    value = value.replace("*", "")
    value = value.replace("_", "")
    return WHITESPACE_RE.sub(" ", value).strip()


def resolve_liquid_relative_urls(text: str) -> str:
    return RELATIVE_URL_RE.sub(lambda match: match.group(2), text)


def parse_front_matter(text: str) -> tuple[dict[str, Any], str]:
    match = FRONT_MATTER_RE.match(text)
    if not match:
        return {}, text
    raw = match.group(1)
    data = yaml.safe_load(raw) or {}
    if not isinstance(data, dict):
        raise ValueError("front matter must decode to a mapping")
    return data, text[match.end() :]


def extract_title(content: str) -> tuple[str | None, str]:
    match = HEADING_RE.search(content)
    if not match:
        return None, content
    title = strip_inline_markdown(match.group(1))
    start = match.start()
    end = match.end()
    while end < len(content) and content[end] == "\n":
        end += 1
    return title, content[:start] + content[end:]


def pretty_name(path: Path) -> str:
    stem = path.stem if path.name != "index.md" else path.parent.name or "Documentation"
    words = stem.replace("-", " ").replace("_", " ").split()
    return " ".join(word.capitalize() for word in words) or "Documentation"


def page_url_from_rel(rel_path: Path) -> str:
    rel_posix = rel_path.as_posix().lstrip("./")
    if rel_posix in {"index.md", "index.html"}:
        return "/"
    if rel_path.name in {"index.md", "index.html"}:
        parent = rel_path.parent.as_posix().strip("/")
        return f"/{parent}/" if parent else "/"
    if rel_path.suffix in {".md", ".html"}:
        base = rel_path.with_suffix("").as_posix().strip("/")
        return f"/{base}/" if base else "/"
    return f"/{rel_posix.lstrip('/')}"


def is_active_nav(item: dict[str, Any], page_url: str, nav_key: str | None) -> bool:
    item_key = item.get("key")
    item_url = str(item.get("url", "/"))
    if nav_key and item_key and nav_key == item_key:
        return True
    if item_url == "/":
        return page_url == "/"
    normalized_item = item_url.rstrip("/") + "/"
    return page_url == normalized_item or page_url.startswith(normalized_item)


def render_sidebar(
    navigation: dict[str, list[dict[str, Any]]],
    site_description: str,
    page_url: str,
    nav_key: str | None,
) -> str:
    parts = [
        '<div class="guide-brand">',
        '  <a class="guide-brand__link" href="/">',
        '    <img class="guide-brand__logo" src="/images/logo.png" alt="sikuli-go logo">',
        '    <span class="guide-brand__text">',
        "      <strong>sikuli-go</strong>",
        "      <span>Documentation</span>",
        "    </span>",
        "  </a>",
        f'  <p class="guide-brand__summary">{html.escape(site_description)}</p>',
        "</div>",
        '<nav class="guide-nav" aria-label="Documentation navigation">',
    ]

    groups = [
        ("Primary Guide", "primary", False),
        ("Secondary", "secondary", False),
        ("Project", "project", True),
    ]
    for label, key, compact in groups:
        items = navigation.get(key, [])
        if not items:
            continue
        list_class = "guide-nav__list guide-nav__list--compact" if compact else "guide-nav__list"
        parts.append('  <div class="guide-nav__group">')
        parts.append(f'    <p class="guide-nav__label">{html.escape(label)}</p>')
        parts.append(f'    <ul class="{list_class}">')
        for item in items:
            title = str(item.get("title", "Untitled"))
            summary = item.get("summary")
            url = str(item.get("url", "/"))
            active_class = " is-active" if is_active_nav(item, page_url, nav_key) else ""
            parts.append(f'      <li class="guide-nav__item{active_class}">')
            parts.append(f'        <a class="guide-nav__link" href="{html.escape(url)}">')
            parts.append(f'          <span class="guide-nav__title">{html.escape(title)}</span>')
            if summary and not compact:
                parts.append(f'          <span class="guide-nav__summary">{html.escape(str(summary))}</span>')
            parts.append("        </a>")
            parts.append("      </li>")
        parts.append("    </ul>")
        parts.append("  </div>")
    parts.append("</nav>")
    return "\n".join(parts)


def render_markdown_page(
    source_path: Path,
    rel_path: Path,
    site_title: str,
    site_description: str,
    navigation: dict[str, list[dict[str, Any]]],
) -> bytes:
    raw_text = source_path.read_text(encoding="utf-8")
    metadata, content = parse_front_matter(raw_text)
    content = resolve_liquid_relative_urls(content)

    title = metadata.get("title")
    if title is not None:
        title = str(title)
    else:
        title, content = extract_title(content)
        title = title or pretty_name(rel_path)

    kicker = metadata.get("kicker")
    lead = metadata.get("lead")
    nav_key = metadata.get("nav_key")
    page_url = page_url_from_rel(rel_path)

    renderer = markdown.Markdown(
        extensions=["extra", "sane_lists"],
        output_format="html5",
    )
    content_html = renderer.convert(content)
    page_title = f"{title} | {site_title}" if title else site_title
    description = strip_tags(str(lead)) if lead else site_description
    sidebar_html = render_sidebar(navigation, site_description, page_url, str(nav_key) if nav_key else None)

    kicker_html = f'<p class="guide-kicker">{html.escape(str(kicker))}</p>' if kicker else ""
    lead_html = f'<p class="guide-lead">{html.escape(str(lead))}</p>' if lead else ""

    document = f"""<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{html.escape(page_title)}</title>
    <meta name="description" content="{html.escape(description)}">
    <link rel="stylesheet" href="/assets/guide.css">
  </head>
  <body>
    <div class="guide-shell">
      <aside class="guide-sidebar">
{sidebar_html}
      </aside>
      <main class="guide-main">
        <header class="guide-page-header">
          {kicker_html}
          <h1 class="guide-title">{html.escape(title)}</h1>
          {lead_html}
        </header>
        <article class="guide-content">
{content_html}
        </article>
      </main>
    </div>
  </body>
</html>
"""
    return document.encode("utf-8")


class ThreadingTCPServer(socketserver.ThreadingMixIn, socketserver.TCPServer):
    allow_reuse_address = True
    daemon_threads = True


class DocsPreviewHandler(http.server.BaseHTTPRequestHandler):
    site_root = Path(".")
    source_root = Path(".")
    site_title = "Documentation"
    site_description = ""
    navigation: dict[str, list[dict[str, Any]]] = {}

    def do_GET(self) -> None:  # noqa: N802
        self.serve(send_body=True)

    def do_HEAD(self) -> None:  # noqa: N802
        self.serve(send_body=False)

    def serve(self, send_body: bool) -> None:
        parsed = urlparse(self.path)
        request_path = unquote(parsed.path)
        resolved = self.resolve_request(request_path)
        if resolved is None:
            self.send_error(404, "File not found")
            return

        kind, source_path, rel_path = resolved
        if kind == "markdown":
            try:
                body = render_markdown_page(
                    source_path=source_path,
                    rel_path=rel_path,
                    site_title=self.site_title,
                    site_description=self.site_description,
                    navigation=self.navigation,
                )
            except Exception as exc:  # pragma: no cover - defensive path
                self.send_error(500, f"Failed to render markdown: {exc}")
                return
            self.send_response(200)
            self.send_header("Content-Type", "text/html; charset=utf-8")
            self.send_header("Content-Length", str(len(body)))
            self.end_headers()
            if send_body:
                self.wfile.write(body)
            return

        content_type = self.guess_type(source_path)
        try:
            data = source_path.read_bytes()
        except OSError as exc:
            self.send_error(404, f"Failed to read file: {exc}")
            return
        self.send_response(200)
        self.send_header("Content-Type", content_type)
        self.send_header("Content-Length", str(len(data)))
        self.end_headers()
        if send_body:
            self.wfile.write(data)

    def guess_type(self, path: Path) -> str:
        guessed, _ = mimetypes.guess_type(str(path))
        return guessed or "application/octet-stream"

    def resolve_request(self, request_path: str) -> tuple[str, Path, Path] | None:
        normalized = request_path if request_path else "/"
        if not normalized.startswith("/"):
            normalized = f"/{normalized}"
        rel = Path(normalized.lstrip("/"))

        site_candidates = self.candidate_paths(rel)
        for candidate in site_candidates:
            path = self.site_root / candidate
            if path.is_file():
                kind = "markdown" if path.suffix == ".md" else "static"
                return kind, path, candidate

        if self.source_root != self.site_root:
            for candidate in self.candidate_paths(rel):
                path = self.source_root / candidate
                if path.is_file():
                    kind = "markdown" if path.suffix == ".md" else "static"
                    return kind, path, candidate
        return None

    @staticmethod
    def candidate_paths(rel: Path) -> list[Path]:
        rel_parts = [part for part in rel.parts if part not in {"", ".", ".."}]
        safe_rel = Path(*rel_parts) if rel_parts else Path()
        if safe_rel == Path():
            return [Path("index.html"), Path("index.md")]

        suffix = safe_rel.suffix.lower()
        candidates: list[Path] = []
        if suffix:
            candidates.append(safe_rel)
            if suffix == ".md":
                candidates.append(safe_rel.with_suffix(".html"))
            elif suffix == ".html":
                candidates.append(safe_rel.with_suffix(".md"))
            return candidates

        candidates.extend(
            [
                safe_rel / "index.html",
                safe_rel / "index.md",
                safe_rel.with_suffix(".html"),
                safe_rel.with_suffix(".md"),
            ]
        )
        return candidates

    def log_message(self, fmt: str, *args: object) -> None:
        sys.stderr.write(f"[docs-preview] {self.address_string()} - {fmt % args}\n")


def load_yaml_mapping(path: Path) -> dict[str, Any]:
    if not path.is_file():
        return {}
    data = yaml.safe_load(path.read_text(encoding="utf-8")) or {}
    if not isinstance(data, dict):
        raise ValueError(f"{path} must decode to a mapping")
    return data


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--host", default="127.0.0.1")
    parser.add_argument("--port", type=int, default=4000)
    parser.add_argument("--site-root", required=True)
    parser.add_argument("--source-root", required=True)
    args = parser.parse_args()

    site_root = Path(args.site_root).resolve()
    source_root = Path(args.source_root).resolve()
    config = load_yaml_mapping(source_root / "_config.yml")
    navigation = load_yaml_mapping(source_root / "_data" / "navigation.yml")

    handler_cls = type(
        "ConfiguredDocsPreviewHandler",
        (DocsPreviewHandler,),
        {
            "site_root": site_root,
            "source_root": source_root,
            "site_title": str(config.get("title", "Documentation")),
            "site_description": str(config.get("description", "")),
            "navigation": navigation,
        },
    )

    with ThreadingTCPServer((args.host, args.port), handler_cls) as server:
        print(f"[docs-preview] serving site_root={site_root}")
        if source_root != site_root:
            print(f"[docs-preview] serving source_root={source_root}")
        print(f"[docs-preview] url=http://{args.host}:{args.port}")
        server.serve_forever()

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
