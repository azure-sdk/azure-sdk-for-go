# Release History

## 2.0.0 (2024-04-17)
### Breaking Changes

- Enum `BillingType` has been removed
- Enum `Capacity` has been removed
- Enum `CapacityReservationLevel` has been removed
- Enum `ClusterEntityStatus` has been removed
- Enum `ClusterSKUNameEnum` has been removed
- Enum `DataIngestionStatus` has been removed
- Enum `DataSourceKind` has been removed
- Enum `DataSourceType` has been removed
- Enum `IdentityType` has been removed
- Enum `LinkedServiceEntityStatus` has been removed
- Enum `PublicNetworkAccessType` has been removed
- Enum `PurgeState` has been removed
- Enum `SKUNameEnum` has been removed
- Enum `SearchSortEnum` has been removed
- Enum `StorageInsightState` has been removed
- Enum `Type` has been removed
- Enum `WorkspaceEntityStatus` has been removed
- Enum `WorkspaceSKUNameEnum` has been removed
- Function `NewAvailableServiceTiersClient` has been removed
- Function `*AvailableServiceTiersClient.ListByWorkspace` has been removed
- Function `*ClientFactory.NewAvailableServiceTiersClient` has been removed
- Function `*ClientFactory.NewClustersClient` has been removed
- Function `*ClientFactory.NewDataExportsClient` has been removed
- Function `*ClientFactory.NewDataSourcesClient` has been removed
- Function `*ClientFactory.NewDeletedWorkspacesClient` has been removed
- Function `*ClientFactory.NewGatewaysClient` has been removed
- Function `*ClientFactory.NewIntelligencePacksClient` has been removed
- Function `*ClientFactory.NewLinkedServicesClient` has been removed
- Function `*ClientFactory.NewLinkedStorageAccountsClient` has been removed
- Function `*ClientFactory.NewManagementGroupsClient` has been removed
- Function `*ClientFactory.NewOperationStatusesClient` has been removed
- Function `*ClientFactory.NewOperationsClient` has been removed
- Function `*ClientFactory.NewSavedSearchesClient` has been removed
- Function `*ClientFactory.NewSchemaClient` has been removed
- Function `*ClientFactory.NewSharedKeysClient` has been removed
- Function `*ClientFactory.NewStorageInsightConfigsClient` has been removed
- Function `*ClientFactory.NewTablesClient` has been removed
- Function `*ClientFactory.NewUsagesClient` has been removed
- Function `*ClientFactory.NewWorkspacePurgeClient` has been removed
- Function `NewClustersClient` has been removed
- Function `*ClustersClient.BeginCreateOrUpdate` has been removed
- Function `*ClustersClient.BeginDelete` has been removed
- Function `*ClustersClient.Get` has been removed
- Function `*ClustersClient.NewListByResourceGroupPager` has been removed
- Function `*ClustersClient.NewListPager` has been removed
- Function `*ClustersClient.BeginUpdate` has been removed
- Function `NewDataExportsClient` has been removed
- Function `*DataExportsClient.CreateOrUpdate` has been removed
- Function `*DataExportsClient.Delete` has been removed
- Function `*DataExportsClient.Get` has been removed
- Function `*DataExportsClient.NewListByWorkspacePager` has been removed
- Function `NewDataSourcesClient` has been removed
- Function `*DataSourcesClient.CreateOrUpdate` has been removed
- Function `*DataSourcesClient.Delete` has been removed
- Function `*DataSourcesClient.Get` has been removed
- Function `*DataSourcesClient.NewListByWorkspacePager` has been removed
- Function `NewDeletedWorkspacesClient` has been removed
- Function `*DeletedWorkspacesClient.NewListByResourceGroupPager` has been removed
- Function `*DeletedWorkspacesClient.NewListPager` has been removed
- Function `NewGatewaysClient` has been removed
- Function `*GatewaysClient.Delete` has been removed
- Function `NewIntelligencePacksClient` has been removed
- Function `*IntelligencePacksClient.Disable` has been removed
- Function `*IntelligencePacksClient.Enable` has been removed
- Function `*IntelligencePacksClient.List` has been removed
- Function `NewLinkedServicesClient` has been removed
- Function `*LinkedServicesClient.BeginCreateOrUpdate` has been removed
- Function `*LinkedServicesClient.BeginDelete` has been removed
- Function `*LinkedServicesClient.Get` has been removed
- Function `*LinkedServicesClient.NewListByWorkspacePager` has been removed
- Function `NewLinkedStorageAccountsClient` has been removed
- Function `*LinkedStorageAccountsClient.CreateOrUpdate` has been removed
- Function `*LinkedStorageAccountsClient.Delete` has been removed
- Function `*LinkedStorageAccountsClient.Get` has been removed
- Function `*LinkedStorageAccountsClient.NewListByWorkspacePager` has been removed
- Function `NewManagementGroupsClient` has been removed
- Function `*ManagementGroupsClient.NewListPager` has been removed
- Function `NewOperationStatusesClient` has been removed
- Function `*OperationStatusesClient.Get` has been removed
- Function `NewOperationsClient` has been removed
- Function `*OperationsClient.NewListPager` has been removed
- Function `NewSavedSearchesClient` has been removed
- Function `*SavedSearchesClient.CreateOrUpdate` has been removed
- Function `*SavedSearchesClient.Delete` has been removed
- Function `*SavedSearchesClient.Get` has been removed
- Function `*SavedSearchesClient.ListByWorkspace` has been removed
- Function `NewSchemaClient` has been removed
- Function `*SchemaClient.Get` has been removed
- Function `NewSharedKeysClient` has been removed
- Function `*SharedKeysClient.GetSharedKeys` has been removed
- Function `*SharedKeysClient.Regenerate` has been removed
- Function `NewStorageInsightConfigsClient` has been removed
- Function `*StorageInsightConfigsClient.CreateOrUpdate` has been removed
- Function `*StorageInsightConfigsClient.Delete` has been removed
- Function `*StorageInsightConfigsClient.Get` has been removed
- Function `*StorageInsightConfigsClient.NewListByWorkspacePager` has been removed
- Function `NewTablesClient` has been removed
- Function `*TablesClient.Get` has been removed
- Function `*TablesClient.NewListByWorkspacePager` has been removed
- Function `*TablesClient.Update` has been removed
- Function `NewUsagesClient` has been removed
- Function `*UsagesClient.NewListPager` has been removed
- Function `NewWorkspacePurgeClient` has been removed
- Function `*WorkspacePurgeClient.GetPurgeStatus` has been removed
- Function `*WorkspacePurgeClient.Purge` has been removed
- Function `*WorkspacesClient.BeginCreateOrUpdate` has been removed
- Function `*WorkspacesClient.BeginDelete` has been removed
- Function `*WorkspacesClient.Get` has been removed
- Function `*WorkspacesClient.NewListByResourceGroupPager` has been removed
- Function `*WorkspacesClient.NewListPager` has been removed
- Function `*WorkspacesClient.Update` has been removed
- Struct `AssociatedWorkspace` has been removed
- Struct `AvailableServiceTier` has been removed
- Struct `AzureEntityResource` has been removed
- Struct `CapacityReservationProperties` has been removed
- Struct `Cluster` has been removed
- Struct `ClusterListResult` has been removed
- Struct `ClusterPatch` has been removed
- Struct `ClusterPatchProperties` has been removed
- Struct `ClusterProperties` has been removed
- Struct `ClusterSKU` has been removed
- Struct `CoreSummary` has been removed
- Struct `DataExport` has been removed
- Struct `DataExportListResult` has been removed
- Struct `DataExportProperties` has been removed
- Struct `DataSource` has been removed
- Struct `DataSourceFilter` has been removed
- Struct `DataSourceListResult` has been removed
- Struct `Destination` has been removed
- Struct `DestinationMetaData` has been removed
- Struct `Identity` has been removed
- Struct `IntelligencePack` has been removed
- Struct `KeyVaultProperties` has been removed
- Struct `LinkedService` has been removed
- Struct `LinkedServiceListResult` has been removed
- Struct `LinkedServiceProperties` has been removed
- Struct `LinkedStorageAccountsListResult` has been removed
- Struct `LinkedStorageAccountsProperties` has been removed
- Struct `LinkedStorageAccountsResource` has been removed
- Struct `ManagementGroup` has been removed
- Struct `ManagementGroupProperties` has been removed
- Struct `MetricName` has been removed
- Struct `Operation` has been removed
- Struct `OperationDisplay` has been removed
- Struct `OperationListResult` has been removed
- Struct `OperationStatus` has been removed
- Struct `PrivateLinkScopedResource` has been removed
- Struct `SavedSearch` has been removed
- Struct `SavedSearchProperties` has been removed
- Struct `SavedSearchesListResult` has been removed
- Struct `SearchGetSchemaResponse` has been removed
- Struct `SearchMetadata` has been removed
- Struct `SearchMetadataSchema` has been removed
- Struct `SearchSchemaValue` has been removed
- Struct `SearchSort` has been removed
- Struct `SharedKeys` has been removed
- Struct `StorageAccount` has been removed
- Struct `StorageInsight` has been removed
- Struct `StorageInsightListResult` has been removed
- Struct `StorageInsightProperties` has been removed
- Struct `StorageInsightStatus` has been removed
- Struct `Table` has been removed
- Struct `TableProperties` has been removed
- Struct `TablesListResult` has been removed
- Struct `Tag` has been removed
- Struct `TrackedResource` has been removed
- Struct `UsageMetric` has been removed
- Struct `UserIdentityProperties` has been removed
- Struct `Workspace` has been removed
- Struct `WorkspaceCapping` has been removed
- Struct `WorkspaceFeatures` has been removed
- Struct `WorkspaceListManagementGroupsResult` has been removed
- Struct `WorkspaceListResult` has been removed
- Struct `WorkspaceListUsagesResult` has been removed
- Struct `WorkspacePatch` has been removed
- Struct `WorkspaceProperties` has been removed
- Struct `WorkspacePurgeBody` has been removed
- Struct `WorkspacePurgeBodyFilters` has been removed
- Struct `WorkspacePurgeResponse` has been removed
- Struct `WorkspacePurgeStatusResponse` has been removed
- Struct `WorkspaceSKU` has been removed

### Features Added

- New function `*WorkspacesClient.GetNSP(context.Context, string, string, string, *WorkspacesClientGetNSPOptions) (WorkspacesClientGetNSPResponse, error)`
- New function `*WorkspacesClient.NewListNSPPager(string, string, *WorkspacesClientListNSPOptions) *runtime.Pager[WorkspacesClientListNSPResponse]`
- New function `*WorkspacesClient.BeginReconcileNSP(context.Context, string, string, string, *WorkspacesClientBeginReconcileNSPOptions) (*runtime.Poller[WorkspacesClientReconcileNSPResponse], error)`
- New struct `NSPConfigAccessRule`
- New struct `NSPConfigAccessRuleProperties`
- New struct `NSPConfigAssociation`
- New struct `NSPConfigNetworkSecurityPerimeterRule`
- New struct `NSPConfigPerimeter`
- New struct `NSPConfigProfile`
- New struct `NSPProvisioningIssue`
- New struct `NSPProvisioningIssueProperties`
- New struct `NetworkSecurityPerimeterConfiguration`
- New struct `NetworkSecurityPerimeterConfigurationListResult`
- New struct `NetworkSecurityPerimeterConfigurationProperties`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-04-06)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module

## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).