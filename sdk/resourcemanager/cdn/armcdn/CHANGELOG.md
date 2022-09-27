# Release History

## 2.0.0-beta.1 (2022-09-27)
### Breaking Changes

- Type of `EndpointProperties.CustomDomains` has been changed from `[]*CustomDomain` to `[]*DeepCreatedCustomDomain`
- Function `*CustomDomainsClient.EnableCustomHTTPS` has been removed
- Function `*CustomDomainsClient.DisableCustomHTTPS` has been removed
- Struct `CustomDomainsClientDisableCustomHTTPSOptions` has been removed
- Struct `CustomDomainsClientEnableCustomHTTPSOptions` has been removed

### Features Added

- New const `ProfileResourceStateMigrated`
- New const `ProfileResourceStateMigrating`
- New const `CanMigrateDefaultSKUStandardAzureFrontDoor`
- New const `ProfileResourceStateAbortingMigration`
- New const `ProfileResourceStateCommittingMigration`
- New const `ProfileResourceStatePendingMigrationCommit`
- New const `CanMigrateDefaultSKUPremiumAzureFrontDoor`
- New type alias `CanMigrateDefaultSKU`
- New function `*CustomDomainsClient.BeginDisableCustomHTTPS(context.Context, string, string, string, string, *CustomDomainsClientBeginDisableCustomHTTPSOptions) (*runtime.Poller[CustomDomainsClientDisableCustomHTTPSResponse], error)`
- New function `*CustomDomainsClient.BeginEnableCustomHTTPS(context.Context, string, string, string, string, *CustomDomainsClientBeginEnableCustomHTTPSOptions) (*runtime.Poller[CustomDomainsClientEnableCustomHTTPSResponse], error)`
- New function `*AFDProfilesClient.ValidateSecret(context.Context, string, string, ValidateSecretInput, *AFDProfilesClientValidateSecretOptions) (AFDProfilesClientValidateSecretResponse, error)`
- New function `*ProfilesClient.CanMigrate(context.Context, string, CanMigrateParameters, *ProfilesClientCanMigrateOptions) (ProfilesClientCanMigrateResponse, error)`
- New function `*ProfilesClient.BeginMigrationCommit(context.Context, string, string, *ProfilesClientBeginMigrationCommitOptions) (*runtime.Poller[ProfilesClientMigrationCommitResponse], error)`
- New function `PossibleCanMigrateDefaultSKUValues() []CanMigrateDefaultSKU`
- New function `*ProfilesClient.BeginMigrate(context.Context, string, MigrationParameters, *ProfilesClientBeginMigrateOptions) (*runtime.Poller[ProfilesClientMigrateResponse], error)`
- New struct `AFDProfilesClientValidateSecretOptions`
- New struct `AFDProfilesClientValidateSecretResponse`
- New struct `CanMigrateParameters`
- New struct `CanMigrateResult`
- New struct `CustomDomainsClientBeginDisableCustomHTTPSOptions`
- New struct `CustomDomainsClientBeginEnableCustomHTTPSOptions`
- New struct `DeepCreatedCustomDomain`
- New struct `DeepCreatedCustomDomainProperties`
- New struct `MigrateResult`
- New struct `MigrationErrorType`
- New struct `MigrationErrorsListResponse`
- New struct `MigrationParameters`
- New struct `MigrationWebApplicationFirewallMapping`
- New struct `ProfilesClientBeginMigrateOptions`
- New struct `ProfilesClientBeginMigrationCommitOptions`
- New struct `ProfilesClientCanMigrateOptions`
- New struct `ProfilesClientCanMigrateResponse`
- New struct `ProfilesClientMigrateResponse`
- New struct `ProfilesClientMigrationCommitResponse`
- New field `ExtendedProperties` in struct `AFDDomainProperties`
- New field `ExtendedProperties` in struct `WebApplicationFirewallPolicyProperties`
- New field `ExtendedProperties` in struct `ProfileProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).