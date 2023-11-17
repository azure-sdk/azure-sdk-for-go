//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

// ActionsClientCreateOrUpdateOptions contains the optional parameters for the ActionsClient.CreateOrUpdate method.
type ActionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ActionsClientDeleteOptions contains the optional parameters for the ActionsClient.Delete method.
type ActionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ActionsClientGetOptions contains the optional parameters for the ActionsClient.Get method.
type ActionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ActionsClientListByAlertRuleOptions contains the optional parameters for the ActionsClient.NewListByAlertRulePager method.
type ActionsClientListByAlertRuleOptions struct {
	// placeholder for future optional parameters
}

// AlertRuleTemplatesClientGetOptions contains the optional parameters for the AlertRuleTemplatesClient.Get method.
type AlertRuleTemplatesClientGetOptions struct {
	// placeholder for future optional parameters
}

// AlertRuleTemplatesClientListOptions contains the optional parameters for the AlertRuleTemplatesClient.NewListPager method.
type AlertRuleTemplatesClientListOptions struct {
	// placeholder for future optional parameters
}

// AlertRulesClientCreateOrUpdateOptions contains the optional parameters for the AlertRulesClient.CreateOrUpdate method.
type AlertRulesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// AlertRulesClientDeleteOptions contains the optional parameters for the AlertRulesClient.Delete method.
type AlertRulesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// AlertRulesClientGetOptions contains the optional parameters for the AlertRulesClient.Get method.
type AlertRulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// AlertRulesClientListOptions contains the optional parameters for the AlertRulesClient.NewListPager method.
type AlertRulesClientListOptions struct {
	// placeholder for future optional parameters
}

// AutomationRulesClientCreateOrUpdateOptions contains the optional parameters for the AutomationRulesClient.CreateOrUpdate
// method.
type AutomationRulesClientCreateOrUpdateOptions struct {
	// The automation rule
	AutomationRuleToUpsert *AutomationRule
}

// AutomationRulesClientDeleteOptions contains the optional parameters for the AutomationRulesClient.Delete method.
type AutomationRulesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// AutomationRulesClientGetOptions contains the optional parameters for the AutomationRulesClient.Get method.
type AutomationRulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// AutomationRulesClientListOptions contains the optional parameters for the AutomationRulesClient.NewListPager method.
type AutomationRulesClientListOptions struct {
	// placeholder for future optional parameters
}

// BookmarksClientCreateOrUpdateOptions contains the optional parameters for the BookmarksClient.CreateOrUpdate method.
type BookmarksClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// BookmarksClientDeleteOptions contains the optional parameters for the BookmarksClient.Delete method.
type BookmarksClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// BookmarksClientGetOptions contains the optional parameters for the BookmarksClient.Get method.
type BookmarksClientGetOptions struct {
	// placeholder for future optional parameters
}

// BookmarksClientListOptions contains the optional parameters for the BookmarksClient.NewListPager method.
type BookmarksClientListOptions struct {
	// placeholder for future optional parameters
}

// DataConnectorsClientCreateOrUpdateOptions contains the optional parameters for the DataConnectorsClient.CreateOrUpdate
// method.
type DataConnectorsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// DataConnectorsClientDeleteOptions contains the optional parameters for the DataConnectorsClient.Delete method.
type DataConnectorsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// DataConnectorsClientGetOptions contains the optional parameters for the DataConnectorsClient.Get method.
type DataConnectorsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DataConnectorsClientListOptions contains the optional parameters for the DataConnectorsClient.NewListPager method.
type DataConnectorsClientListOptions struct {
	// placeholder for future optional parameters
}

// IncidentCommentsClientCreateOrUpdateOptions contains the optional parameters for the IncidentCommentsClient.CreateOrUpdate
// method.
type IncidentCommentsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// IncidentCommentsClientDeleteOptions contains the optional parameters for the IncidentCommentsClient.Delete method.
type IncidentCommentsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// IncidentCommentsClientGetOptions contains the optional parameters for the IncidentCommentsClient.Get method.
type IncidentCommentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// IncidentCommentsClientListOptions contains the optional parameters for the IncidentCommentsClient.NewListPager method.
type IncidentCommentsClientListOptions struct {
	// Filters the results, based on a Boolean condition. Optional.
	Filter *string

	// Sorts the results. Optional.
	Orderby *string

	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skiptoken parameter that
	// specifies a starting point to use for subsequent calls. Optional.
	SkipToken *string

	// Returns only the first n results. Optional.
	Top *int32
}

// IncidentRelationsClientCreateOrUpdateOptions contains the optional parameters for the IncidentRelationsClient.CreateOrUpdate
// method.
type IncidentRelationsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// IncidentRelationsClientDeleteOptions contains the optional parameters for the IncidentRelationsClient.Delete method.
type IncidentRelationsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// IncidentRelationsClientGetOptions contains the optional parameters for the IncidentRelationsClient.Get method.
type IncidentRelationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// IncidentRelationsClientListOptions contains the optional parameters for the IncidentRelationsClient.NewListPager method.
type IncidentRelationsClientListOptions struct {
	// Filters the results, based on a Boolean condition. Optional.
	Filter *string

	// Sorts the results. Optional.
	Orderby *string

	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skiptoken parameter that
	// specifies a starting point to use for subsequent calls. Optional.
	SkipToken *string

	// Returns only the first n results. Optional.
	Top *int32
}

// IncidentsClientCreateOrUpdateOptions contains the optional parameters for the IncidentsClient.CreateOrUpdate method.
type IncidentsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// IncidentsClientDeleteOptions contains the optional parameters for the IncidentsClient.Delete method.
type IncidentsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// IncidentsClientGetOptions contains the optional parameters for the IncidentsClient.Get method.
type IncidentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// IncidentsClientListAlertsOptions contains the optional parameters for the IncidentsClient.ListAlerts method.
type IncidentsClientListAlertsOptions struct {
	// placeholder for future optional parameters
}

// IncidentsClientListBookmarksOptions contains the optional parameters for the IncidentsClient.ListBookmarks method.
type IncidentsClientListBookmarksOptions struct {
	// placeholder for future optional parameters
}

// IncidentsClientListEntitiesOptions contains the optional parameters for the IncidentsClient.ListEntities method.
type IncidentsClientListEntitiesOptions struct {
	// placeholder for future optional parameters
}

// IncidentsClientListOptions contains the optional parameters for the IncidentsClient.NewListPager method.
type IncidentsClientListOptions struct {
	// Filters the results, based on a Boolean condition. Optional.
	Filter *string

	// Sorts the results. Optional.
	Orderby *string

	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skiptoken parameter that
	// specifies a starting point to use for subsequent calls. Optional.
	SkipToken *string

	// Returns only the first n results. Optional.
	Top *int32
}

// MetadataClientCreateOptions contains the optional parameters for the MetadataClient.Create method.
type MetadataClientCreateOptions struct {
	// placeholder for future optional parameters
}

// MetadataClientDeleteOptions contains the optional parameters for the MetadataClient.Delete method.
type MetadataClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// MetadataClientGetOptions contains the optional parameters for the MetadataClient.Get method.
type MetadataClientGetOptions struct {
	// placeholder for future optional parameters
}

// MetadataClientListOptions contains the optional parameters for the MetadataClient.NewListPager method.
type MetadataClientListOptions struct {
	// Filters the results, based on a Boolean condition. Optional.
	Filter *string

	// Sorts the results. Optional.
	Orderby *string

	// Used to skip n elements in the OData query (offset). Returns a nextLink to the next page of results if there are any left.
	Skip *int32

	// Returns only the first n results. Optional.
	Top *int32
}

// MetadataClientUpdateOptions contains the optional parameters for the MetadataClient.Update method.
type MetadataClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// SecurityMLAnalyticsSettingsClientCreateOrUpdateOptions contains the optional parameters for the SecurityMLAnalyticsSettingsClient.CreateOrUpdate
// method.
type SecurityMLAnalyticsSettingsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// SecurityMLAnalyticsSettingsClientDeleteOptions contains the optional parameters for the SecurityMLAnalyticsSettingsClient.Delete
// method.
type SecurityMLAnalyticsSettingsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// SecurityMLAnalyticsSettingsClientGetOptions contains the optional parameters for the SecurityMLAnalyticsSettingsClient.Get
// method.
type SecurityMLAnalyticsSettingsClientGetOptions struct {
	// placeholder for future optional parameters
}

// SecurityMLAnalyticsSettingsClientListOptions contains the optional parameters for the SecurityMLAnalyticsSettingsClient.NewListPager
// method.
type SecurityMLAnalyticsSettingsClientListOptions struct {
	// placeholder for future optional parameters
}

// SentinelOnboardingStatesClientCreateOptions contains the optional parameters for the SentinelOnboardingStatesClient.Create
// method.
type SentinelOnboardingStatesClientCreateOptions struct {
	// The Sentinel onboarding state parameter
	SentinelOnboardingStateParameter *SentinelOnboardingState
}

// SentinelOnboardingStatesClientDeleteOptions contains the optional parameters for the SentinelOnboardingStatesClient.Delete
// method.
type SentinelOnboardingStatesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// SentinelOnboardingStatesClientGetOptions contains the optional parameters for the SentinelOnboardingStatesClient.Get method.
type SentinelOnboardingStatesClientGetOptions struct {
	// placeholder for future optional parameters
}

// SentinelOnboardingStatesClientListOptions contains the optional parameters for the SentinelOnboardingStatesClient.List
// method.
type SentinelOnboardingStatesClientListOptions struct {
	// placeholder for future optional parameters
}

// ThreatIntelligenceIndicatorClientAppendTagsOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.AppendTags
// method.
type ThreatIntelligenceIndicatorClientAppendTagsOptions struct {
	// placeholder for future optional parameters
}

// ThreatIntelligenceIndicatorClientCreateIndicatorOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.CreateIndicator
// method.
type ThreatIntelligenceIndicatorClientCreateIndicatorOptions struct {
	// placeholder for future optional parameters
}

// ThreatIntelligenceIndicatorClientCreateOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.Create
// method.
type ThreatIntelligenceIndicatorClientCreateOptions struct {
	// placeholder for future optional parameters
}

// ThreatIntelligenceIndicatorClientDeleteOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.Delete
// method.
type ThreatIntelligenceIndicatorClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ThreatIntelligenceIndicatorClientGetOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.Get
// method.
type ThreatIntelligenceIndicatorClientGetOptions struct {
	// placeholder for future optional parameters
}

// ThreatIntelligenceIndicatorClientQueryIndicatorsOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.NewQueryIndicatorsPager
// method.
type ThreatIntelligenceIndicatorClientQueryIndicatorsOptions struct {
	// placeholder for future optional parameters
}

// ThreatIntelligenceIndicatorClientReplaceTagsOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.ReplaceTags
// method.
type ThreatIntelligenceIndicatorClientReplaceTagsOptions struct {
	// placeholder for future optional parameters
}

// ThreatIntelligenceIndicatorMetricsClientListOptions contains the optional parameters for the ThreatIntelligenceIndicatorMetricsClient.List
// method.
type ThreatIntelligenceIndicatorMetricsClientListOptions struct {
	// placeholder for future optional parameters
}

// ThreatIntelligenceIndicatorsClientListOptions contains the optional parameters for the ThreatIntelligenceIndicatorsClient.NewListPager
// method.
type ThreatIntelligenceIndicatorsClientListOptions struct {
	// Filters the results, based on a Boolean condition. Optional.
	Filter *string

	// Sorts the results. Optional.
	Orderby *string

	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skiptoken parameter that
	// specifies a starting point to use for subsequent calls. Optional.
	SkipToken *string

	// Returns only the first n results. Optional.
	Top *int32
}

// WatchlistItemsClientCreateOrUpdateOptions contains the optional parameters for the WatchlistItemsClient.CreateOrUpdate
// method.
type WatchlistItemsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// WatchlistItemsClientDeleteOptions contains the optional parameters for the WatchlistItemsClient.Delete method.
type WatchlistItemsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// WatchlistItemsClientGetOptions contains the optional parameters for the WatchlistItemsClient.Get method.
type WatchlistItemsClientGetOptions struct {
	// placeholder for future optional parameters
}

// WatchlistItemsClientListOptions contains the optional parameters for the WatchlistItemsClient.NewListPager method.
type WatchlistItemsClientListOptions struct {
	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skiptoken parameter that
	// specifies a starting point to use for subsequent calls. Optional.
	SkipToken *string
}

// WatchlistsClientCreateOrUpdateOptions contains the optional parameters for the WatchlistsClient.CreateOrUpdate method.
type WatchlistsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// WatchlistsClientDeleteOptions contains the optional parameters for the WatchlistsClient.Delete method.
type WatchlistsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// WatchlistsClientGetOptions contains the optional parameters for the WatchlistsClient.Get method.
type WatchlistsClientGetOptions struct {
	// placeholder for future optional parameters
}

// WatchlistsClientListOptions contains the optional parameters for the WatchlistsClient.NewListPager method.
type WatchlistsClientListOptions struct {
	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skiptoken parameter that
	// specifies a starting point to use for subsequent calls. Optional.
	SkipToken *string
}
