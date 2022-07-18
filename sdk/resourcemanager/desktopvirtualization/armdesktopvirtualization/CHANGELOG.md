# Release History

## 2.0.0-beta.2 (2022-07-18)
### Breaking Changes

- Function `*SessionHostsClient.Update` parameter(s) have been changed from `(context.Context, string, string, string, *SessionHostsClientUpdateOptions)` to `(context.Context, string, string, string, SessionHostPatch, *SessionHostsClientUpdateOptions)`
- Function `*HostPoolsClient.Update` parameter(s) have been changed from `(context.Context, string, string, *HostPoolsClientUpdateOptions)` to `(context.Context, string, string, HostPoolPatch, *HostPoolsClientUpdateOptions)`
- Function `*MSIXPackagesClient.Update` parameter(s) have been changed from `(context.Context, string, string, string, *MSIXPackagesClientUpdateOptions)` to `(context.Context, string, string, string, MSIXPackagePatch, *MSIXPackagesClientUpdateOptions)`
- Function `*ScalingPlansClient.Update` parameter(s) have been changed from `(context.Context, string, string, *ScalingPlansClientUpdateOptions)` to `(context.Context, string, string, ScalingPlanPatch, *ScalingPlansClientUpdateOptions)`
- Function `*ApplicationsClient.Update` parameter(s) have been changed from `(context.Context, string, string, string, *ApplicationsClientUpdateOptions)` to `(context.Context, string, string, string, ApplicationPatch, *ApplicationsClientUpdateOptions)`
- Function `*WorkspacesClient.Update` parameter(s) have been changed from `(context.Context, string, string, *WorkspacesClientUpdateOptions)` to `(context.Context, string, string, WorkspacePatch, *WorkspacesClientUpdateOptions)`
- Function `*ApplicationGroupsClient.Update` parameter(s) have been changed from `(context.Context, string, string, *ApplicationGroupsClientUpdateOptions)` to `(context.Context, string, string, ApplicationGroupPatch, *ApplicationGroupsClientUpdateOptions)`
- Function `*DesktopsClient.Update` parameter(s) have been changed from `(context.Context, string, string, string, *DesktopsClientUpdateOptions)` to `(context.Context, string, string, string, DesktopPatch, *DesktopsClientUpdateOptions)`
- Const `ScalingScheduleDaysOfWeekItemMonday` has been removed
- Const `ScalingScheduleDaysOfWeekItemSaturday` has been removed
- Const `ScalingScheduleDaysOfWeekItemThursday` has been removed
- Const `ScalingScheduleDaysOfWeekItemSunday` has been removed
- Const `ScalingScheduleDaysOfWeekItemWednesday` has been removed
- Const `ScalingScheduleDaysOfWeekItemTuesday` has been removed
- Const `ScalingScheduleDaysOfWeekItemFriday` has been removed
- Function `PossibleScalingScheduleDaysOfWeekItemValues` has been removed
- Struct `ScalingSchedule` has been removed
- Field `Schedules` of struct `ScalingPlanProperties` has been removed
- Field `Desktop` of struct `DesktopsClientUpdateOptions` has been removed
- Field `HostPool` of struct `HostPoolsClientUpdateOptions` has been removed
- Field `Application` of struct `ApplicationsClientUpdateOptions` has been removed
- Field `SessionHost` of struct `SessionHostsClientUpdateOptions` has been removed
- Field `MsixPackage` of struct `MSIXPackagesClientUpdateOptions` has been removed
- Field `ApplicationGroup` of struct `ApplicationGroupsClientUpdateOptions` has been removed
- Field `Workspace` of struct `WorkspacesClientUpdateOptions` has been removed
- Field `ScalingPlan` of struct `ScalingPlansClientUpdateOptions` has been removed
- Field `Schedules` of struct `ScalingPlanPatchProperties` has been removed

### Features Added

- New function `*ScalingPlanPooledSchedulesClient.Get(context.Context, string, string, string, *ScalingPlanPooledSchedulesClientGetOptions) (ScalingPlanPooledSchedulesClientGetResponse, error)`
- New function `NewScalingPlanPooledSchedulesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ScalingPlanPooledSchedulesClient, error)`
- New function `*ScalingPlanPooledSchedulesClient.Delete(context.Context, string, string, string, *ScalingPlanPooledSchedulesClientDeleteOptions) (ScalingPlanPooledSchedulesClientDeleteResponse, error)`
- New function `*ScalingPlanPooledSchedulesClient.NewListPager(string, string, *ScalingPlanPooledSchedulesClientListOptions) *runtime.Pager[ScalingPlanPooledSchedulesClientListResponse]`
- New function `*ScalingPlanPooledSchedulesClient.Update(context.Context, string, string, string, ScalingPlanPooledSchedulePatch, *ScalingPlanPooledSchedulesClientUpdateOptions) (ScalingPlanPooledSchedulesClientUpdateResponse, error)`
- New function `*ScalingPlanPooledSchedulesClient.Create(context.Context, string, string, string, ScalingPlanPooledSchedule, *ScalingPlanPooledSchedulesClientCreateOptions) (ScalingPlanPooledSchedulesClientCreateResponse, error)`
- New struct `ScalingPlanPooledSchedule`
- New struct `ScalingPlanPooledScheduleList`
- New struct `ScalingPlanPooledSchedulePatch`
- New struct `ScalingPlanPooledScheduleProperties`
- New struct `ScalingPlanPooledSchedulesClient`
- New struct `ScalingPlanPooledSchedulesClientCreateOptions`
- New struct `ScalingPlanPooledSchedulesClientCreateResponse`
- New struct `ScalingPlanPooledSchedulesClientDeleteOptions`
- New struct `ScalingPlanPooledSchedulesClientDeleteResponse`
- New struct `ScalingPlanPooledSchedulesClientGetOptions`
- New struct `ScalingPlanPooledSchedulesClientGetResponse`
- New struct `ScalingPlanPooledSchedulesClientListOptions`
- New struct `ScalingPlanPooledSchedulesClientListResponse`
- New struct `ScalingPlanPooledSchedulesClientUpdateOptions`
- New struct `ScalingPlanPooledSchedulesClientUpdateResponse`


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