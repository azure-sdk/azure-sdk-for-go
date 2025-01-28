//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvoiceservices

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/voiceservices/armvoiceservices"
	moduleVersion = "v1.1.0"
)

// APIBridgeActivationState - API Bridge activation state.
type APIBridgeActivationState string

const (
	// APIBridgeActivationStateDisabled - API Bridge is disabled
	APIBridgeActivationStateDisabled APIBridgeActivationState = "disabled"
	// APIBridgeActivationStateEnabled - API Bridge is enabled
	APIBridgeActivationStateEnabled APIBridgeActivationState = "enabled"
)

// PossibleAPIBridgeActivationStateValues returns the possible values for the APIBridgeActivationState const type.
func PossibleAPIBridgeActivationStateValues() []APIBridgeActivationState {
	return []APIBridgeActivationState{
		APIBridgeActivationStateDisabled,
		APIBridgeActivationStateEnabled,
	}
}

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// AutoGeneratedDomainNameLabelScope - Available auto-generated domain name scopes.
type AutoGeneratedDomainNameLabelScope string

const (
	// AutoGeneratedDomainNameLabelScopeNoReuse - Generated domain name label is always unique.
	AutoGeneratedDomainNameLabelScopeNoReuse AutoGeneratedDomainNameLabelScope = "NoReuse"
	// AutoGeneratedDomainNameLabelScopeResourceGroupReuse - Generated domain name label depends on resource name, tenant ID,
	// subscription ID and resource group name.
	AutoGeneratedDomainNameLabelScopeResourceGroupReuse AutoGeneratedDomainNameLabelScope = "ResourceGroupReuse"
	// AutoGeneratedDomainNameLabelScopeSubscriptionReuse - Generated domain name label depends on resource name, tenant ID and
	// subscription ID.
	AutoGeneratedDomainNameLabelScopeSubscriptionReuse AutoGeneratedDomainNameLabelScope = "SubscriptionReuse"
	// AutoGeneratedDomainNameLabelScopeTenantReuse - Generated domain name label depends on resource name and tenant ID.
	AutoGeneratedDomainNameLabelScopeTenantReuse AutoGeneratedDomainNameLabelScope = "TenantReuse"
)

// PossibleAutoGeneratedDomainNameLabelScopeValues returns the possible values for the AutoGeneratedDomainNameLabelScope const type.
func PossibleAutoGeneratedDomainNameLabelScopeValues() []AutoGeneratedDomainNameLabelScope {
	return []AutoGeneratedDomainNameLabelScope{
		AutoGeneratedDomainNameLabelScopeNoReuse,
		AutoGeneratedDomainNameLabelScopeResourceGroupReuse,
		AutoGeneratedDomainNameLabelScopeSubscriptionReuse,
		AutoGeneratedDomainNameLabelScopeTenantReuse,
	}
}

// CheckNameAvailabilityReason - Possible reasons for a name not being available.
type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns the possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{
		CheckNameAvailabilityReasonAlreadyExists,
		CheckNameAvailabilityReasonInvalid,
	}
}

// CommunicationsPlatform - Available platform types.
type CommunicationsPlatform string

const (
	// CommunicationsPlatformOperatorConnect - Operator Connect
	CommunicationsPlatformOperatorConnect CommunicationsPlatform = "OperatorConnect"
	// CommunicationsPlatformTeamsDirectRouting - Teams Direct Routing
	CommunicationsPlatformTeamsDirectRouting CommunicationsPlatform = "TeamsDirectRouting"
	// CommunicationsPlatformTeamsPhoneMobile - Teams Phone Mobile
	CommunicationsPlatformTeamsPhoneMobile CommunicationsPlatform = "TeamsPhoneMobile"
)

// PossibleCommunicationsPlatformValues returns the possible values for the CommunicationsPlatform const type.
func PossibleCommunicationsPlatformValues() []CommunicationsPlatform {
	return []CommunicationsPlatform{
		CommunicationsPlatformOperatorConnect,
		CommunicationsPlatformTeamsDirectRouting,
		CommunicationsPlatformTeamsPhoneMobile,
	}
}

// Connectivity - How this deployment connects back to the operator network
type Connectivity string

const (
	// ConnectivityPublicAddress - This deployment connects to the operator network using a Public IP address, e.g. when using
	// MAPS
	ConnectivityPublicAddress Connectivity = "PublicAddress"
)

// PossibleConnectivityValues returns the possible values for the Connectivity const type.
func PossibleConnectivityValues() []Connectivity {
	return []Connectivity{
		ConnectivityPublicAddress,
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

// E911Type - The method for terminating emergency calls to the PSTN.
type E911Type string

const (
	// E911TypeDirectToEsrp - Emergency calls are routed directly to the ESRP
	E911TypeDirectToEsrp E911Type = "DirectToEsrp"
	// E911TypeStandard - Emergency calls are not handled different from other calls
	E911TypeStandard E911Type = "Standard"
)

// PossibleE911TypeValues returns the possible values for the E911Type const type.
func PossibleE911TypeValues() []E911Type {
	return []E911Type{
		E911TypeDirectToEsrp,
		E911TypeStandard,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned, UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ProvisioningState - Provisioning state of the resource.
type ProvisioningState string

const (
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
	}
}

// SKUTier - This field is required to be implemented by the Resource Provider if the service has more than one tier, but
// is not required on a PUT.
type SKUTier string

const (
	SKUTierBasic    SKUTier = "Basic"
	SKUTierFree     SKUTier = "Free"
	SKUTierPremium  SKUTier = "Premium"
	SKUTierStandard SKUTier = "Standard"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBasic,
		SKUTierFree,
		SKUTierPremium,
		SKUTierStandard,
	}
}

// Status - The status of the current CommunicationsGateway resource.
type Status string

const (
	// StatusChangePending - The resource has been created or updated, but the CommunicationsGateway service has not yet been
	// updated to reflect the changes.
	StatusChangePending Status = "ChangePending"
	// StatusComplete - The CommunicationsGateway service is up and running with the parameters specified in the resource.
	StatusComplete Status = "Complete"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusChangePending,
		StatusComplete,
	}
}

// TeamsCodecs - The voice codecs expected for communication with Teams.
type TeamsCodecs string

const (
	// TeamsCodecsG722 - G.722 wideband audio codec
	TeamsCodecsG722 TeamsCodecs = "G722"
	// TeamsCodecsG7222 - G.722.2 wideband audio codec
	TeamsCodecsG7222 TeamsCodecs = "G722_2"
	// TeamsCodecsPCMA - Pulse code modulation(PCM) U-law narrowband audio codec(G.711u)
	TeamsCodecsPCMA TeamsCodecs = "PCMA"
	// TeamsCodecsPCMU - Pulse code modulation(PCM) U-law narrowband audio codec(G.711u)
	TeamsCodecsPCMU TeamsCodecs = "PCMU"
	// TeamsCodecsSILK16 - SILK/16000 wideband audio codec
	TeamsCodecsSILK16 TeamsCodecs = "SILK_16"
	// TeamsCodecsSILK8 - SILK/8000 narrowband audio codec
	TeamsCodecsSILK8 TeamsCodecs = "SILK_8"
)

// PossibleTeamsCodecsValues returns the possible values for the TeamsCodecs const type.
func PossibleTeamsCodecsValues() []TeamsCodecs {
	return []TeamsCodecs{
		TeamsCodecsG722,
		TeamsCodecsG7222,
		TeamsCodecsPCMA,
		TeamsCodecsPCMU,
		TeamsCodecsSILK16,
		TeamsCodecsSILK8,
	}
}

// TestLinePurpose - The purpose of the TestLine resource.
type TestLinePurpose string

const (
	// TestLinePurposeAutomated - The test line is used for automated testing
	TestLinePurposeAutomated TestLinePurpose = "Automated"
	// TestLinePurposeManual - The test line is used for manual testing
	TestLinePurposeManual TestLinePurpose = "Manual"
)

// PossibleTestLinePurposeValues returns the possible values for the TestLinePurpose const type.
func PossibleTestLinePurposeValues() []TestLinePurpose {
	return []TestLinePurpose{
		TestLinePurposeAutomated,
		TestLinePurposeManual,
	}
}
