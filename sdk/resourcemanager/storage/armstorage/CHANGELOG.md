# Release History

## 2.0.0 (2022-11-09)
### Breaking Changes

- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New const `TargetSKUNameStandardRAGZRS`
- New const `MigrationStatusInProgress`
- New const `MigrationStatusSubmittedForConversion`
- New const `TargetSKUNameStandardGRS`
- New const `TargetSKUNameStandardGZRS`
- New const `ListEncryptionScopesIncludeAll`
- New const `MigrationStatusInvalid`
- New const `TargetSKUNameStandardRAGRS`
- New const `TargetSKUNameStandardZRS`
- New const `MigrationStatusComplete`
- New const `TargetSKUNameStandardLRS`
- New const `MigrationStatusFailed`
- New const `ListEncryptionScopesIncludeDisabled`
- New const `ListEncryptionScopesIncludeEnabled`
- New const `AccountMigrationNameDefault`
- New type alias `TargetSKUName`
- New type alias `MigrationStatus`
- New type alias `AccountMigrationName`
- New type alias `ListEncryptionScopesInclude`
- New function `*AccountMigrationsClient.BeginPut(context.Context, string, string, AccountMigrationName, AccountMigration, *AccountMigrationsClientBeginPutOptions) (*runtime.Poller[AccountMigrationsClientPutResponse], error)`
- New function `*AccountMigrationsClient.Get(context.Context, string, string, AccountMigrationName, *AccountMigrationsClientGetOptions) (AccountMigrationsClientGetResponse, error)`
- New function `PossibleMigrationStatusValues() []MigrationStatus`
- New function `PossibleAccountMigrationNameValues() []AccountMigrationName`
- New function `PossibleListEncryptionScopesIncludeValues() []ListEncryptionScopesInclude`
- New function `PossibleTargetSKUNameValues() []TargetSKUName`
- New function `NewAccountMigrationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AccountMigrationsClient, error)`
- New struct `AccountMigration`
- New struct `AccountMigrationProperties`
- New struct `AccountMigrationsClient`
- New struct `AccountMigrationsClientBeginPutOptions`
- New struct `AccountMigrationsClientGetOptions`
- New struct `AccountMigrationsClientGetResponse`
- New struct `AccountMigrationsClientPutResponse`
- New field `Filter` in struct `EncryptionScopesClientListOptions`
- New field `Include` in struct `EncryptionScopesClientListOptions`
- New field `Maxpagesize` in struct `EncryptionScopesClientListOptions`
- New field `FailoverType` in struct `AccountsClientBeginFailoverOptions`
- New field `TierToHot` in struct `ManagementPolicyBaseBlob`
- New field `TierToCold` in struct `ManagementPolicyBaseBlob`
- New field `AccountMigrationInProgress` in struct `AccountProperties`
- New field `IsSKUConversionBlocked` in struct `AccountProperties`
- New field `TierToCold` in struct `ManagementPolicySnapShot`
- New field `TierToHot` in struct `ManagementPolicySnapShot`
- New field `TierToCold` in struct `ManagementPolicyVersion`
- New field `TierToHot` in struct `ManagementPolicyVersion`


## 1.1.0 (2022-08-10)
### Features Added

- New const `DirectoryServiceOptionsAADKERB`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).