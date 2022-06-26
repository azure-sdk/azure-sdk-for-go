# Release History

## 0.8.0 (2022-06-26)
### Breaking Changes

- Function `*DataCollectionEndpointsClient.Update` parameter(s) have been changed from `(context.Context, string, string, *DataCollectionEndpointsClientUpdateOptions)` to `(context.Context, string, string, ResourceForUpdate, *DataCollectionEndpointsClientUpdateOptions)`
- Function `*DataCollectionEndpointsClient.Create` parameter(s) have been changed from `(context.Context, string, string, *DataCollectionEndpointsClientCreateOptions)` to `(context.Context, string, string, DataCollectionEndpointResource, *DataCollectionEndpointsClientCreateOptions)`
- Function `*DataCollectionRulesClient.Create` parameter(s) have been changed from `(context.Context, string, string, *DataCollectionRulesClientCreateOptions)` to `(context.Context, string, string, DataCollectionRuleResource, *DataCollectionRulesClientCreateOptions)`
- Function `*DataCollectionRulesClient.Update` parameter(s) have been changed from `(context.Context, string, string, *DataCollectionRulesClientUpdateOptions)` to `(context.Context, string, string, ResourceForUpdate, *DataCollectionRulesClientUpdateOptions)`
- Function `*DataCollectionRuleAssociationsClient.Create` parameter(s) have been changed from `(context.Context, string, string, *DataCollectionRuleAssociationsClientCreateOptions)` to `(context.Context, string, string, DataCollectionRuleAssociationProxyOnlyResource, *DataCollectionRuleAssociationsClientCreateOptions)`
- Field `Body` of struct `DataCollectionEndpointsClientUpdateOptions` has been removed
- Field `Body` of struct `DataCollectionRulesClientUpdateOptions` has been removed
- Field `Body` of struct `DataCollectionRulesClientCreateOptions` has been removed
- Field `Body` of struct `DataCollectionEndpointsClientCreateOptions` has been removed
- Field `Body` of struct `DataCollectionRuleAssociationsClientCreateOptions` has been removed


## 0.7.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.7.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).