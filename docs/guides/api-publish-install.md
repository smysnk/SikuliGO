# API Publish and Install (Windows/Linux)

> Start with [Downloads]({{ '/downloads/' | relative_url }}) if you want the primary install chooser. This page remains the lower-level release and manual-install reference.

This document defines practical ways to publish and install the `sikuli-go` API binary for Windows and Linux.

## Publish Targets

- Linux `amd64`: `sikuli-go-linux-amd64.tar.gz`
- Linux `arm64`: `sikuli-go-linux-arm64.tar.gz`
- Windows `amd64`: `sikuli-go-windows-amd64.zip`

## Build Artifacts

From repo root:

```bash
mkdir -p .release/linux-amd64 .release/linux-arm64 .release/windows-amd64
cd packages/api

GOOS=linux GOARCH=amd64 \
  go build -tags "gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d" \
  -trimpath -ldflags="-s -w" -o ../../.release/linux-amd64/sikuli-go ./cmd/sikuli-go

GOOS=linux GOARCH=arm64 \
  go build -tags "gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d" \
  -trimpath -ldflags="-s -w" -o ../../.release/linux-arm64/sikuli-go ./cmd/sikuli-go

GOOS=windows GOARCH=amd64 \
  go build -tags "gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d" \
  -trimpath -ldflags="-s -w" -o ../../.release/windows-amd64/sikuli-go.exe ./cmd/sikuli-go
```

Package artifacts:

```bash
cd .release
tar -C linux-amd64 -czf sikuli-go-linux-amd64.tar.gz sikuli-go
tar -C linux-arm64 -czf sikuli-go-linux-arm64.tar.gz sikuli-go
cd windows-amd64 && zip -q ../sikuli-go-windows-amd64.zip sikuli-go.exe
```

## Publish to GitHub Releases

```bash
TAG="<release-tag>"
gh release create "$TAG" \
  .release/sikuli-go-linux-amd64.tar.gz \
  .release/sikuli-go-linux-arm64.tar.gz \
  .release/sikuli-go-windows-amd64.zip \
  --repo smysnk/SikuliGO \
  --title "$TAG" \
  --notes "sikuli-go API binaries for Linux/Windows."
```

For existing tags, replace `gh release create` with `gh release upload`.

## Install on Linux

Install from a release tarball:

```bash
VERSION="<release-tag>"
ARCH="amd64" # or arm64
curl -fL "https://github.com/smysnk/SikuliGO/releases/download/${VERSION}/sikuli-go-linux-${ARCH}.tar.gz" \
  -o /tmp/sikuli-go.tar.gz
tar -xzf /tmp/sikuli-go.tar.gz -C /tmp
sudo install -m 0755 /tmp/sikuli-go /usr/local/bin/sikuli-go
```

Verify:

```bash
sikuli-go -listen 127.0.0.1:50051 -admin-listen :8080
```

## Install on Windows (PowerShell)

```powershell
$Version = "<release-tag>"
$Url = "https://github.com/smysnk/SikuliGO/releases/download/$Version/sikuli-go-windows-amd64.zip"
$Zip = "$env:TEMP\\sikuli-go.zip"
$Dest = "$env:LOCALAPPDATA\\Programs\\sikuli-go"

Invoke-WebRequest -Uri $Url -OutFile $Zip
New-Item -ItemType Directory -Force -Path $Dest | Out-Null
Expand-Archive -Path $Zip -DestinationPath $Dest -Force
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";$Dest", "User")
```

Open a new PowerShell and run:

```powershell
sikuli-go.exe -listen 127.0.0.1:50051 -admin-listen :8080
```

## Distribution Options

- Current recommended channel: GitHub Releases artifacts (`.tar.gz`/`.zip`).
- Optional later channels:
  - Windows: Winget + Chocolatey package definitions.
  - Linux: APT/YUM repos or container image distribution.
