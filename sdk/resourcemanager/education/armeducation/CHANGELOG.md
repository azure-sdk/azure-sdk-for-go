# Release History

## 0.3.0 (2022-10-11)
### Features Added

- New function `*GrantsClient.NewListAllV2Pager(*GrantsClientListAllV2Options) *runtime.Pager[GrantsClientListAllV2Response]`
- New function `*OperationClient.Status(context.Context, string, *OperationClientStatusOptions) (OperationClientStatusResponse, error)`
- New function `NewGrantClient(azcore.TokenCredential, *arm.ClientOptions) (*GrantClient, error)`
- New function `NewOperationClient(azcore.TokenCredential, *arm.ClientOptions) (*OperationClient, error)`
- New function `*GrantsClient.NewListV2Pager(string, string, *GrantsClientListV2Options) *runtime.Pager[GrantsClientListV2Response]`
- New function `*GrantsClient.GetV2(context.Context, string, string, *GrantsClientGetV2Options) (GrantsClientGetV2Response, error)`
- New function `*GrantClient.Renewal(context.Context, string, string, *GrantClientRenewalOptions) (GrantClientRenewalResponse, error)`
- New struct `GrantClient`
- New struct `GrantClientRenewalOptions`
- New struct `GrantClientRenewalResponse`
- New struct `GrantDetailPropertiesV2`
- New struct `GrantDetailsV2`
- New struct `GrantListResponseV2`
- New struct `GrantsClientGetV2Options`
- New struct `GrantsClientGetV2Response`
- New struct `GrantsClientListAllV2Options`
- New struct `GrantsClientListAllV2Response`
- New struct `GrantsClientListV2Options`
- New struct `GrantsClientListV2Response`
- New struct `OperationClient`
- New struct `OperationClientStatusOptions`
- New struct `OperationClientStatusResponse`


## 0.2.0 (2022-07-04)
### Features Added

- New const `LabStatusPending`


## 0.1.0 (2022-05-24)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.1.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).