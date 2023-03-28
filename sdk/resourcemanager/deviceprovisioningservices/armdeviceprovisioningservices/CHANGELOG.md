# Release History

## 2.0.0-beta.1 (2023-03-28)
### Breaking Changes

- Type of `ErrorDetails.Code` has been changed from `*string` to `*int32`

### Features Added

- New enum type `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewDpsCertificateClient() *DpsCertificateClient`
- New function `*ClientFactory.NewIotDpsResourceClient() *IotDpsResourceClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New struct `ClientFactory`
- New struct `ManagedServiceIdentity`
- New struct `UserAssignedIdentity`
- New field `PortalOperationsHostName` in struct `IotDpsPropertiesDescription`
- New field `Identity` in struct `ProvisioningServiceDescription`
- New field `Resourcegroup` in struct `ProvisioningServiceDescription`
- New field `Subscriptionid` in struct `ProvisioningServiceDescription`
- New field `Resourcegroup` in struct `Resource`
- New field `Subscriptionid` in struct `Resource`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).