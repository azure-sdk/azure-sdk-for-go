# Release History

## 2.0.0-beta.1 (2022-10-18)
### Breaking Changes

- Struct `CloudError` has been removed

### Features Added

- New const `PersistentVolumeRestoreModeRestoreWithVolumeData`
- New const `StorageSettingTypesZoneRedundant`
- New const `ExistingResourcePolicySkip`
- New const `ExistingResourcePolicyPatch`
- New const `PersistentVolumeRestoreModeRestoreWithoutVolumeData`
- New type alias `PersistentVolumeRestoreMode`
- New type alias `ExistingResourcePolicy`
- New function `PossiblePersistentVolumeRestoreModeValues() []PersistentVolumeRestoreMode`
- New function `*KubernetesClusterBackupDatasourceParameters.GetBackupDatasourceParameters() *BackupDatasourceParameters`
- New function `*BackupDatasourceParameters.GetBackupDatasourceParameters() *BackupDatasourceParameters`
- New function `PossibleExistingResourcePolicyValues() []ExistingResourcePolicy`
- New function `*KubernetesClusterRestoreCriteria.GetItemLevelRestoreCriteria() *ItemLevelRestoreCriteria`
- New struct `BackupDatasourceParameters`
- New struct `DppProxyResource`
- New struct `KubernetesClusterBackupDatasourceParameters`
- New struct `KubernetesClusterRestoreCriteria`
- New field `Tags` in struct `BackupInstanceResource`
- New field `BackupDatasourceParametersList` in struct `PolicyParameters`
- New field `IsVaultProtectedByResourceGuard` in struct `BackupVault`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).