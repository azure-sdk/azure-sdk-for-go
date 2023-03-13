# Release History

## 3.0.0-beta.1 (2023-03-13)
### Breaking Changes

- Function `*ResourceGuardsClient.Patch` parameter(s) have been changed from `(context.Context, string, string, PatchResourceGuardInput, *ResourceGuardsClientPatchOptions)` to `(context.Context, string, string, PatchResourceRequestInput, *ResourceGuardsClientPatchOptions)`
- Type of `ResourceGuard.ProvisioningState` has been changed from `*ProvisioningState` to `*ResourceGuardProvisioningState`
- Const `SourceDataStoreTypeOperationalStore` from type alias `SourceDataStoreType` has been removed
- Const `StorageSettingStoreTypesOperationalStore` from type alias `StorageSettingStoreTypes` has been removed
- Type alias `CrossSubscriptionRestoreState` has been removed
- Function `*BlobBackupDatasourceParameters.GetBackupDatasourceParameters` has been removed
- Function `*ItemPathBasedRestoreCriteria.GetItemLevelRestoreCriteria` has been removed
- Operation `*BackupVaultsClient.BeginDelete` has been changed to non-LRO, use `*BackupVaultsClient.Delete` instead.
- Struct `BlobBackupDatasourceParameters` has been removed
- Struct `CrossSubscriptionRestoreSettings` has been removed
- Struct `DppBaseTrackedResource` has been removed
- Struct `FeatureSettings` has been removed
- Struct `ItemPathBasedRestoreCriteria` has been removed
- Struct `PatchResourceGuardInput` has been removed
- Field `ExpiryTime` of struct `AzureBackupDiscreteRecoveryPoint` has been removed
- Field `FeatureSettings` of struct `BackupVault` has been removed
- Field `FeatureSettings` of struct `PatchBackupVaultInput` has been removed
- Field `TargetResourceArmID` of struct `TargetDetails` has been removed

### Features Added

- New value `StorageSettingStoreTypesSnapshotStore` added to type alias `StorageSettingStoreTypes`
- New type alias `ResourceGuardProvisioningState` with values `ResourceGuardProvisioningStateFailed`, `ResourceGuardProvisioningStateProvisioning`, `ResourceGuardProvisioningStateSucceeded`, `ResourceGuardProvisioningStateUnknown`, `ResourceGuardProvisioningStateUpdating`
- New function `NewBackupInstancesExtensionRoutingClient(azcore.TokenCredential, *arm.ClientOptions) (*BackupInstancesExtensionRoutingClient, error)`
- New function `*BackupInstancesExtensionRoutingClient.NewListPager(string, *BackupInstancesExtensionRoutingClientListOptions) *runtime.Pager[BackupInstancesExtensionRoutingClientListResponse]`
- New function `NewDppResourceGuardProxyClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DppResourceGuardProxyClient, error)`
- New function `*DppResourceGuardProxyClient.Delete(context.Context, string, string, string, *DppResourceGuardProxyClientDeleteOptions) (DppResourceGuardProxyClientDeleteResponse, error)`
- New function `*DppResourceGuardProxyClient.Get(context.Context, string, string, string, *DppResourceGuardProxyClientGetOptions) (DppResourceGuardProxyClientGetResponse, error)`
- New function `*DppResourceGuardProxyClient.NewListPager(string, string, *DppResourceGuardProxyClientListOptions) *runtime.Pager[DppResourceGuardProxyClientListResponse]`
- New function `*DppResourceGuardProxyClient.Put(context.Context, string, string, string, ResourceGuardProxyBaseResource, *DppResourceGuardProxyClientPutOptions) (DppResourceGuardProxyClientPutResponse, error)`
- New function `*DppResourceGuardProxyClient.UnlockDelete(context.Context, string, string, string, UnlockDeleteRequest, *DppResourceGuardProxyClientUnlockDeleteOptions) (DppResourceGuardProxyClientUnlockDeleteResponse, error)`
- New struct `BackupInstancesExtensionRoutingClient`
- New struct `DppResourceGuardProxyClient`
- New struct `ResourceGuardOperationDetail`
- New struct `ResourceGuardProxyBase`
- New struct `ResourceGuardProxyBaseResource`
- New struct `ResourceGuardProxyBaseResourceList`
- New struct `UnlockDeleteRequest`
- New struct `UnlockDeleteResponse`
- New field `Identity` in struct `ResourceGuardResource`


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