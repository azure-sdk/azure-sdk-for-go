# Release History

## 0.2.0 (2023-08-25)
### Breaking Changes

- Function `NewClientFactory` parameter(s) have been changed from `(string, string, azcore.TokenCredential, *arm.ClientOptions)` to `(string, azcore.TokenCredential, *arm.ClientOptions)`
- Function `NewServicesClient` parameter(s) have been changed from `(string, string, azcore.TokenCredential, *arm.ClientOptions)` to `(string, azcore.TokenCredential, *arm.ClientOptions)`
- Function `*ServicesClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, Service, *ServicesClientCreateOrUpdateOptions)` to `(context.Context, string, string, Service, *ServicesClientCreateOrUpdateOptions)`
- Function `*ServicesClient.Delete` parameter(s) have been changed from `(context.Context, string, *ServicesClientDeleteOptions)` to `(context.Context, string, string, *ServicesClientDeleteOptions)`
- Function `*ServicesClient.Get` parameter(s) have been changed from `(context.Context, string, *ServicesClientGetOptions)` to `(context.Context, string, string, *ServicesClientGetOptions)`
- Function `*ServicesClient.Update` parameter(s) have been changed from `(context.Context, string, ServiceUpdate, *ServicesClientUpdateOptions)` to `(context.Context, string, string, ServiceUpdate, *ServicesClientUpdateOptions)`


## 0.1.0 (2023-08-25)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apicenter/armapicenter` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).