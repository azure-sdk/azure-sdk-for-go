//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredisenterprise

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
	moduleVersion = "v2.0.0"
)

// AccessKeyType - Which access key to regenerate.
type AccessKeyType string

const (
	AccessKeyTypePrimary   AccessKeyType = "Primary"
	AccessKeyTypeSecondary AccessKeyType = "Secondary"
)

// PossibleAccessKeyTypeValues returns the possible values for the AccessKeyType const type.
func PossibleAccessKeyTypeValues() []AccessKeyType {
	return []AccessKeyType{
		AccessKeyTypePrimary,
		AccessKeyTypeSecondary,
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

// AofFrequency - Sets the frequency at which data is written to disk.
type AofFrequency string

const (
	AofFrequencyAlways AofFrequency = "always"
	AofFrequencyOneS   AofFrequency = "1s"
)

// PossibleAofFrequencyValues returns the possible values for the AofFrequency const type.
func PossibleAofFrequencyValues() []AofFrequency {
	return []AofFrequency{
		AofFrequencyAlways,
		AofFrequencyOneS,
	}
}

// ClusteringPolicy - Clustering policy - default is OSSCluster. Specified at create time.
type ClusteringPolicy string

const (
	ClusteringPolicyEnterpriseCluster ClusteringPolicy = "EnterpriseCluster"
	ClusteringPolicyOSSCluster        ClusteringPolicy = "OSSCluster"
)

// PossibleClusteringPolicyValues returns the possible values for the ClusteringPolicy const type.
func PossibleClusteringPolicyValues() []ClusteringPolicy {
	return []ClusteringPolicy{
		ClusteringPolicyEnterpriseCluster,
		ClusteringPolicyOSSCluster,
	}
}

// CmkIdentityType - Only userAssignedIdentity is supported in this API version; other types may be supported in the future
type CmkIdentityType string

const (
	CmkIdentityTypeSystemAssignedIdentity CmkIdentityType = "systemAssignedIdentity"
	CmkIdentityTypeUserAssignedIdentity   CmkIdentityType = "userAssignedIdentity"
)

// PossibleCmkIdentityTypeValues returns the possible values for the CmkIdentityType const type.
func PossibleCmkIdentityTypeValues() []CmkIdentityType {
	return []CmkIdentityType{
		CmkIdentityTypeSystemAssignedIdentity,
		CmkIdentityTypeUserAssignedIdentity,
	}
}

// EvictionPolicy - Redis eviction policy - default is VolatileLRU
type EvictionPolicy string

const (
	EvictionPolicyAllKeysLFU     EvictionPolicy = "AllKeysLFU"
	EvictionPolicyAllKeysLRU     EvictionPolicy = "AllKeysLRU"
	EvictionPolicyAllKeysRandom  EvictionPolicy = "AllKeysRandom"
	EvictionPolicyNoEviction     EvictionPolicy = "NoEviction"
	EvictionPolicyVolatileLFU    EvictionPolicy = "VolatileLFU"
	EvictionPolicyVolatileLRU    EvictionPolicy = "VolatileLRU"
	EvictionPolicyVolatileRandom EvictionPolicy = "VolatileRandom"
	EvictionPolicyVolatileTTL    EvictionPolicy = "VolatileTTL"
)

// PossibleEvictionPolicyValues returns the possible values for the EvictionPolicy const type.
func PossibleEvictionPolicyValues() []EvictionPolicy {
	return []EvictionPolicy{
		EvictionPolicyAllKeysLFU,
		EvictionPolicyAllKeysLRU,
		EvictionPolicyAllKeysRandom,
		EvictionPolicyNoEviction,
		EvictionPolicyVolatileLFU,
		EvictionPolicyVolatileLRU,
		EvictionPolicyVolatileRandom,
		EvictionPolicyVolatileTTL,
	}
}

// LinkState - State of the link between the database resources.
type LinkState string

const (
	LinkStateLinkFailed   LinkState = "LinkFailed"
	LinkStateLinked       LinkState = "Linked"
	LinkStateLinking      LinkState = "Linking"
	LinkStateUnlinkFailed LinkState = "UnlinkFailed"
	LinkStateUnlinking    LinkState = "Unlinking"
)

// PossibleLinkStateValues returns the possible values for the LinkState const type.
func PossibleLinkStateValues() []LinkState {
	return []LinkState{
		LinkStateLinkFailed,
		LinkStateLinked,
		LinkStateLinking,
		LinkStateUnlinkFailed,
		LinkStateUnlinking,
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

// Protocol - Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted.
type Protocol string

const (
	ProtocolEncrypted Protocol = "Encrypted"
	ProtocolPlaintext Protocol = "Plaintext"
)

// PossibleProtocolValues returns the possible values for the Protocol const type.
func PossibleProtocolValues() []Protocol {
	return []Protocol{
		ProtocolEncrypted,
		ProtocolPlaintext,
	}
}

// ProvisioningState - Current provisioning status
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// RdbFrequency - Sets the frequency at which a snapshot of the database is created.
type RdbFrequency string

const (
	RdbFrequencyOneH    RdbFrequency = "1h"
	RdbFrequencySixH    RdbFrequency = "6h"
	RdbFrequencyTwelveH RdbFrequency = "12h"
)

// PossibleRdbFrequencyValues returns the possible values for the RdbFrequency const type.
func PossibleRdbFrequencyValues() []RdbFrequency {
	return []RdbFrequency{
		RdbFrequencyOneH,
		RdbFrequencySixH,
		RdbFrequencyTwelveH,
	}
}

// ResourceState - Current resource status
type ResourceState string

const (
	ResourceStateCreateFailed  ResourceState = "CreateFailed"
	ResourceStateCreating      ResourceState = "Creating"
	ResourceStateDeleteFailed  ResourceState = "DeleteFailed"
	ResourceStateDeleting      ResourceState = "Deleting"
	ResourceStateDisableFailed ResourceState = "DisableFailed"
	ResourceStateDisabled      ResourceState = "Disabled"
	ResourceStateDisabling     ResourceState = "Disabling"
	ResourceStateEnableFailed  ResourceState = "EnableFailed"
	ResourceStateEnabling      ResourceState = "Enabling"
	ResourceStateRunning       ResourceState = "Running"
	ResourceStateScaling       ResourceState = "Scaling"
	ResourceStateScalingFailed ResourceState = "ScalingFailed"
	ResourceStateUpdateFailed  ResourceState = "UpdateFailed"
	ResourceStateUpdating      ResourceState = "Updating"
)

// PossibleResourceStateValues returns the possible values for the ResourceState const type.
func PossibleResourceStateValues() []ResourceState {
	return []ResourceState{
		ResourceStateCreateFailed,
		ResourceStateCreating,
		ResourceStateDeleteFailed,
		ResourceStateDeleting,
		ResourceStateDisableFailed,
		ResourceStateDisabled,
		ResourceStateDisabling,
		ResourceStateEnableFailed,
		ResourceStateEnabling,
		ResourceStateRunning,
		ResourceStateScaling,
		ResourceStateScalingFailed,
		ResourceStateUpdateFailed,
		ResourceStateUpdating,
	}
}

// SKUName - The type of RedisEnterprise cluster to deploy. Possible values: (EnterpriseE10, EnterpriseFlashF300 etc.)
type SKUName string

const (
	SKUNameEnterpriseE10        SKUName = "Enterprise_E10"
	SKUNameEnterpriseE100       SKUName = "Enterprise_E100"
	SKUNameEnterpriseE20        SKUName = "Enterprise_E20"
	SKUNameEnterpriseE5         SKUName = "Enterprise_E5"
	SKUNameEnterpriseE50        SKUName = "Enterprise_E50"
	SKUNameEnterpriseFlashF1500 SKUName = "EnterpriseFlash_F1500"
	SKUNameEnterpriseFlashF300  SKUName = "EnterpriseFlash_F300"
	SKUNameEnterpriseFlashF700  SKUName = "EnterpriseFlash_F700"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameEnterpriseE10,
		SKUNameEnterpriseE100,
		SKUNameEnterpriseE20,
		SKUNameEnterpriseE5,
		SKUNameEnterpriseE50,
		SKUNameEnterpriseFlashF1500,
		SKUNameEnterpriseFlashF300,
		SKUNameEnterpriseFlashF700,
	}
}

// TLSVersion - The minimum TLS version for the cluster to support, e.g. '1.2'
type TLSVersion string

const (
	TLSVersionOne0 TLSVersion = "1.0"
	TLSVersionOne1 TLSVersion = "1.1"
	TLSVersionOne2 TLSVersion = "1.2"
)

// PossibleTLSVersionValues returns the possible values for the TLSVersion const type.
func PossibleTLSVersionValues() []TLSVersion {
	return []TLSVersion{
		TLSVersionOne0,
		TLSVersionOne1,
		TLSVersionOne2,
	}
}
