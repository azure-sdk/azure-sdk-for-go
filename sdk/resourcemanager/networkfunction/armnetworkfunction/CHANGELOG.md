# Release History

## 2.0.0 (2022-08-26)
### Breaking Changes

- Type of `CollectorPolicy.SystemData` has been changed from `*CollectorPolicySystemData` to `*TrackedResourceSystemData`
- Struct `CollectorPolicySystemData` has been removed

### Features Added

- New const `APIVersionParameterTwoThousandTwentyTwo0501`
- New const `APIVersionParameterTwoThousandTwentyTwo0801`
- New type alias `APIVersionParameter`
- New function `PossibleAPIVersionParameterValues() []APIVersionParameter`
- New field `Location` in struct `CollectorPolicy`
- New field `Tags` in struct `CollectorPolicy`


## 1.0.0 (2022-07-07)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).