# Release History

## 1.1.0 (2022-10-14)
### Features Added

- New function `*SyncSetsClient.Update(context.Context, string, string, string, SyncSetUpdate, *SyncSetsClientUpdateOptions) (SyncSetsClientUpdateResponse, error)`
- New function `*MachinePoolsClient.Delete(context.Context, string, string, string, *MachinePoolsClientDeleteOptions) (MachinePoolsClientDeleteResponse, error)`
- New function `*SyncSetsClient.NewListPager(string, string, *SyncSetsClientListOptions) *runtime.Pager[SyncSetsClientListResponse]`
- New function `NewOpenShiftVersionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*OpenShiftVersionsClient, error)`
- New function `*SyncSetsClient.Delete(context.Context, string, string, string, *SyncSetsClientDeleteOptions) (SyncSetsClientDeleteResponse, error)`
- New function `*SecretsClient.Update(context.Context, string, string, string, SecretUpdate, *SecretsClientUpdateOptions) (SecretsClientUpdateResponse, error)`
- New function `*OpenShiftVersionsClient.NewListPager(string, *OpenShiftVersionsClientListOptions) *runtime.Pager[OpenShiftVersionsClientListResponse]`
- New function `NewSecretsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SecretsClient, error)`
- New function `NewSyncIdentityProvidersClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SyncIdentityProvidersClient, error)`
- New function `*MachinePoolsClient.CreateOrUpdate(context.Context, string, string, string, MachinePool, *MachinePoolsClientCreateOrUpdateOptions) (MachinePoolsClientCreateOrUpdateResponse, error)`
- New function `*SyncIdentityProvidersClient.NewListPager(string, string, *SyncIdentityProvidersClientListOptions) *runtime.Pager[SyncIdentityProvidersClientListResponse]`
- New function `*SyncIdentityProvidersClient.CreateOrUpdate(context.Context, string, string, string, SyncIdentityProvider, *SyncIdentityProvidersClientCreateOrUpdateOptions) (SyncIdentityProvidersClientCreateOrUpdateResponse, error)`
- New function `*SyncSetsClient.CreateOrUpdate(context.Context, string, string, string, SyncSet, *SyncSetsClientCreateOrUpdateOptions) (SyncSetsClientCreateOrUpdateResponse, error)`
- New function `*SecretsClient.Delete(context.Context, string, string, string, *SecretsClientDeleteOptions) (SecretsClientDeleteResponse, error)`
- New function `*MachinePoolsClient.Get(context.Context, string, string, string, *MachinePoolsClientGetOptions) (MachinePoolsClientGetResponse, error)`
- New function `*SyncIdentityProvidersClient.Update(context.Context, string, string, string, SyncIdentityProviderUpdate, *SyncIdentityProvidersClientUpdateOptions) (SyncIdentityProvidersClientUpdateResponse, error)`
- New function `*MachinePoolsClient.NewListPager(string, string, *MachinePoolsClientListOptions) *runtime.Pager[MachinePoolsClientListResponse]`
- New function `*SyncIdentityProvidersClient.Delete(context.Context, string, string, string, *SyncIdentityProvidersClientDeleteOptions) (SyncIdentityProvidersClientDeleteResponse, error)`
- New function `NewSyncSetsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SyncSetsClient, error)`
- New function `NewMachinePoolsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*MachinePoolsClient, error)`
- New function `*SecretsClient.CreateOrUpdate(context.Context, string, string, string, Secret, *SecretsClientCreateOrUpdateOptions) (SecretsClientCreateOrUpdateResponse, error)`
- New function `*SecretsClient.Get(context.Context, string, string, string, *SecretsClientGetOptions) (SecretsClientGetResponse, error)`
- New function `*MachinePoolsClient.Update(context.Context, string, string, string, MachinePoolUpdate, *MachinePoolsClientUpdateOptions) (MachinePoolsClientUpdateResponse, error)`
- New function `*SecretsClient.NewListPager(string, string, *SecretsClientListOptions) *runtime.Pager[SecretsClientListResponse]`
- New function `*SyncSetsClient.Get(context.Context, string, string, string, *SyncSetsClientGetOptions) (SyncSetsClientGetResponse, error)`
- New function `*SyncIdentityProvidersClient.Get(context.Context, string, string, string, *SyncIdentityProvidersClientGetOptions) (SyncIdentityProvidersClientGetResponse, error)`
- New struct `MachinePool`
- New struct `MachinePoolList`
- New struct `MachinePoolProperties`
- New struct `MachinePoolUpdate`
- New struct `MachinePoolsClient`
- New struct `MachinePoolsClientCreateOrUpdateOptions`
- New struct `MachinePoolsClientCreateOrUpdateResponse`
- New struct `MachinePoolsClientDeleteOptions`
- New struct `MachinePoolsClientDeleteResponse`
- New struct `MachinePoolsClientGetOptions`
- New struct `MachinePoolsClientGetResponse`
- New struct `MachinePoolsClientListOptions`
- New struct `MachinePoolsClientListResponse`
- New struct `MachinePoolsClientUpdateOptions`
- New struct `MachinePoolsClientUpdateResponse`
- New struct `OpenShiftVersion`
- New struct `OpenShiftVersionList`
- New struct `OpenShiftVersionProperties`
- New struct `OpenShiftVersionsClient`
- New struct `OpenShiftVersionsClientListOptions`
- New struct `OpenShiftVersionsClientListResponse`
- New struct `ProxyResource`
- New struct `Secret`
- New struct `SecretList`
- New struct `SecretProperties`
- New struct `SecretUpdate`
- New struct `SecretsClient`
- New struct `SecretsClientCreateOrUpdateOptions`
- New struct `SecretsClientCreateOrUpdateResponse`
- New struct `SecretsClientDeleteOptions`
- New struct `SecretsClientDeleteResponse`
- New struct `SecretsClientGetOptions`
- New struct `SecretsClientGetResponse`
- New struct `SecretsClientListOptions`
- New struct `SecretsClientListResponse`
- New struct `SecretsClientUpdateOptions`
- New struct `SecretsClientUpdateResponse`
- New struct `SyncIdentityProvider`
- New struct `SyncIdentityProviderList`
- New struct `SyncIdentityProviderProperties`
- New struct `SyncIdentityProviderUpdate`
- New struct `SyncIdentityProvidersClient`
- New struct `SyncIdentityProvidersClientCreateOrUpdateOptions`
- New struct `SyncIdentityProvidersClientCreateOrUpdateResponse`
- New struct `SyncIdentityProvidersClientDeleteOptions`
- New struct `SyncIdentityProvidersClientDeleteResponse`
- New struct `SyncIdentityProvidersClientGetOptions`
- New struct `SyncIdentityProvidersClientGetResponse`
- New struct `SyncIdentityProvidersClientListOptions`
- New struct `SyncIdentityProvidersClientListResponse`
- New struct `SyncIdentityProvidersClientUpdateOptions`
- New struct `SyncIdentityProvidersClientUpdateResponse`
- New struct `SyncSet`
- New struct `SyncSetList`
- New struct `SyncSetProperties`
- New struct `SyncSetUpdate`
- New struct `SyncSetsClient`
- New struct `SyncSetsClientCreateOrUpdateOptions`
- New struct `SyncSetsClientCreateOrUpdateResponse`
- New struct `SyncSetsClientDeleteOptions`
- New struct `SyncSetsClientDeleteResponse`
- New struct `SyncSetsClientGetOptions`
- New struct `SyncSetsClientGetResponse`
- New struct `SyncSetsClientListOptions`
- New struct `SyncSetsClientListResponse`
- New struct `SyncSetsClientUpdateOptions`
- New struct `SyncSetsClientUpdateResponse`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `TrackedResource`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).