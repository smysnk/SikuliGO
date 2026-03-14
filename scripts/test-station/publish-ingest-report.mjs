#!/usr/bin/env node

import {
  createIngestPayload,
  normalizeStorageOptions,
  publishIngestPayload,
} from "./ingest-report-utils.mjs";

async function main() {
  const reportPath = process.env.TEST_STATION_INGEST_INPUT || "./.test-results/test-station/report.json";
  const endpoint = process.env.TEST_STATION_INGEST_ENDPOINT || "https://test-station.smysnk.com/api/ingest";
  const projectKey = process.env.TEST_STATION_INGEST_PROJECT_KEY || "sikuli-go";
  const sharedKey = process.env.TEST_STATION_INGEST_SHARED_KEY || "";

  if (!sharedKey.trim()) {
    process.stdout.write("Skipping test-station ingest: no shared key provided.\n");
    return;
  }

  const payload = createIngestPayload({
    reportPath,
    projectKey,
    buildStartedAt: process.env.TEST_STATION_BUILD_STARTED_AT,
    buildCompletedAt: process.env.TEST_STATION_BUILD_COMPLETED_AT,
    jobStatus: process.env.TEST_STATION_CI_STATUS,
    storage: normalizeStorageOptions({
      bucket: process.env.S3_BUCKET,
      prefix: process.env.S3_STORAGE_PREFIX,
      baseUrl: process.env.S3_PUBLIC_URL,
    }),
  });

  const response = await publishIngestPayload({
    endpoint,
    sharedKey,
    payload,
  });

  process.stdout.write(`Published ${payload.projectKey}:${payload.source.provider}:${payload.source.runId || "manual"} to ${endpoint}\n`);
  if (response?.runId) {
    process.stdout.write(`runId=${response.runId}\n`);
  }
}

main().catch((error) => {
  process.stderr.write(`${error instanceof Error ? error.stack || error.message : String(error)}\n`);
  process.exit(1);
});
