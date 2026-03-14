#!/usr/bin/env node

import { existsSync, mkdirSync, rmSync } from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { spawnSync } from "node:child_process";
import { createRequire } from "node:module";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const require = createRequire(import.meta.url);
const repoRoot = path.resolve(__dirname, "..", "..");
const outputPath = path.join(repoRoot, "docs", "images", "editor-demo.webp");
const framesDir = path.join(repoRoot, ".tmp", "editor-demo-frames");
const chromePath =
  process.env.PLAYWRIGHT_CHROME_PATH || "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome";
const editorUrl = process.env.SIKULI_GO_EDITOR_DEMO_URL || "http://127.0.0.1:3100/?demo=1";
const frameCount = 23;
const frameDelayMs = 400;
const framePlaybackMs = 850;
const finalFramePlaybackMs = 2200;
const cursorKeyframes = [
  { frame: 0, x: 188, y: 250 },
  { frame: 2, x: 190, y: 282 },
  { frame: 4, x: 900, y: 250 },
  { frame: 6, x: 186, y: 314 },
  { frame: 8, x: 184, y: 404 },
  { frame: 10, x: 952, y: 496 },
  { frame: 12, x: 170, y: 28 },
  { frame: 15, x: 126, y: 214 },
  { frame: 18, x: 126, y: 314 },
  { frame: 20, x: 558, y: 472 },
  { frame: 22, x: 560, y: 520 },
];
const cursorClickFrames = new Set([0, 2, 4, 6, 8, 10, 12, 15, 18, 20, 22]);

if (!existsSync(chromePath)) {
  console.error(`Chrome executable not found at ${chromePath}`);
  process.exit(1);
}

let chromium;
try {
  ({ chromium } = require("playwright-core"));
} catch (error) {
  console.error("playwright-core is required. Run this script with:");
  console.error("  npm install --prefix .tmp/playwright-tools playwright-core");
  console.error("  NODE_PATH=.tmp/playwright-tools/node_modules node scripts/editor/generate-demo-webp.mjs");
  console.error(error instanceof Error ? error.message : String(error));
  process.exit(1);
}

rmSync(framesDir, { recursive: true, force: true });
mkdirSync(framesDir, { recursive: true });
mkdirSync(path.dirname(outputPath), { recursive: true });

function interpolate(start, end, ratio) {
  return Math.round(start + (end - start) * ratio);
}

function buildCursorFrames() {
  return Array.from({ length: frameCount }, (_, index) => {
    const nextKeyframe = cursorKeyframes.find((keyframe) => keyframe.frame >= index) || cursorKeyframes[cursorKeyframes.length - 1];
    const previousKeyframe =
      [...cursorKeyframes].reverse().find((keyframe) => keyframe.frame <= index) || cursorKeyframes[0];
    if (nextKeyframe.frame === previousKeyframe.frame) {
      return { x: nextKeyframe.x, y: nextKeyframe.y, clicking: cursorClickFrames.has(index) };
    }
    const ratio = (index - previousKeyframe.frame) / (nextKeyframe.frame - previousKeyframe.frame);
    return {
      x: interpolate(previousKeyframe.x, nextKeyframe.x, ratio),
      y: interpolate(previousKeyframe.y, nextKeyframe.y, ratio),
      clicking: cursorClickFrames.has(index),
    };
  });
}

async function installCursorOverlay(page) {
  const pointerSvg = `
    <svg class="__sikuli-demo-cursor-pointer" viewBox="0 0 32 48" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
      <path d="M4 2L4 31L11 24L16 40L22 37L17 22L27 22L4 2Z" fill="#f8fafc" stroke="#0f172a" stroke-width="2" stroke-linejoin="round" />
    </svg>
  `;

  await page.evaluate((svgMarkup) => {
    if (document.getElementById("__sikuli-demo-cursor")) {
      return;
    }

    const style = document.createElement("style");
    style.textContent = `
      #__sikuli-demo-cursor {
        position: fixed;
        left: 0;
        top: 0;
        width: 36px;
        height: 48px;
        pointer-events: none;
        z-index: 2147483647;
        transform: translate3d(0, 0, 0);
      }
      #__sikuli-demo-cursor .__sikuli-demo-cursor-pulse {
        position: absolute;
        left: 2px;
        top: 2px;
        width: 26px;
        height: 26px;
        border-radius: 999px;
        border: 2px solid rgba(56, 189, 248, 0.95);
        background: rgba(14, 165, 233, 0.16);
        opacity: 0;
        transform: scale(0.55);
        transform-origin: center;
      }
      #__sikuli-demo-cursor[data-clicking="1"] .__sikuli-demo-cursor-pulse {
        opacity: 1;
        transform: scale(1);
      }
      #__sikuli-demo-cursor .__sikuli-demo-cursor-pointer {
        position: absolute;
        left: 0;
        top: 0;
        width: 32px;
        height: 48px;
        filter: drop-shadow(0 10px 14px rgba(15, 23, 42, 0.45));
      }
    `;
    document.head.append(style);

    const cursor = document.createElement("div");
    cursor.id = "__sikuli-demo-cursor";
    cursor.dataset.clicking = "0";
    cursor.innerHTML = `<div class="__sikuli-demo-cursor-pulse"></div>${svgMarkup}`;
    document.body.append(cursor);

    window.__setSikuliDemoCursor = ({ x, y, clicking }) => {
      cursor.style.transform = `translate3d(${x}px, ${y}px, 0)`;
      cursor.dataset.clicking = clicking ? "1" : "0";
    };
  }, pointerSvg);
}

async function updateCursor(page, frame) {
  await page.evaluate((cursorFrame) => {
    if (typeof window.__setSikuliDemoCursor === "function") {
      window.__setSikuliDemoCursor(cursorFrame);
    }
  }, frame);
}

const browser = await chromium.launch({
  executablePath: chromePath,
  headless: true,
});

const page = await browser.newPage({
  viewport: {
    width: 1600,
    height: 1200,
  },
  deviceScaleFactor: 1,
});

try {
  await page.goto(editorUrl, { waitUntil: "domcontentloaded" });
  await page.waitForSelector("text=Editor walkthrough for capture output", { timeout: 30000 });
  await page.waitForTimeout(800);
  await installCursorOverlay(page);
  const cursorFrames = buildCursorFrames();

  for (let index = 0; index < frameCount; index += 1) {
    await page.waitForTimeout(index === 0 ? 80 : frameDelayMs);
    await updateCursor(page, cursorFrames[index]);
    await page.waitForTimeout(24);

    const framePath = path.join(framesDir, `frame-${String(index).padStart(3, "0")}.png`);
    await page.screenshot({
      path: framePath,
      type: "png",
    });
  }
} finally {
  await browser.close();
}

const img2webpArgs = ["-loop", "0", "-q", "80", "-m", "6"];
for (let index = 0; index < frameCount; index += 1) {
  img2webpArgs.push("-d", index === frameCount - 1 ? String(finalFramePlaybackMs) : String(framePlaybackMs));
  img2webpArgs.push(path.join(framesDir, `frame-${String(index).padStart(3, "0")}.png`));
}
img2webpArgs.push("-o", outputPath);

const img2webp = spawnSync("img2webp", img2webpArgs, {
  cwd: repoRoot,
  stdio: "inherit",
});

if (img2webp.status !== 0) {
  process.exit(img2webp.status || 1);
}

console.log(`Generated ${outputPath}`);
