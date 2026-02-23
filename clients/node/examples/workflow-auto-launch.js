import { Screen, Pattern } from "../src";

async function main() {
  // Workflow 1: client auto-launches sikuligo if needed.
  const screen = await Screen.start();
  try {
    const match = await screen.click(new Pattern("assets/pattern.png").exact());
    console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
  } finally {
    await screen.close();
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
