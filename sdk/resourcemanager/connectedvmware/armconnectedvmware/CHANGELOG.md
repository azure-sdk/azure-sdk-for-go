# Release History

## 0.2.0 (2022-08-24)
### Breaking Changes

- Function `*VCentersClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, VCenter, *VCentersClientBeginCreateOptions)` to `(context.Context, string, string, *VCentersClientBeginCreateOptions)`
- Function `*VCentersClient.Update` parameter(s) have been changed from `(context.Context, string, string, ResourcePatch, *VCentersClientUpdateOptions)` to `(context.Context, string, string, *VCentersClientUpdateOptions)`
- Function `*GuestAgentsClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, string, GuestAgent, *GuestAgentsClientBeginCreateOptions)` to `(context.Context, string, string, string, *GuestAgentsClientBeginCreateOptions)`
- Function `*InventoryItemsClient.Create` parameter(s) have been changed from `(context.Context, string, string, string, InventoryItem, *InventoryItemsClientCreateOptions)` to `(context.Context, string, string, string, *InventoryItemsClientCreateOptions)`
- Function `*DatastoresClient.Update` parameter(s) have been changed from `(context.Context, string, string, ResourcePatch, *DatastoresClientUpdateOptions)` to `(context.Context, string, string, *DatastoresClientUpdateOptions)`
- Function `*HybridIdentityMetadataClient.Create` parameter(s) have been changed from `(context.Context, string, string, string, HybridIdentityMetadata, *HybridIdentityMetadataClientCreateOptions)` to `(context.Context, string, string, string, *HybridIdentityMetadataClientCreateOptions)`
- Function `*VirtualMachinesClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, VirtualMachineUpdate, *VirtualMachinesClientBeginUpdateOptions)` to `(context.Context, string, string, *VirtualMachinesClientBeginUpdateOptions)`
- Function `*VirtualNetworksClient.Update` parameter(s) have been changed from `(context.Context, string, string, ResourcePatch, *VirtualNetworksClientUpdateOptions)` to `(context.Context, string, string, *VirtualNetworksClientUpdateOptions)`
- Function `*DatastoresClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, Datastore, *DatastoresClientBeginCreateOptions)` to `(context.Context, string, string, *DatastoresClientBeginCreateOptions)`
- Function `*HostsClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, Host, *HostsClientBeginCreateOptions)` to `(context.Context, string, string, *HostsClientBeginCreateOptions)`
- Function `*ResourcePoolsClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, ResourcePool, *ResourcePoolsClientBeginCreateOptions)` to `(context.Context, string, string, *ResourcePoolsClientBeginCreateOptions)`
- Function `*HostsClient.Update` parameter(s) have been changed from `(context.Context, string, string, ResourcePatch, *HostsClientUpdateOptions)` to `(context.Context, string, string, *HostsClientUpdateOptions)`
- Function `*VirtualMachineTemplatesClient.Update` parameter(s) have been changed from `(context.Context, string, string, ResourcePatch, *VirtualMachineTemplatesClientUpdateOptions)` to `(context.Context, string, string, *VirtualMachineTemplatesClientUpdateOptions)`
- Function `*ClustersClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, Cluster, *ClustersClientBeginCreateOptions)` to `(context.Context, string, string, *ClustersClientBeginCreateOptions)`
- Function `*ResourcePoolsClient.Update` parameter(s) have been changed from `(context.Context, string, string, ResourcePatch, *ResourcePoolsClientUpdateOptions)` to `(context.Context, string, string, *ResourcePoolsClientUpdateOptions)`
- Function `*VirtualMachineTemplatesClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, VirtualMachineTemplate, *VirtualMachineTemplatesClientBeginCreateOptions)` to `(context.Context, string, string, *VirtualMachineTemplatesClientBeginCreateOptions)`
- Function `*VirtualNetworksClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, VirtualNetwork, *VirtualNetworksClientBeginCreateOptions)` to `(context.Context, string, string, *VirtualNetworksClientBeginCreateOptions)`
- Function `*VirtualMachinesClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, VirtualMachine, *VirtualMachinesClientBeginCreateOptions)` to `(context.Context, string, string, *VirtualMachinesClientBeginCreateOptions)`
- Function `*ClustersClient.Update` parameter(s) have been changed from `(context.Context, string, string, ResourcePatch, *ClustersClientUpdateOptions)` to `(context.Context, string, string, *ClustersClientUpdateOptions)`

### Features Added

- New field `Body` in struct `VirtualMachinesClientBeginUpdateOptions`
- New field `Body` in struct `VCentersClientUpdateOptions`
- New field `Body` in struct `VirtualNetworksClientUpdateOptions`
- New field `Body` in struct `HostsClientBeginCreateOptions`
- New field `Body` in struct `HostsClientUpdateOptions`
- New field `Body` in struct `GuestAgentsClientBeginCreateOptions`
- New field `Body` in struct `ResourcePoolsClientUpdateOptions`
- New field `Body` in struct `VirtualMachinesClientBeginCreateOptions`
- New field `Body` in struct `ClustersClientBeginCreateOptions`
- New field `Body` in struct `VirtualMachineTemplatesClientBeginCreateOptions`
- New field `Body` in struct `HybridIdentityMetadataClientCreateOptions`
- New field `Body` in struct `ClustersClientUpdateOptions`
- New field `Body` in struct `VirtualMachineTemplatesClientUpdateOptions`
- New field `Body` in struct `DatastoresClientUpdateOptions`
- New field `Body` in struct `InventoryItemsClientCreateOptions`
- New field `Body` in struct `DatastoresClientBeginCreateOptions`
- New field `Body` in struct `ResourcePoolsClientBeginCreateOptions`
- New field `Body` in struct `VirtualNetworksClientBeginCreateOptions`
- New field `Body` in struct `VCentersClientBeginCreateOptions`


## 0.1.0 (2022-08-09)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.1.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).