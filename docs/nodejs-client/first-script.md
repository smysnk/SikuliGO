---
layout: guide
title: Node.js Client First Script
nav_key: nodejs-client
kicker: Node.js Client
lead: Write the first Node.js automation script with the current Screen, Pattern, Region, and Match surface.
---

## Minimal Example

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

## What The Main Objects Do

- `Screen()` starts in auto mode and returns a live screen session.
- `Pattern("...")` describes the image target and matching rules.
- `Match` exposes geometry, score, and resolved target coordinates.
- `screen.region(x, y, w, h)` creates a bounded search area when you do not want to search the full screen.

## Common Variations

### Search Before Clicking

```js
const screen = await Screen();
try {
  const match = await screen.find(Pattern("assets/pattern.png").exact());
  console.log(match.score);
} finally {
  await screen.close();
}
```

### Use A Region

```js
const screen = await Screen();
try {
  const region = screen.region(100, 100, 800, 600);
  const match = await region.wait(Pattern("assets/pattern.png"), 3000);
  console.log(match.targetX, match.targetY);
} finally {
  await screen.close();
}
```

## Next Pages

- [Node.js Client: Installation]({{ '/nodejs-client/installation' | relative_url }})
- [Node.js Client: Runtime]({{ '/nodejs-client/runtime' | relative_url }})
- [Node.js Client: Troubleshooting]({{ '/nodejs-client/troubleshooting' | relative_url }})
