# Release History

## 2.0.0 (2022-09-20)
### Breaking Changes

- Function `*Client.CheckFeatureSupport` has been removed
- Struct `CloudError` has been removed

### Features Added

- New const `ProvisioningStateDeleting`
- New const `ProvisioningStateSubscriptionSuspended`
- New const `StorageSettingTypesZoneRedundant`
- New function `*Client.DppResourceGuardProxyGet(context.Context, string, string, string, *ClientDppResourceGuardProxyGetOptions) (ClientDppResourceGuardProxyGetResponse, error)`
- New function `*Client.DppResourceGuardProxyDelete(context.Context, string, string, string, *ClientDppResourceGuardProxyDeleteOptions) (ClientDppResourceGuardProxyDeleteResponse, error)`
- New function `*Client.NewDppResourceGuardProxyListPager(string, string, *ClientDppResourceGuardProxyListOptions) *runtime.Pager[ClientDppResourceGuardProxyListResponse]`
- New function `*Client.DppResourceGuardProxyUnlockDelete(context.Context, string, string, string, UnlockDeleteRequest, *ClientDppResourceGuardProxyUnlockDeleteOptions) (ClientDppResourceGuardProxyUnlockDeleteResponse, error)`
- New function `*Client.DppResourceGuardProxyPut(context.Context, string, string, string, ResourceGuardProxyBaseResource, *ClientDppResourceGuardProxyPutOptions) (ClientDppResourceGuardProxyPutResponse, error)`
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
- New struct `DppProxyResource`
- New struct `ResourceGuardOperationDetail`
- New struct `ResourceGuardProxyBase`
- New struct `ResourceGuardProxyBaseResource`
- New struct `ResourceGuardProxyBaseResourceList`
- New struct `UnlockDeleteRequest`
- New struct `UnlockDeleteResponse`
- New field `Tags` in struct `BackupInstanceResource`
- New field `IsVaultProtectedByResourceGuard` in struct `BackupVault`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).