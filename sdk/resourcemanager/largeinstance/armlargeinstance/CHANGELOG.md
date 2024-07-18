# Release History

## 1.0.0 (2024-07-18)
### Breaking Changes

- Type of `ErrorAdditionalInfo.Info` has been changed from `any` to `*ErrorAdditionalInfoInfo`
- Type of `OperationStatusResult.PercentComplete` has been changed from `*float32` to `*float64`
- Function `NewAzureLargeInstanceClient` has been removed
- Function `*AzureLargeInstanceClient.Get` has been removed
- Function `*AzureLargeInstanceClient.NewListByResourceGroupPager` has been removed
- Function `*AzureLargeInstanceClient.NewListBySubscriptionPager` has been removed
- Function `*AzureLargeInstanceClient.BeginRestart` has been removed
- Function `*AzureLargeInstanceClient.BeginShutdown` has been removed
- Function `*AzureLargeInstanceClient.BeginStart` has been removed
- Function `*AzureLargeInstanceClient.Update` has been removed
- Function `NewAzureLargeStorageInstanceClient` has been removed
- Function `*AzureLargeStorageInstanceClient.Get` has been removed
- Function `*AzureLargeStorageInstanceClient.NewListByResourceGroupPager` has been removed
- Function `*AzureLargeStorageInstanceClient.NewListBySubscriptionPager` has been removed
- Function `*AzureLargeStorageInstanceClient.Update` has been removed
- Function `*ClientFactory.NewAzureLargeInstanceClient` has been removed
- Function `*ClientFactory.NewAzureLargeStorageInstanceClient` has been removed
- Field `PartnerNodeID` of struct `AzureLargeInstanceProperties` has been removed
- Field `ResourceID` of struct `OperationStatusResult` has been removed

### Features Added

- New function `NewAzureLargeInstancesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AzureLargeInstancesClient, error)`
- New function `*AzureLargeInstancesClient.Get(context.Context, string, string, *AzureLargeInstancesClientGetOptions) (AzureLargeInstancesClientGetResponse, error)`
- New function `*AzureLargeInstancesClient.NewListByResourceGroupPager(string, *AzureLargeInstancesClientListByResourceGroupOptions) *runtime.Pager[AzureLargeInstancesClientListByResourceGroupResponse]`
- New function `*AzureLargeInstancesClient.NewListBySubscriptionPager(*AzureLargeInstancesClientListBySubscriptionOptions) *runtime.Pager[AzureLargeInstancesClientListBySubscriptionResponse]`
- New function `*AzureLargeInstancesClient.BeginRestart(context.Context, string, string, *AzureLargeInstancesClientBeginRestartOptions) (*runtime.Poller[AzureLargeInstancesClientRestartResponse], error)`
- New function `*AzureLargeInstancesClient.BeginShutdown(context.Context, string, string, *AzureLargeInstancesClientBeginShutdownOptions) (*runtime.Poller[AzureLargeInstancesClientShutdownResponse], error)`
- New function `*AzureLargeInstancesClient.BeginStart(context.Context, string, string, *AzureLargeInstancesClientBeginStartOptions) (*runtime.Poller[AzureLargeInstancesClientStartResponse], error)`
- New function `*AzureLargeInstancesClient.Update(context.Context, string, string, AzureLargeInstanceTagsUpdate, *AzureLargeInstancesClientUpdateOptions) (AzureLargeInstancesClientUpdateResponse, error)`
- New function `NewAzureLargeStorageInstancesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AzureLargeStorageInstancesClient, error)`
- New function `*AzureLargeStorageInstancesClient.Get(context.Context, string, string, *AzureLargeStorageInstancesClientGetOptions) (AzureLargeStorageInstancesClientGetResponse, error)`
- New function `*AzureLargeStorageInstancesClient.NewListByResourceGroupPager(string, *AzureLargeStorageInstancesClientListByResourceGroupOptions) *runtime.Pager[AzureLargeStorageInstancesClientListByResourceGroupResponse]`
- New function `*AzureLargeStorageInstancesClient.NewListBySubscriptionPager(*AzureLargeStorageInstancesClientListBySubscriptionOptions) *runtime.Pager[AzureLargeStorageInstancesClientListBySubscriptionResponse]`
- New function `*AzureLargeStorageInstancesClient.Update(context.Context, string, string, AzureLargeStorageInstanceTagsUpdate, *AzureLargeStorageInstancesClientUpdateOptions) (AzureLargeStorageInstancesClientUpdateResponse, error)`
- New function `*ClientFactory.NewAzureLargeInstancesClient() *AzureLargeInstancesClient`
- New function `*ClientFactory.NewAzureLargeStorageInstancesClient() *AzureLargeStorageInstancesClient`
- New struct `ErrorAdditionalInfoInfo`


## 0.1.0 (2024-02-23)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/largeinstance/armlargeinstance` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).