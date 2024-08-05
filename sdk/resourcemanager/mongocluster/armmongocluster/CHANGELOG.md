# Release History

## 0.2.0 (2024-08-05)
### Breaking Changes

- Type of `Properties.ClusterStatus` has been changed from `*MongoClusterStatus` to `*Status`
- Enum `MongoClusterStatus` has been removed

### Features Added

- New enum type `Status` with values `StatusDropping`, `StatusProvisioning`, `StatusReady`, `StatusStarting`, `StatusStopped`, `StatusStopping`, `StatusUpdating`


## 0.1.0 (2024-07-05)
### Other Changes

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).