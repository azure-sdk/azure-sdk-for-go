# Release History

## 2.0.0-beta.1 (2022-08-17)
### Breaking Changes

- Type of `MachineExtensionProperties.ProtectedSettings` has been changed from `interface{}` to `map[string]interface{}`
- Type of `MachineExtensionProperties.Settings` has been changed from `interface{}` to `map[string]interface{}`
- Type of `MachineExtensionUpdateProperties.Settings` has been changed from `interface{}` to `map[string]interface{}`
- Type of `MachineExtensionUpdateProperties.ProtectedSettings` has been changed from `interface{}` to `map[string]interface{}`
- Field `Extensions` of struct `MachineProperties` has been removed

### Features Added

- New const `AgentConfigurationModeMonitor`
- New const `AgentConfigurationModeFull`
- New function `PossibleAgentConfigurationModeValues() []AgentConfigurationMode`
- New field `Resources` in struct `Machine`
- New field `ConfigMode` in struct `AgentConfiguration`
- New field `EnableAutomaticUpgrade` in struct `MachineExtensionUpdateProperties`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).