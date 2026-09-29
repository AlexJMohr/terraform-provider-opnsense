// Copyright (c) Matthew Mellor
// SPDX-License-Identifier: MPL-2.0

package kea

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r *dhcpv4SubnetResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Kea DHCPv4 subnet on OPNsense. The subnet's interface must be enabled in the Kea DHCPv4 general settings (`opnsense_kea_dhcpv4_settings`) first.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the subnet.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"subnet": schema.StringAttribute{
				Required: true, MarkdownDescription: "Subnet in CIDR notation (e.g. `192.168.40.0/24`).",
			},
			"allocator": schema.StringAttribute{
				Optional: true, Computed: true,
				Default:             stringdefault.StaticString(""),
				MarkdownDescription: "Lease allocator strategy (empty = default, `iterative`, `random`, or `flq`).",
			},
			"pools": schema.StringAttribute{
				Optional: true, Computed: true,
				Default:             stringdefault.StaticString(""),
				MarkdownDescription: "Address pools for the subnet (newline-separated ranges).",
			},
			"match_client_id": schema.BoolAttribute{
				Optional: true, Computed: true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Match clients by client-id in addition to hardware address. Defaults to `true`.",
			},
			"ping_check": schema.BoolAttribute{
				Optional: true, Computed: true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Ping the address before handing out a lease. Defaults to `false`.",
			},
			"description": schema.StringAttribute{
				Optional: true, Computed: true,
				Default:             stringdefault.StaticString(""),
				MarkdownDescription: "Description of the subnet.",
			},
		},
	}
}
