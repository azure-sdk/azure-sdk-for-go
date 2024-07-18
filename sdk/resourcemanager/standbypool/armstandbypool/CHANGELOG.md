# Release History

## 0.2.0 (2024-07-18)
### Breaking Changes

- Function `*StandbyContainerGroupPoolsClient.Update` parameter(s) have been changed from `(context.Context, string, string, StandbyContainerGroupPoolResourceUpdate, *StandbyContainerGroupPoolsClientUpdateOptions)` to `(context.Context, string, string, StandbyContainerGroupPoolResource, *StandbyContainerGroupPoolsClientUpdateOptions)`
- Function `*StandbyVirtualMachinePoolsClient.Update` parameter(s) have been changed from `(context.Context, string, string, StandbyVirtualMachinePoolResourceUpdate, *StandbyVirtualMachinePoolsClientUpdateOptions)` to `(context.Context, string, string, StandbyVirtualMachinePoolResource, *StandbyVirtualMachinePoolsClientUpdateOptions)`
- Function `*StandbyVirtualMachinesClient.NewListByStandbyVirtualMachinePoolResourcePager` has been removed
- Struct `ContainerGroupProfileUpdate` has been removed
- Struct `ContainerGroupPropertiesUpdate` has been removed
- Struct `StandbyContainerGroupPoolElasticityProfileUpdate` has been removed
- Struct `StandbyContainerGroupPoolResourceUpdate` has been removed
- Struct `StandbyContainerGroupPoolResourceUpdateProperties` has been removed
- Struct `StandbyVirtualMachinePoolElasticityProfileUpdate` has been removed
- Struct `StandbyVirtualMachinePoolResourceUpdate` has been removed
- Struct `StandbyVirtualMachinePoolResourceUpdateProperties` has been removed

### Features Added

- New function `*StandbyVirtualMachinesClient.NewListByStandbyPoolPager(string, string, *StandbyVirtualMachinesClientListByStandbyPoolOptions) *runtime.Pager[StandbyVirtualMachinesClientListByStandbyPoolResponse]`


## 0.1.0 (2024-04-26)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).