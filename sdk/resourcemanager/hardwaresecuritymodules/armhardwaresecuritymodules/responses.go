// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhardwaresecuritymodules

// CloudHsmClusterBackupStatusClientGetResponse contains the response from method CloudHsmClusterBackupStatusClient.Get.
type CloudHsmClusterBackupStatusClientGetResponse struct {
	// Backup operation Result
	BackupResult

	// Location contains the information returned from the Location header response.
	Location *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// CloudHsmClusterPrivateEndpointConnectionsClientCreateResponse contains the response from method CloudHsmClusterPrivateEndpointConnectionsClient.Create.
type CloudHsmClusterPrivateEndpointConnectionsClientCreateResponse struct {
	// The private endpoint connection resource.
	PrivateEndpointConnection
}

// CloudHsmClusterPrivateEndpointConnectionsClientDeleteResponse contains the response from method CloudHsmClusterPrivateEndpointConnectionsClient.BeginDelete.
type CloudHsmClusterPrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// CloudHsmClusterPrivateEndpointConnectionsClientGetResponse contains the response from method CloudHsmClusterPrivateEndpointConnectionsClient.Get.
type CloudHsmClusterPrivateEndpointConnectionsClientGetResponse struct {
	// The private endpoint connection resource.
	PrivateEndpointConnection
}

// CloudHsmClusterPrivateLinkResourcesClientListByCloudHsmClusterResponse contains the response from method CloudHsmClusterPrivateLinkResourcesClient.NewListByCloudHsmClusterPager.
type CloudHsmClusterPrivateLinkResourcesClientListByCloudHsmClusterResponse struct {
	// A list of private link resources.
	PrivateLinkResourceListResult
}

// CloudHsmClusterRestoreStatusClientGetResponse contains the response from method CloudHsmClusterRestoreStatusClient.Get.
type CloudHsmClusterRestoreStatusClientGetResponse struct {
	// Restore operation properties
	RestoreResult

	// Location contains the information returned from the Location header response.
	Location *string

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// CloudHsmClustersClientBackupResponse contains the response from method CloudHsmClustersClient.BeginBackup.
type CloudHsmClustersClientBackupResponse struct {
	// Backup operation Result
	BackupResult
}

// CloudHsmClustersClientCreateOrUpdateResponse contains the response from method CloudHsmClustersClient.BeginCreateOrUpdate.
type CloudHsmClustersClientCreateOrUpdateResponse struct {
	// Resource information with extended details.
	CloudHsmCluster
}

// CloudHsmClustersClientDeleteResponse contains the response from method CloudHsmClustersClient.BeginDelete.
type CloudHsmClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// CloudHsmClustersClientGetResponse contains the response from method CloudHsmClustersClient.Get.
type CloudHsmClustersClientGetResponse struct {
	// Resource information with extended details.
	CloudHsmCluster
}

// CloudHsmClustersClientListByResourceGroupResponse contains the response from method CloudHsmClustersClient.NewListByResourceGroupPager.
type CloudHsmClustersClientListByResourceGroupResponse struct {
	// List of Cloud HSM Clusters
	CloudHsmClusterListResult
}

// CloudHsmClustersClientListBySubscriptionResponse contains the response from method CloudHsmClustersClient.NewListBySubscriptionPager.
type CloudHsmClustersClientListBySubscriptionResponse struct {
	// List of Cloud HSM Clusters
	CloudHsmClusterListResult
}

// CloudHsmClustersClientRestoreResponse contains the response from method CloudHsmClustersClient.BeginRestore.
type CloudHsmClustersClientRestoreResponse struct {
	// Restore operation properties
	RestoreResult
}

// CloudHsmClustersClientUpdateResponse contains the response from method CloudHsmClustersClient.BeginUpdate.
type CloudHsmClustersClientUpdateResponse struct {
	// Resource information with extended details.
	CloudHsmCluster
}

// CloudHsmClustersClientValidateBackupPropertiesResponse contains the response from method CloudHsmClustersClient.BeginValidateBackupProperties.
type CloudHsmClustersClientValidateBackupPropertiesResponse struct {
	// Backup operation Result
	BackupResult
}

// CloudHsmClustersClientValidateRestorePropertiesResponse contains the response from method CloudHsmClustersClient.BeginValidateRestoreProperties.
type CloudHsmClustersClientValidateRestorePropertiesResponse struct {
	// Restore operation properties
	RestoreResult
}

// DedicatedHsmClientCreateOrUpdateResponse contains the response from method DedicatedHsmClient.BeginCreateOrUpdate.
type DedicatedHsmClientCreateOrUpdateResponse struct {
	// Resource information with extended details.
	DedicatedHsm
}

// DedicatedHsmClientDeleteResponse contains the response from method DedicatedHsmClient.BeginDelete.
type DedicatedHsmClientDeleteResponse struct {
	// placeholder for future response values
}

// DedicatedHsmClientGetResponse contains the response from method DedicatedHsmClient.Get.
type DedicatedHsmClientGetResponse struct {
	// Resource information with extended details.
	DedicatedHsm
}

// DedicatedHsmClientListByResourceGroupResponse contains the response from method DedicatedHsmClient.NewListByResourceGroupPager.
type DedicatedHsmClientListByResourceGroupResponse struct {
	// List of dedicated HSMs
	DedicatedHsmListResult
}

// DedicatedHsmClientListBySubscriptionResponse contains the response from method DedicatedHsmClient.NewListBySubscriptionPager.
type DedicatedHsmClientListBySubscriptionResponse struct {
	// List of dedicated HSMs
	DedicatedHsmListResult
}

// DedicatedHsmClientListOutboundNetworkDependenciesEndpointsResponse contains the response from method DedicatedHsmClient.NewListOutboundNetworkDependenciesEndpointsPager.
type DedicatedHsmClientListOutboundNetworkDependenciesEndpointsResponse struct {
	// Collection of OutboundEnvironmentEndpoint
	OutboundEnvironmentEndpointCollection
}

// DedicatedHsmClientUpdateResponse contains the response from method DedicatedHsmClient.BeginUpdate.
type DedicatedHsmClientUpdateResponse struct {
	// Resource information with extended details.
	DedicatedHsm
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// PrivateEndpointConnectionsClientListByCloudHsmClusterResponse contains the response from method PrivateEndpointConnectionsClient.NewListByCloudHsmClusterPager.
type PrivateEndpointConnectionsClientListByCloudHsmClusterResponse struct {
	// List of private endpoint connections associated with the specified resource.
	PrivateEndpointConnectionListResult
}
