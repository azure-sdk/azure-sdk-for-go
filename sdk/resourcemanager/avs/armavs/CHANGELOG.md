# Release History

## 2.0.0 (2022-10-10)
### Breaking Changes

- Struct `CloudError` has been removed

### Features Added

- New const `WorkloadNetworkDhcpProvisioningStateCanceled`
- New const `WorkloadNetworkVMGroupProvisioningStateCanceled`
- New const `NsxPublicIPQuotaRaisedEnumDisabled`
- New const `WorkloadNetworkSegmentProvisioningStateCanceled`
- New const `WorkloadNetworkDNSZoneProvisioningStateCanceled`
- New const `ScriptExecutionProvisioningStateCanceled`
- New const `AzureHybridBenefitTypeSQLHost`
- New const `PrivateCloudProvisioningStateCanceled`
- New const `AddonProvisioningStateCanceled`
- New const `ClusterProvisioningStateCanceled`
- New const `AffinityStrengthShould`
- New const `WorkloadNetworkNameDefault`
- New const `ExpressRouteAuthorizationProvisioningStateCanceled`
- New const `DatastoreProvisioningStateCanceled`
- New const `AddonTypeArc`
- New const `GlobalReachConnectionProvisioningStateCanceled`
- New const `AzureHybridBenefitTypeNone`
- New const `NsxPublicIPQuotaRaisedEnumEnabled`
- New const `PlacementPolicyProvisioningStateCanceled`
- New const `WorkloadNetworkPublicIPProvisioningStateCanceled`
- New const `WorkloadNetworkDNSServiceProvisioningStateCanceled`
- New const `WorkloadNetworkPortMirroringProvisioningStateCanceled`
- New const `AffinityStrengthMust`
- New type alias `AffinityStrength`
- New type alias `WorkloadNetworkName`
- New type alias `AzureHybridBenefitType`
- New type alias `NsxPublicIPQuotaRaisedEnum`
- New function `*ClustersClient.ListZones(context.Context, string, string, string, *ClustersClientListZonesOptions) (ClustersClientListZonesResponse, error)`
- New function `*WorkloadNetworksClient.Get(context.Context, string, string, WorkloadNetworkName, *WorkloadNetworksClientGetOptions) (WorkloadNetworksClientGetResponse, error)`
- New function `*WorkloadNetworksClient.NewListPager(string, string, *WorkloadNetworksClientListOptions) *runtime.Pager[WorkloadNetworksClientListResponse]`
- New function `PossibleAzureHybridBenefitTypeValues() []AzureHybridBenefitType`
- New function `PossibleNsxPublicIPQuotaRaisedEnumValues() []NsxPublicIPQuotaRaisedEnum`
- New function `PossibleAffinityStrengthValues() []AffinityStrength`
- New function `PossibleWorkloadNetworkNameValues() []WorkloadNetworkName`
- New function `*AddonArcProperties.GetAddonProperties() *AddonProperties`
- New struct `AddonArcProperties`
- New struct `ClusterZone`
- New struct `ClusterZoneList`
- New struct `ClustersClientListZonesOptions`
- New struct `ClustersClientListZonesResponse`
- New struct `WorkloadNetwork`
- New struct `WorkloadNetworkList`
- New struct `WorkloadNetworksClientGetOptions`
- New struct `WorkloadNetworksClientGetResponse`
- New struct `WorkloadNetworksClientListOptions`
- New struct `WorkloadNetworksClientListResponse`
- New field `AffinityStrength` in struct `PlacementPolicyUpdateProperties`
- New field `AzureHybridBenefitType` in struct `PlacementPolicyUpdateProperties`
- New field `SKU` in struct `LocationsClientCheckTrialAvailabilityOptions`
- New field `AutoDetectedKeyVersion` in struct `EncryptionKeyVaultProperties`
- New field `NsxPublicIPQuotaRaised` in struct `PrivateCloudProperties`
- New field `Company` in struct `ScriptPackageProperties`
- New field `URI` in struct `ScriptPackageProperties`
- New field `AffinityStrength` in struct `VMHostPlacementPolicyProperties`
- New field `AzureHybridBenefitType` in struct `VMHostPlacementPolicyProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).