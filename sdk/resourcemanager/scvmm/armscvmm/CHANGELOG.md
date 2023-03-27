# Release History

## 0.3.0 (2023-03-27)
### Features Added

- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewAvailabilitySetsClient() *AvailabilitySetsClient`
- New function `*ClientFactory.NewCloudsClient() *CloudsClient`
- New function `*ClientFactory.NewInventoryItemsClient() *InventoryItemsClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewVirtualMachineTemplatesClient() *VirtualMachineTemplatesClient`
- New function `*ClientFactory.NewVirtualMachinesClient() *VirtualMachinesClient`
- New function `*ClientFactory.NewVirtualNetworksClient() *VirtualNetworksClient`
- New function `*ClientFactory.NewVmmServersClient() *VmmServersClient`
- New struct `ClientFactory`


## 0.2.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.2.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).