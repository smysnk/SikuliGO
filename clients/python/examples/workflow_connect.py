from __future__ import annotations

import os

from sikuligo import Pattern, Screen

screen = Screen.connect(address=os.getenv("SIKULI_GRPC_ADDR", "127.0.0.1:50051"))
try:
    match = screen.click(Pattern("assets/pattern.png").exact())
    print(f"clicked match target at ({match.target_x}, {match.target_y})")
finally:
    screen.close()
