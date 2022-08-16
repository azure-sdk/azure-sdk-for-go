# Release History

## 0.8.0 (2022-08-16)
### Breaking Changes

- Function `*ScheduledQueryRulesClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, LogSearchRuleResource, *ScheduledQueryRulesClientCreateOrUpdateOptions)` to `(context.Context, string, string, ScheduledQueryRuleResource, *ScheduledQueryRulesClientCreateOrUpdateOptions)`
- Function `*ScheduledQueryRulesClient.Update` parameter(s) have been changed from `(context.Context, string, string, LogSearchRuleResourcePatch, *ScheduledQueryRulesClientUpdateOptions)` to `(context.Context, string, string, ScheduledQueryRuleResourcePatch, *ScheduledQueryRulesClientUpdateOptions)`
- Type of `OperationStatus.Error` has been changed from `*ErrorResponseCommon` to `*ErrorDetail`
- Type of `PrivateEndpointConnectionProperties.ProvisioningState` has been changed from `*string` to `*PrivateEndpointConnectionProvisioningState`
- Type of `PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState` has been changed from `*PrivateLinkServiceConnectionStateProperty` to `*PrivateLinkServiceConnectionState`
- Type of `PrivateEndpointConnectionProperties.PrivateEndpoint` has been changed from `*PrivateEndpointProperty` to `*PrivateEndpoint`
- Type of `Dimension.Operator` has been changed from `*Operator` to `*DimensionOperator`
- Type of `ErrorContract.Error` has been changed from `*ErrorResponse` to `*ErrorResponseDetails`
- Const `MetricTriggerTypeTotal` has been removed
- Const `MetricTriggerTypeConsecutive` has been removed
- Const `ConditionalOperatorEqual` has been removed
- Const `QueryTypeResultCount` has been removed
- Const `ConditionalOperatorLessThan` has been removed
- Const `ConditionalOperatorGreaterThanOrEqual` has been removed
- Const `ProvisioningStateCanceled` has been removed
- Const `EnabledTrue` has been removed
- Const `ProvisioningStateSucceeded` has been removed
- Const `ConditionalOperatorLessThanOrEqual` has been removed
- Const `ProvisioningStateDeploying` has been removed
- Const `OperatorInclude` has been removed
- Const `ConditionalOperatorGreaterThan` has been removed
- Const `ProvisioningStateFailed` has been removed
- Const `EnabledFalse` has been removed
- Function `*LogToMetricAction.GetAction` has been removed
- Function `*DiagnosticSettingsClient.List` has been removed
- Function `PossibleMetricTriggerTypeValues` has been removed
- Function `PossibleConditionalOperatorValues` has been removed
- Function `PossibleQueryTypeValues` has been removed
- Function `*PrivateLinkResourcesClient.NewListByPrivateLinkScopePager` has been removed
- Function `*Action.GetAction` has been removed
- Function `*DiagnosticSettingsCategoryClient.List` has been removed
- Function `PossibleProvisioningStateValues` has been removed
- Function `PossibleEnabledValues` has been removed
- Function `*AlertingAction.GetAction` has been removed
- Function `*PrivateEndpointConnectionsClient.NewListByPrivateLinkScopePager` has been removed
- Struct `Action` has been removed
- Struct `AlertingAction` has been removed
- Struct `AzNsActionGroup` has been removed
- Struct `Criteria` has been removed
- Struct `ErrorResponseCommon` has been removed
- Struct `LogMetricTrigger` has been removed
- Struct `LogSearchRule` has been removed
- Struct `LogSearchRulePatch` has been removed
- Struct `LogSearchRuleResource` has been removed
- Struct `LogSearchRuleResourceCollection` has been removed
- Struct `LogSearchRuleResourcePatch` has been removed
- Struct `LogToMetricAction` has been removed
- Struct `PrivateEndpointProperty` has been removed
- Struct `PrivateLinkScopesResource` has been removed
- Struct `PrivateLinkServiceConnectionStateProperty` has been removed
- Struct `ProxyOnlyResource` has been removed
- Struct `Schedule` has been removed
- Struct `Source` has been removed
- Struct `TrackedEntityResource` has been removed
- Struct `TriggerCondition` has been removed
- Field `Filter` of struct `ScheduledQueryRulesClientListByResourceGroupOptions` has been removed
- Field `LogSearchRuleResource` of struct `ScheduledQueryRulesClientUpdateResponse` has been removed
- Field `LogSearchRuleResourceCollection` of struct `ScheduledQueryRulesClientListBySubscriptionResponse` has been removed
- Field `LogSearchRuleResource` of struct `ScheduledQueryRulesClientCreateOrUpdateResponse` has been removed
- Field `LogSearchRuleResource` of struct `ScheduledQueryRulesClientGetResponse` has been removed
- Field `Identity` of struct `AzureResource` has been removed
- Field `Kind` of struct `AzureResource` has been removed
- Field `LogSearchRuleResourceCollection` of struct `ScheduledQueryRulesClientListByResourceGroupResponse` has been removed
- Field `Identity` of struct `ActionGroupResource` has been removed
- Field `Kind` of struct `ActionGroupResource` has been removed
- Field `NextLink` of struct `PrivateLinkResourceListResult` has been removed
- Field `NextLink` of struct `PrivateEndpointConnectionListResult` has been removed
- Field `Filter` of struct `ScheduledQueryRulesClientListBySubscriptionOptions` has been removed

### Features Added

- New const `PrivateEndpointConnectionProvisioningStateDeleting`
- New const `PredictiveAutoscalePolicyScaleModeDisabled`
- New const `TimeAggregationCount`
- New const `DimensionOperatorExclude`
- New const `KindLogAlert`
- New const `TimeAggregationTotal`
- New const `PredictiveAutoscalePolicyScaleModeEnabled`
- New const `TimeAggregationMaximum`
- New const `AccessModePrivateOnly`
- New const `PrivateEndpointServiceConnectionStatusApproved`
- New const `PrivateEndpointServiceConnectionStatusPending`
- New const `KindLogToMetric`
- New const `PrivateEndpointConnectionProvisioningStateSucceeded`
- New const `TimeAggregationMinimum`
- New const `PrivateEndpointServiceConnectionStatusRejected`
- New const `TimeAggregationAverage`
- New const `ConditionOperatorEquals`
- New const `DimensionOperatorInclude`
- New const `PredictiveAutoscalePolicyScaleModeForecastOnly`
- New const `AccessModeOpen`
- New const `PrivateEndpointConnectionProvisioningStateCreating`
- New const `PrivateEndpointConnectionProvisioningStateFailed`
- New function `*DiagnosticSettingsCategoryClient.NewListPager(string, *DiagnosticSettingsCategoryClientListOptions) *runtime.Pager[DiagnosticSettingsCategoryClientListResponse]`
- New function `PossibleDimensionOperatorValues() []DimensionOperator`
- New function `PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState`
- New function `*PrivateLinkResourcesClient.ListByPrivateLinkScope(context.Context, string, string, *PrivateLinkResourcesClientListByPrivateLinkScopeOptions) (PrivateLinkResourcesClientListByPrivateLinkScopeResponse, error)`
- New function `PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus`
- New function `PossiblePredictiveAutoscalePolicyScaleModeValues() []PredictiveAutoscalePolicyScaleMode`
- New function `PossibleTimeAggregationValues() []TimeAggregation`
- New function `NewPredictiveMetricClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PredictiveMetricClient, error)`
- New function `*PredictiveMetricClient.Get(context.Context, string, string, string, string, string, string, string, *PredictiveMetricClientGetOptions) (PredictiveMetricClientGetResponse, error)`
- New function `PossibleAccessModeValues() []AccessMode`
- New function `*DiagnosticSettingsClient.NewListPager(string, *DiagnosticSettingsClientListOptions) *runtime.Pager[DiagnosticSettingsClientListResponse]`
- New function `*PrivateEndpointConnectionsClient.ListByPrivateLinkScope(context.Context, string, string, *PrivateEndpointConnectionsClientListByPrivateLinkScopeOptions) (PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse, error)`
- New function `PossibleKindValues() []Kind`
- New struct `AccessModeSettings`
- New struct `AccessModeSettingsExclusion`
- New struct `Actions`
- New struct `AutoscaleErrorResponse`
- New struct `AutoscaleErrorResponseError`
- New struct `Condition`
- New struct `ConditionFailingPeriods`
- New struct `DefaultErrorResponse`
- New struct `ErrorResponseAdditionalInfo`
- New struct `ErrorResponseDetails`
- New struct `PredictiveAutoscalePolicy`
- New struct `PredictiveMetricClient`
- New struct `PredictiveMetricClientGetOptions`
- New struct `PredictiveMetricClientGetResponse`
- New struct `PredictiveResponse`
- New struct `PredictiveValue`
- New struct `PrivateEndpoint`
- New struct `PrivateLinkServiceConnectionState`
- New struct `ProxyResourceAutoGenerated`
- New struct `ResourceAutoGenerated`
- New struct `ResourceAutoGenerated2`
- New struct `ResourceAutoGenerated3`
- New struct `ResourceAutoGenerated4`
- New struct `ScheduledQueryRuleCriteria`
- New struct `ScheduledQueryRuleProperties`
- New struct `ScheduledQueryRuleResource`
- New struct `ScheduledQueryRuleResourceCollection`
- New struct `ScheduledQueryRuleResourcePatch`
- New struct `TrackedResource`
- New field `SystemData` in struct `DiagnosticSettingsResource`
- New anonymous field `TestNotificationDetailsResponse` in struct `ActionGroupsClientCreateNotificationsAtActionGroupResourceLevelResponse`
- New field `AccessModeSettings` in struct `AzureMonitorPrivateLinkScopeProperties`
- New field `CategoryGroups` in struct `DiagnosticSettingsCategory`
- New field `PredictiveAutoscalePolicy` in struct `AutoscaleSetting`
- New field `SystemData` in struct `AutoscaleSettingResource`
- New anonymous field `ScheduledQueryRuleResourceCollection` in struct `ScheduledQueryRulesClientListBySubscriptionResponse`
- New field `SystemData` in struct `DiagnosticSettingsCategoryResource`
- New field `CategoryGroup` in struct `LogSettings`
- New anonymous field `ScheduledQueryRuleResource` in struct `ScheduledQueryRulesClientGetResponse`
- New field `RequiredZoneNames` in struct `PrivateLinkResourceProperties`
- New field `MarketplacePartnerID` in struct `DiagnosticSettings`
- New anonymous field `TestNotificationDetailsResponse` in struct `ActionGroupsClientPostTestNotificationsResponse`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `AzureMonitorPrivateLinkScope`
- New anonymous field `ScheduledQueryRuleResource` in struct `ScheduledQueryRulesClientUpdateResponse`
- New anonymous field `ScheduledQueryRuleResourceCollection` in struct `ScheduledQueryRulesClientListByResourceGroupResponse`
- New field `SystemData` in struct `ScopedResource`
- New anonymous field `ScheduledQueryRuleResource` in struct `ScheduledQueryRulesClientCreateOrUpdateResponse`
- New anonymous field `TestNotificationDetailsResponse` in struct `ActionGroupsClientCreateNotificationsAtResourceGroupLevelResponse`


## 0.7.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.7.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).