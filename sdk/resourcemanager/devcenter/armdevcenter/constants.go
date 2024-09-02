//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevcenter

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
	moduleVersion = "v2.1.0-beta.1"
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

// CatalogConnectionState - The connection state of the catalog.
type CatalogConnectionState string

const (
	CatalogConnectionStateConnected    CatalogConnectionState = "Connected"
	CatalogConnectionStateDisconnected CatalogConnectionState = "Disconnected"
)

// PossibleCatalogConnectionStateValues returns the possible values for the CatalogConnectionState const type.
func PossibleCatalogConnectionStateValues() []CatalogConnectionState {
	return []CatalogConnectionState{
		CatalogConnectionStateConnected,
		CatalogConnectionStateDisconnected,
	}
}

// CatalogItemSyncEnableStatus - Catalog item sync types enable or disable status. Indicates whether project catalogs are
// allowed to sync catalog items under projects associated to this dev center.
type CatalogItemSyncEnableStatus string

const (
	CatalogItemSyncEnableStatusDisabled CatalogItemSyncEnableStatus = "Disabled"
	CatalogItemSyncEnableStatusEnabled  CatalogItemSyncEnableStatus = "Enabled"
)

// PossibleCatalogItemSyncEnableStatusValues returns the possible values for the CatalogItemSyncEnableStatus const type.
func PossibleCatalogItemSyncEnableStatusValues() []CatalogItemSyncEnableStatus {
	return []CatalogItemSyncEnableStatus{
		CatalogItemSyncEnableStatusDisabled,
		CatalogItemSyncEnableStatusEnabled,
	}
}

// CatalogItemType - Indicates catalog item types.
type CatalogItemType string

const (
	CatalogItemTypeEnvironmentDefinition CatalogItemType = "EnvironmentDefinition"
	CatalogItemTypeImageDefinition       CatalogItemType = "ImageDefinition"
)

// PossibleCatalogItemTypeValues returns the possible values for the CatalogItemType const type.
func PossibleCatalogItemTypeValues() []CatalogItemType {
	return []CatalogItemType{
		CatalogItemTypeEnvironmentDefinition,
		CatalogItemTypeImageDefinition,
	}
}

// CatalogResourceValidationStatus - Catalog resource validation status
type CatalogResourceValidationStatus string

const (
	CatalogResourceValidationStatusFailed    CatalogResourceValidationStatus = "Failed"
	CatalogResourceValidationStatusPending   CatalogResourceValidationStatus = "Pending"
	CatalogResourceValidationStatusSucceeded CatalogResourceValidationStatus = "Succeeded"
	CatalogResourceValidationStatusUnknown   CatalogResourceValidationStatus = "Unknown"
)

// PossibleCatalogResourceValidationStatusValues returns the possible values for the CatalogResourceValidationStatus const type.
func PossibleCatalogResourceValidationStatusValues() []CatalogResourceValidationStatus {
	return []CatalogResourceValidationStatus{
		CatalogResourceValidationStatusFailed,
		CatalogResourceValidationStatusPending,
		CatalogResourceValidationStatusSucceeded,
		CatalogResourceValidationStatusUnknown,
	}
}

// CatalogSyncState - The synchronization state of the catalog.
type CatalogSyncState string

const (
	CatalogSyncStateCanceled   CatalogSyncState = "Canceled"
	CatalogSyncStateFailed     CatalogSyncState = "Failed"
	CatalogSyncStateInProgress CatalogSyncState = "InProgress"
	CatalogSyncStateSucceeded  CatalogSyncState = "Succeeded"
)

// PossibleCatalogSyncStateValues returns the possible values for the CatalogSyncState const type.
func PossibleCatalogSyncStateValues() []CatalogSyncState {
	return []CatalogSyncState{
		CatalogSyncStateCanceled,
		CatalogSyncStateFailed,
		CatalogSyncStateInProgress,
		CatalogSyncStateSucceeded,
	}
}

// CatalogSyncType - Indicates the type of sync that is configured for the catalog.
type CatalogSyncType string

const (
	CatalogSyncTypeManual    CatalogSyncType = "Manual"
	CatalogSyncTypeScheduled CatalogSyncType = "Scheduled"
)

// PossibleCatalogSyncTypeValues returns the possible values for the CatalogSyncType const type.
func PossibleCatalogSyncTypeValues() []CatalogSyncType {
	return []CatalogSyncType{
		CatalogSyncTypeManual,
		CatalogSyncTypeScheduled,
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

// CustomizationTaskInputType - Type of the input.
type CustomizationTaskInputType string

const (
	CustomizationTaskInputTypeBoolean CustomizationTaskInputType = "boolean"
	CustomizationTaskInputTypeNumber  CustomizationTaskInputType = "number"
	CustomizationTaskInputTypeString  CustomizationTaskInputType = "string"
)

// PossibleCustomizationTaskInputTypeValues returns the possible values for the CustomizationTaskInputType const type.
func PossibleCustomizationTaskInputTypeValues() []CustomizationTaskInputType {
	return []CustomizationTaskInputType{
		CustomizationTaskInputTypeBoolean,
		CustomizationTaskInputTypeNumber,
		CustomizationTaskInputTypeString,
	}
}

// DevCenterResourceType - Indicates dev center resource types.
type DevCenterResourceType string

const (
	DevCenterResourceTypeAttachedNetworks DevCenterResourceType = "AttachedNetworks"
	DevCenterResourceTypeImages           DevCenterResourceType = "Images"
	DevCenterResourceTypeSKUs             DevCenterResourceType = "Skus"
)

// PossibleDevCenterResourceTypeValues returns the possible values for the DevCenterResourceType const type.
func PossibleDevCenterResourceTypeValues() []DevCenterResourceType {
	return []DevCenterResourceType{
		DevCenterResourceTypeAttachedNetworks,
		DevCenterResourceTypeImages,
		DevCenterResourceTypeSKUs,
	}
}

// DevboxDisksEncryptionEnableStatus - Devbox disk encryption enable or disable status. Indicates if Devbox disks encryption
// is enabled or not.
type DevboxDisksEncryptionEnableStatus string

const (
	DevboxDisksEncryptionEnableStatusDisabled DevboxDisksEncryptionEnableStatus = "Disabled"
	DevboxDisksEncryptionEnableStatusEnabled  DevboxDisksEncryptionEnableStatus = "Enabled"
)

// PossibleDevboxDisksEncryptionEnableStatusValues returns the possible values for the DevboxDisksEncryptionEnableStatus const type.
func PossibleDevboxDisksEncryptionEnableStatusValues() []DevboxDisksEncryptionEnableStatus {
	return []DevboxDisksEncryptionEnableStatus{
		DevboxDisksEncryptionEnableStatusDisabled,
		DevboxDisksEncryptionEnableStatusEnabled,
	}
}

// DomainJoinType - Active Directory join type
type DomainJoinType string

const (
	DomainJoinTypeAzureADJoin       DomainJoinType = "AzureADJoin"
	DomainJoinTypeHybridAzureADJoin DomainJoinType = "HybridAzureADJoin"
	DomainJoinTypeNone              DomainJoinType = "None"
)

// PossibleDomainJoinTypeValues returns the possible values for the DomainJoinType const type.
func PossibleDomainJoinTypeValues() []DomainJoinType {
	return []DomainJoinType{
		DomainJoinTypeAzureADJoin,
		DomainJoinTypeHybridAzureADJoin,
		DomainJoinTypeNone,
	}
}

// EnvironmentTypeEnableStatus - Indicates whether the environment type is either enabled or disabled.
type EnvironmentTypeEnableStatus string

const (
	EnvironmentTypeEnableStatusDisabled EnvironmentTypeEnableStatus = "Disabled"
	EnvironmentTypeEnableStatusEnabled  EnvironmentTypeEnableStatus = "Enabled"
)

// PossibleEnvironmentTypeEnableStatusValues returns the possible values for the EnvironmentTypeEnableStatus const type.
func PossibleEnvironmentTypeEnableStatusValues() []EnvironmentTypeEnableStatus {
	return []EnvironmentTypeEnableStatus{
		EnvironmentTypeEnableStatusDisabled,
		EnvironmentTypeEnableStatusEnabled,
	}
}

// HealthCheckStatus - Health check status values
type HealthCheckStatus string

const (
	HealthCheckStatusFailed  HealthCheckStatus = "Failed"
	HealthCheckStatusPassed  HealthCheckStatus = "Passed"
	HealthCheckStatusPending HealthCheckStatus = "Pending"
	HealthCheckStatusRunning HealthCheckStatus = "Running"
	HealthCheckStatusUnknown HealthCheckStatus = "Unknown"
	HealthCheckStatusWarning HealthCheckStatus = "Warning"
)

// PossibleHealthCheckStatusValues returns the possible values for the HealthCheckStatus const type.
func PossibleHealthCheckStatusValues() []HealthCheckStatus {
	return []HealthCheckStatus{
		HealthCheckStatusFailed,
		HealthCheckStatusPassed,
		HealthCheckStatusPending,
		HealthCheckStatusRunning,
		HealthCheckStatusUnknown,
		HealthCheckStatusWarning,
	}
}

// HealthStatus - Health status indicating whether a pool is available to create Dev Boxes.
type HealthStatus string

const (
	HealthStatusHealthy   HealthStatus = "Healthy"
	HealthStatusPending   HealthStatus = "Pending"
	HealthStatusUnhealthy HealthStatus = "Unhealthy"
	HealthStatusUnknown   HealthStatus = "Unknown"
	HealthStatusWarning   HealthStatus = "Warning"
)

// PossibleHealthStatusValues returns the possible values for the HealthStatus const type.
func PossibleHealthStatusValues() []HealthStatus {
	return []HealthStatus{
		HealthStatusHealthy,
		HealthStatusPending,
		HealthStatusUnhealthy,
		HealthStatusUnknown,
		HealthStatusWarning,
	}
}

// HibernateSupport - Indicates whether hibernate is enabled/disabled.
type HibernateSupport string

const (
	HibernateSupportDisabled HibernateSupport = "Disabled"
	HibernateSupportEnabled  HibernateSupport = "Enabled"
)

// PossibleHibernateSupportValues returns the possible values for the HibernateSupport const type.
func PossibleHibernateSupportValues() []HibernateSupport {
	return []HibernateSupport{
		HibernateSupportDisabled,
		HibernateSupportEnabled,
	}
}

// IdentityType - Values can be systemAssignedIdentity or userAssignedIdentity
type IdentityType string

const (
	IdentityTypeDelegatedResourceIdentity IdentityType = "delegatedResourceIdentity"
	IdentityTypeSystemAssignedIdentity    IdentityType = "systemAssignedIdentity"
	IdentityTypeUserAssignedIdentity      IdentityType = "userAssignedIdentity"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeDelegatedResourceIdentity,
		IdentityTypeSystemAssignedIdentity,
		IdentityTypeUserAssignedIdentity,
	}
}

// ImageDefinitionBuildStatus - The state of an Image Definition Build.
type ImageDefinitionBuildStatus string

const (
	// ImageDefinitionBuildStatusCancelled - The image build has been cancelled.
	ImageDefinitionBuildStatusCancelled ImageDefinitionBuildStatus = "Cancelled"
	// ImageDefinitionBuildStatusFailed - The image build has failed.
	ImageDefinitionBuildStatusFailed ImageDefinitionBuildStatus = "Failed"
	// ImageDefinitionBuildStatusRunning - The image build is running.
	ImageDefinitionBuildStatusRunning ImageDefinitionBuildStatus = "Running"
	// ImageDefinitionBuildStatusSucceeded - The image build has succeeded.
	ImageDefinitionBuildStatusSucceeded ImageDefinitionBuildStatus = "Succeeded"
	// ImageDefinitionBuildStatusTimedOut - The image build has timed out.
	ImageDefinitionBuildStatusTimedOut ImageDefinitionBuildStatus = "TimedOut"
	// ImageDefinitionBuildStatusValidationFailed - The built image has failed validation.
	ImageDefinitionBuildStatusValidationFailed ImageDefinitionBuildStatus = "ValidationFailed"
)

// PossibleImageDefinitionBuildStatusValues returns the possible values for the ImageDefinitionBuildStatus const type.
func PossibleImageDefinitionBuildStatusValues() []ImageDefinitionBuildStatus {
	return []ImageDefinitionBuildStatus{
		ImageDefinitionBuildStatusCancelled,
		ImageDefinitionBuildStatusFailed,
		ImageDefinitionBuildStatusRunning,
		ImageDefinitionBuildStatusSucceeded,
		ImageDefinitionBuildStatusTimedOut,
		ImageDefinitionBuildStatusValidationFailed,
	}
}

// ImageValidationStatus - Image validation status
type ImageValidationStatus string

const (
	ImageValidationStatusFailed    ImageValidationStatus = "Failed"
	ImageValidationStatusPending   ImageValidationStatus = "Pending"
	ImageValidationStatusSucceeded ImageValidationStatus = "Succeeded"
	ImageValidationStatusTimedOut  ImageValidationStatus = "TimedOut"
	ImageValidationStatusUnknown   ImageValidationStatus = "Unknown"
)

// PossibleImageValidationStatusValues returns the possible values for the ImageValidationStatus const type.
func PossibleImageValidationStatusValues() []ImageValidationStatus {
	return []ImageValidationStatus{
		ImageValidationStatusFailed,
		ImageValidationStatusPending,
		ImageValidationStatusSucceeded,
		ImageValidationStatusTimedOut,
		ImageValidationStatusUnknown,
	}
}

// InstallAzureMonitorAgentEnableStatus - Setting to be used when determining whether to install the Azure Monitor Agent service
// on Dev Boxes that belong to this dev center.
type InstallAzureMonitorAgentEnableStatus string

const (
	InstallAzureMonitorAgentEnableStatusDisabled InstallAzureMonitorAgentEnableStatus = "Disabled"
	InstallAzureMonitorAgentEnableStatusEnabled  InstallAzureMonitorAgentEnableStatus = "Enabled"
)

// PossibleInstallAzureMonitorAgentEnableStatusValues returns the possible values for the InstallAzureMonitorAgentEnableStatus const type.
func PossibleInstallAzureMonitorAgentEnableStatusValues() []InstallAzureMonitorAgentEnableStatus {
	return []InstallAzureMonitorAgentEnableStatus{
		InstallAzureMonitorAgentEnableStatusDisabled,
		InstallAzureMonitorAgentEnableStatusEnabled,
	}
}

// LicenseType - License Types
type LicenseType string

const (
	LicenseTypeWindowsClient LicenseType = "Windows_Client"
)

// PossibleLicenseTypeValues returns the possible values for the LicenseType const type.
func PossibleLicenseTypeValues() []LicenseType {
	return []LicenseType{
		LicenseTypeWindowsClient,
	}
}

// LocalAdminStatus - Local Administrator enable or disable status. Indicates whether owners of Dev Boxes are added as local
// administrators on the Dev Box.
type LocalAdminStatus string

const (
	LocalAdminStatusDisabled LocalAdminStatus = "Disabled"
	LocalAdminStatusEnabled  LocalAdminStatus = "Enabled"
)

// PossibleLocalAdminStatusValues returns the possible values for the LocalAdminStatus const type.
func PossibleLocalAdminStatusValues() []LocalAdminStatus {
	return []LocalAdminStatus{
		LocalAdminStatusDisabled,
		LocalAdminStatusEnabled,
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

// MicrosoftHostedNetworkEnableStatus - Indicates whether pools in this Dev Center can use Microsoft Hosted Networks. Defaults
// to Enabled if not set.
type MicrosoftHostedNetworkEnableStatus string

const (
	MicrosoftHostedNetworkEnableStatusDisabled MicrosoftHostedNetworkEnableStatus = "Disabled"
	MicrosoftHostedNetworkEnableStatusEnabled  MicrosoftHostedNetworkEnableStatus = "Enabled"
)

// PossibleMicrosoftHostedNetworkEnableStatusValues returns the possible values for the MicrosoftHostedNetworkEnableStatus const type.
func PossibleMicrosoftHostedNetworkEnableStatusValues() []MicrosoftHostedNetworkEnableStatus {
	return []MicrosoftHostedNetworkEnableStatus{
		MicrosoftHostedNetworkEnableStatusDisabled,
		MicrosoftHostedNetworkEnableStatusEnabled,
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

// ParameterType - The type of data a parameter accepts.
type ParameterType string

const (
	// ParameterTypeArray - The parameter accepts an array of values.
	ParameterTypeArray ParameterType = "array"
	// ParameterTypeBoolean - The parameter accepts a boolean value.
	ParameterTypeBoolean ParameterType = "boolean"
	// ParameterTypeInteger - The parameter accepts an integer value.
	ParameterTypeInteger ParameterType = "integer"
	// ParameterTypeNumber - The parameter accepts a number value.
	ParameterTypeNumber ParameterType = "number"
	// ParameterTypeObject - The parameter accepts an object value.
	ParameterTypeObject ParameterType = "object"
	// ParameterTypeString - The parameter accepts a string value.
	ParameterTypeString ParameterType = "string"
)

// PossibleParameterTypeValues returns the possible values for the ParameterType const type.
func PossibleParameterTypeValues() []ParameterType {
	return []ParameterType{
		ParameterTypeArray,
		ParameterTypeBoolean,
		ParameterTypeInteger,
		ParameterTypeNumber,
		ParameterTypeObject,
		ParameterTypeString,
	}
}

// PlanMemberSyncState - The synchronization state of the plan member.
type PlanMemberSyncState string

const (
	PlanMemberSyncStateCanceled   PlanMemberSyncState = "Canceled"
	PlanMemberSyncStateFailed     PlanMemberSyncState = "Failed"
	PlanMemberSyncStateInProgress PlanMemberSyncState = "InProgress"
	PlanMemberSyncStateNotStarted PlanMemberSyncState = "NotStarted"
	PlanMemberSyncStateSucceeded  PlanMemberSyncState = "Succeeded"
)

// PossiblePlanMemberSyncStateValues returns the possible values for the PlanMemberSyncState const type.
func PossiblePlanMemberSyncStateValues() []PlanMemberSyncState {
	return []PlanMemberSyncState{
		PlanMemberSyncStateCanceled,
		PlanMemberSyncStateFailed,
		PlanMemberSyncStateInProgress,
		PlanMemberSyncStateNotStarted,
		PlanMemberSyncStateSucceeded,
	}
}

// PlanMemberType - The type of the member (user, group)
type PlanMemberType string

const (
	PlanMemberTypeGroup PlanMemberType = "Group"
	PlanMemberTypeUser  PlanMemberType = "User"
)

// PossiblePlanMemberTypeValues returns the possible values for the PlanMemberType const type.
func PossiblePlanMemberTypeValues() []PlanMemberType {
	return []PlanMemberType{
		PlanMemberTypeGroup,
		PlanMemberTypeUser,
	}
}

// PoolDevBoxDefinitionType - Indicates if the pool is created from an existing Dev Box Definition or if one is provided directly.
type PoolDevBoxDefinitionType string

const (
	PoolDevBoxDefinitionTypeReference PoolDevBoxDefinitionType = "Reference"
	PoolDevBoxDefinitionTypeValue     PoolDevBoxDefinitionType = "Value"
)

// PossiblePoolDevBoxDefinitionTypeValues returns the possible values for the PoolDevBoxDefinitionType const type.
func PossiblePoolDevBoxDefinitionTypeValues() []PoolDevBoxDefinitionType {
	return []PoolDevBoxDefinitionType{
		PoolDevBoxDefinitionTypeReference,
		PoolDevBoxDefinitionTypeValue,
	}
}

// ProvisioningState - Provisioning state of the resource.
type ProvisioningState string

const (
	ProvisioningStateAccepted                  ProvisioningState = "Accepted"
	ProvisioningStateCanceled                  ProvisioningState = "Canceled"
	ProvisioningStateCreated                   ProvisioningState = "Created"
	ProvisioningStateCreating                  ProvisioningState = "Creating"
	ProvisioningStateDeleted                   ProvisioningState = "Deleted"
	ProvisioningStateDeleting                  ProvisioningState = "Deleting"
	ProvisioningStateFailed                    ProvisioningState = "Failed"
	ProvisioningStateMovingResources           ProvisioningState = "MovingResources"
	ProvisioningStateNotSpecified              ProvisioningState = "NotSpecified"
	ProvisioningStateRolloutInProgress         ProvisioningState = "RolloutInProgress"
	ProvisioningStateRunning                   ProvisioningState = "Running"
	ProvisioningStateStorageProvisioningFailed ProvisioningState = "StorageProvisioningFailed"
	ProvisioningStateSucceeded                 ProvisioningState = "Succeeded"
	ProvisioningStateTransientFailure          ProvisioningState = "TransientFailure"
	ProvisioningStateUpdated                   ProvisioningState = "Updated"
	ProvisioningStateUpdating                  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreated,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateMovingResources,
		ProvisioningStateNotSpecified,
		ProvisioningStateRolloutInProgress,
		ProvisioningStateRunning,
		ProvisioningStateStorageProvisioningFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateTransientFailure,
		ProvisioningStateUpdated,
		ProvisioningStateUpdating,
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

// ScheduleEnableStatus - Schedule enable or disable status. Indicates whether the schedule applied to is either enabled or
// disabled.
type ScheduleEnableStatus string

const (
	ScheduleEnableStatusDisabled ScheduleEnableStatus = "Disabled"
	ScheduleEnableStatusEnabled  ScheduleEnableStatus = "Enabled"
)

// PossibleScheduleEnableStatusValues returns the possible values for the ScheduleEnableStatus const type.
func PossibleScheduleEnableStatusValues() []ScheduleEnableStatus {
	return []ScheduleEnableStatus{
		ScheduleEnableStatusDisabled,
		ScheduleEnableStatusEnabled,
	}
}

// ScheduledFrequency - The frequency of task execution.
type ScheduledFrequency string

const (
	ScheduledFrequencyDaily ScheduledFrequency = "Daily"
)

// PossibleScheduledFrequencyValues returns the possible values for the ScheduledFrequency const type.
func PossibleScheduledFrequencyValues() []ScheduledFrequency {
	return []ScheduledFrequency{
		ScheduledFrequencyDaily,
	}
}

// ScheduledType - The supported types for a scheduled task.
type ScheduledType string

const (
	ScheduledTypeStopDevBox ScheduledType = "StopDevBox"
)

// PossibleScheduledTypeValues returns the possible values for the ScheduledType const type.
func PossibleScheduledTypeValues() []ScheduledType {
	return []ScheduledType{
		ScheduledTypeStopDevBox,
	}
}

// SingleSignOnStatus - SingleSignOn (SSO) enable or disable status. Indicates whether Dev Boxes in the Pool will have SSO
// enabled or disabled.
type SingleSignOnStatus string

const (
	SingleSignOnStatusDisabled SingleSignOnStatus = "Disabled"
	SingleSignOnStatusEnabled  SingleSignOnStatus = "Enabled"
)

// PossibleSingleSignOnStatusValues returns the possible values for the SingleSignOnStatus const type.
func PossibleSingleSignOnStatusValues() []SingleSignOnStatus {
	return []SingleSignOnStatus{
		SingleSignOnStatusDisabled,
		SingleSignOnStatusEnabled,
	}
}

// StopOnDisconnectEnableStatus - Stop on disconnect enable or disable status. Indicates whether stop on disconnect to is
// either enabled or disabled.
type StopOnDisconnectEnableStatus string

const (
	StopOnDisconnectEnableStatusDisabled StopOnDisconnectEnableStatus = "Disabled"
	StopOnDisconnectEnableStatusEnabled  StopOnDisconnectEnableStatus = "Enabled"
)

// PossibleStopOnDisconnectEnableStatusValues returns the possible values for the StopOnDisconnectEnableStatus const type.
func PossibleStopOnDisconnectEnableStatusValues() []StopOnDisconnectEnableStatus {
	return []StopOnDisconnectEnableStatus{
		StopOnDisconnectEnableStatusDisabled,
		StopOnDisconnectEnableStatusEnabled,
	}
}

// UsageUnit - The unit details.
type UsageUnit string

const (
	UsageUnitCount UsageUnit = "Count"
)

// PossibleUsageUnitValues returns the possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{
		UsageUnitCount,
	}
}

// VirtualNetworkType - Indicates a pool uses a Virtual Network managed by Microsoft (Managed), or a customer provided Network
// (Unmanaged).
type VirtualNetworkType string

const (
	VirtualNetworkTypeManaged   VirtualNetworkType = "Managed"
	VirtualNetworkTypeUnmanaged VirtualNetworkType = "Unmanaged"
)

// PossibleVirtualNetworkTypeValues returns the possible values for the VirtualNetworkType const type.
func PossibleVirtualNetworkTypeValues() []VirtualNetworkType {
	return []VirtualNetworkType{
		VirtualNetworkTypeManaged,
		VirtualNetworkTypeUnmanaged,
	}
}
