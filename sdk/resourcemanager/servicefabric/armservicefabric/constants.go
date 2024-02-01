//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabric

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
	moduleVersion = "v2.0.1"
)

// AddOnFeatures - Available cluster add-on features
type AddOnFeatures string

const (
	AddOnFeaturesBackupRestoreService   AddOnFeatures = "BackupRestoreService"
	AddOnFeaturesDNSService             AddOnFeatures = "DnsService"
	AddOnFeaturesRepairManager          AddOnFeatures = "RepairManager"
	AddOnFeaturesResourceMonitorService AddOnFeatures = "ResourceMonitorService"
)

// PossibleAddOnFeaturesValues returns the possible values for the AddOnFeatures const type.
func PossibleAddOnFeaturesValues() []AddOnFeatures {
	return []AddOnFeatures{
		AddOnFeaturesBackupRestoreService,
		AddOnFeaturesDNSService,
		AddOnFeaturesRepairManager,
		AddOnFeaturesResourceMonitorService,
	}
}

// ArmServicePackageActivationMode - The activation Mode of the service package
type ArmServicePackageActivationMode string

const (
	// ArmServicePackageActivationModeExclusiveProcess - Indicates the application package activation mode will use exclusive
	// process.
	ArmServicePackageActivationModeExclusiveProcess ArmServicePackageActivationMode = "ExclusiveProcess"
	// ArmServicePackageActivationModeSharedProcess - Indicates the application package activation mode will use shared process.
	ArmServicePackageActivationModeSharedProcess ArmServicePackageActivationMode = "SharedProcess"
)

// PossibleArmServicePackageActivationModeValues returns the possible values for the ArmServicePackageActivationMode const type.
func PossibleArmServicePackageActivationModeValues() []ArmServicePackageActivationMode {
	return []ArmServicePackageActivationMode{
		ArmServicePackageActivationModeExclusiveProcess,
		ArmServicePackageActivationModeSharedProcess,
	}
}

// ArmUpgradeFailureAction - The activation Mode of the service package
type ArmUpgradeFailureAction string

const (
	// ArmUpgradeFailureActionManual - Indicates that a manual repair will need to be performed by the administrator if the upgrade
	// fails. Service Fabric will not proceed to the next upgrade domain automatically.
	ArmUpgradeFailureActionManual ArmUpgradeFailureAction = "Manual"
	// ArmUpgradeFailureActionRollback - Indicates that a rollback of the upgrade will be performed by Service Fabric if the upgrade
	// fails.
	ArmUpgradeFailureActionRollback ArmUpgradeFailureAction = "Rollback"
)

// PossibleArmUpgradeFailureActionValues returns the possible values for the ArmUpgradeFailureAction const type.
func PossibleArmUpgradeFailureActionValues() []ArmUpgradeFailureAction {
	return []ArmUpgradeFailureAction{
		ArmUpgradeFailureActionManual,
		ArmUpgradeFailureActionRollback,
	}
}

// ClusterEnvironment - Cluster operating system, the default will be Windows
type ClusterEnvironment string

const (
	ClusterEnvironmentLinux   ClusterEnvironment = "Linux"
	ClusterEnvironmentWindows ClusterEnvironment = "Windows"
)

// PossibleClusterEnvironmentValues returns the possible values for the ClusterEnvironment const type.
func PossibleClusterEnvironmentValues() []ClusterEnvironment {
	return []ClusterEnvironment{
		ClusterEnvironmentLinux,
		ClusterEnvironmentWindows,
	}
}

// ClusterState - The current state of the cluster.
// * WaitingForNodes - Indicates that the cluster resource is created and the resource provider is waiting for Service Fabric
// VM extension to boot up and report to it.
// * Deploying - Indicates that the Service Fabric runtime is being installed on the VMs. Cluster resource will be in this
// state until the cluster boots up and system services are up.
// * BaselineUpgrade - Indicates that the cluster is upgrading to establishes the cluster version. This upgrade is automatically
// initiated when the cluster boots up for the first time.
// * UpdatingUserConfiguration - Indicates that the cluster is being upgraded with the user provided configuration.
// * UpdatingUserCertificate - Indicates that the cluster is being upgraded with the user provided certificate.
// * UpdatingInfrastructure - Indicates that the cluster is being upgraded with the latest Service Fabric runtime version.
// This happens only when the upgradeMode is set to 'Automatic'.
// * EnforcingClusterVersion - Indicates that cluster is on a different version than expected and the cluster is being upgraded
// to the expected version.
// * UpgradeServiceUnreachable - Indicates that the system service in the cluster is no longer polling the Resource Provider.
// Clusters in this state cannot be managed by the Resource Provider.
// * AutoScale - Indicates that the ReliabilityLevel of the cluster is being adjusted.
// * Ready - Indicates that the cluster is in a stable state.
type ClusterState string

const (
	ClusterStateAutoScale                 ClusterState = "AutoScale"
	ClusterStateBaselineUpgrade           ClusterState = "BaselineUpgrade"
	ClusterStateDeploying                 ClusterState = "Deploying"
	ClusterStateEnforcingClusterVersion   ClusterState = "EnforcingClusterVersion"
	ClusterStateReady                     ClusterState = "Ready"
	ClusterStateUpdatingInfrastructure    ClusterState = "UpdatingInfrastructure"
	ClusterStateUpdatingUserCertificate   ClusterState = "UpdatingUserCertificate"
	ClusterStateUpdatingUserConfiguration ClusterState = "UpdatingUserConfiguration"
	ClusterStateUpgradeServiceUnreachable ClusterState = "UpgradeServiceUnreachable"
	ClusterStateWaitingForNodes           ClusterState = "WaitingForNodes"
)

// PossibleClusterStateValues returns the possible values for the ClusterState const type.
func PossibleClusterStateValues() []ClusterState {
	return []ClusterState{
		ClusterStateAutoScale,
		ClusterStateBaselineUpgrade,
		ClusterStateDeploying,
		ClusterStateEnforcingClusterVersion,
		ClusterStateReady,
		ClusterStateUpdatingInfrastructure,
		ClusterStateUpdatingUserCertificate,
		ClusterStateUpdatingUserConfiguration,
		ClusterStateUpgradeServiceUnreachable,
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

type ClusterVersionsEnvironment string

const (
	ClusterVersionsEnvironmentLinux   ClusterVersionsEnvironment = "Linux"
	ClusterVersionsEnvironmentWindows ClusterVersionsEnvironment = "Windows"
)

// PossibleClusterVersionsEnvironmentValues returns the possible values for the ClusterVersionsEnvironment const type.
func PossibleClusterVersionsEnvironmentValues() []ClusterVersionsEnvironment {
	return []ClusterVersionsEnvironment{
		ClusterVersionsEnvironmentLinux,
		ClusterVersionsEnvironmentWindows,
	}
}

// DurabilityLevel - The durability level of the node type. Learn about DurabilityLevel [https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity].
// * Bronze - No privileges. This is the default.
// * Silver - The infrastructure jobs can be paused for a duration of 10 minutes per UD.
// * Gold - The infrastructure jobs can be paused for a duration of 2 hours per UD. Gold durability can be enabled only on
// full node VM skus like D15_V2, G5 etc.
type DurabilityLevel string

const (
	DurabilityLevelBronze DurabilityLevel = "Bronze"
	DurabilityLevelGold   DurabilityLevel = "Gold"
	DurabilityLevelSilver DurabilityLevel = "Silver"
)

// PossibleDurabilityLevelValues returns the possible values for the DurabilityLevel const type.
func PossibleDurabilityLevelValues() []DurabilityLevel {
	return []DurabilityLevel{
		DurabilityLevelBronze,
		DurabilityLevelGold,
		DurabilityLevelSilver,
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

// NotificationCategory - The category of notification.
type NotificationCategory string

const (
	// NotificationCategoryWaveProgress - Notification will be regarding wave progress.
	NotificationCategoryWaveProgress NotificationCategory = "WaveProgress"
)

// PossibleNotificationCategoryValues returns the possible values for the NotificationCategory const type.
func PossibleNotificationCategoryValues() []NotificationCategory {
	return []NotificationCategory{
		NotificationCategoryWaveProgress,
	}
}

// NotificationChannel - The notification channel indicates the type of receivers subscribed to the notification, either user
// or subscription.
type NotificationChannel string

const (
	// NotificationChannelEmailSubscription - For subscription receivers. In this case, the parameter receivers should be a list
	// of roles of the subscription for the cluster (eg. Owner, AccountAdmin, etc) that will receive the notifications.
	NotificationChannelEmailSubscription NotificationChannel = "EmailSubscription"
	// NotificationChannelEmailUser - For email user receivers. In this case, the parameter receivers should be a list of email
	// addresses that will receive the notifications.
	NotificationChannelEmailUser NotificationChannel = "EmailUser"
)

// PossibleNotificationChannelValues returns the possible values for the NotificationChannel const type.
func PossibleNotificationChannelValues() []NotificationChannel {
	return []NotificationChannel{
		NotificationChannelEmailSubscription,
		NotificationChannelEmailUser,
	}
}

// NotificationLevel - The level of notification.
type NotificationLevel string

const (
	// NotificationLevelAll - Receive all notifications.
	NotificationLevelAll NotificationLevel = "All"
	// NotificationLevelCritical - Receive only critical notifications.
	NotificationLevelCritical NotificationLevel = "Critical"
)

// PossibleNotificationLevelValues returns the possible values for the NotificationLevel const type.
func PossibleNotificationLevelValues() []NotificationLevel {
	return []NotificationLevel{
		NotificationLevelAll,
		NotificationLevelCritical,
	}
}

// PartitionScheme - Enumerates the ways that a service can be partitioned.
type PartitionScheme string

const (
	// PartitionSchemeInvalid - Indicates the partition kind is invalid. All Service Fabric enumerations have the invalid type.
	// The value is zero.
	PartitionSchemeInvalid PartitionScheme = "Invalid"
	// PartitionSchemeNamed - Indicates that the partition is based on string names, and is a NamedPartitionSchemeDescription
	// object. The value is 3
	PartitionSchemeNamed PartitionScheme = "Named"
	// PartitionSchemeSingleton - Indicates that the partition is based on string names, and is a SingletonPartitionSchemeDescription
	// object, The value is 1.
	PartitionSchemeSingleton PartitionScheme = "Singleton"
	// PartitionSchemeUniformInt64Range - Indicates that the partition is based on Int64 key ranges, and is a UniformInt64RangePartitionSchemeDescription
	// object. The value is 2.
	PartitionSchemeUniformInt64Range PartitionScheme = "UniformInt64Range"
)

// PossiblePartitionSchemeValues returns the possible values for the PartitionScheme const type.
func PossiblePartitionSchemeValues() []PartitionScheme {
	return []PartitionScheme{
		PartitionSchemeInvalid,
		PartitionSchemeNamed,
		PartitionSchemeSingleton,
		PartitionSchemeUniformInt64Range,
	}
}

// ProvisioningState - The provisioning state of the cluster resource.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ReliabilityLevel - The reliability level sets the replica set size of system services. Learn about ReliabilityLevel [https://docs.microsoft.com/azure/service-fabric/service-fabric-cluster-capacity].
// * None - Run the System services with a target replica set count of 1. This should only be used for test clusters.
// * Bronze - Run the System services with a target replica set count of 3. This should only be used for test clusters.
// * Silver - Run the System services with a target replica set count of 5.
// * Gold - Run the System services with a target replica set count of 7.
// * Platinum - Run the System services with a target replica set count of 9.
type ReliabilityLevel string

const (
	ReliabilityLevelBronze   ReliabilityLevel = "Bronze"
	ReliabilityLevelGold     ReliabilityLevel = "Gold"
	ReliabilityLevelNone     ReliabilityLevel = "None"
	ReliabilityLevelPlatinum ReliabilityLevel = "Platinum"
	ReliabilityLevelSilver   ReliabilityLevel = "Silver"
)

// PossibleReliabilityLevelValues returns the possible values for the ReliabilityLevel const type.
func PossibleReliabilityLevelValues() []ReliabilityLevel {
	return []ReliabilityLevel{
		ReliabilityLevelBronze,
		ReliabilityLevelGold,
		ReliabilityLevelNone,
		ReliabilityLevelPlatinum,
		ReliabilityLevelSilver,
	}
}

// RollingUpgradeMode - The mode used to monitor health during a rolling upgrade. The values are UnmonitoredAuto, UnmonitoredManual,
// and Monitored.
type RollingUpgradeMode string

const (
	// RollingUpgradeModeInvalid - Indicates the upgrade mode is invalid. All Service Fabric enumerations have the invalid type.
	// The value is zero.
	RollingUpgradeModeInvalid RollingUpgradeMode = "Invalid"
	// RollingUpgradeModeMonitored - The upgrade will stop after completing each upgrade domain and automatically monitor health
	// before proceeding. The value is 3
	RollingUpgradeModeMonitored RollingUpgradeMode = "Monitored"
	// RollingUpgradeModeUnmonitoredAuto - The upgrade will proceed automatically without performing any health monitoring. The
	// value is 1
	RollingUpgradeModeUnmonitoredAuto RollingUpgradeMode = "UnmonitoredAuto"
	// RollingUpgradeModeUnmonitoredManual - The upgrade will stop after completing each upgrade domain, giving the opportunity
	// to manually monitor health before proceeding. The value is 2
	RollingUpgradeModeUnmonitoredManual RollingUpgradeMode = "UnmonitoredManual"
)

// PossibleRollingUpgradeModeValues returns the possible values for the RollingUpgradeMode const type.
func PossibleRollingUpgradeModeValues() []RollingUpgradeMode {
	return []RollingUpgradeMode{
		RollingUpgradeModeInvalid,
		RollingUpgradeModeMonitored,
		RollingUpgradeModeUnmonitoredAuto,
		RollingUpgradeModeUnmonitoredManual,
	}
}

// ServiceCorrelationScheme - The service correlation scheme.
type ServiceCorrelationScheme string

const (
	// ServiceCorrelationSchemeAffinity - Indicates that this service has an affinity relationship with another service. Provided
	// for backwards compatibility, consider preferring the Aligned or NonAlignedAffinity options. The value is 1.
	ServiceCorrelationSchemeAffinity ServiceCorrelationScheme = "Affinity"
	// ServiceCorrelationSchemeAlignedAffinity - Aligned affinity ensures that the primaries of the partitions of the affinitized
	// services are collocated on the same nodes. This is the default and is the same as selecting the Affinity scheme. The value
	// is 2.
	ServiceCorrelationSchemeAlignedAffinity ServiceCorrelationScheme = "AlignedAffinity"
	// ServiceCorrelationSchemeInvalid - An invalid correlation scheme. Cannot be used. The value is zero.
	ServiceCorrelationSchemeInvalid ServiceCorrelationScheme = "Invalid"
	// ServiceCorrelationSchemeNonAlignedAffinity - Non-Aligned affinity guarantees that all replicas of each service will be
	// placed on the same nodes. Unlike Aligned Affinity, this does not guarantee that replicas of particular role will be collocated.
	// The value is 3.
	ServiceCorrelationSchemeNonAlignedAffinity ServiceCorrelationScheme = "NonAlignedAffinity"
)

// PossibleServiceCorrelationSchemeValues returns the possible values for the ServiceCorrelationScheme const type.
func PossibleServiceCorrelationSchemeValues() []ServiceCorrelationScheme {
	return []ServiceCorrelationScheme{
		ServiceCorrelationSchemeAffinity,
		ServiceCorrelationSchemeAlignedAffinity,
		ServiceCorrelationSchemeInvalid,
		ServiceCorrelationSchemeNonAlignedAffinity,
	}
}

// ServiceKind - The kind of service (Stateless or Stateful).
type ServiceKind string

const (
	// ServiceKindInvalid - Indicates the service kind is invalid. All Service Fabric enumerations have the invalid type. The
	// value is zero.
	ServiceKindInvalid ServiceKind = "Invalid"
	// ServiceKindStateful - Uses Service Fabric to make its state or part of its state highly available and reliable. The value
	// is 2.
	ServiceKindStateful ServiceKind = "Stateful"
	// ServiceKindStateless - Does not use Service Fabric to make its state highly available or reliable. The value is 1.
	ServiceKindStateless ServiceKind = "Stateless"
)

// PossibleServiceKindValues returns the possible values for the ServiceKind const type.
func PossibleServiceKindValues() []ServiceKind {
	return []ServiceKind{
		ServiceKindInvalid,
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

// ServicePlacementPolicyType - The type of placement policy for a service fabric service. Following are the possible values.
type ServicePlacementPolicyType string

const (
	// ServicePlacementPolicyTypeInvalid - Indicates the type of the placement policy is invalid. All Service Fabric enumerations
	// have the invalid type. The value is zero.
	ServicePlacementPolicyTypeInvalid ServicePlacementPolicyType = "Invalid"
	// ServicePlacementPolicyTypeInvalidDomain - Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementInvalidDomainPolicyDescription,
	// which indicates that a particular fault or upgrade domain cannot be used for placement of this service. The value is 1.
	ServicePlacementPolicyTypeInvalidDomain ServicePlacementPolicyType = "InvalidDomain"
	// ServicePlacementPolicyTypeNonPartiallyPlaceService - Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementNonPartiallyPlaceServicePolicyDescription,
	// which indicates that if possible all replicas of a particular partition of the service should be placed atomically. The
	// value is 5.
	ServicePlacementPolicyTypeNonPartiallyPlaceService ServicePlacementPolicyType = "NonPartiallyPlaceService"
	// ServicePlacementPolicyTypePreferredPrimaryDomain - Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementPreferPrimaryDomainPolicyDescription,
	// which indicates that if possible the Primary replica for the partitions of the service should be located in a particular
	// domain as an optimization. The value is 3.
	ServicePlacementPolicyTypePreferredPrimaryDomain ServicePlacementPolicyType = "PreferredPrimaryDomain"
	// ServicePlacementPolicyTypeRequiredDomain - Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementRequireDomainDistributionPolicyDescription
	// indicating that the replicas of the service must be placed in a specific domain. The value is 2.
	ServicePlacementPolicyTypeRequiredDomain ServicePlacementPolicyType = "RequiredDomain"
	// ServicePlacementPolicyTypeRequiredDomainDistribution - Indicates that the ServicePlacementPolicyDescription is of type
	// ServicePlacementRequireDomainDistributionPolicyDescription, indicating that the system will disallow placement of any two
	// replicas from the same partition in the same domain at any time. The value is 4.
	ServicePlacementPolicyTypeRequiredDomainDistribution ServicePlacementPolicyType = "RequiredDomainDistribution"
)

// PossibleServicePlacementPolicyTypeValues returns the possible values for the ServicePlacementPolicyType const type.
func PossibleServicePlacementPolicyTypeValues() []ServicePlacementPolicyType {
	return []ServicePlacementPolicyType{
		ServicePlacementPolicyTypeInvalid,
		ServicePlacementPolicyTypeInvalidDomain,
		ServicePlacementPolicyTypeNonPartiallyPlaceService,
		ServicePlacementPolicyTypePreferredPrimaryDomain,
		ServicePlacementPolicyTypeRequiredDomain,
		ServicePlacementPolicyTypeRequiredDomainDistribution,
	}
}

// SfZonalUpgradeMode - This property controls the logical grouping of VMs in upgrade domains (UDs). This property can't be
// modified if a node type with multiple Availability Zones is already present in the cluster.
type SfZonalUpgradeMode string

const (
	// SfZonalUpgradeModeHierarchical - If this value is omitted or set to Hierarchical, VMs are grouped to reflect the zonal
	// distribution in up to 15 UDs. Each of the three zones has five UDs. This ensures that the zones are updated one at a time,
	// moving to next zone only after completing five UDs within the first zone. This update process is safer for the cluster
	// and the user application.
	SfZonalUpgradeModeHierarchical SfZonalUpgradeMode = "Hierarchical"
	// SfZonalUpgradeModeParallel - VMs under the node type are grouped into UDs and ignore the zone info in five UDs. This setting
	// causes UDs across all zones to be upgraded at the same time. This deployment mode is faster for upgrades, we don't recommend
	// it because it goes against the SDP guidelines, which state that the updates should be applied to one zone at a time.
	SfZonalUpgradeModeParallel SfZonalUpgradeMode = "Parallel"
)

// PossibleSfZonalUpgradeModeValues returns the possible values for the SfZonalUpgradeMode const type.
func PossibleSfZonalUpgradeModeValues() []SfZonalUpgradeMode {
	return []SfZonalUpgradeMode{
		SfZonalUpgradeModeHierarchical,
		SfZonalUpgradeModeParallel,
	}
}

// StoreName - The local certificate store location.
type StoreName string

const (
	StoreNameAddressBook          StoreName = "AddressBook"
	StoreNameAuthRoot             StoreName = "AuthRoot"
	StoreNameCertificateAuthority StoreName = "CertificateAuthority"
	StoreNameDisallowed           StoreName = "Disallowed"
	StoreNameMy                   StoreName = "My"
	StoreNameRoot                 StoreName = "Root"
	StoreNameTrustedPeople        StoreName = "TrustedPeople"
	StoreNameTrustedPublisher     StoreName = "TrustedPublisher"
)

// PossibleStoreNameValues returns the possible values for the StoreName const type.
func PossibleStoreNameValues() []StoreName {
	return []StoreName{
		StoreNameAddressBook,
		StoreNameAuthRoot,
		StoreNameCertificateAuthority,
		StoreNameDisallowed,
		StoreNameMy,
		StoreNameRoot,
		StoreNameTrustedPeople,
		StoreNameTrustedPublisher,
	}
}

// UpgradeMode - The upgrade mode of the cluster when new Service Fabric runtime version is available.
type UpgradeMode string

const (
	// UpgradeModeAutomatic - The cluster will be automatically upgraded to the latest Service Fabric runtime version, **upgradeWave**
	// will determine when the upgrade starts after the new version becomes available.
	UpgradeModeAutomatic UpgradeMode = "Automatic"
	// UpgradeModeManual - The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster
	// is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
	UpgradeModeManual UpgradeMode = "Manual"
)

// PossibleUpgradeModeValues returns the possible values for the UpgradeMode const type.
func PossibleUpgradeModeValues() []UpgradeMode {
	return []UpgradeMode{
		UpgradeModeAutomatic,
		UpgradeModeManual,
	}
}

// VmssZonalUpgradeMode - This property defines the upgrade mode for the virtual machine scale set, it is mandatory if a node
// type with multiple Availability Zones is added.
type VmssZonalUpgradeMode string

const (
	// VmssZonalUpgradeModeHierarchical - VMs are grouped to reflect the zonal distribution in up to 15 UDs. Each of the three
	// zones has five UDs. This ensures that the zones are updated one at a time, moving to next zone only after completing five
	// UDs within the first zone.
	VmssZonalUpgradeModeHierarchical VmssZonalUpgradeMode = "Hierarchical"
	// VmssZonalUpgradeModeParallel - Updates will happen in all Availability Zones at once for the virtual machine scale sets.
	VmssZonalUpgradeModeParallel VmssZonalUpgradeMode = "Parallel"
)

// PossibleVmssZonalUpgradeModeValues returns the possible values for the VmssZonalUpgradeMode const type.
func PossibleVmssZonalUpgradeModeValues() []VmssZonalUpgradeMode {
	return []VmssZonalUpgradeMode{
		VmssZonalUpgradeModeHierarchical,
		VmssZonalUpgradeModeParallel,
	}
}
