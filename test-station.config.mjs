import path from "node:path";

const rootDir = import.meta.dirname;

export default {
  schemaVersion: "1",
  project: {
    name: "sikuli-go monorepo",
    rootDir,
    outputDir: ".test-results/test-station",
    rawDir: ".test-results/test-station/raw",
  },
  workspaceDiscovery: {
    provider: "manual",
    packages: [
      "quality",
      "api",
      "client-node",
      "client-python",
      "api-electron",
      "editor",
      "client-lua",
    ],
  },
  execution: {
    continueOnError: true,
    defaultCoverage: false,
  },
  manifests: {
    classification: "./test-station.policy.json",
    coverageAttribution: "./test-station.policy.json",
    thresholds: "./test-station.policy.json",
  },
  render: {
    html: true,
    console: true,
    defaultView: "package",
    includeDetailedAnalysisToggle: true,
  },
  suites: [
    {
      id: "check-api-docs",
      label: "API docs gate",
      adapter: "shell-check",
      handler: "./scripts/test-station/shell-check-adapter.mjs",
      package: "quality",
      cwd: rootDir,
      command: ["bash", "./scripts/check-api-docs.sh"],
      module: "repo",
      theme: "quality",
    },
    {
      id: "check-docs-governance",
      label: "Docs governance gate",
      adapter: "shell-check",
      handler: "./scripts/test-station/shell-check-adapter.mjs",
      package: "quality",
      cwd: rootDir,
      command: ["bash", "./scripts/check-docs-governance.sh"],
      module: "repo",
      theme: "quality",
    },
    {
      id: "check-docs-links",
      label: "Docs link gate",
      adapter: "shell-check",
      handler: "./scripts/test-station/shell-check-adapter.mjs",
      package: "quality",
      cwd: rootDir,
      command: ["bash", "./scripts/check-docs-links.sh"],
      module: "repo",
      theme: "quality",
    },
    {
      id: "check-parity-gates",
      label: "Parity gate",
      adapter: "shell-check",
      handler: "./scripts/test-station/shell-check-adapter.mjs",
      package: "quality",
      cwd: rootDir,
      command: ["bash", "./scripts/check-parity-gates.sh"],
      module: "repo",
      theme: "quality",
    },
    {
      id: "check-grpc-stubs",
      label: "gRPC stubs gate",
      adapter: "shell-check",
      handler: "./scripts/test-station/shell-check-adapter.mjs",
      package: "quality",
      cwd: rootDir,
      command: ["bash", "./scripts/check-grpc-stubs.sh"],
      module: "repo",
      theme: "quality",
    },
    {
      id: "go-mod-tidy",
      label: "go mod tidy diff",
      adapter: "shell-check",
      handler: "./scripts/test-station/shell-check-adapter.mjs",
      package: "api",
      cwd: path.join(rootDir, "packages/api"),
      command: ["go", "mod", "tidy", "-diff"],
      module: "runtime",
      theme: "go",
    },
    {
      id: "go-vet",
      label: "go vet",
      adapter: "shell-check",
      handler: "./scripts/test-station/shell-check-adapter.mjs",
      package: "api",
      cwd: path.join(rootDir, "packages/api"),
      command: ["go", "vet", "./..."],
      module: "runtime",
      theme: "go",
    },
    {
      id: "go-tests",
      label: "Go package tests",
      adapter: "go-test",
      handler: "./scripts/test-station/go-test-adapter.mjs",
      package: "api",
      cwd: path.join(rootDir, "packages/api"),
      command: ["go", "test", "-race", "./..."],
      module: "runtime",
      theme: "go",
      coverage: {
        enabled: true
      }
    },
    {
      id: "client-node-build",
      label: "Node client build",
      adapter: "shell-check",
      handler: "./scripts/test-station/shell-check-adapter.mjs",
      package: "client-node",
      cwd: rootDir,
      command: ["bash", "-lc", "bash ./scripts/clients/generate-node-stubs.sh && ./node_modules/.bin/tsc -p packages/client-node/tsconfig.json"],
      module: "clients",
      theme: "nodejs",
    },
    {
      id: "client-node-tests",
      label: "Node client tests",
      adapter: "node-test",
      package: "client-node",
      cwd: path.join(rootDir, "packages/client-node"),
      command: [
        "node",
        "--test",
        "test/binary.test.mjs",
        "test/bootstrap.test.mjs",
        "test/search-semantics.test.mjs",
        "test/launcher.test.mjs"
      ],
      coverage: {
        enabled: true,
        mode: "same-run"
      }
    },
    {
      id: "client-python-tests",
      label: "Python client tests",
      adapter: "python-unittest",
      handler: "./scripts/test-station/python-unittest-adapter.mjs",
      package: "client-python",
      cwd: path.join(rootDir, "packages/client-python"),
      command: ["python3"],
      testDir: "packages/client-python/tests",
      testPattern: "test_*.py",
      module: "clients",
      theme: "python",
      coverage: {
        enabled: true
      }
    },
    {
      id: "api-electron-tests",
      label: "Electron host tests",
      adapter: "node-test",
      package: "api-electron",
      cwd: path.join(rootDir, "packages/api-electron"),
      command: [
        "node",
        "--test",
        "test/examples.test.js",
        "test/runtime-addresses.test.js",
        "test/instrumentation.test.js",
        "test/execution.test.js"
      ],
      coverage: {
        enabled: true,
        mode: "same-run"
      }
    }
  ],
};
