# Release History

## 2.0.0 (2023-01-23)
### Breaking Changes

- Type of `ErrorAdditionalInfo.Info` has been changed from `interface{}` to `any`
- Type of `MachineExtensionProperties.ProtectedSettings` has been changed from `interface{}` to `map[string]any`
- Type of `MachineExtensionProperties.Settings` has been changed from `interface{}` to `map[string]any`
- Type of `MachineExtensionUpdateProperties.Settings` has been changed from `interface{}` to `map[string]any`
- Type of `MachineExtensionUpdateProperties.ProtectedSettings` has been changed from `interface{}` to `map[string]any`

### Features Added

- New type alias `AgentConfigurationMode` with values `AgentConfigurationModeFull`, `AgentConfigurationModeMonitor`
- New function `NewExtensionMetadataClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ExtensionMetadataClient, error)`
- New function `*ExtensionMetadataClient.Get(context.Context, string, string, string, string, *ExtensionMetadataClientGetOptions) (ExtensionMetadataClientGetResponse, error)`
- New function `*ExtensionMetadataClient.NewListPager(string, string, string, *ExtensionMetadataClientListOptions) *runtime.Pager[ExtensionMetadataClientListResponse]`
- New struct `ExtensionMetadataClient`
- New struct `ExtensionMetadataClientListResponse`
- New struct `ExtensionValue`
- New struct `ExtensionValueListResult`
- New struct `ExtensionValueProperties`
- New field `ConfigMode` in struct `AgentConfiguration`
- New field `Resources` in struct `Machine`
- New field `EnableAutomaticUpgrade` in struct `MachineExtensionUpdateProperties`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `TrackedResource`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).