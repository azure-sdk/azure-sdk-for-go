//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconnectedcache

import "time"

// AdditionalCacheNodeProperties - Model representing cache node for connected cache resource
type AdditionalCacheNodeProperties struct {
	// Auto update or fast update version
	AutoUpdateVersion *string

	// Cache node resource Bgp configuration.
	BgpConfiguration *BgpConfiguration

	// issues list to return the issues as part of the additional cache node properties
	CacheNodePropertiesDetailsIssuesList []*string

	// Cache node resource drive configurations.
	DriveConfiguration []*CacheNodeDriveConfiguration

	// Cache node resource requires a proxy
	IsProxyRequired *ProxyRequired

	// Operating system of the cache node
	OSType *OsType

	// Optional property #1 of Mcc response object
	OptionalProperty1 *string

	// Optional property #2 of Mcc response object
	OptionalProperty2 *string

	// Optional property #3 of Mcc response object
	OptionalProperty3 *string

	// Optional property #4 of Mcc response object
	OptionalProperty4 *string

	// Optional property #5 of Mcc response object
	OptionalProperty5 *string

	// Cache node resource Mcc proxy Url
	ProxyURL *string

	// proxyUrl configuration of the cache node
	ProxyURLConfiguration *ProxyURLConfiguration

	// Update Cycle Type
	UpdateCycleType *CycleType

	// Update related information details
	UpdateInfoDetails *string

	// customer requested date time for mcc install of update cycle
	UpdateRequestedDateTime *time.Time

	// READ-ONLY; Cache node resource aggregated status code.
	AggregatedStatusCode *int32

	// READ-ONLY; Cache node resource aggregated status details.
	AggregatedStatusDetails *string

	// READ-ONLY; Cache node resource aggregated status text.
	AggregatedStatusText *string

	// READ-ONLY; Auto update version that is the applied to update on mcc cache node
	AutoUpdateAppliedVersion *string

	// READ-ONLY; Auto update last applied date time of mcc install
	AutoUpdateLastAppliedDateTime *time.Time

	// READ-ONLY; Auto Update status details from the backend after applying the new version details
	AutoUpdateLastAppliedDetails *string

	// READ-ONLY; Last applied auto update state for mcc install of auto update cycle
	AutoUpdateLastAppliedState *string

	// READ-ONLY; Auto update last triggered date time of mcc install
	AutoUpdateLastTriggeredDateTime *time.Time

	// READ-ONLY; Auto update last applied date time of mcc install
	AutoUpdateNextAvailableDateTime *time.Time

	// READ-ONLY; Auto update version that is the Next available version to update on mcc cache node
	AutoUpdateNextAvailableVersion *string

	// READ-ONLY; Cache node resource state as integer.
	CacheNodeState *int32

	// READ-ONLY; Cache node resource detailed state text.
	CacheNodeStateDetailedText *string

	// READ-ONLY; Cache node resource short state text.
	CacheNodeStateShortText *string

	// READ-ONLY; Cache node resource flag indicating if cache node has been physically installed or provisioned on their physical
	// lab.
	IsProvisioned *bool

	// READ-ONLY; Cache node resource Mcc product version.
	ProductVersion *string
}

// AdditionalCustomerProperties - Model representing customer for connected cache resource
type AdditionalCustomerProperties struct {
	// Customer resource Asn (autonomous system number).
	CustomerAsn *string

	// Customer resource contact email.
	CustomerEmail *string

	// Customer resource entitlement expiration date string.
	CustomerEntitlementExpiration *time.Time

	// Customer resource entitlement Sku Guid.
	CustomerEntitlementSKUGUID *string

	// Customer resource entitlement Sku Id.
	CustomerEntitlementSKUID *string

	// Customer resource entitlement Sku name.
	CustomerEntitlementSKUName *string

	// Customer resource transit Asn (autonomous system number).
	CustomerTransitAsn *string

	// Customer resource transit state.
	CustomerTransitState *CustomerTransitState

	// Optional property #1 of Mcc response object.
	OptionalProperty1 *string

	// Optional property #2 of Mcc response object.
	OptionalProperty2 *string

	// Optional property #3 of Mcc response object.
	OptionalProperty3 *string

	// Optional property #4 of Mcc response object.
	OptionalProperty4 *string

	// Optional property #5 of Mcc response object.
	OptionalProperty5 *string

	// READ-ONLY; Customer resource estimated Asn peering peak in Gbps.
	CustomerAsnEstimatedEgressPeekGbps *float32

	// READ-ONLY; Customer resource owner organization name.
	CustomerOrgName *string

	// READ-ONLY; Customer resource average egress in Mbps.
	CustomerPropertiesOverviewAverageEgressMbps *float32

	// READ-ONLY; Customer resource average cache miss throughput in Mbps.
	CustomerPropertiesOverviewAverageMissMbps *float32

	// READ-ONLY; Customer resource cache efficiency.
	CustomerPropertiesOverviewCacheEfficiency *float32

	// READ-ONLY; Customer resource total healthy cache nodes.
	CustomerPropertiesOverviewCacheNodesHealthyCount *int32

	// READ-ONLY; Customer resource total unhealthy cache nodes.
	CustomerPropertiesOverviewCacheNodesUnhealthyCount *int32

	// READ-ONLY; Customer resource maximum egress in Mbps.
	CustomerPropertiesOverviewEgressMbpsMax *float32

	// READ-ONLY; Customer resource peak egress timestamp.
	CustomerPropertiesOverviewEgressMbpsMaxDateTime *time.Time

	// READ-ONLY; Customer resource maximum cache miss throughput in Mbps.
	CustomerPropertiesOverviewMissMbpsMax *float32

	// READ-ONLY; Customer resource peak cache miss throughput timestamp.
	CustomerPropertiesOverviewMissMbpsMaxDateTime *time.Time

	// READ-ONLY; Customer resource last PeeringDB update timestamp.
	PeeringDbLastUpdateDate *time.Time

	// READ-ONLY; Customer resource last PeeringDB update timestamp.
	PeeringDbLastUpdateTime *time.Time

	// READ-ONLY; Customer resource signup phase status code as integer.
	SignupPhaseStatusCode *int32

	// READ-ONLY; Customer resource signup phase status as string text.
	SignupPhaseStatusText *string

	// READ-ONLY; Customer resource signup status as boolean.
	SignupStatus *bool

	// READ-ONLY; Customer resource signup status as integer code.
	SignupStatusCode *int32

	// READ-ONLY; Customer resource signup status as string text.
	SignupStatusText *string
}

// BgpCidrsConfiguration - Mcc cache node Bgp Cidr details.
type BgpCidrsConfiguration struct {
	// READ-ONLY; Mcc cache node Bgp Cidr details.
	BgpCidrs []*string
}

// BgpConfiguration - Bgp configuration of cache node
type BgpConfiguration struct {
	// Asn to ip address mapping
	AsnToIPAddressMapping *string
}

// CacheNodeDriveConfiguration - Drive configuration for cache node
type CacheNodeDriveConfiguration struct {
	// corresponding nginx cache number. Valid cache numbers are 1 - 20
	CacheNumber *int32

	// full binding for corresponding nginx cache drive
	NginxMapping *string

	// physical path location of the folder used for caching content
	PhysicalPath *string

	// physical size of the drive used for caching content
	SizeInGb *int32
}

// CacheNodeEntity - Model representing Cache Node for ConnectedCache resource
type CacheNodeEntity struct {
	// Customer requested day of week for mcc install of auto update cycle
	AutoUpdateRequestedDay *int32

	// Customer requested time of the day for mcc install of auto update cycle, should be hh:mm
	AutoUpdateRequestedTime *string

	// Customer requested week of month for mcc install of auto update cycle
	AutoUpdateRequestedWeek *int32

	// Auto Update Ring Type which is slow or fast etc.
	AutoUpdateRingType *AutoUpdateRingType

	// Cache node resource identifier of the cache node
	CacheNodeID *string

	// Cache node resource name.
	CacheNodeName *string

	// Cache node resource comma separated values of Cidrs.
	CidrCSV []*string

	// Cache node resource current Cidr range precedence selection type.
	CidrSelectionType *int32

	// Cache node resource customer resource Asn (autonomous system number)
	CustomerAsn *int32

	// Cache node resource customer index as string.
	CustomerIndex *string

	// Cache node resource customer resource name.
	CustomerName *string

	// FQDN(fully qualified domain name) value of the mcc cache node
	FullyQualifiedDomainName *string

	// Cache node resource Azure fully qualified resource Id.
	FullyQualifiedResourceID *string

	// Cache node resource Ip address.
	IPAddress *string

	// Cache node resource flag for indicating if cache node is enabled.
	IsEnabled *bool

	// Cache node resource flag for determining if managed by enterprise as boolean.
	IsEnterpriseManaged *bool

	// Cache node resource maximum allowed egress in Mbps.
	MaxAllowableEgressInMbps *int32

	// Cache node resource flag for determining if customer will be migrated.
	ShouldMigrate *bool

	// READ-ONLY; Cache node resource total addressable space defined by the Cidr Csv block.
	AddressSpace *int32

	// READ-ONLY; Cache node resource total addressable space defined by Bgp and Cidr Csv blocks.
	BgpAddressSpace *int32

	// READ-ONLY; Cache node resource Bgp block count.
	BgpCidrBlocksCount *int32

	// READ-ONLY; Cache node resource last Bgp Cidr Csv update timestamp
	BgpCidrCSVLastUpdateTime *time.Time

	// READ-ONLY; Cache node resource bytes truncated from Bgp output file.
	BgpFileBytesTruncated *int32

	// READ-ONLY; Cache node resource last Bgp report timestamp.
	BgpLastReportedTime *time.Time

	// READ-ONLY; Cache node resource Bgp record count.
	BgpNumberOfRecords *int32

	// READ-ONLY; Cache node resource Bgp update count.
	BgpNumberOfTimesUpdated *int32

	// READ-ONLY; Cache node resource Bgp review feedback text.
	BgpReviewFeedback *string

	// READ-ONLY; Cache node resource Bgp review state string text.
	BgpReviewState *BgpReviewStateEnum

	// READ-ONLY; Cache node resource Bgp review state string text in detail.
	BgpReviewStateText *string

	// READ-ONLY; Cache node resource category.
	Category *string

	// READ-ONLY; Cache node resource last Cidr Csv update timestamp
	CidrCSVLastUpdateTime *time.Time

	// READ-ONLY; Cache node resource customer resource client tenant Id of subscription.
	ClientTenantID *string

	// READ-ONLY; Cache node resource configuration state.
	ConfigurationState *ConfigurationState

	// READ-ONLY; Cache node resource configuration state text.
	ConfigurationStateText *string

	// READ-ONLY; Cache node resource container configuration details.
	ContainerConfigurations *string

	// READ-ONLY; Cache node resource Mcc container configuration details re-sync trigger.
	ContainerResyncTrigger *int32

	// READ-ONLY; Cache node resource create async operation Id.
	CreateAsyncOperationID *string

	// READ-ONLY; Cache node resource customer resource GUID Id.
	CustomerID *string

	// READ-ONLY; Cache node resource deletion async operation Id.
	DeleteAsyncOperationID *string

	// READ-ONLY; Cache node resource Mcc Container Id Uri.
	ImageURI *string

	// READ-ONLY; Cache node resource flag for indicating the cache node resource is frozen (not selectable, not editable in UI).
	IsFrozen *bool

	// READ-ONLY; Cache node resource last sync timestamp.
	LastSyncWithAzureTimestamp *time.Time

	// READ-ONLY; Cache node resource last backend updated timestamp.
	LastUpdatedTimestamp *time.Time

	// READ-ONLY; Cache node resource maximum allowed probability of egress.
	MaxAllowableProbability *float32

	// READ-ONLY; Cache node resource release version.
	ReleaseVersion *int32

	// READ-ONLY; Cache node resource review feedback text.
	ReviewFeedback *string

	// READ-ONLY; Cache node resource review process state as integer
	ReviewState *int32

	// READ-ONLY; Cache node resource review state text.
	ReviewStateText *string

	// READ-ONLY; Cache node resource attempts to sync with Azure.
	SynchWithAzureAttemptsCount *int32

	// READ-ONLY; Cache node resource Mcc container deployment worker connection count.
	WorkerConnections *int32

	// READ-ONLY; Cache node resource last updated Mcc container deployment worker connection count timestamp.
	WorkerConnectionsLastUpdatedDateTime *time.Time

	// READ-ONLY; Cache node resource Azure XCid.
	XCid *string
}

// CacheNodeInstallProperties - Mcc cache node resource install script properties.
type CacheNodeInstallProperties struct {
	// Mcc cache node resource Id.
	CacheNodeID *string

	// Mcc customer resource Id.
	CustomerID *string

	// READ-ONLY; Mcc primary account key. Internal to Mcc.
	PrimaryAccountKey *string

	// READ-ONLY; Mcc Iot Central temporary device registration key, used once.
	RegistrationKey *string

	// READ-ONLY; Mcc secondary account key. Internal to Mcc.
	SecondaryAccountKey *string
}

// CacheNodeOldResponse - Model representing Cache Node for ConnectedCache resource
type CacheNodeOldResponse struct {
	// The error details
	Error *ErrorDetail

	// statusCode used to get code details of Mcc response object
	StatusCode *string

	// statusDetails used to get inner details of Mcc response object
	StatusDetails *string

	// statusText used to get status details in string format of Mcc response object
	StatusText *string

	// READ-ONLY; The provisioned state of the resource
	ProvisioningState *ProvisioningState

	// READ-ONLY; status of the HTTP error code
	Status *string
}

// CacheNodePreviewResource - Concrete tracked resource types can be created by aliasing this type using a specific property
// type.
type CacheNodePreviewResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *CacheNodeOldResponse

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

// CacheNodePreviewResourceListResult - The response of a CacheNodePreviewResource list operation.
type CacheNodePreviewResourceListResult struct {
	// REQUIRED; The CacheNodePreviewResource items on this page
	Value []*CacheNodePreviewResource

	// The link to the next page of items
	NextLink *string
}

// CacheNodeProperty - Model representing an Mcc cache node connectedCache resource
type CacheNodeProperty struct {
	// Mcc cache node resource additional properties.
	AdditionalCacheNodeProperties *AdditionalCacheNodeProperties

	// Mcc cache node resource (cache node entity).
	CacheNode *CacheNodeEntity

	// Mcc response error details.
	Error *ErrorDetail

	// Mcc response status code.
	StatusCode *string

	// Mcc response status details for retrieving response inner details.
	StatusDetails *string

	// Mcc response status text as string for retrieving status details.
	StatusText *string

	// READ-ONLY; The provisioned state of the resource
	ProvisioningState *ProvisioningState

	// READ-ONLY; HTTP error status code.
	Status *string
}

// CustomerEntity - Model representing Customer resource for ConnectedCache resource
type CustomerEntity struct {
	// Customer resource client tenant Id of subscription.
	ClientTenantID *string

	// Customer resource contact email.
	ContactEmail *string

	// Customer resource contact full name.
	ContactName *string

	// Customer resource contact phone.
	ContactPhone *string

	// Customer resource name.
	CustomerName *string

	// Customer resource Azure fully qualified resource Id.
	FullyQualifiedResourceID *string

	// Customer resource flag for enterprise management as boolean.
	IsEnterpriseManaged *bool

	// Customer resource entitlement flag as boolean.
	IsEntitled *bool

	// Customer resource Mcc release version.
	ReleaseVersion *int32

	// Customer resource flag for resending signup code as boolean.
	ResendSignupCode *bool

	// Customer resource flag for migration.
	ShouldMigrate *bool

	// Customer resource flag for requiring verification of signup code as boolean.
	VerifySignupCode *bool

	// Customer resource phrase for verifying signup.
	VerifySignupPhrase *string

	// READ-ONLY; Customer resource create async operation Id.
	CreateAsyncOperationID *string

	// READ-ONLY; Customer resource Guid Id.
	CustomerID *string

	// READ-ONLY; Customer resource deletion async operation Id.
	DeleteAsyncOperationID *string

	// READ-ONLY; Customer resource last Azure sync timestamp.
	LastSyncWithAzureTimestamp *time.Time

	// READ-ONLY; Customer resource sync attempts.
	SynchWithAzureAttemptsCount *int32
}

// CustomerProperty - Model representing customer for connectedCache resource
type CustomerProperty struct {
	// Mcc customer resource additional properties.
	AdditionalCustomerProperties *AdditionalCustomerProperties

	// Mcc customer resource (customer entity).
	Customer *CustomerEntity

	// READ-ONLY; Mcc response error details.
	Error *ErrorDetail

	// READ-ONLY; The provisioned state of the resource
	ProvisioningState *ProvisioningState

	// READ-ONLY; HTTP error status code.
	Status *string

	// READ-ONLY; Mcc response status code.
	StatusCode *string

	// READ-ONLY; Mcc response status details for retrieving response inner details.
	StatusDetails *string

	// READ-ONLY; Mcc response status text as string for retrieving status details.
	StatusText *string
}

// EnterpriseMccCacheNodeResource - Represents the high level Nodes needed to provision cache node resources
type EnterpriseMccCacheNodeResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *CacheNodeProperty

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

// EnterpriseMccCacheNodeResourceListResult - The response of a EnterpriseMccCacheNodeResource list operation.
type EnterpriseMccCacheNodeResourceListResult struct {
	// REQUIRED; The EnterpriseMccCacheNodeResource items on this page
	Value []*EnterpriseMccCacheNodeResource

	// The link to the next page of items
	NextLink *string
}

// EnterpriseMccCustomerResource - Represents the high level Nodes needed to provision customer resources
type EnterpriseMccCustomerResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *CustomerProperty

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

// EnterpriseMccCustomerResourceListResult - The response of a EnterpriseMccCustomerResource list operation.
type EnterpriseMccCustomerResourceListResult struct {
	// REQUIRED; The EnterpriseMccCustomerResource items on this page
	Value []*EnterpriseMccCustomerResource

	// The link to the next page of items
	NextLink *string
}

// EnterprisePreviewResource - ConnectedCache Resource. Represents the high level Nodes needed to provision CacheNode and
// customer resources used in private preview
type EnterprisePreviewResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *CacheNodeOldResponse

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

// EnterprisePreviewResourceListResult - The response of a EnterprisePreviewResource list operation.
type EnterprisePreviewResourceListResult struct {
	// REQUIRED; The EnterprisePreviewResource items on this page
	Value []*EnterprisePreviewResource

	// The link to the next page of items
	NextLink *string
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

// IspCacheNodeResource - Represents the high level Nodes needed to provision cache node resources
type IspCacheNodeResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *CacheNodeProperty

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

// IspCacheNodeResourceListResult - The response of a IspCacheNodeResource list operation.
type IspCacheNodeResourceListResult struct {
	// REQUIRED; The IspCacheNodeResource items on this page
	Value []*IspCacheNodeResource

	// The link to the next page of items
	NextLink *string
}

// IspCustomerResource - Represents the high level Nodes needed to provision isp customer resources
type IspCustomerResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *CustomerProperty

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

// IspCustomerResourceListResult - The response of a IspCustomerResource list operation.
type IspCustomerResourceListResult struct {
	// REQUIRED; The IspCustomerResource items on this page
	Value []*IspCustomerResource

	// The link to the next page of items
	NextLink *string
}

// MccCacheNodeBgpCidrDetails - Represents all Cidr details of the Bgp request for a specific cache node resource
type MccCacheNodeBgpCidrDetails struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Mcc cache node resource Bgp Cidr properties.
	Properties *BgpCidrsConfiguration

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

// MccCacheNodeInstallDetails - Mcc cache node resource all install details.
type MccCacheNodeInstallDetails struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Mcc cache node resource install script details.
	Properties *CacheNodeInstallProperties

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

// PatchResource - Mcc PATCH operation properties.
type PatchResource struct {
	// Resource tags.
	Tags map[string]*string
}

// ProxyURLConfiguration - ProxyUrl configuration of cache node
type ProxyURLConfiguration struct {
	// Host Proxy Address configuration along with port number. This can be a proxy or ip address. ex: xx.xx.xx.xxxx:80 or host
	// name http://exampleproxy.com:80
	ProxyURL *string
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
