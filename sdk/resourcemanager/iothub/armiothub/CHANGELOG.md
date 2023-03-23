# Release History

## 1.1.0-beta.1 (2023-03-23)
### Features Added

- New value `RoutingSourceDigitalTwinChangeEvents`, `RoutingSourceMqttBrokerMessages` added to enum type `RoutingSource`
- New enum type `IPVersion` with values `IPVersionIPv4`, `IPVersionIPv4IPv6`, `IPVersionIPv6`
- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewCertificatesClient() *CertificatesClient`
- New function `*ClientFactory.NewClient() *Client`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient`
- New function `*ClientFactory.NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient`
- New function `*ClientFactory.NewResourceClient() *ResourceClient`
- New function `*ClientFactory.NewResourceProviderCommonClient() *ResourceProviderCommonClient`
- New struct `ClientFactory`
- New struct `EncryptionPropertiesDescription`
- New struct `KeyVaultKeyProperties`
- New struct `PropertiesDeviceStreams`
- New struct `RootCertificateProperties`
- New struct `RoutingCosmosDBSQLAPIProperties`
- New field `DeviceStreams` in struct `Properties`
- New field `Encryption` in struct `Properties`
- New field `IPVersion` in struct `Properties`
- New field `RootCertificate` in struct `Properties`
- New field `CosmosDBSQLCollections` in struct `RoutingEndpoints`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).