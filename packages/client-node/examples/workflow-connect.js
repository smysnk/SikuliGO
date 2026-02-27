import { Screen, Pattern } from "../src";

async function main() {
  // Auto-launch workflow.
  const screen = await Screen();
  try {
    const match = await screen.click(Pattern("assets/pattern.png").exact());
    console.log(`clicked match target at (${match.targetX}, ${match.targetY})`);
  } finally {
    await screen.close();
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
