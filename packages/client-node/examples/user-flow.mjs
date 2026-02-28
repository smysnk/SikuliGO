import { Screen, Pattern } from "@sikuligo/sikuligo";

const screen = await Screen({
  startupTimeoutMs: 10_000
});
try {
  const pattern = Pattern("assets/pattern.png").exact();
  const match = await screen.click(pattern);
  console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
  console.log("automation actions sent");
} finally {
  await screen.close();
}
