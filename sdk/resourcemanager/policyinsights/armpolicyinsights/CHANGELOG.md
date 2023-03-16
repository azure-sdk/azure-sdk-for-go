# Release History

## 0.7.0 (2023-03-16)
### Features Added

- New type alias `ComponentPolicyStatesResource` with values `ComponentPolicyStatesResourceLatest`
- New function `NewComponentPolicyStatesClient(azcore.TokenCredential, *arm.ClientOptions) (*ComponentPolicyStatesClient, error)`
- New function `*ComponentPolicyStatesClient.ListQueryResultsForPolicyDefinition(context.Context, string, string, ComponentPolicyStatesResource, *ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionOptions) (ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionResponse, error)`
- New function `*ComponentPolicyStatesClient.ListQueryResultsForResource(context.Context, string, ComponentPolicyStatesResource, *ComponentPolicyStatesClientListQueryResultsForResourceOptions) (ComponentPolicyStatesClientListQueryResultsForResourceResponse, error)`
- New function `*ComponentPolicyStatesClient.ListQueryResultsForResourceGroup(context.Context, string, string, ComponentPolicyStatesResource, *ComponentPolicyStatesClientListQueryResultsForResourceGroupOptions) (ComponentPolicyStatesClientListQueryResultsForResourceGroupResponse, error)`
- New function `*ComponentPolicyStatesClient.ListQueryResultsForResourceGroupLevelPolicyAssignment(context.Context, string, string, string, ComponentPolicyStatesResource, *ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentOptions) (ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentResponse, error)`
- New function `*ComponentPolicyStatesClient.ListQueryResultsForSubscription(context.Context, string, ComponentPolicyStatesResource, *ComponentPolicyStatesClientListQueryResultsForSubscriptionOptions) (ComponentPolicyStatesClientListQueryResultsForSubscriptionResponse, error)`
- New function `*ComponentPolicyStatesClient.ListQueryResultsForSubscriptionLevelPolicyAssignment(context.Context, string, string, ComponentPolicyStatesResource, *ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentOptions) (ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentResponse, error)`
- New struct `ComponentExpressionEvaluationDetails`
- New struct `ComponentPolicyEvaluationDetails`
- New struct `ComponentPolicyState`
- New struct `ComponentPolicyStatesClient`
- New struct `ComponentPolicyStatesQueryResults`
- New field `IsDataAction` in struct `Operation`


## 0.6.0 (2022-10-07)
### Features Added

- New field `Metadata` in struct `AttestationProperties`
- New field `AssessmentDate` in struct `AttestationProperties`


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).