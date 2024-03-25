# Release History

## 0.2.0 (2024-03-25)
### Breaking Changes

- Operation `*SpringbootappsClient.BeginUpdate` has been changed to non-LRO, use `*SpringbootappsClient.Update` instead.
- Operation `*SpringbootserversClient.BeginUpdate` has been changed to non-LRO, use `*SpringbootserversClient.Update` instead.
- Field `Tags` of struct `SpringbootappsModel` has been removed
- Field `Tags` of struct `SpringbootappsPatch` has been removed
- Field `SiteName` of struct `SpringbootappsProperties` has been removed
- Field `Tags` of struct `SpringbootserversModel` has been removed
- Field `Tags` of struct `SpringbootserversPatch` has been removed

### Features Added

- New function `*SpringbootappsClient.BeginCreateOrUpdate(context.Context, string, string, string, SpringbootappsModel, *SpringbootappsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SpringbootappsClientCreateOrUpdateResponse], error)`
- New function `*SpringbootappsClient.BeginDelete(context.Context, string, string, string, *SpringbootappsClientBeginDeleteOptions) (*runtime.Poller[SpringbootappsClientDeleteResponse], error)`
- New field `Labels` in struct `SpringbootappsProperties`
- New field `Labels` in struct `SpringbootserversProperties`


## 0.1.0 (2024-02-01)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/springappdiscovery/armspringappdiscovery` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
