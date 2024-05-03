# Release History

## 5.0.0 (2024-05-03)
### Breaking Changes

- Function `*ClientFactory.NewExtendedUeInformationClient` has been removed
- Function `*ClientFactory.NewUeInformationClient` has been removed
- Function `NewExtendedUeInformationClient` has been removed
- Function `*ExtendedUeInformationClient.Get` has been removed
- Function `*PacketCoreControlPlaneVersionsClient.Get` has been removed
- Function `*PacketCoreControlPlaneVersionsClient.NewListPager` has been removed
- Function `NewUeInformationClient` has been removed
- Function `*UeInformationClient.NewListPager` has been removed
- Struct `Arp` has been removed
- Struct `OperationList` has been removed
- Struct `UeInfoList` has been removed
- Field `OperationList` of struct `OperationsClientListResponse` has been removed

### Features Added

- New value `ProvisioningStateProvisioning`, `ProvisioningStateUpdating` added to enum type `ProvisioningState`
- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `ClusterType` with values `ClusterTypeAks`, `ClusterTypeNexusAks`
- New enum type `NASEncryptionType` with values `NASEncryptionTypeNEA0EEA0`, `NASEncryptionTypeNEA1EEA1`, `NASEncryptionTypeNEA2EEA2`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New enum type `SKUDeploymentType` with values `SKUDeploymentTypeLab`, `SKUDeploymentTypeProduction`
- New function `NewAmfDeploymentsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AmfDeploymentsClient, error)`
- New function `*AmfDeploymentsClient.BeginCreateOrUpdate(context.Context, string, string, AmfDeploymentResource, *AmfDeploymentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AmfDeploymentsClientCreateOrUpdateResponse], error)`
- New function `*AmfDeploymentsClient.BeginDelete(context.Context, string, string, *AmfDeploymentsClientBeginDeleteOptions) (*runtime.Poller[AmfDeploymentsClientDeleteResponse], error)`
- New function `*AmfDeploymentsClient.Get(context.Context, string, string, *AmfDeploymentsClientGetOptions) (AmfDeploymentsClientGetResponse, error)`
- New function `*AmfDeploymentsClient.NewListByResourceGroupPager(string, *AmfDeploymentsClientListByResourceGroupOptions) *runtime.Pager[AmfDeploymentsClientListByResourceGroupResponse]`
- New function `*AmfDeploymentsClient.NewListBySubscriptionPager(*AmfDeploymentsClientListBySubscriptionOptions) *runtime.Pager[AmfDeploymentsClientListBySubscriptionResponse]`
- New function `*AmfDeploymentsClient.UpdateTags(context.Context, string, string, AmfDeploymentResourceTagsUpdate, *AmfDeploymentsClientUpdateTagsOptions) (AmfDeploymentsClientUpdateTagsResponse, error)`
- New function `*ClientFactory.NewAmfDeploymentsClient() *AmfDeploymentsClient`
- New function `*ClientFactory.NewClusterServicesClient() *ClusterServicesClient`
- New function `*ClientFactory.NewExtendedUeInfosClient() *ExtendedUeInfosClient`
- New function `*ClientFactory.NewNrfDeploymentsClient() *NrfDeploymentsClient`
- New function `*ClientFactory.NewNssfDeploymentsClient() *NssfDeploymentsClient`
- New function `*ClientFactory.NewObservabilityServicesClient() *ObservabilityServicesClient`
- New function `*ClientFactory.NewPacketCoreControlPlaneVersionsTenantResourceClient() *PacketCoreControlPlaneVersionsTenantResourceClient`
- New function `*ClientFactory.NewRoutingInfoModelsClient() *RoutingInfoModelsClient`
- New function `*ClientFactory.NewSmfDeploymentsClient() *SmfDeploymentsClient`
- New function `*ClientFactory.NewUesClient() *UesClient`
- New function `*ClientFactory.NewUpfDeploymentsClient() *UpfDeploymentsClient`
- New function `*ClusterServiceAksClusterData.GetClusterServiceClusterTypeSpecificData() *ClusterServiceClusterTypeSpecificData`
- New function `*ClusterServiceClusterTypeSpecificData.GetClusterServiceClusterTypeSpecificData() *ClusterServiceClusterTypeSpecificData`
- New function `*ClusterServiceNexusAksClusterData.GetClusterServiceClusterTypeSpecificData() *ClusterServiceClusterTypeSpecificData`
- New function `NewClusterServicesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ClusterServicesClient, error)`
- New function `*ClusterServicesClient.BeginCreateOrUpdate(context.Context, string, string, ClusterServiceResource, *ClusterServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ClusterServicesClientCreateOrUpdateResponse], error)`
- New function `*ClusterServicesClient.BeginDelete(context.Context, string, string, *ClusterServicesClientBeginDeleteOptions) (*runtime.Poller[ClusterServicesClientDeleteResponse], error)`
- New function `*ClusterServicesClient.Get(context.Context, string, string, *ClusterServicesClientGetOptions) (ClusterServicesClientGetResponse, error)`
- New function `*ClusterServicesClient.NewListByResourceGroupPager(string, *ClusterServicesClientListByResourceGroupOptions) *runtime.Pager[ClusterServicesClientListByResourceGroupResponse]`
- New function `*ClusterServicesClient.NewListBySubscriptionPager(*ClusterServicesClientListBySubscriptionOptions) *runtime.Pager[ClusterServicesClientListBySubscriptionResponse]`
- New function `*ClusterServicesClient.UpdateTags(context.Context, string, string, ClusterServiceResourceTagsUpdate, *ClusterServicesClientUpdateTagsOptions) (ClusterServicesClientUpdateTagsResponse, error)`
- New function `NewExtendedUeInfosClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ExtendedUeInfosClient, error)`
- New function `*ExtendedUeInfosClient.Get(context.Context, string, string, string, *ExtendedUeInfosClientGetOptions) (ExtendedUeInfosClientGetResponse, error)`
- New function `*MobileNetworksClient.NewListSimGroupsPager(string, string, *MobileNetworksClientListSimGroupsOptions) *runtime.Pager[MobileNetworksClientListSimGroupsResponse]`
- New function `NewObservabilityServicesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ObservabilityServicesClient, error)`
- New function `*ObservabilityServicesClient.BeginCreateOrUpdate(context.Context, string, string, ObservabilityServiceResource, *ObservabilityServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ObservabilityServicesClientCreateOrUpdateResponse], error)`
- New function `*ObservabilityServicesClient.BeginDelete(context.Context, string, string, *ObservabilityServicesClientBeginDeleteOptions) (*runtime.Poller[ObservabilityServicesClientDeleteResponse], error)`
- New function `*ObservabilityServicesClient.Get(context.Context, string, string, *ObservabilityServicesClientGetOptions) (ObservabilityServicesClientGetResponse, error)`
- New function `*ObservabilityServicesClient.NewListByResourceGroupPager(string, *ObservabilityServicesClientListByResourceGroupOptions) *runtime.Pager[ObservabilityServicesClientListByResourceGroupResponse]`
- New function `*ObservabilityServicesClient.NewListBySubscriptionPager(*ObservabilityServicesClientListBySubscriptionOptions) *runtime.Pager[ObservabilityServicesClientListBySubscriptionResponse]`
- New function `*ObservabilityServicesClient.UpdateTags(context.Context, string, string, ObservabilityServiceResourceTagsUpdate, *ObservabilityServicesClientUpdateTagsOptions) (ObservabilityServicesClientUpdateTagsResponse, error)`
- New function `NewPacketCoreControlPlaneVersionsTenantResourceClient(azcore.TokenCredential, *arm.ClientOptions) (*PacketCoreControlPlaneVersionsTenantResourceClient, error)`
- New function `*PacketCoreControlPlaneVersionsTenantResourceClient.Get(context.Context, string, *PacketCoreControlPlaneVersionsTenantResourceClientGetOptions) (PacketCoreControlPlaneVersionsTenantResourceClientGetResponse, error)`
- New function `*PacketCoreControlPlaneVersionsTenantResourceClient.NewListByTenantPager(*PacketCoreControlPlaneVersionsTenantResourceClientListByTenantOptions) *runtime.Pager[PacketCoreControlPlaneVersionsTenantResourceClientListByTenantResponse]`
- New function `NewRoutingInfoModelsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*RoutingInfoModelsClient, error)`
- New function `*RoutingInfoModelsClient.Get(context.Context, string, string, *RoutingInfoModelsClientGetOptions) (RoutingInfoModelsClientGetResponse, error)`
- New function `*RoutingInfoModelsClient.NewListByPacketCoreControlPlanePager(string, string, *RoutingInfoModelsClientListByPacketCoreControlPlaneOptions) *runtime.Pager[RoutingInfoModelsClientListByPacketCoreControlPlaneResponse]`
- New function `*SimsClient.BeginClone(context.Context, string, string, SimClone, *SimsClientBeginCloneOptions) (*runtime.Poller[SimsClientCloneResponse], error)`
- New function `*SimsClient.BeginMove(context.Context, string, string, SimMove, *SimsClientBeginMoveOptions) (*runtime.Poller[SimsClientMoveResponse], error)`
- New function `NewSmfDeploymentsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SmfDeploymentsClient, error)`
- New function `*SmfDeploymentsClient.BeginCreateOrUpdate(context.Context, string, string, SmfDeploymentResource, *SmfDeploymentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SmfDeploymentsClientCreateOrUpdateResponse], error)`
- New function `*SmfDeploymentsClient.BeginDelete(context.Context, string, string, *SmfDeploymentsClientBeginDeleteOptions) (*runtime.Poller[SmfDeploymentsClientDeleteResponse], error)`
- New function `*SmfDeploymentsClient.Get(context.Context, string, string, *SmfDeploymentsClientGetOptions) (SmfDeploymentsClientGetResponse, error)`
- New function `*SmfDeploymentsClient.NewListByResourceGroupPager(string, *SmfDeploymentsClientListByResourceGroupOptions) *runtime.Pager[SmfDeploymentsClientListByResourceGroupResponse]`
- New function `*SmfDeploymentsClient.NewListBySubscriptionPager(*SmfDeploymentsClientListBySubscriptionOptions) *runtime.Pager[SmfDeploymentsClientListBySubscriptionResponse]`
- New function `*SmfDeploymentsClient.UpdateTags(context.Context, string, string, SmfDeploymentResourceTagsUpdate, *SmfDeploymentsClientUpdateTagsOptions) (SmfDeploymentsClientUpdateTagsResponse, error)`
- New function `NewUesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*UesClient, error)`
- New function `*UesClient.NewListByPacketCoreControlPlanePager(string, string, *UesClientListByPacketCoreControlPlaneOptions) *runtime.Pager[UesClientListByPacketCoreControlPlaneResponse]`
- New function `NewUpfDeploymentsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*UpfDeploymentsClient, error)`
- New function `*UpfDeploymentsClient.BeginCreateOrUpdate(context.Context, string, string, UpfDeploymentResource, *UpfDeploymentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[UpfDeploymentsClientCreateOrUpdateResponse], error)`
- New function `*UpfDeploymentsClient.BeginDelete(context.Context, string, string, *UpfDeploymentsClientBeginDeleteOptions) (*runtime.Poller[UpfDeploymentsClientDeleteResponse], error)`
- New function `*UpfDeploymentsClient.Get(context.Context, string, string, *UpfDeploymentsClientGetOptions) (UpfDeploymentsClientGetResponse, error)`
- New function `*UpfDeploymentsClient.NewListByResourceGroupPager(string, *UpfDeploymentsClientListByResourceGroupOptions) *runtime.Pager[UpfDeploymentsClientListByResourceGroupResponse]`
- New function `*UpfDeploymentsClient.NewListBySubscriptionPager(*UpfDeploymentsClientListBySubscriptionOptions) *runtime.Pager[UpfDeploymentsClientListBySubscriptionResponse]`
- New function `*UpfDeploymentsClient.UpdateTags(context.Context, string, string, UpfDeploymentResourceTagsUpdate, *UpfDeploymentsClientUpdateTagsOptions) (UpfDeploymentsClientUpdateTagsResponse, error)`
- New function `NewNrfDeploymentsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*NrfDeploymentsClient, error)`
- New function `*NrfDeploymentsClient.BeginCreateOrUpdate(context.Context, string, string, NrfDeploymentResource, *NrfDeploymentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[NrfDeploymentsClientCreateOrUpdateResponse], error)`
- New function `*NrfDeploymentsClient.BeginDelete(context.Context, string, string, *NrfDeploymentsClientBeginDeleteOptions) (*runtime.Poller[NrfDeploymentsClientDeleteResponse], error)`
- New function `*NrfDeploymentsClient.Get(context.Context, string, string, *NrfDeploymentsClientGetOptions) (NrfDeploymentsClientGetResponse, error)`
- New function `*NrfDeploymentsClient.NewListByResourceGroupPager(string, *NrfDeploymentsClientListByResourceGroupOptions) *runtime.Pager[NrfDeploymentsClientListByResourceGroupResponse]`
- New function `*NrfDeploymentsClient.NewListBySubscriptionPager(*NrfDeploymentsClientListBySubscriptionOptions) *runtime.Pager[NrfDeploymentsClientListBySubscriptionResponse]`
- New function `*NrfDeploymentsClient.UpdateTags(context.Context, string, string, NrfDeploymentResourceTagsUpdate, *NrfDeploymentsClientUpdateTagsOptions) (NrfDeploymentsClientUpdateTagsResponse, error)`
- New function `NewNssfDeploymentsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*NssfDeploymentsClient, error)`
- New function `*NssfDeploymentsClient.BeginCreateOrUpdate(context.Context, string, string, NssfDeploymentResource, *NssfDeploymentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[NssfDeploymentsClientCreateOrUpdateResponse], error)`
- New function `*NssfDeploymentsClient.BeginDelete(context.Context, string, string, *NssfDeploymentsClientBeginDeleteOptions) (*runtime.Poller[NssfDeploymentsClientDeleteResponse], error)`
- New function `*NssfDeploymentsClient.Get(context.Context, string, string, *NssfDeploymentsClientGetOptions) (NssfDeploymentsClientGetResponse, error)`
- New function `*NssfDeploymentsClient.NewListByResourceGroupPager(string, *NssfDeploymentsClientListByResourceGroupOptions) *runtime.Pager[NssfDeploymentsClientListByResourceGroupResponse]`
- New function `*NssfDeploymentsClient.NewListBySubscriptionPager(*NssfDeploymentsClientListBySubscriptionOptions) *runtime.Pager[NssfDeploymentsClientListBySubscriptionResponse]`
- New function `*NssfDeploymentsClient.UpdateTags(context.Context, string, string, NssfDeploymentResourceTagsUpdate, *NssfDeploymentsClientUpdateTagsOptions) (NssfDeploymentsClientUpdateTagsResponse, error)`
- New struct `AmfDeploymentResource`
- New struct `AmfDeploymentResourceListResult`
- New struct `AmfDeploymentResourceProperties`
- New struct `AmfDeploymentResourceTagsUpdate`
- New struct `ClusterServiceAksClusterData`
- New struct `ClusterServiceNexusAksClusterData`
- New struct `ClusterServiceResource`
- New struct `ClusterServiceResourceListResult`
- New struct `ClusterServiceResourceProperties`
- New struct `ClusterServiceResourceTagsUpdate`
- New struct `IPv4Route`
- New struct `IPv4RouteNextHop`
- New struct `NrfDeploymentResource`
- New struct `NrfDeploymentResourceListResult`
- New struct `NrfDeploymentResourceProperties`
- New struct `NrfDeploymentResourceTagsUpdate`
- New struct `NssfDeploymentResource`
- New struct `NssfDeploymentResourceListResult`
- New struct `NssfDeploymentResourceProperties`
- New struct `NssfDeploymentResourceTagsUpdate`
- New struct `ObservabilityServiceResource`
- New struct `ObservabilityServiceResourceListResult`
- New struct `ObservabilityServiceResourceProperties`
- New struct `ObservabilityServiceResourceTagsUpdate`
- New struct `OperationListResult`
- New struct `OperationalStatus`
- New struct `QualifiedComponentDeploymentParameters`
- New struct `RoutingInfoModel`
- New struct `RoutingInfoModelListResult`
- New struct `RoutingInfoPropertiesFormat`
- New struct `SimClone`
- New struct `SimMove`
- New struct `SmfDeploymentResource`
- New struct `SmfDeploymentResourceListResult`
- New struct `SmfDeploymentResourceProperties`
- New struct `SmfDeploymentResourceTagsUpdate`
- New struct `UeInfoListResult`
- New struct `UpfDeploymentResource`
- New struct `UpfDeploymentResourceListResult`
- New struct `UpfDeploymentResourceProperties`
- New struct `UpfDeploymentResourceTagsUpdate`
- New struct `UserConsentConfiguration`
- New struct `UserPlaneDataRoutesItem`
- New field `BfdIPv4Endpoints`, `IPv4AddressList`, `VlanID` in struct `InterfaceProperties`
- New field `ActionType`, `Origin` in struct `Operation`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New field `UserConsent` in struct `PacketCoreControlPlanePropertiesFormat`
- New field `HaUpgradesAvailable` in struct `Platform`
- New field `NasEncryption` in struct `SignalingConfiguration`
- New anonymous field `AsyncOperationStatus` in struct `SitesClientDeletePacketCoreResponse`


## 4.0.0 (2024-03-22)
### Breaking Changes

- Function `*MobileNetworksClient.UpdateTags` parameter(s) have been changed from `(context.Context, string, string, TagsObject, *MobileNetworksClientUpdateTagsOptions)` to `(context.Context, string, string, IdentityAndTagsObject, *MobileNetworksClientUpdateTagsOptions)`

### Features Added

- New enum type `HomeNetworkPrivateKeysProvisioningState` with values `HomeNetworkPrivateKeysProvisioningStateFailed`, `HomeNetworkPrivateKeysProvisioningStateNotProvisioned`, `HomeNetworkPrivateKeysProvisioningStateProvisioned`
- New enum type `PdnType` with values `PdnTypeIPV4`
- New enum type `RatType` with values `RatTypeFiveG`, `RatTypeFourG`
- New enum type `RrcEstablishmentCause` with values `RrcEstablishmentCauseEmergency`, `RrcEstablishmentCauseMobileOriginatedData`, `RrcEstablishmentCauseMobileOriginatedSignaling`, `RrcEstablishmentCauseMobileTerminatedData`, `RrcEstablishmentCauseMobileTerminatedSignaling`, `RrcEstablishmentCauseSMS`
- New enum type `UeState` with values `UeStateConnected`, `UeStateDeregistered`, `UeStateDetached`, `UeStateIdle`, `UeStateUnknown`
- New enum type `UeUsageSetting` with values `UeUsageSettingDataCentric`, `UeUsageSettingVoiceCentric`
- New function `*ClientFactory.NewExtendedUeInformationClient() *ExtendedUeInformationClient`
- New function `*ClientFactory.NewUeInformationClient() *UeInformationClient`
- New function `*ExtendedUeInfoProperties.GetExtendedUeInfoProperties() *ExtendedUeInfoProperties`
- New function `NewExtendedUeInformationClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ExtendedUeInformationClient, error)`
- New function `*ExtendedUeInformationClient.Get(context.Context, string, string, string, *ExtendedUeInformationClientGetOptions) (ExtendedUeInformationClientGetResponse, error)`
- New function `*UeInfo4G.GetExtendedUeInfoProperties() *ExtendedUeInfoProperties`
- New function `*UeInfo5G.GetExtendedUeInfoProperties() *ExtendedUeInfoProperties`
- New function `NewUeInformationClient(string, azcore.TokenCredential, *arm.ClientOptions) (*UeInformationClient, error)`
- New function `*UeInformationClient.NewListPager(string, string, *UeInformationClientListOptions) *runtime.Pager[UeInformationClientListResponse]`
- New struct `AmfID`
- New struct `DnnIPPair`
- New struct `ExtendedUeInfo`
- New struct `GNbID`
- New struct `GlobalRanNodeID`
- New struct `Guti4G`
- New struct `Guti5G`
- New struct `HomeNetworkPrivateKeysProvisioning`
- New struct `HomeNetworkPublicKey`
- New struct `MmeID`
- New struct `PublicLandMobileNetwork`
- New struct `PublicLandMobileNetworkHomeNetworkPublicKeys`
- New struct `UeConnectionInfo4G`
- New struct `UeConnectionInfo5G`
- New struct `UeIPAddress`
- New struct `UeInfo`
- New struct `UeInfo4G`
- New struct `UeInfo4GProperties`
- New struct `UeInfo5G`
- New struct `UeInfo5GProperties`
- New struct `UeInfoList`
- New struct `UeInfoPropertiesFormat`
- New struct `UeLocationInfo`
- New struct `UeQOSFlow`
- New struct `UeSessionInfo4G`
- New struct `UeSessionInfo5G`
- New field `Identity` in struct `MobileNetwork`
- New field `HomeNetworkPrivateKeysProvisioning` in struct `PacketCoreControlPlanePropertiesFormat`
- New field `PublicLandMobileNetworks` in struct `PropertiesFormat`


## 3.1.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.
- New value `InstallationReasonControlPlaneAccessInterfaceHasChanged`, `InstallationReasonControlPlaneAccessVirtualIPv4AddressesHasChanged`, `InstallationReasonPublicLandMobileNetworkIdentifierHasChanged`, `InstallationReasonUserPlaneAccessInterfaceHasChanged`, `InstallationReasonUserPlaneAccessVirtualIPv4AddressesHasChanged`, `InstallationReasonUserPlaneDataInterfaceHasChanged` added to enum type `InstallationReason`
- New struct `EventHubConfiguration`
- New struct `NASRerouteConfiguration`
- New struct `SignalingConfiguration`
- New field `OutputFiles` in struct `PacketCapturePropertiesFormat`
- New field `ControlPlaneAccessVirtualIPv4Addresses`, `EventHub`, `Signaling` in struct `PacketCoreControlPlanePropertiesFormat`
- New field `UserPlaneAccessVirtualIPv4Addresses` in struct `PacketCoreDataPlanePropertiesFormat`


## 3.0.0 (2023-07-28)
### Breaking Changes

- Function `*PacketCoreControlPlanesClient.UpdateTags` parameter(s) have been changed from `(context.Context, string, string, TagsObject, *PacketCoreControlPlanesClientUpdateTagsOptions)` to `(context.Context, string, string, IdentityAndTagsObject, *PacketCoreControlPlanesClientUpdateTagsOptions)`
- Function `*SimGroupsClient.UpdateTags` parameter(s) have been changed from `(context.Context, string, string, TagsObject, *SimGroupsClientUpdateTagsOptions)` to `(context.Context, string, string, IdentityAndTagsObject, *SimGroupsClientUpdateTagsOptions)`
- `BillingSKUG3`, `BillingSKUG4` from enum `BillingSKU` has been removed
- `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned` from enum `ManagedServiceIdentityType` has been removed
- Field `PrincipalID`, `TenantID` of struct `ManagedServiceIdentity` has been removed

### Features Added

- New value `CoreNetworkTypeEPC5GC` added to enum type `CoreNetworkType`
- New enum type `DesiredInstallationState` with values `DesiredInstallationStateInstalled`, `DesiredInstallationStateUninstalled`
- New enum type `DiagnosticsPackageStatus` with values `DiagnosticsPackageStatusCollected`, `DiagnosticsPackageStatusCollecting`, `DiagnosticsPackageStatusError`, `DiagnosticsPackageStatusNotStarted`
- New enum type `InstallationReason` with values `InstallationReasonNoAttachedDataNetworks`, `InstallationReasonNoPacketCoreDataPlane`, `InstallationReasonNoSlices`
- New enum type `PacketCaptureStatus` with values `PacketCaptureStatusError`, `PacketCaptureStatusNotStarted`, `PacketCaptureStatusRunning`, `PacketCaptureStatusStopped`
- New enum type `ReinstallRequired` with values `ReinstallRequiredNotRequired`, `ReinstallRequiredRequired`
- New function `*ClientFactory.NewDiagnosticsPackagesClient() *DiagnosticsPackagesClient`
- New function `*ClientFactory.NewPacketCapturesClient() *PacketCapturesClient`
- New function `NewDiagnosticsPackagesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DiagnosticsPackagesClient, error)`
- New function `*DiagnosticsPackagesClient.BeginCreateOrUpdate(context.Context, string, string, string, *DiagnosticsPackagesClientBeginCreateOrUpdateOptions) (*runtime.Poller[DiagnosticsPackagesClientCreateOrUpdateResponse], error)`
- New function `*DiagnosticsPackagesClient.BeginDelete(context.Context, string, string, string, *DiagnosticsPackagesClientBeginDeleteOptions) (*runtime.Poller[DiagnosticsPackagesClientDeleteResponse], error)`
- New function `*DiagnosticsPackagesClient.Get(context.Context, string, string, string, *DiagnosticsPackagesClientGetOptions) (DiagnosticsPackagesClientGetResponse, error)`
- New function `*DiagnosticsPackagesClient.NewListByPacketCoreControlPlanePager(string, string, *DiagnosticsPackagesClientListByPacketCoreControlPlaneOptions) *runtime.Pager[DiagnosticsPackagesClientListByPacketCoreControlPlaneResponse]`
- New function `NewPacketCapturesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PacketCapturesClient, error)`
- New function `*PacketCapturesClient.BeginCreateOrUpdate(context.Context, string, string, string, PacketCapture, *PacketCapturesClientBeginCreateOrUpdateOptions) (*runtime.Poller[PacketCapturesClientCreateOrUpdateResponse], error)`
- New function `*PacketCapturesClient.BeginDelete(context.Context, string, string, string, *PacketCapturesClientBeginDeleteOptions) (*runtime.Poller[PacketCapturesClientDeleteResponse], error)`
- New function `*PacketCapturesClient.Get(context.Context, string, string, string, *PacketCapturesClientGetOptions) (PacketCapturesClientGetResponse, error)`
- New function `*PacketCapturesClient.NewListByPacketCoreControlPlanePager(string, string, *PacketCapturesClientListByPacketCoreControlPlaneOptions) *runtime.Pager[PacketCapturesClientListByPacketCoreControlPlaneResponse]`
- New function `*PacketCapturesClient.BeginStop(context.Context, string, string, string, *PacketCapturesClientBeginStopOptions) (*runtime.Poller[PacketCapturesClientStopResponse], error)`
- New function `*PacketCoreControlPlaneVersionsClient.GetBySubscription(context.Context, string, string, *PacketCoreControlPlaneVersionsClientGetBySubscriptionOptions) (PacketCoreControlPlaneVersionsClientGetBySubscriptionResponse, error)`
- New function `*PacketCoreControlPlaneVersionsClient.NewListBySubscriptionPager(string, *PacketCoreControlPlaneVersionsClientListBySubscriptionOptions) *runtime.Pager[PacketCoreControlPlaneVersionsClientListBySubscriptionResponse]`
- New function `*SitesClient.BeginDeletePacketCore(context.Context, string, string, string, SiteDeletePacketCore, *SitesClientBeginDeletePacketCoreOptions) (*runtime.Poller[SitesClientDeletePacketCoreResponse], error)`
- New struct `DiagnosticsPackage`
- New struct `DiagnosticsPackageListResult`
- New struct `DiagnosticsPackagePropertiesFormat`
- New struct `DiagnosticsUploadConfiguration`
- New struct `IdentityAndTagsObject`
- New struct `PacketCapture`
- New struct `PacketCaptureListResult`
- New struct `PacketCapturePropertiesFormat`
- New struct `PacketCoreControlPlaneResourceID`
- New struct `SiteDeletePacketCore`
- New field `DesiredState`, `Reasons`, `ReinstallRequired` in struct `Installation`
- New field `DiagnosticsUpload`, `InstalledVersion` in struct `PacketCoreControlPlanePropertiesFormat`


## 2.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 2.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 2.0.0 (2023-01-27)
### Breaking Changes

- Function `NewPacketCoreControlPlaneClient` has been removed
- Function `*PacketCoreControlPlaneClient.BeginCollectDiagnosticsPackage` has been removed
- Function `*PacketCoreControlPlaneClient.BeginReinstall` has been removed
- Function `*PacketCoreControlPlaneClient.BeginRollback` has been removed
- Function `NewSimClient` has been removed
- Function `*SimClient.BeginBulkDelete` has been removed
- Function `*SimClient.BeginBulkUpload` has been removed
- Function `*SimClient.BeginBulkUploadEncrypted` has been removed
- Struct `PacketCoreControlPlaneClient` has been removed
- Struct `PacketCoreControlPlaneClientCollectDiagnosticsPackageResponse` has been removed
- Struct `PacketCoreControlPlaneClientReinstallResponse` has been removed
- Struct `PacketCoreControlPlaneClientRollbackResponse` has been removed
- Struct `SimClient` has been removed
- Struct `SimClientBulkDeleteResponse` has been removed
- Struct `SimClientBulkUploadEncryptedResponse` has been removed
- Struct `SimClientBulkUploadResponse` has been removed

### Features Added

- New function `*PacketCoreControlPlanesClient.BeginCollectDiagnosticsPackage(context.Context, string, string, PacketCoreControlPlaneCollectDiagnosticsPackage, *PacketCoreControlPlanesClientBeginCollectDiagnosticsPackageOptions) (*runtime.Poller[PacketCoreControlPlanesClientCollectDiagnosticsPackageResponse], error)`
- New function `*PacketCoreControlPlanesClient.BeginReinstall(context.Context, string, string, *PacketCoreControlPlanesClientBeginReinstallOptions) (*runtime.Poller[PacketCoreControlPlanesClientReinstallResponse], error)`
- New function `*PacketCoreControlPlanesClient.BeginRollback(context.Context, string, string, *PacketCoreControlPlanesClientBeginRollbackOptions) (*runtime.Poller[PacketCoreControlPlanesClientRollbackResponse], error)`
- New function `*SimsClient.BeginBulkDelete(context.Context, string, string, SimDeleteList, *SimsClientBeginBulkDeleteOptions) (*runtime.Poller[SimsClientBulkDeleteResponse], error)`
- New function `*SimsClient.BeginBulkUpload(context.Context, string, string, SimUploadList, *SimsClientBeginBulkUploadOptions) (*runtime.Poller[SimsClientBulkUploadResponse], error)`
- New function `*SimsClient.BeginBulkUploadEncrypted(context.Context, string, string, EncryptedSimUploadList, *SimsClientBeginBulkUploadEncryptedOptions) (*runtime.Poller[SimsClientBulkUploadEncryptedResponse], error)`
- New struct `PacketCoreControlPlanesClientCollectDiagnosticsPackageResponse`
- New struct `PacketCoreControlPlanesClientReinstallResponse`
- New struct `PacketCoreControlPlanesClientRollbackResponse`
- New struct `SimsClientBulkDeleteResponse`
- New struct `SimsClientBulkUploadEncryptedResponse`
- New struct `SimsClientBulkUploadResponse`


## 1.0.0 (2022-12-23)
### Breaking Changes

- Type of `LocalDiagnosticsAccessConfiguration.HTTPSServerCertificate` has been changed from `*KeyVaultCertificate` to `*HTTPSServerCertificate`
- Const `BillingSKUEdgeSite2GBPS`, `BillingSKUEdgeSite3GBPS`, `BillingSKUEdgeSite4GBPS`, `BillingSKUEvaluationPackage`, `BillingSKUFlagshipStarterPackage`, `BillingSKULargePackage`, `BillingSKUMediumPackage` from type alias `BillingSKU` has been removed
- Const `PlatformTypeBaseVM` from type alias `PlatformType` has been removed
- Function `*MobileNetworksClient.BeginListSimIDs` has been removed
- Function `*PacketCoreControlPlaneVersionsClient.NewListByResourceGroupPager` has been removed
- Function `*SimsClient.NewListBySimGroupPager` has been removed
- Struct `KeyVaultCertificate` has been removed
- Struct `MobileNetworksClientListSimIDsResponse` has been removed
- Struct `PacketCoreControlPlaneVersionsClientListByResourceGroupResponse` has been removed
- Struct `SimIDListResult` has been removed
- Struct `SimsClientListBySimGroupResponse` has been removed
- Field `MobileNetwork` of struct `PacketCoreControlPlanePropertiesFormat` has been removed
- Field `RecommendedVersion` of struct `PacketCoreControlPlaneVersionPropertiesFormat` has been removed
- Field `VersionState` of struct `PacketCoreControlPlaneVersionPropertiesFormat` has been removed

### Features Added

- New value `BillingSKUG0`, `BillingSKUG1`, `BillingSKUG10`, `BillingSKUG2`, `BillingSKUG3`, `BillingSKUG4`, `BillingSKUG5` added to type alias `BillingSKU`
- New value `PlatformTypeThreePAZURESTACKHCI` added to type alias `PlatformType`
- New type alias `AuthenticationType` with values `AuthenticationTypeAAD`, `AuthenticationTypePassword`
- New type alias `CertificateProvisioningState` with values `CertificateProvisioningStateFailed`, `CertificateProvisioningStateNotProvisioned`, `CertificateProvisioningStateProvisioned`
- New type alias `InstallationState` with values `InstallationStateFailed`, `InstallationStateInstalled`, `InstallationStateInstalling`, `InstallationStateReinstalling`, `InstallationStateRollingBack`, `InstallationStateUninstalled`, `InstallationStateUninstalling`, `InstallationStateUpdating`, `InstallationStateUpgrading`
- New type alias `ObsoleteVersion` with values `ObsoleteVersionNotObsolete`, `ObsoleteVersionObsolete`
- New type alias `SiteProvisioningState` with values `SiteProvisioningStateAdding`, `SiteProvisioningStateDeleting`, `SiteProvisioningStateFailed`, `SiteProvisioningStateNotApplicable`, `SiteProvisioningStateProvisioned`, `SiteProvisioningStateUpdating`
- New function `NewPacketCoreControlPlaneClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PacketCoreControlPlaneClient, error)`
- New function `*PacketCoreControlPlaneClient.BeginCollectDiagnosticsPackage(context.Context, string, string, PacketCoreControlPlaneCollectDiagnosticsPackage, *PacketCoreControlPlaneClientBeginCollectDiagnosticsPackageOptions) (*runtime.Poller[PacketCoreControlPlaneClientCollectDiagnosticsPackageResponse], error)`
- New function `*PacketCoreControlPlaneClient.BeginReinstall(context.Context, string, string, *PacketCoreControlPlaneClientBeginReinstallOptions) (*runtime.Poller[PacketCoreControlPlaneClientReinstallResponse], error)`
- New function `*PacketCoreControlPlaneClient.BeginRollback(context.Context, string, string, *PacketCoreControlPlaneClientBeginRollbackOptions) (*runtime.Poller[PacketCoreControlPlaneClientRollbackResponse], error)`
- New function `*PacketCoreControlPlaneVersionsClient.NewListPager(*PacketCoreControlPlaneVersionsClientListOptions) *runtime.Pager[PacketCoreControlPlaneVersionsClientListResponse]`
- New function `NewSimClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SimClient, error)`
- New function `*SimClient.BeginBulkDelete(context.Context, string, string, SimDeleteList, *SimClientBeginBulkDeleteOptions) (*runtime.Poller[SimClientBulkDeleteResponse], error)`
- New function `*SimClient.BeginBulkUpload(context.Context, string, string, SimUploadList, *SimClientBeginBulkUploadOptions) (*runtime.Poller[SimClientBulkUploadResponse], error)`
- New function `*SimClient.BeginBulkUploadEncrypted(context.Context, string, string, EncryptedSimUploadList, *SimClientBeginBulkUploadEncryptedOptions) (*runtime.Poller[SimClientBulkUploadEncryptedResponse], error)`
- New function `*SimsClient.NewListByGroupPager(string, string, *SimsClientListByGroupOptions) *runtime.Pager[SimsClientListByGroupResponse]`
- New struct `AsyncOperationID`
- New struct `AsyncOperationStatus`
- New struct `AzureStackHCIClusterResourceID`
- New struct `CertificateProvisioning`
- New struct `CommonSimPropertiesFormat`
- New struct `EncryptedSimPropertiesFormat`
- New struct `EncryptedSimUploadList`
- New struct `HTTPSServerCertificate`
- New struct `Installation`
- New struct `PacketCoreControlPlaneClient`
- New struct `PacketCoreControlPlaneClientCollectDiagnosticsPackageResponse`
- New struct `PacketCoreControlPlaneClientReinstallResponse`
- New struct `PacketCoreControlPlaneClientRollbackResponse`
- New struct `PacketCoreControlPlaneCollectDiagnosticsPackage`
- New struct `PacketCoreControlPlaneVersionsClientListResponse`
- New struct `Platform`
- New struct `SimClient`
- New struct `SimClientBulkDeleteResponse`
- New struct `SimClientBulkUploadEncryptedResponse`
- New struct `SimClientBulkUploadResponse`
- New struct `SimDeleteList`
- New struct `SimNameAndEncryptedProperties`
- New struct `SimNameAndProperties`
- New struct `SimUploadList`
- New struct `SimsClientListByGroupResponse`
- New struct `SiteResourceID`
- New field `MaximumNumberOfBufferedPackets` in struct `DataNetworkConfiguration`
- New field `AuthenticationType` in struct `LocalDiagnosticsAccessConfiguration`
- New field `Installation` in struct `PacketCoreControlPlanePropertiesFormat`
- New field `RollbackVersion` in struct `PacketCoreControlPlanePropertiesFormat`
- New field `Sites` in struct `PacketCoreControlPlanePropertiesFormat`
- New field `UeMtu` in struct `PacketCoreControlPlanePropertiesFormat`
- New field `Platforms` in struct `PacketCoreControlPlaneVersionPropertiesFormat`
- New field `AzureStackEdgeDevices` in struct `PlatformConfiguration`
- New field `AzureStackHciCluster` in struct `PlatformConfiguration`
- New field `SiteProvisioningState` in struct `SimPolicyPropertiesFormat`
- New field `SiteProvisioningState` in struct `SimPropertiesFormat`
- New field `VendorKeyFingerprint` in struct `SimPropertiesFormat`
- New field `VendorName` in struct `SimPropertiesFormat`


## 0.6.0 (2022-07-21)
### Breaking Changes

- Function `*SimsClient.Get` parameter(s) have been changed from `(context.Context, string, string, *SimsClientGetOptions)` to `(context.Context, string, string, string, *SimsClientGetOptions)`
- Function `*SimsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, *SimsClientBeginDeleteOptions)` to `(context.Context, string, string, string, *SimsClientBeginDeleteOptions)`
- Function `*SimsClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, Sim, *SimsClientBeginCreateOrUpdateOptions)` to `(context.Context, string, string, string, Sim, *SimsClientBeginCreateOrUpdateOptions)`
- Function `*SimsClient.NewListByResourceGroupPager` has been removed
- Function `*SimsClient.NewListBySubscriptionPager` has been removed
- Function `*SimsClient.UpdateTags` has been removed
- Struct `SimsClientListByResourceGroupOptions` has been removed
- Struct `SimsClientListByResourceGroupResponse` has been removed
- Struct `SimsClientListBySubscriptionOptions` has been removed
- Struct `SimsClientListBySubscriptionResponse` has been removed
- Struct `SimsClientUpdateTagsOptions` has been removed
- Struct `SimsClientUpdateTagsResponse` has been removed
- Field `MobileNetwork` of struct `SimPropertiesFormat` has been removed
- Field `Tags` of struct `Sim` has been removed
- Field `Location` of struct `Sim` has been removed
- Field `CustomLocation` of struct `PacketCoreControlPlanePropertiesFormat` has been removed

### Features Added

- New const `ManagedServiceIdentityTypeSystemAssigned`
- New const `VersionStateValidationFailed`
- New const `VersionStateUnknown`
- New const `ManagedServiceIdentityTypeNone`
- New const `RecommendedVersionNotRecommended`
- New const `BillingSKUEdgeSite2GBPS`
- New const `BillingSKULargePackage`
- New const `VersionStateActive`
- New const `PlatformTypeAKSHCI`
- New const `BillingSKUFlagshipStarterPackage`
- New const `VersionStateDeprecated`
- New const `BillingSKUEdgeSite4GBPS`
- New const `BillingSKUEdgeSite3GBPS`
- New const `PlatformTypeBaseVM`
- New const `BillingSKUEvaluationPackage`
- New const `ManagedServiceIdentityTypeSystemAssignedUserAssigned`
- New const `BillingSKUMediumPackage`
- New const `ManagedServiceIdentityTypeUserAssigned`
- New const `RecommendedVersionRecommended`
- New const `VersionStatePreview`
- New const `VersionStateValidating`
- New function `PossibleRecommendedVersionValues() []RecommendedVersion`
- New function `PossibleBillingSKUValues() []BillingSKU`
- New function `*SimGroupsClient.UpdateTags(context.Context, string, string, TagsObject, *SimGroupsClientUpdateTagsOptions) (SimGroupsClientUpdateTagsResponse, error)`
- New function `*PacketCoreControlPlaneVersionsClient.NewListByResourceGroupPager(*PacketCoreControlPlaneVersionsClientListByResourceGroupOptions) *runtime.Pager[PacketCoreControlPlaneVersionsClientListByResourceGroupResponse]`
- New function `*SimGroupsClient.NewListByResourceGroupPager(string, *SimGroupsClientListByResourceGroupOptions) *runtime.Pager[SimGroupsClientListByResourceGroupResponse]`
- New function `*PacketCoreControlPlaneVersionsClient.Get(context.Context, string, *PacketCoreControlPlaneVersionsClientGetOptions) (PacketCoreControlPlaneVersionsClientGetResponse, error)`
- New function `*SimGroupsClient.Get(context.Context, string, string, *SimGroupsClientGetOptions) (SimGroupsClientGetResponse, error)`
- New function `PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType`
- New function `*SimGroupsClient.NewListBySubscriptionPager(*SimGroupsClientListBySubscriptionOptions) *runtime.Pager[SimGroupsClientListBySubscriptionResponse]`
- New function `*SimGroupsClient.BeginDelete(context.Context, string, string, *SimGroupsClientBeginDeleteOptions) (*runtime.Poller[SimGroupsClientDeleteResponse], error)`
- New function `NewPacketCoreControlPlaneVersionsClient(azcore.TokenCredential, *arm.ClientOptions) (*PacketCoreControlPlaneVersionsClient, error)`
- New function `PossibleVersionStateValues() []VersionState`
- New function `*SimsClient.NewListBySimGroupPager(string, string, *SimsClientListBySimGroupOptions) *runtime.Pager[SimsClientListBySimGroupResponse]`
- New function `NewSimGroupsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SimGroupsClient, error)`
- New function `PossiblePlatformTypeValues() []PlatformType`
- New function `*SimGroupsClient.BeginCreateOrUpdate(context.Context, string, string, SimGroup, *SimGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SimGroupsClientCreateOrUpdateResponse], error)`
- New struct `AzureStackEdgeDeviceResourceID`
- New struct `ConnectedClusterResourceID`
- New struct `KeyVaultCertificate`
- New struct `KeyVaultKey`
- New struct `LocalDiagnosticsAccessConfiguration`
- New struct `ManagedServiceIdentity`
- New struct `PacketCoreControlPlaneVersion`
- New struct `PacketCoreControlPlaneVersionListResult`
- New struct `PacketCoreControlPlaneVersionPropertiesFormat`
- New struct `PacketCoreControlPlaneVersionsClient`
- New struct `PacketCoreControlPlaneVersionsClientGetOptions`
- New struct `PacketCoreControlPlaneVersionsClientGetResponse`
- New struct `PacketCoreControlPlaneVersionsClientListByResourceGroupOptions`
- New struct `PacketCoreControlPlaneVersionsClientListByResourceGroupResponse`
- New struct `PlatformConfiguration`
- New struct `ProxyResource`
- New struct `SimGroup`
- New struct `SimGroupListResult`
- New struct `SimGroupPropertiesFormat`
- New struct `SimGroupResourceID`
- New struct `SimGroupsClient`
- New struct `SimGroupsClientBeginCreateOrUpdateOptions`
- New struct `SimGroupsClientBeginDeleteOptions`
- New struct `SimGroupsClientCreateOrUpdateResponse`
- New struct `SimGroupsClientDeleteResponse`
- New struct `SimGroupsClientGetOptions`
- New struct `SimGroupsClientGetResponse`
- New struct `SimGroupsClientListByResourceGroupOptions`
- New struct `SimGroupsClientListByResourceGroupResponse`
- New struct `SimGroupsClientListBySubscriptionOptions`
- New struct `SimGroupsClientListBySubscriptionResponse`
- New struct `SimGroupsClientUpdateTagsOptions`
- New struct `SimGroupsClientUpdateTagsResponse`
- New struct `SimsClientListBySimGroupOptions`
- New struct `SimsClientListBySimGroupResponse`
- New struct `UserAssignedIdentity`
- New field `Identity` in struct `PacketCoreControlPlane`
- New field `InteropSettings` in struct `PacketCoreControlPlanePropertiesFormat`
- New field `SKU` in struct `PacketCoreControlPlanePropertiesFormat`
- New field `LocalDiagnosticsAccess` in struct `PacketCoreControlPlanePropertiesFormat`
- New field `Platform` in struct `PacketCoreControlPlanePropertiesFormat`
- New field `DNSAddresses` in struct `AttachedDataNetworkPropertiesFormat`


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).