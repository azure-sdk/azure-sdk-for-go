//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armastro

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/astro/armastro"
	moduleVersion = "v1.0.0"
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

// MarketplaceSubscriptionStatus - Marketplace subscription status of a resource.
type MarketplaceSubscriptionStatus string

const (
	// MarketplaceSubscriptionStatusPendingFulfillmentStart - Purchased but not yet activated
	MarketplaceSubscriptionStatusPendingFulfillmentStart MarketplaceSubscriptionStatus = "PendingFulfillmentStart"
	// MarketplaceSubscriptionStatusSubscribed - Marketplace subscription is activated
	MarketplaceSubscriptionStatusSubscribed MarketplaceSubscriptionStatus = "Subscribed"
	// MarketplaceSubscriptionStatusSuspended - This state indicates that a customer's payment for the Marketplace service was
	// not received
	MarketplaceSubscriptionStatusSuspended MarketplaceSubscriptionStatus = "Suspended"
	// MarketplaceSubscriptionStatusUnsubscribed - Customer has cancelled the subscription
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

// ResourceProvisioningState - The provisioning state of a resource type.
type ResourceProvisioningState string

const (
	// ResourceProvisioningStateCanceled - Resource creation was canceled.
	ResourceProvisioningStateCanceled ResourceProvisioningState = "Canceled"
	// ResourceProvisioningStateFailed - Resource creation failed.
	ResourceProvisioningStateFailed ResourceProvisioningState = "Failed"
	// ResourceProvisioningStateSucceeded - Resource has been created.
	ResourceProvisioningStateSucceeded ResourceProvisioningState = "Succeeded"
)

// PossibleResourceProvisioningStateValues returns the possible values for the ResourceProvisioningState const type.
func PossibleResourceProvisioningStateValues() []ResourceProvisioningState {
	return []ResourceProvisioningState{
		ResourceProvisioningStateCanceled,
		ResourceProvisioningStateFailed,
		ResourceProvisioningStateSucceeded,
	}
}

// SingleSignOnStates - Various states of the SSO resource
type SingleSignOnStates string

const (
	// SingleSignOnStatesDisable - State of the SSO resource when it is disabled
	SingleSignOnStatesDisable SingleSignOnStates = "Disable"
	// SingleSignOnStatesEnable - State of the SSO resource when it is enabled
	SingleSignOnStatesEnable SingleSignOnStates = "Enable"
	// SingleSignOnStatesInitial - Initial state of the SSO resource
	SingleSignOnStatesInitial SingleSignOnStates = "Initial"
)

// PossibleSingleSignOnStatesValues returns the possible values for the SingleSignOnStates const type.
func PossibleSingleSignOnStatesValues() []SingleSignOnStates {
	return []SingleSignOnStates{
		SingleSignOnStatesDisable,
		SingleSignOnStatesEnable,
		SingleSignOnStatesInitial,
	}
}
