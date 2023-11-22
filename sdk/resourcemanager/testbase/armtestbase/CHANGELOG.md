# Release History

## 0.7.0 (2023-11-22)
### Breaking Changes

- Type of `ErrorResponse.Error` has been changed from `*ErrorDefinition` to `*ErrorDetail`
- Type of `Operation.Origin` has been changed from `*string` to `*Origin`
- Struct `ErrorDefinition` has been removed
- Field `Etag` of struct `AccountResource` has been removed
- Field `Properties` of struct `Operation` has been removed
- Field `Etag` of struct `PackageResource` has been removed
- Field `Etag` of struct `TrackedResource` has been removed

### Features Added

- New value `ActionFlowDrivenCustom` added to enum type `Action`
- New value `OsUpdateTypeInplaceUpgrade` added to enum type `OsUpdateType`
- New value `TestTypeFlowDrivenTest` added to enum type `TestType`
- New value `TypeInplaceUpgrade` added to enum type `Type`
- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `ApplicationType` with values `ApplicationTypeWinget`
- New enum type `Architecture` with values `ArchitectureArm64`, `ArchitectureX64`, `ArchitectureX86`
- New enum type `CredentialType` with values `CredentialTypeIntuneAccount`
- New enum type `DraftPackageSourceType` with values `DraftPackageSourceTypeGalleryApp`, `DraftPackageSourceTypeIntuneEnrollment`, `DraftPackageSourceTypeIntuneWin`, `DraftPackageSourceTypeNative`, `DraftPackageSourceTypeTestBasePackage`
- New enum type `Engagements` with values `EngagementsMAPP`, `EngagementsMVI`, `EngagementsMVP`, `EngagementsOther`, `EngagementsSUVP`
- New enum type `ExtractFileType` with values `ExtractFileTypeIntuneWinPackage`, `ExtractFileTypeTestBasePackage`
- New enum type `FileUploadResourceType` with values `FileUploadResourceTypePackage`, `FileUploadResourceTypeVHD`
- New enum type `FreeHourBalanceName` with values `FreeHourBalanceNameSubscriptionLevel`, `FreeHourBalanceNameTenantLevel`
- New enum type `FreeHourStatus` with values `FreeHourStatusEnabled`, `FreeHourStatusSuspended`
- New enum type `FreeHourType` with values `FreeHourTypePermanent`, `FreeHourTypeTemporary`
- New enum type `ImageArchitecture` with values `ImageArchitectureX64`
- New enum type `ImageOSState` with values `ImageOSStateGeneralized`, `ImageOSStateSpecialized`
- New enum type `ImageSecurityType` with values `ImageSecurityTypeStandard`, `ImageSecurityTypeTrustedLaunch`
- New enum type `ImageSource` with values `ImageSourceUnknown`, `ImageSourceVHD`
- New enum type `ImageStatus` with values `ImageStatusFailed`, `ImageStatusReady`, `ImageStatusUnknown`, `ImageStatusValidating`
- New enum type `InteropExecutionMode` with values `InteropExecutionModeFirstPartyApp`, `InteropExecutionModeFirstPartyAppWithTests`
- New enum type `IntuneExtractStatus` with values `IntuneExtractStatusExtractFailed`, `IntuneExtractStatusNoDependencyApp`, `IntuneExtractStatusReady`, `IntuneExtractStatusUploadFailed`, `IntuneExtractStatusUploading`
- New enum type `OrderBy` with values `OrderByPopularity`, `OrderByRelevance`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New enum type `OsProductState` with values `OsProductStateActive`, `OsProductStateDisabled`
- New enum type `PackageStudioTabs` with values `PackageStudioTabsBasicsTab`, `PackageStudioTabsConfigureTestTab`, `PackageStudioTabsEditPackageTab`, `PackageStudioTabsReviewAndCreateTab`, `PackageStudioTabsTagsTab`, `PackageStudioTabsTestMatrixTab`, `PackageStudioTabsUnspecified`
- New enum type `RequestStatus` with values `RequestStatusApproved`, `RequestStatusDeclined`, `RequestStatusInReview`
- New enum type `RequestTypes` with values `RequestTypesPreReleaseAccess`
- New enum type `SystemAssignedServiceIdentityType` with values `SystemAssignedServiceIdentityTypeNone`, `SystemAssignedServiceIdentityTypeSystemAssigned`
- New enum type `VHDStatus` with values `VHDStatusFailed`, `VHDStatusOccupied`, `VHDStatusReady`, `VHDStatusUnknown`, `VHDStatusVerifying`
- New enum type `VerificationStatus` with values `VerificationStatusFailed`, `VerificationStatusPassed`
- New function `NewActionRequestsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ActionRequestsClient, error)`
- New function `*ActionRequestsClient.Delete(context.Context, string, string, string, *ActionRequestsClientDeleteOptions) (ActionRequestsClientDeleteResponse, error)`
- New function `*ActionRequestsClient.Get(context.Context, string, string, string, *ActionRequestsClientGetOptions) (ActionRequestsClientGetResponse, error)`
- New function `*ActionRequestsClient.NewListPager(string, string, *ActionRequestsClientListOptions) *runtime.Pager[ActionRequestsClientListResponse]`
- New function `*ActionRequestsClient.Put(context.Context, string, string, string, *ActionRequestsClientPutOptions) (ActionRequestsClientPutResponse, error)`
- New function `NewAvailableInplaceUpgradeOSClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AvailableInplaceUpgradeOSClient, error)`
- New function `*AvailableInplaceUpgradeOSClient.Get(context.Context, string, string, string, *AvailableInplaceUpgradeOSClientGetOptions) (AvailableInplaceUpgradeOSClientGetResponse, error)`
- New function `*AvailableInplaceUpgradeOSClient.NewListPager(string, string, OsUpdateType, *AvailableInplaceUpgradeOSClientListOptions) *runtime.Pager[AvailableInplaceUpgradeOSClientListResponse]`
- New function `NewBillingHubServiceClient(string, azcore.TokenCredential, *arm.ClientOptions) (*BillingHubServiceClient, error)`
- New function `*BillingHubServiceClient.GetFreeHourBalance(context.Context, string, string, *BillingHubServiceClientGetFreeHourBalanceOptions) (BillingHubServiceClientGetFreeHourBalanceResponse, error)`
- New function `*BillingHubServiceClient.GetUsage(context.Context, string, string, *BillingHubServiceClientGetUsageOptions) (BillingHubServiceClientGetUsageResponse, error)`
- New function `NewChatSessionClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ChatSessionClient, error)`
- New function `*ChatSessionClient.BeginChat(context.Context, string, string, string, ChatRequest, *ChatSessionClientBeginChatOptions) (*runtime.Poller[ChatSessionClientChatResponse], error)`
- New function `NewChatSessionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ChatSessionsClient, error)`
- New function `*ChatSessionsClient.Get(context.Context, string, string, string, *ChatSessionsClientGetOptions) (ChatSessionsClientGetResponse, error)`
- New function `*ChatSessionsClient.NewListPager(string, string, *ChatSessionsClientListOptions) *runtime.Pager[ChatSessionsClientListResponse]`
- New function `*ClientFactory.NewActionRequestsClient() *ActionRequestsClient`
- New function `*ClientFactory.NewAvailableInplaceUpgradeOSClient() *AvailableInplaceUpgradeOSClient`
- New function `*ClientFactory.NewBillingHubServiceClient() *BillingHubServiceClient`
- New function `*ClientFactory.NewChatSessionClient() *ChatSessionClient`
- New function `*ClientFactory.NewChatSessionsClient() *ChatSessionsClient`
- New function `*ClientFactory.NewCredentialClient() *CredentialClient`
- New function `*ClientFactory.NewCredentialsClient() *CredentialsClient`
- New function `*ClientFactory.NewCustomImagesClient() *CustomImagesClient`
- New function `*ClientFactory.NewDraftPackagesClient() *DraftPackagesClient`
- New function `*ClientFactory.NewFeatureUpdateSupportedOsesClient() *FeatureUpdateSupportedOsesClient`
- New function `*ClientFactory.NewFirstPartyAppsClient() *FirstPartyAppsClient`
- New function `*ClientFactory.NewFreeHourBalancesClient() *FreeHourBalancesClient`
- New function `*ClientFactory.NewGalleryAppSKUsClient() *GalleryAppSKUsClient`
- New function `*ClientFactory.NewGalleryAppsClient() *GalleryAppsClient`
- New function `*ClientFactory.NewImageDefinitionsClient() *ImageDefinitionsClient`
- New function `*ClientFactory.NewVHDsClient() *VHDsClient`
- New function `NewCredentialClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CredentialClient, error)`
- New function `*CredentialClient.Get(context.Context, string, string, string, *CredentialClientGetOptions) (CredentialClientGetResponse, error)`
- New function `*CredentialClient.NewListByTestBaseAccountPager(string, string, *CredentialClientListByTestBaseAccountOptions) *runtime.Pager[CredentialClientListByTestBaseAccountResponse]`
- New function `*CredentialProperties.GetCredentialProperties() *CredentialProperties`
- New function `NewCredentialsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CredentialsClient, error)`
- New function `*CredentialsClient.Create(context.Context, string, string, string, CredentialResource, *CredentialsClientCreateOptions) (CredentialsClientCreateResponse, error)`
- New function `*CredentialsClient.Delete(context.Context, string, string, string, *CredentialsClientDeleteOptions) (CredentialsClientDeleteResponse, error)`
- New function `*CredentialsClient.Update(context.Context, string, string, string, CredentialResource, *CredentialsClientUpdateOptions) (CredentialsClientUpdateResponse, error)`
- New function `NewCustomImagesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CustomImagesClient, error)`
- New function `*CustomImagesClient.CheckImageNameAvailability(context.Context, string, string, ImageNameCheckAvailabilityParameters, *CustomImagesClientCheckImageNameAvailabilityOptions) (CustomImagesClientCheckImageNameAvailabilityResponse, error)`
- New function `*CustomImagesClient.BeginCreate(context.Context, string, string, string, CustomImageResource, *CustomImagesClientBeginCreateOptions) (*runtime.Poller[CustomImagesClientCreateResponse], error)`
- New function `*CustomImagesClient.BeginDelete(context.Context, string, string, string, *CustomImagesClientBeginDeleteOptions) (*runtime.Poller[CustomImagesClientDeleteResponse], error)`
- New function `*CustomImagesClient.Get(context.Context, string, string, string, *CustomImagesClientGetOptions) (CustomImagesClientGetResponse, error)`
- New function `*CustomImagesClient.NewListByTestBaseAccountPager(string, string, *CustomImagesClientListByTestBaseAccountOptions) *runtime.Pager[CustomImagesClientListByTestBaseAccountResponse]`
- New function `NewDraftPackagesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DraftPackagesClient, error)`
- New function `*DraftPackagesClient.BeginCopyFromPackage(context.Context, string, string, string, CopyFromPackageOperationParameters, *DraftPackagesClientBeginCopyFromPackageOptions) (*runtime.Poller[DraftPackagesClientCopyFromPackageResponse], error)`
- New function `*DraftPackagesClient.Create(context.Context, string, string, string, DraftPackageResource, *DraftPackagesClientCreateOptions) (DraftPackagesClientCreateResponse, error)`
- New function `*DraftPackagesClient.Delete(context.Context, string, string, string, *DraftPackagesClientDeleteOptions) (DraftPackagesClientDeleteResponse, error)`
- New function `*DraftPackagesClient.BeginExtractFile(context.Context, string, string, string, ExtractFileOperationParameters, *DraftPackagesClientBeginExtractFileOptions) (*runtime.Poller[DraftPackagesClientExtractFileResponse], error)`
- New function `*DraftPackagesClient.BeginGenerateFoldersAndScripts(context.Context, string, string, string, GenerateOperationParameters, *DraftPackagesClientBeginGenerateFoldersAndScriptsOptions) (*runtime.Poller[DraftPackagesClientGenerateFoldersAndScriptsResponse], error)`
- New function `*DraftPackagesClient.Get(context.Context, string, string, string, *DraftPackagesClientGetOptions) (DraftPackagesClientGetResponse, error)`
- New function `*DraftPackagesClient.GetPath(context.Context, string, string, string, *DraftPackagesClientGetPathOptions) (DraftPackagesClientGetPathResponse, error)`
- New function `*DraftPackagesClient.NewListByTestBaseAccountPager(string, string, *DraftPackagesClientListByTestBaseAccountOptions) *runtime.Pager[DraftPackagesClientListByTestBaseAccountResponse]`
- New function `*DraftPackagesClient.Update(context.Context, string, string, string, DraftPackageUpdateParameters, *DraftPackagesClientUpdateOptions) (DraftPackagesClientUpdateResponse, error)`
- New function `NewFeatureUpdateSupportedOsesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*FeatureUpdateSupportedOsesClient, error)`
- New function `*FeatureUpdateSupportedOsesClient.NewListPager(string, string, *FeatureUpdateSupportedOsesClientListOptions) *runtime.Pager[FeatureUpdateSupportedOsesClientListResponse]`
- New function `NewFirstPartyAppsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*FirstPartyAppsClient, error)`
- New function `*FirstPartyAppsClient.Get(context.Context, string, string, string, *FirstPartyAppsClientGetOptions) (FirstPartyAppsClientGetResponse, error)`
- New function `*FirstPartyAppsClient.NewListPager(string, string, *FirstPartyAppsClientListOptions) *runtime.Pager[FirstPartyAppsClientListResponse]`
- New function `NewFreeHourBalancesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*FreeHourBalancesClient, error)`
- New function `*FreeHourBalancesClient.Get(context.Context, string, string, FreeHourBalanceName, *FreeHourBalancesClientGetOptions) (FreeHourBalancesClientGetResponse, error)`
- New function `*FreeHourBalancesClient.NewListPager(string, string, *FreeHourBalancesClientListOptions) *runtime.Pager[FreeHourBalancesClientListResponse]`
- New function `*GalleryAppSKUProperties.GetGalleryAppSKUProperties() *GalleryAppSKUProperties`
- New function `NewGalleryAppSKUsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*GalleryAppSKUsClient, error)`
- New function `*GalleryAppSKUsClient.Get(context.Context, string, string, string, string, *GalleryAppSKUsClientGetOptions) (GalleryAppSKUsClientGetResponse, error)`
- New function `*GalleryAppSKUsClient.NewListPager(string, string, string, *GalleryAppSKUsClientListOptions) *runtime.Pager[GalleryAppSKUsClientListResponse]`
- New function `NewGalleryAppsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*GalleryAppsClient, error)`
- New function `*GalleryAppsClient.Get(context.Context, string, string, string, *GalleryAppsClientGetOptions) (GalleryAppsClientGetResponse, error)`
- New function `*GalleryAppsClient.NewListPager(string, string, *GalleryAppsClientListOptions) *runtime.Pager[GalleryAppsClientListResponse]`
- New function `NewImageDefinitionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ImageDefinitionsClient, error)`
- New function `*ImageDefinitionsClient.Create(context.Context, string, string, string, ImageDefinitionResource, *ImageDefinitionsClientCreateOptions) (ImageDefinitionsClientCreateResponse, error)`
- New function `*ImageDefinitionsClient.Delete(context.Context, string, string, string, *ImageDefinitionsClientDeleteOptions) (ImageDefinitionsClientDeleteResponse, error)`
- New function `*ImageDefinitionsClient.Get(context.Context, string, string, string, *ImageDefinitionsClientGetOptions) (ImageDefinitionsClientGetResponse, error)`
- New function `*ImageDefinitionsClient.NewListByTestBaseAccountPager(string, string, *ImageDefinitionsClientListByTestBaseAccountOptions) *runtime.Pager[ImageDefinitionsClientListByTestBaseAccountResponse]`
- New function `*IntuneSingletonResourceProperties.GetCredentialProperties() *CredentialProperties`
- New function `*PackagesClient.RunTest(context.Context, string, string, string, *PackagesClientRunTestOptions) (PackagesClientRunTestResponse, error)`
- New function `*TestResultsClient.GetConsoleLogDownloadURL(context.Context, string, string, string, string, TestResultConsoleLogDownloadURLParameters, *TestResultsClientGetConsoleLogDownloadURLOptions) (TestResultsClientGetConsoleLogDownloadURLResponse, error)`
- New function `NewVHDsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*VHDsClient, error)`
- New function `*VHDsClient.Delete(context.Context, string, string, string, *VHDsClientDeleteOptions) (VHDsClientDeleteResponse, error)`
- New function `*VHDsClient.Get(context.Context, string, string, string, *VHDsClientGetOptions) (VHDsClientGetResponse, error)`
- New function `*VHDsClient.NewListByTestBaseAccountPager(string, string, *VHDsClientListByTestBaseAccountOptions) *runtime.Pager[VHDsClientListByTestBaseAccountResponse]`
- New function `*WingetAppSKUProperties.GetGalleryAppSKUProperties() *GalleryAppSKUProperties`
- New struct `ActionRequest`
- New struct `ActionRequestProperties`
- New struct `ActionRequests`
- New struct `AnswerCitation`
- New struct `AvailableInplaceUpgradeOSListResult`
- New struct `AvailableInplaceUpgradeOSProperties`
- New struct `AvailableInplaceUpgradeOSResource`
- New struct `BillingHubExecutionUsageDetail`
- New struct `BillingHubFreeHourIncrementEntry`
- New struct `BillingHubGetFreeHourBalanceResponse`
- New struct `BillingHubGetUsageRequest`
- New struct `BillingHubGetUsageResponse`
- New struct `BillingHubPackageUsage`
- New struct `BillingHubUsage`
- New struct `BillingHubUsageGroup`
- New struct `BillingHubUsageGroupedByUpdateType`
- New struct `ChatRequest`
- New struct `ChatResponse`
- New struct `ChatResponseProperties`
- New struct `ChatSessionProperties`
- New struct `ChatSessionResource`
- New struct `ChatSessionResourceListResult`
- New struct `CopyFromPackageOperationParameters`
- New struct `CredentialListResult`
- New struct `CredentialResource`
- New struct `CustomImageListResult`
- New struct `CustomImageProperties`
- New struct `CustomImageResource`
- New struct `DraftPackageGetPathResponse`
- New struct `DraftPackageIntuneAppMetadata`
- New struct `DraftPackageIntuneAppMetadataItem`
- New struct `DraftPackageListResult`
- New struct `DraftPackageProperties`
- New struct `DraftPackageResource`
- New struct `DraftPackageUpdateParameterProperties`
- New struct `DraftPackageUpdateParameters`
- New struct `EnrolledIntuneApp`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ExtractFileOperationParameters`
- New struct `FeatureUpdateSupportedOsesProperties`
- New struct `FeatureUpdateSupportedOsesResource`
- New struct `FeatureUpdateSupportedOsesResult`
- New struct `FirstPartyAppDefinition`
- New struct `FirstPartyAppListResult`
- New struct `FirstPartyAppProperties`
- New struct `FirstPartyAppResource`
- New struct `FreeHourBalanceResource`
- New struct `FreeHourBalancesListResult`
- New struct `GalleryAppDefinition`
- New struct `GalleryAppListResult`
- New struct `GalleryAppProperties`
- New struct `GalleryAppResource`
- New struct `GalleryAppSKUListResult`
- New struct `GalleryAppSKUResource`
- New struct `GenerateOperationParameters`
- New struct `GetImageDefinitionParameters`
- New struct `HighlightedFile`
- New struct `ImageDefinitionListResult`
- New struct `ImageDefinitionProperties`
- New struct `ImageDefinitionResource`
- New struct `ImageDefinitionsListResult`
- New struct `ImageNameCheckAvailabilityParameters`
- New struct `ImageValidationResults`
- New struct `InplaceUpgradeOSInfo`
- New struct `InplaceUpgradeProperties`
- New struct `IntuneEnrollmentMetadata`
- New struct `IntuneSingletonResourceProperties`
- New struct `OsProperties`
- New struct `PackageRunTestParameters`
- New struct `PreReleaseAccessRequestSpec`
- New struct `ReleaseProperties`
- New struct `SystemAssignedServiceIdentity`
- New struct `TabState`
- New struct `TestResultConsoleLogDownloadURLParameters`
- New struct `TestResultFirstPartyAppDefinition`
- New struct `VHDListResult`
- New struct `VHDProperties`
- New struct `VHDResource`
- New struct `VerificationResult`
- New struct `WingetAppSKUProperties`
- New field `Identity` in struct `AccountResource`
- New field `Identity` in struct `AccountUpdateParameters`
- New field `EnrollIntuneBefore`, `Install1PAppBefore`, `PostUpgrade`, `PreUpgrade` in struct `Command`
- New field `ResourceType` in struct `GetFileUploadURLParameters`
- New field `CustomImageDisplayName`, `CustomImageID`, `InplaceUpgradeBaselineProperties` in struct `OSUpdateProperties`
- New field `CustomImageDisplayName`, `CustomImageID`, `InplaceUpgradeBaselineProperties` in struct `OSUpdateTestSummary`
- New field `ActionType` in struct `Operation`
- New field `DraftPackageID`, `FirstPartyApps`, `GalleryApps`, `InplaceUpgradeOSPair`, `IntuneEnrollmentMetadata` in struct `PackageProperties`
- New field `DraftPackageID`, `FirstPartyApps`, `InplaceUpgradeOSPair`, `IntuneEnrollmentMetadata` in struct `PackageUpdateParameterProperties`
- New field `SystemData` in struct `ProxyResource`
- New field `InplaceUpgradeBaselineReliabilityResults` in struct `ReliabilityResultSingletonResourceProperties`
- New field `SystemData` in struct `Resource`
- New field `StderrLogFileName`, `StdoutLogFileName` in struct `ScriptExecutionResult`
- New field `BaselineOSs`, `InsiderChannelIDs`, `TargetOSImageIDs` in struct `TargetOSInfo`
- New field `ValidationResultID` in struct `Test`
- New field `CustomImageDisplayName`, `CustomImageID`, `FirstPartyApps`, `InplaceUpgradeProperties`, `InteropMediaType`, `InteropMediaVersion`, `TestEndTime`, `TestStartTime` in struct `TestResultProperties`
- New field `InplaceUpgradesTestSummary`, `PackageTags` in struct `TestSummaryProperties`
- New field `SystemData` in struct `TrackedResource`


## 0.6.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 0.6.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).