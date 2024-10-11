# Release History

## 2.0.0 (2024-10-11)
### Breaking Changes

- Enum `DenySettingsMode` has been removed
- Enum `DenyStatusMode` has been removed
- Enum `DeploymentStackProvisioningState` has been removed
- Enum `DeploymentStacksDeleteDetachEnum` has been removed
- Enum `ResourceStatusMode` has been removed
- Enum `UnmanageActionManagementGroupMode` has been removed
- Enum `UnmanageActionResourceGroupMode` has been removed
- Enum `UnmanageActionResourceMode` has been removed
- Function `NewClient` has been removed
- Function `*Client.BeginCreateOrUpdateAtManagementGroup` has been removed
- Function `*Client.BeginCreateOrUpdateAtResourceGroup` has been removed
- Function `*Client.BeginCreateOrUpdateAtSubscription` has been removed
- Function `*Client.BeginDeleteAtManagementGroup` has been removed
- Function `*Client.BeginDeleteAtResourceGroup` has been removed
- Function `*Client.BeginDeleteAtSubscription` has been removed
- Function `*Client.ExportTemplateAtManagementGroup` has been removed
- Function `*Client.ExportTemplateAtResourceGroup` has been removed
- Function `*Client.ExportTemplateAtSubscription` has been removed
- Function `*Client.GetAtManagementGroup` has been removed
- Function `*Client.GetAtResourceGroup` has been removed
- Function `*Client.GetAtSubscription` has been removed
- Function `*Client.NewListAtManagementGroupPager` has been removed
- Function `*Client.NewListAtResourceGroupPager` has been removed
- Function `*Client.NewListAtSubscriptionPager` has been removed
- Function `*Client.BeginValidateStackAtManagementGroup` has been removed
- Function `*Client.BeginValidateStackAtResourceGroup` has been removed
- Function `*Client.BeginValidateStackAtSubscription` has been removed
- Function `*ClientFactory.NewClient` has been removed
- Struct `ActionOnUnmanage` has been removed
- Struct `DebugSetting` has been removed
- Struct `DenySettings` has been removed
- Struct `DeploymentParameter` has been removed
- Struct `DeploymentStack` has been removed
- Struct `DeploymentStackListResult` has been removed
- Struct `DeploymentStackProperties` has been removed
- Struct `DeploymentStackTemplateDefinition` has been removed
- Struct `DeploymentStackValidateProperties` has been removed
- Struct `DeploymentStackValidateResult` has been removed
- Struct `ErrorAdditionalInfo` has been removed
- Struct `ErrorDetail` has been removed
- Struct `KeyVaultParameterReference` has been removed
- Struct `KeyVaultReference` has been removed
- Struct `ManagedResourceReference` has been removed
- Struct `ParametersLink` has been removed
- Struct `ResourceReference` has been removed
- Struct `ResourceReferenceExtended` has been removed
- Struct `TemplateLink` has been removed

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


## 1.0.0 (2024-06-21)
### Breaking Changes

- Type of `DeploymentStackProperties.ActionOnUnmanage` has been changed from `*DeploymentStackPropertiesActionOnUnmanage` to `*ActionOnUnmanage`
- Type of `DeploymentStackProperties.Error` has been changed from `*ErrorResponse` to `*ErrorDetail`
- Type of `DeploymentStackProperties.Parameters` has been changed from `any` to `map[string]*DeploymentParameter`
- Type of `ResourceReferenceExtended.Error` has been changed from `*ErrorResponse` to `*ErrorDetail`
- `DeploymentStackProvisioningStateLocking` from enum `DeploymentStackProvisioningState` has been removed
- `ResourceStatusModeNone` from enum `ResourceStatusMode` has been removed
- Struct `DeploymentStackPropertiesActionOnUnmanage` has been removed
- Struct `ErrorResponse` has been removed

### Features Added

- New value `DeploymentStackProvisioningStateUpdatingDenyAssignments` added to enum type `DeploymentStackProvisioningState`
- New function `*Client.BeginValidateStackAtManagementGroup(context.Context, string, string, DeploymentStack, *ClientBeginValidateStackAtManagementGroupOptions) (*runtime.Poller[ClientValidateStackAtManagementGroupResponse], error)`
- New function `*Client.BeginValidateStackAtResourceGroup(context.Context, string, string, DeploymentStack, *ClientBeginValidateStackAtResourceGroupOptions) (*runtime.Poller[ClientValidateStackAtResourceGroupResponse], error)`
- New function `*Client.BeginValidateStackAtSubscription(context.Context, string, DeploymentStack, *ClientBeginValidateStackAtSubscriptionOptions) (*runtime.Poller[ClientValidateStackAtSubscriptionResponse], error)`
- New struct `ActionOnUnmanage`
- New struct `DeploymentParameter`
- New struct `DeploymentStackValidateProperties`
- New struct `DeploymentStackValidateResult`
- New struct `KeyVaultParameterReference`
- New struct `KeyVaultReference`
- New field `BypassStackOutOfSyncError` in struct `ClientBeginDeleteAtManagementGroupOptions`
- New field `BypassStackOutOfSyncError`, `UnmanageActionManagementGroups` in struct `ClientBeginDeleteAtResourceGroupOptions`
- New field `BypassStackOutOfSyncError`, `UnmanageActionManagementGroups` in struct `ClientBeginDeleteAtSubscriptionOptions`
- New field `BypassStackOutOfSyncError`, `CorrelationID` in struct `DeploymentStackProperties`


## 0.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 0.1.0 (2023-08-25)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentstacks` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).