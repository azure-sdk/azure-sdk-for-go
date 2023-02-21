//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetapp

// AccountBackupsClientDeleteResponse contains the response from method AccountBackupsClient.BeginDelete.
type AccountBackupsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountBackupsClientGetResponse contains the response from method AccountBackupsClient.Get.
type AccountBackupsClientGetResponse struct {
	Backup
}

// AccountBackupsClientListResponse contains the response from method AccountBackupsClient.NewListPager.
type AccountBackupsClientListResponse struct {
	BackupsList
}

// AccountsClientCreateOrUpdateResponse contains the response from method AccountsClient.BeginCreateOrUpdate.
type AccountsClientCreateOrUpdateResponse struct {
	Account
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.BeginDelete.
type AccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	Account
}

// AccountsClientListBySubscriptionResponse contains the response from method AccountsClient.NewListBySubscriptionPager.
type AccountsClientListBySubscriptionResponse struct {
	AccountList
}

// AccountsClientListResponse contains the response from method AccountsClient.NewListPager.
type AccountsClientListResponse struct {
	AccountList
}

// AccountsClientRenewCredentialsResponse contains the response from method AccountsClient.BeginRenewCredentials.
type AccountsClientRenewCredentialsResponse struct {
	// placeholder for future response values
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.BeginUpdate.
type AccountsClientUpdateResponse struct {
	Account
}

// BackupPoliciesClientCreateResponse contains the response from method BackupPoliciesClient.BeginCreate.
type BackupPoliciesClientCreateResponse struct {
	BackupPolicy
}

// BackupPoliciesClientDeleteResponse contains the response from method BackupPoliciesClient.BeginDelete.
type BackupPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// BackupPoliciesClientGetResponse contains the response from method BackupPoliciesClient.Get.
type BackupPoliciesClientGetResponse struct {
	BackupPolicy
}

// BackupPoliciesClientListResponse contains the response from method BackupPoliciesClient.NewListPager.
type BackupPoliciesClientListResponse struct {
	BackupPoliciesList
}

// BackupPoliciesClientUpdateResponse contains the response from method BackupPoliciesClient.BeginUpdate.
type BackupPoliciesClientUpdateResponse struct {
	BackupPolicy
}

// BackupsClientCreateResponse contains the response from method BackupsClient.BeginCreate.
type BackupsClientCreateResponse struct {
	Backup
}

// BackupsClientDeleteResponse contains the response from method BackupsClient.BeginDelete.
type BackupsClientDeleteResponse struct {
	// placeholder for future response values
}

// BackupsClientGetResponse contains the response from method BackupsClient.Get.
type BackupsClientGetResponse struct {
	Backup
}

// BackupsClientGetStatusResponse contains the response from method BackupsClient.GetStatus.
type BackupsClientGetStatusResponse struct {
	BackupStatus
}

// BackupsClientGetVolumeRestoreStatusResponse contains the response from method BackupsClient.GetVolumeRestoreStatus.
type BackupsClientGetVolumeRestoreStatusResponse struct {
	RestoreStatus
}

// BackupsClientListResponse contains the response from method BackupsClient.NewListPager.
type BackupsClientListResponse struct {
	BackupsList
}

// BackupsClientRestoreFilesResponse contains the response from method BackupsClient.BeginRestoreFiles.
type BackupsClientRestoreFilesResponse struct {
	// placeholder for future response values
}

// BackupsClientUpdateResponse contains the response from method BackupsClient.BeginUpdate.
type BackupsClientUpdateResponse struct {
	Backup
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// PoolsClientCreateOrUpdateResponse contains the response from method PoolsClient.BeginCreateOrUpdate.
type PoolsClientCreateOrUpdateResponse struct {
	CapacityPool
}

// PoolsClientDeleteResponse contains the response from method PoolsClient.BeginDelete.
type PoolsClientDeleteResponse struct {
	// placeholder for future response values
}

// PoolsClientGetResponse contains the response from method PoolsClient.Get.
type PoolsClientGetResponse struct {
	CapacityPool
}

// PoolsClientListResponse contains the response from method PoolsClient.NewListPager.
type PoolsClientListResponse struct {
	CapacityPoolList
}

// PoolsClientUpdateResponse contains the response from method PoolsClient.BeginUpdate.
type PoolsClientUpdateResponse struct {
	CapacityPool
}

// ResourceClientCheckFilePathAvailabilityResponse contains the response from method ResourceClient.CheckFilePathAvailability.
type ResourceClientCheckFilePathAvailabilityResponse struct {
	CheckAvailabilityResponse
}

// ResourceClientCheckNameAvailabilityResponse contains the response from method ResourceClient.CheckNameAvailability.
type ResourceClientCheckNameAvailabilityResponse struct {
	CheckAvailabilityResponse
}

// ResourceClientCheckQuotaAvailabilityResponse contains the response from method ResourceClient.CheckQuotaAvailability.
type ResourceClientCheckQuotaAvailabilityResponse struct {
	CheckAvailabilityResponse
}

// ResourceClientQueryRegionInfoResponse contains the response from method ResourceClient.QueryRegionInfo.
type ResourceClientQueryRegionInfoResponse struct {
	RegionInfo
}

// ResourceQuotaLimitsClientGetResponse contains the response from method ResourceQuotaLimitsClient.Get.
type ResourceQuotaLimitsClientGetResponse struct {
	SubscriptionQuotaItem
}

// ResourceQuotaLimitsClientListResponse contains the response from method ResourceQuotaLimitsClient.NewListPager.
type ResourceQuotaLimitsClientListResponse struct {
	SubscriptionQuotaItemList
}

// SnapshotPoliciesClientCreateResponse contains the response from method SnapshotPoliciesClient.Create.
type SnapshotPoliciesClientCreateResponse struct {
	SnapshotPolicy
}

// SnapshotPoliciesClientDeleteResponse contains the response from method SnapshotPoliciesClient.BeginDelete.
type SnapshotPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// SnapshotPoliciesClientGetResponse contains the response from method SnapshotPoliciesClient.Get.
type SnapshotPoliciesClientGetResponse struct {
	SnapshotPolicy
}

// SnapshotPoliciesClientListResponse contains the response from method SnapshotPoliciesClient.NewListPager.
type SnapshotPoliciesClientListResponse struct {
	SnapshotPoliciesList
}

// SnapshotPoliciesClientListVolumesResponse contains the response from method SnapshotPoliciesClient.ListVolumes.
type SnapshotPoliciesClientListVolumesResponse struct {
	SnapshotPolicyVolumeList
}

// SnapshotPoliciesClientUpdateResponse contains the response from method SnapshotPoliciesClient.BeginUpdate.
type SnapshotPoliciesClientUpdateResponse struct {
	SnapshotPolicy
}

// SnapshotsClientCreateResponse contains the response from method SnapshotsClient.BeginCreate.
type SnapshotsClientCreateResponse struct {
	Snapshot
}

// SnapshotsClientDeleteResponse contains the response from method SnapshotsClient.BeginDelete.
type SnapshotsClientDeleteResponse struct {
	// placeholder for future response values
}

// SnapshotsClientGetResponse contains the response from method SnapshotsClient.Get.
type SnapshotsClientGetResponse struct {
	Snapshot
}

// SnapshotsClientListResponse contains the response from method SnapshotsClient.NewListPager.
type SnapshotsClientListResponse struct {
	SnapshotsList
}

// SnapshotsClientRestoreFilesResponse contains the response from method SnapshotsClient.BeginRestoreFiles.
type SnapshotsClientRestoreFilesResponse struct {
	// placeholder for future response values
}

// SnapshotsClientUpdateResponse contains the response from method SnapshotsClient.BeginUpdate.
type SnapshotsClientUpdateResponse struct {
	Snapshot
}

// SubvolumesClientCreateResponse contains the response from method SubvolumesClient.BeginCreate.
type SubvolumesClientCreateResponse struct {
	SubvolumeInfo
}

// SubvolumesClientDeleteResponse contains the response from method SubvolumesClient.BeginDelete.
type SubvolumesClientDeleteResponse struct {
	// placeholder for future response values
}

// SubvolumesClientGetMetadataResponse contains the response from method SubvolumesClient.BeginGetMetadata.
type SubvolumesClientGetMetadataResponse struct {
	SubvolumeModel
}

// SubvolumesClientGetResponse contains the response from method SubvolumesClient.Get.
type SubvolumesClientGetResponse struct {
	SubvolumeInfo
}

// SubvolumesClientListByVolumeResponse contains the response from method SubvolumesClient.NewListByVolumePager.
type SubvolumesClientListByVolumeResponse struct {
	SubvolumesList
}

// SubvolumesClientUpdateResponse contains the response from method SubvolumesClient.BeginUpdate.
type SubvolumesClientUpdateResponse struct {
	SubvolumeInfo
}

// VolumeGroupsClientCreateResponse contains the response from method VolumeGroupsClient.BeginCreate.
type VolumeGroupsClientCreateResponse struct {
	VolumeGroupDetails
}

// VolumeGroupsClientDeleteResponse contains the response from method VolumeGroupsClient.BeginDelete.
type VolumeGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// VolumeGroupsClientGetResponse contains the response from method VolumeGroupsClient.Get.
type VolumeGroupsClientGetResponse struct {
	VolumeGroupDetails
}

// VolumeGroupsClientListByNetAppAccountResponse contains the response from method VolumeGroupsClient.NewListByNetAppAccountPager.
type VolumeGroupsClientListByNetAppAccountResponse struct {
	VolumeGroupList
}

// VolumeQuotaRulesClientCreateResponse contains the response from method VolumeQuotaRulesClient.BeginCreate.
type VolumeQuotaRulesClientCreateResponse struct {
	VolumeQuotaRule
}

// VolumeQuotaRulesClientDeleteResponse contains the response from method VolumeQuotaRulesClient.BeginDelete.
type VolumeQuotaRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// VolumeQuotaRulesClientGetResponse contains the response from method VolumeQuotaRulesClient.Get.
type VolumeQuotaRulesClientGetResponse struct {
	VolumeQuotaRule
}

// VolumeQuotaRulesClientListByVolumeResponse contains the response from method VolumeQuotaRulesClient.NewListByVolumePager.
type VolumeQuotaRulesClientListByVolumeResponse struct {
	VolumeQuotaRulesList
}

// VolumeQuotaRulesClientUpdateResponse contains the response from method VolumeQuotaRulesClient.BeginUpdate.
type VolumeQuotaRulesClientUpdateResponse struct {
	VolumeQuotaRule
}

// VolumesClientAuthorizeReplicationResponse contains the response from method VolumesClient.BeginAuthorizeReplication.
type VolumesClientAuthorizeReplicationResponse struct {
	// placeholder for future response values
}

// VolumesClientBreakFileLocksResponse contains the response from method VolumesClient.BeginBreakFileLocks.
type VolumesClientBreakFileLocksResponse struct {
	// placeholder for future response values
}

// VolumesClientBreakReplicationResponse contains the response from method VolumesClient.BeginBreakReplication.
type VolumesClientBreakReplicationResponse struct {
	// placeholder for future response values
}

// VolumesClientCreateOrUpdateResponse contains the response from method VolumesClient.BeginCreateOrUpdate.
type VolumesClientCreateOrUpdateResponse struct {
	Volume
}

// VolumesClientDeleteReplicationResponse contains the response from method VolumesClient.BeginDeleteReplication.
type VolumesClientDeleteReplicationResponse struct {
	// placeholder for future response values
}

// VolumesClientDeleteResponse contains the response from method VolumesClient.BeginDelete.
type VolumesClientDeleteResponse struct {
	// placeholder for future response values
}

// VolumesClientFinalizeRelocationResponse contains the response from method VolumesClient.BeginFinalizeRelocation.
type VolumesClientFinalizeRelocationResponse struct {
	// placeholder for future response values
}

// VolumesClientGetResponse contains the response from method VolumesClient.Get.
type VolumesClientGetResponse struct {
	Volume
}

// VolumesClientListReplicationsResponse contains the response from method VolumesClient.NewListReplicationsPager.
type VolumesClientListReplicationsResponse struct {
	ListReplications
}

// VolumesClientListResponse contains the response from method VolumesClient.NewListPager.
type VolumesClientListResponse struct {
	VolumeList
}

// VolumesClientPoolChangeResponse contains the response from method VolumesClient.BeginPoolChange.
type VolumesClientPoolChangeResponse struct {
	// placeholder for future response values
}

// VolumesClientReInitializeReplicationResponse contains the response from method VolumesClient.BeginReInitializeReplication.
type VolumesClientReInitializeReplicationResponse struct {
	// placeholder for future response values
}

// VolumesClientReestablishReplicationResponse contains the response from method VolumesClient.BeginReestablishReplication.
type VolumesClientReestablishReplicationResponse struct {
	// placeholder for future response values
}

// VolumesClientRelocateResponse contains the response from method VolumesClient.BeginRelocate.
type VolumesClientRelocateResponse struct {
	// placeholder for future response values
}

// VolumesClientReplicationStatusResponse contains the response from method VolumesClient.ReplicationStatus.
type VolumesClientReplicationStatusResponse struct {
	ReplicationStatus
}

// VolumesClientResetCifsPasswordResponse contains the response from method VolumesClient.BeginResetCifsPassword.
type VolumesClientResetCifsPasswordResponse struct {
	// placeholder for future response values
}

// VolumesClientResyncReplicationResponse contains the response from method VolumesClient.BeginResyncReplication.
type VolumesClientResyncReplicationResponse struct {
	// placeholder for future response values
}

// VolumesClientRevertRelocationResponse contains the response from method VolumesClient.BeginRevertRelocation.
type VolumesClientRevertRelocationResponse struct {
	// placeholder for future response values
}

// VolumesClientRevertResponse contains the response from method VolumesClient.BeginRevert.
type VolumesClientRevertResponse struct {
	// placeholder for future response values
}

// VolumesClientUpdateResponse contains the response from method VolumesClient.BeginUpdate.
type VolumesClientUpdateResponse struct {
	Volume
}
