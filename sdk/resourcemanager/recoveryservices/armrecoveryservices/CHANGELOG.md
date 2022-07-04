# Release History

## 1.1.0 (2022-07-04)
### Features Added

- New const `StandardTierStorageRedundancyZoneRedundant`
- New const `StandardTierStorageRedundancyLocallyRedundant`
- New const `StandardTierStorageRedundancyGeoRedundant`
- New const `CrossRegionRestoreEnabled`
- New const `CrossRegionRestoreDisabled`
- New function `PossibleCrossRegionRestoreValues() []CrossRegionRestore`
- New function `PossibleStandardTierStorageRedundancyValues() []StandardTierStorageRedundancy`
- New struct `VaultPropertiesRedundancySettings`
- New field `AADAudience` in struct `ResourceCertificateAndAADDetails`
- New field `RedundancySettings` in struct `VaultProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).