# Release History

## 1.0.0 (2023-03-15)
### Features Added

- New type alias `Result` with values `ResultCanceled`, `ResultFailed`, `ResultSucceeded`
- New type alias `ServiceName` with values `ServiceNameSSH`, `ServiceNameWAC`
- New function `*EndpointsClient.ListIngressGatewayCredentials(context.Context, string, string, *EndpointsClientListIngressGatewayCredentialsOptions) (EndpointsClientListIngressGatewayCredentialsResponse, error)`
- New function `NewServiceconfigurationsClient(azcore.TokenCredential, *arm.ClientOptions) (*ServiceconfigurationsClient, error)`
- New function `*ServiceconfigurationsClient.Create(context.Context, string, string, string, ServiceConfigurationResource, *ServiceconfigurationsClientCreateOptions) (ServiceconfigurationsClientCreateResponse, error)`
- New function `*ServiceconfigurationsClient.Delete(context.Context, string, string, string, *ServiceconfigurationsClientDeleteOptions) (ServiceconfigurationsClientDeleteResponse, error)`
- New function `*ServiceconfigurationsClient.Get(context.Context, string, string, string, *ServiceconfigurationsClientGetOptions) (ServiceconfigurationsClientGetResponse, error)`
- New function `*ServiceconfigurationsClient.NewListByEndpointResourcePager(string, string, *ServiceconfigurationsClientListByEndpointResourceOptions) *runtime.Pager[ServiceconfigurationsClientListByEndpointResourceResponse]`
- New function `*ServiceconfigurationsClient.Update(context.Context, string, string, string, ServiceConfigurationResourcePatch, *ServiceconfigurationsClientUpdateOptions) (ServiceconfigurationsClientUpdateResponse, error)`
- New struct `ListCredentialsRequest`
- New struct `ServiceConfigurationList`
- New struct `ServiceConfigurationProperties`
- New struct `ServiceConfigurationPropertiesPatch`
- New struct `ServiceConfigurationResource`
- New struct `ServiceConfigurationResourcePatch`
- New struct `ServiceconfigurationsClient`
- New field `ListCredentialsRequest` in struct `EndpointsClientListCredentialsOptions`
- New field `ServiceName` in struct `ManagedProxyRequest`
- New field `ServiceConfigurationToken` in struct `RelayNamespaceAccessProperties`


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).