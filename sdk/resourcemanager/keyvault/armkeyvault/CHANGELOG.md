# Release History

## 2.0.0 (2022-09-21)
### Breaking Changes

- Type of `CheckNameAvailabilityResult.Reason` has been changed from `*Reason` to `*ReasonForKeyVault`
- Const `ReasonAlreadyExists` has been removed
- Const `ReasonAccountNameInvalid` has been removed
- Type alias `Reason` has been removed
- Function `PossibleReasonValues` has been removed
- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New const `KeyRotationPolicyActionTypeRotate`
- New const `ReasonForManagedHsmAccountNameInvalid`
- New const `KeyRotationPolicyActionTypeNotify`
- New const `ActivationStatusFailed`
- New const `ReasonForKeyVaultAccountNameInvalid`
- New const `ReasonForManagedHsmAlreadyExists`
- New const `KeyPermissionsSetrotationpolicy`
- New const `ActivationStatusActive`
- New const `ReasonForKeyVaultAlreadyExists`
- New const `ActivationStatusNotActivated`
- New const `ActivationStatusUnknown`
- New const `KeyPermissionsRelease`
- New const `KeyPermissionsGetrotationpolicy`
- New const `JSONWebKeyOperationRelease`
- New const `KeyPermissionsRotate`
- New type alias `ReasonForManagedHsm`
- New type alias `ReasonForKeyVault`
- New type alias `KeyRotationPolicyActionType`
- New type alias `ActivationStatus`
- New function `PossibleReasonForKeyVaultValues() []ReasonForKeyVault`
- New function `PossibleActivationStatusValues() []ActivationStatus`
- New function `PossibleReasonForManagedHsmValues() []ReasonForManagedHsm`
- New function `*ManagedHsmsClient.CheckMhsmNameAvailability(context.Context, CheckMhsmNameAvailabilityParameters, *ManagedHsmsClientCheckMhsmNameAvailabilityOptions) (ManagedHsmsClientCheckMhsmNameAvailabilityResponse, error)`
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
- New field `RotationPolicy` in struct `KeyProperties`
- New field `ReleasePolicy` in struct `KeyProperties`
- New field `Etag` in struct `MHSMPrivateEndpointConnectionItem`
- New field `ID` in struct `MHSMPrivateEndpointConnectionItem`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).