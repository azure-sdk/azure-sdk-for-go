//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmigrate

const (
	moduleName    = "armmigrate"
	moduleVersion = "v2.0.0"
)

type AgentConfigurationRebootStatus string

const (
	AgentConfigurationRebootStatusNotRebooted AgentConfigurationRebootStatus = "notRebooted"
	AgentConfigurationRebootStatusRebooted    AgentConfigurationRebootStatus = "rebooted"
	AgentConfigurationRebootStatusUnknown     AgentConfigurationRebootStatus = "unknown"
)

// PossibleAgentConfigurationRebootStatusValues returns the possible values for the AgentConfigurationRebootStatus const type.
func PossibleAgentConfigurationRebootStatusValues() []AgentConfigurationRebootStatus {
	return []AgentConfigurationRebootStatus{
		AgentConfigurationRebootStatusNotRebooted,
		AgentConfigurationRebootStatusRebooted,
		AgentConfigurationRebootStatusUnknown,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// CredentialType - Credential type of the run as account.
type CredentialType string

const (
	CredentialTypeHyperVFabric  CredentialType = "HyperVFabric"
	CredentialTypeLinuxGuest    CredentialType = "LinuxGuest"
	CredentialTypeLinuxServer   CredentialType = "LinuxServer"
	CredentialTypeVMwareFabric  CredentialType = "VMwareFabric"
	CredentialTypeWindowsGuest  CredentialType = "WindowsGuest"
	CredentialTypeWindowsServer CredentialType = "WindowsServer"
)

// PossibleCredentialTypeValues returns the possible values for the CredentialType const type.
func PossibleCredentialTypeValues() []CredentialType {
	return []CredentialType{
		CredentialTypeHyperVFabric,
		CredentialTypeLinuxGuest,
		CredentialTypeLinuxServer,
		CredentialTypeVMwareFabric,
		CredentialTypeWindowsGuest,
		CredentialTypeWindowsServer,
	}
}

// HighlyAvailable - Value indicating whether the VM is highly available.
type HighlyAvailable string

const (
	HighlyAvailableUnknown HighlyAvailable = "Unknown"
	HighlyAvailableNo      HighlyAvailable = "No"
	HighlyAvailableYes     HighlyAvailable = "Yes"
)

// PossibleHighlyAvailableValues returns the possible values for the HighlyAvailable const type.
func PossibleHighlyAvailableValues() []HighlyAvailable {
	return []HighlyAvailable{
		HighlyAvailableUnknown,
		HighlyAvailableNo,
		HighlyAvailableYes,
	}
}

type HypervisorConfigurationHypervisorType string

const (
	HypervisorConfigurationHypervisorTypeHyperv  HypervisorConfigurationHypervisorType = "hyperv"
	HypervisorConfigurationHypervisorTypeUnknown HypervisorConfigurationHypervisorType = "unknown"
)

// PossibleHypervisorConfigurationHypervisorTypeValues returns the possible values for the HypervisorConfigurationHypervisorType const type.
func PossibleHypervisorConfigurationHypervisorTypeValues() []HypervisorConfigurationHypervisorType {
	return []HypervisorConfigurationHypervisorType{
		HypervisorConfigurationHypervisorTypeHyperv,
		HypervisorConfigurationHypervisorTypeUnknown,
	}
}

type MachinePropertiesMonitoringState string

const (
	MachinePropertiesMonitoringStateDiscovered MachinePropertiesMonitoringState = "discovered"
	MachinePropertiesMonitoringStateMonitored  MachinePropertiesMonitoringState = "monitored"
)

// PossibleMachinePropertiesMonitoringStateValues returns the possible values for the MachinePropertiesMonitoringState const type.
func PossibleMachinePropertiesMonitoringStateValues() []MachinePropertiesMonitoringState {
	return []MachinePropertiesMonitoringState{
		MachinePropertiesMonitoringStateDiscovered,
		MachinePropertiesMonitoringStateMonitored,
	}
}

type MachinePropertiesVirtualizationState string

const (
	MachinePropertiesVirtualizationStateHypervisor MachinePropertiesVirtualizationState = "hypervisor"
	MachinePropertiesVirtualizationStatePhysical   MachinePropertiesVirtualizationState = "physical"
	MachinePropertiesVirtualizationStateUnknown    MachinePropertiesVirtualizationState = "unknown"
	MachinePropertiesVirtualizationStateVirtual    MachinePropertiesVirtualizationState = "virtual"
)

// PossibleMachinePropertiesVirtualizationStateValues returns the possible values for the MachinePropertiesVirtualizationState const type.
func PossibleMachinePropertiesVirtualizationStateValues() []MachinePropertiesVirtualizationState {
	return []MachinePropertiesVirtualizationState{
		MachinePropertiesVirtualizationStateHypervisor,
		MachinePropertiesVirtualizationStatePhysical,
		MachinePropertiesVirtualizationStateUnknown,
		MachinePropertiesVirtualizationStateVirtual,
	}
}

type MachineResourcesConfigurationCPUSpeedAccuracy string

const (
	MachineResourcesConfigurationCPUSpeedAccuracyActual    MachineResourcesConfigurationCPUSpeedAccuracy = "actual"
	MachineResourcesConfigurationCPUSpeedAccuracyEstimated MachineResourcesConfigurationCPUSpeedAccuracy = "estimated"
)

// PossibleMachineResourcesConfigurationCPUSpeedAccuracyValues returns the possible values for the MachineResourcesConfigurationCPUSpeedAccuracy const type.
func PossibleMachineResourcesConfigurationCPUSpeedAccuracyValues() []MachineResourcesConfigurationCPUSpeedAccuracy {
	return []MachineResourcesConfigurationCPUSpeedAccuracy{
		MachineResourcesConfigurationCPUSpeedAccuracyActual,
		MachineResourcesConfigurationCPUSpeedAccuracyEstimated,
	}
}

// MasterSitePropertiesPublicNetworkAccess - Gets or sets the state of public network access.
type MasterSitePropertiesPublicNetworkAccess string

const (
	MasterSitePropertiesPublicNetworkAccessDisabled     MasterSitePropertiesPublicNetworkAccess = "Disabled"
	MasterSitePropertiesPublicNetworkAccessEnabled      MasterSitePropertiesPublicNetworkAccess = "Enabled"
	MasterSitePropertiesPublicNetworkAccessNotSpecified MasterSitePropertiesPublicNetworkAccess = "NotSpecified"
)

// PossibleMasterSitePropertiesPublicNetworkAccessValues returns the possible values for the MasterSitePropertiesPublicNetworkAccess const type.
func PossibleMasterSitePropertiesPublicNetworkAccessValues() []MasterSitePropertiesPublicNetworkAccess {
	return []MasterSitePropertiesPublicNetworkAccess{
		MasterSitePropertiesPublicNetworkAccessDisabled,
		MasterSitePropertiesPublicNetworkAccessEnabled,
		MasterSitePropertiesPublicNetworkAccessNotSpecified,
	}
}

type OperatingSystemConfigurationBitness string

const (
	OperatingSystemConfigurationBitnessSixtyFourBit OperatingSystemConfigurationBitness = "64bit"
	OperatingSystemConfigurationBitnessThirtyTwoBit OperatingSystemConfigurationBitness = "32bit"
)

// PossibleOperatingSystemConfigurationBitnessValues returns the possible values for the OperatingSystemConfigurationBitness const type.
func PossibleOperatingSystemConfigurationBitnessValues() []OperatingSystemConfigurationBitness {
	return []OperatingSystemConfigurationBitness{
		OperatingSystemConfigurationBitnessSixtyFourBit,
		OperatingSystemConfigurationBitnessThirtyTwoBit,
	}
}

type OperatingSystemConfigurationFamily string

const (
	OperatingSystemConfigurationFamilyAix     OperatingSystemConfigurationFamily = "aix"
	OperatingSystemConfigurationFamilyLinux   OperatingSystemConfigurationFamily = "linux"
	OperatingSystemConfigurationFamilySolaris OperatingSystemConfigurationFamily = "solaris"
	OperatingSystemConfigurationFamilyUnknown OperatingSystemConfigurationFamily = "unknown"
	OperatingSystemConfigurationFamilyWindows OperatingSystemConfigurationFamily = "windows"
)

// PossibleOperatingSystemConfigurationFamilyValues returns the possible values for the OperatingSystemConfigurationFamily const type.
func PossibleOperatingSystemConfigurationFamilyValues() []OperatingSystemConfigurationFamily {
	return []OperatingSystemConfigurationFamily{
		OperatingSystemConfigurationFamilyAix,
		OperatingSystemConfigurationFamilyLinux,
		OperatingSystemConfigurationFamilySolaris,
		OperatingSystemConfigurationFamilyUnknown,
		OperatingSystemConfigurationFamilyWindows,
	}
}

type PrivateEndpointConnectionPropertiesProvisioningState string

const (
	PrivateEndpointConnectionPropertiesProvisioningStateAccepted   PrivateEndpointConnectionPropertiesProvisioningState = "Accepted"
	PrivateEndpointConnectionPropertiesProvisioningStateFailed     PrivateEndpointConnectionPropertiesProvisioningState = "Failed"
	PrivateEndpointConnectionPropertiesProvisioningStateInProgress PrivateEndpointConnectionPropertiesProvisioningState = "InProgress"
	PrivateEndpointConnectionPropertiesProvisioningStateSucceeded  PrivateEndpointConnectionPropertiesProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionPropertiesProvisioningStateValues returns the possible values for the PrivateEndpointConnectionPropertiesProvisioningState const type.
func PossiblePrivateEndpointConnectionPropertiesProvisioningStateValues() []PrivateEndpointConnectionPropertiesProvisioningState {
	return []PrivateEndpointConnectionPropertiesProvisioningState{
		PrivateEndpointConnectionPropertiesProvisioningStateAccepted,
		PrivateEndpointConnectionPropertiesProvisioningStateFailed,
		PrivateEndpointConnectionPropertiesProvisioningStateInProgress,
		PrivateEndpointConnectionPropertiesProvisioningStateSucceeded,
	}
}

type PrivateLinkServiceConnectionStateStatus string

const (
	PrivateLinkServiceConnectionStateStatusApproved     PrivateLinkServiceConnectionStateStatus = "Approved"
	PrivateLinkServiceConnectionStateStatusDisconnected PrivateLinkServiceConnectionStateStatus = "Disconnected"
	PrivateLinkServiceConnectionStateStatusPending      PrivateLinkServiceConnectionStateStatus = "Pending"
	PrivateLinkServiceConnectionStateStatusRejected     PrivateLinkServiceConnectionStateStatus = "Rejected"
)

// PossiblePrivateLinkServiceConnectionStateStatusValues returns the possible values for the PrivateLinkServiceConnectionStateStatus const type.
func PossiblePrivateLinkServiceConnectionStateStatusValues() []PrivateLinkServiceConnectionStateStatus {
	return []PrivateLinkServiceConnectionStateStatus{
		PrivateLinkServiceConnectionStateStatusApproved,
		PrivateLinkServiceConnectionStateStatusDisconnected,
		PrivateLinkServiceConnectionStateStatusPending,
		PrivateLinkServiceConnectionStateStatusRejected,
	}
}

// VirtualDiskMode - Disk mode property used for identifying independent disks.
type VirtualDiskMode string

const (
	VirtualDiskModeAppend                   VirtualDiskMode = "append"
	VirtualDiskModeIndependentNonpersistent VirtualDiskMode = "independent_nonpersistent"
	VirtualDiskModeIndependentPersistent    VirtualDiskMode = "independent_persistent"
	VirtualDiskModeNonpersistent            VirtualDiskMode = "nonpersistent"
	VirtualDiskModePersistent               VirtualDiskMode = "persistent"
	VirtualDiskModeUndoable                 VirtualDiskMode = "undoable"
)

// PossibleVirtualDiskModeValues returns the possible values for the VirtualDiskMode const type.
func PossibleVirtualDiskModeValues() []VirtualDiskMode {
	return []VirtualDiskMode{
		VirtualDiskModeAppend,
		VirtualDiskModeIndependentNonpersistent,
		VirtualDiskModeIndependentPersistent,
		VirtualDiskModeNonpersistent,
		VirtualDiskModePersistent,
		VirtualDiskModeUndoable,
	}
}

type VirtualMachineConfigurationVirtualMachineType string

const (
	VirtualMachineConfigurationVirtualMachineTypeHyperv    VirtualMachineConfigurationVirtualMachineType = "hyperv"
	VirtualMachineConfigurationVirtualMachineTypeLdom      VirtualMachineConfigurationVirtualMachineType = "ldom"
	VirtualMachineConfigurationVirtualMachineTypeLpar      VirtualMachineConfigurationVirtualMachineType = "lpar"
	VirtualMachineConfigurationVirtualMachineTypeUnknown   VirtualMachineConfigurationVirtualMachineType = "unknown"
	VirtualMachineConfigurationVirtualMachineTypeVirtualPc VirtualMachineConfigurationVirtualMachineType = "virtualPc"
	VirtualMachineConfigurationVirtualMachineTypeVmware    VirtualMachineConfigurationVirtualMachineType = "vmware"
	VirtualMachineConfigurationVirtualMachineTypeXen       VirtualMachineConfigurationVirtualMachineType = "xen"
)

// PossibleVirtualMachineConfigurationVirtualMachineTypeValues returns the possible values for the VirtualMachineConfigurationVirtualMachineType const type.
func PossibleVirtualMachineConfigurationVirtualMachineTypeValues() []VirtualMachineConfigurationVirtualMachineType {
	return []VirtualMachineConfigurationVirtualMachineType{
		VirtualMachineConfigurationVirtualMachineTypeHyperv,
		VirtualMachineConfigurationVirtualMachineTypeLdom,
		VirtualMachineConfigurationVirtualMachineTypeLpar,
		VirtualMachineConfigurationVirtualMachineTypeUnknown,
		VirtualMachineConfigurationVirtualMachineTypeVirtualPc,
		VirtualMachineConfigurationVirtualMachineTypeVmware,
		VirtualMachineConfigurationVirtualMachineTypeXen,
	}
}
