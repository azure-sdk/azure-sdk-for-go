# Release History

## 0.6.0 (2023-03-17)
### Breaking Changes

- Operation `*ExperimentsClient.Start` has been changed to LRO, use `*ExperimentsClient.BeginStart` instead.

### Features Added

- New value `ResourceIdentityTypeUserAssigned` added to type alias `ResourceIdentityType`
- New function `*ExperimentsClient.Update(context.Context, string, string, ExperimentUpdate, *ExperimentsClientUpdateOptions) (ExperimentsClientUpdateResponse, error)`
- New struct `CapabilityTypePropertiesRuntimeProperties`
- New struct `ComponentsEwb5TmSchemasUserassignedidentitiesAdditionalproperties`
- New struct `ExperimentUpdate`
- New field `Kind` in struct `CapabilityTypeProperties`
- New field `RuntimeProperties` in struct `CapabilityTypeProperties`
- New field `UserAssignedIdentities` in struct `ResourceIdentity`


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).