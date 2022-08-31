# Release History

## 2.2.0-beta.1 (2022-08-31)
### Features Added

- New const `BackendPoolTypeNodeIPConfiguration`
- New const `ModeIPVS`
- New const `ControlledValuesRequestsAndLimits`
- New const `LevelWarning`
- New const `PublicNetworkAccessSecuredByPerimeter`
- New const `LevelOff`
- New const `UpdateModeOff`
- New const `ControlledValuesRequestsOnly`
- New const `TrustedAccessRoleBindingProvisioningStateUpdating`
- New const `IpvsSchedulerRoundRobin`
- New const `OSSKUMariner`
- New const `UpdateModeInitial`
- New const `UpdateModeAuto`
- New const `BackendPoolTypeNodeIP`
- New const `LevelEnforcement`
- New const `TrustedAccessRoleBindingProvisioningStateFailed`
- New const `ModeIPTABLES`
- New const `SnapshotTypeManagedCluster`
- New const `TrustedAccessRoleBindingProvisioningStateDeleting`
- New const `IpvsSchedulerLeastConnection`
- New const `TrustedAccessRoleBindingProvisioningStateSucceeded`
- New const `NetworkPluginModeOverlay`
- New const `UpdateModeRecreate`
- New type alias `IpvsScheduler`
- New type alias `Mode`
- New type alias `ControlledValues`
- New type alias `Level`
- New type alias `UpdateMode`
- New type alias `NetworkPluginMode`
- New type alias `BackendPoolType`
- New type alias `TrustedAccessRoleBindingProvisioningState`
- New function `*ManagedClusterSnapshotsClient.NewListPager(*ManagedClusterSnapshotsClientListOptions) *runtime.Pager[ManagedClusterSnapshotsClientListResponse]`
- New function `NewTrustedAccessRoleBindingsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*TrustedAccessRoleBindingsClient, error)`
- New function `PossibleBackendPoolTypeValues() []BackendPoolType`
- New function `*ManagedClusterSnapshotsClient.Get(context.Context, string, string, *ManagedClusterSnapshotsClientGetOptions) (ManagedClusterSnapshotsClientGetResponse, error)`
- New function `NewTrustedAccessRolesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*TrustedAccessRolesClient, error)`
- New function `PossibleIpvsSchedulerValues() []IpvsScheduler`
- New function `*TrustedAccessRoleBindingsClient.Delete(context.Context, string, string, string, *TrustedAccessRoleBindingsClientDeleteOptions) (TrustedAccessRoleBindingsClientDeleteResponse, error)`
- New function `NewManagedClusterSnapshotsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ManagedClusterSnapshotsClient, error)`
- New function `NewGuardrailsVersionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*GuardrailsVersionsClient, error)`
- New function `*GuardrailsVersionsClient.List(context.Context, string, string, *GuardrailsVersionsClientListOptions) (GuardrailsVersionsClientListResponse, error)`
- New function `PossibleLevelValues() []Level`
- New function `*TrustedAccessRoleBindingsClient.NewListPager(string, string, *TrustedAccessRoleBindingsClientListOptions) *runtime.Pager[TrustedAccessRoleBindingsClientListResponse]`
- New function `PossibleControlledValuesValues() []ControlledValues`
- New function `PossibleUpdateModeValues() []UpdateMode`
- New function `*TrustedAccessRoleBindingsClient.Get(context.Context, string, string, string, *TrustedAccessRoleBindingsClientGetOptions) (TrustedAccessRoleBindingsClientGetResponse, error)`
- New function `PossibleNetworkPluginModeValues() []NetworkPluginMode`
- New function `PossibleTrustedAccessRoleBindingProvisioningStateValues() []TrustedAccessRoleBindingProvisioningState`
- New function `PossibleModeValues() []Mode`
- New function `*ManagedClusterSnapshotsClient.UpdateTags(context.Context, string, string, TagsObject, *ManagedClusterSnapshotsClientUpdateTagsOptions) (ManagedClusterSnapshotsClientUpdateTagsResponse, error)`
- New function `*ManagedClusterSnapshotsClient.CreateOrUpdate(context.Context, string, string, ManagedClusterSnapshot, *ManagedClusterSnapshotsClientCreateOrUpdateOptions) (ManagedClusterSnapshotsClientCreateOrUpdateResponse, error)`
- New function `*ManagedClustersClient.BeginRotateServiceAccountSigningKeys(context.Context, string, string, *ManagedClustersClientBeginRotateServiceAccountSigningKeysOptions) (*runtime.Poller[ManagedClustersClientRotateServiceAccountSigningKeysResponse], error)`
- New function `*TrustedAccessRoleBindingsClient.CreateOrUpdate(context.Context, string, string, string, TrustedAccessRoleBinding, *TrustedAccessRoleBindingsClientCreateOrUpdateOptions) (TrustedAccessRoleBindingsClientCreateOrUpdateResponse, error)`
- New function `*TrustedAccessRolesClient.NewListPager(string, *TrustedAccessRolesClientListOptions) *runtime.Pager[TrustedAccessRolesClientListResponse]`
- New function `*ManagedClustersClient.AbortLatestOperation(context.Context, string, string, *ManagedClustersClientAbortLatestOperationOptions) (ManagedClustersClientAbortLatestOperationResponse, error)`
- New function `*ManagedClusterSnapshotsClient.NewListByResourceGroupPager(string, *ManagedClusterSnapshotsClientListByResourceGroupOptions) *runtime.Pager[ManagedClusterSnapshotsClientListByResourceGroupResponse]`
- New function `*ManagedClusterSnapshotsClient.Delete(context.Context, string, string, *ManagedClusterSnapshotsClientDeleteOptions) (ManagedClusterSnapshotsClientDeleteResponse, error)`
- New function `*AgentPoolsClient.AbortLatestOperation(context.Context, string, string, string, *AgentPoolsClientAbortLatestOperationOptions) (AgentPoolsClientAbortLatestOperationResponse, error)`
- New struct `AgentPoolWindowsProfile`
- New struct `AgentPoolsClientAbortLatestOperationOptions`
- New struct `AgentPoolsClientAbortLatestOperationResponse`
- New struct `GuardrailsProfile`
- New struct `GuardrailsVersionsClient`
- New struct `GuardrailsVersionsClientListOptions`
- New struct `GuardrailsVersionsClientListResponse`
- New struct `ListGuardrailsVersionsResponse`
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
- New field `MessageOfTheDay` in struct `ManagedClusterAgentPoolProfile`
- New field `EnableCustomCATrust` in struct `ManagedClusterAgentPoolProfile`
- New field `CapacityReservationGroupID` in struct `ManagedClusterAgentPoolProfile`
- New field `WindowsProfile` in struct `ManagedClusterAgentPoolProfile`
- New field `IgnorePodDisruptionBudget` in struct `ManagedClustersClientBeginDeleteOptions`
- New field `SubnetID` in struct `ManagedClusterAPIServerAccessProfile`
- New field `EnableVnetIntegration` in struct `ManagedClusterAPIServerAccessProfile`
- New field `WorkloadIdentity` in struct `ManagedClusterSecurityProfile`
- New field `ImageCleaner` in struct `ManagedClusterSecurityProfile`
- New field `NodeRestriction` in struct `ManagedClusterSecurityProfile`
- New field `IgnorePodDisruptionBudget` in struct `AgentPoolsClientBeginDeleteOptions`
- New field `BackendPoolType` in struct `ManagedClusterLoadBalancerProfile`
- New field `Version` in struct `ManagedClusterStorageProfileDiskCSIDriver`
- New field `BlobCSIDriver` in struct `ManagedClusterStorageProfile`
- New field `EffectiveNoProxy` in struct `ManagedClusterHTTPProxyConfig`
- New field `EnableCustomCATrust` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `WindowsProfile` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `MessageOfTheDay` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `CapacityReservationGroupID` in struct `ManagedClusterAgentPoolProfileProperties`
- New field `KubeProxyConfig` in struct `NetworkProfile`
- New field `NetworkPluginMode` in struct `NetworkProfile`
- New field `WorkloadAutoScalerProfile` in struct `ManagedClusterProperties`
- New field `AzureMonitorProfile` in struct `ManagedClusterProperties`
- New field `OidcIssuerProfile` in struct `ManagedClusterProperties`
- New field `EnableNamespaceResources` in struct `ManagedClusterProperties`
- New field `IngressProfile` in struct `ManagedClusterProperties`
- New field `CreationData` in struct `ManagedClusterProperties`


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