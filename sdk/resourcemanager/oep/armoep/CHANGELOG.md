# Release History

## 0.5.0 (2022-07-18)
### Breaking Changes

- Function `*EnergyServicesClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, *EnergyServicesClientBeginCreateOptions)` to `(context.Context, string, string, EnergyService, *EnergyServicesClientBeginCreateOptions)`
- Function `*EnergyServicesClient.Update` parameter(s) have been changed from `(context.Context, string, string, *EnergyServicesClientUpdateOptions)` to `(context.Context, string, string, EnergyResourceUpdate, *EnergyServicesClientUpdateOptions)`
- Field `Name` of struct `DataPartitionAddOrRemoveRequest` has been removed
- Field `Body` of struct `EnergyServicesClientBeginCreateOptions` has been removed
- Field `Body` of struct `EnergyServicesClientUpdateOptions` has been removed
- Field `Value` of struct `DataPartitionsListResult` has been removed

### Features Added

- New anonymous field `DataPartitionAddOrRemoveRequest` in struct `EnergyServicesClientAddPartitionResponse`
- New field `DataPartitionName` in struct `DataPartitionAddOrRemoveRequest`
- New field `DataPartitionInfo` in struct `DataPartitionsListResult`
- New anonymous field `DataPartitionAddOrRemoveRequest` in struct `EnergyServicesClientRemovePartitionResponse`


## 0.4.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oep/armoep` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.4.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).