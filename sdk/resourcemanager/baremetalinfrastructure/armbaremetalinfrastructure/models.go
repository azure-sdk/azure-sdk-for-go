//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbaremetalinfrastructure

import "time"

// AzureBareMetalInstance - AzureBareMetal instance info on Azure (ARM properties and AzureBareMetal properties)
type AzureBareMetalInstance struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// AzureBareMetal instance properties
	Properties *AzureBareMetalInstanceProperties

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

// AzureBareMetalInstanceProperties - Describes the properties of an Azure Bare Metal Instance.
type AzureBareMetalInstanceProperties struct {
	// Specifies the hardware settings for the Azure Bare Metal Instance.
	HardwareProfile *HardwareProfile

	// Specifies the network settings for the Azure Bare Metal Instance.
	NetworkProfile *NetworkProfile

	// Specifies the operating system settings for the Azure Bare Metal Instance.
	OSProfile *OSProfile

	// ARM ID of another AzureBareMetalInstance that will share a network with this AzureBareMetalInstance
	PartnerNodeID *string

	// Specifies the storage settings for the Azure Bare Metal Instance disks.
	StorageProfile *StorageProfile

	// READ-ONLY; Specifies the Azure Bare Metal Instance unique ID.
	AzureBareMetalInstanceID *string

	// READ-ONLY; Hardware revision of an Azure Bare Metal Instance
	HwRevision *string

	// READ-ONLY; Resource power state
	PowerState *AzureBareMetalInstancePowerStateEnum

	// READ-ONLY; State of provisioning of the AzureBareMetalInstance
	ProvisioningState *AzureBareMetalProvisioningStatesEnum

	// READ-ONLY; Resource proximity placement group
	ProximityPlacementGroup *string
}

// AzureBareMetalInstancesListResult - The response from the List Azure Bare Metal Instances operation.
type AzureBareMetalInstancesListResult struct {
	// The URL to get the next set of Azure Bare Metal Instances.
	NextLink *string

	// The list of Azure Bare Metal Instances.
	Value []*AzureBareMetalInstance
}

// AzureBareMetalStorageInstance info on Azure (ARM properties and AzureBareMetalStorage properties)
type AzureBareMetalStorageInstance struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// AzureBareMetalStorageInstance properties
	Properties *AzureBareMetalStorageInstanceProperties

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

// AzureBareMetalStorageInstanceProperties - Describes the properties of an AzureBareMetalStorageInstance.
type AzureBareMetalStorageInstanceProperties struct {
	// Specifies the AzureBareMetaStorageInstance unique ID.
	AzureBareMetalStorageInstanceUniqueIdentifier *string

	// Specifies the storage properties for the AzureBareMetalStorage instance.
	StorageProperties *StorageProperties
}

// AzureBareMetalStorageInstancesListResult - The response from the Get AzureBareMetalStorageInstances operation.
type AzureBareMetalStorageInstancesListResult struct {
	// The URL to get the next set of AzureBareMetalStorage instances.
	NextLink *string

	// The list of AzureBareMetalStorage instances.
	Value []*AzureBareMetalStorageInstance
}

// Disk - Specifies the disk information fo the Azure Bare Metal Instance
type Disk struct {
	// Specifies the size of an empty data disk in gigabytes.
	DiskSizeGB *int32

	// The disk name.
	Name *string

	// READ-ONLY; Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM
	// and therefore must be unique for each data disk attached to a VM.
	Lun *int32
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// ForceState - The active state empowers the server with the ability to forcefully terminate and halt any existing processes
// that may be running on the server
type ForceState struct {
	// Whether to force restart by shutting all processes.
	ForceState *AzureBareMetalInstanceForcePowerState
}

// HardwareProfile - Specifies the hardware settings for the Azure Bare Metal Instance.
type HardwareProfile struct {
	// READ-ONLY; Specifies the Azure Bare Metal Instance SKU.
	AzureBareMetalInstanceSize *AzureBareMetalInstanceSizeNamesEnum

	// READ-ONLY; Name of the hardware type (vendor and/or their product name)
	HardwareType *AzureBareMetalHardwareTypeNamesEnum
}

// NetworkInterface - Specifies the network interfaces of a bare metal resource.
type NetworkInterface struct {
	// Specifies the IP address of the network interface.
	IPAddress *string
}

// NetworkProfile - Specifies the network settings for the Azure Bare Metal Instance disks.
type NetworkProfile struct {
	// Specifies the network interfaces for the Azure Bare Metal Instance.
	NetworkInterfaces []*NetworkInterface

	// READ-ONLY; Specifies the circuit id for connecting to express route.
	CircuitID *string
}

// OSProfile - Specifies the operating system settings for the Azure Bare Metal instance.
type OSProfile struct {
	// Specifies the host OS name of the Azure Bare Metal instance.
	ComputerName *string

	// Specifies the SSH public key used to access the operating system.
	SSHPublicKey *string

	// READ-ONLY; This property allows you to specify the type of the OS.
	OSType *string

	// READ-ONLY; Specifies version of operating system.
	Version *string
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

// OperationStatus - The OperationStatus object returns the state of an asynchronous operation.
type OperationStatus struct {
	// An error from the Azure Bare Metal Infrastructure service.
	Error *OperationStatusError

	// Unique Operation Status Identifier.
	Name *string

	// Start Time when the operation was initially executed.
	StartTime *string

	// Status of the operation.
	Status *AsyncOperationStatus
}

// OperationStatusError - An error from the Azure Bare Metal Infrastructure service.
type OperationStatusError struct {
	// Server-defined set of error codes.
	Code *string

	// Human-readable representation of the error.
	Message *string
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// StorageBillingProperties - Describes the billing related details of the AzureBareMetalStorageInstance.
type StorageBillingProperties struct {
	// the SKU type that is provisioned
	AzureBareMetalStorageInstanceSize *string

	// the billing mode for the storage instance
	BillingMode *string
}

// StorageProfile - Specifies the storage settings for the Azure Bare Metal instance disks.
type StorageProfile struct {
	// Specifies information about the operating system disk used by bare metal instance.
	OSDisks []*Disk

	// READ-ONLY; IP Address to connect to storage.
	NfsIPAddress *string
}

// StorageProperties - described the storage properties of the azure bare metal storage instance
type StorageProperties struct {
	// the kind of storage instance
	Generation *string

	// the hardware type of the storage instance
	HardwareType *string

	// the offering type for which the resource is getting provisioned
	OfferingType *string

	// State of provisioning of the AzureBareMetalStorageInstance
	ProvisioningState *ProvisioningState

	// the billing related information for the resource
	StorageBillingProperties *StorageBillingProperties

	// the storage protocol for which the resource is getting provisioned
	StorageType *string

	// the workload for which the resource is getting provisioned
	WorkloadType *string
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

// Tags field of the AzureBareMetal/AzureBareMetaStorage instance.
type Tags struct {
	// Tags field of the AzureBareMetal/AzureBareMetaStorage instance.
	Tags map[string]*string
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

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
