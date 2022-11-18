# Release History

## 0.7.0 (2022-11-18)
### Breaking Changes

- Struct `CloudError` has been removed

### Features Added

- New const `OverrideKindPolicyEffect`
- New const `SelectorKindResourceLocation`
- New const `SelectorKindPolicyDefinitionReferenceID`
- New const `AssignmentScopeValidationDoNotValidate`
- New const `AssignmentScopeValidationDefault`
- New const `SelectorKindResourceType`
- New const `SelectorKindResourceWithoutLocation`
- New type alias `OverrideKind`
- New type alias `AssignmentScopeValidation`
- New type alias `SelectorKind`
- New function `*VariablesClient.NewListForManagementGroupPager(string, *VariablesClientListForManagementGroupOptions) *runtime.Pager[VariablesClientListForManagementGroupResponse]`
- New function `*VariablesClient.CreateOrUpdateAtManagementGroup(context.Context, string, string, Variable, *VariablesClientCreateOrUpdateAtManagementGroupOptions) (VariablesClientCreateOrUpdateAtManagementGroupResponse, error)`
- New function `*VariablesClient.NewListPager(*VariablesClientListOptions) *runtime.Pager[VariablesClientListResponse]`
- New function `PossibleAssignmentScopeValidationValues() []AssignmentScopeValidation`
- New function `*VariableValuesClient.Delete(context.Context, string, string, *VariableValuesClientDeleteOptions) (VariableValuesClientDeleteResponse, error)`
- New function `*VariableValuesClient.DeleteAtManagementGroup(context.Context, string, string, string, *VariableValuesClientDeleteAtManagementGroupOptions) (VariableValuesClientDeleteAtManagementGroupResponse, error)`
- New function `*VariableValuesClient.Get(context.Context, string, string, *VariableValuesClientGetOptions) (VariableValuesClientGetResponse, error)`
- New function `PossibleOverrideKindValues() []OverrideKind`
- New function `*VariablesClient.Delete(context.Context, string, *VariablesClientDeleteOptions) (VariablesClientDeleteResponse, error)`
- New function `*VariablesClient.GetAtManagementGroup(context.Context, string, string, *VariablesClientGetAtManagementGroupOptions) (VariablesClientGetAtManagementGroupResponse, error)`
- New function `*VariableValuesClient.NewListPager(string, *VariableValuesClientListOptions) *runtime.Pager[VariableValuesClientListResponse]`
- New function `*VariablesClient.CreateOrUpdate(context.Context, string, Variable, *VariablesClientCreateOrUpdateOptions) (VariablesClientCreateOrUpdateResponse, error)`
- New function `*VariableValuesClient.NewListForManagementGroupPager(string, string, *VariableValuesClientListForManagementGroupOptions) *runtime.Pager[VariableValuesClientListForManagementGroupResponse]`
- New function `*VariableValuesClient.CreateOrUpdate(context.Context, string, string, VariableValue, *VariableValuesClientCreateOrUpdateOptions) (VariableValuesClientCreateOrUpdateResponse, error)`
- New function `*VariableValuesClient.GetAtManagementGroup(context.Context, string, string, string, *VariableValuesClientGetAtManagementGroupOptions) (VariableValuesClientGetAtManagementGroupResponse, error)`
- New function `NewVariableValuesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*VariableValuesClient, error)`
- New function `*VariablesClient.DeleteAtManagementGroup(context.Context, string, string, *VariablesClientDeleteAtManagementGroupOptions) (VariablesClientDeleteAtManagementGroupResponse, error)`
- New function `PossibleSelectorKindValues() []SelectorKind`
- New function `*VariablesClient.Get(context.Context, string, *VariablesClientGetOptions) (VariablesClientGetResponse, error)`
- New function `*ExemptionsClient.Update(context.Context, string, string, ExemptionUpdate, *ExemptionsClientUpdateOptions) (ExemptionsClientUpdateResponse, error)`
- New function `*VariableValuesClient.CreateOrUpdateAtManagementGroup(context.Context, string, string, string, VariableValue, *VariableValuesClientCreateOrUpdateAtManagementGroupOptions) (VariableValuesClientCreateOrUpdateAtManagementGroupResponse, error)`
- New function `NewVariablesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*VariablesClient, error)`
- New struct `AssignmentUpdateProperties`
- New struct `ExemptionUpdate`
- New struct `ExemptionUpdateProperties`
- New struct `ExemptionsClientUpdateOptions`
- New struct `ExemptionsClientUpdateResponse`
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
- New struct `VariableValuesClient`
- New struct `VariableValuesClientCreateOrUpdateAtManagementGroupOptions`
- New struct `VariableValuesClientCreateOrUpdateAtManagementGroupResponse`
- New struct `VariableValuesClientCreateOrUpdateOptions`
- New struct `VariableValuesClientCreateOrUpdateResponse`
- New struct `VariableValuesClientDeleteAtManagementGroupOptions`
- New struct `VariableValuesClientDeleteAtManagementGroupResponse`
- New struct `VariableValuesClientDeleteOptions`
- New struct `VariableValuesClientDeleteResponse`
- New struct `VariableValuesClientGetAtManagementGroupOptions`
- New struct `VariableValuesClientGetAtManagementGroupResponse`
- New struct `VariableValuesClientGetOptions`
- New struct `VariableValuesClientGetResponse`
- New struct `VariableValuesClientListForManagementGroupOptions`
- New struct `VariableValuesClientListForManagementGroupResponse`
- New struct `VariableValuesClientListOptions`
- New struct `VariableValuesClientListResponse`
- New struct `VariablesClient`
- New struct `VariablesClientCreateOrUpdateAtManagementGroupOptions`
- New struct `VariablesClientCreateOrUpdateAtManagementGroupResponse`
- New struct `VariablesClientCreateOrUpdateOptions`
- New struct `VariablesClientCreateOrUpdateResponse`
- New struct `VariablesClientDeleteAtManagementGroupOptions`
- New struct `VariablesClientDeleteAtManagementGroupResponse`
- New struct `VariablesClientDeleteOptions`
- New struct `VariablesClientDeleteResponse`
- New struct `VariablesClientGetAtManagementGroupOptions`
- New struct `VariablesClientGetAtManagementGroupResponse`
- New struct `VariablesClientGetOptions`
- New struct `VariablesClientGetResponse`
- New struct `VariablesClientListForManagementGroupOptions`
- New struct `VariablesClientListForManagementGroupResponse`
- New struct `VariablesClientListOptions`
- New struct `VariablesClientListResponse`
- New field `AssignmentScopeValidation` in struct `ExemptionProperties`
- New field `ResourceSelectors` in struct `ExemptionProperties`
- New field `Properties` in struct `AssignmentUpdate`
- New field `ResourceSelectors` in struct `AssignmentProperties`
- New field `Overrides` in struct `AssignmentProperties`


## 0.6.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).