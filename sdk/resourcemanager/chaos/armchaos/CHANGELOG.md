# Release History

## 0.6.0 (2023-03-13)
### Breaking Changes

- Operation `*ExperimentsClient.BeginCancel` has been changed to non-LRO, use `*ExperimentsClient.Cancel` instead.
- Operation `*ExperimentsClient.BeginCreateOrUpdate` has been changed to non-LRO, use `*ExperimentsClient.CreateOrUpdate` instead.

### Features Added

- New type alias `FilterType` with values `FilterTypeSimple`
- New function `*Filter.GetFilter() *Filter`
- New function `*SimpleFilter.GetFilter() *Filter`
- New struct `CapabilityTypePropertiesRuntimeProperties`
- New struct `SimpleFilter`
- New struct `SimpleFilterParameters`
- New field `Kind` in struct `CapabilityTypeProperties`
- New field `RuntimeProperties` in struct `CapabilityTypeProperties`
- New field `Filter` in struct `Selector`


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).