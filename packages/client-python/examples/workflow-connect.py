from __future__ import annotations

from bootstrap import ensure_sikuli_go_on_path
from sikuligo import Pattern, Screen

ensure_sikuli_go_on_path()

# Connect-only workflow (requires sikuli-go already running).
screen = Screen.connect()
try:
    match = screen.click(Pattern("assets/pattern.png").exact())
    print(f"clicked match target at ({match.target_x}, {match.target_y})")
finally:
    screen.close()
