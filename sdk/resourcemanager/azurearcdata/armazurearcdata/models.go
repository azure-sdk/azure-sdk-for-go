//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurearcdata

import "time"

// ActiveDirectoryConnectorDNSDetails - DNS server details
type ActiveDirectoryConnectorDNSDetails struct {
	// REQUIRED; List of Active Directory DNS server IP addresses.
	NameserverIPAddresses []*string

	// DNS domain name for which DNS lookups should be forwarded to the Active Directory DNS servers.
	DomainName *string

	// Flag indicating whether to prefer Kubernetes DNS server response over AD DNS server response for IP address lookups.
	PreferK8SDNSForPtrLookups *bool

	// Replica count for DNS proxy service. Default value is 1.
	Replicas *int64
}

// ActiveDirectoryConnectorDomainDetails - Active Directory domain details
type ActiveDirectoryConnectorDomainDetails struct {
	// REQUIRED; Name (uppercase) of the Active Directory domain that this AD connector will be associated with.
	Realm *string

	// null
	DomainControllers *ActiveDirectoryDomainControllers

	// NETBIOS name of the Active Directory domain.
	NetbiosDomainName *string

	// The distinguished name of the Active Directory Organizational Unit.
	OuDistinguishedName *string

	// The service account provisioning mode for this Active Directory connector.
	ServiceAccountProvisioning *AccountProvisioningMode
}

// ActiveDirectoryConnectorListResult - A list of active directory connectors
type ActiveDirectoryConnectorListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; Array of results.
	Value []*ActiveDirectoryConnectorResource
}

// ActiveDirectoryConnectorProperties - The properties of an Active Directory connector resource
type ActiveDirectoryConnectorProperties struct {
	// REQUIRED; null
	Spec *ActiveDirectoryConnectorSpec

	// Username and password for domain service account authentication.
	DomainServiceAccountLoginInformation *BasicLoginInformation

	// null
	Status *ActiveDirectoryConnectorStatus

	// READ-ONLY; The provisioning state of the Active Directory connector resource.
	ProvisioningState *string
}

// ActiveDirectoryConnectorResource - Active directory connector resource
type ActiveDirectoryConnectorResource struct {
	// REQUIRED; null
	Properties *ActiveDirectoryConnectorProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ActiveDirectoryConnectorSpec - The specifications of the AD Kubernetes resource.
type ActiveDirectoryConnectorSpec struct {
	// REQUIRED; null
	ActiveDirectory *ActiveDirectoryConnectorDomainDetails

	// REQUIRED; null
	DNS *ActiveDirectoryConnectorDNSDetails
}

// ActiveDirectoryConnectorStatus - The status of the Kubernetes custom resource.
type ActiveDirectoryConnectorStatus struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// The time that the custom resource was last updated.
	LastUpdateTime *string

	// The version of the replicaSet associated with the AD connector custom resource.
	ObservedGeneration *int64

	// The state of the AD connector custom resource.
	State *string
}

// ActiveDirectoryDomainController - Information about a domain controller in the AD domain.
type ActiveDirectoryDomainController struct {
	// REQUIRED; Fully-qualified domain name of a domain controller in the AD domain.
	Hostname *string
}

// ActiveDirectoryDomainControllers - Details about the Active Directory domain controllers associated with this AD connector
// instance
type ActiveDirectoryDomainControllers struct {
	// Information about the Primary Domain Controller (PDC) in the AD domain.
	PrimaryDomainController *ActiveDirectoryDomainController

	// null
	SecondaryDomainControllers []*ActiveDirectoryDomainController
}

// ActiveDirectoryInformation - Active Directory information that related to the resource.
type ActiveDirectoryInformation struct {
	// Keytab information that is used for the Sql Managed Instance when Active Directory authentication is used.
	KeytabInformation *KeytabInformation
}

// ArcSQLServerDatabaseListResult - A list of Arc Sql Server database.
type ArcSQLServerDatabaseListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; Array of Arc Sql Server database.
	Value []*SQLServerDatabaseResource
}

// BasicLoginInformation - Username and password for basic login authentication.
type BasicLoginInformation struct {
	// Login password.
	Password *string

	// Login username.
	Username *string
}

// CommonSKU - The resource model definition representing SKU for ARM resources
type CommonSKU struct {
	// REQUIRED; The name of the SKU. It is typically a letter+number code
	Name *string

	// If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the
	// resource this may be omitted.
	Capacity *int32

	// Whether dev/test is enabled. When the dev field is set to true, the resource is used for dev/test purpose.
	Dev *bool

	// If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string

	// The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
	Size *string
}

// DataControllerProperties - The data controller properties.
type DataControllerProperties struct {
	// Deprecated. Azure Arc Data Services data controller no longer expose any endpoint. All traffic are exposed through Kubernetes
	// native API.
	BasicLoginInformation *BasicLoginInformation

	// If a CustomLocation is provided, this contains the ARM id of the connected cluster the custom location belongs to.
	ClusterID *string

	// If a CustomLocation is provided, this contains the ARM id of the extension the custom location belongs to.
	ExtensionID *string

	// The infrastructure the data controller is running on.
	Infrastructure *Infrastructure

	// The raw kubernetes information
	K8SRaw any

	// Last uploaded date from Kubernetes cluster. Defaults to current date time
	LastUploadedDate *time.Time

	// Log analytics workspace id and primary key
	LogAnalyticsWorkspaceConfig *LogAnalyticsWorkspaceConfig

	// Login credential for logs dashboard on the Kubernetes cluster.
	LogsDashboardCredential *BasicLoginInformation

	// Login credential for metrics dashboard on the Kubernetes cluster.
	MetricsDashboardCredential *BasicLoginInformation

	// Properties from the Kubernetes data controller
	OnPremiseProperty *OnPremiseProperty

	// Deprecated. Service principal is deprecated in favor of Arc Kubernetes service extension managed identity.
	UploadServicePrincipal *UploadServicePrincipal

	// Properties on upload watermark. Mostly timestamp for each upload data type
	UploadWatermark *UploadWatermark

	// READ-ONLY; The provisioning state of the Arc Data Controller resource.
	ProvisioningState *string
}

// DataControllerResource - Data controller resource
type DataControllerResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The data controller's properties
	Properties *DataControllerProperties

	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocation

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// DataControllerUpdate - Used for updating a data controller resource.
type DataControllerUpdate struct {
	// The data controller's properties
	Properties *DataControllerProperties

	// Resource tags
	Tags map[string]*string
}

// ErrorResponse - An error response from the Azure Data on Azure Arc service.
type ErrorResponse struct {
	// null
	Error *ErrorResponseBody
}

// ErrorResponseBody - An error response from the Batch service.
type ErrorResponseBody struct {
	// An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string

	// A list of additional details about the error.
	Details []*ErrorResponseBody

	// A message describing the error, intended to be suitable for display in a user interface.
	Message *string

	// The target of the particular error. For example, the name of the property in error.
	Target *string
}

// ExtendedLocation - The complex type of the extended location.
type ExtendedLocation struct {
	// The name of the extended location.
	Name *string

	// The type of the extended location.
	Type *ExtendedLocationTypes
}

// FailoverGroupListResult - A list of failover groups.
type FailoverGroupListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; Array of failover group results.
	Value []*FailoverGroupResource
}

// FailoverGroupProperties - The properties of a failover group resource.
type FailoverGroupProperties struct {
	// REQUIRED; The resource ID of the partner SQL managed instance.
	PartnerManagedInstanceID *string

	// REQUIRED; The specifications of the failover group resource.
	Spec *FailoverGroupSpec

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// The status of the failover group custom resource.
	Status any

	// READ-ONLY; The provisioning state of the failover group resource.
	ProvisioningState *ProvisioningState
}

// FailoverGroupResource - A failover group resource.
type FailoverGroupResource struct {
	// REQUIRED; null
	Properties *FailoverGroupProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// FailoverGroupSpec - The specifications of the failover group resource.
type FailoverGroupSpec struct {
	// REQUIRED; The role of the SQL managed instance in this failover group.
	Role *InstanceFailoverGroupRole

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// The name of the partner SQL managed instance.
	PartnerMI *string

	// The mirroring endpoint public certificate for the partner SQL managed instance. Only PEM format is supported.
	PartnerMirroringCert *string

	// The mirroring endpoint URL of the partner SQL managed instance.
	PartnerMirroringURL *string

	// The partner sync mode of the SQL managed instance.
	PartnerSyncMode *FailoverGroupPartnerSyncMode

	// The shared name of the failover group for this SQL managed instance. Both SQL managed instance and its partner have to
	// use the same shared name.
	SharedName *string

	// The name of the SQL managed instance with this failover group role.
	SourceMI *string
}

// K8SActiveDirectory - The kubernetes active directory information.
type K8SActiveDirectory struct {
	// Account name for AAD
	AccountName *string
	Connector   *K8SActiveDirectoryConnector

	// An array of encryption types
	EncryptionTypes []*string

	// Keytab secret used to authenticate with Active Directory.
	KeytabSecret *string
}

type K8SActiveDirectoryConnector struct {
	// Name of the connector
	Name *string

	// Name space of the connector
	Namespace *string
}

// K8SNetworkSettings - The kubernetes network settings information.
type K8SNetworkSettings struct {
	// If 1, then SQL Server forces all connections to be encrypted. By default, this option is 0
	Forceencryption *int32

	// Specifies which ciphers are allowed by SQL Server for TLS
	Tlsciphers *string

	// A comma-separated list of which TLS protocols are allowed by SQL Server
	Tlsprotocols *string
}

// K8SResourceRequirements - The kubernetes resource limits and requests used to restrict or reserve resource usage.
type K8SResourceRequirements struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// Limits for a kubernetes resource type (e.g 'cpu', 'memory'). The 'cpu' request must be less than or equal to 'cpu' limit.
	// Default 'cpu' is 2, minimum is 1. Default 'memory' is '4Gi', minimum is '2Gi.
	// If sku.tier is GeneralPurpose, maximum 'cpu' is 24 and maximum 'memory' is '128Gi'.
	Limits map[string]*string

	// Requests for a kubernetes resource type (e.g 'cpu', 'memory'). The 'cpu' request must be less than or equal to 'cpu' limit.
	// Default 'cpu' is 2, minimum is 1. Default 'memory' is '4Gi', minimum is
	// '2Gi. If sku.tier is GeneralPurpose, maximum 'cpu' is 24 and maximum 'memory' is '128Gi'.
	Requests map[string]*string
}

// K8SScheduling - The kubernetes scheduling information.
type K8SScheduling struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// The kubernetes scheduling options. It describes restrictions used to help Kubernetes select appropriate nodes to host the
	// database service
	Default *K8SSchedulingOptions
}

// K8SSchedulingOptions - The kubernetes scheduling options. It describes restrictions used to help Kubernetes select appropriate
// nodes to host the database service
type K8SSchedulingOptions struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// The kubernetes resource limits and requests used to restrict or reserve resource usage.
	Resources *K8SResourceRequirements
}

// K8SSecurity - The kubernetes security information.
type K8SSecurity struct {
	// The kubernetes active directory information.
	ActiveDirectory *K8SActiveDirectory

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// Admin login secret key
	AdminLoginSecret *string

	// Service certificate secret used
	ServiceCertificateSecret *string

	// Transparent data encryption information.
	TransparentDataEncryption *K8StransparentDataEncryption
}

// K8SSettings - The kubernetes settings information.
type K8SSettings struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// The kubernetes network settings information.
	Network *K8SNetworkSettings
}

// K8StransparentDataEncryption - Transparent data encryption information.
type K8StransparentDataEncryption struct {
	// Transparent data encryption mode. Can be Service Managed, Customer managed or disabled
	Mode *string

	// Protector secret for customer managed Transparent data encryption mode
	ProtectorSecret *string
}

// KeytabInformation - Keytab used for authenticate with Active Directory.
type KeytabInformation struct {
	// A base64-encoded keytab.
	Keytab *string
}

// LogAnalyticsWorkspaceConfig - Log analytics workspace id and primary key
type LogAnalyticsWorkspaceConfig struct {
	// Primary key of the workspace
	PrimaryKey *string

	// Azure Log Analytics workspace ID
	WorkspaceID *string
}

// OnPremiseProperty - Properties from the Kubernetes data controller
type OnPremiseProperty struct {
	// REQUIRED; A globally unique ID identifying the associated Kubernetes cluster
	ID *string

	// REQUIRED; Certificate that contains the Kubernetes cluster public key used to verify signing
	PublicSigningKey *string

	// Unique thumbprint returned to customer to verify the certificate being uploaded
	SigningCertificateThumbprint *string
}

// Operation - Azure Data Services on Azure Arc operation definition.
type Operation struct {
	// REQUIRED; The localized display information for this particular operation / action.
	Display *OperationDisplay

	// REQUIRED; Indicates whether the operation is a data action
	IsDataAction *bool

	// REQUIRED; The name of the operation being performed on this particular object.
	Name *string

	// READ-ONLY; The intended executor of the operation.
	Origin *OperationOrigin

	// READ-ONLY; Additional descriptions for the operation.
	Properties map[string]any
}

// OperationDisplay - Display metadata associated with the operation.
type OperationDisplay struct {
	// REQUIRED; The localized friendly description for the operation.
	Description *string

	// REQUIRED; The localized friendly name for the operation.
	Operation *string

	// REQUIRED; The localized friendly form of the resource provider name.
	Provider *string

	// REQUIRED; The localized friendly form of the resource type related to this action/operation.
	Resource *string
}

// OperationListResult - Result of the request to list Azure Data Services on Azure Arc operations.
type OperationListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; Array of results.
	Value []*Operation
}

// PageOfDataControllerResource - A list of data controllers.
type PageOfDataControllerResource struct {
	// Link to retrieve next page of results.
	NextLink *string

	// Array of results.
	Value []*DataControllerResource
}

// PostgresInstance - A Postgres Instance.
type PostgresInstance struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; null
	Properties *PostgresInstanceProperties

	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocation

	// Resource sku.
	SKU *PostgresInstanceSKU

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PostgresInstanceListResult - A list of PostgresInstance.
type PostgresInstanceListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; Array of results.
	Value []*PostgresInstance
}

// PostgresInstanceProperties - Postgres Instance properties.
type PostgresInstanceProperties struct {
	// The instance admin
	Admin *string

	// Username and password for basic authentication.
	BasicLoginInformation *BasicLoginInformation

	// The data controller id
	DataControllerID *string

	// The raw kubernetes information
	K8SRaw any

	// Last uploaded date from Kubernetes cluster. Defaults to current date time
	LastUploadedDate *time.Time

	// READ-ONLY; The provisioning state of the Azure Arc-enabled PostgreSQL instance.
	ProvisioningState *string
}

// PostgresInstanceSKU - The resource model definition representing SKU for Azure Database for PostgresSQL - Azure Arc
type PostgresInstanceSKU struct {
	// REQUIRED; The name of the SKU. It is typically a letter+number code
	Name *string

	// If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the
	// resource this may be omitted.
	Capacity *int32

	// Whether dev/test is enabled. When the dev field is set to true, the resource is used for dev/test purpose.
	Dev *bool

	// If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string

	// The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
	Size *string

	// This field is required to be implemented by the Resource Provider if the service has more than one tier.
	Tier *string
}

// PostgresInstanceUpdate - An update to a Postgres Instance.
type PostgresInstanceUpdate struct {
	// Postgres Instance properties.
	Properties *PostgresInstanceProperties

	// Resource tags.
	Tags map[string]*string
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
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
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SQLManagedInstance - A SqlManagedInstance.
type SQLManagedInstance struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; null
	Properties *SQLManagedInstanceProperties

	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocation

	// Resource sku.
	SKU *SQLManagedInstanceSKU

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SQLManagedInstanceK8SRaw - The raw kubernetes information.
type SQLManagedInstanceK8SRaw struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// The kubernetes spec information.
	Spec *SQLManagedInstanceK8SSpec
}

// SQLManagedInstanceK8SSpec - The kubernetes spec information.
type SQLManagedInstanceK8SSpec struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// This option specifies the number of SQL Managed Instance replicas that will be deployed in your Kubernetes cluster for
	// high availability purposes. If sku.tier is BusinessCritical, allowed values are
	// '2' or '3' with default of '3'. If sku.tier is GeneralPurpose, replicas must be '1'.
	Replicas *int32

	// The kubernetes scheduling information.
	Scheduling *K8SScheduling

	// The kubernetes security information.
	Security *K8SSecurity

	// The kubernetes settings information.
	Settings *K8SSettings
}

// SQLManagedInstanceListResult - A list of SqlManagedInstance.
type SQLManagedInstanceListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; Array of results.
	Value []*SQLManagedInstance
}

// SQLManagedInstanceProperties - Properties of sqlManagedInstance.
type SQLManagedInstanceProperties struct {
	// Active Directory information related to this SQL Managed Instance.
	ActiveDirectoryInformation *ActiveDirectoryInformation

	// The instance admin user
	Admin *string

	// Username and password for basic authentication.
	BasicLoginInformation *BasicLoginInformation

	// If a CustomLocation is provided, this contains the ARM id of the connected cluster the custom location belongs to.
	ClusterID *string

	// null
	DataControllerID *string

	// The instance end time
	EndTime *string

	// If a CustomLocation is provided, this contains the ARM id of the extension the custom location belongs to.
	ExtensionID *string

	// The raw kubernetes information
	K8SRaw *SQLManagedInstanceK8SRaw

	// Last uploaded date from Kubernetes cluster. Defaults to current date time
	LastUploadedDate *time.Time

	// The license type to apply for this managed instance.
	LicenseType *ArcSQLManagedInstanceLicenseType

	// The instance start time
	StartTime *string

	// READ-ONLY; The provisioning state of the Arc-enabled SQL Managed Instance resource.
	ProvisioningState *string
}

// SQLManagedInstanceSKU - The resource model definition representing SKU for Azure Managed Instance - Azure Arc
type SQLManagedInstanceSKU struct {
	// CONSTANT; The name of the SKU.
	// Field has constant value "vCore", any specified value is ignored.
	Name *string

	// The SKU capacity
	Capacity *int32

	// Whether dev/test is enabled. When the dev field is set to true, the resource is used for dev/test purpose.
	Dev *bool

	// The SKU family
	Family *string

	// The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
	Size *string

	// The pricing tier for the instance.
	Tier *SQLManagedInstanceSKUTier
}

// SQLManagedInstanceUpdate - An update to a SQL Managed Instance.
type SQLManagedInstanceUpdate struct {
	// Resource tags.
	Tags map[string]*string
}

// SQLServerDatabaseResource - Arc Sql Server database
type SQLServerDatabaseResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; Properties of Arc Sql Server database
	Properties *SQLServerDatabaseResourceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SQLServerDatabaseResourceProperties - The properties of Arc Sql Server database resource
type SQLServerDatabaseResourceProperties struct {
	BackupInformation *SQLServerDatabaseResourcePropertiesBackupInformation

	// Collation of the database.
	CollationName *string

	// Compatibility level of the database
	CompatibilityLevel *int32

	// Creation date of the database.
	DatabaseCreationDate *time.Time

	// List of features that are enabled for the database
	DatabaseOptions *SQLServerDatabaseResourcePropertiesDatabaseOptions

	// Whether the database is read only or not.
	IsReadOnly *bool

	// Status of the database.
	RecoveryMode *RecoveryMode

	// Size of the database.
	SizeMB *float32

	// Space left of the database.
	SpaceAvailableMB *float32

	// State of the database.
	State *DatabaseState

	// READ-ONLY; The provisioning state of the Arc-enabled SQL Server database resource.
	ProvisioningState *string
}

type SQLServerDatabaseResourcePropertiesBackupInformation struct {
	// Date time of last full backup.
	LastFullBackup *time.Time

	// Date time of last log backup.
	LastLogBackup *time.Time
}

// SQLServerDatabaseResourcePropertiesDatabaseOptions - List of features that are enabled for the database
type SQLServerDatabaseResourcePropertiesDatabaseOptions struct {
	IsAutoCloseOn               *bool
	IsAutoCreateStatsOn         *bool
	IsAutoShrinkOn              *bool
	IsAutoUpdateStatsOn         *bool
	IsEncrypted                 *bool
	IsMemoryOptimizationEnabled *bool
	IsRemoteDataArchiveEnabled  *bool
	IsTrustworthyOn             *bool
}

// SQLServerDatabaseUpdate - An update to database resource.
type SQLServerDatabaseUpdate struct {
	// The data controller's properties
	Properties *SQLServerDatabaseResourceProperties

	// Resource tags.
	Tags map[string]*string
}

// SQLServerInstance - A SqlServerInstance.
type SQLServerInstance struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// null
	Properties *SQLServerInstanceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SQLServerInstanceListResult - A list of SqlServerInstance.
type SQLServerInstanceListResult struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; Array of results.
	Value []*SQLServerInstance
}

// SQLServerInstanceProperties - Properties of SqlServerInstance.
type SQLServerInstanceProperties struct {
	// REQUIRED; ARM Resource id of the container resource (Azure Arc for Servers).
	ContainerResourceID *string

	// REQUIRED; The cloud connectivity status.
	Status *ConnectionStatus

	// Status of Azure Defender.
	AzureDefenderStatus *DefenderStatus

	// Timestamp of last Azure Defender status update.
	AzureDefenderStatusLastUpdated *time.Time

	// SQL Server collation.
	Collation *string

	// The number of total cores of the Operating System Environment (OSE) hosting the SQL Server instance.
	Cores *string

	// SQL Server current version.
	CurrentVersion *string

	// SQL Server edition.
	Edition *EditionType

	// Type of host for Azure Arc SQL Server
	HostType *HostType

	// SQL Server instance name.
	InstanceName *string

	// SQL Server license type.
	LicenseType *ArcSQLServerLicenseType

	// SQL Server update level.
	PatchLevel *string

	// SQL Server product ID.
	ProductID *string

	// Dynamic TCP ports used by SQL Server.
	TCPDynamicPorts *string

	// Static TCP ports used by SQL Server.
	TCPStaticPorts *string

	// The number of logical processors used by the SQL Server instance.
	VCore *string

	// SQL Server version.
	Version *SQLVersion

	// READ-ONLY; The time when the resource was created.
	CreateTime *string

	// READ-ONLY; The provisioning state of the Arc-enabled SQL Server resource.
	ProvisioningState *string
}

// SQLServerInstanceUpdate - An update to a SQL Server Instance.
type SQLServerInstanceUpdate struct {
	// Resource tags.
	Tags map[string]*string
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

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// UploadServicePrincipal - Service principal for uploading billing, metrics and logs.
type UploadServicePrincipal struct {
	// Authority for the service principal. Example: https://login.microsoftonline.com/
	Authority *string

	// Client ID of the service principal for uploading data.
	ClientID *string

	// Secret of the service principal
	ClientSecret *string

	// Tenant ID of the service principal.
	TenantID *string
}

// UploadWatermark - Properties on upload watermark. Mostly timestamp for each upload data type
type UploadWatermark struct {
	// Last uploaded date for logs from kubernetes cluster. Defaults to current date time
	Logs *time.Time

	// Last uploaded date for metrics from kubernetes cluster. Defaults to current date time
	Metrics *time.Time

	// Last uploaded date for usages from kubernetes cluster. Defaults to current date time
	Usages *time.Time
}
