from __future__ import annotations

from sikuligo import Pattern, Screen


def main() -> int:
    screen = Screen.connect()
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
