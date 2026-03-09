---
layout: guide
title: Python Client
nav_key: python-client
kicker: Primary Guide
lead: Use the Python package when you want the automation flow, scaffolded examples, and runtime startup centered on Python tooling.
---

## Who This Section Is For

Use this section when Python is your control surface and you want package install, first script, runtime model, and troubleshooting grouped in one place.

## Section Map

<div class="guide-grid">
  <a class="guide-card" href="{{ '/python-client/installation' | relative_url }}">
    <span class="guide-card__eyebrow">Setup</span>
    <span class="guide-card__title">Installation</span>
    <span class="guide-card__body">Install the package, scaffold a project, and optionally install the runtime on PATH.</span>
  </a>
  <a class="guide-card" href="{{ '/python-client/first-script' | relative_url }}">
    <span class="guide-card__eyebrow">Use</span>
    <span class="guide-card__title">First Script</span>
    <span class="guide-card__body">Write the first `Screen` + `Pattern` flow and understand the current wrapper surface.</span>
  </a>
  <a class="guide-card" href="{{ '/python-client/runtime' | relative_url }}">
    <span class="guide-card__eyebrow">Runtime</span>
    <span class="guide-card__title">Runtime</span>
    <span class="guide-card__body">Learn the auto/connect/spawn flow, region scoping, and the environment inputs used by the client.</span>
  </a>
  <a class="guide-card" href="{{ '/python-client/troubleshooting' | relative_url }}">
    <span class="guide-card__eyebrow">Support</span>
    <span class="guide-card__title">Troubleshooting</span>
    <span class="guide-card__body">Resolve package-name, binary-resolution, and desktop-environment issues in the Python workflow.</span>
  </a>
</div>

## Fast Path

```bash
pipx run sikuli-go init:py-examples
cd sikuli-go-demo
python3 examples/click.py
```

## Deeper References

- [PyPI Package](https://pypi.org/project/sikuli-go/)
- [Getting Started]({{ '/getting-started/' | relative_url }})
- [Golang API]({{ '/golang-api/' | relative_url }})
- [Client Strategy]({{ '/strategy/client-strategy' | relative_url }})
