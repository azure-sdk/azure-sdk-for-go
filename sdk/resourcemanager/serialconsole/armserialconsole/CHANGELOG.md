# Release History

## 2.0.0 (2023-12-22)
### Breaking Changes

- Function `*SerialPortsClient.Delete` has been removed
- Field `Value` of struct `MicrosoftSerialConsoleClientDisableConsoleResponse` has been removed
- Field `Value` of struct `MicrosoftSerialConsoleClientEnableConsoleResponse` has been removed
- Field `Value` of struct `MicrosoftSerialConsoleClientGetConsoleStatusResponse` has been removed

### Features Added

- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `SerialPortConnectionState` with values `SerialPortConnectionStateActive`, `SerialPortConnectionStateInactive`
- New struct `SystemData`
- New anonymous field `DisableSerialConsoleResult` in struct `MicrosoftSerialConsoleClientDisableConsoleResponse`
- New anonymous field `EnableSerialConsoleResult` in struct `MicrosoftSerialConsoleClientEnableConsoleResponse`
- New anonymous field `Status` in struct `MicrosoftSerialConsoleClientGetConsoleStatusResponse`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `SerialPort`
- New field `ConnectionState` in struct `SerialPortProperties`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/serialconsole/armserialconsole` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).