# Release History

## 1.3.0-beta.1 (2025-04-25)
### Features Added

- New enum type `Aggregated` with values `AggregatedDay`, `AggregatedMonth`, `AggregatedWeek`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `Duration` with values `DurationFourteen`, `DurationNinety`, `DurationSeven`, `DurationSixty`, `DurationThirty`, `DurationTwentyOne`
- New enum type `PredictionType` with values `PredictionTypePredictiveRightsizing`
- New enum type `Priority` with values `PriorityCritical`, `PriorityHigh`, `PriorityInformational`, `PriorityLow`, `PriorityMedium`
- New enum type `PriorityName` with values `PriorityNameHigh`, `PriorityNameLow`, `PriorityNameMedium`
- New enum type `Reason` with values `ReasonAlternativeSolution`, `ReasonExcessiveInvestment`, `ReasonIncompatible`, `ReasonRiskAccepted`, `ReasonTooComplex`, `ReasonUnclear`
- New enum type `ReasonForRejectionName` with values `ReasonForRejectionNameNotARisk`, `ReasonForRejectionNameRiskAccepted`
- New enum type `RecommendationStatusName` with values `RecommendationStatusNameApproved`, `RecommendationStatusNamePending`, `RecommendationStatusNameRejected`
- New enum type `ReviewStatus` with values `ReviewStatusCompleted`, `ReviewStatusInProgress`, `ReviewStatusNew`, `ReviewStatusTriaged`
- New enum type `State` with values `StateApproved`, `StateCompleted`, `StateDismissed`, `StateInProgress`, `StatePending`, `StatePostponed`, `StateRejected`
- New function `NewAssessmentTypesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AssessmentTypesClient, error)`
- New function `*AssessmentTypesClient.NewListPager(*AssessmentTypesClientListOptions) *runtime.Pager[AssessmentTypesClientListResponse]`
- New function `NewAssessmentsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AssessmentsClient, error)`
- New function `*AssessmentsClient.Delete(context.Context, string, *AssessmentsClientDeleteOptions) (AssessmentsClientDeleteResponse, error)`
- New function `*AssessmentsClient.Get(context.Context, string, *AssessmentsClientGetOptions) (AssessmentsClientGetResponse, error)`
- New function `*AssessmentsClient.NewListPager(*AssessmentsClientListOptions) *runtime.Pager[AssessmentsClientListResponse]`
- New function `*AssessmentsClient.Put(context.Context, string, AssessmentResult, *AssessmentsClientPutOptions) (AssessmentsClientPutResponse, error)`
- New function `*ClientFactory.NewAssessmentTypesClient() *AssessmentTypesClient`
- New function `*ClientFactory.NewAssessmentsClient() *AssessmentsClient`
- New function `*ClientFactory.NewManagementClient() *ManagementClient`
- New function `*ClientFactory.NewResiliencyReviewsClient() *ResiliencyReviewsClient`
- New function `*ClientFactory.NewScoresClient() *ScoresClient`
- New function `*ClientFactory.NewTriageRecommendationsClient() *TriageRecommendationsClient`
- New function `*ClientFactory.NewTriageResourcesClient() *TriageResourcesClient`
- New function `*ClientFactory.NewWorkloadsClient() *WorkloadsClient`
- New function `NewManagementClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagementClient, error)`
- New function `*ManagementClient.Predict(context.Context, PredictionRequest, *ManagementClientPredictOptions) (ManagementClientPredictResponse, error)`
- New function `PossiblePriorityValues() []Priority`
- New function `PossibleReasonValues() []Reason`
- New function `*RecommendationsClient.Patch(context.Context, string, string, TrackedRecommendationPropertiesPayload, *RecommendationsClientPatchOptions) (RecommendationsClientPatchResponse, error)`
- New function `NewResiliencyReviewsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ResiliencyReviewsClient, error)`
- New function `*ResiliencyReviewsClient.Get(context.Context, string, *ResiliencyReviewsClientGetOptions) (ResiliencyReviewsClientGetResponse, error)`
- New function `*ResiliencyReviewsClient.NewListPager(*ResiliencyReviewsClientListOptions) *runtime.Pager[ResiliencyReviewsClientListResponse]`
- New function `NewScoresClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ScoresClient, error)`
- New function `*ScoresClient.Get(context.Context, string, *ScoresClientGetOptions) (ScoresClientGetResponse, error)`
- New function `*ScoresClient.List(context.Context, *ScoresClientListOptions) (ScoresClientListResponse, error)`
- New function `NewTriageRecommendationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*TriageRecommendationsClient, error)`
- New function `*TriageRecommendationsClient.ApproveTriageRecommendation(context.Context, string, string, *TriageRecommendationsClientApproveTriageRecommendationOptions) (TriageRecommendationsClientApproveTriageRecommendationResponse, error)`
- New function `*TriageRecommendationsClient.Get(context.Context, string, string, *TriageRecommendationsClientGetOptions) (TriageRecommendationsClientGetResponse, error)`
- New function `*TriageRecommendationsClient.NewListPager(string, *TriageRecommendationsClientListOptions) *runtime.Pager[TriageRecommendationsClientListResponse]`
- New function `*TriageRecommendationsClient.RejectTriageRecommendation(context.Context, string, string, RecommendationRejectBody, *TriageRecommendationsClientRejectTriageRecommendationOptions) (TriageRecommendationsClientRejectTriageRecommendationResponse, error)`
- New function `*TriageRecommendationsClient.ResetTriageRecommendation(context.Context, string, string, *TriageRecommendationsClientResetTriageRecommendationOptions) (TriageRecommendationsClientResetTriageRecommendationResponse, error)`
- New function `NewTriageResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*TriageResourcesClient, error)`
- New function `*TriageResourcesClient.Get(context.Context, string, string, string, *TriageResourcesClientGetOptions) (TriageResourcesClientGetResponse, error)`
- New function `*TriageResourcesClient.NewListPager(string, string, *TriageResourcesClientListOptions) *runtime.Pager[TriageResourcesClientListResponse]`
- New function `NewWorkloadsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*WorkloadsClient, error)`
- New function `*WorkloadsClient.NewListPager(*WorkloadsClientListOptions) *runtime.Pager[WorkloadsClientListResponse]`
- New struct `AssessmentListResult`
- New struct `AssessmentResult`
- New struct `AssessmentResultProperties`
- New struct `AssessmentTypeListResult`
- New struct `AssessmentTypeResult`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New struct `PredictionRequest`
- New struct `PredictionRequestProperties`
- New struct `PredictionResponse`
- New struct `PredictionResponseProperties`
- New struct `ProxyResource`
- New struct `RecommendationPropertiesResourceWorkload`
- New struct `RecommendationPropertiesReview`
- New struct `RecommendationRejectBody`
- New struct `ResiliencyReview`
- New struct `ResiliencyReviewCollection`
- New struct `ResiliencyReviewProperties`
- New struct `ScoreEntity`
- New struct `ScoreEntityForAdvisor`
- New struct `ScoreEntityForAdvisorProperties`
- New struct `ScoreResponse`
- New struct `SystemData`
- New struct `TimeSeriesEntityItem`
- New struct `TrackedRecommendationProperties`
- New struct `TrackedRecommendationPropertiesPayload`
- New struct `TrackedRecommendationPropertiesPayloadProperties`
- New struct `TriageRecommendation`
- New struct `TriageRecommendationCollection`
- New struct `TriageRecommendationProperties`
- New struct `TriageResource`
- New struct `TriageResourceCollection`
- New struct `TriageResourceProperties`
- New struct `WorkloadListResult`
- New struct `WorkloadResult`
- New field `SystemData` in struct `ConfigData`
- New field `Duration` in struct `ConfigDataProperties`
- New field `Notes`, `ResourceWorkload`, `Review`, `SourceSystem`, `Tracked`, `TrackedProperties` in struct `RecommendationProperties`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `ResourceRecommendationBase`
- New field `SystemData` in struct `SuppressionContract`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.0 (2023-03-27)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).