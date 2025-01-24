//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakestore

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
	moduleVersion = "v1.2.1"
)

// CheckNameAvailabilityParametersType - The resource type. Note: This should not be set by the user, as the constant value
// is Microsoft.DataLakeStore/accounts
type CheckNameAvailabilityParametersType string

const (
	CheckNameAvailabilityParametersTypeMicrosoftDataLakeStoreAccounts CheckNameAvailabilityParametersType = "Microsoft.DataLakeStore/accounts"
)

// PossibleCheckNameAvailabilityParametersTypeValues returns the possible values for the CheckNameAvailabilityParametersType const type.
func PossibleCheckNameAvailabilityParametersTypeValues() []CheckNameAvailabilityParametersType {
	return []CheckNameAvailabilityParametersType{
		CheckNameAvailabilityParametersTypeMicrosoftDataLakeStoreAccounts,
	}
}

// DataLakeStoreAccountState - The state of the Data Lake Store account.
type DataLakeStoreAccountState string

const (
	DataLakeStoreAccountStateActive    DataLakeStoreAccountState = "Active"
	DataLakeStoreAccountStateSuspended DataLakeStoreAccountState = "Suspended"
)

// PossibleDataLakeStoreAccountStateValues returns the possible values for the DataLakeStoreAccountState const type.
func PossibleDataLakeStoreAccountStateValues() []DataLakeStoreAccountState {
	return []DataLakeStoreAccountState{
		DataLakeStoreAccountStateActive,
		DataLakeStoreAccountStateSuspended,
	}
}

// DataLakeStoreAccountStatus - The provisioning status of the Data Lake Store account.
type DataLakeStoreAccountStatus string

const (
	DataLakeStoreAccountStatusCanceled   DataLakeStoreAccountStatus = "Canceled"
	DataLakeStoreAccountStatusCreating   DataLakeStoreAccountStatus = "Creating"
	DataLakeStoreAccountStatusDeleted    DataLakeStoreAccountStatus = "Deleted"
	DataLakeStoreAccountStatusDeleting   DataLakeStoreAccountStatus = "Deleting"
	DataLakeStoreAccountStatusFailed     DataLakeStoreAccountStatus = "Failed"
	DataLakeStoreAccountStatusPatching   DataLakeStoreAccountStatus = "Patching"
	DataLakeStoreAccountStatusResuming   DataLakeStoreAccountStatus = "Resuming"
	DataLakeStoreAccountStatusRunning    DataLakeStoreAccountStatus = "Running"
	DataLakeStoreAccountStatusSucceeded  DataLakeStoreAccountStatus = "Succeeded"
	DataLakeStoreAccountStatusSuspending DataLakeStoreAccountStatus = "Suspending"
	DataLakeStoreAccountStatusUndeleting DataLakeStoreAccountStatus = "Undeleting"
)

// PossibleDataLakeStoreAccountStatusValues returns the possible values for the DataLakeStoreAccountStatus const type.
func PossibleDataLakeStoreAccountStatusValues() []DataLakeStoreAccountStatus {
	return []DataLakeStoreAccountStatus{
		DataLakeStoreAccountStatusCanceled,
		DataLakeStoreAccountStatusCreating,
		DataLakeStoreAccountStatusDeleted,
		DataLakeStoreAccountStatusDeleting,
		DataLakeStoreAccountStatusFailed,
		DataLakeStoreAccountStatusPatching,
		DataLakeStoreAccountStatusResuming,
		DataLakeStoreAccountStatusRunning,
		DataLakeStoreAccountStatusSucceeded,
		DataLakeStoreAccountStatusSuspending,
		DataLakeStoreAccountStatusUndeleting,
	}
}

// EncryptionConfigType - The type of encryption configuration being used. Currently the only supported types are 'UserManaged'
// and 'ServiceManaged'.
type EncryptionConfigType string

const (
	EncryptionConfigTypeServiceManaged EncryptionConfigType = "ServiceManaged"
	EncryptionConfigTypeUserManaged    EncryptionConfigType = "UserManaged"
)

// PossibleEncryptionConfigTypeValues returns the possible values for the EncryptionConfigType const type.
func PossibleEncryptionConfigTypeValues() []EncryptionConfigType {
	return []EncryptionConfigType{
		EncryptionConfigTypeServiceManaged,
		EncryptionConfigTypeUserManaged,
	}
}

// EncryptionProvisioningState - The current state of encryption provisioning for this Data Lake Store account.
type EncryptionProvisioningState string

const (
	EncryptionProvisioningStateCreating  EncryptionProvisioningState = "Creating"
	EncryptionProvisioningStateSucceeded EncryptionProvisioningState = "Succeeded"
)

// PossibleEncryptionProvisioningStateValues returns the possible values for the EncryptionProvisioningState const type.
func PossibleEncryptionProvisioningStateValues() []EncryptionProvisioningState {
	return []EncryptionProvisioningState{
		EncryptionProvisioningStateCreating,
		EncryptionProvisioningStateSucceeded,
	}
}

// EncryptionState - The current state of encryption for this Data Lake Store account.
type EncryptionState string

const (
	EncryptionStateDisabled EncryptionState = "Disabled"
	EncryptionStateEnabled  EncryptionState = "Enabled"
)

// PossibleEncryptionStateValues returns the possible values for the EncryptionState const type.
func PossibleEncryptionStateValues() []EncryptionState {
	return []EncryptionState{
		EncryptionStateDisabled,
		EncryptionStateEnabled,
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

// FirewallState - The current state of the IP address firewall for this Data Lake Store account.
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

// TierType - The commitment tier to use for next month.
type TierType string

const (
	TierTypeCommitment100TB TierType = "Commitment_100TB"
	TierTypeCommitment10TB  TierType = "Commitment_10TB"
	TierTypeCommitment1PB   TierType = "Commitment_1PB"
	TierTypeCommitment1TB   TierType = "Commitment_1TB"
	TierTypeCommitment500TB TierType = "Commitment_500TB"
	TierTypeCommitment5PB   TierType = "Commitment_5PB"
	TierTypeConsumption     TierType = "Consumption"
)

// PossibleTierTypeValues returns the possible values for the TierType const type.
func PossibleTierTypeValues() []TierType {
	return []TierType{
		TierTypeCommitment100TB,
		TierTypeCommitment10TB,
		TierTypeCommitment1PB,
		TierTypeCommitment1TB,
		TierTypeCommitment500TB,
		TierTypeCommitment5PB,
		TierTypeConsumption,
	}
}

// TrustedIDProviderState - The current state of the trusted identity provider feature for this Data Lake Store account.
type TrustedIDProviderState string

const (
	TrustedIDProviderStateDisabled TrustedIDProviderState = "Disabled"
	TrustedIDProviderStateEnabled  TrustedIDProviderState = "Enabled"
)

// PossibleTrustedIDProviderStateValues returns the possible values for the TrustedIDProviderState const type.
func PossibleTrustedIDProviderStateValues() []TrustedIDProviderState {
	return []TrustedIDProviderState{
		TrustedIDProviderStateDisabled,
		TrustedIDProviderStateEnabled,
	}
}

// UsageUnit - Gets the unit of measurement.
type UsageUnit string

const (
	UsageUnitBytes           UsageUnit = "Bytes"
	UsageUnitBytesPerSecond  UsageUnit = "BytesPerSecond"
	UsageUnitCount           UsageUnit = "Count"
	UsageUnitCountsPerSecond UsageUnit = "CountsPerSecond"
	UsageUnitPercent         UsageUnit = "Percent"
	UsageUnitSeconds         UsageUnit = "Seconds"
)

// PossibleUsageUnitValues returns the possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{
		UsageUnitBytes,
		UsageUnitBytesPerSecond,
		UsageUnitCount,
		UsageUnitCountsPerSecond,
		UsageUnitPercent,
		UsageUnitSeconds,
	}
}
