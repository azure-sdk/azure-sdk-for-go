# Release History

## 1.1.0 (2022-11-30)
### Features Added

- New const `DigitalTwinsIdentityTypeSystemAssignedUserAssigned`
- New const `DigitalTwinsIdentityTypeUserAssigned`
- New const `IdentityTypeSystemAssigned`
- New const `IdentityTypeUserAssigned`
- New type alias `IdentityType`
- New function `PossibleIdentityTypeValues() []IdentityType`
- New struct `ManagedIdentityReference`
- New struct `UserAssignedIdentity`
- New field `Identity` in struct `AzureDataExplorerConnectionProperties`
- New field `Identity` in struct `EndpointResourceProperties`
- New field `Identity` in struct `EventGrid`
- New field `Identity` in struct `EventHub`
- New field `UserAssignedIdentities` in struct `Identity`
- New field `Identity` in struct `ServiceBus`
- New field `Identity` in struct `TimeSeriesDatabaseConnectionProperties`


## 1.0.0 (2022-06-14)
### Features Added

- New const `TimeSeriesDatabaseConnectionStateUpdating`
- New const `EndpointProvisioningStateUpdating`


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).