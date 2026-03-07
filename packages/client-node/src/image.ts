import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";
import zlib from "node:zlib";

export type GrayImage = {
  name: string;
  width: number;
  height: number;
  pix: Buffer;
};

export type PatternImage = {
  image: GrayImage;
  mask?: GrayImage;
};

export type ImageFormat = "png";

export type ImageInput =
  | string
  | Buffer
  | Uint8Array
  | {
      bytes: Buffer | Uint8Array;
      format?: ImageFormat;
      name?: string;
    };

export class Image {
  readonly input: ImageInput;

  constructor(input: ImageInput) {
    this.input = input;
  }
}

function paethPredictor(a: number, b: number, c: number): number {
  const p = a + b - c;
  const pa = Math.abs(p - a);
  const pb = Math.abs(p - b);
  const pc = Math.abs(p - c);
  if (pa <= pb && pa <= pc) return a;
  if (pb <= pc) return b;
  return c;
}

function bytesPerPixel(colorType: number): number {
  switch (colorType) {
    case 0:
      return 1;
    case 2:
      return 3;
    case 4:
      return 2;
    case 6:
      return 4;
    default:
      throw new Error(`Unsupported PNG color type: ${colorType}`);
  }
}

function decodePNGToGrayAndMask(data: Buffer, name: string): PatternImage {
  const sig = Buffer.from([0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a]);
  if (!data.subarray(0, 8).equals(sig)) {
    throw new Error(`Not a PNG image: ${name}`);
  }

  let off = 8;
  let width = 0;
  let height = 0;
  let bitDepth = 0;
  let colorType = 0;
  const idatChunks: Buffer[] = [];

  while (off + 12 <= data.length) {
    const length = data.readUInt32BE(off);
    const type = data.toString("ascii", off + 4, off + 8);
    const dataStart = off + 8;
    const dataEnd = dataStart + length;
    if (dataEnd + 4 > data.length) {
      throw new Error(`Corrupt PNG chunk: ${name}`);
    }
    const chunk = data.subarray(dataStart, dataEnd);
    if (type === "IHDR") {
      width = chunk.readUInt32BE(0);
      height = chunk.readUInt32BE(4);
      bitDepth = chunk[8];
      colorType = chunk[9];
      const compression = chunk[10];
      const filter = chunk[11];
      const interlace = chunk[12];
      if (bitDepth !== 8) {
        throw new Error(`Unsupported PNG bit depth ${bitDepth}: ${name}`);
      }
      if (compression !== 0 || filter !== 0 || interlace !== 0) {
        throw new Error(`Unsupported PNG format (compression/filter/interlace): ${name}`);
      }
    } else if (type === "IDAT") {
      idatChunks.push(chunk);
    } else if (type === "IEND") {
      break;
    }
    off = dataEnd + 4;
  }

  if (width <= 0 || height <= 0) {
    throw new Error(`Missing PNG dimensions: ${name}`);
  }

  const bpp = bytesPerPixel(colorType);
  const stride = width * bpp;
  const inflated = zlib.inflateSync(Buffer.concat(idatChunks));
  const expected = (stride + 1) * height;
  if (inflated.length < expected) {
    throw new Error(`Corrupt PNG image payload: ${name}`);
  }

  const raw = Buffer.alloc(stride * height);
  let srcOff = 0;
  for (let y = 0; y < height; y++) {
    const filter = inflated[srcOff++];
    const rowStart = y * stride;
    for (let x = 0; x < stride; x++) {
      const cur = inflated[srcOff++];
      const left = x >= bpp ? raw[rowStart + x - bpp] : 0;
      const up = y > 0 ? raw[rowStart + x - stride] : 0;
      const upLeft = y > 0 && x >= bpp ? raw[rowStart + x - stride - bpp] : 0;
      let out = cur;
      switch (filter) {
        case 0:
          out = cur;
          break;
        case 1:
          out = (cur + left) & 0xff;
          break;
        case 2:
          out = (cur + up) & 0xff;
          break;
        case 3:
          out = (cur + ((left + up) >> 1)) & 0xff;
          break;
        case 4:
          out = (cur + paethPredictor(left, up, upLeft)) & 0xff;
          break;
        default:
          throw new Error(`Unsupported PNG filter type ${filter}: ${name}`);
      }
      raw[rowStart + x] = out;
    }
  }

  const pix = Buffer.alloc(width * height);
  const alphaMask = colorType === 4 || colorType === 6 ? Buffer.alloc(width * height) : undefined;
  let hasTransparency = false;
  for (let i = 0; i < width * height; i++) {
    const p = i * bpp;
    let gray = 0;
    let alpha = 255;
    if (colorType === 0) {
      gray = raw[p];
    } else if (colorType === 2) {
      const r = raw[p];
      const g = raw[p + 1];
      const b = raw[p + 2];
      gray = Math.round(0.299 * r + 0.587 * g + 0.114 * b);
    } else if (colorType === 4) {
      const g = raw[p];
      alpha = raw[p + 1];
      const a = alpha / 255;
      gray = Math.round(g * a + 255 * (1 - a));
    } else {
      const r = raw[p];
      const g = raw[p + 1];
      const b = raw[p + 2];
      alpha = raw[p + 3];
      const a = alpha / 255;
      gray = Math.round((0.299 * r + 0.587 * g + 0.114 * b) * a + 255 * (1 - a));
    }
    pix[i] = gray;
    if (alphaMask) {
      alphaMask[i] = alpha;
      if (alpha < 255) {
        hasTransparency = true;
      }
    }
  }

  const out: PatternImage = {
    image: { name, width, height, pix }
  };
  if (alphaMask && hasTransparency) {
    out.mask = {
      name: `${name}.mask`,
      width,
      height,
      pix: alphaMask
    };
  }
  return out;
}

function isSdkInternalFrame(filePath: string): boolean {
  const normalized = path.normalize(filePath);
  const markers = [
    `${path.sep}clients${path.sep}node${path.sep}src${path.sep}`,
    `${path.sep}clients${path.sep}node${path.sep}dist${path.sep}src${path.sep}`,
    `${path.sep}node_modules${path.sep}@sikuligo${path.sep}sikuli-go${path.sep}dist${path.sep}src${path.sep}`
  ];
  return markers.some((marker) => normalized.includes(marker));
}

function stackFrameFiles(): string[] {
  const stack = new Error().stack ?? "";
  const lines = stack.split("\n").slice(1);
  const files: string[] = [];
  for (const rawLine of lines) {
    const line = rawLine.trim();
    const withParens = line.match(/\((.*):\d+:\d+\)$/);
    const noParens = line.match(/^at (.*):\d+:\d+$/);
    let frame = withParens?.[1] ?? noParens?.[1];
    if (!frame) {
      continue;
    }
    if (frame.startsWith("node:")) {
      continue;
    }
    if (frame.startsWith("file://")) {
      try {
        frame = fileURLToPath(frame);
      } catch {
        continue;
      }
    }
    if (!path.isAbsolute(frame)) {
      continue;
    }
    files.push(path.normalize(frame));
  }
  return files;
}

function resolveImagePath(input: string): string {
  if (path.isAbsolute(input)) {
    return input;
  }
  if (fs.existsSync(input)) {
    return path.resolve(input);
  }
  for (const framePath of stackFrameFiles()) {
    if (isSdkInternalFrame(framePath)) {
      continue;
    }
    const candidate = path.resolve(path.dirname(framePath), input);
    if (fs.existsSync(candidate)) {
      return candidate;
    }
  }
  return path.resolve(input);
}

function inputToBytes(input: ImageInput): { bytes: Buffer; name: string; format: ImageFormat } {
  if (typeof input === "string") {
    const resolvedPath = resolveImagePath(input);
    const bytes = fs.readFileSync(resolvedPath);
    const name = path.basename(resolvedPath);
    return { bytes, name, format: "png" };
  }
  if (Buffer.isBuffer(input)) {
    return { bytes: input, name: "image.png", format: "png" };
  }
  if (input instanceof Uint8Array) {
    return { bytes: Buffer.from(input), name: "image.png", format: "png" };
  }
  const bytes = Buffer.isBuffer(input.bytes) ? input.bytes : Buffer.from(input.bytes);
  const format = input.format ?? "png";
  const name = input.name ?? `image.${format}`;
  return { bytes, name, format };
}

/**
 * Load an image and derive `{image, mask}` for pattern matching.
 * For transparent PNGs, alpha is exported as mask so transparent areas are ignored by matchers.
 */
export function loadPatternImage(input: ImageInput | Image): PatternImage {
  const normalized = input instanceof Image ? input.input : input;
  const { bytes, name, format } = inputToBytes(normalized);
  if (format !== "png") {
    throw new Error(`Unsupported image format: ${format} (currently only png is supported)`);
  }
  return decodePNGToGrayAndMask(bytes, name);
}

/** Load an image as grayscale pixels suitable for direct RPC GrayImage fields. */
export function loadGrayImage(input: ImageInput | Image): GrayImage {
 return loadPatternImage(input).image;
}

export function cropGrayImage(img: GrayImage, x: number, y: number, w: number, h: number): GrayImage {
  const x0 = Math.max(0, x);
  const y0 = Math.max(0, y);
  const x1 = Math.min(img.width, x + w);
  const y1 = Math.min(img.height, y + h);
  if (x1 <= x0 || y1 <= y0) {
    throw new Error("crop region is outside source image");
  }

  const cw = x1 - x0;
  const ch = y1 - y0;
  const pix = Buffer.alloc(cw * ch);
  for (let row = 0; row < ch; row++) {
    const srcStart = (y0 + row) * img.width + x0;
    const dstStart = row * cw;
    img.pix.copy(pix, dstStart, srcStart, srcStart + cw);
  }
  return {
    name: `${img.name}-crop`,
    width: cw,
    height: ch,
    pix
  };
}
