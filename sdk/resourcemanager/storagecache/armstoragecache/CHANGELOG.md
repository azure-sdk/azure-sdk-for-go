# Release History

## 1.1.0 (2022-06-02)
### Features Added

- New const `HealthStateTypeStartFailed`
- New const `PrimingJobStateQueued`
- New const `PrimingJobStatePaused`
- New const `PrimingJobStateComplete`
- New const `PrimingJobStateRunning`
- New const `HealthStateTypeWaitingForKey`
- New const `HealthStateTypeUpgradeFailed`
- New function `*CachesClient.BeginStartPrimingJob(context.Context, string, string, *CachesClientBeginStartPrimingJobOptions) (*runtime.Poller[CachesClientStartPrimingJobResponse], error)`
- New function `*CachesClient.BeginPausePrimingJob(context.Context, string, string, *CachesClientBeginPausePrimingJobOptions) (*runtime.Poller[CachesClientPausePrimingJobResponse], error)`
- New function `*CachesClient.BeginSpaceAllocation(context.Context, string, string, *CachesClientBeginSpaceAllocationOptions) (*runtime.Poller[CachesClientSpaceAllocationResponse], error)`
- New function `*CachesClient.BeginResumePrimingJob(context.Context, string, string, *CachesClientBeginResumePrimingJobOptions) (*runtime.Poller[CachesClientResumePrimingJobResponse], error)`
- New function `*CachesClient.BeginStopPrimingJob(context.Context, string, string, *CachesClientBeginStopPrimingJobOptions) (*runtime.Poller[CachesClientStopPrimingJobResponse], error)`
- New function `PossiblePrimingJobStateValues() []PrimingJobState`
- New struct `CacheUpgradeSettings`
- New struct `CachesClientBeginPausePrimingJobOptions`
- New struct `CachesClientBeginResumePrimingJobOptions`
- New struct `CachesClientBeginSpaceAllocationOptions`
- New struct `CachesClientBeginStartPrimingJobOptions`
- New struct `CachesClientBeginStopPrimingJobOptions`
- New struct `CachesClientPausePrimingJobResponse`
- New struct `CachesClientResumePrimingJobResponse`
- New struct `CachesClientSpaceAllocationResponse`
- New struct `CachesClientStartPrimingJobResponse`
- New struct `CachesClientStopPrimingJobResponse`
- New struct `LogSpecification`
- New struct `PrimingJob`
- New struct `PrimingJobIDParameter`
- New struct `StorageTargetSpaceAllocation`
- New field `AllocationPercentage` in struct `StorageTargetProperties`
- New field `LogSpecifications` in struct `APIOperationPropertiesServiceSpecification`
- New field `UpgradeSettings` in struct `CacheProperties`
- New field `PrimingJobs` in struct `CacheProperties`
- New field `SpaceAllocation` in struct `CacheProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).