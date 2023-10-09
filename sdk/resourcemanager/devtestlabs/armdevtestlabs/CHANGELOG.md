# Release History

## 2.0.0 (2023-10-09)
### Breaking Changes

- Function `*UsersClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, User, *UsersClientBeginCreateOrUpdateOptions)` to `(context.Context, string, string, string, *UsersClientBeginCreateOrUpdateOptions)`
- Type of `LabVirtualMachineCreationParameterProperties.StorageType` has been changed from `*string` to `*StorageType`
- Type of `LabVirtualMachineProperties.StorageType` has been changed from `*string` to `*StorageTypes`
- `NotificationChannelEventTypeCost` from enum `NotificationChannelEventType` has been removed
- `PolicyFactNameLabTargetCost` from enum `PolicyFactName` has been removed
- Enum `CostThresholdStatus` has been removed
- Enum `CostType` has been removed
- Enum `ReportingCycleType` has been removed
- Enum `TargetCostStatus` has been removed
- Function `*ClientFactory.NewCostsClient` has been removed
- Function `NewCostsClient` has been removed
- Function `*CostsClient.CreateOrUpdate` has been removed
- Function `*CostsClient.Get` has been removed
- Operation `*ServiceRunnersClient.Delete` has been changed to LRO, use `*ServiceRunnersClient.BeginDelete` instead.
- Struct `ApplicableScheduleFragment` has been removed
- Struct `CostThresholdProperties` has been removed
- Struct `LabCost` has been removed
- Struct `LabCostDetailsProperties` has been removed
- Struct `LabCostProperties` has been removed
- Struct `LabCostSummaryProperties` has been removed
- Struct `LabResourceCostProperties` has been removed
- Struct `PercentageCostThresholdProperties` has been removed
- Struct `ShutdownNotificationContent` has been removed
- Struct `TargetCostProperties` has been removed

### Features Added

- New value `HTTPStatusCodeAlreadyReported`, `HTTPStatusCodeEarlyHints`, `HTTPStatusCodeFailedDependency`, `HTTPStatusCodeIMUsed`, `HTTPStatusCodeInsufficientStorage`, `HTTPStatusCodeLocked`, `HTTPStatusCodeLoopDetected`, `HTTPStatusCodeMisdirectedRequest`, `HTTPStatusCodeMultiStatus`, `HTTPStatusCodeNetworkAuthenticationRequired`, `HTTPStatusCodeNotExtended`, `HTTPStatusCodePermanentRedirect`, `HTTPStatusCodePreconditionRequired`, `HTTPStatusCodeProcessing`, `HTTPStatusCodeRequestHeaderFieldsTooLarge`, `HTTPStatusCodeTooManyRequests`, `HTTPStatusCodeUnavailableForLegalReasons`, `HTTPStatusCodeUnprocessableEntity`, `HTTPStatusCodeVariantAlsoNegotiates` added to enum type `HTTPStatusCode`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `EnableState` with values `EnableStateDisabled`, `EnableStateEnabled`
- New enum type `EncryptionType` with values `EncryptionTypeEncryptionAtRestWithCustomerKey`, `EncryptionTypeEncryptionAtRestWithPlatformKey`
- New enum type `ImageType` with values `ImageTypeGeneralized`, `ImageTypeSpecialized`
- New enum type `OsType` with values `OsTypeLinux`, `OsTypeWindows`
- New enum type `SecurityTypes` with values `SecurityTypesConfidentialVM`, `SecurityTypesTrustedLaunch`
- New enum type `StorageTypes` with values `StorageTypesPremium`, `StorageTypesStandard`, `StorageTypesStandardSSD`
- New function `NewBastionHostsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*BastionHostsClient, error)`
- New function `*BastionHostsClient.BeginCreateOrUpdate(context.Context, string, string, string, string, BastionHost, *BastionHostsClientBeginCreateOrUpdateOptions) (*runtime.Poller[BastionHostsClientCreateOrUpdateResponse], error)`
- New function `*BastionHostsClient.BeginDelete(context.Context, string, string, string, string, *BastionHostsClientBeginDeleteOptions) (*runtime.Poller[BastionHostsClientDeleteResponse], error)`
- New function `*BastionHostsClient.Get(context.Context, string, string, string, string, *BastionHostsClientGetOptions) (BastionHostsClientGetResponse, error)`
- New function `*BastionHostsClient.NewListPager(string, string, string, *BastionHostsClientListOptions) *runtime.Pager[BastionHostsClientListResponse]`
- New function `*BastionHostsClient.Update(context.Context, string, string, string, string, BastionHostFragment, *BastionHostsClientUpdateOptions) (BastionHostsClientUpdateResponse, error)`
- New function `*ClientFactory.NewBastionHostsClient() *BastionHostsClient`
- New function `*ClientFactory.NewLabSecretsClient() *LabSecretsClient`
- New function `*ClientFactory.NewSharedGalleriesClient() *SharedGalleriesClient`
- New function `*ClientFactory.NewSharedImagesClient() *SharedImagesClient`
- New function `*GalleryImagesClient.Get(context.Context, string, string, string, *GalleryImagesClientGetOptions) (GalleryImagesClientGetResponse, error)`
- New function `NewLabSecretsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*LabSecretsClient, error)`
- New function `*LabSecretsClient.BeginCreateOrUpdate(context.Context, string, string, string, LabSecret, *LabSecretsClientBeginCreateOrUpdateOptions) (*runtime.Poller[LabSecretsClientCreateOrUpdateResponse], error)`
- New function `*LabSecretsClient.BeginDelete(context.Context, string, string, string, *LabSecretsClientBeginDeleteOptions) (*runtime.Poller[LabSecretsClientDeleteResponse], error)`
- New function `*LabSecretsClient.Get(context.Context, string, string, string, *LabSecretsClientGetOptions) (LabSecretsClientGetResponse, error)`
- New function `*LabSecretsClient.NewListPager(string, string, *LabSecretsClientListOptions) *runtime.Pager[LabSecretsClientListResponse]`
- New function `*LabSecretsClient.BeginUpdate(context.Context, string, string, string, SecretFragment, *LabSecretsClientBeginUpdateOptions) (*runtime.Poller[LabSecretsClientUpdateResponse], error)`
- New function `*LabsClient.EnsureCurrentUserProfile(context.Context, string, string, *LabsClientEnsureCurrentUserProfileOptions) (LabsClientEnsureCurrentUserProfileResponse, error)`
- New function `*PolicySetsClient.NewListPager(string, string, *PolicySetsClientListOptions) *runtime.Pager[PolicySetsClientListResponse]`
- New function `*ServiceRunnersClient.NewListPager(string, string, *ServiceRunnersClientListOptions) *runtime.Pager[ServiceRunnersClientListResponse]`
- New function `NewSharedGalleriesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SharedGalleriesClient, error)`
- New function `*SharedGalleriesClient.CreateOrUpdate(context.Context, string, string, string, SharedGallery, *SharedGalleriesClientCreateOrUpdateOptions) (SharedGalleriesClientCreateOrUpdateResponse, error)`
- New function `*SharedGalleriesClient.Delete(context.Context, string, string, string, *SharedGalleriesClientDeleteOptions) (SharedGalleriesClientDeleteResponse, error)`
- New function `*SharedGalleriesClient.Get(context.Context, string, string, string, *SharedGalleriesClientGetOptions) (SharedGalleriesClientGetResponse, error)`
- New function `*SharedGalleriesClient.NewListPager(string, string, *SharedGalleriesClientListOptions) *runtime.Pager[SharedGalleriesClientListResponse]`
- New function `*SharedGalleriesClient.Update(context.Context, string, string, string, SharedGalleryFragment, *SharedGalleriesClientUpdateOptions) (SharedGalleriesClientUpdateResponse, error)`
- New function `NewSharedImagesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SharedImagesClient, error)`
- New function `*SharedImagesClient.CreateOrUpdate(context.Context, string, string, string, string, SharedImage, *SharedImagesClientCreateOrUpdateOptions) (SharedImagesClientCreateOrUpdateResponse, error)`
- New function `*SharedImagesClient.Delete(context.Context, string, string, string, string, *SharedImagesClientDeleteOptions) (SharedImagesClientDeleteResponse, error)`
- New function `*SharedImagesClient.Get(context.Context, string, string, string, string, *SharedImagesClientGetOptions) (SharedImagesClientGetResponse, error)`
- New function `*SharedImagesClient.NewListPager(string, string, string, *SharedImagesClientListOptions) *runtime.Pager[SharedImagesClientListResponse]`
- New function `*SharedImagesClient.Update(context.Context, string, string, string, string, SharedImageFragment, *SharedImagesClientUpdateOptions) (SharedImagesClientUpdateResponse, error)`
- New function `*VirtualMachinesClient.ClearArtifactResults(context.Context, string, string, string, *VirtualMachinesClientClearArtifactResultsOptions) (VirtualMachinesClientClearArtifactResultsResponse, error)`
- New struct `BastionHost`
- New struct `BastionHostFragment`
- New struct `BastionHostList`
- New struct `BastionHostProperties`
- New struct `CustomImagePropertiesFromPlanFragment`
- New struct `Encryption`
- New struct `ImageVersionProperties`
- New struct `LabSecret`
- New struct `LabSecretFragment`
- New struct `LabSecretList`
- New struct `LabSecretProperties`
- New struct `PolicySet`
- New struct `PolicySetList`
- New struct `PolicySetProperties`
- New struct `SecurityProfile`
- New struct `ServiceRunnerProperties`
- New struct `SharedGallery`
- New struct `SharedGalleryFragment`
- New struct `SharedGalleryList`
- New struct `SharedGalleryProperties`
- New struct `SharedImage`
- New struct `SharedImageFragment`
- New struct `SharedImageList`
- New struct `SharedImageProperties`
- New struct `SystemData`
- New struct `UefiSettings`
- New struct `WeekDetailsFragment`
- New field `SystemData` in struct `ApplicableSchedule`
- New field `SystemData` in struct `ArmTemplate`
- New field `SystemData` in struct `Artifact`
- New field `SystemData` in struct `ArtifactSource`
- New field `Identity` in struct `ArtifactSourceFragment`
- New field `SystemData` in struct `CustomImage`
- New field `Identity` in struct `CustomImageFragment`
- New field `SystemData` in struct `Disk`
- New field `Identity` in struct `DiskFragment`
- New field `SystemData` in struct `DtlEnvironment`
- New field `Identity` in struct `DtlEnvironmentFragment`
- New field `SystemData` in struct `Formula`
- New field `Identity` in struct `FormulaFragment`
- New field `SystemData` in struct `GalleryImage`
- New field `UserAssignedIdentities` in struct `IdentityProperties`
- New field `Identity`, `SystemData` in struct `Lab`
- New field `Identity` in struct `LabFragment`
- New field `BrowserConnect`, `DefaultSecretName`, `DisableAutoUpgradeCseMinorVersion`, `Encryption`, `IsolateLabResources`, `ManagementIdentities` in struct `LabProperties`
- New field `SystemData` in struct `LabVirtualMachine`
- New field `ApplicableSchedule`, `ArtifactDeploymentStatus`, `CanApplyArtifacts`, `ComputeID`, `ComputeVM`, `CreatedByUser`, `CreatedByUserID`, `Fqdn`, `GalleryImageVersionID`, `LastKnownPowerState`, `OSDiskSizeGb`, `OSType`, `ProvisioningState`, `SecurityProfile`, `SharedImageID`, `SharedImageVersion`, `UniqueIdentifier`, `VirtualMachineCreationSource` in struct `LabVirtualMachineCreationParameterProperties`
- New field `Identity` in struct `LabVirtualMachineFragment`
- New field `CanApplyArtifacts`, `GalleryImageVersionID`, `OSDiskSizeGb`, `SecurityProfile`, `SharedImageID`, `SharedImageVersion` in struct `LabVirtualMachineProperties`
- New field `SystemData` in struct `NotificationChannel`
- New field `Identity` in struct `NotificationChannelFragment`
- New field `AzureAsyncOperation`, `Location` in struct `OperationsClientGetResponse`
- New field `SystemData` in struct `Policy`
- New field `Identity` in struct `PolicyFragment`
- New field `SystemData` in struct `Schedule`
- New field `CreatedDate`, `ProvisioningState`, `UniqueIdentifier` in struct `ScheduleCreationParameterProperties`
- New field `Identity` in struct `ScheduleFragment`
- New field `SystemData` in struct `Secret`
- New field `Identity` in struct `SecretFragment`
- New field `SystemData` in struct `ServiceFabric`
- New field `Identity` in struct `ServiceFabricFragment`
- New field `Properties`, `SystemData` in struct `ServiceRunner`
- New field `Identity` in struct `UpdateResource`
- New field `SystemData` in struct `User`
- New field `Identity` in struct `UserFragment`
- New field `User` in struct `UsersClientBeginCreateOrUpdateOptions`
- New field `SystemData` in struct `VirtualNetwork`
- New field `Identity` in struct `VirtualNetworkFragment`


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-28)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).