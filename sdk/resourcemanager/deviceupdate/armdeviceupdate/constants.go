//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdeviceupdate

const (
	moduleName    = "armdeviceupdate"
	moduleVersion = "v1.0.1"
)

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

// AuthenticationType - Authentication Type
type AuthenticationType string

const (
	AuthenticationTypeKeyBased AuthenticationType = "KeyBased"
)

// PossibleAuthenticationTypeValues returns the possible values for the AuthenticationType const type.
func PossibleAuthenticationTypeValues() []AuthenticationType {
	return []AuthenticationType{
		AuthenticationTypeKeyBased,
	}
}

// CheckNameAvailabilityReason - The reason why the given name is not available.
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

// GroupIDProvisioningState - The provisioning state of private link group ID.
type GroupIDProvisioningState string

const (
	GroupIDProvisioningStateCanceled  GroupIDProvisioningState = "Canceled"
	GroupIDProvisioningStateFailed    GroupIDProvisioningState = "Failed"
	GroupIDProvisioningStateSucceeded GroupIDProvisioningState = "Succeeded"
)

// PossibleGroupIDProvisioningStateValues returns the possible values for the GroupIDProvisioningState const type.
func PossibleGroupIDProvisioningStateValues() []GroupIDProvisioningState {
	return []GroupIDProvisioningState{
		GroupIDProvisioningStateCanceled,
		GroupIDProvisioningStateFailed,
		GroupIDProvisioningStateSucceeded,
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

// PrivateEndpointConnectionProxyProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProxyProvisioningState string

const (
	PrivateEndpointConnectionProxyProvisioningStateCreating  PrivateEndpointConnectionProxyProvisioningState = "Creating"
	PrivateEndpointConnectionProxyProvisioningStateDeleting  PrivateEndpointConnectionProxyProvisioningState = "Deleting"
	PrivateEndpointConnectionProxyProvisioningStateFailed    PrivateEndpointConnectionProxyProvisioningState = "Failed"
	PrivateEndpointConnectionProxyProvisioningStateSucceeded PrivateEndpointConnectionProxyProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProxyProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProxyProvisioningState const type.
func PossiblePrivateEndpointConnectionProxyProvisioningStateValues() []PrivateEndpointConnectionProxyProvisioningState {
	return []PrivateEndpointConnectionProxyProvisioningState{
		PrivateEndpointConnectionProxyProvisioningStateCreating,
		PrivateEndpointConnectionProxyProvisioningStateDeleting,
		PrivateEndpointConnectionProxyProvisioningStateFailed,
		PrivateEndpointConnectionProxyProvisioningStateSucceeded,
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

// ProvisioningState - Provisioning state.
type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleted   ProvisioningState = "Deleted"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
	}
}

// PublicNetworkAccess - Whether or not public network access is allowed for the account.
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

// Role - Whether the location is primary or failover
type Role string

const (
	RoleFailover Role = "Failover"
	RolePrimary  Role = "Primary"
)

// PossibleRoleValues returns the possible values for the Role const type.
func PossibleRoleValues() []Role {
	return []Role{
		RoleFailover,
		RolePrimary,
	}
}

// SKU - Device Update Sku
type SKU string

const (
	SKUFree     SKU = "Free"
	SKUStandard SKU = "Standard"
)

// PossibleSKUValues returns the possible values for the SKU const type.
func PossibleSKUValues() []SKU {
	return []SKU{
		SKUFree,
		SKUStandard,
	}
}
