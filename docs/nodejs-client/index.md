---
layout: guide
title: Node.js Client
nav_key: nodejs-client
kicker: Primary Guide
lead: Use the Node package when JavaScript is the control surface and you want the runtime, scaffolding, and monitor flow centered around Node.js.
---

## Who This Section Is For

Use this section when JavaScript is your control surface and you want package install, first script, runtime model, and troubleshooting grouped in one place.

## Section Map

<div class="guide-grid">
  <a class="guide-card" href="{{ '/nodejs-client/installation' | relative_url }}">
    <span class="guide-card__eyebrow">Setup</span>
    <span class="guide-card__title">Installation</span>
    <span class="guide-card__body">Install the package, scaffold a project, and optionally install the runtime on PATH.</span>
  </a>
  <a class="guide-card" href="{{ '/nodejs-client/first-script' | relative_url }}">
    <span class="guide-card__eyebrow">Use</span>
    <span class="guide-card__title">First Script</span>
    <span class="guide-card__body">Write the first `Screen` + `Pattern` flow and understand the current wrapper surface.</span>
  </a>
  <a class="guide-card" href="{{ '/nodejs-client/runtime' | relative_url }}">
    <span class="guide-card__eyebrow">Runtime</span>
    <span class="guide-card__title">Runtime</span>
    <span class="guide-card__body">Learn the auto/connect/spawn flow, region scoping, dashboard, and diagnostics.</span>
  </a>
  <a class="guide-card" href="{{ '/nodejs-client/troubleshooting' | relative_url }}">
    <span class="guide-card__eyebrow">Support</span>
    <span class="guide-card__title">Troubleshooting</span>
    <span class="guide-card__body">Resolve binary, permission, and desktop-environment issues in the Node workflow.</span>
  </a>
</div>

## Fast Path

```bash
yarn dlx @sikuligo/sikuli-go init:js-examples
cd sikuli-go-demo
yarn node examples/click.mjs
```

## Deeper References

- [npm Package](https://www.npmjs.com/package/@sikuligo/sikuli-go)
- [Node Package User Flow]({{ '/guides/node-package-user-flow' | relative_url }})
- [Getting Started]({{ '/getting-started/' | relative_url }})
- [Golang API]({{ '/golang-api/' | relative_url }})
- [Client Strategy]({{ '/strategy/client-strategy' | relative_url }})
