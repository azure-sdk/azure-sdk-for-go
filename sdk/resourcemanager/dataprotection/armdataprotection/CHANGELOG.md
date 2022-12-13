# Release History

## 2.0.0 (2022-12-13)
### Breaking Changes

- Const `StorageSettingStoreTypesSnapshotStore` has been removed
- Struct `CloudError` has been removed

### Features Added

- New value `StorageSettingStoreTypesOperationalStore` added to type alias `StorageSettingStoreTypes`
- New value `StorageSettingTypesZoneRedundant` added to type alias `StorageSettingTypes`
- New struct `DppProxyResource`
- New field `Tags` in struct `BackupInstanceResource`
- New field `IsVaultProtectedByResourceGuard` in struct `BackupVault`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).