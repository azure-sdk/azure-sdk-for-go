# Release History

## 1.1.0 (2023-03-22)
### Features Added

- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewApplicationsClient() *ApplicationsClient`
- New function `*ClientFactory.NewClustersClient() *ClustersClient`
- New function `*ClientFactory.NewConfigurationsClient() *ConfigurationsClient`
- New function `*ClientFactory.NewExtensionsClient() *ExtensionsClient`
- New function `*ClientFactory.NewLocationsClient() *LocationsClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient`
- New function `*ClientFactory.NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient`
- New function `*ClientFactory.NewScriptActionsClient() *ScriptActionsClient`
- New function `*ClientFactory.NewScriptExecutionHistoryClient() *ScriptExecutionHistoryClient`
- New function `*ClientFactory.NewVirtualMachinesClient() *VirtualMachinesClient`
- New struct `ClientFactory`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).