import { ensureSikuligoOnPath } from "./bootstrap.mjs";
import { Screen, Pattern } from "@sikuligo/sikuligo";

ensureSikuligoOnPath();

const screen = await Screen();
try {
  const pattern = Pattern("assets/pattern.png").similar(0.9)
  const match = await screen.click(pattern);
  console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
} finally {
  await screen.close();
}
