//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpostgresqlflexibleservers

const (
	moduleName    = "armpostgresqlflexibleservers"
	moduleVersion = "v2.0.0"
)

// ActiveDirectoryAuthEnum - If Enabled, Azure Active Directory authentication is enabled.
type ActiveDirectoryAuthEnum string

const (
	ActiveDirectoryAuthEnumDisabled ActiveDirectoryAuthEnum = "Disabled"
	ActiveDirectoryAuthEnumEnabled  ActiveDirectoryAuthEnum = "Enabled"
)

// PossibleActiveDirectoryAuthEnumValues returns the possible values for the ActiveDirectoryAuthEnum const type.
func PossibleActiveDirectoryAuthEnumValues() []ActiveDirectoryAuthEnum {
	return []ActiveDirectoryAuthEnum{
		ActiveDirectoryAuthEnumDisabled,
		ActiveDirectoryAuthEnumEnabled,
	}
}

// ArmServerKeyType - Data encryption type to depict if it is System assigned vs Azure Key vault.
type ArmServerKeyType string

const (
	ArmServerKeyTypeAzureKeyVault  ArmServerKeyType = "AzureKeyVault"
	ArmServerKeyTypeSystemAssigned ArmServerKeyType = "SystemAssigned"
)

// PossibleArmServerKeyTypeValues returns the possible values for the ArmServerKeyType const type.
func PossibleArmServerKeyTypeValues() []ArmServerKeyType {
	return []ArmServerKeyType{
		ArmServerKeyTypeAzureKeyVault,
		ArmServerKeyTypeSystemAssigned,
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

// ConfigurationDataType - Data type of the configuration.
type ConfigurationDataType string

const (
	ConfigurationDataTypeBoolean     ConfigurationDataType = "Boolean"
	ConfigurationDataTypeEnumeration ConfigurationDataType = "Enumeration"
	ConfigurationDataTypeInteger     ConfigurationDataType = "Integer"
	ConfigurationDataTypeNumeric     ConfigurationDataType = "Numeric"
)

// PossibleConfigurationDataTypeValues returns the possible values for the ConfigurationDataType const type.
func PossibleConfigurationDataTypeValues() []ConfigurationDataType {
	return []ConfigurationDataType{
		ConfigurationDataTypeBoolean,
		ConfigurationDataTypeEnumeration,
		ConfigurationDataTypeInteger,
		ConfigurationDataTypeNumeric,
	}
}

// CreateMode - The mode to create a new PostgreSQL server.
type CreateMode string

const (
	CreateModeCreate             CreateMode = "Create"
	CreateModeDefault            CreateMode = "Default"
	CreateModeGeoRestore         CreateMode = "GeoRestore"
	CreateModePointInTimeRestore CreateMode = "PointInTimeRestore"
	CreateModeReplica            CreateMode = "Replica"
	CreateModeUpdate             CreateMode = "Update"
)

// PossibleCreateModeValues returns the possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{
		CreateModeCreate,
		CreateModeDefault,
		CreateModeGeoRestore,
		CreateModePointInTimeRestore,
		CreateModeReplica,
		CreateModeUpdate,
	}
}

// CreateModeForUpdate - The mode to update a new PostgreSQL server.
type CreateModeForUpdate string

const (
	CreateModeForUpdateDefault CreateModeForUpdate = "Default"
	CreateModeForUpdateUpdate  CreateModeForUpdate = "Update"
)

// PossibleCreateModeForUpdateValues returns the possible values for the CreateModeForUpdate const type.
func PossibleCreateModeForUpdateValues() []CreateModeForUpdate {
	return []CreateModeForUpdate{
		CreateModeForUpdateDefault,
		CreateModeForUpdateUpdate,
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

// FailoverMode - Failover mode.
type FailoverMode string

const (
	FailoverModeForcedFailover    FailoverMode = "ForcedFailover"
	FailoverModeForcedSwitchover  FailoverMode = "ForcedSwitchover"
	FailoverModePlannedFailover   FailoverMode = "PlannedFailover"
	FailoverModePlannedSwitchover FailoverMode = "PlannedSwitchover"
)

// PossibleFailoverModeValues returns the possible values for the FailoverMode const type.
func PossibleFailoverModeValues() []FailoverMode {
	return []FailoverMode{
		FailoverModeForcedFailover,
		FailoverModeForcedSwitchover,
		FailoverModePlannedFailover,
		FailoverModePlannedSwitchover,
	}
}

// GeoRedundantBackupEnum - A value indicating whether Geo-Redundant backup is enabled on the server.
type GeoRedundantBackupEnum string

const (
	GeoRedundantBackupEnumDisabled GeoRedundantBackupEnum = "Disabled"
	GeoRedundantBackupEnumEnabled  GeoRedundantBackupEnum = "Enabled"
)

// PossibleGeoRedundantBackupEnumValues returns the possible values for the GeoRedundantBackupEnum const type.
func PossibleGeoRedundantBackupEnumValues() []GeoRedundantBackupEnum {
	return []GeoRedundantBackupEnum{
		GeoRedundantBackupEnumDisabled,
		GeoRedundantBackupEnumEnabled,
	}
}

// HighAvailabilityMode - The HA mode for the server.
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

// IdentityType - the types of identities associated with this resource; currently restricted to 'SystemAssigned and UserAssigned'
type IdentityType string

const (
	IdentityTypeNone           IdentityType = "None"
	IdentityTypeSystemAssigned IdentityType = "SystemAssigned"
	IdentityTypeUserAssigned   IdentityType = "UserAssigned"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeNone,
		IdentityTypeSystemAssigned,
		IdentityTypeUserAssigned,
	}
}

// OperationOrigin - The intended executor of the operation.
type OperationOrigin string

const (
	OperationOriginNotSpecified OperationOrigin = "NotSpecified"
	OperationOriginSystem       OperationOrigin = "system"
	OperationOriginUser         OperationOrigin = "user"
)

// PossibleOperationOriginValues returns the possible values for the OperationOrigin const type.
func PossibleOperationOriginValues() []OperationOrigin {
	return []OperationOrigin{
		OperationOriginNotSpecified,
		OperationOriginSystem,
		OperationOriginUser,
	}
}

// Origin - Backup type.
type Origin string

const (
	OriginFull Origin = "Full"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginFull,
	}
}

// PasswordAuthEnum - If Enabled, Password authentication is enabled.
type PasswordAuthEnum string

const (
	PasswordAuthEnumDisabled PasswordAuthEnum = "Disabled"
	PasswordAuthEnumEnabled  PasswordAuthEnum = "Enabled"
)

// PossiblePasswordAuthEnumValues returns the possible values for the PasswordAuthEnum const type.
func PossiblePasswordAuthEnumValues() []PasswordAuthEnum {
	return []PasswordAuthEnum{
		PasswordAuthEnumDisabled,
		PasswordAuthEnumEnabled,
	}
}

// PrincipalType - The principal type used to represent the type of Active Directory Administrator.
type PrincipalType string

const (
	PrincipalTypeGroup            PrincipalType = "Group"
	PrincipalTypeServicePrincipal PrincipalType = "ServicePrincipal"
	PrincipalTypeUnknown          PrincipalType = "Unknown"
	PrincipalTypeUser             PrincipalType = "User"
)

// PossiblePrincipalTypeValues returns the possible values for the PrincipalType const type.
func PossiblePrincipalTypeValues() []PrincipalType {
	return []PrincipalType{
		PrincipalTypeGroup,
		PrincipalTypeServicePrincipal,
		PrincipalTypeUnknown,
		PrincipalTypeUser,
	}
}

// ReplicationRole - Used to indicate role of the server in replication set.
type ReplicationRole string

const (
	ReplicationRoleAsyncReplica    ReplicationRole = "AsyncReplica"
	ReplicationRoleGeoAsyncReplica ReplicationRole = "GeoAsyncReplica"
	ReplicationRoleGeoSyncReplica  ReplicationRole = "GeoSyncReplica"
	ReplicationRoleNone            ReplicationRole = "None"
	ReplicationRolePrimary         ReplicationRole = "Primary"
	ReplicationRoleSecondary       ReplicationRole = "Secondary"
	ReplicationRoleSyncReplica     ReplicationRole = "SyncReplica"
	ReplicationRoleWalReplica      ReplicationRole = "WalReplica"
)

// PossibleReplicationRoleValues returns the possible values for the ReplicationRole const type.
func PossibleReplicationRoleValues() []ReplicationRole {
	return []ReplicationRole{
		ReplicationRoleAsyncReplica,
		ReplicationRoleGeoAsyncReplica,
		ReplicationRoleGeoSyncReplica,
		ReplicationRoleNone,
		ReplicationRolePrimary,
		ReplicationRoleSecondary,
		ReplicationRoleSyncReplica,
		ReplicationRoleWalReplica,
	}
}

// SKUTier - The tier of the particular SKU, e.g. Burstable.
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

// ServerHAState - A state of a HA server that is visible to user.
type ServerHAState string

const (
	ServerHAStateCreatingStandby ServerHAState = "CreatingStandby"
	ServerHAStateFailingOver     ServerHAState = "FailingOver"
	ServerHAStateHealthy         ServerHAState = "Healthy"
	ServerHAStateNotEnabled      ServerHAState = "NotEnabled"
	ServerHAStateRemovingStandby ServerHAState = "RemovingStandby"
	ServerHAStateReplicatingData ServerHAState = "ReplicatingData"
)

// PossibleServerHAStateValues returns the possible values for the ServerHAState const type.
func PossibleServerHAStateValues() []ServerHAState {
	return []ServerHAState{
		ServerHAStateCreatingStandby,
		ServerHAStateFailingOver,
		ServerHAStateHealthy,
		ServerHAStateNotEnabled,
		ServerHAStateRemovingStandby,
		ServerHAStateReplicatingData,
	}
}

// ServerPublicNetworkAccessState - public network access is enabled or not
type ServerPublicNetworkAccessState string

const (
	ServerPublicNetworkAccessStateDisabled ServerPublicNetworkAccessState = "Disabled"
	ServerPublicNetworkAccessStateEnabled  ServerPublicNetworkAccessState = "Enabled"
)

// PossibleServerPublicNetworkAccessStateValues returns the possible values for the ServerPublicNetworkAccessState const type.
func PossibleServerPublicNetworkAccessStateValues() []ServerPublicNetworkAccessState {
	return []ServerPublicNetworkAccessState{
		ServerPublicNetworkAccessStateDisabled,
		ServerPublicNetworkAccessStateEnabled,
	}
}

// ServerState - A state of a server that is visible to user.
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
	ServerVersionEleven   ServerVersion = "11"
	ServerVersionFourteen ServerVersion = "14"
	ServerVersionThirteen ServerVersion = "13"
	ServerVersionTwelve   ServerVersion = "12"
)

// PossibleServerVersionValues returns the possible values for the ServerVersion const type.
func PossibleServerVersionValues() []ServerVersion {
	return []ServerVersion{
		ServerVersionEleven,
		ServerVersionFourteen,
		ServerVersionThirteen,
		ServerVersionTwelve,
	}
}
