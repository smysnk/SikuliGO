import { Screen, Pattern } from "../src";

async function main() {
  const screen = await Screen.auto();
  try {
    const pattern = new Pattern("assets/pattern.png").exact();
    const match = await screen.click(pattern);
    console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
  } finally {
    await screen.close();
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
