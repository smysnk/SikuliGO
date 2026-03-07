from __future__ import annotations

from bootstrap_runtime import ensure_project_venv_python

ensure_project_venv_python()

from bootstrap import ensure_sikuli_go_on_path
from generated.sikuli.v1 import sikuli_pb2 as pb
from sikuligo import Screen
from sikuligo.client import gray_image_from_rows

ensure_sikuli_go_on_path()

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

    read_req = pb.ReadTextRequest(source=source, params=pb.OCRParams(language="eng"))
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
