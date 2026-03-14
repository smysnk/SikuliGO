import path from "node:path";
import { fileURLToPath } from "node:url";

import { applySuiteClassification, parseCommandSpec, resolveSuiteEnv, spawnCommand } from "./common.mjs";

const helperPath = fileURLToPath(new URL("./python-unittest-runner.py", import.meta.url));

export default {
  id: "python-unittest",
  description: "Python unittest adapter",
  phase: 1,
  async run({ project, suite, execution }) {
    const pythonSpec = parseCommandSpec(suite.command);
    const args = [
      helperPath,
      "--package-root", path.resolve(suite.cwd || project.rootDir),
      "--test-dir", path.resolve(project.rootDir, suite.testDir || "packages/client-python/tests"),
      "--pattern", suite.testPattern || "test_*.py",
      "--python-executable", pythonSpec.command,
      ...(Boolean(execution?.coverage) && suite?.coverage?.enabled !== false ? ["--coverage"] : []),
    ];

    const result = await spawnCommand(pythonSpec.command, args, {
      cwd: suite.cwd || project.rootDir,
      env: resolveSuiteEnv(suite.env),
    });

    if (result.exitCode !== 0) {
      throw new Error(result.stderr || result.stdout || "Python unittest helper failed.");
    }

    const parsed = JSON.parse(result.stdout);
    return applySuiteClassification(parsed, suite);
  },
};
