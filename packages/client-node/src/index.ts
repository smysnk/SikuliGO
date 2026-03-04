import { Sikuli as SikuliClass } from "./sikuli";
import { Screen as ScreenClass, Region, Pattern as PatternClass, Match } from "./sikulix";
import type { LaunchOptions } from "./launcher";
import type { ImageInput } from "./image";
export { resolveSikuliBinary } from "./binary";
export { launchSikuli, stopSpawnedProcess } from "./launcher";
export { SikuliError } from "./client";
export { Image, loadGrayImage, loadPatternImage } from "./image";
export { Region, Match };

/**
 * Top-level Sikuli constructor in auto mode.
 * Equivalent to `Sikuli.launch(opts)`.
 */
export const Sikuli = Object.assign(
  async (opts: LaunchOptions = {}) => await SikuliClass.launch(opts),
  {
    launch: async (opts: LaunchOptions = {}) => await SikuliClass.launch(opts),
    connect: async (opts: LaunchOptions = {}) => await SikuliClass.connect(opts),
    spawn: async (opts: LaunchOptions = {}) => await SikuliClass.spawn(opts)
  }
);

/**
 * Top-level Screen constructor in auto mode.
 * Equivalent to `Screen.start(opts)`.
 */
export const Screen = Object.assign(
  async (opts: LaunchOptions = {}) => await ScreenClass.start(opts),
  {
    start: async (opts: LaunchOptions = {}) => await ScreenClass.start(opts),
    connect: async (opts: LaunchOptions = {}) => await ScreenClass.connect(opts),
    spawn: async (opts: LaunchOptions = {}) => await ScreenClass.spawn(opts)
  }
);

/** Build a Pattern from local image path or PNG bytes. */
export const Pattern = (image: ImageInput) => new PatternClass(image);

export type {
  LaunchOptions,
  LaunchResult
} from "./launcher";
export type {
  InputOptions,
  MoveMouseRequest,
  ClickRequest,
  TypeTextRequest,
  LaunchResultMeta
} from "./sikuli";
export type {
  RpcMessage,
  MatcherEngine,
  SikuliOptions,
  UnaryCallOptions
} from "./client";
export type {
  GrayImage,
  PatternImage,
  ImageFormat,
  ImageInput
} from "./image";
