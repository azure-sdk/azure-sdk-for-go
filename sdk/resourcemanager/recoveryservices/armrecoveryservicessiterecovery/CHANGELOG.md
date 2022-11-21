# Release History

## 1.1.0 (2022-11-21)
### Features Added

- New const `MigrationStateResumeInProgress`
- New const `MigrationStateMigrationCompletedWithInformation`
- New const `TestMigrationStateTestMigrationCompletedWithInformation`
- New const `MigrationItemOperationPauseReplication`
- New const `MigrationStateMigrationPartiallySucceeded`
- New const `MigrationStateSuspendingProtection`
- New const `MigrationItemOperationResumeReplication`
- New const `TestMigrationStateTestMigrationPartiallySucceeded`
- New const `MigrationStateProtectionSuspended`
- New const `MigrationStateResumeInitiated`
- New function `*ReplicationMigrationItemsClient.BeginResumeReplication(context.Context, string, string, string, ResumeReplicationInput, *ReplicationMigrationItemsClientBeginResumeReplicationOptions) (*runtime.Poller[ReplicationMigrationItemsClientResumeReplicationResponse], error)`
- New function `*VMwareCbtResumeReplicationInput.GetResumeReplicationProviderSpecificInput() *ResumeReplicationProviderSpecificInput`
- New function `*ReplicationMigrationItemsClient.BeginPauseReplication(context.Context, string, string, string, PauseReplicationInput, *ReplicationMigrationItemsClientBeginPauseReplicationOptions) (*runtime.Poller[ReplicationMigrationItemsClientPauseReplicationResponse], error)`
- New function `*ResumeReplicationProviderSpecificInput.GetResumeReplicationProviderSpecificInput() *ResumeReplicationProviderSpecificInput`
- New struct `A2AExtendedLocationDetails`
- New struct `CriticalJobHistoryDetails`
- New struct `PauseReplicationInput`
- New struct `PauseReplicationInputProperties`
- New struct `ReplicationMigrationItemsClientBeginPauseReplicationOptions`
- New struct `ReplicationMigrationItemsClientBeginResumeReplicationOptions`
- New struct `ReplicationMigrationItemsClientPauseReplicationResponse`
- New struct `ReplicationMigrationItemsClientResumeReplicationResponse`
- New struct `ResumeReplicationInput`
- New struct `ResumeReplicationInputProperties`
- New struct `ResumeReplicationProviderSpecificInput`
- New struct `VMwareCbtResumeReplicationInput`
- New field `ResumeProgressPercentage` in struct `VMwareCbtMigrationDetails`
- New field `ResumeRetryCount` in struct `VMwareCbtMigrationDetails`
- New field `TestNetworkID` in struct `VMwareCbtMigrationDetails`
- New field `StorageAccountID` in struct `VMwareCbtMigrationDetails`
- New field `TargetBlobURI` in struct `VMwareCbtProtectedDiskDetails`
- New field `SeedBlobURI` in struct `VMwareCbtProtectedDiskDetails`
- New field `CriticalJobHistory` in struct `MigrationItemProperties`
- New field `RecoveryServicesProviderID` in struct `MigrationItemProperties`
- New field `LastMigrationTime` in struct `MigrationItemProperties`
- New field `LastMigrationStatus` in struct `MigrationItemProperties`
- New field `ReplicationStatus` in struct `MigrationItemProperties`
- New field `ExtendedLocations` in struct `AzureFabricSpecificDetails`
- New field `RecoveryExtendedLocation` in struct `RecoveryPlanA2ADetails`
- New field `PrimaryExtendedLocation` in struct `RecoveryPlanA2ADetails`
- New field `PerformSQLBulkRegistration` in struct `VMwareCbtEnableMigrationInput`
- New field `RoleSizeToNicCountMap` in struct `VMwareCbtProtectionContainerMappingDetails`
- New field `StorageAccountID` in struct `InMageRcmReplicationDetails`
- New field `SeedBlobURI` in struct `InMageRcmProtectedDiskDetails`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).