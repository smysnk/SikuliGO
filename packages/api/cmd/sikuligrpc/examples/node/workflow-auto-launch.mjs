import { ensureSikuliGoOnPath } from "./bootstrap.mjs";
import { Screen, Pattern } from "@sikuligo/sikuli-go";

ensureSikuliGoOnPath();

// Primary constructor: connect first, then spawn fallback.
const screen = await Screen();
try {
  const match = await screen.click(Pattern("assets/pattern.png").exact());
  console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
} finally {
  await screen.close();
}
