# Release History

## 2.0.0 (2022-10-18)
### Breaking Changes

- Function `*RecommendationMetadataClient.Get` has been removed
- Struct `RecommendationMetadataClientGetOptions` has been removed
- Struct `RecommendationMetadataClientGetResponse` has been removed

### Features Added

- New const `DurationTwentyOne`
- New const `DurationFourteen`
- New const `DurationNinety`
- New const `DurationSeven`
- New const `DurationSixty`
- New const `PredictionTypePredictiveRightsizing`
- New const `DurationThirty`
- New type alias `PredictionType`
- New type alias `Duration`
- New function `NewTwelveRecommendationssMetadataClient(azcore.TokenCredential, *arm.ClientOptions) (*TwelveRecommendationssMetadataClient, error)`
- New function `*ManagementClient.Predict(context.Context, PredictionRequest, *ManagementClientPredictOptions) (ManagementClientPredictResponse, error)`
- New function `NewManagementClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagementClient, error)`
- New function `PossibleDurationValues() []Duration`
- New function `PossiblePredictionTypeValues() []PredictionType`
- New function `*TwelveRecommendationssMetadataClient.Get(context.Context, string, *TwelveRecommendationssMetadataClientGetOptions) (TwelveRecommendationssMetadataClientGetResponse, error)`
- New struct `ManagementClient`
- New struct `ManagementClientPredictOptions`
- New struct `ManagementClientPredictResponse`
- New struct `PredictionRequest`
- New struct `PredictionRequestProperties`
- New struct `PredictionResponse`
- New struct `PredictionResponseProperties`
- New struct `TwelveRecommendationssMetadataClient`
- New struct `TwelveRecommendationssMetadataClientGetOptions`
- New struct `TwelveRecommendationssMetadataClientGetResponse`
- New field `Duration` in struct `ConfigDataProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).