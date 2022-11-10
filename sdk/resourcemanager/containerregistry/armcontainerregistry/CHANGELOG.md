# Release History

## 0.7.0 (2022-11-10)
### Features Added

- New const `CredentialNameCredential1`
- New const `CredentialHealthStatusHealthy`
- New const `CredentialHealthStatusUnhealthy`
- New type alias `CredentialName`
- New type alias `CredentialHealthStatus`
- New function `*CredentialSetsClient.BeginDelete(context.Context, string, string, string, *CredentialSetsClientBeginDeleteOptions) (*runtime.Poller[CredentialSetsClientDeleteResponse], error)`
- New function `*CacheRulesClient.GetAsync(context.Context, string, string, string, *CacheRulesClientGetAsyncOptions) (CacheRulesClientGetAsyncResponse, error)`
- New function `*CredentialSetsClient.BeginCreate(context.Context, string, string, string, CredentialSet, *CredentialSetsClientBeginCreateOptions) (*runtime.Poller[CredentialSetsClientCreateResponse], error)`
- New function `NewCacheRulesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CacheRulesClient, error)`
- New function `*CacheRulesClient.BeginDelete(context.Context, string, string, string, *CacheRulesClientBeginDeleteOptions) (*runtime.Poller[CacheRulesClientDeleteResponse], error)`
- New function `*CacheRulesClient.BeginUpdate(context.Context, string, string, string, CacheRuleUpdateParameters, *CacheRulesClientBeginUpdateOptions) (*runtime.Poller[CacheRulesClientUpdateResponse], error)`
- New function `*CacheRulesClient.BeginCreate(context.Context, string, string, string, CacheRule, *CacheRulesClientBeginCreateOptions) (*runtime.Poller[CacheRulesClientCreateResponse], error)`
- New function `PossibleCredentialHealthStatusValues() []CredentialHealthStatus`
- New function `*CredentialSetsClient.List(context.Context, string, string, *CredentialSetsClientListOptions) (CredentialSetsClientListResponse, error)`
- New function `*CredentialSetsClient.BeginUpdate(context.Context, string, string, string, CredentialSetUpdateParameters, *CredentialSetsClientBeginUpdateOptions) (*runtime.Poller[CredentialSetsClientUpdateResponse], error)`
- New function `*CredentialSetsClient.Get(context.Context, string, string, string, *CredentialSetsClientGetOptions) (CredentialSetsClientGetResponse, error)`
- New function `NewCredentialSetsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CredentialSetsClient, error)`
- New function `*CacheRulesClient.NewListAsyncPager(string, string, *CacheRulesClientListAsyncOptions) *runtime.Pager[CacheRulesClientListAsyncResponse]`
- New function `PossibleCredentialNameValues() []CredentialName`
- New struct `AuthCredential`
- New struct `CacheRule`
- New struct `CacheRuleProperties`
- New struct `CacheRuleUpdateParameters`
- New struct `CacheRuleUpdateProperties`
- New struct `CacheRulesClient`
- New struct `CacheRulesClientBeginCreateOptions`
- New struct `CacheRulesClientBeginDeleteOptions`
- New struct `CacheRulesClientBeginUpdateOptions`
- New struct `CacheRulesClientCreateResponse`
- New struct `CacheRulesClientDeleteResponse`
- New struct `CacheRulesClientGetAsyncOptions`
- New struct `CacheRulesClientGetAsyncResponse`
- New struct `CacheRulesClientListAsyncOptions`
- New struct `CacheRulesClientListAsyncResponse`
- New struct `CacheRulesClientUpdateResponse`
- New struct `CacheRulesListResult`
- New struct `CredentialHealth`
- New struct `CredentialSet`
- New struct `CredentialSetListResult`
- New struct `CredentialSetProperties`
- New struct `CredentialSetUpdateParameters`
- New struct `CredentialSetUpdateProperties`
- New struct `CredentialSetsClient`
- New struct `CredentialSetsClientBeginCreateOptions`
- New struct `CredentialSetsClientBeginDeleteOptions`
- New struct `CredentialSetsClientBeginUpdateOptions`
- New struct `CredentialSetsClientCreateResponse`
- New struct `CredentialSetsClientDeleteResponse`
- New struct `CredentialSetsClientGetOptions`
- New struct `CredentialSetsClientGetResponse`
- New struct `CredentialSetsClientListOptions`
- New struct `CredentialSetsClientListResponse`
- New struct `CredentialSetsClientUpdateResponse`


## 0.6.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).