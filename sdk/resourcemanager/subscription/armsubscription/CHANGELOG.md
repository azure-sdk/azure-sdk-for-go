# Release History

## 1.3.0 (2024-03-04)
### Features Added

- New enum type `Provisioning` with values `ProvisioningAccepted`, `ProvisioningPending`, `ProvisioningSucceeded`
- New function `*ClientFactory.NewOperationClient() *OperationClient`
- New function `NewOperationClient(azcore.TokenCredential, *arm.ClientOptions) (*OperationClient, error)`
- New function `*OperationClient.Get(context.Context, string, *OperationClientGetOptions) (OperationClientGetResponse, error)`
- New struct `CreationResult`
- New field `ProvisioningState` in struct `AcceptOwnershipStatusResponse`
- New field `Tags`, `TenantID` in struct `Subscription`
- New field `Country`, `CountryCode`, `DefaultDomain`, `DisplayName`, `Domains`, `TenantCategory`, `TenantType` in struct `TenantIDDescription`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).