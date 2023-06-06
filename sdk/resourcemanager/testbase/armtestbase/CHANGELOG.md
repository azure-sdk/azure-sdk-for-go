# Release History

## 0.7.0 (2023-06-06)
### Breaking Changes

- Type of `ErrorResponse.Error` has been changed from `*ErrorDefinition` to `*ErrorDetail`
- Struct `ErrorDefinition` has been removed
- Field `Etag` of struct `AccountResource` has been removed
- Field `Etag` of struct `PackageResource` has been removed
- Field `Etag` of struct `TrackedResource` has been removed

### Features Added

- New value `ActionFlowDrivenCustom` added to enum type `Action`
- New value `OsUpdateTypeInplaceUpgrade` added to enum type `OsUpdateType`
- New value `TestTypeFlowDrivenTest` added to enum type `TestType`
- New enum type `Architecture` with values `ArchitectureArm`, `ArchitectureArm64`, `ArchitectureIa64`, `ArchitectureX64`, `ArchitectureX86`
- New enum type `DraftPackageSourceType` with values `DraftPackageSourceTypeIntuneWin`, `DraftPackageSourceTypeNative`, `DraftPackageSourceTypeTestBasePackage`
- New enum type `Engagements` with values `EngagementsMAPP`, `EngagementsMVI`, `EngagementsMVP`, `EngagementsOther`, `EngagementsSUVP`
- New enum type `ExtractFileType` with values `ExtractFileTypeIntuneWinPackage`, `ExtractFileTypeTestBasePackage`
- New enum type `InteropExecutionMode` with values `InteropExecutionModeFirstPartyApp`, `InteropExecutionModeFirstPartyAppWithTests`
- New enum type `IntuneExtractStatus` with values `IntuneExtractStatusExtractFailed`, `IntuneExtractStatusNoDependencyApp`, `IntuneExtractStatusReady`, `IntuneExtractStatusUploadFailed`, `IntuneExtractStatusUploading`
- New enum type `OsProductState` with values `OsProductStateActive`, `OsProductStateDisabled`
- New enum type `PackageStudioTabs` with values `PackageStudioTabsBasicsTab`, `PackageStudioTabsConfigureTestTab`, `PackageStudioTabsEditPackageTab`, `PackageStudioTabsReviewAndCreateTab`, `PackageStudioTabsTagsTab`, `PackageStudioTabsTestMatrixTab`, `PackageStudioTabsUnspecified`
- New enum type `RequestStatus` with values `RequestStatusApproved`, `RequestStatusDeclined`, `RequestStatusInReview`
- New enum type `RequestTypes` with values `RequestTypesPreReleaseAccess`
- New function `NewActionRequestsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ActionRequestsClient, error)`
- New function `*ActionRequestsClient.Delete(context.Context, string, string, string, *ActionRequestsClientDeleteOptions) (ActionRequestsClientDeleteResponse, error)`
- New function `*ActionRequestsClient.List(context.Context, string, string, *ActionRequestsClientListOptions) (ActionRequestsClientListResponse, error)`
- New function `*ActionRequestsClient.Put(context.Context, string, string, string, *ActionRequestsClientPutOptions) (ActionRequestsClientPutResponse, error)`
- New function `NewAvailableInplaceUpgradeOSClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AvailableInplaceUpgradeOSClient, error)`
- New function `*AvailableInplaceUpgradeOSClient.Get(context.Context, string, string, string, *AvailableInplaceUpgradeOSClientGetOptions) (AvailableInplaceUpgradeOSClientGetResponse, error)`
- New function `*AvailableInplaceUpgradeOSClient.NewListPager(string, string, OsUpdateType, *AvailableInplaceUpgradeOSClientListOptions) *runtime.Pager[AvailableInplaceUpgradeOSClientListResponse]`
- New function `NewBillingHubServiceClient(string, azcore.TokenCredential, *arm.ClientOptions) (*BillingHubServiceClient, error)`
- New function `*BillingHubServiceClient.GetFreeHourBalance(context.Context, string, string, *BillingHubServiceClientGetFreeHourBalanceOptions) (BillingHubServiceClientGetFreeHourBalanceResponse, error)`
- New function `*BillingHubServiceClient.GetUsage(context.Context, string, string, *BillingHubServiceClientGetUsageOptions) (BillingHubServiceClientGetUsageResponse, error)`
- New function `*ClientFactory.NewActionRequestsClient() *ActionRequestsClient`
- New function `*ClientFactory.NewAvailableInplaceUpgradeOSClient() *AvailableInplaceUpgradeOSClient`
- New function `*ClientFactory.NewBillingHubServiceClient() *BillingHubServiceClient`
- New function `*ClientFactory.NewDraftPackagesClient() *DraftPackagesClient`
- New function `*ClientFactory.NewFeatureUpdateSupportedOsesClient() *FeatureUpdateSupportedOsesClient`
- New function `*ClientFactory.NewFirstPartyAppsClient() *FirstPartyAppsClient`
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
- New function `*PackagesClient.RunTest(context.Context, string, string, string, *PackagesClientRunTestOptions) (PackagesClientRunTestResponse, error)`
- New function `*TestResultsClient.GetConsoleLogDownloadURL(context.Context, string, string, string, string, TestResultConsoleLogDownloadURLParameters, *TestResultsClientGetConsoleLogDownloadURLOptions) (TestResultsClientGetConsoleLogDownloadURLResponse, error)`
- New struct `ActionRequest`
- New struct `ActionRequestProperties`
- New struct `ActionRequests`
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
- New struct `CopyFromPackageOperationParameters`
- New struct `DraftPackageGetPathResponse`
- New struct `DraftPackageIntuneAppMetadata`
- New struct `DraftPackageIntuneAppMetadataItem`
- New struct `DraftPackageListResult`
- New struct `DraftPackageProperties`
- New struct `DraftPackageResource`
- New struct `DraftPackageUpdateParameterProperties`
- New struct `DraftPackageUpdateParameters`
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
- New struct `GenerateOperationParameters`
- New struct `HighlightedFile`
- New struct `InplaceUpgradeOSInfo`
- New struct `InplaceUpgradeProperties`
- New struct `OsProperties`
- New struct `PackageRunTestParameters`
- New struct `PreReleaseAccessRequestSpec`
- New struct `ReleaseProperties`
- New struct `TabState`
- New struct `TestResultConsoleLogDownloadURLParameters`
- New field `Install1PAppBefore`, `PostUpgrade`, `PreUpgrade` in struct `Command`
- New field `InplaceUpgradeBaselineProperties` in struct `OSUpdateProperties`
- New field `InplaceUpgradeBaselineProperties` in struct `OSUpdateTestSummary`
- New field `DraftPackageID`, `FirstPartyApps`, `InplaceUpgradeOSPair` in struct `PackageProperties`
- New field `DraftPackageID`, `FirstPartyApps`, `InplaceUpgradeOSPair` in struct `PackageUpdateParameterProperties`
- New field `SystemData` in struct `ProxyResource`
- New field `InplaceUpgradeBaselineReliabilityResults` in struct `ReliabilityResultSingletonResourceProperties`
- New field `SystemData` in struct `Resource`
- New field `StderrLogFileName`, `StdoutLogFileName` in struct `ScriptExecutionResult`
- New field `BaselineOSs`, `InsiderChannelIDs` in struct `TargetOSInfo`
- New field `ValidationResultID` in struct `Test`
- New field `InplaceUpgradeProperties`, `InteropMediaType`, `InteropMediaVersion` in struct `TestResultProperties`
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