# Release History

## 0.6.0 (2023-03-22)
### Breaking Changes

- Type of `ExperimentProperties.Selectors` has been changed from `[]*Selector` to `[]SelectorClassification`
- Const `SelectorTypePercent`, `SelectorTypeRandom`, `SelectorTypeTag` from type alias `SelectorType` has been removed
- Operation `*ExperimentsClient.BeginCancel` has been changed to non-LRO, use `*ExperimentsClient.Cancel` instead.
- Operation `*ExperimentsClient.BeginCreateOrUpdate` has been changed to non-LRO, use `*ExperimentsClient.CreateOrUpdate` instead.
- Field `Targets` of struct `Selector` has been removed

### Features Added

- New value `SelectorTypeQuery` added to enum type `SelectorType`
- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewCapabilitiesClient() *CapabilitiesClient`
- New function `*ClientFactory.NewCapabilityTypesClient() *CapabilityTypesClient`
- New function `*ClientFactory.NewExperimentsClient() *ExperimentsClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewTargetTypesClient() *TargetTypesClient`
- New function `*ClientFactory.NewTargetsClient() *TargetsClient`
- New function `*Filter.GetFilter() *Filter`
- New function `*ListSelector.GetSelector() *Selector`
- New function `*QuerySelector.GetSelector() *Selector`
- New function `*Selector.GetSelector() *Selector`
- New function `*SimpleFilter.GetFilter() *Filter`
- New struct `CapabilityTypePropertiesRuntimeProperties`
- New struct `ClientFactory`
- New struct `ListSelector`
- New struct `QuerySelector`
- New struct `SimpleFilter`
- New struct `SimpleFilterParameters`
- New field `Kind` in struct `CapabilityTypeProperties`
- New field `RuntimeProperties` in struct `CapabilityTypeProperties`


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).