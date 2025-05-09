# Release History

## 3.0.0 (2025-05-09)
### Breaking Changes

- Function `NewClientFactory` parameter(s) have been changed from `(string, azcore.TokenCredential, *arm.ClientOptions)` to `(string, string, azcore.TokenCredential, *arm.ClientOptions)`
- Type of `PrivateCloudIdentity.Type` has been changed from `*ResourceIdentityType` to `*SystemAssignedServiceIdentityType`
- Type of `ScriptExecutionProperties.NamedOutputs` has been changed from `map[string]any` to `map[string]*ScriptExecutionPropertiesNamedOutput`
- Enum `ResourceIdentityType` has been removed
- Function `NewAddonsClient` has been removed
- Function `*AddonsClient.BeginCreateOrUpdate` has been removed
- Function `*AddonsClient.BeginDelete` has been removed
- Function `*AddonsClient.Get` has been removed
- Function `*AddonsClient.NewListPager` has been removed
- Function `NewAuthorizationsClient` has been removed
- Function `*AuthorizationsClient.BeginCreateOrUpdate` has been removed
- Function `*AuthorizationsClient.BeginDelete` has been removed
- Function `*AuthorizationsClient.Get` has been removed
- Function `*AuthorizationsClient.NewListPager` has been removed
- Function `*ClientFactory.NewAddonsClient` has been removed
- Function `*ClientFactory.NewAuthorizationsClient` has been removed
- Function `*ClientFactory.NewCloudLinksClient` has been removed
- Function `*ClientFactory.NewClustersClient` has been removed
- Function `*ClientFactory.NewDatastoresClient` has been removed
- Function `*ClientFactory.NewGlobalReachConnectionsClient` has been removed
- Function `*ClientFactory.NewHcxEnterpriseSitesClient` has been removed
- Function `*ClientFactory.NewIscsiPathsClient` has been removed
- Function `*ClientFactory.NewLocationsClient` has been removed
- Function `*ClientFactory.NewOperationsClient` has been removed
- Function `*ClientFactory.NewPlacementPoliciesClient` has been removed
- Function `*ClientFactory.NewPrivateCloudsClient` has been removed
- Function `*ClientFactory.NewScriptCmdletsClient` has been removed
- Function `*ClientFactory.NewScriptExecutionsClient` has been removed
- Function `*ClientFactory.NewScriptPackagesClient` has been removed
- Function `*ClientFactory.NewVirtualMachinesClient` has been removed
- Function `*ClientFactory.NewWorkloadNetworksClient` has been removed
- Function `NewCloudLinksClient` has been removed
- Function `*CloudLinksClient.BeginCreateOrUpdate` has been removed
- Function `*CloudLinksClient.BeginDelete` has been removed
- Function `*CloudLinksClient.Get` has been removed
- Function `*CloudLinksClient.NewListPager` has been removed
- Function `NewClustersClient` has been removed
- Function `*ClustersClient.BeginCreateOrUpdate` has been removed
- Function `*ClustersClient.BeginDelete` has been removed
- Function `*ClustersClient.Get` has been removed
- Function `*ClustersClient.NewListPager` has been removed
- Function `*ClustersClient.ListZones` has been removed
- Function `*ClustersClient.BeginUpdate` has been removed
- Function `NewDatastoresClient` has been removed
- Function `*DatastoresClient.BeginCreateOrUpdate` has been removed
- Function `*DatastoresClient.BeginDelete` has been removed
- Function `*DatastoresClient.Get` has been removed
- Function `*DatastoresClient.NewListPager` has been removed
- Function `NewGlobalReachConnectionsClient` has been removed
- Function `*GlobalReachConnectionsClient.BeginCreateOrUpdate` has been removed
- Function `*GlobalReachConnectionsClient.BeginDelete` has been removed
- Function `*GlobalReachConnectionsClient.Get` has been removed
- Function `*GlobalReachConnectionsClient.NewListPager` has been removed
- Function `NewHcxEnterpriseSitesClient` has been removed
- Function `*HcxEnterpriseSitesClient.CreateOrUpdate` has been removed
- Function `*HcxEnterpriseSitesClient.Delete` has been removed
- Function `*HcxEnterpriseSitesClient.Get` has been removed
- Function `*HcxEnterpriseSitesClient.NewListPager` has been removed
- Function `NewIscsiPathsClient` has been removed
- Function `*IscsiPathsClient.BeginCreateOrUpdate` has been removed
- Function `*IscsiPathsClient.BeginDelete` has been removed
- Function `*IscsiPathsClient.Get` has been removed
- Function `*IscsiPathsClient.NewListByPrivateCloudPager` has been removed
- Function `NewLocationsClient` has been removed
- Function `*LocationsClient.CheckQuotaAvailability` has been removed
- Function `*LocationsClient.CheckTrialAvailability` has been removed
- Function `NewOperationsClient` has been removed
- Function `*OperationsClient.NewListPager` has been removed
- Function `NewPlacementPoliciesClient` has been removed
- Function `*PlacementPoliciesClient.BeginCreateOrUpdate` has been removed
- Function `*PlacementPoliciesClient.BeginDelete` has been removed
- Function `*PlacementPoliciesClient.Get` has been removed
- Function `*PlacementPoliciesClient.NewListPager` has been removed
- Function `*PlacementPoliciesClient.BeginUpdate` has been removed
- Function `NewPrivateCloudsClient` has been removed
- Function `*PrivateCloudsClient.BeginCreateOrUpdate` has been removed
- Function `*PrivateCloudsClient.BeginDelete` has been removed
- Function `*PrivateCloudsClient.Get` has been removed
- Function `*PrivateCloudsClient.ListAdminCredentials` has been removed
- Function `*PrivateCloudsClient.NewListInSubscriptionPager` has been removed
- Function `*PrivateCloudsClient.NewListPager` has been removed
- Function `*PrivateCloudsClient.BeginRotateNsxtPassword` has been removed
- Function `*PrivateCloudsClient.BeginRotateVcenterPassword` has been removed
- Function `*PrivateCloudsClient.BeginUpdate` has been removed
- Function `NewScriptCmdletsClient` has been removed
- Function `*ScriptCmdletsClient.Get` has been removed
- Function `*ScriptCmdletsClient.NewListPager` has been removed
- Function `NewScriptExecutionsClient` has been removed
- Function `*ScriptExecutionsClient.BeginCreateOrUpdate` has been removed
- Function `*ScriptExecutionsClient.BeginDelete` has been removed
- Function `*ScriptExecutionsClient.Get` has been removed
- Function `*ScriptExecutionsClient.GetExecutionLogs` has been removed
- Function `*ScriptExecutionsClient.NewListPager` has been removed
- Function `NewScriptPackagesClient` has been removed
- Function `*ScriptPackagesClient.Get` has been removed
- Function `*ScriptPackagesClient.NewListPager` has been removed
- Function `NewVirtualMachinesClient` has been removed
- Function `*VirtualMachinesClient.Get` has been removed
- Function `*VirtualMachinesClient.NewListPager` has been removed
- Function `*VirtualMachinesClient.BeginRestrictMovement` has been removed
- Function `NewWorkloadNetworksClient` has been removed
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
- Function `*WorkloadNetworksClient.Get` has been removed
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
- Function `*WorkloadNetworksClient.NewListPager` has been removed
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

### Features Added

- New enum type `HostKind` with values `HostKindGeneral`, `HostKindSpecialized`
- New enum type `HostMaintenance` with values `HostMaintenanceReplacement`, `HostMaintenanceUpgrade`
- New enum type `HostProvisioningState` with values `HostProvisioningStateCanceled`, `HostProvisioningStateFailed`, `HostProvisioningStateSucceeded`
- New enum type `ProvisionedNetworkProvisioningState` with values `ProvisionedNetworkProvisioningStateCanceled`, `ProvisionedNetworkProvisioningStateFailed`, `ProvisionedNetworkProvisioningStateSucceeded`
- New enum type `ProvisionedNetworkTypes` with values `ProvisionedNetworkTypesEsxManagement`, `ProvisionedNetworkTypesEsxReplication`, `ProvisionedNetworkTypesHcxManagement`, `ProvisionedNetworkTypesHcxUplink`, `ProvisionedNetworkTypesVcenterManagement`, `ProvisionedNetworkTypesVmotion`, `ProvisionedNetworkTypesVsan`
- New enum type `PureStoragePolicyProvisioningState` with values `PureStoragePolicyProvisioningStateCanceled`, `PureStoragePolicyProvisioningStateDeleting`, `PureStoragePolicyProvisioningStateFailed`, `PureStoragePolicyProvisioningStateSucceeded`, `PureStoragePolicyProvisioningStateUpdating`
- New enum type `ResourceSKUResourceType` with values `ResourceSKUResourceTypePrivateClouds`, `ResourceSKUResourceTypePrivateCloudsClusters`
- New enum type `ResourceSKURestrictionsReasonCode` with values `ResourceSKURestrictionsReasonCodeNotAvailableForSubscription`, `ResourceSKURestrictionsReasonCodeQuotaID`
- New enum type `ResourceSKURestrictionsType` with values `ResourceSKURestrictionsTypeLocation`, `ResourceSKURestrictionsTypeZone`
- New enum type `SystemAssignedServiceIdentityType` with values `SystemAssignedServiceIdentityTypeNone`, `SystemAssignedServiceIdentityTypeSystemAssigned`
- New function `NewClient(string, string, azcore.TokenCredential, *arm.ClientOptions) (*Client, error)`
- New function `*Client.CheckQuotaAvailability(context.Context, string, *ClientCheckQuotaAvailabilityOptions) (ClientCheckQuotaAvailabilityResponse, error)`
- New function `*Client.CheckTrialAvailability(context.Context, string, *ClientCheckTrialAvailabilityOptions) (ClientCheckTrialAvailabilityResponse, error)`
- New function `*Client.BeginCreate(context.Context, string, string, string, WorkloadNetworkVMGroup, *ClientBeginCreateOptions) (*runtime.Poller[ClientCreateResponse], error)`
- New function `*Client.BeginCreateOrUpdate(context.Context, string, string, string, PureStoragePolicy, *ClientBeginCreateOrUpdateOptions) (*runtime.Poller[ClientCreateOrUpdateResponse], error)`
- New function `*Client.CreateOrUpdate(context.Context, string, string, string, HcxEnterpriseSite, *ClientCreateOrUpdateOptions) (ClientCreateOrUpdateResponse, error)`
- New function `*Client.BeginDelete(context.Context, string, string, string, *ClientBeginDeleteOptions) (*runtime.Poller[ClientDeleteResponse], error)`
- New function `*Client.Delete(context.Context, string, string, string, *ClientDeleteOptions) (ClientDeleteResponse, error)`
- New function `*Client.BeginDeleteSegment(context.Context, string, string, string, *ClientBeginDeleteSegmentOptions) (*runtime.Poller[ClientDeleteSegmentResponse], error)`
- New function `*Client.Get(context.Context, string, string, string, *ClientGetOptions) (ClientGetResponse, error)`
- New function `*Client.GetExecutionLogs(context.Context, string, string, string, *ClientGetExecutionLogsOptions) (ClientGetExecutionLogsResponse, error)`
- New function `*Client.ListAdminCredentials(context.Context, string, string, *ClientListAdminCredentialsOptions) (ClientListAdminCredentialsResponse, error)`
- New function `*Client.NewListByPrivateCloudPager(string, string, *ClientListByPrivateCloudOptions) *runtime.Pager[ClientListByPrivateCloudResponse]`
- New function `*Client.NewListInSubscriptionPager(*ClientListInSubscriptionOptions) *runtime.Pager[ClientListInSubscriptionResponse]`
- New function `*Client.NewListPager(string, string, *ClientListOptions) *runtime.Pager[ClientListResponse]`
- New function `*Client.ListZones(context.Context, string, string, string, *ClientListZonesOptions) (ClientListZonesResponse, error)`
- New function `*Client.BeginRestrictMovement(context.Context, string, string, string, string, VirtualMachineRestrictMovement, *ClientBeginRestrictMovementOptions) (*runtime.Poller[ClientRestrictMovementResponse], error)`
- New function `*Client.BeginRotateNsxtPassword(context.Context, string, string, *ClientBeginRotateNsxtPasswordOptions) (*runtime.Poller[ClientRotateNsxtPasswordResponse], error)`
- New function `*Client.BeginRotateVcenterPassword(context.Context, string, string, *ClientBeginRotateVcenterPasswordOptions) (*runtime.Poller[ClientRotateVcenterPasswordResponse], error)`
- New function `*Client.BeginUpdate(context.Context, string, string, string, string, PlacementPolicyUpdate, *ClientBeginUpdateOptions) (*runtime.Poller[ClientUpdateResponse], error)`
- New function `*ClientFactory.NewClient() *Client`
- New function `*GeneralHostProperties.GetHostProperties() *HostProperties`
- New function `*HostProperties.GetHostProperties() *HostProperties`
- New function `*SpecializedHostProperties.GetHostProperties() *HostProperties`
- New struct `GeneralHostProperties`
- New struct `Host`
- New struct `HostListResult`
- New struct `PagedResourceSKU`
- New struct `ProvisionedNetwork`
- New struct `ProvisionedNetworkListResult`
- New struct `ProvisionedNetworkProperties`
- New struct `PureStoragePolicy`
- New struct `PureStoragePolicyListResult`
- New struct `PureStoragePolicyProperties`
- New struct `PureStorageVolume`
- New struct `ResourceSKU`
- New struct `ResourceSKUCapabilities`
- New struct `ResourceSKULocationInfo`
- New struct `ResourceSKURestrictionInfo`
- New struct `ResourceSKURestrictions`
- New struct `ResourceSKUZoneDetails`
- New struct `ScriptExecutionPropertiesNamedOutput`
- New struct `SpecializedHostProperties`
- New field `ManagementNetwork`, `UplinkNetwork` in struct `AddonHcxProperties`
- New field `PureStorageVolume` in struct `DatastoreProperties`
- New field `Zones` in struct `PrivateCloud`


## 2.0.0 (2024-09-26)
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