import { Screen, Pattern } from "../src";

async function main() {
  // Workflow 2: user starts sikuligo manually, client connects to it.
  const screen = await Screen.connect({
    address: process.env.SIKULI_GRPC_ADDR ?? "127.0.0.1:50051"
  });
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
