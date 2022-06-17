# Release History

## 2.0.0 (2022-06-17)
### Breaking Changes

- Struct `ExplicitProxySettings` has been removed
- Field `ExplicitProxySettings` of struct `FirewallPolicyPropertiesFormat` has been removed

### Features Added

- New field `Scope` in struct `PacketCaptureResultProperties`
- New field `TargetType` in struct `PacketCaptureResultProperties`
- New field `AutoLearnPrivateRanges` in struct `FirewallPolicySNAT`
- New field `Priority` in struct `ApplicationGatewayRoutingRulePropertiesFormat`
- New field `Scope` in struct `PacketCaptureParameters`
- New field `TargetType` in struct `PacketCaptureParameters`
- New field `VirtualRouterAutoScaleConfiguration` in struct `VirtualHubProperties`
- New field `FlushConnection` in struct `SecurityGroupPropertiesFormat`
- New field `NoInternetAdvertise` in struct `CustomIPPrefixPropertiesFormat`
- New field `ExplicitProxy` in struct `FirewallPolicyPropertiesFormat`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).