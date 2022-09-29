# Release History

## 2.0.0-beta.1 (2022-09-29)
### Breaking Changes

- Type of `ResourceGuard.ProvisioningState` has been changed from `*ProvisioningState` to `*ResourceGuardProvisioningState`
- Function `*Client.CheckFeatureSupport` has been removed
- Struct `CloudError` has been removed

### Features Added

- New const `ImmutabilityStateLocked`
- New const `ResourceGuardProvisioningStateProvisioning`
- New const `SoftDeleteStateOff`
- New const `ResourceGuardProvisioningStateFailed`
- New const `ResourceGuardProvisioningStateUpdating`
- New const `StorageSettingTypesZoneRedundant`
- New const `ResourceGuardProvisioningStateUnknown`
- New const `SoftDeleteStateAlwaysOn`
- New const `ImmutabilityStateDisabled`
- New const `ResourceGuardProvisioningStateSucceeded`
- New const `SoftDeleteStateOn`
- New const `ImmutabilityStateUnlocked`
- New type alias `ResourceGuardProvisioningState`
- New type alias `ImmutabilityState`
- New type alias `SoftDeleteState`
- New function `PossibleSoftDeleteStateValues() []SoftDeleteState`
- New function `*DeletedBackupInstancesClient.NewListPager(string, string, *DeletedBackupInstancesClientListOptions) *runtime.Pager[DeletedBackupInstancesClientListResponse]`
- New function `*DeletedBackupInstancesClient.Get(context.Context, string, string, string, *DeletedBackupInstancesClientGetOptions) (DeletedBackupInstancesClientGetResponse, error)`
- New function `*Client.NewDppResourceGuardProxyListPager(string, string, *ClientDppResourceGuardProxyListOptions) *runtime.Pager[ClientDppResourceGuardProxyListResponse]`
- New function `PossibleResourceGuardProvisioningStateValues() []ResourceGuardProvisioningState`
- New function `*Client.DppResourceGuardProxyUnlockDelete(context.Context, string, string, string, UnlockDeleteRequest, *ClientDppResourceGuardProxyUnlockDeleteOptions) (ClientDppResourceGuardProxyUnlockDeleteResponse, error)`
- New function `*Client.DppResourceGuardProxyDelete(context.Context, string, string, string, *ClientDppResourceGuardProxyDeleteOptions) (ClientDppResourceGuardProxyDeleteResponse, error)`
- New function `*Client.DppResourceGuardProxyGet(context.Context, string, string, string, *ClientDppResourceGuardProxyGetOptions) (ClientDppResourceGuardProxyGetResponse, error)`
- New function `*Client.DppResourceGuardProxyPut(context.Context, string, string, string, ResourceGuardProxyBaseResource, *ClientDppResourceGuardProxyPutOptions) (ClientDppResourceGuardProxyPutResponse, error)`
- New function `*BackupInstancesExtensionRoutingClient.NewListPager(string, *BackupInstancesExtensionRoutingClientListOptions) *runtime.Pager[BackupInstancesExtensionRoutingClientListResponse]`
- New function `PossibleImmutabilityStateValues() []ImmutabilityState`
- New function `*DeletedBackupInstancesClient.BeginUndelete(context.Context, string, string, string, *DeletedBackupInstancesClientBeginUndeleteOptions) (*runtime.Poller[DeletedBackupInstancesClientUndeleteResponse], error)`
- New function `NewBackupInstancesExtensionRoutingClient(azcore.TokenCredential, *arm.ClientOptions) (*BackupInstancesExtensionRoutingClient, error)`
- New function `NewDeletedBackupInstancesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DeletedBackupInstancesClient, error)`
- New struct `BackupInstancesExtensionRoutingClient`
- New struct `BackupInstancesExtensionRoutingClientListOptions`
- New struct `BackupInstancesExtensionRoutingClientListResponse`
- New struct `ClientDppResourceGuardProxyDeleteOptions`
- New struct `ClientDppResourceGuardProxyDeleteResponse`
- New struct `ClientDppResourceGuardProxyGetOptions`
- New struct `ClientDppResourceGuardProxyGetResponse`
- New struct `ClientDppResourceGuardProxyListOptions`
- New struct `ClientDppResourceGuardProxyListResponse`
- New struct `ClientDppResourceGuardProxyPutOptions`
- New struct `ClientDppResourceGuardProxyPutResponse`
- New struct `ClientDppResourceGuardProxyUnlockDeleteOptions`
- New struct `ClientDppResourceGuardProxyUnlockDeleteResponse`
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
- New struct `ImmutabilitySettings`
- New struct `ResourceGuardOperationDetail`
- New struct `ResourceGuardProxyBase`
- New struct `ResourceGuardProxyBaseResource`
- New struct `ResourceGuardProxyBaseResourceList`
- New struct `SecuritySettings`
- New struct `SoftDeleteSettings`
- New struct `UnlockDeleteRequest`
- New struct `UnlockDeleteResponse`
- New field `SecuritySettings` in struct `PatchBackupVaultInput`
- New field `Tags` in struct `BackupInstanceResource`
- New field `IsVaultProtectedByResourceGuard` in struct `BackupVault`
- New field `SecuritySettings` in struct `BackupVault`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).