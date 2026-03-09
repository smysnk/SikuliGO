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

from .client import (
    DEFAULT_ADDR,
    DEFAULT_TIMEOUT_SECONDS,
    ImageInput,
    PointInput,
    RegionInput,
    Sikuli,
    _debug_log,
    _info_log,
)


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
    exe_names = (
        "sikuli-go.exe",
        "sikuli-go",
        "sikuligo.exe",
        "sikuligo",
        "sikuligrpc.exe",
        "sikuligrpc",
    )

    def resolve_path(candidate_path: str | Path, source: str) -> str:
        candidate = Path(candidate_path).expanduser().resolve()
        if candidate.exists() and os.access(candidate, os.X_OK):
            return str(candidate)
        raise FileNotFoundError(f"{source} is not executable: {candidate}")

    if binary_path:
        return resolve_path(binary_path, "configured binary")

    env_path = os.getenv("SIKULI_GO_BINARY_PATH", "").strip()
    if env_path:
        return resolve_path(env_path, "SIKULI_GO_BINARY_PATH")

    packaged_runtime = _resolve_packaged_runtime_binary(exe_names)
    if packaged_runtime:
        return packaged_runtime

    # Try common repo-local locations so examples work without env vars.
    probe_dirs: list[Path] = []
    cwd = Path.cwd()
    probe_dirs.extend([cwd, cwd.parent, cwd.parent.parent])

    this_file = Path(__file__).resolve()
    # .../packages/client-python/sikuligo/sikulix.py -> repo root
    if len(this_file.parents) >= 4:
        repo_root = this_file.parents[3]
        probe_dirs.extend([repo_root, repo_root / "packages" / "api"])

    seen: set[str] = set()
    for probe_dir in probe_dirs:
        key = str(probe_dir)
        if key in seen:
            continue
        seen.add(key)
        for exe_name in exe_names:
            candidate = probe_dir / exe_name
            if candidate.exists() and os.access(candidate, os.X_OK):
                return str(candidate.resolve())

    for name in exe_names:
        found = shutil.which(name)
        if found:
            return found

    raise FileNotFoundError(
        "Unable to resolve sikuli-go binary. Build it in repo root, install it on PATH, or set SIKULI_GO_BINARY_PATH."
    )


def _packaged_runtime_platform_dirs() -> list[str]:
    platform_dirs: list[str] = []
    if os.name == "nt" and os.environ.get("PROCESSOR_ARCHITECTURE", "").lower() in ("amd64", "x86_64"):
        platform_dirs.append("bin-win32-x64")
    elif os.name == "posix":
        uname = os.uname()
        machine = uname.machine.lower()
        sysname = uname.sysname.lower()
        if sysname == "darwin":
            if machine in ("arm64", "aarch64"):
                platform_dirs.append("bin-darwin-arm64")
            elif machine in ("x86_64", "amd64"):
                platform_dirs.append("bin-darwin-x64")
        elif sysname == "linux" and machine in ("x86_64", "amd64"):
            platform_dirs.append("bin-linux-x64")
    return platform_dirs


def _resolve_packaged_runtime_binary(
    exe_names: Sequence[str],
    package_root: Path | None = None,
) -> str | None:
    root = package_root or Path(__file__).resolve().parent
    for platform_dir in _packaged_runtime_platform_dirs():
        runtime_dir = root / "runtime" / platform_dir / "bin"
        for exe_name in exe_names:
            candidate = runtime_dir / exe_name
            if candidate.exists() and os.access(candidate, os.X_OK):
                return str(candidate.resolve())
    return None


def _merge_runtime_path(current_path: str | None) -> str:
    delimiter = ";" if os.name == "nt" else ":"
    existing = [part.strip() for part in (current_path or "").split(delimiter) if part.strip()]
    seen = {part.lower() if os.name == "nt" else part for part in existing}
    home_local = str(Path.home() / ".local" / "bin")
    if os.name == "posix" and os.uname().sysname.lower() == "darwin":
        defaults = ["/opt/homebrew/bin", "/usr/local/bin", "/usr/bin", "/bin", home_local]
    elif os.name == "posix":
        defaults = ["/usr/local/bin", "/usr/bin", "/bin", home_local]
    else:
        defaults = []
    for candidate in defaults:
        key = candidate.lower() if os.name == "nt" else candidate
        if key in seen:
            continue
        seen.add(key)
        existing.append(candidate)
    return delimiter.join(existing)


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
                f"sikuli-go exited before startup completed (code={child.returncode})"
            )
        remaining = deadline - time.monotonic()
        if remaining <= 0:
            raise TimeoutError(
                f"timeout waiting for sikuli-go startup on {session.address}"
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
        engine: str | None = None,
    ) -> Match:
        pattern = _to_pattern(target)
        out = self._session.find_on_screen(
            **pattern._request_kwargs(),
            region=self._bounds,
            timeout_millis=timeout_millis,
            interval_millis=interval_millis,
            timeout_seconds=timeout_seconds,
            engine=engine,
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
        engine: str | None = None,
    ) -> Match | None:
        pattern = _to_pattern(target)
        out = self._session.exists_on_screen(
            **pattern._request_kwargs(),
            region=self._bounds,
            timeout_millis=timeout_millis,
            interval_millis=interval_millis,
            timeout_seconds=timeout_seconds,
            engine=engine,
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
        engine: str | None = None,
    ) -> Match:
        pattern = _to_pattern(target)
        out = self._session.wait_on_screen(
            **pattern._request_kwargs(),
            region=self._bounds,
            timeout_millis=timeout_millis,
            interval_millis=interval_millis,
            timeout_seconds=timeout_seconds,
            engine=engine,
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
        engine: str | None = None,
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
            engine=engine,
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
        engine: str | None = None,
    ) -> Match:
        match = self.find(
            target,
            timeout_millis=timeout_millis,
            interval_millis=interval_millis,
            timeout_seconds=timeout_seconds,
            engine=engine,
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
    def spawn(
        cls,
        *,
        address: str | None = None,
        auth_token: str | None = None,
        trace_id: str | None = None,
        timeout_seconds: float = DEFAULT_TIMEOUT_SECONDS,
        secure: bool = False,
        matcher_engine: str | None = None,
        startup_timeout_seconds: float = DEFAULT_STARTUP_TIMEOUT_SECONDS,
        binary_path: str | None = None,
        admin_listen: str = "",
        sqlite_path: str | None = None,
        server_args: Sequence[str] | None = None,
        stdio: Literal["ignore", "pipe", "inherit"] = "ignore",
    ) -> Screen:
        resolved_address = address or os.getenv("SIKULI_GRPC_ADDR") or f"127.0.0.1:{_find_open_port()}"
        address_source = "option" if address else "env" if os.getenv("SIKULI_GRPC_ADDR") else "generated"
        token = auth_token or os.getenv("SIKULI_GRPC_AUTH_TOKEN") or secrets.token_hex(24)
        binary = _resolve_sikuli_binary(binary_path)
        binary_source = "option" if binary_path else "env" if os.getenv("SIKULI_GO_BINARY_PATH") else "resolver"
        resolved_sqlite_path = (
            sqlite_path
            or os.getenv("SIKULI_GO_SQLITE_PATH", "").strip()
            or "sikuli-go.db"
        )
        effective_stdio = "inherit" if stdio == "ignore" and os.getenv("SIKULI_DEBUG", "").strip().lower() in {"1", "true", "yes", "on"} else stdio
        stdout, stderr = _stdio_targets(effective_stdio)
        _debug_log(
            "launcher.start",
            mode="spawn",
            user_supplied_address="yes" if address_source in {"option", "env"} else "no",
            address_source=address_source,
            address=resolved_address,
            auth_token="yes" if token else "no",
            cwd=Path.cwd(),
            startup_timeout_requested_ms=int(round(startup_timeout_seconds * 1000)),
            startup_timeout_ms=int(round(startup_timeout_seconds * 1000)),
        )

        args = [
            binary,
            "-listen",
            resolved_address,
            "-admin-listen",
            admin_listen,
            "-auth-token",
            token,
            "-enable-reflection=false",
            "-sqlite-path",
            resolved_sqlite_path,
            *(server_args or []),
        ]
        _info_log(
            "launcher.spawn.start",
            binary=binary,
            binary_source=binary_source,
            address=resolved_address,
            address_source=address_source,
            admin_listen=admin_listen,
            sqlite_path=resolved_sqlite_path,
            auth_token="yes" if token else "no",
            server_args_extra_count=len(server_args or ()),
            stdio=effective_stdio,
            startup_timeout_requested_ms=int(round(startup_timeout_seconds * 1000)),
            startup_timeout_ms=int(round(startup_timeout_seconds * 1000)),
        )
        _debug_log(
            "launcher.spawn.args",
            args=" ".join(
                [
                    "-listen",
                    resolved_address,
                    "-admin-listen",
                    admin_listen,
                    "-enable-reflection=false",
                    "-sqlite-path",
                    resolved_sqlite_path,
                ]
            ),
        )
        child = subprocess.Popen(
            args,
            env={
                **os.environ,
                "PATH": _merge_runtime_path(os.environ.get("PATH", "")),
                "SIKULI_GRPC_AUTH_TOKEN": token,
            },
            stdout=stdout,
            stderr=stderr,
        )
        _debug_log("launcher.spawn.pid", pid=child.pid or "unknown")

        session = Sikuli(
            address=resolved_address,
            auth_token=token,
            trace_id=trace_id,
            timeout_seconds=timeout_seconds,
            secure=secure,
            matcher_engine=matcher_engine,
        )
        try:
            _debug_log(
                "launcher.spawn.wait.start",
                pid=child.pid or "unknown",
                startup_timeout_ms=int(round(startup_timeout_seconds * 1000)),
            )
            _wait_for_startup(session, child, startup_timeout_seconds)
        except Exception as exc:
            _debug_log(
                "launcher.spawn.error",
                pid=child.pid or "unknown",
                startup_timeout_ms=int(round(startup_timeout_seconds * 1000)),
                can_fallback_to_connect="no",
                child_exit_code=child.poll() if child.poll() is not None else "nil",
                child_running="yes" if child.poll() is None else "no",
                error=str(exc),
            )
            _stop_spawned_process(child)
            session.close()
            raise
        _info_log("launcher.spawn.ready", address=resolved_address, pid=child.pid or "unknown")

        return cls(
            session,
            child=child,
            meta=LaunchMeta(address=resolved_address, auth_token=token, spawned_server=True),
        )

    @classmethod
    def start(
        cls,
        *,
        address: str | None = None,
        auth_token: str | None = None,
        trace_id: str | None = None,
        timeout_seconds: float = DEFAULT_TIMEOUT_SECONDS,
        secure: bool = False,
        matcher_engine: str | None = None,
        startup_timeout_seconds: float = DEFAULT_STARTUP_TIMEOUT_SECONDS,
        binary_path: str | None = None,
        admin_listen: str = "",
        sqlite_path: str | None = None,
        server_args: Sequence[str] | None = None,
        stdio: Literal["ignore", "pipe", "inherit"] = "ignore",
    ) -> Screen:
        return cls.auto(
            address=address,
            auth_token=auth_token,
            trace_id=trace_id,
            timeout_seconds=timeout_seconds,
            secure=secure,
            matcher_engine=matcher_engine,
            startup_timeout_seconds=startup_timeout_seconds,
            binary_path=binary_path,
            admin_listen=admin_listen,
            sqlite_path=sqlite_path,
            server_args=server_args,
            stdio=stdio,
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
        matcher_engine: str | None = None,
        startup_timeout_seconds: float = DEFAULT_STARTUP_TIMEOUT_SECONDS,
        address_source_hint: str | None = None,
    ) -> Screen:
        resolved_address = address or os.getenv("SIKULI_GRPC_ADDR", DEFAULT_ADDR)
        resolved_auth_token = auth_token if auth_token is not None else os.getenv("SIKULI_GRPC_AUTH_TOKEN", "")
        address_source = address_source_hint or ("option" if address else "env" if os.getenv("SIKULI_GRPC_ADDR") else "default")
        _debug_log(
            "launcher.start",
            mode="connect",
            user_supplied_address="yes" if address_source in {"option", "env"} else "no",
            address_source=address_source,
            address=resolved_address,
            auth_token="yes" if resolved_auth_token else "no",
            cwd=Path.cwd(),
            startup_timeout_requested_ms=int(round(startup_timeout_seconds * 1000)),
            startup_timeout_ms=int(round(startup_timeout_seconds * 1000)),
        )
        _debug_log(
            "launcher.connect.start",
            address=resolved_address,
            address_source=address_source,
            auth_token="yes" if resolved_auth_token else "no",
            startup_timeout_ms=int(round(startup_timeout_seconds * 1000)),
        )
        session = Sikuli(
            address=resolved_address,
            auth_token=resolved_auth_token,
            trace_id=trace_id,
            timeout_seconds=timeout_seconds,
            secure=secure,
            matcher_engine=matcher_engine,
        )
        try:
            session.wait_for_ready(timeout_seconds=startup_timeout_seconds)
        except Exception as exc:
            _debug_log(
                "launcher.connect.error",
                address_source=address_source,
                startup_timeout_ms=int(round(startup_timeout_seconds * 1000)),
                error=str(exc),
            )
            session.close()
            raise
        _debug_log("launcher.connect.ready", address=resolved_address)
        return cls(
            session,
            child=None,
            meta=LaunchMeta(address=resolved_address, auth_token=resolved_auth_token, spawned_server=False),
        )

    @classmethod
    def auto(
        cls,
        *,
        address: str | None = None,
        auth_token: str | None = None,
        trace_id: str | None = None,
        timeout_seconds: float = DEFAULT_TIMEOUT_SECONDS,
        secure: bool = False,
        matcher_engine: str | None = None,
        startup_timeout_seconds: float = DEFAULT_STARTUP_TIMEOUT_SECONDS,
        binary_path: str | None = None,
        admin_listen: str = "",
        sqlite_path: str | None = None,
        server_args: Sequence[str] | None = None,
        stdio: Literal["ignore", "pipe", "inherit"] = "ignore",
    ) -> Screen:
        probe_address = address or os.getenv("SIKULI_GRPC_ADDR", DEFAULT_ADDR)
        _debug_log(
            "launcher.auto.start",
            probe_timeout_ms=1000,
            explicit_address="yes" if address else "no",
            env_address="yes" if os.getenv("SIKULI_GRPC_ADDR") else "no",
        )
        try:
            _debug_log("launcher.auto.probe.connect", probe_timeout_ms=1000)
            connected = cls.connect(
                address=probe_address,
                auth_token=auth_token,
                trace_id=trace_id,
                timeout_seconds=timeout_seconds,
                secure=secure,
                matcher_engine=matcher_engine,
                startup_timeout_seconds=1.0,
                address_source_hint="option" if address else "env" if os.getenv("SIKULI_GRPC_ADDR") else "auto-probe-default",
            )
            _debug_log(
                "launcher.auto.probe.connected_existing",
                spawn_attempted="no",
                spawned_server="yes" if connected.meta.spawned_server else "no",
            )
            return connected
        except Exception as exc:
            _debug_log(
                "launcher.auto.probe.failed",
                probe_timeout_ms=1000,
                reason=str(exc),
            )
            _debug_log(
                "launcher.auto.spawn.start",
                spawn_attempted="yes",
                startup_timeout_ms=int(round(startup_timeout_seconds * 1000)),
            )
            spawned = cls.spawn(
                address=address,
                auth_token=auth_token,
                trace_id=trace_id,
                timeout_seconds=timeout_seconds,
                secure=secure,
                matcher_engine=matcher_engine,
                startup_timeout_seconds=startup_timeout_seconds,
                binary_path=binary_path,
                admin_listen=admin_listen,
                sqlite_path=sqlite_path,
                server_args=server_args,
                stdio=stdio,
            )
            _debug_log(
                "launcher.auto.spawn.ready",
                spawned_server="yes" if spawned.meta.spawned_server else "no",
            )
            return spawned

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
