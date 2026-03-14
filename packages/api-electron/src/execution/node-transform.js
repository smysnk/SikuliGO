const fs = require("fs");
const path = require("path");

function createStripState() {
  return {
    inBlockComment: false,
    inSingleQuote: false,
    inDoubleQuote: false,
    inTemplate: false,
  };
}

function stripLine(line, state) {
  let output = "";

  for (let index = 0; index < line.length; index += 1) {
    const current = line[index];
    const next = line[index + 1];

    if (state.inBlockComment) {
      if (current === "*" && next === "/") {
        state.inBlockComment = false;
        output += "  ";
        index += 1;
      } else {
        output += " ";
      }
      continue;
    }

    if (state.inSingleQuote) {
      output += " ";
      if (current === "\\" && next) {
        output += " ";
        index += 1;
        continue;
      }
      if (current === "'") {
        state.inSingleQuote = false;
      }
      continue;
    }

    if (state.inDoubleQuote) {
      output += " ";
      if (current === "\\" && next) {
        output += " ";
        index += 1;
        continue;
      }
      if (current === "\"") {
        state.inDoubleQuote = false;
      }
      continue;
    }

    if (state.inTemplate) {
      output += " ";
      if (current === "\\" && next) {
        output += " ";
        index += 1;
        continue;
      }
      if (current === "`") {
        state.inTemplate = false;
      }
      continue;
    }

    if (current === "/" && next === "/") {
      output += " ".repeat(line.length - index);
      break;
    }

    if (current === "/" && next === "*") {
      state.inBlockComment = true;
      output += "  ";
      index += 1;
      continue;
    }

    if (current === "'") {
      state.inSingleQuote = true;
      output += " ";
      continue;
    }

    if (current === "\"") {
      state.inDoubleQuote = true;
      output += " ";
      continue;
    }

    if (current === "`") {
      state.inTemplate = true;
      output += " ";
      continue;
    }

    output += current;
  }

  return output;
}

function countDelimiters(line) {
  let paren = 0;
  let bracket = 0;
  let brace = 0;

  for (const character of line) {
    if (character === "(") {
      paren += 1;
    } else if (character === ")") {
      paren -= 1;
    } else if (character === "[") {
      bracket += 1;
    } else if (character === "]") {
      bracket -= 1;
    } else if (character === "{") {
      brace += 1;
    } else if (character === "}") {
      brace -= 1;
    }
  }

  return { paren, bracket, brace };
}

function isFunctionDeclaration(trimmedLine) {
  return /^(export\s+)?(async\s+)?function\b/.test(trimmedLine);
}

function isBlockHeader(trimmedLine) {
  return /^(async\s+function\b|function\b|class\b|if\b|for\b|while\b|switch\b|try\b|catch\b|finally\b|do\b)\b/.test(trimmedLine);
}

function shouldInstrumentLine(trimmedLine) {
  if (!trimmedLine) {
    return false;
  }
  if (/^(import|export)\b/.test(trimmedLine)) {
    return false;
  }
  if (/^(case\b|default:)/.test(trimmedLine)) {
    return false;
  }
  if (trimmedLine === "{" || trimmedLine.startsWith("}")) {
    return false;
  }
  return true;
}

function instrumentNodeSource(source) {
  const lines = source.split("\n");
  const stripState = createStripState();
  let continuationDepth = 0;
  let functionBraceBalance = 0;
  let statementCount = 0;

  const transformed = lines.map((line, lineIndex) => {
    const stripped = stripLine(line, stripState);
    const trimmed = stripped.trim();
    const counts = countDelimiters(stripped);
    const insideFunction = functionBraceBalance > 0;
    const startsFunction = !insideFunction && isFunctionDeclaration(trimmed);
    const canInstrument =
      !insideFunction &&
      !startsFunction &&
      continuationDepth === 0 &&
      shouldInstrumentLine(trimmed);

    let nextLine = line;

    if (canInstrument) {
      const indent = line.match(/^\s*/)?.[0] || "";
      const content = line.slice(indent.length);
      const statementId = `node-${lineIndex + 1}-${statementCount + 1}`;
      const payload = JSON.stringify({
        line: lineIndex + 1,
        column: indent.length + 1,
        statementId,
      });
      nextLine = `${indent}await globalThis.__sikuliStep(${payload}); ${content}`;
      statementCount += 1;
    }

    if (!insideFunction && continuationDepth > 0) {
      continuationDepth = Math.max(0, continuationDepth + counts.paren + counts.bracket + counts.brace);
    } else if (!insideFunction && canInstrument && !isBlockHeader(trimmed)) {
      continuationDepth = Math.max(0, counts.paren + counts.bracket + counts.brace);
    }

    if (insideFunction) {
      functionBraceBalance += counts.brace;
    } else if (startsFunction) {
      functionBraceBalance = counts.brace;
    }

    return nextLine;
  });

  return {
    source: transformed.join("\n"),
    statementCount,
  };
}

function createInstrumentedNodeExample(example, sessionId) {
  const transformed = instrumentNodeSource(example.source);
  const extension = path.extname(example.absolutePath) || ".mjs";
  const baseName = path.basename(example.absolutePath, extension);
  const instrumentedPath = path.join(
    path.dirname(example.absolutePath),
    `.${baseName}.sikuli-ide-${sessionId}${extension}`,
  );

  fs.writeFileSync(instrumentedPath, transformed.source, "utf8");

  return {
    instrumentedPath,
    statementCount: transformed.statementCount,
  };
}

module.exports = {
  createInstrumentedNodeExample,
  instrumentNodeSource,
};
