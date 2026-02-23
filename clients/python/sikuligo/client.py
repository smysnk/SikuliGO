from __future__ import annotations

import inspect
import os
import zlib
from pathlib import Path
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
PNG_SIGNATURE = b"\x89PNG\r\n\x1a\n"

ImageInput = str | bytes | bytearray | memoryview
RegionInput = tuple[int, int, int, int]
PointInput = tuple[int, int]


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
    ) -> None:
        self._address = address or os.getenv("SIKULI_GRPC_ADDR", DEFAULT_ADDR)
        self._auth_token = auth_token if auth_token is not None else os.getenv("SIKULI_GRPC_AUTH_TOKEN", "")
        self._trace_id = trace_id
        self._timeout_seconds = timeout_seconds
        self._channel = (
            grpc.secure_channel(self._address, grpc.ssl_channel_credentials())
            if secure
            else grpc.insecure_channel(self._address)
        )
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
            raise SikuliError(err.code(), err.details() or "", self._trace_id_from_error(err)) from err

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
            ),
        )
        return self._call("FindOnScreen", req, timeout_seconds=timeout_seconds)

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
            ),
        )
        return self._call("ExistsOnScreen", req, timeout_seconds=timeout_seconds)

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
            ),
        )
        return self._call("WaitOnScreen", req, timeout_seconds=timeout_seconds)

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
            ),
            click_opts=click_opts,
        )
        return self._call("ClickOnScreen", req, timeout_seconds=timeout_seconds)

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


def _decode_png_to_gray(data: bytes, name: str) -> pb.GrayImage:
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
    for i in range(width * height):
        p = i * bpp
        if color_type == 0:
            gray = raw[p]
        elif color_type == 2:
            r, g, b = raw[p], raw[p + 1], raw[p + 2]
            gray = round(0.299 * r + 0.587 * g + 0.114 * b)
        elif color_type == 4:
            g = raw[p]
            a = raw[p + 1] / 255.0
            gray = round(g * a + 255 * (1 - a))
        else:
            r, g, b = raw[p], raw[p + 1], raw[p + 2]
            a = raw[p + 3] / 255.0
            gray = round((0.299 * r + 0.587 * g + 0.114 * b) * a + 255 * (1 - a))
        pix[i] = int(gray) & 0xFF

    return pb.GrayImage(name=name, width=width, height=height, pix=bytes(pix))


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
    if isinstance(image, str):
        resolved = _resolve_image_path(image)
        with open(resolved, "rb") as handle:
            data = handle.read()
        image_name = name or resolved.name or "pattern.png"
        return _decode_png_to_gray(data, image_name)
    if isinstance(image, memoryview):
        payload = image.tobytes()
        return _decode_png_to_gray(payload, name or "pattern.png")
    if isinstance(image, (bytes, bytearray)):
        return _decode_png_to_gray(bytes(image), name or "pattern.png")
    raise TypeError("image must be a local path or PNG bytes")


def pattern_from_png(
    image: ImageInput,
    *,
    name: str | None = None,
    similarity: float | None = None,
    exact: bool = False,
    resize_factor: float | None = None,
    target_offset: PointInput | None = None,
) -> pb.Pattern:
    pattern = pb.Pattern(image=gray_image_from_png(image, name=name))
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
) -> pb.ScreenQueryOptions:
    opts = pb.ScreenQueryOptions()
    if region is not None:
        opts.region.CopyFrom(pb.Rect(x=int(region[0]), y=int(region[1]), w=int(region[2]), h=int(region[3])))
    if timeout_millis is not None:
        opts.timeout_millis = int(timeout_millis)
    if interval_millis is not None:
        opts.interval_millis = int(interval_millis)
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
