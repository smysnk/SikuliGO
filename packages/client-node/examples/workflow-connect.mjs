import { Screen, Pattern } from "@sikuligo/sikuligo";

// Connect-only workflow (requires sikuligo already running).
const screen = await Screen.connect();
try {
  const match = await screen.click(Pattern("assets/pattern.png").exact());
  console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
} finally {
  await screen.close();
}
