# Release History

## 2.0.0 (2023-11-11)
### Breaking Changes

- Function `timeRFC3339.MarshalText` has been removed
- Function `*timeRFC3339.Parse` has been removed
- Function `*timeRFC3339.UnmarshalText` has been removed
- Operation `*ApplicationTypeVersionsClient.List` has supported pagination, use `*ApplicationTypeVersionsClient.NewListPager` instead.
- Operation `*ApplicationTypesClient.List` has supported pagination, use `*ApplicationTypesClient.NewListPager` instead.
- Operation `*ApplicationsClient.List` has supported pagination, use `*ApplicationsClient.NewListPager` instead.
- Operation `*ClustersClient.List` has supported pagination, use `*ClustersClient.NewListPager` instead.
- Operation `*ClustersClient.ListByResourceGroup` has supported pagination, use `*ClustersClient.NewListByResourceGroupPager` instead.
- Operation `*ServicesClient.List` has supported pagination, use `*ServicesClient.NewListPager` instead.

### Features Added

- New function `dateTimeRFC3339.MarshalText() ([]byte, error)`
- New function `*dateTimeRFC3339.Parse(string) error`
- New function `*dateTimeRFC3339.UnmarshalText([]byte) error`


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).