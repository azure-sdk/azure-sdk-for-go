# Release History

## 0.8.0 (2022-08-11)
### Breaking Changes

- Type of `OperationStatus.Error` has been changed from `*ErrorResponseCommon` to `*ErrorDetail`
- Type of `PrivateEndpointConnectionProperties.PrivateEndpoint` has been changed from `*PrivateEndpointProperty` to `*PrivateEndpoint`
- Type of `PrivateEndpointConnectionProperties.ProvisioningState` has been changed from `*string` to `*PrivateEndpointConnectionProvisioningState`
- Type of `PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState` has been changed from `*PrivateLinkServiceConnectionStateProperty` to `*PrivateLinkServiceConnectionState`
- Function `*PrivateEndpointConnectionsClient.NewListByPrivateLinkScopePager` has been removed
- Function `*PrivateLinkResourcesClient.NewListByPrivateLinkScopePager` has been removed
- Function `*DiagnosticSettingsCategoryClient.List` has been removed
- Function `*DiagnosticSettingsClient.List` has been removed
- Struct `ErrorResponseCommon` has been removed
- Struct `PrivateEndpointProperty` has been removed
- Struct `PrivateLinkScopesResource` has been removed
- Struct `PrivateLinkServiceConnectionStateProperty` has been removed
- Struct `ProxyOnlyResource` has been removed
- Field `Kind` of struct `ActionGroupResource` has been removed
- Field `Identity` of struct `ActionGroupResource` has been removed
- Field `Identity` of struct `AzureResource` has been removed
- Field `Kind` of struct `AzureResource` has been removed
- Field `NextLink` of struct `PrivateLinkResourceListResult` has been removed
- Field `NextLink` of struct `PrivateEndpointConnectionListResult` has been removed

### Features Added

- New const `PredictiveAutoscalePolicyScaleModeDisabled`
- New const `PredictiveAutoscalePolicyScaleModeForecastOnly`
- New const `PrivateEndpointServiceConnectionStatusRejected`
- New const `OriginUserSystem`
- New const `PrivateEndpointConnectionProvisioningStateCreating`
- New const `ProvisioningStateCreating`
- New const `OriginSystem`
- New const `PrivateEndpointConnectionProvisioningStateSucceeded`
- New const `PrivateEndpointServiceConnectionStatusPending`
- New const `OriginUser`
- New const `ActionTypeInternal`
- New const `PrivateEndpointServiceConnectionStatusApproved`
- New const `AccessModePrivateOnly`
- New const `PrivateEndpointConnectionProvisioningStateFailed`
- New const `PredictiveAutoscalePolicyScaleModeEnabled`
- New const `PrivateEndpointConnectionProvisioningStateDeleting`
- New const `ProvisioningStateDeleting`
- New const `AccessModeOpen`
- New function `*DiagnosticSettingsClient.NewListPager(string, *DiagnosticSettingsClientListOptions) *runtime.Pager[DiagnosticSettingsClientListResponse]`
- New function `*MonitoringAccountsClient.Create(context.Context, string, string, MonitoringAccountResource, *MonitoringAccountsClientCreateOptions) (MonitoringAccountsClientCreateResponse, error)`
- New function `PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus`
- New function `NewClient(azcore.TokenCredential, *arm.ClientOptions) (*Client, error)`
- New function `*DiagnosticSettingsCategoryClient.NewListPager(string, *DiagnosticSettingsCategoryClientListOptions) *runtime.Pager[DiagnosticSettingsCategoryClientListResponse]`
- New function `*MonitoringAccountsClient.NewListBySubscriptionPager(*MonitoringAccountsClientListBySubscriptionOptions) *runtime.Pager[MonitoringAccountsClientListBySubscriptionResponse]`
- New function `*PredictiveMetricClient.Get(context.Context, string, string, string, string, string, string, string, *PredictiveMetricClientGetOptions) (PredictiveMetricClientGetResponse, error)`
- New function `*MonitoringAccountsClient.Get(context.Context, string, string, *MonitoringAccountsClientGetOptions) (MonitoringAccountsClientGetResponse, error)`
- New function `NewMonitoringAccountsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*MonitoringAccountsClient, error)`
- New function `*MonitoringAccountsClient.NewListByResourceGroupPager(string, *MonitoringAccountsClientListByResourceGroupOptions) *runtime.Pager[MonitoringAccountsClientListByResourceGroupResponse]`
- New function `PossibleActionTypeValues() []ActionType`
- New function `PossiblePredictiveAutoscalePolicyScaleModeValues() []PredictiveAutoscalePolicyScaleMode`
- New function `PossibleAccessModeValues() []AccessMode`
- New function `*PrivateLinkResourcesClient.ListByPrivateLinkScope(context.Context, string, string, *PrivateLinkResourcesClientListByPrivateLinkScopeOptions) (PrivateLinkResourcesClientListByPrivateLinkScopeResponse, error)`
- New function `*Client.NewOperationsListPager(*ClientOperationsListOptions) *runtime.Pager[ClientOperationsListResponse]`
- New function `PossibleOriginValues() []Origin`
- New function `PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState`
- New function `*MonitoringAccountsClient.Delete(context.Context, string, string, *MonitoringAccountsClientDeleteOptions) (MonitoringAccountsClientDeleteResponse, error)`
- New function `*PrivateEndpointConnectionsClient.ListByPrivateLinkScope(context.Context, string, string, *PrivateEndpointConnectionsClientListByPrivateLinkScopeOptions) (PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse, error)`
- New function `*MonitoringAccountsClient.Update(context.Context, string, string, *MonitoringAccountsClientUpdateOptions) (MonitoringAccountsClientUpdateResponse, error)`
- New function `NewPredictiveMetricClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PredictiveMetricClient, error)`
- New struct `AccessModeSettings`
- New struct `AccessModeSettingsExclusion`
- New struct `AutoscaleErrorResponse`
- New struct `AutoscaleErrorResponseError`
- New struct `Client`
- New struct `ClientOperationsListOptions`
- New struct `ClientOperationsListResponse`
- New struct `DefaultErrorResponse`
- New struct `ErrorDetailAutoGenerated`
- New struct `ErrorResponseAutoGenerated2`
- New struct `IngestionSettings`
- New struct `Metrics`
- New struct `MonitoringAccount`
- New struct `MonitoringAccountDefaultIngestionSettings`
- New struct `MonitoringAccountMetrics`
- New struct `MonitoringAccountResource`
- New struct `MonitoringAccountResourceForUpdate`
- New struct `MonitoringAccountResourceListResult`
- New struct `MonitoringAccountResourceProperties`
- New struct `MonitoringAccountsClient`
- New struct `MonitoringAccountsClientCreateOptions`
- New struct `MonitoringAccountsClientCreateResponse`
- New struct `MonitoringAccountsClientDeleteOptions`
- New struct `MonitoringAccountsClientDeleteResponse`
- New struct `MonitoringAccountsClientGetOptions`
- New struct `MonitoringAccountsClientGetResponse`
- New struct `MonitoringAccountsClientListByResourceGroupOptions`
- New struct `MonitoringAccountsClientListByResourceGroupResponse`
- New struct `MonitoringAccountsClientListBySubscriptionOptions`
- New struct `MonitoringAccountsClientListBySubscriptionResponse`
- New struct `MonitoringAccountsClientUpdateOptions`
- New struct `MonitoringAccountsClientUpdateResponse`
- New struct `OperationAutoGenerated`
- New struct `OperationDisplayAutoGenerated`
- New struct `OperationListResultAutoGenerated`
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
- New struct `ResourceAutoGenerated5`
- New struct `TrackedResource`
- New struct `TrackedResourceAutoGenerated`
- New field `AccessModeSettings` in struct `AzureMonitorPrivateLinkScopeProperties`
- New field `SystemData` in struct `ScopedResource`
- New field `CategoryGroups` in struct `DiagnosticSettingsCategory`
- New anonymous field `TestNotificationDetailsResponse` in struct `ActionGroupsClientCreateNotificationsAtActionGroupResourceLevelResponse`
- New field `SystemData` in struct `DiagnosticSettingsCategoryResource`
- New field `SystemData` in struct `AzureMonitorPrivateLinkScope`
- New anonymous field `TestNotificationDetailsResponse` in struct `ActionGroupsClientPostTestNotificationsResponse`
- New field `SystemData` in struct `DiagnosticSettingsResource`
- New field `RequiredZoneNames` in struct `PrivateLinkResourceProperties`
- New field `MarketplacePartnerID` in struct `DiagnosticSettings`
- New field `CategoryGroup` in struct `LogSettings`
- New anonymous field `TestNotificationDetailsResponse` in struct `ActionGroupsClientCreateNotificationsAtResourceGroupLevelResponse`
- New field `PredictiveAutoscalePolicy` in struct `AutoscaleSetting`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `AutoscaleSettingResource`


## 0.7.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.7.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).