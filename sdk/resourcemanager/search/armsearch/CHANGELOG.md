# Release History

## 1.2.0 (2023-06-07)
### Features Added

- New value `IdentityTypeSystemAssignedUserAssigned`, `IdentityTypeUserAssigned` added to enum type `IdentityType`
- New enum type `AADAuthFailureMode` with values `AADAuthFailureModeHttp401WithBearerChallenge`, `AADAuthFailureModeHttp403`
- New enum type `PrivateLinkServiceConnectionProvisioningState` with values `PrivateLinkServiceConnectionProvisioningStateCanceled`, `PrivateLinkServiceConnectionProvisioningStateDeleting`, `PrivateLinkServiceConnectionProvisioningStateFailed`, `PrivateLinkServiceConnectionProvisioningStateIncomplete`, `PrivateLinkServiceConnectionProvisioningStateSucceeded`, `PrivateLinkServiceConnectionProvisioningStateUpdating`
- New enum type `SearchEncryptionComplianceStatus` with values `SearchEncryptionComplianceStatusCompliant`, `SearchEncryptionComplianceStatusNonCompliant`
- New enum type `SearchEncryptionWithCmk` with values `SearchEncryptionWithCmkDisabled`, `SearchEncryptionWithCmkEnabled`, `SearchEncryptionWithCmkUnspecified`
- New struct `DataPlaneAADOrAPIKeyAuthOption`
- New struct `DataPlaneAuthOptions`
- New struct `EncryptionWithCmk`
- New struct `UserAssignedManagedIdentity`
- New field `UserAssignedIdentities` in struct `Identity`
- New field `GroupID`, `ProvisioningState` in struct `PrivateEndpointConnectionProperties`
- New field `AuthOptions`, `DisableLocalAuth`, `EncryptionWithCmk` in struct `ServiceProperties`


## 1.1.0 (2023-04-03)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).