# Release History

## 0.7.0 (2022-05-31)
### Breaking Changes

- Function `*DataMaskingPoliciesClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *DataMaskingPoliciesClientGetOptions)` to `(context.Context, string, string, string, DataMaskingPolicyName, *DataMaskingPoliciesClientGetOptions)`
- Function `*DataMaskingPoliciesClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, DataMaskingPolicy, *DataMaskingPoliciesClientCreateOrUpdateOptions)` to `(context.Context, string, string, string, DataMaskingPolicyName, DataMaskingPolicy, *DataMaskingPoliciesClientCreateOrUpdateOptions)`
- Function `*DataMaskingRulesClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, DataMaskingRule, *DataMaskingRulesClientCreateOrUpdateOptions)` to `(context.Context, string, string, string, DataMaskingPolicyName, string, DataMaskingRule, *DataMaskingRulesClientCreateOrUpdateOptions)`
- Function `*DataMaskingRulesClient.NewListByDatabasePager` parameter(s) have been changed from `(string, string, string, *DataMaskingRulesClientListByDatabaseOptions)` to `(string, string, string, DataMaskingPolicyName, *DataMaskingRulesClientListByDatabaseOptions)`
- Function `*GeoBackupPoliciesClient.NewListByDatabasePager` has been removed
- Struct `GeoBackupPoliciesClientListByDatabaseOptions` has been removed
- Struct `GeoBackupPoliciesClientListByDatabaseResponse` has been removed
- Field `IsManagedIdentityInUse` of struct `ExtendedDatabaseBlobAuditingPolicyProperties` has been removed
- Field `IsManagedIdentityInUse` of struct `ServerBlobAuditingPolicyProperties` has been removed
- Field `CurrentValue` of struct `ServerUsage` has been removed
- Field `DisplayName` of struct `ServerUsage` has been removed
- Field `Limit` of struct `ServerUsage` has been removed
- Field `NextResetTime` of struct `ServerUsage` has been removed
- Field `ResourceName` of struct `ServerUsage` has been removed
- Field `Unit` of struct `ServerUsage` has been removed
- Field `IsManagedIdentityInUse` of struct `ExtendedServerBlobAuditingPolicyProperties` has been removed
- Field `IsManagedIdentityInUse` of struct `DatabaseBlobAuditingPolicyProperties` has been removed

### Features Added

- New const `DtcNameCurrent`
- New const `DataMaskingPolicyNameDefault`
- New function `*ManagedDatabaseSensitivityLabelsClient.NewListByDatabasePager(string, string, string, *ManagedDatabaseSensitivityLabelsClientListByDatabaseOptions) *runtime.Pager[ManagedDatabaseSensitivityLabelsClientListByDatabaseResponse]`
- New function `*GeoBackupPoliciesClient.NewListPager(string, string, string, *GeoBackupPoliciesClientListOptions) *runtime.Pager[GeoBackupPoliciesClientListResponse]`
- New function `PossibleDataMaskingPolicyNameValues() []DataMaskingPolicyName`
- New function `*SensitivityLabelsClient.NewListByDatabasePager(string, string, string, *SensitivityLabelsClientListByDatabaseOptions) *runtime.Pager[SensitivityLabelsClientListByDatabaseResponse]`
- New function `PossibleDtcNameValues() []DtcName`
- New struct `GeoBackupPoliciesClientListOptions`
- New struct `GeoBackupPoliciesClientListResponse`
- New struct `ManagedDatabaseSensitivityLabelsClientListByDatabaseOptions`
- New struct `ManagedDatabaseSensitivityLabelsClientListByDatabaseResponse`
- New struct `ManagedInstanceDtc`
- New struct `ManagedInstanceDtcListResult`
- New struct `ManagedInstanceDtcProperties`
- New struct `ManagedInstanceDtcSecuritySettings`
- New struct `ManagedInstanceDtcTransactionManagerCommunicationSettings`
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
- New struct `SensitivityLabelsClientListByDatabaseOptions`
- New struct `SensitivityLabelsClientListByDatabaseResponse`
- New struct `ServerUsageProperties`
- New field `NextLink` in struct `DataMaskingRuleListResult`
- New field `NextLink` in struct `GeoBackupPolicyListResult`
- New field `Type` in struct `ServerUsage`
- New field `Properties` in struct `ServerUsage`
- New field `ID` in struct `ServerUsage`
- New field `NextLink` in struct `ServerUsageListResult`
- New field `NextLink` in struct `RecoverableDatabaseListResult`


## 0.6.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).