# Release History

## 0.2.0 (2024-06-13)
### Breaking Changes

- Function `*AssetEndpointProfilesClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, AssetEndpointProfileUpdate, *AssetEndpointProfilesClientBeginUpdateOptions)` to `(context.Context, string, string, AssetEndpointProfile, *AssetEndpointProfilesClientBeginUpdateOptions)`
- Function `*AssetsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, AssetUpdate, *AssetsClientBeginUpdateOptions)` to `(context.Context, string, string, Asset, *AssetsClientBeginUpdateOptions)`
- Struct `AssetEndpointProfileUpdate` has been removed
- Struct `AssetEndpointProfileUpdateProperties` has been removed
- Struct `AssetUpdate` has been removed
- Struct `AssetUpdateProperties` has been removed
- Struct `TransportAuthenticationUpdate` has been removed
- Struct `UserAuthenticationUpdate` has been removed
- Struct `UsernamePasswordCredentialsUpdate` has been removed
- Struct `X509CredentialsUpdate` has been removed


## 0.1.0 (2024-04-26)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).