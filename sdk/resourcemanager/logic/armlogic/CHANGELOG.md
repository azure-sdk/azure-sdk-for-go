# Release History

## 2.0.0 (2025-03-11)
### Breaking Changes

- Function `*IntegrationServiceEnvironmentManagedApisClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, string, *IntegrationServiceEnvironmentManagedApisClientBeginDeleteOptions)` to `(context.Context, string, string, *IntegrationServiceEnvironmentManagedApisClientBeginDeleteOptions)`
- Function `*IntegrationServiceEnvironmentManagedApisClient.BeginPut` parameter(s) have been changed from `(context.Context, string, string, string, IntegrationServiceEnvironmentManagedAPI, *IntegrationServiceEnvironmentManagedApisClientBeginPutOptions)` to `(context.Context, string, string, IntegrationServiceEnvironmentManagedAPI, *IntegrationServiceEnvironmentManagedApisClientBeginPutOptions)`
- Function `*IntegrationServiceEnvironmentManagedApisClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *IntegrationServiceEnvironmentManagedApisClientGetOptions)` to `(context.Context, string, string, *IntegrationServiceEnvironmentManagedApisClientGetOptions)`
- Function `*IntegrationServiceEnvironmentManagedApisClient.NewListPager` parameter(s) have been changed from `(string, string, *IntegrationServiceEnvironmentManagedApisClientListOptions)` to `(string, *IntegrationServiceEnvironmentManagedApisClientListOptions)`
- Function `*IntegrationServiceEnvironmentNetworkHealthClient.Get` parameter(s) have been changed from `(context.Context, string, string, *IntegrationServiceEnvironmentNetworkHealthClientGetOptions)` to `(context.Context, string, *IntegrationServiceEnvironmentNetworkHealthClientGetOptions)`
- Function `*IntegrationServiceEnvironmentSKUsClient.NewListPager` parameter(s) have been changed from `(string, string, *IntegrationServiceEnvironmentSKUsClientListOptions)` to `(string, *IntegrationServiceEnvironmentSKUsClientListOptions)`
- Function `*IntegrationServiceEnvironmentsClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, IntegrationServiceEnvironment, *IntegrationServiceEnvironmentsClientBeginCreateOrUpdateOptions)` to `(context.Context, string, IntegrationServiceEnvironment, *IntegrationServiceEnvironmentsClientBeginCreateOrUpdateOptions)`
- Function `*IntegrationServiceEnvironmentsClient.Delete` parameter(s) have been changed from `(context.Context, string, string, *IntegrationServiceEnvironmentsClientDeleteOptions)` to `(context.Context, string, *IntegrationServiceEnvironmentsClientDeleteOptions)`
- Function `*IntegrationServiceEnvironmentsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *IntegrationServiceEnvironmentsClientGetOptions)` to `(context.Context, string, *IntegrationServiceEnvironmentsClientGetOptions)`
- Function `*IntegrationServiceEnvironmentsClient.Restart` parameter(s) have been changed from `(context.Context, string, string, *IntegrationServiceEnvironmentsClientRestartOptions)` to `(context.Context, string, *IntegrationServiceEnvironmentsClientRestartOptions)`
- Function `*WorkflowRunsClient.NewListPager` parameter(s) have been changed from `(string, string, *WorkflowRunsClientListOptions)` to `(string, string, string, *WorkflowRunsClientListOptions)`
- Function `*WorkflowsClient.Update` parameter(s) have been changed from `(context.Context, string, string, *WorkflowsClientUpdateOptions)` to `(context.Context, string, string, any, *WorkflowsClientUpdateOptions)`
- Type of `APIReference.Swagger` has been changed from `any` to `map[string]any`
- Type of `APIResourceProperties.ConnectionParameters` has been changed from `map[string]any` to `map[string]map[string]any`
- Type of `AgreementContent.AS2` has been changed from `*AS2AgreementContent` to `*As2AgreementContent`
- Type of `ContentLink.Metadata` has been changed from `any` to `map[string]any`
- Type of `ErrorResponse.Error` has been changed from `*ErrorProperties` to `*ErrorDetail`
- Type of `IntegrationAccountAgreementProperties.Metadata` has been changed from `any` to `map[string]any`
- Type of `IntegrationAccountCertificateProperties.Metadata` has been changed from `any` to `map[string]any`
- Type of `IntegrationAccountMapProperties.Metadata` has been changed from `any` to `map[string]any`
- Type of `IntegrationAccountPartnerProperties.Metadata` has been changed from `any` to `map[string]any`
- Type of `IntegrationAccountSchemaProperties.Metadata` has been changed from `any` to `map[string]any`
- Type of `IntegrationAccountSessionProperties.Content` has been changed from `any` to `map[string]any`
- Type of `IntegrationServiceEnvironmentManagedAPIProperties.ConnectionParameters` has been changed from `map[string]any` to `map[string]map[string]any`
- Type of `Operation.Origin` has been changed from `*string` to `*Origin`
- Type of `OperationResult.Inputs` has been changed from `any` to `map[string]any`
- Type of `OperationResult.Outputs` has been changed from `any` to `map[string]any`
- Type of `OperationResult.TrackedProperties` has been changed from `any` to `map[string]any`
- Type of `Request.Headers` has been changed from `any` to `map[string]any`
- Type of `Response.Headers` has been changed from `any` to `map[string]any`
- Type of `SwaggerCustomDynamicSchema.Parameters` has been changed from `map[string]any` to `map[string]map[string]any`
- Type of `SwaggerCustomDynamicTreeParameter.Value` has been changed from `any` to `map[string]any`
- Type of `SwaggerExternalDocumentation.Extensions` has been changed from `map[string]any` to `map[string]map[string]any`
- Type of `SwaggerSchema.AdditionalProperties` has been changed from `any` to `map[string]any`
- Type of `SwaggerSchema.Example` has been changed from `any` to `map[string]any`
- Type of `SwaggerXML.Extensions` has been changed from `map[string]any` to `map[string]map[string]any`
- Type of `TrackingEvent.Record` has been changed from `any` to `map[string]any`
- Type of `WorkflowOutputParameter.Error` has been changed from `any` to `map[string]any`
- Type of `WorkflowOutputParameter.Metadata` has been changed from `any` to `map[string]any`
- Type of `WorkflowOutputParameter.Value` has been changed from `any` to `map[string]any`
- Type of `WorkflowParameter.Metadata` has been changed from `any` to `map[string]any`
- Type of `WorkflowParameter.Value` has been changed from `any` to `map[string]any`
- Type of `WorkflowProperties.Definition` has been changed from `any` to `map[string]any`
- Type of `WorkflowRunActionProperties.Error` has been changed from `any` to `map[string]any`
- Type of `WorkflowRunActionProperties.TrackedProperties` has been changed from `any` to `map[string]any`
- Type of `WorkflowRunActionRepetitionProperties.Inputs` has been changed from `any` to `map[string]any`
- Type of `WorkflowRunActionRepetitionProperties.Outputs` has been changed from `any` to `map[string]any`
- Type of `WorkflowRunActionRepetitionProperties.TrackedProperties` has been changed from `any` to `map[string]any`
- Type of `WorkflowRunProperties.Error` has been changed from `any` to `map[string]any`
- Type of `WorkflowRunTrigger.Error` has been changed from `any` to `map[string]any`
- Type of `WorkflowRunTrigger.Inputs` has been changed from `any` to `map[string]any`
- Type of `WorkflowRunTrigger.Outputs` has been changed from `any` to `map[string]any`
- Type of `WorkflowRunTrigger.TrackedProperties` has been changed from `any` to `map[string]any`
- Type of `WorkflowTriggerHistoryProperties.Error` has been changed from `any` to `map[string]any`
- Type of `WorkflowVersionProperties.Definition` has been changed from `any` to `map[string]any`
- `TrackingRecordTypeAS2MDN`, `TrackingRecordTypeAS2Message` from enum `TrackingRecordType` has been removed
- Enum `AzureAsyncOperationState` has been removed
- Enum `ErrorResponseCode` has been removed
- Enum `IntegrationServiceEnvironmentNetworkDependencyHealthState` has been removed
- Function `*ClientFactory.NewIntegrationServiceEnvironmentManagedAPIOperationsClient` has been removed
- Function `NewIntegrationServiceEnvironmentManagedAPIOperationsClient` has been removed
- Function `*IntegrationServiceEnvironmentManagedAPIOperationsClient.NewListPager` has been removed
- Function `*WorkflowRunActionRepetitionsClient.NewListPager` has been removed
- Operation `*IntegrationServiceEnvironmentsClient.BeginUpdate` has been changed to non-LRO, use `*IntegrationServiceEnvironmentsClient.Update` instead.
- Operation `*WorkflowTriggerHistoriesClient.Resubmit` has been changed to LRO, use `*WorkflowTriggerHistoriesClient.BeginResubmit` instead.
- Operation `*WorkflowTriggersClient.Run` has been changed to LRO, use `*WorkflowTriggersClient.BeginRun` instead.
- Operation `*IntegrationAccountAssembliesClient.NewListPager` does not support pagination anymore, use `*IntegrationAccountAssembliesClient.List` instead.
- Operation `*IntegrationAccountBatchConfigurationsClient.NewListPager` does not support pagination anymore, use `*IntegrationAccountBatchConfigurationsClient.List` instead.
- Operation `*IntegrationAccountsClient.NewListKeyVaultKeysPager` does not support pagination anymore, use `*IntegrationAccountsClient.ListKeyVaultKeys` instead.
- Operation `*WorkflowRunActionRepetitionsClient.NewListExpressionTracesPager` does not support pagination anymore, use `*WorkflowRunActionRepetitionsClient.ListExpressionTraces` instead.
- Operation `*WorkflowRunActionsClient.NewListExpressionTracesPager` does not support pagination anymore, use `*WorkflowRunActionsClient.ListExpressionTraces` instead.
- Struct `AS2AcknowledgementConnectionSettings` has been removed
- Struct `AS2AgreementContent` has been removed
- Struct `AS2EnvelopeSettings` has been removed
- Struct `AS2ErrorSettings` has been removed
- Struct `AS2MdnSettings` has been removed
- Struct `AS2MessageConnectionSettings` has been removed
- Struct `AS2OneWayAgreement` has been removed
- Struct `AS2ProtocolSettings` has been removed
- Struct `AS2SecuritySettings` has been removed
- Struct `AS2ValidationSettings` has been removed
- Struct `ErrorProperties` has been removed
- Struct `ExtendedErrorInfo` has been removed
- Struct `IntegrationAccountAgreementFilter` has been removed
- Struct `IntegrationAccountMapFilter` has been removed
- Struct `IntegrationAccountPartnerFilter` has been removed
- Struct `IntegrationAccountSchemaFilter` has been removed
- Struct `IntegrationAccountSessionFilter` has been removed
- Struct `IntegrationServiceEnvironmentNetworkDependencyHealth` has been removed
- Struct `IntegrationServiceEnvironmentSubnetNetworkHealth` has been removed
- Struct `ManagedAPI` has been removed
- Struct `ManagedAPIListResult` has been removed
- Struct `SubResource` has been removed
- Struct `WorkflowFilter` has been removed
- Struct `WorkflowRunActionFilter` has been removed
- Struct `WorkflowRunFilter` has been removed
- Struct `WorkflowTriggerFilter` has been removed
- Struct `WorkflowTriggerHistoryFilter` has been removed
- Field `Location`, `Tags` of struct `APIOperation` has been removed
- Field `Name` of struct `APIResourceProperties` has been removed
- Field `ValidateEDITypes`, `ValidateXSDTypes` of struct `EdifactValidationOverride` has been removed
- Field `ValidateEDITypes`, `ValidateXSDTypes` of struct `EdifactValidationSettings` has been removed
- Field `Name` of struct `IntegrationServiceEnvironmentManagedAPIProperties` has been removed
- Field `Value` of struct `IntegrationServiceEnvironmentNetworkHealthClientGetResponse` has been removed
- Field `Properties` of struct `Operation` has been removed
- Field `Location`, `Tags` of struct `Resource` has been removed
- Field `Interface` of struct `WorkflowsClientGenerateUpgradedDefinitionResponse` has been removed
- Field `Interface` of struct `WorkflowsClientListSwaggerResponse` has been removed
- Field `ValidateEDITypes`, `ValidateXSDTypes` of struct `X12ValidationOverride` has been removed
- Field `ValidateEDITypes`, `ValidateXSDTypes` of struct `X12ValidationSettings` has been removed

### Features Added

- New value `ManagedServiceIdentityTypeSystemAssignedUserAssigned` added to enum type `ManagedServiceIdentityType`
- New value `TrackingRecordTypeAs2MDN`, `TrackingRecordTypeAs2Message` added to enum type `TrackingRecordType`
- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New function `NewAPIOperationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*APIOperationsClient, error)`
- New function `*APIOperationsClient.NewListPager(string, string, *APIOperationsClientListOptions) *runtime.Pager[APIOperationsClientListResponse]`
- New function `*ClientFactory.NewAPIOperationsClient() *APIOperationsClient`
- New struct `As2AcknowledgementConnectionSettings`
- New struct `As2AgreementContent`
- New struct `As2EnvelopeSettings`
- New struct `As2ErrorSettings`
- New struct `As2MdnSettings`
- New struct `As2MessageConnectionSettings`
- New struct `As2OneWayAgreement`
- New struct `As2ProtocolSettings`
- New struct `As2SecuritySettings`
- New struct `As2ValidationSettings`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `IntegrationServiceEnvironmentNetworkHealth`
- New struct `ProxyResource`
- New struct `SystemData`
- New struct `TrackedResource`
- New field `SystemData` in struct `APIOperation`
- New field `SystemData` in struct `AssemblyDefinition`
- New field `SystemData` in struct `BatchConfiguration`
- New field `ValidateEdiTypes`, `ValidateXsdTypes` in struct `EdifactValidationOverride`
- New field `ValidateEdiTypes`, `ValidateXsdTypes` in struct `EdifactValidationSettings`
- New field `SystemData` in struct `IntegrationAccount`
- New field `SystemData` in struct `IntegrationAccountAgreement`
- New field `SystemData` in struct `IntegrationAccountCertificate`
- New field `SystemData` in struct `IntegrationAccountMap`
- New field `SystemData` in struct `IntegrationAccountPartner`
- New field `SystemData` in struct `IntegrationAccountSchema`
- New field `SystemData` in struct `IntegrationAccountSession`
- New field `SystemData` in struct `IntegrationServiceEnvironment`
- New field `SystemData` in struct `IntegrationServiceEnvironmentManagedAPI`
- New anonymous field `IntegrationServiceEnvironmentNetworkHealth` in struct `IntegrationServiceEnvironmentNetworkHealthClientGetResponse`
- New field `ActionType`, `IsDataAction` in struct `Operation`
- New field `SystemData` in struct `RequestHistory`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `Workflow`
- New field `SystemData` in struct `WorkflowRun`
- New field `SystemData` in struct `WorkflowRunAction`
- New field `SystemData` in struct `WorkflowRunActionRepetitionDefinition`
- New field `SystemData` in struct `WorkflowTrigger`
- New field `SystemData` in struct `WorkflowTriggerHistory`
- New field `SystemData` in struct `WorkflowVersion`
- New field `Value` in struct `WorkflowsClientGenerateUpgradedDefinitionResponse`
- New field `Value` in struct `WorkflowsClientListSwaggerResponse`
- New field `ValidateEdiTypes`, `ValidateXsdTypes` in struct `X12ValidationOverride`
- New field `ValidateEdiTypes`, `ValidateXsdTypes` in struct `X12ValidationSettings`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).