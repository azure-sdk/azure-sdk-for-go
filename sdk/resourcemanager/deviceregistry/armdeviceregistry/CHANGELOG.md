# Release History

## 0.2.0 (2024-09-09)
### Breaking Changes

- Type of `AssetProperties.Version` has been changed from `*int32` to `*int64`
- Type of `AssetStatus.Version` has been changed from `*int32` to `*int64`
- Type of `DataPoint.ObservabilityMode` has been changed from `*DataPointsObservabilityMode` to `*DataPointObservabilityMode`
- Type of `Event.ObservabilityMode` has been changed from `*EventsObservabilityMode` to `*EventObservabilityMode`
- Enum `DataPointsObservabilityMode` has been removed
- Enum `EventsObservabilityMode` has been removed
- Enum `UserAuthenticationMode` has been removed
- Struct `OwnCertificate` has been removed
- Struct `TransportAuthentication` has been removed
- Struct `TransportAuthenticationUpdate` has been removed
- Struct `UserAuthentication` has been removed
- Struct `UserAuthenticationUpdate` has been removed
- Field `TransportAuthentication`, `UserAuthentication` of struct `AssetEndpointProfileProperties` has been removed
- Field `TransportAuthentication`, `UserAuthentication` of struct `AssetEndpointProfileUpdateProperties` has been removed
- Field `AssetEndpointProfileURI`, `AssetType`, `DataPoints`, `DefaultDataPointsConfiguration` of struct `AssetProperties` has been removed
- Field `AssetType`, `DataPoints`, `DefaultDataPointsConfiguration` of struct `AssetUpdateProperties` has been removed
- Field `CapabilityID` of struct `DataPoint` has been removed
- Field `CapabilityID` of struct `Event` has been removed
- Field `PasswordReference`, `UsernameReference` of struct `UsernamePasswordCredentials` has been removed
- Field `PasswordReference`, `UsernameReference` of struct `UsernamePasswordCredentialsUpdate` has been removed
- Field `CertificateReference` of struct `X509Credentials` has been removed
- Field `CertificateReference` of struct `X509CredentialsUpdate` has been removed

### Features Added

- New value `ProvisioningStateDeleting` added to enum type `ProvisioningState`
- New enum type `AuthenticationMethod` with values `AuthenticationMethodAnonymous`, `AuthenticationMethodCertificate`, `AuthenticationMethodUsernamePassword`
- New enum type `DataPointObservabilityMode` with values `DataPointObservabilityModeCounter`, `DataPointObservabilityModeGauge`, `DataPointObservabilityModeHistogram`, `DataPointObservabilityModeLog`, `DataPointObservabilityModeNone`
- New enum type `EventObservabilityMode` with values `EventObservabilityModeLog`, `EventObservabilityModeNone`
- New enum type `Format` with values `FormatDelta10`, `FormatJSONSchemaDraft7`
- New enum type `SchemaType` with values `SchemaTypeMessageSchema`
- New enum type `SystemAssignedServiceIdentityType` with values `SystemAssignedServiceIdentityTypeNone`, `SystemAssignedServiceIdentityTypeSystemAssigned`
- New enum type `TopicRetainType` with values `TopicRetainTypeKeep`, `TopicRetainTypeNever`
- New function `NewBillingContainersClient(string, azcore.TokenCredential, *arm.ClientOptions) (*BillingContainersClient, error)`
- New function `*BillingContainersClient.Get(context.Context, string, *BillingContainersClientGetOptions) (BillingContainersClientGetResponse, error)`
- New function `*BillingContainersClient.NewListBySubscriptionPager(*BillingContainersClientListBySubscriptionOptions) *runtime.Pager[BillingContainersClientListBySubscriptionResponse]`
- New function `*ClientFactory.NewBillingContainersClient() *BillingContainersClient`
- New function `*ClientFactory.NewDiscoveredAssetEndpointProfilesClient() *DiscoveredAssetEndpointProfilesClient`
- New function `*ClientFactory.NewDiscoveredAssetsClient() *DiscoveredAssetsClient`
- New function `*ClientFactory.NewSchemaRegistriesClient() *SchemaRegistriesClient`
- New function `*ClientFactory.NewSchemaVersionsClient() *SchemaVersionsClient`
- New function `*ClientFactory.NewSchemasClient() *SchemasClient`
- New function `NewDiscoveredAssetEndpointProfilesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DiscoveredAssetEndpointProfilesClient, error)`
- New function `*DiscoveredAssetEndpointProfilesClient.BeginCreateOrReplace(context.Context, string, string, DiscoveredAssetEndpointProfile, *DiscoveredAssetEndpointProfilesClientBeginCreateOrReplaceOptions) (*runtime.Poller[DiscoveredAssetEndpointProfilesClientCreateOrReplaceResponse], error)`
- New function `*DiscoveredAssetEndpointProfilesClient.BeginDelete(context.Context, string, string, *DiscoveredAssetEndpointProfilesClientBeginDeleteOptions) (*runtime.Poller[DiscoveredAssetEndpointProfilesClientDeleteResponse], error)`
- New function `*DiscoveredAssetEndpointProfilesClient.Get(context.Context, string, string, *DiscoveredAssetEndpointProfilesClientGetOptions) (DiscoveredAssetEndpointProfilesClientGetResponse, error)`
- New function `*DiscoveredAssetEndpointProfilesClient.NewListByResourceGroupPager(string, *DiscoveredAssetEndpointProfilesClientListByResourceGroupOptions) *runtime.Pager[DiscoveredAssetEndpointProfilesClientListByResourceGroupResponse]`
- New function `*DiscoveredAssetEndpointProfilesClient.NewListBySubscriptionPager(*DiscoveredAssetEndpointProfilesClientListBySubscriptionOptions) *runtime.Pager[DiscoveredAssetEndpointProfilesClientListBySubscriptionResponse]`
- New function `*DiscoveredAssetEndpointProfilesClient.BeginUpdate(context.Context, string, string, DiscoveredAssetEndpointProfileUpdate, *DiscoveredAssetEndpointProfilesClientBeginUpdateOptions) (*runtime.Poller[DiscoveredAssetEndpointProfilesClientUpdateResponse], error)`
- New function `NewDiscoveredAssetsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DiscoveredAssetsClient, error)`
- New function `*DiscoveredAssetsClient.BeginCreateOrReplace(context.Context, string, string, DiscoveredAsset, *DiscoveredAssetsClientBeginCreateOrReplaceOptions) (*runtime.Poller[DiscoveredAssetsClientCreateOrReplaceResponse], error)`
- New function `*DiscoveredAssetsClient.BeginDelete(context.Context, string, string, *DiscoveredAssetsClientBeginDeleteOptions) (*runtime.Poller[DiscoveredAssetsClientDeleteResponse], error)`
- New function `*DiscoveredAssetsClient.Get(context.Context, string, string, *DiscoveredAssetsClientGetOptions) (DiscoveredAssetsClientGetResponse, error)`
- New function `*DiscoveredAssetsClient.NewListByResourceGroupPager(string, *DiscoveredAssetsClientListByResourceGroupOptions) *runtime.Pager[DiscoveredAssetsClientListByResourceGroupResponse]`
- New function `*DiscoveredAssetsClient.NewListBySubscriptionPager(*DiscoveredAssetsClientListBySubscriptionOptions) *runtime.Pager[DiscoveredAssetsClientListBySubscriptionResponse]`
- New function `*DiscoveredAssetsClient.BeginUpdate(context.Context, string, string, DiscoveredAssetUpdate, *DiscoveredAssetsClientBeginUpdateOptions) (*runtime.Poller[DiscoveredAssetsClientUpdateResponse], error)`
- New function `NewSchemaRegistriesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SchemaRegistriesClient, error)`
- New function `*SchemaRegistriesClient.BeginCreateOrReplace(context.Context, string, string, SchemaRegistry, *SchemaRegistriesClientBeginCreateOrReplaceOptions) (*runtime.Poller[SchemaRegistriesClientCreateOrReplaceResponse], error)`
- New function `*SchemaRegistriesClient.BeginDelete(context.Context, string, string, *SchemaRegistriesClientBeginDeleteOptions) (*runtime.Poller[SchemaRegistriesClientDeleteResponse], error)`
- New function `*SchemaRegistriesClient.Get(context.Context, string, string, *SchemaRegistriesClientGetOptions) (SchemaRegistriesClientGetResponse, error)`
- New function `*SchemaRegistriesClient.NewListByResourceGroupPager(string, *SchemaRegistriesClientListByResourceGroupOptions) *runtime.Pager[SchemaRegistriesClientListByResourceGroupResponse]`
- New function `*SchemaRegistriesClient.NewListBySubscriptionPager(*SchemaRegistriesClientListBySubscriptionOptions) *runtime.Pager[SchemaRegistriesClientListBySubscriptionResponse]`
- New function `*SchemaRegistriesClient.BeginUpdate(context.Context, string, string, SchemaRegistryUpdate, *SchemaRegistriesClientBeginUpdateOptions) (*runtime.Poller[SchemaRegistriesClientUpdateResponse], error)`
- New function `NewSchemaVersionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SchemaVersionsClient, error)`
- New function `*SchemaVersionsClient.CreateOrReplace(context.Context, string, string, string, string, SchemaVersion, *SchemaVersionsClientCreateOrReplaceOptions) (SchemaVersionsClientCreateOrReplaceResponse, error)`
- New function `*SchemaVersionsClient.Delete(context.Context, string, string, string, string, *SchemaVersionsClientDeleteOptions) (SchemaVersionsClientDeleteResponse, error)`
- New function `*SchemaVersionsClient.Get(context.Context, string, string, string, string, *SchemaVersionsClientGetOptions) (SchemaVersionsClientGetResponse, error)`
- New function `*SchemaVersionsClient.NewListBySchemaPager(string, string, string, *SchemaVersionsClientListBySchemaOptions) *runtime.Pager[SchemaVersionsClientListBySchemaResponse]`
- New function `NewSchemasClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SchemasClient, error)`
- New function `*SchemasClient.CreateOrReplace(context.Context, string, string, string, Schema, *SchemasClientCreateOrReplaceOptions) (SchemasClientCreateOrReplaceResponse, error)`
- New function `*SchemasClient.Delete(context.Context, string, string, string, *SchemasClientDeleteOptions) (SchemasClientDeleteResponse, error)`
- New function `*SchemasClient.Get(context.Context, string, string, string, *SchemasClientGetOptions) (SchemasClientGetResponse, error)`
- New function `*SchemasClient.NewListBySchemaRegistryPager(string, string, *SchemasClientListBySchemaRegistryOptions) *runtime.Pager[SchemasClientListBySchemaRegistryResponse]`
- New struct `AssetEndpointProfileStatus`
- New struct `AssetEndpointProfileStatusError`
- New struct `AssetStatusDataset`
- New struct `AssetStatusEvent`
- New struct `Authentication`
- New struct `AuthenticationUpdate`
- New struct `BillingContainer`
- New struct `BillingContainerListResult`
- New struct `BillingContainerProperties`
- New struct `Dataset`
- New struct `DiscoveredAsset`
- New struct `DiscoveredAssetEndpointProfile`
- New struct `DiscoveredAssetEndpointProfileListResult`
- New struct `DiscoveredAssetEndpointProfileProperties`
- New struct `DiscoveredAssetEndpointProfileUpdate`
- New struct `DiscoveredAssetEndpointProfileUpdateProperties`
- New struct `DiscoveredAssetListResult`
- New struct `DiscoveredAssetProperties`
- New struct `DiscoveredAssetUpdate`
- New struct `DiscoveredAssetUpdateProperties`
- New struct `DiscoveredDataPoint`
- New struct `DiscoveredDataset`
- New struct `DiscoveredEvent`
- New struct `MessageSchemaReference`
- New struct `Schema`
- New struct `SchemaListResult`
- New struct `SchemaProperties`
- New struct `SchemaRegistry`
- New struct `SchemaRegistryListResult`
- New struct `SchemaRegistryProperties`
- New struct `SchemaRegistryUpdate`
- New struct `SchemaRegistryUpdateProperties`
- New struct `SchemaVersion`
- New struct `SchemaVersionListResult`
- New struct `SchemaVersionProperties`
- New struct `SystemAssignedServiceIdentity`
- New struct `Topic`
- New struct `TopicUpdate`
- New field `Authentication`, `DiscoveredAssetEndpointProfileRef`, `EndpointProfileType`, `Status` in struct `AssetEndpointProfileProperties`
- New field `Authentication`, `EndpointProfileType` in struct `AssetEndpointProfileUpdateProperties`
- New field `AssetEndpointProfileRef`, `Datasets`, `DefaultDatasetsConfiguration`, `DefaultTopic`, `DiscoveredAssetRefs` in struct `AssetProperties`
- New field `Datasets`, `Events` in struct `AssetStatus`
- New field `Datasets`, `DefaultDatasetsConfiguration`, `DefaultTopic` in struct `AssetUpdateProperties`
- New field `Topic` in struct `Event`
- New field `ResourceID` in struct `OperationStatusResult`
- New field `PasswordSecretName`, `UsernameSecretName` in struct `UsernamePasswordCredentials`
- New field `PasswordSecretName`, `UsernameSecretName` in struct `UsernamePasswordCredentialsUpdate`
- New field `CertificateSecretName` in struct `X509Credentials`
- New field `CertificateSecretName` in struct `X509CredentialsUpdate`


## 0.1.0 (2024-04-26)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).