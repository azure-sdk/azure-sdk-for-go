# Release History

## 2.0.0 (2022-08-03)
### Breaking Changes

- Function `*TagRulesClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, *TagRulesClientCreateOrUpdateOptions)` to `(context.Context, string, string, string, MonitoringTagRules, *TagRulesClientCreateOrUpdateOptions)`
- Function `*MonitorsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, *MonitorsClientBeginUpdateOptions)` to `(context.Context, string, string, MonitorResourceUpdateParameters, *MonitorsClientBeginUpdateOptions)`
- Function `*SingleSignOnConfigurationsClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, *SingleSignOnConfigurationsClientBeginCreateOrUpdateOptions)` to `(context.Context, string, string, string, SingleSignOnResource, *SingleSignOnConfigurationsClientBeginCreateOrUpdateOptions)`
- Function `*MarketplaceAgreementsClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, *MarketplaceAgreementsClientCreateOrUpdateOptions)` to `(context.Context, AgreementResource, *MarketplaceAgreementsClientCreateOrUpdateOptions)`
- Function `*MonitorsClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, *MonitorsClientBeginCreateOptions)` to `(context.Context, string, string, MonitorResource, *MonitorsClientBeginCreateOptions)`
- Field `Body` of struct `MonitorsClientBeginUpdateOptions` has been removed
- Field `Body` of struct `SingleSignOnConfigurationsClientBeginCreateOrUpdateOptions` has been removed
- Field `Body` of struct `MonitorsClientBeginCreateOptions` has been removed
- Field `Body` of struct `MarketplaceAgreementsClientCreateOrUpdateOptions` has been removed
- Field `Body` of struct `TagRulesClientCreateOrUpdateOptions` has been removed


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).