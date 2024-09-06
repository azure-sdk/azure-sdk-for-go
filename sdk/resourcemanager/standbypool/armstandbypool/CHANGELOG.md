# Release History

## 1.0.0 (2024-09-06)
### Features Added

- New function `*ClientFactory.NewStandbyContainerGroupPoolRuntimeViewsClient() *StandbyContainerGroupPoolRuntimeViewsClient`
- New function `*ClientFactory.NewStandbyVirtualMachinePoolRuntimeViewsClient() *StandbyVirtualMachinePoolRuntimeViewsClient`
- New function `NewStandbyContainerGroupPoolRuntimeViewsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*StandbyContainerGroupPoolRuntimeViewsClient, error)`
- New function `*StandbyContainerGroupPoolRuntimeViewsClient.Get(context.Context, string, string, string, *StandbyContainerGroupPoolRuntimeViewsClientGetOptions) (StandbyContainerGroupPoolRuntimeViewsClientGetResponse, error)`
- New function `*StandbyContainerGroupPoolRuntimeViewsClient.NewListByStandbyPoolPager(string, string, *StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolOptions) *runtime.Pager[StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolResponse]`
- New function `NewStandbyVirtualMachinePoolRuntimeViewsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*StandbyVirtualMachinePoolRuntimeViewsClient, error)`
- New function `*StandbyVirtualMachinePoolRuntimeViewsClient.Get(context.Context, string, string, string, *StandbyVirtualMachinePoolRuntimeViewsClientGetOptions) (StandbyVirtualMachinePoolRuntimeViewsClientGetResponse, error)`
- New function `*StandbyVirtualMachinePoolRuntimeViewsClient.NewListByStandbyPoolPager(string, string, *StandbyVirtualMachinePoolRuntimeViewsClientListByStandbyPoolOptions) *runtime.Pager[StandbyVirtualMachinePoolRuntimeViewsClientListByStandbyPoolResponse]`
- New struct `ContainerGroupInstanceCountSummary`
- New struct `PoolResourceStateCount`
- New struct `StandbyContainerGroupPoolRuntimeViewResource`
- New struct `StandbyContainerGroupPoolRuntimeViewResourceListResult`
- New struct `StandbyContainerGroupPoolRuntimeViewResourceProperties`
- New struct `StandbyVirtualMachinePoolRuntimeViewResource`
- New struct `StandbyVirtualMachinePoolRuntimeViewResourceListResult`
- New struct `StandbyVirtualMachinePoolRuntimeViewResourceProperties`
- New struct `VirtualMachineInstanceCountSummary`
- New field `MinReadyCapacity` in struct `StandbyVirtualMachinePoolElasticityProfile`
- New field `MinReadyCapacity` in struct `StandbyVirtualMachinePoolElasticityProfileUpdate`


## 0.1.0 (2024-04-26)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).