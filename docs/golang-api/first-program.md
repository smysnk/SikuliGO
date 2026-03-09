---
layout: guide
title: Golang API First Program
nav_key: golang-api
kicker: Golang API
lead: Connect to a running runtime from Go, load a pattern image, and execute a live screen action against the matched target.
---

## Minimal Runtime-Backed Example

```go
package main

import (
	"image/png"
	"log"
	"os"

	"github.com/smysnk/sikuligo/pkg/sikuli"
)

func main() {
	runtime, err := sikuli.NewRuntime("127.0.0.1:50051")
	if err != nil {
		log.Fatal(err)
	}
	defer runtime.Close()

	screen, err := runtime.PrimaryScreen()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("assets/pattern.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	src, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	img, err := sikuli.NewImageFromAny("pattern", src)
	if err != nil {
		log.Fatal(err)
	}

	pattern, err := sikuli.NewPattern(img)
	if err != nil {
		log.Fatal(err)
	}

	match, err := screen.Find(pattern.Exact())
	if err != nil {
		log.Fatal(err)
	}

	if err := match.Click(sikuli.InputOptions{}); err != nil {
		log.Fatal(err)
	}
}
```

## What This Example Shows

- `NewRuntime` connects to the live runtime over gRPC.
- `PrimaryScreen` gets a live screen-backed `Screen`.
- `NewImageFromAny` and `NewPattern` convert an image file into a match target.
- `screen.Find` returns a live `Match`.
- `match.Click` uses the resolved target point from that live match.

## Before Running It

- Start `sikuli-go` locally.
- Make sure `assets/pattern.png` exists and is visible on the screen you are testing against.
- Run inside an active desktop session.

## Next Pages

- [Golang API: Installation]({{ '/golang-api/installation' | relative_url }})
- [Golang API: Core Types]({{ '/golang-api/core-types' | relative_url }})
- [Golang API: Runtime And Reference]({{ '/golang-api/runtime-and-reference' | relative_url }})
