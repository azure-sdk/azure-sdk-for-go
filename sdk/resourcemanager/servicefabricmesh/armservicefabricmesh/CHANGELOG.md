# Release History

## 0.6.0 (2023-03-29)
### Features Added

- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewApplicationClient() *ApplicationClient`
- New function `*ClientFactory.NewCodePackageClient() *CodePackageClient`
- New function `*ClientFactory.NewGatewayClient() *GatewayClient`
- New function `*ClientFactory.NewNetworkClient() *NetworkClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewSecretClient() *SecretClient`
- New function `*ClientFactory.NewSecretValueClient() *SecretValueClient`
- New function `*ClientFactory.NewServiceClient() *ServiceClient`
- New function `*ClientFactory.NewServiceReplicaClient() *ServiceReplicaClient`
- New function `*ClientFactory.NewVolumeClient() *VolumeClient`
- New struct `AvailableOperationDescriptionProperties`
- New struct `ClientFactory`
- New struct `Dimension`
- New struct `LogSpecification`
- New struct `MetricAvailability`
- New struct `MetricSpecification`
- New struct `ServiceSpecification`
- New field `Properties` in struct `OperationResult`


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).