//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armqumulo

const (
	moduleName    = "armqumulo"
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

// MarketplaceSubscriptionStatus - Marketplace subscription status of the file system resource
type MarketplaceSubscriptionStatus string

const (
	// MarketplaceSubscriptionStatusPendingFulfillmentStart - Fulfillment has not started
	MarketplaceSubscriptionStatusPendingFulfillmentStart MarketplaceSubscriptionStatus = "PendingFulfillmentStart"
	// MarketplaceSubscriptionStatusSubscribed - Marketplace offer is subscribed
	MarketplaceSubscriptionStatusSubscribed MarketplaceSubscriptionStatus = "Subscribed"
	// MarketplaceSubscriptionStatusSuspended - Marketplace offer is suspended because of non payment
	MarketplaceSubscriptionStatusSuspended MarketplaceSubscriptionStatus = "Suspended"
	// MarketplaceSubscriptionStatusUnsubscribed - Marketplace offer is unsubscribed
	MarketplaceSubscriptionStatusUnsubscribed MarketplaceSubscriptionStatus = "Unsubscribed"
)

// PossibleMarketplaceSubscriptionStatusValues returns the possible values for the MarketplaceSubscriptionStatus const type.
func PossibleMarketplaceSubscriptionStatusValues() []MarketplaceSubscriptionStatus {
	return []MarketplaceSubscriptionStatus{
		MarketplaceSubscriptionStatusPendingFulfillmentStart,
		MarketplaceSubscriptionStatusSubscribed,
		MarketplaceSubscriptionStatusSuspended,
		MarketplaceSubscriptionStatusUnsubscribed,
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

// ProvisioningState - Provisioning State of the File system resource
type ProvisioningState string

const (
	// ProvisioningStateAccepted - File system resource creation request accepted
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - File system resource creation canceled
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreating - File system resource creation started
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleted - File system resource is deleted
	ProvisioningStateDeleted ProvisioningState = "Deleted"
	// ProvisioningStateDeleting - File system resource deletion started
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - File system resource creation failed
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateNotSpecified - File system resource state is unknown
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	// ProvisioningStateSucceeded - File system resource creation successful
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - File system resource is being updated
	ProvisioningStateUpdating ProvisioningState = "Updating"
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

// StorageSKU - Storage Sku
type StorageSKU string

const (
	// StorageSKUPerformance - Performance Storage Sku
	StorageSKUPerformance StorageSKU = "Performance"
	// StorageSKUStandard - Standard Storage Sku
	StorageSKUStandard StorageSKU = "Standard"
)

// PossibleStorageSKUValues returns the possible values for the StorageSKU const type.
func PossibleStorageSKUValues() []StorageSKU {
	return []StorageSKU{
		StorageSKUPerformance,
		StorageSKUStandard,
	}
}
