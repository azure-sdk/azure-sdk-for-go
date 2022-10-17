# Release History

## 2.0.0-beta.1 (2022-10-17)
### Breaking Changes

- Type of `AzureSQLContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureIaaSClassicComputeVMContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureStorageContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureStorageProtectableContainer.ProtectableContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `DpmContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureWorkloadContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureVMAppContainerProtectionContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `ProtectionContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureSQLAGWorkloadContainerProtectionContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureVMAppContainerProtectableContainer.ProtectableContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureBackupServerContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `ProtectableContainer.ProtectableContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `IaaSVMContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `MabContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `GenericContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureIaaSComputeVMContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Const `ContainerTypeMicrosoftComputeVirtualMachines` has been removed
- Const `ContainerTypeMicrosoftClassicComputeVirtualMachines` has been removed
- Const `ContainerTypeAzureWorkloadContainer` has been removed
- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New const `ProtectableContainerTypeVCenter`
- New const `ProtectableContainerTypeAzureSQLContainer`
- New const `ProtectableContainerTypeMicrosoftClassicComputeVirtualMachines`
- New const `RestorePointQueryTypeSnapshotCopyOnlyFull`
- New const `BackupItemTypeSAPHanaDBInstance`
- New const `ContainerTypeHanaHSRContainer`
- New const `PolicyTypeSnapshotCopyOnlyFull`
- New const `DataSourceTypeSAPHanaDBInstance`
- New const `ProtectableContainerTypeIaasVMContainer`
- New const `BackupTypeSnapshotCopyOnlyFull`
- New const `RestorePointTypeSnapshotCopyOnlyFull`
- New const `WorkloadTypeSAPHanaDBInstance`
- New const `BackupTypeSnapshotFull`
- New const `RestorePointQueryTypeSnapshotFull`
- New const `ProtectableContainerTypeMicrosoftComputeVirtualMachines`
- New const `WorkloadItemTypeSAPHanaDBInstance`
- New const `TieringModeTierAfter`
- New const `ProtectableContainerTypeUnknown`
- New const `TieringModeTierRecommended`
- New const `ProtectableContainerTypeAzureBackupServerContainer`
- New const `ProtectableContainerTypeDPMContainer`
- New const `ProtectableContainerTypeWindows`
- New const `ProtectableContainerTypeVMAppContainer`
- New const `RestorePointTypeSnapshotFull`
- New const `TieringModeDoNotTier`
- New const `ProtectableContainerTypeStorageContainer`
- New const `ProtectableContainerTypeIaasVMServiceContainer`
- New const `ProtectableContainerTypeInvalid`
- New const `ProtectableContainerTypeGenericContainer`
- New const `TieringModeInvalid`
- New const `PolicyTypeSnapshotFull`
- New const `ProtectableContainerTypeSQLAGWorkLoadContainer`
- New const `ProtectableContainerTypeMABContainer`
- New const `ProtectableContainerTypeAzureWorkloadContainer`
- New const `ProtectableContainerTypeCluster`
- New type alias `ProtectableContainerType`
- New type alias `TieringMode`
- New function `*AzureVMWorkloadSAPHanaDBInstanceProtectedItem.GetProtectedItem() *ProtectedItem`
- New function `*AzureVMWorkloadSAPHanaDBInstanceProtectedItem.GetAzureVMWorkloadProtectedItem() *AzureVMWorkloadProtectedItem`
- New function `PossibleProtectableContainerTypeValues() []ProtectableContainerType`
- New function `*DeletedProtectionContainersClient.NewListPager(string, string, *DeletedProtectionContainersClientListOptions) *runtime.Pager[DeletedProtectionContainersClientListResponse]`
- New function `*AzureVMWorkloadSAPHanaDBInstance.GetWorkloadProtectableItem() *WorkloadProtectableItem`
- New function `NewDeletedProtectionContainersClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DeletedProtectionContainersClient, error)`
- New function `*AzureVMWorkloadSAPHanaHSR.GetWorkloadProtectableItem() *WorkloadProtectableItem`
- New function `*AzureVMWorkloadSAPHanaDBInstance.GetAzureVMWorkloadProtectableItem() *AzureVMWorkloadProtectableItem`
- New function `PossibleTieringModeValues() []TieringMode`
- New function `*AzureVMWorkloadSAPHanaHSR.GetAzureVMWorkloadProtectableItem() *AzureVMWorkloadProtectableItem`
- New struct `AzureVMWorkloadSAPHanaDBInstance`
- New struct `AzureVMWorkloadSAPHanaDBInstanceProtectedItem`
- New struct `AzureVMWorkloadSAPHanaHSR`
- New struct `DeletedProtectionContainersClient`
- New struct `DeletedProtectionContainersClientListOptions`
- New struct `DeletedProtectionContainersClientListResponse`
- New struct `TieringPolicy`
- New field `SoftDeleteRetentionPeriod` in struct `ProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureVMWorkloadSAPHanaDatabaseProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureVMWorkloadSQLDatabaseProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureIaaSComputeVMProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureVMWorkloadSAPAseDatabaseProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureIaaSClassicComputeVMProtectedItem`
- New field `NewestRecoveryPointInArchive` in struct `AzureVMWorkloadProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInArchive` in struct `AzureVMWorkloadProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInVault` in struct `AzureVMWorkloadProtectedItemExtendedInfo`
- New field `SoftDeleteRetentionPeriod` in struct `DPMProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `GenericProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `MabFileFolderProtectedItem`
- New field `NewestRecoveryPointInArchive` in struct `AzureIaaSVMProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInArchive` in struct `AzureIaaSVMProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInVault` in struct `AzureIaaSVMProtectedItemExtendedInfo`
- New field `SoftDeleteRetentionPeriod` in struct `AzureIaaSVMProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureFileshareProtectedItem`
- New field `TieringPolicy` in struct `SubProtectionPolicy`
- New field `SoftDeleteRetentionPeriod` in struct `AzureSQLProtectedItem`
- New field `TieringPolicy` in struct `AzureIaaSVMProtectionPolicy`
- New field `SoftDeleteRetentionPeriod` in struct `AzureVMWorkloadProtectedItem`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).