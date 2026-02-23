from __future__ import annotations

import os
import secrets
import shutil
import socket
import subprocess
import time
from dataclasses import dataclass
from pathlib import Path
from typing import Literal, Sequence

from generated.sikuli.v1 import sikuli_pb2 as pb

from .client import DEFAULT_ADDR, DEFAULT_TIMEOUT_SECONDS, ImageInput, PointInput, RegionInput, Sikuli


DEFAULT_STARTUP_TIMEOUT_SECONDS = 10.0


@dataclass(frozen=True)
class LaunchMeta:
    address: str
    auth_token: str
    spawned_server: bool


class Pattern:
    def __init__(self, image: ImageInput) -> None:
        self.image = image
        self._similarity: float | None = None
        self._exact = False
        self._resize_factor: float | None = None
        self._target_offset: PointInput = (0, 0)

    def similar(self, similarity: float) -> Pattern:
        self._similarity = max(0.0, min(1.0, float(similarity)))
        self._exact = False
        return self

    def exact(self) -> Pattern:
        self._exact = True
        self._similarity = 1.0
        return self

    def target_offset(self, dx: int, dy: int) -> Pattern:
        self._target_offset = (int(dx), int(dy))
        return self

    def resize(self, factor: float) -> Pattern:
        self._resize_factor = float(factor) if factor > 0 else 1.0
        return self

    def _request_kwargs(self) -> dict[str, object]:
        return {
            "image": self.image,
            "similarity": self._similarity,
            "exact": self._exact,
            "resize_factor": self._resize_factor,
            "target_offset": self._target_offset,
        }


class Match:
    def __init__(self, match) -> None:
        self.raw = match
        self.x = int(match.rect.x)
        self.y = int(match.rect.y)
        self.w = int(match.rect.w)
        self.h = int(match.rect.h)
        self.score = float(match.score)
        self.target_x = int(match.target.x)
        self.target_y = int(match.target.y)
        self.index = int(match.index)


def _to_pattern(target: ImageInput | Pattern) -> Pattern:
    if isinstance(target, Pattern):
        return target
    return Pattern(target)


def _find_open_port() -> int:
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
        sock.bind(("127.0.0.1", 0))
        return int(sock.getsockname()[1])


def _resolve_sikuli_binary(binary_path: str | None = None) -> str:
    if binary_path:
        candidate = Path(binary_path).expanduser().resolve()
        if candidate.exists() and os.access(candidate, os.X_OK):
            return str(candidate)
        raise FileNotFoundError(f"configured binary is not executable: {candidate}")

    env_path = os.getenv("SIKULIGO_BINARY_PATH", "").strip()
    if env_path:
        candidate = Path(env_path).expanduser().resolve()
        if candidate.exists() and os.access(candidate, os.X_OK):
            return str(candidate)
        raise FileNotFoundError(f"SIKULIGO_BINARY_PATH is not executable: {candidate}")

    for name in ("sikuligo", "sikuligrpc"):
        found = shutil.which(name)
        if found:
            return found

    raise FileNotFoundError(
        "Unable to resolve sikuligo binary. Install it on PATH or set SIKULIGO_BINARY_PATH."
    )


def _stdio_targets(mode: Literal["ignore", "pipe", "inherit"]) -> tuple[int | None, int | None]:
    if mode == "ignore":
        return subprocess.DEVNULL, subprocess.DEVNULL
    if mode == "pipe":
        return subprocess.PIPE, subprocess.PIPE
    if mode == "inherit":
        return None, None
    raise ValueError("stdio must be one of: ignore, pipe, inherit")


def _stop_spawned_process(child: subprocess.Popen | None, timeout_seconds: float = 3.0) -> None:
    if child is None or child.poll() is not None:
        return
    child.terminate()
    try:
        child.wait(timeout=timeout_seconds)
    except subprocess.TimeoutExpired:
        child.kill()
        child.wait()


def _wait_for_startup(session: Sikuli, child: subprocess.Popen, timeout_seconds: float) -> None:
    deadline = time.monotonic() + timeout_seconds
    while True:
        if child.poll() is not None:
            raise RuntimeError(
                f"sikuligo exited before startup completed (code={child.returncode})"
            )
        remaining = deadline - time.monotonic()
        if remaining <= 0:
            raise TimeoutError(
                f"timeout waiting for sikuligo startup on {session.address}"
            )
        try:
            session.wait_for_ready(timeout_seconds=min(0.2, remaining))
            return
        except TimeoutError:
            continue


class Region:
    def __init__(self, session: Sikuli, bounds: RegionInput | None = None) -> None:
        self._session = session
        self._bounds = bounds

    def find(
        self,
        target: ImageInput | Pattern,
        *,
        timeout_millis: int | None = None,
        interval_millis: int | None = None,
        timeout_seconds: float | None = None,
    ) -> Match:
        pattern = _to_pattern(target)
        out = self._session.find_on_screen(
            **pattern._request_kwargs(),
            region=self._bounds,
            timeout_millis=timeout_millis,
            interval_millis=interval_millis,
            timeout_seconds=timeout_seconds,
        )
        if not out.match:
            raise RuntimeError("match not found")
        return Match(out.match)

    def exists(
        self,
        target: ImageInput | Pattern,
        timeout_millis: int = 0,
        *,
        interval_millis: int | None = None,
        timeout_seconds: float | None = None,
    ) -> Match | None:
        pattern = _to_pattern(target)
        out = self._session.exists_on_screen(
            **pattern._request_kwargs(),
            region=self._bounds,
            timeout_millis=timeout_millis,
            interval_millis=interval_millis,
            timeout_seconds=timeout_seconds,
        )
        if not out.exists or not out.match:
            return None
        return Match(out.match)

    def wait(
        self,
        target: ImageInput | Pattern,
        timeout_millis: int = 3000,
        *,
        interval_millis: int | None = None,
        timeout_seconds: float | None = None,
    ) -> Match:
        pattern = _to_pattern(target)
        out = self._session.wait_on_screen(
            **pattern._request_kwargs(),
            region=self._bounds,
            timeout_millis=timeout_millis,
            interval_millis=interval_millis,
            timeout_seconds=timeout_seconds,
        )
        if not out.match:
            if timeout_millis <= 0:
                raise TimeoutError("wait timeout")
            raise TimeoutError(f"wait timeout after {timeout_millis}ms")
        return Match(out.match)

    def click(
        self,
        target: ImageInput | Pattern,
        *,
        timeout_millis: int | None = None,
        interval_millis: int | None = None,
        button: str | None = None,
        delay_millis: int | None = None,
        timeout_seconds: float | None = None,
    ) -> Match:
        pattern = _to_pattern(target)
        out = self._session.click_on_screen(
            **pattern._request_kwargs(),
            region=self._bounds,
            timeout_millis=timeout_millis,
            interval_millis=interval_millis,
            button=button,
            delay_millis=delay_millis,
            timeout_seconds=timeout_seconds,
        )
        if not out.match:
            raise RuntimeError("match not found")
        return Match(out.match)

    def hover(
        self,
        target: ImageInput | Pattern,
        *,
        timeout_millis: int | None = None,
        interval_millis: int | None = None,
        timeout_seconds: float | None = None,
    ) -> Match:
        match = self.find(
            target,
            timeout_millis=timeout_millis,
            interval_millis=interval_millis,
            timeout_seconds=timeout_seconds,
        )
        self._session.move_mouse(
            pb.MoveMouseRequest(x=match.target_x, y=match.target_y),
            timeout_seconds=timeout_seconds,
        )
        return match


class Screen(Region):
    def __init__(
        self,
        session: Sikuli,
        child: subprocess.Popen | None = None,
        meta: LaunchMeta | None = None,
    ) -> None:
        super().__init__(session)
        self._child = child
        self.meta = meta or LaunchMeta(
            address=session.address,
            auth_token=session.auth_token,
            spawned_server=child is not None,
        )
        self._closed = False

    @property
    def client(self) -> Sikuli:
        return self._session

    @classmethod
    def start(
        cls,
        *,
        address: str | None = None,
        auth_token: str | None = None,
        trace_id: str | None = None,
        timeout_seconds: float = DEFAULT_TIMEOUT_SECONDS,
        secure: bool = False,
        startup_timeout_seconds: float = DEFAULT_STARTUP_TIMEOUT_SECONDS,
        binary_path: str | None = None,
        admin_listen: str = "",
        server_args: Sequence[str] | None = None,
        stdio: Literal["ignore", "pipe", "inherit"] = "ignore",
    ) -> Screen:
        resolved_address = address or os.getenv("SIKULI_GRPC_ADDR") or f"127.0.0.1:{_find_open_port()}"
        token = auth_token or os.getenv("SIKULI_GRPC_AUTH_TOKEN") or secrets.token_hex(24)
        binary = _resolve_sikuli_binary(binary_path)
        stdout, stderr = _stdio_targets(stdio)

        args = [
            binary,
            "-listen",
            resolved_address,
            "-admin-listen",
            admin_listen,
            "-auth-token",
            token,
            "-enable-reflection=false",
            *(server_args or []),
        ]
        child = subprocess.Popen(
            args,
            env={
                **os.environ,
                "SIKULI_GRPC_AUTH_TOKEN": token,
            },
            stdout=stdout,
            stderr=stderr,
        )

        session = Sikuli(
            address=resolved_address,
            auth_token=token,
            trace_id=trace_id,
            timeout_seconds=timeout_seconds,
            secure=secure,
        )
        try:
            _wait_for_startup(session, child, startup_timeout_seconds)
        except Exception:
            _stop_spawned_process(child)
            session.close()
            raise

        return cls(
            session,
            child=child,
            meta=LaunchMeta(address=resolved_address, auth_token=token, spawned_server=True),
        )

    @classmethod
    def connect(
        cls,
        *,
        address: str | None = None,
        auth_token: str | None = None,
        trace_id: str | None = None,
        timeout_seconds: float = DEFAULT_TIMEOUT_SECONDS,
        secure: bool = False,
        startup_timeout_seconds: float = DEFAULT_STARTUP_TIMEOUT_SECONDS,
    ) -> Screen:
        resolved_address = address or os.getenv("SIKULI_GRPC_ADDR", DEFAULT_ADDR)
        resolved_auth_token = auth_token if auth_token is not None else os.getenv("SIKULI_GRPC_AUTH_TOKEN", "")
        session = Sikuli(
            address=resolved_address,
            auth_token=resolved_auth_token,
            trace_id=trace_id,
            timeout_seconds=timeout_seconds,
            secure=secure,
        )
        session.wait_for_ready(timeout_seconds=startup_timeout_seconds)
        return cls(
            session,
            child=None,
            meta=LaunchMeta(address=resolved_address, auth_token=resolved_auth_token, spawned_server=False),
        )

    def region(self, x: int, y: int, w: int, h: int) -> Region:
        return Region(self._session, (int(x), int(y), int(w), int(h)))

    def close(self) -> None:
        if self._closed:
            return
        self._closed = True
        self._session.close()
        _stop_spawned_process(self._child)

    def __enter__(self) -> Screen:
        return self

    def __exit__(self, exc_type, exc, tb) -> None:
        self.close()
