// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysqlflexibleservers

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers"
	moduleVersion = "v2.0.0-beta.4"
)

type AdministratorName string

const (
	AdministratorNameActiveDirectory AdministratorName = "ActiveDirectory"
)

// PossibleAdministratorNameValues returns the possible values for the AdministratorName const type.
func PossibleAdministratorNameValues() []AdministratorName {
	return []AdministratorName{
		AdministratorNameActiveDirectory,
	}
}

// AdministratorType - Type of the sever administrator.
type AdministratorType string

const (
	AdministratorTypeActiveDirectory AdministratorType = "ActiveDirectory"
)

// PossibleAdministratorTypeValues returns the possible values for the AdministratorType const type.
func PossibleAdministratorTypeValues() []AdministratorType {
	return []AdministratorType{
		AdministratorTypeActiveDirectory,
	}
}

type AdvancedThreatProtectionName string

const (
	AdvancedThreatProtectionNameDefault AdvancedThreatProtectionName = "Default"
)

// PossibleAdvancedThreatProtectionNameValues returns the possible values for the AdvancedThreatProtectionName const type.
func PossibleAdvancedThreatProtectionNameValues() []AdvancedThreatProtectionName {
	return []AdvancedThreatProtectionName{
		AdvancedThreatProtectionNameDefault,
	}
}

// AdvancedThreatProtectionProvisioningState - The current provisioning state.
type AdvancedThreatProtectionProvisioningState string

const (
	AdvancedThreatProtectionProvisioningStateCanceled  AdvancedThreatProtectionProvisioningState = "Canceled"
	AdvancedThreatProtectionProvisioningStateFailed    AdvancedThreatProtectionProvisioningState = "Failed"
	AdvancedThreatProtectionProvisioningStateSucceeded AdvancedThreatProtectionProvisioningState = "Succeeded"
	AdvancedThreatProtectionProvisioningStateUpdating  AdvancedThreatProtectionProvisioningState = "Updating"
)

// PossibleAdvancedThreatProtectionProvisioningStateValues returns the possible values for the AdvancedThreatProtectionProvisioningState const type.
func PossibleAdvancedThreatProtectionProvisioningStateValues() []AdvancedThreatProtectionProvisioningState {
	return []AdvancedThreatProtectionProvisioningState{
		AdvancedThreatProtectionProvisioningStateCanceled,
		AdvancedThreatProtectionProvisioningStateFailed,
		AdvancedThreatProtectionProvisioningStateSucceeded,
		AdvancedThreatProtectionProvisioningStateUpdating,
	}
}

// AdvancedThreatProtectionState - Specifies the state of the Advanced Threat Protection, whether it is enabled or disabled
// on the server.
type AdvancedThreatProtectionState string

const (
	AdvancedThreatProtectionStateDisabled AdvancedThreatProtectionState = "Disabled"
	AdvancedThreatProtectionStateEnabled  AdvancedThreatProtectionState = "Enabled"
)

// PossibleAdvancedThreatProtectionStateValues returns the possible values for the AdvancedThreatProtectionState const type.
func PossibleAdvancedThreatProtectionStateValues() []AdvancedThreatProtectionState {
	return []AdvancedThreatProtectionState{
		AdvancedThreatProtectionStateDisabled,
		AdvancedThreatProtectionStateEnabled,
	}
}

// BackupFormat - Backup Format for the current backup. (CollatedFormat is INTERNAL – DO NOT USE)
type BackupFormat string

const (
	BackupFormatCollatedFormat BackupFormat = "CollatedFormat"
	BackupFormatRaw            BackupFormat = "Raw"
)

// PossibleBackupFormatValues returns the possible values for the BackupFormat const type.
func PossibleBackupFormatValues() []BackupFormat {
	return []BackupFormat{
		BackupFormatCollatedFormat,
		BackupFormatRaw,
	}
}

type BackupType string

const (
	BackupTypeFULL BackupType = "FULL"
)

// PossibleBackupTypeValues returns the possible values for the BackupType const type.
func PossibleBackupTypeValues() []BackupType {
	return []BackupType{
		BackupTypeFULL,
	}
}

// ConfigurationSource - Source of the configuration.
type ConfigurationSource string

const (
	ConfigurationSourceSystemDefault ConfigurationSource = "system-default"
	ConfigurationSourceUserOverride  ConfigurationSource = "user-override"
)

// PossibleConfigurationSourceValues returns the possible values for the ConfigurationSource const type.
func PossibleConfigurationSourceValues() []ConfigurationSource {
	return []ConfigurationSource{
		ConfigurationSourceSystemDefault,
		ConfigurationSourceUserOverride,
	}
}

// CreateMode - The mode to create a new MySQL server.
type CreateMode string

const (
	CreateModeDefault            CreateMode = "Default"
	CreateModeGeoRestore         CreateMode = "GeoRestore"
	CreateModePointInTimeRestore CreateMode = "PointInTimeRestore"
	CreateModeReplica            CreateMode = "Replica"
)

// PossibleCreateModeValues returns the possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{
		CreateModeDefault,
		CreateModeGeoRestore,
		CreateModePointInTimeRestore,
		CreateModeReplica,
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

// DataEncryptionType - The key type, AzureKeyVault for enable cmk, SystemManaged for disable cmk.
type DataEncryptionType string

const (
	DataEncryptionTypeAzureKeyVault DataEncryptionType = "AzureKeyVault"
	DataEncryptionTypeSystemManaged DataEncryptionType = "SystemManaged"
)

// PossibleDataEncryptionTypeValues returns the possible values for the DataEncryptionType const type.
func PossibleDataEncryptionTypeValues() []DataEncryptionType {
	return []DataEncryptionType{
		DataEncryptionTypeAzureKeyVault,
		DataEncryptionTypeSystemManaged,
	}
}

// EnableStatusEnum - Enum to indicate whether value is 'Enabled' or 'Disabled'
type EnableStatusEnum string

const (
	EnableStatusEnumDisabled EnableStatusEnum = "Disabled"
	EnableStatusEnumEnabled  EnableStatusEnum = "Enabled"
)

// PossibleEnableStatusEnumValues returns the possible values for the EnableStatusEnum const type.
func PossibleEnableStatusEnumValues() []EnableStatusEnum {
	return []EnableStatusEnum{
		EnableStatusEnumDisabled,
		EnableStatusEnumEnabled,
	}
}

// HighAvailabilityMode - High availability mode for a server.
type HighAvailabilityMode string

const (
	HighAvailabilityModeDisabled      HighAvailabilityMode = "Disabled"
	HighAvailabilityModeSameZone      HighAvailabilityMode = "SameZone"
	HighAvailabilityModeZoneRedundant HighAvailabilityMode = "ZoneRedundant"
)

// PossibleHighAvailabilityModeValues returns the possible values for the HighAvailabilityMode const type.
func PossibleHighAvailabilityModeValues() []HighAvailabilityMode {
	return []HighAvailabilityMode{
		HighAvailabilityModeDisabled,
		HighAvailabilityModeSameZone,
		HighAvailabilityModeZoneRedundant,
	}
}

// HighAvailabilityState - The state of server high availability.
type HighAvailabilityState string

const (
	HighAvailabilityStateCreatingStandby HighAvailabilityState = "CreatingStandby"
	HighAvailabilityStateFailingOver     HighAvailabilityState = "FailingOver"
	HighAvailabilityStateHealthy         HighAvailabilityState = "Healthy"
	HighAvailabilityStateNotEnabled      HighAvailabilityState = "NotEnabled"
	HighAvailabilityStateRemovingStandby HighAvailabilityState = "RemovingStandby"
)

// PossibleHighAvailabilityStateValues returns the possible values for the HighAvailabilityState const type.
func PossibleHighAvailabilityStateValues() []HighAvailabilityState {
	return []HighAvailabilityState{
		HighAvailabilityStateCreatingStandby,
		HighAvailabilityStateFailingOver,
		HighAvailabilityStateHealthy,
		HighAvailabilityStateNotEnabled,
		HighAvailabilityStateRemovingStandby,
	}
}

// ImportSourceStorageType - Storage type of import source.
type ImportSourceStorageType string

const (
	ImportSourceStorageTypeAzureBlob ImportSourceStorageType = "AzureBlob"
)

// PossibleImportSourceStorageTypeValues returns the possible values for the ImportSourceStorageType const type.
func PossibleImportSourceStorageTypeValues() []ImportSourceStorageType {
	return []ImportSourceStorageType{
		ImportSourceStorageTypeAzureBlob,
	}
}

// IsConfigPendingRestart - If is the configuration pending restart or not.
type IsConfigPendingRestart string

const (
	IsConfigPendingRestartFalse IsConfigPendingRestart = "False"
	IsConfigPendingRestartTrue  IsConfigPendingRestart = "True"
)

// PossibleIsConfigPendingRestartValues returns the possible values for the IsConfigPendingRestart const type.
func PossibleIsConfigPendingRestartValues() []IsConfigPendingRestart {
	return []IsConfigPendingRestart{
		IsConfigPendingRestartFalse,
		IsConfigPendingRestartTrue,
	}
}

// IsDynamicConfig - If is the configuration dynamic.
type IsDynamicConfig string

const (
	IsDynamicConfigFalse IsDynamicConfig = "False"
	IsDynamicConfigTrue  IsDynamicConfig = "True"
)

// PossibleIsDynamicConfigValues returns the possible values for the IsDynamicConfig const type.
func PossibleIsDynamicConfigValues() []IsDynamicConfig {
	return []IsDynamicConfig{
		IsDynamicConfigFalse,
		IsDynamicConfigTrue,
	}
}

// IsReadOnly - If is the configuration read only.
type IsReadOnly string

const (
	IsReadOnlyFalse IsReadOnly = "False"
	IsReadOnlyTrue  IsReadOnly = "True"
)

// PossibleIsReadOnlyValues returns the possible values for the IsReadOnly const type.
func PossibleIsReadOnlyValues() []IsReadOnly {
	return []IsReadOnly{
		IsReadOnlyFalse,
		IsReadOnlyTrue,
	}
}

// MaintenanceProvisioningState - The current provisioning state.
type MaintenanceProvisioningState string

const (
	MaintenanceProvisioningStateCreating  MaintenanceProvisioningState = "Creating"
	MaintenanceProvisioningStateDeleting  MaintenanceProvisioningState = "Deleting"
	MaintenanceProvisioningStateFailed    MaintenanceProvisioningState = "Failed"
	MaintenanceProvisioningStateSucceeded MaintenanceProvisioningState = "Succeeded"
)

// PossibleMaintenanceProvisioningStateValues returns the possible values for the MaintenanceProvisioningState const type.
func PossibleMaintenanceProvisioningStateValues() []MaintenanceProvisioningState {
	return []MaintenanceProvisioningState{
		MaintenanceProvisioningStateCreating,
		MaintenanceProvisioningStateDeleting,
		MaintenanceProvisioningStateFailed,
		MaintenanceProvisioningStateSucceeded,
	}
}

// MaintenanceState - The current status of this maintenance.
type MaintenanceState string

const (
	MaintenanceStateCanceled      MaintenanceState = "Canceled"
	MaintenanceStateCompleted     MaintenanceState = "Completed"
	MaintenanceStateInPreparation MaintenanceState = "InPreparation"
	MaintenanceStateProcessing    MaintenanceState = "Processing"
	MaintenanceStateReScheduled   MaintenanceState = "ReScheduled"
	MaintenanceStateScheduled     MaintenanceState = "Scheduled"
)

// PossibleMaintenanceStateValues returns the possible values for the MaintenanceState const type.
func PossibleMaintenanceStateValues() []MaintenanceState {
	return []MaintenanceState{
		MaintenanceStateCanceled,
		MaintenanceStateCompleted,
		MaintenanceStateInPreparation,
		MaintenanceStateProcessing,
		MaintenanceStateReScheduled,
		MaintenanceStateScheduled,
	}
}

// MaintenanceType - The type of this maintenance.
type MaintenanceType string

const (
	MaintenanceTypeHotFixes            MaintenanceType = "HotFixes"
	MaintenanceTypeMinorVersionUpgrade MaintenanceType = "MinorVersionUpgrade"
	MaintenanceTypeRoutineMaintenance  MaintenanceType = "RoutineMaintenance"
	MaintenanceTypeSecurityPatches     MaintenanceType = "SecurityPatches"
)

// PossibleMaintenanceTypeValues returns the possible values for the MaintenanceType const type.
func PossibleMaintenanceTypeValues() []MaintenanceType {
	return []MaintenanceType{
		MaintenanceTypeHotFixes,
		MaintenanceTypeMinorVersionUpgrade,
		MaintenanceTypeRoutineMaintenance,
		MaintenanceTypeSecurityPatches,
	}
}

// ManagedServiceIdentityType - Type of managed service identity.
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeUserAssigned ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// ObjectType - Identifies the type of source operation
type ObjectType string

const (
	ObjectTypeBackupAndExportResponse   ObjectType = "BackupAndExportResponse"
	ObjectTypeImportFromStorageResponse ObjectType = "ImportFromStorageResponse"
)

// PossibleObjectTypeValues returns the possible values for the ObjectType const type.
func PossibleObjectTypeValues() []ObjectType {
	return []ObjectType{
		ObjectTypeBackupAndExportResponse,
		ObjectTypeImportFromStorageResponse,
	}
}

// OperationStatus - The operation status
type OperationStatus string

const (
	// OperationStatusCancelInProgress - The cancellation in progress
	OperationStatusCancelInProgress OperationStatus = "CancelInProgress"
	// OperationStatusCanceled - The operation has been Canceled
	OperationStatusCanceled OperationStatus = "Canceled"
	// OperationStatusFailed - The operation Failed
	OperationStatusFailed OperationStatus = "Failed"
	// OperationStatusInProgress - The operation is running
	OperationStatusInProgress OperationStatus = "InProgress"
	// OperationStatusPending - The operation has been accepted but hasn't started.
	OperationStatusPending OperationStatus = "Pending"
	// OperationStatusSucceeded - The operation Succeeded
	OperationStatusSucceeded OperationStatus = "Succeeded"
)

// PossibleOperationStatusValues returns the possible values for the OperationStatus const type.
func PossibleOperationStatusValues() []OperationStatus {
	return []OperationStatus{
		OperationStatusCancelInProgress,
		OperationStatusCanceled,
		OperationStatusFailed,
		OperationStatusInProgress,
		OperationStatusPending,
		OperationStatusSucceeded,
	}
}

// PatchStrategy - Enum to indicate the patch strategy of a server
type PatchStrategy string

const (
	PatchStrategyRegular       PatchStrategy = "Regular"
	PatchStrategyVirtualCanary PatchStrategy = "VirtualCanary"
)

// PossiblePatchStrategyValues returns the possible values for the PatchStrategy const type.
func PossiblePatchStrategyValues() []PatchStrategy {
	return []PatchStrategy{
		PatchStrategyRegular,
		PatchStrategyVirtualCanary,
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

// ProvisioningState - The current provisioning state.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
	}
}

// ReplicationRole - The replication role.
type ReplicationRole string

const (
	ReplicationRoleNone    ReplicationRole = "None"
	ReplicationRoleReplica ReplicationRole = "Replica"
	ReplicationRoleSource  ReplicationRole = "Source"
)

// PossibleReplicationRoleValues returns the possible values for the ReplicationRole const type.
func PossibleReplicationRoleValues() []ReplicationRole {
	return []ReplicationRole{
		ReplicationRoleNone,
		ReplicationRoleReplica,
		ReplicationRoleSource,
	}
}

// ResetAllToDefault - Whether to reset all server parameters to default.
type ResetAllToDefault string

const (
	ResetAllToDefaultFalse ResetAllToDefault = "False"
	ResetAllToDefaultTrue  ResetAllToDefault = "True"
)

// PossibleResetAllToDefaultValues returns the possible values for the ResetAllToDefault const type.
func PossibleResetAllToDefaultValues() []ResetAllToDefault {
	return []ResetAllToDefault{
		ResetAllToDefaultFalse,
		ResetAllToDefaultTrue,
	}
}

// ServerSKUTier - The tier of the particular SKU, e.g. GeneralPurpose.
type ServerSKUTier string

const (
	ServerSKUTierBurstable       ServerSKUTier = "Burstable"
	ServerSKUTierGeneralPurpose  ServerSKUTier = "GeneralPurpose"
	ServerSKUTierMemoryOptimized ServerSKUTier = "MemoryOptimized"
)

// PossibleServerSKUTierValues returns the possible values for the ServerSKUTier const type.
func PossibleServerSKUTierValues() []ServerSKUTier {
	return []ServerSKUTier{
		ServerSKUTierBurstable,
		ServerSKUTierGeneralPurpose,
		ServerSKUTierMemoryOptimized,
	}
}

// ServerState - The state of a server.
type ServerState string

const (
	ServerStateDisabled ServerState = "Disabled"
	ServerStateDropping ServerState = "Dropping"
	ServerStateReady    ServerState = "Ready"
	ServerStateStarting ServerState = "Starting"
	ServerStateStopped  ServerState = "Stopped"
	ServerStateStopping ServerState = "Stopping"
	ServerStateUpdating ServerState = "Updating"
)

// PossibleServerStateValues returns the possible values for the ServerState const type.
func PossibleServerStateValues() []ServerState {
	return []ServerState{
		ServerStateDisabled,
		ServerStateDropping,
		ServerStateReady,
		ServerStateStarting,
		ServerStateStopped,
		ServerStateStopping,
		ServerStateUpdating,
	}
}

// ServerVersion - The major version of a server. 8.0.21 stands for MySQL 8.0, 5.7.44 stands for MySQL 5.7
type ServerVersion string

const (
	ServerVersionEight021 ServerVersion = "8.0.21"
	ServerVersionFive7    ServerVersion = "5.7"
)

// PossibleServerVersionValues returns the possible values for the ServerVersion const type.
func PossibleServerVersionValues() []ServerVersion {
	return []ServerVersion{
		ServerVersionEight021,
		ServerVersionFive7,
	}
}

// StorageRedundancyEnum - Enum to indicate whether storage sku value is 'ZoneRedundancy' or 'LocalRedundancy'
type StorageRedundancyEnum string

const (
	StorageRedundancyEnumLocalRedundancy StorageRedundancyEnum = "LocalRedundancy"
	StorageRedundancyEnumZoneRedundancy  StorageRedundancyEnum = "ZoneRedundancy"
)

// PossibleStorageRedundancyEnumValues returns the possible values for the StorageRedundancyEnum const type.
func PossibleStorageRedundancyEnumValues() []StorageRedundancyEnum {
	return []StorageRedundancyEnum{
		StorageRedundancyEnumLocalRedundancy,
		StorageRedundancyEnumZoneRedundancy,
	}
}
