# Release History

## 0.8.0 (2024-10-01)
### Breaking Changes

- Type of `ErrorResponse.Error` has been changed from `*ErrorResponseBody` to `*ErrorDetail`
- Operation `*SQLServerInstancesClient.Update` has been changed to LRO, use `*SQLServerInstancesClient.BeginUpdate` instead.
- Struct `ErrorResponseBody` has been removed

### Features Added

- New value `ArcSQLServerLicenseTypeLicenseOnly`, `ArcSQLServerLicenseTypePAYG`, `ArcSQLServerLicenseTypeServerCAL` added to enum type `ArcSQLServerLicenseType`
- New value `EditionTypeBusinessIntelligence` added to enum type `EditionType`
- New value `HostTypeAWSKubernetesService`, `HostTypeAWSVMWareVirtualMachine`, `HostTypeAzureKubernetesService`, `HostTypeAzureVMWareVirtualMachine`, `HostTypeAzureVirtualMachine`, `HostTypeContainer`, `HostTypeGCPKubernetesService`, `HostTypeGCPVMWareVirtualMachine` added to enum type `HostType`
- New enum type `ActivationState` with values `ActivationStateActivated`, `ActivationStateDeactivated`
- New enum type `AggregationType` with values `AggregationTypeAverage`, `AggregationTypeCount`, `AggregationTypeMaximum`, `AggregationTypeMinimum`, `AggregationTypeSum`
- New enum type `AlwaysOnRole` with values `AlwaysOnRoleAvailabilityGroupReplica`, `AlwaysOnRoleFailoverClusterInstance`, `AlwaysOnRoleFailoverClusterNode`, `AlwaysOnRoleNone`
- New enum type `ArcSQLServerAvailabilityMode` with values `ArcSQLServerAvailabilityModeASYNCHRONOUSCOMMIT`, `ArcSQLServerAvailabilityModeSYNCHRONOUSCOMMIT`
- New enum type `ArcSQLServerFailoverMode` with values `ArcSQLServerFailoverModeAUTOMATIC`, `ArcSQLServerFailoverModeEXTERNAL`, `ArcSQLServerFailoverModeMANUAL`
- New enum type `AutomatedBackupPreference` with values `AutomatedBackupPreferenceNONE`, `AutomatedBackupPreferencePRIMARY`, `AutomatedBackupPreferenceSECONDARY`, `AutomatedBackupPreferenceSECONDARYONLY`
- New enum type `BillingPlan` with values `BillingPlanPAYG`, `BillingPlanPaid`
- New enum type `ClusterType` with values `ClusterTypeNONE`, `ClusterTypeWSFC`
- New enum type `ConnectionAuth` with values `ConnectionAuthCertificate`, `ConnectionAuthCertificateWindowsKerberos`, `ConnectionAuthCertificateWindowsNTLM`, `ConnectionAuthCertificateWindowsNegotiate`, `ConnectionAuthWindowsKerberos`, `ConnectionAuthWindowsKerberosCertificate`, `ConnectionAuthWindowsNTLM`, `ConnectionAuthWindowsNTLMCertificate`, `ConnectionAuthWindowsNegotiate`, `ConnectionAuthWindowsNegotiateCertificate`
- New enum type `DatabaseCreateMode` with values `DatabaseCreateModeDefault`, `DatabaseCreateModePointInTimeRestore`
- New enum type `DatabaseState` with values `DatabaseStateCopying`, `DatabaseStateEmergency`, `DatabaseStateOffline`, `DatabaseStateOfflineSecondary`, `DatabaseStateOnline`, `DatabaseStateRecovering`, `DatabaseStateRecoveryPending`, `DatabaseStateRestoring`, `DatabaseStateSuspect`
- New enum type `DbFailover` with values `DbFailoverOFF`, `DbFailoverON`
- New enum type `DifferentialBackupHours` with values `DifferentialBackupHoursTwelve`, `DifferentialBackupHoursTwentyFour`
- New enum type `DtcSupport` with values `DtcSupportNONE`, `DtcSupportPERDB`
- New enum type `EncryptionAlgorithm` with values `EncryptionAlgorithmAES`, `EncryptionAlgorithmAESRC4`, `EncryptionAlgorithmNONE`, `EncryptionAlgorithmNONEAES`, `EncryptionAlgorithmNONEAESRC4`, `EncryptionAlgorithmNONERC4`, `EncryptionAlgorithmNONERC4AES`, `EncryptionAlgorithmRC4`, `EncryptionAlgorithmRC4AES`
- New enum type `ExecutionState` with values `ExecutionStateRunning`, `ExecutionStateWaiting`
- New enum type `FailoverGroupPartnerSyncMode` with values `FailoverGroupPartnerSyncModeAsync`, `FailoverGroupPartnerSyncModeSync`
- New enum type `FailureConditionLevel` with values `FailureConditionLevelFive`, `FailureConditionLevelFour`, `FailureConditionLevelOne`, `FailureConditionLevelThree`, `FailureConditionLevelTwo`
- New enum type `InstanceFailoverGroupRole` with values `InstanceFailoverGroupRoleForcePrimaryAllowDataLoss`, `InstanceFailoverGroupRoleForceSecondary`, `InstanceFailoverGroupRolePrimary`, `InstanceFailoverGroupRoleSecondary`
- New enum type `JobStatus` with values `JobStatusFailed`, `JobStatusInProgress`, `JobStatusNotStarted`, `JobStatusSucceeded`
- New enum type `LastExecutionStatus` with values `LastExecutionStatusCompleted`, `LastExecutionStatusFailed`, `LastExecutionStatusFaulted`, `LastExecutionStatusPostponed`, `LastExecutionStatusRescheduled`, `LastExecutionStatusSucceeded`
- New enum type `LicenseCategory` with values `LicenseCategoryCore`
- New enum type `PrimaryAllowConnections` with values `PrimaryAllowConnectionsALL`, `PrimaryAllowConnectionsREADWRITE`
- New enum type `ProvisioningState` with values `ProvisioningStateAccepted`, `ProvisioningStateCanceled`, `ProvisioningStateFailed`, `ProvisioningStateSucceeded`
- New enum type `RecommendationStatus` with values `RecommendationStatusNotReady`, `RecommendationStatusReady`, `RecommendationStatusUnknown`
- New enum type `RecoveryMode` with values `RecoveryModeBulkLogged`, `RecoveryModeFull`, `RecoveryModeSimple`
- New enum type `ReplicationPartnerType` with values `ReplicationPartnerTypeAzureSQLManagedInstance`, `ReplicationPartnerTypeAzureSQLVM`, `ReplicationPartnerTypeSQLServer`, `ReplicationPartnerTypeUnknown`
- New enum type `Result` with values `ResultFailed`, `ResultNotCompleted`, `ResultSkipped`, `ResultSucceeded`, `ResultTimedOut`
- New enum type `Role` with values `RoleALL`, `RoleNONE`, `RolePARTNER`, `RoleWITNESS`
- New enum type `SQLServerInstanceTelemetryColumnType` with values `SQLServerInstanceTelemetryColumnTypeBool`, `SQLServerInstanceTelemetryColumnTypeDatetime`, `SQLServerInstanceTelemetryColumnTypeDouble`, `SQLServerInstanceTelemetryColumnTypeGUID`, `SQLServerInstanceTelemetryColumnTypeInt`, `SQLServerInstanceTelemetryColumnTypeLong`, `SQLServerInstanceTelemetryColumnTypeString`, `SQLServerInstanceTelemetryColumnTypeTimespan`
- New enum type `ScopeType` with values `ScopeTypeResourceGroup`, `ScopeTypeSubscription`, `ScopeTypeTenant`
- New enum type `SecondaryAllowConnections` with values `SecondaryAllowConnectionsALL`, `SecondaryAllowConnectionsNO`, `SecondaryAllowConnectionsREADONLY`
- New enum type `SeedingMode` with values `SeedingModeAUTOMATIC`, `SeedingModeMANUAL`
- New enum type `SequencerState` with values `SequencerStateCompleted`, `SequencerStateCreatingSuccessors`, `SequencerStateExecutingAction`, `SequencerStateNotStarted`, `SequencerStateWaitingPredecessors`
- New enum type `ServiceType` with values `ServiceTypeEngine`, `ServiceTypePBIRS`, `ServiceTypeSSAS`, `ServiceTypeSSIS`, `ServiceTypeSSRS`
- New enum type `State` with values `StateActive`, `StateCompleted`, `StateDeleted`, `StateDisabled`, `StateEnabled`, `StateFaulted`, `StateInactive`, `StateSuspended`, `StateTerminated`
- New enum type `Version` with values `VersionSQLServer2012`, `VersionSQLServer2014`
- New function `*ClientFactory.NewFailoverGroupsClient() *FailoverGroupsClient`
- New function `*ClientFactory.NewSQLServerAvailabilityGroupsClient() *SQLServerAvailabilityGroupsClient`
- New function `*ClientFactory.NewSQLServerDatabasesClient() *SQLServerDatabasesClient`
- New function `*ClientFactory.NewSQLServerEsuLicensesClient() *SQLServerEsuLicensesClient`
- New function `*ClientFactory.NewSQLServerLicensesClient() *SQLServerLicensesClient`
- New function `NewFailoverGroupsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*FailoverGroupsClient, error)`
- New function `*FailoverGroupsClient.BeginCreate(context.Context, string, string, string, FailoverGroupResource, *FailoverGroupsClientBeginCreateOptions) (*runtime.Poller[FailoverGroupsClientCreateResponse], error)`
- New function `*FailoverGroupsClient.BeginDelete(context.Context, string, string, string, *FailoverGroupsClientBeginDeleteOptions) (*runtime.Poller[FailoverGroupsClientDeleteResponse], error)`
- New function `*FailoverGroupsClient.Get(context.Context, string, string, string, *FailoverGroupsClientGetOptions) (FailoverGroupsClientGetResponse, error)`
- New function `*FailoverGroupsClient.NewListPager(string, string, *FailoverGroupsClientListOptions) *runtime.Pager[FailoverGroupsClientListResponse]`
- New function `NewSQLServerAvailabilityGroupsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLServerAvailabilityGroupsClient, error)`
- New function `*SQLServerAvailabilityGroupsClient.AddDatabases(context.Context, string, string, string, Databases, *SQLServerAvailabilityGroupsClientAddDatabasesOptions) (SQLServerAvailabilityGroupsClientAddDatabasesResponse, error)`
- New function `*SQLServerAvailabilityGroupsClient.Create(context.Context, string, string, string, SQLServerAvailabilityGroupResource, *SQLServerAvailabilityGroupsClientCreateOptions) (SQLServerAvailabilityGroupsClientCreateResponse, error)`
- New function `*SQLServerAvailabilityGroupsClient.BeginCreateAvailabilityGroup(context.Context, string, string, AvailabilityGroupCreateUpdateConfiguration, *SQLServerAvailabilityGroupsClientBeginCreateAvailabilityGroupOptions) (*runtime.Poller[SQLServerAvailabilityGroupsClientCreateAvailabilityGroupResponse], error)`
- New function `*SQLServerAvailabilityGroupsClient.BeginCreateDistributedAvailabilityGroup(context.Context, string, string, DistributedAvailabilityGroupCreateUpdateConfiguration, *SQLServerAvailabilityGroupsClientBeginCreateDistributedAvailabilityGroupOptions) (*runtime.Poller[SQLServerAvailabilityGroupsClientCreateDistributedAvailabilityGroupResponse], error)`
- New function `*SQLServerAvailabilityGroupsClient.BeginDelete(context.Context, string, string, string, *SQLServerAvailabilityGroupsClientBeginDeleteOptions) (*runtime.Poller[SQLServerAvailabilityGroupsClientDeleteResponse], error)`
- New function `*SQLServerAvailabilityGroupsClient.DetailView(context.Context, string, string, string, *SQLServerAvailabilityGroupsClientDetailViewOptions) (SQLServerAvailabilityGroupsClientDetailViewResponse, error)`
- New function `*SQLServerAvailabilityGroupsClient.Failover(context.Context, string, string, string, *SQLServerAvailabilityGroupsClientFailoverOptions) (SQLServerAvailabilityGroupsClientFailoverResponse, error)`
- New function `*SQLServerAvailabilityGroupsClient.ForceFailoverAllowDataLoss(context.Context, string, string, string, *SQLServerAvailabilityGroupsClientForceFailoverAllowDataLossOptions) (SQLServerAvailabilityGroupsClientForceFailoverAllowDataLossResponse, error)`
- New function `*SQLServerAvailabilityGroupsClient.Get(context.Context, string, string, string, *SQLServerAvailabilityGroupsClientGetOptions) (SQLServerAvailabilityGroupsClientGetResponse, error)`
- New function `*SQLServerAvailabilityGroupsClient.NewListPager(string, string, *SQLServerAvailabilityGroupsClientListOptions) *runtime.Pager[SQLServerAvailabilityGroupsClientListResponse]`
- New function `*SQLServerAvailabilityGroupsClient.RemoveDatabases(context.Context, string, string, string, Databases, *SQLServerAvailabilityGroupsClientRemoveDatabasesOptions) (SQLServerAvailabilityGroupsClientRemoveDatabasesResponse, error)`
- New function `*SQLServerAvailabilityGroupsClient.BeginUpdate(context.Context, string, string, string, SQLServerAvailabilityGroupUpdate, *SQLServerAvailabilityGroupsClientBeginUpdateOptions) (*runtime.Poller[SQLServerAvailabilityGroupsClientUpdateResponse], error)`
- New function `NewSQLServerDatabasesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLServerDatabasesClient, error)`
- New function `*SQLServerDatabasesClient.Create(context.Context, string, string, string, SQLServerDatabaseResource, *SQLServerDatabasesClientCreateOptions) (SQLServerDatabasesClientCreateResponse, error)`
- New function `*SQLServerDatabasesClient.BeginDelete(context.Context, string, string, string, *SQLServerDatabasesClientBeginDeleteOptions) (*runtime.Poller[SQLServerDatabasesClientDeleteResponse], error)`
- New function `*SQLServerDatabasesClient.Get(context.Context, string, string, string, *SQLServerDatabasesClientGetOptions) (SQLServerDatabasesClientGetResponse, error)`
- New function `*SQLServerDatabasesClient.NewListPager(string, string, *SQLServerDatabasesClientListOptions) *runtime.Pager[SQLServerDatabasesClientListResponse]`
- New function `*SQLServerDatabasesClient.BeginUpdate(context.Context, string, string, string, SQLServerDatabaseUpdate, *SQLServerDatabasesClientBeginUpdateOptions) (*runtime.Poller[SQLServerDatabasesClientUpdateResponse], error)`
- New function `NewSQLServerEsuLicensesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLServerEsuLicensesClient, error)`
- New function `*SQLServerEsuLicensesClient.Create(context.Context, string, string, SQLServerEsuLicense, *SQLServerEsuLicensesClientCreateOptions) (SQLServerEsuLicensesClientCreateResponse, error)`
- New function `*SQLServerEsuLicensesClient.Delete(context.Context, string, string, *SQLServerEsuLicensesClientDeleteOptions) (SQLServerEsuLicensesClientDeleteResponse, error)`
- New function `*SQLServerEsuLicensesClient.Get(context.Context, string, string, *SQLServerEsuLicensesClientGetOptions) (SQLServerEsuLicensesClientGetResponse, error)`
- New function `*SQLServerEsuLicensesClient.NewListByResourceGroupPager(string, *SQLServerEsuLicensesClientListByResourceGroupOptions) *runtime.Pager[SQLServerEsuLicensesClientListByResourceGroupResponse]`
- New function `*SQLServerEsuLicensesClient.NewListPager(*SQLServerEsuLicensesClientListOptions) *runtime.Pager[SQLServerEsuLicensesClientListResponse]`
- New function `*SQLServerEsuLicensesClient.Update(context.Context, string, string, SQLServerEsuLicenseUpdate, *SQLServerEsuLicensesClientUpdateOptions) (SQLServerEsuLicensesClientUpdateResponse, error)`
- New function `*SQLServerInstancesClient.GetJobsStatus(context.Context, string, string, *SQLServerInstancesClientGetJobsStatusOptions) (SQLServerInstancesClientGetJobsStatusResponse, error)`
- New function `*SQLServerInstancesClient.BeginGetTelemetry(context.Context, string, string, SQLServerInstanceTelemetryRequest, *SQLServerInstancesClientBeginGetTelemetryOptions) (*runtime.Poller[*runtime.Pager[SQLServerInstancesClientGetTelemetryResponse]], error)`
- New function `*SQLServerInstancesClient.PostUpgrade(context.Context, string, string, *SQLServerInstancesClientPostUpgradeOptions) (SQLServerInstancesClientPostUpgradeResponse, error)`
- New function `*SQLServerInstancesClient.PreUpgrade(context.Context, string, string, *SQLServerInstancesClientPreUpgradeOptions) (SQLServerInstancesClientPreUpgradeResponse, error)`
- New function `*SQLServerInstancesClient.RunMigrationAssessment(context.Context, string, string, *SQLServerInstancesClientRunMigrationAssessmentOptions) (SQLServerInstancesClientRunMigrationAssessmentResponse, error)`
- New function `NewSQLServerLicensesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLServerLicensesClient, error)`
- New function `*SQLServerLicensesClient.Create(context.Context, string, string, SQLServerLicense, *SQLServerLicensesClientCreateOptions) (SQLServerLicensesClientCreateResponse, error)`
- New function `*SQLServerLicensesClient.Delete(context.Context, string, string, *SQLServerLicensesClientDeleteOptions) (SQLServerLicensesClientDeleteResponse, error)`
- New function `*SQLServerLicensesClient.Get(context.Context, string, string, *SQLServerLicensesClientGetOptions) (SQLServerLicensesClientGetResponse, error)`
- New function `*SQLServerLicensesClient.NewListByResourceGroupPager(string, *SQLServerLicensesClientListByResourceGroupOptions) *runtime.Pager[SQLServerLicensesClientListByResourceGroupResponse]`
- New function `*SQLServerLicensesClient.NewListPager(*SQLServerLicensesClientListOptions) *runtime.Pager[SQLServerLicensesClientListResponse]`
- New function `*SQLServerLicensesClient.Update(context.Context, string, string, SQLServerLicenseUpdate, *SQLServerLicensesClientUpdateOptions) (SQLServerLicensesClientUpdateResponse, error)`
- New struct `ArcSQLServerAvailabilityGroupListResult`
- New struct `ArcSQLServerDatabaseListResult`
- New struct `AvailabilityGroupConfigure`
- New struct `AvailabilityGroupCreateUpdateConfiguration`
- New struct `AvailabilityGroupCreateUpdateReplicaConfiguration`
- New struct `AvailabilityGroupInfo`
- New struct `AvailabilityGroupState`
- New struct `BackgroundJob`
- New struct `BackupPolicy`
- New struct `ClientConnection`
- New struct `DBMEndpoint`
- New struct `DataBaseMigration`
- New struct `DataBaseMigrationAssessment`
- New struct `DatabaseAssessmentsItem`
- New struct `Databases`
- New struct `DistributedAvailabilityGroupCreateUpdateAvailabilityGroupCertificateConfiguration`
- New struct `DistributedAvailabilityGroupCreateUpdateAvailabilityGroupConfiguration`
- New struct `DistributedAvailabilityGroupCreateUpdateConfiguration`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `FailoverCluster`
- New struct `FailoverGroupListResult`
- New struct `FailoverGroupProperties`
- New struct `FailoverGroupResource`
- New struct `FailoverGroupSpec`
- New struct `K8SActiveDirectory`
- New struct `K8SActiveDirectoryConnector`
- New struct `K8SNetworkSettings`
- New struct `K8SSecurity`
- New struct `K8SSettings`
- New struct `K8StransparentDataEncryption`
- New struct `Migration`
- New struct `MigrationAssessment`
- New struct `Monitoring`
- New struct `ResourceAutoGenerated`
- New struct `SKURecommendationResults`
- New struct `SKURecommendationResultsAzureSQLDatabase`
- New struct `SKURecommendationResultsAzureSQLDatabaseTargetSKU`
- New struct `SKURecommendationResultsAzureSQLDatabaseTargetSKUCategory`
- New struct `SKURecommendationResultsAzureSQLManagedInstance`
- New struct `SKURecommendationResultsAzureSQLManagedInstanceTargetSKU`
- New struct `SKURecommendationResultsAzureSQLManagedInstanceTargetSKUCategory`
- New struct `SKURecommendationResultsAzureSQLVirtualMachine`
- New struct `SKURecommendationResultsAzureSQLVirtualMachineTargetSKU`
- New struct `SKURecommendationResultsAzureSQLVirtualMachineTargetSKUCategory`
- New struct `SKURecommendationResultsMonthlyCost`
- New struct `SKURecommendationSummary`
- New struct `SQLAvailabilityGroupDatabaseReplicaResourceProperties`
- New struct `SQLAvailabilityGroupDhcpListenerProperties`
- New struct `SQLAvailabilityGroupIPV4AddressesAndMasksPropertiesItem`
- New struct `SQLAvailabilityGroupReplicaResourceProperties`
- New struct `SQLAvailabilityGroupStaticIPListenerProperties`
- New struct `SQLServerAvailabilityGroupResource`
- New struct `SQLServerAvailabilityGroupResourceProperties`
- New struct `SQLServerAvailabilityGroupResourcePropertiesDatabases`
- New struct `SQLServerAvailabilityGroupResourcePropertiesReplicas`
- New struct `SQLServerAvailabilityGroupUpdate`
- New struct `SQLServerDatabaseResource`
- New struct `SQLServerDatabaseResourceProperties`
- New struct `SQLServerDatabaseResourcePropertiesBackupInformation`
- New struct `SQLServerDatabaseResourcePropertiesDatabaseOptions`
- New struct `SQLServerDatabaseUpdate`
- New struct `SQLServerEsuLicense`
- New struct `SQLServerEsuLicenseListResult`
- New struct `SQLServerEsuLicenseProperties`
- New struct `SQLServerEsuLicenseUpdate`
- New struct `SQLServerEsuLicenseUpdateProperties`
- New struct `SQLServerInstanceJobStatus`
- New struct `SQLServerInstanceJobsStatusRequest`
- New struct `SQLServerInstanceJobsStatusResponse`
- New struct `SQLServerInstanceRunMigrationAssessmentResponse`
- New struct `SQLServerInstanceTelemetryColumn`
- New struct `SQLServerInstanceTelemetryRequest`
- New struct `SQLServerInstanceTelemetryResponse`
- New struct `SQLServerInstanceUpdateProperties`
- New struct `SQLServerInstancesClientGetTelemetryResponse`
- New struct `SQLServerLicense`
- New struct `SQLServerLicenseListResult`
- New struct `SQLServerLicenseProperties`
- New struct `SQLServerLicenseUpdate`
- New struct `SQLServerLicenseUpdateProperties`
- New struct `SequencerAction`
- New struct `ServerAssessmentsItem`
- New struct `ServerAssessmentsPropertiesItemsItem`
- New struct `TargetReadiness`
- New struct `TrackedResourceAutoGenerated`
- New field `Security`, `Settings` in struct `SQLManagedInstanceK8SSpec`
- New field `AlwaysOnRole`, `BackupPolicy`, `ClientConnection`, `Cores`, `DatabaseMirroringEndpoint`, `DbMasterKeyExists`, `FailoverCluster`, `IsHadrEnabled`, `LastInventoryUploadTime`, `LastUsageUploadTime`, `Migration`, `Monitoring`, `ServiceType`, `TraceFlags`, `UpgradeLockedUntil` in struct `SQLServerInstanceProperties`
- New field `Properties` in struct `SQLServerInstanceUpdate`


## 0.7.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 0.6.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.

## 0.6.0 (2023-03-27)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).