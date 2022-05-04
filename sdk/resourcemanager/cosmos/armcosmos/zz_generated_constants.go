//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

const (
	moduleName    = "armcosmos"
	moduleVersion = "v0.6.0"
)

// APIType - Enum to indicate the API type of the restorable database account.
type APIType string

const (
	APITypeCassandra APIType = "Cassandra"
	APITypeGremlin   APIType = "Gremlin"
	APITypeGremlinV2 APIType = "GremlinV2"
	APITypeMongoDB   APIType = "MongoDB"
	APITypeSQL       APIType = "Sql"
	APITypeTable     APIType = "Table"
)

// PossibleAPITypeValues returns the possible values for the APIType const type.
func PossibleAPITypeValues() []APIType {
	return []APIType{
		APITypeCassandra,
		APITypeGremlin,
		APITypeGremlinV2,
		APITypeMongoDB,
		APITypeSQL,
		APITypeTable,
	}
}

// AnalyticalStorageSchemaType - Describes the types of schema for analytical storage.
type AnalyticalStorageSchemaType string

const (
	AnalyticalStorageSchemaTypeFullFidelity AnalyticalStorageSchemaType = "FullFidelity"
	AnalyticalStorageSchemaTypeWellDefined  AnalyticalStorageSchemaType = "WellDefined"
)

// PossibleAnalyticalStorageSchemaTypeValues returns the possible values for the AnalyticalStorageSchemaType const type.
func PossibleAnalyticalStorageSchemaTypeValues() []AnalyticalStorageSchemaType {
	return []AnalyticalStorageSchemaType{
		AnalyticalStorageSchemaTypeFullFidelity,
		AnalyticalStorageSchemaTypeWellDefined,
	}
}

// AuthenticationMethod - Which authentication method Cassandra should use to authenticate clients. 'None' turns off authentication,
// so should not be used except in emergencies. 'Cassandra' is the default password based
// authentication. The default is 'Cassandra'. 'Ldap' is in preview.
type AuthenticationMethod string

const (
	AuthenticationMethodCassandra AuthenticationMethod = "Cassandra"
	AuthenticationMethodLdap      AuthenticationMethod = "Ldap"
	AuthenticationMethodNone      AuthenticationMethod = "None"
)

// PossibleAuthenticationMethodValues returns the possible values for the AuthenticationMethod const type.
func PossibleAuthenticationMethodValues() []AuthenticationMethod {
	return []AuthenticationMethod{
		AuthenticationMethodCassandra,
		AuthenticationMethodLdap,
		AuthenticationMethodNone,
	}
}

// BackupPolicyMigrationStatus - Describes the status of migration between backup policy types.
type BackupPolicyMigrationStatus string

const (
	BackupPolicyMigrationStatusCompleted  BackupPolicyMigrationStatus = "Completed"
	BackupPolicyMigrationStatusFailed     BackupPolicyMigrationStatus = "Failed"
	BackupPolicyMigrationStatusInProgress BackupPolicyMigrationStatus = "InProgress"
	BackupPolicyMigrationStatusInvalid    BackupPolicyMigrationStatus = "Invalid"
)

// PossibleBackupPolicyMigrationStatusValues returns the possible values for the BackupPolicyMigrationStatus const type.
func PossibleBackupPolicyMigrationStatusValues() []BackupPolicyMigrationStatus {
	return []BackupPolicyMigrationStatus{
		BackupPolicyMigrationStatusCompleted,
		BackupPolicyMigrationStatusFailed,
		BackupPolicyMigrationStatusInProgress,
		BackupPolicyMigrationStatusInvalid,
	}
}

// BackupPolicyType - Describes the mode of backups.
type BackupPolicyType string

const (
	BackupPolicyTypeContinuous BackupPolicyType = "Continuous"
	BackupPolicyTypePeriodic   BackupPolicyType = "Periodic"
)

// PossibleBackupPolicyTypeValues returns the possible values for the BackupPolicyType const type.
func PossibleBackupPolicyTypeValues() []BackupPolicyType {
	return []BackupPolicyType{
		BackupPolicyTypeContinuous,
		BackupPolicyTypePeriodic,
	}
}

// BackupStorageRedundancy - Enum to indicate type of backup storage redundancy.
type BackupStorageRedundancy string

const (
	BackupStorageRedundancyGeo   BackupStorageRedundancy = "Geo"
	BackupStorageRedundancyLocal BackupStorageRedundancy = "Local"
	BackupStorageRedundancyZone  BackupStorageRedundancy = "Zone"
)

// PossibleBackupStorageRedundancyValues returns the possible values for the BackupStorageRedundancy const type.
func PossibleBackupStorageRedundancyValues() []BackupStorageRedundancy {
	return []BackupStorageRedundancy{
		BackupStorageRedundancyGeo,
		BackupStorageRedundancyLocal,
		BackupStorageRedundancyZone,
	}
}

// CompositePathSortOrder - Sort order for composite paths.
type CompositePathSortOrder string

const (
	CompositePathSortOrderAscending  CompositePathSortOrder = "ascending"
	CompositePathSortOrderDescending CompositePathSortOrder = "descending"
)

// PossibleCompositePathSortOrderValues returns the possible values for the CompositePathSortOrder const type.
func PossibleCompositePathSortOrderValues() []CompositePathSortOrder {
	return []CompositePathSortOrder{
		CompositePathSortOrderAscending,
		CompositePathSortOrderDescending,
	}
}

// ConflictResolutionMode - Indicates the conflict resolution mode.
type ConflictResolutionMode string

const (
	ConflictResolutionModeCustom         ConflictResolutionMode = "Custom"
	ConflictResolutionModeLastWriterWins ConflictResolutionMode = "LastWriterWins"
)

// PossibleConflictResolutionModeValues returns the possible values for the ConflictResolutionMode const type.
func PossibleConflictResolutionModeValues() []ConflictResolutionMode {
	return []ConflictResolutionMode{
		ConflictResolutionModeCustom,
		ConflictResolutionModeLastWriterWins,
	}
}

// ConnectionState - The kind of connection error that occurred.
type ConnectionState string

const (
	ConnectionStateDatacenterToDatacenterNetworkError           ConnectionState = "DatacenterToDatacenterNetworkError"
	ConnectionStateInternalError                                ConnectionState = "InternalError"
	ConnectionStateInternalOperatorToDataCenterCertificateError ConnectionState = "InternalOperatorToDataCenterCertificateError"
	ConnectionStateOK                                           ConnectionState = "OK"
	ConnectionStateOperatorToDataCenterNetworkError             ConnectionState = "OperatorToDataCenterNetworkError"
	ConnectionStateUnknown                                      ConnectionState = "Unknown"
)

// PossibleConnectionStateValues returns the possible values for the ConnectionState const type.
func PossibleConnectionStateValues() []ConnectionState {
	return []ConnectionState{
		ConnectionStateDatacenterToDatacenterNetworkError,
		ConnectionStateInternalError,
		ConnectionStateInternalOperatorToDataCenterCertificateError,
		ConnectionStateOK,
		ConnectionStateOperatorToDataCenterNetworkError,
		ConnectionStateUnknown,
	}
}

// ConnectorOffer - The cassandra connector offer type for the Cosmos DB C* database account.
type ConnectorOffer string

const (
	ConnectorOfferSmall ConnectorOffer = "Small"
)

// PossibleConnectorOfferValues returns the possible values for the ConnectorOffer const type.
func PossibleConnectorOfferValues() []ConnectorOffer {
	return []ConnectorOffer{
		ConnectorOfferSmall,
	}
}

// ContinuousTier - Enum to indicate type of Continuous backup tier.
type ContinuousTier string

const (
	ContinuousTierContinuous30Days ContinuousTier = "Continuous30Days"
	ContinuousTierContinuous7Days  ContinuousTier = "Continuous7Days"
)

// PossibleContinuousTierValues returns the possible values for the ContinuousTier const type.
func PossibleContinuousTierValues() []ContinuousTier {
	return []ContinuousTier{
		ContinuousTierContinuous30Days,
		ContinuousTierContinuous7Days,
	}
}

// CreateMode - Enum to indicate the mode of account creation.
type CreateMode string

const (
	CreateModeDefault CreateMode = "Default"
	CreateModeRestore CreateMode = "Restore"
)

// PossibleCreateModeValues returns the possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{
		CreateModeDefault,
		CreateModeRestore,
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

type DataTransferComponent string

const (
	DataTransferComponentAzureBlobStorage  DataTransferComponent = "AzureBlobStorage"
	DataTransferComponentCosmosDBCassandra DataTransferComponent = "CosmosDBCassandra"
	DataTransferComponentCosmosDBSQL       DataTransferComponent = "CosmosDBSql"
)

// PossibleDataTransferComponentValues returns the possible values for the DataTransferComponent const type.
func PossibleDataTransferComponentValues() []DataTransferComponent {
	return []DataTransferComponent{
		DataTransferComponentAzureBlobStorage,
		DataTransferComponentCosmosDBCassandra,
		DataTransferComponentCosmosDBSQL,
	}
}

// DataType - The datatype for which the indexing behavior is applied to.
type DataType string

const (
	DataTypeLineString   DataType = "LineString"
	DataTypeMultiPolygon DataType = "MultiPolygon"
	DataTypeNumber       DataType = "Number"
	DataTypePoint        DataType = "Point"
	DataTypePolygon      DataType = "Polygon"
	DataTypeString       DataType = "String"
)

// PossibleDataTypeValues returns the possible values for the DataType const type.
func PossibleDataTypeValues() []DataType {
	return []DataType{
		DataTypeLineString,
		DataTypeMultiPolygon,
		DataTypeNumber,
		DataTypePoint,
		DataTypePolygon,
		DataTypeString,
	}
}

// DatabaseAccountKind - Indicates the type of database account. This can only be set at database account creation.
type DatabaseAccountKind string

const (
	DatabaseAccountKindGlobalDocumentDB DatabaseAccountKind = "GlobalDocumentDB"
	DatabaseAccountKindMongoDB          DatabaseAccountKind = "MongoDB"
	DatabaseAccountKindParse            DatabaseAccountKind = "Parse"
)

// PossibleDatabaseAccountKindValues returns the possible values for the DatabaseAccountKind const type.
func PossibleDatabaseAccountKindValues() []DatabaseAccountKind {
	return []DatabaseAccountKind{
		DatabaseAccountKindGlobalDocumentDB,
		DatabaseAccountKindMongoDB,
		DatabaseAccountKindParse,
	}
}

// DefaultConsistencyLevel - The default consistency level and configuration settings of the Cosmos DB account.
type DefaultConsistencyLevel string

const (
	DefaultConsistencyLevelEventual         DefaultConsistencyLevel = "Eventual"
	DefaultConsistencyLevelSession          DefaultConsistencyLevel = "Session"
	DefaultConsistencyLevelBoundedStaleness DefaultConsistencyLevel = "BoundedStaleness"
	DefaultConsistencyLevelStrong           DefaultConsistencyLevel = "Strong"
	DefaultConsistencyLevelConsistentPrefix DefaultConsistencyLevel = "ConsistentPrefix"
)

// PossibleDefaultConsistencyLevelValues returns the possible values for the DefaultConsistencyLevel const type.
func PossibleDefaultConsistencyLevelValues() []DefaultConsistencyLevel {
	return []DefaultConsistencyLevel{
		DefaultConsistencyLevelEventual,
		DefaultConsistencyLevelSession,
		DefaultConsistencyLevelBoundedStaleness,
		DefaultConsistencyLevelStrong,
		DefaultConsistencyLevelConsistentPrefix,
	}
}

// EnableFullTextQuery - Describe the level of detail with which queries are to be logged.
type EnableFullTextQuery string

const (
	EnableFullTextQueryNone  EnableFullTextQuery = "None"
	EnableFullTextQueryTrue  EnableFullTextQuery = "True"
	EnableFullTextQueryFalse EnableFullTextQuery = "False"
)

// PossibleEnableFullTextQueryValues returns the possible values for the EnableFullTextQuery const type.
func PossibleEnableFullTextQueryValues() []EnableFullTextQuery {
	return []EnableFullTextQuery{
		EnableFullTextQueryNone,
		EnableFullTextQueryTrue,
		EnableFullTextQueryFalse,
	}
}

// IndexKind - Indicates the type of index.
type IndexKind string

const (
	IndexKindHash    IndexKind = "Hash"
	IndexKindRange   IndexKind = "Range"
	IndexKindSpatial IndexKind = "Spatial"
)

// PossibleIndexKindValues returns the possible values for the IndexKind const type.
func PossibleIndexKindValues() []IndexKind {
	return []IndexKind{
		IndexKindHash,
		IndexKindRange,
		IndexKindSpatial,
	}
}

// IndexingMode - Indicates the indexing mode.
type IndexingMode string

const (
	IndexingModeConsistent IndexingMode = "consistent"
	IndexingModeLazy       IndexingMode = "lazy"
	IndexingModeNone       IndexingMode = "none"
)

// PossibleIndexingModeValues returns the possible values for the IndexingMode const type.
func PossibleIndexingModeValues() []IndexingMode {
	return []IndexingMode{
		IndexingModeConsistent,
		IndexingModeLazy,
		IndexingModeNone,
	}
}

// KeyKind - The access key to regenerate.
type KeyKind string

const (
	KeyKindPrimary           KeyKind = "primary"
	KeyKindPrimaryReadonly   KeyKind = "primaryReadonly"
	KeyKindSecondary         KeyKind = "secondary"
	KeyKindSecondaryReadonly KeyKind = "secondaryReadonly"
)

// PossibleKeyKindValues returns the possible values for the KeyKind const type.
func PossibleKeyKindValues() []KeyKind {
	return []KeyKind{
		KeyKindPrimary,
		KeyKindPrimaryReadonly,
		KeyKindSecondary,
		KeyKindSecondaryReadonly,
	}
}

// ManagedCassandraProvisioningState - The status of the resource at the time the operation was called.
type ManagedCassandraProvisioningState string

const (
	ManagedCassandraProvisioningStateCanceled  ManagedCassandraProvisioningState = "Canceled"
	ManagedCassandraProvisioningStateCreating  ManagedCassandraProvisioningState = "Creating"
	ManagedCassandraProvisioningStateDeleting  ManagedCassandraProvisioningState = "Deleting"
	ManagedCassandraProvisioningStateFailed    ManagedCassandraProvisioningState = "Failed"
	ManagedCassandraProvisioningStateSucceeded ManagedCassandraProvisioningState = "Succeeded"
	ManagedCassandraProvisioningStateUpdating  ManagedCassandraProvisioningState = "Updating"
)

// PossibleManagedCassandraProvisioningStateValues returns the possible values for the ManagedCassandraProvisioningState const type.
func PossibleManagedCassandraProvisioningStateValues() []ManagedCassandraProvisioningState {
	return []ManagedCassandraProvisioningState{
		ManagedCassandraProvisioningStateCanceled,
		ManagedCassandraProvisioningStateCreating,
		ManagedCassandraProvisioningStateDeleting,
		ManagedCassandraProvisioningStateFailed,
		ManagedCassandraProvisioningStateSucceeded,
		ManagedCassandraProvisioningStateUpdating,
	}
}

// ManagedCassandraResourceIdentityType - The type of the resource.
type ManagedCassandraResourceIdentityType string

const (
	ManagedCassandraResourceIdentityTypeNone           ManagedCassandraResourceIdentityType = "None"
	ManagedCassandraResourceIdentityTypeSystemAssigned ManagedCassandraResourceIdentityType = "SystemAssigned"
)

// PossibleManagedCassandraResourceIdentityTypeValues returns the possible values for the ManagedCassandraResourceIdentityType const type.
func PossibleManagedCassandraResourceIdentityTypeValues() []ManagedCassandraResourceIdentityType {
	return []ManagedCassandraResourceIdentityType{
		ManagedCassandraResourceIdentityTypeNone,
		ManagedCassandraResourceIdentityTypeSystemAssigned,
	}
}

// MongoRoleDefinitionType - Indicates whether the Role Definition was built-in or user created.
type MongoRoleDefinitionType string

const (
	MongoRoleDefinitionTypeBuiltInRole MongoRoleDefinitionType = "BuiltInRole"
	MongoRoleDefinitionTypeCustomRole  MongoRoleDefinitionType = "CustomRole"
)

// PossibleMongoRoleDefinitionTypeValues returns the possible values for the MongoRoleDefinitionType const type.
func PossibleMongoRoleDefinitionTypeValues() []MongoRoleDefinitionType {
	return []MongoRoleDefinitionType{
		MongoRoleDefinitionTypeBuiltInRole,
		MongoRoleDefinitionTypeCustomRole,
	}
}

// NetworkACLBypass - Indicates what services are allowed to bypass firewall checks.
type NetworkACLBypass string

const (
	NetworkACLBypassNone          NetworkACLBypass = "None"
	NetworkACLBypassAzureServices NetworkACLBypass = "AzureServices"
)

// PossibleNetworkACLBypassValues returns the possible values for the NetworkACLBypass const type.
func PossibleNetworkACLBypassValues() []NetworkACLBypass {
	return []NetworkACLBypass{
		NetworkACLBypassNone,
		NetworkACLBypassAzureServices,
	}
}

// NodeState - The state of the node in Cassandra ring.
type NodeState string

const (
	NodeStateJoining NodeState = "Joining"
	NodeStateLeaving NodeState = "Leaving"
	NodeStateMoving  NodeState = "Moving"
	NodeStateNormal  NodeState = "Normal"
	NodeStateStopped NodeState = "Stopped"
)

// PossibleNodeStateValues returns the possible values for the NodeState const type.
func PossibleNodeStateValues() []NodeState {
	return []NodeState{
		NodeStateJoining,
		NodeStateLeaving,
		NodeStateMoving,
		NodeStateNormal,
		NodeStateStopped,
	}
}

// NodeStatus - Indicates whether the node is functioning or not.
type NodeStatus string

const (
	NodeStatusDown NodeStatus = "Down"
	NodeStatusUp   NodeStatus = "Up"
)

// PossibleNodeStatusValues returns the possible values for the NodeStatus const type.
func PossibleNodeStatusValues() []NodeStatus {
	return []NodeStatus{
		NodeStatusDown,
		NodeStatusUp,
	}
}

type NotebookWorkspaceName string

const (
	NotebookWorkspaceNameDefault NotebookWorkspaceName = "default"
)

// PossibleNotebookWorkspaceNameValues returns the possible values for the NotebookWorkspaceName const type.
func PossibleNotebookWorkspaceNameValues() []NotebookWorkspaceName {
	return []NotebookWorkspaceName{
		NotebookWorkspaceNameDefault,
	}
}

// OperationType - Enum to indicate the operation type of the event.
type OperationType string

const (
	OperationTypeCreate          OperationType = "Create"
	OperationTypeDelete          OperationType = "Delete"
	OperationTypeReplace         OperationType = "Replace"
	OperationTypeSystemOperation OperationType = "SystemOperation"
)

// PossibleOperationTypeValues returns the possible values for the OperationType const type.
func PossibleOperationTypeValues() []OperationType {
	return []OperationType{
		OperationTypeCreate,
		OperationTypeDelete,
		OperationTypeReplace,
		OperationTypeSystemOperation,
	}
}

// PartitionKind - Indicates the kind of algorithm used for partitioning. For MultiHash, multiple partition keys (upto three
// maximum) are supported for container create
type PartitionKind string

const (
	PartitionKindHash      PartitionKind = "Hash"
	PartitionKindMultiHash PartitionKind = "MultiHash"
	PartitionKindRange     PartitionKind = "Range"
)

// PossiblePartitionKindValues returns the possible values for the PartitionKind const type.
func PossiblePartitionKindValues() []PartitionKind {
	return []PartitionKind{
		PartitionKindHash,
		PartitionKindMultiHash,
		PartitionKindRange,
	}
}

// PrimaryAggregationType - The primary aggregation type of the metric.
type PrimaryAggregationType string

const (
	PrimaryAggregationTypeAverage PrimaryAggregationType = "Average"
	PrimaryAggregationTypeLast    PrimaryAggregationType = "Last"
	PrimaryAggregationTypeMaximum PrimaryAggregationType = "Maximum"
	PrimaryAggregationTypeMinimum PrimaryAggregationType = "Minimum"
	PrimaryAggregationTypeNone    PrimaryAggregationType = "None"
	PrimaryAggregationTypeTotal   PrimaryAggregationType = "Total"
)

// PossiblePrimaryAggregationTypeValues returns the possible values for the PrimaryAggregationType const type.
func PossiblePrimaryAggregationTypeValues() []PrimaryAggregationType {
	return []PrimaryAggregationType{
		PrimaryAggregationTypeAverage,
		PrimaryAggregationTypeLast,
		PrimaryAggregationTypeMaximum,
		PrimaryAggregationTypeMinimum,
		PrimaryAggregationTypeNone,
		PrimaryAggregationTypeTotal,
	}
}

// PublicNetworkAccess - Whether requests from Public Network are allowed
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// ResourceIdentityType - The type of identity used for the resource. The type 'SystemAssigned,UserAssigned' includes both
// an implicitly created identity and a set of user assigned identities. The type 'None' will remove any
// identities from the service.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = "UserAssigned"
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned,UserAssigned"
	ResourceIdentityTypeNone                       ResourceIdentityType = "None"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeUserAssigned,
		ResourceIdentityTypeSystemAssignedUserAssigned,
		ResourceIdentityTypeNone,
	}
}

// RestoreMode - Describes the mode of the restore.
type RestoreMode string

const (
	RestoreModePointInTime RestoreMode = "PointInTime"
)

// PossibleRestoreModeValues returns the possible values for the RestoreMode const type.
func PossibleRestoreModeValues() []RestoreMode {
	return []RestoreMode{
		RestoreModePointInTime,
	}
}

// RoleDefinitionType - Indicates whether the Role Definition was built-in or user created.
type RoleDefinitionType string

const (
	RoleDefinitionTypeBuiltInRole RoleDefinitionType = "BuiltInRole"
	RoleDefinitionTypeCustomRole  RoleDefinitionType = "CustomRole"
)

// PossibleRoleDefinitionTypeValues returns the possible values for the RoleDefinitionType const type.
func PossibleRoleDefinitionTypeValues() []RoleDefinitionType {
	return []RoleDefinitionType{
		RoleDefinitionTypeBuiltInRole,
		RoleDefinitionTypeCustomRole,
	}
}

// ServerVersion - Describes the ServerVersion of an a MongoDB account.
type ServerVersion string

const (
	ServerVersionFour0  ServerVersion = "4.0"
	ServerVersionFour2  ServerVersion = "4.2"
	ServerVersionThree2 ServerVersion = "3.2"
	ServerVersionThree6 ServerVersion = "3.6"
)

// PossibleServerVersionValues returns the possible values for the ServerVersion const type.
func PossibleServerVersionValues() []ServerVersion {
	return []ServerVersion{
		ServerVersionFour0,
		ServerVersionFour2,
		ServerVersionThree2,
		ServerVersionThree6,
	}
}

// ServiceSize - Instance type for the service.
type ServiceSize string

const (
	ServiceSizeCosmosD16S ServiceSize = "Cosmos.D16s"
	ServiceSizeCosmosD4S  ServiceSize = "Cosmos.D4s"
	ServiceSizeCosmosD8S  ServiceSize = "Cosmos.D8s"
)

// PossibleServiceSizeValues returns the possible values for the ServiceSize const type.
func PossibleServiceSizeValues() []ServiceSize {
	return []ServiceSize{
		ServiceSizeCosmosD16S,
		ServiceSizeCosmosD4S,
		ServiceSizeCosmosD8S,
	}
}

// ServiceStatus - Describes the status of a service.
type ServiceStatus string

const (
	ServiceStatusCreating ServiceStatus = "Creating"
	ServiceStatusDeleting ServiceStatus = "Deleting"
	ServiceStatusError    ServiceStatus = "Error"
	ServiceStatusRunning  ServiceStatus = "Running"
	ServiceStatusStopped  ServiceStatus = "Stopped"
	ServiceStatusUpdating ServiceStatus = "Updating"
)

// PossibleServiceStatusValues returns the possible values for the ServiceStatus const type.
func PossibleServiceStatusValues() []ServiceStatus {
	return []ServiceStatus{
		ServiceStatusCreating,
		ServiceStatusDeleting,
		ServiceStatusError,
		ServiceStatusRunning,
		ServiceStatusStopped,
		ServiceStatusUpdating,
	}
}

// ServiceType - ServiceType for the service.
type ServiceType string

const (
	ServiceTypeDataTransfer             ServiceType = "DataTransfer"
	ServiceTypeGraphAPICompute          ServiceType = "GraphAPICompute"
	ServiceTypeMaterializedViewsBuilder ServiceType = "MaterializedViewsBuilder"
	ServiceTypeSQLDedicatedGateway      ServiceType = "SqlDedicatedGateway"
)

// PossibleServiceTypeValues returns the possible values for the ServiceType const type.
func PossibleServiceTypeValues() []ServiceType {
	return []ServiceType{
		ServiceTypeDataTransfer,
		ServiceTypeGraphAPICompute,
		ServiceTypeMaterializedViewsBuilder,
		ServiceTypeSQLDedicatedGateway,
	}
}

// SpatialType - Indicates the spatial type of index.
type SpatialType string

const (
	SpatialTypeLineString   SpatialType = "LineString"
	SpatialTypeMultiPolygon SpatialType = "MultiPolygon"
	SpatialTypePoint        SpatialType = "Point"
	SpatialTypePolygon      SpatialType = "Polygon"
)

// PossibleSpatialTypeValues returns the possible values for the SpatialType const type.
func PossibleSpatialTypeValues() []SpatialType {
	return []SpatialType{
		SpatialTypeLineString,
		SpatialTypeMultiPolygon,
		SpatialTypePoint,
		SpatialTypePolygon,
	}
}

// ThroughputPolicyType - ThroughputPolicy to apply for throughput redistribution
type ThroughputPolicyType string

const (
	ThroughputPolicyTypeCustom ThroughputPolicyType = "custom"
	ThroughputPolicyTypeEqual  ThroughputPolicyType = "equal"
	ThroughputPolicyTypeNone   ThroughputPolicyType = "none"
)

// PossibleThroughputPolicyTypeValues returns the possible values for the ThroughputPolicyType const type.
func PossibleThroughputPolicyTypeValues() []ThroughputPolicyType {
	return []ThroughputPolicyType{
		ThroughputPolicyTypeCustom,
		ThroughputPolicyTypeEqual,
		ThroughputPolicyTypeNone,
	}
}

// TriggerOperation - The operation the trigger is associated with
type TriggerOperation string

const (
	TriggerOperationAll     TriggerOperation = "All"
	TriggerOperationCreate  TriggerOperation = "Create"
	TriggerOperationDelete  TriggerOperation = "Delete"
	TriggerOperationReplace TriggerOperation = "Replace"
	TriggerOperationUpdate  TriggerOperation = "Update"
)

// PossibleTriggerOperationValues returns the possible values for the TriggerOperation const type.
func PossibleTriggerOperationValues() []TriggerOperation {
	return []TriggerOperation{
		TriggerOperationAll,
		TriggerOperationCreate,
		TriggerOperationDelete,
		TriggerOperationReplace,
		TriggerOperationUpdate,
	}
}

// TriggerType - Type of the Trigger
type TriggerType string

const (
	TriggerTypePost TriggerType = "Post"
	TriggerTypePre  TriggerType = "Pre"
)

// PossibleTriggerTypeValues returns the possible values for the TriggerType const type.
func PossibleTriggerTypeValues() []TriggerType {
	return []TriggerType{
		TriggerTypePost,
		TriggerTypePre,
	}
}

// UnitType - The unit of the metric.
type UnitType string

const (
	UnitTypeBytes          UnitType = "Bytes"
	UnitTypeBytesPerSecond UnitType = "BytesPerSecond"
	UnitTypeCount          UnitType = "Count"
	UnitTypeCountPerSecond UnitType = "CountPerSecond"
	UnitTypeMilliseconds   UnitType = "Milliseconds"
	UnitTypePercent        UnitType = "Percent"
	UnitTypeSeconds        UnitType = "Seconds"
)

// PossibleUnitTypeValues returns the possible values for the UnitType const type.
func PossibleUnitTypeValues() []UnitType {
	return []UnitType{
		UnitTypeBytes,
		UnitTypeBytesPerSecond,
		UnitTypeCount,
		UnitTypeCountPerSecond,
		UnitTypeMilliseconds,
		UnitTypePercent,
		UnitTypeSeconds,
	}
}
