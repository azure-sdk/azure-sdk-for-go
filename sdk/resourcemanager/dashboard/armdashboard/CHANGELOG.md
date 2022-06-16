# Release History

## 0.4.0 (2022-06-16)
### Breaking Changes

- Type of `SystemData.LastModifiedByType` has been changed from `*LastModifiedByType` to `*CreatedByType`
- Type of `OperationListResult.Value` has been changed from `[]*OperationResult` to `[]*Operation`
- Type of `ManagedGrafanaUpdateParameters.Identity` has been changed from `*ManagedIdentity` to `*ManagedServiceIdentity`
- Type of `ManagedGrafana.Identity` has been changed from `*ManagedIdentity` to `*ManagedServiceIdentity`
- Const `LastModifiedByTypeManagedIdentity` has been removed
- Const `LastModifiedByTypeUser` has been removed
- Const `LastModifiedByTypeApplication` has been removed
- Const `IdentityTypeNone` has been removed
- Const `LastModifiedByTypeKey` has been removed
- Const `IdentityTypeSystemAssigned` has been removed
- Function `PossibleIdentityTypeValues` has been removed
- Function `PossibleLastModifiedByTypeValues` has been removed
- Struct `ManagedIdentity` has been removed
- Struct `OperationResult` has been removed

### Features Added

- New field `APIKey` in struct `ManagedGrafanaProperties`
- New field `OutboundIPs` in struct `ManagedGrafanaProperties`
- New field `PrivateEndpointConnections` in struct `ManagedGrafanaProperties`
- New field `PublicNetworkAccess` in struct `ManagedGrafanaProperties`
- New field `DeterministicOutboundIP` in struct `ManagedGrafanaProperties`
- New field `Properties` in struct `ManagedGrafanaUpdateParameters`


## 0.3.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.3.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).