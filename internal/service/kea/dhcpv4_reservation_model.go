// Copyright (c) Matthew Mellor
// SPDX-License-Identifier: MPL-2.0

package kea

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/matthew-on-git/terraform-provider-opnsense/pkg/opnsense"
)

// DHCPv4ReservationResourceModel is the Terraform state model for
// opnsense_kea_dhcpv4_reservation.
type DHCPv4ReservationResourceModel struct {
	ID          types.String `tfsdk:"id"`
	SubnetID    types.String `tfsdk:"subnet_id"`
	IPAddress   types.String `tfsdk:"ip_address"`
	HWAddress   types.String `tfsdk:"hw_address"`
	ClientID    types.String `tfsdk:"client_id"`
	Hostname    types.String `tfsdk:"hostname"`
	Description types.String `tfsdk:"description"`
}

type dhcpv4ReservationAPIResponse struct {
	Subnet      opnsense.SelectedMap `json:"subnet"`
	IPAddress   string               `json:"ip_address"`
	HWAddress   string               `json:"hw_address"`
	ClientID    string               `json:"client_id"`
	Hostname    string               `json:"hostname"`
	Description string               `json:"description"`
}

type dhcpv4ReservationAPIRequest struct {
	Subnet      string `json:"subnet"`
	IPAddress   string `json:"ip_address,omitempty"`
	HWAddress   string `json:"hw_address,omitempty"`
	ClientID    string `json:"client_id,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	Description string `json:"description,omitempty"`
}

func (m *DHCPv4ReservationResourceModel) toAPI(_ context.Context) *dhcpv4ReservationAPIRequest {
	return &dhcpv4ReservationAPIRequest{
		Subnet:      m.SubnetID.ValueString(),
		IPAddress:   m.IPAddress.ValueString(),
		HWAddress:   m.HWAddress.ValueString(),
		ClientID:    m.ClientID.ValueString(),
		Hostname:    m.Hostname.ValueString(),
		Description: m.Description.ValueString(),
	}
}

func (m *DHCPv4ReservationResourceModel) fromAPI(_ context.Context, a *dhcpv4ReservationAPIResponse, uuid string) {
	m.ID = types.StringValue(uuid)
	m.SubnetID = types.StringValue(string(a.Subnet))
	m.IPAddress = types.StringValue(a.IPAddress)
	m.HWAddress = types.StringValue(a.HWAddress)
	m.ClientID = types.StringValue(a.ClientID)
	m.Hostname = types.StringValue(a.Hostname)
	m.Description = types.StringValue(a.Description)
}
