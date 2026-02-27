from __future__ import annotations

from sikuligo import Pattern, Screen

# Primary constructor: connect first, spawn fallback.
screen = Screen()
try:
    match = screen.click(Pattern("assets/pattern.png").exact())
    print(f"clicked match target at ({match.target_x}, {match.target_y})")
finally:
    screen.close()
