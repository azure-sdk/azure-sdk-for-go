# Release History

## 0.3.0 (2023-08-29)
### Breaking Changes

- Function `*VirtualMachinesClient.NewListPager` parameter(s) have been changed from `(*VirtualMachinesClientListOptions)` to `(string, *VirtualMachinesClientListOptions)`
- Type of `ClusterInventoryItem.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `ClusterProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `DatastoreInventoryItem.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `GuestAgentProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `HostInventoryItem.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `HostProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `HybridIdentityMetadataProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `InventoryItemProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `ResourcePoolInventoryItem.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `ResourcePoolProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `VCenterProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `VirtualMachineInventoryItem.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `VirtualMachineProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `VirtualMachineTemplateInventoryItem.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `VirtualMachineTemplateProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `VirtualNetworkInventoryItem.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `VirtualNetworkProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Function `*GuestAgentsClient.NewListByVMPager` has been removed
- Function `*HybridIdentityMetadataClient.NewListByVMPager` has been removed
- Function `*VirtualMachinesClient.BeginCreate` has been removed
- Function `*VirtualMachinesClient.NewListByResourceGroupPager` has been removed
- Field `Retain` of struct `VirtualMachinesClientBeginDeleteOptions` has been removed

### Features Added

- New function `NewAzureArcVMwareManagementServiceAPIClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AzureArcVMwareManagementServiceAPIClient, error)`
- New function `*AzureArcVMwareManagementServiceAPIClient.BeginUpgradeExtensions(context.Context, string, string, MachineExtensionUpgrade, *AzureArcVMwareManagementServiceAPIClientBeginUpgradeExtensionsOptions) (*runtime.Poller[AzureArcVMwareManagementServiceAPIClientUpgradeExtensionsResponse], error)`
- New function `*ClientFactory.NewAzureArcVMwareManagementServiceAPIClient() *AzureArcVMwareManagementServiceAPIClient`
- New function `*ClientFactory.NewVMInstanceGuestAgentsClient() *VMInstanceGuestAgentsClient`
- New function `*ClientFactory.NewVMInstanceHybridIdentityMetadataClient() *VMInstanceHybridIdentityMetadataClient`
- New function `*ClientFactory.NewVirtualMachineInstancesClient() *VirtualMachineInstancesClient`
- New function `*GuestAgentsClient.NewListPager(string, string, *GuestAgentsClientListOptions) *runtime.Pager[GuestAgentsClientListResponse]`
- New function `*HybridIdentityMetadataClient.NewListPager(string, string, *HybridIdentityMetadataClientListOptions) *runtime.Pager[HybridIdentityMetadataClientListResponse]`
- New function `NewVMInstanceGuestAgentsClient(azcore.TokenCredential, *arm.ClientOptions) (*VMInstanceGuestAgentsClient, error)`
- New function `*VMInstanceGuestAgentsClient.BeginCreate(context.Context, string, GuestAgent, *VMInstanceGuestAgentsClientBeginCreateOptions) (*runtime.Poller[VMInstanceGuestAgentsClientCreateResponse], error)`
- New function `*VMInstanceGuestAgentsClient.BeginDelete(context.Context, string, *VMInstanceGuestAgentsClientBeginDeleteOptions) (*runtime.Poller[VMInstanceGuestAgentsClientDeleteResponse], error)`
- New function `*VMInstanceGuestAgentsClient.Get(context.Context, string, *VMInstanceGuestAgentsClientGetOptions) (VMInstanceGuestAgentsClientGetResponse, error)`
- New function `*VMInstanceGuestAgentsClient.NewListPager(string, *VMInstanceGuestAgentsClientListOptions) *runtime.Pager[VMInstanceGuestAgentsClientListResponse]`
- New function `NewVMInstanceHybridIdentityMetadataClient(azcore.TokenCredential, *arm.ClientOptions) (*VMInstanceHybridIdentityMetadataClient, error)`
- New function `*VMInstanceHybridIdentityMetadataClient.Get(context.Context, string, *VMInstanceHybridIdentityMetadataClientGetOptions) (VMInstanceHybridIdentityMetadataClientGetResponse, error)`
- New function `*VMInstanceHybridIdentityMetadataClient.NewListPager(string, *VMInstanceHybridIdentityMetadataClientListOptions) *runtime.Pager[VMInstanceHybridIdentityMetadataClientListResponse]`
- New function `NewVirtualMachineInstancesClient(azcore.TokenCredential, *arm.ClientOptions) (*VirtualMachineInstancesClient, error)`
- New function `*VirtualMachineInstancesClient.BeginCreateOrUpdate(context.Context, string, VirtualMachineInstance, *VirtualMachineInstancesClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualMachineInstancesClientCreateOrUpdateResponse], error)`
- New function `*VirtualMachineInstancesClient.BeginDelete(context.Context, string, *VirtualMachineInstancesClientBeginDeleteOptions) (*runtime.Poller[VirtualMachineInstancesClientDeleteResponse], error)`
- New function `*VirtualMachineInstancesClient.Get(context.Context, string, *VirtualMachineInstancesClientGetOptions) (VirtualMachineInstancesClientGetResponse, error)`
- New function `*VirtualMachineInstancesClient.NewListPager(string, *VirtualMachineInstancesClientListOptions) *runtime.Pager[VirtualMachineInstancesClientListResponse]`
- New function `*VirtualMachineInstancesClient.BeginRestart(context.Context, string, *VirtualMachineInstancesClientBeginRestartOptions) (*runtime.Poller[VirtualMachineInstancesClientRestartResponse], error)`
- New function `*VirtualMachineInstancesClient.BeginStart(context.Context, string, *VirtualMachineInstancesClientBeginStartOptions) (*runtime.Poller[VirtualMachineInstancesClientStartResponse], error)`
- New function `*VirtualMachineInstancesClient.BeginStop(context.Context, string, *VirtualMachineInstancesClientBeginStopOptions) (*runtime.Poller[VirtualMachineInstancesClientStopResponse], error)`
- New function `*VirtualMachineInstancesClient.BeginUpdate(context.Context, string, VirtualMachineInstanceUpdate, *VirtualMachineInstancesClientBeginUpdateOptions) (*runtime.Poller[VirtualMachineInstancesClientUpdateResponse], error)`
- New function `*VirtualMachinesClient.BeginCreateOrUpdate(context.Context, string, string, VirtualMachine, *VirtualMachinesClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualMachinesClientCreateOrUpdateResponse], error)`
- New function `*VirtualMachinesClient.NewListAllPager(*VirtualMachinesClientListAllOptions) *runtime.Pager[VirtualMachinesClientListAllResponse]`
- New struct `ErrorAdditionalInfo`
- New struct `ExtensionTargetProperties`
- New struct `GuestAgentProfileUpdate`
- New struct `InfrastructureProfile`
- New struct `MachineExtensionUpgrade`
- New struct `OsProfileForVMInstance`
- New struct `VMInstanceHybridIdentityMetadata`
- New struct `VMInstanceHybridIdentityMetadataList`
- New struct `VMInstanceHybridIdentityMetadataProperties`
- New struct `VirtualMachineInstance`
- New struct `VirtualMachineInstanceProperties`
- New struct `VirtualMachineInstanceUpdate`
- New struct `VirtualMachineInstanceUpdateProperties`
- New struct `VirtualMachineInstancesList`
- New field `TotalCPUMHz`, `TotalMemoryGB`, `UsedCPUMHz`, `UsedMemoryGB` in struct `ClusterProperties`
- New field `CapacityGB`, `FreeSpaceGB` in struct `DatastoreProperties`
- New field `AdditionalInfo` in struct `ErrorDetail`
- New field `ClientPublicKey`, `MssqlDiscovered` in struct `GuestAgentProfile`
- New field `PrivateLinkScopeResourceID` in struct `GuestAgentProperties`
- New field `CPUMhz`, `DatastoreIDs`, `MemorySizeGB`, `NetworkIDs`, `OverallCPUUsageMHz`, `OverallMemoryUsageGB` in struct `HostProperties`
- New field `InventoryType` in struct `InventoryItemDetails`
- New field `CPUCapacityMHz`, `CPUOverallUsageMHz`, `DatastoreIDs`, `MemCapacityGB`, `MemOverallUsageGB`, `NetworkIDs` in struct `ResourcePoolProperties`
- New field `Cluster` in struct `VirtualMachineInventoryItem`
- New field `ToolsVersion`, `ToolsVersionStatus` in struct `VirtualMachineTemplateInventoryItem`
- New field `GuestAgentProfile` in struct `VirtualMachineUpdateProperties`
- New field `DeleteFromHost` in struct `VirtualMachinesClientBeginDeleteOptions`


## 0.2.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 0.2.0 (2023-03-28)
### Breaking Changes

- Struct `Condition` has been removed
- Struct `MachineExtensionInstanceView` has been removed

### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.1.0 (2022-08-09)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.1.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).