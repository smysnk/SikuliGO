from __future__ import annotations

from bootstrap import ensure_sikuli_go_on_path
from generated.sikuli.v1 import sikuli_pb2 as pb
from sikuligo import Screen

ensure_sikuli_go_on_path()

screen = Screen()
client = screen.client
try:
    client.move_mouse(pb.MoveMouseRequest(x=200, y=180, opts=pb.InputOptions(delay_millis=30)))
    client.click(pb.ClickRequest(x=200, y=180, opts=pb.InputOptions(button="left", delay_millis=20)))
    client.type_text(pb.TypeTextRequest(text="hello from python grpc", opts=pb.InputOptions(delay_millis=15)))
    client.hotkey(pb.HotkeyRequest(keys=["cmd", "a"]))
    print("input actions sent")
finally:
    screen.close()
