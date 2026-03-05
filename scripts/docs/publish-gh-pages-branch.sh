#!/usr/bin/env bash
set -euo pipefail

THIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "${THIS_DIR}/../.." && pwd)"

REMOTE="${GH_PAGES_REMOTE:-origin}"
REMOTE_URL_OVERRIDE="${GH_PAGES_REMOTE_URL:-}"
BRANCH="${GH_PAGES_BRANCH:-gh-pages}"
BUILD_SITE="${GH_PAGES_BUILD:-1}"
FORCE_PUSH="${GH_PAGES_FORCE_PUSH:-1}"
INCLUDE_BENCH="${GH_PAGES_INCLUDE_BENCH:-1}"
BENCH_SOURCE="${GH_PAGES_BENCH_SOURCE:-${ROOT_DIR}/.test-results/bench}"
STAGE_DIR="${GH_PAGES_STAGE_DIR:-${ROOT_DIR}/.test-results/docs-pages-publish}"
SITE_DIR="${GH_PAGES_SITE_DIR:-${STAGE_DIR}/site}"
COMMIT_MESSAGE="${GH_PAGES_COMMIT_MESSAGE:-docs: publish GitHub Pages ($(date -u +%Y-%m-%dT%H:%M:%SZ))}"
CNAME_VALUE="${GH_PAGES_CNAME:-}"
CONFIGURE_SOURCE="${GH_PAGES_CONFIGURE_SOURCE:-0}"

is_true() {
  local raw="${1:-}"
  local norm
  norm="$(printf '%s' "${raw}" | tr '[:upper:]' '[:lower:]')"
  [[ "${norm}" =~ ^(1|true|yes|on)$ ]]
}

step() {
  echo "[gh-pages] $1"
}

fail() {
  echo "[gh-pages] ERROR: $1" >&2
  exit 1
}

command -v git >/dev/null 2>&1 || fail "git is required"

if ! git -C "${ROOT_DIR}" rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  fail "ROOT_DIR is not a git repository: ${ROOT_DIR}"
fi

if is_true "${BUILD_SITE}"; then
  step "Build docs site (serve disabled)"
  DOCS_LOCAL_SERVE=0 \
  DOCS_LOCAL_OPEN_BROWSER=0 \
  DOCS_LOCAL_INCLUDE_BENCH="${INCLUDE_BENCH}" \
  DOCS_LOCAL_BENCH_SOURCE="${BENCH_SOURCE}" \
  DOCS_LOCAL_STAGE_DIR="${STAGE_DIR}" \
    "${THIS_DIR}/publish-pages-local.sh"
fi

[[ -d "${SITE_DIR}" ]] || fail "site directory not found: ${SITE_DIR}"
[[ -f "${SITE_DIR}/index.html" || -f "${SITE_DIR}/README.md" ]] || fail "site directory appears empty: ${SITE_DIR}"

REMOTE_URL="$(git -C "${ROOT_DIR}" remote get-url "${REMOTE}" 2>/dev/null || true)"
[[ -n "${REMOTE_URL}" ]] || fail "remote not found: ${REMOTE}"

PUSH_URL="${REMOTE_URL_OVERRIDE}"
if [[ -z "${PUSH_URL}" && -n "${GITHUB_TOKEN:-}" && -n "${GITHUB_REPOSITORY:-}" ]]; then
  PUSH_URL="https://x-access-token:${GITHUB_TOKEN}@github.com/${GITHUB_REPOSITORY}.git"
fi
if [[ -z "${PUSH_URL}" ]]; then
  PUSH_URL="${REMOTE_URL}"
fi

TMP_DIR="$(mktemp -d /tmp/sikuligo-gh-pages.XXXXXX)"
cleanup() {
  rm -rf "${TMP_DIR}"
}
trap cleanup EXIT

step "Stage publish tree"
rsync -a --delete --exclude '.DS_Store' "${SITE_DIR}/" "${TMP_DIR}/"
touch "${TMP_DIR}/.nojekyll"

if [[ -n "${CNAME_VALUE}" ]]; then
  printf '%s\n' "${CNAME_VALUE}" > "${TMP_DIR}/CNAME"
  step "Set CNAME=${CNAME_VALUE}"
fi

step "Create publish commit for branch ${BRANCH}"
git -C "${TMP_DIR}" init -q
git -C "${TMP_DIR}" config user.name "${GIT_AUTHOR_NAME:-sikuligo-pages}"
git -C "${TMP_DIR}" config user.email "${GIT_AUTHOR_EMAIL:-sikuligo-pages@users.noreply.github.com}"
git -C "${TMP_DIR}" checkout -B "${BRANCH}" >/dev/null
git -C "${TMP_DIR}" add -A
git -C "${TMP_DIR}" commit -m "${COMMIT_MESSAGE}" >/dev/null
git -C "${TMP_DIR}" remote add "${REMOTE}" "${PUSH_URL}"

step "Push branch ${BRANCH} to ${REMOTE}"
if is_true "${FORCE_PUSH}"; then
  git -C "${TMP_DIR}" push --force "${REMOTE}" "${BRANCH}"
else
  git -C "${TMP_DIR}" push "${REMOTE}" "${BRANCH}"
fi

if is_true "${CONFIGURE_SOURCE}"; then
  command -v gh >/dev/null 2>&1 || fail "gh is required when GH_PAGES_CONFIGURE_SOURCE=1"
  gh auth status >/dev/null 2>&1 || fail "gh is not authenticated (run: gh auth login)"
  REPO_FULL_NAME="$(gh repo view --json nameWithOwner -q .nameWithOwner)"
  step "Configure GitHub Pages source to ${BRANCH}/"
  gh api \
    --method PUT \
    "/repos/${REPO_FULL_NAME}/pages" \
    -F "source[branch]=${BRANCH}" \
    -F "source[path]=/" >/dev/null
fi

step "Done"
echo "[gh-pages] published branch=${BRANCH} remote=${REMOTE} site_dir=${SITE_DIR}"
