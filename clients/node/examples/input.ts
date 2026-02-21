import { SikuliGrpcClient } from "../src/client";

async function main(): Promise<void> {
  const client = new SikuliGrpcClient();
  try {
    await client.moveMouse({
      x: 200,
      y: 180,
      opts: { delay_millis: 30 }
    });
    await client.click({
      x: 200,
      y: 180,
      opts: { button: "left", delay_millis: 20 }
    });
    await client.typeText({
      text: "hello from node grpc",
      opts: { delay_millis: 15 }
    });
    await client.hotkey({ keys: ["cmd", "a"] });
    console.log("input actions sent");
  } finally {
    client.close();
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
