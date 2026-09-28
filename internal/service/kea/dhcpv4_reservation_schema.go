// Copyright (c) Matthew Mellor
// SPDX-License-Identifier: MPL-2.0

package kea

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r *dhcpv4ReservationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Kea DHCPv4 host reservation on OPNsense, pinning an IPv4 address to a client MAC address within a subnet.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the reservation.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"subnet_id": schema.StringAttribute{
				Required: true, MarkdownDescription: "UUID of the Kea DHCPv4 subnet this reservation belongs to.",
			},
			"ip_address": schema.StringAttribute{
				Required: true, MarkdownDescription: "Reserved IPv4 address (must fall within the subnet).",
			},
			"hw_address": schema.StringAttribute{
				Required: true, MarkdownDescription: "Client MAC address the reservation matches (e.g. `00:11:22:33:44:55`).",
			},
			"client_id": schema.StringAttribute{
				Optional: true, Computed: true,
				Default:             stringdefault.StaticString(""),
				MarkdownDescription: "Client-id the reservation matches (colon-separated hex, alternative to `hw_address`).",
			},
			"hostname": schema.StringAttribute{
				Optional: true, Computed: true,
				Default:             stringdefault.StaticString(""),
				MarkdownDescription: "Hostname assigned to the client.",
			},
			"description": schema.StringAttribute{
				Optional: true, Computed: true,
				Default:             stringdefault.StaticString(""),
				MarkdownDescription: "Description of the reservation.",
			},
		},
	}
}
