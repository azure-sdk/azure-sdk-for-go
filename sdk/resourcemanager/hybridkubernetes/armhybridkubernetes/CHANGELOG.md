# Release History

## 2.0.0 (2024-01-26)
### Breaking Changes

- Type of `ConnectedClusterPatch.Properties` has been changed from `any` to `*ConnectedClusterPatchProperties`

### Features Added

- New enum type `AutoUpgradeOptions` with values `AutoUpgradeOptionsDisabled`, `AutoUpgradeOptionsEnabled`
- New enum type `AzureHybridBenefit` with values `AzureHybridBenefitFalse`, `AzureHybridBenefitNotApplicable`, `AzureHybridBenefitTrue`
- New enum type `ConnectedClusterKind` with values `ConnectedClusterKindProvisionedCluster`
- New enum type `PrivateLinkState` with values `PrivateLinkStateDisabled`, `PrivateLinkStateEnabled`
- New struct `AADProfile`
- New struct `ArcAgentProfile`
- New struct `ConnectedClusterPatchProperties`
- New field `Kind` in struct `ConnectedCluster`
- New field `AADProfile`, `ArcAgentProfile`, `AzureHybridBenefit`, `DistributionVersion`, `MiscellaneousProperties`, `PrivateLinkScopeResourceID`, `PrivateLinkState` in struct `ConnectedClusterProperties`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).