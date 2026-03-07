from __future__ import annotations

from bootstrap import ensure_sikuli_go_on_path
from sikuligo import Pattern, Screen

ensure_sikuli_go_on_path()

screen = Screen(startup_timeout_seconds=10.0)
try:
    match = screen.click(Pattern("assets/pattern.png").exact())
    print(f"clicked match target at ({match.target_x}, {match.target_y})")
    print("automation actions sent")
finally:
    screen.close()
