# Release History

## 1.3.0 (2024-10-04)
### Features Added

- New enum type `Aggregated` with values `AggregatedDay`, `AggregatedMonth`, `AggregatedWeek`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `Duration` with values `DurationFourteen`, `DurationNinety`, `DurationSeven`, `DurationSixty`, `DurationThirty`, `DurationTwentyOne`
- New enum type `PredictionType` with values `PredictionTypePredictiveRightsizing`
- New function `*ClientFactory.NewManagementClient() *ManagementClient`
- New function `*ClientFactory.NewScoresClient() *ScoresClient`
- New function `NewManagementClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagementClient, error)`
- New function `*ManagementClient.Predict(context.Context, PredictionRequest, *ManagementClientPredictOptions) (ManagementClientPredictResponse, error)`
- New function `NewScoresClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ScoresClient, error)`
- New function `*ScoresClient.Get(context.Context, string, *ScoresClientGetOptions) (ScoresClientGetResponse, error)`
- New function `*ScoresClient.List(context.Context, *ScoresClientListOptions) (ScoresClientListResponse, error)`
- New struct `PredictionRequest`
- New struct `PredictionRequestProperties`
- New struct `PredictionResponse`
- New struct `PredictionResponseProperties`
- New struct `ProxyResource`
- New struct `ScoreEntity`
- New struct `ScoreEntityForAdvisor`
- New struct `ScoreEntityForAdvisorProperties`
- New struct `ScoreResponse`
- New struct `SystemData`
- New struct `TimeSeriesEntityItem`
- New field `SystemData` in struct `ConfigData`
- New field `Duration` in struct `ConfigDataProperties`
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