# Release History

## 0.8.0 (2023-06-04)
### Breaking Changes

- Function `*ExtendedSQLPoolBlobAuditingPoliciesClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, ExtendedSQLPoolBlobAuditingPolicy, *ExtendedSQLPoolBlobAuditingPoliciesClientCreateOrUpdateOptions)` to `(context.Context, string, string, string, BlobAuditingPolicyName, ExtendedSQLPoolBlobAuditingPolicy, *ExtendedSQLPoolBlobAuditingPoliciesClientCreateOrUpdateOptions)`
- Function `*ExtendedSQLPoolBlobAuditingPoliciesClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *ExtendedSQLPoolBlobAuditingPoliciesClientGetOptions)` to `(context.Context, string, string, string, BlobAuditingPolicyName, *ExtendedSQLPoolBlobAuditingPoliciesClientGetOptions)`
- Function `*SQLPoolBlobAuditingPoliciesClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, SQLPoolBlobAuditingPolicy, *SQLPoolBlobAuditingPoliciesClientCreateOrUpdateOptions)` to `(context.Context, string, string, string, BlobAuditingPolicyName, SQLPoolBlobAuditingPolicy, *SQLPoolBlobAuditingPoliciesClientCreateOrUpdateOptions)`
- Function `*SQLPoolBlobAuditingPoliciesClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *SQLPoolBlobAuditingPoliciesClientGetOptions)` to `(context.Context, string, string, string, BlobAuditingPolicyName, *SQLPoolBlobAuditingPoliciesClientGetOptions)`
- Type of `AzureADOnlyAuthenticationProperties.State` has been changed from `*StateValue` to `*AuthenticationState`
- Enum `StateValue` has been removed
- Struct `DedicatedSQLminimalTLSSettingsPatchInfo` has been removed
- Field `IsDevopsAuditEnabled` of struct `ExtendedServerBlobAuditingPolicyProperties` has been removed
- Field `IPFirewallRuleInfo` of struct `IPFirewallRulesClientDeleteResponse` has been removed

### Features Added

- Type of `WorkspaceProperties.ExtraProperties` has been changed from `map[string]any` to `any`
- New value `StateMigrated` added to enum type `State`
- New enum type `AuthenticationState` with values `AuthenticationStateConsistent`, `AuthenticationStateInConsistent`, `AuthenticationStateUpdating`
- New enum type `DatawarehousequeriesResourceClass` with values `DatawarehousequeriesResourceClassLargerc`, `DatawarehousequeriesResourceClassMediumrc`, `DatawarehousequeriesResourceClassSmallrc`, `DatawarehousequeriesResourceClassXlargerc`
- New enum type `MigrationClusterRole` with values `MigrationClusterRoleDestination`, `MigrationClusterRoleSource`
- New enum type `StepsDistributionType` with values `StepsDistributionTypeAllAzureEndpoints`, `StepsDistributionTypeAllComputeNodes`, `StepsDistributionTypeAllDistributions`, `StepsDistributionTypeAllNodes`, `StepsDistributionTypeComputeNode`, `StepsDistributionTypeDistribution`, `StepsDistributionTypeSubsetDistributions`, `StepsDistributionTypeSubsetNodes`, `StepsDistributionTypeUnspecified`
- New enum type `StepsLocationType` with values `StepsLocationTypeCompute`, `StepsLocationTypeControl`, `StepsLocationTypeDMS`
- New enum type `StepsOperationType` with values `StepsOperationTypeBackupAzureDatabaseOperation`, `StepsOperationTypeBroadcastMoveOperation`, `StepsOperationTypeCheckEncryptionStatusOperation`, `StepsOperationTypeCopyOperation`, `StepsOperationTypeCreateDiagnosticsOperation`, `StepsOperationTypeDSQLCallBackOperation`, `StepsOperationTypeDbccShowStatisticsOperation`, `StepsOperationTypeDistributeReplicatedTableMoveOperation`, `StepsOperationTypeDistributedExchangeOperation`, `StepsOperationTypeDropDiagnosticsNotifyOperation`, `StepsOperationTypeDropDiagnosticsSynchronizeOperation`, `StepsOperationTypeExternalBroadcastMove`, `StepsOperationTypeExternalExportControlMove`, `StepsOperationTypeExternalExportControlOperation`, `StepsOperationTypeExternalExportDistributedMove`, `StepsOperationTypeExternalExportDistributedOperation`, `StepsOperationTypeExternalExportReplicatedMove`, `StepsOperationTypeExternalExportReplicatedOperation`, `StepsOperationTypeExternalShuffleOperation`, `StepsOperationTypeExternalStatisticsOperation`, `StepsOperationTypeHadoopBroadcastOperation`, `StepsOperationTypeHadoopRoundRobinOperation`, `StepsOperationTypeHadoopShuffleOperation`, `StepsOperationTypeInsertBulkOperation`, `StepsOperationTypeMetaDataCreateOperation`, `StepsOperationTypeMoveOperation`, `StepsOperationTypeOnOperation`, `StepsOperationTypePartitionMoveOperation`, `StepsOperationTypeRandomIDOperation`, `StepsOperationTypeReconfigureOperation`, `StepsOperationTypeReturnOperation`, `StepsOperationTypeRoundRobinMoveOperation`, `StepsOperationTypeSetIdentityInsertOperation`, `StepsOperationTypeSetOptionsOperation`, `StepsOperationTypeShuffleMoveOperation`, `StepsOperationTypeSingleSourceRoundRobinMoveOperation`, `StepsOperationTypeTempTablePropertiesOperation`, `StepsOperationTypeTrimMoveOperation`
- New enum type `WorkspaceProvisioningState` with values `WorkspaceProvisioningStateDeleteError`, `WorkspaceProvisioningStateDeleting`, `WorkspaceProvisioningStateFailed`, `WorkspaceProvisioningStateProvisioning`, `WorkspaceProvisioningStateSucceeded`, `WorkspaceProvisioningStateUpdateError`
- New function `*ClientFactory.NewKustoPoolDatabaseClient() *KustoPoolDatabaseClient`
- New function `*ClientFactory.NewSQLPoolsDatawareHouseQueriesClient() *SQLPoolsDatawareHouseQueriesClient`
- New function `*ClientFactory.NewSQLPoolsDatawareHouseQueriesStepsClient() *SQLPoolsDatawareHouseQueriesStepsClient`
- New function `*ClientFactory.NewScopePoolsListByWorkspaceClient() *ScopePoolsListByWorkspaceClient`
- New function `*ClientFactory.NewWorkspaceCheckDefaultStorageAccountStatusClient() *WorkspaceCheckDefaultStorageAccountStatusClient`
- New function `*ClientFactory.NewWorkspaceTrustedServiceByPassConfigurationClient() *WorkspaceTrustedServiceByPassConfigurationClient`
- New function `NewKustoPoolDatabaseClient(string, azcore.TokenCredential, *arm.ClientOptions) (*KustoPoolDatabaseClient, error)`
- New function `*KustoPoolDatabaseClient.InviteFollower(context.Context, string, string, string, string, DatabaseInviteFollowerRequest, *KustoPoolDatabaseClientInviteFollowerOptions) (KustoPoolDatabaseClientInviteFollowerResponse, error)`
- New function `*KustoPoolsClient.BeginMigrate(context.Context, string, string, string, KustoPoolMigrateRequest, *KustoPoolsClientBeginMigrateOptions) (*runtime.Poller[KustoPoolsClientMigrateResponse], error)`
- New function `NewSQLPoolsDatawareHouseQueriesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLPoolsDatawareHouseQueriesClient, error)`
- New function `*SQLPoolsDatawareHouseQueriesClient.Get(context.Context, string, string, string, *SQLPoolsDatawareHouseQueriesClientGetOptions) (SQLPoolsDatawareHouseQueriesClientGetResponse, error)`
- New function `NewSQLPoolsDatawareHouseQueriesStepsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SQLPoolsDatawareHouseQueriesStepsClient, error)`
- New function `*SQLPoolsDatawareHouseQueriesStepsClient.List(context.Context, string, string, string, string, *SQLPoolsDatawareHouseQueriesStepsClientListOptions) (SQLPoolsDatawareHouseQueriesStepsClientListResponse, error)`
- New function `NewScopePoolsListByWorkspaceClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ScopePoolsListByWorkspaceClient, error)`
- New function `*ScopePoolsListByWorkspaceClient.NewGetPager(string, string, *ScopePoolsListByWorkspaceClientGetOptions) *runtime.Pager[ScopePoolsListByWorkspaceClientGetResponse]`
- New function `NewWorkspaceCheckDefaultStorageAccountStatusClient(string, azcore.TokenCredential, *arm.ClientOptions) (*WorkspaceCheckDefaultStorageAccountStatusClient, error)`
- New function `*WorkspaceCheckDefaultStorageAccountStatusClient.Create(context.Context, string, string, any, *WorkspaceCheckDefaultStorageAccountStatusClientCreateOptions) (WorkspaceCheckDefaultStorageAccountStatusClientCreateResponse, error)`
- New function `NewWorkspaceTrustedServiceByPassConfigurationClient(string, azcore.TokenCredential, *arm.ClientOptions) (*WorkspaceTrustedServiceByPassConfigurationClient, error)`
- New function `*WorkspaceTrustedServiceByPassConfigurationClient.CreateOrUpdate(context.Context, string, string, any, *WorkspaceTrustedServiceByPassConfigurationClientCreateOrUpdateOptions) (WorkspaceTrustedServiceByPassConfigurationClientCreateOrUpdateResponse, error)`
- New struct `CheckDefaultStorageAccountStatus`
- New struct `CheckDefaultStorageAccountStatusProperties`
- New struct `DataWareHouseQueriesProperties`
- New struct `DataWarehouseQueries`
- New struct `DataWarehouseQueriesList`
- New struct `DataWarehouseQueriesSteps`
- New struct `DataWarehouseQueriesStepsList`
- New struct `DataWarehouseQueriesStepsProperties`
- New struct `DatabaseInviteFollowerRequest`
- New struct `DatabaseInviteFollowerResult`
- New struct `KustoPoolMigrateRequest`
- New struct `MigrationClusterProperties`
- New struct `ScopePoolProperties`
- New struct `ScopePoolPropertiesExtendedProperties`
- New struct `ScopePoolPropertiesExtendedPropertiesPoolLimitAndPolicies`
- New struct `ScopePools`
- New struct `ScopePoolsListResult`
- New struct `TrustedServiceByPassConfiguration`
- New struct `TrustedServiceByPassConfigurationProperties`
- New struct `TrustedServiceByPassConfigurationPropertiesTrustedServiceBypassConfigurationInfo`
- New field `Interface` in struct `IPFirewallRulesClientDeleteResponse`
- New field `MigrationCluster` in struct `KustoPoolProperties`
- New field `OSType`, `TargetFramework` in struct `SelfHostedIntegrationRuntimeStatusTypeProperties`
- New field `FunctionsToExclude`, `FunctionsToInclude` in struct `TableLevelSharingProperties`


## 0.7.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 0.7.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.6.0 (2023-02-24)
### Breaking Changes

- Type of `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentity.DesiredState` has been changed from `*ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityDesiredState` to `*DesiredState`
- Type of `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentity.ActualState` has been changed from `*ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualState` to `*ActualState`
- Type alias `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityActualState` has been removed
- Type alias `ManagedIdentitySQLControlSettingsModelPropertiesGrantSQLControlToManagedIdentityDesiredState` has been removed
- Operation `*SQLPoolOperationResultsClient.GetLocationHeaderResult` has been changed to LRO, use `*SQLPoolOperationResultsClient.BeginGetLocationHeaderResult` instead.
- Operation `*SQLPoolsClient.Update` has been changed to LRO, use `*SQLPoolsClient.BeginUpdate` instead.
- Field `Interface` of struct `BigDataPoolsClientDeleteResponse` has been removed
- Field `Interface` of struct `IPFirewallRulesClientDeleteResponse` has been removed
- Field `Interface` of struct `SQLPoolsClientDeleteResponse` has been removed
- Field `Interface` of struct `SQLPoolsClientPauseResponse` has been removed
- Field `Interface` of struct `SQLPoolsClientResumeResponse` has been removed
- Field `Interface` of struct `WorkspacesClientDeleteResponse` has been removed

### Features Added

- New type alias `ActualState` with values `ActualStateDisabled`, `ActualStateDisabling`, `ActualStateEnabled`, `ActualStateEnabling`, `ActualStateUnknown`
- New type alias `DesiredState` with values `DesiredStateDisabled`, `DesiredStateEnabled`
- New type alias `WorkspaceStatus` with values `WorkspaceStatusFailed`, `WorkspaceStatusInProgress`, `WorkspaceStatusSucceeded`
- New function `NewGetClient(string, azcore.TokenCredential, *arm.ClientOptions) (*GetClient, error)`
- New function `*GetClient.IntegrationRuntimeEnableInteractivequery(context.Context, string, string, string, string, *GetClientIntegrationRuntimeEnableInteractivequeryOptions) (GetClientIntegrationRuntimeEnableInteractivequeryResponse, error)`
- New function `*GetClient.IntegrationRuntimeStart(context.Context, string, string, string, string, *GetClientIntegrationRuntimeStartOptions) (GetClientIntegrationRuntimeStartResponse, error)`
- New function `*GetClient.IntegrationRuntimeStop(context.Context, string, string, string, string, *GetClientIntegrationRuntimeStopOptions) (GetClientIntegrationRuntimeStopResponse, error)`
- New function `NewKustoPoolPrivateLinkResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*KustoPoolPrivateLinkResourcesClient, error)`
- New function `*KustoPoolPrivateLinkResourcesClient.NewListPager(string, string, string, *KustoPoolPrivateLinkResourcesClientListOptions) *runtime.Pager[KustoPoolPrivateLinkResourcesClientListResponse]`
- New struct `ErrorResponseAutoGenerated`
- New struct `GetClient`
- New struct `IntegrationRuntimeEnableinteractivequery`
- New struct `IntegrationRuntimeOperationStatus`
- New struct `IntegrationRuntimeStopOperationStatus`
- New struct `KustoPoolPrivateLinkResources`
- New struct `KustoPoolPrivateLinkResourcesClient`
- New struct `KustoPoolPrivateLinkResourcesClientListResponse`
- New struct `PrivateLinkResources`
- New struct `PrivateLinkResourcesProperties`
- New field `IsAutotuneEnabled` in struct `BigDataPoolResourceProperties`
- New anonymous field `BigDataPoolResourceInfo` in struct `BigDataPoolsClientDeleteResponse`
- New field `AADObjectID` in struct `ClusterPrincipalProperties`
- New field `AADObjectID` in struct `DatabasePrincipalProperties`
- New anonymous field `IPFirewallRuleInfo` in struct `IPFirewallRulesClientDeleteResponse`
- New anonymous field `SQLPool` in struct `SQLPoolOperationResultsClientGetLocationHeaderResultResponse`
- New anonymous field `SQLPool` in struct `SQLPoolsClientDeleteResponse`
- New anonymous field `SQLPool` in struct `SQLPoolsClientPauseResponse`
- New anonymous field `SQLPool` in struct `SQLPoolsClientResumeResponse`
- New field `ConfigMergeRule` in struct `SparkConfigurationInfo`
- New anonymous field `Workspace` in struct `WorkspacesClientDeleteResponse`


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).