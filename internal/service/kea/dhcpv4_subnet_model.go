// Copyright (c) Matthew Mellor
// SPDX-License-Identifier: MPL-2.0

package kea

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/matthew-on-git/terraform-provider-opnsense/pkg/opnsense"
)

// DHCPv4SubnetResourceModel is the Terraform state model for opnsense_kea_dhcpv4_subnet.
type DHCPv4SubnetResourceModel struct {
	ID            types.String `tfsdk:"id"`
	Subnet        types.String `tfsdk:"subnet"`
	Allocator     types.String `tfsdk:"allocator"`
	Pools         types.String `tfsdk:"pools"`
	MatchClientID types.Bool   `tfsdk:"match_client_id"`
	PingCheck     types.Bool   `tfsdk:"ping_check"`
	Description   types.String `tfsdk:"description"`
}

type dhcpv4SubnetAPIResponse struct {
	Subnet        string               `json:"subnet"`
	Allocator     opnsense.SelectedMap `json:"allocator"`
	Pools         string               `json:"pools"`
	MatchClientID string               `json:"match-client-id"`
	PingCheck     string               `json:"ping_check"`
	Description   string               `json:"description"`
}

type dhcpv4SubnetAPIRequest struct {
	Subnet        string `json:"subnet"`
	Allocator     string `json:"allocator,omitempty"`
	Pools         string `json:"pools,omitempty"`
	MatchClientID string `json:"match-client-id"`
	PingCheck     string `json:"ping_check"`
	Description   string `json:"description,omitempty"`
}

func (m *DHCPv4SubnetResourceModel) toAPI(_ context.Context) *dhcpv4SubnetAPIRequest {
	return &dhcpv4SubnetAPIRequest{
		Subnet:        m.Subnet.ValueString(),
		Allocator:     m.Allocator.ValueString(),
		Pools:         m.Pools.ValueString(),
		MatchClientID: opnsense.BoolToString(m.MatchClientID.ValueBool()),
		PingCheck:     opnsense.BoolToString(m.PingCheck.ValueBool()),
		Description:   m.Description.ValueString(),
	}
}

func (m *DHCPv4SubnetResourceModel) fromAPI(_ context.Context, a *dhcpv4SubnetAPIResponse, uuid string) {
	m.ID = types.StringValue(uuid)
	m.Subnet = types.StringValue(a.Subnet)
	m.Allocator = types.StringValue(string(a.Allocator))
	m.Pools = types.StringValue(a.Pools)
	m.MatchClientID = types.BoolValue(opnsense.StringToBool(a.MatchClientID))
	m.PingCheck = types.BoolValue(opnsense.StringToBool(a.PingCheck))
	m.Description = types.StringValue(a.Description)
}
