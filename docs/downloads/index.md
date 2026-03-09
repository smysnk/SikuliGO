---
layout: guide
title: Downloads
nav_key: downloads
kicker: Primary Guide
lead: Pick the shortest install path for the runtime, a client package, or a source-based development setup.
---

## Who This Section Is For

Use this page when you are deciding how to get the runtime onto a machine:

- pick a package-driven path for Node.js or Python
- install the runtime binaries directly from a release
- build from source for local development or release work

## Fast Path

| Path | Use When | Command | Next Page |
|---|---|---|---|
| Node.js package | You want the JavaScript client and example scaffolding | `yarn dlx @sikuligo/sikuli-go init:js-examples` | [Node.js Client]({{ '/nodejs-client/' | relative_url }}) |
| Python package | You want the Python client and example scaffolding | `pipx run sikuli-go init:py-examples` | [Python Client]({{ '/python-client/' | relative_url }}) |
| Runtime on PATH | You want `sikuli-go` and `sikuli-go-monitor` available directly | `yarn dlx @sikuligo/sikuli-go install-binary` or `pipx run sikuli-go install-binary` | [Getting Started]({{ '/getting-started/' | relative_url }}) |
| Source build | You are developing in the repo or need generated artifacts | `make` | [Build From Source]({{ '/guides/build-from-source' | relative_url }}) |

## Distribution Channels

<div class="guide-grid">
  <a class="guide-card" href="https://www.npmjs.com/package/@sikuligo/sikuli-go">
    <span class="guide-card__eyebrow">Node.js</span>
    <span class="guide-card__title">npm Package</span>
    <span class="guide-card__body">Use the published package to scaffold `.mjs` examples, bootstrap the runtime, and stay in a Node-first workflow.</span>
  </a>
  <a class="guide-card" href="https://pypi.org/project/sikuli-go/">
    <span class="guide-card__eyebrow">Python</span>
    <span class="guide-card__title">PyPI Package</span>
    <span class="guide-card__body">Use the published package to scaffold Python examples and keep runtime startup close to the script flow.</span>
  </a>
  <a class="guide-card" href="https://github.com/smysnk/SikuliGO/releases">
    <span class="guide-card__eyebrow">Runtime</span>
    <span class="guide-card__title">GitHub Releases</span>
    <span class="guide-card__body">Fetch the published runtime artifacts directly when you only need the binaries.</span>
  </a>
  <a class="guide-card" href="{{ '/guides/build-from-source' | relative_url }}">
    <span class="guide-card__eyebrow">Development</span>
    <span class="guide-card__title">Build From Source</span>
    <span class="guide-card__body">Use the monorepo build and verification flow when you need local edits, generated stubs, or release preparation.</span>
  </a>
</div>

## Common Tasks

### Install The Runtime On PATH

```bash
yarn dlx @sikuligo/sikuli-go install-binary
# or
pipx run sikuli-go install-binary
```

Both flows can add `~/.local/bin` to your shell config. Reload the shell after install if the command is not found immediately.

### Install A Release Binary On Linux

```bash
VERSION="<release-tag>"
ARCH="amd64"
curl -fL "https://github.com/smysnk/SikuliGO/releases/download/${VERSION}/sikuli-go-linux-${ARCH}.tar.gz" \
  -o /tmp/sikuli-go.tar.gz
tar -xzf /tmp/sikuli-go.tar.gz -C /tmp
sudo install -m 0755 /tmp/sikuli-go /usr/local/bin/sikuli-go
```

### Build In The Repository

```bash
make
```

## Troubleshooting And Next Steps

- If `sikuli-go` is not found after `install-binary`, reload `~/.zshrc` or `~/.bash_profile`.
- If you need OS-specific release artifact details, use [API Publish and Install]({{ '/guides/api-publish-install' | relative_url }}).
- If you already know your language path, continue to [Node.js Client]({{ '/nodejs-client/' | relative_url }}) or [Python Client]({{ '/python-client/' | relative_url }}).
- If you want the shortest first-run flow, continue to [Getting Started]({{ '/getting-started/' | relative_url }}).

## Deeper References

- [API Publish and Install]({{ '/guides/api-publish-install' | relative_url }})
- [Build From Source]({{ '/guides/build-from-source' | relative_url }})
- [Node.js Client]({{ '/nodejs-client/' | relative_url }})
- [Python Client]({{ '/python-client/' | relative_url }})
