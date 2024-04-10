# Release History

## 0.10.0 (2024-04-10)
### Features Added

- New function `*ClientFactory.NewDefinitionVersionsClient() *DefinitionVersionsClient`
- New function `*ClientFactory.NewSetDefinitionVersionsClient() *SetDefinitionVersionsClient`
- New function `NewDefinitionVersionsClient(string, string, string, azcore.TokenCredential, *arm.ClientOptions) (*DefinitionVersionsClient, error)`
- New function `*DefinitionVersionsClient.CreateOrUpdate(context.Context, DefinitionVersion, *DefinitionVersionsClientCreateOrUpdateOptions) (DefinitionVersionsClientCreateOrUpdateResponse, error)`
- New function `*DefinitionVersionsClient.CreateOrUpdateAtManagementGroup(context.Context, string, DefinitionVersion, *DefinitionVersionsClientCreateOrUpdateAtManagementGroupOptions) (DefinitionVersionsClientCreateOrUpdateAtManagementGroupResponse, error)`
- New function `*DefinitionVersionsClient.Delete(context.Context, *DefinitionVersionsClientDeleteOptions) (DefinitionVersionsClientDeleteResponse, error)`
- New function `*DefinitionVersionsClient.DeleteAtManagementGroup(context.Context, string, *DefinitionVersionsClientDeleteAtManagementGroupOptions) (DefinitionVersionsClientDeleteAtManagementGroupResponse, error)`
- New function `*DefinitionVersionsClient.Get(context.Context, *DefinitionVersionsClientGetOptions) (DefinitionVersionsClientGetResponse, error)`
- New function `*DefinitionVersionsClient.GetAtManagementGroup(context.Context, string, *DefinitionVersionsClientGetAtManagementGroupOptions) (DefinitionVersionsClientGetAtManagementGroupResponse, error)`
- New function `*DefinitionVersionsClient.GetBuiltIn(context.Context, *DefinitionVersionsClientGetBuiltInOptions) (DefinitionVersionsClientGetBuiltInResponse, error)`
- New function `*DefinitionVersionsClient.ListAll(context.Context, *DefinitionVersionsClientListAllOptions) (DefinitionVersionsClientListAllResponse, error)`
- New function `*DefinitionVersionsClient.ListAllAtManagementGroup(context.Context, string, *DefinitionVersionsClientListAllAtManagementGroupOptions) (DefinitionVersionsClientListAllAtManagementGroupResponse, error)`
- New function `*DefinitionVersionsClient.ListAllBuiltins(context.Context, *DefinitionVersionsClientListAllBuiltinsOptions) (DefinitionVersionsClientListAllBuiltinsResponse, error)`
- New function `*DefinitionVersionsClient.NewListBuiltInPager(*DefinitionVersionsClientListBuiltInOptions) *runtime.Pager[DefinitionVersionsClientListBuiltInResponse]`
- New function `*DefinitionVersionsClient.NewListByManagementGroupPager(string, *DefinitionVersionsClientListByManagementGroupOptions) *runtime.Pager[DefinitionVersionsClientListByManagementGroupResponse]`
- New function `*DefinitionVersionsClient.NewListPager(*DefinitionVersionsClientListOptions) *runtime.Pager[DefinitionVersionsClientListResponse]`
- New function `NewSetDefinitionVersionsClient(string, string, string, azcore.TokenCredential, *arm.ClientOptions) (*SetDefinitionVersionsClient, error)`
- New function `*SetDefinitionVersionsClient.CreateOrUpdate(context.Context, SetDefinitionVersion, *SetDefinitionVersionsClientCreateOrUpdateOptions) (SetDefinitionVersionsClientCreateOrUpdateResponse, error)`
- New function `*SetDefinitionVersionsClient.CreateOrUpdateAtManagementGroup(context.Context, string, SetDefinitionVersion, *SetDefinitionVersionsClientCreateOrUpdateAtManagementGroupOptions) (SetDefinitionVersionsClientCreateOrUpdateAtManagementGroupResponse, error)`
- New function `*SetDefinitionVersionsClient.Delete(context.Context, *SetDefinitionVersionsClientDeleteOptions) (SetDefinitionVersionsClientDeleteResponse, error)`
- New function `*SetDefinitionVersionsClient.DeleteAtManagementGroup(context.Context, string, *SetDefinitionVersionsClientDeleteAtManagementGroupOptions) (SetDefinitionVersionsClientDeleteAtManagementGroupResponse, error)`
- New function `*SetDefinitionVersionsClient.Get(context.Context, *SetDefinitionVersionsClientGetOptions) (SetDefinitionVersionsClientGetResponse, error)`
- New function `*SetDefinitionVersionsClient.GetAtManagementGroup(context.Context, string, *SetDefinitionVersionsClientGetAtManagementGroupOptions) (SetDefinitionVersionsClientGetAtManagementGroupResponse, error)`
- New function `*SetDefinitionVersionsClient.GetBuiltIn(context.Context, *SetDefinitionVersionsClientGetBuiltInOptions) (SetDefinitionVersionsClientGetBuiltInResponse, error)`
- New function `*SetDefinitionVersionsClient.ListAll(context.Context, *SetDefinitionVersionsClientListAllOptions) (SetDefinitionVersionsClientListAllResponse, error)`
- New function `*SetDefinitionVersionsClient.ListAllAtManagementGroup(context.Context, string, *SetDefinitionVersionsClientListAllAtManagementGroupOptions) (SetDefinitionVersionsClientListAllAtManagementGroupResponse, error)`
- New function `*SetDefinitionVersionsClient.ListAllBuiltins(context.Context, *SetDefinitionVersionsClientListAllBuiltinsOptions) (SetDefinitionVersionsClientListAllBuiltinsResponse, error)`
- New function `*SetDefinitionVersionsClient.NewListBuiltInPager(*SetDefinitionVersionsClientListBuiltInOptions) *runtime.Pager[SetDefinitionVersionsClientListBuiltInResponse]`
- New function `*SetDefinitionVersionsClient.NewListByManagementGroupPager(string, *SetDefinitionVersionsClientListByManagementGroupOptions) *runtime.Pager[SetDefinitionVersionsClientListByManagementGroupResponse]`
- New function `*SetDefinitionVersionsClient.NewListPager(*SetDefinitionVersionsClientListOptions) *runtime.Pager[SetDefinitionVersionsClientListResponse]`
- New struct `DefinitionVersion`
- New struct `DefinitionVersionListResult`
- New struct `DefinitionVersionProperties`
- New struct `SetDefinitionVersion`
- New struct `SetDefinitionVersionListResult`
- New struct `SetDefinitionVersionProperties`
- New field `Version`, `Versions` in struct `DefinitionProperties`
- New field `Schema` in struct `ParameterDefinitionsValue`
- New field `Version`, `Versions` in struct `SetDefinitionProperties`


## 0.9.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 0.8.0 (2023-10-27)
### Features Added

- New enum type `AssignmentScopeValidation` with values `AssignmentScopeValidationDefault`, `AssignmentScopeValidationDoNotValidate`
- New enum type `OverrideKind` with values `OverrideKindPolicyEffect`
- New enum type `SelectorKind` with values `SelectorKindPolicyDefinitionReferenceID`, `SelectorKindResourceLocation`, `SelectorKindResourceType`, `SelectorKindResourceWithoutLocation`
- New function `*ClientFactory.NewVariableValuesClient() *VariableValuesClient`
- New function `*ClientFactory.NewVariablesClient() *VariablesClient`
- New function `*ExemptionsClient.Update(context.Context, string, string, ExemptionUpdate, *ExemptionsClientUpdateOptions) (ExemptionsClientUpdateResponse, error)`
- New function `NewVariableValuesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*VariableValuesClient, error)`
- New function `*VariableValuesClient.CreateOrUpdate(context.Context, string, string, VariableValue, *VariableValuesClientCreateOrUpdateOptions) (VariableValuesClientCreateOrUpdateResponse, error)`
- New function `*VariableValuesClient.CreateOrUpdateAtManagementGroup(context.Context, string, string, string, VariableValue, *VariableValuesClientCreateOrUpdateAtManagementGroupOptions) (VariableValuesClientCreateOrUpdateAtManagementGroupResponse, error)`
- New function `*VariableValuesClient.Delete(context.Context, string, string, *VariableValuesClientDeleteOptions) (VariableValuesClientDeleteResponse, error)`
- New function `*VariableValuesClient.DeleteAtManagementGroup(context.Context, string, string, string, *VariableValuesClientDeleteAtManagementGroupOptions) (VariableValuesClientDeleteAtManagementGroupResponse, error)`
- New function `*VariableValuesClient.Get(context.Context, string, string, *VariableValuesClientGetOptions) (VariableValuesClientGetResponse, error)`
- New function `*VariableValuesClient.GetAtManagementGroup(context.Context, string, string, string, *VariableValuesClientGetAtManagementGroupOptions) (VariableValuesClientGetAtManagementGroupResponse, error)`
- New function `*VariableValuesClient.NewListForManagementGroupPager(string, string, *VariableValuesClientListForManagementGroupOptions) *runtime.Pager[VariableValuesClientListForManagementGroupResponse]`
- New function `*VariableValuesClient.NewListPager(string, *VariableValuesClientListOptions) *runtime.Pager[VariableValuesClientListResponse]`
- New function `NewVariablesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*VariablesClient, error)`
- New function `*VariablesClient.CreateOrUpdate(context.Context, string, Variable, *VariablesClientCreateOrUpdateOptions) (VariablesClientCreateOrUpdateResponse, error)`
- New function `*VariablesClient.CreateOrUpdateAtManagementGroup(context.Context, string, string, Variable, *VariablesClientCreateOrUpdateAtManagementGroupOptions) (VariablesClientCreateOrUpdateAtManagementGroupResponse, error)`
- New function `*VariablesClient.Delete(context.Context, string, *VariablesClientDeleteOptions) (VariablesClientDeleteResponse, error)`
- New function `*VariablesClient.DeleteAtManagementGroup(context.Context, string, string, *VariablesClientDeleteAtManagementGroupOptions) (VariablesClientDeleteAtManagementGroupResponse, error)`
- New function `*VariablesClient.Get(context.Context, string, *VariablesClientGetOptions) (VariablesClientGetResponse, error)`
- New function `*VariablesClient.GetAtManagementGroup(context.Context, string, string, *VariablesClientGetAtManagementGroupOptions) (VariablesClientGetAtManagementGroupResponse, error)`
- New function `*VariablesClient.NewListForManagementGroupPager(string, *VariablesClientListForManagementGroupOptions) *runtime.Pager[VariablesClientListForManagementGroupResponse]`
- New function `*VariablesClient.NewListPager(*VariablesClientListOptions) *runtime.Pager[VariablesClientListResponse]`
- New struct `AssignmentUpdateProperties`
- New struct `ExemptionUpdate`
- New struct `ExemptionUpdateProperties`
- New struct `Override`
- New struct `ResourceSelector`
- New struct `Selector`
- New struct `Variable`
- New struct `VariableColumn`
- New struct `VariableListResult`
- New struct `VariableProperties`
- New struct `VariableValue`
- New struct `VariableValueColumnValue`
- New struct `VariableValueListResult`
- New struct `VariableValueProperties`
- New field `Overrides`, `ResourceSelectors` in struct `AssignmentProperties`
- New field `Properties` in struct `AssignmentUpdate`
- New field `AssignmentScopeValidation`, `ResourceSelectors` in struct `ExemptionProperties`


## 0.7.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 0.7.0 (2023-03-27)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.6.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).