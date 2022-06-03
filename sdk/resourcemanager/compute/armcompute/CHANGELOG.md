# Release History

## 2.0.0 (2022-06-03)
### Breaking Changes

- Type of `CloudServiceExtensionProperties.Settings` has been changed from `*string` to `interface{}`
- Type of `CloudServiceExtensionProperties.ProtectedSettings` has been changed from `*string` to `interface{}`
- Type of `DiskRestorePointReplicationStatus.Status` has been changed from `*InstanceViewStatus` to `interface{}`
- Type of `DiskRestorePointInstanceView.ReplicationStatus` has been changed from `*DiskRestorePointReplicationStatus` to `interface{}`
- Const `LinuxVMGuestPatchAutomaticByPlatformRebootSettingIfRequired` has been removed
- Const `WindowsVMGuestPatchAutomaticByPlatformRebootSettingUnknown` has been removed
- Const `LinuxVMGuestPatchAutomaticByPlatformRebootSettingAlways` has been removed
- Const `WindowsVMGuestPatchAutomaticByPlatformRebootSettingIfRequired` has been removed
- Const `LinuxVMGuestPatchAutomaticByPlatformRebootSettingUnknown` has been removed
- Const `WindowsVMGuestPatchAutomaticByPlatformRebootSettingAlways` has been removed
- Const `StorageAccountTypesPremiumV2LRS` has been removed
- Const `WindowsVMGuestPatchAutomaticByPlatformRebootSettingNever` has been removed
- Const `LinuxVMGuestPatchAutomaticByPlatformRebootSettingNever` has been removed
- Function `PossibleLinuxVMGuestPatchAutomaticByPlatformRebootSettingValues` has been removed
- Function `PossibleWindowsVMGuestPatchAutomaticByPlatformRebootSettingValues` has been removed
- Struct `DedicatedHostGroupPropertiesAdditionalCapabilities` has been removed
- Struct `LinuxVMGuestPatchAutomaticByPlatformSettings` has been removed
- Struct `ProximityPlacementGroupPropertiesIntent` has been removed
- Struct `ResourceWithOptionalLocation` has been removed
- Struct `WindowsVMGuestPatchAutomaticByPlatformSettings` has been removed
- Field `AutomaticByPlatformSettings` of struct `LinuxPatchSettings` has been removed
- Field `DeleteOption` of struct `VirtualMachineScaleSetUpdateOSDisk` has been removed
- Field `UseRollingUpgradePolicy` of struct `AutomaticOSUpgradePolicy` has been removed
- Field `Zones` of struct `ProximityPlacementGroup` has been removed
- Field `DeleteOption` of struct `VirtualMachineScaleSetDataDisk` has been removed
- Field `TreatFailureAsDeploymentFailure` of struct `VMGalleryApplication` has been removed
- Field `EnableAutomaticUpgrade` of struct `VMGalleryApplication` has been removed
- Field `AdditionalCapabilities` of struct `DedicatedHostGroupProperties` has been removed
- Field `DeleteOption` of struct `VirtualMachineScaleSetOSDisk` has been removed
- Field `CompletionPercent` of struct `DiskRestorePointReplicationStatus` has been removed
- Field `Intent` of struct `ProximityPlacementGroupProperties` has been removed
- Field `Identity` of struct `VirtualMachineScaleSetVM` has been removed
- Field `AutomaticByPlatformSettings` of struct `PatchSettings` has been removed

### Features Added

- New struct `SystemData`
- New field `SystemData` in struct `CloudService`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).