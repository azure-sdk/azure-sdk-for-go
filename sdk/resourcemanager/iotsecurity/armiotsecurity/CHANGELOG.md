# Release History

## 0.7.0 (2023-08-23)
### Breaking Changes

- Type of `PackageDownloads.Snmp` has been changed from `[]*PackageDownloadInfo` to `[]*PackageBaseDownloadInfo`
- Type of `PackageDownloads.ThreatIntelligence` has been changed from `[]*PackageDownloadInfo` to `[]*PackageBaseDownloadInfo`
- Type of `PackageDownloadsSensor.Full` has been changed from `*PackageDownloadsSensorFull` to `*FullPackageDownloadInfo`
- Struct `PackageDownloadsCentralManagerFullOvf` has been removed
- Struct `PackageDownloadsSensorFull` has been removed
- Struct `PackageDownloadsSensorFullOvf` has been removed
- Field `AuthorizedDevicesImportTemplate`, `DeviceInformationUpdateImportTemplate`, `WmiTool` of struct `PackageDownloads` has been removed
- Field `Ovf` of struct `PackageDownloadsCentralManagerFull` has been removed

### Features Added

- New enum type `Scope` with values `ScopeMajor`, `ScopeMinor`, `ScopePatch`
- New struct `EnterprisePackageDownloadInfo`
- New struct `FullPackageDownloadInfo`
- New struct `PackageBaseDownloadInfo`
- New field `ReleaseDate`, `Scope`, `SupportedUntil` in struct `PackageDownloadInfo`
- New field `Enterprise` in struct `PackageDownloads`
- New field `ReleaseDate`, `Scope`, `SupportedUntil` in struct `UpgradePackageDownloadInfo`


## 0.6.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 0.6.0 (2023-04-03)
### Breaking Changes

- Function `NewDeviceGroupsClient` parameter(s) have been changed from `(string, string, azcore.TokenCredential, *arm.ClientOptions)` to `(string, azcore.TokenCredential, *arm.ClientOptions)`
- Function `*DeviceGroupsClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, DeviceGroupModel, *DeviceGroupsClientCreateOrUpdateOptions)` to `(context.Context, string, string, DeviceGroupModel, *DeviceGroupsClientCreateOrUpdateOptions)`
- Function `*DeviceGroupsClient.Delete` parameter(s) have been changed from `(context.Context, string, *DeviceGroupsClientDeleteOptions)` to `(context.Context, string, string, *DeviceGroupsClientDeleteOptions)`
- Function `*DeviceGroupsClient.Get` parameter(s) have been changed from `(context.Context, string, *DeviceGroupsClientGetOptions)` to `(context.Context, string, string, *DeviceGroupsClientGetOptions)`
- Function `*DeviceGroupsClient.NewListPager` parameter(s) have been changed from `(*DeviceGroupsClientListOptions)` to `(string, *DeviceGroupsClientListOptions)`
- Function `NewDevicesClient` parameter(s) have been changed from `(string, string, azcore.TokenCredential, *arm.ClientOptions)` to `(string, azcore.TokenCredential, *arm.ClientOptions)`
- Function `*DevicesClient.Get` parameter(s) have been changed from `(context.Context, string, string, *DevicesClientGetOptions)` to `(context.Context, string, string, string, *DevicesClientGetOptions)`
- Function `*DevicesClient.NewListPager` parameter(s) have been changed from `(string, *DevicesClientListOptions)` to `(string, string, *DevicesClientListOptions)`
- Function `NewLocationsClient` parameter(s) have been changed from `(string, string, azcore.TokenCredential, *arm.ClientOptions)` to `(string, azcore.TokenCredential, *arm.ClientOptions)`
- Function `*LocationsClient.Get` parameter(s) have been changed from `(context.Context, *LocationsClientGetOptions)` to `(context.Context, string, *LocationsClientGetOptions)`

### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).