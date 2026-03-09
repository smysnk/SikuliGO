---
layout: guide
title: Python Client First Script
nav_key: python-client
kicker: Python Client
lead: Write the first Python automation script with the current Screen, Pattern, Region, and Match wrapper surface.
---

## Minimal Example

```python
from sikuligo import Pattern, Screen

screen = Screen()
try:
    match = screen.click(Pattern("assets/pattern.png").exact())
    print(f"clicked match target at ({match.target_x}, {match.target_y})")
finally:
    screen.close()
```

## What The Main Objects Do

- `Screen()` starts in auto mode and returns a live screen session.
- `Pattern("...")` describes the image target and matching rules.
- `Match` exposes geometry, score, and resolved target coordinates.
- `screen.region(x, y, w, h)` creates a bounded search area when you do not want to search the full screen.

## Common Variations

### Search Before Clicking

```python
from sikuligo import Pattern, Screen

screen = Screen()
try:
    match = screen.find(Pattern("assets/pattern.png").exact())
    print(match.score)
finally:
    screen.close()
```

### Use A Region

```python
from sikuligo import Pattern, Screen

screen = Screen()
try:
    region = screen.region(100, 100, 800, 600)
    match = region.wait(Pattern("assets/pattern.png"), timeout_millis=3000)
    print(match.target_x, match.target_y)
finally:
    screen.close()
```

## Next Pages

- [Python Client: Installation]({{ '/python-client/installation' | relative_url }})
- [Python Client: Runtime]({{ '/python-client/runtime' | relative_url }})
- [Python Client: Troubleshooting]({{ '/python-client/troubleshooting' | relative_url }})
