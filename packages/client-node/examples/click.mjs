import { Screen, Pattern } from "@sikuligo/sikuligo";

const screen = await Screen();
try {
  const pattern = Pattern("assets/pattern.png").exact();
  const match = await screen.click(pattern);
  console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
} finally {
  await screen.close();
}
