# Release History

## 2.0.0 (2022-06-02)
### Breaking Changes

- Type of `GalleryProperties.ProvisioningState` has been changed from `*GalleryPropertiesProvisioningState` to `*GalleryProvisioningState`
- Type of `GalleryImageProperties.ProvisioningState` has been changed from `*GalleryImagePropertiesProvisioningState` to `*GalleryProvisioningState`
- Type of `GalleryApplicationVersionProperties.ProvisioningState` has been changed from `*GalleryApplicationVersionPropertiesProvisioningState` to `*GalleryProvisioningState`
- Type of `GalleryImageVersionProperties.ProvisioningState` has been changed from `*GalleryImageVersionPropertiesProvisioningState` to `*GalleryProvisioningState`
- Type of `VirtualMachineScaleSetIdentity.UserAssignedIdentities` has been changed from `map[string]*VirtualMachineScaleSetIdentityUserAssignedIdentitiesValue` to `map[string]*UserAssignedIdentitiesValue`
- Const `GalleryPropertiesProvisioningStateCreating` has been removed
- Const `GalleryApplicationVersionPropertiesProvisioningStateDeleting` has been removed
- Const `GalleryImageVersionPropertiesProvisioningStateDeleting` has been removed
- Const `GalleryApplicationVersionPropertiesProvisioningStateFailed` has been removed
- Const `SharingProfileGroupTypesCommunity` has been removed
- Const `GalleryImagePropertiesProvisioningStateMigrating` has been removed
- Const `GalleryImagePropertiesProvisioningStateCreating` has been removed
- Const `GalleryImageVersionPropertiesProvisioningStateUpdating` has been removed
- Const `GalleryPropertiesProvisioningStateDeleting` has been removed
- Const `GalleryApplicationVersionPropertiesProvisioningStateCreating` has been removed
- Const `GalleryPropertiesProvisioningStateUpdating` has been removed
- Const `GalleryImageVersionPropertiesProvisioningStateSucceeded` has been removed
- Const `GalleryPropertiesProvisioningStateMigrating` has been removed
- Const `GalleryPropertiesProvisioningStateFailed` has been removed
- Const `GalleryImagePropertiesProvisioningStateSucceeded` has been removed
- Const `GalleryImageVersionPropertiesProvisioningStateFailed` has been removed
- Const `GalleryImagePropertiesProvisioningStateFailed` has been removed
- Const `GalleryApplicationVersionPropertiesProvisioningStateSucceeded` has been removed
- Const `GalleryImageVersionPropertiesProvisioningStateCreating` has been removed
- Const `GalleryImagePropertiesProvisioningStateDeleting` has been removed
- Const `GalleryImagePropertiesProvisioningStateUpdating` has been removed
- Const `GalleryApplicationVersionPropertiesProvisioningStateUpdating` has been removed
- Const `GalleryImageVersionPropertiesProvisioningStateMigrating` has been removed
- Const `GalleryApplicationVersionPropertiesProvisioningStateMigrating` has been removed
- Const `GalleryPropertiesProvisioningStateSucceeded` has been removed
- Function `PossibleGalleryImageVersionPropertiesProvisioningStateValues` has been removed
- Function `PossibleGalleryApplicationVersionPropertiesProvisioningStateValues` has been removed
- Function `PossibleGalleryPropertiesProvisioningStateValues` has been removed
- Function `PossibleGalleryImagePropertiesProvisioningStateValues` has been removed
- Struct `VirtualMachineScaleSetIdentityUserAssignedIdentitiesValue` has been removed

### Features Added

- New const `GallerySharingPermissionTypesCommunity`
- New const `GalleryProvisioningStateUpdating`
- New const `DiskEncryptionSetIdentityTypeUserAssigned`
- New const `SharedGalleryHostCachingReadWrite`
- New const `GalleryProvisioningStateMigrating`
- New const `SharedGalleryHostCachingNone`
- New const `CopyCompletionErrorReasonCopySourceNotFound`
- New const `SharedGalleryHostCachingReadOnly`
- New const `GalleryProvisioningStateFailed`
- New const `DiskStorageAccountTypesPremiumV2LRS`
- New const `GalleryProvisioningStateCreating`
- New const `GalleryProvisioningStateSucceeded`
- New const `DiskEncryptionSetIdentityTypeSystemAssignedUserAssigned`
- New const `GalleryProvisioningStateDeleting`
- New function `PossibleCopyCompletionErrorReasonValues() []CopyCompletionErrorReason`
- New function `*VirtualMachineImagesClient.ListByEdgeZone(context.Context, string, string, *VirtualMachineImagesClientListByEdgeZoneOptions) (VirtualMachineImagesClientListByEdgeZoneResponse, error)`
- New function `PossibleGalleryProvisioningStateValues() []GalleryProvisioningState`
- New function `*CommunityGalleryImagesClient.NewListPager(string, string, *CommunityGalleryImagesClientListOptions) *runtime.Pager[CommunityGalleryImagesClientListResponse]`
- New function `*CommunityGalleryImageVersionsClient.NewListPager(string, string, string, *CommunityGalleryImageVersionsClientListOptions) *runtime.Pager[CommunityGalleryImageVersionsClientListResponse]`
- New function `PossibleSharedGalleryHostCachingValues() []SharedGalleryHostCaching`
- New struct `CommunityGalleryImageList`
- New struct `CommunityGalleryImageVersionList`
- New struct `CommunityGalleryImageVersionsClientListOptions`
- New struct `CommunityGalleryImageVersionsClientListResponse`
- New struct `CommunityGalleryImagesClientListOptions`
- New struct `CommunityGalleryImagesClientListResponse`
- New struct `CopyCompletionError`
- New struct `SharedGalleryDataDiskImage`
- New struct `SharedGalleryDiskImage`
- New struct `SharedGalleryImageVersionStorageProfile`
- New struct `SharedGalleryOSDiskImage`
- New struct `UserArtifactSettings`
- New struct `VMImagesInEdgeZoneListResult`
- New struct `VirtualMachineImagesClientListByEdgeZoneOptions`
- New struct `VirtualMachineImagesClientListByEdgeZoneResponse`
- New field `StorageProfile` in struct `SharedGalleryImageVersionProperties`
- New field `ExcludeFromLatest` in struct `SharedGalleryImageVersionProperties`
- New field `ExcludeFromLatest` in struct `CommunityGalleryImageVersionProperties`
- New field `StorageProfile` in struct `CommunityGalleryImageVersionProperties`
- New field `CopyCompletionError` in struct `SnapshotProperties`
- New field `Architecture` in struct `SharedGalleryImageProperties`
- New field `FederatedClientID` in struct `EncryptionSetProperties`
- New field `AdvancedSettings` in struct `GalleryApplicationVersionPublishingProfile`
- New field `Settings` in struct `GalleryApplicationVersionPublishingProfile`
- New field `SecurityProfile` in struct `DiskRestorePointProperties`
- New field `UserAssignedIdentities` in struct `EncryptionSetIdentity`
- New field `FederatedClientID` in struct `DiskEncryptionSetUpdateProperties`
- New field `Architecture` in struct `CommunityGalleryImageProperties`
- New field `Eula` in struct `CommunityGalleryImageProperties`
- New field `PrivacyStatementURI` in struct `CommunityGalleryImageProperties`
- New field `CommunityGalleryImageID` in struct `ImageDiskReference`
- New field `SharedGalleryImageID` in struct `ImageDiskReference`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).