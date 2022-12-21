# Release History

## 2.0.0-beta.2 (2022-12-21)
### Breaking Changes

- Type alias `Operation` has been removed
- Struct `CloudError` has been removed
- Struct `MigrationRequestProperties` has been removed
- Field `MigrationRequest` of struct `ApplicationGroupProperties` has been removed
- Field `MigrationRequest` of struct `HostPoolProperties` has been removed

### Features Added

- New function `NewScalingPlanPooledSchedulesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ScalingPlanPooledSchedulesClient, error)`
- New function `*ScalingPlanPooledSchedulesClient.Create(context.Context, string, string, string, ScalingPlanPooledSchedule, *ScalingPlanPooledSchedulesClientCreateOptions) (ScalingPlanPooledSchedulesClientCreateResponse, error)`
- New function `*ScalingPlanPooledSchedulesClient.Delete(context.Context, string, string, string, *ScalingPlanPooledSchedulesClientDeleteOptions) (ScalingPlanPooledSchedulesClientDeleteResponse, error)`
- New function `*ScalingPlanPooledSchedulesClient.Get(context.Context, string, string, string, *ScalingPlanPooledSchedulesClientGetOptions) (ScalingPlanPooledSchedulesClientGetResponse, error)`
- New function `*ScalingPlanPooledSchedulesClient.NewListPager(string, string, *ScalingPlanPooledSchedulesClientListOptions) *runtime.Pager[ScalingPlanPooledSchedulesClientListResponse]`
- New function `*ScalingPlanPooledSchedulesClient.Update(context.Context, string, string, string, *ScalingPlanPooledSchedulesClientUpdateOptions) (ScalingPlanPooledSchedulesClientUpdateResponse, error)`
- New struct `ScalingPlanPooledSchedule`
- New struct `ScalingPlanPooledScheduleList`
- New struct `ScalingPlanPooledSchedulePatch`
- New struct `ScalingPlanPooledScheduleProperties`
- New struct `ScalingPlanPooledSchedulesClient`
- New struct `ScalingPlanPooledSchedulesClientListResponse`
- New field `InitialSkip` in struct `ApplicationGroupsClientListByResourceGroupOptions`
- New field `IsDescending` in struct `ApplicationGroupsClientListByResourceGroupOptions`
- New field `PageSize` in struct `ApplicationGroupsClientListByResourceGroupOptions`
- New field `InitialSkip` in struct `ApplicationsClientListOptions`
- New field `IsDescending` in struct `ApplicationsClientListOptions`
- New field `PageSize` in struct `ApplicationsClientListOptions`
- New field `InitialSkip` in struct `DesktopsClientListOptions`
- New field `IsDescending` in struct `DesktopsClientListOptions`
- New field `PageSize` in struct `DesktopsClientListOptions`
- New field `InitialSkip` in struct `HostPoolsClientListByResourceGroupOptions`
- New field `IsDescending` in struct `HostPoolsClientListByResourceGroupOptions`
- New field `PageSize` in struct `HostPoolsClientListByResourceGroupOptions`
- New field `InitialSkip` in struct `HostPoolsClientListOptions`
- New field `IsDescending` in struct `HostPoolsClientListOptions`
- New field `PageSize` in struct `HostPoolsClientListOptions`
- New field `InitialSkip` in struct `MSIXPackagesClientListOptions`
- New field `IsDescending` in struct `MSIXPackagesClientListOptions`
- New field `PageSize` in struct `MSIXPackagesClientListOptions`
- New field `InitialSkip` in struct `PrivateEndpointConnectionsClientListByHostPoolOptions`
- New field `IsDescending` in struct `PrivateEndpointConnectionsClientListByHostPoolOptions`
- New field `PageSize` in struct `PrivateEndpointConnectionsClientListByHostPoolOptions`
- New field `InitialSkip` in struct `PrivateLinkResourcesClientListByHostPoolOptions`
- New field `IsDescending` in struct `PrivateLinkResourcesClientListByHostPoolOptions`
- New field `PageSize` in struct `PrivateLinkResourcesClientListByHostPoolOptions`
- New field `InitialSkip` in struct `PrivateLinkResourcesClientListByWorkspaceOptions`
- New field `IsDescending` in struct `PrivateLinkResourcesClientListByWorkspaceOptions`
- New field `PageSize` in struct `PrivateLinkResourcesClientListByWorkspaceOptions`
- New field `InitialSkip` in struct `ScalingPlansClientListByHostPoolOptions`
- New field `IsDescending` in struct `ScalingPlansClientListByHostPoolOptions`
- New field `PageSize` in struct `ScalingPlansClientListByHostPoolOptions`
- New field `InitialSkip` in struct `ScalingPlansClientListByResourceGroupOptions`
- New field `IsDescending` in struct `ScalingPlansClientListByResourceGroupOptions`
- New field `PageSize` in struct `ScalingPlansClientListByResourceGroupOptions`
- New field `InitialSkip` in struct `ScalingPlansClientListBySubscriptionOptions`
- New field `IsDescending` in struct `ScalingPlansClientListBySubscriptionOptions`
- New field `PageSize` in struct `ScalingPlansClientListBySubscriptionOptions`
- New field `InitialSkip` in struct `SessionHostsClientListOptions`
- New field `IsDescending` in struct `SessionHostsClientListOptions`
- New field `PageSize` in struct `SessionHostsClientListOptions`
- New field `InitialSkip` in struct `StartMenuItemsClientListOptions`
- New field `IsDescending` in struct `StartMenuItemsClientListOptions`
- New field `PageSize` in struct `StartMenuItemsClientListOptions`
- New field `InitialSkip` in struct `UserSessionsClientListByHostPoolOptions`
- New field `IsDescending` in struct `UserSessionsClientListByHostPoolOptions`
- New field `PageSize` in struct `UserSessionsClientListByHostPoolOptions`
- New field `InitialSkip` in struct `UserSessionsClientListOptions`
- New field `IsDescending` in struct `UserSessionsClientListOptions`
- New field `PageSize` in struct `UserSessionsClientListOptions`
- New field `InitialSkip` in struct `WorkspacesClientListByResourceGroupOptions`
- New field `IsDescending` in struct `WorkspacesClientListByResourceGroupOptions`
- New field `PageSize` in struct `WorkspacesClientListByResourceGroupOptions`


## 2.0.0-beta.1 (2022-05-24)
### Breaking Changes

- Type of `ScalingSchedule.RampUpStartTime` has been changed from `*time.Time` to `*Time`
- Type of `ScalingSchedule.RampDownStartTime` has been changed from `*time.Time` to `*Time`
- Type of `ScalingSchedule.OffPeakStartTime` has been changed from `*time.Time` to `*Time`
- Type of `ScalingSchedule.PeakStartTime` has been changed from `*time.Time` to `*Time`
- Type of `ScalingPlanProperties.HostPoolType` has been changed from `*HostPoolType` to `*ScalingHostPoolType`
- Function `*DesktopsClient.List` has been removed
- Function `*ScalingSchedule.UnmarshalJSON` has been removed
- Function `*OperationsClient.List` has been removed
- Field `HostPoolType` of struct `ScalingPlanPatchProperties` has been removed

### Features Added

- New const `PrivateEndpointServiceConnectionStatusRejected`
- New const `HostpoolPublicNetworkAccessEnabledForClientsOnly`
- New const `DayOfWeekMonday`
- New const `CreatedByTypeApplication`
- New const `HostpoolPublicNetworkAccessEnabled`
- New const `PublicNetworkAccessDisabled`
- New const `DayOfWeekSunday`
- New const `DayOfWeekSaturday`
- New const `DayOfWeekFriday`
- New const `ScalingHostPoolTypePooled`
- New const `PrivateEndpointConnectionProvisioningStateDeleting`
- New const `SessionHostComponentUpdateTypeScheduled`
- New const `PrivateEndpointConnectionProvisioningStateSucceeded`
- New const `PrivateEndpointServiceConnectionStatusApproved`
- New const `PrivateEndpointConnectionProvisioningStateCreating`
- New const `PrivateEndpointConnectionProvisioningStateFailed`
- New const `HostpoolPublicNetworkAccessEnabledForSessionHostsOnly`
- New const `DayOfWeekThursday`
- New const `CreatedByTypeKey`
- New const `PrivateEndpointServiceConnectionStatusPending`
- New const `SessionHostComponentUpdateTypeDefault`
- New const `DayOfWeekTuesday`
- New const `HostpoolPublicNetworkAccessDisabled`
- New const `PublicNetworkAccessEnabled`
- New const `CreatedByTypeManagedIdentity`
- New const `DayOfWeekWednesday`
- New const `CreatedByTypeUser`
- New function `PossibleDayOfWeekValues() []DayOfWeek`
- New function `*DesktopsClient.NewListPager(string, string, *DesktopsClientListOptions) *runtime.Pager[DesktopsClientListResponse]`
- New function `PossibleCreatedByTypeValues() []CreatedByType`
- New function `AgentUpdatePatchProperties.MarshalJSON() ([]byte, error)`
- New function `PossiblePublicNetworkAccessValues() []PublicNetworkAccess`
- New function `PossibleScalingHostPoolTypeValues() []ScalingHostPoolType`
- New function `PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus`
- New function `AgentUpdateProperties.MarshalJSON() ([]byte, error)`
- New function `PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState`
- New function `SystemData.MarshalJSON() ([]byte, error)`
- New function `*SystemData.UnmarshalJSON([]byte) error`
- New function `*OperationsClient.NewListPager(*OperationsClientListOptions) *runtime.Pager[OperationsClientListResponse]`
- New function `PossibleSessionHostComponentUpdateTypeValues() []SessionHostComponentUpdateType`
- New function `PrivateLinkResourceProperties.MarshalJSON() ([]byte, error)`
- New function `PossibleHostpoolPublicNetworkAccessValues() []HostpoolPublicNetworkAccess`
- New struct `AgentUpdatePatchProperties`
- New struct `AgentUpdateProperties`
- New struct `MaintenanceWindowPatchProperties`
- New struct `MaintenanceWindowProperties`
- New struct `PrivateEndpoint`
- New struct `PrivateEndpointConnection`
- New struct `PrivateEndpointConnectionListResultWithSystemData`
- New struct `PrivateEndpointConnectionProperties`
- New struct `PrivateEndpointConnectionWithSystemData`
- New struct `PrivateEndpointConnectionsClientDeleteByHostPoolOptions`
- New struct `PrivateEndpointConnectionsClientDeleteByHostPoolResponse`
- New struct `PrivateEndpointConnectionsClientDeleteByWorkspaceOptions`
- New struct `PrivateEndpointConnectionsClientDeleteByWorkspaceResponse`
- New struct `PrivateEndpointConnectionsClientGetByHostPoolOptions`
- New struct `PrivateEndpointConnectionsClientGetByHostPoolResponse`
- New struct `PrivateEndpointConnectionsClientGetByWorkspaceOptions`
- New struct `PrivateEndpointConnectionsClientGetByWorkspaceResponse`
- New struct `PrivateEndpointConnectionsClientListByHostPoolOptions`
- New struct `PrivateEndpointConnectionsClientListByHostPoolResponse`
- New struct `PrivateEndpointConnectionsClientListByWorkspaceOptions`
- New struct `PrivateEndpointConnectionsClientListByWorkspaceResponse`
- New struct `PrivateEndpointConnectionsClientUpdateByHostPoolOptions`
- New struct `PrivateEndpointConnectionsClientUpdateByHostPoolResponse`
- New struct `PrivateEndpointConnectionsClientUpdateByWorkspaceOptions`
- New struct `PrivateEndpointConnectionsClientUpdateByWorkspaceResponse`
- New struct `PrivateLinkResource`
- New struct `PrivateLinkResourceListResult`
- New struct `PrivateLinkResourceProperties`
- New struct `PrivateLinkResourcesClientListByHostPoolOptions`
- New struct `PrivateLinkResourcesClientListByHostPoolResponse`
- New struct `PrivateLinkResourcesClientListByWorkspaceOptions`
- New struct `PrivateLinkResourcesClientListByWorkspaceResponse`
- New struct `PrivateLinkServiceConnectionState`
- New struct `SystemData`
- New struct `Time`
- New field `FriendlyName` in struct `SessionHostProperties`
- New field `SystemData` in struct `UserSession`
- New field `PublicNetworkAccess` in struct `HostPoolPatchProperties`
- New field `AgentUpdate` in struct `HostPoolPatchProperties`
- New field `NextLink` in struct `ResourceProviderOperationList`
- New field `SystemData` in struct `SessionHost`
- New field `SystemData` in struct `MSIXPackage`
- New field `PrivateEndpointConnections` in struct `WorkspaceProperties`
- New field `PublicNetworkAccess` in struct `WorkspaceProperties`
- New field `SystemData` in struct `Workspace`
- New field `SystemData` in struct `Application`
- New field `SystemData` in struct `ScalingPlan`
- New field `PublicNetworkAccess` in struct `HostPoolProperties`
- New field `PrivateEndpointConnections` in struct `HostPoolProperties`
- New field `AgentUpdate` in struct `HostPoolProperties`
- New field `PublicNetworkAccess` in struct `WorkspacePatchProperties`
- New field `Force` in struct `SessionHostsClientUpdateOptions`
- New field `FriendlyName` in struct `SessionHostPatchProperties`
- New field `SystemData` in struct `ApplicationGroup`
- New field `SystemData` in struct `Desktop`
- New field `SystemData` in struct `HostPool`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).