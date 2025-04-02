// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataprotection

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
	moduleVersion = "v4.0.0"
)

type AKSVolumeTypes string

const (
	AKSVolumeTypesAzureDisk         AKSVolumeTypes = "AzureDisk"
	AKSVolumeTypesAzureFileShareSMB AKSVolumeTypes = "AzureFileShareSMB"
)

// PossibleAKSVolumeTypesValues returns the possible values for the AKSVolumeTypes const type.
func PossibleAKSVolumeTypesValues() []AKSVolumeTypes {
	return []AKSVolumeTypes{
		AKSVolumeTypesAzureDisk,
		AKSVolumeTypesAzureFileShareSMB,
	}
}

type AbsoluteMarker string

const (
	AbsoluteMarkerAllBackup    AbsoluteMarker = "AllBackup"
	AbsoluteMarkerFirstOfDay   AbsoluteMarker = "FirstOfDay"
	AbsoluteMarkerFirstOfMonth AbsoluteMarker = "FirstOfMonth"
	AbsoluteMarkerFirstOfWeek  AbsoluteMarker = "FirstOfWeek"
	AbsoluteMarkerFirstOfYear  AbsoluteMarker = "FirstOfYear"
)

// PossibleAbsoluteMarkerValues returns the possible values for the AbsoluteMarker const type.
func PossibleAbsoluteMarkerValues() []AbsoluteMarker {
	return []AbsoluteMarker{
		AbsoluteMarkerAllBackup,
		AbsoluteMarkerFirstOfDay,
		AbsoluteMarkerFirstOfMonth,
		AbsoluteMarkerFirstOfWeek,
		AbsoluteMarkerFirstOfYear,
	}
}

type AlertsState string

const (
	AlertsStateDisabled AlertsState = "Disabled"
	AlertsStateEnabled  AlertsState = "Enabled"
)

// PossibleAlertsStateValues returns the possible values for the AlertsState const type.
func PossibleAlertsStateValues() []AlertsState {
	return []AlertsState{
		AlertsStateDisabled,
		AlertsStateEnabled,
	}
}

// BCDRSecurityLevel - Security Level of Backup Vault
type BCDRSecurityLevel string

const (
	BCDRSecurityLevelExcellent    BCDRSecurityLevel = "Excellent"
	BCDRSecurityLevelFair         BCDRSecurityLevel = "Fair"
	BCDRSecurityLevelGood         BCDRSecurityLevel = "Good"
	BCDRSecurityLevelNotSupported BCDRSecurityLevel = "NotSupported"
	BCDRSecurityLevelPoor         BCDRSecurityLevel = "Poor"
)

// PossibleBCDRSecurityLevelValues returns the possible values for the BCDRSecurityLevel const type.
func PossibleBCDRSecurityLevelValues() []BCDRSecurityLevel {
	return []BCDRSecurityLevel{
		BCDRSecurityLevelExcellent,
		BCDRSecurityLevelFair,
		BCDRSecurityLevelGood,
		BCDRSecurityLevelNotSupported,
		BCDRSecurityLevelPoor,
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

// CrossRegionRestoreState - CrossRegionRestore state
type CrossRegionRestoreState string

const (
	CrossRegionRestoreStateDisabled CrossRegionRestoreState = "Disabled"
	CrossRegionRestoreStateEnabled  CrossRegionRestoreState = "Enabled"
)

// PossibleCrossRegionRestoreStateValues returns the possible values for the CrossRegionRestoreState const type.
func PossibleCrossRegionRestoreStateValues() []CrossRegionRestoreState {
	return []CrossRegionRestoreState{
		CrossRegionRestoreStateDisabled,
		CrossRegionRestoreStateEnabled,
	}
}

// CrossSubscriptionRestoreState - CrossSubscriptionRestore state
type CrossSubscriptionRestoreState string

const (
	CrossSubscriptionRestoreStateDisabled            CrossSubscriptionRestoreState = "Disabled"
	CrossSubscriptionRestoreStateEnabled             CrossSubscriptionRestoreState = "Enabled"
	CrossSubscriptionRestoreStatePermanentlyDisabled CrossSubscriptionRestoreState = "PermanentlyDisabled"
)

// PossibleCrossSubscriptionRestoreStateValues returns the possible values for the CrossSubscriptionRestoreState const type.
func PossibleCrossSubscriptionRestoreStateValues() []CrossSubscriptionRestoreState {
	return []CrossSubscriptionRestoreState{
		CrossSubscriptionRestoreStateDisabled,
		CrossSubscriptionRestoreStateEnabled,
		CrossSubscriptionRestoreStatePermanentlyDisabled,
	}
}

// CurrentProtectionState - Specifies the current protection state of the resource
type CurrentProtectionState string

const (
	CurrentProtectionStateBackupSchedulesSuspended    CurrentProtectionState = "BackupSchedulesSuspended"
	CurrentProtectionStateConfiguringProtection       CurrentProtectionState = "ConfiguringProtection"
	CurrentProtectionStateConfiguringProtectionFailed CurrentProtectionState = "ConfiguringProtectionFailed"
	CurrentProtectionStateInvalid                     CurrentProtectionState = "Invalid"
	CurrentProtectionStateNotProtected                CurrentProtectionState = "NotProtected"
	CurrentProtectionStateProtectionConfigured        CurrentProtectionState = "ProtectionConfigured"
	CurrentProtectionStateProtectionError             CurrentProtectionState = "ProtectionError"
	CurrentProtectionStateProtectionStopped           CurrentProtectionState = "ProtectionStopped"
	CurrentProtectionStateRetentionSchedulesSuspended CurrentProtectionState = "RetentionSchedulesSuspended"
	CurrentProtectionStateSoftDeleted                 CurrentProtectionState = "SoftDeleted"
	CurrentProtectionStateSoftDeleting                CurrentProtectionState = "SoftDeleting"
	CurrentProtectionStateUpdatingProtection          CurrentProtectionState = "UpdatingProtection"
)

// PossibleCurrentProtectionStateValues returns the possible values for the CurrentProtectionState const type.
func PossibleCurrentProtectionStateValues() []CurrentProtectionState {
	return []CurrentProtectionState{
		CurrentProtectionStateBackupSchedulesSuspended,
		CurrentProtectionStateConfiguringProtection,
		CurrentProtectionStateConfiguringProtectionFailed,
		CurrentProtectionStateInvalid,
		CurrentProtectionStateNotProtected,
		CurrentProtectionStateProtectionConfigured,
		CurrentProtectionStateProtectionError,
		CurrentProtectionStateProtectionStopped,
		CurrentProtectionStateRetentionSchedulesSuspended,
		CurrentProtectionStateSoftDeleted,
		CurrentProtectionStateSoftDeleting,
		CurrentProtectionStateUpdatingProtection,
	}
}

// DataStoreTypes - type of datastore; Operational/Vault/Archive
type DataStoreTypes string

const (
	DataStoreTypesArchiveStore     DataStoreTypes = "ArchiveStore"
	DataStoreTypesOperationalStore DataStoreTypes = "OperationalStore"
	DataStoreTypesVaultStore       DataStoreTypes = "VaultStore"
)

// PossibleDataStoreTypesValues returns the possible values for the DataStoreTypes const type.
func PossibleDataStoreTypesValues() []DataStoreTypes {
	return []DataStoreTypes{
		DataStoreTypesArchiveStore,
		DataStoreTypesOperationalStore,
		DataStoreTypesVaultStore,
	}
}

type DayOfWeek string

const (
	DayOfWeekFriday    DayOfWeek = "Friday"
	DayOfWeekMonday    DayOfWeek = "Monday"
	DayOfWeekSaturday  DayOfWeek = "Saturday"
	DayOfWeekSunday    DayOfWeek = "Sunday"
	DayOfWeekThursday  DayOfWeek = "Thursday"
	DayOfWeekTuesday   DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
)

// PossibleDayOfWeekValues returns the possible values for the DayOfWeek const type.
func PossibleDayOfWeekValues() []DayOfWeek {
	return []DayOfWeek{
		DayOfWeekFriday,
		DayOfWeekMonday,
		DayOfWeekSaturday,
		DayOfWeekSunday,
		DayOfWeekThursday,
		DayOfWeekTuesday,
		DayOfWeekWednesday,
	}
}

// EncryptionState - Encryption state of the Backup Vault.
type EncryptionState string

const (
	// EncryptionStateDisabled - CMK encryption is disabled on the Backup Vault. User can not set this state once Encryption State
	// is 'Enabled'.
	EncryptionStateDisabled EncryptionState = "Disabled"
	// EncryptionStateEnabled - CMK encryption is enabled on the Backup Vault
	EncryptionStateEnabled EncryptionState = "Enabled"
	// EncryptionStateInconsistent - CMK encryption is in inconsistent state on the Backup Vault. This state indicates that user
	// needs to retry the encryption settings operation immediately to correct the state.
	EncryptionStateInconsistent EncryptionState = "Inconsistent"
)

// PossibleEncryptionStateValues returns the possible values for the EncryptionState const type.
func PossibleEncryptionStateValues() []EncryptionState {
	return []EncryptionState{
		EncryptionStateDisabled,
		EncryptionStateEnabled,
		EncryptionStateInconsistent,
	}
}

// ExistingResourcePolicy - Gets or sets the Conflict Policy property. This property sets policy during conflict of resources
// during restore.
type ExistingResourcePolicy string

const (
	ExistingResourcePolicyPatch ExistingResourcePolicy = "Patch"
	ExistingResourcePolicySkip  ExistingResourcePolicy = "Skip"
)

// PossibleExistingResourcePolicyValues returns the possible values for the ExistingResourcePolicy const type.
func PossibleExistingResourcePolicyValues() []ExistingResourcePolicy {
	return []ExistingResourcePolicy{
		ExistingResourcePolicyPatch,
		ExistingResourcePolicySkip,
	}
}

// FeatureSupportStatus - feature support status
type FeatureSupportStatus string

const (
	FeatureSupportStatusAlphaPreview       FeatureSupportStatus = "AlphaPreview"
	FeatureSupportStatusGenerallyAvailable FeatureSupportStatus = "GenerallyAvailable"
	FeatureSupportStatusInvalid            FeatureSupportStatus = "Invalid"
	FeatureSupportStatusNotSupported       FeatureSupportStatus = "NotSupported"
	FeatureSupportStatusPrivatePreview     FeatureSupportStatus = "PrivatePreview"
	FeatureSupportStatusPublicPreview      FeatureSupportStatus = "PublicPreview"
)

// PossibleFeatureSupportStatusValues returns the possible values for the FeatureSupportStatus const type.
func PossibleFeatureSupportStatusValues() []FeatureSupportStatus {
	return []FeatureSupportStatus{
		FeatureSupportStatusAlphaPreview,
		FeatureSupportStatusGenerallyAvailable,
		FeatureSupportStatusInvalid,
		FeatureSupportStatusNotSupported,
		FeatureSupportStatusPrivatePreview,
		FeatureSupportStatusPublicPreview,
	}
}

// FeatureType - backup support feature type.
type FeatureType string

const (
	FeatureTypeDataSourceType FeatureType = "DataSourceType"
	FeatureTypeInvalid        FeatureType = "Invalid"
)

// PossibleFeatureTypeValues returns the possible values for the FeatureType const type.
func PossibleFeatureTypeValues() []FeatureType {
	return []FeatureType{
		FeatureTypeDataSourceType,
		FeatureTypeInvalid,
	}
}

// IdentityType - The identity type. 'SystemAssigned' and 'UserAssigned' are mutually exclusive. 'SystemAssigned' will use
// implicitly created managed identity.
type IdentityType string

const (
	IdentityTypeSystemAssigned IdentityType = "SystemAssigned"
	IdentityTypeUserAssigned   IdentityType = "UserAssigned"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeSystemAssigned,
		IdentityTypeUserAssigned,
	}
}

// ImmutabilityState - Immutability state
type ImmutabilityState string

const (
	ImmutabilityStateDisabled ImmutabilityState = "Disabled"
	ImmutabilityStateLocked   ImmutabilityState = "Locked"
	ImmutabilityStateUnlocked ImmutabilityState = "Unlocked"
)

// PossibleImmutabilityStateValues returns the possible values for the ImmutabilityState const type.
func PossibleImmutabilityStateValues() []ImmutabilityState {
	return []ImmutabilityState{
		ImmutabilityStateDisabled,
		ImmutabilityStateLocked,
		ImmutabilityStateUnlocked,
	}
}

// InfrastructureEncryptionState - Enabling/Disabling the Double Encryption state
type InfrastructureEncryptionState string

const (
	InfrastructureEncryptionStateDisabled InfrastructureEncryptionState = "Disabled"
	InfrastructureEncryptionStateEnabled  InfrastructureEncryptionState = "Enabled"
)

// PossibleInfrastructureEncryptionStateValues returns the possible values for the InfrastructureEncryptionState const type.
func PossibleInfrastructureEncryptionStateValues() []InfrastructureEncryptionState {
	return []InfrastructureEncryptionState{
		InfrastructureEncryptionStateDisabled,
		InfrastructureEncryptionStateEnabled,
	}
}

type Month string

const (
	MonthApril     Month = "April"
	MonthAugust    Month = "August"
	MonthDecember  Month = "December"
	MonthFebruary  Month = "February"
	MonthJanuary   Month = "January"
	MonthJuly      Month = "July"
	MonthJune      Month = "June"
	MonthMarch     Month = "March"
	MonthMay       Month = "May"
	MonthNovember  Month = "November"
	MonthOctober   Month = "October"
	MonthSeptember Month = "September"
)

// PossibleMonthValues returns the possible values for the Month const type.
func PossibleMonthValues() []Month {
	return []Month{
		MonthApril,
		MonthAugust,
		MonthDecember,
		MonthFebruary,
		MonthJanuary,
		MonthJuly,
		MonthJune,
		MonthMarch,
		MonthMay,
		MonthNovember,
		MonthOctober,
		MonthSeptember,
	}
}

// PersistentVolumeRestoreMode - Gets or sets the PV (Persistent Volume) Restore Mode property. This property sets whether
// volumes needs to be restored.
type PersistentVolumeRestoreMode string

const (
	PersistentVolumeRestoreModeRestoreWithVolumeData    PersistentVolumeRestoreMode = "RestoreWithVolumeData"
	PersistentVolumeRestoreModeRestoreWithoutVolumeData PersistentVolumeRestoreMode = "RestoreWithoutVolumeData"
)

// PossiblePersistentVolumeRestoreModeValues returns the possible values for the PersistentVolumeRestoreMode const type.
func PossiblePersistentVolumeRestoreModeValues() []PersistentVolumeRestoreMode {
	return []PersistentVolumeRestoreMode{
		PersistentVolumeRestoreModeRestoreWithVolumeData,
		PersistentVolumeRestoreModeRestoreWithoutVolumeData,
	}
}

// ProvisioningState - Provisioning state of the BackupVault resource
type ProvisioningState string

const (
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUnknown      ProvisioningState = "Unknown"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUnknown,
		ProvisioningStateUpdating,
	}
}

// RecoveryOption - Recovery Option
type RecoveryOption string

const (
	RecoveryOptionFailIfExists RecoveryOption = "FailIfExists"
)

// PossibleRecoveryOptionValues returns the possible values for the RecoveryOption const type.
func PossibleRecoveryOptionValues() []RecoveryOption {
	return []RecoveryOption{
		RecoveryOptionFailIfExists,
	}
}

// RecoveryPointCompletionState - Specifies recovery point completeness. Partial (i.e., only some of the intended items were
// backed up), or Completed (i.e., ALL intended items were backed up).
type RecoveryPointCompletionState string

const (
	RecoveryPointCompletionStateCompleted RecoveryPointCompletionState = "Completed"
	RecoveryPointCompletionStatePartial   RecoveryPointCompletionState = "Partial"
)

// PossibleRecoveryPointCompletionStateValues returns the possible values for the RecoveryPointCompletionState const type.
func PossibleRecoveryPointCompletionStateValues() []RecoveryPointCompletionState {
	return []RecoveryPointCompletionState{
		RecoveryPointCompletionStateCompleted,
		RecoveryPointCompletionStatePartial,
	}
}

// RehydrationPriority - Priority to be used for rehydration. Values High or Standard
type RehydrationPriority string

const (
	RehydrationPriorityHigh     RehydrationPriority = "High"
	RehydrationPriorityInvalid  RehydrationPriority = "Invalid"
	RehydrationPriorityStandard RehydrationPriority = "Standard"
)

// PossibleRehydrationPriorityValues returns the possible values for the RehydrationPriority const type.
func PossibleRehydrationPriorityValues() []RehydrationPriority {
	return []RehydrationPriority{
		RehydrationPriorityHigh,
		RehydrationPriorityInvalid,
		RehydrationPriorityStandard,
	}
}

type RehydrationStatus string

const (
	RehydrationStatusCOMPLETED        RehydrationStatus = "COMPLETED"
	RehydrationStatusCREATEINPROGRESS RehydrationStatus = "CREATE_IN_PROGRESS"
	RehydrationStatusDELETED          RehydrationStatus = "DELETED"
	RehydrationStatusDELETEINPROGRESS RehydrationStatus = "DELETE_IN_PROGRESS"
	RehydrationStatusFAILED           RehydrationStatus = "FAILED"
)

// PossibleRehydrationStatusValues returns the possible values for the RehydrationStatus const type.
func PossibleRehydrationStatusValues() []RehydrationStatus {
	return []RehydrationStatus{
		RehydrationStatusCOMPLETED,
		RehydrationStatusCREATEINPROGRESS,
		RehydrationStatusDELETED,
		RehydrationStatusDELETEINPROGRESS,
		RehydrationStatusFAILED,
	}
}

// ResourceMoveState - Resource move state for backup vault
type ResourceMoveState string

const (
	ResourceMoveStateCommitFailed    ResourceMoveState = "CommitFailed"
	ResourceMoveStateCommitTimedout  ResourceMoveState = "CommitTimedout"
	ResourceMoveStateCriticalFailure ResourceMoveState = "CriticalFailure"
	ResourceMoveStateFailed          ResourceMoveState = "Failed"
	ResourceMoveStateInProgress      ResourceMoveState = "InProgress"
	ResourceMoveStateMoveSucceeded   ResourceMoveState = "MoveSucceeded"
	ResourceMoveStatePartialSuccess  ResourceMoveState = "PartialSuccess"
	ResourceMoveStatePrepareFailed   ResourceMoveState = "PrepareFailed"
	ResourceMoveStatePrepareTimedout ResourceMoveState = "PrepareTimedout"
	ResourceMoveStateUnknown         ResourceMoveState = "Unknown"
)

// PossibleResourceMoveStateValues returns the possible values for the ResourceMoveState const type.
func PossibleResourceMoveStateValues() []ResourceMoveState {
	return []ResourceMoveState{
		ResourceMoveStateCommitFailed,
		ResourceMoveStateCommitTimedout,
		ResourceMoveStateCriticalFailure,
		ResourceMoveStateFailed,
		ResourceMoveStateInProgress,
		ResourceMoveStateMoveSucceeded,
		ResourceMoveStatePartialSuccess,
		ResourceMoveStatePrepareFailed,
		ResourceMoveStatePrepareTimedout,
		ResourceMoveStateUnknown,
	}
}

// ResourcePropertiesObjectType - Type of the specific object - used for deserializing
type ResourcePropertiesObjectType string

const (
	ResourcePropertiesObjectTypeDefaultResourceProperties ResourcePropertiesObjectType = "DefaultResourceProperties"
)

// PossibleResourcePropertiesObjectTypeValues returns the possible values for the ResourcePropertiesObjectType const type.
func PossibleResourcePropertiesObjectTypeValues() []ResourcePropertiesObjectType {
	return []ResourcePropertiesObjectType{
		ResourcePropertiesObjectTypeDefaultResourceProperties,
	}
}

// RestoreSourceDataStoreType - Gets or sets the type of the source data store.
type RestoreSourceDataStoreType string

const (
	RestoreSourceDataStoreTypeArchiveStore     RestoreSourceDataStoreType = "ArchiveStore"
	RestoreSourceDataStoreTypeOperationalStore RestoreSourceDataStoreType = "OperationalStore"
)

// PossibleRestoreSourceDataStoreTypeValues returns the possible values for the RestoreSourceDataStoreType const type.
func PossibleRestoreSourceDataStoreTypeValues() []RestoreSourceDataStoreType {
	return []RestoreSourceDataStoreType{
		RestoreSourceDataStoreTypeArchiveStore,
		RestoreSourceDataStoreTypeOperationalStore,
	}
}

// RestoreTargetLocationType - Denotes the target location where the data will be restored, string value for the enum {Microsoft.Internal.AzureBackup.DataProtection.Common.Interface.RestoreTargetLocationType}
type RestoreTargetLocationType string

const (
	RestoreTargetLocationTypeAzureBlobs RestoreTargetLocationType = "AzureBlobs"
	RestoreTargetLocationTypeAzureFiles RestoreTargetLocationType = "AzureFiles"
	RestoreTargetLocationTypeInvalid    RestoreTargetLocationType = "Invalid"
)

// PossibleRestoreTargetLocationTypeValues returns the possible values for the RestoreTargetLocationType const type.
func PossibleRestoreTargetLocationTypeValues() []RestoreTargetLocationType {
	return []RestoreTargetLocationType{
		RestoreTargetLocationTypeAzureBlobs,
		RestoreTargetLocationTypeAzureFiles,
		RestoreTargetLocationTypeInvalid,
	}
}

// SecretStoreType - Gets or sets the type of secret store
type SecretStoreType string

const (
	SecretStoreTypeAzureKeyVault SecretStoreType = "AzureKeyVault"
	SecretStoreTypeInvalid       SecretStoreType = "Invalid"
)

// PossibleSecretStoreTypeValues returns the possible values for the SecretStoreType const type.
func PossibleSecretStoreTypeValues() []SecretStoreType {
	return []SecretStoreType{
		SecretStoreTypeAzureKeyVault,
		SecretStoreTypeInvalid,
	}
}

// SecureScoreLevel - Secure Score of Backup Vault
type SecureScoreLevel string

const (
	SecureScoreLevelAdequate     SecureScoreLevel = "Adequate"
	SecureScoreLevelMaximum      SecureScoreLevel = "Maximum"
	SecureScoreLevelMinimum      SecureScoreLevel = "Minimum"
	SecureScoreLevelNone         SecureScoreLevel = "None"
	SecureScoreLevelNotSupported SecureScoreLevel = "NotSupported"
)

// PossibleSecureScoreLevelValues returns the possible values for the SecureScoreLevel const type.
func PossibleSecureScoreLevelValues() []SecureScoreLevel {
	return []SecureScoreLevel{
		SecureScoreLevelAdequate,
		SecureScoreLevelMaximum,
		SecureScoreLevelMinimum,
		SecureScoreLevelNone,
		SecureScoreLevelNotSupported,
	}
}

// SoftDeleteState - State of soft delete
type SoftDeleteState string

const (
	// SoftDeleteStateAlwaysOn - Soft Delete is permanently enabled for the BackupVault and the setting cannot be changed
	SoftDeleteStateAlwaysOn SoftDeleteState = "AlwaysOn"
	// SoftDeleteStateOff - Soft Delete is turned off for the BackupVault
	SoftDeleteStateOff SoftDeleteState = "Off"
	// SoftDeleteStateOn - Soft Delete is enabled for the BackupVault but can be turned off
	SoftDeleteStateOn SoftDeleteState = "On"
)

// PossibleSoftDeleteStateValues returns the possible values for the SoftDeleteState const type.
func PossibleSoftDeleteStateValues() []SoftDeleteState {
	return []SoftDeleteState{
		SoftDeleteStateAlwaysOn,
		SoftDeleteStateOff,
		SoftDeleteStateOn,
	}
}

// SourceDataStoreType - Gets or sets the type of the source data store.
type SourceDataStoreType string

const (
	SourceDataStoreTypeArchiveStore     SourceDataStoreType = "ArchiveStore"
	SourceDataStoreTypeOperationalStore SourceDataStoreType = "OperationalStore"
	SourceDataStoreTypeSnapshotStore    SourceDataStoreType = "SnapshotStore"
	SourceDataStoreTypeVaultStore       SourceDataStoreType = "VaultStore"
)

// PossibleSourceDataStoreTypeValues returns the possible values for the SourceDataStoreType const type.
func PossibleSourceDataStoreTypeValues() []SourceDataStoreType {
	return []SourceDataStoreType{
		SourceDataStoreTypeArchiveStore,
		SourceDataStoreTypeOperationalStore,
		SourceDataStoreTypeSnapshotStore,
		SourceDataStoreTypeVaultStore,
	}
}

// Status - Specifies the protection status of the resource
type Status string

const (
	StatusConfiguringProtection       Status = "ConfiguringProtection"
	StatusConfiguringProtectionFailed Status = "ConfiguringProtectionFailed"
	StatusProtectionConfigured        Status = "ProtectionConfigured"
	StatusProtectionStopped           Status = "ProtectionStopped"
	StatusSoftDeleted                 Status = "SoftDeleted"
	StatusSoftDeleting                Status = "SoftDeleting"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusConfiguringProtection,
		StatusConfiguringProtectionFailed,
		StatusProtectionConfigured,
		StatusProtectionStopped,
		StatusSoftDeleted,
		StatusSoftDeleting,
	}
}

// StorageSettingStoreTypes - Gets or sets the type of the datastore.
type StorageSettingStoreTypes string

const (
	StorageSettingStoreTypesArchiveStore     StorageSettingStoreTypes = "ArchiveStore"
	StorageSettingStoreTypesOperationalStore StorageSettingStoreTypes = "OperationalStore"
	StorageSettingStoreTypesVaultStore       StorageSettingStoreTypes = "VaultStore"
)

// PossibleStorageSettingStoreTypesValues returns the possible values for the StorageSettingStoreTypes const type.
func PossibleStorageSettingStoreTypesValues() []StorageSettingStoreTypes {
	return []StorageSettingStoreTypes{
		StorageSettingStoreTypesArchiveStore,
		StorageSettingStoreTypesOperationalStore,
		StorageSettingStoreTypesVaultStore,
	}
}

// StorageSettingTypes - Gets or sets the type.
type StorageSettingTypes string

const (
	StorageSettingTypesGeoRedundant     StorageSettingTypes = "GeoRedundant"
	StorageSettingTypesLocallyRedundant StorageSettingTypes = "LocallyRedundant"
	StorageSettingTypesZoneRedundant    StorageSettingTypes = "ZoneRedundant"
)

// PossibleStorageSettingTypesValues returns the possible values for the StorageSettingTypes const type.
func PossibleStorageSettingTypesValues() []StorageSettingTypes {
	return []StorageSettingTypes{
		StorageSettingTypesGeoRedundant,
		StorageSettingTypesLocallyRedundant,
		StorageSettingTypesZoneRedundant,
	}
}

// SyncType - Field indicating sync type e.g. to sync only in case of failure or in all cases
type SyncType string

const (
	SyncTypeDefault     SyncType = "Default"
	SyncTypeForceResync SyncType = "ForceResync"
)

// PossibleSyncTypeValues returns the possible values for the SyncType const type.
func PossibleSyncTypeValues() []SyncType {
	return []SyncType{
		SyncTypeDefault,
		SyncTypeForceResync,
	}
}

// ValidationType - Specifies the type of validation. In case of DeepValidation, all validations from /validateForBackup API
// will run again.
type ValidationType string

const (
	ValidationTypeDeepValidation    ValidationType = "DeepValidation"
	ValidationTypeShallowValidation ValidationType = "ShallowValidation"
)

// PossibleValidationTypeValues returns the possible values for the ValidationType const type.
func PossibleValidationTypeValues() []ValidationType {
	return []ValidationType{
		ValidationTypeDeepValidation,
		ValidationTypeShallowValidation,
	}
}

type WeekNumber string

const (
	WeekNumberFirst  WeekNumber = "First"
	WeekNumberFourth WeekNumber = "Fourth"
	WeekNumberLast   WeekNumber = "Last"
	WeekNumberSecond WeekNumber = "Second"
	WeekNumberThird  WeekNumber = "Third"
)

// PossibleWeekNumberValues returns the possible values for the WeekNumber const type.
func PossibleWeekNumberValues() []WeekNumber {
	return []WeekNumber{
		WeekNumberFirst,
		WeekNumberFourth,
		WeekNumberLast,
		WeekNumberSecond,
		WeekNumberThird,
	}
}
