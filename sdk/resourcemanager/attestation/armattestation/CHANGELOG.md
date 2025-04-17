# Release History

## 1.3.0 (2025-04-17)
### Features Added

- New enum type `PublicNetworkAccessType` with values `PublicNetworkAccessTypeDisabled`, `PublicNetworkAccessTypeEnabled`
- New enum type `TpmAttestationAuthenticationType` with values `TpmAttestationAuthenticationTypeDisabled`, `TpmAttestationAuthenticationTypeEnabled`
- New function `*ClientFactory.NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient`
- New function `NewPrivateLinkResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrivateLinkResourcesClient, error)`
- New function `*PrivateLinkResourcesClient.ListByProvider(context.Context, string, string, *PrivateLinkResourcesClientListByProviderOptions) (PrivateLinkResourcesClientListByProviderResponse, error)`
- New struct `LogSpecification`
- New struct `OperationProperties`
- New struct `PrivateLinkResource`
- New struct `PrivateLinkResourceListResult`
- New struct `PrivateLinkResourceProperties`
- New struct `ServicePatchSpecificParams`
- New struct `ServiceSpecification`
- New field `Properties` in struct `OperationsDefinition`
- New field `PublicNetworkAccess`, `TpmAttestationAuthentication` in struct `ServiceCreationSpecificParams`
- New field `Properties` in struct `ServicePatchParams`
- New field `PublicNetworkAccess`, `TpmAttestationAuthentication` in struct `StatusResult`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.0 (2023-03-27)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/attestation/armattestation` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).