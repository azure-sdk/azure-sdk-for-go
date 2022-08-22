# Release History

## 1.1.0 (2022-08-22)
### Features Added

- New const `RunLocationUsgovvirginia`
- New const `JobDefinitionStateQueueing`
- New const `JobStatusFailedRetried`
- New const `JobStatusCancelledSkipped`
- New const `JobStatusFailedSkipped`
- New const `RunLocationUsgovtexas`
- New const `JobStatusCancelledRetried`
- New const `JobDefinitionStateWaiting`
- New const `JobDefinitionStateComplete`
- New const `JobDefinitionStateCancelled`
- New const `JobDefinitionStatePaused`
- New const `JobStatusWaiting`
- New const `JobDefinitionStateRunning`
- New const `JobDefinitionStateCreated`
- New const `JobDefinitionStateCancelling`
- New type alias `JobDefinitionState`
- New function `*JobDefinitionsClient.SkipJob(context.Context, string, string, string, string, *JobDefinitionsClientSkipJobOptions) (JobDefinitionsClientSkipJobResponse, error)`
- New function `*JobDefinitionsClient.RetryJob(context.Context, string, string, string, string, *JobDefinitionsClientRetryJobOptions) (JobDefinitionsClientRetryJobResponse, error)`
- New function `PossibleJobDefinitionStateValues() []JobDefinitionState`
- New function `*JobDefinitionsClient.Cancel(context.Context, string, string, string, string, *JobDefinitionsClientCancelOptions) (JobDefinitionsClientCancelResponse, error)`
- New struct `CloudError`
- New struct `CloudErrorBody`
- New struct `JobDefinitionsClientCancelOptions`
- New struct `JobDefinitionsClientCancelResponse`
- New struct `JobDefinitionsClientRetryJobOptions`
- New struct `JobDefinitionsClientRetryJobResponse`
- New struct `JobDefinitionsClientSkipJobOptions`
- New struct `JobDefinitionsClientSkipJobResponse`
- New field `JobDefinitionSchemaVersion` in struct `JobDefinitionProperties`
- New field `JobDefinitionState` in struct `JobDefinitionProperties`
- New field `ActionType` in struct `AvailableProviderOperation`
- New field `IsDataAction` in struct `AvailableProviderOperation`
- New field `DirectoryParserEndTime` in struct `JobProperties`
- New field `PostEstimationEndTime` in struct `JobProperties`
- New field `PostCopyEndTime` in struct `JobProperties`
- New field `CleanupComputeStartTime` in struct `JobProperties`
- New field `BackupEndTime` in struct `JobProperties`
- New field `CopyStartTime` in struct `JobProperties`
- New field `EstimationEndTime` in struct `JobProperties`
- New field `PurgeDetectionEndTime` in struct `JobProperties`
- New field `CopyEndTime` in struct `JobProperties`
- New field `PostCopyStartTime` in struct `JobProperties`
- New field `CopyErrors` in struct `JobProperties`
- New field `TotalItemsToDelete` in struct `JobProperties`
- New field `EstimationStartTime` in struct `JobProperties`
- New field `DirectoryParserStartTime` in struct `JobProperties`
- New field `CleanupComputeEndTime` in struct `JobProperties`
- New field `PostEstimationStartTime` in struct `JobProperties`
- New field `BackupStartTime` in struct `JobProperties`
- New field `UnsupportedFiles` in struct `JobProperties`
- New field `CopyConfigureComputeEndTime` in struct `JobProperties`
- New field `EstimationConfigureComputeEndTime` in struct `JobProperties`
- New field `EstimationConfigureComputeStartTime` in struct `JobProperties`
- New field `PurgeDetectionStartTime` in struct `JobProperties`
- New field `CopyConfigureComputeStartTime` in struct `JobProperties`
- New field `FilesDeletedCounter` in struct `JobProperties`
- New field `FilesCopiedCounter` in struct `JobProperties`
- New field `BytesCopied` in struct `JobProperties`
- New field `ItemsDeleted` in struct `JobProperties`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).