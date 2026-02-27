# API Publish and Install (Windows/Linux)

This document defines practical ways to publish and install the `sikuligo` API binary for Windows and Linux.

## Publish Targets

- Linux `amd64`: `sikuligo-linux-amd64.tar.gz`
- Linux `arm64`: `sikuligo-linux-arm64.tar.gz`
- Windows `amd64`: `sikuligo-windows-amd64.zip`

## Build Artifacts

From repo root:

```bash
mkdir -p .release/linux-amd64 .release/linux-arm64 .release/windows-amd64
cd packages/api

GOOS=linux GOARCH=amd64 \
  go build -tags "gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d" \
  -trimpath -ldflags="-s -w" -o ../../.release/linux-amd64/sikuligo ./cmd/sikuligrpc

GOOS=linux GOARCH=arm64 \
  go build -tags "gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d" \
  -trimpath -ldflags="-s -w" -o ../../.release/linux-arm64/sikuligo ./cmd/sikuligrpc

GOOS=windows GOARCH=amd64 \
  go build -tags "gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d" \
  -trimpath -ldflags="-s -w" -o ../../.release/windows-amd64/sikuligo.exe ./cmd/sikuligrpc
```

Package artifacts:

```bash
cd .release
tar -C linux-amd64 -czf sikuligo-linux-amd64.tar.gz sikuligo
tar -C linux-arm64 -czf sikuligo-linux-arm64.tar.gz sikuligo
cd windows-amd64 && zip -q ../sikuligo-windows-amd64.zip sikuligo.exe
```

## Publish to GitHub Releases

```bash
TAG="v0.1.0"
gh release create "$TAG" \
  .release/sikuligo-linux-amd64.tar.gz \
  .release/sikuligo-linux-arm64.tar.gz \
  .release/sikuligo-windows-amd64.zip \
  --repo smysnk/SikuliGO \
  --title "$TAG" \
  --notes "SikuliGO API binaries for Linux/Windows."
```

For existing tags, replace `gh release create` with `gh release upload`.

## Install on Linux

Install from a release tarball:

```bash
VERSION="v0.1.0"
ARCH="amd64" # or arm64
curl -fL "https://github.com/smysnk/SikuliGO/releases/download/${VERSION}/sikuligo-linux-${ARCH}.tar.gz" \
  -o /tmp/sikuligo.tar.gz
tar -xzf /tmp/sikuligo.tar.gz -C /tmp
sudo install -m 0755 /tmp/sikuligo /usr/local/bin/sikuligo
```

Verify:

```bash
sikuligo -listen 127.0.0.1:50051 -admin-listen :8080
```

## Install on Windows (PowerShell)

```powershell
$Version = "v0.1.0"
$Url = "https://github.com/smysnk/SikuliGO/releases/download/$Version/sikuligo-windows-amd64.zip"
$Zip = "$env:TEMP\\sikuligo.zip"
$Dest = "$env:LOCALAPPDATA\\Programs\\sikuligo"

Invoke-WebRequest -Uri $Url -OutFile $Zip
New-Item -ItemType Directory -Force -Path $Dest | Out-Null
Expand-Archive -Path $Zip -DestinationPath $Dest -Force
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";$Dest", "User")
```

Open a new PowerShell and run:

```powershell
sikuligo.exe -listen 127.0.0.1:50051 -admin-listen :8080
```

## Distribution Options

- Current recommended channel: GitHub Releases artifacts (`.tar.gz`/`.zip`).
- Optional later channels:
  - Windows: Winget + Chocolatey package definitions.
  - Linux: APT/YUM repos or container image distribution.
