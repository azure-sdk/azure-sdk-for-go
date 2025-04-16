// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakeanalytics

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
	moduleVersion = "v0.8.1"
)

// AADObjectType - The type of AAD object the object identifier refers to.
type AADObjectType string

const (
	AADObjectTypeGroup            AADObjectType = "Group"
	AADObjectTypeServicePrincipal AADObjectType = "ServicePrincipal"
	AADObjectTypeUser             AADObjectType = "User"
)

// PossibleAADObjectTypeValues returns the possible values for the AADObjectType const type.
func PossibleAADObjectTypeValues() []AADObjectType {
	return []AADObjectType{
		AADObjectTypeGroup,
		AADObjectTypeServicePrincipal,
		AADObjectTypeUser,
	}
}

// CheckNameAvailabilityParametersType - The resource type. Note: This should not be set by the user, as the constant value
// is Microsoft.DataLakeAnalytics/accounts
type CheckNameAvailabilityParametersType string

const (
	CheckNameAvailabilityParametersTypeMicrosoftDataLakeAnalyticsAccounts CheckNameAvailabilityParametersType = "Microsoft.DataLakeAnalytics/accounts"
)

// PossibleCheckNameAvailabilityParametersTypeValues returns the possible values for the CheckNameAvailabilityParametersType const type.
func PossibleCheckNameAvailabilityParametersTypeValues() []CheckNameAvailabilityParametersType {
	return []CheckNameAvailabilityParametersType{
		CheckNameAvailabilityParametersTypeMicrosoftDataLakeAnalyticsAccounts,
	}
}

// DataLakeAnalyticsAccountState - The state of the Data Lake Analytics account.
type DataLakeAnalyticsAccountState string

const (
	DataLakeAnalyticsAccountStateActive    DataLakeAnalyticsAccountState = "Active"
	DataLakeAnalyticsAccountStateSuspended DataLakeAnalyticsAccountState = "Suspended"
)

// PossibleDataLakeAnalyticsAccountStateValues returns the possible values for the DataLakeAnalyticsAccountState const type.
func PossibleDataLakeAnalyticsAccountStateValues() []DataLakeAnalyticsAccountState {
	return []DataLakeAnalyticsAccountState{
		DataLakeAnalyticsAccountStateActive,
		DataLakeAnalyticsAccountStateSuspended,
	}
}

// DataLakeAnalyticsAccountStatus - The provisioning status of the Data Lake Analytics account.
type DataLakeAnalyticsAccountStatus string

const (
	DataLakeAnalyticsAccountStatusCanceled   DataLakeAnalyticsAccountStatus = "Canceled"
	DataLakeAnalyticsAccountStatusCreating   DataLakeAnalyticsAccountStatus = "Creating"
	DataLakeAnalyticsAccountStatusDeleted    DataLakeAnalyticsAccountStatus = "Deleted"
	DataLakeAnalyticsAccountStatusDeleting   DataLakeAnalyticsAccountStatus = "Deleting"
	DataLakeAnalyticsAccountStatusFailed     DataLakeAnalyticsAccountStatus = "Failed"
	DataLakeAnalyticsAccountStatusPatching   DataLakeAnalyticsAccountStatus = "Patching"
	DataLakeAnalyticsAccountStatusResuming   DataLakeAnalyticsAccountStatus = "Resuming"
	DataLakeAnalyticsAccountStatusRunning    DataLakeAnalyticsAccountStatus = "Running"
	DataLakeAnalyticsAccountStatusSucceeded  DataLakeAnalyticsAccountStatus = "Succeeded"
	DataLakeAnalyticsAccountStatusSuspending DataLakeAnalyticsAccountStatus = "Suspending"
	DataLakeAnalyticsAccountStatusUndeleting DataLakeAnalyticsAccountStatus = "Undeleting"
)

// PossibleDataLakeAnalyticsAccountStatusValues returns the possible values for the DataLakeAnalyticsAccountStatus const type.
func PossibleDataLakeAnalyticsAccountStatusValues() []DataLakeAnalyticsAccountStatus {
	return []DataLakeAnalyticsAccountStatus{
		DataLakeAnalyticsAccountStatusCanceled,
		DataLakeAnalyticsAccountStatusCreating,
		DataLakeAnalyticsAccountStatusDeleted,
		DataLakeAnalyticsAccountStatusDeleting,
		DataLakeAnalyticsAccountStatusFailed,
		DataLakeAnalyticsAccountStatusPatching,
		DataLakeAnalyticsAccountStatusResuming,
		DataLakeAnalyticsAccountStatusRunning,
		DataLakeAnalyticsAccountStatusSucceeded,
		DataLakeAnalyticsAccountStatusSuspending,
		DataLakeAnalyticsAccountStatusUndeleting,
	}
}

// DebugDataAccessLevel - The current state of the DebugDataAccessLevel for this account.
type DebugDataAccessLevel string

const (
	DebugDataAccessLevelAll      DebugDataAccessLevel = "All"
	DebugDataAccessLevelCustomer DebugDataAccessLevel = "Customer"
	DebugDataAccessLevelNone     DebugDataAccessLevel = "None"
)

// PossibleDebugDataAccessLevelValues returns the possible values for the DebugDataAccessLevel const type.
func PossibleDebugDataAccessLevelValues() []DebugDataAccessLevel {
	return []DebugDataAccessLevel{
		DebugDataAccessLevelAll,
		DebugDataAccessLevelCustomer,
		DebugDataAccessLevelNone,
	}
}

// FirewallAllowAzureIPsState - The current state of allowing or disallowing IPs originating within Azure through the firewall.
// If the firewall is disabled, this is not enforced.
type FirewallAllowAzureIPsState string

const (
	FirewallAllowAzureIPsStateDisabled FirewallAllowAzureIPsState = "Disabled"
	FirewallAllowAzureIPsStateEnabled  FirewallAllowAzureIPsState = "Enabled"
)

// PossibleFirewallAllowAzureIPsStateValues returns the possible values for the FirewallAllowAzureIPsState const type.
func PossibleFirewallAllowAzureIPsStateValues() []FirewallAllowAzureIPsState {
	return []FirewallAllowAzureIPsState{
		FirewallAllowAzureIPsStateDisabled,
		FirewallAllowAzureIPsStateEnabled,
	}
}

// FirewallState - The current state of the IP address firewall for this account.
type FirewallState string

const (
	FirewallStateDisabled FirewallState = "Disabled"
	FirewallStateEnabled  FirewallState = "Enabled"
)

// PossibleFirewallStateValues returns the possible values for the FirewallState const type.
func PossibleFirewallStateValues() []FirewallState {
	return []FirewallState{
		FirewallStateDisabled,
		FirewallStateEnabled,
	}
}

// NestedResourceProvisioningState - The current state of the NestedResourceProvisioning for this account.
type NestedResourceProvisioningState string

const (
	NestedResourceProvisioningStateCanceled  NestedResourceProvisioningState = "Canceled"
	NestedResourceProvisioningStateFailed    NestedResourceProvisioningState = "Failed"
	NestedResourceProvisioningStateSucceeded NestedResourceProvisioningState = "Succeeded"
)

// PossibleNestedResourceProvisioningStateValues returns the possible values for the NestedResourceProvisioningState const type.
func PossibleNestedResourceProvisioningStateValues() []NestedResourceProvisioningState {
	return []NestedResourceProvisioningState{
		NestedResourceProvisioningStateCanceled,
		NestedResourceProvisioningStateFailed,
		NestedResourceProvisioningStateSucceeded,
	}
}

// OperationOrigin - The intended executor of the operation.
type OperationOrigin string

const (
	OperationOriginSystem     OperationOrigin = "system"
	OperationOriginUser       OperationOrigin = "user"
	OperationOriginUserSystem OperationOrigin = "user,system"
)

// PossibleOperationOriginValues returns the possible values for the OperationOrigin const type.
func PossibleOperationOriginValues() []OperationOrigin {
	return []OperationOrigin{
		OperationOriginSystem,
		OperationOriginUser,
		OperationOriginUserSystem,
	}
}

// SubscriptionState - The subscription state.
type SubscriptionState string

const (
	SubscriptionStateDeleted      SubscriptionState = "Deleted"
	SubscriptionStateRegistered   SubscriptionState = "Registered"
	SubscriptionStateSuspended    SubscriptionState = "Suspended"
	SubscriptionStateUnregistered SubscriptionState = "Unregistered"
	SubscriptionStateWarned       SubscriptionState = "Warned"
)

// PossibleSubscriptionStateValues returns the possible values for the SubscriptionState const type.
func PossibleSubscriptionStateValues() []SubscriptionState {
	return []SubscriptionState{
		SubscriptionStateDeleted,
		SubscriptionStateRegistered,
		SubscriptionStateSuspended,
		SubscriptionStateUnregistered,
		SubscriptionStateWarned,
	}
}

// TierType - The commitment tier for the next month.
type TierType string

const (
	TierTypeCommitment100000AUHours TierType = "Commitment_100000AUHours"
	TierTypeCommitment10000AUHours  TierType = "Commitment_10000AUHours"
	TierTypeCommitment1000AUHours   TierType = "Commitment_1000AUHours"
	TierTypeCommitment100AUHours    TierType = "Commitment_100AUHours"
	TierTypeCommitment500000AUHours TierType = "Commitment_500000AUHours"
	TierTypeCommitment50000AUHours  TierType = "Commitment_50000AUHours"
	TierTypeCommitment5000AUHours   TierType = "Commitment_5000AUHours"
	TierTypeCommitment500AUHours    TierType = "Commitment_500AUHours"
	TierTypeConsumption             TierType = "Consumption"
)

// PossibleTierTypeValues returns the possible values for the TierType const type.
func PossibleTierTypeValues() []TierType {
	return []TierType{
		TierTypeCommitment100000AUHours,
		TierTypeCommitment10000AUHours,
		TierTypeCommitment1000AUHours,
		TierTypeCommitment100AUHours,
		TierTypeCommitment500000AUHours,
		TierTypeCommitment50000AUHours,
		TierTypeCommitment5000AUHours,
		TierTypeCommitment500AUHours,
		TierTypeConsumption,
	}
}

// VirtualNetworkRuleState - The current state of the VirtualNetworkRule for this account.
type VirtualNetworkRuleState string

const (
	VirtualNetworkRuleStateActive               VirtualNetworkRuleState = "Active"
	VirtualNetworkRuleStateFailed               VirtualNetworkRuleState = "Failed"
	VirtualNetworkRuleStateNetworkSourceDeleted VirtualNetworkRuleState = "NetworkSourceDeleted"
)

// PossibleVirtualNetworkRuleStateValues returns the possible values for the VirtualNetworkRuleState const type.
func PossibleVirtualNetworkRuleStateValues() []VirtualNetworkRuleState {
	return []VirtualNetworkRuleState{
		VirtualNetworkRuleStateActive,
		VirtualNetworkRuleStateFailed,
		VirtualNetworkRuleStateNetworkSourceDeleted,
	}
}
