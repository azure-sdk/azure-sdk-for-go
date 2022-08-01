//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourcehealth

import "time"

// AvailabilityStatus - availabilityStatus of a resource.
type AvailabilityStatus struct {
	// Azure Resource Manager Identity for the availabilityStatuses resource.
	ID *string `json:"id,omitempty"`

	// Azure Resource Manager geo location of the resource.
	Location *string `json:"location,omitempty"`

	// current.
	Name *string `json:"name,omitempty"`

	// Properties of availability state.
	Properties *AvailabilityStatusProperties `json:"properties,omitempty"`

	// Microsoft.ResourceHealth/AvailabilityStatuses.
	Type *string `json:"type,omitempty"`
}

// AvailabilityStatusListResult - The List availabilityStatus operation response.
type AvailabilityStatusListResult struct {
	// REQUIRED; The list of availabilityStatuses.
	Value []*AvailabilityStatus `json:"value,omitempty"`

	// The URI to fetch the next page of availabilityStatuses. Call ListNext() with this URI to fetch the next page of availabilityStatuses.
	NextLink *string `json:"nextLink,omitempty"`
}

// AvailabilityStatusProperties - Properties of availability state.
type AvailabilityStatusProperties struct {
	// Availability status of the resource. When it is null, this availabilityStatus object represents an availability impacting
	// event
	AvailabilityState *AvailabilityStateValues `json:"availabilityState,omitempty"`

	// Details of the availability status.
	DetailedStatus *string `json:"detailedStatus,omitempty"`

	// In case of an availability impacting event, it describes the category of a PlatformInitiated health impacting event. Examples
	// are Planned, Unplanned etc.
	HealthEventCategory *string `json:"healthEventCategory,omitempty"`

	// In case of an availability impacting event, it describes where the health impacting event was originated. Examples are
	// PlatformInitiated, UserInitiated etc.
	HealthEventCause *string `json:"healthEventCause,omitempty"`

	// It is a unique Id that identifies the event
	HealthEventID *string `json:"healthEventId,omitempty"`

	// In case of an availability impacting event, it describes when the health impacting event was originated. Examples are Lifecycle,
	// Downtime, Fault Analysis etc.
	HealthEventType *string `json:"healthEventType,omitempty"`

	// Timestamp for when last change in health status occurred.
	OccuredTime *time.Time `json:"occuredTime,omitempty"`

	// Chronicity of the availability transition.
	ReasonChronicity *ReasonChronicityTypes `json:"reasonChronicity,omitempty"`

	// When the resource's availabilityState is Unavailable, it describes where the health impacting event was originated. Examples
	// are planned, unplanned, user initiated or an outage etc.
	ReasonType *string `json:"reasonType,omitempty"`

	// An annotation describing a change in the availabilityState to Available from Unavailable with a reasonType of type Unplanned
	RecentlyResolved *AvailabilityStatusPropertiesRecentlyResolved `json:"recentlyResolved,omitempty"`

	// Lists actions the user can take based on the current availabilityState of the resource.
	RecommendedActions []*RecommendedAction `json:"recommendedActions,omitempty"`

	// Timestamp for when the health was last checked.
	ReportedTime *time.Time `json:"reportedTime,omitempty"`

	// When the resource's availabilityState is Unavailable and the reasonType is not User Initiated, it provides the date and
	// time for when the issue is expected to be resolved.
	ResolutionETA *time.Time `json:"resolutionETA,omitempty"`

	// When the resource's availabilityState is Unavailable, it provides the Timestamp for when the health impacting event was
	// received.
	RootCauseAttributionTime *time.Time `json:"rootCauseAttributionTime,omitempty"`

	// Lists the service impacting events that may be affecting the health of the resource.
	ServiceImpactingEvents []*ServiceImpactingEvent `json:"serviceImpactingEvents,omitempty"`

	// Summary description of the availability status.
	Summary *string `json:"summary,omitempty"`

	// Title description of the availability status.
	Title *string `json:"title,omitempty"`
}

// AvailabilityStatusPropertiesRecentlyResolved - An annotation describing a change in the availabilityState to Available
// from Unavailable with a reasonType of type Unplanned
type AvailabilityStatusPropertiesRecentlyResolved struct {
	// Timestamp when the availabilityState changes to Available.
	ResolvedTime *time.Time `json:"resolvedTime,omitempty"`

	// Timestamp for when the availabilityState changed to Unavailable
	UnavailableOccuredTime *time.Time `json:"unavailableOccuredTime,omitempty"`

	// Brief description of cause of the resource becoming unavailable.
	UnavailableSummary *string `json:"unavailableSummary,omitempty"`
}

// AvailabilityStatusesClientGetByResourceOptions contains the optional parameters for the AvailabilityStatusesClient.GetByResource
// method.
type AvailabilityStatusesClientGetByResourceOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// AvailabilityStatusesClientListByResourceGroupOptions contains the optional parameters for the AvailabilityStatusesClient.ListByResourceGroup
// method.
type AvailabilityStatusesClientListByResourceGroupOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// AvailabilityStatusesClientListBySubscriptionIDOptions contains the optional parameters for the AvailabilityStatusesClient.ListBySubscriptionID
// method.
type AvailabilityStatusesClientListBySubscriptionIDOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// AvailabilityStatusesClientListOptions contains the optional parameters for the AvailabilityStatusesClient.List method.
type AvailabilityStatusesClientListOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// ChildAvailabilityStatusesClientGetByResourceOptions contains the optional parameters for the ChildAvailabilityStatusesClient.GetByResource
// method.
type ChildAvailabilityStatusesClientGetByResourceOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// ChildAvailabilityStatusesClientListOptions contains the optional parameters for the ChildAvailabilityStatusesClient.List
// method.
type ChildAvailabilityStatusesClientListOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// ChildResourcesClientListOptions contains the optional parameters for the ChildResourcesClient.List method.
type ChildResourcesClientListOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// EmergingIssue - On-going emerging issue from azure status.
type EmergingIssue struct {
	// Timestamp for when last time refreshed for ongoing emerging issue.
	RefreshTimestamp *time.Time `json:"refreshTimestamp,omitempty"`

	// The list of emerging issues of active event type.
	StatusActiveEvents []*StatusActiveEvent `json:"statusActiveEvents,omitempty"`

	// The list of emerging issues of banner type.
	StatusBanners []*StatusBanner `json:"statusBanners,omitempty"`
}

// EmergingIssueImpact - Object of the emerging issue impact on services and regions.
type EmergingIssueImpact struct {
	// The impacted service id.
	ID *string `json:"id,omitempty"`

	// The impacted service name.
	Name *string `json:"name,omitempty"`

	// The list of impacted regions for corresponding emerging issues.
	Regions []*ImpactedRegion `json:"regions,omitempty"`
}

// EmergingIssueListResult - The list of emerging issues.
type EmergingIssueListResult struct {
	// The link used to get the next page of emerging issues.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of emerging issues.
	Value []*EmergingIssuesGetResult `json:"value,omitempty"`
}

// EmergingIssuesClientGetOptions contains the optional parameters for the EmergingIssuesClient.Get method.
type EmergingIssuesClientGetOptions struct {
	// placeholder for future optional parameters
}

// EmergingIssuesClientListOptions contains the optional parameters for the EmergingIssuesClient.List method.
type EmergingIssuesClientListOptions struct {
	// placeholder for future optional parameters
}

// EmergingIssuesGetResult - The Get EmergingIssues operation response.
type EmergingIssuesGetResult struct {
	// The emerging issue entity properties.
	Properties *EmergingIssue `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty" azure:"ro"`

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

// Event - Service health event
type Event struct {
	// Properties of event.
	Properties *EventProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// EventProperties - Properties of event.
type EventProperties struct {
	// Contains additional information regarding the event.
	AdditionalInformation *EventPropertiesAdditionalInformation `json:"additionalInformation,omitempty"`

	// Article of event.
	Article *EventPropertiesArticle `json:"article,omitempty"`

	// Contains the communication message for the event, that could include summary, root cause and other details.
	Description *string `json:"description,omitempty"`

	// The expected duration of the event in seconds.
	Duration *int32 `json:"duration,omitempty"`

	// Tells if we want to enable or disable Microsoft Support for this event.
	EnableChatWithUs *bool `json:"enableChatWithUs,omitempty"`

	// Tells if we want to enable or disable Microsoft Support for this event.
	EnableMicrosoftSupport *bool `json:"enableMicrosoftSupport,omitempty"`

	// Level of event.
	EventLevel *EventLevelValues `json:"eventLevel,omitempty"`

	// Source of event.
	EventSource *EventSourceValues `json:"eventSource,omitempty"`

	// Type of event.
	EventType *EventTypeValues `json:"eventType,omitempty"`

	// Frequently asked questions for the service health event.
	Faqs []*Faq `json:"faqs,omitempty"`

	// Header text of event.
	Header *string `json:"header,omitempty"`

	// Stage for HIR Document
	HirStage *string `json:"hirStage,omitempty"`

	// List services impacted by the service health event.
	Impact []*Impact `json:"impact,omitempty"`

	// It provides the Timestamp for when the health impacting event resolved.
	ImpactMitigationTime *time.Time `json:"impactMitigationTime,omitempty"`

	// It provides the Timestamp for when the health impacting event started.
	ImpactStartTime *time.Time `json:"impactStartTime,omitempty"`

	// The event's impact type
	ImpactType *string `json:"impactType,omitempty"`

	// It provides information if the event is High incident rate event or not.
	IsHIR *bool `json:"isHIR,omitempty"`

	// It provides the Timestamp for when the health impacting event was last updated.
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`

	// Level of insight.
	Level *LevelValues `json:"level,omitempty"`

	// Useful links of event.
	Links []*Link `json:"links,omitempty"`

	// Is true if the event is platform initiated.
	PlatformInitiated *bool `json:"platformInitiated,omitempty"`

	// Priority level of the event. Has value from 0 to 23. 0 is the highest priority. Service issue events have higher priority
	// followed by planned maintenance and health advisory. Critical events have
	// higher priority followed by error, warning and informational. Furthermore, active events have higher priority than resolved.
	Priority *int32 `json:"priority,omitempty"`

	// Recommended actions of event.
	RecommendedActions *EventPropertiesRecommendedActions `json:"recommendedActions,omitempty"`

	// Current status of event.
	Status *EventStatusValues `json:"status,omitempty"`

	// Summary text of event.
	Summary *string `json:"summary,omitempty"`

	// Title text of event.
	Title *string `json:"title,omitempty"`
}

// EventPropertiesAdditionalInformation - Contains additional information regarding the event.
type EventPropertiesAdditionalInformation struct {
	// Additional information message.
	Message *string `json:"message,omitempty"`
}

// EventPropertiesArticle - Article of event.
type EventPropertiesArticle struct {
	// Article content of event.
	ArticleContent *string `json:"articleContent,omitempty"`
}

// EventPropertiesRecommendedActions - Recommended actions of event.
type EventPropertiesRecommendedActions struct {
	// Recommended actions for the service health event.
	Actions []*EventPropertiesRecommendedActionsItem `json:"actions,omitempty"`

	// Recommended action locale for the service health event.
	LocaleCode *string `json:"localeCode,omitempty"`

	// Recommended action title for the service health event.
	Message *string `json:"message,omitempty"`
}

// EventPropertiesRecommendedActionsItem - Recommended action for the service health event.
type EventPropertiesRecommendedActionsItem struct {
	// Recommended action text
	ActionText *string `json:"actionText,omitempty"`

	// Recommended action group Id for the service health event.
	GroupID *int32 `json:"groupId,omitempty"`
}

// Events - The List events operation response.
type Events struct {
	// REQUIRED; The list of event.
	Value []*Event `json:"value,omitempty"`

	// The URI to fetch the next page of events. Call ListNext() with this URI to fetch the next page of events.
	NextLink *string `json:"nextLink,omitempty"`
}

// EventsClientListBySingleResourceOptions contains the optional parameters for the EventsClient.ListBySingleResource method.
type EventsClientListBySingleResourceOptions struct {
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
	// setting view=full expands the article parameters
	View *string
}

// EventsClientListBySubscriptionIDOptions contains the optional parameters for the EventsClient.ListBySubscriptionID method.
type EventsClientListBySubscriptionIDOptions struct {
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
	// Specifies from when to return events, based on the lastUpdateTime property. For example, queryStartTime = 7/24/2020 OR
	// queryStartTime=7%2F24%2F2020
	QueryStartTime *string
	// setting view=full expands the article parameters
	View *string
}

// Faq - Frequently asked question for the service health event
type Faq struct {
	// FAQ answer for the service health event.
	Answer *string `json:"answer,omitempty"`

	// FAQ locale for the service health event.
	LocaleCode *string `json:"localeCode,omitempty"`

	// FAQ question for the service health event.
	Question *string `json:"question,omitempty"`
}

// Impact - Azure service impacted by the service health event.
type Impact struct {
	// List regions impacted by the service health event.
	ImpactedRegions []*ImpactedServiceRegion `json:"impactedRegions,omitempty"`

	// Impacted service name.
	ImpactedService *string `json:"impactedService,omitempty"`
}

// ImpactedRegion - Object of impacted region.
type ImpactedRegion struct {
	// The impacted region id.
	ID *string `json:"id,omitempty"`

	// The impacted region name.
	Name *string `json:"name,omitempty"`
}

// ImpactedResourceListResult - The List impactedResourceStatus operation response.
type ImpactedResourceListResult struct {
	// REQUIRED; The list of impactedResourceStatus.
	Value []*ImpactedResourceStatus `json:"value,omitempty"`

	// The URI to fetch the next page of impactedResourceStatus.
	NextLink *string `json:"nextLink,omitempty"`
}

// ImpactedResourceStatus - impactedResource with health status
type ImpactedResourceStatus struct {
	// Properties of impacted resource status.
	Properties *ImpactedResourceStatusProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ImpactedResourceStatusProperties - Properties of impacted resource status.
type ImpactedResourceStatusProperties struct {
	// Impacted resource status of the resource.
	AvailabilityState *AvailabilityStateValues `json:"availabilityState,omitempty"`

	// Timestamp for when last change in health status occurred.
	OccuredTime *time.Time `json:"occuredTime,omitempty"`

	// When the resource's availabilityState is Unavailable, it describes where the health impacting event was originated.
	ReasonType *ReasonTypeValues `json:"reasonType,omitempty"`

	// Summary description of the impacted resource status.
	Summary *string `json:"summary,omitempty"`

	// Title description of the impacted resource status.
	Title *string `json:"title,omitempty"`
}

// ImpactedResourcesClientListBySubscriptionIDOptions contains the optional parameters for the ImpactedResourcesClient.ListBySubscriptionID
// method.
type ImpactedResourcesClientListBySubscriptionIDOptions struct {
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// ImpactedServiceRegion - Azure region impacted by the service health event.
type ImpactedServiceRegion struct {
	// Impacted region name.
	ImpactedRegion *string `json:"impactedRegion,omitempty"`

	// List subscription impacted by the service health event.
	ImpactedSubscriptions []*string `json:"impactedSubscriptions,omitempty"`

	// It provides the Timestamp for when the last update for the service health event.
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`

	// Current status of event in the region.
	Status *EventStatusValues `json:"status,omitempty"`

	// List of updates for given service health event.
	Updates []*Update `json:"updates,omitempty"`
}

// Link - Useful links for service health event.
type Link struct {
	// It provides the name of portal extension blade to produce link for given service health event.
	BladeName *string `json:"bladeName,omitempty"`

	// Display text of link.
	DisplayText *LinkDisplayText `json:"displayText,omitempty"`

	// It provides the name of portal extension to produce link for given service health event.
	ExtensionName *string `json:"extensionName,omitempty"`

	// It provides a map of parameter name and value for portal extension blade to produce lik for given service health event.
	Parameters interface{} `json:"parameters,omitempty"`

	// Type of link.
	Type *LinkTypeValues `json:"type,omitempty"`
}

// LinkDisplayText - Display text of link.
type LinkDisplayText struct {
	// Localized display text of link.
	LocalizedValue *string `json:"localizedValue,omitempty"`

	// Display text of link.
	Value *string `json:"value,omitempty"`
}

// MetadataClientGetEntityOptions contains the optional parameters for the MetadataClient.GetEntity method.
type MetadataClientGetEntityOptions struct {
	// placeholder for future optional parameters
}

// MetadataClientListOptions contains the optional parameters for the MetadataClient.List method.
type MetadataClientListOptions struct {
	// placeholder for future optional parameters
}

// MetadataEntity - The metadata entity contract.
type MetadataEntity struct {
	// The metadata entity properties.
	Properties *MetadataEntityProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MetadataEntityListResult - The list of metadata entities
type MetadataEntityListResult struct {
	// The link used to get the next page of metadata.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of metadata entities.
	Value []*MetadataEntity `json:"value,omitempty"`
}

// MetadataEntityProperties - The metadata entity properties
type MetadataEntityProperties struct {
	// The list of scenarios applicable to this metadata entity.
	ApplicableScenarios []*Scenario `json:"applicableScenarios,omitempty"`

	// The list of keys on which this entity depends on.
	DependsOn []*string `json:"dependsOn,omitempty"`

	// The display name.
	DisplayName *string `json:"displayName,omitempty"`

	// The list of supported values.
	SupportedValues []*MetadataSupportedValueDetail `json:"supportedValues,omitempty"`
}

// MetadataSupportedValueDetail - The metadata supported value detail.
type MetadataSupportedValueDetail struct {
	// The display name.
	DisplayName *string `json:"displayName,omitempty"`

	// The id.
	ID *string `json:"id,omitempty"`
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

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// RecommendedAction - Lists actions the user can take based on the current availabilityState of the resource.
type RecommendedAction struct {
	// Recommended action.
	Action *string `json:"action,omitempty"`

	// Link to the action
	ActionURL *string `json:"actionUrl,omitempty"`

	// Substring of action, it describes which text should host the action url.
	ActionURLText *string `json:"actionUrlText,omitempty"`
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ServiceImpactingEvent - Lists the service impacting events that may be affecting the health of the resource.
type ServiceImpactingEvent struct {
	// Correlation id for the event
	CorrelationID *string `json:"correlationId,omitempty"`

	// Timestamp for when the event started.
	EventStartTime *time.Time `json:"eventStartTime,omitempty"`

	// Timestamp for when event was submitted/detected.
	EventStatusLastModifiedTime *time.Time `json:"eventStatusLastModifiedTime,omitempty"`

	// Properties of the service impacting event.
	IncidentProperties *ServiceImpactingEventIncidentProperties `json:"incidentProperties,omitempty"`

	// Status of the service impacting event.
	Status *ServiceImpactingEventStatus `json:"status,omitempty"`
}

// ServiceImpactingEventIncidentProperties - Properties of the service impacting event.
type ServiceImpactingEventIncidentProperties struct {
	// Type of Event.
	IncidentType *string `json:"incidentType,omitempty"`

	// Region impacted by the event.
	Region *string `json:"region,omitempty"`

	// Service impacted by the event.
	Service *string `json:"service,omitempty"`

	// Title of the incident.
	Title *string `json:"title,omitempty"`
}

// ServiceImpactingEventStatus - Status of the service impacting event.
type ServiceImpactingEventStatus struct {
	// Current status of the event
	Value *string `json:"value,omitempty"`
}

// StatusActiveEvent - Active event type of emerging issue.
type StatusActiveEvent struct {
	// The cloud type of this active event.
	Cloud *string `json:"cloud,omitempty"`

	// The details of active event.
	Description *string `json:"description,omitempty"`

	// The list of emerging issues impacts.
	Impacts []*EmergingIssueImpact `json:"impacts,omitempty"`

	// The last time modified on this banner.
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`

	// The boolean value of this active event if published or not.
	Published *bool `json:"published,omitempty"`

	// The severity level of this active event.
	Severity *SeverityValues `json:"severity,omitempty"`

	// The stage of this active event.
	Stage *StageValues `json:"stage,omitempty"`

	// The impact start time on this active event.
	StartTime *time.Time `json:"startTime,omitempty"`

	// The active event title.
	Title *string `json:"title,omitempty"`

	// The tracking id of this active event.
	TrackingID *string `json:"trackingId,omitempty"`
}

// StatusBanner - Banner type of emerging issue.
type StatusBanner struct {
	// The cloud type of this banner.
	Cloud *string `json:"cloud,omitempty"`

	// The last time modified on this banner.
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`

	// The details of banner.
	Message *string `json:"message,omitempty"`

	// The banner title.
	Title *string `json:"title,omitempty"`
}

// Update for service health event.
type Update struct {
	// Summary text for the given update for the service health event.
	Summary *string `json:"summary,omitempty"`

	// It provides the Timestamp for the given update for the service health event.
	UpdateDateTime *time.Time `json:"updateDateTime,omitempty"`
}
