import test from "node:test";
import assert from "node:assert/strict";
import * as grpc from "@grpc/grpc-js";

import { Region } from "../dist/src/sikulix.js";
import { SikuliError } from "../dist/src/client.js";

const PNG_BYTES = Buffer.from(
  "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO3Z8V0AAAAASUVORK5CYII=",
  "base64"
);

function presentMatch() {
  return {
    rect: { x: 10, y: 20, w: 30, h: 40 },
    target: { x: 25, y: 40 },
    score: 0.95,
    index: 0
  };
}

test("Region.find preserves transport not-found errors", async () => {
  const err = new SikuliError(grpc.status.NOT_FOUND, "sikuli: find failed");
  const region = new Region({
    findOnScreen: async () => {
      throw err;
    }
  });
  await assert.rejects(region.find(PNG_BYTES), (actual) => actual === err);
});

test("Region.wait preserves transport deadline errors", async () => {
  const err = new SikuliError(grpc.status.DEADLINE_EXCEEDED, "deadline exceeded");
  const region = new Region({
    waitOnScreen: async () => {
      throw err;
    }
  });
  await assert.rejects(region.wait(PNG_BYTES, 25), (actual) => actual === err);
});

test("Region.waitVanish returns true immediately when target is already absent", async () => {
  let calls = 0;
  const region = new Region({
    existsOnScreen: async () => {
      calls += 1;
      return { exists: false };
    }
  });
  const vanished = await region.waitVanish(PNG_BYTES, 0);
  assert.equal(vanished, true);
  assert.equal(calls, 1);
});

test("Region.waitVanish retries until the target disappears", async () => {
  let calls = 0;
  const region = new Region({
    existsOnScreen: async () => {
      calls += 1;
      if (calls < 3) {
        return { exists: true, match: presentMatch() };
      }
      return { exists: false };
    }
  });
  const vanished = await region.waitVanish(PNG_BYTES, 50, 1);
  assert.equal(vanished, true);
  assert.equal(calls, 3);
});

test("Region.waitVanish returns false on timeout without throwing", async () => {
  const region = new Region({
    existsOnScreen: async () => ({ exists: true, match: presentMatch() })
  });
  const vanished = await region.waitVanish(PNG_BYTES, 5, 1);
  assert.equal(vanished, false);
});
