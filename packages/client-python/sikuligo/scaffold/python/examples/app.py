from __future__ import annotations

from bootstrap_runtime import ensure_project_venv_python

ensure_project_venv_python()

import os

from bootstrap import ensure_sikuli_go_on_path
from generated.sikuli.v1 import sikuli_pb2 as pb
from sikuligo import Screen

ensure_sikuli_go_on_path()

screen = Screen()
client = screen.client
app_name = os.getenv("SIKULI_APP_NAME", "Calculator")
try:
    client.open_app(pb.AppActionRequest(name=app_name))
    running = client.is_app_running(pb.AppActionRequest(name=app_name))
    print(f"is running => {running.running}")

    windows = client.list_windows(pb.AppActionRequest(name=app_name))
    print(f"windows => {len(windows.windows)}")

    client.focus_app(pb.AppActionRequest(name=app_name))
    client.close_app(pb.AppActionRequest(name=app_name))
    print("app control actions sent")
finally:
    screen.close()
