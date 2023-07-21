//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazuresiterecovery

import "encoding/json"

func unmarshalDraModelCustomPropertiesClassification(rawMsg json.RawMessage) (DraModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DraModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "VMwareDraModelCustomProperties":
		b = &VMwareDraModelCustomProperties{}
	default:
		b = &DraModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalEventModelCustomPropertiesClassification(rawMsg json.RawMessage) (EventModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EventModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "HyperVToAzStackHCIEventModelCustomProperties":
		b = &HyperVToAzStackHCIEventModelCustomProperties{}
	default:
		b = &EventModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalFabricModelCustomPropertiesClassification(rawMsg json.RawMessage) (FabricModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FabricModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "AzStackHCIFabricModelCustomProperties":
		b = &AzStackHCIFabricModelCustomProperties{}
	case "HyperVMigrateFabricModelCustomProperties":
		b = &HyperVMigrateFabricModelCustomProperties{}
	case "VMwareMigrateFabricModelCustomProperties":
		b = &VMwareMigrateFabricModelCustomProperties{}
	default:
		b = &FabricModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalPlannedFailoverModelCustomPropertiesClassification(rawMsg json.RawMessage) (PlannedFailoverModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PlannedFailoverModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "HyperVToAzStackHCIPlannedFailoverModelCustomProperties":
		b = &HyperVToAzStackHCIPlannedFailoverModelCustomProperties{}
	case "VMwareToAzStackHCIPlannedFailoverModelCustomProperties":
		b = &VMwareToAzStackHCIPlannedFailoverModelCustomProperties{}
	default:
		b = &PlannedFailoverModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalPolicyModelCustomPropertiesClassification(rawMsg json.RawMessage) (PolicyModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PolicyModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "HyperVToAzStackHCIPolicyModelCustomProperties":
		b = &HyperVToAzStackHCIPolicyModelCustomProperties{}
	case "VMwareToAzStackHCIPolicyModelCustomProperties":
		b = &VMwareToAzStackHCIPolicyModelCustomProperties{}
	default:
		b = &PolicyModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalProtectedItemModelCustomPropertiesClassification(rawMsg json.RawMessage) (ProtectedItemModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ProtectedItemModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "HyperVToAzStackHCIProtectedItemModelCustomProperties":
		b = &HyperVToAzStackHCIProtectedItemModelCustomProperties{}
	case "VMwareToAzStackHCIProtectedItemModelCustomProperties":
		b = &VMwareToAzStackHCIProtectedItemModelCustomProperties{}
	default:
		b = &ProtectedItemModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalRecoveryPointModelCustomPropertiesClassification(rawMsg json.RawMessage) (RecoveryPointModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RecoveryPointModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "HyperVToAzStackHCIRecoveryPointModelCustomProperties":
		b = &HyperVToAzStackHCIRecoveryPointModelCustomProperties{}
	default:
		b = &RecoveryPointModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalReplicationExtensionModelCustomPropertiesClassification(rawMsg json.RawMessage) (ReplicationExtensionModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ReplicationExtensionModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "HyperVToAzStackHCIReplicationExtensionModelCustomProperties":
		b = &HyperVToAzStackHCIReplicationExtensionModelCustomProperties{}
	case "VMwareToAzStackHCIReplicationExtensionModelCustomProperties":
		b = &VMwareToAzStackHCIReplicationExtensionModelCustomProperties{}
	default:
		b = &ReplicationExtensionModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalWorkflowModelCustomPropertiesClassification(rawMsg json.RawMessage) (WorkflowModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b WorkflowModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "FailoverWorkflowModelCustomProperties":
		b = &FailoverWorkflowModelCustomProperties{}
	case "TestFailoverCleanupWorkflowModelCustomProperties":
		b = &TestFailoverCleanupWorkflowModelCustomProperties{}
	case "TestFailoverWorkflowModelCustomProperties":
		b = &TestFailoverWorkflowModelCustomProperties{}
	default:
		b = &WorkflowModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
