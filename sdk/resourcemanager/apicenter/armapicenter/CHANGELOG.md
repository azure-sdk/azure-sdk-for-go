# Release History

## 1.0.0 (2023-10-30)
### Breaking Changes

- Function `NewClientFactory` parameter(s) have been changed from `(string, string, azcore.TokenCredential, *arm.ClientOptions)` to `(string, string, string, string, string, string, string, string, *string, azcore.TokenCredential, *arm.ClientOptions)`
- Function `NewServicesClient` parameter(s) have been changed from `(string, string, azcore.TokenCredential, *arm.ClientOptions)` to `(string, azcore.TokenCredential, *arm.ClientOptions)`
- Function `*ServicesClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, Service, *ServicesClientCreateOrUpdateOptions)` to `(context.Context, string, string, Service, *ServicesClientCreateOrUpdateOptions)`
- Function `*ServicesClient.Delete` parameter(s) have been changed from `(context.Context, string, *ServicesClientDeleteOptions)` to `(context.Context, string, string, *ServicesClientDeleteOptions)`
- Function `*ServicesClient.Get` parameter(s) have been changed from `(context.Context, string, *ServicesClientGetOptions)` to `(context.Context, string, string, *ServicesClientGetOptions)`
- Function `*ServicesClient.Update` parameter(s) have been changed from `(context.Context, string, ServiceUpdate, *ServicesClientUpdateOptions)` to `(context.Context, string, string, ServiceUpdate, *ServicesClientUpdateOptions)`

### Features Added

- New enum type `APIKind` with values `APIKindGraphql`, `APIKindGrpc`, `APIKindRest`, `APIKindSoap`, `APIKindWebhook`, `APIKindWebsocket`
- New enum type `APISpecExportResultFormat` with values `APISpecExportResultFormatInline`, `APISpecExportResultFormatLink`
- New enum type `APISpecImportSourceFormat` with values `APISpecImportSourceFormatInline`, `APISpecImportSourceFormatLink`
- New enum type `DeploymentState` with values `DeploymentStateActive`, `DeploymentStateInactive`
- New enum type `EnvironmentKind` with values `EnvironmentKindDevelopment`, `EnvironmentKindProduction`, `EnvironmentKindStaging`, `EnvironmentKindTesting`
- New enum type `EnvironmentServerType` with values `EnvironmentServerTypeAWSAPIGateway`, `EnvironmentServerTypeApigeeAPIManagement`, `EnvironmentServerTypeAzureAPIManagement`, `EnvironmentServerTypeAzureComputeService`, `EnvironmentServerTypeKongAPIGateway`, `EnvironmentServerTypeKubernetes`, `EnvironmentServerTypeMuleSoftAPIManagement`
- New enum type `LifecycleStage` with values `LifecycleStageDeprecated`, `LifecycleStageDesign`, `LifecycleStageDevelopment`, `LifecycleStagePreview`, `LifecycleStageProduction`, `LifecycleStageRetired`, `LifecycleStageTesting`
- New enum type `MetadataAssignmentEntity` with values `MetadataAssignmentEntityAPI`, `MetadataAssignmentEntityDeployment`, `MetadataAssignmentEntityEnvironment`
- New enum type `MetadataSchemaExportFormat` with values `MetadataSchemaExportFormatInline`, `MetadataSchemaExportFormatLink`
- New function `NewAPIDefinitionsClient(string, string, string, string, string, *string, azcore.TokenCredential, *arm.ClientOptions) (*APIDefinitionsClient, error)`
- New function `*APIDefinitionsClient.CreateOrUpdate(context.Context, string, string, APIDefinition, *APIDefinitionsClientCreateOrUpdateOptions) (APIDefinitionsClientCreateOrUpdateResponse, error)`
- New function `*APIDefinitionsClient.Delete(context.Context, string, string, *APIDefinitionsClientDeleteOptions) (APIDefinitionsClientDeleteResponse, error)`
- New function `*APIDefinitionsClient.ExportSpecification(context.Context, string, string, *APIDefinitionsClientExportSpecificationOptions) (APIDefinitionsClientExportSpecificationResponse, error)`
- New function `*APIDefinitionsClient.Get(context.Context, string, string, *APIDefinitionsClientGetOptions) (APIDefinitionsClientGetResponse, error)`
- New function `*APIDefinitionsClient.Head(context.Context, string, string, *APIDefinitionsClientHeadOptions) (APIDefinitionsClientHeadResponse, error)`
- New function `*APIDefinitionsClient.ImportSpecification(context.Context, string, string, APISpecImportRequest, *APIDefinitionsClientImportSpecificationOptions) (APIDefinitionsClientImportSpecificationResponse, error)`
- New function `*APIDefinitionsClient.NewListPager(string, string, *APIDefinitionsClientListOptions) *runtime.Pager[APIDefinitionsClientListResponse]`
- New function `NewAPIVersionsClient(string, string, string, string, *string, azcore.TokenCredential, *arm.ClientOptions) (*APIVersionsClient, error)`
- New function `*APIVersionsClient.CreateOrUpdate(context.Context, string, string, APIVersion, *APIVersionsClientCreateOrUpdateOptions) (APIVersionsClientCreateOrUpdateResponse, error)`
- New function `*APIVersionsClient.Delete(context.Context, string, string, *APIVersionsClientDeleteOptions) (APIVersionsClientDeleteResponse, error)`
- New function `*APIVersionsClient.Get(context.Context, string, string, *APIVersionsClientGetOptions) (APIVersionsClientGetResponse, error)`
- New function `*APIVersionsClient.Head(context.Context, string, string, *APIVersionsClientHeadOptions) (APIVersionsClientHeadResponse, error)`
- New function `*APIVersionsClient.NewListPager(string, string, *APIVersionsClientListOptions) *runtime.Pager[APIVersionsClientListResponse]`
- New function `NewApisClient(string, string, string, *string, azcore.TokenCredential, *arm.ClientOptions) (*ApisClient, error)`
- New function `*ApisClient.CreateOrUpdate(context.Context, string, string, API, *ApisClientCreateOrUpdateOptions) (ApisClientCreateOrUpdateResponse, error)`
- New function `*ApisClient.Delete(context.Context, string, string, *ApisClientDeleteOptions) (ApisClientDeleteResponse, error)`
- New function `*ApisClient.Get(context.Context, string, string, *ApisClientGetOptions) (ApisClientGetResponse, error)`
- New function `*ApisClient.Head(context.Context, string, string, *ApisClientHeadOptions) (ApisClientHeadResponse, error)`
- New function `*ApisClient.NewListPager(string, string, *ApisClientListOptions) *runtime.Pager[ApisClientListResponse]`
- New function `*ClientFactory.NewAPIDefinitionsClient() *APIDefinitionsClient`
- New function `*ClientFactory.NewAPIVersionsClient() *APIVersionsClient`
- New function `*ClientFactory.NewApisClient() *ApisClient`
- New function `*ClientFactory.NewDeploymentsClient() *DeploymentsClient`
- New function `*ClientFactory.NewEnvironmentsClient() *EnvironmentsClient`
- New function `*ClientFactory.NewMetadataSchemasClient() *MetadataSchemasClient`
- New function `*ClientFactory.NewWorkspacesClient() *WorkspacesClient`
- New function `NewDeploymentsClient(string, string, string, string, *string, azcore.TokenCredential, *arm.ClientOptions) (*DeploymentsClient, error)`
- New function `*DeploymentsClient.CreateOrUpdate(context.Context, string, string, Deployment, *DeploymentsClientCreateOrUpdateOptions) (DeploymentsClientCreateOrUpdateResponse, error)`
- New function `*DeploymentsClient.Delete(context.Context, string, string, *DeploymentsClientDeleteOptions) (DeploymentsClientDeleteResponse, error)`
- New function `*DeploymentsClient.Get(context.Context, string, string, *DeploymentsClientGetOptions) (DeploymentsClientGetResponse, error)`
- New function `*DeploymentsClient.Head(context.Context, string, string, *DeploymentsClientHeadOptions) (DeploymentsClientHeadResponse, error)`
- New function `*DeploymentsClient.NewListPager(string, string, *DeploymentsClientListOptions) *runtime.Pager[DeploymentsClientListResponse]`
- New function `NewEnvironmentsClient(string, string, string, *string, azcore.TokenCredential, *arm.ClientOptions) (*EnvironmentsClient, error)`
- New function `*EnvironmentsClient.CreateOrUpdate(context.Context, string, string, Environment, *EnvironmentsClientCreateOrUpdateOptions) (EnvironmentsClientCreateOrUpdateResponse, error)`
- New function `*EnvironmentsClient.Delete(context.Context, string, string, *EnvironmentsClientDeleteOptions) (EnvironmentsClientDeleteResponse, error)`
- New function `*EnvironmentsClient.Get(context.Context, string, string, *EnvironmentsClientGetOptions) (EnvironmentsClientGetResponse, error)`
- New function `*EnvironmentsClient.Head(context.Context, string, string, *EnvironmentsClientHeadOptions) (EnvironmentsClientHeadResponse, error)`
- New function `*EnvironmentsClient.NewListPager(string, string, *EnvironmentsClientListOptions) *runtime.Pager[EnvironmentsClientListResponse]`
- New function `NewMetadataSchemasClient(string, string, *string, azcore.TokenCredential, *arm.ClientOptions) (*MetadataSchemasClient, error)`
- New function `*MetadataSchemasClient.CreateOrUpdate(context.Context, string, string, MetadataSchema, *MetadataSchemasClientCreateOrUpdateOptions) (MetadataSchemasClientCreateOrUpdateResponse, error)`
- New function `*MetadataSchemasClient.Delete(context.Context, string, string, *MetadataSchemasClientDeleteOptions) (MetadataSchemasClientDeleteResponse, error)`
- New function `*MetadataSchemasClient.Get(context.Context, string, string, *MetadataSchemasClientGetOptions) (MetadataSchemasClientGetResponse, error)`
- New function `*MetadataSchemasClient.Head(context.Context, string, string, *MetadataSchemasClientHeadOptions) (MetadataSchemasClientHeadResponse, error)`
- New function `*MetadataSchemasClient.NewListPager(string, string, *MetadataSchemasClientListOptions) *runtime.Pager[MetadataSchemasClientListResponse]`
- New function `*ServicesClient.ExportMetadataSchema(context.Context, string, string, MetadataSchemaExportRequest, *ServicesClientExportMetadataSchemaOptions) (ServicesClientExportMetadataSchemaResponse, error)`
- New function `NewWorkspacesClient(string, string, *string, azcore.TokenCredential, *arm.ClientOptions) (*WorkspacesClient, error)`
- New function `*WorkspacesClient.CreateOrUpdate(context.Context, string, string, Workspace, *WorkspacesClientCreateOrUpdateOptions) (WorkspacesClientCreateOrUpdateResponse, error)`
- New function `*WorkspacesClient.Delete(context.Context, string, string, *WorkspacesClientDeleteOptions) (WorkspacesClientDeleteResponse, error)`
- New function `*WorkspacesClient.Get(context.Context, string, string, *WorkspacesClientGetOptions) (WorkspacesClientGetResponse, error)`
- New function `*WorkspacesClient.Head(context.Context, string, string, *WorkspacesClientHeadOptions) (WorkspacesClientHeadResponse, error)`
- New function `*WorkspacesClient.NewListPager(string, string, *WorkspacesClientListOptions) *runtime.Pager[WorkspacesClientListResponse]`
- New struct `API`
- New struct `APICollection`
- New struct `APIDefinition`
- New struct `APIDefinitionCollection`
- New struct `APIDefinitionProperties`
- New struct `APIDefinitionPropertiesSpecification`
- New struct `APIProperties`
- New struct `APISpecExportResult`
- New struct `APISpecImportRequest`
- New struct `APISpecImportRequestSpecification`
- New struct `APIVersion`
- New struct `APIVersionCollection`
- New struct `APIVersionProperties`
- New struct `Contact`
- New struct `Deployment`
- New struct `DeploymentCollection`
- New struct `DeploymentProperties`
- New struct `DeploymentServer`
- New struct `Environment`
- New struct `EnvironmentCollection`
- New struct `EnvironmentProperties`
- New struct `EnvironmentServer`
- New struct `ExternalDocumentation`
- New struct `License`
- New struct `MetadataAssignment`
- New struct `MetadataSchema`
- New struct `MetadataSchemaCollection`
- New struct `MetadataSchemaExportRequest`
- New struct `MetadataSchemaExportResult`
- New struct `MetadataSchemaProperties`
- New struct `Onboarding`
- New struct `TermsOfService`
- New struct `Workspace`
- New struct `WorkspaceCollection`
- New struct `WorkspaceProperties`


## 0.1.0 (2023-08-25)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apicenter/armapicenter` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).