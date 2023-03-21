# Release History

## 0.2.0 (2023-03-21)
### Breaking Changes

- Operation `*PrivateEndpointConnectionsClient.BeginCreateOrUpdate` has been changed to non-LRO, use `*PrivateEndpointConnectionsClient.CreateOrUpdate` instead.
- Operation `*PrivateEndpointConnectionsClient.BeginDelete` has been changed to non-LRO, use `*PrivateEndpointConnectionsClient.Delete` instead.

### Features Added

- New value `EnterprisePolicyKindIdentity` added to enum type `EnterprisePolicyKind`
- New enum type `HealthStatus` with values `HealthStatusHealthy`, `HealthStatusUndetermined`, `HealthStatusUnhealthy`
- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewAccountsClient() *AccountsClient`
- New function `*ClientFactory.NewEnterprisePoliciesClient() *EnterprisePoliciesClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient`
- New function `*ClientFactory.NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient`
- New struct `ClientFactory`
- New field `SystemID` in struct `AccountProperties`
- New field `HealthStatus` in struct `Properties`
- New field `SystemID` in struct `Properties`


## 0.1.0 (2022-06-10)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.1.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).