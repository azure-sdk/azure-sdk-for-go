//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicessiterecovery

import "encoding/json"

func unmarshalAddDisksProviderSpecificInputClassification(rawMsg json.RawMessage) (AddDisksProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AddDisksProviderSpecificInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2AAddDisksInput{}
	default:
		b = &AddDisksProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalApplianceSpecificDetailsClassification(rawMsg json.RawMessage) (ApplianceSpecificDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ApplianceSpecificDetailsClassification
	switch m["instanceType"] {
	case "InMageRcm":
		b = &InMageRcmApplianceSpecificDetails{}
	default:
		b = &ApplianceSpecificDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalApplyRecoveryPointProviderSpecificInputClassification(rawMsg json.RawMessage) (ApplyRecoveryPointProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ApplyRecoveryPointProviderSpecificInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2AApplyRecoveryPointInput{}
	case "A2ACrossClusterMigration":
		b = &A2ACrossClusterMigrationApplyRecoveryPointInput{}
	case "HyperVReplicaAzure":
		b = &HyperVReplicaAzureApplyRecoveryPointInput{}
	case "InMageAzureV2":
		b = &InMageAzureV2ApplyRecoveryPointInput{}
	case "InMageRcm":
		b = &InMageRcmApplyRecoveryPointInput{}
	default:
		b = &ApplyRecoveryPointProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalClusterUnplannedFailoverProviderSpecificInputClassification(rawMsg json.RawMessage) (ClusterUnplannedFailoverProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ClusterUnplannedFailoverProviderSpecificInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2AClusterUnplannedFailoverInput{}
	default:
		b = &ClusterUnplannedFailoverProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalConfigurationSettingsClassification(rawMsg json.RawMessage) (ConfigurationSettingsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ConfigurationSettingsClassification
	switch m["instanceType"] {
	case "HyperVVirtualMachine":
		b = &HyperVVirtualMachineDetails{}
	case "ReplicationGroupDetails":
		b = &ReplicationGroupDetails{}
	case "VMwareVirtualMachine":
		b = &VMwareVirtualMachineDetails{}
	case "VmmVirtualMachine":
		b = &VmmVirtualMachineDetails{}
	default:
		b = &ConfigurationSettings{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalCreateProtectionIntentProviderSpecificDetailsClassification(rawMsg json.RawMessage) (CreateProtectionIntentProviderSpecificDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CreateProtectionIntentProviderSpecificDetailsClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2ACreateProtectionIntentInput{}
	default:
		b = &CreateProtectionIntentProviderSpecificDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDisableProtectionProviderSpecificInputClassification(rawMsg json.RawMessage) (DisableProtectionProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DisableProtectionProviderSpecificInputClassification
	switch m["instanceType"] {
	case "InMage":
		b = &InMageDisableProtectionProviderSpecificInput{}
	default:
		b = &DisableProtectionProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalEnableMigrationProviderSpecificInputClassification(rawMsg json.RawMessage) (EnableMigrationProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EnableMigrationProviderSpecificInputClassification
	switch m["instanceType"] {
	case "VMwareCbt":
		b = &VMwareCbtEnableMigrationInput{}
	default:
		b = &EnableMigrationProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalEnableProtectionProviderSpecificInputClassification(rawMsg json.RawMessage) (EnableProtectionProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EnableProtectionProviderSpecificInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2AEnableProtectionInput{}
	case "A2ACrossClusterMigration":
		b = &A2ACrossClusterMigrationEnableProtectionInput{}
	case "HyperVReplicaAzure":
		b = &HyperVReplicaAzureEnableProtectionInput{}
	case "InMage":
		b = &InMageEnableProtectionInput{}
	case "InMageAzureV2":
		b = &InMageAzureV2EnableProtectionInput{}
	case "InMageRcm":
		b = &InMageRcmEnableProtectionInput{}
	default:
		b = &EnableProtectionProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalEventProviderSpecificDetailsClassification(rawMsg json.RawMessage) (EventProviderSpecificDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EventProviderSpecificDetailsClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2AEventDetails{}
	case "HyperVReplica2012":
		b = &HyperVReplica2012EventDetails{}
	case "HyperVReplica2012R2":
		b = &HyperVReplica2012R2EventDetails{}
	case "HyperVReplicaAzure":
		b = &HyperVReplicaAzureEventDetails{}
	case "HyperVReplicaBaseEventDetails":
		b = &HyperVReplicaBaseEventDetails{}
	case "InMageAzureV2":
		b = &InMageAzureV2EventDetails{}
	case "InMageRcm":
		b = &InMageRcmEventDetails{}
	case "InMageRcmFailback":
		b = &InMageRcmFailbackEventDetails{}
	case "VMwareCbt":
		b = &VMwareCbtEventDetails{}
	default:
		b = &EventProviderSpecificDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalEventSpecificDetailsClassification(rawMsg json.RawMessage) (EventSpecificDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EventSpecificDetailsClassification
	switch m["instanceType"] {
	case "JobStatus":
		b = &JobStatusEventDetails{}
	default:
		b = &EventSpecificDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalFabricSpecificCreateNetworkMappingInputClassification(rawMsg json.RawMessage) (FabricSpecificCreateNetworkMappingInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FabricSpecificCreateNetworkMappingInputClassification
	switch m["instanceType"] {
	case "AzureToAzure":
		b = &AzureToAzureCreateNetworkMappingInput{}
	case "VmmToAzure":
		b = &VmmToAzureCreateNetworkMappingInput{}
	case "VmmToVmm":
		b = &VmmToVmmCreateNetworkMappingInput{}
	default:
		b = &FabricSpecificCreateNetworkMappingInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalFabricSpecificCreationInputClassification(rawMsg json.RawMessage) (FabricSpecificCreationInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FabricSpecificCreationInputClassification
	switch m["instanceType"] {
	case "Azure":
		b = &AzureFabricCreationInput{}
	case "InMageRcm":
		b = &InMageRcmFabricCreationInput{}
	case "VMwareV2":
		b = &VMwareV2FabricCreationInput{}
	default:
		b = &FabricSpecificCreationInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalFabricSpecificDetailsClassification(rawMsg json.RawMessage) (FabricSpecificDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FabricSpecificDetailsClassification
	switch m["instanceType"] {
	case "Azure":
		b = &AzureFabricSpecificDetails{}
	case "HyperVSite":
		b = &HyperVSiteDetails{}
	case "InMageRcm":
		b = &InMageRcmFabricSpecificDetails{}
	case "VMM":
		b = &VmmDetails{}
	case "VMware":
		b = &VMwareDetails{}
	case "VMwareV2":
		b = &VMwareV2FabricSpecificDetails{}
	default:
		b = &FabricSpecificDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalFabricSpecificUpdateNetworkMappingInputClassification(rawMsg json.RawMessage) (FabricSpecificUpdateNetworkMappingInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FabricSpecificUpdateNetworkMappingInputClassification
	switch m["instanceType"] {
	case "AzureToAzure":
		b = &AzureToAzureUpdateNetworkMappingInput{}
	case "VmmToAzure":
		b = &VmmToAzureUpdateNetworkMappingInput{}
	case "VmmToVmm":
		b = &VmmToVmmUpdateNetworkMappingInput{}
	default:
		b = &FabricSpecificUpdateNetworkMappingInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalGroupTaskDetailsClassification(rawMsg json.RawMessage) (GroupTaskDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b GroupTaskDetailsClassification
	switch m["instanceType"] {
	case "InlineWorkflowTaskDetails":
		b = &InlineWorkflowTaskDetails{}
	case "RecoveryPlanGroupTaskDetails":
		b = &RecoveryPlanGroupTaskDetails{}
	case "RecoveryPlanShutdownGroupTaskDetails":
		b = &RecoveryPlanShutdownGroupTaskDetails{}
	default:
		b = &GroupTaskDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalJobDetailsClassification(rawMsg json.RawMessage) (JobDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b JobDetailsClassification
	switch m["instanceType"] {
	case "AsrJobDetails":
		b = &AsrJobDetails{}
	case "ExportJobDetails":
		b = &ExportJobDetails{}
	case "FailoverJobDetails":
		b = &FailoverJobDetails{}
	case "SwitchProtectionJobDetails":
		b = &SwitchProtectionJobDetails{}
	case "TestFailoverJobDetails":
		b = &TestFailoverJobDetails{}
	default:
		b = &JobDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalMigrateProviderSpecificInputClassification(rawMsg json.RawMessage) (MigrateProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MigrateProviderSpecificInputClassification
	switch m["instanceType"] {
	case "VMwareCbt":
		b = &VMwareCbtMigrateInput{}
	default:
		b = &MigrateProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalMigrationProviderSpecificSettingsClassification(rawMsg json.RawMessage) (MigrationProviderSpecificSettingsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MigrationProviderSpecificSettingsClassification
	switch m["instanceType"] {
	case "VMwareCbt":
		b = &VMwareCbtMigrationDetails{}
	default:
		b = &MigrationProviderSpecificSettings{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalNetworkMappingFabricSpecificSettingsClassification(rawMsg json.RawMessage) (NetworkMappingFabricSpecificSettingsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b NetworkMappingFabricSpecificSettingsClassification
	switch m["instanceType"] {
	case "AzureToAzure":
		b = &AzureToAzureNetworkMappingSettings{}
	case "VmmToAzure":
		b = &VmmToAzureNetworkMappingSettings{}
	case "VmmToVmm":
		b = &VmmToVmmNetworkMappingSettings{}
	default:
		b = &NetworkMappingFabricSpecificSettings{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalPlannedFailoverProviderSpecificFailoverInputClassification(rawMsg json.RawMessage) (PlannedFailoverProviderSpecificFailoverInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PlannedFailoverProviderSpecificFailoverInputClassification
	switch m["instanceType"] {
	case "HyperVReplicaAzure":
		b = &HyperVReplicaAzurePlannedFailoverProviderInput{}
	case "HyperVReplicaAzureFailback":
		b = &HyperVReplicaAzureFailbackProviderInput{}
	case "InMageRcmFailback":
		b = &InMageRcmFailbackPlannedFailoverProviderInput{}
	default:
		b = &PlannedFailoverProviderSpecificFailoverInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalPolicyProviderSpecificDetailsClassification(rawMsg json.RawMessage) (PolicyProviderSpecificDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PolicyProviderSpecificDetailsClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2APolicyDetails{}
	case "HyperVReplica2012":
		b = &HyperVReplicaPolicyDetails{}
	case "HyperVReplica2012R2":
		b = &HyperVReplicaBluePolicyDetails{}
	case "HyperVReplicaAzure":
		b = &HyperVReplicaAzurePolicyDetails{}
	case "HyperVReplicaBasePolicyDetails":
		b = &HyperVReplicaBasePolicyDetails{}
	case "InMage":
		b = &InMagePolicyDetails{}
	case "InMageAzureV2":
		b = &InMageAzureV2PolicyDetails{}
	case "InMageBasePolicyDetails":
		b = &InMageBasePolicyDetails{}
	case "InMageRcm":
		b = &InMageRcmPolicyDetails{}
	case "InMageRcmFailback":
		b = &InMageRcmFailbackPolicyDetails{}
	case "VMwareCbt":
		b = &VmwareCbtPolicyDetails{}
	default:
		b = &PolicyProviderSpecificDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalPolicyProviderSpecificInputClassification(rawMsg json.RawMessage) (PolicyProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PolicyProviderSpecificInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2APolicyCreationInput{}
	case "A2ACrossClusterMigration":
		b = &A2ACrossClusterMigrationPolicyCreationInput{}
	case "HyperVReplica2012":
		b = &HyperVReplicaPolicyInput{}
	case "HyperVReplica2012R2":
		b = &HyperVReplicaBluePolicyInput{}
	case "HyperVReplicaAzure":
		b = &HyperVReplicaAzurePolicyInput{}
	case "InMage":
		b = &InMagePolicyInput{}
	case "InMageAzureV2":
		b = &InMageAzureV2PolicyInput{}
	case "InMageRcm":
		b = &InMageRcmPolicyCreationInput{}
	case "InMageRcmFailback":
		b = &InMageRcmFailbackPolicyCreationInput{}
	case "VMwareCbt":
		b = &VMwareCbtPolicyCreationInput{}
	default:
		b = &PolicyProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalProtectionContainerMappingProviderSpecificDetailsClassification(rawMsg json.RawMessage) (ProtectionContainerMappingProviderSpecificDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ProtectionContainerMappingProviderSpecificDetailsClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2AProtectionContainerMappingDetails{}
	case "InMageRcm":
		b = &InMageRcmProtectionContainerMappingDetails{}
	case "VMwareCbt":
		b = &VMwareCbtProtectionContainerMappingDetails{}
	default:
		b = &ProtectionContainerMappingProviderSpecificDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalProtectionProfileCustomDetailsClassification(rawMsg json.RawMessage) (ProtectionProfileCustomDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ProtectionProfileCustomDetailsClassification
	switch m["resourceType"] {
	case "Existing":
		b = &ExistingProtectionProfile{}
	case "New":
		b = &NewProtectionProfile{}
	default:
		b = &ProtectionProfileCustomDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalProviderSpecificRecoveryPointDetailsClassification(rawMsg json.RawMessage) (ProviderSpecificRecoveryPointDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ProviderSpecificRecoveryPointDetailsClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2ARecoveryPointDetails{}
	case "InMageAzureV2":
		b = &InMageAzureV2RecoveryPointDetails{}
	case "InMageRcm":
		b = &InMageRcmRecoveryPointDetails{}
	default:
		b = &ProviderSpecificRecoveryPointDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalRecoveryAvailabilitySetCustomDetailsClassification(rawMsg json.RawMessage) (RecoveryAvailabilitySetCustomDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RecoveryAvailabilitySetCustomDetailsClassification
	switch m["resourceType"] {
	case "Existing":
		b = &ExistingRecoveryAvailabilitySet{}
	default:
		b = &RecoveryAvailabilitySetCustomDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalRecoveryPlanActionDetailsClassification(rawMsg json.RawMessage) (RecoveryPlanActionDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RecoveryPlanActionDetailsClassification
	switch m["instanceType"] {
	case "AutomationRunbookActionDetails":
		b = &RecoveryPlanAutomationRunbookActionDetails{}
	case "ManualActionDetails":
		b = &RecoveryPlanManualActionDetails{}
	case "ScriptActionDetails":
		b = &RecoveryPlanScriptActionDetails{}
	default:
		b = &RecoveryPlanActionDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalRecoveryPlanProviderSpecificDetailsClassification(rawMsg json.RawMessage) (RecoveryPlanProviderSpecificDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RecoveryPlanProviderSpecificDetailsClassification
	switch m["instanceType"] {
	case "A2A":
		b = &RecoveryPlanA2ADetails{}
	default:
		b = &RecoveryPlanProviderSpecificDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalRecoveryPlanProviderSpecificDetailsClassificationArray(rawMsg json.RawMessage) ([]RecoveryPlanProviderSpecificDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]RecoveryPlanProviderSpecificDetailsClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalRecoveryPlanProviderSpecificDetailsClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalRecoveryPlanProviderSpecificFailoverInputClassification(rawMsg json.RawMessage) (RecoveryPlanProviderSpecificFailoverInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RecoveryPlanProviderSpecificFailoverInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &RecoveryPlanA2AFailoverInput{}
	case "HyperVReplicaAzure":
		b = &RecoveryPlanHyperVReplicaAzureFailoverInput{}
	case "HyperVReplicaAzureFailback":
		b = &RecoveryPlanHyperVReplicaAzureFailbackInput{}
	case "InMage":
		b = &RecoveryPlanInMageFailoverInput{}
	case "InMageAzureV2":
		b = &RecoveryPlanInMageAzureV2FailoverInput{}
	case "InMageRcm":
		b = &RecoveryPlanInMageRcmFailoverInput{}
	case "InMageRcmFailback":
		b = &RecoveryPlanInMageRcmFailbackFailoverInput{}
	default:
		b = &RecoveryPlanProviderSpecificFailoverInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalRecoveryPlanProviderSpecificFailoverInputClassificationArray(rawMsg json.RawMessage) ([]RecoveryPlanProviderSpecificFailoverInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]RecoveryPlanProviderSpecificFailoverInputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalRecoveryPlanProviderSpecificFailoverInputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalRecoveryPlanProviderSpecificInputClassification(rawMsg json.RawMessage) (RecoveryPlanProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RecoveryPlanProviderSpecificInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &RecoveryPlanA2AInput{}
	default:
		b = &RecoveryPlanProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalRecoveryPlanProviderSpecificInputClassificationArray(rawMsg json.RawMessage) ([]RecoveryPlanProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]RecoveryPlanProviderSpecificInputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalRecoveryPlanProviderSpecificInputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalRecoveryProximityPlacementGroupCustomDetailsClassification(rawMsg json.RawMessage) (RecoveryProximityPlacementGroupCustomDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RecoveryProximityPlacementGroupCustomDetailsClassification
	switch m["resourceType"] {
	case "Existing":
		b = &ExistingRecoveryProximityPlacementGroup{}
	default:
		b = &RecoveryProximityPlacementGroupCustomDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalRecoveryResourceGroupCustomDetailsClassification(rawMsg json.RawMessage) (RecoveryResourceGroupCustomDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RecoveryResourceGroupCustomDetailsClassification
	switch m["resourceType"] {
	case "Existing":
		b = &ExistingRecoveryResourceGroup{}
	default:
		b = &RecoveryResourceGroupCustomDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalRecoveryVirtualNetworkCustomDetailsClassification(rawMsg json.RawMessage) (RecoveryVirtualNetworkCustomDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RecoveryVirtualNetworkCustomDetailsClassification
	switch m["resourceType"] {
	case "Existing":
		b = &ExistingRecoveryVirtualNetwork{}
	case "New":
		b = &NewRecoveryVirtualNetwork{}
	default:
		b = &RecoveryVirtualNetworkCustomDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalRemoveDisksProviderSpecificInputClassification(rawMsg json.RawMessage) (RemoveDisksProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RemoveDisksProviderSpecificInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2ARemoveDisksInput{}
	default:
		b = &RemoveDisksProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalReplicationClusterProviderSpecificSettingsClassification(rawMsg json.RawMessage) (ReplicationClusterProviderSpecificSettingsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ReplicationClusterProviderSpecificSettingsClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2AReplicationProtectionClusterDetails{}
	default:
		b = &ReplicationClusterProviderSpecificSettings{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalReplicationProtectionIntentProviderSpecificSettingsClassification(rawMsg json.RawMessage) (ReplicationProtectionIntentProviderSpecificSettingsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ReplicationProtectionIntentProviderSpecificSettingsClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2AReplicationIntentDetails{}
	default:
		b = &ReplicationProtectionIntentProviderSpecificSettings{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalReplicationProviderSpecificContainerCreationInputClassification(rawMsg json.RawMessage) (ReplicationProviderSpecificContainerCreationInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ReplicationProviderSpecificContainerCreationInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2AContainerCreationInput{}
	case "A2ACrossClusterMigration":
		b = &A2ACrossClusterMigrationContainerCreationInput{}
	case "VMwareCbt":
		b = &VMwareCbtContainerCreationInput{}
	default:
		b = &ReplicationProviderSpecificContainerCreationInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalReplicationProviderSpecificContainerCreationInputClassificationArray(rawMsg json.RawMessage) ([]ReplicationProviderSpecificContainerCreationInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ReplicationProviderSpecificContainerCreationInputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalReplicationProviderSpecificContainerCreationInputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalReplicationProviderSpecificContainerMappingInputClassification(rawMsg json.RawMessage) (ReplicationProviderSpecificContainerMappingInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ReplicationProviderSpecificContainerMappingInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2AContainerMappingInput{}
	case "VMwareCbt":
		b = &VMwareCbtContainerMappingInput{}
	default:
		b = &ReplicationProviderSpecificContainerMappingInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalReplicationProviderSpecificSettingsClassification(rawMsg json.RawMessage) (ReplicationProviderSpecificSettingsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ReplicationProviderSpecificSettingsClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2AReplicationDetails{}
	case "A2ACrossClusterMigration":
		b = &A2ACrossClusterMigrationReplicationDetails{}
	case "HyperVReplica2012":
		b = &HyperVReplicaReplicationDetails{}
	case "HyperVReplica2012R2":
		b = &HyperVReplicaBlueReplicationDetails{}
	case "HyperVReplicaAzure":
		b = &HyperVReplicaAzureReplicationDetails{}
	case "HyperVReplicaBaseReplicationDetails":
		b = &HyperVReplicaBaseReplicationDetails{}
	case "InMage":
		b = &InMageReplicationDetails{}
	case "InMageAzureV2":
		b = &InMageAzureV2ReplicationDetails{}
	case "InMageRcm":
		b = &InMageRcmReplicationDetails{}
	case "InMageRcmFailback":
		b = &InMageRcmFailbackReplicationDetails{}
	default:
		b = &ReplicationProviderSpecificSettings{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalReplicationProviderSpecificUpdateContainerMappingInputClassification(rawMsg json.RawMessage) (ReplicationProviderSpecificUpdateContainerMappingInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ReplicationProviderSpecificUpdateContainerMappingInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2AUpdateContainerMappingInput{}
	case "InMageRcm":
		b = &InMageRcmUpdateContainerMappingInput{}
	default:
		b = &ReplicationProviderSpecificUpdateContainerMappingInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalResumeReplicationProviderSpecificInputClassification(rawMsg json.RawMessage) (ResumeReplicationProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ResumeReplicationProviderSpecificInputClassification
	switch m["instanceType"] {
	case "VMwareCbt":
		b = &VMwareCbtResumeReplicationInput{}
	default:
		b = &ResumeReplicationProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalResyncProviderSpecificInputClassification(rawMsg json.RawMessage) (ResyncProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ResyncProviderSpecificInputClassification
	switch m["instanceType"] {
	case "VMwareCbt":
		b = &VMwareCbtResyncInput{}
	default:
		b = &ResyncProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalReverseReplicationProviderSpecificInputClassification(rawMsg json.RawMessage) (ReverseReplicationProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ReverseReplicationProviderSpecificInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2AReprotectInput{}
	case "HyperVReplicaAzure":
		b = &HyperVReplicaAzureReprotectInput{}
	case "InMage":
		b = &InMageReprotectInput{}
	case "InMageAzureV2":
		b = &InMageAzureV2ReprotectInput{}
	case "InMageRcm":
		b = &InMageRcmReprotectInput{}
	case "InMageRcmFailback":
		b = &InMageRcmFailbackReprotectInput{}
	default:
		b = &ReverseReplicationProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSharedDiskReplicationProviderSpecificSettingsClassification(rawMsg json.RawMessage) (SharedDiskReplicationProviderSpecificSettingsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SharedDiskReplicationProviderSpecificSettingsClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2ASharedDiskReplicationDetails{}
	default:
		b = &SharedDiskReplicationProviderSpecificSettings{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalStorageAccountCustomDetailsClassification(rawMsg json.RawMessage) (StorageAccountCustomDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b StorageAccountCustomDetailsClassification
	switch m["resourceType"] {
	case "Existing":
		b = &ExistingStorageAccount{}
	default:
		b = &StorageAccountCustomDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSwitchProtectionProviderSpecificInputClassification(rawMsg json.RawMessage) (SwitchProtectionProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SwitchProtectionProviderSpecificInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2ASwitchProtectionInput{}
	default:
		b = &SwitchProtectionProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSwitchProviderSpecificInputClassification(rawMsg json.RawMessage) (SwitchProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SwitchProviderSpecificInputClassification
	switch m["instanceType"] {
	case "InMageAzureV2":
		b = &InMageAzureV2SwitchProviderInput{}
	default:
		b = &SwitchProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalTaskTypeDetailsClassification(rawMsg json.RawMessage) (TaskTypeDetailsClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TaskTypeDetailsClassification
	switch m["instanceType"] {
	case "AutomationRunbookTaskDetails":
		b = &AutomationRunbookTaskDetails{}
	case "ConsistencyCheckTaskDetails":
		b = &ConsistencyCheckTaskDetails{}
	case "FabricReplicationGroupTaskDetails":
		b = &FabricReplicationGroupTaskDetails{}
	case "JobTaskDetails":
		b = &JobTaskDetails{}
	case "ManualActionTaskDetails":
		b = &ManualActionTaskDetails{}
	case "ScriptActionTaskDetails":
		b = &ScriptActionTaskDetails{}
	case "VirtualMachineTaskDetails":
		b = &VirtualMachineTaskDetails{}
	case "VmNicUpdatesTaskDetails":
		b = &VMNicUpdatesTaskDetails{}
	default:
		b = &TaskTypeDetails{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalTestFailoverProviderSpecificInputClassification(rawMsg json.RawMessage) (TestFailoverProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TestFailoverProviderSpecificInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2ATestFailoverInput{}
	case "HyperVReplicaAzure":
		b = &HyperVReplicaAzureTestFailoverInput{}
	case "InMage":
		b = &InMageTestFailoverInput{}
	case "InMageAzureV2":
		b = &InMageAzureV2TestFailoverInput{}
	case "InMageRcm":
		b = &InMageRcmTestFailoverInput{}
	default:
		b = &TestFailoverProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalTestMigrateProviderSpecificInputClassification(rawMsg json.RawMessage) (TestMigrateProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TestMigrateProviderSpecificInputClassification
	switch m["instanceType"] {
	case "VMwareCbt":
		b = &VMwareCbtTestMigrateInput{}
	default:
		b = &TestMigrateProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalUnplannedFailoverProviderSpecificInputClassification(rawMsg json.RawMessage) (UnplannedFailoverProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b UnplannedFailoverProviderSpecificInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2AUnplannedFailoverInput{}
	case "HyperVReplicaAzure":
		b = &HyperVReplicaAzureUnplannedFailoverInput{}
	case "InMage":
		b = &InMageUnplannedFailoverInput{}
	case "InMageAzureV2":
		b = &InMageAzureV2UnplannedFailoverInput{}
	case "InMageRcm":
		b = &InMageRcmUnplannedFailoverInput{}
	default:
		b = &UnplannedFailoverProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalUpdateApplianceForReplicationProtectedItemProviderSpecificInputClassification(rawMsg json.RawMessage) (UpdateApplianceForReplicationProtectedItemProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b UpdateApplianceForReplicationProtectedItemProviderSpecificInputClassification
	switch m["instanceType"] {
	case "InMageRcm":
		b = &InMageRcmUpdateApplianceForReplicationProtectedItemInput{}
	default:
		b = &UpdateApplianceForReplicationProtectedItemProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalUpdateMigrationItemProviderSpecificInputClassification(rawMsg json.RawMessage) (UpdateMigrationItemProviderSpecificInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b UpdateMigrationItemProviderSpecificInputClassification
	switch m["instanceType"] {
	case "VMwareCbt":
		b = &VMwareCbtUpdateMigrationItemInput{}
	default:
		b = &UpdateMigrationItemProviderSpecificInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalUpdateReplicationProtectedItemProviderInputClassification(rawMsg json.RawMessage) (UpdateReplicationProtectedItemProviderInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b UpdateReplicationProtectedItemProviderInputClassification
	switch m["instanceType"] {
	case "A2A":
		b = &A2AUpdateReplicationProtectedItemInput{}
	case "HyperVReplicaAzure":
		b = &HyperVReplicaAzureUpdateReplicationProtectedItemInput{}
	case "InMageAzureV2":
		b = &InMageAzureV2UpdateReplicationProtectedItemInput{}
	case "InMageRcm":
		b = &InMageRcmUpdateReplicationProtectedItemInput{}
	default:
		b = &UpdateReplicationProtectedItemProviderInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}
