# Release History

## 2.0.0-beta.1 (2025-04-11)
### Breaking Changes

- Type of `Operation.Origin` has been changed from `*string` to `*Origin`
- `MachineBootTypeBIOS` from enum `MachineBootType` has been removed
- `ProvisioningStateCreating`, `ProvisioningStateMoving` from enum `ProvisioningState` has been removed
- Enum `AssessmentStage` has been removed
- Enum `AzureDiskSize` has been removed
- Enum `AzureDiskSuitabilityDetail` has been removed
- Enum `AzureDiskSuitabilityExplanation` has been removed
- Enum `AzureNetworkAdapterSuitabilityDetail` has been removed
- Enum `AzureNetworkAdapterSuitabilityExplanation` has been removed
- Enum `AzureOfferCode` has been removed
- Enum `AzureVMSize` has been removed
- Enum `AzureVMSuitabilityDetail` has been removed
- Enum `AzureVMSuitabilityExplanation` has been removed
- Enum `CloudSuitability` has been removed
- Enum `Currency` has been removed
- Enum `GroupStatus` has been removed
- Enum `GroupUpdateOperation` has been removed
- Enum `PrivateEndpointConnectionPropertiesProvisioningState` has been removed
- Enum `PrivateLinkServiceConnectionStateStatus` has been removed
- Enum `ProjectStatus` has been removed
- Enum `ReservedInstance` has been removed
- Function `NewAssessedMachinesClient` has been removed
- Function `*AssessedMachinesClient.Get` has been removed
- Function `*AssessedMachinesClient.NewListByAssessmentPager` has been removed
- Function `NewAssessmentsClient` has been removed
- Function `*AssessmentsClient.Create` has been removed
- Function `*AssessmentsClient.Delete` has been removed
- Function `*AssessmentsClient.Get` has been removed
- Function `*AssessmentsClient.GetReportDownloadURL` has been removed
- Function `*AssessmentsClient.NewListByGroupPager` has been removed
- Function `*AssessmentsClient.NewListByProjectPager` has been removed
- Function `*ClientFactory.NewAssessedMachinesClient` has been removed
- Function `*ClientFactory.NewAssessmentsClient` has been removed
- Function `*ClientFactory.NewGroupsClient` has been removed
- Function `*ClientFactory.NewHyperVCollectorsClient` has been removed
- Function `*ClientFactory.NewImportCollectorsClient` has been removed
- Function `*ClientFactory.NewMachinesClient` has been removed
- Function `*ClientFactory.NewPrivateEndpointConnectionClient` has been removed
- Function `*ClientFactory.NewPrivateLinkResourceClient` has been removed
- Function `*ClientFactory.NewProjectsClient` has been removed
- Function `*ClientFactory.NewServerCollectorsClient` has been removed
- Function `*ClientFactory.NewVMwareCollectorsClient` has been removed
- Function `NewGroupsClient` has been removed
- Function `*GroupsClient.Create` has been removed
- Function `*GroupsClient.Delete` has been removed
- Function `*GroupsClient.Get` has been removed
- Function `*GroupsClient.NewListByProjectPager` has been removed
- Function `*GroupsClient.UpdateMachines` has been removed
- Function `NewHyperVCollectorsClient` has been removed
- Function `*HyperVCollectorsClient.Create` has been removed
- Function `*HyperVCollectorsClient.Delete` has been removed
- Function `*HyperVCollectorsClient.Get` has been removed
- Function `*HyperVCollectorsClient.NewListByProjectPager` has been removed
- Function `NewImportCollectorsClient` has been removed
- Function `*ImportCollectorsClient.Create` has been removed
- Function `*ImportCollectorsClient.Delete` has been removed
- Function `*ImportCollectorsClient.Get` has been removed
- Function `*ImportCollectorsClient.NewListByProjectPager` has been removed
- Function `NewMachinesClient` has been removed
- Function `*MachinesClient.Get` has been removed
- Function `*MachinesClient.NewListByProjectPager` has been removed
- Function `NewPrivateEndpointConnectionClient` has been removed
- Function `*PrivateEndpointConnectionClient.Delete` has been removed
- Function `*PrivateEndpointConnectionClient.Get` has been removed
- Function `*PrivateEndpointConnectionClient.ListByProject` has been removed
- Function `*PrivateEndpointConnectionClient.Update` has been removed
- Function `NewPrivateLinkResourceClient` has been removed
- Function `*PrivateLinkResourceClient.Get` has been removed
- Function `*PrivateLinkResourceClient.ListByProject` has been removed
- Function `NewProjectsClient` has been removed
- Function `*ProjectsClient.AssessmentOptions` has been removed
- Function `*ProjectsClient.NewAssessmentOptionsListPager` has been removed
- Function `*ProjectsClient.Create` has been removed
- Function `*ProjectsClient.Delete` has been removed
- Function `*ProjectsClient.Get` has been removed
- Function `*ProjectsClient.NewListBySubscriptionPager` has been removed
- Function `*ProjectsClient.NewListPager` has been removed
- Function `*ProjectsClient.Update` has been removed
- Function `NewServerCollectorsClient` has been removed
- Function `*ServerCollectorsClient.Create` has been removed
- Function `*ServerCollectorsClient.Delete` has been removed
- Function `*ServerCollectorsClient.Get` has been removed
- Function `*ServerCollectorsClient.NewListByProjectPager` has been removed
- Function `NewVMwareCollectorsClient` has been removed
- Function `*VMwareCollectorsClient.Create` has been removed
- Function `*VMwareCollectorsClient.Delete` has been removed
- Function `*VMwareCollectorsClient.Get` has been removed
- Function `*VMwareCollectorsClient.NewListByProjectPager` has been removed
- Struct `AssessedDisk` has been removed
- Struct `AssessedMachine` has been removed
- Struct `AssessedMachineProperties` has been removed
- Struct `AssessedMachineResultList` has been removed
- Struct `AssessedNetworkAdapter` has been removed
- Struct `Assessment` has been removed
- Struct `AssessmentOptions` has been removed
- Struct `AssessmentOptionsProperties` has been removed
- Struct `AssessmentOptionsResultList` has been removed
- Struct `AssessmentProperties` has been removed
- Struct `AssessmentResultList` has been removed
- Struct `CollectorAgentProperties` has been removed
- Struct `CollectorBodyAgentSpnProperties` has been removed
- Struct `CollectorProperties` has been removed
- Struct `Disk` has been removed
- Struct `Group` has been removed
- Struct `GroupBodyProperties` has been removed
- Struct `GroupProperties` has been removed
- Struct `GroupResultList` has been removed
- Struct `HyperVCollector` has been removed
- Struct `HyperVCollectorList` has been removed
- Struct `ImportCollector` has been removed
- Struct `ImportCollectorList` has been removed
- Struct `ImportCollectorProperties` has been removed
- Struct `Machine` has been removed
- Struct `MachineProperties` has been removed
- Struct `MachineResultList` has been removed
- Struct `NetworkAdapter` has been removed
- Struct `OperationResultList` has been removed
- Struct `PrivateEndpointConnection` has been removed
- Struct `PrivateEndpointConnectionCollection` has been removed
- Struct `PrivateEndpointConnectionProperties` has been removed
- Struct `PrivateLinkResource` has been removed
- Struct `PrivateLinkResourceCollection` has been removed
- Struct `PrivateLinkResourceProperties` has been removed
- Struct `PrivateLinkServiceConnectionState` has been removed
- Struct `Project` has been removed
- Struct `ProjectProperties` has been removed
- Struct `ProjectResultList` has been removed
- Struct `ResourceID` has been removed
- Struct `ServerCollector` has been removed
- Struct `ServerCollectorList` has been removed
- Struct `UpdateGroupBody` has been removed
- Struct `VMFamily` has been removed
- Struct `VMwareCollector` has been removed
- Struct `VMwareCollectorList` has been removed
- Field `OperationResultList` of struct `OperationsClientListResponse` has been removed

### Features Added

- New value `AssessmentStatusDeleted`, `AssessmentStatusFailed` added to enum type `AssessmentStatus`
- New value `AzureDiskTypePremiumV2`, `AzureDiskTypeUltra` added to enum type `AzureDiskType`
- New value `AzureLocationAustraliaCentral`, `AzureLocationAustraliaCentral2`, `AzureLocationChinaEast2`, `AzureLocationChinaNorth2`, `AzureLocationFranceCentral`, `AzureLocationFranceSouth`, `AzureLocationGermanyNorth`, `AzureLocationGermanyWestCentral`, `AzureLocationIsraelCentral`, `AzureLocationItalyNorth`, `AzureLocationJioIndiaWest`, `AzureLocationMexicoCentral`, `AzureLocationNewZealandNorth`, `AzureLocationNorwayEast`, `AzureLocationNorwayWest`, `AzureLocationPolandCentral`, `AzureLocationQatarCentral`, `AzureLocationSouthAfricaNorth`, `AzureLocationSouthAfricaWest`, `AzureLocationSpainCentral`, `AzureLocationSwedenCentral`, `AzureLocationSwitzerlandNorth`, `AzureLocationSwitzerlandWest`, `AzureLocationUAECentral`, `AzureLocationUAENorth`, `AzureLocationUsNatEast`, `AzureLocationUsNatWest`, `AzureLocationUsSecCentral`, `AzureLocationUsSecEast`, `AzureLocationUsSecWest` added to enum type `AzureLocation`
- New value `AzureVMFamilyDadsv5Series`, `AzureVMFamilyDasv4Series`, `AzureVMFamilyDasv5Series`, `AzureVMFamilyDav4Series`, `AzureVMFamilyDdsv4Series`, `AzureVMFamilyDdsv5Series`, `AzureVMFamilyDdv4Series`, `AzureVMFamilyDdv5Series`, `AzureVMFamilyDsv4Series`, `AzureVMFamilyDsv5Series`, `AzureVMFamilyDv4Series`, `AzureVMFamilyDv5Series`, `AzureVMFamilyEadsv5Series`, `AzureVMFamilyEasv4Series`, `AzureVMFamilyEasv5Series`, `AzureVMFamilyEav4Series`, `AzureVMFamilyEbdsv5Series`, `AzureVMFamilyEbsv5Series`, `AzureVMFamilyEdsv4Series`, `AzureVMFamilyEdsv5Series`, `AzureVMFamilyEdv4Series`, `AzureVMFamilyEdv5Series`, `AzureVMFamilyEsv4Series`, `AzureVMFamilyEsv5Series`, `AzureVMFamilyEv4Series`, `AzureVMFamilyEv5Series`, `AzureVMFamilyLsv2Series`, `AzureVMFamilyMdsv2Series`, `AzureVMFamilyMsv2Series`, `AzureVMFamilyMv2Series` added to enum type `AzureVMFamily`
- New value `MachineBootTypeBios`, `MachineBootTypeNotSpecified` added to enum type `MachineBootType`
- New value `PercentilePercentileUnknown` added to enum type `Percentile`
- New value `ProvisioningStateCanceled`, `ProvisioningStateProvisioning`, `ProvisioningStateUpdating` added to enum type `ProvisioningState`
- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `AssessedMachineType` with values `AssessedMachineTypeAssessedMachine`, `AssessedMachineTypeAvsAssessedMachine`, `AssessedMachineTypeSQLAssessedMachine`, `AssessedMachineTypeUnknown`
- New enum type `AssessmentSource` with values `AssessmentSourceIIS`, `AssessmentSourceMachine`, `AssessmentSourceMySQLServer`, `AssessmentSourceOracleDatabase`, `AssessmentSourceOracleServer`, `AssessmentSourceSAPInstance`, `AssessmentSourceSQLDatabase`, `AssessmentSourceSQLInstance`, `AssessmentSourceSpringbootApplication`, `AssessmentSourceTomCat`, `AssessmentSourceUnknown`, `AssessmentSourceWebApps`
- New enum type `AzureCurrency` with values `AzureCurrencyARS`, `AzureCurrencyAUD`, `AzureCurrencyBRL`, `AzureCurrencyCAD`, `AzureCurrencyCHF`, `AzureCurrencyCNY`, `AzureCurrencyDKK`, `AzureCurrencyEUR`, `AzureCurrencyGBP`, `AzureCurrencyHKD`, `AzureCurrencyIDR`, `AzureCurrencyINR`, `AzureCurrencyJPY`, `AzureCurrencyKRW`, `AzureCurrencyMXN`, `AzureCurrencyMYR`, `AzureCurrencyNOK`, `AzureCurrencyNZD`, `AzureCurrencyRUB`, `AzureCurrencySAR`, `AzureCurrencySEK`, `AzureCurrencyTRY`, `AzureCurrencyTWD`, `AzureCurrencyUSD`, `AzureCurrencyUnknown`, `AzureCurrencyZAR`
- New enum type `AzureManagementOfferingType` with values `AzureManagementOfferingTypeAUM`, `AzureManagementOfferingTypeAzMon`, `AzureManagementOfferingTypeAzureBackup`, `AzureManagementOfferingTypeNo`, `AzureManagementOfferingTypeSCOMMI`
- New enum type `AzureOffer` with values `AzureOfferEA`, `AzureOfferMsazr0003P`, `AzureOfferMsazr0023P`, `AzureOfferMsazrusgov0003P`, `AzureOfferMsmcazr0044P`, `AzureOfferMsmcazr0059P`, `AzureOfferMsmcazr0060P`, `AzureOfferMsmcazr0063P`, `AzureOfferSavingsPlan1Year`, `AzureOfferSavingsPlan3Year`, `AzureOfferUnknown`
- New enum type `AzureSecurityOfferingType` with values `AzureSecurityOfferingTypeMDC`, `AzureSecurityOfferingTypeNO`
- New enum type `AzureTarget` with values `AzureTargetAKS`, `AzureTargetAvs`, `AzureTargetAzureAppService`, `AzureTargetAzureAppServiceContainer`, `AzureTargetAzureSQLVM`, `AzureTargetAzureSpringApps`, `AzureTargetAzureVM`, `AzureTargetFlexServerPG`, `AzureTargetMySQLAzureFlexServer`, `AzureTargetOracleIaasVM`, `AzureTargetSAPAzureInstance`, `AzureTargetSQLDatabase`, `AzureTargetSQLMI`, `AzureTargetUnknown`
- New enum type `AzureVMSecuritySuitabilityDetail` with values `AzureVMSecuritySuitabilityDetailAllDiskSizeExceeds2TB`, `AzureVMSecuritySuitabilityDetailAllDiskSizeExceeds4TB`, `AzureVMSecuritySuitabilityDetailAnyDiskSizeExceeds4TB`, `AzureVMSecuritySuitabilityDetailAnyDiskSizeExceeds4TBConditional`, `AzureVMSecuritySuitabilityDetailBootTypeNotSupported`, `AzureVMSecuritySuitabilityDetailDiskSize2TBConditionalReadyBiosBoot`, `AzureVMSecuritySuitabilityDetailDiskSizeLarge`, `AzureVMSecuritySuitabilityDetailGuestOperatingSystemNotSupported`, `AzureVMSecuritySuitabilityDetailNotApplicable`, `AzureVMSecuritySuitabilityDetailOSNameCannotBeRead`, `AzureVMSecuritySuitabilityDetailOSNotSupported`, `AzureVMSecuritySuitabilityDetailStandardCanMigrate`, `AzureVMSecuritySuitabilityDetailStandardCannotMigrateOSNotSupported`, `AzureVMSecuritySuitabilityDetailStandardNotReady2TBBiosBoot`, `AzureVMSecuritySuitabilityDetailStandardNotReadyOSNotSupported`, `AzureVMSecuritySuitabilityDetailStandardReady`, `AzureVMSecuritySuitabilityDetailTVMCanMigrateConditional`, `AzureVMSecuritySuitabilityDetailTVMCannotMigrateOSNotSupported`, `AzureVMSecuritySuitabilityDetailTVMNotReadyOSNotSupported`, `AzureVMSecuritySuitabilityDetailTVMNotSupportedForBiosBoot`, `AzureVMSecuritySuitabilityDetailTVMReady`, `AzureVMSecuritySuitabilityDetailUnknown`
- New enum type `AzureVMSecurityType` with values `AzureVMSecurityTypeCVM`, `AzureVMSecurityTypeStandard`, `AzureVMSecurityTypeTVM`, `AzureVMSecurityTypeUnknown`
- New enum type `CloudSuitabilityCommon` with values `CloudSuitabilityCommonConditionallySuitable`, `CloudSuitabilityCommonNotSuitable`, `CloudSuitabilityCommonReadinessUnknown`, `CloudSuitabilityCommonSuitable`, `CloudSuitabilityCommonSuitableWithWarnings`, `CloudSuitabilityCommonUnknown`
- New enum type `CostType` with values `CostTypeDataProtectionService`, `CostTypeMonitoringService`, `CostTypeMonthlyAvsExternalStorageCost`, `CostTypeMonthlyAvsNetworkCost`, `CostTypeMonthlyAvsNodeCost`, `CostTypeMonthlyAzureHybridCost`, `CostTypeMonthlyBandwidthCost`, `CostTypeMonthlyComputeCost`, `CostTypeMonthlyLicensingCost`, `CostTypeMonthlyLinuxAzureHybridCost`, `CostTypeMonthlyManagementCost`, `CostTypeMonthlyPremiumStorageCost`, `CostTypeMonthlyPremiumV2StorageCost`, `CostTypeMonthlySecurityCost`, `CostTypeMonthlyStandardHddStorageCost`, `CostTypeMonthlyStandardSsdStorageCost`, `CostTypeMonthlyStorageCost`, `CostTypeMonthlyUltraDiskCost`, `CostTypeMonthlyUltraStorageCost`, `CostTypePatchingService`, `CostTypeTotalMonthlyCost`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `EnvironmentType` with values `EnvironmentTypeDevTest`, `EnvironmentTypeProduction`, `EnvironmentTypeUnknown`
- New enum type `GuestOperatingSystemArchitecture` with values `GuestOperatingSystemArchitectureUnknown`, `GuestOperatingSystemArchitectureX64`, `GuestOperatingSystemArchitectureX86`
- New enum type `LicensingProgram` with values `LicensingProgramEA`, `LicensingProgramMCA`, `LicensingProgramRetail`, `LicensingProgramUnknown`
- New enum type `LinkageKind` with values `LinkageKindDatabase`, `LinkageKindInstance`, `LinkageKindMachine`, `LinkageKindServer`, `LinkageKindUnknown`, `LinkageKindWebApplication`, `LinkageKindWebServer`
- New enum type `LinkageType` with values `LinkageTypeParent`, `LinkageTypeSource`
- New enum type `MigrationIssuesCategory` with values `MigrationIssuesCategoryInternal`, `MigrationIssuesCategoryIssue`, `MigrationIssuesCategoryWarning`
- New enum type `MigrationPlatform` with values `MigrationPlatformIaaS`, `MigrationPlatformPaaS`, `MigrationPlatformSaaS`, `MigrationPlatformUnknown`
- New enum type `MigrationType` with values `MigrationTypeRearchitect`, `MigrationTypeRehost`, `MigrationTypeReplatform`, `MigrationTypeRetain`, `MigrationTypeUnknown`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New enum type `PremiumDiskSupport` with values `PremiumDiskSupportPremiumDiskNotSupported`, `PremiumDiskSupportPremiumDiskSupported`
- New enum type `SKUKind` with values `SKUKindDetails`, `SKUKindEndpoint`
- New enum type `SKUType` with values `SKUTypeAppServicePlan`, `SKUTypeAzureSpringApps`, `SKUTypeCompute`, `SKUTypeNetwork`, `SKUTypeStorage`, `SKUTypeUnknown`
- New enum type `SavingsOptions` with values `SavingsOptionsCustomAzureOfferCode`, `SavingsOptionsNone`, `SavingsOptionsRI1Year`, `SavingsOptionsRI3Year`, `SavingsOptionsSavingsPlan1Year`, `SavingsOptionsSavingsPlan3Year`
- New enum type `SavingsType` with values `SavingsTypeMonthlyAhubSQLCostSavings`, `SavingsTypeMonthlyAzureHybridCostSavings`, `SavingsTypeMonthlyLinuxAzureHybridCostSavings`, `SavingsTypeMonthlyVcfByolCostDifference`
- New enum type `ScopeType` with values `ScopeTypeAzureResourceGraphQuery`, `ScopeTypeServerGroupID`
- New enum type `StrategyType` with values `StrategyTypeCostOptimized`, `StrategyTypeEffortOptimized`, `StrategyTypePaaSPreferred`
- New enum type `SummaryType` with values `SummaryTypeStrategy`, `SummaryTypeTarget`, `SummaryTypeUnknown`
- New enum type `UltraDiskSupport` with values `UltraDiskSupportUltraDiskNotSupported`, `UltraDiskSupportUltraDiskSupported`
- New enum type `VMFamilyCategoryItem` with values `VMFamilyCategoryItemComputeOptimized`, `VMFamilyCategoryItemConfidential`, `VMFamilyCategoryItemGeneralPurpose`, `VMFamilyCategoryItemGpuOptimized`, `VMFamilyCategoryItemHighPerformanceCompute`, `VMFamilyCategoryItemMemoryOptimized`, `VMFamilyCategoryItemStorageOptimized`, `VMFamilyCategoryItemSupportsPremiumStorage`, `VMFamilyCategoryItemSupportsUltraDiskStorage`, `VMFamilyCategoryItemUnknown`
- New enum type `WorkloadType` with values `WorkloadTypeDatabase`, `WorkloadTypeHost`, `WorkloadTypeInstance`, `WorkloadTypeMachine`, `WorkloadTypeManagementServer`, `WorkloadTypeServer`, `WorkloadTypeUnknown`, `WorkloadTypeWebApplication`, `WorkloadTypeWebServer`
- New function `NewAssessedMachinesV2OperationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AssessedMachinesV2OperationsClient, error)`
- New function `*AssessedMachinesV2OperationsClient.Get(context.Context, string, string, string, string, *AssessedMachinesV2OperationsClientGetOptions) (AssessedMachinesV2OperationsClientGetResponse, error)`
- New function `*AssessedMachinesV2OperationsClient.NewListByParentPager(string, string, string, *AssessedMachinesV2OperationsClientListByParentOptions) *runtime.Pager[AssessedMachinesV2OperationsClientListByParentResponse]`
- New function `*ClientFactory.NewAssessedMachinesV2OperationsClient() *AssessedMachinesV2OperationsClient`
- New function `*ClientFactory.NewMachineAssessmentV2SummaryOperationsClient() *MachineAssessmentV2SummaryOperationsClient`
- New function `*ClientFactory.NewMachineAssessmentsV2OperationsClient() *MachineAssessmentsV2OperationsClient`
- New function `*ClientFactory.NewMachineGraphAssessmentOptionsOperationsClient() *MachineGraphAssessmentOptionsOperationsClient`
- New function `NewMachineAssessmentV2SummaryOperationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*MachineAssessmentV2SummaryOperationsClient, error)`
- New function `*MachineAssessmentV2SummaryOperationsClient.Get(context.Context, string, string, string, string, *MachineAssessmentV2SummaryOperationsClientGetOptions) (MachineAssessmentV2SummaryOperationsClientGetResponse, error)`
- New function `*MachineAssessmentV2SummaryOperationsClient.NewListByParentPager(string, string, string, *MachineAssessmentV2SummaryOperationsClientListByParentOptions) *runtime.Pager[MachineAssessmentV2SummaryOperationsClientListByParentResponse]`
- New function `NewMachineAssessmentsV2OperationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*MachineAssessmentsV2OperationsClient, error)`
- New function `*MachineAssessmentsV2OperationsClient.BeginCreate(context.Context, string, string, string, MachineAssessmentV2, *MachineAssessmentsV2OperationsClientBeginCreateOptions) (*runtime.Poller[MachineAssessmentsV2OperationsClientCreateResponse], error)`
- New function `*MachineAssessmentsV2OperationsClient.Delete(context.Context, string, string, string, *MachineAssessmentsV2OperationsClientDeleteOptions) (MachineAssessmentsV2OperationsClientDeleteResponse, error)`
- New function `*MachineAssessmentsV2OperationsClient.BeginDownloadURL(context.Context, string, string, string, DownloadURLRequest, *MachineAssessmentsV2OperationsClientBeginDownloadURLOptions) (*runtime.Poller[MachineAssessmentsV2OperationsClientDownloadURLResponse], error)`
- New function `*MachineAssessmentsV2OperationsClient.Get(context.Context, string, string, string, *MachineAssessmentsV2OperationsClientGetOptions) (MachineAssessmentsV2OperationsClientGetResponse, error)`
- New function `*MachineAssessmentsV2OperationsClient.NewListByParentPager(string, string, *MachineAssessmentsV2OperationsClientListByParentOptions) *runtime.Pager[MachineAssessmentsV2OperationsClientListByParentResponse]`
- New function `NewMachineGraphAssessmentOptionsOperationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*MachineGraphAssessmentOptionsOperationsClient, error)`
- New function `*MachineGraphAssessmentOptionsOperationsClient.Get(context.Context, string, string, *MachineGraphAssessmentOptionsOperationsClientGetOptions) (MachineGraphAssessmentOptionsOperationsClientGetResponse, error)`
- New function `*MachineGraphAssessmentOptionsOperationsClient.NewListByParentPager(string, string, *MachineGraphAssessmentOptionsOperationsClientListByParentOptions) *runtime.Pager[MachineGraphAssessmentOptionsOperationsClientListByParentResponse]`
- New struct `AssessedDiskDataV3`
- New struct `AssessedMachineExtendedDetails`
- New struct `AssessedMachineNetworkAdaptersV3`
- New struct `AssessedMachineV2`
- New struct `AssessedMachineV2ListResult`
- New struct `AssessedMachineV2Properties`
- New struct `AssessmentDetails`
- New struct `BillingSettings`
- New struct `Configuration`
- New struct `CostDetailsCommon`
- New struct `DownloadURLRequest`
- New struct `Error`
- New struct `Linkages`
- New struct `MachineAssessmentOptionsOutboundEdgeGroup`
- New struct `MachineAssessmentRecommendation`
- New struct `MachineAssessmentSettings`
- New struct `MachineAssessmentV2`
- New struct `MachineAssessmentV2ListResult`
- New struct `MachineAssessmentV2Properties`
- New struct `MachineAssessmentV2Summary`
- New struct `MachineAssessmentV2SummaryListResult`
- New struct `MachineGraphAssessmentOptions`
- New struct `MachineGraphAssessmentOptionsListResult`
- New struct `MachineGraphAssessmentOptionsProperties`
- New struct `ManagementDetails`
- New struct `MigrationDetails`
- New struct `MigrationIssues`
- New struct `MigrationSuitability`
- New struct `MoreInformation`
- New struct `NameValuePair`
- New struct `NameValuePairCloudSuitabilityCommon`
- New struct `NameValuePairCostType`
- New struct `NameValuePairSavingsType`
- New struct `OperationListResult`
- New struct `PerformanceData`
- New struct `ProcessorInfo`
- New struct `ProductSupportStatus`
- New struct `RecommendedFor`
- New struct `SKUDetails`
- New struct `SKUsMigrationSuitability`
- New struct `SavingsDetailsCommon`
- New struct `SavingsSettings`
- New struct `Scope`
- New struct `SecuritySuitability`
- New struct `SourceDetails`
- New struct `SourceRecommendationMigrationSuitability`
- New struct `SummaryProperties`
- New struct `SystemData`
- New struct `TargetDetails`
- New struct `TargetSourcePair`
- New struct `VMSecuritySuitability`
- New field `ActionType`, `IsDataAction` in struct `Operation`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-06-10)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).