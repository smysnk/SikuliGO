import { Sikuli } from "../src";

async function main(): Promise<void> {
  const bot = await Sikuli.launch({
    startupTimeoutMs: 10_000
  });

  try {
    await bot.click({ x: 300, y: 220 });
    await bot.typeText("hello from sikuligo");
    await bot.hotkey(["cmd", "enter"]);
    console.log("automation actions sent");
  } finally {
    await bot.close();
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
