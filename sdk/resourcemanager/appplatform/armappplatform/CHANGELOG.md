# Release History

## 2.0.0 (2023-03-13)
### Breaking Changes

- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New type alias `APIPortalProvisioningState` with values `APIPortalProvisioningStateCreating`, `APIPortalProvisioningStateDeleting`, `APIPortalProvisioningStateFailed`, `APIPortalProvisioningStateSucceeded`, `APIPortalProvisioningStateUpdating`
- New type alias `BackendProtocol` with values `BackendProtocolDefault`, `BackendProtocolGRPC`
- New type alias `CertificateResourceProvisioningState` with values `CertificateResourceProvisioningStateCreating`, `CertificateResourceProvisioningStateDeleting`, `CertificateResourceProvisioningStateFailed`, `CertificateResourceProvisioningStateSucceeded`, `CertificateResourceProvisioningStateUpdating`
- New type alias `CustomDomainResourceProvisioningState` with values `CustomDomainResourceProvisioningStateCreating`, `CustomDomainResourceProvisioningStateDeleting`, `CustomDomainResourceProvisioningStateFailed`, `CustomDomainResourceProvisioningStateSucceeded`, `CustomDomainResourceProvisioningStateUpdating`
- New type alias `GatewayProvisioningState` with values `GatewayProvisioningStateCreating`, `GatewayProvisioningStateDeleting`, `GatewayProvisioningStateFailed`, `GatewayProvisioningStateSucceeded`, `GatewayProvisioningStateUpdating`
- New type alias `GatewayRouteConfigProtocol` with values `GatewayRouteConfigProtocolHTTP`, `GatewayRouteConfigProtocolHTTPS`
- New type alias `HTTPSchemeType` with values `HTTPSchemeTypeHTTP`, `HTTPSchemeTypeHTTPS`
- New type alias `PowerState` with values `PowerStateRunning`, `PowerStateStopped`
- New type alias `ProbeActionType` with values `ProbeActionTypeExecAction`, `ProbeActionTypeHTTPGetAction`, `ProbeActionTypeTCPSocketAction`
- New type alias `SessionAffinity` with values `SessionAffinityCookie`, `SessionAffinityNone`
- New type alias `StorageType` with values `StorageTypeStorageAccount`
- New type alias `Type` with values `TypeAzureFileVolume`
- New function `NewAPIPortalCustomDomainsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*APIPortalCustomDomainsClient, error)`
- New function `*APIPortalCustomDomainsClient.BeginCreateOrUpdate(context.Context, string, string, string, string, APIPortalCustomDomainResource, *APIPortalCustomDomainsClientBeginCreateOrUpdateOptions) (*runtime.Poller[APIPortalCustomDomainsClientCreateOrUpdateResponse], error)`
- New function `*APIPortalCustomDomainsClient.BeginDelete(context.Context, string, string, string, string, *APIPortalCustomDomainsClientBeginDeleteOptions) (*runtime.Poller[APIPortalCustomDomainsClientDeleteResponse], error)`
- New function `*APIPortalCustomDomainsClient.Get(context.Context, string, string, string, string, *APIPortalCustomDomainsClientGetOptions) (APIPortalCustomDomainsClientGetResponse, error)`
- New function `*APIPortalCustomDomainsClient.NewListPager(string, string, string, *APIPortalCustomDomainsClientListOptions) *runtime.Pager[APIPortalCustomDomainsClientListResponse]`
- New function `NewAPIPortalsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*APIPortalsClient, error)`
- New function `*APIPortalsClient.BeginCreateOrUpdate(context.Context, string, string, string, APIPortalResource, *APIPortalsClientBeginCreateOrUpdateOptions) (*runtime.Poller[APIPortalsClientCreateOrUpdateResponse], error)`
- New function `*APIPortalsClient.BeginDelete(context.Context, string, string, string, *APIPortalsClientBeginDeleteOptions) (*runtime.Poller[APIPortalsClientDeleteResponse], error)`
- New function `*APIPortalsClient.Get(context.Context, string, string, string, *APIPortalsClientGetOptions) (APIPortalsClientGetResponse, error)`
- New function `*APIPortalsClient.NewListPager(string, string, *APIPortalsClientListOptions) *runtime.Pager[APIPortalsClientListResponse]`
- New function `*APIPortalsClient.ValidateDomain(context.Context, string, string, string, CustomDomainValidatePayload, *APIPortalsClientValidateDomainOptions) (APIPortalsClientValidateDomainResponse, error)`
- New function `*AzureFileVolume.GetCustomPersistentDiskProperties() *CustomPersistentDiskProperties`
- New function `*BuildServiceBuilderClient.ListDeployments(context.Context, string, string, string, string, *BuildServiceBuilderClientListDeploymentsOptions) (BuildServiceBuilderClientListDeploymentsResponse, error)`
- New function `*CustomContainerUserSourceInfo.GetUserSourceInfo() *UserSourceInfo`
- New function `*CustomPersistentDiskProperties.GetCustomPersistentDiskProperties() *CustomPersistentDiskProperties`
- New function `*DeploymentsClient.BeginDisableRemoteDebugging(context.Context, string, string, string, string, *DeploymentsClientBeginDisableRemoteDebuggingOptions) (*runtime.Poller[DeploymentsClientDisableRemoteDebuggingResponse], error)`
- New function `*DeploymentsClient.BeginEnableRemoteDebugging(context.Context, string, string, string, string, *DeploymentsClientBeginEnableRemoteDebuggingOptions) (*runtime.Poller[DeploymentsClientEnableRemoteDebuggingResponse], error)`
- New function `*DeploymentsClient.GetRemoteDebuggingConfig(context.Context, string, string, string, string, *DeploymentsClientGetRemoteDebuggingConfigOptions) (DeploymentsClientGetRemoteDebuggingConfigResponse, error)`
- New function `*ExecAction.GetProbeAction() *ProbeAction`
- New function `NewGatewayCustomDomainsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*GatewayCustomDomainsClient, error)`
- New function `*GatewayCustomDomainsClient.BeginCreateOrUpdate(context.Context, string, string, string, string, GatewayCustomDomainResource, *GatewayCustomDomainsClientBeginCreateOrUpdateOptions) (*runtime.Poller[GatewayCustomDomainsClientCreateOrUpdateResponse], error)`
- New function `*GatewayCustomDomainsClient.BeginDelete(context.Context, string, string, string, string, *GatewayCustomDomainsClientBeginDeleteOptions) (*runtime.Poller[GatewayCustomDomainsClientDeleteResponse], error)`
- New function `*GatewayCustomDomainsClient.Get(context.Context, string, string, string, string, *GatewayCustomDomainsClientGetOptions) (GatewayCustomDomainsClientGetResponse, error)`
- New function `*GatewayCustomDomainsClient.NewListPager(string, string, string, *GatewayCustomDomainsClientListOptions) *runtime.Pager[GatewayCustomDomainsClientListResponse]`
- New function `NewGatewayRouteConfigsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*GatewayRouteConfigsClient, error)`
- New function `*GatewayRouteConfigsClient.BeginCreateOrUpdate(context.Context, string, string, string, string, GatewayRouteConfigResource, *GatewayRouteConfigsClientBeginCreateOrUpdateOptions) (*runtime.Poller[GatewayRouteConfigsClientCreateOrUpdateResponse], error)`
- New function `*GatewayRouteConfigsClient.BeginDelete(context.Context, string, string, string, string, *GatewayRouteConfigsClientBeginDeleteOptions) (*runtime.Poller[GatewayRouteConfigsClientDeleteResponse], error)`
- New function `*GatewayRouteConfigsClient.Get(context.Context, string, string, string, string, *GatewayRouteConfigsClientGetOptions) (GatewayRouteConfigsClientGetResponse, error)`
- New function `*GatewayRouteConfigsClient.NewListPager(string, string, string, *GatewayRouteConfigsClientListOptions) *runtime.Pager[GatewayRouteConfigsClientListResponse]`
- New function `NewGatewaysClient(string, azcore.TokenCredential, *arm.ClientOptions) (*GatewaysClient, error)`
- New function `*GatewaysClient.BeginCreateOrUpdate(context.Context, string, string, string, GatewayResource, *GatewaysClientBeginCreateOrUpdateOptions) (*runtime.Poller[GatewaysClientCreateOrUpdateResponse], error)`
- New function `*GatewaysClient.BeginDelete(context.Context, string, string, string, *GatewaysClientBeginDeleteOptions) (*runtime.Poller[GatewaysClientDeleteResponse], error)`
- New function `*GatewaysClient.Get(context.Context, string, string, string, *GatewaysClientGetOptions) (GatewaysClientGetResponse, error)`
- New function `*GatewaysClient.NewListPager(string, string, *GatewaysClientListOptions) *runtime.Pager[GatewaysClientListResponse]`
- New function `*GatewaysClient.ValidateDomain(context.Context, string, string, string, CustomDomainValidatePayload, *GatewaysClientValidateDomainOptions) (GatewaysClientValidateDomainResponse, error)`
- New function `*HTTPGetAction.GetProbeAction() *ProbeAction`
- New function `*ProbeAction.GetProbeAction() *ProbeAction`
- New function `*ServicesClient.BeginStart(context.Context, string, string, *ServicesClientBeginStartOptions) (*runtime.Poller[ServicesClientStartResponse], error)`
- New function `*ServicesClient.BeginStop(context.Context, string, string, *ServicesClientBeginStopOptions) (*runtime.Poller[ServicesClientStopResponse], error)`
- New function `*StorageAccount.GetStorageProperties() *StorageProperties`
- New function `*StorageProperties.GetStorageProperties() *StorageProperties`
- New function `NewStoragesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*StoragesClient, error)`
- New function `*StoragesClient.BeginCreateOrUpdate(context.Context, string, string, string, StorageResource, *StoragesClientBeginCreateOrUpdateOptions) (*runtime.Poller[StoragesClientCreateOrUpdateResponse], error)`
- New function `*StoragesClient.BeginDelete(context.Context, string, string, string, *StoragesClientBeginDeleteOptions) (*runtime.Poller[StoragesClientDeleteResponse], error)`
- New function `*StoragesClient.Get(context.Context, string, string, string, *StoragesClientGetOptions) (StoragesClientGetResponse, error)`
- New function `*StoragesClient.NewListPager(string, string, *StoragesClientListOptions) *runtime.Pager[StoragesClientListResponse]`
- New function `*TCPSocketAction.GetProbeAction() *ProbeAction`
- New struct `APIPortalCustomDomainProperties`
- New struct `APIPortalCustomDomainResource`
- New struct `APIPortalCustomDomainResourceCollection`
- New struct `APIPortalCustomDomainsClient`
- New struct `APIPortalInstance`
- New struct `APIPortalProperties`
- New struct `APIPortalResource`
- New struct `APIPortalResourceCollection`
- New struct `APIPortalResourceRequests`
- New struct `APIPortalsClient`
- New struct `AppVNetAddons`
- New struct `AzureFileVolume`
- New struct `BuildResourceRequests`
- New struct `ContainerProbeSettings`
- New struct `CustomContainer`
- New struct `CustomContainerUserSourceInfo`
- New struct `CustomPersistentDiskResource`
- New struct `DeploymentList`
- New struct `ExecAction`
- New struct `GatewayAPIMetadataProperties`
- New struct `GatewayAPIRoute`
- New struct `GatewayCorsProperties`
- New struct `GatewayCustomDomainProperties`
- New struct `GatewayCustomDomainResource`
- New struct `GatewayCustomDomainResourceCollection`
- New struct `GatewayCustomDomainsClient`
- New struct `GatewayInstance`
- New struct `GatewayOperatorProperties`
- New struct `GatewayOperatorResourceRequests`
- New struct `GatewayProperties`
- New struct `GatewayResource`
- New struct `GatewayResourceCollection`
- New struct `GatewayResourceRequests`
- New struct `GatewayRouteConfigOpenAPIProperties`
- New struct `GatewayRouteConfigProperties`
- New struct `GatewayRouteConfigResource`
- New struct `GatewayRouteConfigResourceCollection`
- New struct `GatewayRouteConfigsClient`
- New struct `GatewaysClient`
- New struct `HTTPGetAction`
- New struct `ImageRegistryCredential`
- New struct `IngressConfig`
- New struct `IngressSettings`
- New struct `IngressSettingsClientAuth`
- New struct `Probe`
- New struct `RemoteDebugging`
- New struct `RemoteDebuggingPayload`
- New struct `ServiceVNetAddons`
- New struct `SsoProperties`
- New struct `StorageAccount`
- New struct `StorageResource`
- New struct `StorageResourceCollection`
- New struct `StoragesClient`
- New struct `TCPSocketAction`
- New struct `UserAssignedManagedIdentity`
- New field `CustomPersistentDisks` in struct `AppResourceProperties`
- New field `IngressSettings` in struct `AppResourceProperties`
- New field `VnetAddons` in struct `AppResourceProperties`
- New field `ResourceRequests` in struct `BuildProperties`
- New field `ProvisioningState` in struct `CertificateProperties`
- New field `PowerState` in struct `ClusterResourceProperties`
- New field `VnetAddons` in struct `ClusterResourceProperties`
- New field `ProvisioningState` in struct `ContentCertificateProperties`
- New field `ProvisioningState` in struct `CustomDomainProperties`
- New field `ContainerProbeSettings` in struct `DeploymentSettings`
- New field `LivenessProbe` in struct `DeploymentSettings`
- New field `ReadinessProbe` in struct `DeploymentSettings`
- New field `StartupProbe` in struct `DeploymentSettings`
- New field `TerminationGracePeriodSeconds` in struct `DeploymentSettings`
- New field `ProvisioningState` in struct `KeyVaultCertificateProperties`
- New field `UserAssignedIdentities` in struct `ManagedIdentityProperties`
- New field `IngressConfig` in struct `NetworkProfile`
- New field `OutboundType` in struct `NetworkProfile`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).