# Unreleased

## Additive Changes

### New Constants

1. EndpointType.EndpointTypeAzureArcVM
1. EndpointType.EndpointTypeAzureVMSS
1. PacketCaptureTargetType.PacketCaptureTargetTypeAzureVM
1. PacketCaptureTargetType.PacketCaptureTargetTypeAzureVMSS

### New Funcs

1. PossiblePacketCaptureTargetTypeValues() []PacketCaptureTargetType

### Struct Changes

#### New Structs

1. PacketCaptureMachineScope

#### New Struct Fields

1. PacketCaptureParameters.Scope
1. PacketCaptureParameters.TargetType
1. PacketCaptureResultProperties.Scope
1. PacketCaptureResultProperties.TargetType
