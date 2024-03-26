# Release History

## 1.0.0 (2024-03-26)
### Breaking Changes

- Type of `ErrorResponse.Error` has been changed from `*ErrorResponseBody` to `*ErrorDetail`
- Operation `*SQLServerInstancesClient.Update` has been changed to LRO, use `*SQLServerInstancesClient.BeginUpdate` instead.
- Struct `ErrorResponseBody` has been removed

### Features Added

- New value `ArcSQLServerLicenseTypeLicenseOnly`, `ArcSQLServerLicenseTypePAYG`, `ArcSQLServerLicenseTypeServerCAL` added to enum type `ArcSQLServerLicenseType`
- New value `EditionTypeBusinessIntelligence` added to enum type `EditionType`
- New value `HostTypeAWSKubernetesService`, `HostTypeAWSVMWareVirtualMachine`, `HostTypeAzureKubernetesService`, `HostTypeAzureVMWareVirtualMachine`, `HostTypeAzureVirtualMachine`, `HostTypeContainer`, `HostTypeGCPKubernetesService`, `HostTypeGCPVMWareVirtualMachine` added to enum type `HostType`
- New enum type `AggregationType` with values `AggregationTypeAverage`, `AggregationTypeCount`, `AggregationTypeMaximum`, `AggregationTypeMinimum`, `AggregationTypeSum`
- New enum type `AlwaysOnRole` with values `AlwaysOnRoleAvailabilityGroupReplica`, `AlwaysOnRoleFailoverClusterInstance`, `AlwaysOnRoleFailoverClusterNode`, `AlwaysOnRoleNone`
- New enum type `DatabaseCreateMode` with values `DatabaseCreateModeDefault`, `DatabaseCreateModePointInTimeRestore`
- New enum type `DatabaseState` with values `DatabaseStateCopying`, `DatabaseStateEmergency`, `DatabaseStateOffline`, `DatabaseStateOfflineSecondary`, `DatabaseStateOnline`, `DatabaseStateRecovering`, `DatabaseStateRecoveryPending`, `DatabaseStateRestoring`, `DatabaseStateSuspect`
- New enum type `DifferentialBackupHours` with values `DifferentialBackupHoursTwelve`, `DifferentialBackupHoursTwentyFour`
- New enum type `FailoverGroupPartnerSyncMode` with values `FailoverGroupPartnerSyncModeAsync`, `FailoverGroupPartnerSyncModeSync`
- New enum type `InstanceFailoverGroupRole` with values `InstanceFailoverGroupRoleForcePrimaryAllowDataLoss`, `InstanceFailoverGroupRoleForceSecondary`, `InstanceFailoverGroupRolePrimary`, `InstanceFailoverGroupRoleSecondary`
- New enum type `ProvisioningState` with values `ProvisioningStateAccepted`, `ProvisioningStateCanceled`, `ProvisioningStateFailed`, `ProvisioningStateSucceeded`
- New enum type `RecoveryMode` with values `RecoveryModeBulkLogged`, `RecoveryModeFull`, `RecoveryModeSimple`
- New enum type `ReplicationPartnerType` with values `ReplicationPartnerTypeAzureSQLManagedInstance`, `ReplicationPartnerTypeAzureSQLVM`, `ReplicationPartnerTypeSQLServer`, `ReplicationPartnerTypeUnknown`
- New enum type `SQLServerInstanceTelemetryColumnType` with values `SQLServerInstanceTelemetryColumnTypeBool`, `SQLServerInstanceTelemetryColumnTypeDatetime`, `SQLServerInstanceTelemetryColumnTypeDouble`, `SQLServerInstanceTelemetryColumnTypeGUID`, `SQLServerInstanceTelemetryColumnTypeInt`, `SQLServerInstanceTelemetryColumnTypeLong`, `SQLServerInstanceTelemetryColumnTypeString`, `SQLServerInstanceTelemetryColumnTypeTimespan`
- New function `*ClientFactory.NewFailoverGroupsClient() *FailoverGroupsClient`
- New function `*ClientFactory.NewSQLServerAvailabilityGroupsClient() *SQLServerAvailabilityGroupsClient`
- New function `*ClientFactory.NewSQLServerDatabasesClient() *SQLServerDatabasesClient`
- New function `NewFailoverGroupsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*FailoverGroupsClient, error)`
- New function `*FailoverGroupsClient.BeginCreate(context.Context, string, string, string, FailoverGroupResource, *FailoverGroupsClientBeginCreateOptions) (*runtime.Poller[FailoverGroupsClientCreateResponse], error)`
- New function `*FailoverGroupsClient.BeginDelete(context.Context, string, string, string, *FailoverGroupsClientBeginDeleteOptions) (*runtime.Poller[FailoverGroupsClientDeleteResponse], error)`
- New function `*FailoverGroupsClient.Get(context.Context, string, string, string, *FailoverGroupsClientGetOptions) (FailoverGroupsClientGetResponse, error)`
- New function `*FailoverGroupsClient.NewListPager(string, string, *FailoverGroupsClientListOptions) *runtime.Pager[FailoverGroupsClientListResponse]`
- New function `NewSQLServerAvailabilityGroupsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLServerAvailabilityGroupsClient, error)`
- New function `*SQLServerAvailabilityGroupsClient.Create(context.Context, string, string, string, SQLServerAvailabilityGroupResource, *SQLServerAvailabilityGroupsClientCreateOptions) (SQLServerAvailabilityGroupsClientCreateResponse, error)`
- New function `*SQLServerAvailabilityGroupsClient.BeginDelete(context.Context, string, string, string, *SQLServerAvailabilityGroupsClientBeginDeleteOptions) (*runtime.Poller[SQLServerAvailabilityGroupsClientDeleteResponse], error)`
- New function `*SQLServerAvailabilityGroupsClient.DetailView(context.Context, string, string, string, *SQLServerAvailabilityGroupsClientDetailViewOptions) (SQLServerAvailabilityGroupsClientDetailViewResponse, error)`
- New function `*SQLServerAvailabilityGroupsClient.Failover(context.Context, string, string, string, *SQLServerAvailabilityGroupsClientFailoverOptions) (SQLServerAvailabilityGroupsClientFailoverResponse, error)`
- New function `*SQLServerAvailabilityGroupsClient.ForceFailoverAllowDataLoss(context.Context, string, string, string, *SQLServerAvailabilityGroupsClientForceFailoverAllowDataLossOptions) (SQLServerAvailabilityGroupsClientForceFailoverAllowDataLossResponse, error)`
- New function `*SQLServerAvailabilityGroupsClient.Get(context.Context, string, string, string, *SQLServerAvailabilityGroupsClientGetOptions) (SQLServerAvailabilityGroupsClientGetResponse, error)`
- New function `*SQLServerAvailabilityGroupsClient.NewListPager(string, string, *SQLServerAvailabilityGroupsClientListOptions) *runtime.Pager[SQLServerAvailabilityGroupsClientListResponse]`
- New function `*SQLServerAvailabilityGroupsClient.BeginUpdate(context.Context, string, string, string, SQLServerAvailabilityGroupUpdate, *SQLServerAvailabilityGroupsClientBeginUpdateOptions) (*runtime.Poller[SQLServerAvailabilityGroupsClientUpdateResponse], error)`
- New function `NewSQLServerDatabasesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLServerDatabasesClient, error)`
- New function `*SQLServerDatabasesClient.Create(context.Context, string, string, string, SQLServerDatabaseResource, *SQLServerDatabasesClientCreateOptions) (SQLServerDatabasesClientCreateResponse, error)`
- New function `*SQLServerDatabasesClient.BeginDelete(context.Context, string, string, string, *SQLServerDatabasesClientBeginDeleteOptions) (*runtime.Poller[SQLServerDatabasesClientDeleteResponse], error)`
- New function `*SQLServerDatabasesClient.Get(context.Context, string, string, string, *SQLServerDatabasesClientGetOptions) (SQLServerDatabasesClientGetResponse, error)`
- New function `*SQLServerDatabasesClient.NewListPager(string, string, *SQLServerDatabasesClientListOptions) *runtime.Pager[SQLServerDatabasesClientListResponse]`
- New function `*SQLServerDatabasesClient.BeginUpdate(context.Context, string, string, string, SQLServerDatabaseUpdate, *SQLServerDatabasesClientBeginUpdateOptions) (*runtime.Poller[SQLServerDatabasesClientUpdateResponse], error)`
- New function `*SQLServerInstancesClient.BeginGetTelemetry(context.Context, string, string, SQLServerInstanceTelemetryRequest, *SQLServerInstancesClientBeginGetTelemetryOptions) (*runtime.Poller[*runtime.Pager[SQLServerInstancesClientGetTelemetryResponse]], error)`
- New struct `ArcSQLServerAvailabilityGroupListResult`
- New struct `ArcSQLServerDatabaseListResult`
- New struct `AvailabilityGroupConfigure`
- New struct `AvailabilityGroupInfo`
- New struct `AvailabilityGroupState`
- New struct `BackupPolicy`
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
- New struct `Monitoring`
- New struct `SQLAvailabilityGroupDatabaseReplicaResourceProperties`
- New struct `SQLAvailabilityGroupReplicaResourceProperties`
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
- New struct `SQLServerInstanceTelemetryColumn`
- New struct `SQLServerInstanceTelemetryRequest`
- New struct `SQLServerInstanceTelemetryResponse`
- New struct `SQLServerInstanceUpdateProperties`
- New struct `SQLServerInstancesClientGetTelemetryResponse`
- New field `Security`, `Settings` in struct `SQLManagedInstanceK8SSpec`
- New field `AlwaysOnRole`, `BackupPolicy`, `Cores`, `FailoverCluster`, `LastInventoryUploadTime`, `LastUsageUploadTime`, `Monitoring`, `UpgradeLockedUntil` in struct `SQLServerInstanceProperties`
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