# Release History

## 2.0.0 (2022-06-20)
### Breaking Changes

- Type of `OriginGroupProperties.ProvisioningState` has been changed from `*OriginGroupProvisioningState` to `*string`
- Type of `OriginProperties.ProvisioningState` has been changed from `*OriginProvisioningState` to `*string`
- Type of `ProfileProperties.ProvisioningState` has been changed from `*ProfileProvisioningState` to `*string`
- Type of `EndpointProperties.ProvisioningState` has been changed from `*EndpointProvisioningState` to `*string`
- Type of `CustomDomainProperties.ProvisioningState` has been changed from `*CustomHTTPSProvisioningState` to `*string`
- Type of `ResourceUsage.Unit` has been changed from `*ResourceUsageUnit` to `*string`
- Const `OriginProvisioningStateUpdating` has been removed
- Const `ProfileProvisioningStateCreating` has been removed
- Const `OriginGroupProvisioningStateCreating` has been removed
- Const `EndpointProvisioningStateCreating` has been removed
- Const `OriginProvisioningStateDeleting` has been removed
- Const `ProfileProvisioningStateDeleting` has been removed
- Const `ProfileProvisioningStateFailed` has been removed
- Const `EndpointProvisioningStateFailed` has been removed
- Const `OriginGroupProvisioningStateUpdating` has been removed
- Const `OriginGroupProvisioningStateFailed` has been removed
- Const `OriginGroupProvisioningStateDeleting` has been removed
- Const `OriginGroupProvisioningStateSucceeded` has been removed
- Const `OriginProvisioningStateSucceeded` has been removed
- Const `OriginProvisioningStateFailed` has been removed
- Const `ProfileProvisioningStateSucceeded` has been removed
- Const `OriginProvisioningStateCreating` has been removed
- Const `EndpointProvisioningStateDeleting` has been removed
- Const `EndpointProvisioningStateSucceeded` has been removed
- Const `EndpointProvisioningStateUpdating` has been removed
- Const `ProfileProvisioningStateUpdating` has been removed
- Const `ResourceUsageUnitCount` has been removed
- Function `PossibleEndpointProvisioningStateValues` has been removed
- Function `PossibleOriginGroupProvisioningStateValues` has been removed
- Function `PossibleProfileProvisioningStateValues` has been removed
- Function `PossibleResourceUsageUnitValues` has been removed
- Function `PossibleOriginProvisioningStateValues` has been removed


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).