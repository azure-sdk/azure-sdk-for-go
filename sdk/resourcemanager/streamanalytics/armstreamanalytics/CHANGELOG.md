# Release History

## 2.0.0 (2023-06-10)
### Breaking Changes

- Type of `OutputProperties.SizeWindow` has been changed from `*float32` to `*int32`

### Features Added

- New function `*FileReferenceInputDataSource.GetReferenceInputDataSource() *ReferenceInputDataSource`
- New function `*GatewayMessageBusOutputDataSource.GetOutputDataSource() *OutputDataSource`
- New function `*GatewayMessageBusStreamInputDataSource.GetStreamInputDataSource() *StreamInputDataSource`
- New struct `FileReferenceInputDataSource`
- New struct `FileReferenceInputDataSourceProperties`
- New struct `GatewayMessageBusOutputDataSource`
- New struct `GatewayMessageBusOutputDataSourceProperties`
- New struct `GatewayMessageBusSourceProperties`
- New struct `GatewayMessageBusStreamInputDataSource`
- New struct `GatewayMessageBusStreamInputDataSourceProperties`
- New field `BlobPathPrefix` in struct `BlobOutputDataSourceProperties`


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).