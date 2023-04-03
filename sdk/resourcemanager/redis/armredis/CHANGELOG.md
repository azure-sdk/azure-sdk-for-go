# Release History

## 2.3.0 (2023-04-03)
### Features Added

- New enum type `AccessPolicyAssignmentProvisioningState` with values `AccessPolicyAssignmentProvisioningStateCanceled`, `AccessPolicyAssignmentProvisioningStateDeleted`, `AccessPolicyAssignmentProvisioningStateDeleting`, `AccessPolicyAssignmentProvisioningStateFailed`, `AccessPolicyAssignmentProvisioningStateSucceeded`, `AccessPolicyAssignmentProvisioningStateUpdating`
- New enum type `AccessPolicyProvisioningState` with values `AccessPolicyProvisioningStateCanceled`, `AccessPolicyProvisioningStateDeleted`, `AccessPolicyProvisioningStateDeleting`, `AccessPolicyProvisioningStateFailed`, `AccessPolicyProvisioningStateSucceeded`, `AccessPolicyProvisioningStateUpdating`
- New enum type `AccessPolicyType` with values `AccessPolicyTypeBuiltIn`, `AccessPolicyTypeCustom`
- New function `NewAccessPolicyAssignmentClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AccessPolicyAssignmentClient, error)`
- New function `*AccessPolicyAssignmentClient.BeginCreateUpdate(context.Context, string, string, string, CacheAccessPolicyAssignmentSet, *AccessPolicyAssignmentClientBeginCreateUpdateOptions) (*runtime.Poller[AccessPolicyAssignmentClientCreateUpdateResponse], error)`
- New function `*AccessPolicyAssignmentClient.BeginDelete(context.Context, string, string, string, *AccessPolicyAssignmentClientBeginDeleteOptions) (*runtime.Poller[AccessPolicyAssignmentClientDeleteResponse], error)`
- New function `*AccessPolicyAssignmentClient.Get(context.Context, string, string, string, *AccessPolicyAssignmentClientGetOptions) (AccessPolicyAssignmentClientGetResponse, error)`
- New function `*AccessPolicyAssignmentClient.NewListPager(string, string, *AccessPolicyAssignmentClientListOptions) *runtime.Pager[AccessPolicyAssignmentClientListResponse]`
- New function `NewAccessPolicyClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AccessPolicyClient, error)`
- New function `*AccessPolicyClient.BeginCreateUpdate(context.Context, string, string, string, CacheAccessPolicy, *AccessPolicyClientBeginCreateUpdateOptions) (*runtime.Poller[AccessPolicyClientCreateUpdateResponse], error)`
- New function `*AccessPolicyClient.BeginDelete(context.Context, string, string, string, *AccessPolicyClientBeginDeleteOptions) (*runtime.Poller[AccessPolicyClientDeleteResponse], error)`
- New function `*AccessPolicyClient.Get(context.Context, string, string, string, *AccessPolicyClientGetOptions) (AccessPolicyClientGetResponse, error)`
- New function `*AccessPolicyClient.NewListPager(string, string, *AccessPolicyClientListOptions) *runtime.Pager[AccessPolicyClientListResponse]`
- New function `*ClientFactory.NewAccessPolicyAssignmentClient() *AccessPolicyAssignmentClient`
- New function `*ClientFactory.NewAccessPolicyClient() *AccessPolicyClient`
- New struct `CacheAccessPolicy`
- New struct `CacheAccessPolicyAssignment`
- New struct `CacheAccessPolicyAssignmentList`
- New struct `CacheAccessPolicyAssignmentSet`
- New struct `CacheAccessPolicyAssignmentSetProperties`
- New struct `CacheAccessPolicyList`
- New struct `CacheAccessPolicyProperties`
- New field `AADEnabled` in struct `CommonPropertiesRedisConfiguration`


## 2.2.0 (2023-03-27)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 2.1.0 (2022-10-14)

### Features Added

- New field `PrimaryHostName` in struct `LinkedServerCreateProperties`
- New field `GeoReplicatedPrimaryHostName` in struct `LinkedServerCreateProperties`
- New field `GeoReplicatedPrimaryHostName` in struct `LinkedServerProperties`
- New field `PrimaryHostName` in struct `LinkedServerProperties`


## 2.0.0 (2022-09-01)
### Breaking Changes

- Operation `LinkedServerClient.Delete` has been changed to LRO, use `LinkedServerClient.BeginDelete` instead
- Operation `*Client.Update` has been changed to LRO, use `Client.BeginUpdate` instead

### Features Added

- New field `Authnotrequired` in struct `CommonPropertiesRedisConfiguration`
- New field `AofBackupEnabled` in struct `CommonPropertiesRedisConfiguration`
- New field `PreferredDataArchiveAuthMethod` in struct `ImportRDBParameters`
- New field `PreferredDataArchiveAuthMethod` in struct `ExportRDBParameters`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
