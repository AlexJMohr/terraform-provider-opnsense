# Story 30.1: Track OPNsense Assignment API Release

## Status

backlog

## Story

As a provider maintainer,
I want an automated release/API tracker for the expanded OPNsense `interfaces/assignment` API,
so that we know exactly when interface IP configuration can be safely implemented for a supported release.

## Acceptance Criteria

1. A repeatable probe checks OPNsense `master`, active `stable/*` branches, and latest release tags for `NetworkInterface.xml` fields required by provider-managed interface settings.
2. The probe records whether `type4`, `ipaddr`, `type6`, `ipaddrv6`, gateway fields, DHCP fields, assignment `set_item`, and assignment `reconfigure` exist.
3. The probe distinguishes `master` evidence from stable release support.
4. Results update or feed `_bmad-output/planning-artifacts/support-matrix.md` and `docs/upstream-blocked.md`.
5. The tracker captures the first released tag/branch that contains the expanded API and links to the upstream commits.
6. No provider resource is marked stable until the API is present in a released OPNsense version or explicitly documented as snapshot-only.

## Dev Notes

- September 2026 finding: the duplicate upstream PR was closed because OPNsense `master` already has the assignment API expansion.
- `stable/26.7`, `26.7.4`, and `27.1.a` did not contain the full expanded assignment API when checked.
- Likely stable target is the next major train after the checked tags, expected around `27.1` unless OPNsense backports the work.
- Do not create a workaround based on legacy `interfaces.php`, XML writes, or GUI scraping.

## Validation

- Probe output is deterministic and committed or stored as a planning artifact.
- `git diff --check` passes.
- If scripts are added, `make check` must pass before completion.
