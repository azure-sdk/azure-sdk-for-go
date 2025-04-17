// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
	moduleVersion = "v1.2.0"
)

// AggregationType - Metric's aggregation type for e.g. (Average, Total)
type AggregationType string

const (
	AggregationTypeAverage AggregationType = "Average"
	AggregationTypeTotal   AggregationType = "Total"
)

// PossibleAggregationTypeValues returns the possible values for the AggregationType const type.
func PossibleAggregationTypeValues() []AggregationType {
	return []AggregationType{
		AggregationTypeAverage,
		AggregationTypeTotal,
	}
}

// CustomizationHostNameType - Type of host name
type CustomizationHostNameType string

const (
	CustomizationHostNameTypeCUSTOMNAME         CustomizationHostNameType = "CUSTOM_NAME"
	CustomizationHostNameTypeFIXED              CustomizationHostNameType = "FIXED"
	CustomizationHostNameTypePREFIXBASED        CustomizationHostNameType = "PREFIX_BASED"
	CustomizationHostNameTypeUSERDEFINED        CustomizationHostNameType = "USER_DEFINED"
	CustomizationHostNameTypeVIRTUALMACHINENAME CustomizationHostNameType = "VIRTUAL_MACHINE_NAME"
)

// PossibleCustomizationHostNameTypeValues returns the possible values for the CustomizationHostNameType const type.
func PossibleCustomizationHostNameTypeValues() []CustomizationHostNameType {
	return []CustomizationHostNameType{
		CustomizationHostNameTypeCUSTOMNAME,
		CustomizationHostNameTypeFIXED,
		CustomizationHostNameTypePREFIXBASED,
		CustomizationHostNameTypeUSERDEFINED,
		CustomizationHostNameTypeVIRTUALMACHINENAME,
	}
}

// CustomizationIPAddressType - Customization Specification ip type
type CustomizationIPAddressType string

const (
	CustomizationIPAddressTypeCUSTOM      CustomizationIPAddressType = "CUSTOM"
	CustomizationIPAddressTypeDHCPIP      CustomizationIPAddressType = "DHCP_IP"
	CustomizationIPAddressTypeFIXEDIP     CustomizationIPAddressType = "FIXED_IP"
	CustomizationIPAddressTypeUSERDEFINED CustomizationIPAddressType = "USER_DEFINED"
)

// PossibleCustomizationIPAddressTypeValues returns the possible values for the CustomizationIPAddressType const type.
func PossibleCustomizationIPAddressTypeValues() []CustomizationIPAddressType {
	return []CustomizationIPAddressType{
		CustomizationIPAddressTypeCUSTOM,
		CustomizationIPAddressTypeDHCPIP,
		CustomizationIPAddressTypeFIXEDIP,
		CustomizationIPAddressTypeUSERDEFINED,
	}
}

// CustomizationIdentityType - Identity type
type CustomizationIdentityType string

const (
	CustomizationIdentityTypeLINUX       CustomizationIdentityType = "LINUX"
	CustomizationIdentityTypeWINDOWS     CustomizationIdentityType = "WINDOWS"
	CustomizationIdentityTypeWINDOWSTEXT CustomizationIdentityType = "WINDOWS_TEXT"
)

// PossibleCustomizationIdentityTypeValues returns the possible values for the CustomizationIdentityType const type.
func PossibleCustomizationIdentityTypeValues() []CustomizationIdentityType {
	return []CustomizationIdentityType{
		CustomizationIdentityTypeLINUX,
		CustomizationIdentityTypeWINDOWS,
		CustomizationIdentityTypeWINDOWSTEXT,
	}
}

// CustomizationPolicyPropertiesType - The type of customization (Linux or Windows)
type CustomizationPolicyPropertiesType string

const (
	CustomizationPolicyPropertiesTypeLINUX   CustomizationPolicyPropertiesType = "LINUX"
	CustomizationPolicyPropertiesTypeWINDOWS CustomizationPolicyPropertiesType = "WINDOWS"
)

// PossibleCustomizationPolicyPropertiesTypeValues returns the possible values for the CustomizationPolicyPropertiesType const type.
func PossibleCustomizationPolicyPropertiesTypeValues() []CustomizationPolicyPropertiesType {
	return []CustomizationPolicyPropertiesType{
		CustomizationPolicyPropertiesTypeLINUX,
		CustomizationPolicyPropertiesTypeWINDOWS,
	}
}

// DiskIndependenceMode - Disk's independence mode type
type DiskIndependenceMode string

const (
	DiskIndependenceModeIndependentNonpersistent DiskIndependenceMode = "independent_nonpersistent"
	DiskIndependenceModeIndependentPersistent    DiskIndependenceMode = "independent_persistent"
	DiskIndependenceModePersistent               DiskIndependenceMode = "persistent"
)

// PossibleDiskIndependenceModeValues returns the possible values for the DiskIndependenceMode const type.
func PossibleDiskIndependenceModeValues() []DiskIndependenceMode {
	return []DiskIndependenceMode{
		DiskIndependenceModeIndependentNonpersistent,
		DiskIndependenceModeIndependentPersistent,
		DiskIndependenceModePersistent,
	}
}

// GuestOSNICCustomizationAllocation - IP address allocation method
type GuestOSNICCustomizationAllocation string

const (
	GuestOSNICCustomizationAllocationDynamic GuestOSNICCustomizationAllocation = "dynamic"
	GuestOSNICCustomizationAllocationStatic  GuestOSNICCustomizationAllocation = "static"
)

// PossibleGuestOSNICCustomizationAllocationValues returns the possible values for the GuestOSNICCustomizationAllocation const type.
func PossibleGuestOSNICCustomizationAllocationValues() []GuestOSNICCustomizationAllocation {
	return []GuestOSNICCustomizationAllocation{
		GuestOSNICCustomizationAllocationDynamic,
		GuestOSNICCustomizationAllocationStatic,
	}
}

// GuestOSType - The Guest OS type
type GuestOSType string

const (
	GuestOSTypeLinux   GuestOSType = "linux"
	GuestOSTypeOther   GuestOSType = "other"
	GuestOSTypeWindows GuestOSType = "windows"
)

// PossibleGuestOSTypeValues returns the possible values for the GuestOSType const type.
func PossibleGuestOSTypeValues() []GuestOSType {
	return []GuestOSType{
		GuestOSTypeLinux,
		GuestOSTypeOther,
		GuestOSTypeWindows,
	}
}

// NICType - NIC type
type NICType string

const (
	NICTypeE1000   NICType = "E1000"
	NICTypeE1000E  NICType = "E1000E"
	NICTypePCNET32 NICType = "PCNET32"
	NICTypeVMXNET  NICType = "VMXNET"
	NICTypeVMXNET2 NICType = "VMXNET2"
	NICTypeVMXNET3 NICType = "VMXNET3"
)

// PossibleNICTypeValues returns the possible values for the NICType const type.
func PossibleNICTypeValues() []NICType {
	return []NICType{
		NICTypeE1000,
		NICTypeE1000E,
		NICTypePCNET32,
		NICTypeVMXNET,
		NICTypeVMXNET2,
		NICTypeVMXNET3,
	}
}

// NodeStatus - Node status, indicates is private cloud set up on this node or not
type NodeStatus string

const (
	NodeStatusUnused NodeStatus = "unused"
	NodeStatusUsed   NodeStatus = "used"
)

// PossibleNodeStatusValues returns the possible values for the NodeStatus const type.
func PossibleNodeStatusValues() []NodeStatus {
	return []NodeStatus{
		NodeStatusUnused,
		NodeStatusUsed,
	}
}

// OnboardingStatus - indicates whether account onboarded or not in a given region
type OnboardingStatus string

const (
	OnboardingStatusNotOnBoarded     OnboardingStatus = "notOnBoarded"
	OnboardingStatusOnBoarded        OnboardingStatus = "onBoarded"
	OnboardingStatusOnBoarding       OnboardingStatus = "onBoarding"
	OnboardingStatusOnBoardingFailed OnboardingStatus = "onBoardingFailed"
)

// PossibleOnboardingStatusValues returns the possible values for the OnboardingStatus const type.
func PossibleOnboardingStatusValues() []OnboardingStatus {
	return []OnboardingStatus{
		OnboardingStatusNotOnBoarded,
		OnboardingStatusOnBoarded,
		OnboardingStatusOnBoarding,
		OnboardingStatusOnBoardingFailed,
	}
}

// OperationOrigin - The origin of operation
type OperationOrigin string

const (
	OperationOriginSystem     OperationOrigin = "system"
	OperationOriginUser       OperationOrigin = "user"
	OperationOriginUserSystem OperationOrigin = "user,system"
)

// PossibleOperationOriginValues returns the possible values for the OperationOrigin const type.
func PossibleOperationOriginValues() []OperationOrigin {
	return []OperationOrigin{
		OperationOriginSystem,
		OperationOriginUser,
		OperationOriginUserSystem,
	}
}

// StopMode - mode indicates a type of stop operation - reboot, suspend, shutdown or power-off
type StopMode string

const (
	StopModePoweroff StopMode = "poweroff"
	StopModeReboot   StopMode = "reboot"
	StopModeShutdown StopMode = "shutdown"
	StopModeSuspend  StopMode = "suspend"
)

// PossibleStopModeValues returns the possible values for the StopMode const type.
func PossibleStopModeValues() []StopMode {
	return []StopMode{
		StopModePoweroff,
		StopModeReboot,
		StopModeShutdown,
		StopModeSuspend,
	}
}

// UsageCount - The usages' unit
type UsageCount string

const (
	UsageCountBytes          UsageCount = "Bytes"
	UsageCountBytesPerSecond UsageCount = "BytesPerSecond"
	UsageCountCount          UsageCount = "Count"
	UsageCountCountPerSecond UsageCount = "CountPerSecond"
	UsageCountPercent        UsageCount = "Percent"
	UsageCountSeconds        UsageCount = "Seconds"
)

// PossibleUsageCountValues returns the possible values for the UsageCount const type.
func PossibleUsageCountValues() []UsageCount {
	return []UsageCount{
		UsageCountBytes,
		UsageCountBytesPerSecond,
		UsageCountCount,
		UsageCountCountPerSecond,
		UsageCountPercent,
		UsageCountSeconds,
	}
}

// VirtualMachineStatus - The status of Virtual machine
type VirtualMachineStatus string

const (
	VirtualMachineStatusDeallocating VirtualMachineStatus = "deallocating"
	VirtualMachineStatusDeleting     VirtualMachineStatus = "deleting"
	VirtualMachineStatusPoweredoff   VirtualMachineStatus = "poweredoff"
	VirtualMachineStatusRunning      VirtualMachineStatus = "running"
	VirtualMachineStatusSuspended    VirtualMachineStatus = "suspended"
	VirtualMachineStatusUpdating     VirtualMachineStatus = "updating"
)

// PossibleVirtualMachineStatusValues returns the possible values for the VirtualMachineStatus const type.
func PossibleVirtualMachineStatusValues() []VirtualMachineStatus {
	return []VirtualMachineStatus{
		VirtualMachineStatusDeallocating,
		VirtualMachineStatusDeleting,
		VirtualMachineStatusPoweredoff,
		VirtualMachineStatusRunning,
		VirtualMachineStatusSuspended,
		VirtualMachineStatusUpdating,
	}
}
