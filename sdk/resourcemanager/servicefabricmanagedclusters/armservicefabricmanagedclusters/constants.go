//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmanagedclusters

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
	moduleVersion = "v0.3.0"
)

// Access - The network traffic is allowed or denied.
type Access string

const (
	AccessAllow Access = "allow"
	AccessDeny  Access = "deny"
)

// PossibleAccessValues returns the possible values for the Access const type.
func PossibleAccessValues() []Access {
	return []Access{
		AccessAllow,
		AccessDeny,
	}
}

// AutoGeneratedDomainNameLabelScope - This enum is the entrypoint to using a certificate from a public CA for your cluster.
// This property was introduced to solve the domain squatting problem with new domains. A domain name will be
// generated in the following format: .... The hash portion comes from Azure DNS' Deterministic Name Library. The library
// creates a hash using the cluster's Tenant, Subscription, Resource Group and
// Resource Name using the AutoGeneratedDomainNameLabelScope/reuse policy chosen.
type AutoGeneratedDomainNameLabelScope string

const (
	// AutoGeneratedDomainNameLabelScopeNoReuse - NoReuse will create a new hash regardless of the Subscription, Resource Group,
	// Tenant and Resource name.
	AutoGeneratedDomainNameLabelScopeNoReuse AutoGeneratedDomainNameLabelScope = "NoReuse"
	// AutoGeneratedDomainNameLabelScopeResourceGroupReuse - ResourceGroupReuse allows for the same hash to be created if the
	// resource is created in the same Resource Group with the same resource name.
	AutoGeneratedDomainNameLabelScopeResourceGroupReuse AutoGeneratedDomainNameLabelScope = "ResourceGroupReuse"
	// AutoGeneratedDomainNameLabelScopeSubscriptionReuse - SubscriptionReuse allows for the same hash to be created if the resource
	// is created in the same Subscription with the same resource name.
	AutoGeneratedDomainNameLabelScopeSubscriptionReuse AutoGeneratedDomainNameLabelScope = "SubscriptionReuse"
	// AutoGeneratedDomainNameLabelScopeTenantReuse - TenantReuse allows for the same hash to be created if the resource is created
	// in the same Tenant with the same resource name.
	AutoGeneratedDomainNameLabelScopeTenantReuse AutoGeneratedDomainNameLabelScope = "TenantReuse"
)

// PossibleAutoGeneratedDomainNameLabelScopeValues returns the possible values for the AutoGeneratedDomainNameLabelScope const type.
func PossibleAutoGeneratedDomainNameLabelScopeValues() []AutoGeneratedDomainNameLabelScope {
	return []AutoGeneratedDomainNameLabelScope{
		AutoGeneratedDomainNameLabelScopeNoReuse,
		AutoGeneratedDomainNameLabelScopeResourceGroupReuse,
		AutoGeneratedDomainNameLabelScopeSubscriptionReuse,
		AutoGeneratedDomainNameLabelScopeTenantReuse,
	}
}

// ClusterState - The current state of the cluster.
type ClusterState string

const (
	// ClusterStateBaselineUpgrade - Indicates that the cluster is upgrading to establishes the cluster version. This upgrade
	// is automatically initiated when the cluster boots up for the first time.
	ClusterStateBaselineUpgrade ClusterState = "BaselineUpgrade"
	// ClusterStateDeploying - Indicates that the Service Fabric runtime is being installed on the VMs. Cluster resource will
	// be in this state until the cluster boots up and system services are up.
	ClusterStateDeploying ClusterState = "Deploying"
	// ClusterStateReady - Indicates that the cluster is in a stable state.
	ClusterStateReady ClusterState = "Ready"
	// ClusterStateUpgradeFailed - Indicates that the last upgrade for the cluster has failed.
	ClusterStateUpgradeFailed ClusterState = "UpgradeFailed"
	// ClusterStateUpgrading - Indicates that the cluster is being upgraded with the user provided configuration.
	ClusterStateUpgrading ClusterState = "Upgrading"
	// ClusterStateWaitingForNodes - Indicates that the cluster resource is created and the resource provider is waiting for Service
	// Fabric VM extension to boot up and report to it.
	ClusterStateWaitingForNodes ClusterState = "WaitingForNodes"
)

// PossibleClusterStateValues returns the possible values for the ClusterState const type.
func PossibleClusterStateValues() []ClusterState {
	return []ClusterState{
		ClusterStateBaselineUpgrade,
		ClusterStateDeploying,
		ClusterStateReady,
		ClusterStateUpgradeFailed,
		ClusterStateUpgrading,
		ClusterStateWaitingForNodes,
	}
}

// ClusterUpgradeCadence - Indicates when new cluster runtime version upgrades will be applied after they are released. By
// default is Wave0.
type ClusterUpgradeCadence string

const (
	// ClusterUpgradeCadenceWave0 - Cluster upgrade starts immediately after a new version is rolled out. Recommended for Test/Dev
	// clusters.
	ClusterUpgradeCadenceWave0 ClusterUpgradeCadence = "Wave0"
	// ClusterUpgradeCadenceWave1 - Cluster upgrade starts 7 days after a new version is rolled out. Recommended for Pre-prod
	// clusters.
	ClusterUpgradeCadenceWave1 ClusterUpgradeCadence = "Wave1"
	// ClusterUpgradeCadenceWave2 - Cluster upgrade starts 14 days after a new version is rolled out. Recommended for Production
	// clusters.
	ClusterUpgradeCadenceWave2 ClusterUpgradeCadence = "Wave2"
)

// PossibleClusterUpgradeCadenceValues returns the possible values for the ClusterUpgradeCadence const type.
func PossibleClusterUpgradeCadenceValues() []ClusterUpgradeCadence {
	return []ClusterUpgradeCadence{
		ClusterUpgradeCadenceWave0,
		ClusterUpgradeCadenceWave1,
		ClusterUpgradeCadenceWave2,
	}
}

// ClusterUpgradeMode - The upgrade mode of the cluster when new Service Fabric runtime version is available.
type ClusterUpgradeMode string

const (
	// ClusterUpgradeModeAutomatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version,
	// **clusterUpgradeCadence** will determine when the upgrade starts after the new version becomes available.
	ClusterUpgradeModeAutomatic ClusterUpgradeMode = "Automatic"
	// ClusterUpgradeModeManual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version.
	// The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
	ClusterUpgradeModeManual ClusterUpgradeMode = "Manual"
)

// PossibleClusterUpgradeModeValues returns the possible values for the ClusterUpgradeMode const type.
func PossibleClusterUpgradeModeValues() []ClusterUpgradeMode {
	return []ClusterUpgradeMode{
		ClusterUpgradeModeAutomatic,
		ClusterUpgradeModeManual,
	}
}

// Direction - Network security rule direction.
type Direction string

const (
	DirectionInbound  Direction = "inbound"
	DirectionOutbound Direction = "outbound"
)

// PossibleDirectionValues returns the possible values for the Direction const type.
func PossibleDirectionValues() []Direction {
	return []Direction{
		DirectionInbound,
		DirectionOutbound,
	}
}

// DiskType - Managed data disk type. IOPS and throughput are given by the disk size, to see more information go to https://docs.microsoft.com/en-us/azure/virtual-machines/disks-types.
type DiskType string

const (
	// DiskTypePremiumLRS - Premium SSD locally redundant storage. Best for production and performance sensitive workloads.
	DiskTypePremiumLRS DiskType = "Premium_LRS"
	// DiskTypeStandardLRS - Standard HDD locally redundant storage. Best for backup, non-critical, and infrequent access.
	DiskTypeStandardLRS DiskType = "Standard_LRS"
	// DiskTypeStandardSSDLRS - Standard SSD locally redundant storage. Best for web servers, lightly used enterprise applications
	// and dev/test.
	DiskTypeStandardSSDLRS DiskType = "StandardSSD_LRS"
)

// PossibleDiskTypeValues returns the possible values for the DiskType const type.
func PossibleDiskTypeValues() []DiskType {
	return []DiskType{
		DiskTypePremiumLRS,
		DiskTypeStandardLRS,
		DiskTypeStandardSSDLRS,
	}
}

// EvictionPolicyType - Specifies the eviction policy for virtual machines in a SPOT node type.
type EvictionPolicyType string

const (
	// EvictionPolicyTypeDeallocate - Eviction policy will be Deallocate for SPOT vms.
	EvictionPolicyTypeDeallocate EvictionPolicyType = "Deallocate"
	// EvictionPolicyTypeDelete - Eviction policy will be Delete for SPOT vms.
	EvictionPolicyTypeDelete EvictionPolicyType = "Delete"
)

// PossibleEvictionPolicyTypeValues returns the possible values for the EvictionPolicyType const type.
func PossibleEvictionPolicyTypeValues() []EvictionPolicyType {
	return []EvictionPolicyType{
		EvictionPolicyTypeDeallocate,
		EvictionPolicyTypeDelete,
	}
}

// FailureAction - The compensating action to perform when a Monitored upgrade encounters monitoring policy or health policy
// violations. Rollback specifies that the upgrade will start rolling back automatically. Manual
// indicates that the upgrade will switch to UnmonitoredManual upgrade mode.
type FailureAction string

const (
	// FailureActionManual - The upgrade will switch to UnmonitoredManual upgrade mode. The value is 1
	FailureActionManual FailureAction = "Manual"
	// FailureActionRollback - The upgrade will start rolling back automatically. The value is 0
	FailureActionRollback FailureAction = "Rollback"
)

// PossibleFailureActionValues returns the possible values for the FailureAction const type.
func PossibleFailureActionValues() []FailureAction {
	return []FailureAction{
		FailureActionManual,
		FailureActionRollback,
	}
}

// IPAddressType - The IP address type.
type IPAddressType string

const (
	// IPAddressTypeIPv4 - IPv4 address type.
	IPAddressTypeIPv4 IPAddressType = "IPv4"
	// IPAddressTypeIPv6 - IPv6 address type.
	IPAddressTypeIPv6 IPAddressType = "IPv6"
)

// PossibleIPAddressTypeValues returns the possible values for the IPAddressType const type.
func PossibleIPAddressTypeValues() []IPAddressType {
	return []IPAddressType{
		IPAddressTypeIPv4,
		IPAddressTypeIPv6,
	}
}

// ManagedClusterAddOnFeature - Available cluster add-on features
type ManagedClusterAddOnFeature string

const (
	// ManagedClusterAddOnFeatureBackupRestoreService - Backup and restore service
	ManagedClusterAddOnFeatureBackupRestoreService ManagedClusterAddOnFeature = "BackupRestoreService"
	// ManagedClusterAddOnFeatureDNSService - Dns service
	ManagedClusterAddOnFeatureDNSService ManagedClusterAddOnFeature = "DnsService"
	// ManagedClusterAddOnFeatureResourceMonitorService - Resource monitor service
	ManagedClusterAddOnFeatureResourceMonitorService ManagedClusterAddOnFeature = "ResourceMonitorService"
)

// PossibleManagedClusterAddOnFeatureValues returns the possible values for the ManagedClusterAddOnFeature const type.
func PossibleManagedClusterAddOnFeatureValues() []ManagedClusterAddOnFeature {
	return []ManagedClusterAddOnFeature{
		ManagedClusterAddOnFeatureBackupRestoreService,
		ManagedClusterAddOnFeatureDNSService,
		ManagedClusterAddOnFeatureResourceMonitorService,
	}
}

type ManagedClusterVersionEnvironment string

const (
	// ManagedClusterVersionEnvironmentWindows - Windows.
	ManagedClusterVersionEnvironmentWindows ManagedClusterVersionEnvironment = "Windows"
)

// PossibleManagedClusterVersionEnvironmentValues returns the possible values for the ManagedClusterVersionEnvironment const type.
func PossibleManagedClusterVersionEnvironmentValues() []ManagedClusterVersionEnvironment {
	return []ManagedClusterVersionEnvironment{
		ManagedClusterVersionEnvironmentWindows,
	}
}

// ManagedIdentityType - The type of managed identity for the resource.
type ManagedIdentityType string

const (
	// ManagedIdentityTypeNone - Indicates that no identity is associated with the resource.
	ManagedIdentityTypeNone ManagedIdentityType = "None"
	// ManagedIdentityTypeSystemAssigned - Indicates that system assigned identity is associated with the resource.
	ManagedIdentityTypeSystemAssigned ManagedIdentityType = "SystemAssigned"
	// ManagedIdentityTypeSystemAssignedUserAssigned - Indicates that both system assigned and user assigned identity are associated
	// with the resource.
	ManagedIdentityTypeSystemAssignedUserAssigned ManagedIdentityType = "SystemAssigned, UserAssigned"
	// ManagedIdentityTypeUserAssigned - Indicates that user assigned identity is associated with the resource.
	ManagedIdentityTypeUserAssigned ManagedIdentityType = "UserAssigned"
)

// PossibleManagedIdentityTypeValues returns the possible values for the ManagedIdentityType const type.
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return []ManagedIdentityType{
		ManagedIdentityTypeNone,
		ManagedIdentityTypeSystemAssigned,
		ManagedIdentityTypeSystemAssignedUserAssigned,
		ManagedIdentityTypeUserAssigned,
	}
}

// ManagedResourceProvisioningState - The provisioning state of the managed resource.
type ManagedResourceProvisioningState string

const (
	ManagedResourceProvisioningStateCanceled  ManagedResourceProvisioningState = "Canceled"
	ManagedResourceProvisioningStateCreated   ManagedResourceProvisioningState = "Created"
	ManagedResourceProvisioningStateCreating  ManagedResourceProvisioningState = "Creating"
	ManagedResourceProvisioningStateDeleted   ManagedResourceProvisioningState = "Deleted"
	ManagedResourceProvisioningStateDeleting  ManagedResourceProvisioningState = "Deleting"
	ManagedResourceProvisioningStateFailed    ManagedResourceProvisioningState = "Failed"
	ManagedResourceProvisioningStateNone      ManagedResourceProvisioningState = "None"
	ManagedResourceProvisioningStateOther     ManagedResourceProvisioningState = "Other"
	ManagedResourceProvisioningStateSucceeded ManagedResourceProvisioningState = "Succeeded"
	ManagedResourceProvisioningStateUpdating  ManagedResourceProvisioningState = "Updating"
)

// PossibleManagedResourceProvisioningStateValues returns the possible values for the ManagedResourceProvisioningState const type.
func PossibleManagedResourceProvisioningStateValues() []ManagedResourceProvisioningState {
	return []ManagedResourceProvisioningState{
		ManagedResourceProvisioningStateCanceled,
		ManagedResourceProvisioningStateCreated,
		ManagedResourceProvisioningStateCreating,
		ManagedResourceProvisioningStateDeleted,
		ManagedResourceProvisioningStateDeleting,
		ManagedResourceProvisioningStateFailed,
		ManagedResourceProvisioningStateNone,
		ManagedResourceProvisioningStateOther,
		ManagedResourceProvisioningStateSucceeded,
		ManagedResourceProvisioningStateUpdating,
	}
}

// MoveCost - Specifies the move cost for the service.
type MoveCost string

const (
	// MoveCostHigh - Specifies the move cost of the service as High. The value is 3.
	MoveCostHigh MoveCost = "High"
	// MoveCostLow - Specifies the move cost of the service as Low. The value is 1.
	MoveCostLow MoveCost = "Low"
	// MoveCostMedium - Specifies the move cost of the service as Medium. The value is 2.
	MoveCostMedium MoveCost = "Medium"
	// MoveCostZero - Zero move cost. This value is zero.
	MoveCostZero MoveCost = "Zero"
)

// PossibleMoveCostValues returns the possible values for the MoveCost const type.
func PossibleMoveCostValues() []MoveCost {
	return []MoveCost{
		MoveCostHigh,
		MoveCostLow,
		MoveCostMedium,
		MoveCostZero,
	}
}

// NodeTypeSKUScaleType - Node type capacity scale type.
type NodeTypeSKUScaleType string

const (
	// NodeTypeSKUScaleTypeAutomatic - Automatic scale is allowed.
	NodeTypeSKUScaleTypeAutomatic NodeTypeSKUScaleType = "Automatic"
	// NodeTypeSKUScaleTypeManual - The user must manually scale out/in.
	NodeTypeSKUScaleTypeManual NodeTypeSKUScaleType = "Manual"
	// NodeTypeSKUScaleTypeNone - Node count is not adjustable in any way (e.g. it is fixed).
	NodeTypeSKUScaleTypeNone NodeTypeSKUScaleType = "None"
)

// PossibleNodeTypeSKUScaleTypeValues returns the possible values for the NodeTypeSKUScaleType const type.
func PossibleNodeTypeSKUScaleTypeValues() []NodeTypeSKUScaleType {
	return []NodeTypeSKUScaleType{
		NodeTypeSKUScaleTypeAutomatic,
		NodeTypeSKUScaleTypeManual,
		NodeTypeSKUScaleTypeNone,
	}
}

// NsgProtocol - Network protocol this rule applies to.
type NsgProtocol string

const (
	NsgProtocolAh    NsgProtocol = "ah"
	NsgProtocolEsp   NsgProtocol = "esp"
	NsgProtocolHTTP  NsgProtocol = "http"
	NsgProtocolHTTPS NsgProtocol = "https"
	NsgProtocolIcmp  NsgProtocol = "icmp"
	NsgProtocolTCP   NsgProtocol = "tcp"
	NsgProtocolUDP   NsgProtocol = "udp"
)

// PossibleNsgProtocolValues returns the possible values for the NsgProtocol const type.
func PossibleNsgProtocolValues() []NsgProtocol {
	return []NsgProtocol{
		NsgProtocolAh,
		NsgProtocolEsp,
		NsgProtocolHTTP,
		NsgProtocolHTTPS,
		NsgProtocolIcmp,
		NsgProtocolTCP,
		NsgProtocolUDP,
	}
}

// OsType - Cluster operating system, the default will be Windows
type OsType string

const (
	// OsTypeWindows - Indicates os is Windows.
	OsTypeWindows OsType = "Windows"
)

// PossibleOsTypeValues returns the possible values for the OsType const type.
func PossibleOsTypeValues() []OsType {
	return []OsType{
		OsTypeWindows,
	}
}

// PartitionScheme - Enumerates the ways that a service can be partitioned.
type PartitionScheme string

const (
	// PartitionSchemeNamed - Indicates that the partition is based on string names, and is a NamedPartitionScheme object. The
	// value is 2.
	PartitionSchemeNamed PartitionScheme = "Named"
	// PartitionSchemeSingleton - Indicates that the partition is based on string names, and is a SingletonPartitionScheme object,
	// The value is 0.
	PartitionSchemeSingleton PartitionScheme = "Singleton"
	// PartitionSchemeUniformInt64Range - Indicates that the partition is based on Int64 key ranges, and is a UniformInt64RangePartitionScheme
	// object. The value is 1.
	PartitionSchemeUniformInt64Range PartitionScheme = "UniformInt64Range"
)

// PossiblePartitionSchemeValues returns the possible values for the PartitionScheme const type.
func PossiblePartitionSchemeValues() []PartitionScheme {
	return []PartitionScheme{
		PartitionSchemeNamed,
		PartitionSchemeSingleton,
		PartitionSchemeUniformInt64Range,
	}
}

// PrivateEndpointNetworkPolicies - Enable or Disable apply network policies on private end point in the subnet.
type PrivateEndpointNetworkPolicies string

const (
	PrivateEndpointNetworkPoliciesDisabled PrivateEndpointNetworkPolicies = "disabled"
	PrivateEndpointNetworkPoliciesEnabled  PrivateEndpointNetworkPolicies = "enabled"
)

// PossiblePrivateEndpointNetworkPoliciesValues returns the possible values for the PrivateEndpointNetworkPolicies const type.
func PossiblePrivateEndpointNetworkPoliciesValues() []PrivateEndpointNetworkPolicies {
	return []PrivateEndpointNetworkPolicies{
		PrivateEndpointNetworkPoliciesDisabled,
		PrivateEndpointNetworkPoliciesEnabled,
	}
}

// PrivateIPAddressVersion - Specifies whether the IP configuration's private IP is IPv4 or IPv6. Default is IPv4.
type PrivateIPAddressVersion string

const (
	PrivateIPAddressVersionIPv4 PrivateIPAddressVersion = "IPv4"
	PrivateIPAddressVersionIPv6 PrivateIPAddressVersion = "IPv6"
)

// PossiblePrivateIPAddressVersionValues returns the possible values for the PrivateIPAddressVersion const type.
func PossiblePrivateIPAddressVersionValues() []PrivateIPAddressVersion {
	return []PrivateIPAddressVersion{
		PrivateIPAddressVersionIPv4,
		PrivateIPAddressVersionIPv6,
	}
}

// PrivateLinkServiceNetworkPolicies - Enable or Disable apply network policies on private link service in the subnet.
type PrivateLinkServiceNetworkPolicies string

const (
	PrivateLinkServiceNetworkPoliciesDisabled PrivateLinkServiceNetworkPolicies = "disabled"
	PrivateLinkServiceNetworkPoliciesEnabled  PrivateLinkServiceNetworkPolicies = "enabled"
)

// PossiblePrivateLinkServiceNetworkPoliciesValues returns the possible values for the PrivateLinkServiceNetworkPolicies const type.
func PossiblePrivateLinkServiceNetworkPoliciesValues() []PrivateLinkServiceNetworkPolicies {
	return []PrivateLinkServiceNetworkPolicies{
		PrivateLinkServiceNetworkPoliciesDisabled,
		PrivateLinkServiceNetworkPoliciesEnabled,
	}
}

// ProbeProtocol - the reference to the load balancer probe used by the load balancing rule.
type ProbeProtocol string

const (
	ProbeProtocolHTTP  ProbeProtocol = "http"
	ProbeProtocolHTTPS ProbeProtocol = "https"
	ProbeProtocolTCP   ProbeProtocol = "tcp"
)

// PossibleProbeProtocolValues returns the possible values for the ProbeProtocol const type.
func PossibleProbeProtocolValues() []ProbeProtocol {
	return []ProbeProtocol{
		ProbeProtocolHTTP,
		ProbeProtocolHTTPS,
		ProbeProtocolTCP,
	}
}

// Protocol - The reference to the transport protocol used by the load balancing rule.
type Protocol string

const (
	ProtocolTCP Protocol = "tcp"
	ProtocolUDP Protocol = "udp"
)

// PossibleProtocolValues returns the possible values for the Protocol const type.
func PossibleProtocolValues() []Protocol {
	return []Protocol{
		ProtocolTCP,
		ProtocolUDP,
	}
}

// PublicIPAddressVersion - Specifies whether the IP configuration's public IP is IPv4 or IPv6. Default is IPv4.
type PublicIPAddressVersion string

const (
	PublicIPAddressVersionIPv4 PublicIPAddressVersion = "IPv4"
	PublicIPAddressVersionIPv6 PublicIPAddressVersion = "IPv6"
)

// PossiblePublicIPAddressVersionValues returns the possible values for the PublicIPAddressVersion const type.
func PossiblePublicIPAddressVersionValues() []PublicIPAddressVersion {
	return []PublicIPAddressVersion{
		PublicIPAddressVersionIPv4,
		PublicIPAddressVersionIPv6,
	}
}

// RollingUpgradeMode - The mode used to monitor health during a rolling upgrade. The values are Monitored, and UnmonitoredAuto.
type RollingUpgradeMode string

const (
	// RollingUpgradeModeMonitored - The upgrade will stop after completing each upgrade domain and automatically monitor health
	// before proceeding. The value is 0.
	RollingUpgradeModeMonitored RollingUpgradeMode = "Monitored"
	// RollingUpgradeModeUnmonitoredAuto - The upgrade will proceed automatically without performing any health monitoring. The
	// value is 1.
	RollingUpgradeModeUnmonitoredAuto RollingUpgradeMode = "UnmonitoredAuto"
)

// PossibleRollingUpgradeModeValues returns the possible values for the RollingUpgradeMode const type.
func PossibleRollingUpgradeModeValues() []RollingUpgradeMode {
	return []RollingUpgradeMode{
		RollingUpgradeModeMonitored,
		RollingUpgradeModeUnmonitoredAuto,
	}
}

// SKUName - Sku Name.
type SKUName string

const (
	// SKUNameBasic - Basic requires a minimum of 3 nodes and allows only 1 node type.
	SKUNameBasic SKUName = "Basic"
	// SKUNameStandard - Requires a minimum of 5 nodes and allows 1 or more node type.
	SKUNameStandard SKUName = "Standard"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameBasic,
		SKUNameStandard,
	}
}

// SecurityType - Specifies the security type of the nodeType. Only Standard and TrustedLaunch are currently supported
type SecurityType string

const (
	// SecurityTypeStandard - Standard is the default security type for all machines.
	SecurityTypeStandard SecurityType = "Standard"
	// SecurityTypeTrustedLaunch - Trusted Launch is a security type that secures generation 2 virtual machines.
	SecurityTypeTrustedLaunch SecurityType = "TrustedLaunch"
)

// PossibleSecurityTypeValues returns the possible values for the SecurityType const type.
func PossibleSecurityTypeValues() []SecurityType {
	return []SecurityType{
		SecurityTypeStandard,
		SecurityTypeTrustedLaunch,
	}
}

// ServiceCorrelationScheme - The service correlation scheme.
type ServiceCorrelationScheme string

const (
	// ServiceCorrelationSchemeAlignedAffinity - Aligned affinity ensures that the primaries of the partitions of the affinitized
	// services are collocated on the same nodes. This is the default and is the same as selecting the Affinity scheme. The value
	// is 0.
	ServiceCorrelationSchemeAlignedAffinity ServiceCorrelationScheme = "AlignedAffinity"
	// ServiceCorrelationSchemeNonAlignedAffinity - Non-Aligned affinity guarantees that all replicas of each service will be
	// placed on the same nodes. Unlike Aligned Affinity, this does not guarantee that replicas of particular role will be collocated.
	// The value is 1.
	ServiceCorrelationSchemeNonAlignedAffinity ServiceCorrelationScheme = "NonAlignedAffinity"
)

// PossibleServiceCorrelationSchemeValues returns the possible values for the ServiceCorrelationScheme const type.
func PossibleServiceCorrelationSchemeValues() []ServiceCorrelationScheme {
	return []ServiceCorrelationScheme{
		ServiceCorrelationSchemeAlignedAffinity,
		ServiceCorrelationSchemeNonAlignedAffinity,
	}
}

// ServiceKind - The kind of service (Stateless or Stateful).
type ServiceKind string

const (
	// ServiceKindStateful - Uses Service Fabric to make its state or part of its state highly available and reliable. The value
	// is 1.
	ServiceKindStateful ServiceKind = "Stateful"
	// ServiceKindStateless - Does not use Service Fabric to make its state highly available or reliable. The value is 0.
	ServiceKindStateless ServiceKind = "Stateless"
)

// PossibleServiceKindValues returns the possible values for the ServiceKind const type.
func PossibleServiceKindValues() []ServiceKind {
	return []ServiceKind{
		ServiceKindStateful,
		ServiceKindStateless,
	}
}

// ServiceLoadMetricWeight - Determines the metric weight relative to the other metrics that are configured for this service.
// During runtime, if two metrics end up in conflict, the Cluster Resource Manager prefers the metric with
// the higher weight.
type ServiceLoadMetricWeight string

const (
	// ServiceLoadMetricWeightHigh - Specifies the metric weight of the service load as High. The value is 3.
	ServiceLoadMetricWeightHigh ServiceLoadMetricWeight = "High"
	// ServiceLoadMetricWeightLow - Specifies the metric weight of the service load as Low. The value is 1.
	ServiceLoadMetricWeightLow ServiceLoadMetricWeight = "Low"
	// ServiceLoadMetricWeightMedium - Specifies the metric weight of the service load as Medium. The value is 2.
	ServiceLoadMetricWeightMedium ServiceLoadMetricWeight = "Medium"
	// ServiceLoadMetricWeightZero - Disables resource balancing for this metric. This value is zero.
	ServiceLoadMetricWeightZero ServiceLoadMetricWeight = "Zero"
)

// PossibleServiceLoadMetricWeightValues returns the possible values for the ServiceLoadMetricWeight const type.
func PossibleServiceLoadMetricWeightValues() []ServiceLoadMetricWeight {
	return []ServiceLoadMetricWeight{
		ServiceLoadMetricWeightHigh,
		ServiceLoadMetricWeightLow,
		ServiceLoadMetricWeightMedium,
		ServiceLoadMetricWeightZero,
	}
}

// ServicePackageActivationMode - The activation Mode of the service package
type ServicePackageActivationMode string

const (
	// ServicePackageActivationModeExclusiveProcess - Indicates the application package activation mode will use exclusive process.
	ServicePackageActivationModeExclusiveProcess ServicePackageActivationMode = "ExclusiveProcess"
	// ServicePackageActivationModeSharedProcess - Indicates the application package activation mode will use shared process.
	ServicePackageActivationModeSharedProcess ServicePackageActivationMode = "SharedProcess"
)

// PossibleServicePackageActivationModeValues returns the possible values for the ServicePackageActivationMode const type.
func PossibleServicePackageActivationModeValues() []ServicePackageActivationMode {
	return []ServicePackageActivationMode{
		ServicePackageActivationModeExclusiveProcess,
		ServicePackageActivationModeSharedProcess,
	}
}

// ServicePlacementPolicyType - The type of placement policy for a service fabric service. Following are the possible values.
type ServicePlacementPolicyType string

const (
	// ServicePlacementPolicyTypeInvalidDomain - Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementInvalidDomainPolicyDescription,
	// which indicates that a particular fault or upgrade domain cannot be used for placement of this service. The value is 0.
	ServicePlacementPolicyTypeInvalidDomain ServicePlacementPolicyType = "InvalidDomain"
	// ServicePlacementPolicyTypeNonPartiallyPlaceService - Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementNonPartiallyPlaceServicePolicyDescription,
	// which indicates that if possible all replicas of a particular partition of the service should be placed atomically. The
	// value is 4.
	ServicePlacementPolicyTypeNonPartiallyPlaceService ServicePlacementPolicyType = "NonPartiallyPlaceService"
	// ServicePlacementPolicyTypePreferredPrimaryDomain - Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementPreferPrimaryDomainPolicyDescription,
	// which indicates that if possible the Primary replica for the partitions of the service should be located in a particular
	// domain as an optimization. The value is 2.
	ServicePlacementPolicyTypePreferredPrimaryDomain ServicePlacementPolicyType = "PreferredPrimaryDomain"
	// ServicePlacementPolicyTypeRequiredDomain - Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementRequireDomainDistributionPolicyDescription
	// indicating that the replicas of the service must be placed in a specific domain. The value is 1.
	ServicePlacementPolicyTypeRequiredDomain ServicePlacementPolicyType = "RequiredDomain"
	// ServicePlacementPolicyTypeRequiredDomainDistribution - Indicates that the ServicePlacementPolicyDescription is of type
	// ServicePlacementRequireDomainDistributionPolicyDescription, indicating that the system will disallow placement of any two
	// replicas from the same partition in the same domain at any time. The value is 3.
	ServicePlacementPolicyTypeRequiredDomainDistribution ServicePlacementPolicyType = "RequiredDomainDistribution"
)

// PossibleServicePlacementPolicyTypeValues returns the possible values for the ServicePlacementPolicyType const type.
func PossibleServicePlacementPolicyTypeValues() []ServicePlacementPolicyType {
	return []ServicePlacementPolicyType{
		ServicePlacementPolicyTypeInvalidDomain,
		ServicePlacementPolicyTypeNonPartiallyPlaceService,
		ServicePlacementPolicyTypePreferredPrimaryDomain,
		ServicePlacementPolicyTypeRequiredDomain,
		ServicePlacementPolicyTypeRequiredDomainDistribution,
	}
}

// ServiceScalingMechanismKind - Enumerates the ways that a service can be partitioned.
type ServiceScalingMechanismKind string

const (
	// ServiceScalingMechanismKindAddRemoveIncrementalNamedPartition - Represents a scaling mechanism for adding or removing named
	// partitions of a stateless service. The value is 1.
	ServiceScalingMechanismKindAddRemoveIncrementalNamedPartition ServiceScalingMechanismKind = "AddRemoveIncrementalNamedPartition"
	// ServiceScalingMechanismKindScalePartitionInstanceCount - Represents a scaling mechanism for adding or removing instances
	// of stateless service partition. The value is 0.
	ServiceScalingMechanismKindScalePartitionInstanceCount ServiceScalingMechanismKind = "ScalePartitionInstanceCount"
)

// PossibleServiceScalingMechanismKindValues returns the possible values for the ServiceScalingMechanismKind const type.
func PossibleServiceScalingMechanismKindValues() []ServiceScalingMechanismKind {
	return []ServiceScalingMechanismKind{
		ServiceScalingMechanismKindAddRemoveIncrementalNamedPartition,
		ServiceScalingMechanismKindScalePartitionInstanceCount,
	}
}

// ServiceScalingTriggerKind - Enumerates the ways that a service can be partitioned.
type ServiceScalingTriggerKind string

const (
	// ServiceScalingTriggerKindAveragePartitionLoadTrigger - Represents a scaling trigger related to an average load of a metric/resource
	// of a partition. The value is 0.
	ServiceScalingTriggerKindAveragePartitionLoadTrigger ServiceScalingTriggerKind = "AveragePartitionLoadTrigger"
	// ServiceScalingTriggerKindAverageServiceLoadTrigger - Represents a scaling policy related to an average load of a metric/resource
	// of a service. The value is 1.
	ServiceScalingTriggerKindAverageServiceLoadTrigger ServiceScalingTriggerKind = "AverageServiceLoadTrigger"
)

// PossibleServiceScalingTriggerKindValues returns the possible values for the ServiceScalingTriggerKind const type.
func PossibleServiceScalingTriggerKindValues() []ServiceScalingTriggerKind {
	return []ServiceScalingTriggerKind{
		ServiceScalingTriggerKindAveragePartitionLoadTrigger,
		ServiceScalingTriggerKindAverageServiceLoadTrigger,
	}
}

// UpdateType - Specifies the way the operation will be performed.
type UpdateType string

const (
	// UpdateTypeByUpgradeDomain - The operation will proceed one upgrade domain at a time, checking the health in between each
	// to continue.
	UpdateTypeByUpgradeDomain UpdateType = "ByUpgradeDomain"
	// UpdateTypeDefault - The operation will proceed in all specified nodes at the same time.
	UpdateTypeDefault UpdateType = "Default"
)

// PossibleUpdateTypeValues returns the possible values for the UpdateType const type.
func PossibleUpdateTypeValues() []UpdateType {
	return []UpdateType{
		UpdateTypeByUpgradeDomain,
		UpdateTypeDefault,
	}
}

// VMSetupAction - action to be performed on the vms before bootstrapping the service fabric runtime.
type VMSetupAction string

const (
	// VMSetupActionEnableContainers - Enable windows containers feature.
	VMSetupActionEnableContainers VMSetupAction = "EnableContainers"
	// VMSetupActionEnableHyperV - Enables windows HyperV feature.
	VMSetupActionEnableHyperV VMSetupAction = "EnableHyperV"
)

// PossibleVMSetupActionValues returns the possible values for the VMSetupAction const type.
func PossibleVMSetupActionValues() []VMSetupAction {
	return []VMSetupAction{
		VMSetupActionEnableContainers,
		VMSetupActionEnableHyperV,
	}
}

// VmssExtensionSetupOrder - Vm extension setup order.
type VmssExtensionSetupOrder string

const (
	// VmssExtensionSetupOrderBeforeSFRuntime - Indicates that the vm extension should run before the service fabric runtime starts.
	VmssExtensionSetupOrderBeforeSFRuntime VmssExtensionSetupOrder = "BeforeSFRuntime"
)

// PossibleVmssExtensionSetupOrderValues returns the possible values for the VmssExtensionSetupOrder const type.
func PossibleVmssExtensionSetupOrderValues() []VmssExtensionSetupOrder {
	return []VmssExtensionSetupOrder{
		VmssExtensionSetupOrderBeforeSFRuntime,
	}
}

// ZonalUpdateMode - Indicates the update mode for Cross Az clusters.
type ZonalUpdateMode string

const (
	// ZonalUpdateModeFast - The cluster will use a maximum of 3 upgrade domains per zone instead of 5 for Cross Az Node types
	// for faster deployments.
	ZonalUpdateModeFast ZonalUpdateMode = "Fast"
	// ZonalUpdateModeStandard - The cluster will use 5 upgrade domains for Cross Az Node types.
	ZonalUpdateModeStandard ZonalUpdateMode = "Standard"
)

// PossibleZonalUpdateModeValues returns the possible values for the ZonalUpdateMode const type.
func PossibleZonalUpdateModeValues() []ZonalUpdateMode {
	return []ZonalUpdateMode{
		ZonalUpdateModeFast,
		ZonalUpdateModeStandard,
	}
}
