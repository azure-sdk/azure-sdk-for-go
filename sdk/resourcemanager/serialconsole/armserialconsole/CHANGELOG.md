# Release History

## 2.0.0 (2022-08-11)
### Breaking Changes

- Field `Disabled` of struct `DisableSerialConsoleResult` has been removed
- Field `Disabled` of struct `Status` has been removed
- Field `Disabled` of struct `EnableSerialConsoleResult` has been removed
- Field `Operations` of struct `MicrosoftSerialConsoleClientListOperationsResponse` has been removed

### Features Added

- New struct `OperationsForbidden`
- New struct `State`
- New field `Properties` in struct `DisableSerialConsoleResult`
- New field `Properties` in struct `EnableSerialConsoleResult`
- New field `Properties` in struct `Status`
- New field `Value` in struct `MicrosoftSerialConsoleClientListOperationsResponse`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/serialconsole/armserialconsole` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).