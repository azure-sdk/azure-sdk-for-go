# Release History

## 1.0.0 (2023-03-20)
### Breaking Changes

- Function `*ElasticSansClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, ElasticSan, *ElasticSansClientBeginCreateOptions)` to `(context.Context, string, string, CreateParameter, *ElasticSansClientBeginCreateOptions)`
- Function `*VolumesClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, string, string, Volume, *VolumesClientBeginCreateOptions)` to `(context.Context, string, string, string, string, VolumeCreateParameter, *VolumesClientBeginCreateOptions)`

### Features Added

- New function `NewSnapshotsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SnapshotsClient, error)`
- New function `*SnapshotsClient.BeginCreate(context.Context, string, string, string, string, SnapshotCreateParameter, *SnapshotsClientBeginCreateOptions) (*runtime.Poller[SnapshotsClientCreateResponse], error)`
- New function `*SnapshotsClient.BeginDelete(context.Context, string, string, string, string, *SnapshotsClientBeginDeleteOptions) (*runtime.Poller[SnapshotsClientDeleteResponse], error)`
- New function `*SnapshotsClient.Get(context.Context, string, string, string, string, *SnapshotsClientGetOptions) (SnapshotsClientGetResponse, error)`
- New function `*SnapshotsClient.NewListByVolumeGroupPager(string, string, string, *SnapshotsClientListByVolumeGroupOptions) *runtime.Pager[SnapshotsClientListByVolumeGroupResponse]`
- New function `*SnapshotsClient.Update(context.Context, string, string, string, string, SnapshotUpdate, *SnapshotsClientUpdateOptions) (SnapshotsClientUpdateResponse, error)`
- New struct `CreateParameter`
- New struct `Snapshot`
- New struct `SnapshotCreateParameter`
- New struct `SnapshotCreationData`
- New struct `SnapshotList`
- New struct `SnapshotProperties`
- New struct `SnapshotUpdate`
- New struct `SnapshotsClient`
- New struct `VolumeCreateParameter`


## 0.1.0 (2022-10-21)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).