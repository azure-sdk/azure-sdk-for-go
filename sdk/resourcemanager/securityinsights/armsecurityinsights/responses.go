//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

// ActionsClientCreateOrUpdateResponse contains the response from method ActionsClient.CreateOrUpdate.
type ActionsClientCreateOrUpdateResponse struct {
	// Action for alert rule.
	ActionResponse
}

// ActionsClientDeleteResponse contains the response from method ActionsClient.Delete.
type ActionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ActionsClientGetResponse contains the response from method ActionsClient.Get.
type ActionsClientGetResponse struct {
	// Action for alert rule.
	ActionResponse
}

// ActionsClientListByAlertRuleResponse contains the response from method ActionsClient.NewListByAlertRulePager.
type ActionsClientListByAlertRuleResponse struct {
	// List all the actions.
	ActionsList
}

// AlertRuleTemplatesClientGetResponse contains the response from method AlertRuleTemplatesClient.Get.
type AlertRuleTemplatesClientGetResponse struct {
	// Alert rule template.
	AlertRuleTemplateClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AlertRuleTemplatesClientGetResponse.
func (a *AlertRuleTemplatesClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalAlertRuleTemplateClassification(data)
	if err != nil {
		return err
	}
	a.AlertRuleTemplateClassification = res
	return nil
}

// AlertRuleTemplatesClientListResponse contains the response from method AlertRuleTemplatesClient.NewListPager.
type AlertRuleTemplatesClientListResponse struct {
	// List all the alert rule templates.
	AlertRuleTemplatesList
}

// AlertRulesClientCreateOrUpdateResponse contains the response from method AlertRulesClient.CreateOrUpdate.
type AlertRulesClientCreateOrUpdateResponse struct {
	// Alert rule.
	AlertRuleClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AlertRulesClientCreateOrUpdateResponse.
func (a *AlertRulesClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalAlertRuleClassification(data)
	if err != nil {
		return err
	}
	a.AlertRuleClassification = res
	return nil
}

// AlertRulesClientDeleteResponse contains the response from method AlertRulesClient.Delete.
type AlertRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// AlertRulesClientGetResponse contains the response from method AlertRulesClient.Get.
type AlertRulesClientGetResponse struct {
	// Alert rule.
	AlertRuleClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AlertRulesClientGetResponse.
func (a *AlertRulesClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalAlertRuleClassification(data)
	if err != nil {
		return err
	}
	a.AlertRuleClassification = res
	return nil
}

// AlertRulesClientListResponse contains the response from method AlertRulesClient.NewListPager.
type AlertRulesClientListResponse struct {
	// List all the alert rules.
	AlertRulesList
}

// AutomationRulesClientCreateOrUpdateResponse contains the response from method AutomationRulesClient.CreateOrUpdate.
type AutomationRulesClientCreateOrUpdateResponse struct {
	AutomationRule
}

// AutomationRulesClientDeleteResponse contains the response from method AutomationRulesClient.Delete.
type AutomationRulesClientDeleteResponse struct {
	// Anything
	Interface any
}

// AutomationRulesClientGetResponse contains the response from method AutomationRulesClient.Get.
type AutomationRulesClientGetResponse struct {
	AutomationRule
}

// AutomationRulesClientListResponse contains the response from method AutomationRulesClient.NewListPager.
type AutomationRulesClientListResponse struct {
	AutomationRulesList
}

// BookmarksClientCreateOrUpdateResponse contains the response from method BookmarksClient.CreateOrUpdate.
type BookmarksClientCreateOrUpdateResponse struct {
	// Represents a bookmark in Azure Security Insights.
	Bookmark
}

// BookmarksClientDeleteResponse contains the response from method BookmarksClient.Delete.
type BookmarksClientDeleteResponse struct {
	// placeholder for future response values
}

// BookmarksClientGetResponse contains the response from method BookmarksClient.Get.
type BookmarksClientGetResponse struct {
	// Represents a bookmark in Azure Security Insights.
	Bookmark
}

// BookmarksClientListResponse contains the response from method BookmarksClient.NewListPager.
type BookmarksClientListResponse struct {
	// List all the bookmarks.
	BookmarkList
}

// ContentPackageClientInstallResponse contains the response from method ContentPackageClient.Install.
type ContentPackageClientInstallResponse struct {
	// Represents a Package in Azure Security Insights.
	PackageModel
}

// ContentPackageClientUninstallResponse contains the response from method ContentPackageClient.Uninstall.
type ContentPackageClientUninstallResponse struct {
	// placeholder for future response values
}

// ContentPackagesClientGetResponse contains the response from method ContentPackagesClient.Get.
type ContentPackagesClientGetResponse struct {
	// Represents a Package in Azure Security Insights.
	PackageModel
}

// ContentPackagesClientListResponse contains the response from method ContentPackagesClient.NewListPager.
type ContentPackagesClientListResponse struct {
	// List available packages.
	PackageList
}

// ContentTemplateClientDeleteResponse contains the response from method ContentTemplateClient.Delete.
type ContentTemplateClientDeleteResponse struct {
	// placeholder for future response values
}

// ContentTemplateClientGetResponse contains the response from method ContentTemplateClient.Get.
type ContentTemplateClientGetResponse struct {
	// Template resource definition.
	TemplateModel
}

// ContentTemplateClientInstallResponse contains the response from method ContentTemplateClient.Install.
type ContentTemplateClientInstallResponse struct {
	// Template resource definition.
	TemplateModel
}

// ContentTemplatesClientListResponse contains the response from method ContentTemplatesClient.NewListPager.
type ContentTemplatesClientListResponse struct {
	// List of all the template.
	TemplateList
}

// DataConnectorDefinitionsClientCreateOrUpdateResponse contains the response from method DataConnectorDefinitionsClient.CreateOrUpdate.
type DataConnectorDefinitionsClientCreateOrUpdateResponse struct {
	// An Azure resource, which encapsulate the entire info requires to display a data connector page in Azure portal,
	// and the info required to define data connections.
	DataConnectorDefinitionClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DataConnectorDefinitionsClientCreateOrUpdateResponse.
func (d *DataConnectorDefinitionsClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDataConnectorDefinitionClassification(data)
	if err != nil {
		return err
	}
	d.DataConnectorDefinitionClassification = res
	return nil
}

// DataConnectorDefinitionsClientDeleteResponse contains the response from method DataConnectorDefinitionsClient.Delete.
type DataConnectorDefinitionsClientDeleteResponse struct {
	// placeholder for future response values
}

// DataConnectorDefinitionsClientGetResponse contains the response from method DataConnectorDefinitionsClient.Get.
type DataConnectorDefinitionsClientGetResponse struct {
	// An Azure resource, which encapsulate the entire info requires to display a data connector page in Azure portal,
	// and the info required to define data connections.
	DataConnectorDefinitionClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DataConnectorDefinitionsClientGetResponse.
func (d *DataConnectorDefinitionsClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDataConnectorDefinitionClassification(data)
	if err != nil {
		return err
	}
	d.DataConnectorDefinitionClassification = res
	return nil
}

// DataConnectorDefinitionsClientListResponse contains the response from method DataConnectorDefinitionsClient.NewListPager.
type DataConnectorDefinitionsClientListResponse struct {
	// Encapsulate the data connector definition object
	DataConnectorDefinitionArmCollectionWrapper
}

// DataConnectorsClientCreateOrUpdateResponse contains the response from method DataConnectorsClient.CreateOrUpdate.
type DataConnectorsClientCreateOrUpdateResponse struct {
	// Data connector.
	DataConnectorClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DataConnectorsClientCreateOrUpdateResponse.
func (d *DataConnectorsClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDataConnectorClassification(data)
	if err != nil {
		return err
	}
	d.DataConnectorClassification = res
	return nil
}

// DataConnectorsClientDeleteResponse contains the response from method DataConnectorsClient.Delete.
type DataConnectorsClientDeleteResponse struct {
	// placeholder for future response values
}

// DataConnectorsClientGetResponse contains the response from method DataConnectorsClient.Get.
type DataConnectorsClientGetResponse struct {
	// Data connector.
	DataConnectorClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DataConnectorsClientGetResponse.
func (d *DataConnectorsClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDataConnectorClassification(data)
	if err != nil {
		return err
	}
	d.DataConnectorClassification = res
	return nil
}

// DataConnectorsClientListResponse contains the response from method DataConnectorsClient.NewListPager.
type DataConnectorsClientListResponse struct {
	// List all the data connectors.
	DataConnectorList
}

// EntitiesClientRunPlaybookResponse contains the response from method EntitiesClient.RunPlaybook.
type EntitiesClientRunPlaybookResponse struct {
	// placeholder for future response values
}

// IncidentCommentsClientCreateOrUpdateResponse contains the response from method IncidentCommentsClient.CreateOrUpdate.
type IncidentCommentsClientCreateOrUpdateResponse struct {
	// Represents an incident comment
	IncidentComment
}

// IncidentCommentsClientDeleteResponse contains the response from method IncidentCommentsClient.Delete.
type IncidentCommentsClientDeleteResponse struct {
	// placeholder for future response values
}

// IncidentCommentsClientGetResponse contains the response from method IncidentCommentsClient.Get.
type IncidentCommentsClientGetResponse struct {
	// Represents an incident comment
	IncidentComment
}

// IncidentCommentsClientListResponse contains the response from method IncidentCommentsClient.NewListPager.
type IncidentCommentsClientListResponse struct {
	// List of incident comments.
	IncidentCommentList
}

// IncidentRelationsClientCreateOrUpdateResponse contains the response from method IncidentRelationsClient.CreateOrUpdate.
type IncidentRelationsClientCreateOrUpdateResponse struct {
	// Represents a relation between two resources
	Relation
}

// IncidentRelationsClientDeleteResponse contains the response from method IncidentRelationsClient.Delete.
type IncidentRelationsClientDeleteResponse struct {
	// placeholder for future response values
}

// IncidentRelationsClientGetResponse contains the response from method IncidentRelationsClient.Get.
type IncidentRelationsClientGetResponse struct {
	// Represents a relation between two resources
	Relation
}

// IncidentRelationsClientListResponse contains the response from method IncidentRelationsClient.NewListPager.
type IncidentRelationsClientListResponse struct {
	// List of relations.
	RelationList
}

// IncidentTasksClientCreateOrUpdateResponse contains the response from method IncidentTasksClient.CreateOrUpdate.
type IncidentTasksClientCreateOrUpdateResponse struct {
	// Describes incident task properties
	IncidentTask
}

// IncidentTasksClientDeleteResponse contains the response from method IncidentTasksClient.Delete.
type IncidentTasksClientDeleteResponse struct {
	// placeholder for future response values
}

// IncidentTasksClientGetResponse contains the response from method IncidentTasksClient.Get.
type IncidentTasksClientGetResponse struct {
	// Describes incident task properties
	IncidentTask
}

// IncidentTasksClientListResponse contains the response from method IncidentTasksClient.NewListPager.
type IncidentTasksClientListResponse struct {
	// List of incident tasks
	IncidentTaskList
}

// IncidentsClientCreateOrUpdateResponse contains the response from method IncidentsClient.CreateOrUpdate.
type IncidentsClientCreateOrUpdateResponse struct {
	// Represents an incident in Azure Security Insights.
	Incident
}

// IncidentsClientDeleteResponse contains the response from method IncidentsClient.Delete.
type IncidentsClientDeleteResponse struct {
	// placeholder for future response values
}

// IncidentsClientGetResponse contains the response from method IncidentsClient.Get.
type IncidentsClientGetResponse struct {
	// Represents an incident in Azure Security Insights.
	Incident
}

// IncidentsClientListAlertsResponse contains the response from method IncidentsClient.ListAlerts.
type IncidentsClientListAlertsResponse struct {
	// List of incident alerts.
	IncidentAlertList
}

// IncidentsClientListBookmarksResponse contains the response from method IncidentsClient.ListBookmarks.
type IncidentsClientListBookmarksResponse struct {
	// List of incident bookmarks.
	IncidentBookmarkList
}

// IncidentsClientListEntitiesResponse contains the response from method IncidentsClient.ListEntities.
type IncidentsClientListEntitiesResponse struct {
	// The incident related entities response.
	IncidentEntitiesResponse
}

// IncidentsClientListResponse contains the response from method IncidentsClient.NewListPager.
type IncidentsClientListResponse struct {
	// List all the incidents.
	IncidentList
}

// IncidentsClientRunPlaybookResponse contains the response from method IncidentsClient.RunPlaybook.
type IncidentsClientRunPlaybookResponse struct {
	// placeholder for future response values
}

// MetadataClientCreateResponse contains the response from method MetadataClient.Create.
type MetadataClientCreateResponse struct {
	// Metadata resource definition.
	MetadataModel
}

// MetadataClientDeleteResponse contains the response from method MetadataClient.Delete.
type MetadataClientDeleteResponse struct {
	// placeholder for future response values
}

// MetadataClientGetResponse contains the response from method MetadataClient.Get.
type MetadataClientGetResponse struct {
	// Metadata resource definition.
	MetadataModel
}

// MetadataClientListResponse contains the response from method MetadataClient.NewListPager.
type MetadataClientListResponse struct {
	// List of all the metadata.
	MetadataList
}

// MetadataClientUpdateResponse contains the response from method MetadataClient.Update.
type MetadataClientUpdateResponse struct {
	// Metadata resource definition.
	MetadataModel
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Lists the operations available in the SecurityInsights RP.
	OperationsList
}

// ProductPackageClientGetResponse contains the response from method ProductPackageClient.Get.
type ProductPackageClientGetResponse struct {
	// Represents a Package in Azure Security Insights.
	ProductPackageModel
}

// ProductPackagesClientListResponse contains the response from method ProductPackagesClient.NewListPager.
type ProductPackagesClientListResponse struct {
	// List available packages.
	ProductPackageList
}

// ProductTemplateClientGetResponse contains the response from method ProductTemplateClient.Get.
type ProductTemplateClientGetResponse struct {
	// Template resource definition.
	ProductTemplateModel
}

// ProductTemplatesClientListResponse contains the response from method ProductTemplatesClient.NewListPager.
type ProductTemplatesClientListResponse struct {
	// List of all the template.
	ProductTemplateList
}

// SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse contains the response from method SecurityMLAnalyticsSettingsClient.CreateOrUpdate.
type SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse struct {
	// Security ML Analytics Setting
	SecurityMLAnalyticsSettingClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse.
func (s *SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalSecurityMLAnalyticsSettingClassification(data)
	if err != nil {
		return err
	}
	s.SecurityMLAnalyticsSettingClassification = res
	return nil
}

// SecurityMLAnalyticsSettingsClientDeleteResponse contains the response from method SecurityMLAnalyticsSettingsClient.Delete.
type SecurityMLAnalyticsSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// SecurityMLAnalyticsSettingsClientGetResponse contains the response from method SecurityMLAnalyticsSettingsClient.Get.
type SecurityMLAnalyticsSettingsClientGetResponse struct {
	// Security ML Analytics Setting
	SecurityMLAnalyticsSettingClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SecurityMLAnalyticsSettingsClientGetResponse.
func (s *SecurityMLAnalyticsSettingsClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalSecurityMLAnalyticsSettingClassification(data)
	if err != nil {
		return err
	}
	s.SecurityMLAnalyticsSettingClassification = res
	return nil
}

// SecurityMLAnalyticsSettingsClientListResponse contains the response from method SecurityMLAnalyticsSettingsClient.NewListPager.
type SecurityMLAnalyticsSettingsClientListResponse struct {
	// List all the SecurityMLAnalyticsSettings
	SecurityMLAnalyticsSettingsList
}

// SentinelOnboardingStatesClientCreateResponse contains the response from method SentinelOnboardingStatesClient.Create.
type SentinelOnboardingStatesClientCreateResponse struct {
	// Sentinel onboarding state
	SentinelOnboardingState
}

// SentinelOnboardingStatesClientDeleteResponse contains the response from method SentinelOnboardingStatesClient.Delete.
type SentinelOnboardingStatesClientDeleteResponse struct {
	// placeholder for future response values
}

// SentinelOnboardingStatesClientGetResponse contains the response from method SentinelOnboardingStatesClient.Get.
type SentinelOnboardingStatesClientGetResponse struct {
	// Sentinel onboarding state
	SentinelOnboardingState
}

// SentinelOnboardingStatesClientListResponse contains the response from method SentinelOnboardingStatesClient.List.
type SentinelOnboardingStatesClientListResponse struct {
	// List of the Sentinel onboarding states
	SentinelOnboardingStatesList
}

// SourceControlClientListRepositoriesResponse contains the response from method SourceControlClient.NewListRepositoriesPager.
type SourceControlClientListRepositoriesResponse struct {
	// List all the source controls.
	RepoList
}

// SourceControlsClientCreateResponse contains the response from method SourceControlsClient.Create.
type SourceControlsClientCreateResponse struct {
	// Represents a SourceControl in Azure Security Insights.
	SourceControl
}

// SourceControlsClientDeleteResponse contains the response from method SourceControlsClient.Delete.
type SourceControlsClientDeleteResponse struct {
	// Warning response structure.
	Warning
}

// SourceControlsClientGetResponse contains the response from method SourceControlsClient.Get.
type SourceControlsClientGetResponse struct {
	// Represents a SourceControl in Azure Security Insights.
	SourceControl
}

// SourceControlsClientListResponse contains the response from method SourceControlsClient.NewListPager.
type SourceControlsClientListResponse struct {
	// List all the source controls.
	SourceControlList
}

// ThreatIntelligenceIndicatorClientAppendTagsResponse contains the response from method ThreatIntelligenceIndicatorClient.AppendTags.
type ThreatIntelligenceIndicatorClientAppendTagsResponse struct {
	// placeholder for future response values
}

// ThreatIntelligenceIndicatorClientCreateIndicatorResponse contains the response from method ThreatIntelligenceIndicatorClient.CreateIndicator.
type ThreatIntelligenceIndicatorClientCreateIndicatorResponse struct {
	// Threat intelligence information object.
	ThreatIntelligenceInformationClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ThreatIntelligenceIndicatorClientCreateIndicatorResponse.
func (t *ThreatIntelligenceIndicatorClientCreateIndicatorResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalThreatIntelligenceInformationClassification(data)
	if err != nil {
		return err
	}
	t.ThreatIntelligenceInformationClassification = res
	return nil
}

// ThreatIntelligenceIndicatorClientCreateResponse contains the response from method ThreatIntelligenceIndicatorClient.Create.
type ThreatIntelligenceIndicatorClientCreateResponse struct {
	// Threat intelligence information object.
	ThreatIntelligenceInformationClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ThreatIntelligenceIndicatorClientCreateResponse.
func (t *ThreatIntelligenceIndicatorClientCreateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalThreatIntelligenceInformationClassification(data)
	if err != nil {
		return err
	}
	t.ThreatIntelligenceInformationClassification = res
	return nil
}

// ThreatIntelligenceIndicatorClientDeleteResponse contains the response from method ThreatIntelligenceIndicatorClient.Delete.
type ThreatIntelligenceIndicatorClientDeleteResponse struct {
	// placeholder for future response values
}

// ThreatIntelligenceIndicatorClientGetResponse contains the response from method ThreatIntelligenceIndicatorClient.Get.
type ThreatIntelligenceIndicatorClientGetResponse struct {
	// Threat intelligence information object.
	ThreatIntelligenceInformationClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ThreatIntelligenceIndicatorClientGetResponse.
func (t *ThreatIntelligenceIndicatorClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalThreatIntelligenceInformationClassification(data)
	if err != nil {
		return err
	}
	t.ThreatIntelligenceInformationClassification = res
	return nil
}

// ThreatIntelligenceIndicatorClientQueryIndicatorsResponse contains the response from method ThreatIntelligenceIndicatorClient.NewQueryIndicatorsPager.
type ThreatIntelligenceIndicatorClientQueryIndicatorsResponse struct {
	// List of all the threat intelligence information objects.
	ThreatIntelligenceInformationList
}

// ThreatIntelligenceIndicatorClientReplaceTagsResponse contains the response from method ThreatIntelligenceIndicatorClient.ReplaceTags.
type ThreatIntelligenceIndicatorClientReplaceTagsResponse struct {
	// Threat intelligence information object.
	ThreatIntelligenceInformationClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ThreatIntelligenceIndicatorClientReplaceTagsResponse.
func (t *ThreatIntelligenceIndicatorClientReplaceTagsResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalThreatIntelligenceInformationClassification(data)
	if err != nil {
		return err
	}
	t.ThreatIntelligenceInformationClassification = res
	return nil
}

// ThreatIntelligenceIndicatorMetricsClientListResponse contains the response from method ThreatIntelligenceIndicatorMetricsClient.List.
type ThreatIntelligenceIndicatorMetricsClientListResponse struct {
	// List of all the threat intelligence metric fields (type/threat type/source).
	ThreatIntelligenceMetricsList
}

// ThreatIntelligenceIndicatorsClientListResponse contains the response from method ThreatIntelligenceIndicatorsClient.NewListPager.
type ThreatIntelligenceIndicatorsClientListResponse struct {
	// List of all the threat intelligence information objects.
	ThreatIntelligenceInformationList
}

// WatchlistItemsClientCreateOrUpdateResponse contains the response from method WatchlistItemsClient.CreateOrUpdate.
type WatchlistItemsClientCreateOrUpdateResponse struct {
	// Represents a Watchlist Item in Azure Security Insights.
	WatchlistItem
}

// WatchlistItemsClientDeleteResponse contains the response from method WatchlistItemsClient.Delete.
type WatchlistItemsClientDeleteResponse struct {
	// placeholder for future response values
}

// WatchlistItemsClientGetResponse contains the response from method WatchlistItemsClient.Get.
type WatchlistItemsClientGetResponse struct {
	// Represents a Watchlist Item in Azure Security Insights.
	WatchlistItem
}

// WatchlistItemsClientListResponse contains the response from method WatchlistItemsClient.NewListPager.
type WatchlistItemsClientListResponse struct {
	// List all the watchlist items.
	WatchlistItemList
}

// WatchlistsClientCreateOrUpdateResponse contains the response from method WatchlistsClient.CreateOrUpdate.
type WatchlistsClientCreateOrUpdateResponse struct {
	// Represents a Watchlist in Azure Security Insights.
	Watchlist
}

// WatchlistsClientDeleteResponse contains the response from method WatchlistsClient.Delete.
type WatchlistsClientDeleteResponse struct {
	// placeholder for future response values
}

// WatchlistsClientGetResponse contains the response from method WatchlistsClient.Get.
type WatchlistsClientGetResponse struct {
	// Represents a Watchlist in Azure Security Insights.
	Watchlist
}

// WatchlistsClientListResponse contains the response from method WatchlistsClient.NewListPager.
type WatchlistsClientListResponse struct {
	// List all the watchlists.
	WatchlistList
}
