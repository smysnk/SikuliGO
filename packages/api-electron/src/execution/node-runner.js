const { pathToFileURL } = require("node:url");

const instrumentedPath = process.argv[2];

const state = {
  currentStep: null,
  pauseBeforeNext: process.env.SIKULI_GO_STEP_START_PAUSED === "1",
  pauseBeforeRuntimeCall: false,
  runToLine: null,
  pendingResume: null,
};

function sendEvent(payload) {
  if (typeof process.send === "function") {
    process.send(payload);
  }
}

function normalizeStepPayload(payload) {
  return {
    line: payload?.line ?? null,
    column: payload?.column ?? null,
    statementId: payload?.statementId ?? null,
  };
}

function setCurrentStep(payload) {
  state.currentStep = normalizeStepPayload(payload);
  globalThis.__sikuliIdeCurrentStep = state.currentStep;
}

function clearCurrentStep() {
  state.currentStep = null;
  globalThis.__sikuliIdeCurrentStep = null;
}

function finishCurrentStep() {
  if (!state.currentStep) {
    return;
  }
  sendEvent({
    type: "step:end",
    ...state.currentStep,
  });
  clearCurrentStep();
}

async function pauseAtCurrentStep(reason, extras = {}) {
  const currentStep = state.currentStep;
  if (!currentStep) {
    return;
  }

  sendEvent({
    type: "step:pause",
    reason,
    ...extras,
    ...currentStep,
  });

  await new Promise((resolve) => {
    state.pendingResume = {
      resolve,
      step: currentStep,
    };
  });
}

function releasePendingResume(mode, extras = {}) {
  if (!state.pendingResume) {
    return false;
  }

  const pending = state.pendingResume;
  state.pendingResume = null;
  sendEvent({
    type: "step:resume",
    mode,
    ...extras,
    ...pending.step,
  });
  pending.resolve();
  return true;
}

globalThis.__sikuliStep = async (payload) => {
  finishCurrentStep();
  setCurrentStep(payload);

  sendEvent({
    type: "step:start",
    ...state.currentStep,
  });

  if (state.runToLine !== null && state.currentStep?.line === state.runToLine) {
    const targetLine = state.runToLine;
    state.runToLine = null;
    await pauseAtCurrentStep("run-to-line", {
      targetLine,
    });
    return;
  }

  if (state.pauseBeforeNext) {
    state.pauseBeforeNext = false;
    await pauseAtCurrentStep("step");
  }
};

globalThis.__sikuliRuntimeEvent = async (payload) => {
  if (!payload || typeof payload !== "object") {
    return;
  }

  const eventPayload = {
    ...payload,
    ...normalizeStepPayload(state.currentStep),
  };
  sendEvent(eventPayload);

  if (payload.type === "runtime:call:start" && state.pauseBeforeRuntimeCall) {
    state.pauseBeforeRuntimeCall = false;
    await pauseAtCurrentStep("runtime-call", {
      runtimeMethod: payload.method ?? null,
      runtimeCallId: payload.callId ?? null,
    });
  }
};

process.on("message", (message) => {
  if (!message || typeof message !== "object") {
    return;
  }

  if (message.command === "pause") {
    if (!state.pendingResume) {
      state.pauseBeforeNext = true;
      state.pauseBeforeRuntimeCall = false;
      state.runToLine = null;
    }
    return;
  }

  if (message.command === "resume") {
    state.pauseBeforeNext = false;
    state.pauseBeforeRuntimeCall = false;
    state.runToLine = null;
    releasePendingResume("resume");
    return;
  }

  if (message.command === "step") {
    state.pauseBeforeNext = true;
    state.pauseBeforeRuntimeCall = false;
    state.runToLine = null;
    releasePendingResume("step");
    return;
  }

  if (message.command === "run-to-line") {
    const targetLine = Number(message.line);
    if (!Number.isFinite(targetLine) || targetLine < 1) {
      return;
    }
    state.runToLine = Math.floor(targetLine);
    state.pauseBeforeNext = false;
    state.pauseBeforeRuntimeCall = false;
    releasePendingResume("run-to-line", {
      targetLine: state.runToLine,
    });
    return;
  }

  if (message.command === "continue-to-runtime-call") {
    state.pauseBeforeRuntimeCall = true;
    state.pauseBeforeNext = false;
    state.runToLine = null;
    releasePendingResume("runtime-call");
  }
});

(async () => {
  try {
    await import(pathToFileURL(instrumentedPath).href);
    finishCurrentStep();
    if (process.connected) {
      process.disconnect();
    }
  } catch (error) {
    if (state.currentStep) {
      sendEvent({
        type: "step:error",
        error: error instanceof Error ? error.message : String(error),
        ...state.currentStep,
      });
    }
    console.error(error);
    if (process.connected) {
      process.disconnect();
    }
    process.exitCode = 1;
  }
})();
