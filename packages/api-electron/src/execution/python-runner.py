from __future__ import annotations

import json
import os
import runpy
import socket
import sys
import threading
import builtins
from pathlib import Path
from typing import Any


class StepController:
    def __init__(self, host: str, port: int, *, start_paused: bool) -> None:
        self.socket = socket.create_connection((host, port))
        self.socket_reader = self.socket.makefile("r", encoding="utf-8")
        self.socket_lock = threading.Lock()
        self.state_lock = threading.Lock()
        self.pause_before_next = start_paused
        self.pause_before_runtime_call = False
        self.run_to_line: int | None = None
        self.current_step: dict[str, Any] | None = None
        self.pending_resume: threading.Event | None = None
        self.closed = False

        self.listener = threading.Thread(target=self._listen, name="sikuli-python-step-listener", daemon=True)
        self.listener.start()

    def _send_event(self, payload: dict[str, Any]) -> None:
        if self.closed:
            return
        line = f"{json.dumps(payload)}\n".encode("utf-8")
        with self.socket_lock:
            self.socket.sendall(line)

    def _listen(self) -> None:
        try:
            for line in self.socket_reader:
                if not line.strip():
                    continue
                message = json.loads(line)
                if isinstance(message, dict):
                    self._handle_command(message)
        except OSError:
            pass
        finally:
            self._release_pending_resume()

    def _release_pending_resume(self) -> None:
        with self.state_lock:
            pending = self.pending_resume
            self.pending_resume = None
            self.pause_before_next = False
        if pending is not None:
            pending.set()

    def _pause_for_reason(self, reason: str, **extra: Any) -> None:
        with self.state_lock:
            step = dict(self.current_step or {})
            pending = self.pending_resume
            if pending is None:
                pending = threading.Event()
                self.pending_resume = pending

        self._send_event(
            {
                "type": "step:pause",
                "reason": reason,
                **extra,
                **step,
            }
        )
        pending.wait()

    def _handle_command(self, message: dict[str, Any]) -> None:
        command = str(message.get("command", ""))
        if command == "pause":
            with self.state_lock:
                if self.pending_resume is None:
                    self.pause_before_next = True
                    self.pause_before_runtime_call = False
                    self.run_to_line = None
            return

        if command not in {"resume", "step", "run-to-line", "continue-to-runtime-call"}:
            return

        target_line: int | None = None
        with self.state_lock:
            pending = self.pending_resume
            step = dict(self.current_step or {})
            if command == "resume":
                self.pause_before_next = False
                self.pause_before_runtime_call = False
                self.run_to_line = None
            elif command == "step":
                self.pause_before_next = True
                self.pause_before_runtime_call = False
                self.run_to_line = None
            elif command == "continue-to-runtime-call":
                self.pause_before_next = False
                self.pause_before_runtime_call = True
                self.run_to_line = None
            else:
                target_line = int(message.get("line", 0))
                if target_line < 1:
                    return
                self.pause_before_next = False
                self.pause_before_runtime_call = False
                self.run_to_line = target_line

            if pending is None:
                return
            self.pending_resume = None

        self._send_event(
            {
                "type": "step:resume",
                "mode": "runtime-call" if command == "continue-to-runtime-call" else "run-to-line" if command == "run-to-line" else command,
                "targetLine": target_line,
                **step,
            }
        )
        pending.set()

    def finish_current_step(self) -> None:
        with self.state_lock:
            step = self.current_step
            self.current_step = None
            builtins.__sikuli_ide_current_step__ = None
        if step is None:
            return
        self._send_event({"type": "step:end", **step})

    def step_hook(self, line: int | None, column: int | None, statement_id: str | None) -> None:
        self.finish_current_step()

        step = {
            "line": line,
            "column": column,
            "statementId": statement_id,
        }

        with self.state_lock:
            self.current_step = step
            builtins.__sikuli_ide_current_step__ = dict(step)
            target_line = self.run_to_line
            should_pause = self.pause_before_next
            if target_line is not None and line == target_line:
                self.run_to_line = None
                self.pause_before_next = False
                self.pause_before_runtime_call = False
                pending = None
            elif should_pause:
                self.pause_before_next = False
                pending = None
            else:
                pending = None

        self._send_event({"type": "step:start", **step})

        if target_line is not None and line == target_line:
            self._pause_for_reason("run-to-line", targetLine=target_line)
            return

        if should_pause:
            self._pause_for_reason("step")
            return

    def runtime_event_hook(self, payload: dict[str, Any]) -> None:
        if not isinstance(payload, dict):
            return

        with self.state_lock:
            step = dict(self.current_step or {})
            should_pause = self.pause_before_runtime_call and payload.get("type") == "runtime:call:start"
            if should_pause:
                self.pause_before_runtime_call = False

        event_payload = dict(payload)
        for key, value in step.items():
            event_payload.setdefault(key, value)
        self._send_event(event_payload)

        if should_pause:
            self._pause_for_reason(
                "runtime-call",
                runtimeMethod=event_payload.get("method"),
                runtimeCallId=event_payload.get("callId"),
            )

    def send_step_error(self, message: str) -> None:
        with self.state_lock:
            step = dict(self.current_step or {})
        self._send_event({"type": "step:error", "error": message, **step})

    def close(self) -> None:
        if self.closed:
            return
        self.closed = True
        self._release_pending_resume()
        try:
            self.socket.shutdown(socket.SHUT_RDWR)
        except OSError:
            pass
        finally:
            self.socket.close()


def main(argv: list[str]) -> int:
    if len(argv) != 2:
        print("usage: python-runner.py <instrumented.py>", file=sys.stderr)
        return 2

    instrumented_path = Path(argv[1]).resolve()
    script_dir = str(instrumented_path.parent)
    if script_dir not in sys.path:
        sys.path.insert(0, script_dir)
    cwd = os.getcwd()
    if cwd not in sys.path:
        sys.path.insert(0, cwd)

    controller = StepController(
        os.environ["SIKULI_GO_STEP_HOST"],
        int(os.environ["SIKULI_GO_STEP_PORT"]),
        start_paused=os.environ.get("SIKULI_GO_STEP_START_PAUSED") == "1",
    )
    builtins.__sikuli_runtime_event__ = controller.runtime_event_hook
    builtins.__sikuli_ide_current_step__ = None

    try:
        runpy.run_path(
            str(instrumented_path),
            run_name="__main__",
            init_globals={"__sikuli_step": controller.step_hook},
        )
        controller.finish_current_step()
        return 0
    except SystemExit as error:
        code = error.code
        exit_code = 0 if code is None else int(code) if isinstance(code, int) else 1
        if exit_code == 0:
            controller.finish_current_step()
            return 0
        controller.send_step_error(str(code))
        raise
    except BaseException as error:  # noqa: BLE001
        controller.send_step_error(str(error))
        raise
    finally:
        controller.close()


if __name__ == "__main__":
    raise SystemExit(main(sys.argv))
