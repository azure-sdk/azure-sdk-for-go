# Release History

## 3.0.0-beta.1 (2022-09-14)
### Breaking Changes

- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New const `NetworkPluginModeOverlay`
- New const `FleetMemberProvisioningStateJoining`
- New const `FleetProvisioningStateCanceled`
- New const `BackendPoolTypeNodeIP`
- New const `ControlledValuesRequestsOnly`
- New const `FleetMemberProvisioningStateLeaving`
- New const `IpvsSchedulerLeastConnection`
- New const `FleetProvisioningStateFailed`
- New const `LevelOff`
- New const `TrustedAccessRoleBindingProvisioningStateDeleting`
- New const `UpdateModeInitial`
- New const `FleetProvisioningStateSucceeded`
- New const `TrustedAccessRoleBindingProvisioningStateSucceeded`
- New const `FleetMemberProvisioningStateUpdating`
- New const `FleetMemberProvisioningStateCanceled`
- New const `FleetMemberProvisioningStateSucceeded`
- New const `TrustedAccessRoleBindingProvisioningStateFailed`
- New const `UpdateModeAuto`
- New const `ModeIPVS`
- New const `LevelEnforcement`
- New const `UpdateModeRecreate`
- New const `SnapshotTypeManagedCluster`
- New const `TrustedAccessRoleBindingProvisioningStateUpdating`
- New const `LevelWarning`
- New const `FleetProvisioningStateUpdating`
- New const `FleetMemberProvisioningStateFailed`
- New const `OSSKUMariner`
- New const `FleetProvisioningStateCreating`
- New const `PublicNetworkAccessSecuredByPerimeter`
- New const `BackendPoolTypeNodeIPConfiguration`
- New const `FleetProvisioningStateDeleting`
- New const `ModeIPTABLES`
- New const `ControlledValuesRequestsAndLimits`
- New const `IpvsSchedulerRoundRobin`
- New const `UpdateModeOff`
- New type alias `UpdateMode`
- New type alias `IpvsScheduler`
- New type alias `TrustedAccessRoleBindingProvisioningState`
- New type alias `Level`
- New type alias `BackendPoolType`
- New type alias `NetworkPluginMode`
- New type alias `FleetProvisioningState`
- New type alias `FleetMemberProvisioningState`
- New type alias `Mode`
- New type alias `ControlledValues`
- New function `NewTrustedAccessRoleBindingsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*TrustedAccessRoleBindingsClient, error)`
- New function `NewFleetMembersClient(string, azcore.TokenCredential, *arm.ClientOptions) (*FleetMembersClient, error)`
- New function `PossibleFleetMemberProvisioningStateValues() []FleetMemberProvisioningState`
- New function `*TrustedAccessRoleBindingsClient.CreateOrUpdate(context.Context, string, string, string, TrustedAccessRoleBinding, *TrustedAccessRoleBindingsClientCreateOrUpdateOptions) (TrustedAccessRoleBindingsClientCreateOrUpdateResponse, error)`
- New function `*FleetsClient.ListCredentials(context.Context, string, string, *FleetsClientListCredentialsOptions) (FleetsClientListCredentialsResponse, error)`
- New function `*FleetMembersClient.Get(context.Context, string, string, string, *FleetMembersClientGetOptions) (FleetMembersClientGetResponse, error)`
- New function `PossibleControlledValuesValues() []ControlledValues`
- New function `NewTrustedAccessRolesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*TrustedAccessRolesClient, error)`
- New function `NewManagedClusterSnapshotsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagedClusterSnapshotsClient, error)`
- New function `PossibleModeValues() []Mode`
- New function `*ManagedClusterSnapshotsClient.UpdateTags(context.Context, string, string, TagsObject, *ManagedClusterSnapshotsClientUpdateTagsOptions) (ManagedClusterSnapshotsClientUpdateTagsResponse, error)`
- New function `NewFleetsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*FleetsClient, error)`
- New function `*FleetMembersClient.NewListByFleetPager(string, string, *FleetMembersClientListByFleetOptions) *runtime.Pager[FleetMembersClientListByFleetResponse]`
- New function `*TrustedAccessRoleBindingsClient.Get(context.Context, string, string, string, *TrustedAccessRoleBindingsClientGetOptions) (TrustedAccessRoleBindingsClientGetResponse, error)`
- New function `*AgentPoolsClient.AbortLatestOperation(context.Context, string, string, string, *AgentPoolsClientAbortLatestOperationOptions) (AgentPoolsClientAbortLatestOperationResponse, error)`
- New function `PossibleLevelValues() []Level`
- New function `PossibleNetworkPluginModeValues() []NetworkPluginMode`
- New function `PossibleBackendPoolTypeValues() []BackendPoolType`
- New function `*FleetsClient.NewListByResourceGroupPager(string, *FleetsClientListByResourceGroupOptions) *runtime.Pager[FleetsClientListByResourceGroupResponse]`
- New function `*ManagedClusterSnapshotsClient.NewListPager(*ManagedClusterSnapshotsClientListOptions) *runtime.Pager[ManagedClusterSnapshotsClientListResponse]`
- New function `*ManagedClusterSnapshotsClient.Get(context.Context, string, string, *ManagedClusterSnapshotsClientGetOptions) (ManagedClusterSnapshotsClientGetResponse, error)`
- New function `*ManagedClusterSnapshotsClient.Delete(context.Context, string, string, *ManagedClusterSnapshotsClientDeleteOptions) (ManagedClusterSnapshotsClientDeleteResponse, error)`
- New function `*FleetMembersClient.BeginDelete(context.Context, string, string, string, *FleetMembersClientBeginDeleteOptions) (*runtime.Poller[FleetMembersClientDeleteResponse], error)`
- New function `*FleetsClient.BeginDelete(context.Context, string, string, *FleetsClientBeginDeleteOptions) (*runtime.Poller[FleetsClientDeleteResponse], error)`
- New function `*FleetMembersClient.BeginCreateOrUpdate(context.Context, string, string, string, FleetMember, *FleetMembersClientBeginCreateOrUpdateOptions) (*runtime.Poller[FleetMembersClientCreateOrUpdateResponse], error)`
- New function `*FleetsClient.BeginCreateOrUpdate(context.Context, string, string, Fleet, *FleetsClientBeginCreateOrUpdateOptions) (*runtime.Poller[FleetsClientCreateOrUpdateResponse], error)`
- New function `*ManagedClusterSnapshotsClient.NewListByResourceGroupPager(string, *ManagedClusterSnapshotsClientListByResourceGroupOptions) *runtime.Pager[ManagedClusterSnapshotsClientListByResourceGroupResponse]`
- New function `*FleetsClient.Get(context.Context, string, string, *FleetsClientGetOptions) (FleetsClientGetResponse, error)`
- New function `*ManagedClustersClient.AbortLatestOperation(context.Context, string, string, *ManagedClustersClientAbortLatestOperationOptions) (ManagedClustersClientAbortLatestOperationResponse, error)`
- New function `*ManagedClustersClient.BeginRotateServiceAccountSigningKeys(context.Context, string, string, *ManagedClustersClientBeginRotateServiceAccountSigningKeysOptions) (*runtime.Poller[ManagedClustersClientRotateServiceAccountSigningKeysResponse], error)`
- New function `*TrustedAccessRoleBindingsClient.NewListPager(string, string, *TrustedAccessRoleBindingsClientListOptions) *runtime.Pager[TrustedAccessRoleBindingsClientListResponse]`
- New function `PossibleTrustedAccessRoleBindingProvisioningStateValues() []TrustedAccessRoleBindingProvisioningState`
- New function `*TrustedAccessRolesClient.NewListPager(string, *TrustedAccessRolesClientListOptions) *runtime.Pager[TrustedAccessRolesClientListResponse]`
- New function `PossibleUpdateModeValues() []UpdateMode`
- New function `*FleetsClient.Update(context.Context, string, string, *FleetsClientUpdateOptions) (FleetsClientUpdateResponse, error)`
- New function `PossibleIpvsSchedulerValues() []IpvsScheduler`
- New function `*ManagedClusterSnapshotsClient.CreateOrUpdate(context.Context, string, string, ManagedClusterSnapshot, *ManagedClusterSnapshotsClientCreateOrUpdateOptions) (ManagedClusterSnapshotsClientCreateOrUpdateResponse, error)`
- New function `*FleetsClient.NewListPager(*FleetsClientListOptions) *runtime.Pager[FleetsClientListResponse]`
- New function `PossibleFleetProvisioningStateValues() []FleetProvisioningState`
- New function `*TrustedAccessRoleBindingsClient.Delete(context.Context, string, string, string, *TrustedAccessRoleBindingsClientDeleteOptions) (TrustedAccessRoleBindingsClientDeleteResponse, error)`
- New struct `AgentPoolWindowsProfile`
- New struct `AgentPoolsClientAbortLatestOperationOptions`
- New struct `AgentPoolsClientAbortLatestOperationResponse`
- New struct `AzureEntityResource`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New struct `Fleet`
- New struct `FleetCredentialResult`
- New struct `FleetCredentialResults`
- New struct `FleetHubProfile`
- New struct `FleetListResult`
- New struct `FleetMember`
- New struct `FleetMemberProperties`
- New struct `FleetMembersClient`
- New struct `FleetMembersClientBeginCreateOrUpdateOptions`
- New struct `FleetMembersClientBeginDeleteOptions`
- New struct `FleetMembersClientCreateOrUpdateResponse`
- New struct `FleetMembersClientDeleteResponse`
- New struct `FleetMembersClientGetOptions`
- New struct `FleetMembersClientGetResponse`
- New struct `FleetMembersClientListByFleetOptions`
- New struct `FleetMembersClientListByFleetResponse`
- New struct `FleetMembersListResult`
- New struct `FleetPatch`
- New struct `FleetProperties`
- New struct `FleetsClient`
- New struct `FleetsClientBeginCreateOrUpdateOptions`
- New struct `FleetsClientBeginDeleteOptions`
- New struct `FleetsClientCreateOrUpdateResponse`
- New struct `FleetsClientDeleteResponse`
- New struct `FleetsClientGetOptions`
- New struct `FleetsClientGetResponse`
- New struct `FleetsClientListByResourceGroupOptions`
- New struct `FleetsClientListByResourceGroupResponse`
- New struct `FleetsClientListCredentialsOptions`
- New struct `FleetsClientListCredentialsResponse`
- New struct `FleetsClientListOptions`
- New struct `FleetsClientListResponse`
- New struct `FleetsClientUpdateOptions`
- New struct `FleetsClientUpdateResponse`
- New struct `GuardrailsProfile`
- New struct `ManagedClusterAzureMonitorProfile`
- New struct `ManagedClusterAzureMonitorProfileKubeStateMetrics`
- New struct `ManagedClusterAzureMonitorProfileMetrics`
- New struct `ManagedClusterIngressProfile`
- New struct `ManagedClusterIngressProfileWebAppRouting`
- New struct `ManagedClusterOIDCIssuerProfile`
- New struct `ManagedClusterPropertiesForSnapshot`
- New struct `ManagedClusterSecurityProfileImageCleaner`
- New struct `ManagedClusterSecurityProfileNodeRestriction`
- New struct `ManagedClusterSecurityProfileWorkloadIdentity`
- New struct `ManagedClusterSnapshot`
- New struct `ManagedClusterSnapshotListResult`
- New struct `ManagedClusterSnapshotProperties`
- New struct `ManagedClusterSnapshotsClient`
- New struct `ManagedClusterSnapshotsClientCreateOrUpdateOptions`
- New struct `ManagedClusterSnapshotsClientCreateOrUpdateResponse`
- New struct `ManagedClusterSnapshotsClientDeleteOptions`
- New struct `ManagedClusterSnapshotsClientDeleteResponse`
- New struct `ManagedClusterSnapshotsClientGetOptions`
- New struct `ManagedClusterSnapshotsClientGetResponse`
- New struct `ManagedClusterSnapshotsClientListByResourceGroupOptions`
- New struct `ManagedClusterSnapshotsClientListByResourceGroupResponse`
- New struct `ManagedClusterSnapshotsClientListOptions`
- New struct `ManagedClusterSnapshotsClientListResponse`
- New struct `ManagedClusterSnapshotsClientUpdateTagsOptions`
- New struct `ManagedClusterSnapshotsClientUpdateTagsResponse`
- New struct `ManagedClusterStorageProfileBlobCSIDriver`
- New struct `ManagedClusterWorkloadAutoScalerProfile`
- New struct `ManagedClusterWorkloadAutoScalerProfileKeda`
- New struct `ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler`
- New struct `ManagedClustersClientAbortLatestOperationOptions`
- New struct `ManagedClustersClientAbortLatestOperationResponse`
- New struct `ManagedClustersClientBeginRotateServiceAccountSigningKeysOptions`
- New struct `ManagedClustersClientRotateServiceAccountSigningKeysResponse`
- New struct `NetworkProfileForSnapshot`
- New struct `NetworkProfileKubeProxyConfig`
- New struct `NetworkProfileKubeProxyConfigIpvsConfig`
- New struct `TrustedAccessRole`
- New struct `TrustedAccessRoleBinding`
- New struct `TrustedAccessRoleBindingListResult`
- New struct `TrustedAccessRoleBindingProperties`
- New struct `TrustedAccessRoleBindingsClient`
- New struct `TrustedAccessRoleBindingsClientCreateOrUpdateOptions`
- New struct `TrustedAccessRoleBindingsClientCreateOrUpdateResponse`
- New struct `TrustedAccessRoleBindingsClientDeleteOptions`
- New struct `TrustedAccessRoleBindingsClientDeleteResponse`
- New struct `TrustedAccessRoleBindingsClientGetOptions`
- New struct `TrustedAccessRoleBindingsClientGetResponse`
- New struct `TrustedAccessRoleBindingsClientListOptions`
- New struct `TrustedAccessRoleBindingsClientListResponse`
- New struct `TrustedAccessRoleListResult`
- New struct `TrustedAccessRoleRule`
- New struct `TrustedAccessRolesClient`
- New struct `TrustedAccessRolesClientListOptions`
- New struct `TrustedAccessRolesClientListResponse`
- New field `BackendPoolType` in struct `ManagedClusterLoadBalancerProfile`
- New field `EffectiveNoProxy` in struct `ManagedClusterHTTPProxyConfig`
- New field `NodeRestriction` in struct `ManagedClusterSecurityProfile`
- New field `WorkloadIdentity` in struct `ManagedClusterSecurityProfile`
- New field `ImageCleaner` in struct `ManagedClusterSecurityProfile`
- New field `IgnorePodDisruptionBudget` in struct `ManagedClustersClientBeginDeleteOptions`
- New field `MessageOfTheDay` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `EnableCustomCATrust` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `CapacityReservationGroupID` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `WindowsProfile` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `IgnorePodDisruptionBudget` in struct `AgentPoolsClientBeginDeleteOptions`
- New field `KubeProxyConfig` in struct `NetworkProfile`
- New field `NetworkPluginMode` in struct `NetworkProfile`
- New field `Version` in struct `ManagedClusterStorageProfileDiskCSIDriver`
- New field `AzureMonitorProfile` in struct `ManagedClusterProperties`
- New field `OidcIssuerProfile` in struct `ManagedClusterProperties`
- New field `GuardrailsProfile` in struct `ManagedClusterProperties`
- New field `WorkloadAutoScalerProfile` in struct `ManagedClusterProperties`
- New field `CreationData` in struct `ManagedClusterProperties`
- New field `EnableNamespaceResources` in struct `ManagedClusterProperties`
- New field `IngressProfile` in struct `ManagedClusterProperties`
- New field `EnableVnetIntegration` in struct `ManagedClusterAPIServerAccessProfile`
- New field `SubnetID` in struct `ManagedClusterAPIServerAccessProfile`
- New field `MessageOfTheDay` in struct `ManagedClusterAgentPoolProfile`
- New field `EnableCustomCATrust` in struct `ManagedClusterAgentPoolProfile`
- New field `WindowsProfile` in struct `ManagedClusterAgentPoolProfile`
- New field `CapacityReservationGroupID` in struct `ManagedClusterAgentPoolProfile`
- New field `BlobCSIDriver` in struct `ManagedClusterStorageProfile`


## 2.1.0 (2022-08-25)
### Features Added

- New const `OSSKUWindows2019`
- New const `OSSKUWindows2022`


## 2.0.0 (2022-07-22)
### Breaking Changes

- Struct `ManagedClusterSecurityProfileAzureDefender` has been removed
- Field `AzureDefender` of struct `ManagedClusterSecurityProfile` has been removed

### Features Added

- New const `KeyVaultNetworkAccessTypesPrivate`
- New const `NetworkPluginNone`
- New const `KeyVaultNetworkAccessTypesPublic`
- New function `PossibleKeyVaultNetworkAccessTypesValues() []KeyVaultNetworkAccessTypes`
- New struct `AzureKeyVaultKms`
- New struct `ManagedClusterSecurityProfileDefender`
- New struct `ManagedClusterSecurityProfileDefenderSecurityMonitoring`
- New field `HostGroupID` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `HostGroupID` in struct `ManagedClusterAgentPoolProfile`
- New field `AzureKeyVaultKms` in struct `ManagedClusterSecurityProfile`
- New field `Defender` in struct `ManagedClusterSecurityProfile`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).