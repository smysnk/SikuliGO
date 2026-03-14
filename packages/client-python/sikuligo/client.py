from __future__ import annotations

import builtins
import inspect
import itertools
import os
import sys
import time
import zlib
from pathlib import Path
from typing import Iterable, Literal, Mapping, Sequence

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
PNG_SIGNATURE = b"\x89PNG\r\n\x1a\n"
DEBUG_ENABLED = str(os.getenv("SIKULI_DEBUG", "")).strip().lower() in {"1", "true", "yes", "on"}

ImageInput = str | bytes | bytearray | memoryview
RegionInput = tuple[int, int, int, int]
PointInput = tuple[int, int]
MatcherEngine = Literal["template", "orb", "akaze", "brisk", "kaze", "sift", "hybrid"]
_RUNTIME_CALL_COUNTER = itertools.count(1)


def _format_log_suffix(fields: Mapping[str, object]) -> str:
    parts = [
        f"{key}={value}"
        for key, value in fields.items()
        if key != "address" and value not in (None, "", ())
    ]
    return f" {' '.join(parts)}" if parts else ""


def _debug_log(message: str, **fields: object) -> None:
    if not DEBUG_ENABLED:
        return
    print(f"[debug] {message}{_format_log_suffix(fields)}", file=sys.stderr)


def _info_log(message: str, **fields: object) -> None:
    print(f"[info] {message}{_format_log_suffix(fields)}", file=sys.stderr)


def _error_log(message: str, **fields: object) -> None:
    print(f"[error] {message}{_format_log_suffix(fields)}", file=sys.stderr)


def _startup_wait_error_message(timeout_seconds: float) -> str:
    timeout_ms = int(round(timeout_seconds * 1000))
    return (
        f"Failed to connect before the deadline; startup/connect probe timed out after {timeout_ms}ms. "
        "Set SIKULI_DEBUG=1 to log launcher startup phases, binary resolution, and gRPC readiness timing. "
        "Examples should auto-launch sikuli-go when the packaged runtime can start on this machine. "
        "If you expect local auto-launch, verify sikuli-go is installed and executable. "
        "If you expect to connect to an existing server, verify a sikuli-go process is already listening at the configured address."
    )


def _normalize_matcher_engine(raw: str | None) -> MatcherEngine:
    normalized = str(raw or "").strip().lower()
    if normalized in ("orb", "akaze", "brisk", "kaze", "sift", "hybrid"):
        return normalized  # type: ignore[return-value]
    return "template"


def _matcher_engine_proto_value(raw: str | None) -> int:
    normalized = _normalize_matcher_engine(raw)
    if normalized == "orb":
        return int(pb.MATCHER_ENGINE_ORB)
    if normalized == "akaze":
        return int(pb.MATCHER_ENGINE_AKAZE)
    if normalized == "brisk":
        return int(pb.MATCHER_ENGINE_BRISK)
    if normalized == "kaze":
        return int(pb.MATCHER_ENGINE_KAZE)
    if normalized == "sift":
        return int(pb.MATCHER_ENGINE_SIFT)
    if normalized == "hybrid":
        return int(pb.MATCHER_ENGINE_HYBRID)
    return int(pb.MATCHER_ENGINE_TEMPLATE)


def _runtime_trace_hook():
    hook = getattr(builtins, "__sikuli_runtime_event__", None)
    return hook if callable(hook) else None


def _emit_runtime_event(payload: Mapping[str, object]) -> None:
    hook = _runtime_trace_hook()
    if hook is None:
        return
    try:
        hook(dict(payload))
    except Exception:
        # Ignore IDE trace hook failures so client behavior stays unchanged outside the editor.
        return


def _preview_text(value: object, max_length: int = 48) -> str:
    text = " ".join(str(value or "").split())
    if not text:
        return ""
    if len(text) <= max_length:
        return text
    return f"{text[: max_length - 3]}..."


def _rect_string(value: object) -> str:
    if value is None:
        return ""
    width = int(getattr(value, "w", 0) or 0)
    height = int(getattr(value, "h", 0) or 0)
    if width <= 0 or height <= 0:
        return ""
    x = int(getattr(value, "x", 0) or 0)
    y = int(getattr(value, "y", 0) or 0)
    return f"{x},{y},{width},{height}"


def _point_string(value: object) -> str:
    if value is None:
        return ""
    x = getattr(value, "x", None)
    y = getattr(value, "y", None)
    if x is None or y is None:
        return ""
    return f"{int(x)},{int(y)}"


def _request_summary_for_method(method_name: str, request: object) -> str:
    if method_name in ("FindOnScreen", "ExistsOnScreen", "WaitOnScreen", "ClickOnScreen"):
        pattern = getattr(request, "pattern", None)
        image = getattr(pattern, "image", None)
        opts = getattr(request, "opts", None)
        region = _rect_string(getattr(opts, "region", None))
        width = int(getattr(image, "width", 0) or 0)
        height = int(getattr(image, "height", 0) or 0)
        timeout_ms = int(getattr(opts, "timeout_millis", 0) or 0)
        parts = [
            f"region={region}" if region else "region=full_screen",
            f"pattern={width}x{height}" if width > 0 and height > 0 else "",
            "exact" if bool(getattr(pattern, "exact", False)) else "",
            (
                f"sim={float(getattr(pattern, 'similarity', 0.0)):.2f}"
                if float(getattr(pattern, "similarity", 0.0) or 0.0) > 0
                else ""
            ),
            f"timeout={timeout_ms}ms" if timeout_ms > 0 else "",
        ]
        if method_name == "ClickOnScreen":
            click_opts = getattr(request, "click_opts", None)
            button = getattr(click_opts, "button", "")
            if button:
                parts.append(f"button={button}")
        return " ".join(part for part in parts if part)

    if method_name in ("MoveMouse", "Click"):
        point = _point_string(request)
        parts = [f"point={point}" if point else ""]
        opts = getattr(request, "opts", None)
        button = getattr(opts, "button", "")
        if method_name == "Click" and button:
            parts.append(f"button={button}")
        return " ".join(part for part in parts if part)

    if method_name == "TypeText":
        text = _preview_text(getattr(request, "text", ""))
        return f'text="{text}" len={len(str(getattr(request, "text", "") or ""))}' if text else ""

    if method_name == "Hotkey":
        keys = list(getattr(request, "keys", []) or [])
        return f"keys={'+'.join(str(key) for key in keys)}" if keys else ""

    if method_name in ("OpenApp", "FocusApp", "CloseApp", "IsAppRunning", "ListWindows"):
        name = str(getattr(request, "name", "") or "")
        args = list(getattr(request, "args", []) or [])
        return f"app={name}{f' args={len(args)}' if args else ''}" if name else ""

    if method_name in ("ReadText", "FindText"):
        return "ocr request"

    return ""


def _response_details_for_method(method_name: str, response: object) -> dict[str, object]:
    if method_name in ("FindOnScreen", "ExistsOnScreen", "WaitOnScreen", "ClickOnScreen"):
        match = getattr(response, "match", None)
        match_rect = _rect_string(getattr(match, "rect", None))
        target_point = _point_string(getattr(match, "target", None))
        if match_rect:
            score = float(getattr(match, "score", 0.0) or 0.0)
            parts = [f"match={match_rect}"]
            if target_point:
                parts.append(f"target={target_point}")
            if score > 0:
                parts.append(f"score={score:.3f}")
            return {
                "responseSummary": " ".join(parts),
                "matchRect": match_rect,
                "targetPoint": target_point,
                "score": score,
            }
        if method_name == "ExistsOnScreen":
            exists = bool(getattr(response, "exists", False))
            return {
                "responseSummary": f"exists={'yes' if exists else 'no'}",
                "exists": exists,
            }
        return {
            "responseSummary": "no match",
            "exists": False,
        }

    if method_name == "IsAppRunning":
        running = bool(getattr(response, "running", False))
        return {
            "responseSummary": f"running={'yes' if running else 'no'}",
            "exists": running,
        }

    if method_name == "ListWindows":
        windows = getattr(response, "windows", None)
        count = len(windows) if windows is not None else 0
        return {
            "responseSummary": f"windows={count}",
        }

    if method_name == "ReadText":
        text = _preview_text(getattr(response, "text", ""))
        return {
            "responseSummary": f'text="{text}"' if text else "text=<empty>",
        }

    if method_name == "FindText":
        matches = getattr(response, "matches", None)
        count = len(matches) if matches is not None else 0
        return {
            "responseSummary": f"matches={count}",
        }

    return {}


class SikuliError(RuntimeError):
    def __init__(self, code: grpc.StatusCode, details: str, trace_id: str = "") -> None:
        suffix = f" trace_id={trace_id}" if trace_id else ""
        super().__init__(f"{code.name}: {details}{suffix}")
        self.code = code
        self.details = details
        self.trace_id = trace_id


class Sikuli:
    def __init__(
        self,
        *,
        address: str | None = None,
        auth_token: str | None = None,
        trace_id: str | None = None,
        timeout_seconds: float = DEFAULT_TIMEOUT_SECONDS,
        secure: bool = False,
        matcher_engine: str | None = None,
    ) -> None:
        self._address = address or os.getenv("SIKULI_GRPC_ADDR", DEFAULT_ADDR)
        self._auth_token = auth_token if auth_token is not None else os.getenv("SIKULI_GRPC_AUTH_TOKEN", "")
        self._trace_id = trace_id
        self._timeout_seconds = timeout_seconds
        self._matcher_engine = _normalize_matcher_engine(matcher_engine or os.getenv("SIKULI_MATCHER_ENGINE"))
        self._channel = (
            grpc.secure_channel(self._address, grpc.ssl_channel_credentials())
            if secure
            else grpc.insecure_channel(self._address)
        )
        self._stub = pb_grpc.SikuliServiceStub(self._channel)

    @property
    def address(self) -> str:
        return self._address

    @property
    def auth_token(self) -> str:
        return self._auth_token

    def close(self) -> None:
        _debug_log("grpc.close", address=self._address)
        self._channel.close()

    def wait_for_ready(self, timeout_seconds: float = DEFAULT_TIMEOUT_SECONDS) -> None:
        started_at = time.time()
        timeout_ms = int(round(timeout_seconds * 1000))
        _debug_log("grpc.wait_for_ready.start", address=self._address, timeout_ms=timeout_ms)
        try:
            grpc.channel_ready_future(self._channel).result(timeout=timeout_seconds)
        except grpc.FutureTimeoutError as exc:
            _debug_log(
                "grpc.wait_for_ready.error",
                address=self._address,
                timeout_ms=timeout_ms,
                duration_ms=int(round((time.time() - started_at) * 1000)),
                error="Failed to connect before the deadline",
            )
            raise TimeoutError(_startup_wait_error_message(timeout_seconds)) from exc
        _debug_log(
            "grpc.wait_for_ready.ok",
            address=self._address,
            duration_ms=int(round((time.time() - started_at) * 1000)),
        )

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
        matcher_engine: str | None = None,
    ):
        method = getattr(self._stub, method_name)
        request = self._with_matcher_engine(method_name, request, matcher_engine)
        call_id = f"{method_name.lower()}-{int(time.time() * 1000)}-{next(_RUNTIME_CALL_COUNTER)}"
        request_summary = _request_summary_for_method(method_name, request)
        _emit_runtime_event(
            {
                "type": "runtime:call:start",
                "callId": call_id,
                "method": method_name,
                "requestSummary": request_summary,
                "traceId": self._trace_id or "",
            }
        )
        started_at = time.time()
        try:
            response = method(
                request,
                timeout=timeout_seconds if timeout_seconds is not None else self._timeout_seconds,
                metadata=self._metadata(metadata),
            )
        except grpc.RpcError as err:
            trace_id = self._trace_id_from_error(err)
            runtime_error = SikuliError(err.code(), err.details() or "", trace_id)
            _emit_runtime_event(
                {
                    "type": "runtime:call:error",
                    "callId": call_id,
                    "method": method_name,
                    "requestSummary": request_summary,
                    "durationMs": int(round((time.time() - started_at) * 1000)),
                    "traceId": trace_id or self._trace_id or "",
                    "error": str(runtime_error),
                }
            )
            raise runtime_error from err
        _emit_runtime_event(
            {
                "type": "runtime:call:end",
                "callId": call_id,
                "method": method_name,
                "requestSummary": request_summary,
                "durationMs": int(round((time.time() - started_at) * 1000)),
                "traceId": self._trace_id or "",
                **_response_details_for_method(method_name, response),
            }
        )
        return response

    def _with_matcher_engine(self, method_name: str, request: object, matcher_engine: str | None):
        value = _matcher_engine_proto_value(matcher_engine or self._matcher_engine)
        if method_name in ("Find", "FindAll") and hasattr(request, "matcher_engine"):
            if int(getattr(request, "matcher_engine")) == int(pb.MATCHER_ENGINE_UNSPECIFIED):
                setattr(request, "matcher_engine", value)
            return request

        if method_name in ("FindOnScreen", "ExistsOnScreen", "WaitOnScreen", "ClickOnScreen"):
            opts = getattr(request, "opts", None)
            if isinstance(opts, pb.ScreenQueryOptions):
                if int(opts.matcher_engine) == int(pb.MATCHER_ENGINE_UNSPECIFIED):
                    opts.matcher_engine = value
            return request

        return request

    def find_on_screen(
        self,
        image: ImageInput,
        *,
        name: str | None = None,
        similarity: float | None = None,
        exact: bool = False,
        resize_factor: float | None = None,
        target_offset: PointInput | None = None,
        region: RegionInput | None = None,
        timeout_millis: int | None = None,
        interval_millis: int | None = None,
        timeout_seconds: float | None = None,
        engine: str | None = None,
    ):
        req = pb.FindOnScreenRequest(
            pattern=pattern_from_png(
                image,
                name=name,
                similarity=similarity,
                exact=exact,
                resize_factor=resize_factor,
                target_offset=target_offset,
            ),
            opts=screen_query_options(
                region=region,
                timeout_millis=timeout_millis,
                interval_millis=interval_millis,
                matcher_engine=engine,
            ),
        )
        return self._call("FindOnScreen", req, timeout_seconds=timeout_seconds, matcher_engine=engine)

    def exists_on_screen(
        self,
        image: ImageInput,
        *,
        name: str | None = None,
        similarity: float | None = None,
        exact: bool = False,
        resize_factor: float | None = None,
        target_offset: PointInput | None = None,
        region: RegionInput | None = None,
        timeout_millis: int | None = None,
        interval_millis: int | None = None,
        timeout_seconds: float | None = None,
        engine: str | None = None,
    ):
        req = pb.ExistsOnScreenRequest(
            pattern=pattern_from_png(
                image,
                name=name,
                similarity=similarity,
                exact=exact,
                resize_factor=resize_factor,
                target_offset=target_offset,
            ),
            opts=screen_query_options(
                region=region,
                timeout_millis=timeout_millis,
                interval_millis=interval_millis,
                matcher_engine=engine,
            ),
        )
        return self._call("ExistsOnScreen", req, timeout_seconds=timeout_seconds, matcher_engine=engine)

    def wait_on_screen(
        self,
        image: ImageInput,
        *,
        name: str | None = None,
        similarity: float | None = None,
        exact: bool = False,
        resize_factor: float | None = None,
        target_offset: PointInput | None = None,
        region: RegionInput | None = None,
        timeout_millis: int | None = None,
        interval_millis: int | None = None,
        timeout_seconds: float | None = None,
        engine: str | None = None,
    ):
        req = pb.WaitOnScreenRequest(
            pattern=pattern_from_png(
                image,
                name=name,
                similarity=similarity,
                exact=exact,
                resize_factor=resize_factor,
                target_offset=target_offset,
            ),
            opts=screen_query_options(
                region=region,
                timeout_millis=timeout_millis,
                interval_millis=interval_millis,
                matcher_engine=engine,
            ),
        )
        return self._call("WaitOnScreen", req, timeout_seconds=timeout_seconds, matcher_engine=engine)

    def click_on_screen(
        self,
        image: ImageInput,
        *,
        name: str | None = None,
        similarity: float | None = None,
        exact: bool = False,
        resize_factor: float | None = None,
        target_offset: PointInput | None = None,
        region: RegionInput | None = None,
        timeout_millis: int | None = None,
        interval_millis: int | None = None,
        button: str | None = None,
        delay_millis: int | None = None,
        timeout_seconds: float | None = None,
        engine: str | None = None,
    ):
        click_opts = pb.InputOptions()
        if button:
            click_opts.button = button
        if delay_millis is not None:
            click_opts.delay_millis = delay_millis
        req = pb.ClickOnScreenRequest(
            pattern=pattern_from_png(
                image,
                name=name,
                similarity=similarity,
                exact=exact,
                resize_factor=resize_factor,
                target_offset=target_offset,
            ),
            opts=screen_query_options(
                region=region,
                timeout_millis=timeout_millis,
                interval_millis=interval_millis,
                matcher_engine=engine,
            ),
            click_opts=click_opts,
        )
        return self._call("ClickOnScreen", req, timeout_seconds=timeout_seconds, matcher_engine=engine)

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


def _png_paeth(a: int, b: int, c: int) -> int:
    p = a + b - c
    pa = abs(p - a)
    pbv = abs(p - b)
    pc = abs(p - c)
    if pa <= pbv and pa <= pc:
        return a
    if pbv <= pc:
        return b
    return c


def _bytes_per_pixel(color_type: int) -> int:
    if color_type == 0:
        return 1
    if color_type == 2:
        return 3
    if color_type == 4:
        return 2
    if color_type == 6:
        return 4
    raise ValueError(f"unsupported PNG color type: {color_type}")


def _decode_png_to_gray_and_mask(data: bytes, name: str) -> tuple[pb.GrayImage, pb.GrayImage | None]:
    if not data.startswith(PNG_SIGNATURE):
        raise ValueError(f"not a PNG image: {name}")

    off = len(PNG_SIGNATURE)
    width = 0
    height = 0
    bit_depth = 0
    color_type = 0
    idat_chunks: list[bytes] = []

    while off + 12 <= len(data):
        length = int.from_bytes(data[off : off + 4], "big")
        chunk_type = data[off + 4 : off + 8]
        data_start = off + 8
        data_end = data_start + length
        if data_end + 4 > len(data):
            raise ValueError(f"corrupt PNG chunk: {name}")
        chunk = data[data_start:data_end]
        if chunk_type == b"IHDR":
            width = int.from_bytes(chunk[0:4], "big")
            height = int.from_bytes(chunk[4:8], "big")
            bit_depth = chunk[8]
            color_type = chunk[9]
            compression = chunk[10]
            png_filter = chunk[11]
            interlace = chunk[12]
            if bit_depth != 8:
                raise ValueError(f"unsupported PNG bit depth {bit_depth}: {name}")
            if compression != 0 or png_filter != 0 or interlace != 0:
                raise ValueError(f"unsupported PNG format (compression/filter/interlace): {name}")
        elif chunk_type == b"IDAT":
            idat_chunks.append(chunk)
        elif chunk_type == b"IEND":
            break
        off = data_end + 4

    if width <= 0 or height <= 0:
        raise ValueError(f"missing PNG dimensions: {name}")

    bpp = _bytes_per_pixel(color_type)
    stride = width * bpp
    inflated = zlib.decompress(b"".join(idat_chunks))
    expected_len = (stride + 1) * height
    if len(inflated) < expected_len:
        raise ValueError(f"corrupt PNG image payload: {name}")

    raw = bytearray(stride * height)
    src_off = 0
    for y in range(height):
        filt = inflated[src_off]
        src_off += 1
        row_start = y * stride
        for x in range(stride):
            cur = inflated[src_off]
            src_off += 1
            left = raw[row_start + x - bpp] if x >= bpp else 0
            up = raw[row_start + x - stride] if y > 0 else 0
            up_left = raw[row_start + x - stride - bpp] if y > 0 and x >= bpp else 0
            if filt == 0:
                out = cur
            elif filt == 1:
                out = (cur + left) & 0xFF
            elif filt == 2:
                out = (cur + up) & 0xFF
            elif filt == 3:
                out = (cur + ((left + up) // 2)) & 0xFF
            elif filt == 4:
                out = (cur + _png_paeth(left, up, up_left)) & 0xFF
            else:
                raise ValueError(f"unsupported PNG filter type {filt}: {name}")
            raw[row_start + x] = out

    pix = bytearray(width * height)
    alpha_mask = bytearray(width * height) if color_type in (4, 6) else None
    has_transparency = False
    for i in range(width * height):
        p = i * bpp
        alpha = 255
        if color_type == 0:
            gray = raw[p]
        elif color_type == 2:
            r, g, b = raw[p], raw[p + 1], raw[p + 2]
            gray = round(0.299 * r + 0.587 * g + 0.114 * b)
        elif color_type == 4:
            g = raw[p]
            alpha = raw[p + 1]
            a = alpha / 255.0
            gray = round(g * a + 255 * (1 - a))
        else:
            r, g, b = raw[p], raw[p + 1], raw[p + 2]
            alpha = raw[p + 3]
            a = alpha / 255.0
            gray = round((0.299 * r + 0.587 * g + 0.114 * b) * a + 255 * (1 - a))
        pix[i] = int(gray) & 0xFF
        if alpha_mask is not None:
            alpha_mask[i] = alpha & 0xFF
            if alpha < 255:
                has_transparency = True

    image = pb.GrayImage(name=name, width=width, height=height, pix=bytes(pix))
    if alpha_mask is not None and has_transparency:
        mask = pb.GrayImage(name=f"{name}.mask", width=width, height=height, pix=bytes(alpha_mask))
        return image, mask
    return image, None


def _pattern_image_from_png(image: ImageInput, *, name: str | None = None) -> tuple[pb.GrayImage, pb.GrayImage | None]:
    if isinstance(image, str):
        resolved = _resolve_image_path(image)
        with open(resolved, "rb") as handle:
            data = handle.read()
        image_name = name or resolved.name or "pattern.png"
        return _decode_png_to_gray_and_mask(data, image_name)
    if isinstance(image, memoryview):
        payload = image.tobytes()
        return _decode_png_to_gray_and_mask(payload, name or "pattern.png")
    if isinstance(image, (bytes, bytearray)):
        return _decode_png_to_gray_and_mask(bytes(image), name or "pattern.png")
    raise TypeError("image must be a local path or PNG bytes")


def _resolve_image_path(image: str) -> Path:
    raw = Path(image)
    if raw.is_absolute():
        return raw
    cwd_candidate = (Path.cwd() / raw).resolve()
    if cwd_candidate.exists():
        return cwd_candidate
    this_file = Path(__file__).resolve()
    for frame in inspect.stack()[1:]:
        frame_path = Path(frame.filename).resolve()
        if frame_path == this_file:
            continue
        candidate = (frame_path.parent / raw).resolve()
        if candidate.exists():
            return candidate
    return cwd_candidate


def gray_image_from_png(image: ImageInput, *, name: str | None = None) -> pb.GrayImage:
    gray, _ = _pattern_image_from_png(image, name=name)
    return gray


def pattern_from_png(
    image: ImageInput,
    *,
    name: str | None = None,
    similarity: float | None = None,
    exact: bool = False,
    resize_factor: float | None = None,
    target_offset: PointInput | None = None,
) -> pb.Pattern:
    gray, mask = _pattern_image_from_png(image, name=name)
    pattern = pb.Pattern(image=gray)
    if mask is not None:
        pattern.mask.CopyFrom(mask)
    if exact:
        pattern.exact = True
    elif similarity is not None:
        pattern.similarity = max(0.0, min(1.0, float(similarity)))
    if resize_factor is not None:
        pattern.resize_factor = float(resize_factor)
    if target_offset is not None:
        pattern.target_offset.CopyFrom(pb.Point(x=int(target_offset[0]), y=int(target_offset[1])))
    return pattern


def screen_query_options(
    *,
    region: RegionInput | None = None,
    timeout_millis: int | None = None,
    interval_millis: int | None = None,
    matcher_engine: str | None = None,
) -> pb.ScreenQueryOptions:
    opts = pb.ScreenQueryOptions()
    if region is not None:
        opts.region.CopyFrom(pb.Rect(x=int(region[0]), y=int(region[1]), w=int(region[2]), h=int(region[3])))
    if timeout_millis is not None:
        opts.timeout_millis = int(timeout_millis)
    if interval_millis is not None:
        opts.interval_millis = int(interval_millis)
    if matcher_engine is not None:
        opts.matcher_engine = _matcher_engine_proto_value(matcher_engine)
    return opts


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
