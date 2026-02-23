import { Sikuli } from "../src";

async function main() {
  const client = await Sikuli.launch();
  try {
    await client.moveMouse({
      x: 200,
      y: 180,
      opts: { delayMillis: 30 }
    });
    await client.click({
      x: 200,
      y: 180,
      button: "left",
      delayMillis: 20
    });
    await client.typeText({
      text: "hello from node grpc",
      delayMillis: 15
    });
    await client.hotkey(["cmd", "a"]);
    console.log("input actions sent");
  } finally {
    await client.close();
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
