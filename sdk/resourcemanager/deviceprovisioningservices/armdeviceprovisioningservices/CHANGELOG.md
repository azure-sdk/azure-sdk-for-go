# Release History

## 2.0.0-beta.1 (2023-05-02)
### Breaking Changes

- Type of `ErrorDetails.Code` has been changed from `*string` to `*int32`
- Field `EnableDataResidency` of struct `IotDpsPropertiesDescription` has been removed

### Features Added

- New enum type `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New function `*IotDpsResourceClient.Failover(context.Context, string, string, CustomerInitiatedFailoverInput, *IotDpsResourceClientFailoverOptions) (IotDpsResourceClientFailoverResponse, error)`
- New struct `CustomerInitiatedFailoverInput`
- New struct `IotDpsPropertiesDescriptionDpsFailoverDescription`
- New struct `ManagedServiceIdentity`
- New struct `UserAssignedIdentity`
- New field `DpsFailoverDescription`, `EnableCustomerInitiatedFailover`, `PortalOperationsHostName` in struct `IotDpsPropertiesDescription`
- New field `Identity`, `Resourcegroup`, `Subscriptionid` in struct `ProvisioningServiceDescription`
- New field `Resourcegroup`, `Subscriptionid` in struct `Resource`


## 1.1.0 (2023-03-28)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).