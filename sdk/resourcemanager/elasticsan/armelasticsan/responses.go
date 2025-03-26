// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armelasticsan

// ElasticSansClientCreateResponse contains the response from method ElasticSansClient.BeginCreate.
type ElasticSansClientCreateResponse struct {
	// Response for ElasticSan request.
	ElasticSan
}

// ElasticSansClientDeleteResponse contains the response from method ElasticSansClient.BeginDelete.
type ElasticSansClientDeleteResponse struct {
	// placeholder for future response values
}

// ElasticSansClientGetResponse contains the response from method ElasticSansClient.Get.
type ElasticSansClientGetResponse struct {
	// Response for ElasticSan request.
	ElasticSan
}

// ElasticSansClientListByResourceGroupResponse contains the response from method ElasticSansClient.NewListByResourceGroupPager.
type ElasticSansClientListByResourceGroupResponse struct {
	// List of Elastic Sans
	List
}

// ElasticSansClientListBySubscriptionResponse contains the response from method ElasticSansClient.NewListBySubscriptionPager.
type ElasticSansClientListBySubscriptionResponse struct {
	// List of Elastic Sans
	List
}

// ElasticSansClientUpdateResponse contains the response from method ElasticSansClient.BeginUpdate.
type ElasticSansClientUpdateResponse struct {
	// Response for ElasticSan request.
	ElasticSan
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreate.
type PrivateEndpointConnectionsClientCreateResponse struct {
	// Response for PrivateEndpoint Connection object
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// Response for PrivateEndpoint Connection object
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.NewListPager.
type PrivateEndpointConnectionsClientListResponse struct {
	// List of private endpoint connections associated with SAN
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientListByElasticSanResponse contains the response from method PrivateLinkResourcesClient.ListByElasticSan.
type PrivateLinkResourcesClientListByElasticSanResponse struct {
	// A list of private link resources
	PrivateLinkResourceListResult
}

// SKUsClientListResponse contains the response from method SKUsClient.NewListPager.
type SKUsClientListResponse struct {
	// List of SKU Information objects
	SKUInformationList
}

// VolumeGroupsClientCreateResponse contains the response from method VolumeGroupsClient.BeginCreate.
type VolumeGroupsClientCreateResponse struct {
	// Response for Volume Group request.
	VolumeGroup
}

// VolumeGroupsClientDeleteResponse contains the response from method VolumeGroupsClient.BeginDelete.
type VolumeGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// VolumeGroupsClientGetResponse contains the response from method VolumeGroupsClient.Get.
type VolumeGroupsClientGetResponse struct {
	// Response for Volume Group request.
	VolumeGroup
}

// VolumeGroupsClientListByElasticSanResponse contains the response from method VolumeGroupsClient.NewListByElasticSanPager.
type VolumeGroupsClientListByElasticSanResponse struct {
	// List of Volume Groups
	VolumeGroupList
}

// VolumeGroupsClientPreBackupResponse contains the response from method VolumeGroupsClient.BeginPreBackup.
type VolumeGroupsClientPreBackupResponse struct {
	// response object for pre validation api
	PreValidationResponse
}

// VolumeGroupsClientPreRestoreResponse contains the response from method VolumeGroupsClient.BeginPreRestore.
type VolumeGroupsClientPreRestoreResponse struct {
	// response object for pre validation api
	PreValidationResponse
}

// VolumeGroupsClientUpdateResponse contains the response from method VolumeGroupsClient.BeginUpdate.
type VolumeGroupsClientUpdateResponse struct {
	// Response for Volume Group request.
	VolumeGroup
}

// VolumeSnapshotsClientCreateResponse contains the response from method VolumeSnapshotsClient.BeginCreate.
type VolumeSnapshotsClientCreateResponse struct {
	// Response for Volume Snapshot request.
	Snapshot
}

// VolumeSnapshotsClientDeleteResponse contains the response from method VolumeSnapshotsClient.BeginDelete.
type VolumeSnapshotsClientDeleteResponse struct {
	// placeholder for future response values
}

// VolumeSnapshotsClientGetResponse contains the response from method VolumeSnapshotsClient.Get.
type VolumeSnapshotsClientGetResponse struct {
	// Response for Volume Snapshot request.
	Snapshot
}

// VolumeSnapshotsClientListByVolumeGroupResponse contains the response from method VolumeSnapshotsClient.NewListByVolumeGroupPager.
type VolumeSnapshotsClientListByVolumeGroupResponse struct {
	// List of Snapshots
	SnapshotList
}

// VolumesClientCreateResponse contains the response from method VolumesClient.BeginCreate.
type VolumesClientCreateResponse struct {
	// Response for Volume request.
	Volume
}

// VolumesClientDeleteResponse contains the response from method VolumesClient.BeginDelete.
type VolumesClientDeleteResponse struct {
	// placeholder for future response values
}

// VolumesClientGetResponse contains the response from method VolumesClient.Get.
type VolumesClientGetResponse struct {
	// Response for Volume request.
	Volume
}

// VolumesClientListByVolumeGroupResponse contains the response from method VolumesClient.NewListByVolumeGroupPager.
type VolumesClientListByVolumeGroupResponse struct {
	// List of Volumes
	VolumeList
}

// VolumesClientUpdateResponse contains the response from method VolumesClient.BeginUpdate.
type VolumesClientUpdateResponse struct {
	// Response for Volume request.
	Volume
}
