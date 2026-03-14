import { applySuiteClassification, createSummary, parseCommandSpec, resolveSuiteEnv, slugify, spawnCommand, trimForReport } from "./common.mjs";

export default {
  id: "shell-check",
  description: "Single-check shell adapter",
  phase: 1,
  async run({ project, suite }) {
    const commandSpec = parseCommandSpec(suite.command);
    const execution = await spawnCommand(commandSpec.command, commandSpec.args, {
      cwd: suite.cwd || project.rootDir,
      env: resolveSuiteEnv(suite.env),
    });

    const combinedOutput = [execution.stdout, execution.stderr].filter(Boolean).join("\n");
    const passed = execution.exitCode === 0;

    return applySuiteClassification({
      status: passed ? "passed" : "failed",
      durationMs: execution.durationMs,
      summary: createSummary({
        total: 1,
        passed: passed ? 1 : 0,
        failed: passed ? 0 : 1,
        skipped: 0,
      }),
      coverage: null,
      tests: [
        {
          name: suite.label || suite.id || "shell check",
          fullName: `${suite.packageName || "default"} ${suite.label || suite.id || "shell check"}`,
          status: passed ? "passed" : "failed",
          durationMs: execution.durationMs,
          file: null,
          line: null,
          column: null,
          assertions: ["Command completed successfully."],
          setup: [],
          mocks: [],
          failureMessages: passed ? [] : [trimForReport(execution.stderr || execution.stdout || "Shell command failed.")],
          rawDetails: {
            command: [commandSpec.command, ...commandSpec.args].join(" "),
            stdout: trimForReport(execution.stdout, 2000),
            stderr: trimForReport(execution.stderr, 2000),
          },
        },
      ],
      warnings: [],
      output: {
        stdout: execution.stdout,
        stderr: execution.stderr,
      },
      rawArtifacts: [
        {
          relativePath: `${slugify(suite.packageName || "default")}-${slugify(suite.id || "shell-check")}.log`,
          content: combinedOutput,
        },
      ],
    }, suite);
  },
};
