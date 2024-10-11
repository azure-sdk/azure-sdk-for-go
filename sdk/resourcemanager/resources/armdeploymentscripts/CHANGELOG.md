# Release History

## 3.0.0 (2024-10-11)
### Breaking Changes

- Enum `CleanupOptions` has been removed
- Enum `ManagedServiceIdentityType` has been removed
- Enum `ScriptProvisioningState` has been removed
- Enum `ScriptType` has been removed
- Function `*AzureCliScript.GetDeploymentScript` has been removed
- Function `*AzurePowerShellScript.GetDeploymentScript` has been removed
- Function `NewClient` has been removed
- Function `*Client.BeginCreate` has been removed
- Function `*Client.Delete` has been removed
- Function `*Client.Get` has been removed
- Function `*Client.GetLogs` has been removed
- Function `*Client.GetLogsDefault` has been removed
- Function `*Client.NewListByResourceGroupPager` has been removed
- Function `*Client.NewListBySubscriptionPager` has been removed
- Function `*Client.Update` has been removed
- Function `*ClientFactory.NewClient` has been removed
- Function `*DeploymentScript.GetDeploymentScript` has been removed
- Struct `AzureCliScript` has been removed
- Struct `AzureCliScriptProperties` has been removed
- Struct `AzurePowerShellScript` has been removed
- Struct `AzurePowerShellScriptProperties` has been removed
- Struct `AzureResourceBase` has been removed
- Struct `ContainerConfiguration` has been removed
- Struct `DeploymentScriptListResult` has been removed
- Struct `DeploymentScriptUpdateParameter` has been removed
- Struct `EnvironmentVariable` has been removed
- Struct `Error` has been removed
- Struct `LogProperties` has been removed
- Struct `ManagedServiceIdentity` has been removed
- Struct `ScriptLog` has been removed
- Struct `ScriptLogsList` has been removed
- Struct `ScriptStatus` has been removed
- Struct `StorageAccountConfiguration` has been removed
- Struct `UserAssignedIdentity` has been removed

### Features Added

- New enum type `AssignmentType` with values `AssignmentTypeCustom`, `AssignmentTypeNotSpecified`, `AssignmentTypeSystem`, `AssignmentTypeSystemHidden`
- New enum type `EnforcementMode` with values `EnforcementModeDefault`, `EnforcementModeDoNotEnforce`
- New enum type `OverrideKind` with values `OverrideKindDefinitionVersion`, `OverrideKindPolicyEffect`
- New enum type `ParameterType` with values `ParameterTypeArray`, `ParameterTypeBoolean`, `ParameterTypeDateTime`, `ParameterTypeFloat`, `ParameterTypeInteger`, `ParameterTypeObject`, `ParameterTypeString`
- New enum type `PolicyType` with values `PolicyTypeBuiltIn`, `PolicyTypeCustom`, `PolicyTypeNotSpecified`, `PolicyTypeStatic`
- New enum type `ResourceIdentityType` with values `ResourceIdentityTypeNone`, `ResourceIdentityTypeSystemAssigned`, `ResourceIdentityTypeUserAssigned`
- New enum type `SelectorKind` with values `SelectorKindPolicyDefinitionReferenceID`, `SelectorKindResourceLocation`, `SelectorKindResourceType`, `SelectorKindResourceWithoutLocation`
- New function `*ClientFactory.NewPolicyAssignmentsClient() *PolicyAssignmentsClient`
- New function `*ClientFactory.NewPolicyDefinitionVersionsClient() *PolicyDefinitionVersionsClient`
- New function `*ClientFactory.NewPolicyDefinitionsClient() *PolicyDefinitionsClient`
- New function `*ClientFactory.NewPolicySetDefinitionVersionsClient() *PolicySetDefinitionVersionsClient`
- New function `*ClientFactory.NewPolicySetDefinitionsClient() *PolicySetDefinitionsClient`
- New function `NewPolicyAssignmentsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PolicyAssignmentsClient, error)`
- New function `*PolicyAssignmentsClient.Create(context.Context, string, string, PolicyAssignment, *PolicyAssignmentsClientCreateOptions) (PolicyAssignmentsClientCreateResponse, error)`
- New function `*PolicyAssignmentsClient.CreateByID(context.Context, string, PolicyAssignment, *PolicyAssignmentsClientCreateByIDOptions) (PolicyAssignmentsClientCreateByIDResponse, error)`
- New function `*PolicyAssignmentsClient.Delete(context.Context, string, string, *PolicyAssignmentsClientDeleteOptions) (PolicyAssignmentsClientDeleteResponse, error)`
- New function `*PolicyAssignmentsClient.DeleteByID(context.Context, string, *PolicyAssignmentsClientDeleteByIDOptions) (PolicyAssignmentsClientDeleteByIDResponse, error)`
- New function `*PolicyAssignmentsClient.Get(context.Context, string, string, *PolicyAssignmentsClientGetOptions) (PolicyAssignmentsClientGetResponse, error)`
- New function `*PolicyAssignmentsClient.GetByID(context.Context, string, *PolicyAssignmentsClientGetByIDOptions) (PolicyAssignmentsClientGetByIDResponse, error)`
- New function `*PolicyAssignmentsClient.NewListForManagementGroupPager(string, *PolicyAssignmentsClientListForManagementGroupOptions) *runtime.Pager[PolicyAssignmentsClientListForManagementGroupResponse]`
- New function `*PolicyAssignmentsClient.NewListForResourceGroupPager(string, *PolicyAssignmentsClientListForResourceGroupOptions) *runtime.Pager[PolicyAssignmentsClientListForResourceGroupResponse]`
- New function `*PolicyAssignmentsClient.NewListForResourcePager(string, string, string, string, string, *PolicyAssignmentsClientListForResourceOptions) *runtime.Pager[PolicyAssignmentsClientListForResourceResponse]`
- New function `*PolicyAssignmentsClient.NewListPager(*PolicyAssignmentsClientListOptions) *runtime.Pager[PolicyAssignmentsClientListResponse]`
- New function `*PolicyAssignmentsClient.Update(context.Context, string, string, PolicyAssignmentUpdate, *PolicyAssignmentsClientUpdateOptions) (PolicyAssignmentsClientUpdateResponse, error)`
- New function `*PolicyAssignmentsClient.UpdateByID(context.Context, string, PolicyAssignmentUpdate, *PolicyAssignmentsClientUpdateByIDOptions) (PolicyAssignmentsClientUpdateByIDResponse, error)`
- New function `NewPolicyDefinitionVersionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PolicyDefinitionVersionsClient, error)`
- New function `*PolicyDefinitionVersionsClient.CreateOrUpdate(context.Context, string, string, PolicyDefinitionVersion, *PolicyDefinitionVersionsClientCreateOrUpdateOptions) (PolicyDefinitionVersionsClientCreateOrUpdateResponse, error)`
- New function `*PolicyDefinitionVersionsClient.CreateOrUpdateAtManagementGroup(context.Context, string, string, string, PolicyDefinitionVersion, *PolicyDefinitionVersionsClientCreateOrUpdateAtManagementGroupOptions) (PolicyDefinitionVersionsClientCreateOrUpdateAtManagementGroupResponse, error)`
- New function `*PolicyDefinitionVersionsClient.Delete(context.Context, string, string, *PolicyDefinitionVersionsClientDeleteOptions) (PolicyDefinitionVersionsClientDeleteResponse, error)`
- New function `*PolicyDefinitionVersionsClient.DeleteAtManagementGroup(context.Context, string, string, string, *PolicyDefinitionVersionsClientDeleteAtManagementGroupOptions) (PolicyDefinitionVersionsClientDeleteAtManagementGroupResponse, error)`
- New function `*PolicyDefinitionVersionsClient.Get(context.Context, string, string, *PolicyDefinitionVersionsClientGetOptions) (PolicyDefinitionVersionsClientGetResponse, error)`
- New function `*PolicyDefinitionVersionsClient.GetAtManagementGroup(context.Context, string, string, string, *PolicyDefinitionVersionsClientGetAtManagementGroupOptions) (PolicyDefinitionVersionsClientGetAtManagementGroupResponse, error)`
- New function `*PolicyDefinitionVersionsClient.GetBuiltIn(context.Context, string, string, *PolicyDefinitionVersionsClientGetBuiltInOptions) (PolicyDefinitionVersionsClientGetBuiltInResponse, error)`
- New function `*PolicyDefinitionVersionsClient.ListAll(context.Context, *PolicyDefinitionVersionsClientListAllOptions) (PolicyDefinitionVersionsClientListAllResponse, error)`
- New function `*PolicyDefinitionVersionsClient.ListAllAtManagementGroup(context.Context, string, *PolicyDefinitionVersionsClientListAllAtManagementGroupOptions) (PolicyDefinitionVersionsClientListAllAtManagementGroupResponse, error)`
- New function `*PolicyDefinitionVersionsClient.ListAllBuiltins(context.Context, *PolicyDefinitionVersionsClientListAllBuiltinsOptions) (PolicyDefinitionVersionsClientListAllBuiltinsResponse, error)`
- New function `*PolicyDefinitionVersionsClient.NewListBuiltInPager(string, *PolicyDefinitionVersionsClientListBuiltInOptions) *runtime.Pager[PolicyDefinitionVersionsClientListBuiltInResponse]`
- New function `*PolicyDefinitionVersionsClient.NewListByManagementGroupPager(string, string, *PolicyDefinitionVersionsClientListByManagementGroupOptions) *runtime.Pager[PolicyDefinitionVersionsClientListByManagementGroupResponse]`
- New function `*PolicyDefinitionVersionsClient.NewListPager(string, *PolicyDefinitionVersionsClientListOptions) *runtime.Pager[PolicyDefinitionVersionsClientListResponse]`
- New function `NewPolicyDefinitionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PolicyDefinitionsClient, error)`
- New function `*PolicyDefinitionsClient.CreateOrUpdate(context.Context, string, PolicyDefinition, *PolicyDefinitionsClientCreateOrUpdateOptions) (PolicyDefinitionsClientCreateOrUpdateResponse, error)`
- New function `*PolicyDefinitionsClient.CreateOrUpdateAtManagementGroup(context.Context, string, string, PolicyDefinition, *PolicyDefinitionsClientCreateOrUpdateAtManagementGroupOptions) (PolicyDefinitionsClientCreateOrUpdateAtManagementGroupResponse, error)`
- New function `*PolicyDefinitionsClient.Delete(context.Context, string, *PolicyDefinitionsClientDeleteOptions) (PolicyDefinitionsClientDeleteResponse, error)`
- New function `*PolicyDefinitionsClient.DeleteAtManagementGroup(context.Context, string, string, *PolicyDefinitionsClientDeleteAtManagementGroupOptions) (PolicyDefinitionsClientDeleteAtManagementGroupResponse, error)`
- New function `*PolicyDefinitionsClient.Get(context.Context, string, *PolicyDefinitionsClientGetOptions) (PolicyDefinitionsClientGetResponse, error)`
- New function `*PolicyDefinitionsClient.GetAtManagementGroup(context.Context, string, string, *PolicyDefinitionsClientGetAtManagementGroupOptions) (PolicyDefinitionsClientGetAtManagementGroupResponse, error)`
- New function `*PolicyDefinitionsClient.GetBuiltIn(context.Context, string, *PolicyDefinitionsClientGetBuiltInOptions) (PolicyDefinitionsClientGetBuiltInResponse, error)`
- New function `*PolicyDefinitionsClient.NewListBuiltInPager(*PolicyDefinitionsClientListBuiltInOptions) *runtime.Pager[PolicyDefinitionsClientListBuiltInResponse]`
- New function `*PolicyDefinitionsClient.NewListByManagementGroupPager(string, *PolicyDefinitionsClientListByManagementGroupOptions) *runtime.Pager[PolicyDefinitionsClientListByManagementGroupResponse]`
- New function `*PolicyDefinitionsClient.NewListPager(*PolicyDefinitionsClientListOptions) *runtime.Pager[PolicyDefinitionsClientListResponse]`
- New function `NewPolicySetDefinitionVersionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PolicySetDefinitionVersionsClient, error)`
- New function `*PolicySetDefinitionVersionsClient.CreateOrUpdate(context.Context, string, string, PolicySetDefinitionVersion, *PolicySetDefinitionVersionsClientCreateOrUpdateOptions) (PolicySetDefinitionVersionsClientCreateOrUpdateResponse, error)`
- New function `*PolicySetDefinitionVersionsClient.CreateOrUpdateAtManagementGroup(context.Context, string, string, string, PolicySetDefinitionVersion, *PolicySetDefinitionVersionsClientCreateOrUpdateAtManagementGroupOptions) (PolicySetDefinitionVersionsClientCreateOrUpdateAtManagementGroupResponse, error)`
- New function `*PolicySetDefinitionVersionsClient.Delete(context.Context, string, string, *PolicySetDefinitionVersionsClientDeleteOptions) (PolicySetDefinitionVersionsClientDeleteResponse, error)`
- New function `*PolicySetDefinitionVersionsClient.DeleteAtManagementGroup(context.Context, string, string, string, *PolicySetDefinitionVersionsClientDeleteAtManagementGroupOptions) (PolicySetDefinitionVersionsClientDeleteAtManagementGroupResponse, error)`
- New function `*PolicySetDefinitionVersionsClient.Get(context.Context, string, string, *PolicySetDefinitionVersionsClientGetOptions) (PolicySetDefinitionVersionsClientGetResponse, error)`
- New function `*PolicySetDefinitionVersionsClient.GetAtManagementGroup(context.Context, string, string, string, *PolicySetDefinitionVersionsClientGetAtManagementGroupOptions) (PolicySetDefinitionVersionsClientGetAtManagementGroupResponse, error)`
- New function `*PolicySetDefinitionVersionsClient.GetBuiltIn(context.Context, string, string, *PolicySetDefinitionVersionsClientGetBuiltInOptions) (PolicySetDefinitionVersionsClientGetBuiltInResponse, error)`
- New function `*PolicySetDefinitionVersionsClient.ListAll(context.Context, *PolicySetDefinitionVersionsClientListAllOptions) (PolicySetDefinitionVersionsClientListAllResponse, error)`
- New function `*PolicySetDefinitionVersionsClient.ListAllAtManagementGroup(context.Context, string, *PolicySetDefinitionVersionsClientListAllAtManagementGroupOptions) (PolicySetDefinitionVersionsClientListAllAtManagementGroupResponse, error)`
- New function `*PolicySetDefinitionVersionsClient.ListAllBuiltins(context.Context, *PolicySetDefinitionVersionsClientListAllBuiltinsOptions) (PolicySetDefinitionVersionsClientListAllBuiltinsResponse, error)`
- New function `*PolicySetDefinitionVersionsClient.NewListBuiltInPager(string, *PolicySetDefinitionVersionsClientListBuiltInOptions) *runtime.Pager[PolicySetDefinitionVersionsClientListBuiltInResponse]`
- New function `*PolicySetDefinitionVersionsClient.NewListByManagementGroupPager(string, string, *PolicySetDefinitionVersionsClientListByManagementGroupOptions) *runtime.Pager[PolicySetDefinitionVersionsClientListByManagementGroupResponse]`
- New function `*PolicySetDefinitionVersionsClient.NewListPager(string, *PolicySetDefinitionVersionsClientListOptions) *runtime.Pager[PolicySetDefinitionVersionsClientListResponse]`
- New function `NewPolicySetDefinitionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PolicySetDefinitionsClient, error)`
- New function `*PolicySetDefinitionsClient.CreateOrUpdate(context.Context, string, PolicySetDefinition, *PolicySetDefinitionsClientCreateOrUpdateOptions) (PolicySetDefinitionsClientCreateOrUpdateResponse, error)`
- New function `*PolicySetDefinitionsClient.CreateOrUpdateAtManagementGroup(context.Context, string, string, PolicySetDefinition, *PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupOptions) (PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupResponse, error)`
- New function `*PolicySetDefinitionsClient.Delete(context.Context, string, *PolicySetDefinitionsClientDeleteOptions) (PolicySetDefinitionsClientDeleteResponse, error)`
- New function `*PolicySetDefinitionsClient.DeleteAtManagementGroup(context.Context, string, string, *PolicySetDefinitionsClientDeleteAtManagementGroupOptions) (PolicySetDefinitionsClientDeleteAtManagementGroupResponse, error)`
- New function `*PolicySetDefinitionsClient.Get(context.Context, string, *PolicySetDefinitionsClientGetOptions) (PolicySetDefinitionsClientGetResponse, error)`
- New function `*PolicySetDefinitionsClient.GetAtManagementGroup(context.Context, string, string, *PolicySetDefinitionsClientGetAtManagementGroupOptions) (PolicySetDefinitionsClientGetAtManagementGroupResponse, error)`
- New function `*PolicySetDefinitionsClient.GetBuiltIn(context.Context, string, *PolicySetDefinitionsClientGetBuiltInOptions) (PolicySetDefinitionsClientGetBuiltInResponse, error)`
- New function `*PolicySetDefinitionsClient.NewListBuiltInPager(*PolicySetDefinitionsClientListBuiltInOptions) *runtime.Pager[PolicySetDefinitionsClientListBuiltInResponse]`
- New function `*PolicySetDefinitionsClient.NewListByManagementGroupPager(string, *PolicySetDefinitionsClientListByManagementGroupOptions) *runtime.Pager[PolicySetDefinitionsClientListByManagementGroupResponse]`
- New function `*PolicySetDefinitionsClient.NewListPager(*PolicySetDefinitionsClientListOptions) *runtime.Pager[PolicySetDefinitionsClientListResponse]`
- New struct `Identity`
- New struct `NonComplianceMessage`
- New struct `Override`
- New struct `ParameterDefinitionsValue`
- New struct `ParameterDefinitionsValueMetadata`
- New struct `ParameterValuesValue`
- New struct `PolicyAssignment`
- New struct `PolicyAssignmentListResult`
- New struct `PolicyAssignmentProperties`
- New struct `PolicyAssignmentUpdate`
- New struct `PolicyAssignmentUpdateProperties`
- New struct `PolicyDefinition`
- New struct `PolicyDefinitionGroup`
- New struct `PolicyDefinitionListResult`
- New struct `PolicyDefinitionProperties`
- New struct `PolicyDefinitionReference`
- New struct `PolicyDefinitionVersion`
- New struct `PolicyDefinitionVersionListResult`
- New struct `PolicyDefinitionVersionProperties`
- New struct `PolicySetDefinition`
- New struct `PolicySetDefinitionListResult`
- New struct `PolicySetDefinitionProperties`
- New struct `PolicySetDefinitionVersion`
- New struct `PolicySetDefinitionVersionListResult`
- New struct `PolicySetDefinitionVersionProperties`
- New struct `ResourceSelector`
- New struct `Selector`
- New struct `UserAssignedIdentitiesValue`


## 2.1.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 2.0.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 2.0.0 (2023-03-27)
### Breaking Changes

- Struct `DeploymentScriptPropertiesBase` has been removed
- Struct `ScriptConfigurationBase` has been removed

### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module

## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentscripts` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).