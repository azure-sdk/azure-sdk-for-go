# Release History

## 0.4.0 (2025-03-04)
### Breaking Changes

- Type of `ManagedClusterVersionDetails.SupportExpiryUTC` has been changed from `*string` to `*time.Time`
- Type of `SystemData.CreatedByType` has been changed from `*string` to `*CreatedByType`
- Type of `SystemData.LastModifiedByType` has been changed from `*string` to `*CreatedByType`
- Type of `VMSSExtensionProperties.ProtectedSettings` has been changed from `any` to `*VMSSExtensionPropertiesProtectedSettings`
- Type of `VMSSExtensionProperties.Settings` has been changed from `any` to `*VMSSExtensionPropertiesSettings`
- Function `*ClientFactory.NewManagedApplyMaintenanceWindowClient` has been removed
- Function `*ClientFactory.NewManagedAzResiliencyStatusClient` has been removed
- Function `*ClientFactory.NewManagedMaintenanceWindowStatusClient` has been removed
- Function `*ClientFactory.NewManagedUnsupportedVMSizesClient` has been removed
- Function `*ClientFactory.NewNodeTypeSKUsClient` has been removed
- Function `NewManagedApplyMaintenanceWindowClient` has been removed
- Function `*ManagedApplyMaintenanceWindowClient.Post` has been removed
- Function `NewManagedAzResiliencyStatusClient` has been removed
- Function `*ManagedAzResiliencyStatusClient.Get` has been removed
- Function `NewManagedMaintenanceWindowStatusClient` has been removed
- Function `*ManagedMaintenanceWindowStatusClient.Get` has been removed
- Function `NewManagedUnsupportedVMSizesClient` has been removed
- Function `*ManagedUnsupportedVMSizesClient.Get` has been removed
- Function `*ManagedUnsupportedVMSizesClient.NewListPager` has been removed
- Function `NewNodeTypeSKUsClient` has been removed
- Function `*NodeTypeSKUsClient.NewListPager` has been removed
- Operation `*OperationResultsClient.Get` has been changed to LRO, use `*OperationResultsClient.BeginGet` instead.
- Field `ApplicationTypeVersionResource` of struct `ApplicationTypeVersionsClientCreateOrUpdateResponse` has been removed
- Field `ApplicationResource` of struct `ApplicationsClientCreateOrUpdateResponse` has been removed
- Field `ManagedClusterCodeVersionResultArray` of struct `ManagedClusterVersionClientListByEnvironmentResponse` has been removed
- Field `ManagedClusterCodeVersionResultArray` of struct `ManagedClusterVersionClientListResponse` has been removed
- Field `ManagedCluster` of struct `ManagedClustersClientCreateOrUpdateResponse` has been removed
- Field `NodeType` of struct `NodeTypesClientCreateOrUpdateResponse` has been removed
- Field `NodeType` of struct `NodeTypesClientUpdateResponse` has been removed
- Field `ServiceResource` of struct `ServicesClientCreateOrUpdateResponse` has been removed

### Features Added

- New value `DiskTypePremiumV2LRS`, `DiskTypePremiumZRS`, `DiskTypeStandardSSDZRS` added to enum type `DiskType`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `FaultKind` with values `FaultKindZone`
- New enum type `FaultSimulationStatus` with values `FaultSimulationStatusActive`, `FaultSimulationStatusDone`, `FaultSimulationStatusStartFailed`, `FaultSimulationStatusStarting`, `FaultSimulationStatusStopFailed`, `FaultSimulationStatusStopping`
- New enum type `SfmcOperationStatus` with values `SfmcOperationStatusAborted`, `SfmcOperationStatusCanceled`, `SfmcOperationStatusCreated`, `SfmcOperationStatusFailed`, `SfmcOperationStatusStarted`, `SfmcOperationStatusSucceeded`
- New function `*ClientFactory.NewManagedVMSizesClient() *ManagedVMSizesClient`
- New function `*FaultSimulationContent.GetFaultSimulationContent() *FaultSimulationContent`
- New function `*ManagedClustersClient.GetFaultSimulation(context.Context, string, string, FaultSimulationIDContent, *ManagedClustersClientGetFaultSimulationOptions) (ManagedClustersClientGetFaultSimulationResponse, error)`
- New function `*ManagedClustersClient.NewListFaultSimulationPager(string, string, *ManagedClustersClientListFaultSimulationOptions) *runtime.Pager[ManagedClustersClientListFaultSimulationResponse]`
- New function `*ManagedClustersClient.ManagedAzResiliencyStatusGet(context.Context, string, string, *ManagedClustersClientManagedAzResiliencyStatusGetOptions) (ManagedClustersClientManagedAzResiliencyStatusGetResponse, error)`
- New function `*ManagedClustersClient.ManagedMaintenanceWindowStatusGet(context.Context, string, string, *ManagedClustersClientManagedMaintenanceWindowStatusGetOptions) (ManagedClustersClientManagedMaintenanceWindowStatusGetResponse, error)`
- New function `*ManagedClustersClient.Post(context.Context, string, string, *ManagedClustersClientPostOptions) (ManagedClustersClientPostResponse, error)`
- New function `*ManagedClustersClient.BeginStartFaultSimulation(context.Context, string, string, FaultSimulationContentClassification, *ManagedClustersClientBeginStartFaultSimulationOptions) (*runtime.Poller[ManagedClustersClientStartFaultSimulationResponse], error)`
- New function `*ManagedClustersClient.BeginStopFaultSimulation(context.Context, string, string, FaultSimulationIDContent, *ManagedClustersClientBeginStopFaultSimulationOptions) (*runtime.Poller[ManagedClustersClientStopFaultSimulationResponse], error)`
- New function `NewManagedVMSizesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagedVMSizesClient, error)`
- New function `*ManagedVMSizesClient.Get(context.Context, string, string, *ManagedVMSizesClientGetOptions) (ManagedVMSizesClientGetResponse, error)`
- New function `*ManagedVMSizesClient.NewListPager(string, *ManagedVMSizesClientListOptions) *runtime.Pager[ManagedVMSizesClientListResponse]`
- New function `*ZoneFaultSimulationContent.GetFaultSimulationContent() *FaultSimulationContent`
- New function `*NodeTypesClient.BeginDeallocate(context.Context, string, string, string, NodeTypeActionParameters, *NodeTypesClientBeginDeallocateOptions) (*runtime.Poller[NodeTypesClientDeallocateResponse], error)`
- New function `*NodeTypesClient.GetFaultSimulation(context.Context, string, string, string, FaultSimulationIDContent, *NodeTypesClientGetFaultSimulationOptions) (NodeTypesClientGetFaultSimulationResponse, error)`
- New function `*NodeTypesClient.NewListFaultSimulationPager(string, string, string, *NodeTypesClientListFaultSimulationOptions) *runtime.Pager[NodeTypesClientListFaultSimulationResponse]`
- New function `*NodeTypesClient.NewListPager(string, string, string, *NodeTypesClientListOptions) *runtime.Pager[NodeTypesClientListResponse]`
- New function `*NodeTypesClient.BeginRedeploy(context.Context, string, string, string, NodeTypeActionParameters, *NodeTypesClientBeginRedeployOptions) (*runtime.Poller[NodeTypesClientRedeployResponse], error)`
- New function `*NodeTypesClient.BeginStart(context.Context, string, string, string, NodeTypeActionParameters, *NodeTypesClientBeginStartOptions) (*runtime.Poller[NodeTypesClientStartResponse], error)`
- New function `*NodeTypesClient.BeginStartFaultSimulation(context.Context, string, string, string, FaultSimulationContentClassification, *NodeTypesClientBeginStartFaultSimulationOptions) (*runtime.Poller[NodeTypesClientStartFaultSimulationResponse], error)`
- New function `*NodeTypesClient.BeginStopFaultSimulation(context.Context, string, string, string, FaultSimulationIDContent, *NodeTypesClientBeginStopFaultSimulationOptions) (*runtime.Poller[NodeTypesClientStopFaultSimulationResponse], error)`
- New struct `FaultSimulation`
- New struct `FaultSimulationConstraints`
- New struct `FaultSimulationDetails`
- New struct `FaultSimulationIDContent`
- New struct `FaultSimulationListResult`
- New struct `NodeTypeFaultSimulation`
- New struct `VMSSExtensionPropertiesProtectedSettings`
- New struct `VMSSExtensionPropertiesSettings`
- New struct `ZoneFaultSimulationContent`
- New field `Value` in struct `ManagedClusterVersionClientListByEnvironmentResponse`
- New field `Value` in struct `ManagedClusterVersionClientListResponse`


## 0.3.0 (2024-12-27)
### Breaking Changes

- Operation `*NodeTypesClient.Update` has been changed to LRO, use `*NodeTypesClient.BeginUpdate` instead.
- Field `CustomFqdn` of struct `ManagedClusterProperties` has been removed

### Features Added

- New field `AllocatedOutboundPorts` in struct `ManagedClusterProperties`


## 0.2.0 (2024-10-23)
### Features Added

- New enum type `AutoGeneratedDomainNameLabelScope` with values `AutoGeneratedDomainNameLabelScopeNoReuse`, `AutoGeneratedDomainNameLabelScopeResourceGroupReuse`, `AutoGeneratedDomainNameLabelScopeSubscriptionReuse`, `AutoGeneratedDomainNameLabelScopeTenantReuse`
- New struct `VMApplication`
- New field `AutoGeneratedDomainNameLabelScope`, `CustomFqdn` in struct `ManagedClusterProperties`
- New field `VMApplications` in struct `NodeTypeProperties`


## 0.1.0 (2024-07-29)
### Other Changes

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).