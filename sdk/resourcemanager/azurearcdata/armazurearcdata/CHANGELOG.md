# Release History

## 0.6.0 (2023-01-17)
### Breaking Changes

- Type of `ActiveDirectoryConnectorStatus.AdditionalProperties` has been changed from `map[string]interface{}` to `map[string]any`
- Type of `DataControllerProperties.K8SRaw` has been changed from `interface{}` to `any`
- Type of `K8SResourceRequirements.AdditionalProperties` has been changed from `map[string]interface{}` to `map[string]any`
- Type of `K8SScheduling.AdditionalProperties` has been changed from `map[string]interface{}` to `map[string]any`
- Type of `K8SSchedulingOptions.AdditionalProperties` has been changed from `map[string]interface{}` to `map[string]any`
- Type of `Operation.Properties` has been changed from `map[string]interface{}` to `map[string]any`
- Type of `PostgresInstanceProperties.K8SRaw` has been changed from `interface{}` to `any`
- Type of `SQLManagedInstanceK8SRaw.AdditionalProperties` has been changed from `map[string]interface{}` to `map[string]any`
- Type of `SQLManagedInstanceK8SSpec.AdditionalProperties` has been changed from `map[string]interface{}` to `map[string]any`

### Features Added

- New value `ArcSQLServerLicenseTypeLicenseOnly`, `ArcSQLServerLicenseTypePAYG`, `ArcSQLServerLicenseTypeServerCAL` added to type alias `ArcSQLServerLicenseType`
- New value `EditionTypeBusinessIntelligence` added to type alias `EditionType`
- New value `HostTypeAWSKubernetesService`, `HostTypeAWSVMWareVirtualMachine`, `HostTypeAzureKubernetesService`, `HostTypeAzureVMWareVirtualMachine`, `HostTypeAzureVirtualMachine`, `HostTypeContainer`, `HostTypeGCPKubernetesService`, `HostTypeGCPVMWareVirtualMachine` added to type alias `HostType`
- New type alias `DatabaseState` with values `DatabaseStateCopying`, `DatabaseStateEmergency`, `DatabaseStateOffline`, `DatabaseStateOfflineSecondary`, `DatabaseStateOnline`, `DatabaseStateRecovering`, `DatabaseStateRecoveryPending`, `DatabaseStateRestoring`, `DatabaseStateSuspect`
- New type alias `RecoveryMode` with values `RecoveryModeBulkLogged`, `RecoveryModeFull`, `RecoveryModeSimple`
- New function `NewSQLServerDatabasesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLServerDatabasesClient, error)`
- New function `*SQLServerDatabasesClient.Create(context.Context, string, string, string, SQLServerDatabaseResource, *SQLServerDatabasesClientCreateOptions) (SQLServerDatabasesClientCreateResponse, error)`
- New function `*SQLServerDatabasesClient.Delete(context.Context, string, string, string, *SQLServerDatabasesClientDeleteOptions) (SQLServerDatabasesClientDeleteResponse, error)`
- New function `*SQLServerDatabasesClient.Get(context.Context, string, string, string, *SQLServerDatabasesClientGetOptions) (SQLServerDatabasesClientGetResponse, error)`
- New function `*SQLServerDatabasesClient.NewListPager(string, string, *SQLServerDatabasesClientListOptions) *runtime.Pager[SQLServerDatabasesClientListResponse]`
- New function `*SQLServerDatabasesClient.Update(context.Context, string, string, string, SQLServerDatabaseUpdate, *SQLServerDatabasesClientUpdateOptions) (SQLServerDatabasesClientUpdateResponse, error)`
- New struct `ArcSQLServerDatabaseListResult`
- New struct `SQLServerDatabaseResource`
- New struct `SQLServerDatabaseResourceProperties`
- New struct `SQLServerDatabaseResourcePropertiesBackupInformation`
- New struct `SQLServerDatabaseResourcePropertiesDatabaseOptions`
- New struct `SQLServerDatabaseUpdate`
- New struct `SQLServerDatabasesClient`
- New struct `SQLServerDatabasesClientListResponse`
- New field `Cores` in struct `SQLServerInstanceProperties`


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).