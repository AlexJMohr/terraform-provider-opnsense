# Story 32.2: Provider Preflight and Version Diagnostics

## Status

backlog

## Story

As a Terraform user,
I want provider preflight diagnostics for OPNsense version, plugins, endpoints, and permissions,
so that I get actionable errors before CRUD operations fail halfway through an apply.

## Acceptance Criteria

1. Provider exposes a data source or configure-time diagnostics that report OPNsense version and plugin availability.
2. Diagnostics identify missing plugins required by configured resources.
3. Diagnostics identify missing endpoint/version support for release-gated resources such as interface assignment/IP settings.
4. Diagnostics identify likely ACL/permission gaps and suggest required API privileges.
5. Existing resources use the diagnostics path for clearer error messages where practical.

## Dev Notes

- This should build on `opnsense_system_info` if possible.
- Avoid noisy provider configure failures for configurations that do not use optional plugins.

## Validation

- Unit tests cover missing plugin, missing endpoint, and missing permission diagnostics.
- Acceptance tests cover at least one happy path and one missing-capability path.
- `make check` passes.
