# Release History

## 2.0.0-beta.1 (2022-10-04)
### Breaking Changes

- Type of `ConnectedClusterPatch.Properties` has been changed from `interface{}` to `*ConnectedClusterPatchProperties`

### Features Added

- New const `AzureHybridBenefitFalse`
- New const `AzureHybridBenefitNotApplicable`
- New const `PrivateLinkStateDisabled`
- New const `PrivateLinkStateEnabled`
- New const `AzureHybridBenefitTrue`
- New type alias `PrivateLinkState`
- New type alias `AzureHybridBenefit`
- New function `PossibleAzureHybridBenefitValues() []AzureHybridBenefit`
- New function `PossiblePrivateLinkStateValues() []PrivateLinkState`
- New struct `ConnectedClusterPatchProperties`
- New field `MiscellaneousProperties` in struct `ConnectedClusterProperties`
- New field `AzureHybridBenefit` in struct `ConnectedClusterProperties`
- New field `DistributionVersion` in struct `ConnectedClusterProperties`
- New field `PrivateLinkState` in struct `ConnectedClusterProperties`
- New field `PrivateLinkScopeResourceID` in struct `ConnectedClusterProperties`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).