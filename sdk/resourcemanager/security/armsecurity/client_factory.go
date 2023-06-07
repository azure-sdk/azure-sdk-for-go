//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewMdeOnboardingsClient() *MdeOnboardingsClient {
	subClient, _ := NewMdeOnboardingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCustomAssessmentAutomationsClient() *CustomAssessmentAutomationsClient {
	subClient, _ := NewCustomAssessmentAutomationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCustomEntityStoreAssignmentsClient() *CustomEntityStoreAssignmentsClient {
	subClient, _ := NewCustomEntityStoreAssignmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewComplianceResultsClient() *ComplianceResultsClient {
	subClient, _ := NewComplianceResultsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPricingsClient() *PricingsClient {
	subClient, _ := NewPricingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAdvancedThreatProtectionClient() *AdvancedThreatProtectionClient {
	subClient, _ := NewAdvancedThreatProtectionClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDeviceSecurityGroupsClient() *DeviceSecurityGroupsClient {
	subClient, _ := NewDeviceSecurityGroupsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIotSecuritySolutionClient() *IotSecuritySolutionClient {
	subClient, _ := NewIotSecuritySolutionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIotSecuritySolutionAnalyticsClient() *IotSecuritySolutionAnalyticsClient {
	subClient, _ := NewIotSecuritySolutionAnalyticsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIotSecuritySolutionsAnalyticsAggregatedAlertClient() *IotSecuritySolutionsAnalyticsAggregatedAlertClient {
	subClient, _ := NewIotSecuritySolutionsAnalyticsAggregatedAlertClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIotSecuritySolutionsAnalyticsRecommendationClient() *IotSecuritySolutionsAnalyticsRecommendationClient {
	subClient, _ := NewIotSecuritySolutionsAnalyticsRecommendationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLocationsClient() *LocationsClient {
	subClient, _ := NewLocationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTasksClient() *TasksClient {
	subClient, _ := NewTasksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAutoProvisioningSettingsClient() *AutoProvisioningSettingsClient {
	subClient, _ := NewAutoProvisioningSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCompliancesClient() *CompliancesClient {
	subClient, _ := NewCompliancesClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewInformationProtectionPoliciesClient() *InformationProtectionPoliciesClient {
	subClient, _ := NewInformationProtectionPoliciesClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewContactsClient() *ContactsClient {
	subClient, _ := NewContactsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkspaceSettingsClient() *WorkspaceSettingsClient {
	subClient, _ := NewWorkspaceSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegulatoryComplianceStandardsClient() *RegulatoryComplianceStandardsClient {
	subClient, _ := NewRegulatoryComplianceStandardsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegulatoryComplianceControlsClient() *RegulatoryComplianceControlsClient {
	subClient, _ := NewRegulatoryComplianceControlsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegulatoryComplianceAssessmentsClient() *RegulatoryComplianceAssessmentsClient {
	subClient, _ := NewRegulatoryComplianceAssessmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSubAssessmentsClient() *SubAssessmentsClient {
	subClient, _ := NewSubAssessmentsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAutomationsClient() *AutomationsClient {
	subClient, _ := NewAutomationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAlertsSuppressionRulesClient() *AlertsSuppressionRulesClient {
	subClient, _ := NewAlertsSuppressionRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewServerVulnerabilityAssessmentClient() *ServerVulnerabilityAssessmentClient {
	subClient, _ := NewServerVulnerabilityAssessmentClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAssessmentsMetadataClient() *AssessmentsMetadataClient {
	subClient, _ := NewAssessmentsMetadataClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAssessmentsClient() *AssessmentsClient {
	subClient, _ := NewAssessmentsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAdaptiveApplicationControlsClient() *AdaptiveApplicationControlsClient {
	subClient, _ := NewAdaptiveApplicationControlsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAdaptiveNetworkHardeningsClient() *AdaptiveNetworkHardeningsClient {
	subClient, _ := NewAdaptiveNetworkHardeningsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAllowedConnectionsClient() *AllowedConnectionsClient {
	subClient, _ := NewAllowedConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTopologyClient() *TopologyClient {
	subClient, _ := NewTopologyClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJitNetworkAccessPoliciesClient() *JitNetworkAccessPoliciesClient {
	subClient, _ := NewJitNetworkAccessPoliciesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDiscoveredSecuritySolutionsClient() *DiscoveredSecuritySolutionsClient {
	subClient, _ := NewDiscoveredSecuritySolutionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSolutionsReferenceDataClient() *SolutionsReferenceDataClient {
	subClient, _ := NewSolutionsReferenceDataClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewExternalSecuritySolutionsClient() *ExternalSecuritySolutionsClient {
	subClient, _ := NewExternalSecuritySolutionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSecureScoresClient() *SecureScoresClient {
	subClient, _ := NewSecureScoresClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSecureScoreControlsClient() *SecureScoreControlsClient {
	subClient, _ := NewSecureScoreControlsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSecureScoreControlDefinitionsClient() *SecureScoreControlDefinitionsClient {
	subClient, _ := NewSecureScoreControlDefinitionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSolutionsClient() *SolutionsClient {
	subClient, _ := NewSolutionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAccountConnectorsClient() *AccountConnectorsClient {
	subClient, _ := NewAccountConnectorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCenterClient() *CenterClient {
	subClient, _ := NewCenterClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSensitivitySettingsClient() *SensitivitySettingsClient {
	subClient, _ := NewSensitivitySettingsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAlertsClient() *AlertsClient {
	subClient, _ := NewAlertsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSettingsClient() *SettingsClient {
	subClient, _ := NewSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIngestionSettingsClient() *IngestionSettingsClient {
	subClient, _ := NewIngestionSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSoftwareInventoriesClient() *SoftwareInventoriesClient {
	subClient, _ := NewSoftwareInventoriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGovernanceRulesClient() *GovernanceRulesClient {
	subClient, _ := NewGovernanceRulesClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGovernanceAssignmentsClient() *GovernanceAssignmentsClient {
	subClient, _ := NewGovernanceAssignmentsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewApplicationsClient() *ApplicationsClient {
	subClient, _ := NewApplicationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewApplicationClient() *ApplicationClient {
	subClient, _ := NewApplicationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConnectorApplicationsClient() *ConnectorApplicationsClient {
	subClient, _ := NewConnectorApplicationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConnectorApplicationClient() *ConnectorApplicationClient {
	subClient, _ := NewConnectorApplicationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPICollectionClient() *APICollectionClient {
	subClient, _ := NewAPICollectionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPICollectionOnboardingClient() *APICollectionOnboardingClient {
	subClient, _ := NewAPICollectionOnboardingClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPICollectionOffboardingClient() *APICollectionOffboardingClient {
	subClient, _ := NewAPICollectionOffboardingClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSQLVulnerabilityAssessmentScansClient() *SQLVulnerabilityAssessmentScansClient {
	subClient, _ := NewSQLVulnerabilityAssessmentScansClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSQLVulnerabilityAssessmentScanResultsClient() *SQLVulnerabilityAssessmentScanResultsClient {
	subClient, _ := NewSQLVulnerabilityAssessmentScanResultsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSQLVulnerabilityAssessmentBaselineRulesClient() *SQLVulnerabilityAssessmentBaselineRulesClient {
	subClient, _ := NewSQLVulnerabilityAssessmentBaselineRulesClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConnectorsClient() *ConnectorsClient {
	subClient, _ := NewConnectorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperatorsClient() *OperatorsClient {
	subClient, _ := NewOperatorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewHealthReportsClient() *HealthReportsClient {
	subClient, _ := NewHealthReportsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDefenderForStorageClient() *DefenderForStorageClient {
	subClient, _ := NewDefenderForStorageClient(c.credential, c.options)
	return subClient
}
