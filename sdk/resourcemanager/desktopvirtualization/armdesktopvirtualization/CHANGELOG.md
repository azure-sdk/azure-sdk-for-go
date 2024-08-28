# Release History

## 3.0.0-beta.1 (2024-08-28)
### Breaking Changes

- Type of `ApplicationGroup.Identity` has been changed from `*ResourceModelWithAllowedPropertySetIdentity` to `*ManagedServiceIdentity`
- Type of `ApplicationGroup.Plan` has been changed from `*ResourceModelWithAllowedPropertySetPlan` to `*Plan`
- Type of `ApplicationGroup.SKU` has been changed from `*ResourceModelWithAllowedPropertySetSKU` to `*SKU`
- Type of `HostPool.Identity` has been changed from `*ResourceModelWithAllowedPropertySetIdentity` to `*ManagedServiceIdentity`
- Type of `HostPool.Plan` has been changed from `*ResourceModelWithAllowedPropertySetPlan` to `*Plan`
- Type of `HostPool.SKU` has been changed from `*ResourceModelWithAllowedPropertySetSKU` to `*SKU`
- Type of `ScalingPlan.Identity` has been changed from `*ResourceModelWithAllowedPropertySetIdentity` to `*ManagedServiceIdentity`
- Type of `ScalingPlan.Plan` has been changed from `*ResourceModelWithAllowedPropertySetPlan` to `*Plan`
- Type of `ScalingPlan.SKU` has been changed from `*ResourceModelWithAllowedPropertySetSKU` to `*SKU`
- Type of `Workspace.Identity` has been changed from `*ResourceModelWithAllowedPropertySetIdentity` to `*ManagedServiceIdentity`
- Type of `Workspace.Plan` has been changed from `*ResourceModelWithAllowedPropertySetPlan` to `*Plan`
- Type of `Workspace.SKU` has been changed from `*ResourceModelWithAllowedPropertySetSKU` to `*SKU`
- Struct `ResourceModelWithAllowedPropertySetIdentity` has been removed
- Struct `ResourceModelWithAllowedPropertySetPlan` has been removed
- Struct `ResourceModelWithAllowedPropertySetSKU` has been removed

### Features Added

- New value `LoadBalancerTypeMultiplePersistent` added to enum type `LoadBalancerType`
- New value `ScalingHostPoolTypePersonal` added to enum type `ScalingHostPoolType`
- New enum type `AppAttachPackageArchitectures` with values `AppAttachPackageArchitecturesALL`, `AppAttachPackageArchitecturesARM`, `AppAttachPackageArchitecturesARM64`, `AppAttachPackageArchitecturesNeutral`, `AppAttachPackageArchitecturesX64`, `AppAttachPackageArchitecturesX86`, `AppAttachPackageArchitecturesX86A64`
- New enum type `DirectUDP` with values `DirectUDPDefault`, `DirectUDPDisabled`, `DirectUDPEnabled`
- New enum type `DomainJoinType` with values `DomainJoinTypeActiveDirectory`, `DomainJoinTypeAzureActiveDirectory`
- New enum type `FailHealthCheckOnStagingFailure` with values `FailHealthCheckOnStagingFailureDoNotFail`, `FailHealthCheckOnStagingFailureNeedsAssistance`, `FailHealthCheckOnStagingFailureUnhealthy`
- New enum type `HostPoolUpdateAction` with values `HostPoolUpdateActionCancel`, `HostPoolUpdateActionPause`, `HostPoolUpdateActionResume`, `HostPoolUpdateActionRetry`, `HostPoolUpdateActionStart`
- New enum type `ManagedPrivateUDP` with values `ManagedPrivateUDPDefault`, `ManagedPrivateUDPDisabled`, `ManagedPrivateUDPEnabled`
- New enum type `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New enum type `ManagementType` with values `ManagementTypeAutomated`, `ManagementTypeStandard`
- New enum type `PackageTimestamped` with values `PackageTimestampedNotTimestamped`, `PackageTimestampedTimestamped`
- New enum type `ProvisioningState` with values `ProvisioningStateCanceled`, `ProvisioningStateFailed`, `ProvisioningStateProvisioning`, `ProvisioningStateSucceeded`
- New enum type `ProvisioningStateSHC` with values `ProvisioningStateSHCCanceled`, `ProvisioningStateSHCFailed`, `ProvisioningStateSHCProvisioning`, `ProvisioningStateSHCSucceeded`
- New enum type `PublicUDP` with values `PublicUDPDefault`, `PublicUDPDisabled`, `PublicUDPEnabled`
- New enum type `RelayUDP` with values `RelayUDPDefault`, `RelayUDPDisabled`, `RelayUDPEnabled`
- New enum type `Type` with values `TypeCustom`, `TypeMarketplace`
- New enum type `VirtualMachineDiskType` with values `VirtualMachineDiskTypePremiumLRS`, `VirtualMachineDiskTypeStandardLRS`, `VirtualMachineDiskTypeStandardSSDLRS`
- New enum type `VirtualMachineSecurityType` with values `VirtualMachineSecurityTypeConfidentialVM`, `VirtualMachineSecurityTypeStandard`, `VirtualMachineSecurityTypeTrustedLaunch`
- New function `NewActiveSessionHostConfigurationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ActiveSessionHostConfigurationsClient, error)`
- New function `*ActiveSessionHostConfigurationsClient.Get(context.Context, string, string, *ActiveSessionHostConfigurationsClientGetOptions) (ActiveSessionHostConfigurationsClientGetResponse, error)`
- New function `*ActiveSessionHostConfigurationsClient.NewListByHostPoolPager(string, string, *ActiveSessionHostConfigurationsClientListByHostPoolOptions) *runtime.Pager[ActiveSessionHostConfigurationsClientListByHostPoolResponse]`
- New function `NewAppAttachPackageClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AppAttachPackageClient, error)`
- New function `*AppAttachPackageClient.CreateOrUpdate(context.Context, string, string, AppAttachPackage, *AppAttachPackageClientCreateOrUpdateOptions) (AppAttachPackageClientCreateOrUpdateResponse, error)`
- New function `*AppAttachPackageClient.Delete(context.Context, string, string, *AppAttachPackageClientDeleteOptions) (AppAttachPackageClientDeleteResponse, error)`
- New function `*AppAttachPackageClient.Get(context.Context, string, string, *AppAttachPackageClientGetOptions) (AppAttachPackageClientGetResponse, error)`
- New function `*AppAttachPackageClient.NewListByResourceGroupPager(string, *AppAttachPackageClientListByResourceGroupOptions) *runtime.Pager[AppAttachPackageClientListByResourceGroupResponse]`
- New function `*AppAttachPackageClient.NewListBySubscriptionPager(*AppAttachPackageClientListBySubscriptionOptions) *runtime.Pager[AppAttachPackageClientListBySubscriptionResponse]`
- New function `*AppAttachPackageClient.Update(context.Context, string, string, *AppAttachPackageClientUpdateOptions) (AppAttachPackageClientUpdateResponse, error)`
- New function `NewAppAttachPackageInfoClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AppAttachPackageInfoClient, error)`
- New function `*AppAttachPackageInfoClient.NewImportPager(string, string, ImportPackageInfoRequest, *AppAttachPackageInfoClientImportOptions) *runtime.Pager[AppAttachPackageInfoClientImportResponse]`
- New function `*ClientFactory.NewActiveSessionHostConfigurationsClient() *ActiveSessionHostConfigurationsClient`
- New function `*ClientFactory.NewAppAttachPackageClient() *AppAttachPackageClient`
- New function `*ClientFactory.NewAppAttachPackageInfoClient() *AppAttachPackageInfoClient`
- New function `*ClientFactory.NewControlSessionHostUpdateClient() *ControlSessionHostUpdateClient`
- New function `*ClientFactory.NewInitiateSessionHostUpdateClient() *InitiateSessionHostUpdateClient`
- New function `*ClientFactory.NewSessionHostClient() *SessionHostClient`
- New function `*ClientFactory.NewSessionHostConfigurationsClient() *SessionHostConfigurationsClient`
- New function `*ClientFactory.NewSessionHostManagementsClient() *SessionHostManagementsClient`
- New function `*ClientFactory.NewSessionHostManagementsUpdateStatusClient() *SessionHostManagementsUpdateStatusClient`
- New function `NewControlSessionHostUpdateClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ControlSessionHostUpdateClient, error)`
- New function `*ControlSessionHostUpdateClient.BeginPost(context.Context, string, string, HostPoolUpdateControlParameter, *ControlSessionHostUpdateClientBeginPostOptions) (*runtime.Poller[ControlSessionHostUpdateClientPostResponse], error)`
- New function `*HostPoolsClient.ListRegistrationTokens(context.Context, string, string, *HostPoolsClientListRegistrationTokensOptions) (HostPoolsClientListRegistrationTokensResponse, error)`
- New function `NewInitiateSessionHostUpdateClient(string, azcore.TokenCredential, *arm.ClientOptions) (*InitiateSessionHostUpdateClient, error)`
- New function `*InitiateSessionHostUpdateClient.Post(context.Context, string, string, *InitiateSessionHostUpdateClientPostOptions) (InitiateSessionHostUpdateClientPostResponse, error)`
- New function `NewSessionHostClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SessionHostClient, error)`
- New function `*SessionHostClient.RetryProvisioning(context.Context, string, string, string, *SessionHostClientRetryProvisioningOptions) (SessionHostClientRetryProvisioningResponse, error)`
- New function `NewSessionHostConfigurationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SessionHostConfigurationsClient, error)`
- New function `*SessionHostConfigurationsClient.CreateOrUpdate(context.Context, string, string, SessionHostConfiguration, *SessionHostConfigurationsClientCreateOrUpdateOptions) (SessionHostConfigurationsClientCreateOrUpdateResponse, error)`
- New function `*SessionHostConfigurationsClient.Get(context.Context, string, string, *SessionHostConfigurationsClientGetOptions) (SessionHostConfigurationsClientGetResponse, error)`
- New function `*SessionHostConfigurationsClient.NewListByHostPoolPager(string, string, *SessionHostConfigurationsClientListByHostPoolOptions) *runtime.Pager[SessionHostConfigurationsClientListByHostPoolResponse]`
- New function `*SessionHostConfigurationsClient.BeginUpdate(context.Context, string, string, *SessionHostConfigurationsClientBeginUpdateOptions) (*runtime.Poller[SessionHostConfigurationsClientUpdateResponse], error)`
- New function `NewSessionHostManagementsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SessionHostManagementsClient, error)`
- New function `*SessionHostManagementsClient.CreateOrUpdate(context.Context, string, string, SessionHostManagement, *SessionHostManagementsClientCreateOrUpdateOptions) (SessionHostManagementsClientCreateOrUpdateResponse, error)`
- New function `*SessionHostManagementsClient.Get(context.Context, string, string, *SessionHostManagementsClientGetOptions) (SessionHostManagementsClientGetResponse, error)`
- New function `*SessionHostManagementsClient.NewListByHostPoolPager(string, string, *SessionHostManagementsClientListByHostPoolOptions) *runtime.Pager[SessionHostManagementsClientListByHostPoolResponse]`
- New function `*SessionHostManagementsClient.Update(context.Context, string, string, *SessionHostManagementsClientUpdateOptions) (SessionHostManagementsClientUpdateResponse, error)`
- New function `NewSessionHostManagementsUpdateStatusClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SessionHostManagementsUpdateStatusClient, error)`
- New function `*SessionHostManagementsUpdateStatusClient.Get(context.Context, string, string, *SessionHostManagementsUpdateStatusClientGetOptions) (SessionHostManagementsUpdateStatusClientGetResponse, error)`
- New struct `ActiveDirectoryInfoPatchProperties`
- New struct `ActiveDirectoryInfoProperties`
- New struct `ActiveSessionHostConfiguration`
- New struct `ActiveSessionHostConfigurationList`
- New struct `ActiveSessionHostConfigurationProperties`
- New struct `AppAttachPackage`
- New struct `AppAttachPackageInfoProperties`
- New struct `AppAttachPackageList`
- New struct `AppAttachPackagePatch`
- New struct `AppAttachPackagePatchProperties`
- New struct `AppAttachPackageProperties`
- New struct `AzureActiveDirectoryInfoProperties`
- New struct `BootDiagnosticsInfoPatchProperties`
- New struct `BootDiagnosticsInfoProperties`
- New struct `CustomInfoPatchProperties`
- New struct `CustomInfoProperties`
- New struct `DiskInfoPatchProperties`
- New struct `DiskInfoProperties`
- New struct `DomainInfoPatchProperties`
- New struct `DomainInfoProperties`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `HostPoolUpdateConfigurationPatchProperties`
- New struct `HostPoolUpdateConfigurationProperties`
- New struct `HostPoolUpdateControlParameter`
- New struct `ImageInfoPatchProperties`
- New struct `ImageInfoProperties`
- New struct `ImportPackageInfoRequest`
- New struct `KeyVaultCredentialsPatchProperties`
- New struct `KeyVaultCredentialsProperties`
- New struct `ManagedServiceIdentity`
- New struct `MarketplaceInfoPatchProperties`
- New struct `MarketplaceInfoProperties`
- New struct `NetworkInfoPatchProperties`
- New struct `NetworkInfoProperties`
- New struct `Plan`
- New struct `RegistrationTokenList`
- New struct `RegistrationTokenMinimal`
- New struct `SKU`
- New struct `SecurityInfoPatchProperties`
- New struct `SecurityInfoProperties`
- New struct `SessionHostConfiguration`
- New struct `SessionHostConfigurationList`
- New struct `SessionHostConfigurationPatch`
- New struct `SessionHostConfigurationPatchProperties`
- New struct `SessionHostConfigurationProperties`
- New struct `SessionHostManagement`
- New struct `SessionHostManagementList`
- New struct `SessionHostManagementOperationProgress`
- New struct `SessionHostManagementPatch`
- New struct `SessionHostManagementPatchProperties`
- New struct `SessionHostManagementProperties`
- New struct `SessionHostManagementUpdateStatus`
- New struct `SessionHostManagementUpdateStatusProperties`
- New struct `UpdateSessionHostsRequestBody`
- New struct `UserAssignedIdentity`
- New field `SystemData` in struct `ApplicationGroupPatch`
- New field `SystemData` in struct `ExpandMsixImage`
- New field `CertificateExpiry`, `CertificateName` in struct `ExpandMsixImageProperties`
- New field `SystemData` in struct `HostPoolPatch`
- New field `DirectUDP`, `ManagedPrivateUDP`, `PublicUDP`, `RelayUDP` in struct `HostPoolPatchProperties`
- New field `AppAttachPackageReferences`, `DirectUDP`, `ManagedPrivateUDP`, `ManagementType`, `PublicUDP`, `RelayUDP` in struct `HostPoolProperties`
- New field `SystemData` in struct `MSIXPackagePatch`
- New field `SystemData` in struct `PrivateEndpointConnection`
- New field `GroupIDs` in struct `PrivateEndpointConnectionProperties`
- New field `SystemData` in struct `PrivateLinkResource`
- New field `SystemData` in struct `ScalingPlanPooledSchedulePatch`
- New field `SystemData` in struct `SessionHostPatch`
- New field `LastSessionHostUpdateTime`, `SessionHostConfiguration` in struct `SessionHostProperties`
- New field `VMPath` in struct `SessionHostsClientListOptions`
- New field `SystemData` in struct `StartMenuItem`


## 2.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 2.1.0 (2023-10-27)
### Features Added

- New enum type `HostpoolPublicNetworkAccess` with values `HostpoolPublicNetworkAccessDisabled`, `HostpoolPublicNetworkAccessEnabled`, `HostpoolPublicNetworkAccessEnabledForClientsOnly`, `HostpoolPublicNetworkAccessEnabledForSessionHostsOnly`
- New enum type `PrivateEndpointConnectionProvisioningState` with values `PrivateEndpointConnectionProvisioningStateCreating`, `PrivateEndpointConnectionProvisioningStateDeleting`, `PrivateEndpointConnectionProvisioningStateFailed`, `PrivateEndpointConnectionProvisioningStateSucceeded`
- New enum type `PrivateEndpointServiceConnectionStatus` with values `PrivateEndpointServiceConnectionStatusApproved`, `PrivateEndpointServiceConnectionStatusPending`, `PrivateEndpointServiceConnectionStatusRejected`
- New enum type `PublicNetworkAccess` with values `PublicNetworkAccessDisabled`, `PublicNetworkAccessEnabled`
- New enum type `SessionHandlingOperation` with values `SessionHandlingOperationDeallocate`, `SessionHandlingOperationHibernate`, `SessionHandlingOperationNone`
- New enum type `SetStartVMOnConnect` with values `SetStartVMOnConnectDisable`, `SetStartVMOnConnectEnable`
- New enum type `StartupBehavior` with values `StartupBehaviorAll`, `StartupBehaviorNone`, `StartupBehaviorWithAssignedUser`
- New function `*ClientFactory.NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient`
- New function `*ClientFactory.NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient`
- New function `*ClientFactory.NewScalingPlanPersonalSchedulesClient() *ScalingPlanPersonalSchedulesClient`
- New function `NewPrivateEndpointConnectionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrivateEndpointConnectionsClient, error)`
- New function `*PrivateEndpointConnectionsClient.DeleteByHostPool(context.Context, string, string, string, *PrivateEndpointConnectionsClientDeleteByHostPoolOptions) (PrivateEndpointConnectionsClientDeleteByHostPoolResponse, error)`
- New function `*PrivateEndpointConnectionsClient.DeleteByWorkspace(context.Context, string, string, string, *PrivateEndpointConnectionsClientDeleteByWorkspaceOptions) (PrivateEndpointConnectionsClientDeleteByWorkspaceResponse, error)`
- New function `*PrivateEndpointConnectionsClient.GetByHostPool(context.Context, string, string, string, *PrivateEndpointConnectionsClientGetByHostPoolOptions) (PrivateEndpointConnectionsClientGetByHostPoolResponse, error)`
- New function `*PrivateEndpointConnectionsClient.GetByWorkspace(context.Context, string, string, string, *PrivateEndpointConnectionsClientGetByWorkspaceOptions) (PrivateEndpointConnectionsClientGetByWorkspaceResponse, error)`
- New function `*PrivateEndpointConnectionsClient.NewListByHostPoolPager(string, string, *PrivateEndpointConnectionsClientListByHostPoolOptions) *runtime.Pager[PrivateEndpointConnectionsClientListByHostPoolResponse]`
- New function `*PrivateEndpointConnectionsClient.NewListByWorkspacePager(string, string, *PrivateEndpointConnectionsClientListByWorkspaceOptions) *runtime.Pager[PrivateEndpointConnectionsClientListByWorkspaceResponse]`
- New function `*PrivateEndpointConnectionsClient.UpdateByHostPool(context.Context, string, string, string, PrivateEndpointConnection, *PrivateEndpointConnectionsClientUpdateByHostPoolOptions) (PrivateEndpointConnectionsClientUpdateByHostPoolResponse, error)`
- New function `*PrivateEndpointConnectionsClient.UpdateByWorkspace(context.Context, string, string, string, PrivateEndpointConnection, *PrivateEndpointConnectionsClientUpdateByWorkspaceOptions) (PrivateEndpointConnectionsClientUpdateByWorkspaceResponse, error)`
- New function `NewPrivateLinkResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrivateLinkResourcesClient, error)`
- New function `*PrivateLinkResourcesClient.NewListByHostPoolPager(string, string, *PrivateLinkResourcesClientListByHostPoolOptions) *runtime.Pager[PrivateLinkResourcesClientListByHostPoolResponse]`
- New function `*PrivateLinkResourcesClient.NewListByWorkspacePager(string, string, *PrivateLinkResourcesClientListByWorkspaceOptions) *runtime.Pager[PrivateLinkResourcesClientListByWorkspaceResponse]`
- New function `NewScalingPlanPersonalSchedulesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ScalingPlanPersonalSchedulesClient, error)`
- New function `*ScalingPlanPersonalSchedulesClient.Create(context.Context, string, string, string, ScalingPlanPersonalSchedule, *ScalingPlanPersonalSchedulesClientCreateOptions) (ScalingPlanPersonalSchedulesClientCreateResponse, error)`
- New function `*ScalingPlanPersonalSchedulesClient.Delete(context.Context, string, string, string, *ScalingPlanPersonalSchedulesClientDeleteOptions) (ScalingPlanPersonalSchedulesClientDeleteResponse, error)`
- New function `*ScalingPlanPersonalSchedulesClient.Get(context.Context, string, string, string, *ScalingPlanPersonalSchedulesClientGetOptions) (ScalingPlanPersonalSchedulesClientGetResponse, error)`
- New function `*ScalingPlanPersonalSchedulesClient.NewListPager(string, string, *ScalingPlanPersonalSchedulesClientListOptions) *runtime.Pager[ScalingPlanPersonalSchedulesClientListResponse]`
- New function `*ScalingPlanPersonalSchedulesClient.Update(context.Context, string, string, string, *ScalingPlanPersonalSchedulesClientUpdateOptions) (ScalingPlanPersonalSchedulesClientUpdateResponse, error)`
- New struct `PrivateEndpoint`
- New struct `PrivateEndpointConnection`
- New struct `PrivateEndpointConnectionListResultWithSystemData`
- New struct `PrivateEndpointConnectionProperties`
- New struct `PrivateEndpointConnectionWithSystemData`
- New struct `PrivateLinkResource`
- New struct `PrivateLinkResourceListResult`
- New struct `PrivateLinkResourceProperties`
- New struct `PrivateLinkServiceConnectionState`
- New struct `ScalingPlanPersonalSchedule`
- New struct `ScalingPlanPersonalScheduleList`
- New struct `ScalingPlanPersonalSchedulePatch`
- New struct `ScalingPlanPersonalScheduleProperties`
- New field `ShowInFeed` in struct `ApplicationGroupPatchProperties`
- New field `ShowInFeed` in struct `ApplicationGroupProperties`
- New field `PublicNetworkAccess` in struct `HostPoolPatchProperties`
- New field `PrivateEndpointConnections`, `PublicNetworkAccess` in struct `HostPoolProperties`
- New field `PublicNetworkAccess` in struct `WorkspacePatchProperties`
- New field `PrivateEndpointConnections`, `PublicNetworkAccess` in struct `WorkspaceProperties`


## 2.0.0 (2023-03-24)
### Breaking Changes

- Type of `ScalingPlanProperties.HostPoolType` has been changed from `*HostPoolType` to `*ScalingHostPoolType`
- Type of `ScalingSchedule.OffPeakStartTime` has been changed from `*time.Time` to `*Time`
- Type of `ScalingSchedule.PeakStartTime` has been changed from `*time.Time` to `*Time`
- Type of `ScalingSchedule.RampDownStartTime` has been changed from `*time.Time` to `*Time`
- Type of `ScalingSchedule.RampUpStartTime` has been changed from `*time.Time` to `*Time`
- Type alias `Operation` has been removed
- Operation `*DesktopsClient.List` has supported pagination, use `*DesktopsClient.NewListPager` instead.
- Operation `*OperationsClient.List` has supported pagination, use `*OperationsClient.NewListPager` instead.
- Struct `MigrationRequestProperties` has been removed
- Field `MigrationRequest` of struct `ApplicationGroupProperties` has been removed
- Field `MigrationRequest` of struct `HostPoolProperties` has been removed
- Field `HostPoolType` of struct `ScalingPlanPatchProperties` has been removed

### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `DayOfWeek` with values `DayOfWeekFriday`, `DayOfWeekMonday`, `DayOfWeekSaturday`, `DayOfWeekSunday`, `DayOfWeekThursday`, `DayOfWeekTuesday`, `DayOfWeekWednesday`
- New enum type `ScalingHostPoolType` with values `ScalingHostPoolTypePooled`
- New enum type `SessionHostComponentUpdateType` with values `SessionHostComponentUpdateTypeDefault`, `SessionHostComponentUpdateTypeScheduled`
- New function `NewScalingPlanPooledSchedulesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ScalingPlanPooledSchedulesClient, error)`
- New function `*ScalingPlanPooledSchedulesClient.Create(context.Context, string, string, string, ScalingPlanPooledSchedule, *ScalingPlanPooledSchedulesClientCreateOptions) (ScalingPlanPooledSchedulesClientCreateResponse, error)`
- New function `*ScalingPlanPooledSchedulesClient.Delete(context.Context, string, string, string, *ScalingPlanPooledSchedulesClientDeleteOptions) (ScalingPlanPooledSchedulesClientDeleteResponse, error)`
- New function `*ScalingPlanPooledSchedulesClient.Get(context.Context, string, string, string, *ScalingPlanPooledSchedulesClientGetOptions) (ScalingPlanPooledSchedulesClientGetResponse, error)`
- New function `*ScalingPlanPooledSchedulesClient.NewListPager(string, string, *ScalingPlanPooledSchedulesClientListOptions) *runtime.Pager[ScalingPlanPooledSchedulesClientListResponse]`
- New function `*ScalingPlanPooledSchedulesClient.Update(context.Context, string, string, string, *ScalingPlanPooledSchedulesClientUpdateOptions) (ScalingPlanPooledSchedulesClientUpdateResponse, error)`
- New struct `AgentUpdatePatchProperties`
- New struct `AgentUpdateProperties`
- New struct `MaintenanceWindowPatchProperties`
- New struct `MaintenanceWindowProperties`
- New struct `ScalingPlanPooledSchedule`
- New struct `ScalingPlanPooledScheduleList`
- New struct `ScalingPlanPooledSchedulePatch`
- New struct `ScalingPlanPooledScheduleProperties`
- New struct `SystemData`
- New struct `Time`
- New field `SystemData` in struct `Application`
- New field `SystemData` in struct `ApplicationGroup`
- New field `InitialSkip` in struct `ApplicationGroupsClientListByResourceGroupOptions`
- New field `IsDescending` in struct `ApplicationGroupsClientListByResourceGroupOptions`
- New field `PageSize` in struct `ApplicationGroupsClientListByResourceGroupOptions`
- New field `InitialSkip` in struct `ApplicationsClientListOptions`
- New field `IsDescending` in struct `ApplicationsClientListOptions`
- New field `PageSize` in struct `ApplicationsClientListOptions`
- New field `SystemData` in struct `Desktop`
- New field `SystemData` in struct `HostPool`
- New field `AgentUpdate` in struct `HostPoolPatchProperties`
- New field `AgentUpdate` in struct `HostPoolProperties`
- New field `InitialSkip` in struct `HostPoolsClientListByResourceGroupOptions`
- New field `IsDescending` in struct `HostPoolsClientListByResourceGroupOptions`
- New field `PageSize` in struct `HostPoolsClientListByResourceGroupOptions`
- New field `InitialSkip` in struct `HostPoolsClientListOptions`
- New field `IsDescending` in struct `HostPoolsClientListOptions`
- New field `PageSize` in struct `HostPoolsClientListOptions`
- New field `SystemData` in struct `MSIXPackage`
- New field `InitialSkip` in struct `MSIXPackagesClientListOptions`
- New field `IsDescending` in struct `MSIXPackagesClientListOptions`
- New field `PageSize` in struct `MSIXPackagesClientListOptions`
- New field `NextLink` in struct `ResourceProviderOperationList`
- New field `SystemData` in struct `ScalingPlan`
- New field `InitialSkip` in struct `ScalingPlansClientListByHostPoolOptions`
- New field `IsDescending` in struct `ScalingPlansClientListByHostPoolOptions`
- New field `PageSize` in struct `ScalingPlansClientListByHostPoolOptions`
- New field `InitialSkip` in struct `ScalingPlansClientListByResourceGroupOptions`
- New field `IsDescending` in struct `ScalingPlansClientListByResourceGroupOptions`
- New field `PageSize` in struct `ScalingPlansClientListByResourceGroupOptions`
- New field `InitialSkip` in struct `ScalingPlansClientListBySubscriptionOptions`
- New field `IsDescending` in struct `ScalingPlansClientListBySubscriptionOptions`
- New field `PageSize` in struct `ScalingPlansClientListBySubscriptionOptions`
- New field `SystemData` in struct `SessionHost`
- New field `FriendlyName` in struct `SessionHostPatchProperties`
- New field `FriendlyName` in struct `SessionHostProperties`
- New field `InitialSkip` in struct `SessionHostsClientListOptions`
- New field `IsDescending` in struct `SessionHostsClientListOptions`
- New field `PageSize` in struct `SessionHostsClientListOptions`
- New field `Force` in struct `SessionHostsClientUpdateOptions`
- New field `InitialSkip` in struct `StartMenuItemsClientListOptions`
- New field `IsDescending` in struct `StartMenuItemsClientListOptions`
- New field `PageSize` in struct `StartMenuItemsClientListOptions`
- New field `SystemData` in struct `UserSession`
- New field `InitialSkip` in struct `UserSessionsClientListByHostPoolOptions`
- New field `IsDescending` in struct `UserSessionsClientListByHostPoolOptions`
- New field `PageSize` in struct `UserSessionsClientListByHostPoolOptions`
- New field `InitialSkip` in struct `UserSessionsClientListOptions`
- New field `IsDescending` in struct `UserSessionsClientListOptions`
- New field `PageSize` in struct `UserSessionsClientListOptions`
- New field `SystemData` in struct `Workspace`
- New field `InitialSkip` in struct `WorkspacesClientListByResourceGroupOptions`
- New field `IsDescending` in struct `WorkspacesClientListByResourceGroupOptions`
- New field `PageSize` in struct `WorkspacesClientListByResourceGroupOptions`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).