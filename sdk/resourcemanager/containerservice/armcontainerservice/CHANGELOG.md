# Release History

## 2.0.0-beta.3 (2022-07-25)
### Features Added

- New const `UpdateModeRecreate`
- New const `ControlledValuesRequestsOnly`
- New const `UpdateModeOff`
- New const `UpdateModeAuto`
- New const `ControlledValuesRequestsAndLimits`
- New const `UpdateModeInitial`
- New function `PossibleControlledValuesValues() []ControlledValues`
- New function `PossibleUpdateModeValues() []UpdateMode`
- New struct `ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler`
- New field `VerticalPodAutoscaler` in struct `ManagedClusterWorkloadAutoScalerProfile`


## 2.0.0 (2022-07-22)
### Breaking Changes

- Struct `ManagedClusterSecurityProfileAzureDefender` has been removed
- Field `AzureDefender` of struct `ManagedClusterSecurityProfile` has been removed

### Features Added

- New const `KeyVaultNetworkAccessTypesPrivate`
- New const `NetworkPluginNone`
- New const `KeyVaultNetworkAccessTypesPublic`
- New function `PossibleKeyVaultNetworkAccessTypesValues() []KeyVaultNetworkAccessTypes`
- New struct `AzureKeyVaultKms`
- New struct `ManagedClusterSecurityProfileDefender`
- New struct `ManagedClusterSecurityProfileDefenderSecurityMonitoring`
- New field `HostGroupID` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `HostGroupID` in struct `ManagedClusterAgentPoolProfile`
- New field `AzureKeyVaultKms` in struct `ManagedClusterSecurityProfile`
- New field `Defender` in struct `ManagedClusterSecurityProfile`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).