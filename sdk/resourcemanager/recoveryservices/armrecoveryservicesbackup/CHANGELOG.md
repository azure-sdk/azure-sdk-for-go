# Release History

## 2.0.0-beta.1 (2022-11-23)
### Breaking Changes

- Type of `AzureBackupServerContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureIaaSClassicComputeVMContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureIaaSComputeVMContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureSQLAGWorkloadContainerProtectionContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureSQLContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureStorageContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureStorageProtectableContainer.ProtectableContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureVMAppContainerProtectableContainer.ProtectableContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureVMAppContainerProtectionContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureWorkloadContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `DpmContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `GenericContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `IaaSVMContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `MabContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `ProtectableContainer.ProtectableContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `ProtectionContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Const `ContainerTypeAzureWorkloadContainer` has been removed
- Const `ContainerTypeMicrosoftClassicComputeVirtualMachines` has been removed
- Const `ContainerTypeMicrosoftComputeVirtualMachines` has been removed
- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New const `BackupItemTypeSAPHanaDBInstance`
- New const `BackupTypeSnapshotCopyOnlyFull`
- New const `BackupTypeSnapshotFull`
- New const `ContainerTypeHanaHSRContainer`
- New const `DataSourceTypeSAPHanaDBInstance`
- New const `PolicyTypeSnapshotCopyOnlyFull`
- New const `PolicyTypeSnapshotFull`
- New const `ProtectableContainerTypeAzureBackupServerContainer`
- New const `ProtectableContainerTypeAzureSQLContainer`
- New const `ProtectableContainerTypeAzureWorkloadContainer`
- New const `ProtectableContainerTypeCluster`
- New const `ProtectableContainerTypeDPMContainer`
- New const `ProtectableContainerTypeGenericContainer`
- New const `ProtectableContainerTypeIaasVMContainer`
- New const `ProtectableContainerTypeIaasVMServiceContainer`
- New const `ProtectableContainerTypeInvalid`
- New const `ProtectableContainerTypeMABContainer`
- New const `ProtectableContainerTypeMicrosoftClassicComputeVirtualMachines`
- New const `ProtectableContainerTypeMicrosoftComputeVirtualMachines`
- New const `ProtectableContainerTypeSQLAGWorkLoadContainer`
- New const `ProtectableContainerTypeStorageContainer`
- New const `ProtectableContainerTypeUnknown`
- New const `ProtectableContainerTypeVCenter`
- New const `ProtectableContainerTypeVMAppContainer`
- New const `ProtectableContainerTypeWindows`
- New const `RestorePointQueryTypeSnapshotCopyOnlyFull`
- New const `RestorePointQueryTypeSnapshotFull`
- New const `RestorePointTypeSnapshotCopyOnlyFull`
- New const `RestorePointTypeSnapshotFull`
- New const `TieringModeDoNotTier`
- New const `TieringModeInvalid`
- New const `TieringModeTierAfter`
- New const `TieringModeTierRecommended`
- New const `WorkloadItemTypeSAPHanaDBInstance`
- New const `WorkloadTypeSAPHanaDBInstance`
- New type alias `ProtectableContainerType`
- New type alias `TieringMode`
- New function `*AzureVMWorkloadSAPHanaDBInstance.GetAzureVMWorkloadProtectableItem() *AzureVMWorkloadProtectableItem`
- New function `*AzureVMWorkloadSAPHanaDBInstance.GetWorkloadProtectableItem() *WorkloadProtectableItem`
- New function `*AzureVMWorkloadSAPHanaDBInstanceProtectedItem.GetAzureVMWorkloadProtectedItem() *AzureVMWorkloadProtectedItem`
- New function `*AzureVMWorkloadSAPHanaDBInstanceProtectedItem.GetProtectedItem() *ProtectedItem`
- New function `*AzureVMWorkloadSAPHanaHSR.GetAzureVMWorkloadProtectableItem() *AzureVMWorkloadProtectableItem`
- New function `*AzureVMWorkloadSAPHanaHSR.GetWorkloadProtectableItem() *WorkloadProtectableItem`
- New function `NewDeletedProtectionContainersClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DeletedProtectionContainersClient, error)`
- New function `*DeletedProtectionContainersClient.NewListPager(string, string, *DeletedProtectionContainersClientListOptions) *runtime.Pager[DeletedProtectionContainersClientListResponse]`
- New function `PossibleProtectableContainerTypeValues() []ProtectableContainerType`
- New function `PossibleTieringModeValues() []TieringMode`
- New struct `AzureVMWorkloadSAPHanaDBInstance`
- New struct `AzureVMWorkloadSAPHanaDBInstanceProtectedItem`
- New struct `AzureVMWorkloadSAPHanaHSR`
- New struct `DeletedProtectionContainersClient`
- New struct `DeletedProtectionContainersClientListOptions`
- New struct `DeletedProtectionContainersClientListResponse`
- New struct `TieringPolicy`
- New field `SoftDeleteRetentionPeriod` in struct `AzureFileshareProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureIaaSClassicComputeVMProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureIaaSComputeVMProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureIaaSVMProtectedItem`
- New field `NewestRecoveryPointInArchive` in struct `AzureIaaSVMProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInArchive` in struct `AzureIaaSVMProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInVault` in struct `AzureIaaSVMProtectedItemExtendedInfo`
- New field `TieringPolicy` in struct `AzureIaaSVMProtectionPolicy`
- New field `SoftDeleteRetentionPeriod` in struct `AzureSQLProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureVMWorkloadProtectedItem`
- New field `NewestRecoveryPointInArchive` in struct `AzureVMWorkloadProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInArchive` in struct `AzureVMWorkloadProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInVault` in struct `AzureVMWorkloadProtectedItemExtendedInfo`
- New field `SoftDeleteRetentionPeriod` in struct `AzureVMWorkloadSAPAseDatabaseProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureVMWorkloadSAPHanaDatabaseProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureVMWorkloadSQLDatabaseProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `DPMProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `GenericProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `MabFileFolderProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `ProtectedItem`
- New field `TieringPolicy` in struct `SubProtectionPolicy`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).