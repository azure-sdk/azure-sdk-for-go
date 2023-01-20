# Release History

## 2.0.0 (2023-01-20)
### Breaking Changes

- Type of `RecommendationProperties.Metadata` has been changed from `map[string]interface{}` to `map[string]any`
- Type of `RecommendationProperties.ExposedMetadataProperties` has been changed from `map[string]interface{}` to `map[string]any`
- Type of `RecommendationProperties.Remediation` has been changed from `map[string]interface{}` to `map[string]any`
- Type of `RecommendationProperties.Actions` has been changed from `[]map[string]interface{}` to `[]map[string]any`
- Type of `ResourceMetadata.Action` has been changed from `map[string]interface{}` to `map[string]any`

### Features Added

- New type alias `Duration` with values `DurationFourteen`, `DurationNinety`, `DurationSeven`, `DurationSixty`, `DurationThirty`, `DurationTwentyOne`
- New type alias `PredictionType` with values `PredictionTypePredictiveRightsizing`
- New function `NewManagementClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagementClient, error)`
- New function `*ManagementClient.Predict(context.Context, PredictionRequest, *ManagementClientPredictOptions) (ManagementClientPredictResponse, error)`
- New struct `ManagementClient`
- New struct `PredictionRequest`
- New struct `PredictionRequestProperties`
- New struct `PredictionResponse`
- New struct `PredictionResponseProperties`
- New field `Duration` in struct `ConfigDataProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).