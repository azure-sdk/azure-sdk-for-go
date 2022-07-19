# Release History

## 2.0.0-beta.1 (2022-07-19)
### Breaking Changes

- Const `DNSRefreshConfigurationPropertiesStatusFailed` has been removed
- Const `DNSRefreshConfigurationPropertiesStatusSucceeded` has been removed
- Function `PossibleDNSRefreshConfigurationPropertiesStatusValues` has been removed
- Function `*VirtualClustersClient.UpdateDNSServers` has been removed
- Struct `DNSRefreshConfigurationProperties` has been removed
- Struct `UpdateManagedInstanceDNSServersOperation` has been removed
- Struct `VirtualClustersClientUpdateDNSServersOptions` has been removed
- Field `MaintenanceConfigurationID` of struct `VirtualClusterProperties` has been removed
- Field `Family` of struct `VirtualClusterProperties` has been removed
- Field `UpdateManagedInstanceDNSServersOperation` of struct `VirtualClustersClientUpdateDNSServersResponse` has been removed

### Features Added

- New const `DNSRefreshOperationStatusInProgress`
- New const `DNSRefreshOperationStatusFailed`
- New const `DNSRefreshOperationStatusSucceeded`
- New function `*ManagedServerDNSAliasesClient.Get(context.Context, string, string, string, *ManagedServerDNSAliasesClientGetOptions) (ManagedServerDNSAliasesClientGetResponse, error)`
- New function `NewManagedServerDNSAliasesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagedServerDNSAliasesClient, error)`
- New function `*ManagedServerDNSAliasesClient.NewListByManagedInstancePager(string, string, *ManagedServerDNSAliasesClientListByManagedInstanceOptions) *runtime.Pager[ManagedServerDNSAliasesClientListByManagedInstanceResponse]`
- New function `*ManagedServerDNSAliasesClient.BeginAcquire(context.Context, string, string, string, ManagedServerDNSAliasAcquisition, *ManagedServerDNSAliasesClientBeginAcquireOptions) (*runtime.Poller[ManagedServerDNSAliasesClientAcquireResponse], error)`
- New function `*ManagedServerDNSAliasesClient.BeginDelete(context.Context, string, string, string, *ManagedServerDNSAliasesClientBeginDeleteOptions) (*runtime.Poller[ManagedServerDNSAliasesClientDeleteResponse], error)`
- New function `PossibleDNSRefreshOperationStatusValues() []DNSRefreshOperationStatus`
- New function `*VirtualClustersClient.BeginUpdateDNSServers(context.Context, string, string, *VirtualClustersClientBeginUpdateDNSServersOptions) (*runtime.Poller[VirtualClustersClientUpdateDNSServersResponse], error)`
- New function `*ManagedServerDNSAliasesClient.BeginCreateOrUpdate(context.Context, string, string, string, ManagedServerDNSAliasCreation, *ManagedServerDNSAliasesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedServerDNSAliasesClientCreateOrUpdateResponse], error)`
- New struct `ManagedServerDNSAlias`
- New struct `ManagedServerDNSAliasAcquisition`
- New struct `ManagedServerDNSAliasCreation`
- New struct `ManagedServerDNSAliasListResult`
- New struct `ManagedServerDNSAliasProperties`
- New struct `ManagedServerDNSAliasesClient`
- New struct `ManagedServerDNSAliasesClientAcquireResponse`
- New struct `ManagedServerDNSAliasesClientBeginAcquireOptions`
- New struct `ManagedServerDNSAliasesClientBeginCreateOrUpdateOptions`
- New struct `ManagedServerDNSAliasesClientBeginDeleteOptions`
- New struct `ManagedServerDNSAliasesClientCreateOrUpdateResponse`
- New struct `ManagedServerDNSAliasesClientDeleteResponse`
- New struct `ManagedServerDNSAliasesClientGetOptions`
- New struct `ManagedServerDNSAliasesClientGetResponse`
- New struct `ManagedServerDNSAliasesClientListByManagedInstanceOptions`
- New struct `ManagedServerDNSAliasesClientListByManagedInstanceResponse`
- New struct `UpdateVirtualClusterDNSServersOperation`
- New struct `VirtualClusterDNSServersProperties`
- New struct `VirtualClustersClientBeginUpdateDNSServersOptions`
- New field `Version` in struct `VirtualClusterProperties`
- New anonymous field `UpdateVirtualClusterDNSServersOperation` in struct `VirtualClustersClientUpdateDNSServersResponse`


## 1.0.0 (2022-06-02)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).