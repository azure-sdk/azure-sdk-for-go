//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

// ActivityRunsClientQueryByPipelineRunResponse contains the response from method ActivityRunsClient.QueryByPipelineRun.
type ActivityRunsClientQueryByPipelineRunResponse struct {
	// A list activity runs.
	ActivityRunsQueryResponse
}

// ChangeDataCaptureClientCreateOrUpdateResponse contains the response from method ChangeDataCaptureClient.CreateOrUpdate.
type ChangeDataCaptureClientCreateOrUpdateResponse struct {
	// Change data capture resource type.
	ChangeDataCaptureResource
}

// ChangeDataCaptureClientDeleteResponse contains the response from method ChangeDataCaptureClient.Delete.
type ChangeDataCaptureClientDeleteResponse struct {
	// placeholder for future response values
}

// ChangeDataCaptureClientGetResponse contains the response from method ChangeDataCaptureClient.Get.
type ChangeDataCaptureClientGetResponse struct {
	// Change data capture resource type.
	ChangeDataCaptureResource
}

// ChangeDataCaptureClientListByFactoryResponse contains the response from method ChangeDataCaptureClient.NewListByFactoryPager.
type ChangeDataCaptureClientListByFactoryResponse struct {
	// A list of change data capture resources.
	ChangeDataCaptureListResponse
}

// ChangeDataCaptureClientStartResponse contains the response from method ChangeDataCaptureClient.Start.
type ChangeDataCaptureClientStartResponse struct {
	// placeholder for future response values
}

// ChangeDataCaptureClientStatusResponse contains the response from method ChangeDataCaptureClient.Status.
type ChangeDataCaptureClientStatusResponse struct {
	// Current status of the change data capture resource.
	Value *string
}

// ChangeDataCaptureClientStopResponse contains the response from method ChangeDataCaptureClient.Stop.
type ChangeDataCaptureClientStopResponse struct {
	// placeholder for future response values
}

// CredentialOperationsClientCreateOrUpdateResponse contains the response from method CredentialOperationsClient.CreateOrUpdate.
type CredentialOperationsClientCreateOrUpdateResponse struct {
	// Credential resource type.
	CredentialResource
}

// CredentialOperationsClientDeleteResponse contains the response from method CredentialOperationsClient.Delete.
type CredentialOperationsClientDeleteResponse struct {
	// placeholder for future response values
}

// CredentialOperationsClientGetResponse contains the response from method CredentialOperationsClient.Get.
type CredentialOperationsClientGetResponse struct {
	// Credential resource type.
	CredentialResource
}

// CredentialOperationsClientListByFactoryResponse contains the response from method CredentialOperationsClient.NewListByFactoryPager.
type CredentialOperationsClientListByFactoryResponse struct {
	// A list of credential resources.
	CredentialListResponse
}

// DataFlowDebugSessionClientAddDataFlowResponse contains the response from method DataFlowDebugSessionClient.AddDataFlow.
type DataFlowDebugSessionClientAddDataFlowResponse struct {
	// Response body structure for starting data flow debug session.
	AddDataFlowToDebugSessionResponse
}

// DataFlowDebugSessionClientCreateResponse contains the response from method DataFlowDebugSessionClient.BeginCreate.
type DataFlowDebugSessionClientCreateResponse struct {
	// Response body structure for creating data flow debug session.
	CreateDataFlowDebugSessionResponse
}

// DataFlowDebugSessionClientDeleteResponse contains the response from method DataFlowDebugSessionClient.Delete.
type DataFlowDebugSessionClientDeleteResponse struct {
	// placeholder for future response values
}

// DataFlowDebugSessionClientExecuteCommandResponse contains the response from method DataFlowDebugSessionClient.BeginExecuteCommand.
type DataFlowDebugSessionClientExecuteCommandResponse struct {
	// Response body structure of data flow result for data preview, statistics or expression preview.
	DataFlowDebugCommandResponse
}

// DataFlowDebugSessionClientQueryByFactoryResponse contains the response from method DataFlowDebugSessionClient.NewQueryByFactoryPager.
type DataFlowDebugSessionClientQueryByFactoryResponse struct {
	// A list of active debug sessions.
	QueryDataFlowDebugSessionsResponse
}

// DataFlowsClientCreateOrUpdateResponse contains the response from method DataFlowsClient.CreateOrUpdate.
type DataFlowsClientCreateOrUpdateResponse struct {
	// Data flow resource type.
	DataFlowResource
}

// DataFlowsClientDeleteResponse contains the response from method DataFlowsClient.Delete.
type DataFlowsClientDeleteResponse struct {
	// placeholder for future response values
}

// DataFlowsClientGetResponse contains the response from method DataFlowsClient.Get.
type DataFlowsClientGetResponse struct {
	// Data flow resource type.
	DataFlowResource
}

// DataFlowsClientListByFactoryResponse contains the response from method DataFlowsClient.NewListByFactoryPager.
type DataFlowsClientListByFactoryResponse struct {
	// A list of data flow resources.
	DataFlowListResponse
}

// DatasetsClientCreateOrUpdateResponse contains the response from method DatasetsClient.CreateOrUpdate.
type DatasetsClientCreateOrUpdateResponse struct {
	// Dataset resource type.
	DatasetResource
}

// DatasetsClientDeleteResponse contains the response from method DatasetsClient.Delete.
type DatasetsClientDeleteResponse struct {
	// placeholder for future response values
}

// DatasetsClientGetResponse contains the response from method DatasetsClient.Get.
type DatasetsClientGetResponse struct {
	// Dataset resource type.
	DatasetResource
}

// DatasetsClientListByFactoryResponse contains the response from method DatasetsClient.NewListByFactoryPager.
type DatasetsClientListByFactoryResponse struct {
	// A list of dataset resources.
	DatasetListResponse
}

// ExposureControlClientGetFeatureValueByFactoryResponse contains the response from method ExposureControlClient.GetFeatureValueByFactory.
type ExposureControlClientGetFeatureValueByFactoryResponse struct {
	// The exposure control response.
	ExposureControlResponse
}

// ExposureControlClientGetFeatureValueResponse contains the response from method ExposureControlClient.GetFeatureValue.
type ExposureControlClientGetFeatureValueResponse struct {
	// The exposure control response.
	ExposureControlResponse
}

// ExposureControlClientQueryFeatureValuesByFactoryResponse contains the response from method ExposureControlClient.QueryFeatureValuesByFactory.
type ExposureControlClientQueryFeatureValuesByFactoryResponse struct {
	// A list of exposure control feature values.
	ExposureControlBatchResponse
}

// FactoriesClientConfigureFactoryRepoResponse contains the response from method FactoriesClient.ConfigureFactoryRepo.
type FactoriesClientConfigureFactoryRepoResponse struct {
	// Factory resource type.
	Factory
}

// FactoriesClientCreateOrUpdateResponse contains the response from method FactoriesClient.CreateOrUpdate.
type FactoriesClientCreateOrUpdateResponse struct {
	// Factory resource type.
	Factory
}

// FactoriesClientDeleteResponse contains the response from method FactoriesClient.Delete.
type FactoriesClientDeleteResponse struct {
	// placeholder for future response values
}

// FactoriesClientGetDataPlaneAccessResponse contains the response from method FactoriesClient.GetDataPlaneAccess.
type FactoriesClientGetDataPlaneAccessResponse struct {
	// Get Data Plane read only token response definition.
	AccessPolicyResponse
}

// FactoriesClientGetGitHubAccessTokenResponse contains the response from method FactoriesClient.GetGitHubAccessToken.
type FactoriesClientGetGitHubAccessTokenResponse struct {
	// Get GitHub access token response definition.
	GitHubAccessTokenResponse
}

// FactoriesClientGetResponse contains the response from method FactoriesClient.Get.
type FactoriesClientGetResponse struct {
	// Factory resource type.
	Factory
}

// FactoriesClientListByResourceGroupResponse contains the response from method FactoriesClient.NewListByResourceGroupPager.
type FactoriesClientListByResourceGroupResponse struct {
	// A list of factory resources.
	FactoryListResponse
}

// FactoriesClientListResponse contains the response from method FactoriesClient.NewListPager.
type FactoriesClientListResponse struct {
	// A list of factory resources.
	FactoryListResponse
}

// FactoriesClientUpdateResponse contains the response from method FactoriesClient.Update.
type FactoriesClientUpdateResponse struct {
	// Factory resource type.
	Factory
}

// GlobalParametersClientCreateOrUpdateResponse contains the response from method GlobalParametersClient.CreateOrUpdate.
type GlobalParametersClientCreateOrUpdateResponse struct {
	// Global parameters resource type.
	GlobalParameterResource
}

// GlobalParametersClientDeleteResponse contains the response from method GlobalParametersClient.Delete.
type GlobalParametersClientDeleteResponse struct {
	// placeholder for future response values
}

// GlobalParametersClientGetResponse contains the response from method GlobalParametersClient.Get.
type GlobalParametersClientGetResponse struct {
	// Global parameters resource type.
	GlobalParameterResource
}

// GlobalParametersClientListByFactoryResponse contains the response from method GlobalParametersClient.NewListByFactoryPager.
type GlobalParametersClientListByFactoryResponse struct {
	// A list of Global parameters.
	GlobalParameterListResponse
}

// IntegrationRuntimeNodesClientDeleteResponse contains the response from method IntegrationRuntimeNodesClient.Delete.
type IntegrationRuntimeNodesClientDeleteResponse struct {
	// placeholder for future response values
}

// IntegrationRuntimeNodesClientGetIPAddressResponse contains the response from method IntegrationRuntimeNodesClient.GetIPAddress.
type IntegrationRuntimeNodesClientGetIPAddressResponse struct {
	// The IP address of self-hosted integration runtime node.
	IntegrationRuntimeNodeIPAddress
}

// IntegrationRuntimeNodesClientGetResponse contains the response from method IntegrationRuntimeNodesClient.Get.
type IntegrationRuntimeNodesClientGetResponse struct {
	// Properties of Self-hosted integration runtime node.
	SelfHostedIntegrationRuntimeNode
}

// IntegrationRuntimeNodesClientUpdateResponse contains the response from method IntegrationRuntimeNodesClient.Update.
type IntegrationRuntimeNodesClientUpdateResponse struct {
	// Properties of Self-hosted integration runtime node.
	SelfHostedIntegrationRuntimeNode
}

// IntegrationRuntimeObjectMetadataClientGetResponse contains the response from method IntegrationRuntimeObjectMetadataClient.Get.
type IntegrationRuntimeObjectMetadataClientGetResponse struct {
	// A list of SSIS object metadata.
	SsisObjectMetadataListResponse
}

// IntegrationRuntimeObjectMetadataClientRefreshResponse contains the response from method IntegrationRuntimeObjectMetadataClient.BeginRefresh.
type IntegrationRuntimeObjectMetadataClientRefreshResponse struct {
	// The status of the operation.
	SsisObjectMetadataStatusResponse
}

// IntegrationRuntimesClientCreateLinkedIntegrationRuntimeResponse contains the response from method IntegrationRuntimesClient.CreateLinkedIntegrationRuntime.
type IntegrationRuntimesClientCreateLinkedIntegrationRuntimeResponse struct {
	// Integration runtime status response.
	IntegrationRuntimeStatusResponse
}

// IntegrationRuntimesClientCreateOrUpdateResponse contains the response from method IntegrationRuntimesClient.CreateOrUpdate.
type IntegrationRuntimesClientCreateOrUpdateResponse struct {
	// Integration runtime resource type.
	IntegrationRuntimeResource
}

// IntegrationRuntimesClientDeleteResponse contains the response from method IntegrationRuntimesClient.Delete.
type IntegrationRuntimesClientDeleteResponse struct {
	// placeholder for future response values
}

// IntegrationRuntimesClientGetConnectionInfoResponse contains the response from method IntegrationRuntimesClient.GetConnectionInfo.
type IntegrationRuntimesClientGetConnectionInfoResponse struct {
	// Connection information for encrypting the on-premises data source credentials.
	IntegrationRuntimeConnectionInfo
}

// IntegrationRuntimesClientGetMonitoringDataResponse contains the response from method IntegrationRuntimesClient.GetMonitoringData.
type IntegrationRuntimesClientGetMonitoringDataResponse struct {
	// Get monitoring data response.
	IntegrationRuntimeMonitoringData
}

// IntegrationRuntimesClientGetResponse contains the response from method IntegrationRuntimesClient.Get.
type IntegrationRuntimesClientGetResponse struct {
	// Integration runtime resource type.
	IntegrationRuntimeResource
}

// IntegrationRuntimesClientGetStatusResponse contains the response from method IntegrationRuntimesClient.GetStatus.
type IntegrationRuntimesClientGetStatusResponse struct {
	// Integration runtime status response.
	IntegrationRuntimeStatusResponse
}

// IntegrationRuntimesClientListAuthKeysResponse contains the response from method IntegrationRuntimesClient.ListAuthKeys.
type IntegrationRuntimesClientListAuthKeysResponse struct {
	// The integration runtime authentication keys.
	IntegrationRuntimeAuthKeys
}

// IntegrationRuntimesClientListByFactoryResponse contains the response from method IntegrationRuntimesClient.NewListByFactoryPager.
type IntegrationRuntimesClientListByFactoryResponse struct {
	// A list of integration runtime resources.
	IntegrationRuntimeListResponse
}

// IntegrationRuntimesClientListOutboundNetworkDependenciesEndpointsResponse contains the response from method IntegrationRuntimesClient.ListOutboundNetworkDependenciesEndpoints.
type IntegrationRuntimesClientListOutboundNetworkDependenciesEndpointsResponse struct {
	// Azure-SSIS integration runtime outbound network dependency endpoints.
	IntegrationRuntimeOutboundNetworkDependenciesEndpointsResponse
}

// IntegrationRuntimesClientRegenerateAuthKeyResponse contains the response from method IntegrationRuntimesClient.RegenerateAuthKey.
type IntegrationRuntimesClientRegenerateAuthKeyResponse struct {
	// The integration runtime authentication keys.
	IntegrationRuntimeAuthKeys
}

// IntegrationRuntimesClientRemoveLinksResponse contains the response from method IntegrationRuntimesClient.RemoveLinks.
type IntegrationRuntimesClientRemoveLinksResponse struct {
	// placeholder for future response values
}

// IntegrationRuntimesClientStartResponse contains the response from method IntegrationRuntimesClient.BeginStart.
type IntegrationRuntimesClientStartResponse struct {
	// Integration runtime status response.
	IntegrationRuntimeStatusResponse
}

// IntegrationRuntimesClientStopResponse contains the response from method IntegrationRuntimesClient.BeginStop.
type IntegrationRuntimesClientStopResponse struct {
	// placeholder for future response values
}

// IntegrationRuntimesClientSyncCredentialsResponse contains the response from method IntegrationRuntimesClient.SyncCredentials.
type IntegrationRuntimesClientSyncCredentialsResponse struct {
	// placeholder for future response values
}

// IntegrationRuntimesClientUpdateResponse contains the response from method IntegrationRuntimesClient.Update.
type IntegrationRuntimesClientUpdateResponse struct {
	// Integration runtime resource type.
	IntegrationRuntimeResource
}

// IntegrationRuntimesClientUpgradeResponse contains the response from method IntegrationRuntimesClient.Upgrade.
type IntegrationRuntimesClientUpgradeResponse struct {
	// placeholder for future response values
}

// LinkedServicesClientCreateOrUpdateResponse contains the response from method LinkedServicesClient.CreateOrUpdate.
type LinkedServicesClientCreateOrUpdateResponse struct {
	// Linked service resource type.
	LinkedServiceResource
}

// LinkedServicesClientDeleteResponse contains the response from method LinkedServicesClient.Delete.
type LinkedServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// LinkedServicesClientGetResponse contains the response from method LinkedServicesClient.Get.
type LinkedServicesClientGetResponse struct {
	// Linked service resource type.
	LinkedServiceResource
}

// LinkedServicesClientListByFactoryResponse contains the response from method LinkedServicesClient.NewListByFactoryPager.
type LinkedServicesClientListByFactoryResponse struct {
	// A list of linked service resources.
	LinkedServiceListResponse
}

// ManagedPrivateEndpointsClientCreateOrUpdateResponse contains the response from method ManagedPrivateEndpointsClient.CreateOrUpdate.
type ManagedPrivateEndpointsClientCreateOrUpdateResponse struct {
	// Managed private endpoint resource type.
	ManagedPrivateEndpointResource
}

// ManagedPrivateEndpointsClientDeleteResponse contains the response from method ManagedPrivateEndpointsClient.Delete.
type ManagedPrivateEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagedPrivateEndpointsClientGetResponse contains the response from method ManagedPrivateEndpointsClient.Get.
type ManagedPrivateEndpointsClientGetResponse struct {
	// Managed private endpoint resource type.
	ManagedPrivateEndpointResource
}

// ManagedPrivateEndpointsClientListByFactoryResponse contains the response from method ManagedPrivateEndpointsClient.NewListByFactoryPager.
type ManagedPrivateEndpointsClientListByFactoryResponse struct {
	// A list of managed private endpoint resources.
	ManagedPrivateEndpointListResponse
}

// ManagedVirtualNetworksClientCreateOrUpdateResponse contains the response from method ManagedVirtualNetworksClient.CreateOrUpdate.
type ManagedVirtualNetworksClientCreateOrUpdateResponse struct {
	// Managed Virtual Network resource type.
	ManagedVirtualNetworkResource
}

// ManagedVirtualNetworksClientGetResponse contains the response from method ManagedVirtualNetworksClient.Get.
type ManagedVirtualNetworksClientGetResponse struct {
	// Managed Virtual Network resource type.
	ManagedVirtualNetworkResource
}

// ManagedVirtualNetworksClientListByFactoryResponse contains the response from method ManagedVirtualNetworksClient.NewListByFactoryPager.
type ManagedVirtualNetworksClientListByFactoryResponse struct {
	// A list of managed Virtual Network resources.
	ManagedVirtualNetworkListResponse
}

// PipelineRunsClientCancelResponse contains the response from method PipelineRunsClient.Cancel.
type PipelineRunsClientCancelResponse struct {
	// placeholder for future response values
}

// PipelineRunsClientGetResponse contains the response from method PipelineRunsClient.Get.
type PipelineRunsClientGetResponse struct {
	// Information about a pipeline run.
	PipelineRun
}

// PipelineRunsClientQueryByFactoryResponse contains the response from method PipelineRunsClient.QueryByFactory.
type PipelineRunsClientQueryByFactoryResponse struct {
	// A list pipeline runs.
	PipelineRunsQueryResponse
}

// PipelinesClientCreateOrUpdateResponse contains the response from method PipelinesClient.CreateOrUpdate.
type PipelinesClientCreateOrUpdateResponse struct {
	// Pipeline resource type.
	PipelineResource
}

// PipelinesClientCreateRunResponse contains the response from method PipelinesClient.CreateRun.
type PipelinesClientCreateRunResponse struct {
	// Response body with a run identifier.
	CreateRunResponse
}

// PipelinesClientDeleteResponse contains the response from method PipelinesClient.Delete.
type PipelinesClientDeleteResponse struct {
	// placeholder for future response values
}

// PipelinesClientGetResponse contains the response from method PipelinesClient.Get.
type PipelinesClientGetResponse struct {
	// Pipeline resource type.
	PipelineResource
}

// PipelinesClientListByFactoryResponse contains the response from method PipelinesClient.NewListByFactoryPager.
type PipelinesClientListByFactoryResponse struct {
	// A list of pipeline resources.
	PipelineListResponse
}

// PrivateEndPointConnectionsClientListByFactoryResponse contains the response from method PrivateEndPointConnectionsClient.NewListByFactoryPager.
type PrivateEndPointConnectionsClientListByFactoryResponse struct {
	// A list of linked service resources.
	PrivateEndpointConnectionListResponse
}

// PrivateEndpointConnectionClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionClient.CreateOrUpdate.
type PrivateEndpointConnectionClientCreateOrUpdateResponse struct {
	// Private Endpoint Connection ARM resource.
	PrivateEndpointConnectionResource
}

// PrivateEndpointConnectionClientDeleteResponse contains the response from method PrivateEndpointConnectionClient.Delete.
type PrivateEndpointConnectionClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionClientGetResponse contains the response from method PrivateEndpointConnectionClient.Get.
type PrivateEndpointConnectionClientGetResponse struct {
	// Private Endpoint Connection ARM resource.
	PrivateEndpointConnectionResource
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	// Wrapper for a collection of private link resources
	PrivateLinkResourcesWrapper
}

// TriggerRunsClientCancelResponse contains the response from method TriggerRunsClient.Cancel.
type TriggerRunsClientCancelResponse struct {
	// placeholder for future response values
}

// TriggerRunsClientQueryByFactoryResponse contains the response from method TriggerRunsClient.QueryByFactory.
type TriggerRunsClientQueryByFactoryResponse struct {
	// A list of trigger runs.
	TriggerRunsQueryResponse
}

// TriggerRunsClientRerunResponse contains the response from method TriggerRunsClient.Rerun.
type TriggerRunsClientRerunResponse struct {
	// placeholder for future response values
}

// TriggersClientCreateOrUpdateResponse contains the response from method TriggersClient.CreateOrUpdate.
type TriggersClientCreateOrUpdateResponse struct {
	// Trigger resource type.
	TriggerResource
}

// TriggersClientDeleteResponse contains the response from method TriggersClient.Delete.
type TriggersClientDeleteResponse struct {
	// placeholder for future response values
}

// TriggersClientGetEventSubscriptionStatusResponse contains the response from method TriggersClient.GetEventSubscriptionStatus.
type TriggersClientGetEventSubscriptionStatusResponse struct {
	// Defines the response of a trigger subscription operation.
	TriggerSubscriptionOperationStatus
}

// TriggersClientGetResponse contains the response from method TriggersClient.Get.
type TriggersClientGetResponse struct {
	// Trigger resource type.
	TriggerResource
}

// TriggersClientListByFactoryResponse contains the response from method TriggersClient.NewListByFactoryPager.
type TriggersClientListByFactoryResponse struct {
	// A list of trigger resources.
	TriggerListResponse
}

// TriggersClientQueryByFactoryResponse contains the response from method TriggersClient.QueryByFactory.
type TriggersClientQueryByFactoryResponse struct {
	// A query of triggers.
	TriggerQueryResponse
}

// TriggersClientStartResponse contains the response from method TriggersClient.BeginStart.
type TriggersClientStartResponse struct {
	// placeholder for future response values
}

// TriggersClientStopResponse contains the response from method TriggersClient.BeginStop.
type TriggersClientStopResponse struct {
	// placeholder for future response values
}

// TriggersClientSubscribeToEventsResponse contains the response from method TriggersClient.BeginSubscribeToEvents.
type TriggersClientSubscribeToEventsResponse struct {
	// Defines the response of a trigger subscription operation.
	TriggerSubscriptionOperationStatus
}

// TriggersClientUnsubscribeFromEventsResponse contains the response from method TriggersClient.BeginUnsubscribeFromEvents.
type TriggersClientUnsubscribeFromEventsResponse struct {
	// Defines the response of a trigger subscription operation.
	TriggerSubscriptionOperationStatus
}
