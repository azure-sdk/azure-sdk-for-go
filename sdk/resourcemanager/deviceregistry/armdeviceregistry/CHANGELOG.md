# Release History

## 0.2.0 (2024-07-17)
### Breaking Changes

- Type of `AssetEndpointProfileUpdateProperties.TransportAuthentication` has been changed from `*TransportAuthenticationUpdate` to `*TransportAuthentication`
- Type of `AssetEndpointProfileUpdateProperties.UserAuthentication` has been changed from `*UserAuthenticationUpdate` to `*UserAuthentication`
- Type of `ErrorAdditionalInfo.Info` has been changed from `any` to `*ErrorAdditionalInfoInfo`
- Type of `OperationStatusResult.PercentComplete` has been changed from `*float32` to `*float64`
- Struct `TransportAuthenticationUpdate` has been removed
- Struct `UserAuthenticationUpdate` has been removed
- Struct `UsernamePasswordCredentialsUpdate` has been removed
- Struct `X509CredentialsUpdate` has been removed

### Features Added

- New struct `ErrorAdditionalInfoInfo`


## 0.1.0 (2024-04-26)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).