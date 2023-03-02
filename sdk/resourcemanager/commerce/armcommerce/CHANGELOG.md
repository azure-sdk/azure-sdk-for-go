# Release History

## 0.2.0 (2023-03-02)
### Breaking Changes

- Struct `InfoField` has been removed
- Field `UsageAggregationListResult` of struct `UsageAggregatesClientListResponse` has been removed

### Features Added

- Type of `UsageSample.InfoFields` has been changed from `*InfoField` to `any`
- New field `Value` in struct `UsageAggregatesClientListResponse`


## 0.1.0 (2022-06-10)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commerce/armcommerce` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.1.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).