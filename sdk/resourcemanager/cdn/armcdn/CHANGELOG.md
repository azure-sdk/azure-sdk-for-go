# Release History

## 2.0.0-beta.1 (2022-08-22)
### Breaking Changes

- Type of `ProfileProperties.ProvisioningState` has been changed from `*ProfileProvisioningState` to `*string`
- Type of `CustomDomainProperties.ProvisioningState` has been changed from `*CustomHTTPSProvisioningState` to `*string`
- Type of `EndpointProperties.ProvisioningState` has been changed from `*EndpointProvisioningState` to `*string`
- Type of `ResourceUsage.Unit` has been changed from `*ResourceUsageUnit` to `*string`
- Type of `OriginGroupProperties.ProvisioningState` has been changed from `*OriginGroupProvisioningState` to `*string`
- Type of `OriginProperties.ProvisioningState` has been changed from `*OriginProvisioningState` to `*string`
- Const `EndpointProvisioningStateCreating` has been removed
- Const `OriginProvisioningStateCreating` has been removed
- Const `OriginGroupProvisioningStateFailed` has been removed
- Const `OriginProvisioningStateFailed` has been removed
- Const `OriginGroupProvisioningStateCreating` has been removed
- Const `EndpointProvisioningStateSucceeded` has been removed
- Const `EndpointProvisioningStateFailed` has been removed
- Const `OriginGroupProvisioningStateDeleting` has been removed
- Const `OriginGroupProvisioningStateUpdating` has been removed
- Const `ProfileProvisioningStateFailed` has been removed
- Const `ProfileProvisioningStateDeleting` has been removed
- Const `EndpointProvisioningStateUpdating` has been removed
- Const `ResourceUsageUnitCount` has been removed
- Const `OriginGroupProvisioningStateSucceeded` has been removed
- Const `EndpointProvisioningStateDeleting` has been removed
- Const `OriginProvisioningStateSucceeded` has been removed
- Const `OriginProvisioningStateDeleting` has been removed
- Const `ProfileProvisioningStateSucceeded` has been removed
- Const `ProfileProvisioningStateCreating` has been removed
- Const `ProfileProvisioningStateUpdating` has been removed
- Const `OriginProvisioningStateUpdating` has been removed
- Type alias `EndpointProvisioningState` has been removed
- Type alias `ProfileProvisioningState` has been removed
- Type alias `OriginProvisioningState` has been removed
- Type alias `OriginGroupProvisioningState` has been removed
- Type alias `ResourceUsageUnit` has been removed
- Function `PossibleEndpointProvisioningStateValues` has been removed
- Function `PossibleProfileProvisioningStateValues` has been removed
- Function `*CustomDomainsClient.EnableCustomHTTPS` has been removed
- Function `PossibleOriginProvisioningStateValues` has been removed
- Function `PossibleResourceUsageUnitValues` has been removed
- Function `PossibleOriginGroupProvisioningStateValues` has been removed
- Function `*CustomDomainsClient.DisableCustomHTTPS` has been removed
- Struct `CustomDomainsClientDisableCustomHTTPSOptions` has been removed
- Struct `CustomDomainsClientEnableCustomHTTPSOptions` has been removed

### Features Added

- New const `ProfileResourceStateCommittingMigration`
- New const `ProfileResourceStateAbortingMigration`
- New const `ProfileResourceStatePendingMigrationCommit`
- New const `CanMigrateDefaultSKUPremiumAzureFrontDoor`
- New const `CanMigrateDefaultSKUStandardAzureFrontDoor`
- New const `ProfileResourceStateMigrated`
- New const `ProfileResourceStateMigrating`
- New type alias `CanMigrateDefaultSKU`
- New function `PossibleCanMigrateDefaultSKUValues() []CanMigrateDefaultSKU`
- New function `*CustomDomainsClient.BeginDisableCustomHTTPS(context.Context, string, string, string, string, *CustomDomainsClientBeginDisableCustomHTTPSOptions) (*runtime.Poller[CustomDomainsClientDisableCustomHTTPSResponse], error)`
- New function `*ProfilesClient.CanMigrate(context.Context, string, CanMigrateParameters, *ProfilesClientCanMigrateOptions) (ProfilesClientCanMigrateResponse, error)`
- New function `*ProfilesClient.BeginMigrate(context.Context, string, MigrationParameters, *ProfilesClientBeginMigrateOptions) (*runtime.Poller[ProfilesClientMigrateResponse], error)`
- New function `*ProfilesClient.BeginMigrationCommit(context.Context, string, string, *ProfilesClientBeginMigrationCommitOptions) (*runtime.Poller[ProfilesClientMigrationCommitResponse], error)`
- New function `*CustomDomainsClient.BeginEnableCustomHTTPS(context.Context, string, string, string, string, *CustomDomainsClientBeginEnableCustomHTTPSOptions) (*runtime.Poller[CustomDomainsClientEnableCustomHTTPSResponse], error)`
- New struct `CanMigrateParameters`
- New struct `CanMigrateResult`
- New struct `CustomDomainsClientBeginDisableCustomHTTPSOptions`
- New struct `CustomDomainsClientBeginEnableCustomHTTPSOptions`
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
- New field `ExtendedProperties` in struct `ProfileProperties`
- New field `ExtendedProperties` in struct `WebApplicationFirewallPolicyProperties`
- New field `ExtendedProperties` in struct `AFDDomainProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).