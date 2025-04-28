# Release History

## 2.0.0 (2025-04-28)
### Breaking Changes

- Type of `SystemData.CreatedByType` has been changed from `*IdentityType` to `*CreatedByType`
- Type of `SystemData.LastModifiedByType` has been changed from `*IdentityType` to `*CreatedByType`
- Enum `IdentityType` has been removed
- Operation `*CapacitiesClient.NewListByResourceGroupPager` does not support pagination anymore, use `*CapacitiesClient.ListByResourceGroup` instead.
- Operation `*CapacitiesClient.NewListPager` does not support pagination anymore, use `*CapacitiesClient.List` instead.
- Field `Location`, `Tags` of struct `Resource` has been removed

### Features Added

- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New struct `TrackedResource`
- New field `NextLink` in struct `AutoScaleVCoreListResult`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).