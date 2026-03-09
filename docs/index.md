---
layout: guide
title: Documentation
nav_key: home
kicker: Public Guide
lead: Task-first documentation for downloads, first-run setup, the Node.js client, the Python client, and the base Go API.
---

<div class="guide-grid">
  <a class="guide-card" href="{{ '/getting-started/' | relative_url }}">
    <span class="guide-card__eyebrow">Start Here</span>
    <span class="guide-card__title">Getting Started</span>
    <span class="guide-card__body">Choose Node.js, Python, or Go and get to a first working script quickly.</span>
  </a>
  <a class="guide-card" href="{{ '/downloads/' | relative_url }}">
    <span class="guide-card__eyebrow">Install</span>
    <span class="guide-card__title">Downloads</span>
    <span class="guide-card__body">Pick the right package, binary, or source path for your environment.</span>
  </a>
  <a class="guide-card" href="{{ '/golang-api/' | relative_url }}">
    <span class="guide-card__eyebrow">Base Runtime</span>
    <span class="guide-card__title">Golang API</span>
    <span class="guide-card__body">Use the core implementation and generated API reference as the base contract.</span>
  </a>
</div>

## Choose Your Path

<div class="guide-grid">
  <a class="guide-card" href="{{ '/nodejs-client/' | relative_url }}">
    <span class="guide-card__eyebrow">Client Guide</span>
    <span class="guide-card__title">Node.js Client</span>
    <span class="guide-card__body">Package install flow, scaffolded examples, dashboard startup, and current runtime references.</span>
  </a>
  <a class="guide-card" href="{{ '/python-client/' | relative_url }}">
    <span class="guide-card__eyebrow">Client Guide</span>
    <span class="guide-card__title">Python Client</span>
    <span class="guide-card__body">pipx-based setup, example scaffolding, runtime behavior, and current Python-facing references.</span>
  </a>
  <a class="guide-card" href="{{ '/getting-help/' | relative_url }}">
    <span class="guide-card__eyebrow">Support</span>
    <span class="guide-card__title">Getting Help</span>
    <span class="guide-card__body">Use the issue-reporting checklist, troubleshooting references, and support entry points.</span>
  </a>
</div>

<div class="guide-callout">
  <strong>Guide structure</strong>
  Start with the primary sections above. Use Guides, Reference, Strategy, and Benchmarks when you need deeper implementation detail, generated API output, or contributor planning material.
</div>

## What Lives In Each Layer

- `Downloads` is the install chooser for npm, PyPI, GitHub releases, and source builds.
- `Getting Started` is the shortest path to the first working automation flow.
- `Node.js Client`, `Python Client`, and `Golang API` are the language-specific landing areas for the public surface.
- `Getting Help`, `Contribution`, and `License` hold operational support and project metadata.
- `Reference`, `Guides`, `Strategy`, and `Benchmarks` remain the deeper supporting layers.

## Deeper Layers

<div class="guide-grid">
  <a class="guide-card guide-card--subtle" href="{{ '/reference/' | relative_url }}">
    <span class="guide-card__eyebrow">Reference</span>
    <span class="guide-card__title">API And Parity Material</span>
    <span class="guide-card__body">Generated package docs and parity references for detailed contract work.</span>
  </a>
  <a class="guide-card guide-card--subtle" href="{{ '/guides/' | relative_url }}">
    <span class="guide-card__eyebrow">Guides</span>
    <span class="guide-card__title">Implementation Guides</span>
    <span class="guide-card__body">Build, OCR, input, app control, and publishing guides that remain valid during the rewrite.</span>
  </a>
  <a class="guide-card guide-card--subtle" href="{{ '/strategy/' | relative_url }}">
    <span class="guide-card__eyebrow">Strategy</span>
    <span class="guide-card__title">Plans And Ownership</span>
    <span class="guide-card__body">Docs plan, source ownership, and architecture strategy for contributors and maintainers.</span>
  </a>
  <a class="guide-card guide-card--subtle" href="{{ '/bench/' | relative_url }}">
    <span class="guide-card__eyebrow">Benchmarks</span>
    <span class="guide-card__title">Published Performance Evidence</span>
    <span class="guide-card__body">Scenario definitions, benchmark reports, and generated visual artifacts.</span>
  </a>
</div>

## Historical References

Use these external references for project background and upstream history:

- [SikuliX Official Site](https://sikulix.github.io/)
- [Wikipedia](https://de.wikipedia.org/wiki/Sikuli_(Software))
- [Original Sikuli Github](https://github.com/sikuli/sikuli)
- [Sikuli Framework](https://github.com/smysnk/sikuli-framework)

## Generated Reference Note

API reference pages are generated from source using `go doc` and validated with `./scripts/check-api-docs.sh`.
