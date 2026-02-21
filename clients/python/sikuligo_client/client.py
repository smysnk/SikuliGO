from __future__ import annotations

import os
from typing import Iterable, Mapping, Sequence

import grpc

try:
    from generated.sikuli.v1 import sikuli_pb2 as pb
    from generated.sikuli.v1 import sikuli_pb2_grpc as pb_grpc
except ImportError as exc:  # pragma: no cover - runtime setup validation
    raise ImportError(
        "Missing generated Python gRPC stubs. Run ./scripts/clients/generate-python-stubs.sh first."
    ) from exc


DEFAULT_ADDR = "127.0.0.1:50051"
DEFAULT_TIMEOUT_SECONDS = 5.0
TRACE_HEADER = "x-trace-id"


class SikuliGrpcError(RuntimeError):
    def __init__(self, code: grpc.StatusCode, details: str, trace_id: str = "") -> None:
        suffix = f" trace_id={trace_id}" if trace_id else ""
        super().__init__(f"{code.name}: {details}{suffix}")
        self.code = code
        self.details = details
        self.trace_id = trace_id


class SikuliGrpcClient:
    def __init__(
        self,
        *,
        address: str | None = None,
        auth_token: str | None = None,
        trace_id: str | None = None,
        timeout_seconds: float = DEFAULT_TIMEOUT_SECONDS,
        secure: bool = False,
    ) -> None:
        self._address = address or os.getenv("SIKULI_GRPC_ADDR", DEFAULT_ADDR)
        self._auth_token = auth_token if auth_token is not None else os.getenv("SIKULI_GRPC_AUTH_TOKEN", "")
        self._trace_id = trace_id
        self._timeout_seconds = timeout_seconds
        self._channel = grpc.secure_channel(self._address, grpc.ssl_channel_credentials()) if secure else grpc.insecure_channel(self._address)
        self._stub = pb_grpc.SikuliServiceStub(self._channel)

    def close(self) -> None:
        self._channel.close()

    def _metadata(self, extra: Mapping[str, str] | None = None) -> Sequence[tuple[str, str]]:
        out: list[tuple[str, str]] = []
        if self._auth_token:
            out.append(("x-api-key", self._auth_token))
        if self._trace_id:
            out.append((TRACE_HEADER, self._trace_id))
        if extra:
            out.extend((k, v) for k, v in extra.items() if v)
        return out

    @staticmethod
    def _trace_id_from_error(err: grpc.RpcError) -> str:
        for getter in ("initial_metadata", "trailing_metadata"):
            fn = getattr(err, getter, None)
            if fn is None:
                continue
            md = fn()
            if not md:
                continue
            for key, value in md:
                if key.lower() == TRACE_HEADER:
                    return value
        return ""

    def _call(
        self,
        method_name: str,
        request: object,
        *,
        timeout_seconds: float | None = None,
        metadata: Mapping[str, str] | None = None,
    ):
        method = getattr(self._stub, method_name)
        try:
            return method(
                request,
                timeout=timeout_seconds if timeout_seconds is not None else self._timeout_seconds,
                metadata=self._metadata(metadata),
            )
        except grpc.RpcError as err:
            raise SikuliGrpcError(err.code(), err.details() or "", self._trace_id_from_error(err)) from err

    def find(self, request: pb.FindRequest, *, timeout_seconds: float | None = None):
        return self._call("Find", request, timeout_seconds=timeout_seconds)

    def find_all(self, request: pb.FindRequest, *, timeout_seconds: float | None = None):
        return self._call("FindAll", request, timeout_seconds=timeout_seconds)

    def read_text(self, request: pb.ReadTextRequest, *, timeout_seconds: float | None = None):
        return self._call("ReadText", request, timeout_seconds=timeout_seconds)

    def find_text(self, request: pb.FindTextRequest, *, timeout_seconds: float | None = None):
        return self._call("FindText", request, timeout_seconds=timeout_seconds)

    def move_mouse(self, request: pb.MoveMouseRequest, *, timeout_seconds: float | None = None):
        return self._call("MoveMouse", request, timeout_seconds=timeout_seconds)

    def click(self, request: pb.ClickRequest, *, timeout_seconds: float | None = None):
        return self._call("Click", request, timeout_seconds=timeout_seconds)

    def type_text(self, request: pb.TypeTextRequest, *, timeout_seconds: float | None = None):
        return self._call("TypeText", request, timeout_seconds=timeout_seconds)

    def hotkey(self, request: pb.HotkeyRequest, *, timeout_seconds: float | None = None):
        return self._call("Hotkey", request, timeout_seconds=timeout_seconds)

    def open_app(self, request: pb.AppActionRequest, *, timeout_seconds: float | None = None):
        return self._call("OpenApp", request, timeout_seconds=timeout_seconds)

    def focus_app(self, request: pb.AppActionRequest, *, timeout_seconds: float | None = None):
        return self._call("FocusApp", request, timeout_seconds=timeout_seconds)

    def close_app(self, request: pb.AppActionRequest, *, timeout_seconds: float | None = None):
        return self._call("CloseApp", request, timeout_seconds=timeout_seconds)

    def is_app_running(self, request: pb.AppActionRequest, *, timeout_seconds: float | None = None):
        return self._call("IsAppRunning", request, timeout_seconds=timeout_seconds)

    def list_windows(self, request: pb.AppActionRequest, *, timeout_seconds: float | None = None):
        return self._call("ListWindows", request, timeout_seconds=timeout_seconds)


def gray_image_from_rows(name: str, rows: Sequence[Sequence[int]]) -> pb.GrayImage:
    if not rows:
        raise ValueError("rows must not be empty")
    width = len(rows[0])
    if width == 0:
        raise ValueError("rows must have at least one column")
    pix = bytearray()
    for row in rows:
        if len(row) != width:
            raise ValueError("all rows must have the same width")
        pix.extend(v & 0xFF for v in row)
    return pb.GrayImage(name=name, width=width, height=len(rows), pix=bytes(pix))


def hotkey_request(keys: Iterable[str]) -> pb.HotkeyRequest:
    return pb.HotkeyRequest(keys=list(keys))
