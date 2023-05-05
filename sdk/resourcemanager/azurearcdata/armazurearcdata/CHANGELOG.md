# Release History

## 0.7.0 (2023-05-05)
### Features Added

- New value `ArcSQLServerLicenseTypeLicenseOnly`, `ArcSQLServerLicenseTypePAYG`, `ArcSQLServerLicenseTypeServerCAL` added to enum type `ArcSQLServerLicenseType`
- New value `EditionTypeBusinessIntelligence` added to enum type `EditionType`
- New value `HostTypeAWSKubernetesService`, `HostTypeAWSVMWareVirtualMachine`, `HostTypeAzureKubernetesService`, `HostTypeAzureVMWareVirtualMachine`, `HostTypeAzureVirtualMachine`, `HostTypeContainer`, `HostTypeGCPKubernetesService`, `HostTypeGCPVMWareVirtualMachine` added to enum type `HostType`
- New enum type `DatabaseState` with values `DatabaseStateCopying`, `DatabaseStateEmergency`, `DatabaseStateOffline`, `DatabaseStateOfflineSecondary`, `DatabaseStateOnline`, `DatabaseStateRecovering`, `DatabaseStateRecoveryPending`, `DatabaseStateRestoring`, `DatabaseStateSuspect`
- New enum type `FailoverGroupPartnerSyncMode` with values `FailoverGroupPartnerSyncModeAsync`, `FailoverGroupPartnerSyncModeSync`
- New enum type `InstanceFailoverGroupRole` with values `InstanceFailoverGroupRoleForcePrimaryAllowDataLoss`, `InstanceFailoverGroupRoleForceSecondary`, `InstanceFailoverGroupRolePrimary`, `InstanceFailoverGroupRoleSecondary`
- New enum type `ProvisioningState` with values `ProvisioningStateAccepted`, `ProvisioningStateCanceled`, `ProvisioningStateFailed`, `ProvisioningStateSucceeded`
- New enum type `RecoveryMode` with values `RecoveryModeBulkLogged`, `RecoveryModeFull`, `RecoveryModeSimple`
- New function `*ClientFactory.NewFailoverGroupsClient() *FailoverGroupsClient`
- New function `*ClientFactory.NewSQLServerDatabasesClient() *SQLServerDatabasesClient`
- New function `NewFailoverGroupsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*FailoverGroupsClient, error)`
- New function `*FailoverGroupsClient.BeginCreate(context.Context, string, string, string, FailoverGroupResource, *FailoverGroupsClientBeginCreateOptions) (*runtime.Poller[FailoverGroupsClientCreateResponse], error)`
- New function `*FailoverGroupsClient.BeginDelete(context.Context, string, string, string, *FailoverGroupsClientBeginDeleteOptions) (*runtime.Poller[FailoverGroupsClientDeleteResponse], error)`
- New function `*FailoverGroupsClient.Get(context.Context, string, string, string, *FailoverGroupsClientGetOptions) (FailoverGroupsClientGetResponse, error)`
- New function `*FailoverGroupsClient.NewListPager(string, string, *FailoverGroupsClientListOptions) *runtime.Pager[FailoverGroupsClientListResponse]`
- New function `NewSQLServerDatabasesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLServerDatabasesClient, error)`
- New function `*SQLServerDatabasesClient.Create(context.Context, string, string, string, SQLServerDatabaseResource, *SQLServerDatabasesClientCreateOptions) (SQLServerDatabasesClientCreateResponse, error)`
- New function `*SQLServerDatabasesClient.Delete(context.Context, string, string, string, *SQLServerDatabasesClientDeleteOptions) (SQLServerDatabasesClientDeleteResponse, error)`
- New function `*SQLServerDatabasesClient.Get(context.Context, string, string, string, *SQLServerDatabasesClientGetOptions) (SQLServerDatabasesClientGetResponse, error)`
- New function `*SQLServerDatabasesClient.NewListPager(string, string, *SQLServerDatabasesClientListOptions) *runtime.Pager[SQLServerDatabasesClientListResponse]`
- New function `*SQLServerDatabasesClient.Update(context.Context, string, string, string, SQLServerDatabaseUpdate, *SQLServerDatabasesClientUpdateOptions) (SQLServerDatabasesClientUpdateResponse, error)`
- New struct `ArcSQLServerDatabaseListResult`
- New struct `FailoverGroupListResult`
- New struct `FailoverGroupProperties`
- New struct `FailoverGroupResource`
- New struct `FailoverGroupSpec`
- New struct `SQLServerDatabaseResource`
- New struct `SQLServerDatabaseResourceProperties`
- New struct `SQLServerDatabaseResourcePropertiesBackupInformation`
- New struct `SQLServerDatabaseResourcePropertiesDatabaseOptions`
- New struct `SQLServerDatabaseUpdate`
- New field `Cores` in struct `SQLServerInstanceProperties`


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