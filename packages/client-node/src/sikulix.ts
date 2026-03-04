import { LaunchOptions } from "./launcher";
import { Sikuli } from "./sikuli";
import { Image, ImageInput, loadPatternImage } from "./image";
import type { MatcherEngine, UnaryCallOptions } from "./client";

type ProtoPoint = { x?: number; y?: number };
type ProtoRect = { x?: number; y?: number; w?: number; h?: number };
type ProtoMatch = { rect?: ProtoRect; target?: ProtoPoint; score?: number; index?: number };
type ScreenQueryOptions = { region?: RegionBounds; timeout_millis?: number };

/**
 * Pattern mirrors SikuliX Pattern semantics:
 * - `similar(x)` tunes threshold
 * - `exact()` sets exact matching
 * - `targetOffset(dx, dy)` shifts click anchor from match center
 * - `resize(factor)` scales pattern before matching
 */
export class Pattern {
  readonly image: ImageInput | Image;
  private similarityValue?: number;
  private exactValue = false;
  private resizeValue?: number;
  private offsetValue = { x: 0, y: 0 };

  constructor(image: ImageInput | Image) {
    this.image = image;
  }

  /** Set similarity threshold in [0, 1]. Higher means stricter match. */
  similar(similarity: number): Pattern {
    this.similarityValue = Math.max(0, Math.min(1, similarity));
    this.exactValue = false;
    return this;
  }

  /** Convenience for exact matching (`similar(1)`). */
  exact(): Pattern {
    this.exactValue = true;
    this.similarityValue = 1;
    return this;
  }

  /** Shift click anchor from the default center target. */
  targetOffset(dx: number, dy: number): Pattern {
    this.offsetValue = { x: dx, y: dy };
    return this;
  }

  /** Scale the pattern before search (for DPI/zoom variance). */
  resize(factor: number): Pattern {
    this.resizeValue = factor > 0 ? factor : 1;
    return this;
  }

  toRequestPattern() {
    const decoded = loadPatternImage(this.image);
    return {
      image: decoded.image,
      mask: decoded.mask,
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

/**
 * Region is a scoped search surface similar to SikuliX Region.
 * All find/wait/click calls run within this region when bounds are set.
 */
export class Region {
  protected readonly session: Sikuli;
  protected readonly bounds?: RegionBounds;

  constructor(session: Sikuli, bounds?: RegionBounds) {
    this.session = session;
    this.bounds = bounds;
  }

  /** Find first match for image/pattern in this region. */
  async find(target: ImageInput | Image | Pattern, engine?: MatcherEngine): Promise<Match> {
    const pattern = toPattern(target);
    const out = (await this.session.findOnScreen({
      pattern: pattern.toRequestPattern(),
      opts: this.toScreenQueryOptions()
    }, this.toUnaryOptions(engine))) as { match?: ProtoMatch };
    if (!out.match) {
      throw new Error("match not found");
    }
    return new Match(out.match);
  }

  /** Check for presence with optional timeout; returns null if not found. */
  async exists(target: ImageInput | Image | Pattern, timeoutMs = 0, engine?: MatcherEngine): Promise<Match | null> {
    const pattern = toPattern(target);
    const out = (await this.session.existsOnScreen({
      pattern: pattern.toRequestPattern(),
      opts: this.toScreenQueryOptions(timeoutMs)
    }, this.toUnaryOptions(engine))) as { exists?: boolean; match?: ProtoMatch };
    if (!out.exists || !out.match) {
      return null;
    }
    return new Match(out.match);
  }

  /** Wait until pattern appears or timeout expires. */
  async wait(target: ImageInput | Image | Pattern, timeoutMs = 3000, engine?: MatcherEngine): Promise<Match> {
    const pattern = toPattern(target);
    const out = (await this.session.waitOnScreen({
      pattern: pattern.toRequestPattern(),
      opts: this.toScreenQueryOptions(timeoutMs)
    }, this.toUnaryOptions(engine))) as { match?: ProtoMatch };
    if (!out.match) {
      const msg = timeoutMs <= 0 ? "wait timeout" : `wait timeout after ${timeoutMs}ms`;
      throw new Error(msg);
    }
    return new Match(out.match);
  }

  /** Wait for match and click its target point. */
  async click(target: ImageInput | Image | Pattern, engine?: MatcherEngine): Promise<Match> {
    const pattern = toPattern(target);
    const out = (await this.session.clickOnScreen({
      pattern: pattern.toRequestPattern(),
      opts: this.toScreenQueryOptions()
    }, this.toUnaryOptions(engine))) as { match?: ProtoMatch };
    if (!out.match) {
      throw new Error("match not found");
    }
    return new Match(out.match);
  }

  /** Move mouse to matched target point. */
  async hover(target: ImageInput | Image | Pattern, engine?: MatcherEngine): Promise<Match> {
    const match = await this.find(target, engine);
    await this.session.moveMouse({ x: match.targetX, y: match.targetY });
    return match;
  }

  private toUnaryOptions(engine?: MatcherEngine): UnaryCallOptions | undefined {
    if (!engine) {
      return undefined;
    }
    return { matcherEngine: engine };
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

/**
 * Screen is the top-level automation entry point.
 * Use `Screen()` / `Screen.start()` for auto mode, `connect()` to attach to a running API, or `spawn()` to force spawn.
 */
export class Screen extends Region {
  private constructor(session: Sikuli) {
    super(session);
  }

  /** Force-launch a new SikuliGO API process and connect. */
  static async spawn(opts: LaunchOptions = {}): Promise<Screen> {
    const session = await Sikuli.spawn(opts);
    return new Screen(session);
  }

  /** Auto mode: connect first, spawn on fallback. */
  static async start(opts: LaunchOptions = {}): Promise<Screen> {
    const session = await Sikuli.launch(opts);
    return new Screen(session);
  }

  /** Alias for `start()` to match explicit auto semantics. */
  static async auto(opts: LaunchOptions = {}): Promise<Screen> {
    const session = await Sikuli.auto(opts);
    return new Screen(session);
  }

  /** Connect to an already-running API process. */
  static async connect(opts: LaunchOptions = {}): Promise<Screen> {
    const session = await Sikuli.connect(opts);
    return new Screen(session);
  }

  /** Create a bounded child region from this screen. */
  region(x: number, y: number, w: number, h: number): Region {
    return new Region(this.session, { x, y, w, h });
  }

  /** Close gRPC transport and stop spawned process if owned by this client. */
  async close(): Promise<void> {
    await this.session.close();
  }
}
