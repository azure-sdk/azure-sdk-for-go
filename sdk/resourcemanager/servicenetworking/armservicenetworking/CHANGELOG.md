# Release History

## 1.0.0 (2023-02-16)
### Breaking Changes

- Type alias `FrontendIPAddressVersion` has been removed
- Field `IPAddressVersion` of struct `FrontendProperties` has been removed
- Field `IPAddressVersion` of struct `FrontendUpdateProperties` has been removed

### Features Added

- New type alias `FrontendIPVersion` with values `FrontendIPVersionIPv4`, `FrontendIPVersionIPv6`
- New field `IPVersion` in struct `FrontendProperties`
- New field `IPVersion` in struct `FrontendUpdateProperties`


## 0.1.0 (2023-01-11)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).