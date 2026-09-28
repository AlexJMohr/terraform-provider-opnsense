// Copyright (c) Matthew Mellor
// SPDX-License-Identifier: MPL-2.0

package firewall

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestNatPortForwardRequestUsesNestedSourceDestination(t *testing.T) {
	m := NatPortForwardResourceModel{
		Enabled:         types.BoolValue(true),
		Interface:       types.StringValue("wan"),
		IPProtocol:      types.StringValue("inet"),
		Protocol:        types.StringValue("tcp"),
		SourceNet:       types.StringValue("any"),
		SourcePort:      types.StringValue(""),
		SourceNot:       types.BoolValue(false),
		DestinationNet:  types.StringValue("wanip"),
		DestinationPort: types.StringValue("3074"),
		DestinationNot:  types.BoolValue(false),
		Target:          types.StringValue("192.168.1.8"),
		LocalPort:       types.StringValue("3074"),
		Log:             types.BoolValue(false),
		Description:     types.StringValue("probe"),
		Categories:      types.SetNull(types.StringType),
	}

	req := m.toAPI(context.Background())

	raw, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	source, ok := payload["source"].(map[string]any)
	if !ok {
		t.Fatalf("payload has no nested \"source\" object; got %T: %s", payload["source"], raw)
	}
	if got := source["network"]; got != "any" {
		t.Errorf(`source.network = %v, want "any"`, got)
	}
	if got := source["port"]; got != "" {
		t.Errorf(`source.port = %v, want ""`, got)
	}
	if got := source["not"]; got != "0" {
		t.Errorf(`source.not = %v, want "0"`, got)
	}

	destination, ok := payload["destination"].(map[string]any)
	if !ok {
		t.Fatalf("payload has no nested \"destination\" object; got %T: %s", payload["destination"], raw)
	}
	if got := destination["network"]; got != "wanip" {
		t.Errorf(`destination.network = %v, want "wanip"`, got)
	}
	if got := destination["port"]; got != "3074" {
		t.Errorf(`destination.port = %v, want "3074"`, got)
	}

	if _, present := payload["source.network"]; present {
		t.Errorf(`flat "source.network" key still present in payload: %s`, raw)
	}
	if _, present := payload["destination.port"]; present {
		t.Errorf(`flat "destination.port" key still present in payload: %s`, raw)
	}
}

func TestNatPortForwardResponseNestedRoundTrip(t *testing.T) {
	const resp = `{
		"disabled": "0",
		"interface": {"wan": {"value": "WAN", "selected": 1}},
		"ipprotocol": {"inet": {"value": "IPv4", "selected": 1}},
		"protocol": {"tcp": {"value": "TCP", "selected": 1}},
		"source": {"network": "any", "port": "", "not": "0"},
		"destination": {"network": "wanip", "port": "3074", "not": "0"},
		"target": "192.168.1.8",
		"local-port": "3074",
		"log": "0",
		"descr": "probe",
		"categories": []
	}`

	var a natPortForwardAPIResponse
	if err := json.Unmarshal([]byte(resp), &a); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}

	var m NatPortForwardResourceModel
	m.fromAPI(context.Background(), &a, "test-uuid")

	if got := m.SourceNet.ValueString(); got != "any" {
		t.Errorf("SourceNet = %q, want %q", got, "any")
	}
	if got := m.DestinationPort.ValueString(); got != "3074" {
		t.Errorf("DestinationPort = %q, want %q", got, "3074")
	}
	if got := m.Target.ValueString(); got != "192.168.1.8" {
		t.Errorf("Target = %q, want %q", got, "192.168.1.8")
	}
	if m.Enabled.ValueBool() != true {
		t.Errorf("Enabled = %v, want true (disabled=\"0\")", m.Enabled.ValueBool())
	}
}
