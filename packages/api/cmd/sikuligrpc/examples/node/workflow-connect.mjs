import { ensureSikuliGoOnPath } from "./bootstrap.mjs";
import { Screen, Pattern } from "@sikuligo/sikuli-go";

ensureSikuliGoOnPath();

// Connect-only workflow (requires sikuli-go already running).
const screen = await Screen.connect();
try {
  const match = await screen.click(Pattern("assets/pattern.png").exact());
  console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
} finally {
  await screen.close();
}
