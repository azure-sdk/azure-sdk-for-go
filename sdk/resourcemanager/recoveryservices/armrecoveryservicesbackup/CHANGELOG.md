# Release History

## 2.1.0 (2023-03-22)
### Features Added

- New enum type `TargetDiskNetworkAccessOption` with values `TargetDiskNetworkAccessOptionEnablePrivateAccessForAllDisks`, `TargetDiskNetworkAccessOptionEnablePublicAccessForAllDisks`, `TargetDiskNetworkAccessOptionSameAsOnSourceDisks`
- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewBMSPrepareDataMoveOperationResultClient() *BMSPrepareDataMoveOperationResultClient`
- New function `*ClientFactory.NewBackupEnginesClient() *BackupEnginesClient`
- New function `*ClientFactory.NewBackupJobsClient() *BackupJobsClient`
- New function `*ClientFactory.NewBackupOperationResultsClient() *BackupOperationResultsClient`
- New function `*ClientFactory.NewBackupOperationStatusesClient() *BackupOperationStatusesClient`
- New function `*ClientFactory.NewBackupPoliciesClient() *BackupPoliciesClient`
- New function `*ClientFactory.NewBackupProtectableItemsClient() *BackupProtectableItemsClient`
- New function `*ClientFactory.NewBackupProtectedItemsClient() *BackupProtectedItemsClient`
- New function `*ClientFactory.NewBackupProtectionContainersClient() *BackupProtectionContainersClient`
- New function `*ClientFactory.NewBackupProtectionIntentClient() *BackupProtectionIntentClient`
- New function `*ClientFactory.NewBackupResourceEncryptionConfigsClient() *BackupResourceEncryptionConfigsClient`
- New function `*ClientFactory.NewBackupResourceStorageConfigsNonCRRClient() *BackupResourceStorageConfigsNonCRRClient`
- New function `*ClientFactory.NewBackupResourceVaultConfigsClient() *BackupResourceVaultConfigsClient`
- New function `*ClientFactory.NewBackupStatusClient() *BackupStatusClient`
- New function `*ClientFactory.NewBackupUsageSummariesClient() *BackupUsageSummariesClient`
- New function `*ClientFactory.NewBackupWorkloadItemsClient() *BackupWorkloadItemsClient`
- New function `*ClientFactory.NewBackupsClient() *BackupsClient`
- New function `*ClientFactory.NewClient() *Client`
- New function `*ClientFactory.NewDeletedProtectionContainersClient() *DeletedProtectionContainersClient`
- New function `*ClientFactory.NewExportJobsOperationResultsClient() *ExportJobsOperationResultsClient`
- New function `*ClientFactory.NewFeatureSupportClient() *FeatureSupportClient`
- New function `*ClientFactory.NewItemLevelRecoveryConnectionsClient() *ItemLevelRecoveryConnectionsClient`
- New function `*ClientFactory.NewJobCancellationsClient() *JobCancellationsClient`
- New function `*ClientFactory.NewJobDetailsClient() *JobDetailsClient`
- New function `*ClientFactory.NewJobOperationResultsClient() *JobOperationResultsClient`
- New function `*ClientFactory.NewJobsClient() *JobsClient`
- New function `*ClientFactory.NewOperationClient() *OperationClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewPrivateEndpointClient() *PrivateEndpointClient`
- New function `*ClientFactory.NewPrivateEndpointConnectionClient() *PrivateEndpointConnectionClient`
- New function `*ClientFactory.NewProtectableContainersClient() *ProtectableContainersClient`
- New function `*ClientFactory.NewProtectedItemOperationResultsClient() *ProtectedItemOperationResultsClient`
- New function `*ClientFactory.NewProtectedItemOperationStatusesClient() *ProtectedItemOperationStatusesClient`
- New function `*ClientFactory.NewProtectedItemsClient() *ProtectedItemsClient`
- New function `*ClientFactory.NewProtectionContainerOperationResultsClient() *ProtectionContainerOperationResultsClient`
- New function `*ClientFactory.NewProtectionContainerRefreshOperationResultsClient() *ProtectionContainerRefreshOperationResultsClient`
- New function `*ClientFactory.NewProtectionContainersClient() *ProtectionContainersClient`
- New function `*ClientFactory.NewProtectionIntentClient() *ProtectionIntentClient`
- New function `*ClientFactory.NewProtectionPoliciesClient() *ProtectionPoliciesClient`
- New function `*ClientFactory.NewProtectionPolicyOperationResultsClient() *ProtectionPolicyOperationResultsClient`
- New function `*ClientFactory.NewProtectionPolicyOperationStatusesClient() *ProtectionPolicyOperationStatusesClient`
- New function `*ClientFactory.NewRecoveryPointsClient() *RecoveryPointsClient`
- New function `*ClientFactory.NewRecoveryPointsRecommendedForMoveClient() *RecoveryPointsRecommendedForMoveClient`
- New function `*ClientFactory.NewResourceGuardProxiesClient() *ResourceGuardProxiesClient`
- New function `*ClientFactory.NewResourceGuardProxyClient() *ResourceGuardProxyClient`
- New function `*ClientFactory.NewRestoresClient() *RestoresClient`
- New function `*ClientFactory.NewSecurityPINsClient() *SecurityPINsClient`
- New function `*ClientFactory.NewValidateOperationClient() *ValidateOperationClient`
- New function `*ClientFactory.NewValidateOperationResultsClient() *ValidateOperationResultsClient`
- New function `*ClientFactory.NewValidateOperationStatusesClient() *ValidateOperationStatusesClient`
- New struct `ClientFactory`
- New struct `ExtendedLocation`
- New struct `SecuredVMDetails`
- New struct `TargetDiskNetworkAccessSettings`
- New field `IncludeSoftDeletedRP` in struct `BMSRPQueryObject`
- New field `IsPrivateAccessEnabledOnAnyDisk` in struct `IaasVMRecoveryPoint`
- New field `SecurityType` in struct `IaasVMRecoveryPoint`
- New field `ExtendedLocation` in struct `IaasVMRestoreRequest`
- New field `SecuredVMDetails` in struct `IaasVMRestoreRequest`
- New field `TargetDiskNetworkAccessSettings` in struct `IaasVMRestoreRequest`
- New field `ExtendedLocation` in struct `IaasVMRestoreWithRehydrationRequest`
- New field `SecuredVMDetails` in struct `IaasVMRestoreWithRehydrationRequest`
- New field `TargetDiskNetworkAccessSettings` in struct `IaasVMRestoreWithRehydrationRequest`
- New field `IsSoftDeleted` in struct `RecoveryPointProperties`


## 2.0.0 (2023-01-19)
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
- Const `ContainerTypeAzureWorkloadContainer`, `ContainerTypeMicrosoftClassicComputeVirtualMachines`, `ContainerTypeMicrosoftComputeVirtualMachines` from type alias `ContainerType` has been removed

### Features Added

- New value `BackupItemTypeSAPHanaDBInstance` added to type alias `BackupItemType`
- New value `BackupTypeSnapshotCopyOnlyFull`, `BackupTypeSnapshotFull` added to type alias `BackupType`
- New value `ContainerTypeHanaHSRContainer` added to type alias `ContainerType`
- New value `DataSourceTypeSAPHanaDBInstance` added to type alias `DataSourceType`
- New value `PolicyTypeSnapshotCopyOnlyFull`, `PolicyTypeSnapshotFull` added to type alias `PolicyType`
- New value `ProtectedItemStateBackupsSuspended` added to type alias `ProtectedItemState`
- New value `ProtectionStateBackupsSuspended` added to type alias `ProtectionState`
- New value `RestorePointQueryTypeSnapshotCopyOnlyFull`, `RestorePointQueryTypeSnapshotFull` added to type alias `RestorePointQueryType`
- New value `RestorePointTypeSnapshotCopyOnlyFull`, `RestorePointTypeSnapshotFull` added to type alias `RestorePointType`
- New value `WorkloadItemTypeSAPHanaDBInstance` added to type alias `WorkloadItemType`
- New value `WorkloadTypeSAPHanaDBInstance` added to type alias `WorkloadType`
- New type alias `ProtectableContainerType` with values `ProtectableContainerTypeAzureBackupServerContainer`, `ProtectableContainerTypeAzureSQLContainer`, `ProtectableContainerTypeAzureWorkloadContainer`, `ProtectableContainerTypeCluster`, `ProtectableContainerTypeDPMContainer`, `ProtectableContainerTypeGenericContainer`, `ProtectableContainerTypeIaasVMContainer`, `ProtectableContainerTypeIaasVMServiceContainer`, `ProtectableContainerTypeInvalid`, `ProtectableContainerTypeMABContainer`, `ProtectableContainerTypeMicrosoftClassicComputeVirtualMachines`, `ProtectableContainerTypeMicrosoftComputeVirtualMachines`, `ProtectableContainerTypeSQLAGWorkLoadContainer`, `ProtectableContainerTypeStorageContainer`, `ProtectableContainerTypeUnknown`, `ProtectableContainerTypeVCenter`, `ProtectableContainerTypeVMAppContainer`, `ProtectableContainerTypeWindows`
- New type alias `TieringMode` with values `TieringModeDoNotTier`, `TieringModeInvalid`, `TieringModeTierAfter`, `TieringModeTierRecommended`
- New function `*AzureVMWorkloadSAPHanaDBInstance.GetAzureVMWorkloadProtectableItem() *AzureVMWorkloadProtectableItem`
- New function `*AzureVMWorkloadSAPHanaDBInstance.GetWorkloadProtectableItem() *WorkloadProtectableItem`
- New function `*AzureVMWorkloadSAPHanaDBInstanceProtectedItem.GetAzureVMWorkloadProtectedItem() *AzureVMWorkloadProtectedItem`
- New function `*AzureVMWorkloadSAPHanaDBInstanceProtectedItem.GetProtectedItem() *ProtectedItem`
- New function `*AzureVMWorkloadSAPHanaHSR.GetAzureVMWorkloadProtectableItem() *AzureVMWorkloadProtectableItem`
- New function `*AzureVMWorkloadSAPHanaHSR.GetWorkloadProtectableItem() *WorkloadProtectableItem`
- New function `NewDeletedProtectionContainersClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DeletedProtectionContainersClient, error)`
- New function `*DeletedProtectionContainersClient.NewListPager(string, string, *DeletedProtectionContainersClientListOptions) *runtime.Pager[DeletedProtectionContainersClientListResponse]`
- New struct `AzureVMWorkloadSAPHanaDBInstance`
- New struct `AzureVMWorkloadSAPHanaDBInstanceProtectedItem`
- New struct `AzureVMWorkloadSAPHanaHSR`
- New struct `DeletedProtectionContainersClient`
- New struct `DeletedProtectionContainersClientListResponse`
- New struct `RecoveryPointProperties`
- New struct `TieringPolicy`
- New field `RecoveryPointProperties` in struct `AzureFileShareRecoveryPoint`
- New field `SoftDeleteRetentionPeriod` in struct `AzureFileshareProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureIaaSClassicComputeVMProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureIaaSComputeVMProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureIaaSVMProtectedItem`
- New field `NewestRecoveryPointInArchive` in struct `AzureIaaSVMProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInArchive` in struct `AzureIaaSVMProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInVault` in struct `AzureIaaSVMProtectedItemExtendedInfo`
- New field `TieringPolicy` in struct `AzureIaaSVMProtectionPolicy`
- New field `SoftDeleteRetentionPeriod` in struct `AzureSQLProtectedItem`
- New field `NewestRecoveryPointInArchive` in struct `AzureVMWorkloadProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInArchive` in struct `AzureVMWorkloadProtectedItemExtendedInfo`
- New field `OldestRecoveryPointInVault` in struct `AzureVMWorkloadProtectedItemExtendedInfo`
- New field `SoftDeleteRetentionPeriod` in struct `AzureVMWorkloadSAPAseDatabaseProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureVMWorkloadSAPHanaDatabaseProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `AzureVMWorkloadSQLDatabaseProtectedItem`
- New field `RecoveryPointProperties` in struct `AzureWorkloadPointInTimeRecoveryPoint`
- New field `RecoveryPointProperties` in struct `AzureWorkloadRecoveryPoint`
- New field `RecoveryPointProperties` in struct `AzureWorkloadSAPHanaPointInTimeRecoveryPoint`
- New field `RecoveryPointProperties` in struct `AzureWorkloadSAPHanaRecoveryPoint`
- New field `RecoveryPointProperties` in struct `AzureWorkloadSQLPointInTimeRecoveryPoint`
- New field `RecoveryPointProperties` in struct `AzureWorkloadSQLRecoveryPoint`
- New field `SoftDeleteRetentionPeriod` in struct `DPMProtectedItem`
- New field `SoftDeleteRetentionPeriod` in struct `GenericProtectedItem`
- New field `RecoveryPointProperties` in struct `GenericRecoveryPoint`
- New field `RecoveryPointProperties` in struct `IaasVMRecoveryPoint`
- New field `SoftDeleteRetentionPeriod` in struct `MabFileFolderProtectedItem`
- New field `TieringPolicy` in struct `SubProtectionPolicy`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
