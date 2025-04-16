// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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
//   - subscriptionID - The ID of the target subscription.
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

// NewAzureADOnlyAuthenticationsClient creates a new instance of AzureADOnlyAuthenticationsClient.
func (c *ClientFactory) NewAzureADOnlyAuthenticationsClient() *AzureADOnlyAuthenticationsClient {
	return &AzureADOnlyAuthenticationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBigDataPoolsClient creates a new instance of BigDataPoolsClient.
func (c *ClientFactory) NewBigDataPoolsClient() *BigDataPoolsClient {
	return &BigDataPoolsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDataMaskingPoliciesClient creates a new instance of DataMaskingPoliciesClient.
func (c *ClientFactory) NewDataMaskingPoliciesClient() *DataMaskingPoliciesClient {
	return &DataMaskingPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDataMaskingRulesClient creates a new instance of DataMaskingRulesClient.
func (c *ClientFactory) NewDataMaskingRulesClient() *DataMaskingRulesClient {
	return &DataMaskingRulesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewExtendedSQLPoolBlobAuditingPoliciesClient creates a new instance of ExtendedSQLPoolBlobAuditingPoliciesClient.
func (c *ClientFactory) NewExtendedSQLPoolBlobAuditingPoliciesClient() *ExtendedSQLPoolBlobAuditingPoliciesClient {
	return &ExtendedSQLPoolBlobAuditingPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGetClient creates a new instance of GetClient.
func (c *ClientFactory) NewGetClient() *GetClient {
	return &GetClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIPFirewallRulesClient creates a new instance of IPFirewallRulesClient.
func (c *ClientFactory) NewIPFirewallRulesClient() *IPFirewallRulesClient {
	return &IPFirewallRulesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIntegrationRuntimeAuthKeysClient creates a new instance of IntegrationRuntimeAuthKeysClient.
func (c *ClientFactory) NewIntegrationRuntimeAuthKeysClient() *IntegrationRuntimeAuthKeysClient {
	return &IntegrationRuntimeAuthKeysClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIntegrationRuntimeConnectionInfosClient creates a new instance of IntegrationRuntimeConnectionInfosClient.
func (c *ClientFactory) NewIntegrationRuntimeConnectionInfosClient() *IntegrationRuntimeConnectionInfosClient {
	return &IntegrationRuntimeConnectionInfosClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIntegrationRuntimeCredentialsClient creates a new instance of IntegrationRuntimeCredentialsClient.
func (c *ClientFactory) NewIntegrationRuntimeCredentialsClient() *IntegrationRuntimeCredentialsClient {
	return &IntegrationRuntimeCredentialsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIntegrationRuntimeMonitoringDataClient creates a new instance of IntegrationRuntimeMonitoringDataClient.
func (c *ClientFactory) NewIntegrationRuntimeMonitoringDataClient() *IntegrationRuntimeMonitoringDataClient {
	return &IntegrationRuntimeMonitoringDataClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIntegrationRuntimeNodeIPAddressClient creates a new instance of IntegrationRuntimeNodeIPAddressClient.
func (c *ClientFactory) NewIntegrationRuntimeNodeIPAddressClient() *IntegrationRuntimeNodeIPAddressClient {
	return &IntegrationRuntimeNodeIPAddressClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIntegrationRuntimeNodesClient creates a new instance of IntegrationRuntimeNodesClient.
func (c *ClientFactory) NewIntegrationRuntimeNodesClient() *IntegrationRuntimeNodesClient {
	return &IntegrationRuntimeNodesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIntegrationRuntimeObjectMetadataClient creates a new instance of IntegrationRuntimeObjectMetadataClient.
func (c *ClientFactory) NewIntegrationRuntimeObjectMetadataClient() *IntegrationRuntimeObjectMetadataClient {
	return &IntegrationRuntimeObjectMetadataClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIntegrationRuntimeStatusClient creates a new instance of IntegrationRuntimeStatusClient.
func (c *ClientFactory) NewIntegrationRuntimeStatusClient() *IntegrationRuntimeStatusClient {
	return &IntegrationRuntimeStatusClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIntegrationRuntimesClient creates a new instance of IntegrationRuntimesClient.
func (c *ClientFactory) NewIntegrationRuntimesClient() *IntegrationRuntimesClient {
	return &IntegrationRuntimesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewKeysClient creates a new instance of KeysClient.
func (c *ClientFactory) NewKeysClient() *KeysClient {
	return &KeysClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewKustoOperationsClient creates a new instance of KustoOperationsClient.
func (c *ClientFactory) NewKustoOperationsClient() *KustoOperationsClient {
	return &KustoOperationsClient{
		internal: c.internal,
	}
}

// NewKustoPoolAttachedDatabaseConfigurationsClient creates a new instance of KustoPoolAttachedDatabaseConfigurationsClient.
func (c *ClientFactory) NewKustoPoolAttachedDatabaseConfigurationsClient() *KustoPoolAttachedDatabaseConfigurationsClient {
	return &KustoPoolAttachedDatabaseConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewKustoPoolChildResourceClient creates a new instance of KustoPoolChildResourceClient.
func (c *ClientFactory) NewKustoPoolChildResourceClient() *KustoPoolChildResourceClient {
	return &KustoPoolChildResourceClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewKustoPoolDataConnectionsClient creates a new instance of KustoPoolDataConnectionsClient.
func (c *ClientFactory) NewKustoPoolDataConnectionsClient() *KustoPoolDataConnectionsClient {
	return &KustoPoolDataConnectionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewKustoPoolDatabasePrincipalAssignmentsClient creates a new instance of KustoPoolDatabasePrincipalAssignmentsClient.
func (c *ClientFactory) NewKustoPoolDatabasePrincipalAssignmentsClient() *KustoPoolDatabasePrincipalAssignmentsClient {
	return &KustoPoolDatabasePrincipalAssignmentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewKustoPoolDatabasesClient creates a new instance of KustoPoolDatabasesClient.
func (c *ClientFactory) NewKustoPoolDatabasesClient() *KustoPoolDatabasesClient {
	return &KustoPoolDatabasesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewKustoPoolPrincipalAssignmentsClient creates a new instance of KustoPoolPrincipalAssignmentsClient.
func (c *ClientFactory) NewKustoPoolPrincipalAssignmentsClient() *KustoPoolPrincipalAssignmentsClient {
	return &KustoPoolPrincipalAssignmentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewKustoPoolPrivateLinkResourcesClient creates a new instance of KustoPoolPrivateLinkResourcesClient.
func (c *ClientFactory) NewKustoPoolPrivateLinkResourcesClient() *KustoPoolPrivateLinkResourcesClient {
	return &KustoPoolPrivateLinkResourcesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewKustoPoolsClient creates a new instance of KustoPoolsClient.
func (c *ClientFactory) NewKustoPoolsClient() *KustoPoolsClient {
	return &KustoPoolsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewLibrariesClient creates a new instance of LibrariesClient.
func (c *ClientFactory) NewLibrariesClient() *LibrariesClient {
	return &LibrariesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewLibraryClient creates a new instance of LibraryClient.
func (c *ClientFactory) NewLibraryClient() *LibraryClient {
	return &LibraryClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient.
func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	return &PrivateEndpointConnectionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateEndpointConnectionsPrivateLinkHubClient creates a new instance of PrivateEndpointConnectionsPrivateLinkHubClient.
func (c *ClientFactory) NewPrivateEndpointConnectionsPrivateLinkHubClient() *PrivateEndpointConnectionsPrivateLinkHubClient {
	return &PrivateEndpointConnectionsPrivateLinkHubClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateLinkHubPrivateLinkResourcesClient creates a new instance of PrivateLinkHubPrivateLinkResourcesClient.
func (c *ClientFactory) NewPrivateLinkHubPrivateLinkResourcesClient() *PrivateLinkHubPrivateLinkResourcesClient {
	return &PrivateLinkHubPrivateLinkResourcesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateLinkHubsClient creates a new instance of PrivateLinkHubsClient.
func (c *ClientFactory) NewPrivateLinkHubsClient() *PrivateLinkHubsClient {
	return &PrivateLinkHubsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient.
func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	return &PrivateLinkResourcesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewRestorableDroppedSQLPoolsClient creates a new instance of RestorableDroppedSQLPoolsClient.
func (c *ClientFactory) NewRestorableDroppedSQLPoolsClient() *RestorableDroppedSQLPoolsClient {
	return &RestorableDroppedSQLPoolsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolBlobAuditingPoliciesClient creates a new instance of SQLPoolBlobAuditingPoliciesClient.
func (c *ClientFactory) NewSQLPoolBlobAuditingPoliciesClient() *SQLPoolBlobAuditingPoliciesClient {
	return &SQLPoolBlobAuditingPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolColumnsClient creates a new instance of SQLPoolColumnsClient.
func (c *ClientFactory) NewSQLPoolColumnsClient() *SQLPoolColumnsClient {
	return &SQLPoolColumnsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolConnectionPoliciesClient creates a new instance of SQLPoolConnectionPoliciesClient.
func (c *ClientFactory) NewSQLPoolConnectionPoliciesClient() *SQLPoolConnectionPoliciesClient {
	return &SQLPoolConnectionPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolDataWarehouseUserActivitiesClient creates a new instance of SQLPoolDataWarehouseUserActivitiesClient.
func (c *ClientFactory) NewSQLPoolDataWarehouseUserActivitiesClient() *SQLPoolDataWarehouseUserActivitiesClient {
	return &SQLPoolDataWarehouseUserActivitiesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolGeoBackupPoliciesClient creates a new instance of SQLPoolGeoBackupPoliciesClient.
func (c *ClientFactory) NewSQLPoolGeoBackupPoliciesClient() *SQLPoolGeoBackupPoliciesClient {
	return &SQLPoolGeoBackupPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolMaintenanceWindowOptionsClient creates a new instance of SQLPoolMaintenanceWindowOptionsClient.
func (c *ClientFactory) NewSQLPoolMaintenanceWindowOptionsClient() *SQLPoolMaintenanceWindowOptionsClient {
	return &SQLPoolMaintenanceWindowOptionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolMaintenanceWindowsClient creates a new instance of SQLPoolMaintenanceWindowsClient.
func (c *ClientFactory) NewSQLPoolMaintenanceWindowsClient() *SQLPoolMaintenanceWindowsClient {
	return &SQLPoolMaintenanceWindowsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolMetadataSyncConfigsClient creates a new instance of SQLPoolMetadataSyncConfigsClient.
func (c *ClientFactory) NewSQLPoolMetadataSyncConfigsClient() *SQLPoolMetadataSyncConfigsClient {
	return &SQLPoolMetadataSyncConfigsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolOperationResultsClient creates a new instance of SQLPoolOperationResultsClient.
func (c *ClientFactory) NewSQLPoolOperationResultsClient() *SQLPoolOperationResultsClient {
	return &SQLPoolOperationResultsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolOperationsClient creates a new instance of SQLPoolOperationsClient.
func (c *ClientFactory) NewSQLPoolOperationsClient() *SQLPoolOperationsClient {
	return &SQLPoolOperationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolRecommendedSensitivityLabelsClient creates a new instance of SQLPoolRecommendedSensitivityLabelsClient.
func (c *ClientFactory) NewSQLPoolRecommendedSensitivityLabelsClient() *SQLPoolRecommendedSensitivityLabelsClient {
	return &SQLPoolRecommendedSensitivityLabelsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolReplicationLinksClient creates a new instance of SQLPoolReplicationLinksClient.
func (c *ClientFactory) NewSQLPoolReplicationLinksClient() *SQLPoolReplicationLinksClient {
	return &SQLPoolReplicationLinksClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolRestorePointsClient creates a new instance of SQLPoolRestorePointsClient.
func (c *ClientFactory) NewSQLPoolRestorePointsClient() *SQLPoolRestorePointsClient {
	return &SQLPoolRestorePointsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolSchemasClient creates a new instance of SQLPoolSchemasClient.
func (c *ClientFactory) NewSQLPoolSchemasClient() *SQLPoolSchemasClient {
	return &SQLPoolSchemasClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolSecurityAlertPoliciesClient creates a new instance of SQLPoolSecurityAlertPoliciesClient.
func (c *ClientFactory) NewSQLPoolSecurityAlertPoliciesClient() *SQLPoolSecurityAlertPoliciesClient {
	return &SQLPoolSecurityAlertPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolSensitivityLabelsClient creates a new instance of SQLPoolSensitivityLabelsClient.
func (c *ClientFactory) NewSQLPoolSensitivityLabelsClient() *SQLPoolSensitivityLabelsClient {
	return &SQLPoolSensitivityLabelsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolTableColumnsClient creates a new instance of SQLPoolTableColumnsClient.
func (c *ClientFactory) NewSQLPoolTableColumnsClient() *SQLPoolTableColumnsClient {
	return &SQLPoolTableColumnsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolTablesClient creates a new instance of SQLPoolTablesClient.
func (c *ClientFactory) NewSQLPoolTablesClient() *SQLPoolTablesClient {
	return &SQLPoolTablesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolTransparentDataEncryptionsClient creates a new instance of SQLPoolTransparentDataEncryptionsClient.
func (c *ClientFactory) NewSQLPoolTransparentDataEncryptionsClient() *SQLPoolTransparentDataEncryptionsClient {
	return &SQLPoolTransparentDataEncryptionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolUsagesClient creates a new instance of SQLPoolUsagesClient.
func (c *ClientFactory) NewSQLPoolUsagesClient() *SQLPoolUsagesClient {
	return &SQLPoolUsagesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolVulnerabilityAssessmentRuleBaselinesClient creates a new instance of SQLPoolVulnerabilityAssessmentRuleBaselinesClient.
func (c *ClientFactory) NewSQLPoolVulnerabilityAssessmentRuleBaselinesClient() *SQLPoolVulnerabilityAssessmentRuleBaselinesClient {
	return &SQLPoolVulnerabilityAssessmentRuleBaselinesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolVulnerabilityAssessmentScansClient creates a new instance of SQLPoolVulnerabilityAssessmentScansClient.
func (c *ClientFactory) NewSQLPoolVulnerabilityAssessmentScansClient() *SQLPoolVulnerabilityAssessmentScansClient {
	return &SQLPoolVulnerabilityAssessmentScansClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolVulnerabilityAssessmentsClient creates a new instance of SQLPoolVulnerabilityAssessmentsClient.
func (c *ClientFactory) NewSQLPoolVulnerabilityAssessmentsClient() *SQLPoolVulnerabilityAssessmentsClient {
	return &SQLPoolVulnerabilityAssessmentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolWorkloadClassifierClient creates a new instance of SQLPoolWorkloadClassifierClient.
func (c *ClientFactory) NewSQLPoolWorkloadClassifierClient() *SQLPoolWorkloadClassifierClient {
	return &SQLPoolWorkloadClassifierClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolWorkloadGroupClient creates a new instance of SQLPoolWorkloadGroupClient.
func (c *ClientFactory) NewSQLPoolWorkloadGroupClient() *SQLPoolWorkloadGroupClient {
	return &SQLPoolWorkloadGroupClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSQLPoolsClient creates a new instance of SQLPoolsClient.
func (c *ClientFactory) NewSQLPoolsClient() *SQLPoolsClient {
	return &SQLPoolsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSparkConfigurationClient creates a new instance of SparkConfigurationClient.
func (c *ClientFactory) NewSparkConfigurationClient() *SparkConfigurationClient {
	return &SparkConfigurationClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSparkConfigurationsClient creates a new instance of SparkConfigurationsClient.
func (c *ClientFactory) NewSparkConfigurationsClient() *SparkConfigurationsClient {
	return &SparkConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspaceAADAdminsClient creates a new instance of WorkspaceAADAdminsClient.
func (c *ClientFactory) NewWorkspaceAADAdminsClient() *WorkspaceAADAdminsClient {
	return &WorkspaceAADAdminsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspaceManagedIdentitySQLControlSettingsClient creates a new instance of WorkspaceManagedIdentitySQLControlSettingsClient.
func (c *ClientFactory) NewWorkspaceManagedIdentitySQLControlSettingsClient() *WorkspaceManagedIdentitySQLControlSettingsClient {
	return &WorkspaceManagedIdentitySQLControlSettingsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient creates a new instance of WorkspaceManagedSQLServerBlobAuditingPoliciesClient.
func (c *ClientFactory) NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient() *WorkspaceManagedSQLServerBlobAuditingPoliciesClient {
	return &WorkspaceManagedSQLServerBlobAuditingPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient creates a new instance of WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient.
func (c *ClientFactory) NewWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient() *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient {
	return &WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspaceManagedSQLServerEncryptionProtectorClient creates a new instance of WorkspaceManagedSQLServerEncryptionProtectorClient.
func (c *ClientFactory) NewWorkspaceManagedSQLServerEncryptionProtectorClient() *WorkspaceManagedSQLServerEncryptionProtectorClient {
	return &WorkspaceManagedSQLServerEncryptionProtectorClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient creates a new instance of WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient.
func (c *ClientFactory) NewWorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient() *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient {
	return &WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspaceManagedSQLServerRecoverableSQLPoolsClient creates a new instance of WorkspaceManagedSQLServerRecoverableSQLPoolsClient.
func (c *ClientFactory) NewWorkspaceManagedSQLServerRecoverableSQLPoolsClient() *WorkspaceManagedSQLServerRecoverableSQLPoolsClient {
	return &WorkspaceManagedSQLServerRecoverableSQLPoolsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspaceManagedSQLServerSecurityAlertPolicyClient creates a new instance of WorkspaceManagedSQLServerSecurityAlertPolicyClient.
func (c *ClientFactory) NewWorkspaceManagedSQLServerSecurityAlertPolicyClient() *WorkspaceManagedSQLServerSecurityAlertPolicyClient {
	return &WorkspaceManagedSQLServerSecurityAlertPolicyClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspaceManagedSQLServerUsagesClient creates a new instance of WorkspaceManagedSQLServerUsagesClient.
func (c *ClientFactory) NewWorkspaceManagedSQLServerUsagesClient() *WorkspaceManagedSQLServerUsagesClient {
	return &WorkspaceManagedSQLServerUsagesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspaceManagedSQLServerVulnerabilityAssessmentsClient creates a new instance of WorkspaceManagedSQLServerVulnerabilityAssessmentsClient.
func (c *ClientFactory) NewWorkspaceManagedSQLServerVulnerabilityAssessmentsClient() *WorkspaceManagedSQLServerVulnerabilityAssessmentsClient {
	return &WorkspaceManagedSQLServerVulnerabilityAssessmentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspaceSQLAADAdminsClient creates a new instance of WorkspaceSQLAADAdminsClient.
func (c *ClientFactory) NewWorkspaceSQLAADAdminsClient() *WorkspaceSQLAADAdminsClient {
	return &WorkspaceSQLAADAdminsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspacesClient creates a new instance of WorkspacesClient.
func (c *ClientFactory) NewWorkspacesClient() *WorkspacesClient {
	return &WorkspacesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
