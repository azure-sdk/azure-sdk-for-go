# Release History

## 2.0.0 (2022-10-28)
### Breaking Changes

- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New struct `LogSpecification`
- New struct `OperationProperties`
- New struct `ServiceSpecification`
- New field `IsDataAction` in struct `OperationsDefinition`
- New field `Origin` in struct `OperationsDefinition`
- New field `Properties` in struct `OperationsDefinition`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/attestation/armattestation` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).