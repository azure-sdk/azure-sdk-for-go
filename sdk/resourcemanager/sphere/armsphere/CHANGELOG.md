# Release History

## 0.3.0 (2023-11-30)
### Breaking Changes

- Function `*CatalogsClient.CountDevices` parameter(s) have been changed from `(context.Context, string, string, *CatalogsClientCountDevicesOptions)` to `(context.Context, string, string, any, *CatalogsClientCountDevicesOptions)`
- Function `*CatalogsClient.NewListDeploymentsPager` parameter(s) have been changed from `(string, string, *CatalogsClientListDeploymentsOptions)` to `(string, string, any, *CatalogsClientListDeploymentsOptions)`
- Function `*CatalogsClient.NewListDeviceInsightsPager` parameter(s) have been changed from `(string, string, *CatalogsClientListDeviceInsightsOptions)` to `(string, string, any, *CatalogsClientListDeviceInsightsOptions)`
- Function `*CatalogsClient.NewListDevicesPager` parameter(s) have been changed from `(string, string, *CatalogsClientListDevicesOptions)` to `(string, string, any, *CatalogsClientListDevicesOptions)`
- Function `*CatalogsClient.Update` parameter(s) have been changed from `(context.Context, string, string, CatalogUpdate, *CatalogsClientUpdateOptions)` to `(context.Context, string, string, CatalogTagsUpdate, *CatalogsClientUpdateOptions)`
- Function `*CertificatesClient.RetrieveCertChain` parameter(s) have been changed from `(context.Context, string, string, string, *CertificatesClientRetrieveCertChainOptions)` to `(context.Context, string, string, string, any, *CertificatesClientRetrieveCertChainOptions)`
- Function `*DeviceGroupsClient.CountDevices` parameter(s) have been changed from `(context.Context, string, string, string, string, *DeviceGroupsClientCountDevicesOptions)` to `(context.Context, string, string, string, string, any, *DeviceGroupsClientCountDevicesOptions)`
- Function `*DevicesClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, string, DeviceUpdate, *DevicesClientBeginUpdateOptions)` to `(context.Context, string, string, string, string, string, any, *DevicesClientBeginUpdateOptions)`
- Function `*ProductsClient.CountDevices` parameter(s) have been changed from `(context.Context, string, string, string, *ProductsClientCountDevicesOptions)` to `(context.Context, string, string, string, any, *ProductsClientCountDevicesOptions)`
- Function `*ProductsClient.NewGenerateDefaultDeviceGroupsPager` parameter(s) have been changed from `(string, string, string, *ProductsClientGenerateDefaultDeviceGroupsOptions)` to `(string, string, string, any, *ProductsClientGenerateDefaultDeviceGroupsOptions)`
- Struct `CatalogUpdate` has been removed
- Struct `DeviceUpdate` has been removed
- Struct `DeviceUpdateProperties` has been removed

### Features Added

- New struct `CatalogTagsUpdate`


## 0.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 0.1.0 (2023-07-28)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sphere/armsphere` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).