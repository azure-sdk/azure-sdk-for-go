// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnewrelicobservability

import "time"

// AccountInfo - Account Info of the NewRelic account
type AccountInfo struct {
	// Account id
	AccountID *string

	// ingestion key of account
	IngestionKey *string

	// Region where New Relic account is present
	Region *string
}

// AccountProperties - List of all the New relic accounts for the given user
type AccountProperties struct {
	// account id
	AccountID *string

	// account name
	AccountName *string

	// organization id
	OrganizationID *string

	// Region where New Relic account is present
	Region *string
}

// AccountPropertiesForNewRelic - Properties of the NewRelic account
type AccountPropertiesForNewRelic struct {
	// NewRelic Account Information
	AccountInfo *AccountInfo

	// NewRelic Organization Information
	OrganizationInfo *OrganizationInfo

	// date when plan was applied
	SingleSignOnProperties *NewRelicSingleSignOnProperties

	// User id
	UserID *string
}

// AccountResource - The details of a account resource.
type AccountResource struct {
	// The resource-specific properties for this resource.
	Properties *AccountProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// AccountsListResponse - Response of get all accounts Operation.
type AccountsListResponse struct {
	// REQUIRED; The AccountResource items on this page
	Value []*AccountResource

	// The link to the next page of items
	NextLink *string
}

// AppServiceInfo - Details of VM Resource having NewRelic OneAgent installed
type AppServiceInfo struct {
	// Status of the NewRelic agent installed on the App service.
	AgentStatus *string

	// Version of the NewRelic agent installed on the App service.
	AgentVersion *string

	// Azure App service resource ID
	AzureResourceID *string
}

// AppServicesGetRequest - Request of a app services get Operation.
type AppServicesGetRequest struct {
	// REQUIRED; User Email
	UserEmail *string

	// Azure resource IDs
	AzureResourceIDs []*string
}

// AppServicesListResponse - Response of a list app services Operation.
type AppServicesListResponse struct {
	// REQUIRED; The AppServiceInfo items on this page
	Value []*AppServiceInfo

	// The link to the next page of items
	NextLink *string
}

// BillingInfoResponse - Marketplace Subscription and Organization details to which resource gets billed into.
type BillingInfoResponse struct {
	// Marketplace Subscription details
	MarketplaceSaasInfo *MarketplaceSaaSInfo

	// Partner Billing Entity details: Organization Info
	PartnerBillingEntity *PartnerBillingEntity
}

// ConnectedPartnerResourceProperties - Connected Partner Resource Properties
type ConnectedPartnerResourceProperties struct {
	// NewRelic Account Id
	AccountID *string

	// NewRelic account name
	AccountName *string

	// The azure resource Id of the deployment.
	AzureResourceID *string

	// The location of the deployment.
	Location *string
}

// ConnectedPartnerResourcesListFormat - Connected Partner Resources List Format
type ConnectedPartnerResourcesListFormat struct {
	// Connected Partner Resource Properties
	Properties *ConnectedPartnerResourceProperties
}

// ConnectedPartnerResourcesListResponse - List of all active newrelic deployments.
type ConnectedPartnerResourcesListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*ConnectedPartnerResourcesListFormat
}

// FilteringTag - The definition of a filtering tag. Filtering tags are used for capturing resources and include/exclude them
// from being monitored.
type FilteringTag struct {
	// Valid actions for a filtering tag. Exclusion takes priority over inclusion.
	Action *TagAction

	// The name (also known as the key) of the tag.
	Name *string

	// The value of the tag.
	Value *string
}

// HostsGetRequest - Request of a Hosts get Operation.
type HostsGetRequest struct {
	// REQUIRED; User Email
	UserEmail *string

	// VM resource IDs
	VMIDs []*string
}

// LinkedResource - The definition of a linked resource.
type LinkedResource struct {
	// The ARM id of the linked resource.
	ID *string
}

// LinkedResourceListResponse - Response of a list operation.
type LinkedResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*LinkedResource
}

// LogRules - Set of rules for sending logs for the Monitor resource.
type LogRules struct {
	// List of filtering tags to be used for capturing logs. This only takes effect if SendActivityLogs flag is enabled. If empty,
	// all resources will be captured. If only Exclude action is specified, the
	// rules will apply to the list of all available resources. If Include actions are specified, the rules will only include
	// resources with the associated tags.
	FilteringTags []*FilteringTag

	// Flag specifying if AAD logs should be sent for the Monitor resource.
	SendAADLogs *SendAADLogsStatus

	// Flag specifying if activity logs from Azure resources should be sent for the Monitor resource.
	SendActivityLogs *SendActivityLogsStatus

	// Flag specifying if subscription logs should be sent for the Monitor resource.
	SendSubscriptionLogs *SendSubscriptionLogsStatus
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

// MarketplaceSaaSInfo - Marketplace SAAS Info of the resource.
type MarketplaceSaaSInfo struct {
	// The Azure Subscription ID to which the Marketplace Subscription belongs and gets billed into.
	BilledAzureSubscriptionID *string

	// Marketplace Subscription Details: Resource URI
	MarketplaceResourceID *string

	// Marketplace Subscription Details: SaaS Subscription Status
	MarketplaceStatus *string

	// Marketplace Subscription Id. This is a GUID-formatted string.
	MarketplaceSubscriptionID *string

	// Marketplace Subscription Details: SAAS Name
	MarketplaceSubscriptionName *string

	// Offer Id of the Marketplace offer,
	OfferID *string

	// Publisher Id of the Marketplace offer.
	PublisherID *string
}

// MetricRules - Set of rules for sending metrics for the Monitor resource.
type MetricRules struct {
	// List of filtering tags to be used for capturing metrics.
	FilteringTags []*FilteringTag

	// Flag specifying if metrics should be sent for the Monitor resource.
	SendMetrics *SendMetricsStatus

	// User Email
	UserEmail *string
}

// MetricsRequest - Request of get metrics Operation.
type MetricsRequest struct {
	// REQUIRED; User Email
	UserEmail *string
}

// MetricsStatusRequest - Request of get metrics status Operation.
type MetricsStatusRequest struct {
	// REQUIRED; User Email
	UserEmail *string

	// Azure resource IDs
	AzureResourceIDs []*string
}

// MetricsStatusResponse - Response of get metrics status Operation.
type MetricsStatusResponse struct {
	// Azure resource IDs
	AzureResourceIDs []*string
}

// MonitorProperties - Properties specific to the NewRelic Monitor resource
type MonitorProperties struct {
	// Source of account creation
	AccountCreationSource *AccountCreationSource

	// MarketplaceSubscriptionStatus of the resource
	NewRelicAccountProperties *AccountPropertiesForNewRelic

	// Source of org creation
	OrgCreationSource *OrgCreationSource

	// Plan details
	PlanData *PlanData

	// Status of Azure Subscription where Marketplace SaaS is located.
	SaaSAzureSubscriptionStatus *string

	// State of the Azure Subscription containing the monitor resource
	SubscriptionState *string

	// User Info
	UserInfo *UserInfo

	// READ-ONLY; Liftr resource category
	LiftrResourceCategory *LiftrResourceCategories

	// READ-ONLY; Liftr resource preference. The priority of the resource.
	LiftrResourcePreference *int32

	// READ-ONLY; Marketplace Subscription Id
	MarketplaceSubscriptionID *string

	// READ-ONLY; NewRelic Organization properties of the resource
	MarketplaceSubscriptionStatus *MarketplaceSubscriptionStatus

	// READ-ONLY; MonitoringStatus of the resource
	MonitoringStatus *MonitoringStatus

	// READ-ONLY; Provisioning State of the resource
	ProvisioningState *ProvisioningState
}

// MonitoredResource - Details of resource being monitored by NewRelic monitor resource
type MonitoredResource struct {
	// The ARM id of the resource.
	ID *string

	// Reason for why the resource is sending logs (or why it is not sending).
	ReasonForLogsStatus *string

	// Reason for why the resource is sending metrics (or why it is not sending).
	ReasonForMetricsStatus *string

	// Flag indicating if resource is sending logs to NewRelic.
	SendingLogs *SendingLogsStatus

	// Flag indicating if resource is sending metrics to NewRelic.
	SendingMetrics *SendingMetricsStatus
}

// MonitoredResourceListResponse - List of all the resources being monitored by NewRelic monitor resource
type MonitoredResourceListResponse struct {
	// REQUIRED; The MonitoredResource items on this page
	Value []*MonitoredResource

	// The link to the next page of items
	NextLink *string
}

// MonitoredSubscription - The list of subscriptions and it's monitoring status by current NewRelic monitor.
type MonitoredSubscription struct {
	// The reason of not monitoring the subscription.
	Error *string

	// The state of monitoring.
	Status *Status

	// The subscriptionId to be monitored.
	SubscriptionID *string

	// The resource-specific properties for this resource.
	TagRules *MonitoringTagRulesProperties
}

// MonitoredSubscriptionProperties - The request to update subscriptions needed to be monitored by the NewRelic monitor resource.
type MonitoredSubscriptionProperties struct {
	// The request to update subscriptions needed to be monitored by the NewRelic monitor resource.
	Properties *SubscriptionList

	// READ-ONLY; The id of the monitored subscription resource.
	ID *string

	// READ-ONLY; Name of the monitored subscription resource.
	Name *string

	// READ-ONLY; The type of the monitored subscription resource.
	Type *string
}

type MonitoredSubscriptionPropertiesList struct {
	// The link to the next page of items
	NextLink *string
	Value    []*MonitoredSubscriptionProperties
}

// MonitoringTagRulesProperties - The resource-specific properties for this resource.
type MonitoringTagRulesProperties struct {
	// Set of rules for sending logs for the Monitor resource.
	LogRules *LogRules

	// Set of rules for sending metrics for the Monitor resource.
	MetricRules *MetricRules

	// READ-ONLY; Provisioning State of the resource
	ProvisioningState *ProvisioningState
}

// NewRelicMonitorResource - A Monitor Resource by NewRelic
type NewRelicMonitorResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The resource-specific properties for this resource.
	Properties *MonitorProperties

	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

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

// NewRelicMonitorResourceListResult - The response of a NewRelicMonitorResource list operation.
type NewRelicMonitorResourceListResult struct {
	// REQUIRED; The NewRelicMonitorResource items on this page
	Value []*NewRelicMonitorResource

	// The link to the next page of items
	NextLink *string
}

// NewRelicMonitorResourceUpdate - The type used for update operations of the NewRelicMonitorResource.
type NewRelicMonitorResourceUpdate struct {
	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

	// The updatable properties of the NewRelicMonitorResource.
	Properties *NewRelicMonitorResourceUpdateProperties

	// Resource tags.
	Tags map[string]*string
}

// NewRelicMonitorResourceUpdateProperties - The updatable properties of the NewRelicMonitorResource.
type NewRelicMonitorResourceUpdateProperties struct {
	// Source of account creation
	AccountCreationSource *AccountCreationSource

	// MarketplaceSubscriptionStatus of the resource
	NewRelicAccountProperties *AccountPropertiesForNewRelic

	// Source of org creation
	OrgCreationSource *OrgCreationSource

	// Plan details
	PlanData *PlanData

	// User Info
	UserInfo *UserInfo
}

// NewRelicSingleSignOnProperties - Single sign on Info of the NewRelic account
type NewRelicSingleSignOnProperties struct {
	// The Id of the Enterprise App used for Single sign-on.
	EnterpriseAppID *string

	// Provisioning state
	ProvisioningState *ProvisioningState

	// Single sign-on state
	SingleSignOnState *SingleSignOnStates

	// The login URL specific to this NewRelic Organization
	SingleSignOnURL *string
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

// OrganizationInfo - Organization Info of the NewRelic account
type OrganizationInfo struct {
	// Organization id
	OrganizationID *string
}

// OrganizationProperties - Details of Organizations
type OrganizationProperties struct {
	// Billing source
	BillingSource *BillingSource

	// organization id
	OrganizationID *string

	// organization name
	OrganizationName *string
}

// OrganizationResource - The details of a Organization resource.
type OrganizationResource struct {
	// The resource-specific properties for this resource.
	Properties *OrganizationProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// OrganizationsListResponse - Response of get all organizations Operation.
type OrganizationsListResponse struct {
	// REQUIRED; The OrganizationResource items on this page
	Value []*OrganizationResource

	// The link to the next page of items
	NextLink *string
}

// PartnerBillingEntity - Partner Billing details associated with the resource.
type PartnerBillingEntity struct {
	// The New Relic Organization Id.
	OrganizationID *string

	// The New Relic Organization Name.
	OrganizationName *string
}

// PlanData - Plan data of NewRelic Monitor resource
type PlanData struct {
	// Different billing cycles like Monthly/Weekly.
	BillingCycle *string

	// date when plan was applied
	EffectiveDate *time.Time

	// plan id as published by NewRelic
	PlanDetails *string

	// Different usage type like PAYG/COMMITTED. this could be enum
	UsageType *UsageType
}

// PlanDataListResponse - Response of get all plan data Operation.
type PlanDataListResponse struct {
	// REQUIRED; The PlanDataResource items on this page
	Value []*PlanDataResource

	// The link to the next page of items
	NextLink *string
}

// PlanDataProperties - Plan details
type PlanDataProperties struct {
	// Source of account creation
	AccountCreationSource *AccountCreationSource

	// Source of org creation
	OrgCreationSource *OrgCreationSource

	// Plan details
	PlanData *PlanData
}

// PlanDataResource - The details of a PlanData resource.
type PlanDataResource struct {
	// The resource-specific properties for this resource.
	Properties *PlanDataProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ResubscribeProperties - Resubscribe Properties
type ResubscribeProperties struct {
	// Offer Id of the NewRelic offer that needs to be resubscribed
	OfferID *string

	// Organization Id of the NewRelic Organization that needs to be resubscribed
	OrganizationID *string

	// Newly selected plan Id to create the new Marketplace subscription for Resubscribe
	PlanID *string

	// Publisher Id of the NewRelic offer that needs to be resubscribed
	PublisherID *string

	// Newly selected Azure resource group in which the new Marketplace subscription will be created for Resubscribe
	ResourceGroup *string

	// Newly selected Azure Subscription Id in which the new Marketplace subscription will be created for Resubscribe
	SubscriptionID *string

	// Newly selected term Id to create the new Marketplace subscription for Resubscribe
	TermID *string
}

// SubscriptionList - The request to update subscriptions needed to be monitored by the NewRelic monitor resource.
type SubscriptionList struct {
	// List of subscriptions and the state of the monitoring.
	MonitoredSubscriptionList []*MonitoredSubscription

	// The operation for the patch on the resource.
	PatchOperation *PatchOperation

	// READ-ONLY; Provisioning State of the resource
	ProvisioningState *ProvisioningState
}

// SwitchBillingRequest - Request of a switch billing Operation.
type SwitchBillingRequest struct {
	// REQUIRED; User Email
	UserEmail *string

	// Azure resource Id
	AzureResourceID *string

	// Organization id
	OrganizationID *string

	// Plan details
	PlanData *PlanData
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

// TagRule - A tag rule belonging to NewRelic account
type TagRule struct {
	// REQUIRED; The resource-specific properties for this resource.
	Properties *MonitoringTagRulesProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// TagRuleListResult - The response of a TagRule list operation.
type TagRuleListResult struct {
	// REQUIRED; The TagRule items on this page
	Value []*TagRule

	// The link to the next page of items
	NextLink *string
}

// TagRuleUpdate - The type used for update operations of the TagRule.
type TagRuleUpdate struct {
	// The updatable properties of the TagRule.
	Properties *TagRuleUpdateProperties
}

// TagRuleUpdateProperties - The updatable properties of the TagRule.
type TagRuleUpdateProperties struct {
	// Set of rules for sending logs for the Monitor resource.
	LogRules *LogRules

	// Set of rules for sending metrics for the Monitor resource.
	MetricRules *MetricRules
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}

// UserInfo - User Info of NewRelic Monitor resource
type UserInfo struct {
	// country if user
	Country *string

	// User Email
	EmailAddress *string

	// First name
	FirstName *string

	// Last name
	LastName *string

	// Contact phone number
	PhoneNumber *string
}

// VMExtensionPayload - Response of payload to be passed while installing VM agent.
type VMExtensionPayload struct {
	// Ingestion key of the account
	IngestionKey *string
}

// VMHostsListResponse - Response of a list VM Host Operation.
type VMHostsListResponse struct {
	// REQUIRED; The VMInfo items on this page
	Value []*VMInfo

	// The link to the next page of items
	NextLink *string
}

// VMInfo - Details of VM Resource having NewRelic OneAgent installed
type VMInfo struct {
	// Status of the NewRelic agent installed on the VM.
	AgentStatus *string

	// Version of the NewRelic agent installed on the VM.
	AgentVersion *string

	// Azure VM resource ID
	VMID *string
}
