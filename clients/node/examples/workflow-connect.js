import { Screen, Pattern } from "../src";

async function main() {
  // Auto-launch workflow.
  const screen = await Screen.auto();
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
