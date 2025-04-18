// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcontainerservice

import "time"

// AddonStatusProfile - The status profile of the addons and other kubernetes components
type AddonStatusProfile struct {
	// Observed error message from the addon or component
	ErrorMessage *string

	// Name of the addon or component
	Name *string

	// Observed phase of the addon or component on the provisioned cluster. Possible values include: 'pending', 'provisioning',
	// 'provisioning {HelmChartInstalled}', 'provisioning {MSICertificateDownloaded}',
	// 'provisioned', 'deleting', 'failed', 'upgrading'
	Phase *AddonPhase

	// Indicates whether the addon or component is ready
	Ready *bool
}

// AgentPool - The agentPool resource definition
type AgentPool struct {
	// Extended location pointing to the underlying infrastructure
	ExtendedLocation *ExtendedLocation

	// Properties of the agent pool resource
	Properties *AgentPoolProperties

	// Resource tags
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// AgentPoolListResult - List of all agent pool resources associated with the provisioned cluster.
type AgentPoolListResult struct {
	NextLink *string
	Value    []*AgentPool
}

// AgentPoolProperties - Properties of the agent pool resource
type AgentPoolProperties struct {
	// Number of nodes in the agent pool. The default value is 1.
	Count *int32

	// Whether to enable auto-scaler. Default value is false
	EnableAutoScaling *bool

	// The maximum number of nodes for auto-scaling
	MaxCount *int32

	// The maximum number of pods that can run on a node.
	MaxPods *int32

	// The minimum number of nodes for auto-scaling
	MinCount *int32

	// The node labels to be persisted across all nodes in agent pool.
	NodeLabels map[string]*string

	// Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints []*string

	// Specifies the OS SKU used by the agent pool. The default is CBLMariner if OSType is Linux. The default is Windows2019 when
	// OSType is Windows.
	OSSKU *OSSKU

	// The particular KubernetesVersion Image OS Type (Linux, Windows)
	OSType *OsType

	// The observed status of the agent pool.
	Status *AgentPoolProvisioningStatusStatus

	// The VM sku size of the agent pool node VMs.
	VMSize *string

	// READ-ONLY; Version of Kubernetes in use by the agent pool. This is inherited from the kubernetesVersion of the provisioned
	// cluster.
	KubernetesVersion *string

	// READ-ONLY; The status of the latest long running operation for the agent pool.
	ProvisioningState *ResourceProvisioningState
}

// AgentPoolProvisioningStatusStatus - The observed status of the agent pool.
type AgentPoolProvisioningStatusStatus struct {
	// Error messages during an agent pool operation or steady state.
	ErrorMessage  *string
	ReadyReplicas []*AgentPoolUpdateProfile

	// READ-ONLY; The current state of the agent pool.
	CurrentState *ResourceProvisioningState
}

// AgentPoolUpdateProfile - Profile for agent pool properties that can be updated
type AgentPoolUpdateProfile struct {
	// Number of nodes in the agent pool. The default value is 1.
	Count *int32

	// The VM sku size of the agent pool node VMs.
	VMSize *string

	// READ-ONLY; Version of Kubernetes in use by the agent pool. This is inherited from the kubernetesVersion of the provisioned
	// cluster.
	KubernetesVersion *string
}

// CloudProviderProfile - The profile for the underlying cloud infrastructure provider for the provisioned cluster.
type CloudProviderProfile struct {
	// The profile for the infrastructure networks used by the provisioned cluster
	InfraNetworkProfile *CloudProviderProfileInfraNetworkProfile
}

// CloudProviderProfileInfraNetworkProfile - The profile for the infrastructure networks used by the provisioned cluster
type CloudProviderProfileInfraNetworkProfile struct {
	// List of ARM resource Ids (maximum 1) for the infrastructure network object e.g.
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/logicalNetworks/{logicalNetworkName}
	VnetSubnetIDs []*string
}

// ClusterVMAccessProfile - The SSH restricted access profile for the VMs in the provisioned cluster.
type ClusterVMAccessProfile struct {
	// IP Address or CIDR for SSH access to VMs in the provisioned cluster
	AuthorizedIPRanges *string
}

// ControlPlaneProfile - The properties of the control plane nodes of the provisioned cluster
type ControlPlaneProfile struct {
	// IP Address of the Kubernetes API server
	ControlPlaneEndpoint *ControlPlaneProfileControlPlaneEndpoint

	// Number of control plane nodes. The default value is 1, and the count should be an odd number
	Count *int32

	// VM sku size of the control plane nodes
	VMSize *string
}

// ControlPlaneProfileControlPlaneEndpoint - IP Address of the Kubernetes API server
type ControlPlaneProfileControlPlaneEndpoint struct {
	// IP address of the Kubernetes API server
	HostIP *string
}

// CredentialResult - The credential result response.
type CredentialResult struct {
	// READ-ONLY; The name of the credential.
	Name *string

	// READ-ONLY; Base64-encoded Kubernetes configuration file.
	Value []byte
}

// ExtendedLocation - Extended location pointing to the underlying infrastructure
type ExtendedLocation struct {
	// ARM Id of the extended location.
	Name *string

	// The extended location type. Allowed value: 'CustomLocation'
	Type *ExtendedLocationTypes
}

// HybridIdentityMetadata - Defines the hybridIdentityMetadata.
type HybridIdentityMetadata struct {
	// REQUIRED; Resource properties.
	Properties *HybridIdentityMetadataProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// HybridIdentityMetadataList - List of hybridIdentityMetadata.
type HybridIdentityMetadataList struct {
	// REQUIRED; Array of hybridIdentityMetadata
	Value []*HybridIdentityMetadata

	// Url to follow for getting next page of hybridIdentityMetadata.
	NextLink *string
}

// HybridIdentityMetadataProperties - Defines the resource properties for the hybrid identity metadata.
type HybridIdentityMetadataProperties struct {
	// Onboarding public key for provisioning the Managed identity for the connected cluster.
	PublicKey *string

	// Unique id of the parent provisioned cluster resource.
	ResourceUID *string

	// READ-ONLY; Provisioning state of the resource
	ProvisioningState *ResourceProvisioningState
}

// KubernetesPatchVersions - Kubernetes Patch Version profile
type KubernetesPatchVersions struct {
	// Indicates whether the kubernetes version image is ready or not
	Readiness []*KubernetesVersionReadiness

	// Possible upgrade paths for given patch version
	Upgrades []*string
}

// KubernetesVersionProfile - The supported kubernetes versions.
type KubernetesVersionProfile struct {
	// Extended location pointing to the underlying infrastructure
	ExtendedLocation *ExtendedLocation

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY
	Properties *KubernetesVersionProfileProperties

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// KubernetesVersionProfileList - List of supported kubernetes versions.
type KubernetesVersionProfileList struct {
	NextLink *string
	Value    []*KubernetesVersionProfile
}

type KubernetesVersionProfileProperties struct {
	// List of supported Kubernetes versions
	Values []*KubernetesVersionProperties

	// READ-ONLY; Provisioning state of the resource
	ProvisioningState *ResourceProvisioningState
}

// KubernetesVersionProperties - Kubernetes version profile for given major.minor release
type KubernetesVersionProperties struct {
	// READ-ONLY; Whether this version is in preview mode.
	IsPreview *bool

	// READ-ONLY; Patch versions of a Kubernetes release
	PatchVersions map[string]*KubernetesPatchVersions

	// READ-ONLY; major.minor version of Kubernetes release
	Version *string
}

// KubernetesVersionReadiness - Indicates whether the kubernetes version image is ready or not
type KubernetesVersionReadiness struct {
	// Specifies the OS SKU used by the agent pool. The default is CBLMariner if OSType is Linux. The default is Windows2019 when
	// OSType is Windows.
	OSSKU *OSSKU

	// READ-ONLY; The error message for version not being ready
	ErrorMessage *string

	// READ-ONLY; The particular KubernetesVersion Image OS Type (Linux, Windows)
	OSType *OsType

	// READ-ONLY; Whether the kubernetes version image is ready or not
	Ready *bool
}

// LinuxProfileProperties - SSH profile for control plane and nodepool VMs of the provisioned cluster.
type LinuxProfileProperties struct {
	// SSH configuration for VMs of the provisioned cluster.
	SSH *LinuxProfilePropertiesSSH
}

// LinuxProfilePropertiesSSH - SSH configuration for VMs of the provisioned cluster.
type LinuxProfilePropertiesSSH struct {
	// The list of SSH public keys used to authenticate with VMs. A maximum of 1 key may be specified.
	PublicKeys []*LinuxProfilePropertiesSSHPublicKeysItem
}

type LinuxProfilePropertiesSSHPublicKeysItem struct {
	// Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without
	// headers.
	KeyData *string
}

// ListCredentialResponse - The list kubeconfig result response.
type ListCredentialResponse struct {
	Error      *ListCredentialResponseError
	Properties *ListCredentialResponseProperties

	// READ-ONLY; Operation Id
	ID *string

	// READ-ONLY; Operation Name
	Name *string

	// READ-ONLY; ARM Resource Id of the provisioned cluster instance
	ResourceID *string

	// READ-ONLY; Provisioning state of the resource
	Status *ResourceProvisioningState
}

type ListCredentialResponseError struct {
	Code    *string
	Message *string
}

type ListCredentialResponseProperties struct {
	// READ-ONLY; Base64-encoded Kubernetes configuration file.
	Kubeconfigs []*CredentialResult
}

// NamedAgentPoolProfile - Profile of the default agent pool along with a name parameter
type NamedAgentPoolProfile struct {
	// Number of nodes in the agent pool. The default value is 1.
	Count *int32

	// Whether to enable auto-scaler. Default value is false
	EnableAutoScaling *bool

	// The maximum number of nodes for auto-scaling
	MaxCount *int32

	// The maximum number of pods that can run on a node.
	MaxPods *int32

	// The minimum number of nodes for auto-scaling
	MinCount *int32

	// Unique name of the default agent pool in the context of the provisioned cluster. Default value is -nodepool1
	Name *string

	// The node labels to be persisted across all nodes in agent pool.
	NodeLabels map[string]*string

	// Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints []*string

	// Specifies the OS SKU used by the agent pool. The default is CBLMariner if OSType is Linux. The default is Windows2019 when
	// OSType is Windows.
	OSSKU *OSSKU

	// The particular KubernetesVersion Image OS Type (Linux, Windows)
	OSType *OsType

	// The VM sku size of the agent pool node VMs.
	VMSize *string

	// READ-ONLY; Version of Kubernetes in use by the agent pool. This is inherited from the kubernetesVersion of the provisioned
	// cluster.
	KubernetesVersion *string
}

// NetworkProfile - The network configuration profile for the provisioned cluster.
type NetworkProfile struct {
	// Profile of the HA Proxy load balancer.
	LoadBalancerProfile *NetworkProfileLoadBalancerProfile

	// Network policy used for building Kubernetes network. Possible values include: 'calico'.
	NetworkPolicy *NetworkPolicy

	// A CIDR notation IP Address range from which to assign pod IPs.
	PodCidr *string
}

// NetworkProfileLoadBalancerProfile - Profile of the HA Proxy load balancer.
type NetworkProfileLoadBalancerProfile struct {
	// Number of HA Proxy load balancer VMs. The default value is 0.
	Count *int32
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// ProvisionedCluster - The provisioned cluster resource definition.
type ProvisionedCluster struct {
	// Extended location pointing to the underlying infrastructure
	ExtendedLocation *ExtendedLocation

	// Properties of the provisioned cluster.
	Properties *ProvisionedClusterProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ProvisionedClusterLicenseProfile - The license profile of the provisioned cluster.
type ProvisionedClusterLicenseProfile struct {
	// Indicates whether Azure Hybrid Benefit is opted in. Default value is false
	AzureHybridBenefit *AzureHybridBenefit
}

// ProvisionedClusterListResult - Lists the ProvisionedClusterInstance resource associated with the ConnectedCluster.
type ProvisionedClusterListResult struct {
	NextLink *string
	Value    []*ProvisionedCluster
}

// ProvisionedClusterPoolUpgradeProfile - The list of available kubernetes versions for upgrade.
type ProvisionedClusterPoolUpgradeProfile struct {
	// List of available kubernetes versions for upgrade.
	Upgrades []*ProvisionedClusterPoolUpgradeProfileProperties

	// READ-ONLY; The Kubernetes version (major.minor.patch).
	KubernetesVersion *string

	// READ-ONLY; The particular KubernetesVersion Image OS Type (Linux, Windows)
	OSType *OsType
}

// ProvisionedClusterPoolUpgradeProfileProperties - The upgrade properties.
type ProvisionedClusterPoolUpgradeProfileProperties struct {
	// READ-ONLY; Whether the Kubernetes version is currently in preview.
	IsPreview *bool

	// READ-ONLY; The Kubernetes version (major.minor.patch).
	KubernetesVersion *string
}

// ProvisionedClusterProperties - Properties of the provisioned cluster.
type ProvisionedClusterProperties struct {
	// The agent pool properties for the provisioned cluster.
	AgentPoolProfiles []*NamedAgentPoolProfile

	// Parameters to be applied to the cluster-autoscaler when auto scaling is enabled for the provisioned cluster.
	AutoScalerProfile *ProvisionedClusterPropertiesAutoScalerProfile

	// The profile for the underlying cloud infrastructure provider for the provisioned cluster.
	CloudProviderProfile *CloudProviderProfile

	// The SSH restricted access profile for the VMs in the provisioned cluster.
	ClusterVMAccessProfile *ClusterVMAccessProfile

	// The profile for control plane of the provisioned cluster.
	ControlPlane *ControlPlaneProfile

	// The version of Kubernetes in use by the provisioned cluster.
	KubernetesVersion *string

	// The license profile of the provisioned cluster.
	LicenseProfile *ProvisionedClusterLicenseProfile

	// The profile for Linux VMs in the provisioned cluster.
	LinuxProfile *LinuxProfileProperties

	// The network configuration profile for the provisioned cluster.
	NetworkProfile *NetworkProfile

	// The storage configuration profile for the provisioned cluster.
	StorageProfile *StorageProfile

	// READ-ONLY; The status of the latest long running operation for the provisioned cluster.
	ProvisioningState *ResourceProvisioningState

	// READ-ONLY; The observed status of the provisioned cluster.
	Status *ProvisionedClusterPropertiesStatus
}

// ProvisionedClusterPropertiesAutoScalerProfile - Parameters to be applied to the cluster-autoscaler when auto scaling is
// enabled for the provisioned cluster.
type ProvisionedClusterPropertiesAutoScalerProfile struct {
	// Valid values are 'true' and 'false'
	BalanceSimilarNodeGroups *string

	// If not specified, the default is 'random'. See expanders [https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders]
	// for more information.
	Expander *Expander

	// The default is 10.
	MaxEmptyBulkDelete *string

	// The default is 600.
	MaxGracefulTerminationSec *string

	// The default is '15m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.
	MaxNodeProvisionTime *string

	// The default is 45. The maximum is 100 and the minimum is 0.
	MaxTotalUnreadyPercentage *string

	// For scenarios like burst/batch scale where you don't want CA to act before the kubernetes scheduler could schedule all
	// the pods, you can tell CA to ignore unscheduled pods before they're a certain
	// age. The default is '0s'. Values must be an integer followed by a unit ('s' for seconds, 'm' for minutes, 'h' for hours,
	// etc).
	NewPodScaleUpDelay *string

	// This must be an integer. The default is 3.
	OkTotalUnreadyCount *string

	// The default is '10m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.
	ScaleDownDelayAfterAdd *string

	// The default is the scan-interval. Values must be an integer followed by an 'm'. No unit of time other than minutes (m)
	// is supported.
	ScaleDownDelayAfterDelete *string

	// The default is '3m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.
	ScaleDownDelayAfterFailure *string

	// The default is '10m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.
	ScaleDownUnneededTime *string

	// The default is '20m'. Values must be an integer followed by an 'm'. No unit of time other than minutes (m) is supported.
	ScaleDownUnreadyTime *string

	// The default is '0.5'.
	ScaleDownUtilizationThreshold *string

	// The default is '10'. Values must be an integer number of seconds.
	ScanInterval *string

	// The default is true.
	SkipNodesWithLocalStorage *string

	// The default is true.
	SkipNodesWithSystemPods *string
}

// ProvisionedClusterPropertiesStatus - The observed status of the provisioned cluster.
type ProvisionedClusterPropertiesStatus struct {
	// The detailed status of the provisioned cluster components including addons.
	ControlPlaneStatus []*AddonStatusProfile

	// Error messages during a provisioned cluster operation or steady state.
	ErrorMessage *string

	// READ-ONLY; The current state of the provisioned cluster.
	CurrentState *ResourceProvisioningState
}

// ProvisionedClusterUpgradeProfile - The list of available kubernetes version upgrades for the provisioned cluster.
type ProvisionedClusterUpgradeProfile struct {
	// REQUIRED; The properties of the upgrade profile.
	Properties *ProvisionedClusterUpgradeProfileProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ProvisionedClusterUpgradeProfileProperties - Control plane and agent pool upgrade profiles.
type ProvisionedClusterUpgradeProfileProperties struct {
	// REQUIRED; The list of available kubernetes version upgrades for the control plane.
	ControlPlaneProfile *ProvisionedClusterPoolUpgradeProfile

	// READ-ONLY; Provisioning state of the resource
	ProvisioningState *ResourceProvisioningState
}

// StorageProfile - The storage configuration profile for the provisioned cluster.
type StorageProfile struct {
	// NFS CSI Driver settings for the storage profile.
	NfsCsiDriver *StorageProfileNfsCSIDriver

	// SMB CSI Driver settings for the storage profile.
	SmbCsiDriver *StorageProfileSmbCSIDriver
}

// StorageProfileNfsCSIDriver - NFS CSI Driver settings for the storage profile.
type StorageProfileNfsCSIDriver struct {
	// Indicates whether to enable NFS CSI Driver. The default value is true.
	Enabled *bool
}

// StorageProfileSmbCSIDriver - SMB CSI Driver settings for the storage profile.
type StorageProfileSmbCSIDriver struct {
	// Indicates whether to enable SMB CSI Driver. The default value is true.
	Enabled *bool
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// VMSKUCapabilities - Describes the VM SKU capabilities like MemoryGB, vCPUs, etc.
type VMSKUCapabilities struct {
	// READ-ONLY; Name of the VM SKU capability
	Name *string

	// READ-ONLY; Value of the VM SKU capability
	Value *string
}

// VMSKUProfile - The list of supported VM SKUs.
type VMSKUProfile struct {
	// Extended location pointing to the underlying infrastructure
	ExtendedLocation *ExtendedLocation

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY
	Properties *VMSKUProfileProperties

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// VMSKUProfileList - The list of supported VM SKUs.
type VMSKUProfileList struct {
	NextLink *string
	Value    []*VMSKUProfile
}

type VMSKUProfileProperties struct {
	// List of supported VM SKUs.
	Values []*VMSKUProperties

	// READ-ONLY; Provisioning state of the resource
	ProvisioningState *ResourceProvisioningState
}

// VMSKUProperties - The profile for supported VM SKUs
type VMSKUProperties struct {
	// READ-ONLY; The list of name-value pairs to describe VM SKU capabilities like MemoryGB, vCPUs, etc.
	Capabilities []*VMSKUCapabilities

	// READ-ONLY; The name of the VM SKU
	Name *string

	// READ-ONLY; The type of resource the SKU applies to.
	ResourceType *string

	// READ-ONLY; The size of the VM SKU
	Size *string

	// READ-ONLY; The tier of the VM SKU
	Tier *string
}

// VirtualNetwork - The Virtual Network resource definition.
type VirtualNetwork struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Extended location pointing to the underlying infrastructure
	ExtendedLocation *VirtualNetworkExtendedLocation

	// Properties of the virtual network resource
	Properties *VirtualNetworkProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// VirtualNetworkExtendedLocation - Extended location pointing to the underlying infrastructure
type VirtualNetworkExtendedLocation struct {
	// ARM Id of the extended location.
	Name *string

	// The extended location type. Allowed value: 'CustomLocation'
	Type *ExtendedLocationTypes
}

// VirtualNetworkProperties - Properties of the virtual network resource
type VirtualNetworkProperties struct {
	// List of DNS server IP Addresses associated with the network
	DNSServers []*string

	// IP Address of the Gateway associated with the network
	Gateway *string

	// IP Address Prefix of the network
	IPAddressPrefix  *string
	InfraVnetProfile *VirtualNetworkPropertiesInfraVnetProfile

	// Range of IP Addresses for Kubernetes API Server and services if using HA Proxy load balancer
	VipPool []*VirtualNetworkPropertiesVipPoolItem

	// VLAN Id used by the network
	VlanID *int32

	// Range of IP Addresses for Kubernetes node VMs
	VmipPool []*VirtualNetworkPropertiesVmipPoolItem

	// READ-ONLY
	ProvisioningState *ProvisioningState

	// READ-ONLY; Status of the virtual network resource
	Status *VirtualNetworkPropertiesStatus
}

type VirtualNetworkPropertiesInfraVnetProfile struct {
	// Infrastructure network profile for HCI platform
	Hci *VirtualNetworkPropertiesInfraVnetProfileHci
}

// VirtualNetworkPropertiesInfraVnetProfileHci - Infrastructure network profile for HCI platform
type VirtualNetworkPropertiesInfraVnetProfileHci struct {
	// Group in MOC(Microsoft On-premises Cloud)
	MocGroup *string

	// Location in MOC(Microsoft On-premises Cloud)
	MocLocation *string

	// Virtual Network name in MOC(Microsoft On-premises Cloud)
	MocVnetName *string
}

// VirtualNetworkPropertiesStatus - Status of the virtual network resource
type VirtualNetworkPropertiesStatus struct {
	// The detailed status of the long running operation.
	OperationStatus *VirtualNetworkPropertiesStatusOperationStatus
}

// VirtualNetworkPropertiesStatusOperationStatus - The detailed status of the long running operation.
type VirtualNetworkPropertiesStatusOperationStatus struct {
	// The error if any from the operation.
	Error *VirtualNetworkPropertiesStatusOperationStatusError

	// The identifier of the operation.
	OperationID *string

	// The status of the operation.
	Status *string
}

// VirtualNetworkPropertiesStatusOperationStatusError - The error if any from the operation.
type VirtualNetworkPropertiesStatusOperationStatusError struct {
	// The error code from the operation.
	Code *string

	// The error message from the operation.
	Message *string
}

type VirtualNetworkPropertiesVipPoolItem struct {
	// Ending IP address for the IP Pool
	EndIP *string

	// Starting IP address for the IP Pool
	StartIP *string
}

type VirtualNetworkPropertiesVmipPoolItem struct {
	// Ending IP address for the IP Pool
	EndIP *string

	// Starting IP address for the IP Pool
	StartIP *string
}

// VirtualNetworksListResult - A list of virtual network resources.
type VirtualNetworksListResult struct {
	NextLink *string
	Value    []*VirtualNetwork
}

// VirtualNetworksPatch - The Virtual Network resource patch definition.
type VirtualNetworksPatch struct {
	// Resource tags
	Tags map[string]*string
}
