# Release History

## 1.1.0 (2022-07-02)
### Features Added

- New const `CrossRegionRestoreEnabled`
- New const `StandardTierStorageRedundancyLocallyRedundant`
- New const `StandardTierStorageRedundancyZoneRedundant`
- New const `StandardTierStorageRedundancyGeoRedundant`
- New const `ImmutabilityStateUnlocked`
- New const `ImmutabilityStateLocked`
- New const `ImmutabilityStateDisabled`
- New const `CrossRegionRestoreDisabled`
- New function `PossibleCrossRegionRestoreValues() []CrossRegionRestore`
- New function `PossibleImmutabilityStateValues() []ImmutabilityState`
- New function `PossibleStandardTierStorageRedundancyValues() []StandardTierStorageRedundancy`
- New struct `ImmutabilitySettings`
- New struct `SecuritySettings`
- New struct `VaultPropertiesRedundancySettings`
- New field `RedundancySettings` in struct `VaultProperties`
- New field `SecuritySettings` in struct `VaultProperties`
- New field `AADAudience` in struct `ResourceCertificateAndAADDetails`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).