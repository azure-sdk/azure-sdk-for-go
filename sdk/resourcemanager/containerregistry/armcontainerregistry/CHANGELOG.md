# Release History

## 0.7.0 (2022-12-12)
### Breaking Changes

- Type alias `ActivationStatus` has been removed
- Type alias `AuditLogStatus` has been removed
- Type alias `AzureADAuthenticationAsArmPolicyStatus` has been removed
- Type alias `CertificateType` has been removed
- Type alias `ConnectedRegistryMode` has been removed
- Type alias `ConnectionState` has been removed
- Type alias `LogLevel` has been removed
- Type alias `PipelineOptions` has been removed
- Type alias `PipelineRunSourceType` has been removed
- Type alias `PipelineRunTargetType` has been removed
- Type alias `PipelineSourceType` has been removed
- Type alias `TLSStatus` has been removed
- Function `NewConnectedRegistriesClient` has been removed
- Function `*ConnectedRegistriesClient.BeginCreate` has been removed
- Function `*ConnectedRegistriesClient.BeginDeactivate` has been removed
- Function `*ConnectedRegistriesClient.BeginDelete` has been removed
- Function `*ConnectedRegistriesClient.Get` has been removed
- Function `*ConnectedRegistriesClient.NewListPager` has been removed
- Function `*ConnectedRegistriesClient.BeginUpdate` has been removed
- Function `NewExportPipelinesClient` has been removed
- Function `*ExportPipelinesClient.BeginCreate` has been removed
- Function `*ExportPipelinesClient.BeginDelete` has been removed
- Function `*ExportPipelinesClient.Get` has been removed
- Function `*ExportPipelinesClient.NewListPager` has been removed
- Function `NewImportPipelinesClient` has been removed
- Function `*ImportPipelinesClient.BeginCreate` has been removed
- Function `*ImportPipelinesClient.BeginDelete` has been removed
- Function `*ImportPipelinesClient.Get` has been removed
- Function `*ImportPipelinesClient.NewListPager` has been removed
- Function `NewPipelineRunsClient` has been removed
- Function `*PipelineRunsClient.BeginCreate` has been removed
- Function `*PipelineRunsClient.BeginDelete` has been removed
- Function `*PipelineRunsClient.Get` has been removed
- Function `*PipelineRunsClient.NewListPager` has been removed
- Struct `ActivationProperties` has been removed
- Struct `AzureADAuthenticationAsArmPolicy` has been removed
- Struct `ConnectedRegistriesClient` has been removed
- Struct `ConnectedRegistriesClientCreateResponse` has been removed
- Struct `ConnectedRegistriesClientDeactivateResponse` has been removed
- Struct `ConnectedRegistriesClientDeleteResponse` has been removed
- Struct `ConnectedRegistriesClientListResponse` has been removed
- Struct `ConnectedRegistriesClientUpdateResponse` has been removed
- Struct `ConnectedRegistry` has been removed
- Struct `ConnectedRegistryListResult` has been removed
- Struct `ConnectedRegistryProperties` has been removed
- Struct `ConnectedRegistryUpdateParameters` has been removed
- Struct `ConnectedRegistryUpdateProperties` has been removed
- Struct `ExportPipeline` has been removed
- Struct `ExportPipelineListResult` has been removed
- Struct `ExportPipelineProperties` has been removed
- Struct `ExportPipelineTargetProperties` has been removed
- Struct `ExportPipelinesClient` has been removed
- Struct `ExportPipelinesClientCreateResponse` has been removed
- Struct `ExportPipelinesClientDeleteResponse` has been removed
- Struct `ExportPipelinesClientListResponse` has been removed
- Struct `ImportPipeline` has been removed
- Struct `ImportPipelineListResult` has been removed
- Struct `ImportPipelineProperties` has been removed
- Struct `ImportPipelineSourceProperties` has been removed
- Struct `ImportPipelinesClient` has been removed
- Struct `ImportPipelinesClientCreateResponse` has been removed
- Struct `ImportPipelinesClientDeleteResponse` has been removed
- Struct `ImportPipelinesClientListResponse` has been removed
- Struct `LoggingProperties` has been removed
- Struct `LoginServerProperties` has been removed
- Struct `ParentProperties` has been removed
- Struct `PipelineRun` has been removed
- Struct `PipelineRunListResult` has been removed
- Struct `PipelineRunProperties` has been removed
- Struct `PipelineRunRequest` has been removed
- Struct `PipelineRunResponse` has been removed
- Struct `PipelineRunSourceProperties` has been removed
- Struct `PipelineRunTargetProperties` has been removed
- Struct `PipelineRunsClient` has been removed
- Struct `PipelineRunsClientCreateResponse` has been removed
- Struct `PipelineRunsClientDeleteResponse` has been removed
- Struct `PipelineRunsClientListResponse` has been removed
- Struct `PipelineSourceTriggerDescriptor` has been removed
- Struct `PipelineSourceTriggerProperties` has been removed
- Struct `PipelineTriggerDescriptor` has been removed
- Struct `PipelineTriggerProperties` has been removed
- Struct `ProgressProperties` has been removed
- Struct `SoftDeletePolicy` has been removed
- Struct `StatusDetailProperties` has been removed
- Struct `SyncProperties` has been removed
- Struct `SyncUpdateProperties` has been removed
- Struct `TLSCertificateProperties` has been removed
- Struct `TLSProperties` has been removed
- Field `AzureADAuthenticationAsArmPolicy` of struct `Policies` has been removed
- Field `SoftDeletePolicy` of struct `Policies` has been removed
- Field `AnonymousPullEnabled` of struct `RegistryProperties` has been removed
- Field `AnonymousPullEnabled` of struct `RegistryPropertiesUpdateParameters` has been removed


## 0.6.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).