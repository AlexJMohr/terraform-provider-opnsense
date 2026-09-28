# Story 30.3: Update Interface Support Matrix and Public Blocker Docs

## Status

backlog

## Story

As a provider evaluator,
I want the public support matrix and upstream-blocked register to reflect the September 2026 OPNsense assignment API pivot,
so that I understand what is supported now, what is release-gated, and what remains blocked.

## Acceptance Criteria

1. `_bmad-output/planning-artifacts/support-matrix.md` distinguishes current supported interface-type resources from release-gated base interface assignment/IP settings.
2. `docs/upstream-blocked.md` no longer points at stale PR #8436 as the main source of truth for interface assignment.
3. `docs/index.md`, `docs/migration-import.md`, and `templates/index.md.tmpl` are updated if their public support text is stale.
4. The public docs state that OPNsense `master` has the expanded assignment API, while stable release support remains gated until the API lands in a released branch/tag.
5. The docs explicitly keep PPPoE and any unsupported hardware flags out of the first provider story unless upstream stable semantics are verified.

## Dev Notes

- This story can be done before the implementation prototype if docs currently mislead users.
- Cite the local release probe from Story 30.1 once it exists.

## Validation

- `make docs` or the repository's docs generation path is run if templates change.
- `git diff --check` passes.
- `make check` passes if generated docs or code paths change.
