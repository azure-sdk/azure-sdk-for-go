# Release History

## 2.0.0-beta.1 (2022-10-27)
### Breaking Changes

- Type of `ResourceGuard.ProvisioningState` has been changed from `*ProvisioningState` to `*ResourceGuardProvisioningState`
- Struct `CloudError` has been removed

### Features Added

- New const `ResourceGuardProvisioningStateUnknown`
- New const `StorageSettingTypesZoneRedundant`
- New const `ExistingResourcePolicyPatch`
- New const `SoftDeleteStateOn`
- New const `ImmutabilityStateUnlocked`
- New const `ImmutabilityStateDisabled`
- New const `PersistentVolumeRestoreModeRestoreWithoutVolumeData`
- New const `PersistentVolumeRestoreModeRestoreWithVolumeData`
- New const `SoftDeleteStateAlwaysOn`
- New const `ResourceGuardProvisioningStateUpdating`
- New const `ResourceGuardProvisioningStateFailed`
- New const `SoftDeleteStateOff`
- New const `ImmutabilityStateLocked`
- New const `ResourceGuardProvisioningStateSucceeded`
- New const `ExistingResourcePolicySkip`
- New const `ResourceGuardProvisioningStateProvisioning`
- New type alias `ExistingResourcePolicy`
- New type alias `ResourceGuardProvisioningState`
- New type alias `SoftDeleteState`
- New type alias `ImmutabilityState`
- New type alias `PersistentVolumeRestoreMode`
- New function `PossiblePersistentVolumeRestoreModeValues() []PersistentVolumeRestoreMode`
- New function `*DppResourceGuardProxyClient.Put(context.Context, string, string, string, ResourceGuardProxyBaseResource, *DppResourceGuardProxyClientPutOptions) (DppResourceGuardProxyClientPutResponse, error)`
- New function `PossibleImmutabilityStateValues() []ImmutabilityState`
- New function `NewDppResourceGuardProxyClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DppResourceGuardProxyClient, error)`
- New function `*DeletedBackupInstancesClient.NewListPager(string, string, *DeletedBackupInstancesClientListOptions) *runtime.Pager[DeletedBackupInstancesClientListResponse]`
- New function `*DppResourceGuardProxyClient.Get(context.Context, string, string, string, *DppResourceGuardProxyClientGetOptions) (DppResourceGuardProxyClientGetResponse, error)`
- New function `NewBackupInstancesExtensionRoutingClient(azcore.TokenCredential, *arm.ClientOptions) (*BackupInstancesExtensionRoutingClient, error)`
- New function `NewDeletedBackupInstancesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DeletedBackupInstancesClient, error)`
- New function `*KubernetesClusterBackupDatasourceParameters.GetBackupDatasourceParameters() *BackupDatasourceParameters`
- New function `PossibleResourceGuardProvisioningStateValues() []ResourceGuardProvisioningState`
- New function `*KubernetesClusterRestoreCriteria.GetItemLevelRestoreCriteria() *ItemLevelRestoreCriteria`
- New function `*BackupDatasourceParameters.GetBackupDatasourceParameters() *BackupDatasourceParameters`
- New function `PossibleExistingResourcePolicyValues() []ExistingResourcePolicy`
- New function `*DppResourceGuardProxyClient.UnlockDelete(context.Context, string, string, string, UnlockDeleteRequest, *DppResourceGuardProxyClientUnlockDeleteOptions) (DppResourceGuardProxyClientUnlockDeleteResponse, error)`
- New function `*DeletedBackupInstancesClient.BeginUndelete(context.Context, string, string, string, *DeletedBackupInstancesClientBeginUndeleteOptions) (*runtime.Poller[DeletedBackupInstancesClientUndeleteResponse], error)`
- New function `*DppResourceGuardProxyClient.NewListPager(string, string, *DppResourceGuardProxyClientListOptions) *runtime.Pager[DppResourceGuardProxyClientListResponse]`
- New function `*DppResourceGuardProxyClient.Delete(context.Context, string, string, string, *DppResourceGuardProxyClientDeleteOptions) (DppResourceGuardProxyClientDeleteResponse, error)`
- New function `PossibleSoftDeleteStateValues() []SoftDeleteState`
- New function `*BackupInstancesExtensionRoutingClient.NewListPager(string, *BackupInstancesExtensionRoutingClientListOptions) *runtime.Pager[BackupInstancesExtensionRoutingClientListResponse]`
- New function `*DeletedBackupInstancesClient.Get(context.Context, string, string, string, *DeletedBackupInstancesClientGetOptions) (DeletedBackupInstancesClientGetResponse, error)`
- New struct `BackupDatasourceParameters`
- New struct `BackupInstancesExtensionRoutingClient`
- New struct `BackupInstancesExtensionRoutingClientListOptions`
- New struct `BackupInstancesExtensionRoutingClientListResponse`
- New struct `DeletedBackupInstance`
- New struct `DeletedBackupInstanceResource`
- New struct `DeletedBackupInstanceResourceList`
- New struct `DeletedBackupInstancesClient`
- New struct `DeletedBackupInstancesClientBeginUndeleteOptions`
- New struct `DeletedBackupInstancesClientGetOptions`
- New struct `DeletedBackupInstancesClientGetResponse`
- New struct `DeletedBackupInstancesClientListOptions`
- New struct `DeletedBackupInstancesClientListResponse`
- New struct `DeletedBackupInstancesClientUndeleteResponse`
- New struct `DeletionInfo`
- New struct `DppProxyResource`
- New struct `DppResourceGuardProxyClient`
- New struct `DppResourceGuardProxyClientDeleteOptions`
- New struct `DppResourceGuardProxyClientDeleteResponse`
- New struct `DppResourceGuardProxyClientGetOptions`
- New struct `DppResourceGuardProxyClientGetResponse`
- New struct `DppResourceGuardProxyClientListOptions`
- New struct `DppResourceGuardProxyClientListResponse`
- New struct `DppResourceGuardProxyClientPutOptions`
- New struct `DppResourceGuardProxyClientPutResponse`
- New struct `DppResourceGuardProxyClientUnlockDeleteOptions`
- New struct `DppResourceGuardProxyClientUnlockDeleteResponse`
- New struct `ImmutabilitySettings`
- New struct `KubernetesClusterBackupDatasourceParameters`
- New struct `KubernetesClusterRestoreCriteria`
- New struct `ResourceGuardOperationDetail`
- New struct `ResourceGuardProxyBase`
- New struct `ResourceGuardProxyBaseResource`
- New struct `ResourceGuardProxyBaseResourceList`
- New struct `SecuritySettings`
- New struct `SoftDeleteSettings`
- New struct `UnlockDeleteRequest`
- New struct `UnlockDeleteResponse`
- New field `BackupDatasourceParametersList` in struct `PolicyParameters`
- New field `SecuritySettings` in struct `PatchBackupVaultInput`
- New field `Tags` in struct `BackupInstanceResource`
- New field `IsVaultProtectedByResourceGuard` in struct `BackupVault`
- New field `SecuritySettings` in struct `BackupVault`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).