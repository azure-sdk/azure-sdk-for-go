// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewAPICollectionsClient creates a new instance of APICollectionsClient.
func (c *ClientFactory) NewAPICollectionsClient() *APICollectionsClient {
	return &APICollectionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAccountConnectorsClient creates a new instance of AccountConnectorsClient.
func (c *ClientFactory) NewAccountConnectorsClient() *AccountConnectorsClient {
	return &AccountConnectorsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAdvancedThreatProtectionClient creates a new instance of AdvancedThreatProtectionClient.
func (c *ClientFactory) NewAdvancedThreatProtectionClient() *AdvancedThreatProtectionClient {
	return &AdvancedThreatProtectionClient{
		internal: c.internal,
	}
}

// NewAlertsClient creates a new instance of AlertsClient.
func (c *ClientFactory) NewAlertsClient() *AlertsClient {
	return &AlertsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAlertsSuppressionRulesClient creates a new instance of AlertsSuppressionRulesClient.
func (c *ClientFactory) NewAlertsSuppressionRulesClient() *AlertsSuppressionRulesClient {
	return &AlertsSuppressionRulesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAllowedConnectionsClient creates a new instance of AllowedConnectionsClient.
func (c *ClientFactory) NewAllowedConnectionsClient() *AllowedConnectionsClient {
	return &AllowedConnectionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewApplicationClient creates a new instance of ApplicationClient.
func (c *ClientFactory) NewApplicationClient() *ApplicationClient {
	return &ApplicationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewApplicationsClient creates a new instance of ApplicationsClient.
func (c *ClientFactory) NewApplicationsClient() *ApplicationsClient {
	return &ApplicationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAssessmentsClient creates a new instance of AssessmentsClient.
func (c *ClientFactory) NewAssessmentsClient() *AssessmentsClient {
	return &AssessmentsClient{
		internal: c.internal,
	}
}

// NewAssessmentsMetadataClient creates a new instance of AssessmentsMetadataClient.
func (c *ClientFactory) NewAssessmentsMetadataClient() *AssessmentsMetadataClient {
	return &AssessmentsMetadataClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAutoProvisioningSettingsClient creates a new instance of AutoProvisioningSettingsClient.
func (c *ClientFactory) NewAutoProvisioningSettingsClient() *AutoProvisioningSettingsClient {
	return &AutoProvisioningSettingsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAutomationsClient creates a new instance of AutomationsClient.
func (c *ClientFactory) NewAutomationsClient() *AutomationsClient {
	return &AutomationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAzureDevOpsOrgsClient creates a new instance of AzureDevOpsOrgsClient.
func (c *ClientFactory) NewAzureDevOpsOrgsClient() *AzureDevOpsOrgsClient {
	return &AzureDevOpsOrgsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAzureDevOpsProjectsClient creates a new instance of AzureDevOpsProjectsClient.
func (c *ClientFactory) NewAzureDevOpsProjectsClient() *AzureDevOpsProjectsClient {
	return &AzureDevOpsProjectsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAzureDevOpsReposClient creates a new instance of AzureDevOpsReposClient.
func (c *ClientFactory) NewAzureDevOpsReposClient() *AzureDevOpsReposClient {
	return &AzureDevOpsReposClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewComplianceResultsClient creates a new instance of ComplianceResultsClient.
func (c *ClientFactory) NewComplianceResultsClient() *ComplianceResultsClient {
	return &ComplianceResultsClient{
		internal: c.internal,
	}
}

// NewCompliancesClient creates a new instance of CompliancesClient.
func (c *ClientFactory) NewCompliancesClient() *CompliancesClient {
	return &CompliancesClient{
		internal: c.internal,
	}
}

// NewConnectorApplicationClient creates a new instance of ConnectorApplicationClient.
func (c *ClientFactory) NewConnectorApplicationClient() *ConnectorApplicationClient {
	return &ConnectorApplicationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConnectorApplicationsClient creates a new instance of ConnectorApplicationsClient.
func (c *ClientFactory) NewConnectorApplicationsClient() *ConnectorApplicationsClient {
	return &ConnectorApplicationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConnectorsClient creates a new instance of ConnectorsClient.
func (c *ClientFactory) NewConnectorsClient() *ConnectorsClient {
	return &ConnectorsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewContactsClient creates a new instance of ContactsClient.
func (c *ClientFactory) NewContactsClient() *ContactsClient {
	return &ContactsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCustomAssessmentAutomationsClient creates a new instance of CustomAssessmentAutomationsClient.
func (c *ClientFactory) NewCustomAssessmentAutomationsClient() *CustomAssessmentAutomationsClient {
	return &CustomAssessmentAutomationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCustomEntityStoreAssignmentsClient creates a new instance of CustomEntityStoreAssignmentsClient.
func (c *ClientFactory) NewCustomEntityStoreAssignmentsClient() *CustomEntityStoreAssignmentsClient {
	return &CustomEntityStoreAssignmentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCustomRecommendationsClient creates a new instance of CustomRecommendationsClient.
func (c *ClientFactory) NewCustomRecommendationsClient() *CustomRecommendationsClient {
	return &CustomRecommendationsClient{
		internal: c.internal,
	}
}

// NewDefenderForStorageClient creates a new instance of DefenderForStorageClient.
func (c *ClientFactory) NewDefenderForStorageClient() *DefenderForStorageClient {
	return &DefenderForStorageClient{
		internal: c.internal,
	}
}

// NewDevOpsConfigurationsClient creates a new instance of DevOpsConfigurationsClient.
func (c *ClientFactory) NewDevOpsConfigurationsClient() *DevOpsConfigurationsClient {
	return &DevOpsConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDevOpsOperationResultsClient creates a new instance of DevOpsOperationResultsClient.
func (c *ClientFactory) NewDevOpsOperationResultsClient() *DevOpsOperationResultsClient {
	return &DevOpsOperationResultsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDevOpsPoliciesClient creates a new instance of DevOpsPoliciesClient.
func (c *ClientFactory) NewDevOpsPoliciesClient() *DevOpsPoliciesClient {
	return &DevOpsPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDevOpsPolicyAssignmentsClient creates a new instance of DevOpsPolicyAssignmentsClient.
func (c *ClientFactory) NewDevOpsPolicyAssignmentsClient() *DevOpsPolicyAssignmentsClient {
	return &DevOpsPolicyAssignmentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDeviceSecurityGroupsClient creates a new instance of DeviceSecurityGroupsClient.
func (c *ClientFactory) NewDeviceSecurityGroupsClient() *DeviceSecurityGroupsClient {
	return &DeviceSecurityGroupsClient{
		internal: c.internal,
	}
}

// NewDiscoveredSecuritySolutionsClient creates a new instance of DiscoveredSecuritySolutionsClient.
func (c *ClientFactory) NewDiscoveredSecuritySolutionsClient() *DiscoveredSecuritySolutionsClient {
	return &DiscoveredSecuritySolutionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewExternalSecuritySolutionsClient creates a new instance of ExternalSecuritySolutionsClient.
func (c *ClientFactory) NewExternalSecuritySolutionsClient() *ExternalSecuritySolutionsClient {
	return &ExternalSecuritySolutionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGitHubOwnersClient creates a new instance of GitHubOwnersClient.
func (c *ClientFactory) NewGitHubOwnersClient() *GitHubOwnersClient {
	return &GitHubOwnersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGitHubReposClient creates a new instance of GitHubReposClient.
func (c *ClientFactory) NewGitHubReposClient() *GitHubReposClient {
	return &GitHubReposClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGitLabGroupsClient creates a new instance of GitLabGroupsClient.
func (c *ClientFactory) NewGitLabGroupsClient() *GitLabGroupsClient {
	return &GitLabGroupsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGitLabProjectsClient creates a new instance of GitLabProjectsClient.
func (c *ClientFactory) NewGitLabProjectsClient() *GitLabProjectsClient {
	return &GitLabProjectsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGitLabSubgroupsClient creates a new instance of GitLabSubgroupsClient.
func (c *ClientFactory) NewGitLabSubgroupsClient() *GitLabSubgroupsClient {
	return &GitLabSubgroupsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGovernanceAssignmentsClient creates a new instance of GovernanceAssignmentsClient.
func (c *ClientFactory) NewGovernanceAssignmentsClient() *GovernanceAssignmentsClient {
	return &GovernanceAssignmentsClient{
		internal: c.internal,
	}
}

// NewGovernanceRulesClient creates a new instance of GovernanceRulesClient.
func (c *ClientFactory) NewGovernanceRulesClient() *GovernanceRulesClient {
	return &GovernanceRulesClient{
		internal: c.internal,
	}
}

// NewHealthReportsClient creates a new instance of HealthReportsClient.
func (c *ClientFactory) NewHealthReportsClient() *HealthReportsClient {
	return &HealthReportsClient{
		internal: c.internal,
	}
}

// NewInformationProtectionPoliciesClient creates a new instance of InformationProtectionPoliciesClient.
func (c *ClientFactory) NewInformationProtectionPoliciesClient() *InformationProtectionPoliciesClient {
	return &InformationProtectionPoliciesClient{
		internal: c.internal,
	}
}

// NewIotSecuritySolutionAnalyticsClient creates a new instance of IotSecuritySolutionAnalyticsClient.
func (c *ClientFactory) NewIotSecuritySolutionAnalyticsClient() *IotSecuritySolutionAnalyticsClient {
	return &IotSecuritySolutionAnalyticsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIotSecuritySolutionClient creates a new instance of IotSecuritySolutionClient.
func (c *ClientFactory) NewIotSecuritySolutionClient() *IotSecuritySolutionClient {
	return &IotSecuritySolutionClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIotSecuritySolutionsAnalyticsAggregatedAlertClient creates a new instance of IotSecuritySolutionsAnalyticsAggregatedAlertClient.
func (c *ClientFactory) NewIotSecuritySolutionsAnalyticsAggregatedAlertClient() *IotSecuritySolutionsAnalyticsAggregatedAlertClient {
	return &IotSecuritySolutionsAnalyticsAggregatedAlertClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIotSecuritySolutionsAnalyticsRecommendationClient creates a new instance of IotSecuritySolutionsAnalyticsRecommendationClient.
func (c *ClientFactory) NewIotSecuritySolutionsAnalyticsRecommendationClient() *IotSecuritySolutionsAnalyticsRecommendationClient {
	return &IotSecuritySolutionsAnalyticsRecommendationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewJitNetworkAccessPoliciesClient creates a new instance of JitNetworkAccessPoliciesClient.
func (c *ClientFactory) NewJitNetworkAccessPoliciesClient() *JitNetworkAccessPoliciesClient {
	return &JitNetworkAccessPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewLocationsClient creates a new instance of LocationsClient.
func (c *ClientFactory) NewLocationsClient() *LocationsClient {
	return &LocationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMdeOnboardingsClient creates a new instance of MdeOnboardingsClient.
func (c *ClientFactory) NewMdeOnboardingsClient() *MdeOnboardingsClient {
	return &MdeOnboardingsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewOperatorsClient creates a new instance of OperatorsClient.
func (c *ClientFactory) NewOperatorsClient() *OperatorsClient {
	return &OperatorsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPricingsClient creates a new instance of PricingsClient.
func (c *ClientFactory) NewPricingsClient() *PricingsClient {
	return &PricingsClient{
		internal: c.internal,
	}
}

// NewRegulatoryComplianceAssessmentsClient creates a new instance of RegulatoryComplianceAssessmentsClient.
func (c *ClientFactory) NewRegulatoryComplianceAssessmentsClient() *RegulatoryComplianceAssessmentsClient {
	return &RegulatoryComplianceAssessmentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewRegulatoryComplianceControlsClient creates a new instance of RegulatoryComplianceControlsClient.
func (c *ClientFactory) NewRegulatoryComplianceControlsClient() *RegulatoryComplianceControlsClient {
	return &RegulatoryComplianceControlsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewRegulatoryComplianceStandardsClient creates a new instance of RegulatoryComplianceStandardsClient.
func (c *ClientFactory) NewRegulatoryComplianceStandardsClient() *RegulatoryComplianceStandardsClient {
	return &RegulatoryComplianceStandardsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLVulnerabilityAssessmentBaselineRulesClient creates a new instance of SQLVulnerabilityAssessmentBaselineRulesClient.
func (c *ClientFactory) NewSQLVulnerabilityAssessmentBaselineRulesClient() *SQLVulnerabilityAssessmentBaselineRulesClient {
	return &SQLVulnerabilityAssessmentBaselineRulesClient{
		internal: c.internal,
	}
}

// NewSQLVulnerabilityAssessmentScanResultsClient creates a new instance of SQLVulnerabilityAssessmentScanResultsClient.
func (c *ClientFactory) NewSQLVulnerabilityAssessmentScanResultsClient() *SQLVulnerabilityAssessmentScanResultsClient {
	return &SQLVulnerabilityAssessmentScanResultsClient{
		internal: c.internal,
	}
}

// NewSQLVulnerabilityAssessmentScansClient creates a new instance of SQLVulnerabilityAssessmentScansClient.
func (c *ClientFactory) NewSQLVulnerabilityAssessmentScansClient() *SQLVulnerabilityAssessmentScansClient {
	return &SQLVulnerabilityAssessmentScansClient{
		internal: c.internal,
	}
}

// NewSecureScoreControlDefinitionsClient creates a new instance of SecureScoreControlDefinitionsClient.
func (c *ClientFactory) NewSecureScoreControlDefinitionsClient() *SecureScoreControlDefinitionsClient {
	return &SecureScoreControlDefinitionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSecureScoreControlsClient creates a new instance of SecureScoreControlsClient.
func (c *ClientFactory) NewSecureScoreControlsClient() *SecureScoreControlsClient {
	return &SecureScoreControlsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSecureScoresClient creates a new instance of SecureScoresClient.
func (c *ClientFactory) NewSecureScoresClient() *SecureScoresClient {
	return &SecureScoresClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSensitivitySettingsClient creates a new instance of SensitivitySettingsClient.
func (c *ClientFactory) NewSensitivitySettingsClient() *SensitivitySettingsClient {
	return &SensitivitySettingsClient{
		internal: c.internal,
	}
}

// NewServerVulnerabilityAssessmentClient creates a new instance of ServerVulnerabilityAssessmentClient.
func (c *ClientFactory) NewServerVulnerabilityAssessmentClient() *ServerVulnerabilityAssessmentClient {
	return &ServerVulnerabilityAssessmentClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewServerVulnerabilityAssessmentsSettingsClient creates a new instance of ServerVulnerabilityAssessmentsSettingsClient.
func (c *ClientFactory) NewServerVulnerabilityAssessmentsSettingsClient() *ServerVulnerabilityAssessmentsSettingsClient {
	return &ServerVulnerabilityAssessmentsSettingsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSettingsClient creates a new instance of SettingsClient.
func (c *ClientFactory) NewSettingsClient() *SettingsClient {
	return &SettingsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSoftwareInventoriesClient creates a new instance of SoftwareInventoriesClient.
func (c *ClientFactory) NewSoftwareInventoriesClient() *SoftwareInventoriesClient {
	return &SoftwareInventoriesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSolutionsClient creates a new instance of SolutionsClient.
func (c *ClientFactory) NewSolutionsClient() *SolutionsClient {
	return &SolutionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSolutionsReferenceDataClient creates a new instance of SolutionsReferenceDataClient.
func (c *ClientFactory) NewSolutionsReferenceDataClient() *SolutionsReferenceDataClient {
	return &SolutionsReferenceDataClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewStandardAssignmentsClient creates a new instance of StandardAssignmentsClient.
func (c *ClientFactory) NewStandardAssignmentsClient() *StandardAssignmentsClient {
	return &StandardAssignmentsClient{
		internal: c.internal,
	}
}

// NewStandardsClient creates a new instance of StandardsClient.
func (c *ClientFactory) NewStandardsClient() *StandardsClient {
	return &StandardsClient{
		internal: c.internal,
	}
}

// NewSubAssessmentsClient creates a new instance of SubAssessmentsClient.
func (c *ClientFactory) NewSubAssessmentsClient() *SubAssessmentsClient {
	return &SubAssessmentsClient{
		internal: c.internal,
	}
}

// NewTasksClient creates a new instance of TasksClient.
func (c *ClientFactory) NewTasksClient() *TasksClient {
	return &TasksClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewTopologyClient creates a new instance of TopologyClient.
func (c *ClientFactory) NewTopologyClient() *TopologyClient {
	return &TopologyClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspaceSettingsClient creates a new instance of WorkspaceSettingsClient.
func (c *ClientFactory) NewWorkspaceSettingsClient() *WorkspaceSettingsClient {
	return &WorkspaceSettingsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
