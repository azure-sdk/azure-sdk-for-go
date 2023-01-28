# Release History

## 2.0.0-beta.1 (2023-01-28)
### Breaking Changes

- Type of `EventHubConsumerGroupInfo.Properties` has been changed from `map[string]interface{}` to `map[string]any`
- Type of `RoutingTwin.Tags` has been changed from `interface{}` to `any`
- Type of `RoutingTwinProperties.Desired` has been changed from `interface{}` to `any`
- Type of `RoutingTwinProperties.Reported` has been changed from `interface{}` to `any`

### Features Added

- New value `RoutingSourceDigitalTwinChangeEvents`, `RoutingSourceMqttBrokerMessages` added to type alias `RoutingSource`
- New struct `EncryptionPropertiesDescription`
- New struct `KeyVaultKeyProperties`
- New struct `PropertiesDeviceStreams`
- New struct `RootCertificateProperties`
- New struct `RoutingCosmosDBSQLAPIProperties`
- New field `DeviceStreams` in struct `Properties`
- New field `Encryption` in struct `Properties`
- New field `RootCertificate` in struct `Properties`
- New field `CosmosDBSQLCollections` in struct `RoutingEndpoints`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).