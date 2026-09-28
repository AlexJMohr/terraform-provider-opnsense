---
title: Provider Management Roadmap - Complete Provider Coverage and OPNsense Assignment API
date: 2026-09-28
author: BMad PM
status: current
inputs:
  - prd.md
  - post-release-epics.md
  - support-matrix.md
  - docs/upstream-blocked.md
  - Provider ecosystem coverage comparison, 2026-09-28
  - OrgDocs h155 OPNsense provider/API gap note
---

# Provider Management Roadmap

## PM Decision

We should manage `matthew-on-git/opnsense` as the broad, modern, brownfield-friendly OPNsense provider and make it complete enough to be the default provider choice for new OPNsense automation.

Why this matters: this provider now has broad OPNsense surface coverage, but breadth only becomes trustworthy when it is backed by clear support status, release-gated API adoption, acceptance-test proof, and a coverage matrix that proves users can standardize on one provider rather than run multiple providers side-by-side. Migration content matters, but it is secondary to completeness.

## Current Facts

| Signal | `matthew-on-git/opnsense` | Established provider baseline |
|---|---:|---:|
| Current public release checked | `v0.4.5` | `v0.26.0` |
| Registry downloads checked | about 8.1K | about 1.33M |
| Resources / data sources | 102 / 88 | about 47 / 45 |
| Primary advantage | breadth, plugin/service coverage, generated docs | adoption, community recognition, polished positioning |
| Main product gap | proof and positioning | coverage breadth |

Coverage advantage already exists in HAProxy, ACME, Monit, syslog, traffic shaping, FRR/OSPF, auth/trust, interface types, and many service/plugin domains. The weak spots are not raw resource count; they are provable field-level coverage, release narrative, acceptance proof, v1 criteria, and clear migration guidance.

## OPNsense Assignment API Release Tracking

The September upstream PR detour changed the plan. We closed the duplicate upstream `address_settings` PR because current OPNsense `master` already has the relevant work under the existing assignment API:

- `api/interfaces/assignment/set_item/{ifname}`
- `api/interfaces/assignment/reconfigure`
- expanded `NetworkInterface.xml` fields including `enable`, `type4`, `ipaddr`, `type6`, `ipaddrv6`, gateways, DHCP fields, and hardware flags
- `/tmp/.interfaces.todo` staging and assignment `reconfigure` finalization

Release timing conclusion from the 2026-09-28 check:

- Not present in `stable/26.7`.
- Not present in tags checked through `26.7.4` or `27.1.a`.
- Present on `master` only, with relevant commits around Aug 20, Aug 27, and Sep 24.
- Unless OPNsense backports the work, the first stable provider target should be the next major train that carries the expanded assignment API, likely `27.1`.

Product rule: provider implementation may prototype against OPNsense `master` or snapshot, but stable release support must stay gated until the endpoint appears in a stable branch or released tag with durable semantics verified.

## Public Positioning Strategy

### Positioning

Use this message in README, Registry docs, and release notes:

> `matthew-on-git/opnsense` focuses on broad, modern OPNsense appliance management: core networking plus the plugin/service areas operators actually use at the edge, with brownfield import guidance and release-gated API support.

Avoid public comparisons that invite maintainer drama. Claim "broad modern coverage" or "complete appliance coverage goal" only where the support matrix backs it.

### Product Principles

1. Version-gate new OPNsense APIs instead of pretending `master` equals stable.
2. Treat acceptance proof as product surface, not private maintainer confidence.
3. Treat established ecosystem coverage as a product gate, including field-level schema depth where providers expose similar concepts.
4. Make migration from existing providers explicit and low-friction after the coverage position is understood.
5. Publish v1 criteria before declaring v1 readiness.
6. Do not ship workaround resources that scrape legacy PHP pages, edit XML directly, or bypass OPNsense API ownership.

## Backlog Epics

## Epic 30: OPNsense Assignment API Adoption

Goal: turn the upstream assignment API work into a safe provider feature as soon as it lands in a supported OPNsense release.

| Story | Priority | Status | File |
|---|---|---|---|
| 30.1 Track OPNsense Assignment API Release | P0 | backlog | `_bmad-output/implementation-artifacts/30-1-track-opnsense-assignment-api-release.md` |
| 30.2 Prototype Interface Assignment/IP Resource Against OPNsense Master | P0 | backlog | `_bmad-output/implementation-artifacts/30-2-prototype-interface-assignment-ip-resource.md` |
| 30.3 Update Interface Support Matrix and Public Blocker Docs | P0 | backlog | `_bmad-output/implementation-artifacts/30-3-update-interface-support-matrix.md` |

## Epic 31: Provider Completeness, Positioning, and Migration

Goal: make this provider the standard OPNsense provider by closing meaningful coverage gaps, then turning that completeness into user trust and adoption.

| Story | Priority | Status | File |
|---|---|---|---|
| 31.0 Complete Provider Coverage Audit | P0 | backlog | `_bmad-output/implementation-artifacts/31-0-complete-provider-coverage-audit.md` |
| 31.1 Complete Provider Positioning README and Registry Pass | P1 | backlog | `_bmad-output/implementation-artifacts/31-1-complete-provider-positioning-docs.md` |
| 31.2 Existing Provider Migration Guide | P1 | backlog | `_bmad-output/implementation-artifacts/31-2-existing-provider-migration-guide.md` |
| 31.3 Example Stacks for Adoption | P2 | backlog | `_bmad-output/implementation-artifacts/31-3-example-stacks-for-adoption.md` |

## Epic 32: Release Credibility and v1 Readiness

Goal: make releases feel boring, tested, and intentionally versioned.

| Story | Priority | Status | File |
|---|---|---|---|
| 32.1 Acceptance Test Proof Artifact | P1 | backlog | `_bmad-output/implementation-artifacts/32-1-acceptance-test-proof-artifact.md` |
| 32.2 Provider Preflight and Version Diagnostics | P2 | backlog | `_bmad-output/implementation-artifacts/32-2-provider-preflight-version-diagnostics.md` |
| 32.3 v1 Stability Roadmap | P2 | backlog | `_bmad-output/implementation-artifacts/32-3-v1-stability-roadmap.md` |

## Epic 33: Remaining Ecosystem Gap Mining

Goal: mine broader ecosystem-only areas without distracting from release tracking and the primary complete-provider coverage objective.

Candidates:

- captive portal templates/zones
- firewall groups
- NPTv6
- firewall automation filter/source NAT variants

This should stay P1/P2 until Epic 30 has a working release gate and Epic 31 has fixed the public positioning gap.

## Recommended Sequence

1. Story 30.1: create the release/API probe and support-matrix evidence.
2. Story 30.2: prototype the interface assignment/IP resource against OPNsense `master` with version guardrails.
3. Story 30.3: update public docs so "blocked" becomes "release-gated prototype" where accurate.
4. Story 31.0: perform a field-level complete-provider coverage audit and split implementation stories for any true product gaps.
5. Story 31.1: make README/Registry positioning match the actual provider advantage.
6. Story 31.2: publish the existing-provider migration guide.
7. Story 32.1: publish repeatable acceptance-test proof.
8. Story 32.3: define v1 criteria before v1 work is declared.
