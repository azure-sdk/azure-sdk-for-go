//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

const (
	moduleName    = "armhybridcompute"
	moduleVersion = "v1.0.0"
)

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
