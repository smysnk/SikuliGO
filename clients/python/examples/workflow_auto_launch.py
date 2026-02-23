from __future__ import annotations

from sikuligo import Pattern, Screen

# Workflow: client auto-launches sikuligo if needed.
screen = Screen.start()
try:
    match = screen.click(Pattern("assets/pattern.png").exact())
    print(f"clicked match target at ({match.target_x}, {match.target_y})")
finally:
    screen.close()
