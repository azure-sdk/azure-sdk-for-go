# Release History

## 1.2.0 (2022-06-07)
### Features Added

- New const `RestServiceAuthenticationTypeOAuth2ClientCredential`
- New function `*SapODPSource.GetCopySource() *CopySource`
- New function `*SapODPLinkedService.GetLinkedService() *LinkedService`
- New function `*SapODPSource.GetTabularSource() *TabularSource`
- New function `*SapODPResourceDataset.GetDataset() *Dataset`
- New struct `SapODPLinkedService`
- New struct `SapODPLinkedServiceTypeProperties`
- New struct `SapODPResourceDataset`
- New struct `SapODPResourceDatasetTypeProperties`
- New struct `SapODPSource`
- New field `Scope` in struct `RestServiceLinkedServiceTypeProperties`
- New field `ClientID` in struct `RestServiceLinkedServiceTypeProperties`
- New field `TokenEndpoint` in struct `RestServiceLinkedServiceTypeProperties`
- New field `ClientSecret` in struct `RestServiceLinkedServiceTypeProperties`
- New field `Resource` in struct `RestServiceLinkedServiceTypeProperties`


## 1.1.0 (2022-05-30)
### Features Added

- New function `GlobalParameterResource.MarshalJSON() ([]byte, error)`
- New struct `GlobalParameterListResponse`
- New struct `GlobalParameterResource`
- New struct `GlobalParametersClientCreateOrUpdateOptions`
- New struct `GlobalParametersClientCreateOrUpdateResponse`
- New struct `GlobalParametersClientDeleteOptions`
- New struct `GlobalParametersClientDeleteResponse`
- New struct `GlobalParametersClientGetOptions`
- New struct `GlobalParametersClientGetResponse`
- New struct `GlobalParametersClientListByFactoryOptions`
- New struct `GlobalParametersClientListByFactoryResponse`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).