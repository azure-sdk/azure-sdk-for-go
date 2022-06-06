# Release History

## 1.1.0 (2022-06-06)
### Features Added

- New const `PacketCaptureTargetTypeAzureVMSS`
- New const `EndpointTypeAzureVMSS`
- New const `EndpointTypeAzureArcVM`
- New const `PacketCaptureTargetTypeAzureVM`
- New function `PossiblePacketCaptureTargetTypeValues() []PacketCaptureTargetType`
- New struct `PacketCaptureMachineScope`
- New field `TargetType` in struct `PacketCaptureResultProperties`
- New field `Scope` in struct `PacketCaptureResultProperties`
- New field `Scope` in struct `PacketCaptureParameters`
- New field `TargetType` in struct `PacketCaptureParameters`
- New field `Priority` in struct `ApplicationGatewayRoutingRulePropertiesFormat`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).