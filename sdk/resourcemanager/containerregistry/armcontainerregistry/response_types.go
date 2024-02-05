//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

// AgentPoolsClientCreateResponse contains the response from method AgentPoolsClient.BeginCreate.
type AgentPoolsClientCreateResponse struct {
	// The agentpool that has the ARM resource and properties.
	// The agentpool will have all information to create an agent pool.
	AgentPool
}

// AgentPoolsClientDeleteResponse contains the response from method AgentPoolsClient.BeginDelete.
type AgentPoolsClientDeleteResponse struct {
	// placeholder for future response values
}

// AgentPoolsClientGetQueueStatusResponse contains the response from method AgentPoolsClient.GetQueueStatus.
type AgentPoolsClientGetQueueStatusResponse struct {
	// The QueueStatus of Agent Pool
	AgentPoolQueueStatus
}

// AgentPoolsClientGetResponse contains the response from method AgentPoolsClient.Get.
type AgentPoolsClientGetResponse struct {
	// The agentpool that has the ARM resource and properties.
	// The agentpool will have all information to create an agent pool.
	AgentPool
}

// AgentPoolsClientListResponse contains the response from method AgentPoolsClient.NewListPager.
type AgentPoolsClientListResponse struct {
	// The collection of agent pools.
	AgentPoolListResult
}

// AgentPoolsClientUpdateResponse contains the response from method AgentPoolsClient.BeginUpdate.
type AgentPoolsClientUpdateResponse struct {
	// The agentpool that has the ARM resource and properties.
	// The agentpool will have all information to create an agent pool.
	AgentPool
}

// ArchiveVersionsClientCreateResponse contains the response from method ArchiveVersionsClient.BeginCreate.
type ArchiveVersionsClientCreateResponse struct {
	// An object that represents an export pipeline for a container registry.
	ArchiveVersion
}

// ArchiveVersionsClientDeleteResponse contains the response from method ArchiveVersionsClient.BeginDelete.
type ArchiveVersionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ArchiveVersionsClientGetResponse contains the response from method ArchiveVersionsClient.Get.
type ArchiveVersionsClientGetResponse struct {
	// An object that represents an export pipeline for a container registry.
	ArchiveVersion
}

// ArchiveVersionsClientListResponse contains the response from method ArchiveVersionsClient.NewListPager.
type ArchiveVersionsClientListResponse struct {
	// The result of a request to list export pipelines for a container registry.
	ArchiveVersionListResult
}

// ArchivesClientCreateResponse contains the response from method ArchivesClient.BeginCreate.
type ArchivesClientCreateResponse struct {
	// An object that represents a archive for a container registry.
	Archive
}

// ArchivesClientDeleteResponse contains the response from method ArchivesClient.BeginDelete.
type ArchivesClientDeleteResponse struct {
	// placeholder for future response values
}

// ArchivesClientGetResponse contains the response from method ArchivesClient.Get.
type ArchivesClientGetResponse struct {
	// An object that represents a archive for a container registry.
	Archive
}

// ArchivesClientListResponse contains the response from method ArchivesClient.NewListPager.
type ArchivesClientListResponse struct {
	// The result of a request to list archives for a container registry.
	ArchiveListResult
}

// ArchivesClientUpdateResponse contains the response from method ArchivesClient.Update.
type ArchivesClientUpdateResponse struct {
	// An object that represents a archive for a container registry.
	Archive
}

// CacheRulesClientCreateResponse contains the response from method CacheRulesClient.BeginCreate.
type CacheRulesClientCreateResponse struct {
	// An object that represents a cache rule for a container registry.
	CacheRule
}

// CacheRulesClientDeleteResponse contains the response from method CacheRulesClient.BeginDelete.
type CacheRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// CacheRulesClientGetResponse contains the response from method CacheRulesClient.Get.
type CacheRulesClientGetResponse struct {
	// An object that represents a cache rule for a container registry.
	CacheRule
}

// CacheRulesClientListResponse contains the response from method CacheRulesClient.NewListPager.
type CacheRulesClientListResponse struct {
	// The result of a request to list cache rules for a container registry.
	CacheRulesListResult
}

// CacheRulesClientUpdateResponse contains the response from method CacheRulesClient.BeginUpdate.
type CacheRulesClientUpdateResponse struct {
	// An object that represents a cache rule for a container registry.
	CacheRule
}

// ConnectedRegistriesClientCreateResponse contains the response from method ConnectedRegistriesClient.BeginCreate.
type ConnectedRegistriesClientCreateResponse struct {
	// An object that represents a connected registry for a container registry.
	ConnectedRegistry
}

// ConnectedRegistriesClientDeactivateResponse contains the response from method ConnectedRegistriesClient.BeginDeactivate.
type ConnectedRegistriesClientDeactivateResponse struct {
	// placeholder for future response values
}

// ConnectedRegistriesClientDeleteResponse contains the response from method ConnectedRegistriesClient.BeginDelete.
type ConnectedRegistriesClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectedRegistriesClientGetResponse contains the response from method ConnectedRegistriesClient.Get.
type ConnectedRegistriesClientGetResponse struct {
	// An object that represents a connected registry for a container registry.
	ConnectedRegistry
}

// ConnectedRegistriesClientListResponse contains the response from method ConnectedRegistriesClient.NewListPager.
type ConnectedRegistriesClientListResponse struct {
	// The result of a request to list connected registries for a container registry.
	ConnectedRegistryListResult
}

// ConnectedRegistriesClientUpdateResponse contains the response from method ConnectedRegistriesClient.BeginUpdate.
type ConnectedRegistriesClientUpdateResponse struct {
	// An object that represents a connected registry for a container registry.
	ConnectedRegistry
}

// CredentialSetsClientCreateResponse contains the response from method CredentialSetsClient.BeginCreate.
type CredentialSetsClientCreateResponse struct {
	// An object that represents a credential set resource for a container registry.
	CredentialSet
}

// CredentialSetsClientDeleteResponse contains the response from method CredentialSetsClient.BeginDelete.
type CredentialSetsClientDeleteResponse struct {
	// placeholder for future response values
}

// CredentialSetsClientGetResponse contains the response from method CredentialSetsClient.Get.
type CredentialSetsClientGetResponse struct {
	// An object that represents a credential set resource for a container registry.
	CredentialSet
}

// CredentialSetsClientListResponse contains the response from method CredentialSetsClient.NewListPager.
type CredentialSetsClientListResponse struct {
	// The result of a request to list credential sets for a container registry.
	CredentialSetListResult
}

// CredentialSetsClientUpdateResponse contains the response from method CredentialSetsClient.BeginUpdate.
type CredentialSetsClientUpdateResponse struct {
	// An object that represents a credential set resource for a container registry.
	CredentialSet
}

// ExportPipelinesClientCreateResponse contains the response from method ExportPipelinesClient.BeginCreate.
type ExportPipelinesClientCreateResponse struct {
	// An object that represents an export pipeline for a container registry.
	ExportPipeline
}

// ExportPipelinesClientDeleteResponse contains the response from method ExportPipelinesClient.BeginDelete.
type ExportPipelinesClientDeleteResponse struct {
	// placeholder for future response values
}

// ExportPipelinesClientGetResponse contains the response from method ExportPipelinesClient.Get.
type ExportPipelinesClientGetResponse struct {
	// An object that represents an export pipeline for a container registry.
	ExportPipeline
}

// ExportPipelinesClientListResponse contains the response from method ExportPipelinesClient.NewListPager.
type ExportPipelinesClientListResponse struct {
	// The result of a request to list export pipelines for a container registry.
	ExportPipelineListResult
}

// ImportPipelinesClientCreateResponse contains the response from method ImportPipelinesClient.BeginCreate.
type ImportPipelinesClientCreateResponse struct {
	// An object that represents an import pipeline for a container registry.
	ImportPipeline
}

// ImportPipelinesClientDeleteResponse contains the response from method ImportPipelinesClient.BeginDelete.
type ImportPipelinesClientDeleteResponse struct {
	// placeholder for future response values
}

// ImportPipelinesClientGetResponse contains the response from method ImportPipelinesClient.Get.
type ImportPipelinesClientGetResponse struct {
	// An object that represents an import pipeline for a container registry.
	ImportPipeline
}

// ImportPipelinesClientListResponse contains the response from method ImportPipelinesClient.NewListPager.
type ImportPipelinesClientListResponse struct {
	// The result of a request to list import pipelines for a container registry.
	ImportPipelineListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// The result of a request to list container registry operations.
	OperationListResult
}

// PipelineRunsClientCreateResponse contains the response from method PipelineRunsClient.BeginCreate.
type PipelineRunsClientCreateResponse struct {
	// An object that represents a pipeline run for a container registry.
	PipelineRun
}

// PipelineRunsClientDeleteResponse contains the response from method PipelineRunsClient.BeginDelete.
type PipelineRunsClientDeleteResponse struct {
	// placeholder for future response values
}

// PipelineRunsClientGetResponse contains the response from method PipelineRunsClient.Get.
type PipelineRunsClientGetResponse struct {
	// An object that represents a pipeline run for a container registry.
	PipelineRun
}

// PipelineRunsClientListResponse contains the response from method PipelineRunsClient.NewListPager.
type PipelineRunsClientListResponse struct {
	// The result of a request to list pipeline runs for a container registry.
	PipelineRunListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	// An object that represents a private endpoint connection for a container registry.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// An object that represents a private endpoint connection for a container registry.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.NewListPager.
type PrivateEndpointConnectionsClientListResponse struct {
	// The result of a request to list private endpoint connections for a container registry.
	PrivateEndpointConnectionListResult
}

// RegistriesClientCheckCacheRuleArtifactSyncEstimateResponse contains the response from method RegistriesClient.CheckCacheRuleArtifactSyncEstimate.
type RegistriesClientCheckCacheRuleArtifactSyncEstimateResponse struct {
	// The result of a request to check the cache rule artifact sync estimate
	CacheRuleArtifactSyncEstimateResult
}

// RegistriesClientCheckNameAvailabilityResponse contains the response from method RegistriesClient.CheckNameAvailability.
type RegistriesClientCheckNameAvailabilityResponse struct {
	// The result of a request to check the availability of a container registry name.
	RegistryNameStatus
}

// RegistriesClientCreateResponse contains the response from method RegistriesClient.BeginCreate.
type RegistriesClientCreateResponse struct {
	// An object that represents a container registry.
	Registry
}

// RegistriesClientDeleteResponse contains the response from method RegistriesClient.BeginDelete.
type RegistriesClientDeleteResponse struct {
	// placeholder for future response values
}

// RegistriesClientGenerateCredentialsResponse contains the response from method RegistriesClient.BeginGenerateCredentials.
type RegistriesClientGenerateCredentialsResponse struct {
	// The response from the GenerateCredentials operation.
	GenerateCredentialsResult
}

// RegistriesClientGetBuildSourceUploadURLResponse contains the response from method RegistriesClient.GetBuildSourceUploadURL.
type RegistriesClientGetBuildSourceUploadURLResponse struct {
	// The properties of a response to source upload request.
	SourceUploadDefinition
}

// RegistriesClientGetPrivateLinkResourceResponse contains the response from method RegistriesClient.GetPrivateLinkResource.
type RegistriesClientGetPrivateLinkResourceResponse struct {
	// A resource that supports private link capabilities.
	PrivateLinkResource
}

// RegistriesClientGetResponse contains the response from method RegistriesClient.Get.
type RegistriesClientGetResponse struct {
	// An object that represents a container registry.
	Registry
}

// RegistriesClientImportImageResponse contains the response from method RegistriesClient.BeginImportImage.
type RegistriesClientImportImageResponse struct {
	// placeholder for future response values
}

// RegistriesClientListByResourceGroupResponse contains the response from method RegistriesClient.NewListByResourceGroupPager.
type RegistriesClientListByResourceGroupResponse struct {
	// The result of a request to list container registries.
	RegistryListResult
}

// RegistriesClientListCredentialsResponse contains the response from method RegistriesClient.ListCredentials.
type RegistriesClientListCredentialsResponse struct {
	// The response from the ListCredentials operation.
	RegistryListCredentialsResult
}

// RegistriesClientListPrivateLinkResourcesResponse contains the response from method RegistriesClient.NewListPrivateLinkResourcesPager.
type RegistriesClientListPrivateLinkResourcesResponse struct {
	// The result of a request to list private link resources for a container registry.
	PrivateLinkResourceListResult
}

// RegistriesClientListResponse contains the response from method RegistriesClient.NewListPager.
type RegistriesClientListResponse struct {
	// The result of a request to list container registries.
	RegistryListResult
}

// RegistriesClientListUsagesResponse contains the response from method RegistriesClient.ListUsages.
type RegistriesClientListUsagesResponse struct {
	// The result of a request to get container registry quota usages.
	RegistryUsageListResult
}

// RegistriesClientRegenerateCredentialResponse contains the response from method RegistriesClient.RegenerateCredential.
type RegistriesClientRegenerateCredentialResponse struct {
	// The response from the ListCredentials operation.
	RegistryListCredentialsResult
}

// RegistriesClientScheduleRunResponse contains the response from method RegistriesClient.BeginScheduleRun.
type RegistriesClientScheduleRunResponse struct {
	// Run resource properties
	Run
}

// RegistriesClientUpdateResponse contains the response from method RegistriesClient.BeginUpdate.
type RegistriesClientUpdateResponse struct {
	// An object that represents a container registry.
	Registry
}

// ReplicationsClientCreateResponse contains the response from method ReplicationsClient.BeginCreate.
type ReplicationsClientCreateResponse struct {
	// An object that represents a replication for a container registry.
	Replication
}

// ReplicationsClientDeleteResponse contains the response from method ReplicationsClient.BeginDelete.
type ReplicationsClientDeleteResponse struct {
	// placeholder for future response values
}

// ReplicationsClientGetResponse contains the response from method ReplicationsClient.Get.
type ReplicationsClientGetResponse struct {
	// An object that represents a replication for a container registry.
	Replication
}

// ReplicationsClientListResponse contains the response from method ReplicationsClient.NewListPager.
type ReplicationsClientListResponse struct {
	// The result of a request to list replications for a container registry.
	ReplicationListResult
}

// ReplicationsClientUpdateResponse contains the response from method ReplicationsClient.BeginUpdate.
type ReplicationsClientUpdateResponse struct {
	// An object that represents a replication for a container registry.
	Replication
}

// RunsClientCancelResponse contains the response from method RunsClient.BeginCancel.
type RunsClientCancelResponse struct {
	// placeholder for future response values
}

// RunsClientGetLogSasURLResponse contains the response from method RunsClient.GetLogSasURL.
type RunsClientGetLogSasURLResponse struct {
	// The result of get log link operation.
	RunGetLogResult
}

// RunsClientGetResponse contains the response from method RunsClient.Get.
type RunsClientGetResponse struct {
	// Run resource properties
	Run
}

// RunsClientListResponse contains the response from method RunsClient.NewListPager.
type RunsClientListResponse struct {
	// Collection of runs.
	RunListResult
}

// RunsClientUpdateResponse contains the response from method RunsClient.BeginUpdate.
type RunsClientUpdateResponse struct {
	// Run resource properties
	Run
}

// ScopeMapsClientCreateResponse contains the response from method ScopeMapsClient.BeginCreate.
type ScopeMapsClientCreateResponse struct {
	// An object that represents a scope map for a container registry.
	ScopeMap
}

// ScopeMapsClientDeleteResponse contains the response from method ScopeMapsClient.BeginDelete.
type ScopeMapsClientDeleteResponse struct {
	// placeholder for future response values
}

// ScopeMapsClientGetResponse contains the response from method ScopeMapsClient.Get.
type ScopeMapsClientGetResponse struct {
	// An object that represents a scope map for a container registry.
	ScopeMap
}

// ScopeMapsClientListResponse contains the response from method ScopeMapsClient.NewListPager.
type ScopeMapsClientListResponse struct {
	// The result of a request to list scope maps for a container registry.
	ScopeMapListResult
}

// ScopeMapsClientUpdateResponse contains the response from method ScopeMapsClient.BeginUpdate.
type ScopeMapsClientUpdateResponse struct {
	// An object that represents a scope map for a container registry.
	ScopeMap
}

// TaskRunsClientCreateResponse contains the response from method TaskRunsClient.BeginCreate.
type TaskRunsClientCreateResponse struct {
	// The task run that has the ARM resource and properties.
	// The task run will have the information of request and result of a run.
	TaskRun
}

// TaskRunsClientDeleteResponse contains the response from method TaskRunsClient.BeginDelete.
type TaskRunsClientDeleteResponse struct {
	// placeholder for future response values
}

// TaskRunsClientGetDetailsResponse contains the response from method TaskRunsClient.GetDetails.
type TaskRunsClientGetDetailsResponse struct {
	// The task run that has the ARM resource and properties.
	// The task run will have the information of request and result of a run.
	TaskRun
}

// TaskRunsClientGetResponse contains the response from method TaskRunsClient.Get.
type TaskRunsClientGetResponse struct {
	// The task run that has the ARM resource and properties.
	// The task run will have the information of request and result of a run.
	TaskRun
}

// TaskRunsClientListResponse contains the response from method TaskRunsClient.NewListPager.
type TaskRunsClientListResponse struct {
	// The collection of task runs.
	TaskRunListResult
}

// TaskRunsClientUpdateResponse contains the response from method TaskRunsClient.BeginUpdate.
type TaskRunsClientUpdateResponse struct {
	// The task run that has the ARM resource and properties.
	// The task run will have the information of request and result of a run.
	TaskRun
}

// TasksClientCreateResponse contains the response from method TasksClient.BeginCreate.
type TasksClientCreateResponse struct {
	// The task that has the ARM resource and task properties.
	// The task will have all information to schedule a run against it.
	Task
}

// TasksClientDeleteResponse contains the response from method TasksClient.BeginDelete.
type TasksClientDeleteResponse struct {
	// placeholder for future response values
}

// TasksClientGetDetailsResponse contains the response from method TasksClient.GetDetails.
type TasksClientGetDetailsResponse struct {
	// The task that has the ARM resource and task properties.
	// The task will have all information to schedule a run against it.
	Task
}

// TasksClientGetResponse contains the response from method TasksClient.Get.
type TasksClientGetResponse struct {
	// The task that has the ARM resource and task properties.
	// The task will have all information to schedule a run against it.
	Task
}

// TasksClientListResponse contains the response from method TasksClient.NewListPager.
type TasksClientListResponse struct {
	// The collection of tasks.
	TaskListResult
}

// TasksClientUpdateResponse contains the response from method TasksClient.BeginUpdate.
type TasksClientUpdateResponse struct {
	// The task that has the ARM resource and task properties.
	// The task will have all information to schedule a run against it.
	Task
}

// TokensClientCreateResponse contains the response from method TokensClient.BeginCreate.
type TokensClientCreateResponse struct {
	// An object that represents a token for a container registry.
	Token
}

// TokensClientDeleteResponse contains the response from method TokensClient.BeginDelete.
type TokensClientDeleteResponse struct {
	// placeholder for future response values
}

// TokensClientGetResponse contains the response from method TokensClient.Get.
type TokensClientGetResponse struct {
	// An object that represents a token for a container registry.
	Token
}

// TokensClientListResponse contains the response from method TokensClient.NewListPager.
type TokensClientListResponse struct {
	// The result of a request to list tokens for a container registry.
	TokenListResult
}

// TokensClientUpdateResponse contains the response from method TokensClient.BeginUpdate.
type TokensClientUpdateResponse struct {
	// An object that represents a token for a container registry.
	Token
}

// WebhooksClientCreateResponse contains the response from method WebhooksClient.BeginCreate.
type WebhooksClientCreateResponse struct {
	// An object that represents a webhook for a container registry.
	Webhook
}

// WebhooksClientDeleteResponse contains the response from method WebhooksClient.BeginDelete.
type WebhooksClientDeleteResponse struct {
	// placeholder for future response values
}

// WebhooksClientGetCallbackConfigResponse contains the response from method WebhooksClient.GetCallbackConfig.
type WebhooksClientGetCallbackConfigResponse struct {
	// The configuration of service URI and custom headers for the webhook.
	CallbackConfig
}

// WebhooksClientGetResponse contains the response from method WebhooksClient.Get.
type WebhooksClientGetResponse struct {
	// An object that represents a webhook for a container registry.
	Webhook
}

// WebhooksClientListEventsResponse contains the response from method WebhooksClient.NewListEventsPager.
type WebhooksClientListEventsResponse struct {
	// The result of a request to list events for a webhook.
	EventListResult
}

// WebhooksClientListResponse contains the response from method WebhooksClient.NewListPager.
type WebhooksClientListResponse struct {
	// The result of a request to list webhooks for a container registry.
	WebhookListResult
}

// WebhooksClientPingResponse contains the response from method WebhooksClient.Ping.
type WebhooksClientPingResponse struct {
	// The basic information of an event.
	EventInfo
}

// WebhooksClientUpdateResponse contains the response from method WebhooksClient.BeginUpdate.
type WebhooksClientUpdateResponse struct {
	// An object that represents a webhook for a container registry.
	Webhook
}
