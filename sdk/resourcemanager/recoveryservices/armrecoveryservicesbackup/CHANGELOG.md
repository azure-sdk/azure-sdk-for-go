# Release History

## 2.0.0-beta.1 (2022-08-23)
### Breaking Changes

- Type of `AzureStorageProtectableContainer.ProtectableContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureWorkloadContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `DpmContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureVMAppContainerProtectableContainer.ProtectableContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureSQLAGWorkloadContainerProtectionContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureIaaSClassicComputeVMContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `MabContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `IaaSVMContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `ProtectableContainer.ProtectableContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureIaaSComputeVMContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `ProtectionContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureSQLContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `GenericContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureVMAppContainerProtectionContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureStorageContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureBackupServerContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Const `ContainerTypeMicrosoftClassicComputeVirtualMachines` has been removed
- Const `ContainerTypeAzureWorkloadContainer` has been removed
- Const `ContainerTypeMicrosoftComputeVirtualMachines` has been removed

### Features Added

- New const `TieringModeDoNotTier`
- New const `RestorePointTypeSnapshotCopyOnlyFull`
- New const `BackupTypeSnapshotCopyOnlyFull`
- New const `ProtectableContainerTypeIaasVMContainer`
- New const `ProtectableContainerTypeDPMContainer`
- New const `RestorePointTypeSnapshotFull`
- New const `ProtectableContainerTypeInvalid`
- New const `ProtectableContainerTypeCluster`
- New const `BackupTypeSnapshotFull`
- New const `DataSourceTypeSAPHanaDBInstance`
- New const `ProtectableContainerTypeVCenter`
- New const `ProtectableContainerTypeMicrosoftComputeVirtualMachines`
- New const `PolicyTypeSnapshotCopyOnlyFull`
- New const `ProtectableContainerTypeMicrosoftClassicComputeVirtualMachines`
- New const `ProtectableContainerTypeSQLAGWorkLoadContainer`
- New const `ProtectableContainerTypeIaasVMServiceContainer`
- New const `PolicyTypeSnapshotFull`
- New const `ProtectableContainerTypeWindows`
- New const `ProtectableContainerTypeAzureWorkloadContainer`
- New const `TieringModeTierAfter`
- New const `TieringModeTierRecommended`
- New const `RestorePointQueryTypeSnapshotCopyOnlyFull`
- New const `ProtectableContainerTypeStorageContainer`
- New const `BackupItemTypeSAPHanaDBInstance`
- New const `ProtectableContainerTypeMABContainer`
- New const `ProtectableContainerTypeGenericContainer`
- New const `ProtectableContainerTypeAzureBackupServerContainer`
- New const `ContainerTypeHanaHSRContainer`
- New const `ProtectableContainerTypeUnknown`
- New const `ProtectableContainerTypeAzureSQLContainer`
- New const `WorkloadItemTypeSAPHanaDBInstance`
- New const `TieringModeInvalid`
- New const `RestorePointQueryTypeSnapshotFull`
- New const `ProtectableContainerTypeVMAppContainer`
- New const `WorkloadTypeSAPHanaDBInstance`
- New type alias `TieringMode`
- New type alias `ProtectableContainerType`
- New function `*AzureVMWorkloadSAPHanaHSR.GetAzureVMWorkloadProtectableItem() *AzureVMWorkloadProtectableItem`
- New function `PossibleTieringModeValues() []TieringMode`
- New function `*AzureVMWorkloadSAPHanaDBInstance.GetAzureVMWorkloadProtectableItem() *AzureVMWorkloadProtectableItem`
- New function `*AzureVMWorkloadSAPHanaHSR.GetWorkloadProtectableItem() *WorkloadProtectableItem`
- New function `*AzureVMWorkloadSAPHanaDBInstance.GetWorkloadProtectableItem() *WorkloadProtectableItem`
- New function `*AzureVMWorkloadSAPHanaDBInstanceProtectedItem.GetAzureVMWorkloadProtectedItem() *AzureVMWorkloadProtectedItem`
- New function `*AzureVMWorkloadSAPHanaDBInstanceProtectedItem.GetProtectedItem() *ProtectedItem`
- New function `PossibleProtectableContainerTypeValues() []ProtectableContainerType`
- New struct `AzureVMWorkloadSAPHanaDBInstance`
- New struct `AzureVMWorkloadSAPHanaDBInstanceProtectedItem`
- New struct `AzureVMWorkloadSAPHanaHSR`
- New struct `TieringPolicy`
- New field `TieringPolicy` in struct `SubProtectionPolicy`
- New field `OldestRecoveryPointInVault` in struct `AzureVMWorkloadProtectedItemExtendedInfo`
- New field `NewestRecoveryPointInArchive` in struct `AzureVMWorkloadProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInArchive` in struct `AzureVMWorkloadProtectedItemExtendedInfo`
- New field `TieringPolicy` in struct `AzureIaaSVMProtectionPolicy`
- New field `NewestRecoveryPointInArchive` in struct `AzureIaaSVMProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInArchive` in struct `AzureIaaSVMProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInVault` in struct `AzureIaaSVMProtectedItemExtendedInfo`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).