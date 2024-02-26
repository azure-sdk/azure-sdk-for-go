# Release History

## 1.2.0-beta.1 (2024-02-26)
### Features Added

- New enum type `EnableStatus` with values `EnableStatusDisabled`, `EnableStatusEnabled`
- New enum type `ProductSerialStatusValues` with values `ProductSerialStatusValuesAllocated`, `ProductSerialStatusValuesInProgress`
- New enum type `RegistrationStatus` with values `RegistrationStatusNotRegistered`, `RegistrationStatusRegistered`
- New function `*ClientFactory.NewPaloAltoNetworksCloudngfwClient() *PaloAltoNetworksCloudngfwClient`
- New function `NewPaloAltoNetworksCloudngfwClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PaloAltoNetworksCloudngfwClient, error)`
- New function `*PaloAltoNetworksCloudngfwClient.CreateProductSerialNumber(context.Context, *PaloAltoNetworksCloudngfwClientCreateProductSerialNumberOptions) (PaloAltoNetworksCloudngfwClientCreateProductSerialNumberResponse, error)`
- New function `*PaloAltoNetworksCloudngfwClient.ListCloudManagerTenants(context.Context, *PaloAltoNetworksCloudngfwClientListCloudManagerTenantsOptions) (PaloAltoNetworksCloudngfwClientListCloudManagerTenantsResponse, error)`
- New function `*PaloAltoNetworksCloudngfwClient.ListProductSerialNumberStatus(context.Context, *PaloAltoNetworksCloudngfwClientListProductSerialNumberStatusOptions) (PaloAltoNetworksCloudngfwClientListProductSerialNumberStatusResponse, error)`
- New function `*PaloAltoNetworksCloudngfwClient.ListSupportInfo(context.Context, *PaloAltoNetworksCloudngfwClientListSupportInfoOptions) (PaloAltoNetworksCloudngfwClientListSupportInfoResponse, error)`
- New struct `CloudManagerTenantList`
- New struct `ProductSerialNumberRequestStatus`
- New struct `ProductSerialNumberStatus`
- New struct `StrataCloudManagerConfig`
- New struct `StrataCloudManagerInfo`
- New struct `SupportInfoModel`
- New field `IsStrataCloudManaged`, `StrataCloudManagerConfig` in struct `FirewallDeploymentProperties`
- New field `IsStrataCloudManaged`, `StrataCloudManagerConfig` in struct `FirewallResourceUpdateProperties`
- New field `IsStrataCloudManaged`, `StrataCloudManagerInfo` in struct `FirewallStatusProperty`
- New field `PrivateSourceNatRulesDestination` in struct `NetworkProfile`


## 1.1.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.
- New field `TrustedRanges` in struct `NetworkProfile`


## 1.0.0 (2023-07-14)
### Other Changes

- Release stable version.

## 0.1.0 (2023-04-28)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).