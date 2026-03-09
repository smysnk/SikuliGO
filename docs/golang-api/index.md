---
layout: guide
title: Golang API
nav_key: golang-api
kicker: Primary Guide
lead: Use this section when you need the base implementation, the runtime contract, and the generated reference that client wrappers build on top of.
---

## Who This Section Is For

Use this section when you need the base implementation rather than a language wrapper, or when you want the authoritative generated contract behind the Node.js and Python clients.

## Section Map

<div class="guide-grid">
  <a class="guide-card" href="{{ '/golang-api/installation' | relative_url }}">
    <span class="guide-card__eyebrow">Setup</span>
    <span class="guide-card__title">Installation</span>
    <span class="guide-card__body">Prepare the runtime and module path for direct Go usage from the repo or a published runtime.</span>
  </a>
  <a class="guide-card" href="{{ '/golang-api/first-program' | relative_url }}">
    <span class="guide-card__eyebrow">Use</span>
    <span class="guide-card__title">First Program</span>
    <span class="guide-card__body">Connect to the runtime, load a pattern, and execute a live action from Go.</span>
  </a>
  <a class="guide-card" href="{{ '/golang-api/core-types' | relative_url }}">
    <span class="guide-card__eyebrow">Surface</span>
    <span class="guide-card__title">Core Types</span>
    <span class="guide-card__body">Understand `Runtime`, `Screen`, `LiveRegion`, `Pattern`, `Match`, and the related value types.</span>
  </a>
  <a class="guide-card" href="{{ '/golang-api/runtime-and-reference' | relative_url }}">
    <span class="guide-card__eyebrow">Reference</span>
    <span class="guide-card__title">Runtime And Reference</span>
    <span class="guide-card__body">Connect the runtime options to the generated reference and deeper integration guides.</span>
  </a>
</div>

## Fast Path

```bash
make
```

## Deeper References

- [API Reference]({{ '/reference/api/' | relative_url }})
- [Build From Source]({{ '/guides/build-from-source' | relative_url }})
- [OCR Integration]({{ '/guides/ocr-integration' | relative_url }})
- [Input Automation]({{ '/guides/input-automation' | relative_url }})
- [App Control]({{ '/guides/app-control' | relative_url }})
