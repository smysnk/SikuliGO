from __future__ import annotations

from bootstrap_runtime import ensure_project_venv_python

ensure_project_venv_python()

from bootstrap import ensure_sikuli_go_on_path
from sikuligo import Pattern, Screen

ensure_sikuli_go_on_path()

# Primary constructor: connect first, then spawn fallback.
screen = Screen()
try:
    match = screen.click(Pattern("assets/pattern.png").exact())
    print(f"clicked match target at ({match.target_x}, {match.target_y})")
finally:
    screen.close()
