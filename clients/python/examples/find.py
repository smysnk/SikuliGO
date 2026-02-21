from generated.sikuli.v1 import sikuli_pb2 as pb
from sikuligo_client.client import SikuliGrpcClient, gray_image_from_rows


def main() -> int:
    client = SikuliGrpcClient()
    try:
        source = gray_image_from_rows(
            "source",
            [
                [10, 10, 10, 10, 10, 10, 10, 10],
                [10, 0, 255, 10, 10, 10, 10, 10],
                [10, 255, 0, 10, 0, 255, 10, 10],
                [10, 10, 10, 10, 255, 0, 10, 10],
                [10, 10, 10, 10, 10, 10, 10, 10],
            ],
        )
        needle = gray_image_from_rows(
            "needle",
            [
                [0, 255],
                [255, 0],
            ],
        )
        req = pb.FindRequest(
            source=source,
            pattern=pb.Pattern(image=needle, exact=True),
        )
        res = client.find(req)
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
