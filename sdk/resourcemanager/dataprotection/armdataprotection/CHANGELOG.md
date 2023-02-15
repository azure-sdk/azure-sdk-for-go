# Release History

## 2.0.0-beta.2 (2023-02-15)
### Breaking Changes

- Function `*ResourceGuardsClient.Patch` parameter(s) have been changed from `(context.Context, string, string, PatchResourceRequestInput, *ResourceGuardsClientPatchOptions)` to `(context.Context, string, string, PatchResourceGuardInput, *ResourceGuardsClientPatchOptions)`
- Type of `ResourceGuard.ProvisioningState` has been changed from `*ResourceGuardProvisioningState` to `*ProvisioningState`
- Const `StorageSettingStoreTypesSnapshotStore` from type alias `StorageSettingStoreTypes` has been removed
- Type alias `ResourceGuardProvisioningState` has been removed
- Function `NewBackupInstancesExtensionRoutingClient` has been removed
- Function `*BackupInstancesExtensionRoutingClient.NewListPager` has been removed
- Function `NewDppResourceGuardProxyClient` has been removed
- Function `*DppResourceGuardProxyClient.Delete` has been removed
- Function `*DppResourceGuardProxyClient.Get` has been removed
- Function `*DppResourceGuardProxyClient.NewListPager` has been removed
- Function `*DppResourceGuardProxyClient.Put` has been removed
- Function `*DppResourceGuardProxyClient.UnlockDelete` has been removed
- Operation `*BackupVaultsClient.Delete` has been changed to LRO, use `*BackupVaultsClient.BeginDelete` instead.
- Struct `BackupInstancesExtensionRoutingClient` has been removed
- Struct `DppResourceGuardProxyClient` has been removed
- Struct `ResourceGuardOperationDetail` has been removed
- Struct `ResourceGuardProxyBase` has been removed
- Struct `ResourceGuardProxyBaseResource` has been removed
- Struct `ResourceGuardProxyBaseResourceList` has been removed
- Struct `UnlockDeleteRequest` has been removed
- Struct `UnlockDeleteResponse` has been removed
- Field `Identity` of struct `ResourceGuardResource` has been removed

### Features Added

- New value `SourceDataStoreTypeOperationalStore` added to type alias `SourceDataStoreType`
- New value `StorageSettingStoreTypesOperationalStore` added to type alias `StorageSettingStoreTypes`
- New type alias `CrossSubscriptionRestoreState` with values `CrossSubscriptionRestoreStateDisabled`, `CrossSubscriptionRestoreStateEnabled`, `CrossSubscriptionRestoreStatePermanentlyDisabled`
- New type alias `EncryptionState` with values `EncryptionStateDisabled`, `EncryptionStateEnabled`, `EncryptionStateInconsistent`
- New type alias `IdentityType` with values `IdentityTypeSystemAssigned`, `IdentityTypeUserAssigned`
- New type alias `InfrastructureEncryptionState` with values `InfrastructureEncryptionStateDisabled`, `InfrastructureEncryptionStateEnabled`
- New struct `CmkKekIdentity`
- New struct `CmkKeyVaultProperties`
- New struct `CrossSubscriptionRestoreSettings`
- New struct `DppBaseTrackedResource`
- New struct `EncryptionSettings`
- New struct `FeatureSettings`
- New struct `PatchResourceGuardInput`
- New field `ExpiryTime` in struct `AzureBackupDiscreteRecoveryPoint`
- New field `FeatureSettings` in struct `BackupVault`
- New field `FeatureSettings` in struct `PatchBackupVaultInput`
- New field `EncryptionSettings` in struct `SecuritySettings`
- New field `TargetResourceArmID` in struct `TargetDetails`


## 2.0.0-beta.1 (2023-01-13)
### Breaking Changes

- Type of `ResourceGuard.ProvisioningState` has been changed from `*ProvisioningState` to `*ResourceGuardProvisioningState`

### Features Added

- New value `StorageSettingTypesZoneRedundant` added to type alias `StorageSettingTypes`
- New type alias `ExistingResourcePolicy` with values `ExistingResourcePolicyPatch`, `ExistingResourcePolicySkip`
- New type alias `ImmutabilityState` with values `ImmutabilityStateDisabled`, `ImmutabilityStateLocked`, `ImmutabilityStateUnlocked`
- New type alias `PersistentVolumeRestoreMode` with values `PersistentVolumeRestoreModeRestoreWithVolumeData`, `PersistentVolumeRestoreModeRestoreWithoutVolumeData`
- New type alias `ResourceGuardProvisioningState` with values `ResourceGuardProvisioningStateFailed`, `ResourceGuardProvisioningStateProvisioning`, `ResourceGuardProvisioningStateSucceeded`, `ResourceGuardProvisioningStateUnknown`, `ResourceGuardProvisioningStateUpdating`
- New type alias `SoftDeleteState` with values `SoftDeleteStateAlwaysOn`, `SoftDeleteStateOff`, `SoftDeleteStateOn`
- New function `*BackupDatasourceParameters.GetBackupDatasourceParameters() *BackupDatasourceParameters`
- New function `NewBackupInstancesExtensionRoutingClient(azcore.TokenCredential, *arm.ClientOptions) (*BackupInstancesExtensionRoutingClient, error)`
- New function `*BackupInstancesExtensionRoutingClient.NewListPager(string, *BackupInstancesExtensionRoutingClientListOptions) *runtime.Pager[BackupInstancesExtensionRoutingClientListResponse]`
- New function `*BlobBackupDatasourceParameters.GetBackupDatasourceParameters() *BackupDatasourceParameters`
- New function `NewDeletedBackupInstancesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DeletedBackupInstancesClient, error)`
- New function `*DeletedBackupInstancesClient.Get(context.Context, string, string, string, *DeletedBackupInstancesClientGetOptions) (DeletedBackupInstancesClientGetResponse, error)`
- New function `*DeletedBackupInstancesClient.NewListPager(string, string, *DeletedBackupInstancesClientListOptions) *runtime.Pager[DeletedBackupInstancesClientListResponse]`
- New function `*DeletedBackupInstancesClient.BeginUndelete(context.Context, string, string, string, *DeletedBackupInstancesClientBeginUndeleteOptions) (*runtime.Poller[DeletedBackupInstancesClientUndeleteResponse], error)`
- New function `NewDppResourceGuardProxyClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DppResourceGuardProxyClient, error)`
- New function `*DppResourceGuardProxyClient.Delete(context.Context, string, string, string, *DppResourceGuardProxyClientDeleteOptions) (DppResourceGuardProxyClientDeleteResponse, error)`
- New function `*DppResourceGuardProxyClient.Get(context.Context, string, string, string, *DppResourceGuardProxyClientGetOptions) (DppResourceGuardProxyClientGetResponse, error)`
- New function `*DppResourceGuardProxyClient.NewListPager(string, string, *DppResourceGuardProxyClientListOptions) *runtime.Pager[DppResourceGuardProxyClientListResponse]`
- New function `*DppResourceGuardProxyClient.Put(context.Context, string, string, string, ResourceGuardProxyBaseResource, *DppResourceGuardProxyClientPutOptions) (DppResourceGuardProxyClientPutResponse, error)`
- New function `*DppResourceGuardProxyClient.UnlockDelete(context.Context, string, string, string, UnlockDeleteRequest, *DppResourceGuardProxyClientUnlockDeleteOptions) (DppResourceGuardProxyClientUnlockDeleteResponse, error)`
- New function `*ItemPathBasedRestoreCriteria.GetItemLevelRestoreCriteria() *ItemLevelRestoreCriteria`
- New function `*KubernetesClusterBackupDatasourceParameters.GetBackupDatasourceParameters() *BackupDatasourceParameters`
- New function `*KubernetesClusterRestoreCriteria.GetItemLevelRestoreCriteria() *ItemLevelRestoreCriteria`
- New struct `BackupInstancesExtensionRoutingClient`
- New struct `BackupInstancesExtensionRoutingClientListResponse`
- New struct `BlobBackupDatasourceParameters`
- New struct `DeletedBackupInstance`
- New struct `DeletedBackupInstanceResource`
- New struct `DeletedBackupInstanceResourceList`
- New struct `DeletedBackupInstancesClient`
- New struct `DeletedBackupInstancesClientListResponse`
- New struct `DeletedBackupInstancesClientUndeleteResponse`
- New struct `DeletionInfo`
- New struct `DppProxyResource`
- New struct `DppResourceGuardProxyClient`
- New struct `DppResourceGuardProxyClientListResponse`
- New struct `ImmutabilitySettings`
- New struct `ItemPathBasedRestoreCriteria`
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
- New field `Tags` in struct `BackupInstanceResource`
- New field `IsVaultProtectedByResourceGuard` in struct `BackupVault`
- New field `SecuritySettings` in struct `BackupVault`
- New field `SecuritySettings` in struct `PatchBackupVaultInput`
- New field `BackupDatasourceParametersList` in struct `PolicyParameters`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).