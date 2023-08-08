# Release History

## 0.7.0 (2023-08-08)
### Breaking Changes

- Enum `CreatedByType` has been removed
- Enum `ResourceIdentityType` has been removed
- Enum `SigningKey` has been removed
- Function `NewAccountsClient` has been removed
- Function `*AccountsClient.CreateOrUpdate` has been removed
- Function `*AccountsClient.Delete` has been removed
- Function `*AccountsClient.Get` has been removed
- Function `*AccountsClient.NewListByResourceGroupPager` has been removed
- Function `*AccountsClient.NewListBySubscriptionPager` has been removed
- Function `*AccountsClient.ListKeys` has been removed
- Function `*AccountsClient.ListSas` has been removed
- Function `*AccountsClient.RegenerateKeys` has been removed
- Function `*AccountsClient.Update` has been removed
- Function `NewClient` has been removed
- Function `*Client.NewListOperationsPager` has been removed
- Function `*Client.NewListSubscriptionOperationsPager` has been removed
- Function `*ClientFactory.NewAccountsClient` has been removed
- Function `*ClientFactory.NewClient` has been removed
- Function `timeRFC3339.MarshalText` has been removed
- Function `*timeRFC3339.Parse` has been removed
- Function `*timeRFC3339.UnmarshalText` has been removed
- Struct `Account` has been removed
- Struct `AccountKeys` has been removed
- Struct `AccountProperties` has been removed
- Struct `AccountSasParameters` has been removed
- Struct `AccountSasToken` has been removed
- Struct `AccountUpdateParameters` has been removed
- Struct `Accounts` has been removed
- Struct `Components1Jq1T4ISchemasManagedserviceidentityPropertiesUserassignedidentitiesAdditionalproperties` has been removed
- Struct `CorsRule` has been removed
- Struct `CorsRules` has been removed
- Struct `LinkedResource` has been removed
- Struct `ManagedServiceIdentity` has been removed
- Struct `SystemData` has been removed
- Field `SystemData` of struct `Creator` has been removed
- Field `InternalMetricName`, `SourceMdmAccount` of struct `MetricSpecification` has been removed

### Features Added

- New field `ConsumedStorageUnitSizeInBytes`, `TotalStorageUnitSizeInBytes` in struct `CreatorProperties`


## 0.6.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 0.6.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).