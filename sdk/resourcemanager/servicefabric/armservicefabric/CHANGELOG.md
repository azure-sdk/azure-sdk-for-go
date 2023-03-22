# Release History

## 1.1.0 (2023-03-22)
### Features Added

- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewApplicationTypeVersionsClient() *ApplicationTypeVersionsClient`
- New function `*ClientFactory.NewApplicationTypesClient() *ApplicationTypesClient`
- New function `*ClientFactory.NewApplicationsClient() *ApplicationsClient`
- New function `*ClientFactory.NewClusterVersionsClient() *ClusterVersionsClient`
- New function `*ClientFactory.NewClustersClient() *ClustersClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewServicesClient() *ServicesClient`
- New struct `ClientFactory`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).