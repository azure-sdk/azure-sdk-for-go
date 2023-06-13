//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerinstance

const (
	moduleName    = "armcontainerinstance"
	moduleVersion = "v2.3.1"
)

// ContainerGroupIPAddressType - Specifies if the IP is exposed to the public internet or private VNET.
type ContainerGroupIPAddressType string

const (
	ContainerGroupIPAddressTypePrivate ContainerGroupIPAddressType = "Private"
	ContainerGroupIPAddressTypePublic  ContainerGroupIPAddressType = "Public"
)

// PossibleContainerGroupIPAddressTypeValues returns the possible values for the ContainerGroupIPAddressType const type.
func PossibleContainerGroupIPAddressTypeValues() []ContainerGroupIPAddressType {
	return []ContainerGroupIPAddressType{
		ContainerGroupIPAddressTypePrivate,
		ContainerGroupIPAddressTypePublic,
	}
}

// ContainerGroupNetworkProtocol - The protocol associated with the port.
type ContainerGroupNetworkProtocol string

const (
	ContainerGroupNetworkProtocolTCP ContainerGroupNetworkProtocol = "TCP"
	ContainerGroupNetworkProtocolUDP ContainerGroupNetworkProtocol = "UDP"
)

// PossibleContainerGroupNetworkProtocolValues returns the possible values for the ContainerGroupNetworkProtocol const type.
func PossibleContainerGroupNetworkProtocolValues() []ContainerGroupNetworkProtocol {
	return []ContainerGroupNetworkProtocol{
		ContainerGroupNetworkProtocolTCP,
		ContainerGroupNetworkProtocolUDP,
	}
}

// ContainerGroupPriority - The priority of the container group.
type ContainerGroupPriority string

const (
	ContainerGroupPriorityRegular ContainerGroupPriority = "Regular"
	ContainerGroupPrioritySpot    ContainerGroupPriority = "Spot"
)

// PossibleContainerGroupPriorityValues returns the possible values for the ContainerGroupPriority const type.
func PossibleContainerGroupPriorityValues() []ContainerGroupPriority {
	return []ContainerGroupPriority{
		ContainerGroupPriorityRegular,
		ContainerGroupPrioritySpot,
	}
}

// ContainerGroupRestartPolicy - Restart policy for all containers within the container group.
// * Always Always restart
// * OnFailure Restart on failure
// * Never Never restart
type ContainerGroupRestartPolicy string

const (
	ContainerGroupRestartPolicyAlways    ContainerGroupRestartPolicy = "Always"
	ContainerGroupRestartPolicyNever     ContainerGroupRestartPolicy = "Never"
	ContainerGroupRestartPolicyOnFailure ContainerGroupRestartPolicy = "OnFailure"
)

// PossibleContainerGroupRestartPolicyValues returns the possible values for the ContainerGroupRestartPolicy const type.
func PossibleContainerGroupRestartPolicyValues() []ContainerGroupRestartPolicy {
	return []ContainerGroupRestartPolicy{
		ContainerGroupRestartPolicyAlways,
		ContainerGroupRestartPolicyNever,
		ContainerGroupRestartPolicyOnFailure,
	}
}

// ContainerGroupSKU - The container group SKU.
type ContainerGroupSKU string

const (
	ContainerGroupSKUConfidential ContainerGroupSKU = "Confidential"
	ContainerGroupSKUDedicated    ContainerGroupSKU = "Dedicated"
	ContainerGroupSKUStandard     ContainerGroupSKU = "Standard"
)

// PossibleContainerGroupSKUValues returns the possible values for the ContainerGroupSKU const type.
func PossibleContainerGroupSKUValues() []ContainerGroupSKU {
	return []ContainerGroupSKU{
		ContainerGroupSKUConfidential,
		ContainerGroupSKUDedicated,
		ContainerGroupSKUStandard,
	}
}

// ContainerInstanceOperationsOrigin - The intended executor of the operation.
type ContainerInstanceOperationsOrigin string

const (
	ContainerInstanceOperationsOriginSystem ContainerInstanceOperationsOrigin = "System"
	ContainerInstanceOperationsOriginUser   ContainerInstanceOperationsOrigin = "User"
)

// PossibleContainerInstanceOperationsOriginValues returns the possible values for the ContainerInstanceOperationsOrigin const type.
func PossibleContainerInstanceOperationsOriginValues() []ContainerInstanceOperationsOrigin {
	return []ContainerInstanceOperationsOrigin{
		ContainerInstanceOperationsOriginSystem,
		ContainerInstanceOperationsOriginUser,
	}
}

// ContainerNetworkProtocol - The protocol associated with the port.
type ContainerNetworkProtocol string

const (
	ContainerNetworkProtocolTCP ContainerNetworkProtocol = "TCP"
	ContainerNetworkProtocolUDP ContainerNetworkProtocol = "UDP"
)

// PossibleContainerNetworkProtocolValues returns the possible values for the ContainerNetworkProtocol const type.
func PossibleContainerNetworkProtocolValues() []ContainerNetworkProtocol {
	return []ContainerNetworkProtocol{
		ContainerNetworkProtocolTCP,
		ContainerNetworkProtocolUDP,
	}
}

// DNSNameLabelReusePolicy - The value representing the security enum. The 'Unsecure' value is the default value if not selected
// and means the object's domain name label is not secured against subdomain takeover. The
// 'TenantReuse' value is the default value if selected and means the object's domain name label can be reused within the
// same tenant. The 'SubscriptionReuse' value means the object's domain name label
// can be reused within the same subscription. The 'ResourceGroupReuse' value means the object's domain name label can be
// reused within the same resource group. The 'NoReuse' value means the object's
// domain name label cannot be reused within the same resource group, subscription, or tenant.
type DNSNameLabelReusePolicy string

const (
	DNSNameLabelReusePolicyNoreuse            DNSNameLabelReusePolicy = "Noreuse"
	DNSNameLabelReusePolicyResourceGroupReuse DNSNameLabelReusePolicy = "ResourceGroupReuse"
	DNSNameLabelReusePolicySubscriptionReuse  DNSNameLabelReusePolicy = "SubscriptionReuse"
	DNSNameLabelReusePolicyTenantReuse        DNSNameLabelReusePolicy = "TenantReuse"
	DNSNameLabelReusePolicyUnsecure           DNSNameLabelReusePolicy = "Unsecure"
)

// PossibleDNSNameLabelReusePolicyValues returns the possible values for the DNSNameLabelReusePolicy const type.
func PossibleDNSNameLabelReusePolicyValues() []DNSNameLabelReusePolicy {
	return []DNSNameLabelReusePolicy{
		DNSNameLabelReusePolicyNoreuse,
		DNSNameLabelReusePolicyResourceGroupReuse,
		DNSNameLabelReusePolicySubscriptionReuse,
		DNSNameLabelReusePolicyTenantReuse,
		DNSNameLabelReusePolicyUnsecure,
	}
}

// GpuSKU - The SKU of the GPU resource.
type GpuSKU string

const (
	GpuSKUK80  GpuSKU = "K80"
	GpuSKUP100 GpuSKU = "P100"
	GpuSKUV100 GpuSKU = "V100"
)

// PossibleGpuSKUValues returns the possible values for the GpuSKU const type.
func PossibleGpuSKUValues() []GpuSKU {
	return []GpuSKU{
		GpuSKUK80,
		GpuSKUP100,
		GpuSKUV100,
	}
}

// LogAnalyticsLogType - The log type to be used.
type LogAnalyticsLogType string

const (
	LogAnalyticsLogTypeContainerInsights     LogAnalyticsLogType = "ContainerInsights"
	LogAnalyticsLogTypeContainerInstanceLogs LogAnalyticsLogType = "ContainerInstanceLogs"
)

// PossibleLogAnalyticsLogTypeValues returns the possible values for the LogAnalyticsLogType const type.
func PossibleLogAnalyticsLogTypeValues() []LogAnalyticsLogType {
	return []LogAnalyticsLogType{
		LogAnalyticsLogTypeContainerInsights,
		LogAnalyticsLogTypeContainerInstanceLogs,
	}
}

// OperatingSystemTypes - The operating system type required by the containers in the container group.
type OperatingSystemTypes string

const (
	OperatingSystemTypesLinux   OperatingSystemTypes = "Linux"
	OperatingSystemTypesWindows OperatingSystemTypes = "Windows"
)

// PossibleOperatingSystemTypesValues returns the possible values for the OperatingSystemTypes const type.
func PossibleOperatingSystemTypesValues() []OperatingSystemTypes {
	return []OperatingSystemTypes{
		OperatingSystemTypesLinux,
		OperatingSystemTypesWindows,
	}
}

// ResourceIdentityType - The type of identity used for the container group. The type 'SystemAssigned, UserAssigned' includes
// both an implicitly created identity and a set of user assigned identities. The type 'None' will
// remove any identities from the container group.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = "UserAssigned"
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	ResourceIdentityTypeNone                       ResourceIdentityType = "None"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeUserAssigned,
		ResourceIdentityTypeSystemAssignedUserAssigned,
		ResourceIdentityTypeNone,
	}
}

// Scheme - The scheme.
type Scheme string

const (
	SchemeHTTP  Scheme = "http"
	SchemeHTTPS Scheme = "https"
)

// PossibleSchemeValues returns the possible values for the Scheme const type.
func PossibleSchemeValues() []Scheme {
	return []Scheme{
		SchemeHTTP,
		SchemeHTTPS,
	}
}
