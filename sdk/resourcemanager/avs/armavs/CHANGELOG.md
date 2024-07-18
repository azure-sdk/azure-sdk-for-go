# Release History

## 2.0.0 (2024-07-18)
### Breaking Changes

- Function `*WorkloadNetworksClient.Get` parameter(s) have been changed from `(context.Context, string, string, WorkloadNetworkName, *WorkloadNetworksClientGetOptions)` to `(context.Context, string, string, *WorkloadNetworksClientGetOptions)`
- Type of `Operation.Origin` has been changed from `*string` to `*Origin`
- Type of `PrivateCloud.Identity` has been changed from `*PrivateCloudIdentity` to `*SystemAssignedServiceIdentity`
- Type of `PrivateCloudUpdate.Identity` has been changed from `*PrivateCloudIdentity` to `*SystemAssignedServiceIdentity`
- Type of `ScriptExecutionProperties.NamedOutputs` has been changed from `map[string]any` to `map[string]*ScriptExecutionPropertiesNamedOutput`
- Enum `ResourceIdentityType` has been removed
- Enum `WorkloadNetworkName` has been removed
- Function `*VMPlacementPolicyProperties.GetPlacementPolicyProperties` has been removed
- Function `*WorkloadNetworksClient.BeginCreateDNSService` has been removed
- Function `*WorkloadNetworksClient.BeginCreateDNSZone` has been removed
- Function `*WorkloadNetworksClient.BeginCreateDhcp` has been removed
- Function `*WorkloadNetworksClient.BeginCreatePortMirroring` has been removed
- Function `*WorkloadNetworksClient.BeginCreatePublicIP` has been removed
- Function `*WorkloadNetworksClient.BeginCreateSegments` has been removed
- Function `*WorkloadNetworksClient.BeginCreateVMGroup` has been removed
- Function `*WorkloadNetworksClient.BeginDeleteDNSService` has been removed
- Function `*WorkloadNetworksClient.BeginDeleteDNSZone` has been removed
- Function `*WorkloadNetworksClient.BeginDeleteDhcp` has been removed
- Function `*WorkloadNetworksClient.BeginDeletePortMirroring` has been removed
- Function `*WorkloadNetworksClient.BeginDeletePublicIP` has been removed
- Function `*WorkloadNetworksClient.BeginDeleteSegment` has been removed
- Function `*WorkloadNetworksClient.BeginDeleteVMGroup` has been removed
- Function `*WorkloadNetworksClient.GetDNSService` has been removed
- Function `*WorkloadNetworksClient.GetDNSZone` has been removed
- Function `*WorkloadNetworksClient.GetDhcp` has been removed
- Function `*WorkloadNetworksClient.GetGateway` has been removed
- Function `*WorkloadNetworksClient.GetPortMirroring` has been removed
- Function `*WorkloadNetworksClient.GetPublicIP` has been removed
- Function `*WorkloadNetworksClient.GetSegment` has been removed
- Function `*WorkloadNetworksClient.GetVMGroup` has been removed
- Function `*WorkloadNetworksClient.GetVirtualMachine` has been removed
- Function `*WorkloadNetworksClient.NewListDNSServicesPager` has been removed
- Function `*WorkloadNetworksClient.NewListDNSZonesPager` has been removed
- Function `*WorkloadNetworksClient.NewListDhcpPager` has been removed
- Function `*WorkloadNetworksClient.NewListGatewaysPager` has been removed
- Function `*WorkloadNetworksClient.NewListPortMirroringPager` has been removed
- Function `*WorkloadNetworksClient.NewListPublicIPsPager` has been removed
- Function `*WorkloadNetworksClient.NewListSegmentsPager` has been removed
- Function `*WorkloadNetworksClient.NewListVMGroupsPager` has been removed
- Function `*WorkloadNetworksClient.NewListVirtualMachinesPager` has been removed
- Function `*WorkloadNetworksClient.BeginUpdateDNSService` has been removed
- Function `*WorkloadNetworksClient.BeginUpdateDNSZone` has been removed
- Function `*WorkloadNetworksClient.BeginUpdateDhcp` has been removed
- Function `*WorkloadNetworksClient.BeginUpdatePortMirroring` has been removed
- Function `*WorkloadNetworksClient.BeginUpdateSegments` has been removed
- Function `*WorkloadNetworksClient.BeginUpdateVMGroup` has been removed
- Operation `*ClustersClient.BeginUpdate` has been changed to non-LRO, use `*ClustersClient.Update` instead.
- Operation `*PlacementPoliciesClient.BeginUpdate` has been changed to non-LRO, use `*PlacementPoliciesClient.Update` instead.
- Operation `*PrivateCloudsClient.BeginUpdate` has been changed to non-LRO, use `*PrivateCloudsClient.Update` instead.
- Struct `LogSpecification` has been removed
- Struct `MetricDimension` has been removed
- Struct `MetricSpecification` has been removed
- Struct `OperationList` has been removed
- Struct `OperationProperties` has been removed
- Struct `PrivateCloudIdentity` has been removed
- Struct `ServiceSpecification` has been removed
- Struct `VMPlacementPolicyProperties` has been removed
- Field `SKU` of struct `LocationsClientCheckTrialAvailabilityOptions` has been removed
- Field `Properties` of struct `Operation` has been removed
- Field `OperationList` of struct `OperationsClientListResponse` has been removed

### Features Added

- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `CloudLinkProvisioningState` with values `CloudLinkProvisioningStateCanceled`, `CloudLinkProvisioningStateFailed`, `CloudLinkProvisioningStateSucceeded`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `DNSZoneType` with values `DNSZoneTypePrivate`, `DNSZoneTypePublic`
- New enum type `HcxEnterpriseSiteProvisioningState` with values `HcxEnterpriseSiteProvisioningStateCanceled`, `HcxEnterpriseSiteProvisioningStateFailed`, `HcxEnterpriseSiteProvisioningStateSucceeded`
- New enum type `IscsiPathProvisioningState` with values `IscsiPathProvisioningStateBuilding`, `IscsiPathProvisioningStateCanceled`, `IscsiPathProvisioningStateDeleting`, `IscsiPathProvisioningStateFailed`, `IscsiPathProvisioningStatePending`, `IscsiPathProvisioningStateSucceeded`, `IscsiPathProvisioningStateUpdating`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New enum type `SKUTier` with values `SKUTierBasic`, `SKUTierFree`, `SKUTierPremium`, `SKUTierStandard`
- New enum type `ScriptCmdletAudience` with values `ScriptCmdletAudienceAny`, `ScriptCmdletAudienceAutomation`
- New enum type `ScriptCmdletProvisioningState` with values `ScriptCmdletProvisioningStateCanceled`, `ScriptCmdletProvisioningStateFailed`, `ScriptCmdletProvisioningStateSucceeded`
- New enum type `ScriptPackageProvisioningState` with values `ScriptPackageProvisioningStateCanceled`, `ScriptPackageProvisioningStateFailed`, `ScriptPackageProvisioningStateSucceeded`
- New enum type `SystemAssignedServiceIdentityType` with values `SystemAssignedServiceIdentityTypeNone`, `SystemAssignedServiceIdentityTypeSystemAssigned`
- New enum type `VirtualMachineProvisioningState` with values `VirtualMachineProvisioningStateCanceled`, `VirtualMachineProvisioningStateFailed`, `VirtualMachineProvisioningStateSucceeded`
- New enum type `WorkloadNetworkProvisioningState` with values `WorkloadNetworkProvisioningStateBuilding`, `WorkloadNetworkProvisioningStateCanceled`, `WorkloadNetworkProvisioningStateDeleting`, `WorkloadNetworkProvisioningStateFailed`, `WorkloadNetworkProvisioningStateSucceeded`, `WorkloadNetworkProvisioningStateUpdating`
- New function `*ClientFactory.NewIscsiPathsClient() *IscsiPathsClient`
- New function `*ClientFactory.NewWorkloadNetworkDNSServicesClient() *WorkloadNetworkDNSServicesClient`
- New function `*ClientFactory.NewWorkloadNetworkDNSZonesClient() *WorkloadNetworkDNSZonesClient`
- New function `*ClientFactory.NewWorkloadNetworkDhcpConfigurationsClient() *WorkloadNetworkDhcpConfigurationsClient`
- New function `*ClientFactory.NewWorkloadNetworkGatewaysClient() *WorkloadNetworkGatewaysClient`
- New function `*ClientFactory.NewWorkloadNetworkPortMirroringProfilesClient() *WorkloadNetworkPortMirroringProfilesClient`
- New function `*ClientFactory.NewWorkloadNetworkPublicIPsClient() *WorkloadNetworkPublicIPsClient`
- New function `*ClientFactory.NewWorkloadNetworkSegmentsClient() *WorkloadNetworkSegmentsClient`
- New function `*ClientFactory.NewWorkloadNetworkVMGroupsClient() *WorkloadNetworkVMGroupsClient`
- New function `*ClientFactory.NewWorkloadNetworkVirtualMachinesClient() *WorkloadNetworkVirtualMachinesClient`
- New function `NewIscsiPathsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*IscsiPathsClient, error)`
- New function `*IscsiPathsClient.BeginCreateOrUpdate(context.Context, string, string, IscsiPath, *IscsiPathsClientBeginCreateOrUpdateOptions) (*runtime.Poller[IscsiPathsClientCreateOrUpdateResponse], error)`
- New function `*IscsiPathsClient.BeginDelete(context.Context, string, string, *IscsiPathsClientBeginDeleteOptions) (*runtime.Poller[IscsiPathsClientDeleteResponse], error)`
- New function `*IscsiPathsClient.Get(context.Context, string, string, *IscsiPathsClientGetOptions) (IscsiPathsClientGetResponse, error)`
- New function `*IscsiPathsClient.NewListPager(string, string, *IscsiPathsClientListOptions) *runtime.Pager[IscsiPathsClientListResponse]`
- New function `*VMVMPlacementPolicyProperties.GetPlacementPolicyProperties() *PlacementPolicyProperties`
- New function `NewWorkloadNetworkDNSServicesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*WorkloadNetworkDNSServicesClient, error)`
- New function `*WorkloadNetworkDNSServicesClient.BeginCreate(context.Context, string, string, string, WorkloadNetworkDNSService, *WorkloadNetworkDNSServicesClientBeginCreateOptions) (*runtime.Poller[WorkloadNetworkDNSServicesClientCreateResponse], error)`
- New function `*WorkloadNetworkDNSServicesClient.BeginDelete(context.Context, string, string, string, *WorkloadNetworkDNSServicesClientBeginDeleteOptions) (*runtime.Poller[WorkloadNetworkDNSServicesClientDeleteResponse], error)`
- New function `*WorkloadNetworkDNSServicesClient.Get(context.Context, string, string, string, *WorkloadNetworkDNSServicesClientGetOptions) (WorkloadNetworkDNSServicesClientGetResponse, error)`
- New function `*WorkloadNetworkDNSServicesClient.NewListPager(string, string, *WorkloadNetworkDNSServicesClientListOptions) *runtime.Pager[WorkloadNetworkDNSServicesClientListResponse]`
- New function `*WorkloadNetworkDNSServicesClient.BeginUpdate(context.Context, string, string, string, WorkloadNetworkDNSService, *WorkloadNetworkDNSServicesClientBeginUpdateOptions) (*runtime.Poller[WorkloadNetworkDNSServicesClientUpdateResponse], error)`
- New function `NewWorkloadNetworkDNSZonesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*WorkloadNetworkDNSZonesClient, error)`
- New function `*WorkloadNetworkDNSZonesClient.BeginCreate(context.Context, string, string, string, WorkloadNetworkDNSZone, *WorkloadNetworkDNSZonesClientBeginCreateOptions) (*runtime.Poller[WorkloadNetworkDNSZonesClientCreateResponse], error)`
- New function `*WorkloadNetworkDNSZonesClient.BeginDelete(context.Context, string, string, string, *WorkloadNetworkDNSZonesClientBeginDeleteOptions) (*runtime.Poller[WorkloadNetworkDNSZonesClientDeleteResponse], error)`
- New function `*WorkloadNetworkDNSZonesClient.Get(context.Context, string, string, string, *WorkloadNetworkDNSZonesClientGetOptions) (WorkloadNetworkDNSZonesClientGetResponse, error)`
- New function `*WorkloadNetworkDNSZonesClient.NewListPager(string, string, *WorkloadNetworkDNSZonesClientListOptions) *runtime.Pager[WorkloadNetworkDNSZonesClientListResponse]`
- New function `*WorkloadNetworkDNSZonesClient.BeginUpdate(context.Context, string, string, string, WorkloadNetworkDNSZone, *WorkloadNetworkDNSZonesClientBeginUpdateOptions) (*runtime.Poller[WorkloadNetworkDNSZonesClientUpdateResponse], error)`
- New function `NewWorkloadNetworkDhcpConfigurationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*WorkloadNetworkDhcpConfigurationsClient, error)`
- New function `*WorkloadNetworkDhcpConfigurationsClient.BeginCreate(context.Context, string, string, string, WorkloadNetworkDhcp, *WorkloadNetworkDhcpConfigurationsClientBeginCreateOptions) (*runtime.Poller[WorkloadNetworkDhcpConfigurationsClientCreateResponse], error)`
- New function `*WorkloadNetworkDhcpConfigurationsClient.BeginDelete(context.Context, string, string, string, *WorkloadNetworkDhcpConfigurationsClientBeginDeleteOptions) (*runtime.Poller[WorkloadNetworkDhcpConfigurationsClientDeleteResponse], error)`
- New function `*WorkloadNetworkDhcpConfigurationsClient.Get(context.Context, string, string, string, *WorkloadNetworkDhcpConfigurationsClientGetOptions) (WorkloadNetworkDhcpConfigurationsClientGetResponse, error)`
- New function `*WorkloadNetworkDhcpConfigurationsClient.NewListPager(string, string, *WorkloadNetworkDhcpConfigurationsClientListOptions) *runtime.Pager[WorkloadNetworkDhcpConfigurationsClientListResponse]`
- New function `*WorkloadNetworkDhcpConfigurationsClient.BeginUpdate(context.Context, string, string, string, WorkloadNetworkDhcp, *WorkloadNetworkDhcpConfigurationsClientBeginUpdateOptions) (*runtime.Poller[WorkloadNetworkDhcpConfigurationsClientUpdateResponse], error)`
- New function `NewWorkloadNetworkGatewaysClient(string, azcore.TokenCredential, *arm.ClientOptions) (*WorkloadNetworkGatewaysClient, error)`
- New function `*WorkloadNetworkGatewaysClient.Get(context.Context, string, string, string, *WorkloadNetworkGatewaysClientGetOptions) (WorkloadNetworkGatewaysClientGetResponse, error)`
- New function `*WorkloadNetworkGatewaysClient.NewListPager(string, string, *WorkloadNetworkGatewaysClientListOptions) *runtime.Pager[WorkloadNetworkGatewaysClientListResponse]`
- New function `NewWorkloadNetworkPortMirroringProfilesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*WorkloadNetworkPortMirroringProfilesClient, error)`
- New function `*WorkloadNetworkPortMirroringProfilesClient.BeginCreate(context.Context, string, string, string, WorkloadNetworkPortMirroring, *WorkloadNetworkPortMirroringProfilesClientBeginCreateOptions) (*runtime.Poller[WorkloadNetworkPortMirroringProfilesClientCreateResponse], error)`
- New function `*WorkloadNetworkPortMirroringProfilesClient.BeginDelete(context.Context, string, string, string, *WorkloadNetworkPortMirroringProfilesClientBeginDeleteOptions) (*runtime.Poller[WorkloadNetworkPortMirroringProfilesClientDeleteResponse], error)`
- New function `*WorkloadNetworkPortMirroringProfilesClient.Get(context.Context, string, string, string, *WorkloadNetworkPortMirroringProfilesClientGetOptions) (WorkloadNetworkPortMirroringProfilesClientGetResponse, error)`
- New function `*WorkloadNetworkPortMirroringProfilesClient.NewListPager(string, string, *WorkloadNetworkPortMirroringProfilesClientListOptions) *runtime.Pager[WorkloadNetworkPortMirroringProfilesClientListResponse]`
- New function `*WorkloadNetworkPortMirroringProfilesClient.BeginUpdate(context.Context, string, string, string, WorkloadNetworkPortMirroring, *WorkloadNetworkPortMirroringProfilesClientBeginUpdateOptions) (*runtime.Poller[WorkloadNetworkPortMirroringProfilesClientUpdateResponse], error)`
- New function `NewWorkloadNetworkPublicIPsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*WorkloadNetworkPublicIPsClient, error)`
- New function `*WorkloadNetworkPublicIPsClient.BeginCreate(context.Context, string, string, string, WorkloadNetworkPublicIP, *WorkloadNetworkPublicIPsClientBeginCreateOptions) (*runtime.Poller[WorkloadNetworkPublicIPsClientCreateResponse], error)`
- New function `*WorkloadNetworkPublicIPsClient.BeginDelete(context.Context, string, string, string, *WorkloadNetworkPublicIPsClientBeginDeleteOptions) (*runtime.Poller[WorkloadNetworkPublicIPsClientDeleteResponse], error)`
- New function `*WorkloadNetworkPublicIPsClient.Get(context.Context, string, string, string, *WorkloadNetworkPublicIPsClientGetOptions) (WorkloadNetworkPublicIPsClientGetResponse, error)`
- New function `*WorkloadNetworkPublicIPsClient.NewListPager(string, string, *WorkloadNetworkPublicIPsClientListOptions) *runtime.Pager[WorkloadNetworkPublicIPsClientListResponse]`
- New function `NewWorkloadNetworkSegmentsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*WorkloadNetworkSegmentsClient, error)`
- New function `*WorkloadNetworkSegmentsClient.BeginCreate(context.Context, string, string, string, WorkloadNetworkSegment, *WorkloadNetworkSegmentsClientBeginCreateOptions) (*runtime.Poller[WorkloadNetworkSegmentsClientCreateResponse], error)`
- New function `*WorkloadNetworkSegmentsClient.BeginDeleteSegment(context.Context, string, string, string, *WorkloadNetworkSegmentsClientBeginDeleteSegmentOptions) (*runtime.Poller[WorkloadNetworkSegmentsClientDeleteSegmentResponse], error)`
- New function `*WorkloadNetworkSegmentsClient.Get(context.Context, string, string, string, *WorkloadNetworkSegmentsClientGetOptions) (WorkloadNetworkSegmentsClientGetResponse, error)`
- New function `*WorkloadNetworkSegmentsClient.NewListPager(string, string, *WorkloadNetworkSegmentsClientListOptions) *runtime.Pager[WorkloadNetworkSegmentsClientListResponse]`
- New function `*WorkloadNetworkSegmentsClient.BeginUpdate(context.Context, string, string, string, WorkloadNetworkSegment, *WorkloadNetworkSegmentsClientBeginUpdateOptions) (*runtime.Poller[WorkloadNetworkSegmentsClientUpdateResponse], error)`
- New function `NewWorkloadNetworkVMGroupsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*WorkloadNetworkVMGroupsClient, error)`
- New function `*WorkloadNetworkVMGroupsClient.BeginCreate(context.Context, string, string, string, WorkloadNetworkVMGroup, *WorkloadNetworkVMGroupsClientBeginCreateOptions) (*runtime.Poller[WorkloadNetworkVMGroupsClientCreateResponse], error)`
- New function `*WorkloadNetworkVMGroupsClient.BeginDelete(context.Context, string, string, string, *WorkloadNetworkVMGroupsClientBeginDeleteOptions) (*runtime.Poller[WorkloadNetworkVMGroupsClientDeleteResponse], error)`
- New function `*WorkloadNetworkVMGroupsClient.Get(context.Context, string, string, string, *WorkloadNetworkVMGroupsClientGetOptions) (WorkloadNetworkVMGroupsClientGetResponse, error)`
- New function `*WorkloadNetworkVMGroupsClient.NewListPager(string, string, *WorkloadNetworkVMGroupsClientListOptions) *runtime.Pager[WorkloadNetworkVMGroupsClientListResponse]`
- New function `*WorkloadNetworkVMGroupsClient.BeginUpdate(context.Context, string, string, string, WorkloadNetworkVMGroup, *WorkloadNetworkVMGroupsClientBeginUpdateOptions) (*runtime.Poller[WorkloadNetworkVMGroupsClientUpdateResponse], error)`
- New function `NewWorkloadNetworkVirtualMachinesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*WorkloadNetworkVirtualMachinesClient, error)`
- New function `*WorkloadNetworkVirtualMachinesClient.Get(context.Context, string, string, string, *WorkloadNetworkVirtualMachinesClientGetOptions) (WorkloadNetworkVirtualMachinesClientGetResponse, error)`
- New function `*WorkloadNetworkVirtualMachinesClient.NewListPager(string, string, *WorkloadNetworkVirtualMachinesClientListOptions) *runtime.Pager[WorkloadNetworkVirtualMachinesClientListResponse]`
- New struct `ElasticSanVolume`
- New struct `IscsiPath`
- New struct `IscsiPathListResult`
- New struct `IscsiPathProperties`
- New struct `OperationListResult`
- New struct `ScriptExecutionPropertiesNamedOutput`
- New struct `SystemAssignedServiceIdentity`
- New struct `SystemData`
- New struct `VMVMPlacementPolicyProperties`
- New struct `WorkloadNetworkProperties`
- New field `SystemData` in struct `Addon`
- New field `SystemData` in struct `CloudLink`
- New field `ProvisioningState` in struct `CloudLinkProperties`
- New field `SystemData` in struct `Cluster`
- New field `VsanDatastoreName` in struct `ClusterProperties`
- New field `SKU` in struct `ClusterUpdate`
- New field `SystemData` in struct `Datastore`
- New field `ElasticSanVolume` in struct `DatastoreProperties`
- New field `HcxCloudManagerIP`, `NsxtManagerIP`, `VcenterIP` in struct `Endpoints`
- New field `SystemData` in struct `ExpressRouteAuthorization`
- New field `SystemData` in struct `GlobalReachConnection`
- New field `SystemData` in struct `HcxEnterpriseSite`
- New field `ProvisioningState` in struct `HcxEnterpriseSiteProperties`
- New field `Sku` in struct `LocationsClientCheckTrialAvailabilityOptions`
- New field `VsanDatastoreName` in struct `ManagementCluster`
- New field `ActionType` in struct `Operation`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New field `SystemData` in struct `PlacementPolicy`
- New field `SystemData` in struct `PrivateCloud`
- New field `DNSZoneType`, `VirtualNetworkID` in struct `PrivateCloudProperties`
- New field `SKU` in struct `PrivateCloudUpdate`
- New field `DNSZoneType` in struct `PrivateCloudUpdateProperties`
- New field `Capacity`, `Family`, `Size`, `Tier` in struct `SKU`
- New field `SystemData` in struct `ScriptCmdlet`
- New field `Audience`, `ProvisioningState` in struct `ScriptCmdletProperties`
- New field `SystemData` in struct `ScriptExecution`
- New field `SystemData` in struct `ScriptPackage`
- New field `ProvisioningState` in struct `ScriptPackageProperties`
- New field `SystemData` in struct `VirtualMachine`
- New field `ProvisioningState` in struct `VirtualMachineProperties`
- New field `Properties`, `SystemData` in struct `WorkloadNetwork`
- New field `SystemData` in struct `WorkloadNetworkDNSService`
- New field `SystemData` in struct `WorkloadNetworkDNSZone`
- New field `SystemData` in struct `WorkloadNetworkDhcp`
- New field `SystemData` in struct `WorkloadNetworkGateway`
- New field `ProvisioningState` in struct `WorkloadNetworkGatewayProperties`
- New field `SystemData` in struct `WorkloadNetworkPortMirroring`
- New field `SystemData` in struct `WorkloadNetworkPublicIP`
- New field `SystemData` in struct `WorkloadNetworkSegment`
- New field `SystemData` in struct `WorkloadNetworkVMGroup`
- New field `SystemData` in struct `WorkloadNetworkVirtualMachine`
- New field `ProvisioningState` in struct `WorkloadNetworkVirtualMachineProperties`


## 2.0.0-beta.1 (2024-06-28)
### Breaking Changes

- Function `*WorkloadNetworksClient.Get` parameter(s) have been changed from `(context.Context, string, string, WorkloadNetworkName, *WorkloadNetworksClientGetOptions)` to `(context.Context, string, string, *WorkloadNetworksClientGetOptions)`
- Type of `Operation.Origin` has been changed from `*string` to `*Origin`
- Enum `WorkloadNetworkName` has been removed
- Struct `LogSpecification` has been removed
- Struct `MetricDimension` has been removed
- Struct `MetricSpecification` has been removed
- Struct `OperationList` has been removed
- Struct `OperationProperties` has been removed
- Struct `ServiceSpecification` has been removed
- Field `Properties` of struct `Operation` has been removed
- Field `OperationList` of struct `OperationsClientListResponse` has been removed

### Features Added

- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `CloudLinkProvisioningState` with values `CloudLinkProvisioningStateCanceled`, `CloudLinkProvisioningStateFailed`, `CloudLinkProvisioningStateSucceeded`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `DNSZoneType` with values `DNSZoneTypePrivate`, `DNSZoneTypePublic`
- New enum type `HcxEnterpriseSiteProvisioningState` with values `HcxEnterpriseSiteProvisioningStateCanceled`, `HcxEnterpriseSiteProvisioningStateFailed`, `HcxEnterpriseSiteProvisioningStateSucceeded`
- New enum type `IscsiPathProvisioningState` with values `IscsiPathProvisioningStateBuilding`, `IscsiPathProvisioningStateCanceled`, `IscsiPathProvisioningStateDeleting`, `IscsiPathProvisioningStateFailed`, `IscsiPathProvisioningStatePending`, `IscsiPathProvisioningStateSucceeded`, `IscsiPathProvisioningStateUpdating`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New enum type `SKUTier` with values `SKUTierBasic`, `SKUTierFree`, `SKUTierPremium`, `SKUTierStandard`
- New enum type `ScriptCmdletAudience` with values `ScriptCmdletAudienceAny`, `ScriptCmdletAudienceAutomation`
- New enum type `ScriptCmdletProvisioningState` with values `ScriptCmdletProvisioningStateCanceled`, `ScriptCmdletProvisioningStateFailed`, `ScriptCmdletProvisioningStateSucceeded`
- New enum type `ScriptPackageProvisioningState` with values `ScriptPackageProvisioningStateCanceled`, `ScriptPackageProvisioningStateFailed`, `ScriptPackageProvisioningStateSucceeded`
- New enum type `VirtualMachineProvisioningState` with values `VirtualMachineProvisioningStateCanceled`, `VirtualMachineProvisioningStateFailed`, `VirtualMachineProvisioningStateSucceeded`
- New enum type `WorkloadNetworkProvisioningState` with values `WorkloadNetworkProvisioningStateBuilding`, `WorkloadNetworkProvisioningStateCanceled`, `WorkloadNetworkProvisioningStateDeleting`, `WorkloadNetworkProvisioningStateFailed`, `WorkloadNetworkProvisioningStateSucceeded`, `WorkloadNetworkProvisioningStateUpdating`
- New function `*ClientFactory.NewIscsiPathsClient() *IscsiPathsClient`
- New function `NewIscsiPathsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*IscsiPathsClient, error)`
- New function `*IscsiPathsClient.BeginCreateOrUpdate(context.Context, string, string, IscsiPath, *IscsiPathsClientBeginCreateOrUpdateOptions) (*runtime.Poller[IscsiPathsClientCreateOrUpdateResponse], error)`
- New function `*IscsiPathsClient.BeginDelete(context.Context, string, string, *IscsiPathsClientBeginDeleteOptions) (*runtime.Poller[IscsiPathsClientDeleteResponse], error)`
- New function `*IscsiPathsClient.Get(context.Context, string, string, *IscsiPathsClientGetOptions) (IscsiPathsClientGetResponse, error)`
- New function `*IscsiPathsClient.NewListByPrivateCloudPager(string, string, *IscsiPathsClientListByPrivateCloudOptions) *runtime.Pager[IscsiPathsClientListByPrivateCloudResponse]`
- New struct `ElasticSanVolume`
- New struct `IscsiPath`
- New struct `IscsiPathListResult`
- New struct `IscsiPathProperties`
- New struct `OperationListResult`
- New struct `SystemData`
- New struct `WorkloadNetworkProperties`
- New field `SystemData` in struct `Addon`
- New field `SystemData` in struct `CloudLink`
- New field `ProvisioningState` in struct `CloudLinkProperties`
- New field `SystemData` in struct `Cluster`
- New field `VsanDatastoreName` in struct `ClusterProperties`
- New field `SKU` in struct `ClusterUpdate`
- New field `SystemData` in struct `Datastore`
- New field `ElasticSanVolume` in struct `DatastoreProperties`
- New field `HcxCloudManagerIP`, `NsxtManagerIP`, `VcenterIP` in struct `Endpoints`
- New field `SystemData` in struct `ExpressRouteAuthorization`
- New field `SystemData` in struct `GlobalReachConnection`
- New field `SystemData` in struct `HcxEnterpriseSite`
- New field `ProvisioningState` in struct `HcxEnterpriseSiteProperties`
- New field `VsanDatastoreName` in struct `ManagementCluster`
- New field `ActionType` in struct `Operation`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New field `SystemData` in struct `PlacementPolicy`
- New field `SystemData` in struct `PrivateCloud`
- New field `DNSZoneType`, `VirtualNetworkID` in struct `PrivateCloudProperties`
- New field `SKU` in struct `PrivateCloudUpdate`
- New field `DNSZoneType` in struct `PrivateCloudUpdateProperties`
- New field `Capacity`, `Family`, `Size`, `Tier` in struct `SKU`
- New field `SystemData` in struct `ScriptCmdlet`
- New field `Audience`, `ProvisioningState` in struct `ScriptCmdletProperties`
- New field `SystemData` in struct `ScriptExecution`
- New field `SystemData` in struct `ScriptPackage`
- New field `ProvisioningState` in struct `ScriptPackageProperties`
- New field `SystemData` in struct `VirtualMachine`
- New field `ProvisioningState` in struct `VirtualMachineProperties`
- New field `Properties`, `SystemData` in struct `WorkloadNetwork`
- New field `SystemData` in struct `WorkloadNetworkDNSService`
- New field `SystemData` in struct `WorkloadNetworkDNSZone`
- New field `SystemData` in struct `WorkloadNetworkDhcp`
- New field `SystemData` in struct `WorkloadNetworkGateway`
- New field `ProvisioningState` in struct `WorkloadNetworkGatewayProperties`
- New field `SystemData` in struct `WorkloadNetworkPortMirroring`
- New field `SystemData` in struct `WorkloadNetworkPublicIP`
- New field `SystemData` in struct `WorkloadNetworkSegment`
- New field `SystemData` in struct `WorkloadNetworkVMGroup`
- New field `SystemData` in struct `WorkloadNetworkVirtualMachine`
- New field `ProvisioningState` in struct `WorkloadNetworkVirtualMachineProperties`


## 1.4.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.3.0 (2023-08-25)

### Features Added

- New field `ExtendedNetworkBlocks` in struct `PrivateCloudProperties`
- New field `ExtendedNetworkBlocks` in struct `PrivateCloudUpdateProperties`


## 1.2.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.

## 1.2.0 (2023-03-27)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.1.0 (2022-10-13)

### Features Added

- New const `ExpressRouteAuthorizationProvisioningStateCanceled`
- New const `AffinityStrengthShould`
- New const `AddonTypeArc`
- New const `PrivateCloudProvisioningStateCanceled`
- New const `NsxPublicIPQuotaRaisedEnumEnabled`
- New const `AzureHybridBenefitTypeNone`
- New const `WorkloadNetworkPublicIPProvisioningStateCanceled`
- New const `WorkloadNetworkDNSServiceProvisioningStateCanceled`
- New const `WorkloadNetworkSegmentProvisioningStateCanceled`
- New const `WorkloadNetworkDNSZoneProvisioningStateCanceled`
- New const `WorkloadNetworkNameDefault`
- New const `PlacementPolicyProvisioningStateCanceled`
- New const `WorkloadNetworkDhcpProvisioningStateCanceled`
- New const `WorkloadNetworkPortMirroringProvisioningStateCanceled`
- New const `WorkloadNetworkVMGroupProvisioningStateCanceled`
- New const `NsxPublicIPQuotaRaisedEnumDisabled`
- New const `DatastoreProvisioningStateCanceled`
- New const `AzureHybridBenefitTypeSQLHost`
- New const `AddonProvisioningStateCanceled`
- New const `ClusterProvisioningStateCanceled`
- New const `AffinityStrengthMust`
- New const `GlobalReachConnectionProvisioningStateCanceled`
- New const `ScriptExecutionProvisioningStateCanceled`
- New type alias `NsxPublicIPQuotaRaisedEnum`
- New type alias `AzureHybridBenefitType`
- New type alias `AffinityStrength`
- New type alias `WorkloadNetworkName`
- New function `PossibleAzureHybridBenefitTypeValues() []AzureHybridBenefitType`
- New function `*WorkloadNetworksClient.Get(context.Context, string, string, WorkloadNetworkName, *WorkloadNetworksClientGetOptions) (WorkloadNetworksClientGetResponse, error)`
- New function `*ClustersClient.ListZones(context.Context, string, string, string, *ClustersClientListZonesOptions) (ClustersClientListZonesResponse, error)`
- New function `PossibleNsxPublicIPQuotaRaisedEnumValues() []NsxPublicIPQuotaRaisedEnum`
- New function `PossibleWorkloadNetworkNameValues() []WorkloadNetworkName`
- New function `*WorkloadNetworksClient.NewListPager(string, string, *WorkloadNetworksClientListOptions) *runtime.Pager[WorkloadNetworksClientListResponse]`
- New function `*AddonArcProperties.GetAddonProperties() *AddonProperties`
- New function `PossibleAffinityStrengthValues() []AffinityStrength`
- New struct `AddonArcProperties`
- New struct `ClusterZone`
- New struct `ClusterZoneList`
- New struct `ClustersClientListZonesOptions`
- New struct `ClustersClientListZonesResponse`
- New struct `WorkloadNetwork`
- New struct `WorkloadNetworkList`
- New struct `WorkloadNetworksClientGetOptions`
- New struct `WorkloadNetworksClientGetResponse`
- New struct `WorkloadNetworksClientListOptions`
- New struct `WorkloadNetworksClientListResponse`
- New field `AffinityStrength` in struct `PlacementPolicyUpdateProperties`
- New field `AzureHybridBenefitType` in struct `PlacementPolicyUpdateProperties`
- New field `AutoDetectedKeyVersion` in struct `EncryptionKeyVaultProperties`
- New field `SKU` in struct `LocationsClientCheckTrialAvailabilityOptions`
- New field `AzureHybridBenefitType` in struct `VMHostPlacementPolicyProperties`
- New field `AffinityStrength` in struct `VMHostPlacementPolicyProperties`
- New field `NsxPublicIPQuotaRaised` in struct `PrivateCloudProperties`
- New field `Company` in struct `ScriptPackageProperties`
- New field `URI` in struct `ScriptPackageProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).