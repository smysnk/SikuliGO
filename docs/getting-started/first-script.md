---
layout: guide
title: Getting Started First Script
nav_key: getting-started
kicker: Getting Started
lead: Run the shortest working automation example from a scaffolded Node.js or Python project.
---

## Who This Page Is For

Use this page after installation when you want the first successful automation command, not the deeper runtime details.

## Node.js Example

Scaffold and run:

```bash
yarn dlx @sikuligo/sikuli-go init:js-examples
cd sikuli-go-demo
yarn node examples/click.mjs
```

Core example:

```js
import { Screen, Pattern } from "@sikuligo/sikuli-go";

const screen = await Screen();
try {
  const match = await screen.click(Pattern("assets/pattern.png").exact());
  console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
} finally {
  await screen.close();
}
```

## Python Example

Scaffold and run:

```bash
pipx run sikuli-go init:py-examples
cd sikuli-go-demo
python3 examples/click.py
```

Core example:

```python
from sikuligo import Pattern, Screen

screen = Screen()
try:
    match = screen.click(Pattern("assets/pattern.png").exact())
    print(f"clicked match target at ({match.target_x}, {match.target_y})")
finally:
    screen.close()
```

## What To Expect

- The scaffolded project includes `assets/pattern.png`.
- The runtime is started automatically by the client in auto mode if an existing server is not available.
- The click example reports the resolved target coordinates after a successful match.

## If You Want The Go Path

The base implementation has its own runtime-facing example flow:

- [Golang API: First Program]({{ '/golang-api/first-program' | relative_url }})

## Next Pages

- [Getting Started: Dashboard]({{ '/getting-started/dashboard' | relative_url }})
- [Node.js Client: First Script]({{ '/nodejs-client/first-script' | relative_url }})
- [Python Client: First Script]({{ '/python-client/first-script' | relative_url }})
