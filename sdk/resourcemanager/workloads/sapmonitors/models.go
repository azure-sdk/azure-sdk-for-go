//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package sapmonitors

import "time"

// AppServicePlanConfiguration - Configuration details of app service plan
type AppServicePlanConfiguration struct {
	// The number of workers in app service plan. If this is not set or set to 0, auto scale will be configured for the app service
	// plan, otherwise, instance count is set to this number.
	Capacity *int32

	// The App Service plan tier.
	Tier *AppServicePlanTier
}

// Db2ProviderInstanceProperties - Gets or sets the DB2 provider properties.
type Db2ProviderInstanceProperties struct {
	// REQUIRED; The provider type. For example, the value can be SapHana.
	ProviderType *string

	// Gets or sets the db2 database name.
	DbName *string

	// Gets or sets the db2 database password.
	DbPassword *string

	// Gets or sets the key vault URI to secret with the database password.
	DbPasswordURI *string

	// Gets or sets the db2 database sql port.
	DbPort *string

	// Gets or sets the db2 database user name.
	DbUsername *string

	// Gets or sets the target virtual machine name.
	Hostname *string

	// Gets or sets the blob URI to SSL certificate for the DB2 Database.
	SSLCertificateURI *string

	// Gets or sets certificate preference if secure communication is enabled.
	SSLPreference *SSLPreference

	// Gets or sets the SAP System Identifier
	SapSid *string
}

// GetProviderSpecificProperties implements the ProviderSpecificPropertiesClassification interface for type Db2ProviderInstanceProperties.
func (d *Db2ProviderInstanceProperties) GetProviderSpecificProperties() *ProviderSpecificProperties {
	return &ProviderSpecificProperties{
		ProviderType: d.ProviderType,
	}
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

// HanaDbProviderInstanceProperties - Gets or sets the provider properties.
type HanaDbProviderInstanceProperties struct {
	// REQUIRED; The provider type. For example, the value can be SapHana.
	ProviderType *string

	// Gets or sets the hana database name.
	DbName *string

	// Gets or sets the database password.
	DbPassword *string

	// Gets or sets the key vault URI to secret with the database password.
	DbPasswordURI *string

	// Gets or sets the database user name.
	DbUsername *string

	// Gets or sets the target virtual machine size.
	Hostname *string

	// Gets or sets the database instance number.
	InstanceNumber *string

	// Gets or sets the database sql port.
	SQLPort *string

	// Gets or sets the blob URI to SSL certificate for the DB.
	SSLCertificateURI *string

	// Gets or sets the hostname(s) in the SSL certificate.
	SSLHostNameInCertificate *string

	// Gets or sets certificate preference if secure communication is enabled.
	SSLPreference *SSLPreference

	// Gets or sets the SAP System Identifier.
	SapSid *string
}

// GetProviderSpecificProperties implements the ProviderSpecificPropertiesClassification interface for type HanaDbProviderInstanceProperties.
func (h *HanaDbProviderInstanceProperties) GetProviderSpecificProperties() *ProviderSpecificProperties {
	return &ProviderSpecificProperties{
		ProviderType: h.ProviderType,
	}
}

// Health - Resource health details
type Health struct {
	// READ-ONLY; Health state of the resource
	HealthState *WorkloadProviderInstanceHealthState

	// READ-ONLY; Reasons impacting health state
	ImpactingReasons *string
}

// ManagedResourceGroupConfiguration - Managed resource group configuration
type ManagedResourceGroupConfiguration struct {
	// Managed resource group name
	Name *string
}

// ManagedServiceIdentity - The Managed service identity.
type ManagedServiceIdentity struct {
	// REQUIRED; The type of managed service identity.
	Type *ManagedServiceIdentityType

	// The user assigned identities.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
}

// Monitor - SAP monitor info on Azure (ARM properties and SAP monitor properties)
type Monitor struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

	// The resource-specific properties for this resource.
	Properties *MonitorProperties

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

// MonitorListResult - The response of a Monitor list operation.
type MonitorListResult struct {
	// REQUIRED; The Monitor items on this page
	Value []*Monitor

	// The link to the next page of items
	NextLink *string
}

// MonitorProperties - Describes the properties of a SAP monitor.
type MonitorProperties struct {
	// The SAP monitor resources will be deployed in the SAP monitoring region. The subnet region should be same as the SAP monitoring
	// region.
	AppLocation *string

	// App service plan configuration
	AppServicePlanConfiguration *AppServicePlanConfiguration

	// The ARM ID of the Log Analytics Workspace that is used for SAP monitoring.
	LogAnalyticsWorkspaceArmID *string

	// Managed resource group configuration
	ManagedResourceGroupConfiguration *ManagedResourceGroupConfiguration

	// The subnet which the SAP monitor will be deployed in
	MonitorSubnet *string

	// Sets the routing preference of the SAP monitor. By default only RFC1918 traffic is routed to the customer VNET.
	RoutingPreference *RoutingPreference

	// Sets the preference for zone redundancy on resources created for the SAP monitor. By default resources will be created
	// which do not support zone redundancy.
	ZoneRedundancyPreference *string

	// READ-ONLY; Defines the SAP monitor errors.
	Errors *ErrorDetail

	// READ-ONLY; The ARM ID of the MSI used for SAP monitoring.
	MsiArmID *string

	// READ-ONLY; State of provisioning of the SAP monitor.
	ProvisioningState *WorkloadMonitorProvisioningState

	// READ-ONLY; The ARM ID of the Storage account used for SAP monitoring.
	StorageAccountArmID *string
}

// MsSQLServerProviderInstanceProperties - Gets or sets the SQL server provider properties.
type MsSQLServerProviderInstanceProperties struct {
	// REQUIRED; The provider type. For example, the value can be SapHana.
	ProviderType *string

	// Gets or sets the database password.
	DbPassword *string

	// Gets or sets the key vault URI to secret with the database password.
	DbPasswordURI *string

	// Gets or sets the database sql port.
	DbPort *string

	// Gets or sets the database user name.
	DbUsername *string

	// Gets or sets the SQL server host name.
	Hostname *string

	// Gets or sets the blob URI to SSL certificate for the SQL Database.
	SSLCertificateURI *string

	// Gets or sets certificate preference if secure communication is enabled.
	SSLPreference *SSLPreference

	// Gets or sets the SAP System Identifier
	SapSid *string
}

// GetProviderSpecificProperties implements the ProviderSpecificPropertiesClassification interface for type MsSQLServerProviderInstanceProperties.
func (m *MsSQLServerProviderInstanceProperties) GetProviderSpecificProperties() *ProviderSpecificProperties {
	return &ProviderSpecificProperties{
		ProviderType: m.ProviderType,
	}
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// PrometheusHaClusterProviderInstanceProperties - Gets or sets the PrometheusHaCluster provider properties.
type PrometheusHaClusterProviderInstanceProperties struct {
	// REQUIRED; The provider type. For example, the value can be SapHana.
	ProviderType *string

	// Gets or sets the clusterName.
	ClusterName *string

	// Gets or sets the target machine name.
	Hostname *string

	// URL of the Node Exporter endpoint.
	PrometheusURL *string

	// Gets or sets the blob URI to SSL certificate for the HA cluster exporter.
	SSLCertificateURI *string

	// Gets or sets certificate preference if secure communication is enabled.
	SSLPreference *SSLPreference

	// Gets or sets the cluster sid.
	Sid *string
}

// GetProviderSpecificProperties implements the ProviderSpecificPropertiesClassification interface for type PrometheusHaClusterProviderInstanceProperties.
func (p *PrometheusHaClusterProviderInstanceProperties) GetProviderSpecificProperties() *ProviderSpecificProperties {
	return &ProviderSpecificProperties{
		ProviderType: p.ProviderType,
	}
}

// PrometheusOsProviderInstanceProperties - Gets or sets the PrometheusOS provider properties.
type PrometheusOsProviderInstanceProperties struct {
	// REQUIRED; The provider type. For example, the value can be SapHana.
	ProviderType *string

	// URL of the Node Exporter endpoint
	PrometheusURL *string

	// Gets or sets the blob URI to SSL certificate for the prometheus node exporter.
	SSLCertificateURI *string

	// Gets or sets certificate preference if secure communication is enabled.
	SSLPreference *SSLPreference

	// Gets or sets the SAP System Identifier
	SapSid *string
}

// GetProviderSpecificProperties implements the ProviderSpecificPropertiesClassification interface for type PrometheusOsProviderInstanceProperties.
func (p *PrometheusOsProviderInstanceProperties) GetProviderSpecificProperties() *ProviderSpecificProperties {
	return &ProviderSpecificProperties{
		ProviderType: p.ProviderType,
	}
}

// ProviderInstance - A provider instance associated with SAP monitor.
type ProviderInstance struct {
	// The resource-specific properties for this resource.
	Properties *ProviderInstanceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ProviderInstanceListResult - The response of a ProviderInstance list operation.
type ProviderInstanceListResult struct {
	// REQUIRED; The ProviderInstance items on this page
	Value []*ProviderInstance

	// The link to the next page of items
	NextLink *string
}

// ProviderInstanceProperties - Describes the properties of a provider instance.
type ProviderInstanceProperties struct {
	// Defines the provider specific properties.
	ProviderSettings ProviderSpecificPropertiesClassification

	// READ-ONLY; Defines the provider instance errors.
	Errors *ErrorDetail

	// READ-ONLY; Resource health details
	Health *Health

	// READ-ONLY; State of provisioning of the provider instance
	ProvisioningState *WorkloadMonitorProvisioningState
}

// ProviderSpecificProperties - Gets or sets the provider specific properties.
type ProviderSpecificProperties struct {
	// REQUIRED; The provider type. For example, the value can be SapHana.
	ProviderType *string
}

// GetProviderSpecificProperties implements the ProviderSpecificPropertiesClassification interface for type ProviderSpecificProperties.
func (p *ProviderSpecificProperties) GetProviderSpecificProperties() *ProviderSpecificProperties {
	return p
}

// SapLandscapeMonitorMetricThresholds - Gets or sets the Threshold Values for Top Metrics Health.
type SapLandscapeMonitorMetricThresholds struct {
	// Gets or sets the threshold value for Green.
	Green *float32

	// Gets or sets the name of the threshold.
	Name *string

	// Gets or sets the threshold value for Red.
	Red *float32

	// Gets or sets the threshold value for Yellow.
	Yellow *float32
}

// SapLandscapeMonitorProperties - Gets or sets the properties for Sap Landscape Monitor Dashboard.
type SapLandscapeMonitorProperties struct {
	// Gets or sets the SID groupings by landscape and Environment.
	Grouping *SapLandscapeMonitorPropertiesGrouping

	// Gets or sets the list Top Metric Thresholds for SAP Landscape Monitor Dashboard
	TopMetricsThresholds []*SapLandscapeMonitorMetricThresholds

	// READ-ONLY; State of provisioning of the SAP monitor.
	ProvisioningState *SapLandscapeMonitorProvisioningState
}

// SapLandscapeMonitorPropertiesGrouping - Gets or sets the SID groupings by landscape and Environment.
type SapLandscapeMonitorPropertiesGrouping struct {
	// Gets or sets the list of landscape to SID mappings.
	Landscape []*SapLandscapeMonitorSidMapping

	// Gets or sets the list of Sap Applications to SID mappings.
	SapApplication []*SapLandscapeMonitorSidMapping
}

// SapLandscapeMonitorResource - configuration associated with SAP Landscape Monitor Dashboard.
type SapLandscapeMonitorResource struct {
	// The resource-specific properties for this resource.
	Properties *SapLandscapeMonitorProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SapLandscapeMonitorResourceListResult - The response of a SapLandscapeMonitorResource list operation.
type SapLandscapeMonitorResourceListResult struct {
	// REQUIRED; The SapLandscapeMonitorResource items on this page
	Value []*SapLandscapeMonitorResource

	// The link to the next page of items
	NextLink *string
}

// SapLandscapeMonitorSidMapping - Gets or sets the mapping for SID to Environment/Applications.
type SapLandscapeMonitorSidMapping struct {
	// Gets or sets the name of the grouping.
	Name *string

	// Gets or sets the list of SID's.
	TopSid []*string
}

// SapNetWeaverProviderInstanceProperties - Gets or sets the provider properties.
type SapNetWeaverProviderInstanceProperties struct {
	// REQUIRED; The provider type. For example, the value can be SapHana.
	ProviderType *string

	// Gets or sets the blob URI to SSL certificate for the SAP system.
	SSLCertificateURI *string

	// Gets or sets certificate preference if secure communication is enabled.
	SSLPreference *SSLPreference

	// Gets or sets the SAP Client ID.
	SapClientID *string

	// Gets or sets the list of HostFile Entries
	SapHostFileEntries []*string

	// Gets or sets the target virtual machine IP Address/FQDN.
	SapHostname *string

	// Gets or sets the instance number of SAP NetWeaver.
	SapInstanceNr *string

	// Sets the SAP password.
	SapPassword *string

	// Gets or sets the key vault URI to secret with the SAP password.
	SapPasswordURI *string

	// Gets or sets the SAP HTTP port number.
	SapPortNumber *string

	// Gets or sets the SAP System Identifier
	SapSid *string

	// Gets or sets the SAP user name.
	SapUsername *string
}

// GetProviderSpecificProperties implements the ProviderSpecificPropertiesClassification interface for type SapNetWeaverProviderInstanceProperties.
func (s *SapNetWeaverProviderInstanceProperties) GetProviderSpecificProperties() *ProviderSpecificProperties {
	return &ProviderSpecificProperties{
		ProviderType: s.ProviderType,
	}
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

// UpdateMonitorRequest - Defines the request body for updating SAP monitor resource.
type UpdateMonitorRequest struct {
	// The Managed service identity.
	Identity *ManagedServiceIdentity

	// Resource tags.
	Tags map[string]*string
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}
