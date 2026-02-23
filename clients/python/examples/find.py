from __future__ import annotations

from sikuligo_client.client import SikuliGrpcClient


def main() -> int:
    client = SikuliGrpcClient()
    try:
        res = client.find_on_screen("assets/pattern.png", exact=True, timeout_millis=3000)
        match = res.match
        print(
            f"match rect=({match.rect.x},{match.rect.y},{match.rect.w},{match.rect.h}) "
            f"score={match.score:.3f} target=({match.target.x},{match.target.y})"
        )
    finally:
        client.close()
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
