# Release History

## 0.2.0 (2024-07-17)
### Breaking Changes

- Type of `ErrorAdditionalInfo.Info` has been changed from `any` to `*ErrorAdditionalInfoInfo`
- Type of `OperationStatusResult.PercentComplete` has been changed from `*float32` to `*float64`
- Enum `ActionType` has been removed
- Enum `Origin` has been removed
- Function `*ClientFactory.NewOperationsClient` has been removed
- Function `*ClientFactory.NewSapDiscoverySitesClient` has been removed
- Function `*ClientFactory.NewSapInstancesClient` has been removed
- Function `NewOperationsClient` has been removed
- Function `*OperationsClient.NewListPager` has been removed
- Function `NewSapDiscoverySitesClient` has been removed
- Function `*SapDiscoverySitesClient.BeginCreate` has been removed
- Function `*SapDiscoverySitesClient.BeginDelete` has been removed
- Function `*SapDiscoverySitesClient.Get` has been removed
- Function `*SapDiscoverySitesClient.BeginImportEntities` has been removed
- Function `*SapDiscoverySitesClient.NewListByResourceGroupPager` has been removed
- Function `*SapDiscoverySitesClient.NewListBySubscriptionPager` has been removed
- Function `*SapDiscoverySitesClient.Update` has been removed
- Function `NewSapInstancesClient` has been removed
- Function `*SapInstancesClient.BeginCreate` has been removed
- Function `*SapInstancesClient.BeginDelete` has been removed
- Function `*SapInstancesClient.Get` has been removed
- Function `*SapInstancesClient.NewListBySapDiscoverySitePager` has been removed
- Function `*SapInstancesClient.Update` has been removed
- Function `*ServerInstancesClient.NewListBySapInstancePager` has been removed
- Struct `Operation` has been removed
- Struct `OperationDisplay` has been removed
- Struct `OperationListResult` has been removed

### Features Added

- New function `*ClientFactory.NewSAPDiscoverySitesClient() *SAPDiscoverySitesClient`
- New function `*ClientFactory.NewSAPInstancesClient() *SAPInstancesClient`
- New function `NewSAPDiscoverySitesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SAPDiscoverySitesClient, error)`
- New function `*SAPDiscoverySitesClient.BeginCreate(context.Context, string, string, SAPDiscoverySite, *SAPDiscoverySitesClientBeginCreateOptions) (*runtime.Poller[SAPDiscoverySitesClientCreateResponse], error)`
- New function `*SAPDiscoverySitesClient.BeginDelete(context.Context, string, string, *SAPDiscoverySitesClientBeginDeleteOptions) (*runtime.Poller[SAPDiscoverySitesClientDeleteResponse], error)`
- New function `*SAPDiscoverySitesClient.Get(context.Context, string, string, *SAPDiscoverySitesClientGetOptions) (SAPDiscoverySitesClientGetResponse, error)`
- New function `*SAPDiscoverySitesClient.BeginImportEntities(context.Context, string, string, *SAPDiscoverySitesClientBeginImportEntitiesOptions) (*runtime.Poller[SAPDiscoverySitesClientImportEntitiesResponse], error)`
- New function `*SAPDiscoverySitesClient.NewListByResourceGroupPager(string, *SAPDiscoverySitesClientListByResourceGroupOptions) *runtime.Pager[SAPDiscoverySitesClientListByResourceGroupResponse]`
- New function `*SAPDiscoverySitesClient.NewListBySubscriptionPager(*SAPDiscoverySitesClientListBySubscriptionOptions) *runtime.Pager[SAPDiscoverySitesClientListBySubscriptionResponse]`
- New function `*SAPDiscoverySitesClient.Update(context.Context, string, string, SAPDiscoverySiteTagsUpdate, *SAPDiscoverySitesClientUpdateOptions) (SAPDiscoverySitesClientUpdateResponse, error)`
- New function `NewSAPInstancesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SAPInstancesClient, error)`
- New function `*SAPInstancesClient.BeginCreate(context.Context, string, string, string, SAPInstance, *SAPInstancesClientBeginCreateOptions) (*runtime.Poller[SAPInstancesClientCreateResponse], error)`
- New function `*SAPInstancesClient.BeginDelete(context.Context, string, string, string, *SAPInstancesClientBeginDeleteOptions) (*runtime.Poller[SAPInstancesClientDeleteResponse], error)`
- New function `*SAPInstancesClient.Get(context.Context, string, string, string, *SAPInstancesClientGetOptions) (SAPInstancesClientGetResponse, error)`
- New function `*SAPInstancesClient.NewListBySAPDiscoverySitePager(string, string, *SAPInstancesClientListBySAPDiscoverySiteOptions) *runtime.Pager[SAPInstancesClientListBySAPDiscoverySiteResponse]`
- New function `*SAPInstancesClient.Update(context.Context, string, string, string, SAPInstanceTagsUpdate, *SAPInstancesClientUpdateOptions) (SAPInstancesClientUpdateResponse, error)`
- New function `*ServerInstancesClient.NewListBySAPInstancePager(string, string, string, *ServerInstancesClientListBySAPInstanceOptions) *runtime.Pager[ServerInstancesClientListBySAPInstanceResponse]`
- New struct `ErrorAdditionalInfoInfo`


## 0.1.0 (2024-03-22)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationdiscovery/armmigrationdiscoverysap` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).