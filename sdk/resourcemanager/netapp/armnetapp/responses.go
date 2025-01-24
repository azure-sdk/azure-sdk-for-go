//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

// AccountsClientCreateOrUpdateResponse contains the response from method AccountsClient.BeginCreateOrUpdate.
type AccountsClientCreateOrUpdateResponse struct {
	// NetApp account resource
	Account
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.BeginDelete.
type AccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	// NetApp account resource
	Account
}

// AccountsClientListBySubscriptionResponse contains the response from method AccountsClient.NewListBySubscriptionPager.
type AccountsClientListBySubscriptionResponse struct {
	// List of NetApp account resources
	AccountList
}

// AccountsClientListResponse contains the response from method AccountsClient.NewListPager.
type AccountsClientListResponse struct {
	// List of NetApp account resources
	AccountList
}

// AccountsClientRenewCredentialsResponse contains the response from method AccountsClient.BeginRenewCredentials.
type AccountsClientRenewCredentialsResponse struct {
	// placeholder for future response values
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.BeginUpdate.
type AccountsClientUpdateResponse struct {
	// NetApp account resource
	Account
}

// BackupPoliciesClientCreateResponse contains the response from method BackupPoliciesClient.BeginCreate.
type BackupPoliciesClientCreateResponse struct {
	// Backup policy information
	BackupPolicy
}

// BackupPoliciesClientDeleteResponse contains the response from method BackupPoliciesClient.BeginDelete.
type BackupPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// BackupPoliciesClientGetResponse contains the response from method BackupPoliciesClient.Get.
type BackupPoliciesClientGetResponse struct {
	// Backup policy information
	BackupPolicy
}

// BackupPoliciesClientListResponse contains the response from method BackupPoliciesClient.NewListPager.
type BackupPoliciesClientListResponse struct {
	// List of Backup Policies
	BackupPoliciesList
}

// BackupPoliciesClientUpdateResponse contains the response from method BackupPoliciesClient.BeginUpdate.
type BackupPoliciesClientUpdateResponse struct {
	// Backup policy information
	BackupPolicy
}

// BackupVaultsClientCreateOrUpdateResponse contains the response from method BackupVaultsClient.BeginCreateOrUpdate.
type BackupVaultsClientCreateOrUpdateResponse struct {
	// Backup Vault information
	BackupVault
}

// BackupVaultsClientDeleteResponse contains the response from method BackupVaultsClient.BeginDelete.
type BackupVaultsClientDeleteResponse struct {
	// placeholder for future response values
}

// BackupVaultsClientGetResponse contains the response from method BackupVaultsClient.Get.
type BackupVaultsClientGetResponse struct {
	// Backup Vault information
	BackupVault
}

// BackupVaultsClientListByNetAppAccountResponse contains the response from method BackupVaultsClient.NewListByNetAppAccountPager.
type BackupVaultsClientListByNetAppAccountResponse struct {
	// List of Backup Vaults
	BackupVaultsList
}

// BackupVaultsClientUpdateResponse contains the response from method BackupVaultsClient.BeginUpdate.
type BackupVaultsClientUpdateResponse struct {
	// Backup Vault information
	BackupVault
}

// BackupsClientCreateResponse contains the response from method BackupsClient.BeginCreate.
type BackupsClientCreateResponse struct {
	// Backup under a Backup Vault
	Backup
}

// BackupsClientDeleteResponse contains the response from method BackupsClient.BeginDelete.
type BackupsClientDeleteResponse struct {
	// placeholder for future response values
}

// BackupsClientGetLatestStatusResponse contains the response from method BackupsClient.GetLatestStatus.
type BackupsClientGetLatestStatusResponse struct {
	// Backup status
	BackupStatus
}

// BackupsClientGetResponse contains the response from method BackupsClient.Get.
type BackupsClientGetResponse struct {
	// Backup under a Backup Vault
	Backup
}

// BackupsClientGetVolumeLatestRestoreStatusResponse contains the response from method BackupsClient.GetVolumeLatestRestoreStatus.
type BackupsClientGetVolumeLatestRestoreStatusResponse struct {
	// Restore status
	RestoreStatus
}

// BackupsClientListByVaultResponse contains the response from method BackupsClient.NewListByVaultPager.
type BackupsClientListByVaultResponse struct {
	// List of Backups
	BackupsList
}

// BackupsClientUpdateResponse contains the response from method BackupsClient.BeginUpdate.
type BackupsClientUpdateResponse struct {
	// Backup under a Backup Vault
	Backup
}

// BackupsUnderAccountClientMigrateBackupsResponse contains the response from method BackupsUnderAccountClient.BeginMigrateBackups.
type BackupsUnderAccountClientMigrateBackupsResponse struct {
	// placeholder for future response values
}

// BackupsUnderBackupVaultClientRestoreFilesResponse contains the response from method BackupsUnderBackupVaultClient.BeginRestoreFiles.
type BackupsUnderBackupVaultClientRestoreFilesResponse struct {
	// placeholder for future response values
}

// BackupsUnderVolumeClientMigrateBackupsResponse contains the response from method BackupsUnderVolumeClient.BeginMigrateBackups.
type BackupsUnderVolumeClientMigrateBackupsResponse struct {
	// placeholder for future response values
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of the request to list Cloud Volume operations. It contains a list of operations and a URL link to get the next
	// set of results.
	OperationListResult
}

// PoolsClientCreateOrUpdateResponse contains the response from method PoolsClient.BeginCreateOrUpdate.
type PoolsClientCreateOrUpdateResponse struct {
	// Capacity pool resource
	CapacityPool
}

// PoolsClientDeleteResponse contains the response from method PoolsClient.BeginDelete.
type PoolsClientDeleteResponse struct {
	// placeholder for future response values
}

// PoolsClientGetResponse contains the response from method PoolsClient.Get.
type PoolsClientGetResponse struct {
	// Capacity pool resource
	CapacityPool
}

// PoolsClientListResponse contains the response from method PoolsClient.NewListPager.
type PoolsClientListResponse struct {
	// List of capacity pool resources
	CapacityPoolList
}

// PoolsClientUpdateResponse contains the response from method PoolsClient.BeginUpdate.
type PoolsClientUpdateResponse struct {
	// Capacity pool resource
	CapacityPool
}

// ResourceClientCheckFilePathAvailabilityResponse contains the response from method ResourceClient.CheckFilePathAvailability.
type ResourceClientCheckFilePathAvailabilityResponse struct {
	// Information regarding availability of a resource.
	CheckAvailabilityResponse
}

// ResourceClientCheckNameAvailabilityResponse contains the response from method ResourceClient.CheckNameAvailability.
type ResourceClientCheckNameAvailabilityResponse struct {
	// Information regarding availability of a resource.
	CheckAvailabilityResponse
}

// ResourceClientCheckQuotaAvailabilityResponse contains the response from method ResourceClient.CheckQuotaAvailability.
type ResourceClientCheckQuotaAvailabilityResponse struct {
	// Information regarding availability of a resource.
	CheckAvailabilityResponse
}

// ResourceClientQueryNetworkSiblingSetResponse contains the response from method ResourceClient.QueryNetworkSiblingSet.
type ResourceClientQueryNetworkSiblingSetResponse struct {
	// Describes the contents of a network sibling set.
	NetworkSiblingSet
}

// ResourceClientQueryRegionInfoResponse contains the response from method ResourceClient.QueryRegionInfo.
type ResourceClientQueryRegionInfoResponse struct {
	// Provides region specific information.
	RegionInfo
}

// ResourceClientUpdateNetworkSiblingSetResponse contains the response from method ResourceClient.BeginUpdateNetworkSiblingSet.
type ResourceClientUpdateNetworkSiblingSetResponse struct {
	// Describes the contents of a network sibling set.
	NetworkSiblingSet
}

// ResourceQuotaLimitsClientGetResponse contains the response from method ResourceQuotaLimitsClient.Get.
type ResourceQuotaLimitsClientGetResponse struct {
	// Information regarding Subscription Quota Item.
	SubscriptionQuotaItem
}

// ResourceQuotaLimitsClientListResponse contains the response from method ResourceQuotaLimitsClient.NewListPager.
type ResourceQuotaLimitsClientListResponse struct {
	// List of Subscription Quota Items
	SubscriptionQuotaItemList
}

// ResourceRegionInfosClientGetResponse contains the response from method ResourceRegionInfosClient.Get.
type ResourceRegionInfosClientGetResponse struct {
	// Information regarding regionInfo Item.
	RegionInfoResource
}

// ResourceRegionInfosClientListResponse contains the response from method ResourceRegionInfosClient.NewListPager.
type ResourceRegionInfosClientListResponse struct {
	// List of regionInfo resources
	RegionInfosList
}

// SnapshotPoliciesClientCreateResponse contains the response from method SnapshotPoliciesClient.Create.
type SnapshotPoliciesClientCreateResponse struct {
	// Snapshot policy information
	SnapshotPolicy
}

// SnapshotPoliciesClientDeleteResponse contains the response from method SnapshotPoliciesClient.BeginDelete.
type SnapshotPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// SnapshotPoliciesClientGetResponse contains the response from method SnapshotPoliciesClient.Get.
type SnapshotPoliciesClientGetResponse struct {
	// Snapshot policy information
	SnapshotPolicy
}

// SnapshotPoliciesClientListResponse contains the response from method SnapshotPoliciesClient.NewListPager.
type SnapshotPoliciesClientListResponse struct {
	// List of Snapshot Policies
	SnapshotPoliciesList
}

// SnapshotPoliciesClientListVolumesResponse contains the response from method SnapshotPoliciesClient.ListVolumes.
type SnapshotPoliciesClientListVolumesResponse struct {
	// Volumes associated with snapshot policy
	SnapshotPolicyVolumeList
}

// SnapshotPoliciesClientUpdateResponse contains the response from method SnapshotPoliciesClient.BeginUpdate.
type SnapshotPoliciesClientUpdateResponse struct {
	// Snapshot policy information
	SnapshotPolicy
}

// SnapshotsClientCreateResponse contains the response from method SnapshotsClient.BeginCreate.
type SnapshotsClientCreateResponse struct {
	// Snapshot of a Volume
	Snapshot
}

// SnapshotsClientDeleteResponse contains the response from method SnapshotsClient.BeginDelete.
type SnapshotsClientDeleteResponse struct {
	// placeholder for future response values
}

// SnapshotsClientGetResponse contains the response from method SnapshotsClient.Get.
type SnapshotsClientGetResponse struct {
	// Snapshot of a Volume
	Snapshot
}

// SnapshotsClientListResponse contains the response from method SnapshotsClient.NewListPager.
type SnapshotsClientListResponse struct {
	// List of Snapshots
	SnapshotsList
}

// SnapshotsClientRestoreFilesResponse contains the response from method SnapshotsClient.BeginRestoreFiles.
type SnapshotsClientRestoreFilesResponse struct {
	// placeholder for future response values
}

// SnapshotsClientUpdateResponse contains the response from method SnapshotsClient.BeginUpdate.
type SnapshotsClientUpdateResponse struct {
	// Snapshot of a Volume
	Snapshot
}

// SubvolumesClientCreateResponse contains the response from method SubvolumesClient.BeginCreate.
type SubvolumesClientCreateResponse struct {
	// Subvolume Information properties
	SubvolumeInfo
}

// SubvolumesClientDeleteResponse contains the response from method SubvolumesClient.BeginDelete.
type SubvolumesClientDeleteResponse struct {
	// placeholder for future response values
}

// SubvolumesClientGetMetadataResponse contains the response from method SubvolumesClient.BeginGetMetadata.
type SubvolumesClientGetMetadataResponse struct {
	// Result of the post subvolume and action is to get metadata of the subvolume.
	SubvolumeModel
}

// SubvolumesClientGetResponse contains the response from method SubvolumesClient.Get.
type SubvolumesClientGetResponse struct {
	// Subvolume Information properties
	SubvolumeInfo
}

// SubvolumesClientListByVolumeResponse contains the response from method SubvolumesClient.NewListByVolumePager.
type SubvolumesClientListByVolumeResponse struct {
	// List of Subvolumes
	SubvolumesList
}

// SubvolumesClientUpdateResponse contains the response from method SubvolumesClient.BeginUpdate.
type SubvolumesClientUpdateResponse struct {
	// Subvolume Information properties
	SubvolumeInfo
}

// VolumeGroupsClientCreateResponse contains the response from method VolumeGroupsClient.BeginCreate.
type VolumeGroupsClientCreateResponse struct {
	// Volume group resource for create
	VolumeGroupDetails
}

// VolumeGroupsClientDeleteResponse contains the response from method VolumeGroupsClient.BeginDelete.
type VolumeGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// VolumeGroupsClientGetResponse contains the response from method VolumeGroupsClient.Get.
type VolumeGroupsClientGetResponse struct {
	// Volume group resource for create
	VolumeGroupDetails
}

// VolumeGroupsClientListByNetAppAccountResponse contains the response from method VolumeGroupsClient.NewListByNetAppAccountPager.
type VolumeGroupsClientListByNetAppAccountResponse struct {
	// List of volume group resources
	VolumeGroupList
}

// VolumeQuotaRulesClientCreateResponse contains the response from method VolumeQuotaRulesClient.BeginCreate.
type VolumeQuotaRulesClientCreateResponse struct {
	// Quota Rule of a Volume
	VolumeQuotaRule
}

// VolumeQuotaRulesClientDeleteResponse contains the response from method VolumeQuotaRulesClient.BeginDelete.
type VolumeQuotaRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// VolumeQuotaRulesClientGetResponse contains the response from method VolumeQuotaRulesClient.Get.
type VolumeQuotaRulesClientGetResponse struct {
	// Quota Rule of a Volume
	VolumeQuotaRule
}

// VolumeQuotaRulesClientListByVolumeResponse contains the response from method VolumeQuotaRulesClient.NewListByVolumePager.
type VolumeQuotaRulesClientListByVolumeResponse struct {
	// List of Volume Quota Rules
	VolumeQuotaRulesList
}

// VolumeQuotaRulesClientUpdateResponse contains the response from method VolumeQuotaRulesClient.BeginUpdate.
type VolumeQuotaRulesClientUpdateResponse struct {
	// Quota Rule of a Volume
	VolumeQuotaRule
}

// VolumesClientAuthorizeExternalReplicationResponse contains the response from method VolumesClient.BeginAuthorizeExternalReplication.
type VolumesClientAuthorizeExternalReplicationResponse struct {
	// Information about svm peering process
	SvmPeerCommandResponse
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
	// Volume resource
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

// VolumesClientFinalizeExternalReplicationResponse contains the response from method VolumesClient.BeginFinalizeExternalReplication.
type VolumesClientFinalizeExternalReplicationResponse struct {
	// placeholder for future response values
}

// VolumesClientFinalizeRelocationResponse contains the response from method VolumesClient.BeginFinalizeRelocation.
type VolumesClientFinalizeRelocationResponse struct {
	// placeholder for future response values
}

// VolumesClientGetResponse contains the response from method VolumesClient.Get.
type VolumesClientGetResponse struct {
	// Volume resource
	Volume
}

// VolumesClientListGetGroupIDListForLdapUserResponse contains the response from method VolumesClient.BeginListGetGroupIDListForLdapUser.
type VolumesClientListGetGroupIDListForLdapUserResponse struct {
	// Group Id list for Ldap user
	GetGroupIDListForLDAPUserResponse
}

// VolumesClientListReplicationsResponse contains the response from method VolumesClient.NewListReplicationsPager.
type VolumesClientListReplicationsResponse struct {
	// List Replications
	ListReplications
}

// VolumesClientListResponse contains the response from method VolumesClient.NewListPager.
type VolumesClientListResponse struct {
	// List of volume resources
	VolumeList
}

// VolumesClientPeerExternalClusterResponse contains the response from method VolumesClient.BeginPeerExternalCluster.
type VolumesClientPeerExternalClusterResponse struct {
	// Information about cluster peering process
	ClusterPeerCommandResponse
}

// VolumesClientPerformReplicationTransferResponse contains the response from method VolumesClient.BeginPerformReplicationTransfer.
type VolumesClientPerformReplicationTransferResponse struct {
	// placeholder for future response values
}

// VolumesClientPoolChangeResponse contains the response from method VolumesClient.BeginPoolChange.
type VolumesClientPoolChangeResponse struct {
	// placeholder for future response values
}

// VolumesClientPopulateAvailabilityZoneResponse contains the response from method VolumesClient.BeginPopulateAvailabilityZone.
type VolumesClientPopulateAvailabilityZoneResponse struct {
	// Volume resource
	Volume
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
	// Replication status
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
	// Volume resource
	Volume
}
