# Release History

## 1.0.0 (2024-09-26)
### Breaking Changes

- Enum `ComplianceState` has been removed
- Enum `FieldRestrictionResult` has been removed
- Enum `PolicyEventsResourceType` has been removed
- Enum `PolicyStatesResource` has been removed
- Enum `PolicyStatesSummaryResourceType` has been removed
- Enum `PolicyTrackedResourcesResourceType` has been removed
- Function `NewAttestationsClient` has been removed
- Function `*AttestationsClient.BeginCreateOrUpdateAtResource` has been removed
- Function `*AttestationsClient.BeginCreateOrUpdateAtResourceGroup` has been removed
- Function `*AttestationsClient.BeginCreateOrUpdateAtSubscription` has been removed
- Function `*AttestationsClient.DeleteAtResource` has been removed
- Function `*AttestationsClient.DeleteAtResourceGroup` has been removed
- Function `*AttestationsClient.DeleteAtSubscription` has been removed
- Function `*AttestationsClient.GetAtResource` has been removed
- Function `*AttestationsClient.GetAtResourceGroup` has been removed
- Function `*AttestationsClient.GetAtSubscription` has been removed
- Function `*AttestationsClient.NewListForResourceGroupPager` has been removed
- Function `*AttestationsClient.NewListForResourcePager` has been removed
- Function `*AttestationsClient.NewListForSubscriptionPager` has been removed
- Function `*ClientFactory.NewAttestationsClient` has been removed
- Function `*ClientFactory.NewOperationsClient` has been removed
- Function `*ClientFactory.NewPolicyEventsClient` has been removed
- Function `*ClientFactory.NewPolicyMetadataClient` has been removed
- Function `*ClientFactory.NewPolicyRestrictionsClient` has been removed
- Function `*ClientFactory.NewPolicyStatesClient` has been removed
- Function `*ClientFactory.NewPolicyTrackedResourcesClient` has been removed
- Function `NewOperationsClient` has been removed
- Function `*OperationsClient.List` has been removed
- Function `NewPolicyEventsClient` has been removed
- Function `*PolicyEventsClient.NewListQueryResultsForManagementGroupPager` has been removed
- Function `*PolicyEventsClient.NewListQueryResultsForPolicyDefinitionPager` has been removed
- Function `*PolicyEventsClient.NewListQueryResultsForPolicySetDefinitionPager` has been removed
- Function `*PolicyEventsClient.NewListQueryResultsForResourceGroupLevelPolicyAssignmentPager` has been removed
- Function `*PolicyEventsClient.NewListQueryResultsForResourceGroupPager` has been removed
- Function `*PolicyEventsClient.NewListQueryResultsForResourcePager` has been removed
- Function `*PolicyEventsClient.NewListQueryResultsForSubscriptionLevelPolicyAssignmentPager` has been removed
- Function `*PolicyEventsClient.NewListQueryResultsForSubscriptionPager` has been removed
- Function `NewPolicyMetadataClient` has been removed
- Function `*PolicyMetadataClient.GetResource` has been removed
- Function `*PolicyMetadataClient.NewListPager` has been removed
- Function `NewPolicyRestrictionsClient` has been removed
- Function `*PolicyRestrictionsClient.CheckAtManagementGroupScope` has been removed
- Function `*PolicyRestrictionsClient.CheckAtResourceGroupScope` has been removed
- Function `*PolicyRestrictionsClient.CheckAtSubscriptionScope` has been removed
- Function `NewPolicyStatesClient` has been removed
- Function `*PolicyStatesClient.NewListQueryResultsForManagementGroupPager` has been removed
- Function `*PolicyStatesClient.NewListQueryResultsForPolicyDefinitionPager` has been removed
- Function `*PolicyStatesClient.NewListQueryResultsForPolicySetDefinitionPager` has been removed
- Function `*PolicyStatesClient.NewListQueryResultsForResourceGroupLevelPolicyAssignmentPager` has been removed
- Function `*PolicyStatesClient.NewListQueryResultsForResourceGroupPager` has been removed
- Function `*PolicyStatesClient.NewListQueryResultsForResourcePager` has been removed
- Function `*PolicyStatesClient.NewListQueryResultsForSubscriptionLevelPolicyAssignmentPager` has been removed
- Function `*PolicyStatesClient.NewListQueryResultsForSubscriptionPager` has been removed
- Function `*PolicyStatesClient.SummarizeForManagementGroup` has been removed
- Function `*PolicyStatesClient.SummarizeForPolicyDefinition` has been removed
- Function `*PolicyStatesClient.SummarizeForPolicySetDefinition` has been removed
- Function `*PolicyStatesClient.SummarizeForResource` has been removed
- Function `*PolicyStatesClient.SummarizeForResourceGroup` has been removed
- Function `*PolicyStatesClient.SummarizeForResourceGroupLevelPolicyAssignment` has been removed
- Function `*PolicyStatesClient.SummarizeForSubscription` has been removed
- Function `*PolicyStatesClient.SummarizeForSubscriptionLevelPolicyAssignment` has been removed
- Function `*PolicyStatesClient.BeginTriggerResourceGroupEvaluation` has been removed
- Function `*PolicyStatesClient.BeginTriggerSubscriptionEvaluation` has been removed
- Function `NewPolicyTrackedResourcesClient` has been removed
- Function `*PolicyTrackedResourcesClient.NewListQueryResultsForManagementGroupPager` has been removed
- Function `*PolicyTrackedResourcesClient.NewListQueryResultsForResourceGroupPager` has been removed
- Function `*PolicyTrackedResourcesClient.NewListQueryResultsForResourcePager` has been removed
- Function `*PolicyTrackedResourcesClient.NewListQueryResultsForSubscriptionPager` has been removed
- Struct `Attestation` has been removed
- Struct `AttestationEvidence` has been removed
- Struct `AttestationListResult` has been removed
- Struct `AttestationProperties` has been removed
- Struct `CheckManagementGroupRestrictionsRequest` has been removed
- Struct `CheckRestrictionsRequest` has been removed
- Struct `CheckRestrictionsResourceDetails` has been removed
- Struct `CheckRestrictionsResult` has been removed
- Struct `CheckRestrictionsResultContentEvaluationResult` has been removed
- Struct `ComplianceDetail` has been removed
- Struct `ComponentEventDetails` has been removed
- Struct `ComponentStateDetails` has been removed
- Struct `ErrorDefinitionAutoGenerated` has been removed
- Struct `ErrorDefinitionAutoGenerated2` has been removed
- Struct `ErrorResponseAutoGenerated` has been removed
- Struct `ErrorResponseAutoGenerated2` has been removed
- Struct `ExpressionEvaluationDetails` has been removed
- Struct `FieldRestriction` has been removed
- Struct `FieldRestrictions` has been removed
- Struct `IfNotExistsEvaluationDetails` has been removed
- Struct `Operation` has been removed
- Struct `OperationDisplay` has been removed
- Struct `OperationsListResults` has been removed
- Struct `PendingField` has been removed
- Struct `PolicyAssignmentSummary` has been removed
- Struct `PolicyDefinitionSummary` has been removed
- Struct `PolicyDetails` has been removed
- Struct `PolicyEvaluationDetails` has been removed
- Struct `PolicyEvaluationResult` has been removed
- Struct `PolicyEvent` has been removed
- Struct `PolicyEventsQueryResults` has been removed
- Struct `PolicyGroupSummary` has been removed
- Struct `PolicyMetadata` has been removed
- Struct `PolicyMetadataCollection` has been removed
- Struct `PolicyMetadataProperties` has been removed
- Struct `PolicyMetadataSlimProperties` has been removed
- Struct `PolicyReference` has been removed
- Struct `PolicyState` has been removed
- Struct `PolicyStatesQueryResults` has been removed
- Struct `PolicyTrackedResource` has been removed
- Struct `PolicyTrackedResourcesQueryResults` has been removed
- Struct `QueryFailure` has been removed
- Struct `QueryFailureError` has been removed
- Struct `Resource` has been removed
- Struct `SlimPolicyMetadata` has been removed
- Struct `SummarizeResults` has been removed
- Struct `Summary` has been removed
- Struct `SummaryResults` has been removed
- Struct `TrackedResourceModificationDetails` has been removed
- Field `Apply`, `Expand`, `From`, `OrderBy`, `Select`, `SkipToken`, `To` of struct `QueryOptions` has been removed

### Features Added

- New field `ResourceIDs` in struct `RemediationFilters`


## 0.8.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 0.7.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 0.7.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.6.0 (2022-10-07)
### Features Added

- New field `Metadata` in struct `AttestationProperties`
- New field `AssessmentDate` in struct `AttestationProperties`


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).