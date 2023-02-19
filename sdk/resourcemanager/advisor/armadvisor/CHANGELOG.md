# Release History

## 1.1.0 (2023-02-19)
### Features Added

- New type alias `Aggregated` with values `AggregatedDay`, `AggregatedMonth`, `AggregatedWeek`
- New type alias `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New type alias `Duration` with values `DurationFourteen`, `DurationNinety`, `DurationSeven`, `DurationSixty`, `DurationThirty`, `DurationTwentyOne`
- New type alias `PredictionType` with values `PredictionTypePredictiveRightsizing`
- New function `NewManagementClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagementClient, error)`
- New function `*ManagementClient.Predict(context.Context, PredictionRequest, *ManagementClientPredictOptions) (ManagementClientPredictResponse, error)`
- New function `NewScoresClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ScoresClient, error)`
- New function `*ScoresClient.Get(context.Context, string, *ScoresClientGetOptions) (ScoresClientGetResponse, error)`
- New function `*ScoresClient.List(context.Context, *ScoresClientListOptions) (ScoresClientListResponse, error)`
- New struct `ManagementClient`
- New struct `PredictionRequest`
- New struct `PredictionRequestProperties`
- New struct `PredictionResponse`
- New struct `PredictionResponseProperties`
- New struct `ProxyResource`
- New struct `ScoreEntity`
- New struct `ScoreEntityForAdvisor`
- New struct `ScoreEntityForAdvisorProperties`
- New struct `ScoreResponse`
- New struct `ScoresClient`
- New struct `SystemData`
- New struct `TimeSeriesEntityItem`
- New field `SystemData` in struct `ConfigData`
- New field `Duration` in struct `ConfigDataProperties`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `ResourceRecommendationBase`
- New field `SystemData` in struct `SuppressionContract`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).