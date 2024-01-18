//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

// ActionGroupsClientCreateNotificationsAtActionGroupResourceLevelResponse contains the response from method ActionGroupsClient.BeginCreateNotificationsAtActionGroupResourceLevel.
type ActionGroupsClientCreateNotificationsAtActionGroupResourceLevelResponse struct {
	// The details of the test notification results.
	TestNotificationDetailsResponse
}

// ActionGroupsClientCreateOrUpdateResponse contains the response from method ActionGroupsClient.CreateOrUpdate.
type ActionGroupsClientCreateOrUpdateResponse struct {
	// An action group resource.
	ActionGroupResource
}

// ActionGroupsClientDeleteResponse contains the response from method ActionGroupsClient.Delete.
type ActionGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// ActionGroupsClientEnableReceiverResponse contains the response from method ActionGroupsClient.EnableReceiver.
type ActionGroupsClientEnableReceiverResponse struct {
	// placeholder for future response values
}

// ActionGroupsClientGetResponse contains the response from method ActionGroupsClient.Get.
type ActionGroupsClientGetResponse struct {
	// An action group resource.
	ActionGroupResource
}

// ActionGroupsClientGetTestNotificationsAtActionGroupResourceLevelResponse contains the response from method ActionGroupsClient.GetTestNotificationsAtActionGroupResourceLevel.
type ActionGroupsClientGetTestNotificationsAtActionGroupResourceLevelResponse struct {
	// The details of the test notification results.
	TestNotificationDetailsResponse
}

// ActionGroupsClientListByResourceGroupResponse contains the response from method ActionGroupsClient.NewListByResourceGroupPager.
type ActionGroupsClientListByResourceGroupResponse struct {
	// A list of action groups.
	ActionGroupList
}

// ActionGroupsClientListBySubscriptionIDResponse contains the response from method ActionGroupsClient.NewListBySubscriptionIDPager.
type ActionGroupsClientListBySubscriptionIDResponse struct {
	// A list of action groups.
	ActionGroupList
}

// ActionGroupsClientUpdateResponse contains the response from method ActionGroupsClient.Update.
type ActionGroupsClientUpdateResponse struct {
	// An action group resource.
	ActionGroupResource
}

// ActivityLogAlertsClientCreateOrUpdateResponse contains the response from method ActivityLogAlertsClient.CreateOrUpdate.
type ActivityLogAlertsClientCreateOrUpdateResponse struct {
	// An Activity Log Alert rule resource.
	ActivityLogAlertResource
}

// ActivityLogAlertsClientDeleteResponse contains the response from method ActivityLogAlertsClient.Delete.
type ActivityLogAlertsClientDeleteResponse struct {
	// placeholder for future response values
}

// ActivityLogAlertsClientGetResponse contains the response from method ActivityLogAlertsClient.Get.
type ActivityLogAlertsClientGetResponse struct {
	// An Activity Log Alert rule resource.
	ActivityLogAlertResource
}

// ActivityLogAlertsClientListByResourceGroupResponse contains the response from method ActivityLogAlertsClient.NewListByResourceGroupPager.
type ActivityLogAlertsClientListByResourceGroupResponse struct {
	// A list of Activity Log Alert rules.
	AlertRuleList
}

// ActivityLogAlertsClientListBySubscriptionIDResponse contains the response from method ActivityLogAlertsClient.NewListBySubscriptionIDPager.
type ActivityLogAlertsClientListBySubscriptionIDResponse struct {
	// A list of Activity Log Alert rules.
	AlertRuleList
}

// ActivityLogAlertsClientUpdateResponse contains the response from method ActivityLogAlertsClient.Update.
type ActivityLogAlertsClientUpdateResponse struct {
	// An Activity Log Alert rule resource.
	ActivityLogAlertResource
}

// ActivityLogsClientListResponse contains the response from method ActivityLogsClient.NewListPager.
type ActivityLogsClientListResponse struct {
	// Represents collection of events.
	EventDataCollection
}

// AlertRuleIncidentsClientGetResponse contains the response from method AlertRuleIncidentsClient.Get.
type AlertRuleIncidentsClientGetResponse struct {
	// An alert incident indicates the activation status of an alert rule.
	Incident
}

// AlertRuleIncidentsClientListByAlertRuleResponse contains the response from method AlertRuleIncidentsClient.NewListByAlertRulePager.
type AlertRuleIncidentsClientListByAlertRuleResponse struct {
	// The List incidents operation response.
	IncidentListResult
}

// AlertRulesClientCreateOrUpdateResponse contains the response from method AlertRulesClient.CreateOrUpdate.
type AlertRulesClientCreateOrUpdateResponse struct {
	// The alert rule resource.
	AlertRuleResource
}

// AlertRulesClientDeleteResponse contains the response from method AlertRulesClient.Delete.
type AlertRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// AlertRulesClientGetResponse contains the response from method AlertRulesClient.Get.
type AlertRulesClientGetResponse struct {
	// The alert rule resource.
	AlertRuleResource
}

// AlertRulesClientListByResourceGroupResponse contains the response from method AlertRulesClient.NewListByResourceGroupPager.
type AlertRulesClientListByResourceGroupResponse struct {
	// Represents a collection of alert rule resources.
	AlertRuleResourceCollection
}

// AlertRulesClientListBySubscriptionResponse contains the response from method AlertRulesClient.NewListBySubscriptionPager.
type AlertRulesClientListBySubscriptionResponse struct {
	// Represents a collection of alert rule resources.
	AlertRuleResourceCollection
}

// AlertRulesClientUpdateResponse contains the response from method AlertRulesClient.Update.
type AlertRulesClientUpdateResponse struct {
	// The alert rule resource.
	AlertRuleResource
}

// AutoscaleSettingsClientCreateOrUpdateResponse contains the response from method AutoscaleSettingsClient.CreateOrUpdate.
type AutoscaleSettingsClientCreateOrUpdateResponse struct {
	// The autoscale setting resource.
	AutoscaleSettingResource
}

// AutoscaleSettingsClientDeleteResponse contains the response from method AutoscaleSettingsClient.Delete.
type AutoscaleSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// AutoscaleSettingsClientGetResponse contains the response from method AutoscaleSettingsClient.Get.
type AutoscaleSettingsClientGetResponse struct {
	// The autoscale setting resource.
	AutoscaleSettingResource
}

// AutoscaleSettingsClientListByResourceGroupResponse contains the response from method AutoscaleSettingsClient.NewListByResourceGroupPager.
type AutoscaleSettingsClientListByResourceGroupResponse struct {
	// Represents a collection of autoscale setting resources.
	AutoscaleSettingResourceCollection
}

// AutoscaleSettingsClientListBySubscriptionResponse contains the response from method AutoscaleSettingsClient.NewListBySubscriptionPager.
type AutoscaleSettingsClientListBySubscriptionResponse struct {
	// Represents a collection of autoscale setting resources.
	AutoscaleSettingResourceCollection
}

// AutoscaleSettingsClientUpdateResponse contains the response from method AutoscaleSettingsClient.Update.
type AutoscaleSettingsClientUpdateResponse struct {
	// The autoscale setting resource.
	AutoscaleSettingResource
}

// AzureMonitorWorkspacesClientCreateResponse contains the response from method AzureMonitorWorkspacesClient.Create.
type AzureMonitorWorkspacesClientCreateResponse struct {
	// An Azure Monitor Workspace definition
	AzureMonitorWorkspaceResource
}

// AzureMonitorWorkspacesClientDeleteResponse contains the response from method AzureMonitorWorkspacesClient.BeginDelete.
type AzureMonitorWorkspacesClientDeleteResponse struct {
	// placeholder for future response values
}

// AzureMonitorWorkspacesClientGetResponse contains the response from method AzureMonitorWorkspacesClient.Get.
type AzureMonitorWorkspacesClientGetResponse struct {
	// An Azure Monitor Workspace definition
	AzureMonitorWorkspaceResource
}

// AzureMonitorWorkspacesClientListByResourceGroupResponse contains the response from method AzureMonitorWorkspacesClient.NewListByResourceGroupPager.
type AzureMonitorWorkspacesClientListByResourceGroupResponse struct {
	// A pageable list of resources
	AzureMonitorWorkspaceResourceListResult
}

// AzureMonitorWorkspacesClientListBySubscriptionResponse contains the response from method AzureMonitorWorkspacesClient.NewListBySubscriptionPager.
type AzureMonitorWorkspacesClientListBySubscriptionResponse struct {
	// A pageable list of resources
	AzureMonitorWorkspaceResourceListResult
}

// AzureMonitorWorkspacesClientUpdateResponse contains the response from method AzureMonitorWorkspacesClient.Update.
type AzureMonitorWorkspacesClientUpdateResponse struct {
	// An Azure Monitor Workspace definition
	AzureMonitorWorkspaceResource
}

// BaselinesClientListResponse contains the response from method BaselinesClient.NewListPager.
type BaselinesClientListResponse struct {
	// A list of metric baselines.
	MetricBaselinesResponse
}

// ClientCreateNotificationsAtTenantActionGroupResourceLevelResponse contains the response from method Client.BeginCreateNotificationsAtTenantActionGroupResourceLevel.
type ClientCreateNotificationsAtTenantActionGroupResourceLevelResponse struct {
	// The details of the test notification results.
	TestNotificationDetailsResponseAutoGenerated
}

// ClientGetTestNotificationsAtTenantActionGroupResourceLevelResponse contains the response from method Client.GetTestNotificationsAtTenantActionGroupResourceLevel.
type ClientGetTestNotificationsAtTenantActionGroupResourceLevelResponse struct {
	// The details of the test notification results.
	TestNotificationDetailsResponseAutoGenerated
}

// DataCollectionEndpointsClientCreateResponse contains the response from method DataCollectionEndpointsClient.Create.
type DataCollectionEndpointsClientCreateResponse struct {
	// Definition of ARM tracked top level resource.
	DataCollectionEndpointResource
}

// DataCollectionEndpointsClientDeleteResponse contains the response from method DataCollectionEndpointsClient.Delete.
type DataCollectionEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// DataCollectionEndpointsClientGetResponse contains the response from method DataCollectionEndpointsClient.Get.
type DataCollectionEndpointsClientGetResponse struct {
	// Definition of ARM tracked top level resource.
	DataCollectionEndpointResource
}

// DataCollectionEndpointsClientListByResourceGroupResponse contains the response from method DataCollectionEndpointsClient.NewListByResourceGroupPager.
type DataCollectionEndpointsClientListByResourceGroupResponse struct {
	// A pageable list of resources.
	DataCollectionEndpointResourceListResult
}

// DataCollectionEndpointsClientListBySubscriptionResponse contains the response from method DataCollectionEndpointsClient.NewListBySubscriptionPager.
type DataCollectionEndpointsClientListBySubscriptionResponse struct {
	// A pageable list of resources.
	DataCollectionEndpointResourceListResult
}

// DataCollectionEndpointsClientUpdateResponse contains the response from method DataCollectionEndpointsClient.Update.
type DataCollectionEndpointsClientUpdateResponse struct {
	// Definition of ARM tracked top level resource.
	DataCollectionEndpointResource
}

// DataCollectionRuleAssociationsClientCreateResponse contains the response from method DataCollectionRuleAssociationsClient.Create.
type DataCollectionRuleAssociationsClientCreateResponse struct {
	// Definition of generic ARM proxy resource.
	DataCollectionRuleAssociationProxyOnlyResource
}

// DataCollectionRuleAssociationsClientDeleteResponse contains the response from method DataCollectionRuleAssociationsClient.Delete.
type DataCollectionRuleAssociationsClientDeleteResponse struct {
	// placeholder for future response values
}

// DataCollectionRuleAssociationsClientGetResponse contains the response from method DataCollectionRuleAssociationsClient.Get.
type DataCollectionRuleAssociationsClientGetResponse struct {
	// Definition of generic ARM proxy resource.
	DataCollectionRuleAssociationProxyOnlyResource
}

// DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse contains the response from method DataCollectionRuleAssociationsClient.NewListByDataCollectionEndpointPager.
type DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse struct {
	// A pageable list of resources.
	DataCollectionRuleAssociationProxyOnlyResourceListResult
}

// DataCollectionRuleAssociationsClientListByResourceResponse contains the response from method DataCollectionRuleAssociationsClient.NewListByResourcePager.
type DataCollectionRuleAssociationsClientListByResourceResponse struct {
	// A pageable list of resources.
	DataCollectionRuleAssociationProxyOnlyResourceListResult
}

// DataCollectionRuleAssociationsClientListByRuleResponse contains the response from method DataCollectionRuleAssociationsClient.NewListByRulePager.
type DataCollectionRuleAssociationsClientListByRuleResponse struct {
	// A pageable list of resources.
	DataCollectionRuleAssociationProxyOnlyResourceListResult
}

// DataCollectionRulesClientCreateResponse contains the response from method DataCollectionRulesClient.Create.
type DataCollectionRulesClientCreateResponse struct {
	// Definition of ARM tracked top level resource.
	DataCollectionRuleResource
}

// DataCollectionRulesClientDeleteResponse contains the response from method DataCollectionRulesClient.Delete.
type DataCollectionRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// DataCollectionRulesClientGetResponse contains the response from method DataCollectionRulesClient.Get.
type DataCollectionRulesClientGetResponse struct {
	// Definition of ARM tracked top level resource.
	DataCollectionRuleResource
}

// DataCollectionRulesClientListByResourceGroupResponse contains the response from method DataCollectionRulesClient.NewListByResourceGroupPager.
type DataCollectionRulesClientListByResourceGroupResponse struct {
	// A pageable list of resources.
	DataCollectionRuleResourceListResult
}

// DataCollectionRulesClientListBySubscriptionResponse contains the response from method DataCollectionRulesClient.NewListBySubscriptionPager.
type DataCollectionRulesClientListBySubscriptionResponse struct {
	// A pageable list of resources.
	DataCollectionRuleResourceListResult
}

// DataCollectionRulesClientUpdateResponse contains the response from method DataCollectionRulesClient.Update.
type DataCollectionRulesClientUpdateResponse struct {
	// Definition of ARM tracked top level resource.
	DataCollectionRuleResource
}

// DiagnosticSettingsCategoryClientGetResponse contains the response from method DiagnosticSettingsCategoryClient.Get.
type DiagnosticSettingsCategoryClientGetResponse struct {
	// The diagnostic settings category resource.
	DiagnosticSettingsCategoryResource
}

// DiagnosticSettingsCategoryClientListResponse contains the response from method DiagnosticSettingsCategoryClient.NewListPager.
type DiagnosticSettingsCategoryClientListResponse struct {
	// Represents a collection of diagnostic setting category resources.
	DiagnosticSettingsCategoryResourceCollection
}

// DiagnosticSettingsClientCreateOrUpdateResponse contains the response from method DiagnosticSettingsClient.CreateOrUpdate.
type DiagnosticSettingsClientCreateOrUpdateResponse struct {
	// The diagnostic setting resource.
	DiagnosticSettingsResource
}

// DiagnosticSettingsClientDeleteResponse contains the response from method DiagnosticSettingsClient.Delete.
type DiagnosticSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// DiagnosticSettingsClientGetResponse contains the response from method DiagnosticSettingsClient.Get.
type DiagnosticSettingsClientGetResponse struct {
	// The diagnostic setting resource.
	DiagnosticSettingsResource
}

// DiagnosticSettingsClientListResponse contains the response from method DiagnosticSettingsClient.NewListPager.
type DiagnosticSettingsClientListResponse struct {
	// Represents a collection of alert rule resources.
	DiagnosticSettingsResourceCollection
}

// EventCategoriesClientListResponse contains the response from method EventCategoriesClient.NewListPager.
type EventCategoriesClientListResponse struct {
	// A collection of event categories. Currently possible values are: Administrative, Security, ServiceHealth, Alert, Recommendation,
	// Policy.
	EventCategoryCollection
}

// LogProfilesClientCreateOrUpdateResponse contains the response from method LogProfilesClient.CreateOrUpdate.
type LogProfilesClientCreateOrUpdateResponse struct {
	// The log profile resource.
	LogProfileResource
}

// LogProfilesClientDeleteResponse contains the response from method LogProfilesClient.Delete.
type LogProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// LogProfilesClientGetResponse contains the response from method LogProfilesClient.Get.
type LogProfilesClientGetResponse struct {
	// The log profile resource.
	LogProfileResource
}

// LogProfilesClientListResponse contains the response from method LogProfilesClient.NewListPager.
type LogProfilesClientListResponse struct {
	// Represents a collection of log profiles.
	LogProfileCollection
}

// LogProfilesClientUpdateResponse contains the response from method LogProfilesClient.Update.
type LogProfilesClientUpdateResponse struct {
	// The log profile resource.
	LogProfileResource
}

// MetricAlertsClientCreateOrUpdateResponse contains the response from method MetricAlertsClient.CreateOrUpdate.
type MetricAlertsClientCreateOrUpdateResponse struct {
	// The metric alert resource.
	MetricAlertResource
}

// MetricAlertsClientDeleteResponse contains the response from method MetricAlertsClient.Delete.
type MetricAlertsClientDeleteResponse struct {
	// placeholder for future response values
}

// MetricAlertsClientGetResponse contains the response from method MetricAlertsClient.Get.
type MetricAlertsClientGetResponse struct {
	// The metric alert resource.
	MetricAlertResource
}

// MetricAlertsClientListByResourceGroupResponse contains the response from method MetricAlertsClient.NewListByResourceGroupPager.
type MetricAlertsClientListByResourceGroupResponse struct {
	// Represents a collection of alert rule resources.
	MetricAlertResourceCollection
}

// MetricAlertsClientListBySubscriptionResponse contains the response from method MetricAlertsClient.NewListBySubscriptionPager.
type MetricAlertsClientListBySubscriptionResponse struct {
	// Represents a collection of alert rule resources.
	MetricAlertResourceCollection
}

// MetricAlertsClientUpdateResponse contains the response from method MetricAlertsClient.Update.
type MetricAlertsClientUpdateResponse struct {
	// The metric alert resource.
	MetricAlertResource
}

// MetricAlertsStatusClientListByNameResponse contains the response from method MetricAlertsStatusClient.ListByName.
type MetricAlertsStatusClientListByNameResponse struct {
	// Represents a collection of alert rule resources.
	MetricAlertStatusCollection
}

// MetricAlertsStatusClientListResponse contains the response from method MetricAlertsStatusClient.List.
type MetricAlertsStatusClientListResponse struct {
	// Represents a collection of alert rule resources.
	MetricAlertStatusCollection
}

// MetricDefinitionsClientListAtSubscriptionScopeResponse contains the response from method MetricDefinitionsClient.NewListAtSubscriptionScopePager.
type MetricDefinitionsClientListAtSubscriptionScopeResponse struct {
	// Represents collection of metric definitions.
	SubscriptionScopeMetricDefinitionCollection
}

// MetricDefinitionsClientListResponse contains the response from method MetricDefinitionsClient.NewListPager.
type MetricDefinitionsClientListResponse struct {
	// Represents collection of metric definitions.
	MetricDefinitionCollection
}

// MetricNamespacesClientListResponse contains the response from method MetricNamespacesClient.NewListPager.
type MetricNamespacesClientListResponse struct {
	// Represents collection of metric namespaces.
	MetricNamespaceCollection
}

// MetricsClientListAtSubscriptionScopePostResponse contains the response from method MetricsClient.ListAtSubscriptionScopePost.
type MetricsClientListAtSubscriptionScopePostResponse struct {
	// The response to a metrics query.
	Response
}

// MetricsClientListAtSubscriptionScopeResponse contains the response from method MetricsClient.ListAtSubscriptionScope.
type MetricsClientListAtSubscriptionScopeResponse struct {
	// The response to a metrics query.
	Response
}

// MetricsClientListResponse contains the response from method MetricsClient.List.
type MetricsClientListResponse struct {
	// The response to a metrics query.
	Response
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	// Result of the request to list Microsoft.Insights operations. It contains a list of operations and a URL link to get the
	// next set of results.
	OperationListResultAutoGenerated
}

// OperationsForMonitorClientListResponse contains the response from method OperationsForMonitorClient.NewListPager.
type OperationsForMonitorClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// PredictiveMetricClientGetResponse contains the response from method PredictiveMetricClient.Get.
type PredictiveMetricClientGetResponse struct {
	// The response to a metrics query.
	PredictiveResponse
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	// The Private Endpoint Connection resource.
	PrivateEndpointConnectionAutoGenerated
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// The Private Endpoint Connection resource.
	PrivateEndpointConnectionAutoGenerated
}

// PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse contains the response from method PrivateEndpointConnectionsClient.ListByPrivateLinkScope.
type PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse struct {
	// List of private endpoint connection associated with the specified storage account
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	// A private link resource
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByPrivateLinkScopeResponse contains the response from method PrivateLinkResourcesClient.ListByPrivateLinkScope.
type PrivateLinkResourcesClientListByPrivateLinkScopeResponse struct {
	// A list of private link resources
	PrivateLinkResourceListResult
}

// PrivateLinkScopeOperationStatusClientGetResponse contains the response from method PrivateLinkScopeOperationStatusClient.Get.
type PrivateLinkScopeOperationStatusClientGetResponse struct {
	// The status of operation.
	OperationStatus
}

// PrivateLinkScopedResourcesClientCreateOrUpdateResponse contains the response from method PrivateLinkScopedResourcesClient.BeginCreateOrUpdate.
type PrivateLinkScopedResourcesClientCreateOrUpdateResponse struct {
	// A private link scoped resource
	ScopedResource
}

// PrivateLinkScopedResourcesClientDeleteResponse contains the response from method PrivateLinkScopedResourcesClient.BeginDelete.
type PrivateLinkScopedResourcesClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateLinkScopedResourcesClientGetResponse contains the response from method PrivateLinkScopedResourcesClient.Get.
type PrivateLinkScopedResourcesClientGetResponse struct {
	// A private link scoped resource
	ScopedResource
}

// PrivateLinkScopedResourcesClientListByPrivateLinkScopeResponse contains the response from method PrivateLinkScopedResourcesClient.NewListByPrivateLinkScopePager.
type PrivateLinkScopedResourcesClientListByPrivateLinkScopeResponse struct {
	// A list of scoped resources in a private link scope.
	ScopedResourceListResult
}

// PrivateLinkScopesClientCreateOrUpdateResponse contains the response from method PrivateLinkScopesClient.CreateOrUpdate.
type PrivateLinkScopesClientCreateOrUpdateResponse struct {
	// An Azure Monitor PrivateLinkScope definition.
	AzureMonitorPrivateLinkScope
}

// PrivateLinkScopesClientDeleteResponse contains the response from method PrivateLinkScopesClient.BeginDelete.
type PrivateLinkScopesClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateLinkScopesClientGetResponse contains the response from method PrivateLinkScopesClient.Get.
type PrivateLinkScopesClientGetResponse struct {
	// An Azure Monitor PrivateLinkScope definition.
	AzureMonitorPrivateLinkScope
}

// PrivateLinkScopesClientListByResourceGroupResponse contains the response from method PrivateLinkScopesClient.NewListByResourceGroupPager.
type PrivateLinkScopesClientListByResourceGroupResponse struct {
	// Describes the list of Azure Monitor PrivateLinkScope resources.
	AzureMonitorPrivateLinkScopeListResult
}

// PrivateLinkScopesClientListResponse contains the response from method PrivateLinkScopesClient.NewListPager.
type PrivateLinkScopesClientListResponse struct {
	// Describes the list of Azure Monitor PrivateLinkScope resources.
	AzureMonitorPrivateLinkScopeListResult
}

// PrivateLinkScopesClientUpdateTagsResponse contains the response from method PrivateLinkScopesClient.UpdateTags.
type PrivateLinkScopesClientUpdateTagsResponse struct {
	// An Azure Monitor PrivateLinkScope definition.
	AzureMonitorPrivateLinkScope
}

// ScheduledQueryRulesClientCreateOrUpdateResponse contains the response from method ScheduledQueryRulesClient.CreateOrUpdate.
type ScheduledQueryRulesClientCreateOrUpdateResponse struct {
	// The scheduled query rule resource.
	ScheduledQueryRuleResource
}

// ScheduledQueryRulesClientDeleteResponse contains the response from method ScheduledQueryRulesClient.Delete.
type ScheduledQueryRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// ScheduledQueryRulesClientGetResponse contains the response from method ScheduledQueryRulesClient.Get.
type ScheduledQueryRulesClientGetResponse struct {
	// The scheduled query rule resource.
	ScheduledQueryRuleResource
}

// ScheduledQueryRulesClientListByResourceGroupResponse contains the response from method ScheduledQueryRulesClient.NewListByResourceGroupPager.
type ScheduledQueryRulesClientListByResourceGroupResponse struct {
	// Represents a collection of scheduled query rule resources.
	ScheduledQueryRuleResourceCollection
}

// ScheduledQueryRulesClientListBySubscriptionResponse contains the response from method ScheduledQueryRulesClient.NewListBySubscriptionPager.
type ScheduledQueryRulesClientListBySubscriptionResponse struct {
	// Represents a collection of scheduled query rule resources.
	ScheduledQueryRuleResourceCollection
}

// ScheduledQueryRulesClientUpdateResponse contains the response from method ScheduledQueryRulesClient.Update.
type ScheduledQueryRulesClientUpdateResponse struct {
	// The scheduled query rule resource.
	ScheduledQueryRuleResource
}

// TenantActionGroupsClientCreateOrUpdateResponse contains the response from method TenantActionGroupsClient.CreateOrUpdate.
type TenantActionGroupsClientCreateOrUpdateResponse struct {
	// A tenant action group resource.
	TenantActionGroupResource
}

// TenantActionGroupsClientDeleteResponse contains the response from method TenantActionGroupsClient.Delete.
type TenantActionGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// TenantActionGroupsClientGetResponse contains the response from method TenantActionGroupsClient.Get.
type TenantActionGroupsClientGetResponse struct {
	// A tenant action group resource.
	TenantActionGroupResource
}

// TenantActionGroupsClientListByManagementGroupIDResponse contains the response from method TenantActionGroupsClient.NewListByManagementGroupIDPager.
type TenantActionGroupsClientListByManagementGroupIDResponse struct {
	// A list of tenant action groups.
	TenantActionGroupList
}

// TenantActionGroupsClientUpdateResponse contains the response from method TenantActionGroupsClient.Update.
type TenantActionGroupsClientUpdateResponse struct {
	// A tenant action group resource.
	TenantActionGroupResource
}

// TenantActivityLogsClientListResponse contains the response from method TenantActivityLogsClient.NewListPager.
type TenantActivityLogsClientListResponse struct {
	// Represents collection of events.
	EventDataCollection
}

// VMInsightsClientGetOnboardingStatusResponse contains the response from method VMInsightsClient.GetOnboardingStatus.
type VMInsightsClientGetOnboardingStatusResponse struct {
	// VM Insights onboarding status for a resource.
	VMInsightsOnboardingStatus
}
