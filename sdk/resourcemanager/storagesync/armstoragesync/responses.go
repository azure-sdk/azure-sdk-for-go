// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragesync

// CloudEndpointsClientAfsShareMetadataCertificatePublicKeysResponse contains the response from method CloudEndpointsClient.AfsShareMetadataCertificatePublicKeys.
type CloudEndpointsClientAfsShareMetadataCertificatePublicKeysResponse struct {
	// Cloud endpoint AFS file share metadata signing certificate public keys.
	CloudEndpointAfsShareMetadataCertificatePublicKeys

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// CloudEndpointsClientCreateResponse contains the response from method CloudEndpointsClient.BeginCreate.
type CloudEndpointsClientCreateResponse struct {
	// Cloud Endpoint object.
	CloudEndpoint
}

// CloudEndpointsClientDeleteResponse contains the response from method CloudEndpointsClient.BeginDelete.
type CloudEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// CloudEndpointsClientGetResponse contains the response from method CloudEndpointsClient.Get.
type CloudEndpointsClientGetResponse struct {
	// Cloud Endpoint object.
	CloudEndpoint

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// CloudEndpointsClientListBySyncGroupResponse contains the response from method CloudEndpointsClient.NewListBySyncGroupPager.
type CloudEndpointsClientListBySyncGroupResponse struct {
	// Array of CloudEndpoint
	CloudEndpointArray

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// CloudEndpointsClientPostBackupResponse contains the response from method CloudEndpointsClient.BeginPostBackup.
type CloudEndpointsClientPostBackupResponse struct {
	// Post Backup Response
	PostBackupResponse
}

// CloudEndpointsClientPostRestoreResponse contains the response from method CloudEndpointsClient.BeginPostRestore.
type CloudEndpointsClientPostRestoreResponse struct {
	// placeholder for future response values
}

// CloudEndpointsClientPreBackupResponse contains the response from method CloudEndpointsClient.BeginPreBackup.
type CloudEndpointsClientPreBackupResponse struct {
	// placeholder for future response values
}

// CloudEndpointsClientPreRestoreResponse contains the response from method CloudEndpointsClient.BeginPreRestore.
type CloudEndpointsClientPreRestoreResponse struct {
	// placeholder for future response values
}

// CloudEndpointsClientRestoreheartbeatResponse contains the response from method CloudEndpointsClient.Restoreheartbeat.
type CloudEndpointsClientRestoreheartbeatResponse struct {
	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// CloudEndpointsClientTriggerChangeDetectionResponse contains the response from method CloudEndpointsClient.BeginTriggerChangeDetection.
type CloudEndpointsClientTriggerChangeDetectionResponse struct {
	// placeholder for future response values
}

// MicrosoftStorageSyncClientLocationOperationStatusResponse contains the response from method MicrosoftStorageSyncClient.LocationOperationStatus.
type MicrosoftStorageSyncClientLocationOperationStatusResponse struct {
	// Operation status object
	LocationOperationStatus

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// OperationStatusClientGetResponse contains the response from method OperationStatusClient.Get.
type OperationStatusClientGetResponse struct {
	// Operation status object
	OperationStatus

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// The list of storage sync operations.
	OperationEntityListResult

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// PrivateEndpointConnectionsClientCreateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreate.
type PrivateEndpointConnectionsClientCreateResponse struct {
	// The private endpoint connection resource.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// The private endpoint connection resource.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByStorageSyncServiceResponse contains the response from method PrivateEndpointConnectionsClient.NewListByStorageSyncServicePager.
type PrivateEndpointConnectionsClientListByStorageSyncServiceResponse struct {
	// List of private endpoint connections associated with the specified resource.
	PrivateEndpointConnectionListResult

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// PrivateLinkResourcesClientListByStorageSyncServiceResponse contains the response from method PrivateLinkResourcesClient.ListByStorageSyncService.
type PrivateLinkResourcesClientListByStorageSyncServiceResponse struct {
	// A list of private link resources.
	PrivateLinkResourceListResult
}

// RegisteredServersClientCreateResponse contains the response from method RegisteredServersClient.BeginCreate.
type RegisteredServersClientCreateResponse struct {
	// Registered Server resource.
	RegisteredServer
}

// RegisteredServersClientDeleteResponse contains the response from method RegisteredServersClient.BeginDelete.
type RegisteredServersClientDeleteResponse struct {
	// placeholder for future response values
}

// RegisteredServersClientGetResponse contains the response from method RegisteredServersClient.Get.
type RegisteredServersClientGetResponse struct {
	// Registered Server resource.
	RegisteredServer

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// RegisteredServersClientListByStorageSyncServiceResponse contains the response from method RegisteredServersClient.NewListByStorageSyncServicePager.
type RegisteredServersClientListByStorageSyncServiceResponse struct {
	// Array of RegisteredServer
	RegisteredServerArray

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// RegisteredServersClientTriggerRolloverResponse contains the response from method RegisteredServersClient.BeginTriggerRollover.
type RegisteredServersClientTriggerRolloverResponse struct {
	// placeholder for future response values
}

// RegisteredServersClientUpdateResponse contains the response from method RegisteredServersClient.BeginUpdate.
type RegisteredServersClientUpdateResponse struct {
	// Registered Server resource.
	RegisteredServer
}

// ServerEndpointsClientCreateResponse contains the response from method ServerEndpointsClient.BeginCreate.
type ServerEndpointsClientCreateResponse struct {
	// Server Endpoint object.
	ServerEndpoint
}

// ServerEndpointsClientDeleteResponse contains the response from method ServerEndpointsClient.BeginDelete.
type ServerEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// ServerEndpointsClientGetResponse contains the response from method ServerEndpointsClient.Get.
type ServerEndpointsClientGetResponse struct {
	// Server Endpoint object.
	ServerEndpoint

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ServerEndpointsClientListBySyncGroupResponse contains the response from method ServerEndpointsClient.NewListBySyncGroupPager.
type ServerEndpointsClientListBySyncGroupResponse struct {
	// Array of ServerEndpoint
	ServerEndpointArray

	// Location contains the information returned from the Location header response.
	Location *string

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ServerEndpointsClientRecallActionResponse contains the response from method ServerEndpointsClient.BeginRecallAction.
type ServerEndpointsClientRecallActionResponse struct {
	// placeholder for future response values
}

// ServerEndpointsClientUpdateResponse contains the response from method ServerEndpointsClient.BeginUpdate.
type ServerEndpointsClientUpdateResponse struct {
	// Server Endpoint object.
	ServerEndpoint
}

// ServicesClientCheckNameAvailabilityResponse contains the response from method ServicesClient.CheckNameAvailability.
type ServicesClientCheckNameAvailabilityResponse struct {
	// The CheckNameAvailability operation response.
	CheckNameAvailabilityResult
}

// ServicesClientCreateResponse contains the response from method ServicesClient.BeginCreate.
type ServicesClientCreateResponse struct {
	// Storage Sync Service object.
	Service
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.BeginDelete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	// Storage Sync Service object.
	Service

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ServicesClientListByResourceGroupResponse contains the response from method ServicesClient.NewListByResourceGroupPager.
type ServicesClientListByResourceGroupResponse struct {
	// Array of StorageSyncServices
	ServiceArray

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ServicesClientListBySubscriptionResponse contains the response from method ServicesClient.NewListBySubscriptionPager.
type ServicesClientListBySubscriptionResponse struct {
	// Array of StorageSyncServices
	ServiceArray

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.BeginUpdate.
type ServicesClientUpdateResponse struct {
	// Storage Sync Service object.
	Service
}

// SyncGroupsClientCreateResponse contains the response from method SyncGroupsClient.Create.
type SyncGroupsClientCreateResponse struct {
	// Sync Group object.
	SyncGroup

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// SyncGroupsClientDeleteResponse contains the response from method SyncGroupsClient.Delete.
type SyncGroupsClientDeleteResponse struct {
	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// SyncGroupsClientGetResponse contains the response from method SyncGroupsClient.Get.
type SyncGroupsClientGetResponse struct {
	// Sync Group object.
	SyncGroup

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// SyncGroupsClientListByStorageSyncServiceResponse contains the response from method SyncGroupsClient.NewListByStorageSyncServicePager.
type SyncGroupsClientListByStorageSyncServiceResponse struct {
	// Array of SyncGroup
	SyncGroupArray

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// WorkflowsClientAbortResponse contains the response from method WorkflowsClient.Abort.
type WorkflowsClientAbortResponse struct {
	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// WorkflowsClientGetResponse contains the response from method WorkflowsClient.Get.
type WorkflowsClientGetResponse struct {
	// Workflow resource.
	Workflow

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// WorkflowsClientListByStorageSyncServiceResponse contains the response from method WorkflowsClient.NewListByStorageSyncServicePager.
type WorkflowsClientListByStorageSyncServiceResponse struct {
	// Array of Workflow
	WorkflowArray

	// XMSCorrelationRequestID contains the information returned from the x-ms-correlation-request-id header response.
	XMSCorrelationRequestID *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}
