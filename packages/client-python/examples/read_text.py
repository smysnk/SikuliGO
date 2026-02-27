from generated.sikuli.v1 import sikuli_pb2 as pb
from sikuligo import Screen
from sikuligo.client import gray_image_from_rows


def main() -> int:
    screen = Screen()
    client = screen.client
    try:
        source = gray_image_from_rows(
            "ocr-source",
            [
                [220, 220, 220, 220],
                [220, 20, 20, 220],
                [220, 220, 220, 220],
            ],
        )

        read_req = pb.ReadTextRequest(
            source=source,
            params=pb.OCRParams(language="eng"),
        )
        read_res = client.read_text(read_req, timeout_seconds=10.0)
        print(f"read_text => {read_res.text!r}")

        find_req = pb.FindTextRequest(
            source=source,
            query="example",
            params=pb.OCRParams(language="eng", case_sensitive=False),
        )
        find_res = client.find_text(find_req, timeout_seconds=10.0)
        print(f"find_text matches => {len(find_res.matches)}")
    finally:
        screen.close()
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
