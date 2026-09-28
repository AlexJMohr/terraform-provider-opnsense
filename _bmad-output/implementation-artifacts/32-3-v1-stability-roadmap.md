# Story 32.3: v1 Stability Roadmap

## Status

backlog

## Story

As a provider user,
I want a clear v1 stability roadmap,
so that I know when schemas, import IDs, version support, and breaking-change policy become stable enough for long-lived production use.

## Acceptance Criteria

1. Public roadmap defines the criteria for `v1.0.0`.
2. Criteria include schema freeze rules, import ID stability, deprecation policy, minimum OPNsense version policy, and acceptance coverage threshold.
3. Roadmap distinguishes pre-v1 experimentation from post-v1 semver guarantees.
4. Roadmap lists domains that must be resolved, declared release-gated, or explicitly excluded before v1.
5. README and Registry docs link to the roadmap.

## Dev Notes

- Use existing technical research that already defines v1 as production-ready after schema stabilization and acceptance coverage.
- Do not promise `v1.0.0` until the acceptance proof story and support matrix are credible.

## Validation

- Roadmap is linked from public docs.
- `git diff --check` passes.
- Docs generation runs if templates change.
