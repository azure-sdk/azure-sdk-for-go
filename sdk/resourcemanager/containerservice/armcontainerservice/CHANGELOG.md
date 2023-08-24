# Release History

## 4.2.0-beta.3 (2023-08-24)
### Breaking Changes

- Enum `ControlPlaneUpgradeOverride` has been removed
- Enum `ControlledValues` has been removed
- Enum `UpdateMode` has been removed
- Field `DNSZoneResourceID` of struct `ManagedClusterIngressProfileWebAppRouting` has been removed
- Field `ControlledValues`, `UpdateMode` of struct `ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler` has been removed
- Field `ControlPlaneOverrides` of struct `UpgradeOverrideSettings` has been removed

### Features Added

- New enum type `GuardrailsSupport` with values `GuardrailsSupportPreview`, `GuardrailsSupportStable`
- New function `*ClientFactory.NewMachinesClient() *MachinesClient`
- New function `NewMachinesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*MachinesClient, error)`
- New function `*MachinesClient.Get(context.Context, string, string, string, string, *MachinesClientGetOptions) (MachinesClientGetResponse, error)`
- New function `*MachinesClient.NewListPager(string, string, string, *MachinesClientListOptions) *runtime.Pager[MachinesClientListResponse]`
- New function `*ManagedClustersClient.GetGuardrailsVersions(context.Context, string, string, *ManagedClustersClientGetGuardrailsVersionsOptions) (ManagedClustersClientGetGuardrailsVersionsResponse, error)`
- New function `*ManagedClustersClient.NewListGuardrailsVersionsPager(string, *ManagedClustersClientListGuardrailsVersionsOptions) *runtime.Pager[ManagedClustersClientListGuardrailsVersionsResponse]`
- New struct `GuardrailsAvailableVersion`
- New struct `GuardrailsAvailableVersionsList`
- New struct `GuardrailsAvailableVersionsProperties`
- New struct `IstioEgressGateway`
- New struct `Machine`
- New struct `MachineIPAddress`
- New struct `MachineListResult`
- New struct `MachineNetworkProperties`
- New struct `MachineProperties`
- New struct `ManagedClusterAzureMonitorProfileAppMonitoring`
- New struct `ManagedClusterAzureMonitorProfileAppMonitoringOpenTelemetryMetrics`
- New struct `ManagedClusterAzureMonitorProfileContainerInsights`
- New struct `ManagedClusterAzureMonitorProfileLogs`
- New struct `ManagedClusterAzureMonitorProfileWindowsHostLogs`
- New struct `ManagedClusterCostAnalysis`
- New struct `ManagedClusterMetricsProfile`
- New field `EgressGateways` in struct `IstioComponents`
- New field `Logs` in struct `ManagedClusterAzureMonitorProfile`
- New field `AppMonitoringOpenTelemetryMetrics` in struct `ManagedClusterAzureMonitorProfileMetrics`
- New field `DNSZoneResourceIDs` in struct `ManagedClusterIngressProfileWebAppRouting`
- New field `MetricsProfile`, `ResourceUID` in struct `ManagedClusterProperties`
- New field `ForceUpgrade` in struct `UpgradeOverrideSettings`


## 4.2.0 (2023-08-25)
### Features Added

- New enum type `NodeOSUpgradeChannel` with values `NodeOSUpgradeChannelNodeImage`, `NodeOSUpgradeChannelNone`, `NodeOSUpgradeChannelUnmanaged`
- New struct `DelegatedResource`
- New struct `ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler`
- New field `DrainTimeoutInMinutes` in struct `AgentPoolUpgradeSettings`
- New field `NodeOSUpgradeChannel` in struct `ManagedClusterAutoUpgradeProfile`
- New field `DelegatedResources` in struct `ManagedClusterIdentity`
- New field `VerticalPodAutoscaler` in struct `ManagedClusterWorkloadAutoScalerProfile`


## 4.1.0 (2023-07-28)
### Features Added

- New enum type `Type` with values `TypeFirst`, `TypeFourth`, `TypeLast`, `TypeSecond`, `TypeThird`
- New struct `AbsoluteMonthlySchedule`
- New struct `DailySchedule`
- New struct `DateSpan`
- New struct `MaintenanceWindow`
- New struct `RelativeMonthlySchedule`
- New struct `Schedule`
- New struct `WeeklySchedule`
- New field `MaintenanceWindow` in struct `MaintenanceConfigurationProperties`


## 4.0.0 (2023-05-26)
### Breaking Changes

- Field `DockerBridgeCidr` of struct `NetworkProfile` has been removed

### Features Added

- New value `OSSKUAzureLinux` added to enum type `OSSKU`


## 3.0.0 (2023-04-28)
### Breaking Changes

- Const `ManagedClusterSKUNameBasic` from type alias `ManagedClusterSKUName` has been removed
- Const `ManagedClusterSKUTierPaid` from type alias `ManagedClusterSKUTier` has been removed

### Features Added

- New value `ManagedClusterSKUTierPremium` added to enum type `ManagedClusterSKUTier`
- New value `NetworkPolicyCilium` added to enum type `NetworkPolicy`
- New enum type `KubernetesSupportPlan` with values `KubernetesSupportPlanAKSLongTermSupport`, `KubernetesSupportPlanKubernetesOfficial`
- New enum type `NetworkDataplane` with values `NetworkDataplaneAzure`, `NetworkDataplaneCilium`
- New enum type `NetworkPluginMode` with values `NetworkPluginModeOverlay`
- New function `*ManagedClustersClient.ListKubernetesVersions(context.Context, string, *ManagedClustersClientListKubernetesVersionsOptions) (ManagedClustersClientListKubernetesVersionsResponse, error)`
- New struct `KubernetesPatchVersion`
- New struct `KubernetesVersion`
- New struct `KubernetesVersionCapabilities`
- New struct `KubernetesVersionListResult`
- New struct `ManagedClusterSecurityProfileImageCleaner`
- New struct `ManagedClusterSecurityProfileWorkloadIdentity`
- New field `SupportPlan` in struct `ManagedClusterProperties`
- New field `ImageCleaner` in struct `ManagedClusterSecurityProfile`
- New field `WorkloadIdentity` in struct `ManagedClusterSecurityProfile`
- New field `NetworkDataplane` in struct `NetworkProfile`
- New field `NetworkPluginMode` in struct `NetworkProfile`


## 2.4.0 (2023-03-24)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module
- New value `ManagedClusterSKUNameBase` added to enum type `ManagedClusterSKUName`
- New value `ManagedClusterSKUTierStandard` added to enum type `ManagedClusterSKUTier`
- New function `*AgentPoolsClient.BeginAbortLatestOperation(context.Context, string, string, string, *AgentPoolsClientBeginAbortLatestOperationOptions) (*runtime.Poller[AgentPoolsClientAbortLatestOperationResponse], error)`
- New function `*ManagedClustersClient.BeginAbortLatestOperation(context.Context, string, string, *ManagedClustersClientBeginAbortLatestOperationOptions) (*runtime.Poller[ManagedClustersClientAbortLatestOperationResponse], error)`
- New struct `ManagedClusterAzureMonitorProfile`
- New struct `ManagedClusterAzureMonitorProfileKubeStateMetrics`
- New struct `ManagedClusterAzureMonitorProfileMetrics`
- New field `AzureMonitorProfile` in struct `ManagedClusterProperties`


## 2.3.0 (2023-01-27)
### Features Added

- New value `ManagedClusterPodIdentityProvisioningStateCanceled`, `ManagedClusterPodIdentityProvisioningStateSucceeded` added to type alias `ManagedClusterPodIdentityProvisioningState`
- New value `PrivateEndpointConnectionProvisioningStateCanceled` added to type alias `PrivateEndpointConnectionProvisioningState`
- New struct `ManagedClusterWorkloadAutoScalerProfile`
- New struct `ManagedClusterWorkloadAutoScalerProfileKeda`
- New field `WorkloadAutoScalerProfile` in struct `ManagedClusterProperties`
- New field `Location` in struct `ManagedClustersClientGetCommandResultResponse`


## 2.2.0 (2022-10-26)
### Features Added

- New function `*ManagedClustersClient.BeginRotateServiceAccountSigningKeys(context.Context, string, string, *ManagedClustersClientBeginRotateServiceAccountSigningKeysOptions) (*runtime.Poller[ManagedClustersClientRotateServiceAccountSigningKeysResponse], error)`
- New struct `ManagedClusterOIDCIssuerProfile`
- New struct `ManagedClusterStorageProfileBlobCSIDriver`
- New struct `ManagedClustersClientBeginRotateServiceAccountSigningKeysOptions`
- New struct `ManagedClustersClientRotateServiceAccountSigningKeysResponse`
- New field `BlobCSIDriver` in struct `ManagedClusterStorageProfile`
- New field `OidcIssuerProfile` in struct `ManagedClusterProperties`


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
