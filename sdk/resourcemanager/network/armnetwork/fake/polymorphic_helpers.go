//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

func unmarshalActiveBaseSecurityAdminRuleClassification(rawMsg json.RawMessage) (armnetwork.ActiveBaseSecurityAdminRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armnetwork.ActiveBaseSecurityAdminRuleClassification
	switch m["kind"] {
	case string(armnetwork.EffectiveAdminRuleKindCustom):
		b = &armnetwork.ActiveSecurityAdminRule{}
	case string(armnetwork.EffectiveAdminRuleKindDefault):
		b = &armnetwork.ActiveDefaultSecurityAdminRule{}
	default:
		b = &armnetwork.ActiveBaseSecurityAdminRule{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalActiveBaseSecurityAdminRuleClassificationArray(rawMsg json.RawMessage) ([]armnetwork.ActiveBaseSecurityAdminRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]armnetwork.ActiveBaseSecurityAdminRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalActiveBaseSecurityAdminRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalBaseAdminRuleClassification(rawMsg json.RawMessage) (armnetwork.BaseAdminRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armnetwork.BaseAdminRuleClassification
	switch m["kind"] {
	case string(armnetwork.AdminRuleKindCustom):
		b = &armnetwork.AdminRule{}
	case string(armnetwork.AdminRuleKindDefault):
		b = &armnetwork.DefaultAdminRule{}
	default:
		b = &armnetwork.BaseAdminRule{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalBaseAdminRuleClassificationArray(rawMsg json.RawMessage) ([]armnetwork.BaseAdminRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]armnetwork.BaseAdminRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalBaseAdminRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalEffectiveBaseSecurityAdminRuleClassification(rawMsg json.RawMessage) (armnetwork.EffectiveBaseSecurityAdminRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armnetwork.EffectiveBaseSecurityAdminRuleClassification
	switch m["kind"] {
	case string(armnetwork.EffectiveAdminRuleKindCustom):
		b = &armnetwork.EffectiveSecurityAdminRule{}
	case string(armnetwork.EffectiveAdminRuleKindDefault):
		b = &armnetwork.EffectiveDefaultSecurityAdminRule{}
	default:
		b = &armnetwork.EffectiveBaseSecurityAdminRule{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalEffectiveBaseSecurityAdminRuleClassificationArray(rawMsg json.RawMessage) ([]armnetwork.EffectiveBaseSecurityAdminRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]armnetwork.EffectiveBaseSecurityAdminRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalEffectiveBaseSecurityAdminRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalFirewallPolicyRuleClassification(rawMsg json.RawMessage) (armnetwork.FirewallPolicyRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armnetwork.FirewallPolicyRuleClassification
	switch m["ruleType"] {
	case string(armnetwork.FirewallPolicyRuleTypeApplicationRule):
		b = &armnetwork.ApplicationRule{}
	case string(armnetwork.FirewallPolicyRuleTypeNatRule):
		b = &armnetwork.NatRule{}
	case string(armnetwork.FirewallPolicyRuleTypeNetworkRule):
		b = &armnetwork.Rule{}
	default:
		b = &armnetwork.FirewallPolicyRule{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalFirewallPolicyRuleClassificationArray(rawMsg json.RawMessage) ([]armnetwork.FirewallPolicyRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]armnetwork.FirewallPolicyRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalFirewallPolicyRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalFirewallPolicyRuleCollectionClassification(rawMsg json.RawMessage) (armnetwork.FirewallPolicyRuleCollectionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armnetwork.FirewallPolicyRuleCollectionClassification
	switch m["ruleCollectionType"] {
	case string(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyFilterRuleCollection):
		b = &armnetwork.FirewallPolicyFilterRuleCollection{}
	case string(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyNatRuleCollection):
		b = &armnetwork.FirewallPolicyNatRuleCollection{}
	default:
		b = &armnetwork.FirewallPolicyRuleCollection{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalFirewallPolicyRuleCollectionClassificationArray(rawMsg json.RawMessage) ([]armnetwork.FirewallPolicyRuleCollectionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]armnetwork.FirewallPolicyRuleCollectionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalFirewallPolicyRuleCollectionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}
