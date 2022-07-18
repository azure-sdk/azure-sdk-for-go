# Release History

## 1.1.0-beta.1 (2022-07-18)
### Features Added

- New const `DtcNameCurrent`
- New function `*ManagedInstanceDtcsClient.NewListByManagedInstancePager(string, string, *ManagedInstanceDtcsClientListByManagedInstanceOptions) *runtime.Pager[ManagedInstanceDtcsClientListByManagedInstanceResponse]`
- New function `*ManagedInstanceDtcsClient.Get(context.Context, string, string, DtcName, *ManagedInstanceDtcsClientGetOptions) (ManagedInstanceDtcsClientGetResponse, error)`
- New function `NewManagedInstanceDtcsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagedInstanceDtcsClient, error)`
- New function `*ManagedInstanceDtcsClient.BeginCreateOrUpdate(context.Context, string, string, DtcName, ManagedInstanceDtc, *ManagedInstanceDtcsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedInstanceDtcsClientCreateOrUpdateResponse], error)`
- New function `PossibleDtcNameValues() []DtcName`
- New struct `ManagedInstanceDtc`
- New struct `ManagedInstanceDtcListResult`
- New struct `ManagedInstanceDtcProperties`
- New struct `ManagedInstanceDtcSecuritySettings`
- New struct `ManagedInstanceDtcTransactionManagerCommunicationSettings`
- New struct `ManagedInstanceDtcsClient`
- New struct `ManagedInstanceDtcsClientBeginCreateOrUpdateOptions`
- New struct `ManagedInstanceDtcsClientCreateOrUpdateResponse`
- New struct `ManagedInstanceDtcsClientGetOptions`
- New struct `ManagedInstanceDtcsClientGetResponse`
- New struct `ManagedInstanceDtcsClientListByManagedInstanceOptions`
- New struct `ManagedInstanceDtcsClientListByManagedInstanceResponse`


## 1.0.0 (2022-06-02)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).