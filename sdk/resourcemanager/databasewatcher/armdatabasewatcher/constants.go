//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabasewatcher

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasewatcher/armdatabasewatcher"
	moduleVersion = "v0.1.0"
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

// DatabaseWatcherProvisioningState - The status of the last provisioning operation performed on the resource.
type DatabaseWatcherProvisioningState string

const (
	// DatabaseWatcherProvisioningStateCanceled - Resource creation was canceled.
	DatabaseWatcherProvisioningStateCanceled DatabaseWatcherProvisioningState = "Canceled"
	// DatabaseWatcherProvisioningStateFailed - Resource creation failed.
	DatabaseWatcherProvisioningStateFailed DatabaseWatcherProvisioningState = "Failed"
	// DatabaseWatcherProvisioningStateSucceeded - Resource has been created.
	DatabaseWatcherProvisioningStateSucceeded DatabaseWatcherProvisioningState = "Succeeded"
)

// PossibleDatabaseWatcherProvisioningStateValues returns the possible values for the DatabaseWatcherProvisioningState const type.
func PossibleDatabaseWatcherProvisioningStateValues() []DatabaseWatcherProvisioningState {
	return []DatabaseWatcherProvisioningState{
		DatabaseWatcherProvisioningStateCanceled,
		DatabaseWatcherProvisioningStateFailed,
		DatabaseWatcherProvisioningStateSucceeded,
	}
}

// KustoOfferingType - The type of Kusto offering.
type KustoOfferingType string

const (
	// KustoOfferingTypeAdx - The Azure Data Explorer cluster Kusto offering.
	KustoOfferingTypeAdx KustoOfferingType = "adx"
	// KustoOfferingTypeFabric - The Fabric Real-Time Analytics Kusto offering.
	KustoOfferingTypeFabric KustoOfferingType = "fabric"
	// KustoOfferingTypeFree - The free Azure Data Explorer cluster Kusto offering.
	KustoOfferingTypeFree KustoOfferingType = "free"
)

// PossibleKustoOfferingTypeValues returns the possible values for the KustoOfferingType const type.
func PossibleKustoOfferingTypeValues() []KustoOfferingType {
	return []KustoOfferingType{
		KustoOfferingTypeAdx,
		KustoOfferingTypeFabric,
		KustoOfferingTypeFree,
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

// SharedPrivateLinkResourceStatus - Status of the shared private link resource. Can be Pending, Approved, Rejected or Disconnected.
type SharedPrivateLinkResourceStatus string

const (
	// SharedPrivateLinkResourceStatusApproved - The shared private link connection request was approved by the resource owner.
	SharedPrivateLinkResourceStatusApproved SharedPrivateLinkResourceStatus = "Approved"
	// SharedPrivateLinkResourceStatusDisconnected - The shared private link connection request was disconnected by the resource
	// owner.
	SharedPrivateLinkResourceStatusDisconnected SharedPrivateLinkResourceStatus = "Disconnected"
	// SharedPrivateLinkResourceStatusPending - The shared private link connection request was not yet authorized by the resource
	// owner.
	SharedPrivateLinkResourceStatusPending SharedPrivateLinkResourceStatus = "Pending"
	// SharedPrivateLinkResourceStatusRejected - The shared private link connection request was rejected by the resource owner.
	SharedPrivateLinkResourceStatusRejected SharedPrivateLinkResourceStatus = "Rejected"
)

// PossibleSharedPrivateLinkResourceStatusValues returns the possible values for the SharedPrivateLinkResourceStatus const type.
func PossibleSharedPrivateLinkResourceStatusValues() []SharedPrivateLinkResourceStatus {
	return []SharedPrivateLinkResourceStatus{
		SharedPrivateLinkResourceStatusApproved,
		SharedPrivateLinkResourceStatusDisconnected,
		SharedPrivateLinkResourceStatusPending,
		SharedPrivateLinkResourceStatusRejected,
	}
}

// String - The target collection status.
type String string

const (
	// StringDisabled - Denotes a target that is disabled.
	StringDisabled String = "Disabled"
	// StringEnabled - Denotes a target that is enabled.
	StringEnabled String = "Enabled"
)

// PossibleStringValues returns the possible values for the String const type.
func PossibleStringValues() []String {
	return []String{
		StringDisabled,
		StringEnabled,
	}
}

// TargetAuthenticationType - The type of authentication to use when connecting to a target.
type TargetAuthenticationType string

const (
	// TargetAuthenticationTypeAAD - The Azure Active Directory authentication.
	TargetAuthenticationTypeAAD TargetAuthenticationType = "Aad"
	// TargetAuthenticationTypeSQL - The SQL password authentication.
	TargetAuthenticationTypeSQL TargetAuthenticationType = "Sql"
)

// PossibleTargetAuthenticationTypeValues returns the possible values for the TargetAuthenticationType const type.
func PossibleTargetAuthenticationTypeValues() []TargetAuthenticationType {
	return []TargetAuthenticationType{
		TargetAuthenticationTypeAAD,
		TargetAuthenticationTypeSQL,
	}
}

// WatcherStatus - The monitoring collection status of a watcher.
type WatcherStatus string

const (
	// WatcherStatusDeleting - Denotes the watcher is in a deleting state.
	WatcherStatusDeleting WatcherStatus = "Deleting"
	// WatcherStatusRunning - Denotes the watcher is in a running state.
	WatcherStatusRunning WatcherStatus = "Running"
	// WatcherStatusStarting - Denotes the watcher is in a starting state.
	WatcherStatusStarting WatcherStatus = "Starting"
	// WatcherStatusStopped - Denotes the watcher is in a stopped state.
	WatcherStatusStopped WatcherStatus = "Stopped"
	// WatcherStatusStopping - Denotes the watcher is in a stopping state.
	WatcherStatusStopping WatcherStatus = "Stopping"
)

// PossibleWatcherStatusValues returns the possible values for the WatcherStatus const type.
func PossibleWatcherStatusValues() []WatcherStatus {
	return []WatcherStatus{
		WatcherStatusDeleting,
		WatcherStatusRunning,
		WatcherStatusStarting,
		WatcherStatusStopped,
		WatcherStatusStopping,
	}
}
