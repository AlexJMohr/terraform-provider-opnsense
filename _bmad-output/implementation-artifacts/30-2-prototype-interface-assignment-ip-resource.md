# Story 30.2: Prototype Interface Assignment/IP Resource Against OPNsense Master

## Status

backlog

## Story

As an OPNsense Terraform user,
I want an experimental interface assignment/settings resource guarded by OPNsense version capability checks,
so that I can manage interface enablement and static/DHCP address settings once upstream ships the expanded assignment API.

## Acceptance Criteria

1. Prototype resource name is selected deliberately, with a preference for `opnsense_interface_assignment` or `opnsense_interface_settings`.
2. Resource can import an existing `wan`, `lan`, or `optN` assignment by interface name or stable upstream identifier.
3. Resource supports enable/disable and IPv4 mode/address/prefix/gateway where upstream supports it.
4. Resource supports IPv6 mode/address/prefix/gateway where upstream supports it.
5. Create/update calls use `api/interfaces/assignment/set_item/{ifname}` followed by `api/interfaces/assignment/reconfigure`.
6. Terraform plan is no-change after apply against an OPNsense `master`/snapshot appliance with the expanded API.
7. Provider fails clearly on OPNsense versions that lack the expanded assignment API.
8. Tests prove both the supported path and the "API missing" diagnostic path.

## Dev Notes

- This story is a prototype until Story 30.1 identifies the first stable release carrying the API.
- Avoid overload with API-backed virtual interface resources already supported by the provider, such as VLAN, bridge, GIF, GRE, LAGG, loopback, neighbor, and VXLAN.
- Base assignment/IP config has stricter blast radius than most resources. Acceptance tests should run on disposable interfaces in a QEMU/Vagrant appliance, not production firewall interfaces.
- Preserve global mutation serialization and reconfigure behavior.

## Validation

- Unit tests cover schema/version guard behavior.
- Acceptance test evidence is recorded against a disposable OPNsense `master`/snapshot VM.
- `make check` passes.
