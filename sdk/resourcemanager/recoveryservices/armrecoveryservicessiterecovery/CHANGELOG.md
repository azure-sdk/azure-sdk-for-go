# Release History

## 1.2.0 (2023-03-21)
### Features Added

- New enum type `SecurityType` with values `SecurityTypeConfidentialVM`, `SecurityTypeNone`, `SecurityTypeTrustedLaunch`
- New function `NewClientFactory(string, string, string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewMigrationRecoveryPointsClient() *MigrationRecoveryPointsClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewRecoveryPointsClient() *RecoveryPointsClient`
- New function `*ClientFactory.NewReplicationAlertSettingsClient() *ReplicationAlertSettingsClient`
- New function `*ClientFactory.NewReplicationAppliancesClient() *ReplicationAppliancesClient`
- New function `*ClientFactory.NewReplicationEligibilityResultsClient() *ReplicationEligibilityResultsClient`
- New function `*ClientFactory.NewReplicationEventsClient() *ReplicationEventsClient`
- New function `*ClientFactory.NewReplicationFabricsClient() *ReplicationFabricsClient`
- New function `*ClientFactory.NewReplicationJobsClient() *ReplicationJobsClient`
- New function `*ClientFactory.NewReplicationLogicalNetworksClient() *ReplicationLogicalNetworksClient`
- New function `*ClientFactory.NewReplicationMigrationItemsClient() *ReplicationMigrationItemsClient`
- New function `*ClientFactory.NewReplicationNetworkMappingsClient() *ReplicationNetworkMappingsClient`
- New function `*ClientFactory.NewReplicationNetworksClient() *ReplicationNetworksClient`
- New function `*ClientFactory.NewReplicationPoliciesClient() *ReplicationPoliciesClient`
- New function `*ClientFactory.NewReplicationProtectableItemsClient() *ReplicationProtectableItemsClient`
- New function `*ClientFactory.NewReplicationProtectedItemsClient() *ReplicationProtectedItemsClient`
- New function `*ClientFactory.NewReplicationProtectionContainerMappingsClient() *ReplicationProtectionContainerMappingsClient`
- New function `*ClientFactory.NewReplicationProtectionContainersClient() *ReplicationProtectionContainersClient`
- New function `*ClientFactory.NewReplicationProtectionIntentsClient() *ReplicationProtectionIntentsClient`
- New function `*ClientFactory.NewReplicationRecoveryPlansClient() *ReplicationRecoveryPlansClient`
- New function `*ClientFactory.NewReplicationRecoveryServicesProvidersClient() *ReplicationRecoveryServicesProvidersClient`
- New function `*ClientFactory.NewReplicationStorageClassificationMappingsClient() *ReplicationStorageClassificationMappingsClient`
- New function `*ClientFactory.NewReplicationStorageClassificationsClient() *ReplicationStorageClassificationsClient`
- New function `*ClientFactory.NewReplicationVaultHealthClient() *ReplicationVaultHealthClient`
- New function `*ClientFactory.NewReplicationVaultSettingClient() *ReplicationVaultSettingClient`
- New function `*ClientFactory.NewReplicationvCentersClient() *ReplicationvCentersClient`
- New function `*ClientFactory.NewSupportedOperatingSystemsClient() *SupportedOperatingSystemsClient`
- New function `*ClientFactory.NewTargetComputeSizesClient() *TargetComputeSizesClient`
- New struct `A2AFabricSpecificLocationDetails`
- New struct `ClientFactory`
- New struct `VMwareCbtSecurityProfileProperties`
- New field `LocationDetails` in struct `AzureFabricSpecificDetails`
- New field `ExtendedLocationMappings` in struct `FabricQueryParameter`
- New field `LocationDetails` in struct `FabricQueryParameter`
- New field `ConfidentialVMKeyVaultID` in struct `VMwareCbtEnableMigrationInput`
- New field `TargetVMSecurityProfile` in struct `VMwareCbtEnableMigrationInput`
- New field `OSUpgradeVersion` in struct `VMwareCbtMigrateInput`
- New field `ConfidentialVMKeyVaultID` in struct `VMwareCbtMigrationDetails`
- New field `OSName` in struct `VMwareCbtMigrationDetails`
- New field `SupportedOSVersions` in struct `VMwareCbtMigrationDetails`
- New field `TargetVMSecurityProfile` in struct `VMwareCbtMigrationDetails`
- New field `ExcludedSKUs` in struct `VMwareCbtProtectionContainerMappingDetails`
- New field `OSUpgradeVersion` in struct `VMwareCbtTestMigrateInput`


## 1.1.0 (2022-12-23)
### Features Added

- New value `MigrationItemOperationPauseReplication`, `MigrationItemOperationResumeReplication` added to type alias `MigrationItemOperation`
- New value `MigrationStateMigrationCompletedWithInformation`, `MigrationStateMigrationPartiallySucceeded`, `MigrationStateProtectionSuspended`, `MigrationStateResumeInProgress`, `MigrationStateResumeInitiated`, `MigrationStateSuspendingProtection` added to type alias `MigrationState`
- New value `TestMigrationStateTestMigrationCompletedWithInformation`, `TestMigrationStateTestMigrationPartiallySucceeded` added to type alias `TestMigrationState`
- New function `*ReplicationMigrationItemsClient.BeginPauseReplication(context.Context, string, string, string, PauseReplicationInput, *ReplicationMigrationItemsClientBeginPauseReplicationOptions) (*runtime.Poller[ReplicationMigrationItemsClientPauseReplicationResponse], error)`
- New function `*ReplicationMigrationItemsClient.BeginResumeReplication(context.Context, string, string, string, ResumeReplicationInput, *ReplicationMigrationItemsClientBeginResumeReplicationOptions) (*runtime.Poller[ReplicationMigrationItemsClientResumeReplicationResponse], error)`
- New function `*ResumeReplicationProviderSpecificInput.GetResumeReplicationProviderSpecificInput() *ResumeReplicationProviderSpecificInput`
- New function `*VMwareCbtResumeReplicationInput.GetResumeReplicationProviderSpecificInput() *ResumeReplicationProviderSpecificInput`
- New struct `A2AExtendedLocationDetails`
- New struct `CriticalJobHistoryDetails`
- New struct `PauseReplicationInput`
- New struct `PauseReplicationInputProperties`
- New struct `ReplicationMigrationItemsClientPauseReplicationResponse`
- New struct `ReplicationMigrationItemsClientResumeReplicationResponse`
- New struct `ResumeReplicationInput`
- New struct `ResumeReplicationInputProperties`
- New struct `VMwareCbtResumeReplicationInput`
- New field `ExtendedLocations` in struct `AzureFabricSpecificDetails`
- New field `SeedBlobURI` in struct `InMageRcmProtectedDiskDetails`
- New field `StorageAccountID` in struct `InMageRcmReplicationDetails`
- New field `CriticalJobHistory` in struct `MigrationItemProperties`
- New field `LastMigrationStatus` in struct `MigrationItemProperties`
- New field `LastMigrationTime` in struct `MigrationItemProperties`
- New field `RecoveryServicesProviderID` in struct `MigrationItemProperties`
- New field `ReplicationStatus` in struct `MigrationItemProperties`
- New field `PrimaryExtendedLocation` in struct `RecoveryPlanA2ADetails`
- New field `RecoveryExtendedLocation` in struct `RecoveryPlanA2ADetails`
- New field `PerformSQLBulkRegistration` in struct `VMwareCbtEnableMigrationInput`
- New field `ResumeProgressPercentage` in struct `VMwareCbtMigrationDetails`
- New field `ResumeRetryCount` in struct `VMwareCbtMigrationDetails`
- New field `StorageAccountID` in struct `VMwareCbtMigrationDetails`
- New field `TestNetworkID` in struct `VMwareCbtMigrationDetails`
- New field `SeedBlobURI` in struct `VMwareCbtProtectedDiskDetails`
- New field `TargetBlobURI` in struct `VMwareCbtProtectedDiskDetails`
- New field `RoleSizeToNicCountMap` in struct `VMwareCbtProtectionContainerMappingDetails`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).