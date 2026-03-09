---
layout: guide
title: Contribution Development Setup
nav_key: contribution
kicker: Contribution
lead: Set up the repository, build the runtime, run the local verification flows, and preview the docs before changing public behavior.
---

## Prerequisites

- Go `1.24+`
- Node.js `20+`
- Yarn `4+`
- Python `3.10+`
- `protoc`

## Install Workspace Dependencies

```bash
yarn install
```

## Build The Project

```bash
make
```

For explicit runtime builds:

```bash
cd packages/api
go build -tags "gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d" \
  -trimpath -ldflags="-s -w" -o ../../sikuli-go ./cmd/sikuli-go
go build -tags "gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d" \
  -trimpath -ldflags="-s -w" -o ../../sikuli-go-monitor ./cmd/sikuli-go-monitor
```

## Verify The Repo

Quick publish-oriented verification:

```bash
make test-publish
```

Broader integration verification:

```bash
make test-integration
```

## Docs And Generated Artifacts

Run the docs governance check:

```bash
./scripts/check-docs-governance.sh
```

Run the internal docs link check:

```bash
./scripts/check-docs-links.sh
```

Regenerate API docs when needed:

```bash
./scripts/generate-api-docs.sh
```

Regenerate parity docs when needed:

```bash
./scripts/generate-parity-docs.sh
```

Preview the docs locally:

```bash
make docs
```

The local preview server serves staged markdown through the guide shell when Jekyll is not available, so the browser view stays close to the published Pages layout.

For a non-interactive local check without opening a browser or keeping the server running:

```bash
DOCS_LOCAL_OPEN_BROWSER=0 DOCS_LOCAL_SERVE=0 make docs
```

## Key Workflow References

- [Documentation Workflow]({{ '/contribution/docs-workflow' | relative_url }})
- [Build From Source]({{ '/guides/build-from-source' | relative_url }})
- [Guide-Format Documentation Plan]({{ '/strategy/guide-format-documentation-plan' | relative_url }})
- [Go Test Workflow](https://github.com/smysnk/SikuliGO/blob/master/.github/workflows/go-test.yml)
