// Copyright (c) Matthew Mellor
// SPDX-License-Identifier: MPL-2.0

package kea

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/matthew-on-git/terraform-provider-opnsense/pkg/opnsense"
)

// Ensure dhcpv4ReservationResource satisfies the resource interfaces.
var (
	_ resource.Resource                = &dhcpv4ReservationResource{}
	_ resource.ResourceWithImportState = &dhcpv4ReservationResource{}
)

// dhcpv4ReservationReqOpts configures the OPNsense API endpoints for Kea DHCPv4 reservations.
var dhcpv4ReservationReqOpts = opnsense.ReqOpts{
	AddEndpoint:         "/api/kea/dhcpv4/add_reservation",
	GetEndpoint:         "/api/kea/dhcpv4/get_reservation",
	UpdateEndpoint:      "/api/kea/dhcpv4/set_reservation",
	DeleteEndpoint:      "/api/kea/dhcpv4/del_reservation",
	SearchEndpoint:      "/api/kea/dhcpv4/search_reservation",
	ReconfigureEndpoint: "/api/kea/service/reconfigure",
	Monad:               "reservation",
}

// dhcpv4ReservationResource implements the opnsense_kea_dhcpv4_reservation resource.
type dhcpv4ReservationResource struct {
	client *opnsense.Client
}

func newDHCPv4ReservationResource() resource.Resource {
	return &dhcpv4ReservationResource{}
}

// Metadata sets the resource type name.
func (r *dhcpv4ReservationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kea_dhcpv4_reservation"
}

// Configure extracts the OPNsense API client from provider data.
func (r *dhcpv4ReservationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*opnsense.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Provider Data",
			"Expected *opnsense.Client, got something else.",
		)
		return
	}
	r.client = client
}

// Create creates a new Kea DHCPv4 reservation via the OPNsense API.
func (r *dhcpv4ReservationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DHCPv4ReservationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq := plan.toAPI(ctx)

	uuid, err := opnsense.Add(ctx, r.client, dhcpv4ReservationReqOpts, apiReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Kea DHCPv4 reservation",
			fmt.Sprintf("Could not create Kea DHCPv4 reservation: %s", err),
		)
		return
	}

	result, err := opnsense.Get[dhcpv4ReservationAPIResponse](ctx, r.client, dhcpv4ReservationReqOpts, uuid)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Kea DHCPv4 reservation after create",
			fmt.Sprintf("Created reservation %s but could not read it back: %s", uuid, err),
		)
		return
	}

	plan.fromAPI(ctx, result, uuid)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state from the OPNsense API.
func (r *dhcpv4ReservationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DHCPv4ReservationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	result, err := opnsense.Get[dhcpv4ReservationAPIResponse](ctx, r.client, dhcpv4ReservationReqOpts, state.ID.ValueString())
	if err != nil {
		var notFoundErr *opnsense.NotFoundError
		if errors.As(err, &notFoundErr) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error reading Kea DHCPv4 reservation",
			fmt.Sprintf("Could not read Kea DHCPv4 reservation %s: %s", state.ID.ValueString(), err),
		)
		return
	}

	state.fromAPI(ctx, result, state.ID.ValueString())
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update modifies an existing Kea DHCPv4 reservation via the OPNsense API.
func (r *dhcpv4ReservationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DHCPv4ReservationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state DHCPv4ReservationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq := plan.toAPI(ctx)
	id := state.ID.ValueString()

	err := opnsense.Update(ctx, r.client, dhcpv4ReservationReqOpts, apiReq, id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Kea DHCPv4 reservation",
			fmt.Sprintf("Could not update Kea DHCPv4 reservation %s: %s", id, err),
		)
		return
	}

	result, err := opnsense.Get[dhcpv4ReservationAPIResponse](ctx, r.client, dhcpv4ReservationReqOpts, id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Kea DHCPv4 reservation after update",
			fmt.Sprintf("Updated reservation %s but could not read it back: %s", id, err),
		)
		return
	}

	plan.fromAPI(ctx, result, id)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Delete removes a Kea DHCPv4 reservation from the OPNsense API.
func (r *dhcpv4ReservationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DHCPv4ReservationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := opnsense.Delete(ctx, r.client, dhcpv4ReservationReqOpts, state.ID.ValueString())
	if err != nil {
		var notFoundErr *opnsense.NotFoundError
		if errors.As(err, &notFoundErr) {
			return
		}
		resp.Diagnostics.AddError(
			"Error deleting Kea DHCPv4 reservation",
			fmt.Sprintf("Could not delete Kea DHCPv4 reservation %s: %s", state.ID.ValueString(), err),
		)
	}
}

// ImportState imports an existing Kea DHCPv4 reservation by UUID.
func (r *dhcpv4ReservationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
