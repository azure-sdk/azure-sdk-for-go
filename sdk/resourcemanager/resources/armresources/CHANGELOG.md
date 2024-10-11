# Release History

## 2.0.0 (2024-10-11)
### Breaking Changes

- Type of `Identity.UserAssignedIdentities` has been changed from `map[string]*IdentityUserAssignedIdentitiesValue` to `map[string]*UserAssignedIdentitiesValue`
- `ResourceIdentityTypeSystemAssignedUserAssigned` from enum `ResourceIdentityType` has been removed
- Enum `AliasPathAttributes` has been removed
- Enum `AliasPathTokenType` has been removed
- Enum `AliasPatternType` has been removed
- Enum `AliasType` has been removed
- Enum `ChangeType` has been removed
- Enum `DeploymentMode` has been removed
- Enum `ExpressionEvaluationOptionsScopeType` has been removed
- Enum `ExtendedLocationType` has been removed
- Enum `OnErrorDeploymentType` has been removed
- Enum `PropertyChangeType` has been removed
- Enum `ProviderAuthorizationConsentState` has been removed
- Enum `ProvisioningOperation` has been removed
- Enum `ProvisioningState` has been removed
- Enum `TagsPatchOperation` has been removed
- Enum `WhatIfResultFormat` has been removed
- Function `NewClient` has been removed
- Function `*Client.CheckExistence` has been removed
- Function `*Client.CheckExistenceByID` has been removed
- Function `*Client.BeginCreateOrUpdate` has been removed
- Function `*Client.BeginCreateOrUpdateByID` has been removed
- Function `*Client.BeginDelete` has been removed
- Function `*Client.BeginDeleteByID` has been removed
- Function `*Client.Get` has been removed
- Function `*Client.GetByID` has been removed
- Function `*Client.NewListByResourceGroupPager` has been removed
- Function `*Client.NewListPager` has been removed
- Function `*Client.BeginMoveResources` has been removed
- Function `*Client.BeginUpdate` has been removed
- Function `*Client.BeginUpdateByID` has been removed
- Function `*Client.BeginValidateMoveResources` has been removed
- Function `*ClientFactory.NewClient` has been removed
- Function `*ClientFactory.NewDeploymentOperationsClient` has been removed
- Function `*ClientFactory.NewDeploymentsClient` has been removed
- Function `*ClientFactory.NewOperationsClient` has been removed
- Function `*ClientFactory.NewProviderResourceTypesClient` has been removed
- Function `*ClientFactory.NewProvidersClient` has been removed
- Function `*ClientFactory.NewResourceGroupsClient` has been removed
- Function `*ClientFactory.NewTagsClient` has been removed
- Function `NewDeploymentOperationsClient` has been removed
- Function `*DeploymentOperationsClient.Get` has been removed
- Function `*DeploymentOperationsClient.GetAtManagementGroupScope` has been removed
- Function `*DeploymentOperationsClient.GetAtScope` has been removed
- Function `*DeploymentOperationsClient.GetAtSubscriptionScope` has been removed
- Function `*DeploymentOperationsClient.GetAtTenantScope` has been removed
- Function `*DeploymentOperationsClient.NewListAtManagementGroupScopePager` has been removed
- Function `*DeploymentOperationsClient.NewListAtScopePager` has been removed
- Function `*DeploymentOperationsClient.NewListAtSubscriptionScopePager` has been removed
- Function `*DeploymentOperationsClient.NewListAtTenantScopePager` has been removed
- Function `*DeploymentOperationsClient.NewListPager` has been removed
- Function `NewDeploymentsClient` has been removed
- Function `*DeploymentsClient.CalculateTemplateHash` has been removed
- Function `*DeploymentsClient.Cancel` has been removed
- Function `*DeploymentsClient.CancelAtManagementGroupScope` has been removed
- Function `*DeploymentsClient.CancelAtScope` has been removed
- Function `*DeploymentsClient.CancelAtSubscriptionScope` has been removed
- Function `*DeploymentsClient.CancelAtTenantScope` has been removed
- Function `*DeploymentsClient.CheckExistence` has been removed
- Function `*DeploymentsClient.CheckExistenceAtManagementGroupScope` has been removed
- Function `*DeploymentsClient.CheckExistenceAtScope` has been removed
- Function `*DeploymentsClient.CheckExistenceAtSubscriptionScope` has been removed
- Function `*DeploymentsClient.CheckExistenceAtTenantScope` has been removed
- Function `*DeploymentsClient.BeginCreateOrUpdate` has been removed
- Function `*DeploymentsClient.BeginCreateOrUpdateAtManagementGroupScope` has been removed
- Function `*DeploymentsClient.BeginCreateOrUpdateAtScope` has been removed
- Function `*DeploymentsClient.BeginCreateOrUpdateAtSubscriptionScope` has been removed
- Function `*DeploymentsClient.BeginCreateOrUpdateAtTenantScope` has been removed
- Function `*DeploymentsClient.BeginDelete` has been removed
- Function `*DeploymentsClient.BeginDeleteAtManagementGroupScope` has been removed
- Function `*DeploymentsClient.BeginDeleteAtScope` has been removed
- Function `*DeploymentsClient.BeginDeleteAtSubscriptionScope` has been removed
- Function `*DeploymentsClient.BeginDeleteAtTenantScope` has been removed
- Function `*DeploymentsClient.ExportTemplate` has been removed
- Function `*DeploymentsClient.ExportTemplateAtManagementGroupScope` has been removed
- Function `*DeploymentsClient.ExportTemplateAtScope` has been removed
- Function `*DeploymentsClient.ExportTemplateAtSubscriptionScope` has been removed
- Function `*DeploymentsClient.ExportTemplateAtTenantScope` has been removed
- Function `*DeploymentsClient.Get` has been removed
- Function `*DeploymentsClient.GetAtManagementGroupScope` has been removed
- Function `*DeploymentsClient.GetAtScope` has been removed
- Function `*DeploymentsClient.GetAtSubscriptionScope` has been removed
- Function `*DeploymentsClient.GetAtTenantScope` has been removed
- Function `*DeploymentsClient.NewListAtManagementGroupScopePager` has been removed
- Function `*DeploymentsClient.NewListAtScopePager` has been removed
- Function `*DeploymentsClient.NewListAtSubscriptionScopePager` has been removed
- Function `*DeploymentsClient.NewListAtTenantScopePager` has been removed
- Function `*DeploymentsClient.NewListByResourceGroupPager` has been removed
- Function `*DeploymentsClient.BeginValidate` has been removed
- Function `*DeploymentsClient.BeginValidateAtManagementGroupScope` has been removed
- Function `*DeploymentsClient.BeginValidateAtScope` has been removed
- Function `*DeploymentsClient.BeginValidateAtSubscriptionScope` has been removed
- Function `*DeploymentsClient.BeginValidateAtTenantScope` has been removed
- Function `*DeploymentsClient.BeginWhatIf` has been removed
- Function `*DeploymentsClient.BeginWhatIfAtManagementGroupScope` has been removed
- Function `*DeploymentsClient.BeginWhatIfAtSubscriptionScope` has been removed
- Function `*DeploymentsClient.BeginWhatIfAtTenantScope` has been removed
- Function `NewOperationsClient` has been removed
- Function `*OperationsClient.NewListPager` has been removed
- Function `NewProviderResourceTypesClient` has been removed
- Function `*ProviderResourceTypesClient.List` has been removed
- Function `NewProvidersClient` has been removed
- Function `*ProvidersClient.Get` has been removed
- Function `*ProvidersClient.GetAtTenantScope` has been removed
- Function `*ProvidersClient.NewListAtTenantScopePager` has been removed
- Function `*ProvidersClient.NewListPager` has been removed
- Function `*ProvidersClient.ProviderPermissions` has been removed
- Function `*ProvidersClient.Register` has been removed
- Function `*ProvidersClient.RegisterAtManagementGroupScope` has been removed
- Function `*ProvidersClient.Unregister` has been removed
- Function `NewResourceGroupsClient` has been removed
- Function `*ResourceGroupsClient.CheckExistence` has been removed
- Function `*ResourceGroupsClient.CreateOrUpdate` has been removed
- Function `*ResourceGroupsClient.BeginDelete` has been removed
- Function `*ResourceGroupsClient.BeginExportTemplate` has been removed
- Function `*ResourceGroupsClient.Get` has been removed
- Function `*ResourceGroupsClient.NewListPager` has been removed
- Function `*ResourceGroupsClient.Update` has been removed
- Function `NewTagsClient` has been removed
- Function `*TagsClient.CreateOrUpdate` has been removed
- Function `*TagsClient.CreateOrUpdateAtScope` has been removed
- Function `*TagsClient.CreateOrUpdateValue` has been removed
- Function `*TagsClient.Delete` has been removed
- Function `*TagsClient.DeleteAtScope` has been removed
- Function `*TagsClient.DeleteValue` has been removed
- Function `*TagsClient.GetAtScope` has been removed
- Function `*TagsClient.NewListPager` has been removed
- Function `*TagsClient.UpdateAtScope` has been removed
- Struct `APIProfile` has been removed
- Struct `Alias` has been removed
- Struct `AliasPath` has been removed
- Struct `AliasPathMetadata` has been removed
- Struct `AliasPattern` has been removed
- Struct `BasicDependency` has been removed
- Struct `DebugSetting` has been removed
- Struct `Dependency` has been removed
- Struct `Deployment` has been removed
- Struct `DeploymentExportResult` has been removed
- Struct `DeploymentExtended` has been removed
- Struct `DeploymentExtendedFilter` has been removed
- Struct `DeploymentListResult` has been removed
- Struct `DeploymentOperation` has been removed
- Struct `DeploymentOperationProperties` has been removed
- Struct `DeploymentOperationsListResult` has been removed
- Struct `DeploymentProperties` has been removed
- Struct `DeploymentPropertiesExtended` has been removed
- Struct `DeploymentValidateResult` has been removed
- Struct `DeploymentWhatIf` has been removed
- Struct `DeploymentWhatIfProperties` has been removed
- Struct `DeploymentWhatIfSettings` has been removed
- Struct `ExportTemplateRequest` has been removed
- Struct `ExpressionEvaluationOptions` has been removed
- Struct `ExtendedLocation` has been removed
- Struct `GenericResource` has been removed
- Struct `GenericResourceExpanded` has been removed
- Struct `GenericResourceFilter` has been removed
- Struct `HTTPMessage` has been removed
- Struct `IdentityUserAssignedIdentitiesValue` has been removed
- Struct `MoveInfo` has been removed
- Struct `OnErrorDeployment` has been removed
- Struct `OnErrorDeploymentExtended` has been removed
- Struct `Operation` has been removed
- Struct `OperationDisplay` has been removed
- Struct `OperationListResult` has been removed
- Struct `ParametersLink` has been removed
- Struct `Permission` has been removed
- Struct `Plan` has been removed
- Struct `Provider` has been removed
- Struct `ProviderConsentDefinition` has been removed
- Struct `ProviderExtendedLocation` has been removed
- Struct `ProviderListResult` has been removed
- Struct `ProviderPermission` has been removed
- Struct `ProviderPermissionListResult` has been removed
- Struct `ProviderRegistrationRequest` has been removed
- Struct `ProviderResourceType` has been removed
- Struct `ProviderResourceTypeListResult` has been removed
- Struct `Resource` has been removed
- Struct `ResourceGroup` has been removed
- Struct `ResourceGroupExportResult` has been removed
- Struct `ResourceGroupFilter` has been removed
- Struct `ResourceGroupListResult` has been removed
- Struct `ResourceGroupPatchable` has been removed
- Struct `ResourceGroupProperties` has been removed
- Struct `ResourceListResult` has been removed
- Struct `ResourceProviderOperationDisplayProperties` has been removed
- Struct `ResourceReference` has been removed
- Struct `RoleDefinition` has been removed
- Struct `SKU` has been removed
- Struct `ScopedDeployment` has been removed
- Struct `ScopedDeploymentWhatIf` has been removed
- Struct `StatusMessage` has been removed
- Struct `SubResource` has been removed
- Struct `TagCount` has been removed
- Struct `TagDetails` has been removed
- Struct `TagValue` has been removed
- Struct `Tags` has been removed
- Struct `TagsListResult` has been removed
- Struct `TagsPatchResource` has been removed
- Struct `TagsResource` has been removed
- Struct `TargetResource` has been removed
- Struct `TemplateHashResult` has been removed
- Struct `TemplateLink` has been removed
- Struct `WhatIfChange` has been removed
- Struct `WhatIfOperationProperties` has been removed
- Struct `WhatIfOperationResult` has been removed
- Struct `WhatIfPropertyChange` has been removed
- Struct `ZoneMapping` has been removed

### Features Added

- New enum type `AssignmentType` with values `AssignmentTypeCustom`, `AssignmentTypeNotSpecified`, `AssignmentTypeSystem`, `AssignmentTypeSystemHidden`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `EnforcementMode` with values `EnforcementModeDefault`, `EnforcementModeDoNotEnforce`
- New enum type `OverrideKind` with values `OverrideKindDefinitionVersion`, `OverrideKindPolicyEffect`
- New enum type `ParameterType` with values `ParameterTypeArray`, `ParameterTypeBoolean`, `ParameterTypeDateTime`, `ParameterTypeFloat`, `ParameterTypeInteger`, `ParameterTypeObject`, `ParameterTypeString`
- New enum type `PolicyType` with values `PolicyTypeBuiltIn`, `PolicyTypeCustom`, `PolicyTypeNotSpecified`, `PolicyTypeStatic`
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
- New struct `SystemData`
- New struct `UserAssignedIdentitiesValue`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.2.0-beta.3 (2023-10-09)

### Other Changes

- Updated to latest `azcore` beta.

## 1.2.0-beta.2 (2023-07-19)

### Bug Fixes

- Fixed a potential panic in faked paged and long-running operations.

## 1.2.0-beta.1 (2023-06-12)

### Features Added

- Support for test fakes and OpenTelemetry trace spans.

## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-27)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).