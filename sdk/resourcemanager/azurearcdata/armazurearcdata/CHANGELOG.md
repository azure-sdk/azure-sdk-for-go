# Release History

## 0.6.0 (2022-11-07)
### Features Added

- New const `DatabaseStateSuspect`
- New const `RecoveryModeSimple`
- New const `DatabaseStateEmergency`
- New const `HostTypeAWSKubernetesService`
- New const `HostTypeAzureVMWareVirtualMachine`
- New const `HostTypeAWSVMWareVirtualMachine`
- New const `DatabaseStateRecovering`
- New const `DatabaseStateRecoveryPending`
- New const `ArcSQLServerLicenseTypeLicenseOnly`
- New const `HostTypeAzureKubernetesService`
- New const `DatabaseStateOfflineSecondary`
- New const `DatabaseStateOnline`
- New const `DatabaseStateOffline`
- New const `DatabaseStateCopying`
- New const `DatabaseStateRestoring`
- New const `RecoveryModeBulkLogged`
- New const `HostTypeGCPKubernetesService`
- New const `ArcSQLServerLicenseTypePAYG`
- New const `HostTypeContainer`
- New const `HostTypeAzureVirtualMachine`
- New const `ArcSQLServerLicenseTypeServerCAL`
- New const `RecoveryModeFull`
- New const `HostTypeGCPVMWareVirtualMachine`
- New type alias `DatabaseState`
- New type alias `RecoveryMode`
- New function `*SQLServerDatabasesClient.NewListPager(string, string, *SQLServerDatabasesClientListOptions) *runtime.Pager[SQLServerDatabasesClientListResponse]`
- New function `*SQLServerDatabasesClient.Create(context.Context, string, string, string, SQLServerDatabaseResource, *SQLServerDatabasesClientCreateOptions) (SQLServerDatabasesClientCreateResponse, error)`
- New function `PossibleRecoveryModeValues() []RecoveryMode`
- New function `NewSQLServerDatabasesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLServerDatabasesClient, error)`
- New function `*SQLServerDatabasesClient.Get(context.Context, string, string, string, *SQLServerDatabasesClientGetOptions) (SQLServerDatabasesClientGetResponse, error)`
- New function `PossibleDatabaseStateValues() []DatabaseState`
- New function `*SQLServerDatabasesClient.Delete(context.Context, string, string, string, *SQLServerDatabasesClientDeleteOptions) (SQLServerDatabasesClientDeleteResponse, error)`
- New function `*SQLServerDatabasesClient.Update(context.Context, string, string, string, SQLServerDatabaseUpdate, *SQLServerDatabasesClientUpdateOptions) (SQLServerDatabasesClientUpdateResponse, error)`
- New struct `ArcSQLServerDatabaseListResult`
- New struct `SQLServerDatabaseResource`
- New struct `SQLServerDatabaseResourceProperties`
- New struct `SQLServerDatabaseResourcePropertiesBackupInformation`
- New struct `SQLServerDatabaseResourcePropertiesDatabaseOptions`
- New struct `SQLServerDatabaseUpdate`
- New struct `SQLServerDatabasesClient`
- New struct `SQLServerDatabasesClientCreateOptions`
- New struct `SQLServerDatabasesClientCreateResponse`
- New struct `SQLServerDatabasesClientDeleteOptions`
- New struct `SQLServerDatabasesClientDeleteResponse`
- New struct `SQLServerDatabasesClientGetOptions`
- New struct `SQLServerDatabasesClientGetResponse`
- New struct `SQLServerDatabasesClientListOptions`
- New struct `SQLServerDatabasesClientListResponse`
- New struct `SQLServerDatabasesClientUpdateOptions`
- New struct `SQLServerDatabasesClientUpdateResponse`
- New field `Cores` in struct `SQLServerInstanceProperties`


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).