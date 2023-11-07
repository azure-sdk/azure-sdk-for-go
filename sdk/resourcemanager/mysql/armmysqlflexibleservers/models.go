//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysqlflexibleservers

import "time"

// AdministratorListResult - A List of azure ad administrators.
type AdministratorListResult struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of azure ad administrator of a server
	Value []*AzureADAdministrator
}

// AdministratorProperties - The properties of an administrator.
type AdministratorProperties struct {
	// Type of the sever administrator.
	AdministratorType *AdministratorType

	// The resource id of the identity used for AAD Authentication.
	IdentityResourceID *string

	// Login name of the server administrator.
	Login *string

	// SID (object ID) of the server administrator.
	Sid *string

	// Tenant ID of the administrator.
	TenantID *string
}

// AdvancedThreatProtection - A server's Advanced Threat Protection.
type AdvancedThreatProtection struct {
	// Resource properties.
	Properties *AdvancedThreatProtectionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// AdvancedThreatProtectionForUpdate - Parameters allowed to update advanced threat protection for a server.
type AdvancedThreatProtectionForUpdate struct {
	// Resource update properties.
	Properties *AdvancedThreatProtectionUpdateProperties
}

// AdvancedThreatProtectionListResult - A list of the server's Advanced Threat Protection configurations.
type AdvancedThreatProtectionListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; Array of results.
	Value []*AdvancedThreatProtection
}

// AdvancedThreatProtectionProperties - Properties of an Advanced Threat Protection setting.
type AdvancedThreatProtectionProperties struct {
	// Specifies the state of the Advanced Threat Protection, whether it is enabled or disabled or a state has not been applied
	// yet on the specific database or server.
	State *AdvancedThreatProtectionState

	// READ-ONLY; Specifies the UTC creation time of the policy.
	CreationTime *time.Time

	// READ-ONLY; Provisioning state of the Threat Protection.
	ProvisioningState *AdvancedThreatProtectionProvisioningState
}

// AdvancedThreatProtectionUpdateProperties - Properties of Advanced Threat Protection that can be updated.
type AdvancedThreatProtectionUpdateProperties struct {
	// REQUIRED; Specifies the state of the Advanced Threat Protection, whether it is enabled or disabled or a state has not been
	// applied yet on the specific database or server.
	State *AdvancedThreatProtectionState
}

// AzureADAdministrator - Represents a Administrator.
type AzureADAdministrator struct {
	// The properties of an administrator.
	Properties *AdministratorProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// Backup - Storage Profile properties of a server
type Backup struct {
	// Backup retention days for the server.
	BackupRetentionDays *int32

	// Whether or not geo redundant backup is enabled.
	GeoRedundantBackup *EnableStatusEnum

	// READ-ONLY; Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate *time.Time
}

// BackupAndExportRequest - BackupAndExport API Request
type BackupAndExportRequest struct {
	// REQUIRED; Backup Settings
	BackupSettings *BackupSettings

	// REQUIRED; Backup Target Store Details
	TargetDetails BackupStoreDetailsClassification
}

// BackupAndExportResponse - Represents BackupAndExport API Response
type BackupAndExportResponse struct {
	// End time
	EndTime *time.Time

	// The error object.
	Error *ErrorDetail

	// Operation progress (0-100).
	PercentComplete *float64

	// The response properties of a backup and export operation.
	Properties *BackupAndExportResponseProperties

	// Start time
	StartTime *time.Time

	// The operation status
	Status *OperationStatus

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// BackupAndExportResponseProperties - BackupAndExport Response Properties
type BackupAndExportResponseProperties struct {
	// Metadata related to backup to be stored for restoring resource in key-value pairs.
	BackupMetadata *string

	// Data transferred in bytes
	DataTransferredInBytes *int64

	// Size of datasource in bytes
	DatasourceSizeInBytes *int64
}

// BackupRequestBase is the base for all backup request.
type BackupRequestBase struct {
	// REQUIRED; Backup Settings
	BackupSettings *BackupSettings
}

// BackupSettings - Backup Settings
type BackupSettings struct {
	// REQUIRED; The name of the backup.
	BackupName *string

	// Backup Format for the current backup. (CollatedFormat is INTERNAL – DO NOT USE)
	BackupFormat *BackupFormat
}

// BackupStoreDetails - Details about the target where the backup content will be stored.
type BackupStoreDetails struct {
	// REQUIRED; Type of the specific object - used for deserializing
	ObjectType *string
}

// GetBackupStoreDetails implements the BackupStoreDetailsClassification interface for type BackupStoreDetails.
func (b *BackupStoreDetails) GetBackupStoreDetails() *BackupStoreDetails { return b }

// CapabilitiesListResult - location capability
type CapabilitiesListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; A list of supported capabilities.
	Value []*CapabilityProperties
}

// Capability - Represents a location capability set.
type Capability struct {
	// The properties of a location capability set.
	Properties *CapabilityPropertiesV2

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// CapabilityProperties - Location capabilities.
type CapabilityProperties struct {
	// READ-ONLY; A list of supported flexible server editions.
	SupportedFlexibleServerEditions []*ServerEditionCapability

	// READ-ONLY; supported geo backup regions
	SupportedGeoBackupRegions []*string

	// READ-ONLY; Supported high availability mode
	SupportedHAMode []*string

	// READ-ONLY; zone name
	Zone *string
}

// CapabilityPropertiesV2 - Location capability.
type CapabilityPropertiesV2 struct {
	// READ-ONLY; A list of supported flexible server editions.
	SupportedFlexibleServerEditions []*ServerEditionCapabilityV2

	// READ-ONLY; supported geo backup regions
	SupportedGeoBackupRegions []*string

	// READ-ONLY; A list of supported server versions.
	SupportedServerVersions []*ServerVersionCapabilityV2
}

// CapabilitySetsList - location capability set
type CapabilitySetsList struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; A list of supported capability sets.
	Value []*Capability
}

// Configuration - Represents a Configuration.
type Configuration struct {
	// The properties of a configuration.
	Properties *ConfigurationProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ConfigurationForBatchUpdate - Represents a Configuration.
type ConfigurationForBatchUpdate struct {
	// Name of the configuration.
	Name *string

	// The properties can be updated for a configuration.
	Properties *ConfigurationForBatchUpdateProperties
}

// ConfigurationForBatchUpdateProperties - The properties can be updated for a configuration.
type ConfigurationForBatchUpdateProperties struct {
	// Source of the configuration.
	Source *string

	// Value of the configuration.
	Value *string
}

// ConfigurationListForBatchUpdate - A list of server configurations to update.
type ConfigurationListForBatchUpdate struct {
	// Whether to reset all server parameters to default.
	ResetAllToDefault *ResetAllToDefault

	// The list of server configurations.
	Value []*ConfigurationForBatchUpdate
}

// ConfigurationListResult - A list of server configurations.
type ConfigurationListResult struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of server configurations.
	Value []*Configuration
}

// ConfigurationProperties - The properties of a configuration.
type ConfigurationProperties struct {
	// Current value of the configuration.
	CurrentValue *string

	// Source of the configuration.
	Source *ConfigurationSource

	// Value of the configuration.
	Value *string

	// READ-ONLY; Allowed values of the configuration.
	AllowedValues *string

	// READ-ONLY; Data type of the configuration.
	DataType *string

	// READ-ONLY; Default value of the configuration.
	DefaultValue *string

	// READ-ONLY; Description of the configuration.
	Description *string

	// READ-ONLY; The link used to get the document from community or Azure site.
	DocumentationLink *string

	// READ-ONLY; If is the configuration pending restart or not.
	IsConfigPendingRestart *IsConfigPendingRestart

	// READ-ONLY; If is the configuration dynamic.
	IsDynamicConfig *IsDynamicConfig

	// READ-ONLY; If is the configuration read only.
	IsReadOnly *IsReadOnly
}

// DataEncryption - The date encryption for cmk.
type DataEncryption struct {
	// Geo backup key uri as key vault can't cross region, need cmk in same region as geo backup
	GeoBackupKeyURI *string

	// Geo backup user identity resource id as identity can't cross region, need identity in same region as geo backup
	GeoBackupUserAssignedIdentityID *string

	// Primary key uri
	PrimaryKeyURI *string

	// Primary user identity resource id
	PrimaryUserAssignedIdentityID *string

	// The key type, AzureKeyVault for enable cmk, SystemManaged for disable cmk.
	Type *DataEncryptionType
}

// Database - Represents a Database.
type Database struct {
	// The properties of a database.
	Properties *DatabaseProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// DatabaseListResult - A List of databases.
type DatabaseListResult struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of databases housed in a server
	Value []*Database
}

// DatabaseProperties - The properties of a database.
type DatabaseProperties struct {
	// The charset of the database.
	Charset *string

	// The collation of the database.
	Collation *string
}

// DelegatedSubnetUsage - Delegated subnet usage data.
type DelegatedSubnetUsage struct {
	// READ-ONLY; name of the subnet
	SubnetName *string

	// READ-ONLY; Number of used delegated subnets
	Usage *int64
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// FirewallRule - Represents a server firewall rule.
type FirewallRule struct {
	// REQUIRED; The properties of a firewall rule.
	Properties *FirewallRuleProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// FirewallRuleListResult - A list of firewall rules.
type FirewallRuleListResult struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of firewall rules in a server.
	Value []*FirewallRule
}

// FirewallRuleProperties - The properties of a server firewall rule.
type FirewallRuleProperties struct {
	// REQUIRED; The end IP address of the server firewall rule. Must be IPv4 format.
	EndIPAddress *string

	// REQUIRED; The start IP address of the server firewall rule. Must be IPv4 format.
	StartIPAddress *string
}

// FullBackupStoreDetails is used for scenarios where backup data is streamed/copied over to a storage destination.
type FullBackupStoreDetails struct {
	// REQUIRED; Type of the specific object - used for deserializing
	ObjectType *string

	// REQUIRED; SASUriList of storage containers where backup data is to be streamed/copied.
	SasURIList []*string
}

// GetBackupStoreDetails implements the BackupStoreDetailsClassification interface for type FullBackupStoreDetails.
func (f *FullBackupStoreDetails) GetBackupStoreDetails() *BackupStoreDetails {
	return &BackupStoreDetails{
		ObjectType: f.ObjectType,
	}
}

// GetPrivateDNSZoneSuffixResponse - The response of get private dns zone suffix.
type GetPrivateDNSZoneSuffixResponse struct {
	// Represents the private DNS zone suffix.
	PrivateDNSZoneSuffix *string
}

// HighAvailability - Network related properties of a server
type HighAvailability struct {
	// High availability mode for a server.
	Mode *HighAvailabilityMode

	// Availability zone of the standby server.
	StandbyAvailabilityZone *string

	// READ-ONLY; The state of server high availability.
	State *HighAvailabilityState
}

// ImportSourceProperties - Import source related properties.
type ImportSourceProperties struct {
	// Relative path of data directory in storage.
	DataDirPath *string

	// Sas token for accessing source storage. Read and list permissions are required for sas token.
	SasToken *string

	// Storage type of import source.
	StorageType *ImportSourceStorageType

	// Uri of the import source storage.
	StorageURL *string
}

// LogFile - Represents a logFile.
type LogFile struct {
	// The properties of a logFile.
	Properties *LogFileProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// LogFileListResult - A List of logFiles.
type LogFileListResult struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of logFiles in a server
	Value []*LogFile
}

// LogFileProperties - The properties of a logFile.
type LogFileProperties struct {
	// Creation timestamp of the log file.
	CreatedTime *time.Time

	// Last modified timestamp of the log file.
	LastModifiedTime *time.Time

	// The size in kb of the logFile.
	SizeInKB *int64

	// Type of the log file.
	Type *string

	// The url to download the log file from.
	URL *string
}

// MaintenanceWindow - Maintenance window of a server.
type MaintenanceWindow struct {
	// indicates whether custom window is enabled or disabled
	CustomWindow *string

	// day of week for maintenance window
	DayOfWeek *int32

	// start hour for maintenance window
	StartHour *int32

	// start minute for maintenance window
	StartMinute *int32
}

// MySQLServerIdentity - Properties to configure Identity for Bring your Own Keys
type MySQLServerIdentity struct {
	// Type of managed service identity.
	Type *ManagedServiceIdentityType

	// Metadata of user assigned identity.
	UserAssignedIdentities map[string]any

	// READ-ONLY; ObjectId from the KeyVault
	PrincipalID *string

	// READ-ONLY; TenantId from the KeyVault
	TenantID *string
}

// MySQLServerSKU - Billing information related properties of a server.
type MySQLServerSKU struct {
	// REQUIRED; The name of the sku, e.g. StandardD32sv3.
	Name *string

	// REQUIRED; The tier of the particular SKU, e.g. GeneralPurpose.
	Tier *ServerSKUTier
}

// NameAvailability - Represents a resource name availability.
type NameAvailability struct {
	// Error Message.
	Message *string

	// Indicates whether the resource name is available.
	NameAvailable *bool

	// Reason for name being unavailable.
	Reason *string
}

// NameAvailabilityRequest - Request from client to check resource name availability.
type NameAvailabilityRequest struct {
	// REQUIRED; Resource name to verify.
	Name *string

	// Resource type used for verification.
	Type *string
}

// Network related properties of a server
type Network struct {
	// Delegated subnet resource id used to setup vnet for a server.
	DelegatedSubnetResourceID *string

	// Private DNS zone resource id.
	PrivateDNSZoneResourceID *string

	// Whether or not public network access is allowed for this server. Value is 'Disabled' when server has VNet integration.
	PublicNetworkAccess *EnableStatusEnum
}

// Operation - REST API operation definition.
type Operation struct {
	// The localized display information for this particular operation or action.
	Display *OperationDisplay

	// The name of the operation being performed on this particular object.
	Name *string

	// The intended executor of the operation.
	Origin *string

	// Additional descriptions for the operation.
	Properties map[string]any
}

// OperationDisplay - Display metadata associated with the operation.
type OperationDisplay struct {
	// Operation description.
	Description *string

	// Localized friendly name for the operation.
	Operation *string

	// Operation resource provider name.
	Provider *string

	// Resource on which the operation is performed.
	Resource *string
}

// OperationListResult - A list of resource provider operations.
type OperationListResult struct {
	// URL client should use to fetch the next page (per server side paging).
	NextLink *string

	// Collection of available operation details
	Value []*Operation
}

// OperationStatusExtendedResult - Represents Operation Results API Response
type OperationStatusExtendedResult struct {
	// REQUIRED; Operation status.
	Status *string

	// The end time of the operation.
	EndTime *time.Time

	// If present, details of the operation error.
	Error *ErrorDetail

	// Fully qualified ID for the async operation.
	ID *string

	// Name of the async operation.
	Name *string

	// The operations list.
	Operations []*OperationStatusResult

	// Percent of the operation that is complete.
	PercentComplete *float32

	// The extended properties of Operation Results
	Properties map[string]any

	// The start time of the operation.
	StartTime *time.Time

	// READ-ONLY; Fully qualified ID of the resource against which the original async operation was started.
	ResourceID *string
}

// OperationStatusResult - The current status of an async operation.
type OperationStatusResult struct {
	// REQUIRED; Operation status.
	Status *string

	// The end time of the operation.
	EndTime *time.Time

	// If present, details of the operation error.
	Error *ErrorDetail

	// Fully qualified ID for the async operation.
	ID *string

	// Name of the async operation.
	Name *string

	// The operations list.
	Operations []*OperationStatusResult

	// Percent of the operation that is complete.
	PercentComplete *float32

	// The start time of the operation.
	StartTime *time.Time

	// READ-ONLY; Fully qualified ID of the resource against which the original async operation was started.
	ResourceID *string
}

// PrivateEndpoint - The private endpoint resource.
type PrivateEndpoint struct {
	// READ-ONLY; The ARM identifier for private endpoint.
	ID *string
}

// PrivateEndpointConnection - The private endpoint connection resource.
type PrivateEndpointConnection struct {
	// Resource properties.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateEndpointConnectionProperties - Properties of the private endpoint connection.
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState

	// The private endpoint resource.
	PrivateEndpoint *PrivateEndpoint

	// READ-ONLY; The group ids for the private endpoint resource.
	GroupIDs []*string

	// READ-ONLY; The provisioning state of the private endpoint connection resource.
	ProvisioningState *PrivateEndpointConnectionProvisioningState
}

// PrivateLinkServiceConnectionState - A collection of information about the state of the connection between service consumer
// and provider.
type PrivateLinkServiceConnectionState struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string

	// The reason for approval/rejection of the connection.
	Description *string

	// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SKUCapability - Sku capability
type SKUCapability struct {
	// READ-ONLY; vCore name
	Name *string

	// READ-ONLY; supported IOPS
	SupportedIops *int64

	// READ-ONLY; supported memory per vCore in MB
	SupportedMemoryPerVCoreMB *int64

	// READ-ONLY; supported vCores
	VCores *int64
}

// SKUCapabilityV2 - Sku capability
type SKUCapabilityV2 struct {
	// READ-ONLY; vCore name
	Name *string

	// READ-ONLY; Supported high availability mode
	SupportedHAMode []*string

	// READ-ONLY; supported IOPS
	SupportedIops *int64

	// READ-ONLY; supported memory per vCore in MB
	SupportedMemoryPerVCoreMB *int64

	// READ-ONLY; Supported zones
	SupportedZones []*string

	// READ-ONLY; supported vCores
	VCores *int64
}

// Server - Represents a server.
type Server struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The cmk identity for the server.
	Identity *MySQLServerIdentity

	// Properties of the server.
	Properties *ServerProperties

	// The SKU (pricing tier) of the server.
	SKU *MySQLServerSKU

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ServerBackup - Server backup properties
type ServerBackup struct {
	// The properties of a server backup.
	Properties *ServerBackupProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ServerBackupListResult - A list of server backups.
type ServerBackupListResult struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of backups of a server.
	Value []*ServerBackup
}

// ServerBackupProperties - The properties of a server backup.
type ServerBackupProperties struct {
	// Backup type.
	BackupType *string

	// Backup completed time (ISO8601 format).
	CompletedTime *time.Time

	// Backup source
	Source *string
}

// ServerEditionCapability - Server edition capabilities.
type ServerEditionCapability struct {
	// READ-ONLY; Server edition name
	Name *string

	// READ-ONLY; A list of supported server versions.
	SupportedServerVersions []*ServerVersionCapability

	// READ-ONLY; A list of supported storage editions
	SupportedStorageEditions []*StorageEditionCapability
}

// ServerEditionCapabilityV2 - Server edition capabilities.
type ServerEditionCapabilityV2 struct {
	// READ-ONLY; Default Sku name
	DefaultSKU *string

	// READ-ONLY; Default storage size
	DefaultStorageSize *int32

	// READ-ONLY; Server edition name
	Name *string

	// READ-ONLY; A list of supported Skus
	SupportedSKUs []*SKUCapabilityV2

	// READ-ONLY; A list of supported storage editions
	SupportedStorageEditions []*StorageEditionCapability
}

// ServerForUpdate - Parameters allowed to update for a server.
type ServerForUpdate struct {
	// The cmk identity for the server.
	Identity *MySQLServerIdentity

	// The properties that can be updated for a server.
	Properties *ServerPropertiesForUpdate

	// The SKU (pricing tier) of the server.
	SKU *MySQLServerSKU

	// Application-specific metadata in the form of key-value pairs.
	Tags map[string]*string
}

// ServerGtidSetParameter - Server Gtid set parameters.
type ServerGtidSetParameter struct {
	// The Gtid set of server.
	GtidSet *string
}

// ServerListResult - A list of servers.
type ServerListResult struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of servers
	Value []*Server
}

// ServerProperties - The properties of a server.
type ServerProperties struct {
	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for
	// creation).
	AdministratorLogin *string

	// The password of the administrator login (required for server creation).
	AdministratorLoginPassword *string

	// availability Zone information of the server.
	AvailabilityZone *string

	// Backup related properties of a server.
	Backup *Backup

	// The mode to create a new MySQL server.
	CreateMode *CreateMode

	// The Data Encryption for CMK.
	DataEncryption *DataEncryption

	// High availability related properties of a server.
	HighAvailability *HighAvailability

	// Source properties for import from storage.
	ImportSourceProperties *ImportSourceProperties

	// Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindow

	// Network related properties of a server.
	Network *Network

	// The replication role.
	ReplicationRole *ReplicationRole

	// Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime *time.Time

	// The source MySQL server id.
	SourceServerResourceID *string

	// Storage related properties of a server.
	Storage *Storage

	// Server version.
	Version *ServerVersion

	// READ-ONLY; The fully qualified domain name of a server.
	FullyQualifiedDomainName *string

	// READ-ONLY; PrivateEndpointConnections related properties of a server.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// READ-ONLY; The maximum number of replicas that a primary server can have.
	ReplicaCapacity *int32

	// READ-ONLY; The state of a server.
	State *ServerState
}

// ServerPropertiesForUpdate - The properties that can be updated for a server.
type ServerPropertiesForUpdate struct {
	// The password of the administrator login.
	AdministratorLoginPassword *string

	// Backup related properties of a server.
	Backup *Backup

	// The Data Encryption for CMK.
	DataEncryption *DataEncryption

	// High availability related properties of a server.
	HighAvailability *HighAvailability

	// Maintenance window of a server.
	MaintenanceWindow *MaintenanceWindow

	// Network related properties of a server
	Network *Network

	// The replication role of the server.
	ReplicationRole *ReplicationRole

	// Storage related properties of a server.
	Storage *Storage

	// Server version.
	Version *ServerVersion
}

// ServerRestartParameter - Server restart parameters.
type ServerRestartParameter struct {
	// The maximum allowed failover time in seconds.
	MaxFailoverSeconds *int32

	// Whether or not failover to standby server when restarting a server with high availability enabled.
	RestartWithFailover *EnableStatusEnum
}

// ServerVersionCapability - Server version capabilities.
type ServerVersionCapability struct {
	// READ-ONLY; server version
	Name *string

	// READ-ONLY; A list of supported Skus
	SupportedSKUs []*SKUCapability
}

// ServerVersionCapabilityV2 - Server version capabilities.
type ServerVersionCapabilityV2 struct {
	// READ-ONLY; server version
	Name *string
}

// Storage Profile properties of a server
type Storage struct {
	// Enable Storage Auto Grow or not.
	AutoGrow *EnableStatusEnum

	// Enable IO Auto Scaling or not.
	AutoIoScaling *EnableStatusEnum

	// Storage IOPS for a server.
	Iops *int32

	// Enable Log On Disk or not.
	LogOnDisk *EnableStatusEnum

	// Max storage size allowed for a server.
	StorageSizeGB *int32

	// READ-ONLY; The sku name of the server storage.
	StorageSKU *string
}

// StorageEditionCapability - storage edition capability
type StorageEditionCapability struct {
	// READ-ONLY; Maximum backup retention days
	MaxBackupRetentionDays *int64

	// READ-ONLY; The maximum supported storage size.
	MaxStorageSize *int64

	// READ-ONLY; Minimal backup retention days
	MinBackupRetentionDays *int64

	// READ-ONLY; The minimal supported storage size.
	MinStorageSize *int64

	// READ-ONLY; storage edition name
	Name *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// UserAssignedIdentity - Metadata of user assigned identity.
type UserAssignedIdentity struct {
	// READ-ONLY; Client Id of user assigned identity
	ClientID *string

	// READ-ONLY; Principal Id of user assigned identity
	PrincipalID *string
}

// ValidateBackupResponse - Represents ValidateBackup API Response
type ValidateBackupResponse struct {
	// The response properties of a pre backup operation.
	Properties *ValidateBackupResponseProperties
}

// ValidateBackupResponseProperties - ValidateBackup Response Properties
type ValidateBackupResponseProperties struct {
	// Estimated no of storage containers required for resource data to be backed up.
	NumberOfContainers *int32
}

// VirtualNetworkSubnetUsageParameter - Virtual network subnet usage parameter
type VirtualNetworkSubnetUsageParameter struct {
	// Virtual network resource id.
	VirtualNetworkResourceID *string
}

// VirtualNetworkSubnetUsageResult - Virtual network subnet usage data.
type VirtualNetworkSubnetUsageResult struct {
	// READ-ONLY; A list of delegated subnet usage
	DelegatedSubnetsUsage []*DelegatedSubnetUsage

	// READ-ONLY; The location name.
	Location *string

	// READ-ONLY; The subscription id.
	SubscriptionID *string
}
