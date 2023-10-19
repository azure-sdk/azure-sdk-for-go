# Release History

## 1.2.0-beta.1 (2023-10-19)
### Features Added

- New enum type `ExportDetailCategory` with values `ExportDetailCategoryInformation`, `ExportDetailCategoryNotSpecified`, `ExportDetailCategoryRequiredStep`
- New enum type `ExportDetailCode` with values `ExportDetailCodeConnectionCloningUnsupported`, `ExportDetailCodeNotSpecified`
- New enum type `ValidationState` with values `ValidationStateFailed`, `ValidationStateNotSpecified`, `ValidationStateSucceeded`, `ValidationStateSucceededWithWarning`
- New enum type `WorkflowExportOptions` with values `WorkflowExportOptionsCloneConnections`, `WorkflowExportOptionsCloneConnectionsGenerateInfrastructureTemplates`, `WorkflowExportOptionsGenerateInfrastructureTemplates`
- New function `*ClientFactory.NewLocationsClient() *LocationsClient`
- New function `NewLocationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*LocationsClient, error)`
- New function `*LocationsClient.ValidateWorkflowExport(context.Context, string, WorkflowExportRequest, *LocationsClientValidateWorkflowExportOptions) (LocationsClientValidateWorkflowExportResponse, error)`
- New function `*LocationsClient.WorkflowExport(context.Context, string, WorkflowExportRequest, *LocationsClientWorkflowExportOptions) (LocationsClientWorkflowExportResponse, error)`
- New struct `ConnectionExportValidity`
- New struct `ExportDetail`
- New struct `ResourceExportValidity`
- New struct `WorkflowExportRequest`
- New struct `WorkflowExportResult`
- New struct `WorkflowExportValidity`
- New struct `WorkflowExportValidityResult`


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