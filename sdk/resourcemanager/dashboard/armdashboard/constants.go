//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdashboard

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard"
	moduleVersion = "v1.3.0"
)

type APIKey string

const (
	APIKeyDisabled APIKey = "Disabled"
	APIKeyEnabled  APIKey = "Enabled"
)

// PossibleAPIKeyValues returns the possible values for the APIKey const type.
func PossibleAPIKeyValues() []APIKey {
	return []APIKey{
		APIKeyDisabled,
		APIKeyEnabled,
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

// AutoGeneratedDomainNameLabelScope - Scope for dns deterministic name hash calculation
type AutoGeneratedDomainNameLabelScope string

const (
	AutoGeneratedDomainNameLabelScopeTenantReuse AutoGeneratedDomainNameLabelScope = "TenantReuse"
)

// PossibleAutoGeneratedDomainNameLabelScopeValues returns the possible values for the AutoGeneratedDomainNameLabelScope const type.
func PossibleAutoGeneratedDomainNameLabelScopeValues() []AutoGeneratedDomainNameLabelScope {
	return []AutoGeneratedDomainNameLabelScope{
		AutoGeneratedDomainNameLabelScopeTenantReuse,
	}
}

type AvailablePromotion string

const (
	AvailablePromotionFreeTrial AvailablePromotion = "FreeTrial"
	AvailablePromotionNone      AvailablePromotion = "None"
)

// PossibleAvailablePromotionValues returns the possible values for the AvailablePromotion const type.
func PossibleAvailablePromotionValues() []AvailablePromotion {
	return []AvailablePromotion{
		AvailablePromotionFreeTrial,
		AvailablePromotionNone,
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

type DeterministicOutboundIP string

const (
	DeterministicOutboundIPDisabled DeterministicOutboundIP = "Disabled"
	DeterministicOutboundIPEnabled  DeterministicOutboundIP = "Enabled"
)

// PossibleDeterministicOutboundIPValues returns the possible values for the DeterministicOutboundIP const type.
func PossibleDeterministicOutboundIPValues() []DeterministicOutboundIP {
	return []DeterministicOutboundIP{
		DeterministicOutboundIPDisabled,
		DeterministicOutboundIPEnabled,
	}
}

// ManagedPrivateEndpointConnectionStatus - The approval/rejection status of managed private endpoint connection.
type ManagedPrivateEndpointConnectionStatus string

const (
	ManagedPrivateEndpointConnectionStatusApproved     ManagedPrivateEndpointConnectionStatus = "Approved"
	ManagedPrivateEndpointConnectionStatusDisconnected ManagedPrivateEndpointConnectionStatus = "Disconnected"
	ManagedPrivateEndpointConnectionStatusPending      ManagedPrivateEndpointConnectionStatus = "Pending"
	ManagedPrivateEndpointConnectionStatusRejected     ManagedPrivateEndpointConnectionStatus = "Rejected"
)

// PossibleManagedPrivateEndpointConnectionStatusValues returns the possible values for the ManagedPrivateEndpointConnectionStatus const type.
func PossibleManagedPrivateEndpointConnectionStatusValues() []ManagedPrivateEndpointConnectionStatus {
	return []ManagedPrivateEndpointConnectionStatus{
		ManagedPrivateEndpointConnectionStatusApproved,
		ManagedPrivateEndpointConnectionStatusDisconnected,
		ManagedPrivateEndpointConnectionStatusPending,
		ManagedPrivateEndpointConnectionStatusRejected,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
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

// MarketplaceAutoRenew - The AutoRenew setting of the Enterprise subscription
type MarketplaceAutoRenew string

const (
	MarketplaceAutoRenewDisabled MarketplaceAutoRenew = "Disabled"
	MarketplaceAutoRenewEnabled  MarketplaceAutoRenew = "Enabled"
)

// PossibleMarketplaceAutoRenewValues returns the possible values for the MarketplaceAutoRenew const type.
func PossibleMarketplaceAutoRenewValues() []MarketplaceAutoRenew {
	return []MarketplaceAutoRenew{
		MarketplaceAutoRenewDisabled,
		MarketplaceAutoRenewEnabled,
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

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateNotSpecified,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// PublicNetworkAccess - Indicate the state for enable or disable traffic over the public interface.
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// StartTLSPolicy - The StartTLSPolicy setting of the SMTP configuration https://pkg.go.dev/github.com/go-mail/mail#StartTLSPolicy
type StartTLSPolicy string

const (
	StartTLSPolicyMandatoryStartTLS     StartTLSPolicy = "MandatoryStartTLS"
	StartTLSPolicyNoStartTLS            StartTLSPolicy = "NoStartTLS"
	StartTLSPolicyOpportunisticStartTLS StartTLSPolicy = "OpportunisticStartTLS"
)

// PossibleStartTLSPolicyValues returns the possible values for the StartTLSPolicy const type.
func PossibleStartTLSPolicyValues() []StartTLSPolicy {
	return []StartTLSPolicy{
		StartTLSPolicyMandatoryStartTLS,
		StartTLSPolicyNoStartTLS,
		StartTLSPolicyOpportunisticStartTLS,
	}
}

type ZoneRedundancy string

const (
	ZoneRedundancyDisabled ZoneRedundancy = "Disabled"
	ZoneRedundancyEnabled  ZoneRedundancy = "Enabled"
)

// PossibleZoneRedundancyValues returns the possible values for the ZoneRedundancy const type.
func PossibleZoneRedundancyValues() []ZoneRedundancy {
	return []ZoneRedundancy{
		ZoneRedundancyDisabled,
		ZoneRedundancyEnabled,
	}
}
