# Release History

## 2.0.0 (2023-04-03)
### Breaking Changes

- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New value `SearchServiceStatusStopped` added to enum type `SearchServiceStatus`
- New enum type `AADAuthFailureMode` with values `AADAuthFailureModeHttp401WithBearerChallenge`, `AADAuthFailureModeHttp403`
- New enum type `PrivateLinkServiceConnectionProvisioningState` with values `PrivateLinkServiceConnectionProvisioningStateCanceled`, `PrivateLinkServiceConnectionProvisioningStateDeleting`, `PrivateLinkServiceConnectionProvisioningStateFailed`, `PrivateLinkServiceConnectionProvisioningStateIncomplete`, `PrivateLinkServiceConnectionProvisioningStateSucceeded`, `PrivateLinkServiceConnectionProvisioningStateUpdating`
- New enum type `SearchEncryptionComplianceStatus` with values `SearchEncryptionComplianceStatusCompliant`, `SearchEncryptionComplianceStatusNonCompliant`
- New enum type `SearchEncryptionWithCmk` with values `SearchEncryptionWithCmkDisabled`, `SearchEncryptionWithCmkEnabled`, `SearchEncryptionWithCmkUnspecified`
- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewAdminKeysClient() *AdminKeysClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient`
- New function `*ClientFactory.NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient`
- New function `*ClientFactory.NewQueryKeysClient() *QueryKeysClient`
- New function `*ClientFactory.NewServicesClient() *ServicesClient`
- New function `*ClientFactory.NewSharedPrivateLinkResourcesClient() *SharedPrivateLinkResourcesClient`
- New struct `ClientFactory`
- New struct `DataPlaneAADOrAPIKeyAuthOption`
- New struct `DataPlaneAuthOptions`
- New struct `EncryptionWithCmk`
- New field `GroupID` in struct `PrivateEndpointConnectionProperties`
- New field `ProvisioningState` in struct `PrivateEndpointConnectionProperties`
- New field `AuthOptions` in struct `ServiceProperties`
- New field `DisableLocalAuth` in struct `ServiceProperties`
- New field `EncryptionWithCmk` in struct `ServiceProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).