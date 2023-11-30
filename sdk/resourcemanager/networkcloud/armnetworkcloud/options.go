//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetworkcloud

// AgentPoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the AgentPoolsClient.BeginCreateOrUpdate
// method.
type AgentPoolsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AgentPoolsClientBeginDeleteOptions contains the optional parameters for the AgentPoolsClient.BeginDelete method.
type AgentPoolsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AgentPoolsClientBeginUpdateOptions contains the optional parameters for the AgentPoolsClient.BeginUpdate method.
type AgentPoolsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AgentPoolsClientGetOptions contains the optional parameters for the AgentPoolsClient.Get method.
type AgentPoolsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AgentPoolsClientListByKubernetesClusterOptions contains the optional parameters for the AgentPoolsClient.NewListByKubernetesClusterPager
// method.
type AgentPoolsClientListByKubernetesClusterOptions struct {
	// placeholder for future optional parameters
}

// BareMetalMachineKeySetsClientBeginCreateOrUpdateOptions contains the optional parameters for the BareMetalMachineKeySetsClient.BeginCreateOrUpdate
// method.
type BareMetalMachineKeySetsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachineKeySetsClientBeginDeleteOptions contains the optional parameters for the BareMetalMachineKeySetsClient.BeginDelete
// method.
type BareMetalMachineKeySetsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachineKeySetsClientBeginUpdateOptions contains the optional parameters for the BareMetalMachineKeySetsClient.BeginUpdate
// method.
type BareMetalMachineKeySetsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachineKeySetsClientGetOptions contains the optional parameters for the BareMetalMachineKeySetsClient.Get method.
type BareMetalMachineKeySetsClientGetOptions struct {
	// placeholder for future optional parameters
}

// BareMetalMachineKeySetsClientListByClusterOptions contains the optional parameters for the BareMetalMachineKeySetsClient.NewListByClusterPager
// method.
type BareMetalMachineKeySetsClientListByClusterOptions struct {
	// placeholder for future optional parameters
}

// BareMetalMachinesClientBeginCordonOptions contains the optional parameters for the BareMetalMachinesClient.BeginCordon
// method.
type BareMetalMachinesClientBeginCordonOptions struct {
	// The request body.
	BareMetalMachineCordonParameters *BareMetalMachineCordonParameters

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachinesClientBeginCreateOrUpdateOptions contains the optional parameters for the BareMetalMachinesClient.BeginCreateOrUpdate
// method.
type BareMetalMachinesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachinesClientBeginDeleteOptions contains the optional parameters for the BareMetalMachinesClient.BeginDelete
// method.
type BareMetalMachinesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachinesClientBeginPowerOffOptions contains the optional parameters for the BareMetalMachinesClient.BeginPowerOff
// method.
type BareMetalMachinesClientBeginPowerOffOptions struct {
	// The request body.
	BareMetalMachinePowerOffParameters *BareMetalMachinePowerOffParameters

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachinesClientBeginReimageOptions contains the optional parameters for the BareMetalMachinesClient.BeginReimage
// method.
type BareMetalMachinesClientBeginReimageOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachinesClientBeginReplaceOptions contains the optional parameters for the BareMetalMachinesClient.BeginReplace
// method.
type BareMetalMachinesClientBeginReplaceOptions struct {
	// The request body.
	BareMetalMachineReplaceParameters *BareMetalMachineReplaceParameters

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachinesClientBeginRestartOptions contains the optional parameters for the BareMetalMachinesClient.BeginRestart
// method.
type BareMetalMachinesClientBeginRestartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachinesClientBeginRunCommandOptions contains the optional parameters for the BareMetalMachinesClient.BeginRunCommand
// method.
type BareMetalMachinesClientBeginRunCommandOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachinesClientBeginRunDataExtractsOptions contains the optional parameters for the BareMetalMachinesClient.BeginRunDataExtracts
// method.
type BareMetalMachinesClientBeginRunDataExtractsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachinesClientBeginRunReadCommandsOptions contains the optional parameters for the BareMetalMachinesClient.BeginRunReadCommands
// method.
type BareMetalMachinesClientBeginRunReadCommandsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachinesClientBeginStartOptions contains the optional parameters for the BareMetalMachinesClient.BeginStart method.
type BareMetalMachinesClientBeginStartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachinesClientBeginUncordonOptions contains the optional parameters for the BareMetalMachinesClient.BeginUncordon
// method.
type BareMetalMachinesClientBeginUncordonOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachinesClientBeginUpdateOptions contains the optional parameters for the BareMetalMachinesClient.BeginUpdate
// method.
type BareMetalMachinesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BareMetalMachinesClientGetOptions contains the optional parameters for the BareMetalMachinesClient.Get method.
type BareMetalMachinesClientGetOptions struct {
	// placeholder for future optional parameters
}

// BareMetalMachinesClientListByResourceGroupOptions contains the optional parameters for the BareMetalMachinesClient.NewListByResourceGroupPager
// method.
type BareMetalMachinesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// BareMetalMachinesClientListBySubscriptionOptions contains the optional parameters for the BareMetalMachinesClient.NewListBySubscriptionPager
// method.
type BareMetalMachinesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// BmcKeySetsClientBeginCreateOrUpdateOptions contains the optional parameters for the BmcKeySetsClient.BeginCreateOrUpdate
// method.
type BmcKeySetsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BmcKeySetsClientBeginDeleteOptions contains the optional parameters for the BmcKeySetsClient.BeginDelete method.
type BmcKeySetsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BmcKeySetsClientBeginUpdateOptions contains the optional parameters for the BmcKeySetsClient.BeginUpdate method.
type BmcKeySetsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BmcKeySetsClientGetOptions contains the optional parameters for the BmcKeySetsClient.Get method.
type BmcKeySetsClientGetOptions struct {
	// placeholder for future optional parameters
}

// BmcKeySetsClientListByClusterOptions contains the optional parameters for the BmcKeySetsClient.NewListByClusterPager method.
type BmcKeySetsClientListByClusterOptions struct {
	// placeholder for future optional parameters
}

// CloudServicesNetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the CloudServicesNetworksClient.BeginCreateOrUpdate
// method.
type CloudServicesNetworksClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CloudServicesNetworksClientBeginDeleteOptions contains the optional parameters for the CloudServicesNetworksClient.BeginDelete
// method.
type CloudServicesNetworksClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CloudServicesNetworksClientBeginUpdateOptions contains the optional parameters for the CloudServicesNetworksClient.BeginUpdate
// method.
type CloudServicesNetworksClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CloudServicesNetworksClientGetOptions contains the optional parameters for the CloudServicesNetworksClient.Get method.
type CloudServicesNetworksClientGetOptions struct {
	// placeholder for future optional parameters
}

// CloudServicesNetworksClientListByResourceGroupOptions contains the optional parameters for the CloudServicesNetworksClient.NewListByResourceGroupPager
// method.
type CloudServicesNetworksClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// CloudServicesNetworksClientListBySubscriptionOptions contains the optional parameters for the CloudServicesNetworksClient.NewListBySubscriptionPager
// method.
type CloudServicesNetworksClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ClusterManagersClientBeginCreateOrUpdateOptions contains the optional parameters for the ClusterManagersClient.BeginCreateOrUpdate
// method.
type ClusterManagersClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClusterManagersClientBeginDeleteOptions contains the optional parameters for the ClusterManagersClient.BeginDelete method.
type ClusterManagersClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClusterManagersClientGetOptions contains the optional parameters for the ClusterManagersClient.Get method.
type ClusterManagersClientGetOptions struct {
	// placeholder for future optional parameters
}

// ClusterManagersClientListByResourceGroupOptions contains the optional parameters for the ClusterManagersClient.NewListByResourceGroupPager
// method.
type ClusterManagersClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ClusterManagersClientListBySubscriptionOptions contains the optional parameters for the ClusterManagersClient.NewListBySubscriptionPager
// method.
type ClusterManagersClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ClusterManagersClientUpdateOptions contains the optional parameters for the ClusterManagersClient.Update method.
type ClusterManagersClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientBeginCreateOrUpdateOptions contains the optional parameters for the ClustersClient.BeginCreateOrUpdate method.
type ClustersClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientBeginDeleteOptions contains the optional parameters for the ClustersClient.BeginDelete method.
type ClustersClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientBeginDeployOptions contains the optional parameters for the ClustersClient.BeginDeploy method.
type ClustersClientBeginDeployOptions struct {
	// The request body.
	ClusterDeployParameters *ClusterDeployParameters

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientBeginScanRuntimeOptions contains the optional parameters for the ClustersClient.BeginScanRuntime method.
type ClustersClientBeginScanRuntimeOptions struct {
	// The request body.
	ClusterScanRuntimeParameters *ClusterScanRuntimeParameters

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientBeginUpdateOptions contains the optional parameters for the ClustersClient.BeginUpdate method.
type ClustersClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientBeginUpdateVersionOptions contains the optional parameters for the ClustersClient.BeginUpdateVersion method.
type ClustersClientBeginUpdateVersionOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClustersClientGetOptions contains the optional parameters for the ClustersClient.Get method.
type ClustersClientGetOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientListByResourceGroupOptions contains the optional parameters for the ClustersClient.NewListByResourceGroupPager
// method.
type ClustersClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ClustersClientListBySubscriptionOptions contains the optional parameters for the ClustersClient.NewListBySubscriptionPager
// method.
type ClustersClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ConsolesClientBeginCreateOrUpdateOptions contains the optional parameters for the ConsolesClient.BeginCreateOrUpdate method.
type ConsolesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConsolesClientBeginDeleteOptions contains the optional parameters for the ConsolesClient.BeginDelete method.
type ConsolesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConsolesClientBeginUpdateOptions contains the optional parameters for the ConsolesClient.BeginUpdate method.
type ConsolesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConsolesClientGetOptions contains the optional parameters for the ConsolesClient.Get method.
type ConsolesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ConsolesClientListByVirtualMachineOptions contains the optional parameters for the ConsolesClient.NewListByVirtualMachinePager
// method.
type ConsolesClientListByVirtualMachineOptions struct {
	// placeholder for future optional parameters
}

// KubernetesClustersClientBeginCreateOrUpdateOptions contains the optional parameters for the KubernetesClustersClient.BeginCreateOrUpdate
// method.
type KubernetesClustersClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// KubernetesClustersClientBeginDeleteOptions contains the optional parameters for the KubernetesClustersClient.BeginDelete
// method.
type KubernetesClustersClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// KubernetesClustersClientBeginRestartNodeOptions contains the optional parameters for the KubernetesClustersClient.BeginRestartNode
// method.
type KubernetesClustersClientBeginRestartNodeOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// KubernetesClustersClientBeginUpdateOptions contains the optional parameters for the KubernetesClustersClient.BeginUpdate
// method.
type KubernetesClustersClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// KubernetesClustersClientGetOptions contains the optional parameters for the KubernetesClustersClient.Get method.
type KubernetesClustersClientGetOptions struct {
	// placeholder for future optional parameters
}

// KubernetesClustersClientListByResourceGroupOptions contains the optional parameters for the KubernetesClustersClient.NewListByResourceGroupPager
// method.
type KubernetesClustersClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// KubernetesClustersClientListBySubscriptionOptions contains the optional parameters for the KubernetesClustersClient.NewListBySubscriptionPager
// method.
type KubernetesClustersClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// L2NetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the L2NetworksClient.BeginCreateOrUpdate
// method.
type L2NetworksClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// L2NetworksClientBeginDeleteOptions contains the optional parameters for the L2NetworksClient.BeginDelete method.
type L2NetworksClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// L2NetworksClientGetOptions contains the optional parameters for the L2NetworksClient.Get method.
type L2NetworksClientGetOptions struct {
	// placeholder for future optional parameters
}

// L2NetworksClientListByResourceGroupOptions contains the optional parameters for the L2NetworksClient.NewListByResourceGroupPager
// method.
type L2NetworksClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// L2NetworksClientListBySubscriptionOptions contains the optional parameters for the L2NetworksClient.NewListBySubscriptionPager
// method.
type L2NetworksClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// L2NetworksClientUpdateOptions contains the optional parameters for the L2NetworksClient.Update method.
type L2NetworksClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// L3NetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the L3NetworksClient.BeginCreateOrUpdate
// method.
type L3NetworksClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// L3NetworksClientBeginDeleteOptions contains the optional parameters for the L3NetworksClient.BeginDelete method.
type L3NetworksClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// L3NetworksClientGetOptions contains the optional parameters for the L3NetworksClient.Get method.
type L3NetworksClientGetOptions struct {
	// placeholder for future optional parameters
}

// L3NetworksClientListByResourceGroupOptions contains the optional parameters for the L3NetworksClient.NewListByResourceGroupPager
// method.
type L3NetworksClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// L3NetworksClientListBySubscriptionOptions contains the optional parameters for the L3NetworksClient.NewListBySubscriptionPager
// method.
type L3NetworksClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// L3NetworksClientUpdateOptions contains the optional parameters for the L3NetworksClient.Update method.
type L3NetworksClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// MetricsConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the MetricsConfigurationsClient.BeginCreateOrUpdate
// method.
type MetricsConfigurationsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MetricsConfigurationsClientBeginDeleteOptions contains the optional parameters for the MetricsConfigurationsClient.BeginDelete
// method.
type MetricsConfigurationsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MetricsConfigurationsClientBeginUpdateOptions contains the optional parameters for the MetricsConfigurationsClient.BeginUpdate
// method.
type MetricsConfigurationsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MetricsConfigurationsClientGetOptions contains the optional parameters for the MetricsConfigurationsClient.Get method.
type MetricsConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MetricsConfigurationsClientListByClusterOptions contains the optional parameters for the MetricsConfigurationsClient.NewListByClusterPager
// method.
type MetricsConfigurationsClientListByClusterOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// RackSKUsClientGetOptions contains the optional parameters for the RackSKUsClient.Get method.
type RackSKUsClientGetOptions struct {
	// placeholder for future optional parameters
}

// RackSKUsClientListBySubscriptionOptions contains the optional parameters for the RackSKUsClient.NewListBySubscriptionPager
// method.
type RackSKUsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// RacksClientBeginCreateOrUpdateOptions contains the optional parameters for the RacksClient.BeginCreateOrUpdate method.
type RacksClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// RacksClientBeginDeleteOptions contains the optional parameters for the RacksClient.BeginDelete method.
type RacksClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// RacksClientBeginUpdateOptions contains the optional parameters for the RacksClient.BeginUpdate method.
type RacksClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// RacksClientGetOptions contains the optional parameters for the RacksClient.Get method.
type RacksClientGetOptions struct {
	// placeholder for future optional parameters
}

// RacksClientListByResourceGroupOptions contains the optional parameters for the RacksClient.NewListByResourceGroupPager
// method.
type RacksClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// RacksClientListBySubscriptionOptions contains the optional parameters for the RacksClient.NewListBySubscriptionPager method.
type RacksClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// StorageAppliancesClientBeginCreateOrUpdateOptions contains the optional parameters for the StorageAppliancesClient.BeginCreateOrUpdate
// method.
type StorageAppliancesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageAppliancesClientBeginDeleteOptions contains the optional parameters for the StorageAppliancesClient.BeginDelete
// method.
type StorageAppliancesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageAppliancesClientBeginDisableRemoteVendorManagementOptions contains the optional parameters for the StorageAppliancesClient.BeginDisableRemoteVendorManagement
// method.
type StorageAppliancesClientBeginDisableRemoteVendorManagementOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageAppliancesClientBeginEnableRemoteVendorManagementOptions contains the optional parameters for the StorageAppliancesClient.BeginEnableRemoteVendorManagement
// method.
type StorageAppliancesClientBeginEnableRemoteVendorManagementOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string

	// The request body.
	StorageApplianceEnableRemoteVendorManagementParameters *StorageApplianceEnableRemoteVendorManagementParameters
}

// StorageAppliancesClientBeginUpdateOptions contains the optional parameters for the StorageAppliancesClient.BeginUpdate
// method.
type StorageAppliancesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageAppliancesClientGetOptions contains the optional parameters for the StorageAppliancesClient.Get method.
type StorageAppliancesClientGetOptions struct {
	// placeholder for future optional parameters
}

// StorageAppliancesClientListByResourceGroupOptions contains the optional parameters for the StorageAppliancesClient.NewListByResourceGroupPager
// method.
type StorageAppliancesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// StorageAppliancesClientListBySubscriptionOptions contains the optional parameters for the StorageAppliancesClient.NewListBySubscriptionPager
// method.
type StorageAppliancesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// TrunkedNetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the TrunkedNetworksClient.BeginCreateOrUpdate
// method.
type TrunkedNetworksClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TrunkedNetworksClientBeginDeleteOptions contains the optional parameters for the TrunkedNetworksClient.BeginDelete method.
type TrunkedNetworksClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TrunkedNetworksClientGetOptions contains the optional parameters for the TrunkedNetworksClient.Get method.
type TrunkedNetworksClientGetOptions struct {
	// placeholder for future optional parameters
}

// TrunkedNetworksClientListByResourceGroupOptions contains the optional parameters for the TrunkedNetworksClient.NewListByResourceGroupPager
// method.
type TrunkedNetworksClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// TrunkedNetworksClientListBySubscriptionOptions contains the optional parameters for the TrunkedNetworksClient.NewListBySubscriptionPager
// method.
type TrunkedNetworksClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// TrunkedNetworksClientUpdateOptions contains the optional parameters for the TrunkedNetworksClient.Update method.
type TrunkedNetworksClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// VirtualMachinesClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualMachinesClient.BeginCreateOrUpdate
// method.
type VirtualMachinesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VirtualMachinesClientBeginDeleteOptions contains the optional parameters for the VirtualMachinesClient.BeginDelete method.
type VirtualMachinesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VirtualMachinesClientBeginPowerOffOptions contains the optional parameters for the VirtualMachinesClient.BeginPowerOff
// method.
type VirtualMachinesClientBeginPowerOffOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string

	// The request body.
	VirtualMachinePowerOffParameters *VirtualMachinePowerOffParameters
}

// VirtualMachinesClientBeginReimageOptions contains the optional parameters for the VirtualMachinesClient.BeginReimage method.
type VirtualMachinesClientBeginReimageOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VirtualMachinesClientBeginRestartOptions contains the optional parameters for the VirtualMachinesClient.BeginRestart method.
type VirtualMachinesClientBeginRestartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VirtualMachinesClientBeginStartOptions contains the optional parameters for the VirtualMachinesClient.BeginStart method.
type VirtualMachinesClientBeginStartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VirtualMachinesClientBeginUpdateOptions contains the optional parameters for the VirtualMachinesClient.BeginUpdate method.
type VirtualMachinesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VirtualMachinesClientGetOptions contains the optional parameters for the VirtualMachinesClient.Get method.
type VirtualMachinesClientGetOptions struct {
	// placeholder for future optional parameters
}

// VirtualMachinesClientListByResourceGroupOptions contains the optional parameters for the VirtualMachinesClient.NewListByResourceGroupPager
// method.
type VirtualMachinesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// VirtualMachinesClientListBySubscriptionOptions contains the optional parameters for the VirtualMachinesClient.NewListBySubscriptionPager
// method.
type VirtualMachinesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// VolumesClientBeginCreateOrUpdateOptions contains the optional parameters for the VolumesClient.BeginCreateOrUpdate method.
type VolumesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VolumesClientBeginDeleteOptions contains the optional parameters for the VolumesClient.BeginDelete method.
type VolumesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VolumesClientGetOptions contains the optional parameters for the VolumesClient.Get method.
type VolumesClientGetOptions struct {
	// placeholder for future optional parameters
}

// VolumesClientListByResourceGroupOptions contains the optional parameters for the VolumesClient.NewListByResourceGroupPager
// method.
type VolumesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// VolumesClientListBySubscriptionOptions contains the optional parameters for the VolumesClient.NewListBySubscriptionPager
// method.
type VolumesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// VolumesClientUpdateOptions contains the optional parameters for the VolumesClient.Update method.
type VolumesClientUpdateOptions struct {
	// placeholder for future optional parameters
}
