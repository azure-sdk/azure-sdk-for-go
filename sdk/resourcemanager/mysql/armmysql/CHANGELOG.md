# Release History

## 2.0.0 (2023-03-21)
### Breaking Changes

- Struct `CloudError` has been removed

### Features Added

- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewAdvisorsClient() *AdvisorsClient`
- New function `*ClientFactory.NewCheckNameAvailabilityClient() *CheckNameAvailabilityClient`
- New function `*ClientFactory.NewConfigurationsClient() *ConfigurationsClient`
- New function `*ClientFactory.NewDatabasesClient() *DatabasesClient`
- New function `*ClientFactory.NewFirewallRulesClient() *FirewallRulesClient`
- New function `*ClientFactory.NewLocationBasedPerformanceTierClient() *LocationBasedPerformanceTierClient`
- New function `*ClientFactory.NewLocationBasedRecommendedActionSessionsOperationStatusClient() *LocationBasedRecommendedActionSessionsOperationStatusClient`
- New function `*ClientFactory.NewLocationBasedRecommendedActionSessionsResultClient() *LocationBasedRecommendedActionSessionsResultClient`
- New function `*ClientFactory.NewLogFilesClient() *LogFilesClient`
- New function `*ClientFactory.NewManagementClient() *ManagementClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient`
- New function `*ClientFactory.NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient`
- New function `*ClientFactory.NewQueryTextsClient() *QueryTextsClient`
- New function `*ClientFactory.NewRecommendedActionsClient() *RecommendedActionsClient`
- New function `*ClientFactory.NewRecoverableServersClient() *RecoverableServersClient`
- New function `*ClientFactory.NewReplicasClient() *ReplicasClient`
- New function `*ClientFactory.NewServerAdministratorsClient() *ServerAdministratorsClient`
- New function `*ClientFactory.NewServerBasedPerformanceTierClient() *ServerBasedPerformanceTierClient`
- New function `*ClientFactory.NewServerKeysClient() *ServerKeysClient`
- New function `*ClientFactory.NewServerParametersClient() *ServerParametersClient`
- New function `*ClientFactory.NewServerSecurityAlertPoliciesClient() *ServerSecurityAlertPoliciesClient`
- New function `*ClientFactory.NewServersClient() *ServersClient`
- New function `*ClientFactory.NewTopQueryStatisticsClient() *TopQueryStatisticsClient`
- New function `*ClientFactory.NewVirtualNetworkRulesClient() *VirtualNetworkRulesClient`
- New function `*ClientFactory.NewWaitStatisticsClient() *WaitStatisticsClient`
- New struct `ClientFactory`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).