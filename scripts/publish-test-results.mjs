#!/usr/bin/env node

import fs from "node:fs";
import path from "node:path";

const ROOT = process.cwd();
const SUMMARY_PATH = path.join(ROOT, ".test-results", "summary.json");

const webhookUrl =
  process.env.DISCORD_TEST_WEBHOOK_URL || process.env.DISCORD_WEBHOOK_URL || "";

if (!webhookUrl) {
  console.log("Skipping Discord publish: no webhook URL provided.");
  process.exit(0);
}

const githubRepository = process.env.GITHUB_REPOSITORY || null;
const githubRunId = process.env.GITHUB_RUN_ID || null;
const githubRefName =
  process.env.GITHUB_HEAD_REF || process.env.GITHUB_REF_NAME || process.env.GITHUB_REF || "unknown";
const githubSha = process.env.GITHUB_SHA || "unknown";
const runUrl =
  githubRepository && githubRunId
    ? `https://github.com/${githubRepository}/actions/runs/${githubRunId}`
    : null;

let report = null;
if (fs.existsSync(SUMMARY_PATH)) {
  try {
    report = JSON.parse(fs.readFileSync(SUMMARY_PATH, "utf8"));
  } catch (error) {
    console.error("Failed to parse workspace test summary JSON:", error);
  }
}

const summary = report?.summary || {};
const packages = Array.isArray(report?.packages) ? report.packages : [];
const totalTests = Number(summary.tests || 0);
const totalPassed = Number(summary.passed || 0);
const totalFailed = Number(summary.failed || 0);
const totalSkipped = Number(summary.skipped || 0);
const durationMs = Number(summary.durationMs || 0);
const packageCount = Number(summary.packageCount || packages.length || 0);
const failedPackages = Number(
  summary.failedPackages || packages.filter((item) => item?.status === "failed").length || 0,
);

const hasReport = Boolean(report);
const overallStatus = hasReport && failedPackages === 0 && totalFailed === 0 ? "passed" : "failed";
const statusText = overallStatus === "passed" ? "PASSED" : "FAILED";
const statusIcon = overallStatus === "passed" ? "✅" : "❌";
const color = overallStatus === "passed" ? 0x57f287 : 0xed4245;

const packageLines = packages.map((item) => {
  const isPassed = String(item?.status || "").toLowerCase() === "passed";
  const icon = isPassed ? "✅" : "❌";
  const name = item?.package || "unknown";
  const tests = Number(item?.tests || 0);
  const passed = Number(item?.passed || 0);
  const failed = Number(item?.failed || 0);
  const skipped = Number(item?.skipped || 0);
  const seconds = (Number(item?.durationMs || 0) / 1000).toFixed(2);
  return `${icon} **${name}** • tests: ${tests} • pass: ${passed} • fail: ${failed} • skip: ${skipped} • ${seconds}s`;
});

const descriptionLines = [
  `**Result:** ${statusIcon} ${statusText}`,
  `**Branch:** \`${githubRefName}\``,
  `**Commit:** \`${String(githubSha).slice(0, 12)}\``,
  `**Packages:** ${packageCount} (failed: ${failedPackages})`,
  `**Total Tests:** ${totalTests} (pass: ${totalPassed}, fail: ${totalFailed}, skip: ${totalSkipped})`,
  `**Duration:** ${(durationMs / 1000).toFixed(2)}s`,
];

if (!hasReport) {
  descriptionLines.push("", `No report found at \`${path.relative(ROOT, SUMMARY_PATH)}\`.`);
}

if (packageLines.length > 0) {
  descriptionLines.push("", "**Per Package**", ...packageLines);
}

if (runUrl) {
  descriptionLines.push("", `[View workflow run](${runUrl})`);
}

const payload = {
  username: "SikuliGO CI",
  embeds: [
    {
      title: "Workspace Test Results",
      description: descriptionLines.join("\n"),
      color,
      timestamp: new Date().toISOString(),
    },
  ],
};

const response = await fetch(webhookUrl, {
  method: "POST",
  headers: {
    "content-type": "application/json",
  },
  body: JSON.stringify(payload),
});

if (!response.ok) {
  const text = await response.text();
  throw new Error(`Discord webhook request failed (${response.status}): ${text}`);
}

console.log("Published workspace test summary to Discord webhook.");
