# Release History

## 1.1.0-beta.1 (2022-08-25)
### Features Added

- New const `MoveOperationModeMove`
- New const `MoveOperationModeCopy`
- New const `DtcNameCurrent`
- New type alias `MoveOperationMode`
- New type alias `DtcName`
- New function `NewManagedInstanceAdvancedThreatProtectionSettingsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagedInstanceAdvancedThreatProtectionSettingsClient, error)`
- New function `*ManagedServerDNSAliasesClient.BeginDelete(context.Context, string, string, string, *ManagedServerDNSAliasesClientBeginDeleteOptions) (*runtime.Poller[ManagedServerDNSAliasesClientDeleteResponse], error)`
- New function `*ManagedDatabaseMoveOperationsClient.Get(context.Context, string, string, string, *ManagedDatabaseMoveOperationsClientGetOptions) (ManagedDatabaseMoveOperationsClientGetResponse, error)`
- New function `*ManagedDatabaseMoveOperationsClient.NewListByLocationPager(string, string, *ManagedDatabaseMoveOperationsClientListByLocationOptions) *runtime.Pager[ManagedDatabaseMoveOperationsClientListByLocationResponse]`
- New function `*ManagedInstanceDtcsClient.BeginCreateOrUpdate(context.Context, string, string, DtcName, ManagedInstanceDtc, *ManagedInstanceDtcsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedInstanceDtcsClientCreateOrUpdateResponse], error)`
- New function `*ManagedInstanceAdvancedThreatProtectionSettingsClient.BeginCreateOrUpdate(context.Context, string, string, AdvancedThreatProtectionName, ManagedInstanceAdvancedThreatProtection, *ManagedInstanceAdvancedThreatProtectionSettingsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedInstanceAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse], error)`
- New function `*ManagedDatabaseAdvancedThreatProtectionSettingsClient.NewListByDatabasePager(string, string, string, *ManagedDatabaseAdvancedThreatProtectionSettingsClientListByDatabaseOptions) *runtime.Pager[ManagedDatabaseAdvancedThreatProtectionSettingsClientListByDatabaseResponse]`
- New function `NewManagedDatabaseAdvancedThreatProtectionSettingsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagedDatabaseAdvancedThreatProtectionSettingsClient, error)`
- New function `*ManagedServerDNSAliasesClient.NewListByManagedInstancePager(string, string, *ManagedServerDNSAliasesClientListByManagedInstanceOptions) *runtime.Pager[ManagedServerDNSAliasesClientListByManagedInstanceResponse]`
- New function `NewManagedServerDNSAliasesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagedServerDNSAliasesClient, error)`
- New function `PossibleMoveOperationModeValues() []MoveOperationMode`
- New function `*ManagedInstanceDtcsClient.Get(context.Context, string, string, DtcName, *ManagedInstanceDtcsClientGetOptions) (ManagedInstanceDtcsClientGetResponse, error)`
- New function `NewManagedInstanceDtcsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagedInstanceDtcsClient, error)`
- New function `NewManagedDatabaseMoveOperationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagedDatabaseMoveOperationsClient, error)`
- New function `*ManagedServerDNSAliasesClient.BeginAcquire(context.Context, string, string, string, ManagedServerDNSAliasAcquisition, *ManagedServerDNSAliasesClientBeginAcquireOptions) (*runtime.Poller[ManagedServerDNSAliasesClientAcquireResponse], error)`
- New function `*ManagedDatabaseAdvancedThreatProtectionSettingsClient.CreateOrUpdate(context.Context, string, string, string, AdvancedThreatProtectionName, ManagedDatabaseAdvancedThreatProtection, *ManagedDatabaseAdvancedThreatProtectionSettingsClientCreateOrUpdateOptions) (ManagedDatabaseAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse, error)`
- New function `PossibleDtcNameValues() []DtcName`
- New function `*ManagedDatabaseAdvancedThreatProtectionSettingsClient.Get(context.Context, string, string, string, AdvancedThreatProtectionName, *ManagedDatabaseAdvancedThreatProtectionSettingsClientGetOptions) (ManagedDatabaseAdvancedThreatProtectionSettingsClientGetResponse, error)`
- New function `*ManagedServerDNSAliasesClient.Get(context.Context, string, string, string, *ManagedServerDNSAliasesClientGetOptions) (ManagedServerDNSAliasesClientGetResponse, error)`
- New function `*ManagedInstanceAdvancedThreatProtectionSettingsClient.NewListByInstancePager(string, string, *ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceOptions) *runtime.Pager[ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceResponse]`
- New function `*ManagedInstanceAdvancedThreatProtectionSettingsClient.Get(context.Context, string, string, AdvancedThreatProtectionName, *ManagedInstanceAdvancedThreatProtectionSettingsClientGetOptions) (ManagedInstanceAdvancedThreatProtectionSettingsClientGetResponse, error)`
- New function `*ManagedServerDNSAliasesClient.BeginCreateOrUpdate(context.Context, string, string, string, ManagedServerDNSAliasCreation, *ManagedServerDNSAliasesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedServerDNSAliasesClientCreateOrUpdateResponse], error)`
- New function `*ManagedInstanceDtcsClient.NewListByManagedInstancePager(string, string, *ManagedInstanceDtcsClientListByManagedInstanceOptions) *runtime.Pager[ManagedInstanceDtcsClientListByManagedInstanceResponse]`
- New struct `ManagedDatabaseAdvancedThreatProtection`
- New struct `ManagedDatabaseAdvancedThreatProtectionListResult`
- New struct `ManagedDatabaseAdvancedThreatProtectionSettingsClient`
- New struct `ManagedDatabaseAdvancedThreatProtectionSettingsClientCreateOrUpdateOptions`
- New struct `ManagedDatabaseAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse`
- New struct `ManagedDatabaseAdvancedThreatProtectionSettingsClientGetOptions`
- New struct `ManagedDatabaseAdvancedThreatProtectionSettingsClientGetResponse`
- New struct `ManagedDatabaseAdvancedThreatProtectionSettingsClientListByDatabaseOptions`
- New struct `ManagedDatabaseAdvancedThreatProtectionSettingsClientListByDatabaseResponse`
- New struct `ManagedDatabaseMoveOperationListResult`
- New struct `ManagedDatabaseMoveOperationResult`
- New struct `ManagedDatabaseMoveOperationResultProperties`
- New struct `ManagedDatabaseMoveOperationsClient`
- New struct `ManagedDatabaseMoveOperationsClientGetOptions`
- New struct `ManagedDatabaseMoveOperationsClientGetResponse`
- New struct `ManagedDatabaseMoveOperationsClientListByLocationOptions`
- New struct `ManagedDatabaseMoveOperationsClientListByLocationResponse`
- New struct `ManagedInstanceAdvancedThreatProtection`
- New struct `ManagedInstanceAdvancedThreatProtectionListResult`
- New struct `ManagedInstanceAdvancedThreatProtectionSettingsClient`
- New struct `ManagedInstanceAdvancedThreatProtectionSettingsClientBeginCreateOrUpdateOptions`
- New struct `ManagedInstanceAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse`
- New struct `ManagedInstanceAdvancedThreatProtectionSettingsClientGetOptions`
- New struct `ManagedInstanceAdvancedThreatProtectionSettingsClientGetResponse`
- New struct `ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceOptions`
- New struct `ManagedInstanceAdvancedThreatProtectionSettingsClientListByInstanceResponse`
- New struct `ManagedInstanceDtc`
- New struct `ManagedInstanceDtcListResult`
- New struct `ManagedInstanceDtcProperties`
- New struct `ManagedInstanceDtcSecuritySettings`
- New struct `ManagedInstanceDtcTransactionManagerCommunicationSettings`
- New struct `ManagedInstanceDtcsClient`
- New struct `ManagedInstanceDtcsClientBeginCreateOrUpdateOptions`
- New struct `ManagedInstanceDtcsClientCreateOrUpdateResponse`
- New struct `ManagedInstanceDtcsClientGetOptions`
- New struct `ManagedInstanceDtcsClientGetResponse`
- New struct `ManagedInstanceDtcsClientListByManagedInstanceOptions`
- New struct `ManagedInstanceDtcsClientListByManagedInstanceResponse`
- New struct `ManagedServerDNSAlias`
- New struct `ManagedServerDNSAliasAcquisition`
- New struct `ManagedServerDNSAliasCreation`
- New struct `ManagedServerDNSAliasListResult`
- New struct `ManagedServerDNSAliasProperties`
- New struct `ManagedServerDNSAliasesClient`
- New struct `ManagedServerDNSAliasesClientAcquireResponse`
- New struct `ManagedServerDNSAliasesClientBeginAcquireOptions`
- New struct `ManagedServerDNSAliasesClientBeginCreateOrUpdateOptions`
- New struct `ManagedServerDNSAliasesClientBeginDeleteOptions`
- New struct `ManagedServerDNSAliasesClientCreateOrUpdateResponse`
- New struct `ManagedServerDNSAliasesClientDeleteResponse`
- New struct `ManagedServerDNSAliasesClientGetOptions`
- New struct `ManagedServerDNSAliasesClientGetResponse`
- New struct `ManagedServerDNSAliasesClientListByManagedInstanceOptions`
- New struct `ManagedServerDNSAliasesClientListByManagedInstanceResponse`


## 1.0.0 (2022-06-02)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).