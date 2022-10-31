# Release History

## 0.6.0 (2022-10-31)
### Breaking Changes

- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New function `*ServerGroupClient.BeginPromoteReadReplica(context.Context, string, string, *ServerGroupClientBeginPromoteReadReplicaOptions) (*runtime.Poller[ServerGroupClientPromoteReadReplicaResponse], error)`
- New function `NewServerGroupClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ServerGroupClient, error)`
- New struct `ServerGroupClient`
- New struct `ServerGroupClientBeginPromoteReadReplicaOptions`
- New struct `ServerGroupClientPromoteReadReplicaResponse`


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).