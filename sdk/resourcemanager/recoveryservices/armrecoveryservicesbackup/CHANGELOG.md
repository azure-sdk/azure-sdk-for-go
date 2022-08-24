# Release History

## 2.0.0-beta.1 (2022-08-24)
### Breaking Changes

- Type of `AzureSQLContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `GenericContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureStorageProtectableContainer.ProtectableContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureIaaSClassicComputeVMContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureWorkloadContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `DpmContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureIaaSComputeVMContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureSQLAGWorkloadContainerProtectionContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `ProtectionContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureVMAppContainerProtectableContainer.ProtectableContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `MabContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureVMAppContainerProtectionContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureBackupServerContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `ProtectableContainer.ProtectableContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `AzureStorageContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Type of `IaaSVMContainer.ContainerType` has been changed from `*ContainerType` to `*ProtectableContainerType`
- Const `ContainerTypeMicrosoftComputeVirtualMachines` has been removed
- Const `ContainerTypeMicrosoftClassicComputeVirtualMachines` has been removed
- Const `ContainerTypeAzureWorkloadContainer` has been removed

### Features Added

- New const `TieringModeTierAfter`
- New const `ProtectableContainerTypeAzureSQLContainer`
- New const `ProtectableContainerTypeUnknown`
- New const `WorkloadItemTypeSAPHanaDBInstance`
- New const `ProtectableContainerTypeMicrosoftComputeVirtualMachines`
- New const `ProtectableContainerTypeMABContainer`
- New const `WorkloadTypeSAPHanaDBInstance`
- New const `BackupItemTypeSAPHanaDBInstance`
- New const `ContainerTypeHanaHSRContainer`
- New const `ProtectableContainerTypeSQLAGWorkLoadContainer`
- New const `PolicyTypeSnapshotFull`
- New const `ProtectableContainerTypeAzureBackupServerContainer`
- New const `ProtectableContainerTypeInvalid`
- New const `BackupTypeSnapshotFull`
- New const `ProtectableContainerTypeStorageContainer`
- New const `ProtectableContainerTypeCluster`
- New const `ProtectableContainerTypeDPMContainer`
- New const `RestorePointQueryTypeSnapshotFull`
- New const `ProtectableContainerTypeAzureWorkloadContainer`
- New const `ProtectableContainerTypeGenericContainer`
- New const `RestorePointTypeSnapshotFull`
- New const `TieringModeTierRecommended`
- New const `ProtectableContainerTypeVCenter`
- New const `RestorePointQueryTypeSnapshotCopyOnlyFull`
- New const `ProtectableContainerTypeIaasVMContainer`
- New const `ProtectableContainerTypeWindows`
- New const `ProtectableContainerTypeVMAppContainer`
- New const `ProtectableContainerTypeMicrosoftClassicComputeVirtualMachines`
- New const `ProtectableContainerTypeIaasVMServiceContainer`
- New const `TieringModeInvalid`
- New const `PolicyTypeSnapshotCopyOnlyFull`
- New const `RestorePointTypeSnapshotCopyOnlyFull`
- New const `BackupTypeSnapshotCopyOnlyFull`
- New const `DataSourceTypeSAPHanaDBInstance`
- New const `TieringModeDoNotTier`
- New type alias `ProtectableContainerType`
- New type alias `TieringMode`
- New function `*AzureVMWorkloadSAPHanaDBInstanceProtectedItem.GetAzureVMWorkloadProtectedItem() *AzureVMWorkloadProtectedItem`
- New function `PossibleTieringModeValues() []TieringMode`
- New function `PossibleProtectableContainerTypeValues() []ProtectableContainerType`
- New function `*AzureVMWorkloadSAPHanaHSR.GetWorkloadProtectableItem() *WorkloadProtectableItem`
- New function `*AzureVMWorkloadSAPHanaDBInstance.GetWorkloadProtectableItem() *WorkloadProtectableItem`
- New function `*AzureVMWorkloadSAPHanaHSR.GetAzureVMWorkloadProtectableItem() *AzureVMWorkloadProtectableItem`
- New function `*AzureVMWorkloadSAPHanaDBInstanceProtectedItem.GetProtectedItem() *ProtectedItem`
- New function `*AzureVMWorkloadSAPHanaDBInstance.GetAzureVMWorkloadProtectableItem() *AzureVMWorkloadProtectableItem`
- New struct `AzureVMWorkloadSAPHanaDBInstance`
- New struct `AzureVMWorkloadSAPHanaDBInstanceProtectedItem`
- New struct `AzureVMWorkloadSAPHanaHSR`
- New struct `TieringPolicy`
- New field `TieringPolicy` in struct `AzureIaaSVMProtectionPolicy`
- New field `NewestRecoveryPointInArchive` in struct `AzureIaaSVMProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInArchive` in struct `AzureIaaSVMProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInVault` in struct `AzureIaaSVMProtectedItemExtendedInfo`
- New field `TieringPolicy` in struct `SubProtectionPolicy`
- New field `NewestRecoveryPointInArchive` in struct `AzureVMWorkloadProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInArchive` in struct `AzureVMWorkloadProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInVault` in struct `AzureVMWorkloadProtectedItemExtendedInfo`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).