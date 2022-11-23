# Release History

## 0.2.0 (2022-11-23)
### Breaking Changes

- Function `*ClustersClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, Cluster, *ClustersClientBeginCreateOptions)` to `(context.Context, string, string, *ClustersClientBeginCreateOptions)`
- Function `*ClustersClient.Update` parameter(s) have been changed from `(context.Context, string, string, ResourcePatch, *ClustersClientUpdateOptions)` to `(context.Context, string, string, *ClustersClientUpdateOptions)`
- Function `*DatastoresClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, Datastore, *DatastoresClientBeginCreateOptions)` to `(context.Context, string, string, *DatastoresClientBeginCreateOptions)`
- Function `*DatastoresClient.Update` parameter(s) have been changed from `(context.Context, string, string, ResourcePatch, *DatastoresClientUpdateOptions)` to `(context.Context, string, string, *DatastoresClientUpdateOptions)`
- Function `*GuestAgentsClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, string, GuestAgent, *GuestAgentsClientBeginCreateOptions)` to `(context.Context, string, string, string, *GuestAgentsClientBeginCreateOptions)`
- Function `*HostsClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, Host, *HostsClientBeginCreateOptions)` to `(context.Context, string, string, *HostsClientBeginCreateOptions)`
- Function `*HostsClient.Update` parameter(s) have been changed from `(context.Context, string, string, ResourcePatch, *HostsClientUpdateOptions)` to `(context.Context, string, string, *HostsClientUpdateOptions)`
- Function `*HybridIdentityMetadataClient.Create` parameter(s) have been changed from `(context.Context, string, string, string, HybridIdentityMetadata, *HybridIdentityMetadataClientCreateOptions)` to `(context.Context, string, string, string, *HybridIdentityMetadataClientCreateOptions)`
- Function `*InventoryItemsClient.Create` parameter(s) have been changed from `(context.Context, string, string, string, InventoryItem, *InventoryItemsClientCreateOptions)` to `(context.Context, string, string, string, *InventoryItemsClientCreateOptions)`
- Function `*ResourcePoolsClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, ResourcePool, *ResourcePoolsClientBeginCreateOptions)` to `(context.Context, string, string, *ResourcePoolsClientBeginCreateOptions)`
- Function `*ResourcePoolsClient.Update` parameter(s) have been changed from `(context.Context, string, string, ResourcePatch, *ResourcePoolsClientUpdateOptions)` to `(context.Context, string, string, *ResourcePoolsClientUpdateOptions)`
- Function `*VCentersClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, VCenter, *VCentersClientBeginCreateOptions)` to `(context.Context, string, string, *VCentersClientBeginCreateOptions)`
- Function `*VCentersClient.Update` parameter(s) have been changed from `(context.Context, string, string, ResourcePatch, *VCentersClientUpdateOptions)` to `(context.Context, string, string, *VCentersClientUpdateOptions)`
- Function `*VirtualMachineTemplatesClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, VirtualMachineTemplate, *VirtualMachineTemplatesClientBeginCreateOptions)` to `(context.Context, string, string, *VirtualMachineTemplatesClientBeginCreateOptions)`
- Function `*VirtualMachineTemplatesClient.Update` parameter(s) have been changed from `(context.Context, string, string, ResourcePatch, *VirtualMachineTemplatesClientUpdateOptions)` to `(context.Context, string, string, *VirtualMachineTemplatesClientUpdateOptions)`
- Function `*VirtualMachinesClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, VirtualMachineUpdate, *VirtualMachinesClientBeginUpdateOptions)` to `(context.Context, string, string, *VirtualMachinesClientBeginUpdateOptions)`
- Function `*VirtualMachinesClient.NewListPager` parameter(s) have been changed from `(*VirtualMachinesClientListOptions)` to `(string, *VirtualMachinesClientListOptions)`
- Function `*VirtualNetworksClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, VirtualNetwork, *VirtualNetworksClientBeginCreateOptions)` to `(context.Context, string, string, *VirtualNetworksClientBeginCreateOptions)`
- Function `*VirtualNetworksClient.Update` parameter(s) have been changed from `(context.Context, string, string, ResourcePatch, *VirtualNetworksClientUpdateOptions)` to `(context.Context, string, string, *VirtualNetworksClientUpdateOptions)`
- Type of `ErrorResponse.Error` has been changed from `*ErrorDefinition` to `*ErrorDetail`
- Function `*GuestAgentsClient.NewListByVMPager` has been removed
- Function `*HybridIdentityMetadataClient.NewListByVMPager` has been removed
- Function `*VirtualMachinesClient.BeginCreate` has been removed
- Function `*VirtualMachinesClient.NewListByResourceGroupPager` has been removed
- Struct `ErrorDefinition` has been removed
- Struct `GuestAgentsClientListByVMOptions` has been removed
- Struct `GuestAgentsClientListByVMResponse` has been removed
- Struct `HybridIdentityMetadataClientListByVMOptions` has been removed
- Struct `HybridIdentityMetadataClientListByVMResponse` has been removed
- Struct `VirtualMachinesClientBeginCreateOptions` has been removed
- Struct `VirtualMachinesClientCreateResponse` has been removed
- Struct `VirtualMachinesClientListByResourceGroupOptions` has been removed
- Struct `VirtualMachinesClientListByResourceGroupResponse` has been removed

### Features Added

- New function `NewAzureArcVMwareManagementServiceAPIClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AzureArcVMwareManagementServiceAPIClient, error)`
- New function `*AzureArcVMwareManagementServiceAPIClient.BeginUpgradeExtensions(context.Context, string, string, MachineExtensionUpgrade, *AzureArcVMwareManagementServiceAPIClientBeginUpgradeExtensionsOptions) (*runtime.Poller[AzureArcVMwareManagementServiceAPIClientUpgradeExtensionsResponse], error)`
- New function `*GuestAgentsClient.NewListPager(string, string, *GuestAgentsClientListOptions) *runtime.Pager[GuestAgentsClientListResponse]`
- New function `*HybridIdentityMetadataClient.NewListPager(string, string, *HybridIdentityMetadataClientListOptions) *runtime.Pager[HybridIdentityMetadataClientListResponse]`
- New function `*VirtualMachinesClient.BeginCreateOrUpdate(context.Context, string, string, *VirtualMachinesClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualMachinesClientCreateOrUpdateResponse], error)`
- New function `*VirtualMachinesClient.NewListAllPager(*VirtualMachinesClientListAllOptions) *runtime.Pager[VirtualMachinesClientListAllResponse]`
- New struct `AzureArcVMwareManagementServiceAPIClient`
- New struct `AzureArcVMwareManagementServiceAPIClientBeginUpgradeExtensionsOptions`
- New struct `AzureArcVMwareManagementServiceAPIClientUpgradeExtensionsResponse`
- New struct `ErrorAdditionalInfo`
- New struct `ExtensionTargetProperties`
- New struct `GuestAgentProfileUpdate`
- New struct `GuestAgentsClientListOptions`
- New struct `GuestAgentsClientListResponse`
- New struct `HybridIdentityMetadataClientListOptions`
- New struct `HybridIdentityMetadataClientListResponse`
- New struct `MachineExtensionUpgrade`
- New struct `VirtualMachinesClientBeginCreateOrUpdateOptions`
- New struct `VirtualMachinesClientCreateOrUpdateResponse`
- New struct `VirtualMachinesClientListAllOptions`
- New struct `VirtualMachinesClientListAllResponse`
- New field `Body` in struct `ClustersClientBeginCreateOptions`
- New field `Body` in struct `ClustersClientUpdateOptions`
- New field `CapacityGB` in struct `DatastoreProperties`
- New field `FreeSpaceGB` in struct `DatastoreProperties`
- New field `Body` in struct `DatastoresClientBeginCreateOptions`
- New field `Body` in struct `DatastoresClientUpdateOptions`
- New field `AdditionalInfo` in struct `ErrorDetail`
- New field `ClientPublicKey` in struct `GuestAgentProfile`
- New field `MssqlDiscovered` in struct `GuestAgentProfile`
- New field `Body` in struct `GuestAgentsClientBeginCreateOptions`
- New field `DatastoreIDs` in struct `HostProperties`
- New field `NetworkIDs` in struct `HostProperties`
- New field `Body` in struct `HostsClientBeginCreateOptions`
- New field `Body` in struct `HostsClientUpdateOptions`
- New field `Body` in struct `HybridIdentityMetadataClientCreateOptions`
- New field `InventoryType` in struct `InventoryItemDetails`
- New field `Body` in struct `InventoryItemsClientCreateOptions`
- New field `DatastoreIDs` in struct `ResourcePoolProperties`
- New field `NetworkIDs` in struct `ResourcePoolProperties`
- New field `Body` in struct `ResourcePoolsClientBeginCreateOptions`
- New field `Body` in struct `ResourcePoolsClientUpdateOptions`
- New field `Body` in struct `VCentersClientBeginCreateOptions`
- New field `Body` in struct `VCentersClientUpdateOptions`
- New field `PublicKey` in struct `VirtualMachineProperties`
- New field `ToolsVersion` in struct `VirtualMachineTemplateInventoryItem`
- New field `ToolsVersionStatus` in struct `VirtualMachineTemplateInventoryItem`
- New field `Body` in struct `VirtualMachineTemplatesClientBeginCreateOptions`
- New field `Body` in struct `VirtualMachineTemplatesClientUpdateOptions`
- New field `GuestAgentProfile` in struct `VirtualMachineUpdateProperties`
- New field `Body` in struct `VirtualMachinesClientBeginUpdateOptions`
- New field `Body` in struct `VirtualNetworksClientBeginCreateOptions`
- New field `Body` in struct `VirtualNetworksClientUpdateOptions`


## 0.1.0 (2022-08-09)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.1.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).