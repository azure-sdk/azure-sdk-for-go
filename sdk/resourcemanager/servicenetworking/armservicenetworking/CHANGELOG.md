# Release History

## 2.0.0-beta.1 (2024-07-18)
### Breaking Changes

- Function `*AssociationsInterfaceClient.Update` parameter(s) have been changed from `(context.Context, string, string, string, AssociationUpdate, *AssociationsInterfaceClientUpdateOptions)` to `(context.Context, string, string, string, Association, *AssociationsInterfaceClientUpdateOptions)`
- Function `*FrontendsInterfaceClient.Update` parameter(s) have been changed from `(context.Context, string, string, string, FrontendUpdate, *FrontendsInterfaceClientUpdateOptions)` to `(context.Context, string, string, string, Frontend, *FrontendsInterfaceClientUpdateOptions)`
- Function `*TrafficControllerInterfaceClient.Update` parameter(s) have been changed from `(context.Context, string, string, TrafficControllerUpdate, *TrafficControllerInterfaceClientUpdateOptions)` to `(context.Context, string, string, TrafficController, *TrafficControllerInterfaceClientUpdateOptions)`
- Function `*AssociationsInterfaceClient.NewListByTrafficControllerPager` has been removed
- Function `*FrontendsInterfaceClient.NewListByTrafficControllerPager` has been removed
- Struct `AssociationSubnetUpdate` has been removed
- Struct `AssociationUpdate` has been removed
- Struct `AssociationUpdateProperties` has been removed
- Struct `FrontendUpdate` has been removed
- Struct `TrafficControllerUpdate` has been removed

### Features Added

- New enum type `PolicyType` with values `PolicyTypeWAF`
- New function `*AssociationsInterfaceClient.NewListByResourceGroupPager(string, string, *AssociationsInterfaceClientListByResourceGroupOptions) *runtime.Pager[AssociationsInterfaceClientListByResourceGroupResponse]`
- New function `*ClientFactory.NewSecurityPoliciesInterfaceClient() *SecurityPoliciesInterfaceClient`
- New function `*FrontendsInterfaceClient.NewListByResourceGroupPager(string, string, *FrontendsInterfaceClientListByResourceGroupOptions) *runtime.Pager[FrontendsInterfaceClientListByResourceGroupResponse]`
- New function `NewSecurityPoliciesInterfaceClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SecurityPoliciesInterfaceClient, error)`
- New function `*SecurityPoliciesInterfaceClient.BeginCreateOrUpdate(context.Context, string, string, string, SecurityPolicy, *SecurityPoliciesInterfaceClientBeginCreateOrUpdateOptions) (*runtime.Poller[SecurityPoliciesInterfaceClientCreateOrUpdateResponse], error)`
- New function `*SecurityPoliciesInterfaceClient.BeginDelete(context.Context, string, string, string, *SecurityPoliciesInterfaceClientBeginDeleteOptions) (*runtime.Poller[SecurityPoliciesInterfaceClientDeleteResponse], error)`
- New function `*SecurityPoliciesInterfaceClient.Get(context.Context, string, string, string, *SecurityPoliciesInterfaceClientGetOptions) (SecurityPoliciesInterfaceClientGetResponse, error)`
- New function `*SecurityPoliciesInterfaceClient.NewListByResourceGroupPager(string, string, *SecurityPoliciesInterfaceClientListByResourceGroupOptions) *runtime.Pager[SecurityPoliciesInterfaceClientListByResourceGroupResponse]`
- New function `*SecurityPoliciesInterfaceClient.Update(context.Context, string, string, string, SecurityPolicy, *SecurityPoliciesInterfaceClientUpdateOptions) (SecurityPoliciesInterfaceClientUpdateResponse, error)`
- New struct `SecurityPolicy`
- New struct `SecurityPolicyConfigurations`
- New struct `SecurityPolicyListResult`
- New struct `SecurityPolicyProperties`
- New struct `WafPolicy`
- New struct `WafSecurityPolicy`
- New field `SecurityPolicies`, `SecurityPolicyConfigurations` in struct `TrafficControllerProperties`


## 1.0.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.

### Other Changes

- Release stable version.


## 0.3.0 (2023-05-26)
### Breaking Changes

- Type of `AssociationProperties.AssociationType` has been changed from `*string` to `*AssociationType`
- Type of `AssociationUpdateProperties.AssociationType` has been changed from `*string` to `*AssociationType`
- Type of `AssociationUpdateProperties.Subnet` has been changed from `*AssociationSubnet` to `*AssociationSubnetUpdate`
- Enum `FrontendIPAddressVersion` has been removed
- Struct `FrontendPropertiesIPAddress` has been removed
- Struct `FrontendUpdateProperties` has been removed
- Field `IPAddressVersion`, `Mode`, `PublicIPAddress` of struct `FrontendProperties` has been removed
- Field `Properties` of struct `FrontendUpdate` has been removed
- Field `Properties` of struct `TrafficControllerUpdate` has been removed

### Features Added

- New enum type `AssociationType` with values `AssociationTypeSubnets`
- New struct `AssociationSubnetUpdate`
- New field `Fqdn` in struct `FrontendProperties`


## 0.2.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 0.2.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.1.0 (2023-01-11)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).