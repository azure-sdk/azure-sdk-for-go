# Release History

## 0.4.0 (2023-09-11)
### Features Added

- New enum type `IdentityType` with values `IdentityTypeNone`, `IdentityTypeSystemAssigned`
- New enum type `ProvisioningAction` with values `ProvisioningActionInstall`, `ProvisioningActionRepair`, `ProvisioningActionUninstall`
- New enum type `StatusLevelTypes` with values `StatusLevelTypesError`, `StatusLevelTypesInfo`, `StatusLevelTypesWarning`
- New enum type `StatusTypes` with values `StatusTypesConnected`, `StatusTypesDisconnected`, `StatusTypesError`
- New function `*ClientFactory.NewGuestAgentsClient() *GuestAgentsClient`
- New function `*ClientFactory.NewHybridIdentityMetadatasClient() *HybridIdentityMetadatasClient`
- New function `*ClientFactory.NewMachineExtensionsClient() *MachineExtensionsClient`
- New function `*ClientFactory.NewVMInstanceGuestAgentsClient() *VMInstanceGuestAgentsClient`
- New function `*ClientFactory.NewVirtualMachineInstanceHybridIdentityMetadataClient() *VirtualMachineInstanceHybridIdentityMetadataClient`
- New function `*ClientFactory.NewVirtualMachineInstancesClient() *VirtualMachineInstancesClient`
- New function `NewGuestAgentsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*GuestAgentsClient, error)`
- New function `*GuestAgentsClient.BeginCreate(context.Context, string, string, string, *GuestAgentsClientBeginCreateOptions) (*runtime.Poller[GuestAgentsClientCreateResponse], error)`
- New function `*GuestAgentsClient.BeginDelete(context.Context, string, string, string, *GuestAgentsClientBeginDeleteOptions) (*runtime.Poller[GuestAgentsClientDeleteResponse], error)`
- New function `*GuestAgentsClient.Get(context.Context, string, string, string, *GuestAgentsClientGetOptions) (GuestAgentsClientGetResponse, error)`
- New function `*GuestAgentsClient.NewListByVMPager(string, string, *GuestAgentsClientListByVMOptions) *runtime.Pager[GuestAgentsClientListByVMResponse]`
- New function `NewHybridIdentityMetadatasClient(string, azcore.TokenCredential, *arm.ClientOptions) (*HybridIdentityMetadatasClient, error)`
- New function `*HybridIdentityMetadatasClient.Create(context.Context, string, string, string, *HybridIdentityMetadatasClientCreateOptions) (HybridIdentityMetadatasClientCreateResponse, error)`
- New function `*HybridIdentityMetadatasClient.Delete(context.Context, string, string, string, *HybridIdentityMetadatasClientDeleteOptions) (HybridIdentityMetadatasClientDeleteResponse, error)`
- New function `*HybridIdentityMetadatasClient.Get(context.Context, string, string, string, *HybridIdentityMetadatasClientGetOptions) (HybridIdentityMetadatasClientGetResponse, error)`
- New function `*HybridIdentityMetadatasClient.NewListByVMPager(string, string, *HybridIdentityMetadatasClientListByVMOptions) *runtime.Pager[HybridIdentityMetadatasClientListByVMResponse]`
- New function `NewMachineExtensionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*MachineExtensionsClient, error)`
- New function `*MachineExtensionsClient.BeginCreateOrUpdate(context.Context, string, string, string, MachineExtension, *MachineExtensionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[MachineExtensionsClientCreateOrUpdateResponse], error)`
- New function `*MachineExtensionsClient.BeginDelete(context.Context, string, string, string, *MachineExtensionsClientBeginDeleteOptions) (*runtime.Poller[MachineExtensionsClientDeleteResponse], error)`
- New function `*MachineExtensionsClient.Get(context.Context, string, string, string, *MachineExtensionsClientGetOptions) (MachineExtensionsClientGetResponse, error)`
- New function `*MachineExtensionsClient.NewListPager(string, string, *MachineExtensionsClientListOptions) *runtime.Pager[MachineExtensionsClientListResponse]`
- New function `*MachineExtensionsClient.BeginUpdate(context.Context, string, string, string, MachineExtensionUpdate, *MachineExtensionsClientBeginUpdateOptions) (*runtime.Poller[MachineExtensionsClientUpdateResponse], error)`
- New function `NewVMInstanceGuestAgentsClient(azcore.TokenCredential, *arm.ClientOptions) (*VMInstanceGuestAgentsClient, error)`
- New function `*VMInstanceGuestAgentsClient.BeginCreate(context.Context, string, *VMInstanceGuestAgentsClientBeginCreateOptions) (*runtime.Poller[VMInstanceGuestAgentsClientCreateResponse], error)`
- New function `*VMInstanceGuestAgentsClient.BeginDelete(context.Context, string, *VMInstanceGuestAgentsClientBeginDeleteOptions) (*runtime.Poller[VMInstanceGuestAgentsClientDeleteResponse], error)`
- New function `*VMInstanceGuestAgentsClient.Get(context.Context, string, *VMInstanceGuestAgentsClientGetOptions) (VMInstanceGuestAgentsClientGetResponse, error)`
- New function `*VMInstanceGuestAgentsClient.NewListPager(string, *VMInstanceGuestAgentsClientListOptions) *runtime.Pager[VMInstanceGuestAgentsClientListResponse]`
- New function `NewVirtualMachineInstanceHybridIdentityMetadataClient(azcore.TokenCredential, *arm.ClientOptions) (*VirtualMachineInstanceHybridIdentityMetadataClient, error)`
- New function `*VirtualMachineInstanceHybridIdentityMetadataClient.Get(context.Context, string, *VirtualMachineInstanceHybridIdentityMetadataClientGetOptions) (VirtualMachineInstanceHybridIdentityMetadataClientGetResponse, error)`
- New function `*VirtualMachineInstanceHybridIdentityMetadataClient.NewListPager(string, *VirtualMachineInstanceHybridIdentityMetadataClientListOptions) *runtime.Pager[VirtualMachineInstanceHybridIdentityMetadataClientListResponse]`
- New function `NewVirtualMachineInstancesClient(azcore.TokenCredential, *arm.ClientOptions) (*VirtualMachineInstancesClient, error)`
- New function `*VirtualMachineInstancesClient.BeginCreateCheckpoint(context.Context, string, *VirtualMachineInstancesClientBeginCreateCheckpointOptions) (*runtime.Poller[VirtualMachineInstancesClientCreateCheckpointResponse], error)`
- New function `*VirtualMachineInstancesClient.BeginCreateOrUpdate(context.Context, string, *VirtualMachineInstancesClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualMachineInstancesClientCreateOrUpdateResponse], error)`
- New function `*VirtualMachineInstancesClient.BeginDelete(context.Context, string, *VirtualMachineInstancesClientBeginDeleteOptions) (*runtime.Poller[VirtualMachineInstancesClientDeleteResponse], error)`
- New function `*VirtualMachineInstancesClient.BeginDeleteCheckpoint(context.Context, string, *VirtualMachineInstancesClientBeginDeleteCheckpointOptions) (*runtime.Poller[VirtualMachineInstancesClientDeleteCheckpointResponse], error)`
- New function `*VirtualMachineInstancesClient.Get(context.Context, string, *VirtualMachineInstancesClientGetOptions) (VirtualMachineInstancesClientGetResponse, error)`
- New function `*VirtualMachineInstancesClient.List(context.Context, string, *VirtualMachineInstancesClientListOptions) (VirtualMachineInstancesClientListResponse, error)`
- New function `*VirtualMachineInstancesClient.BeginRestart(context.Context, string, *VirtualMachineInstancesClientBeginRestartOptions) (*runtime.Poller[VirtualMachineInstancesClientRestartResponse], error)`
- New function `*VirtualMachineInstancesClient.BeginRestoreCheckpoint(context.Context, string, *VirtualMachineInstancesClientBeginRestoreCheckpointOptions) (*runtime.Poller[VirtualMachineInstancesClientRestoreCheckpointResponse], error)`
- New function `*VirtualMachineInstancesClient.BeginStart(context.Context, string, *VirtualMachineInstancesClientBeginStartOptions) (*runtime.Poller[VirtualMachineInstancesClientStartResponse], error)`
- New function `*VirtualMachineInstancesClient.BeginStop(context.Context, string, *VirtualMachineInstancesClientBeginStopOptions) (*runtime.Poller[VirtualMachineInstancesClientStopResponse], error)`
- New function `*VirtualMachineInstancesClient.BeginUpdate(context.Context, string, *VirtualMachineInstancesClientBeginUpdateOptions) (*runtime.Poller[VirtualMachineInstancesClientUpdateResponse], error)`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `GuestAgent`
- New struct `GuestAgentList`
- New struct `GuestAgentProfile`
- New struct `GuestAgentProperties`
- New struct `GuestCredential`
- New struct `HTTPProxyConfiguration`
- New struct `HybridIdentityMetadata`
- New struct `HybridIdentityMetadataList`
- New struct `HybridIdentityMetadataProperties`
- New struct `Identity`
- New struct `InfrastructureProfile`
- New struct `InfrastructureProfileUpdate`
- New struct `MachineExtension`
- New struct `MachineExtensionInstanceView`
- New struct `MachineExtensionInstanceViewStatus`
- New struct `MachineExtensionProperties`
- New struct `MachineExtensionPropertiesInstanceView`
- New struct `MachineExtensionUpdate`
- New struct `MachineExtensionUpdateProperties`
- New struct `MachineExtensionsListResult`
- New struct `OsProfileForVMInstance`
- New struct `TrackedResource`
- New struct `VMInstanceHybridIdentityMetadata`
- New struct `VMInstanceHybridIdentityMetadataList`
- New struct `VMInstanceHybridIdentityMetadataProperties`
- New struct `VirtualMachineInstance`
- New struct `VirtualMachineInstanceListResult`
- New struct `VirtualMachineInstanceProperties`
- New struct `VirtualMachineInstanceUpdate`
- New struct `VirtualMachineInstanceUpdateProperties`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`
- New field `Identity` in struct `VirtualMachine`
- New field `BiosGUID`, `ManagedMachineResourceID`, `OSVersion` in struct `VirtualMachineInventoryItem`
- New field `GuestAgentProfile`, `LastRestoredVMCheckpoint` in struct `VirtualMachineProperties`
- New field `Identity` in struct `VirtualMachineUpdate`
- New field `CheckpointType` in struct `VirtualMachineUpdateProperties`
- New field `DeleteFromHost` in struct `VirtualMachinesClientBeginDeleteOptions`


## 0.3.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.2.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.2.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).