//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

const (
	moduleName    = "armhybridcompute"
	moduleVersion = "v2.0.0-beta.1"
)

// AgentConfigurationMode - Name of configuration mode to use. Modes are pre-defined configurations of security controls,
// extension allowlists and guest configuration, maintained by Microsoft.
type AgentConfigurationMode string

const (
	AgentConfigurationModeFull    AgentConfigurationMode = "full"
	AgentConfigurationModeMonitor AgentConfigurationMode = "monitor"
)

// PossibleAgentConfigurationModeValues returns the possible values for the AgentConfigurationMode const type.
func PossibleAgentConfigurationModeValues() []AgentConfigurationMode {
	return []AgentConfigurationMode{
		AgentConfigurationModeFull,
		AgentConfigurationModeMonitor,
	}
}

// AssessmentModeTypes - Specifies the assessment mode.
type AssessmentModeTypes string

const (
	AssessmentModeTypesAutomaticByPlatform AssessmentModeTypes = "AutomaticByPlatform"
	AssessmentModeTypesImageDefault        AssessmentModeTypes = "ImageDefault"
)

// PossibleAssessmentModeTypesValues returns the possible values for the AssessmentModeTypes const type.
func PossibleAssessmentModeTypesValues() []AssessmentModeTypes {
	return []AssessmentModeTypes{
		AssessmentModeTypesAutomaticByPlatform,
		AssessmentModeTypesImageDefault,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

type InstanceViewTypes string

const (
	InstanceViewTypesInstanceView InstanceViewTypes = "instanceView"
)

// PossibleInstanceViewTypesValues returns the possible values for the InstanceViewTypes const type.
func PossibleInstanceViewTypesValues() []InstanceViewTypes {
	return []InstanceViewTypes{
		InstanceViewTypesInstanceView,
	}
}

// LastAttemptStatusEnum - Specifies the status of Agent Upgrade.
type LastAttemptStatusEnum string

const (
	LastAttemptStatusEnumFailed  LastAttemptStatusEnum = "Failed"
	LastAttemptStatusEnumSuccess LastAttemptStatusEnum = "Success"
)

// PossibleLastAttemptStatusEnumValues returns the possible values for the LastAttemptStatusEnum const type.
func PossibleLastAttemptStatusEnumValues() []LastAttemptStatusEnum {
	return []LastAttemptStatusEnum{
		LastAttemptStatusEnumFailed,
		LastAttemptStatusEnumSuccess,
	}
}

// PatchModeTypes - Specifies the patch mode.
type PatchModeTypes string

const (
	PatchModeTypesAutomaticByOS       PatchModeTypes = "AutomaticByOS"
	PatchModeTypesAutomaticByPlatform PatchModeTypes = "AutomaticByPlatform"
	PatchModeTypesImageDefault        PatchModeTypes = "ImageDefault"
	PatchModeTypesManual              PatchModeTypes = "Manual"
)

// PossiblePatchModeTypesValues returns the possible values for the PatchModeTypes const type.
func PossiblePatchModeTypesValues() []PatchModeTypes {
	return []PatchModeTypes{
		PatchModeTypesAutomaticByOS,
		PatchModeTypesAutomaticByPlatform,
		PatchModeTypesImageDefault,
		PatchModeTypesManual,
	}
}

// PrivateCloudKind - Indicates which kind of VM fabric the instance is an instance of, such as HCI or SCVMM etc.
type PrivateCloudKind string

const (
	PrivateCloudKindAVS    PrivateCloudKind = "AVS"
	PrivateCloudKindHCI    PrivateCloudKind = "HCI"
	PrivateCloudKindSCVMM  PrivateCloudKind = "SCVMM"
	PrivateCloudKindVMware PrivateCloudKind = "VMware"
)

// PossiblePrivateCloudKindValues returns the possible values for the PrivateCloudKind const type.
func PossiblePrivateCloudKindValues() []PrivateCloudKind {
	return []PrivateCloudKind{
		PrivateCloudKindAVS,
		PrivateCloudKindHCI,
		PrivateCloudKindSCVMM,
		PrivateCloudKindVMware,
	}
}

// PublicNetworkAccessType - The network access policy to determine if Azure Arc agents can use public Azure Arc service endpoints.
// Defaults to disabled (access to Azure Arc services only via private link).
type PublicNetworkAccessType string

const (
	// PublicNetworkAccessTypeDisabled - Does not allow Azure Arc agents to communicate with Azure Arc services over public (internet)
	// endpoints. The agents must use the private link.
	PublicNetworkAccessTypeDisabled PublicNetworkAccessType = "Disabled"
	// PublicNetworkAccessTypeEnabled - Allows Azure Arc agents to communicate with Azure Arc services over both public (internet)
	// and private endpoints.
	PublicNetworkAccessTypeEnabled PublicNetworkAccessType = "Enabled"
)

// PossiblePublicNetworkAccessTypeValues returns the possible values for the PublicNetworkAccessType const type.
func PossiblePublicNetworkAccessTypeValues() []PublicNetworkAccessType {
	return []PublicNetworkAccessType{
		PublicNetworkAccessTypeDisabled,
		PublicNetworkAccessTypeEnabled,
	}
}

// StatusLevelTypes - The level code.
type StatusLevelTypes string

const (
	StatusLevelTypesError   StatusLevelTypes = "Error"
	StatusLevelTypesInfo    StatusLevelTypes = "Info"
	StatusLevelTypesWarning StatusLevelTypes = "Warning"
)

// PossibleStatusLevelTypesValues returns the possible values for the StatusLevelTypes const type.
func PossibleStatusLevelTypesValues() []StatusLevelTypes {
	return []StatusLevelTypes{
		StatusLevelTypesError,
		StatusLevelTypesInfo,
		StatusLevelTypesWarning,
	}
}

// StatusTypes - The status of the hybrid machine agent.
type StatusTypes string

const (
	StatusTypesConnected    StatusTypes = "Connected"
	StatusTypesDisconnected StatusTypes = "Disconnected"
	StatusTypesError        StatusTypes = "Error"
)

// PossibleStatusTypesValues returns the possible values for the StatusTypes const type.
func PossibleStatusTypesValues() []StatusTypes {
	return []StatusTypes{
		StatusTypesConnected,
		StatusTypesDisconnected,
		StatusTypesError,
	}
}
