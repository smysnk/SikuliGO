import { Sikuli } from "../src";

async function main() {
  const client = await Sikuli.launch();
  const appName = process.env.SIKULI_APP_NAME ?? "Calculator";

  try {
    await client.openApp({
      name: appName,
      args: []
    });

    const running = await client.isAppRunning(appName);
    console.log("isAppRunning", JSON.stringify(running, null, 2));

    const windows = await client.listWindows(appName);
    console.log("listWindows", JSON.stringify(windows, null, 2));

    await client.focusApp(appName);
    await client.closeApp(appName);
    console.log("app control actions sent");
  } finally {
    await client.close();
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
