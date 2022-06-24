# Release History

## 0.6.0 (2022-06-24)
### Breaking Changes

- Function `*PolicyEventsClient.NewListQueryResultsForResourceGroupLevelPolicyAssignmentPager` parameter(s) have been changed from `(PolicyEventsResourceType, string, string, string, *QueryOptions, *PolicyEventsClientListQueryResultsForResourceGroupLevelPolicyAssignmentOptions)` to `(Enum3, string, string, Enum5, string, *QueryOptions, *PolicyEventsClientListQueryResultsForResourceGroupLevelPolicyAssignmentOptions)`
- Function `*PolicyEventsClient.NewListQueryResultsForSubscriptionPager` parameter(s) have been changed from `(PolicyEventsResourceType, string, *QueryOptions, *PolicyEventsClientListQueryResultsForSubscriptionOptions)` to `(Enum3, string, *QueryOptions, *PolicyEventsClientListQueryResultsForSubscriptionOptions)`
- Function `*PolicyStatesClient.NewListQueryResultsForPolicySetDefinitionPager` parameter(s) have been changed from `(PolicyStatesResource, string, string, *QueryOptions, *PolicyStatesClientListQueryResultsForPolicySetDefinitionOptions)` to `(PolicyStatesResource, string, Enum5, string, *QueryOptions, *PolicyStatesClientListQueryResultsForPolicySetDefinitionOptions)`
- Function `*PolicyStatesClient.NewListQueryResultsForPolicyDefinitionPager` parameter(s) have been changed from `(PolicyStatesResource, string, string, *QueryOptions, *PolicyStatesClientListQueryResultsForPolicyDefinitionOptions)` to `(PolicyStatesResource, string, Enum5, string, *QueryOptions, *PolicyStatesClientListQueryResultsForPolicyDefinitionOptions)`
- Function `*PolicyEventsClient.NewListQueryResultsForResourceGroupPager` parameter(s) have been changed from `(PolicyEventsResourceType, string, string, *QueryOptions, *PolicyEventsClientListQueryResultsForResourceGroupOptions)` to `(Enum3, string, string, *QueryOptions, *PolicyEventsClientListQueryResultsForResourceGroupOptions)`
- Function `*PolicyStatesClient.SummarizeForResource` parameter(s) have been changed from `(context.Context, PolicyStatesSummaryResourceType, string, *QueryOptions, *PolicyStatesClientSummarizeForResourceOptions)` to `(context.Context, Enum7, string, *QueryOptions, *PolicyStatesClientSummarizeForResourceOptions)`
- Function `*PolicyEventsClient.NewListQueryResultsForPolicyDefinitionPager` parameter(s) have been changed from `(PolicyEventsResourceType, string, string, *QueryOptions, *PolicyEventsClientListQueryResultsForPolicyDefinitionOptions)` to `(Enum3, string, Enum5, string, *QueryOptions, *PolicyEventsClientListQueryResultsForPolicyDefinitionOptions)`
- Function `*PolicyStatesClient.NewListQueryResultsForManagementGroupPager` parameter(s) have been changed from `(PolicyStatesResource, string, *QueryOptions, *PolicyStatesClientListQueryResultsForManagementGroupOptions)` to `(PolicyStatesResource, Enum4, string, *QueryOptions, *PolicyStatesClientListQueryResultsForManagementGroupOptions)`
- Function `*PolicyStatesClient.SummarizeForPolicyDefinition` parameter(s) have been changed from `(context.Context, PolicyStatesSummaryResourceType, string, string, *QueryOptions, *PolicyStatesClientSummarizeForPolicyDefinitionOptions)` to `(context.Context, Enum7, string, Enum5, string, *QueryOptions, *PolicyStatesClientSummarizeForPolicyDefinitionOptions)`
- Function `*PolicyStatesClient.SummarizeForResourceGroupLevelPolicyAssignment` parameter(s) have been changed from `(context.Context, PolicyStatesSummaryResourceType, string, string, string, *QueryOptions, *PolicyStatesClientSummarizeForResourceGroupLevelPolicyAssignmentOptions)` to `(context.Context, Enum7, string, string, Enum5, string, *QueryOptions, *PolicyStatesClientSummarizeForResourceGroupLevelPolicyAssignmentOptions)`
- Function `*PolicyStatesClient.SummarizeForPolicySetDefinition` parameter(s) have been changed from `(context.Context, PolicyStatesSummaryResourceType, string, string, *QueryOptions, *PolicyStatesClientSummarizeForPolicySetDefinitionOptions)` to `(context.Context, Enum7, string, Enum5, string, *QueryOptions, *PolicyStatesClientSummarizeForPolicySetDefinitionOptions)`
- Function `*PolicyEventsClient.NewListQueryResultsForPolicySetDefinitionPager` parameter(s) have been changed from `(PolicyEventsResourceType, string, string, *QueryOptions, *PolicyEventsClientListQueryResultsForPolicySetDefinitionOptions)` to `(Enum3, string, Enum5, string, *QueryOptions, *PolicyEventsClientListQueryResultsForPolicySetDefinitionOptions)`
- Function `*PolicyStatesClient.SummarizeForSubscriptionLevelPolicyAssignment` parameter(s) have been changed from `(context.Context, PolicyStatesSummaryResourceType, string, string, *QueryOptions, *PolicyStatesClientSummarizeForSubscriptionLevelPolicyAssignmentOptions)` to `(context.Context, Enum7, string, Enum5, string, *QueryOptions, *PolicyStatesClientSummarizeForSubscriptionLevelPolicyAssignmentOptions)`
- Function `*PolicyEventsClient.NewListQueryResultsForResourcePager` parameter(s) have been changed from `(PolicyEventsResourceType, string, *QueryOptions, *PolicyEventsClientListQueryResultsForResourceOptions)` to `(Enum3, string, *QueryOptions, *PolicyEventsClientListQueryResultsForResourceOptions)`
- Function `*PolicyEventsClient.NewListQueryResultsForSubscriptionLevelPolicyAssignmentPager` parameter(s) have been changed from `(PolicyEventsResourceType, string, string, *QueryOptions, *PolicyEventsClientListQueryResultsForSubscriptionLevelPolicyAssignmentOptions)` to `(Enum3, string, Enum5, string, *QueryOptions, *PolicyEventsClientListQueryResultsForSubscriptionLevelPolicyAssignmentOptions)`
- Function `*PolicyStatesClient.SummarizeForResourceGroup` parameter(s) have been changed from `(context.Context, PolicyStatesSummaryResourceType, string, string, *QueryOptions, *PolicyStatesClientSummarizeForResourceGroupOptions)` to `(context.Context, Enum7, string, string, *QueryOptions, *PolicyStatesClientSummarizeForResourceGroupOptions)`
- Function `*PolicyStatesClient.SummarizeForSubscription` parameter(s) have been changed from `(context.Context, PolicyStatesSummaryResourceType, string, *QueryOptions, *PolicyStatesClientSummarizeForSubscriptionOptions)` to `(context.Context, Enum7, string, *QueryOptions, *PolicyStatesClientSummarizeForSubscriptionOptions)`
- Function `*PolicyStatesClient.NewListQueryResultsForSubscriptionLevelPolicyAssignmentPager` parameter(s) have been changed from `(PolicyStatesResource, string, string, *QueryOptions, *PolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentOptions)` to `(PolicyStatesResource, string, Enum5, string, *QueryOptions, *PolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentOptions)`
- Function `*PolicyStatesClient.NewListQueryResultsForResourceGroupLevelPolicyAssignmentPager` parameter(s) have been changed from `(PolicyStatesResource, string, string, string, *QueryOptions, *PolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentOptions)` to `(PolicyStatesResource, string, string, Enum5, string, *QueryOptions, *PolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentOptions)`
- Function `*PolicyStatesClient.SummarizeForManagementGroup` parameter(s) have been changed from `(context.Context, PolicyStatesSummaryResourceType, string, *QueryOptions, *PolicyStatesClientSummarizeForManagementGroupOptions)` to `(context.Context, Enum7, Enum4, string, *QueryOptions, *PolicyStatesClientSummarizeForManagementGroupOptions)`
- Function `*PolicyEventsClient.NewListQueryResultsForManagementGroupPager` parameter(s) have been changed from `(PolicyEventsResourceType, string, *QueryOptions, *PolicyEventsClientListQueryResultsForManagementGroupOptions)` to `(Enum3, Enum4, string, *QueryOptions, *PolicyEventsClientListQueryResultsForManagementGroupOptions)`
- Type of `PolicyEvaluationResult.EvaluationDetails` has been changed from `*PolicyEvaluationDetails` to `*PolicyEvaluationDetailsAutoGenerated`
- Const `PolicyEventsResourceTypeDefault` has been removed
- Const `PolicyStatesSummaryResourceTypeLatest` has been removed
- Function `PossiblePolicyEventsResourceTypeValues` has been removed
- Function `*PolicyRestrictionsClient.CheckAtManagementGroupScope` has been removed
- Function `PossiblePolicyStatesSummaryResourceTypeValues` has been removed
- Struct `CheckManagementGroupRestrictionsRequest` has been removed
- Struct `PolicyRestrictionsClientCheckAtManagementGroupScopeOptions` has been removed
- Struct `PolicyRestrictionsClientCheckAtManagementGroupScopeResponse` has been removed

### Features Added

- New const `Enum3Default`
- New const `Enum7Latest`
- New const `Enum5MicrosoftAuthorization`
- New const `Enum4MicrosoftManagement`
- New function `PossibleEnum3Values() []Enum3`
- New function `PossibleEnum4Values() []Enum4`
- New function `PossibleEnum5Values() []Enum5`
- New function `PossibleEnum7Values() []Enum7`
- New struct `PolicyEvaluationDetailsAutoGenerated`


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).