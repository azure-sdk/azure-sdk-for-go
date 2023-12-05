# Release History

## 2.0.0 (2023-12-05)
### Breaking Changes

- Function `*UsersClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, User, *UsersClientBeginCreateOrUpdateOptions)` to `(context.Context, string, string, string, *UsersClientBeginCreateOrUpdateOptions)`
- Type of `LabVirtualMachineCreationParameterProperties.StorageType` has been changed from `*string` to `*StorageType`
- Type of `LabVirtualMachineProperties.StorageType` has been changed from `*string` to `*StorageTypes`
- Operation `*ServiceRunnersClient.CreateOrUpdate` has been changed to LRO, use `*ServiceRunnersClient.BeginCreateOrUpdate` instead.
- Operation `*ServiceRunnersClient.Delete` has been changed to LRO, use `*ServiceRunnersClient.BeginDelete` instead.
- Struct `ApplicableScheduleFragment` has been removed
- Struct `ShutdownNotificationContent` has been removed

### Features Added

- New value `HTTPStatusCodeAlreadyReported`, `HTTPStatusCodeEarlyHints`, `HTTPStatusCodeFailedDependency`, `HTTPStatusCodeIMUsed`, `HTTPStatusCodeInsufficientStorage`, `HTTPStatusCodeLocked`, `HTTPStatusCodeLoopDetected`, `HTTPStatusCodeMisdirectedRequest`, `HTTPStatusCodeMultiStatus`, `HTTPStatusCodeNetworkAuthenticationRequired`, `HTTPStatusCodeNotExtended`, `HTTPStatusCodePermanentRedirect`, `HTTPStatusCodePreconditionRequired`, `HTTPStatusCodeProcessing`, `HTTPStatusCodeRequestHeaderFieldsTooLarge`, `HTTPStatusCodeTooManyRequests`, `HTTPStatusCodeUnavailableForLegalReasons`, `HTTPStatusCodeUnprocessableEntity`, `HTTPStatusCodeVariantAlsoNegotiates` added to enum type `HTTPStatusCode`
- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `CheckNameAvailabilityReason` with values `CheckNameAvailabilityReasonAlreadyExists`, `CheckNameAvailabilityReasonInvalid`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `EnableState` with values `EnableStateDisabled`, `EnableStateEnabled`
- New enum type `EncryptionStatus` with values `EncryptionStatusDisabled`, `EncryptionStatusEnabled`
- New enum type `EncryptionType` with values `EncryptionTypeEncryptionAtRestWithCustomerKey`, `EncryptionTypeEncryptionAtRestWithPlatformKey`
- New enum type `ImageType` with values `ImageTypeGeneralized`, `ImageTypeSpecialized`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New enum type `OsType` with values `OsTypeLinux`, `OsTypeWindows`
- New enum type `SKUTier` with values `SKUTierBasic`, `SKUTierFree`, `SKUTierPremium`, `SKUTierStandard`
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
- New function `*LabSecretsClient.Update(context.Context, string, string, string, SecretFragment, *LabSecretsClientUpdateOptions) (LabSecretsClientUpdateResponse, error)`
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
- New struct `AzureEntityResource`
- New struct `BastionHost`
- New struct `BastionHostFragment`
- New struct `BastionHostList`
- New struct `BastionHostProperties`
- New struct `CheckNameAvailabilityRequest`
- New struct `CheckNameAvailabilityResponse`
- New struct `Cost`
- New struct `CustomImagePropertiesFromPlanFragment`
- New struct `Encryption`
- New struct `EncryptionProperties`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New struct `Identity`
- New struct `ImageVersionProperties`
- New struct `KeyVaultProperties`
- New struct `LabCostList`
- New struct `LabSecret`
- New struct `LabSecretFragment`
- New struct `LabSecretList`
- New struct `LabSecretProperties`
- New struct `LocationData`
- New struct `Operation`
- New struct `OperationDisplay`
- New struct `OperationListResult`
- New struct `OperationStatusResult`
- New struct `Plan`
- New struct `PolicySet`
- New struct `PolicySetList`
- New struct `PolicySetProperties`
- New struct `ProxyResource`
- New struct `ResourceModelWithAllowedPropertySet`
- New struct `ResourceModelWithAllowedPropertySetIdentity`
- New struct `ResourceModelWithAllowedPropertySetPlan`
- New struct `ResourceModelWithAllowedPropertySetSKU`
- New struct `SKU`
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
- New struct `WeekDetailsFragment`
- New field `SystemData` in struct `ApplicableSchedule`
- New field `SystemData` in struct `ArmTemplate`
- New field `SystemData` in struct `Artifact`
- New field `SystemData` in struct `ArtifactSource`
- New field `SystemData` in struct `CustomImage`
- New field `SystemData` in struct `Disk`
- New field `SystemData` in struct `DtlEnvironment`
- New field `SystemData` in struct `Formula`
- New field `SystemData` in struct `GalleryImage`
- New field `UserAssignedIdentities` in struct `IdentityProperties`
- New field `Identity`, `SystemData` in struct `Lab`
- New field `SystemData` in struct `LabCost`
- New field `BrowserConnect`, `DefaultSecretName`, `DisableAutoUpgradeCseMinorVersion`, `Encryption`, `IsolateLabResources`, `ManagementIdentities` in struct `LabProperties`
- New field `SystemData` in struct `LabVirtualMachine`
- New field `ApplicableSchedule`, `ArtifactDeploymentStatus`, `CanApplyArtifacts`, `ComputeID`, `ComputeVM`, `CreatedByUser`, `CreatedByUserID`, `Fqdn`, `GalleryImageVersionID`, `LastKnownPowerState`, `OSDiskSizeGb`, `OSType`, `ProvisioningState`, `SharedImageID`, `SharedImageVersion`, `UniqueIdentifier`, `VirtualMachineCreationSource` in struct `LabVirtualMachineCreationParameterProperties`
- New field `CanApplyArtifacts`, `GalleryImageVersionID`, `OSDiskSizeGb`, `SharedImageID`, `SharedImageVersion` in struct `LabVirtualMachineProperties`
- New field `SystemData` in struct `NotificationChannel`
- New field `SystemData` in struct `Policy`
- New field `SystemData` in struct `Schedule`
- New field `CreatedDate`, `ProvisioningState`, `UniqueIdentifier` in struct `ScheduleCreationParameterProperties`
- New field `SystemData` in struct `Secret`
- New field `SystemData` in struct `ServiceFabric`
- New field `Properties`, `SystemData` in struct `ServiceRunner`
- New field `SystemData` in struct `User`
- New field `User` in struct `UsersClientBeginCreateOrUpdateOptions`
- New field `SystemData` in struct `VirtualNetwork`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


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