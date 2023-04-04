//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmysqlflexibleservers

const (
	moduleName    = "armmysqlflexibleservers"
	moduleVersion = "v1.2.0-beta.1"
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

// SKUTier - The tier of the particular SKU, e.g. GeneralPurpose.
type SKUTier string

const (
	SKUTierBurstable       SKUTier = "Burstable"
	SKUTierGeneralPurpose  SKUTier = "GeneralPurpose"
	SKUTierMemoryOptimized SKUTier = "MemoryOptimized"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBurstable,
		SKUTierGeneralPurpose,
		SKUTierMemoryOptimized,
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

// ServerVersion - The version of a server.
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
