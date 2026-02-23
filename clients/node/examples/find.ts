import { Screen, Pattern } from "../src";

async function main(): Promise<void> {
  const screen = await Screen.start();
  try {
    const pattern = new Pattern("assets/pattern.png").exact();
    const match = await screen.find(pattern);
    console.log(`found rect=(${match.x},${match.y},${match.w},${match.h}) target=(${match.targetX},${match.targetY})`);
  } finally {
    await screen.close();
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
