# Release History

## 0.6.0 (2023-03-22)
### Features Added

- New value `ArcSQLServerLicenseTypeLicenseOnly`, `ArcSQLServerLicenseTypePAYG`, `ArcSQLServerLicenseTypeServerCAL` added to enum type `ArcSQLServerLicenseType`
- New value `EditionTypeBusinessIntelligence` added to enum type `EditionType`
- New value `HostTypeAWSKubernetesService`, `HostTypeAWSVMWareVirtualMachine`, `HostTypeAzureKubernetesService`, `HostTypeAzureVMWareVirtualMachine`, `HostTypeAzureVirtualMachine`, `HostTypeContainer`, `HostTypeGCPKubernetesService`, `HostTypeGCPVMWareVirtualMachine` added to enum type `HostType`
- New enum type `DatabaseState` with values `DatabaseStateCopying`, `DatabaseStateEmergency`, `DatabaseStateOffline`, `DatabaseStateOfflineSecondary`, `DatabaseStateOnline`, `DatabaseStateRecovering`, `DatabaseStateRecoveryPending`, `DatabaseStateRestoring`, `DatabaseStateSuspect`
- New enum type `FailoverGroupPartnerSyncMode` with values `FailoverGroupPartnerSyncModeAsync`, `FailoverGroupPartnerSyncModeSync`
- New enum type `InstanceFailoverGroupRole` with values `InstanceFailoverGroupRoleForcePrimaryAllowDataLoss`, `InstanceFailoverGroupRoleForceSecondary`, `InstanceFailoverGroupRolePrimary`, `InstanceFailoverGroupRoleSecondary`
- New enum type `ProvisioningState` with values `ProvisioningStateAccepted`, `ProvisioningStateCanceled`, `ProvisioningStateFailed`, `ProvisioningStateSucceeded`
- New enum type `RecoveryMode` with values `RecoveryModeBulkLogged`, `RecoveryModeFull`, `RecoveryModeSimple`
- New enum type `State` with values `StateFailed`, `StateSucceeded`, `StateWaiting`
- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewActiveDirectoryConnectorsClient() *ActiveDirectoryConnectorsClient`
- New function `*ClientFactory.NewDataControllersClient() *DataControllersClient`
- New function `*ClientFactory.NewFailoverGroupsClient() *FailoverGroupsClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewPostgresInstancesClient() *PostgresInstancesClient`
- New function `*ClientFactory.NewSQLAvailabilityGroupDatabasesClient() *SQLAvailabilityGroupDatabasesClient`
- New function `*ClientFactory.NewSQLAvailabilityGroupReplicasClient() *SQLAvailabilityGroupReplicasClient`
- New function `*ClientFactory.NewSQLAvailabilityGroupsClient() *SQLAvailabilityGroupsClient`
- New function `*ClientFactory.NewSQLManagedInstancesClient() *SQLManagedInstancesClient`
- New function `*ClientFactory.NewSQLServerAvailabilityGroupsClient() *SQLServerAvailabilityGroupsClient`
- New function `*ClientFactory.NewSQLServerDatabasesClient() *SQLServerDatabasesClient`
- New function `*ClientFactory.NewSQLServerInstancesClient() *SQLServerInstancesClient`
- New function `NewFailoverGroupsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*FailoverGroupsClient, error)`
- New function `*FailoverGroupsClient.BeginCreate(context.Context, string, string, string, FailoverGroupResource, *FailoverGroupsClientBeginCreateOptions) (*runtime.Poller[FailoverGroupsClientCreateResponse], error)`
- New function `*FailoverGroupsClient.BeginDelete(context.Context, string, string, string, *FailoverGroupsClientBeginDeleteOptions) (*runtime.Poller[FailoverGroupsClientDeleteResponse], error)`
- New function `*FailoverGroupsClient.Get(context.Context, string, string, string, *FailoverGroupsClientGetOptions) (FailoverGroupsClientGetResponse, error)`
- New function `*FailoverGroupsClient.NewListPager(string, string, *FailoverGroupsClientListOptions) *runtime.Pager[FailoverGroupsClientListResponse]`
- New function `NewSQLAvailabilityGroupDatabasesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLAvailabilityGroupDatabasesClient, error)`
- New function `*SQLAvailabilityGroupDatabasesClient.Create(context.Context, string, string, string, SQLAvailabilityGroupDatabaseResource, *SQLAvailabilityGroupDatabasesClientCreateOptions) (SQLAvailabilityGroupDatabasesClientCreateResponse, error)`
- New function `*SQLAvailabilityGroupDatabasesClient.Delete(context.Context, string, string, string, *SQLAvailabilityGroupDatabasesClientDeleteOptions) (SQLAvailabilityGroupDatabasesClientDeleteResponse, error)`
- New function `*SQLAvailabilityGroupDatabasesClient.Get(context.Context, string, string, string, *SQLAvailabilityGroupDatabasesClientGetOptions) (SQLAvailabilityGroupDatabasesClientGetResponse, error)`
- New function `*SQLAvailabilityGroupDatabasesClient.NewListPager(string, string, *SQLAvailabilityGroupDatabasesClientListOptions) *runtime.Pager[SQLAvailabilityGroupDatabasesClientListResponse]`
- New function `*SQLAvailabilityGroupDatabasesClient.Update(context.Context, string, string, string, SQLAvailabilityGroupDatabaseUpdate, *SQLAvailabilityGroupDatabasesClientUpdateOptions) (SQLAvailabilityGroupDatabasesClientUpdateResponse, error)`
- New function `NewSQLAvailabilityGroupReplicasClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLAvailabilityGroupReplicasClient, error)`
- New function `*SQLAvailabilityGroupReplicasClient.Create(context.Context, string, string, string, SQLAvailabilityGroupReplicaResource, *SQLAvailabilityGroupReplicasClientCreateOptions) (SQLAvailabilityGroupReplicasClientCreateResponse, error)`
- New function `*SQLAvailabilityGroupReplicasClient.Delete(context.Context, string, string, string, *SQLAvailabilityGroupReplicasClientDeleteOptions) (SQLAvailabilityGroupReplicasClientDeleteResponse, error)`
- New function `*SQLAvailabilityGroupReplicasClient.Get(context.Context, string, string, string, *SQLAvailabilityGroupReplicasClientGetOptions) (SQLAvailabilityGroupReplicasClientGetResponse, error)`
- New function `*SQLAvailabilityGroupReplicasClient.NewListPager(string, string, *SQLAvailabilityGroupReplicasClientListOptions) *runtime.Pager[SQLAvailabilityGroupReplicasClientListResponse]`
- New function `*SQLAvailabilityGroupReplicasClient.Update(context.Context, string, string, string, SQLAvailabilityGroupReplicaUpdate, *SQLAvailabilityGroupReplicasClientUpdateOptions) (SQLAvailabilityGroupReplicasClientUpdateResponse, error)`
- New function `NewSQLAvailabilityGroupsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLAvailabilityGroupsClient, error)`
- New function `*SQLAvailabilityGroupsClient.BeginCreate(context.Context, string, string, SQLAvailabilityGroup, *SQLAvailabilityGroupsClientBeginCreateOptions) (*runtime.Poller[SQLAvailabilityGroupsClientCreateResponse], error)`
- New function `*SQLAvailabilityGroupsClient.BeginDelete(context.Context, string, string, *SQLAvailabilityGroupsClientBeginDeleteOptions) (*runtime.Poller[SQLAvailabilityGroupsClientDeleteResponse], error)`
- New function `*SQLAvailabilityGroupsClient.Get(context.Context, string, string, *SQLAvailabilityGroupsClientGetOptions) (SQLAvailabilityGroupsClientGetResponse, error)`
- New function `*SQLAvailabilityGroupsClient.NewListByResourceGroupPager(string, *SQLAvailabilityGroupsClientListByResourceGroupOptions) *runtime.Pager[SQLAvailabilityGroupsClientListByResourceGroupResponse]`
- New function `*SQLAvailabilityGroupsClient.NewListPager(*SQLAvailabilityGroupsClientListOptions) *runtime.Pager[SQLAvailabilityGroupsClientListResponse]`
- New function `*SQLAvailabilityGroupsClient.Update(context.Context, string, string, SQLAvailabilityGroupUpdate, *SQLAvailabilityGroupsClientUpdateOptions) (SQLAvailabilityGroupsClientUpdateResponse, error)`
- New function `NewSQLServerAvailabilityGroupsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLServerAvailabilityGroupsClient, error)`
- New function `*SQLServerAvailabilityGroupsClient.Create(context.Context, string, string, string, SQLServerAvailabilityGroupResource, *SQLServerAvailabilityGroupsClientCreateOptions) (SQLServerAvailabilityGroupsClientCreateResponse, error)`
- New function `*SQLServerAvailabilityGroupsClient.Delete(context.Context, string, string, string, *SQLServerAvailabilityGroupsClientDeleteOptions) (SQLServerAvailabilityGroupsClientDeleteResponse, error)`
- New function `*SQLServerAvailabilityGroupsClient.Get(context.Context, string, string, string, *SQLServerAvailabilityGroupsClientGetOptions) (SQLServerAvailabilityGroupsClientGetResponse, error)`
- New function `*SQLServerAvailabilityGroupsClient.NewListPager(string, string, *SQLServerAvailabilityGroupsClientListOptions) *runtime.Pager[SQLServerAvailabilityGroupsClientListResponse]`
- New function `*SQLServerAvailabilityGroupsClient.Update(context.Context, string, string, string, SQLServerAvailabilityGroupUpdate, *SQLServerAvailabilityGroupsClientUpdateOptions) (SQLServerAvailabilityGroupsClientUpdateResponse, error)`
- New function `NewSQLServerDatabasesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLServerDatabasesClient, error)`
- New function `*SQLServerDatabasesClient.Create(context.Context, string, string, string, SQLServerDatabaseResource, *SQLServerDatabasesClientCreateOptions) (SQLServerDatabasesClientCreateResponse, error)`
- New function `*SQLServerDatabasesClient.Delete(context.Context, string, string, string, *SQLServerDatabasesClientDeleteOptions) (SQLServerDatabasesClientDeleteResponse, error)`
- New function `*SQLServerDatabasesClient.Get(context.Context, string, string, string, *SQLServerDatabasesClientGetOptions) (SQLServerDatabasesClientGetResponse, error)`
- New function `*SQLServerDatabasesClient.NewListPager(string, string, *SQLServerDatabasesClientListOptions) *runtime.Pager[SQLServerDatabasesClientListResponse]`
- New function `*SQLServerDatabasesClient.Update(context.Context, string, string, string, SQLServerDatabaseUpdate, *SQLServerDatabasesClientUpdateOptions) (SQLServerDatabasesClientUpdateResponse, error)`
- New struct `AGReplicas`
- New struct `ArcSQLAvailabilityGroupDatabasesListResult`
- New struct `ArcSQLAvailabilityGroupReplicaListResult`
- New struct `ArcSQLServerAvailabilityGroupListResult`
- New struct `ArcSQLServerDatabaseListResult`
- New struct `AvailabilityGroupConfigure`
- New struct `AvailabilityGroupState`
- New struct `ClientFactory`
- New struct `FailoverGroupListResult`
- New struct `FailoverGroupProperties`
- New struct `FailoverGroupResource`
- New struct `FailoverGroupSpec`
- New struct `FailoverGroupStatus`
- New struct `SQLAvailabilityGroup`
- New struct `SQLAvailabilityGroupDatabaseReplicaResourceProperties`
- New struct `SQLAvailabilityGroupDatabaseResource`
- New struct `SQLAvailabilityGroupDatabaseUpdate`
- New struct `SQLAvailabilityGroupListResult`
- New struct `SQLAvailabilityGroupMultiDatabaseReplicaResourceProperties`
- New struct `SQLAvailabilityGroupProperties`
- New struct `SQLAvailabilityGroupReplicaResource`
- New struct `SQLAvailabilityGroupReplicaResourceProperties`
- New struct `SQLAvailabilityGroupReplicaUpdate`
- New struct `SQLAvailabilityGroupUpdate`
- New struct `SQLServerAvailabilityGroupResource`
- New struct `SQLServerAvailabilityGroupResourceProperties`
- New struct `SQLServerAvailabilityGroupUpdate`
- New struct `SQLServerDatabaseResource`
- New struct `SQLServerDatabaseResourceProperties`
- New struct `SQLServerDatabaseResourcePropertiesBackupInformation`
- New struct `SQLServerDatabaseResourcePropertiesDatabaseOptions`
- New struct `SQLServerDatabaseUpdate`
- New field `Cores` in struct `SQLServerInstanceProperties`


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).