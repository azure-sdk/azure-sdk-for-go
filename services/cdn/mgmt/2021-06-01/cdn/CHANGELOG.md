# Unreleased

## Breaking Changes

### Signature Changes

#### Struct Fields

1. CustomDomainProperties.ProvisioningState changed type from *string to CustomHTTPSProvisioningState
1. EndpointProperties.ProvisioningState changed type from *string to EndpointProvisioningState
1. OriginGroupProperties.ProvisioningState changed type from *string to OriginGroupProvisioningState
1. OriginProperties.ProvisioningState changed type from *string to OriginProvisioningState
1. ProfileProperties.ProvisioningState changed type from *string to ProfileProvisioningState
1. ResourceUsage.Unit changed type from *string to ResourceUsageUnit

## Additive Changes

### New Constants

1. EndpointProvisioningState.EndpointProvisioningStateCreating
1. EndpointProvisioningState.EndpointProvisioningStateDeleting
1. EndpointProvisioningState.EndpointProvisioningStateFailed
1. EndpointProvisioningState.EndpointProvisioningStateSucceeded
1. EndpointProvisioningState.EndpointProvisioningStateUpdating
1. OriginGroupProvisioningState.OriginGroupProvisioningStateCreating
1. OriginGroupProvisioningState.OriginGroupProvisioningStateDeleting
1. OriginGroupProvisioningState.OriginGroupProvisioningStateFailed
1. OriginGroupProvisioningState.OriginGroupProvisioningStateSucceeded
1. OriginGroupProvisioningState.OriginGroupProvisioningStateUpdating
1. OriginProvisioningState.OriginProvisioningStateCreating
1. OriginProvisioningState.OriginProvisioningStateDeleting
1. OriginProvisioningState.OriginProvisioningStateFailed
1. OriginProvisioningState.OriginProvisioningStateSucceeded
1. OriginProvisioningState.OriginProvisioningStateUpdating
1. ProfileProvisioningState.ProfileProvisioningStateCreating
1. ProfileProvisioningState.ProfileProvisioningStateDeleting
1. ProfileProvisioningState.ProfileProvisioningStateFailed
1. ProfileProvisioningState.ProfileProvisioningStateSucceeded
1. ProfileProvisioningState.ProfileProvisioningStateUpdating
1. ResourceUsageUnit.ResourceUsageUnitCount

### New Funcs

1. PossibleEndpointProvisioningStateValues() []EndpointProvisioningState
1. PossibleOriginGroupProvisioningStateValues() []OriginGroupProvisioningState
1. PossibleOriginProvisioningStateValues() []OriginProvisioningState
1. PossibleProfileProvisioningStateValues() []ProfileProvisioningState
1. PossibleResourceUsageUnitValues() []ResourceUsageUnit
