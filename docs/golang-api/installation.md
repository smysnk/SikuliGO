---
layout: guide
title: Golang API Installation
nav_key: golang-api
kicker: Golang API
lead: Prepare the runtime and API module for direct Go usage, either from published binaries or from the repository build flow.
---

## Who This Page Is For

Use this page when you want to work with the base implementation directly instead of starting from the Node.js or Python wrappers.

## Runtime Setup Options

### Build From Source

```bash
make
```

Or build the runtime binaries explicitly:

```bash
cd packages/api
go build -tags "gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d" \
  -trimpath -ldflags="-s -w" -o ../../sikuli-go ./cmd/sikuli-go
go build -tags "gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d" \
  -trimpath -ldflags="-s -w" -o ../../sikuli-go-monitor ./cmd/sikuli-go-monitor
```

### Install A Published Runtime

Use [Downloads]({{ '/downloads/' | relative_url }}) and [API Publish and Install]({{ '/guides/api-publish-install' | relative_url }}) if you want a published binary instead of a local build.

## Module Path

The Go package surfaces in this repository live under:

```go
github.com/smysnk/sikuligo/pkg/sikuli
```

## Verify The Runtime

```bash
sikuli-go -listen 127.0.0.1:50051 -admin-listen :8080
```

Then use `sikuli.NewRuntime("127.0.0.1:50051")` from your Go program.

## Next Pages

- [Golang API: First Program]({{ '/golang-api/first-program' | relative_url }})
- [Golang API: Core Types]({{ '/golang-api/core-types' | relative_url }})
- [Golang API: Runtime And Reference]({{ '/golang-api/runtime-and-reference' | relative_url }})
