// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armelasticsan

// ElasticSansClientBeginCreateOptions contains the optional parameters for the ElasticSansClient.BeginCreate method.
type ElasticSansClientBeginCreateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ElasticSansClientBeginDeleteOptions contains the optional parameters for the ElasticSansClient.BeginDelete method.
type ElasticSansClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ElasticSansClientBeginUpdateOptions contains the optional parameters for the ElasticSansClient.BeginUpdate method.
type ElasticSansClientBeginUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ElasticSansClientGetOptions contains the optional parameters for the ElasticSansClient.Get method.
type ElasticSansClientGetOptions struct {
	// placeholder for future optional parameters
}

// ElasticSansClientListByResourceGroupOptions contains the optional parameters for the ElasticSansClient.NewListByResourceGroupPager
// method.
type ElasticSansClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ElasticSansClientListBySubscriptionOptions contains the optional parameters for the ElasticSansClient.NewListBySubscriptionPager
// method.
type ElasticSansClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientBeginCreateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginCreate
// method.
type PrivateEndpointConnectionsClientBeginCreateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
type PrivateEndpointConnectionsClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListPager
// method.
type PrivateEndpointConnectionsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListByElasticSanOptions contains the optional parameters for the PrivateLinkResourcesClient.ListByElasticSan
// method.
type PrivateLinkResourcesClientListByElasticSanOptions struct {
	// placeholder for future optional parameters
}

// SKUsClientListOptions contains the optional parameters for the SKUsClient.NewListPager method.
type SKUsClientListOptions struct {
	// Specify $filter='location eq ' to filter on location.
	Filter *string
}

// VolumeGroupsClientBeginCreateOptions contains the optional parameters for the VolumeGroupsClient.BeginCreate method.
type VolumeGroupsClientBeginCreateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// VolumeGroupsClientBeginDeleteOptions contains the optional parameters for the VolumeGroupsClient.BeginDelete method.
type VolumeGroupsClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// VolumeGroupsClientBeginPreBackupOptions contains the optional parameters for the VolumeGroupsClient.BeginPreBackup method.
type VolumeGroupsClientBeginPreBackupOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// VolumeGroupsClientBeginPreRestoreOptions contains the optional parameters for the VolumeGroupsClient.BeginPreRestore method.
type VolumeGroupsClientBeginPreRestoreOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// VolumeGroupsClientBeginUpdateOptions contains the optional parameters for the VolumeGroupsClient.BeginUpdate method.
type VolumeGroupsClientBeginUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// VolumeGroupsClientGetOptions contains the optional parameters for the VolumeGroupsClient.Get method.
type VolumeGroupsClientGetOptions struct {
	// placeholder for future optional parameters
}

// VolumeGroupsClientListByElasticSanOptions contains the optional parameters for the VolumeGroupsClient.NewListByElasticSanPager
// method.
type VolumeGroupsClientListByElasticSanOptions struct {
	// placeholder for future optional parameters
}

// VolumeSnapshotsClientBeginCreateOptions contains the optional parameters for the VolumeSnapshotsClient.BeginCreate method.
type VolumeSnapshotsClientBeginCreateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// VolumeSnapshotsClientBeginDeleteOptions contains the optional parameters for the VolumeSnapshotsClient.BeginDelete method.
type VolumeSnapshotsClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// VolumeSnapshotsClientGetOptions contains the optional parameters for the VolumeSnapshotsClient.Get method.
type VolumeSnapshotsClientGetOptions struct {
	// placeholder for future optional parameters
}

// VolumeSnapshotsClientListByVolumeGroupOptions contains the optional parameters for the VolumeSnapshotsClient.NewListByVolumeGroupPager
// method.
type VolumeSnapshotsClientListByVolumeGroupOptions struct {
	// Specify $filter='volumeName eq ' to filter on volume.
	Filter *string
}

// VolumesClientBeginCreateOptions contains the optional parameters for the VolumesClient.BeginCreate method.
type VolumesClientBeginCreateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// VolumesClientBeginDeleteOptions contains the optional parameters for the VolumesClient.BeginDelete method.
type VolumesClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string

	// Optional, used to delete snapshots under volume. Allowed value are only true or false. Default value is false.
	XMSDeleteSnapshots *XMSDeleteSnapshots

	// Optional, used to delete volume if active sessions present. Allowed value are only true or false. Default value is false.
	XMSForceDelete *XMSForceDelete
}

// VolumesClientBeginUpdateOptions contains the optional parameters for the VolumesClient.BeginUpdate method.
type VolumesClientBeginUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// VolumesClientGetOptions contains the optional parameters for the VolumesClient.Get method.
type VolumesClientGetOptions struct {
	// placeholder for future optional parameters
}

// VolumesClientListByVolumeGroupOptions contains the optional parameters for the VolumesClient.NewListByVolumeGroupPager
// method.
type VolumesClientListByVolumeGroupOptions struct {
	// placeholder for future optional parameters
}
