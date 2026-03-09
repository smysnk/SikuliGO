---
layout: guide
title: Documentation Workflow
nav_key: contribution
kicker: Contributor Reference
lead: The working rules for editing end-user documentation while the guide-format rewrite is still moving content out of READMEs and deep technical pages.
---

This page implements the phase 0 freeze for end-user documentation.

## Goal

Keep the current documentation usable while the guide-format rewrite is in progress, and prevent user-facing content from drifting across package READMEs and docs pages.

## Source Of Truth

- Use `docs/strategy/documentation-source-map.md` to identify the current canonical source for a section.
- Treat the `Target page` in that file as the required destination for the guide-format rewrite.
- If a target page already exists, update it in the same change when you modify user-facing content.

## Package README Rules

- Do not land end-user documentation changes only in a package README.
- Keep package READMEs focused on quickstarts, install notes, and package-local behavior.
- When a package README changes user-facing behavior, update the matching docs target or its placeholder page in the same change.
- Preserve the `DOCS_CANONICAL_TARGET`, `DOCS_SOURCE_MAP`, and `DOCS_WORKFLOW` markers at the top of tracked READMEs.

## Generated Docs Rules

- Do not hand-edit generated API pages under `docs/reference/api/`.
- Use the existing generators and verification scripts for generated docs.
- Keep strategy and parity docs as supporting material, not as the only place a new user-facing behavior is explained.

## Update Sequence

1. Open `docs/strategy/documentation-source-map.md` and find the section you are touching.
2. Edit the current canonical source listed there.
3. Edit the matching target page under `docs/`.
4. If the target page is still a placeholder, keep it aligned with the new source material and note the supporting references.
5. Run `./scripts/check-docs-governance.sh`.
6. Run `./scripts/check-docs-links.sh`.

## Review Checklist

- The changed section still has one clear canonical source.
- The matching target page exists and was reviewed in the same change.
- No generated docs were edited by hand.
- Package README changes did not become the only place a user can learn a new behavior.
- Internal docs links still resolve through the published-site path model.
