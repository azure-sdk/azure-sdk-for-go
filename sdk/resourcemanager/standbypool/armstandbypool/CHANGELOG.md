# Release History

## 0.2.0 (2024-07-17)
### Breaking Changes

- Type of `StandbyContainerGroupPoolResourceUpdateProperties.ContainerGroupProperties` has been changed from `*ContainerGroupPropertiesUpdate` to `*ContainerGroupProperties`
- Type of `StandbyContainerGroupPoolResourceUpdateProperties.ElasticityProfile` has been changed from `*StandbyContainerGroupPoolElasticityProfileUpdate` to `*StandbyContainerGroupPoolElasticityProfile`
- Type of `StandbyVirtualMachinePoolResourceUpdateProperties.ElasticityProfile` has been changed from `*StandbyVirtualMachinePoolElasticityProfileUpdate` to `*StandbyVirtualMachinePoolElasticityProfile`
- Struct `ContainerGroupProfileUpdate` has been removed
- Struct `ContainerGroupPropertiesUpdate` has been removed
- Struct `StandbyContainerGroupPoolElasticityProfileUpdate` has been removed
- Struct `StandbyVirtualMachinePoolElasticityProfileUpdate` has been removed


## 0.1.0 (2024-04-26)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).