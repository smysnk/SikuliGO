const fs = require("fs");
const path = require("path");

const REPO_ROOT = path.resolve(__dirname, "../../..");

const EXAMPLE_SOURCES = [
  {
    language: "nodejs",
    label: "Node.js",
    packageRoot: path.join(REPO_ROOT, "packages", "client-node"),
    examplesDir: path.join(REPO_ROOT, "packages", "client-node", "examples"),
    extension: ".mjs",
    ignoredFiles: new Set(["bootstrap.mjs"]),
  },
  {
    language: "python",
    label: "Python",
    packageRoot: path.join(REPO_ROOT, "packages", "client-python"),
    examplesDir: path.join(REPO_ROOT, "packages", "client-python", "examples"),
    extension: ".py",
    ignoredFiles: new Set(["bootstrap.py"]),
  },
];

const TEXT_EXTENSIONS = new Set([
  ".cjs",
  ".css",
  ".csv",
  ".html",
  ".js",
  ".json",
  ".md",
  ".mjs",
  ".py",
  ".sh",
  ".svg",
  ".ts",
  ".tsx",
  ".txt",
  ".xml",
  ".yaml",
  ".yml",
]);
const IMAGE_MIME_TYPES = new Map([
  [".gif", "image/gif"],
  [".jpeg", "image/jpeg"],
  [".jpg", "image/jpeg"],
  [".png", "image/png"],
  [".svg", "image/svg+xml"],
  [".webp", "image/webp"],
]);

const DEFAULT_SCRIPT_TEMPLATES = {
  nodejs: [
    'import { Screen, Pattern } from "@sikuligo/sikuli-go";',
    "",
    "const screen = await Screen();",
    "try {",
    '  const pattern = Pattern("assets/pattern.png").exact();',
    "  const match = await screen.find(pattern);",
    "  console.log(match);",
    "} finally {",
    "  await screen.close();",
    "}",
    "",
  ].join("\n"),
  python: [
    "from sikuligo import Pattern, Screen",
    "",
    "",
    "def main() -> int:",
    "    screen = Screen()",
    "    try:",
    '        match = screen.find(Pattern("assets/pattern.png").exact())',
    "        print(match)",
    "    finally:",
    "        screen.close()",
    "    return 0",
    "",
    "",
    'if __name__ == "__main__":',
    "    raise SystemExit(main())",
    "",
  ].join("\n"),
};

function titleCaseFromFileName(fileName) {
  return fileName
    .replace(/\.[^.]+$/, "")
    .split(/[-_]+/)
    .filter(Boolean)
    .map((part) => part.charAt(0).toUpperCase() + part.slice(1))
    .join(" ");
}

function commandPreview(config, relativePath) {
  if (config.language === "nodejs") {
    return `node ${relativePath}`;
  }
  return `python ${relativePath}`;
}

function sortByRelativePath(left, right) {
  return left.relativePath.localeCompare(right.relativePath);
}

function sanitizeFileName(rawName, extension) {
  const trimmed = String(rawName || "").trim();
  if (!trimmed) {
    throw new Error("File name is required");
  }
  const collapsed = trimmed.replace(/\s+/g, "-").replace(/[^A-Za-z0-9._-]/g, "-");
  if (!collapsed || collapsed === "." || collapsed === "..") {
    throw new Error("File name must contain letters, numbers, dashes, underscores, or dots");
  }
  if (collapsed.includes("/") || collapsed.includes("\\")) {
    throw new Error("File name must not contain path separators");
  }
  const withExtension = collapsed.endsWith(extension) ? collapsed : `${collapsed}${extension}`;
  if (withExtension.startsWith(".")) {
    throw new Error("File name must not start with a dot");
  }
  return withExtension;
}

function ensureInside(rootDir, absolutePath) {
  const normalizedRoot = path.resolve(rootDir);
  const normalizedPath = path.resolve(absolutePath);
  const relativePath = path.relative(normalizedRoot, normalizedPath);
  if (!relativePath || relativePath === "") {
    return normalizedPath;
  }
  if (relativePath.startsWith("..") || path.isAbsolute(relativePath)) {
    throw new Error(`Path escapes workspace root: ${absolutePath}`);
  }
  return normalizedPath;
}

function isIgnoredWorkspaceEntry(config, name) {
  if (!name || name === ".DS_Store") {
    return true;
  }
  if (name === "__pycache__") {
    return true;
  }
  if (name.startsWith(".")) {
    return !name.includes(".sikuli-ide-");
  }
  return config.ignoredFiles.has(name);
}

function shouldSkipWorkspaceFile(config, name) {
  if (isIgnoredWorkspaceEntry(config, name)) {
    return true;
  }
  if (name.includes(".sikuli-ide-")) {
    return true;
  }
  if (name.endsWith(".pyc")) {
    return true;
  }
  return false;
}

function buildExampleRecord(config, fileName) {
  const relativePath = path.posix.join("examples", fileName);
  const absolutePath = ensureInside(config.examplesDir, path.join(config.examplesDir, fileName));
  const source = fs.readFileSync(absolutePath, "utf8");
  return {
    id: `${config.language}:${fileName}`,
    language: config.language,
    languageLabel: config.label,
    name: titleCaseFromFileName(fileName),
    fileName,
    relativePath,
    absolutePath,
    packageRoot: config.packageRoot,
    commandPreview: commandPreview(config, relativePath),
    source,
  };
}

function buildWorkspaceFileRecord(config, absolutePath) {
  const resolvedPath = ensureInside(config.examplesDir, absolutePath);
  const relativeWorkspacePath = path.relative(config.examplesDir, resolvedPath).split(path.sep).join(path.posix.sep);
  const relativePath = path.posix.join("examples", relativeWorkspacePath);
  const extension = path.extname(relativeWorkspacePath).toLowerCase();
  const stat = fs.statSync(resolvedPath);
  return {
    id: `${config.language}:${relativeWorkspacePath}`,
    language: config.language,
    languageLabel: config.label,
    fileName: path.basename(relativeWorkspacePath),
    relativeWorkspacePath,
    relativePath,
    absolutePath: resolvedPath,
    packageRoot: config.packageRoot,
    workspaceRoot: config.examplesDir,
    kind: extension === config.extension ? "script" : "asset",
    extension,
    isImage: IMAGE_MIME_TYPES.has(extension),
    isText: TEXT_EXTENSIONS.has(extension),
    size: stat.size,
    updatedAt: stat.mtime.toISOString(),
  };
}

function readDataUrl(filePath, extension) {
  const mimeType = IMAGE_MIME_TYPES.get(extension);
  if (!mimeType) {
    return "";
  }
  const payload = fs.readFileSync(filePath);
  return `data:${mimeType};base64,${payload.toString("base64")}`;
}

function listWorkspaceEntries(config) {
  if (!fs.existsSync(config.examplesDir)) {
    return [];
  }

  const entries = [];
  const queue = [config.examplesDir];

  while (queue.length > 0) {
    const currentDir = queue.shift();
    if (!currentDir) {
      continue;
    }
    const directoryEntries = fs.readdirSync(currentDir, { withFileTypes: true });
    for (const entry of directoryEntries) {
      if (entry.isDirectory()) {
        if (isIgnoredWorkspaceEntry(config, entry.name)) {
          continue;
        }
        queue.push(path.join(currentDir, entry.name));
        continue;
      }
      if (!entry.isFile() || shouldSkipWorkspaceFile(config, entry.name)) {
        continue;
      }
      entries.push(buildWorkspaceFileRecord(config, path.join(currentDir, entry.name)));
    }
  }

  entries.sort(sortByRelativePath);
  return entries;
}

function fileNameFromExampleId(exampleId) {
  const [, fileName] = String(exampleId || "").split(":");
  return fileName || "";
}

function defaultSourceForLanguage(language) {
  return DEFAULT_SCRIPT_TEMPLATES[language] || "";
}

function createExampleStore(options = {}) {
  const sources = options.sources || EXAMPLE_SOURCES;

  function configForLanguage(language) {
    const config = sources.find((candidate) => candidate.language === language);
    if (!config) {
      throw new Error(`Unknown language: ${language}`);
    }
    return config;
  }

  function configForExampleId(exampleId) {
    const [language] = String(exampleId || "").split(":");
    return configForLanguage(language);
  }

  function listExamples() {
    const examples = [];
    for (const config of sources) {
      if (!fs.existsSync(config.examplesDir)) {
        continue;
      }
      const files = fs
        .readdirSync(config.examplesDir)
        .filter((fileName) => fileName.endsWith(config.extension))
        .filter((fileName) => !shouldSkipWorkspaceFile(config, fileName))
        .sort((left, right) => left.localeCompare(right));
      for (const fileName of files) {
        examples.push(buildExampleRecord(config, fileName));
      }
    }
    return examples;
  }

  function getExampleById(exampleId) {
    return listExamples().find((example) => example.id === exampleId) || null;
  }

  function readExample(exampleId) {
    const example = getExampleById(exampleId);
    if (!example) {
      throw new Error(`Unknown example: ${exampleId}`);
    }
    return example;
  }

  function saveExample(exampleId, source) {
    const example = readExample(exampleId);
    fs.writeFileSync(example.absolutePath, String(source), "utf8");
    return readExample(exampleId);
  }

  function createExample(language, fileName, source) {
    const config = configForLanguage(language);
    const normalizedFileName = sanitizeFileName(fileName, config.extension);
    const absolutePath = ensureInside(config.examplesDir, path.join(config.examplesDir, normalizedFileName));
    if (fs.existsSync(absolutePath)) {
      throw new Error(`File already exists: ${normalizedFileName}`);
    }
    fs.mkdirSync(config.examplesDir, { recursive: true });
    fs.writeFileSync(absolutePath, source ?? defaultSourceForLanguage(language), "utf8");
    return readExample(`${language}:${normalizedFileName}`);
  }

  function cloneExample(exampleId, fileName, source) {
    const example = readExample(exampleId);
    return createExample(example.language, fileName, source ?? example.source);
  }

  function renameExample(exampleId, fileName, source) {
    const example = readExample(exampleId);
    const config = configForExampleId(exampleId);
    const normalizedFileName = sanitizeFileName(fileName, config.extension);
    if (normalizedFileName === example.fileName) {
      if (typeof source === "string") {
        return saveExample(exampleId, source);
      }
      return example;
    }
    const nextAbsolutePath = ensureInside(config.examplesDir, path.join(config.examplesDir, normalizedFileName));
    if (fs.existsSync(nextAbsolutePath)) {
      throw new Error(`File already exists: ${normalizedFileName}`);
    }
    fs.renameSync(example.absolutePath, nextAbsolutePath);
    const nextId = `${example.language}:${normalizedFileName}`;
    if (typeof source === "string") {
      fs.writeFileSync(nextAbsolutePath, source, "utf8");
    }
    return readExample(nextId);
  }

  function listWorkspaceFiles(exampleId) {
    const config = configForExampleId(exampleId);
    return listWorkspaceEntries(config);
  }

  function readWorkspaceFile(exampleId, relativeWorkspacePath) {
    const config = configForExampleId(exampleId);
    const normalizedRelativePath = String(relativeWorkspacePath || "").replace(/\\/g, "/").replace(/^\/+/, "");
    if (!normalizedRelativePath) {
      throw new Error("Workspace file path is required");
    }
    const absolutePath = ensureInside(config.examplesDir, path.join(config.examplesDir, normalizedRelativePath));
    if (!fs.existsSync(absolutePath) || !fs.statSync(absolutePath).isFile()) {
      throw new Error(`Unknown workspace file: ${normalizedRelativePath}`);
    }

    const record = buildWorkspaceFileRecord(config, absolutePath);
    if (record.isText) {
      return {
        ...record,
        content: fs.readFileSync(absolutePath, "utf8"),
        dataUrl: "",
      };
    }
    if (record.isImage) {
      return {
        ...record,
        content: "",
        dataUrl: readDataUrl(absolutePath, record.extension),
      };
    }
    return {
      ...record,
      content: "",
      dataUrl: "",
    };
  }

  return {
    configForLanguage,
    createExample,
    cloneExample,
    fileNameFromExampleId,
    getExampleById,
    listExamples,
    listWorkspaceFiles,
    readExample,
    readWorkspaceFile,
    renameExample,
    saveExample,
  };
}

const defaultStore = createExampleStore();

module.exports = {
  DEFAULT_SCRIPT_TEMPLATES,
  EXAMPLE_SOURCES,
  REPO_ROOT,
  createExampleStore,
  createExample: defaultStore.createExample,
  cloneExample: defaultStore.cloneExample,
  fileNameFromExampleId: defaultStore.fileNameFromExampleId,
  getExampleById: defaultStore.getExampleById,
  listExamples: defaultStore.listExamples,
  listWorkspaceFiles: defaultStore.listWorkspaceFiles,
  readExample: defaultStore.readExample,
  readWorkspaceFile: defaultStore.readWorkspaceFile,
  renameExample: defaultStore.renameExample,
  saveExample: defaultStore.saveExample,
};
