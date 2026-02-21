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
  const source = grayImageFromRows("ocr-source", [
    [220, 220, 220, 220],
    [220, 20, 20, 220],
    [220, 220, 220, 220]
  ]);

  try {
    const readText = await client.readText({
      source,
      params: {
        language: "eng"
      }
    });
    console.log("readText", JSON.stringify(readText, null, 2));

    const findText = await client.findText({
      source,
      query: "example",
      params: {
        language: "eng",
        case_sensitive: false
      }
    });
    console.log("findText", JSON.stringify(findText, null, 2));
  } finally {
    client.close();
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
