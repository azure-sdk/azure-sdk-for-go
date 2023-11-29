# Release History

## 2.0.0 (2023-11-29)
### Breaking Changes

- Function `*ServersClient.BeginResume` parameter(s) have been changed from `(context.Context, string, string, *ServersClientBeginResumeOptions)` to `(context.Context, string, string, any, *ServersClientBeginResumeOptions)`
- Function `*ServersClient.BeginSuspend` parameter(s) have been changed from `(context.Context, string, string, *ServersClientBeginSuspendOptions)` to `(context.Context, string, string, any, *ServersClientBeginSuspendOptions)`
- Function `*ServersClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, ServerUpdateParameters, *ServersClientBeginUpdateOptions)` to `(context.Context, string, string, ServerUpdate, *ServersClientBeginUpdateOptions)`
- Function `*ServersClient.DissociateGateway` parameter(s) have been changed from `(context.Context, string, string, *ServersClientDissociateGatewayOptions)` to `(context.Context, string, string, any, *ServersClientDissociateGatewayOptions)`
- Function `*ServersClient.ListGatewayStatus` parameter(s) have been changed from `(context.Context, string, string, *ServersClientListGatewayStatusOptions)` to `(context.Context, string, string, any, *ServersClientListGatewayStatusOptions)`
- Function `*ServersClient.ListSKUsForExisting` parameter(s) have been changed from `(context.Context, string, string, *ServersClientListSKUsForExistingOptions)` to `(context.Context, string, string, string, *ServersClientListSKUsForExistingOptions)`
- Type of `GatewayListStatusLive.Status` has been changed from `*int32` to `*GatewayListStatusLiveStatus`
- Type of `Operation.Origin` has been changed from `*string` to `*Origin`
- Type of `SKUDetailsForExistingResource.SKU` has been changed from `*ResourceSKU` to `*AzureResourceManagerResourceSKU`
- Type of `ServerProperties.SKU` has been changed from `*ResourceSKU` to `*AzureResourceManagerResourceSKU`
- Enum `ManagedMode` has been removed
- Enum `ServerMonitorMode` has been removed
- Function `*ServersClient.CheckNameAvailability` has been removed
- Function `*ServersClient.ListOperationResults` has been removed
- Function `*ServersClient.ListOperationStatuses` has been removed
- Function `*ServersClient.NewListPager` has been removed
- Function `*ServersClient.ListSKUsForNew` has been removed
- Struct `GatewayListStatusError` has been removed
- Struct `ResourceSKU` has been removed
- Struct `SKUEnumerationForNewResourceResult` has been removed
- Struct `ServerMutableProperties` has been removed
- Struct `ServerUpdateParameters` has been removed
- Struct `Servers` has been removed
- Field `HTTPStatusCode`, `SubCode`, `TimeStamp` of struct `ErrorDetail` has been removed
- Field `Properties` of struct `Operation` has been removed
- Field `Error` of struct `OperationStatus` has been removed
- Field `Location`, `SKU`, `Tags` of struct `Resource` has been removed
- Field `SKU` of struct `Server` has been removed
- Field `AsAdministrators`, `BackupBlobContainerURI`, `GatewayDetails`, `IPV4FirewallSettings`, `ManagedMode`, `QuerypoolConnectionMode`, `ServerMonitorMode` of struct `ServerProperties` has been removed
- Field `Servers` of struct `ServersClientListByResourceGroupResponse` has been removed

### Features Added

- New value `SKUTierFree`, `SKUTierPremium` added to enum type `SKUTier`
- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `GatewayListStatusLiveStatus` with values `GatewayListStatusLiveStatusZero`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New enum type `Versions` with values `VersionsV20170801`
- New struct `AzureCoreFoundationsError`
- New struct `AzureCoreFoundationsErrorResponse`
- New struct `AzureCoreFoundationsInnerError`
- New struct `AzureResourceManagerResourceSKU`
- New struct `AzureResourceManagerResourceSKUUpdate`
- New struct `SKU`
- New struct `ServerListResult`
- New struct `ServerUpdate`
- New struct `ServerUpdateProperties`
- New struct `SystemData`
- New struct `TrackedResource`
- New field `ActionType`, `IsDataAction` in struct `Operation`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `Server`
- New field `Interface` in struct `ServersClientDissociateGatewayResponse`
- New anonymous field `ServerListResult` in struct `ServersClientListByResourceGroupResponse`
- New field `Interface` in struct `ServersClientResumeResponse`
- New field `Interface` in struct `ServersClientSuspendResponse`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.

## 1.1.0 (2023-03-27)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).