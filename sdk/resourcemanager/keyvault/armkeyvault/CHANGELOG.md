# Release History

## 1.1.0 (2022-08-17)
### Features Added

- New const `KeyRotationPolicyActionTypeRotate`
- New const `ActivationStatusFailed`
- New const `ActivationStatusNotActivated`
- New const `KeyRotationPolicyActionTypeNotify`
- New const `KeyPermissionsGetrotationpolicy`
- New const `KeyPermissionsRelease`
- New const `KeyPermissionsRotate`
- New const `ActivationStatusUnknown`
- New const `ActivationStatusActive`
- New const `KeyPermissionsSetrotationpolicy`
- New const `JSONWebKeyOperationRelease`
- New function `*ManagedHsmsClient.CheckMhsmNameAvailability(context.Context, CheckMhsmNameAvailabilityParameters, *ManagedHsmsClientCheckMhsmNameAvailabilityOptions) (ManagedHsmsClientCheckMhsmNameAvailabilityResponse, error)`
- New function `PossibleActivationStatusValues() []ActivationStatus`
- New function `PossibleKeyRotationPolicyActionTypeValues() []KeyRotationPolicyActionType`
- New struct `Action`
- New struct `CheckMhsmNameAvailabilityParameters`
- New struct `CheckMhsmNameAvailabilityResult`
- New struct `KeyReleasePolicy`
- New struct `KeyRotationPolicyAttributes`
- New struct `LifetimeAction`
- New struct `ManagedHSMSecurityDomainProperties`
- New struct `ManagedHsmsClientCheckMhsmNameAvailabilityOptions`
- New struct `ManagedHsmsClientCheckMhsmNameAvailabilityResponse`
- New struct `RotationPolicy`
- New struct `Trigger`
- New field `ID` in struct `MHSMPrivateEndpointConnectionItem`
- New field `Etag` in struct `MHSMPrivateEndpointConnectionItem`
- New field `ReleasePolicy` in struct `KeyProperties`
- New field `RotationPolicy` in struct `KeyProperties`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).