# Release History

## 1.0.0 (2025-01-24)
### Breaking Changes

- Type of `OperationErrorDetails.ErrorDetails` has been changed from `*time.Time` to `*string`

### Features Added

- New field `AzureOperationName`, `Timestamp` in struct `OperationErrorDetails`
- New field `Timezone` in struct `ResourceOperationDetails`
- New field `Deadline`, `Timezone` in struct `Schedule`


## 0.1.0 (2024-09-27)
### Other Changes

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).