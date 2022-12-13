# Release History

## 2.0.0 (2022-12-13)
### Breaking Changes

- Function `*DatabasesClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, string, DatabaseUpdate, *DatabasesClientBeginUpdateOptions)` to `(context.Context, string, string, string, interface{}, *DatabasesClientBeginUpdateOptions)`
- Type of `ClusterUpdate.Properties` has been changed from `*ClusterProperties` to `*ClusterUpdateProperties`
- Type of `DatabaseProperties.ClientProtocol` has been changed from `*Protocol` to `*ClientProtocol`
- Type of `RegenerateKeyParameters.KeyType` has been changed from `*AccessKeyType` to `*KeyType`
- Const `AofFrequencyOneS` has been removed
- Const `ProvisioningStateCreating` has been removed
- Const `RdbFrequencyOneH` has been removed
- Const `RdbFrequencySixH` has been removed
- Const `RdbFrequencyTwelveH` has been removed
- Const `TLSVersionOne0` has been removed
- Const `TLSVersionOne1` has been removed
- Const `TLSVersionOne2` has been removed
- Type alias `AccessKeyType` has been removed
- Type alias `ClusteringPolicy` has been removed
- Type alias `LinkState` has been removed
- Type alias `Protocol` has been removed
- Function `*Client.NewListPager` has been removed
- Function `*DatabasesClient.BeginExport` has been removed
- Function `*DatabasesClient.BeginImport` has been removed
- Function `*PrivateEndpointConnectionsClient.NewListPager` has been removed
- Function `*PrivateEndpointConnectionsClient.BeginPut` has been removed
- Operation `*PrivateEndpointConnectionsClient.Delete` has been changed to LRO, use `*PrivateEndpointConnectionsClient.BeginDelete` instead.
- Struct `ClientListResponse` has been removed
- Struct `ClusterList` has been removed
- Struct `DatabaseList` has been removed
- Struct `DatabasePropertiesGeoReplication` has been removed
- Struct `DatabaseUpdate` has been removed
- Struct `DatabasesClientExportResponse` has been removed
- Struct `DatabasesClientImportResponse` has been removed
- Struct `ExportClusterParameters` has been removed
- Struct `ImportClusterParameters` has been removed
- Struct `LinkedDatabase` has been removed
- Struct `Module` has been removed
- Struct `PrivateEndpointConnectionsClientListResponse` has been removed
- Struct `PrivateEndpointConnectionsClientPutResponse` has been removed
- Struct `PrivateLinkResource` has been removed
- Struct `PrivateLinkResourceListResult` has been removed
- Struct `PrivateLinkResourceProperties` has been removed
- Field `ClusterList` of struct `ClientListByResourceGroupResponse` has been removed
- Field `Zones` of struct `Cluster` has been removed
- Field `MinimumTLSVersion` of struct `ClusterProperties` has been removed
- Field `PrivateEndpointConnections` of struct `ClusterProperties` has been removed
- Field `ClusteringPolicy` of struct `DatabaseProperties` has been removed
- Field `GeoReplication` of struct `DatabaseProperties` has been removed
- Field `Modules` of struct `DatabaseProperties` has been removed
- Field `ResourceState` of struct `DatabaseProperties` has been removed
- Field `DatabaseList` of struct `DatabasesClientListByClusterResponse` has been removed
- Field `EndTime` of struct `OperationStatus` has been removed
- Field `Error` of struct `OperationStatus` has been removed
- Field `StartTime` of struct `OperationStatus` has been removed
- Field `Status` of struct `OperationStatus` has been removed
- Field `PrivateLinkResourceListResult` of struct `PrivateLinkResourcesClientListByClusterResponse` has been removed

### Features Added

- New value `AofFrequencyPerSecond` added to type alias `AofFrequency`
- New value `PrivateEndpointConnectionProvisioningStateCanceled` added to type alias `PrivateEndpointConnectionProvisioningState`
- New value `ProvisioningStateAccepted`, `ProvisioningStateProvisioning` added to type alias `ProvisioningState`
- New value `RdbFrequencyPer12Hours`, `RdbFrequencyPer6Hours`, `RdbFrequencyPerHour` added to type alias `RdbFrequency`
- New value `TLSVersionOnePointOne`, `TLSVersionOnePointTwo`, `TLSVersionOnePointZero` added to type alias `TLSVersion`
- New type alias `ClientProtocol` with values `ClientProtocolEncrypted`, `ClientProtocolPlaintext`
- New type alias `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New type alias `KeyType` with values `KeyTypePrimary`, `KeyTypeSecondary`
- New function `*Client.NewListBySubscriptionPager(*ClientListBySubscriptionOptions) *runtime.Pager[ClientListBySubscriptionResponse]`
- New function `*DatabasesClient.BeginExportRdb(context.Context, string, string, string, ExportParameters, *DatabasesClientBeginExportRdbOptions) (*runtime.Poller[DatabasesClientExportRdbResponse], error)`
- New function `*DatabasesClient.BeginImportRdb(context.Context, string, string, string, ImportParameters, *DatabasesClientBeginImportRdbOptions) (*runtime.Poller[DatabasesClientImportRdbResponse], error)`
- New function `*PrivateEndpointConnectionsClient.BeginCreate(context.Context, string, string, string, PrivateEndpointConnection, *PrivateEndpointConnectionsClientBeginCreateOptions) (*runtime.Poller[PrivateEndpointConnectionsClientCreateResponse], error)`
- New function `*PrivateEndpointConnectionsClient.NewListByClusterPager(string, string, *PrivateEndpointConnectionsClientListByClusterOptions) *runtime.Pager[PrivateEndpointConnectionsClientListByClusterResponse]`
- New function `*PrivateEndpointConnectionsClient.BeginUpdate(context.Context, string, string, string, interface{}, *PrivateEndpointConnectionsClientBeginUpdateOptions) (*runtime.Poller[PrivateEndpointConnectionsClientUpdateResponse], error)`
- New function `timeRFC3339.MarshalText() ([]byte, error)`
- New function `*timeRFC3339.Parse(string) error`
- New function `*timeRFC3339.UnmarshalText([]byte) error`
- New struct `ClientListBySubscriptionResponse`
- New struct `ClusterListResult`
- New struct `ClusterUpdateProperties`
- New struct `DatabaseListResult`
- New struct `DatabasesClientExportRdbResponse`
- New struct `DatabasesClientImportRdbResponse`
- New struct `ExportParameters`
- New struct `ImportParameters`
- New struct `PrivateEndpointConnectionsClientCreateResponse`
- New struct `PrivateEndpointConnectionsClientListByClusterResponse`
- New struct `PrivateEndpointConnectionsClientUpdateResponse`
- New struct `PrivateLink`
- New struct `PrivateLinkListResult`
- New struct `PrivateLinkProperties`
- New struct `SystemData`
- New anonymous field `ClusterListResult` in struct `ClientListByResourceGroupResponse`
- New field `SystemData` in struct `Cluster`
- New field `MinTLSVersion` in struct `ClusterProperties`
- New field `SystemData` in struct `Database`
- New anonymous field `DatabaseListResult` in struct `DatabasesClientListByClusterResponse`
- New field `Properties` in struct `OperationStatus`
- New field `SystemData` in struct `OperationStatus`
- New field `Type` in struct `OperationStatus`
- New field `SystemData` in struct `PrivateEndpointConnection`
- New field `NextLink` in struct `PrivateEndpointConnectionListResult`
- New anonymous field `PrivateLinkListResult` in struct `PrivateLinkResourcesClientListByClusterResponse`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `TrackedResource`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).