// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdashboard

import "time"

// AzureMonitorWorkspaceIntegration - Integrations for Azure Monitor Workspace.
type AzureMonitorWorkspaceIntegration struct {
	// The resource Id of the connected Azure Monitor Workspace.
	AzureMonitorWorkspaceResourceID *string
}

// EnterpriseConfigurations - Enterprise settings of a Grafana instance
type EnterpriseConfigurations struct {
	// The AutoRenew setting of the Enterprise subscription
	MarketplaceAutoRenew *MarketplaceAutoRenew

	// The Plan Id of the Azure Marketplace subscription for the Enterprise plugins
	MarketplacePlanID *string
}

// EnterpriseDetails - Enterprise details of a Grafana instance
type EnterpriseDetails struct {
	// The allocation details of the per subscription free trial slot of the subscription.
	MarketplaceTrialQuota *MarketplaceTrialQuota

	// SaaS subscription details of a Grafana instance
	SaasSubscriptionDetails *SaasSubscriptionDetails
}

// GrafanaAvailablePlugin - Available plugins of grafana
type GrafanaAvailablePlugin struct {
	// READ-ONLY; Grafana plugin display name
	Name *string

	// READ-ONLY; Grafana plugin id
	PluginID *string
}

type GrafanaAvailablePluginListResponse struct {
	NextLink *string
	Value    []*GrafanaAvailablePlugin
}

// GrafanaConfigurations - Server configurations of a Grafana instance
type GrafanaConfigurations struct {
	// Email server settings. https://grafana.com/docs/grafana/v9.0/setup-grafana/configure-grafana/#smtp
	SMTP *SMTP

	// Grafana security settings
	Security *Security

	// Grafana Snapshots settings
	Snapshots *Snapshots

	// Grafana users settings
	Users *Users
}

// GrafanaIntegrations is a bundled observability experience (e.g. pre-configured data source, tailored Grafana dashboards,
// alerting defaults) for common monitoring scenarios.
type GrafanaIntegrations struct {
	AzureMonitorWorkspaceIntegrations []*AzureMonitorWorkspaceIntegration
}

// GrafanaPlugin - Plugin of Grafana
type GrafanaPlugin struct {
	// READ-ONLY; Grafana plugin id
	PluginID *string
}

// IntegrationFabric - The integration fabric resource type.
type IntegrationFabric struct {
	// REQUIRED; The geo-location where the resource lives
	Location   *string
	Properties *IntegrationFabricProperties

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

type IntegrationFabricListResponse struct {
	NextLink *string
	Value    []*IntegrationFabric
}

type IntegrationFabricProperties struct {
	// The resource Id of the Azure resource which is used to configure Grafana data source. E.g., an Azure Monitor Workspace,
	// an Azure Data Explorer cluster, etc.
	DataSourceResourceID *string

	// A list of integration scenarios covered by this integration fabric
	Scenarios []*string

	// The resource Id of the Azure resource being integrated with Azure Managed Grafana. E.g., an Azure Kubernetes Service cluster.
	TargetResourceID *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

type IntegrationFabricPropertiesUpdateParameters struct {
	// The new integration scenarios covered by this integration fabric.
	Scenarios []*string
}

// IntegrationFabricUpdateParameters - The parameters for a PATCH request to a Integration Fabric resource.
type IntegrationFabricUpdateParameters struct {
	// The new properties of this Integration Fabric resource
	Properties *IntegrationFabricPropertiesUpdateParameters

	// The new tags of the Integration Fabric resource.
	Tags map[string]*string
}

// ManagedGrafana - The grafana resource type.
type ManagedGrafana struct {
	// The managed identity of the grafana resource.
	Identity *ManagedServiceIdentity

	// The geo-location where the grafana resource lives
	Location *string

	// Properties specific to the grafana resource.
	Properties *ManagedGrafanaProperties

	// The Sku of the grafana resource.
	SKU *ResourceSKU

	// The tags for grafana resource.
	Tags map[string]*string

	// READ-ONLY; ARM id of the grafana resource
	ID *string

	// READ-ONLY; Name of the grafana resource.
	Name *string

	// READ-ONLY; The system meta data relating to this grafana resource.
	SystemData *SystemData

	// READ-ONLY; The type of the grafana resource.
	Type *string
}

type ManagedGrafanaListResponse struct {
	NextLink *string
	Value    []*ManagedGrafana
}

// ManagedGrafanaProperties - Properties specific to the grafana resource.
type ManagedGrafanaProperties struct {
	// The api key setting of the Grafana instance.
	APIKey *APIKey

	// Scope for dns deterministic name hash calculation.
	AutoGeneratedDomainNameLabelScope *AutoGeneratedDomainNameLabelScope

	// Whether a Grafana instance uses deterministic outbound IPs.
	DeterministicOutboundIP *DeterministicOutboundIP

	// Enterprise settings of a Grafana instance
	EnterpriseConfigurations *EnterpriseConfigurations

	// Server configurations of a Grafana instance
	GrafanaConfigurations *GrafanaConfigurations

	// GrafanaIntegrations is a bundled observability experience (e.g. pre-configured data source, tailored Grafana dashboards,
	// alerting defaults) for common monitoring scenarios.
	GrafanaIntegrations *GrafanaIntegrations

	// The major Grafana software version to target.
	GrafanaMajorVersion *string

	// Installed plugin list of the Grafana instance. Key is plugin id, value is plugin definition.
	GrafanaPlugins map[string]*GrafanaPlugin

	// Indicate the state for enable or disable traffic over the public interface.
	PublicNetworkAccess *PublicNetworkAccess

	// The zone redundancy setting of the Grafana instance.
	ZoneRedundancy *ZoneRedundancy

	// READ-ONLY; The endpoint of the Grafana instance.
	Endpoint *string

	// READ-ONLY; The Grafana software version.
	GrafanaVersion *string

	// READ-ONLY; List of outbound IPs if deterministicOutboundIP is enabled.
	OutboundIPs []*string

	// READ-ONLY; The private endpoint connections of the Grafana instance.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// ManagedGrafanaPropertiesUpdateParameters - The properties parameters for a PATCH request to a grafana resource.
type ManagedGrafanaPropertiesUpdateParameters struct {
	// The api key setting of the Grafana instance.
	APIKey *APIKey

	// Whether a Grafana instance uses deterministic outbound IPs.
	DeterministicOutboundIP *DeterministicOutboundIP

	// Enterprise settings of a Grafana instance
	EnterpriseConfigurations *EnterpriseConfigurations

	// Server configurations of a Grafana instance
	GrafanaConfigurations *GrafanaConfigurations

	// GrafanaIntegrations is a bundled observability experience (e.g. pre-configured data source, tailored Grafana dashboards,
	// alerting defaults) for common monitoring scenarios.
	GrafanaIntegrations *GrafanaIntegrations

	// The major Grafana software version to target.
	GrafanaMajorVersion *string

	// Update of Grafana plugin. Key is plugin id, value is plugin definition. If plugin definition is null, plugin with given
	// plugin id will be removed. Otherwise, given plugin will be installed.
	GrafanaPlugins map[string]*GrafanaPlugin

	// Indicate the state for enable or disable traffic over the public interface.
	PublicNetworkAccess *PublicNetworkAccess

	// The zone redundancy setting of the Grafana instance.
	ZoneRedundancy *ZoneRedundancy
}

// ManagedGrafanaUpdateParameters - The parameters for a PATCH request to a grafana resource.
type ManagedGrafanaUpdateParameters struct {
	// The managed identity of the grafana resource.
	Identity *ManagedServiceIdentity

	// Properties specific to the managed grafana resource.
	Properties *ManagedGrafanaPropertiesUpdateParameters
	SKU        *ResourceSKU

	// The new tags of the grafana resource.
	Tags map[string]*string
}

// ManagedPrivateEndpointConnectionState - The state of managed private endpoint connection.
type ManagedPrivateEndpointConnectionState struct {
	// READ-ONLY; Gets or sets the reason for approval/rejection of the connection.
	Description *string

	// READ-ONLY; The approval/rejection status of managed private endpoint connection.
	Status *ManagedPrivateEndpointConnectionStatus
}

// ManagedPrivateEndpointModel - The managed private endpoint resource type.
type ManagedPrivateEndpointModel struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Resource properties.
	Properties *ManagedPrivateEndpointModelProperties

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

// ManagedPrivateEndpointModelListResponse - The list of managed private endpoints of a grafana resource
type ManagedPrivateEndpointModelListResponse struct {
	NextLink *string
	Value    []*ManagedPrivateEndpointModel
}

// ManagedPrivateEndpointModelProperties - Properties specific to the managed private endpoint.
type ManagedPrivateEndpointModelProperties struct {
	// The group Ids of the managed private endpoint.
	GroupIDs []*string

	// The ARM resource ID of the resource for which the managed private endpoint is pointing to.
	PrivateLinkResourceID *string

	// The region of the resource to which the managed private endpoint is pointing to.
	PrivateLinkResourceRegion *string

	// The URL of the data store behind the private link service. It would be the URL in the Grafana data source configuration
	// page without the protocol and port.
	PrivateLinkServiceURL *string

	// User input request message of the managed private endpoint.
	RequestMessage *string

	// READ-ONLY; The state of managed private endpoint connection.
	ConnectionState *ManagedPrivateEndpointConnectionState

	// READ-ONLY; The private IP of private endpoint after approval. This property is empty before connection is approved.
	PrivateLinkServicePrivateIP *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// ManagedPrivateEndpointUpdateParameters - The parameters for a PATCH request to a managed private endpoint.
type ManagedPrivateEndpointUpdateParameters struct {
	// The new tags of the managed private endpoint.
	Tags map[string]*string
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
}

// MarketplaceTrialQuota - The allocation details of the per subscription free trial slot of the subscription.
type MarketplaceTrialQuota struct {
	// Available enterprise promotion for the subscription
	AvailablePromotion *AvailablePromotion

	// Resource Id of the Grafana resource which is doing the trial.
	GrafanaResourceID *string

	// The date and time in UTC of when the trial ends.
	TrialEndAt *time.Time

	// The date and time in UTC of when the trial starts.
	TrialStartAt *time.Time
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

// PrivateEndpoint - The Private Endpoint resource.
type PrivateEndpoint struct {
	// READ-ONLY; The ARM identifier for Private Endpoint
	ID *string
}

// PrivateEndpointConnection - The Private Endpoint Connection resource.
type PrivateEndpointConnection struct {
	// Resource properties.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateEndpointConnectionListResult - List of private endpoint connection associated with the specified storage account
type PrivateEndpointConnectionListResult struct {
	// Array of private endpoint connections
	Value []*PrivateEndpointConnection

	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string
}

// PrivateEndpointConnectionProperties - Properties of the PrivateEndpointConnectProperties.
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState

	// The private endpoint connection group ids.
	GroupIDs []*string

	// The resource of private end point.
	PrivateEndpoint *PrivateEndpoint

	// READ-ONLY; The provisioning state of the private endpoint connection resource.
	ProvisioningState *PrivateEndpointConnectionProvisioningState
}

// PrivateLinkResource - A private link resource
type PrivateLinkResource struct {
	// Resource properties.
	Properties *PrivateLinkResourceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateLinkResourceListResult - A list of private link resources
type PrivateLinkResourceListResult struct {
	// Array of private link resources
	Value []*PrivateLinkResource

	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string
}

// PrivateLinkResourceProperties - Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// The private link resource Private link DNS zone name.
	RequiredZoneNames []*string

	// READ-ONLY; The private link resource group id.
	GroupID *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string
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

type ResourceSKU struct {
	// REQUIRED
	Name *string
}

// SMTP - Email server settings. https://grafana.com/docs/grafana/v9.0/setup-grafana/configure-grafana/#smtp
type SMTP struct {
	// Enable this to allow Grafana to send email. Default is false
	Enabled *bool

	// Address used when sending out emails https://pkg.go.dev/net/mail#Address
	FromAddress *string

	// Name to be used when sending out emails. Default is "Azure Managed Grafana Notification" https://pkg.go.dev/net/mail#Address
	FromName *string

	// SMTP server hostname with port, e.g. test.email.net:587
	Host *string

	// Password of SMTP auth. If the password contains # or ;, then you have to wrap it with triple quotes
	Password *string

	// Verify SSL for SMTP server. Default is false https://pkg.go.dev/crypto/tls#Config
	SkipVerify *bool

	// The StartTLSPolicy setting of the SMTP configuration https://pkg.go.dev/github.com/go-mail/mail#StartTLSPolicy
	StartTLSPolicy *StartTLSPolicy

	// User of SMTP auth
	User *string
}

// SaasSubscriptionDetails - SaaS subscription details of a Grafana instance
type SaasSubscriptionDetails struct {
	// The offer Id of the SaaS subscription.
	OfferID *string

	// The plan Id of the SaaS subscription.
	PlanID *string

	// The publisher Id of the SaaS subscription.
	PublisherID *string

	// The billing term of the SaaS Subscription.
	Term *SubscriptionTerm
}

// Security - Grafana security settings
type Security struct {
	// Set to true to execute the CSRF check even if the login cookie is not in a request (default false).
	CsrfAlwaysCheck *bool
}

// Snapshots - Grafana Snapshots settings
type Snapshots struct {
	// Set to false to disable external snapshot publish endpoint
	ExternalEnabled *bool
}

// SubscriptionTerm - The current billing term of the SaaS Subscription.
type SubscriptionTerm struct {
	// The date and time in UTC of when the billing term ends.
	EndDate *time.Time

	// The date and time in UTC of when the billing term starts.
	StartDate *time.Time

	// The unit of the billing term.
	TermUnit *string
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

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}

// Users - Grafana users settings
type Users struct {
	// Set to true so viewers can access and use explore and perform temporary edits on panels in dashboards they have access
	// to. They cannot save their changes.
	ViewersCanEdit *bool
}
