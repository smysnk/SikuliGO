import { SikuliGrpcClient } from "../src/client";

function grayImageFromRows(name: string, rows: number[][]) {
  const height = rows.length;
  const width = rows[0].length;
  const pix = rows.flat().map((v) => v & 0xff);
  return {
    name,
    width,
    height,
    pix: Buffer.from(pix)
  };
}

async function main(): Promise<void> {
  const client = new SikuliGrpcClient();
  const source = grayImageFromRows("source", [
    [10, 10, 10, 10, 10, 10, 10, 10],
    [10, 0, 255, 10, 10, 10, 10, 10],
    [10, 255, 0, 10, 0, 255, 10, 10],
    [10, 10, 10, 10, 255, 0, 10, 10],
    [10, 10, 10, 10, 10, 10, 10, 10]
  ]);
  const needle = grayImageFromRows("needle", [
    [0, 255],
    [255, 0]
  ]);

  try {
    const response = await client.find({
      source,
      pattern: {
        image: needle,
        exact: true
      }
    });
    console.log(JSON.stringify(response, null, 2));
  } finally {
    client.close();
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
