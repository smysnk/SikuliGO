import { LaunchOptions } from "./launcher";
import { Sikuli } from "./sikuli";
import { Image, ImageInput, loadGrayImage } from "./image";

type ProtoPoint = { x?: number; y?: number };
type ProtoRect = { x?: number; y?: number; w?: number; h?: number };
type ProtoMatch = { rect?: ProtoRect; target?: ProtoPoint; score?: number; index?: number };
type ScreenQueryOptions = { region?: RegionBounds; timeout_millis?: number };

export class Pattern {
  readonly image: ImageInput | Image;
  private similarityValue?: number;
  private exactValue = false;
  private resizeValue?: number;
  private offsetValue = { x: 0, y: 0 };

  constructor(image: ImageInput | Image) {
    this.image = image;
  }

  similar(similarity: number): Pattern {
    this.similarityValue = Math.max(0, Math.min(1, similarity));
    this.exactValue = false;
    return this;
  }

  exact(): Pattern {
    this.exactValue = true;
    this.similarityValue = 1;
    return this;
  }

  targetOffset(dx: number, dy: number): Pattern {
    this.offsetValue = { x: dx, y: dy };
    return this;
  }

  resize(factor: number): Pattern {
    this.resizeValue = factor > 0 ? factor : 1;
    return this;
  }

  toRequestPattern() {
    const image = loadGrayImage(this.image);
    return {
      image,
      similarity: this.similarityValue,
      exact: this.exactValue,
      resize_factor: this.resizeValue,
      target_offset: this.offsetValue
    };
  }
}

export class Match {
  readonly x: number;
  readonly y: number;
  readonly w: number;
  readonly h: number;
  readonly score: number;
  readonly targetX: number;
  readonly targetY: number;
  readonly index: number;

  constructor(match: ProtoMatch, offset = { x: 0, y: 0 }) {
    this.x = (match.rect?.x ?? 0) + offset.x;
    this.y = (match.rect?.y ?? 0) + offset.y;
    this.w = match.rect?.w ?? 0;
    this.h = match.rect?.h ?? 0;
    this.score = match.score ?? 0;
    this.targetX = (match.target?.x ?? 0) + offset.x;
    this.targetY = (match.target?.y ?? 0) + offset.y;
    this.index = match.index ?? 0;
  }
}

type RegionBounds = { x: number; y: number; w: number; h: number };

function toPattern(target: ImageInput | Image | Pattern): Pattern {
  if (target instanceof Pattern) {
    return target;
  }
  if (target instanceof Image) {
    return new Pattern(target);
  }
  return new Pattern(target);
}

export class Region {
  protected readonly session: Sikuli;
  protected readonly bounds?: RegionBounds;

  constructor(session: Sikuli, bounds?: RegionBounds) {
    this.session = session;
    this.bounds = bounds;
  }

  async find(target: ImageInput | Image | Pattern): Promise<Match> {
    const pattern = toPattern(target);
    const out = (await this.session.findOnScreen({
      pattern: pattern.toRequestPattern(),
      opts: this.toScreenQueryOptions()
    })) as { match?: ProtoMatch };
    if (!out.match) {
      throw new Error("match not found");
    }
    return new Match(out.match);
  }

  async exists(target: ImageInput | Image | Pattern, timeoutMs = 0): Promise<Match | null> {
    const pattern = toPattern(target);
    const out = (await this.session.existsOnScreen({
      pattern: pattern.toRequestPattern(),
      opts: this.toScreenQueryOptions(timeoutMs)
    })) as { exists?: boolean; match?: ProtoMatch };
    if (!out.exists || !out.match) {
      return null;
    }
    return new Match(out.match);
  }

  async wait(target: ImageInput | Image | Pattern, timeoutMs = 3000): Promise<Match> {
    const pattern = toPattern(target);
    const out = (await this.session.waitOnScreen({
      pattern: pattern.toRequestPattern(),
      opts: this.toScreenQueryOptions(timeoutMs)
    })) as { match?: ProtoMatch };
    if (!out.match) {
      const msg = timeoutMs <= 0 ? "wait timeout" : `wait timeout after ${timeoutMs}ms`;
      throw new Error(msg);
    }
    return new Match(out.match);
  }

  async click(target: ImageInput | Image | Pattern): Promise<Match> {
    const pattern = toPattern(target);
    const out = (await this.session.clickOnScreen({
      pattern: pattern.toRequestPattern(),
      opts: this.toScreenQueryOptions()
    })) as { match?: ProtoMatch };
    if (!out.match) {
      throw new Error("match not found");
    }
    return new Match(out.match);
  }

  async hover(target: ImageInput | Image | Pattern): Promise<Match> {
    const match = await this.find(target);
    await this.session.moveMouse({ x: match.targetX, y: match.targetY });
    return match;
  }

  private toScreenQueryOptions(timeoutMs?: number): ScreenQueryOptions {
    const opts: ScreenQueryOptions = {};
    if (this.bounds && this.bounds.w > 0 && this.bounds.h > 0) {
      opts.region = { ...this.bounds };
    }
    if (typeof timeoutMs === "number" && timeoutMs > 0) {
      opts.timeout_millis = timeoutMs;
    }
    return opts;
  }
}

export class Screen extends Region {
  private constructor(session: Sikuli) {
    super(session);
  }

  static async start(opts: LaunchOptions = {}): Promise<Screen> {
    const session = await Sikuli.launch(opts);
    return new Screen(session);
  }

  static async connect(opts: LaunchOptions = {}): Promise<Screen> {
    const session = await Sikuli.connect(opts);
    return new Screen(session);
  }

  static async auto(opts: LaunchOptions = {}): Promise<Screen> {
    const session = await Sikuli.auto(opts);
    return new Screen(session);
  }

  region(x: number, y: number, w: number, h: number): Region {
    return new Region(this.session, { x, y, w, h });
  }

  async close(): Promise<void> {
    await this.session.close();
  }
}
