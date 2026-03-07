from __future__ import annotations

from bootstrap_runtime import ensure_project_venv_python

ensure_project_venv_python()

from bootstrap import ensure_sikuli_go_on_path
from sikuligo import Pattern, Screen

ensure_sikuli_go_on_path()


def main() -> int:
    screen = Screen()
    try:
        match = screen.find(Pattern("assets/pattern.png").exact(), timeout_millis=3000)
        print(
            f"match rect=({match.x},{match.y},{match.w},{match.h}) "
            f"score={match.score:.3f} target=({match.target_x},{match.target_y})"
        )
    finally:
        screen.close()
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
