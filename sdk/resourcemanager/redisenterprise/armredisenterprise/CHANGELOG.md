# Release History

## 1.1.0-beta.2 (2023-03-29)
### Breaking Changes

- Function `*DatabasesClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, string, Database, *DatabasesClientBeginCreateOptions)` to `(context.Context, string, string, string, DatabaseCreateOrUpdate, *DatabasesClientBeginCreateOptions)`
- Function `*DatabasesClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, string, DatabaseUpdate, *DatabasesClientBeginUpdateOptions)` to `(context.Context, string, string, string, any, *DatabasesClientBeginUpdateOptions)`
- Function `NewOperationsStatusClient` parameter(s) have been changed from `(string, azcore.TokenCredential, *arm.ClientOptions)` to `(azcore.TokenCredential, *arm.ClientOptions)`
- Function `*OperationsStatusClient.Get` parameter(s) have been changed from `(context.Context, string, string, *OperationsStatusClientGetOptions)` to `(context.Context, string, string, string, *OperationsStatusClientGetOptions)`
- Type of `ClusterUpdate.Properties` has been changed from `*ClusterProperties` to `*ClusterUpdateProperties`
- Type of `ClusterUpdate.SKU` has been changed from `*SKU` to `*SKUUpdate`
- Type of `DatabaseProperties.ClientProtocol` has been changed from `*Protocol` to `*ClientProtocol`
- Type of `RegenerateKeyParameters.KeyType` has been changed from `*AccessKeyType` to `*KeyType`
- Const `AofFrequencyOneS` from type alias `AofFrequency` has been removed
- Const `ProvisioningStateCreating` from type alias `ProvisioningState` has been removed
- Const `RdbFrequencyOneH`, `RdbFrequencySixH`, `RdbFrequencyTwelveH` from type alias `RdbFrequency` has been removed
- Const `TLSVersionOne0`, `TLSVersionOne1`, `TLSVersionOne2` from type alias `TLSVersion` has been removed
- Type alias `AccessKeyType` has been removed
- Type alias `ClusteringPolicy` has been removed
- Type alias `CmkIdentityType` has been removed
- Type alias `LinkState` has been removed
- Type alias `ManagedServiceIdentityType` has been removed
- Type alias `Protocol` has been removed
- Function `*Client.NewListPager` has been removed
- Function `*DatabasesClient.BeginExport` has been removed
- Function `*DatabasesClient.BeginFlush` has been removed
- Function `*DatabasesClient.BeginImport` has been removed
- Function `*PrivateEndpointConnectionsClient.NewListPager` has been removed
- Function `*PrivateEndpointConnectionsClient.BeginPut` has been removed
- Function `*SKUsClient.NewListPager` has been removed
- Operation `*PrivateEndpointConnectionsClient.Delete` has been changed to LRO, use `*PrivateEndpointConnectionsClient.BeginDelete` instead.
- Struct `Capability` has been removed
- Struct `ClusterList` has been removed
- Struct `ClusterPropertiesEncryption` has been removed
- Struct `ClusterPropertiesEncryptionCustomerManagedKeyEncryption` has been removed
- Struct `ClusterPropertiesEncryptionCustomerManagedKeyEncryptionKeyIdentity` has been removed
- Struct `DatabaseList` has been removed
- Struct `DatabasePropertiesGeoReplication` has been removed
- Struct `DatabaseUpdate` has been removed
- Struct `ExportClusterParameters` has been removed
- Struct `ImportClusterParameters` has been removed
- Struct `LinkedDatabase` has been removed
- Struct `LocationInfo` has been removed
- Struct `ManagedServiceIdentity` has been removed
- Struct `Module` has been removed
- Struct `OperationStatus` has been removed
- Struct `PrivateLinkResource` has been removed
- Struct `PrivateLinkResourceListResult` has been removed
- Struct `PrivateLinkResourceProperties` has been removed
- Struct `RegionSKUDetail` has been removed
- Struct `RegionSKUDetails` has been removed
- Struct `SKUDetail` has been removed
- Struct `UserAssignedIdentity` has been removed
- Field `ClusterList` of struct `ClientListByResourceGroupResponse` has been removed
- Field `Identity` of struct `Cluster` has been removed
- Field `Zones` of struct `Cluster` has been removed
- Field `Encryption` of struct `ClusterProperties` has been removed
- Field `MinimumTLSVersion` of struct `ClusterProperties` has been removed
- Field `PrivateEndpointConnections` of struct `ClusterProperties` has been removed
- Field `Identity` of struct `ClusterUpdate` has been removed
- Field `ClusteringPolicy` of struct `DatabaseProperties` has been removed
- Field `EvictionPolicy` of struct `DatabaseProperties` has been removed
- Field `GeoReplication` of struct `DatabaseProperties` has been removed
- Field `Modules` of struct `DatabaseProperties` has been removed
- Field `ResourceState` of struct `DatabaseProperties` has been removed
- Field `DatabaseList` of struct `DatabasesClientListByClusterResponse` has been removed
- Field `OperationStatus` of struct `OperationsStatusClientGetResponse` has been removed
- Field `PrivateLinkResourceListResult` of struct `PrivateLinkResourcesClientListByClusterResponse` has been removed

### Features Added

- New value `AofFrequencyPerSecond` added to enum type `AofFrequency`
- New value `PrivateEndpointConnectionProvisioningStateCanceled` added to enum type `PrivateEndpointConnectionProvisioningState`
- New value `ProvisioningStateAccepted`, `ProvisioningStateProvisioning` added to enum type `ProvisioningState`
- New value `RdbFrequencyPer12Hours`, `RdbFrequencyPer6Hours`, `RdbFrequencyPerHour` added to enum type `RdbFrequency`
- New value `TLSVersionOnePointOne`, `TLSVersionOnePointTwo`, `TLSVersionOnePointZero` added to enum type `TLSVersion`
- New enum type `ClientProtocol` with values `ClientProtocolEncrypted`, `ClientProtocolPlaintext`
- New enum type `KeyType` with values `KeyTypePrimary`, `KeyTypeSecondary`
- New function `*Client.NewListBySubscriptionPager(*ClientListBySubscriptionOptions) *runtime.Pager[ClientListBySubscriptionResponse]`
- New function `*DatabasesClient.BeginExportRdb(context.Context, string, string, string, ExportParameters, *DatabasesClientBeginExportRdbOptions) (*runtime.Poller[DatabasesClientExportRdbResponse], error)`
- New function `*DatabasesClient.BeginFush(context.Context, string, string, string, FlushParameters, *DatabasesClientBeginFushOptions) (*runtime.Poller[DatabasesClientFushResponse], error)`
- New function `*DatabasesClient.BeginImportRdb(context.Context, string, string, string, ImportParameters, *DatabasesClientBeginImportRdbOptions) (*runtime.Poller[DatabasesClientImportRdbResponse], error)`
- New function `*PrivateEndpointConnectionsClient.BeginCreate(context.Context, string, string, string, PrivateEndpointConnection, *PrivateEndpointConnectionsClientBeginCreateOptions) (*runtime.Poller[PrivateEndpointConnectionsClientCreateResponse], error)`
- New function `*PrivateEndpointConnectionsClient.NewListByClusterPager(string, string, *PrivateEndpointConnectionsClientListByClusterOptions) *runtime.Pager[PrivateEndpointConnectionsClientListByClusterResponse]`
- New function `*PrivateEndpointConnectionsClient.BeginUpdate(context.Context, string, string, string, any, *PrivateEndpointConnectionsClientBeginUpdateOptions) (*runtime.Poller[PrivateEndpointConnectionsClientUpdateResponse], error)`
- New function `*SKUsClient.NewListBySubscriptionPager(*SKUsClientListBySubscriptionOptions) *runtime.Pager[SKUsClientListBySubscriptionResponse]`
- New struct `ClusterListResult`
- New struct `ClusterNameParameter`
- New struct `ClusterUpdateProperties`
- New struct `DatabaseCreateOrUpdate`
- New struct `DatabaseListResult`
- New struct `DatabasePropertiesCreateOrUpdate`
- New struct `ExportParameters`
- New struct `ImportParameters`
- New struct `OperationStatusResult`
- New struct `PrivateLink`
- New struct `PrivateLinkListResult`
- New struct `PrivateLinkProperties`
- New struct `SKUDetails`
- New struct `SKUDetailsListResult`
- New struct `SKUUpdate`
- New anonymous field `ClusterListResult` in struct `ClientListByResourceGroupResponse`
- New field `MinTLSVersion` in struct `ClusterProperties`
- New anonymous field `DatabaseListResult` in struct `DatabasesClientListByClusterResponse`
- New anonymous field `OperationStatusResult` in struct `OperationsStatusClientGetResponse`
- New field `NextLink` in struct `PrivateEndpointConnectionListResult`
- New anonymous field `PrivateLinkListResult` in struct `PrivateLinkResourcesClientListByClusterResponse`


## 1.1.0-beta.1 (2023-03-24)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module
- New enum type `CmkIdentityType` with values `CmkIdentityTypeSystemAssignedIdentity`, `CmkIdentityTypeUserAssignedIdentity`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New function `*DatabasesClient.BeginFlush(context.Context, string, string, string, FlushParameters, *DatabasesClientBeginFlushOptions) (*runtime.Poller[DatabasesClientFlushResponse], error)`
- New function `NewSKUsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SKUsClient, error)`
- New function `*SKUsClient.NewListPager(string, *SKUsClientListOptions) *runtime.Pager[SKUsClientListResponse]`
- New function `timeRFC3339.MarshalText() ([]byte, error)`
- New function `*timeRFC3339.Parse(string) error`
- New function `*timeRFC3339.UnmarshalText([]byte) error`
- New struct `Capability`
- New struct `ClusterPropertiesEncryption`
- New struct `ClusterPropertiesEncryptionCustomerManagedKeyEncryption`
- New struct `ClusterPropertiesEncryptionCustomerManagedKeyEncryptionKeyIdentity`
- New struct `FlushParameters`
- New struct `LocationInfo`
- New struct `ManagedServiceIdentity`
- New struct `RegionSKUDetail`
- New struct `RegionSKUDetails`
- New struct `SKUDetail`
- New struct `SystemData`
- New struct `UserAssignedIdentity`
- New field `Identity` in struct `Cluster`
- New field `SystemData` in struct `Cluster`
- New field `Encryption` in struct `ClusterProperties`
- New field `Identity` in struct `ClusterUpdate`
- New field `SystemData` in struct `Database`
- New field `SystemData` in struct `PrivateEndpointConnection`
- New field `SystemData` in struct `PrivateLinkResource`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `TrackedResource`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).