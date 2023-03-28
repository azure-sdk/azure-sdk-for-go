# Release History

## 2.1.0 (2023-03-28)
### Features Added

- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewBackupInstancesClient() *BackupInstancesClient`
- New function `*ClientFactory.NewBackupPoliciesClient() *BackupPoliciesClient`
- New function `*ClientFactory.NewBackupVaultOperationResultsClient() *BackupVaultOperationResultsClient`
- New function `*ClientFactory.NewBackupVaultsClient() *BackupVaultsClient`
- New function `*ClientFactory.NewClient() *Client`
- New function `*ClientFactory.NewDeletedBackupInstancesClient() *DeletedBackupInstancesClient`
- New function `*ClientFactory.NewDppResourceGuardProxyClient() *DppResourceGuardProxyClient`
- New function `*ClientFactory.NewExportJobsClient() *ExportJobsClient`
- New function `*ClientFactory.NewExportJobsOperationResultClient() *ExportJobsOperationResultClient`
- New function `*ClientFactory.NewJobsClient() *JobsClient`
- New function `*ClientFactory.NewOperationResultClient() *OperationResultClient`
- New function `*ClientFactory.NewOperationStatusBackupVaultContextClient() *OperationStatusBackupVaultContextClient`
- New function `*ClientFactory.NewOperationStatusClient() *OperationStatusClient`
- New function `*ClientFactory.NewOperationStatusResourceGroupContextClient() *OperationStatusResourceGroupContextClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewRecoveryPointsClient() *RecoveryPointsClient`
- New function `*ClientFactory.NewResourceGuardsClient() *ResourceGuardsClient`
- New function `*ClientFactory.NewRestorableTimeRangesClient() *RestorableTimeRangesClient`
- New function `NewDppResourceGuardProxyClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DppResourceGuardProxyClient, error)`
- New function `*DppResourceGuardProxyClient.CreateOrUpdate(context.Context, string, string, string, ResourceGuardProxyBaseResource, *DppResourceGuardProxyClientCreateOrUpdateOptions) (DppResourceGuardProxyClientCreateOrUpdateResponse, error)`
- New function `*DppResourceGuardProxyClient.Delete(context.Context, string, string, string, *DppResourceGuardProxyClientDeleteOptions) (DppResourceGuardProxyClientDeleteResponse, error)`
- New function `*DppResourceGuardProxyClient.Get(context.Context, string, string, string, *DppResourceGuardProxyClientGetOptions) (DppResourceGuardProxyClientGetResponse, error)`
- New function `*DppResourceGuardProxyClient.NewListPager(string, string, *DppResourceGuardProxyClientListOptions) *runtime.Pager[DppResourceGuardProxyClientListResponse]`
- New function `*DppResourceGuardProxyClient.UnlockDelete(context.Context, string, string, string, UnlockDeleteRequest, *DppResourceGuardProxyClientUnlockDeleteOptions) (DppResourceGuardProxyClientUnlockDeleteResponse, error)`
- New struct `ClientFactory`
- New struct `ResourceGuardOperationDetail`
- New struct `ResourceGuardProxyBase`
- New struct `ResourceGuardProxyBaseResource`
- New struct `ResourceGuardProxyBaseResourceList`
- New struct `UnlockDeleteRequest`
- New struct `UnlockDeleteResponse`


## 2.0.0 (2023-02-24)
### Breaking Changes

- Function `*ResourceGuardsClient.Patch` parameter(s) have been changed from `(context.Context, string, string, PatchResourceRequestInput, *ResourceGuardsClientPatchOptions)` to `(context.Context, string, string, PatchResourceGuardInput, *ResourceGuardsClientPatchOptions)`
- Const `StorageSettingStoreTypesSnapshotStore` from type alias `StorageSettingStoreTypes` has been removed
- Operation `*BackupVaultsClient.Delete` has been changed to LRO, use `*BackupVaultsClient.BeginDelete` instead.
- Field `Identity` of struct `ResourceGuardResource` has been removed

### Features Added

- New value `SourceDataStoreTypeOperationalStore` added to type alias `SourceDataStoreType`
- New value `StorageSettingStoreTypesOperationalStore` added to type alias `StorageSettingStoreTypes`
- New value `StorageSettingTypesZoneRedundant` added to type alias `StorageSettingTypes`
- New type alias `CrossSubscriptionRestoreState` with values `CrossSubscriptionRestoreStateDisabled`, `CrossSubscriptionRestoreStateEnabled`, `CrossSubscriptionRestoreStatePermanentlyDisabled`
- New type alias `ExistingResourcePolicy` with values `ExistingResourcePolicyPatch`, `ExistingResourcePolicySkip`
- New type alias `ImmutabilityState` with values `ImmutabilityStateDisabled`, `ImmutabilityStateLocked`, `ImmutabilityStateUnlocked`
- New type alias `PersistentVolumeRestoreMode` with values `PersistentVolumeRestoreModeRestoreWithVolumeData`, `PersistentVolumeRestoreModeRestoreWithoutVolumeData`
- New type alias `SoftDeleteState` with values `SoftDeleteStateAlwaysOn`, `SoftDeleteStateOff`, `SoftDeleteStateOn`
- New function `*BackupDatasourceParameters.GetBackupDatasourceParameters() *BackupDatasourceParameters`
- New function `*BlobBackupDatasourceParameters.GetBackupDatasourceParameters() *BackupDatasourceParameters`
- New function `NewDeletedBackupInstancesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DeletedBackupInstancesClient, error)`
- New function `*DeletedBackupInstancesClient.Get(context.Context, string, string, string, *DeletedBackupInstancesClientGetOptions) (DeletedBackupInstancesClientGetResponse, error)`
- New function `*DeletedBackupInstancesClient.NewListPager(string, string, *DeletedBackupInstancesClientListOptions) *runtime.Pager[DeletedBackupInstancesClientListResponse]`
- New function `*DeletedBackupInstancesClient.BeginUndelete(context.Context, string, string, string, *DeletedBackupInstancesClientBeginUndeleteOptions) (*runtime.Poller[DeletedBackupInstancesClientUndeleteResponse], error)`
- New function `*ItemPathBasedRestoreCriteria.GetItemLevelRestoreCriteria() *ItemLevelRestoreCriteria`
- New function `*KubernetesClusterBackupDatasourceParameters.GetBackupDatasourceParameters() *BackupDatasourceParameters`
- New function `*KubernetesClusterRestoreCriteria.GetItemLevelRestoreCriteria() *ItemLevelRestoreCriteria`
- New struct `BlobBackupDatasourceParameters`
- New struct `CrossSubscriptionRestoreSettings`
- New struct `DeletedBackupInstance`
- New struct `DeletedBackupInstanceResource`
- New struct `DeletedBackupInstanceResourceList`
- New struct `DeletedBackupInstancesClient`
- New struct `DeletedBackupInstancesClientListResponse`
- New struct `DeletedBackupInstancesClientUndeleteResponse`
- New struct `DeletionInfo`
- New struct `DppBaseTrackedResource`
- New struct `DppProxyResource`
- New struct `FeatureSettings`
- New struct `ImmutabilitySettings`
- New struct `ItemPathBasedRestoreCriteria`
- New struct `KubernetesClusterBackupDatasourceParameters`
- New struct `KubernetesClusterRestoreCriteria`
- New struct `PatchResourceGuardInput`
- New struct `SecuritySettings`
- New struct `SoftDeleteSettings`
- New field `ExpiryTime` in struct `AzureBackupDiscreteRecoveryPoint`
- New field `Tags` in struct `BackupInstanceResource`
- New field `FeatureSettings` in struct `BackupVault`
- New field `IsVaultProtectedByResourceGuard` in struct `BackupVault`
- New field `SecuritySettings` in struct `BackupVault`
- New field `FeatureSettings` in struct `PatchBackupVaultInput`
- New field `SecuritySettings` in struct `PatchBackupVaultInput`
- New field `BackupDatasourceParametersList` in struct `PolicyParameters`
- New field `TargetResourceArmID` in struct `TargetDetails`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).