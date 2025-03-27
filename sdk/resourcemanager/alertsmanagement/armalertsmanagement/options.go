// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armalertsmanagement

// AlertProcessingRulesClientCreateOrUpdateOptions contains the optional parameters for the AlertProcessingRulesClient.CreateOrUpdate
// method.
type AlertProcessingRulesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// AlertProcessingRulesClientDeleteOptions contains the optional parameters for the AlertProcessingRulesClient.Delete method.
type AlertProcessingRulesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// AlertProcessingRulesClientGetByNameOptions contains the optional parameters for the AlertProcessingRulesClient.GetByName
// method.
type AlertProcessingRulesClientGetByNameOptions struct {
	// placeholder for future optional parameters
}

// AlertProcessingRulesClientListByResourceGroupOptions contains the optional parameters for the AlertProcessingRulesClient.NewListByResourceGroupPager
// method.
type AlertProcessingRulesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// AlertProcessingRulesClientListBySubscriptionOptions contains the optional parameters for the AlertProcessingRulesClient.NewListBySubscriptionPager
// method.
type AlertProcessingRulesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// AlertProcessingRulesClientUpdateOptions contains the optional parameters for the AlertProcessingRulesClient.Update method.
type AlertProcessingRulesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// AlertRuleRecommendationsClientListByResourceOptions contains the optional parameters for the AlertRuleRecommendationsClient.NewListByResourcePager
// method.
type AlertRuleRecommendationsClientListByResourceOptions struct {
	// placeholder for future optional parameters
}

// AlertRuleRecommendationsClientListByTargetTypeOptions contains the optional parameters for the AlertRuleRecommendationsClient.NewListByTargetTypePager
// method.
type AlertRuleRecommendationsClientListByTargetTypeOptions struct {
	// placeholder for future optional parameters
}

// AlertsClientChangeStateOptions contains the optional parameters for the AlertsClient.ChangeState method.
type AlertsClientChangeStateOptions struct {
	// reason of change alert state
	Comment *Comments
}

// AlertsClientGetAllOptions contains the optional parameters for the AlertsClient.NewGetAllPager method.
type AlertsClientGetAllOptions struct {
	// Filter by specific alert rule. Default value is to select all.
	AlertRule *string

	// Filter by state of the alert instance. Default value is to select all.
	AlertState *AlertState

	// Filter by custom time range in the format / where time is in (ISO-8601 format)'. Permissible values is within 30 days from
	// query time. Either timeRange or customTimeRange could be used but not both.
	// Default is none.
	CustomTimeRange *string

	// Include context which has contextual data specific to the monitor service. Default value is false'
	IncludeContext *bool

	// Include egress config which would be used for displaying the content in portal. Default value is 'false'.
	IncludeEgressConfig *bool

	// Filter by monitor condition which is either 'Fired' or 'Resolved'. Default value is to select all.
	MonitorCondition *MonitorCondition

	// Filter by monitor service which generates the alert instance. Default value is select all.
	MonitorService *MonitorService

	// Determines number of alerts returned per page in response. Permissible value is between 1 to 250. When the "includeContent"
	// filter is selected, maximum value allowed is 25. Default value is 25.
	PageCount *int64

	// This filter allows to selection of the fields(comma separated) which would be part of the essential section. This would
	// allow to project only the required fields rather than getting entire content.
	// Default is to fetch all the fields in the essentials section.
	Select *string

	// Filter by severity. Default value is select all.
	Severity *Severity

	// Filter the alerts list by the Smart Group Id. Default value is none.
	SmartGroupID *string

	// Sort the query results by input field, Default value is 'lastModifiedDateTime'.
	SortBy *AlertsSortByFields

	// Sort the query results order in either ascending or descending. Default value is 'desc' for time fields and 'asc' for others.
	SortOrder *SortOrder

	// Filter by target resource( which is full ARM ID) Default value is select all.
	TargetResource *string

	// Filter by target resource group name. Default value is select all.
	TargetResourceGroup *string

	// Filter by target resource type. Default value is select all.
	TargetResourceType *string

	// Filter by time range by below listed values. Default value is 1 day.
	TimeRange *TimeRange
}

// AlertsClientGetByIDOptions contains the optional parameters for the AlertsClient.GetByID method.
type AlertsClientGetByIDOptions struct {
	// placeholder for future optional parameters
}

// AlertsClientGetEnrichmentsOptions contains the optional parameters for the AlertsClient.GetEnrichments method.
type AlertsClientGetEnrichmentsOptions struct {
	// placeholder for future optional parameters
}

// AlertsClientGetHistoryOptions contains the optional parameters for the AlertsClient.GetHistory method.
type AlertsClientGetHistoryOptions struct {
	// placeholder for future optional parameters
}

// AlertsClientGetSummaryOptions contains the optional parameters for the AlertsClient.GetSummary method.
type AlertsClientGetSummaryOptions struct {
	// Filter by specific alert rule. Default value is to select all.
	AlertRule *string

	// Filter by state of the alert instance. Default value is to select all.
	AlertState *AlertState

	// Filter by custom time range in the format / where time is in (ISO-8601 format)'. Permissible values is within 30 days from
	// query time. Either timeRange or customTimeRange could be used but not both.
	// Default is none.
	CustomTimeRange *string

	// Include count of the SmartGroups as part of the summary. Default value is 'false'.
	IncludeSmartGroupsCount *bool

	// Filter by monitor condition which is either 'Fired' or 'Resolved'. Default value is to select all.
	MonitorCondition *MonitorCondition

	// Filter by monitor service which generates the alert instance. Default value is select all.
	MonitorService *MonitorService

	// Filter by severity. Default value is select all.
	Severity *Severity

	// Filter by target resource( which is full ARM ID) Default value is select all.
	TargetResource *string

	// Filter by target resource group name. Default value is select all.
	TargetResourceGroup *string

	// Filter by target resource type. Default value is select all.
	TargetResourceType *string

	// Filter by time range by below listed values. Default value is 1 day.
	TimeRange *TimeRange
}

// AlertsClientListEnrichmentsOptions contains the optional parameters for the AlertsClient.NewListEnrichmentsPager method.
type AlertsClientListEnrichmentsOptions struct {
	// placeholder for future optional parameters
}

// AlertsClientMetaDataOptions contains the optional parameters for the AlertsClient.MetaData method.
type AlertsClientMetaDataOptions struct {
	// placeholder for future optional parameters
}

// IssueClientAddOrUpdateAlertsOptions contains the optional parameters for the IssueClient.AddOrUpdateAlerts method.
type IssueClientAddOrUpdateAlertsOptions struct {
	// placeholder for future optional parameters
}

// IssueClientAddOrUpdateResourcesOptions contains the optional parameters for the IssueClient.AddOrUpdateResources method.
type IssueClientAddOrUpdateResourcesOptions struct {
	// placeholder for future optional parameters
}

// IssueClientCreateOptions contains the optional parameters for the IssueClient.Create method.
type IssueClientCreateOptions struct {
	// placeholder for future optional parameters
}

// IssueClientDeleteOptions contains the optional parameters for the IssueClient.Delete method.
type IssueClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// IssueClientFetchInvestigationResultOptions contains the optional parameters for the IssueClient.FetchInvestigationResult
// method.
type IssueClientFetchInvestigationResultOptions struct {
	// placeholder for future optional parameters
}

// IssueClientGetOptions contains the optional parameters for the IssueClient.Get method.
type IssueClientGetOptions struct {
	// placeholder for future optional parameters
}

// IssueClientListAlertsOptions contains the optional parameters for the IssueClient.NewListAlertsPager method.
type IssueClientListAlertsOptions struct {
	// placeholder for future optional parameters
}

// IssueClientListOptions contains the optional parameters for the IssueClient.NewListPager method.
type IssueClientListOptions struct {
	// placeholder for future optional parameters
}

// IssueClientListResourcesOptions contains the optional parameters for the IssueClient.NewListResourcesPager method.
type IssueClientListResourcesOptions struct {
	// placeholder for future optional parameters
}

// IssueClientStartInvestigationOptions contains the optional parameters for the IssueClient.StartInvestigation method.
type IssueClientStartInvestigationOptions struct {
	// placeholder for future optional parameters
}

// IssueClientUpdateOptions contains the optional parameters for the IssueClient.Update method.
type IssueClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrometheusRuleGroupsClientCreateOrUpdateOptions contains the optional parameters for the PrometheusRuleGroupsClient.CreateOrUpdate
// method.
type PrometheusRuleGroupsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// PrometheusRuleGroupsClientDeleteOptions contains the optional parameters for the PrometheusRuleGroupsClient.Delete method.
type PrometheusRuleGroupsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrometheusRuleGroupsClientGetOptions contains the optional parameters for the PrometheusRuleGroupsClient.Get method.
type PrometheusRuleGroupsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrometheusRuleGroupsClientListByResourceGroupOptions contains the optional parameters for the PrometheusRuleGroupsClient.NewListByResourceGroupPager
// method.
type PrometheusRuleGroupsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// PrometheusRuleGroupsClientListBySubscriptionOptions contains the optional parameters for the PrometheusRuleGroupsClient.NewListBySubscriptionPager
// method.
type PrometheusRuleGroupsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// PrometheusRuleGroupsClientUpdateOptions contains the optional parameters for the PrometheusRuleGroupsClient.Update method.
type PrometheusRuleGroupsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// SmartGroupsClientChangeStateOptions contains the optional parameters for the SmartGroupsClient.ChangeState method.
type SmartGroupsClientChangeStateOptions struct {
	// placeholder for future optional parameters
}

// SmartGroupsClientGetAllOptions contains the optional parameters for the SmartGroupsClient.NewGetAllPager method.
type SmartGroupsClientGetAllOptions struct {
	// Filter by monitor condition which is either 'Fired' or 'Resolved'. Default value is to select all.
	MonitorCondition *MonitorCondition

	// Filter by monitor service which generates the alert instance. Default value is select all.
	MonitorService *MonitorService

	// Determines number of alerts returned per page in response. Permissible value is between 1 to 250. When the "includeContent"
	// filter is selected, maximum value allowed is 25. Default value is 25.
	PageCount *int64

	// Filter by severity. Default value is select all.
	Severity *Severity

	// Filter by state of the smart group. Default value is to select all.
	SmartGroupState *AlertState

	// Sort the query results by input field. Default value is sort by 'lastModifiedDateTime'.
	SortBy *SmartGroupsSortByFields

	// Sort the query results order in either ascending or descending. Default value is 'desc' for time fields and 'asc' for others.
	SortOrder *SortOrder

	// Filter by target resource( which is full ARM ID) Default value is select all.
	TargetResource *string

	// Filter by target resource group name. Default value is select all.
	TargetResourceGroup *string

	// Filter by target resource type. Default value is select all.
	TargetResourceType *string

	// Filter by time range by below listed values. Default value is 1 day.
	TimeRange *TimeRange
}

// SmartGroupsClientGetByIDOptions contains the optional parameters for the SmartGroupsClient.GetByID method.
type SmartGroupsClientGetByIDOptions struct {
	// placeholder for future optional parameters
}

// SmartGroupsClientGetHistoryOptions contains the optional parameters for the SmartGroupsClient.GetHistory method.
type SmartGroupsClientGetHistoryOptions struct {
	// placeholder for future optional parameters
}
