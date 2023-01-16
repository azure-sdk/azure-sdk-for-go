//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdashboard

import "time"

// AzureMonitorWorkspaceIntegration - Integrations for Azure Monitor Workspace.
type AzureMonitorWorkspaceIntegration struct {
	// The resource Id of the connected Azure Monitor Workspace.
	AzureMonitorWorkspaceResourceID *string `json:"azureMonitorWorkspaceResourceId,omitempty"`
}

// EnterpriseConfigurations - Enterprise settings of a Grafana instance
type EnterpriseConfigurations struct {
	// The AutoRenew setting of the Enterprise subscription
	MarketplaceAutoRenew *MarketplaceAutoRenew `json:"marketplaceAutoRenew,omitempty"`

	// The Plan Id of the Azure Marketplace subscription for the Enterprise plugins
	MarketplacePlanID *string `json:"marketplacePlanId,omitempty"`
}

// EnterpriseDetails - Enterprise details of a Grafana instance
type EnterpriseDetails struct {
	// The allocation details of the per subscription free trial slot of the subscription.
	MarketplaceTrialQuota *MarketplaceTrialQuota `json:"marketplaceTrialQuota,omitempty"`

	// SaaS subscription details of a Grafana instance
	SaasSubscriptionDetails *SaasSubscriptionDetails `json:"saasSubscriptionDetails,omitempty"`
}

// EnterpriseDetailsClientPostOptions contains the optional parameters for the EnterpriseDetailsClient.Post method.
type EnterpriseDetailsClientPostOptions struct {
	// placeholder for future optional parameters
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// GrafanaClientBeginCreateOptions contains the optional parameters for the GrafanaClient.BeginCreate method.
type GrafanaClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GrafanaClientBeginDeleteOptions contains the optional parameters for the GrafanaClient.BeginDelete method.
type GrafanaClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GrafanaClientGetOptions contains the optional parameters for the GrafanaClient.Get method.
type GrafanaClientGetOptions struct {
	// placeholder for future optional parameters
}

// GrafanaClientListByResourceGroupOptions contains the optional parameters for the GrafanaClient.NewListByResourceGroupPager
// method.
type GrafanaClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// GrafanaClientListOptions contains the optional parameters for the GrafanaClient.NewListPager method.
type GrafanaClientListOptions struct {
	// placeholder for future optional parameters
}

// GrafanaClientUpdateOptions contains the optional parameters for the GrafanaClient.Update method.
type GrafanaClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// GrafanaConfigurations - Server configurations of a Grafana instance
type GrafanaConfigurations struct {
	// Email server settings. https://grafana.com/docs/grafana/v9.0/setup-grafana/configure-grafana/#smtp
	SMTP *SMTP `json:"smtp,omitempty"`
}

// GrafanaIntegrations is a bundled observability experience (e.g. pre-configured data source, tailored Grafana dashboards,
// alerting defaults) for common monitoring scenarios.
type GrafanaIntegrations struct {
	AzureMonitorWorkspaceIntegrations []*AzureMonitorWorkspaceIntegration `json:"azureMonitorWorkspaceIntegrations,omitempty"`
}

// ManagedGrafana - The grafana resource type.
type ManagedGrafana struct {
	// The managed identity of the grafana resource.
	Identity *ManagedServiceIdentity `json:"identity,omitempty"`

	// The geo-location where the grafana resource lives
	Location *string `json:"location,omitempty"`

	// Properties specific to the grafana resource.
	Properties *ManagedGrafanaProperties `json:"properties,omitempty"`

	// The Sku of the grafana resource.
	SKU *ResourceSKU `json:"sku,omitempty"`

	// The tags for grafana resource.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; ARM id of the grafana resource
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the grafana resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system meta data relating to this grafana resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the grafana resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

type ManagedGrafanaListResponse struct {
	NextLink *string           `json:"nextLink,omitempty"`
	Value    []*ManagedGrafana `json:"value,omitempty"`
}

// ManagedGrafanaProperties - Properties specific to the grafana resource.
type ManagedGrafanaProperties struct {
	// The api key setting of the Grafana instance.
	APIKey *APIKey `json:"apiKey,omitempty"`

	// Scope for dns deterministic name hash calculation.
	AutoGeneratedDomainNameLabelScope *AutoGeneratedDomainNameLabelScope `json:"autoGeneratedDomainNameLabelScope,omitempty"`

	// Whether a Grafana instance uses deterministic outbound IPs.
	DeterministicOutboundIP *DeterministicOutboundIP `json:"deterministicOutboundIP,omitempty"`

	// Enterprise settings of a Grafana instance
	EnterpriseConfigurations *EnterpriseConfigurations `json:"enterpriseConfigurations,omitempty"`

	// Server configurations of a Grafana instance
	GrafanaConfigurations *GrafanaConfigurations `json:"grafanaConfigurations,omitempty"`

	// GrafanaIntegrations is a bundled observability experience (e.g. pre-configured data source, tailored Grafana dashboards,
	// alerting defaults) for common monitoring scenarios.
	GrafanaIntegrations *GrafanaIntegrations `json:"grafanaIntegrations,omitempty"`

	// Indicate the state for enable or disable traffic over the public interface.
	PublicNetworkAccess *PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// The zone redundancy setting of the Grafana instance.
	ZoneRedundancy *ZoneRedundancy `json:"zoneRedundancy,omitempty"`

	// READ-ONLY; The endpoint of the Grafana instance.
	Endpoint *string `json:"endpoint,omitempty" azure:"ro"`

	// READ-ONLY; The Grafana software version.
	GrafanaVersion *string `json:"grafanaVersion,omitempty" azure:"ro"`

	// READ-ONLY; List of outbound IPs if deterministicOutboundIP is enabled.
	OutboundIPs []*string `json:"outboundIPs,omitempty" azure:"ro"`

	// READ-ONLY; The private endpoint connections of the Grafana instance.
	PrivateEndpointConnections []*PrivateEndpointConnection `json:"privateEndpointConnections,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// ManagedGrafanaPropertiesUpdateParameters - The properties parameters for a PATCH request to a grafana resource.
type ManagedGrafanaPropertiesUpdateParameters struct {
	// The api key setting of the Grafana instance.
	APIKey *APIKey `json:"apiKey,omitempty"`

	// Whether a Grafana instance uses deterministic outbound IPs.
	DeterministicOutboundIP *DeterministicOutboundIP `json:"deterministicOutboundIP,omitempty"`

	// Enterprise settings of a Grafana instance
	EnterpriseConfigurations *EnterpriseConfigurations `json:"enterpriseConfigurations,omitempty"`

	// Server configurations of a Grafana instance
	GrafanaConfigurations *GrafanaConfigurations `json:"grafanaConfigurations,omitempty"`

	// GrafanaIntegrations is a bundled observability experience (e.g. pre-configured data source, tailored Grafana dashboards,
	// alerting defaults) for common monitoring scenarios.
	GrafanaIntegrations *GrafanaIntegrations `json:"grafanaIntegrations,omitempty"`

	// Indicate the state for enable or disable traffic over the public interface.
	PublicNetworkAccess *PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// The zone redundancy setting of the Grafana instance.
	ZoneRedundancy *ZoneRedundancy `json:"zoneRedundancy,omitempty"`
}

// ManagedGrafanaUpdateParameters - The parameters for a PATCH request to a grafana resource.
type ManagedGrafanaUpdateParameters struct {
	// The managed identity of the grafana resource.
	Identity *ManagedServiceIdentity `json:"identity,omitempty"`

	// Properties specific to the managed grafana resource.
	Properties *ManagedGrafanaPropertiesUpdateParameters `json:"properties,omitempty"`

	// The new tags of the grafana resource.
	Tags map[string]*string `json:"tags,omitempty"`
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType `json:"type,omitempty"`

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity `json:"userAssignedIdentities,omitempty"`

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// MarketplaceTrialQuota - The allocation details of the per subscription free trial slot of the subscription.
type MarketplaceTrialQuota struct {
	// Available enterprise promotion for the subscription
	AvailablePromotion *AvailablePromotion `json:"availablePromotion,omitempty"`

	// Resource Id of the Grafana resource which is doing the trial.
	GrafanaResourceID *string `json:"grafanaResourceId,omitempty"`

	// The date and time in UTC of when the trial ends.
	TrialEndAt *time.Time `json:"trialEndAt,omitempty"`

	// The date and time in UTC of when the trial starts.
	TrialStartAt *time.Time `json:"trialStartAt,omitempty"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpoint - The Private Endpoint resource.
type PrivateEndpoint struct {
	// READ-ONLY; The ARM identifier for Private Endpoint
	ID *string `json:"id,omitempty" azure:"ro"`
}

// PrivateEndpointConnection - The Private Endpoint Connection resource.
type PrivateEndpointConnection struct {
	// Resource properties.
	Properties *PrivateEndpointConnectionProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionListResult - List of private endpoint connection associated with the specified storage account
type PrivateEndpointConnectionListResult struct {
	// Array of private endpoint connections
	Value []*PrivateEndpointConnection `json:"value,omitempty"`

	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionProperties - Properties of the PrivateEndpointConnectProperties.
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `json:"privateLinkServiceConnectionState,omitempty"`

	// The private endpoint connection group ids.
	GroupIDs []*string `json:"groupIds,omitempty"`

	// The resource of private end point.
	PrivateEndpoint *PrivateEndpoint `json:"privateEndpoint,omitempty"`

	// READ-ONLY; The provisioning state of the private endpoint connection resource.
	ProvisioningState *PrivateEndpointConnectionProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionsClientBeginApproveOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginApprove
// method.
type PrivateEndpointConnectionsClientBeginApproveOptions struct {
	Body *PrivateEndpointConnection
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
type PrivateEndpointConnectionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListPager
// method.
type PrivateEndpointConnectionsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResource - A private link resource
type PrivateLinkResource struct {
	// Resource properties.
	Properties *PrivateLinkResourceProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateLinkResourceListResult - A list of private link resources
type PrivateLinkResourceListResult struct {
	// Array of private link resources
	Value []*PrivateLinkResource `json:"value,omitempty"`

	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// PrivateLinkResourceProperties - Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// The private link resource Private link DNS zone name.
	RequiredZoneNames []*string `json:"requiredZoneNames,omitempty"`

	// READ-ONLY; The private link resource group id.
	GroupID *string `json:"groupId,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string `json:"requiredMembers,omitempty" azure:"ro"`
}

// PrivateLinkResourcesClientGetOptions contains the optional parameters for the PrivateLinkResourcesClient.Get method.
type PrivateLinkResourcesClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListOptions contains the optional parameters for the PrivateLinkResourcesClient.NewListPager
// method.
type PrivateLinkResourcesClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkServiceConnectionState - A collection of information about the state of the connection between service consumer
// and provider.
type PrivateLinkServiceConnectionState struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string `json:"actionsRequired,omitempty"`

	// The reason for approval/rejection of the connection.
	Description *string `json:"description,omitempty"`

	// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus `json:"status,omitempty"`
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

type ResourceSKU struct {
	// REQUIRED
	Name *string `json:"name,omitempty"`
}

// SMTP - Email server settings. https://grafana.com/docs/grafana/v9.0/setup-grafana/configure-grafana/#smtp
type SMTP struct {
	// Enable this to allow Grafana to send email. Default is false
	Enabled *bool `json:"enabled,omitempty"`

	// Address used when sending out emails https://pkg.go.dev/net/mail#Address
	FromAddress *string `json:"fromAddress,omitempty"`

	// Name to be used when sending out emails. Default is "Azure Managed Grafana Notification" https://pkg.go.dev/net/mail#Address
	FromName *string `json:"fromName,omitempty"`

	// SMTP server hostname with port, e.g. test.email.net:587
	Host *string `json:"host,omitempty"`

	// Password of SMTP auth. If the password contains # or ;, then you have to wrap it with triple quotes. Example: """#password;"""
	Password *string `json:"password,omitempty"`

	// Verify SSL for SMTP server. Default is false https://pkg.go.dev/crypto/tls#Config
	SkipVerify *bool `json:"skipVerify,omitempty"`

	// The StartTLSPolicy setting of the SMTP configuration https://pkg.go.dev/github.com/go-mail/mail#StartTLSPolicy
	StartTLSPolicy *StartTLSPolicy `json:"startTLSPolicy,omitempty"`

	// User of SMTP auth
	User *string `json:"user,omitempty"`
}

// SaasSubscriptionDetails - SaaS subscription details of a Grafana instance
type SaasSubscriptionDetails struct {
	OfferID     *string `json:"offerId,omitempty"`
	PlanID      *string `json:"planId,omitempty"`
	PublisherID *string `json:"publisherId,omitempty"`

	// The current billing term of the SaaS Subscription.
	Term *SubscriptionTerm `json:"term,omitempty"`
}

// SubscriptionTerm - The current billing term of the SaaS Subscription.
type SubscriptionTerm struct {
	// The date and time in UTC of when the billing term ends.
	EndDate *time.Time `json:"endDate,omitempty"`

	// The date and time in UTC of when the billing term starts.
	StartDate *time.Time `json:"startDate,omitempty"`
	TermUnit  *string    `json:"termUnit,omitempty"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string `json:"clientId,omitempty" azure:"ro"`

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`
}
