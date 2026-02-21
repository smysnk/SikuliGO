import { SikuliGrpcClient } from "../src/client";

async function main(): Promise<void> {
  const client = new SikuliGrpcClient();
  const appName = process.env.SIKULI_APP_NAME ?? "Calculator";

  try {
    await client.openApp({
      name: appName,
      args: []
    });

    const running = await client.isAppRunning({ name: appName });
    console.log("isAppRunning", JSON.stringify(running, null, 2));

    const windows = await client.listWindows({ name: appName });
    console.log("listWindows", JSON.stringify(windows, null, 2));

    await client.focusApp({ name: appName });
    await client.closeApp({ name: appName });
    console.log("app control actions sent");
  } finally {
    client.close();
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
