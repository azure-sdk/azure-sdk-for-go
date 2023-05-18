//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system metadata relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// AzureBareMetalInstanceProperties - Describes the properties of an AzureBareMetal instance.
type AzureBareMetalInstanceProperties struct {
	// Specifies the hardware settings for the AzureBareMetal instance.
	HardwareProfile *HardwareProfile

	// Specifies the network settings for the AzureBareMetal instance.
	NetworkProfile *NetworkProfile

	// Specifies the operating system settings for the AzureBareMetal instance.
	OSProfile *OSProfile

	// ARM ID of another AzureBareMetalInstance that will share a network with this AzureBareMetalInstance
	PartnerNodeID *string

	// Specifies the storage settings for the AzureBareMetal instance disks.
	StorageProfile *StorageProfile

	// READ-ONLY; Specifies the AzureBareMetal instance unique ID.
	AzureBareMetalInstanceID *string

	// READ-ONLY; Hardware revision of an AzureBareMetal instance
	HwRevision *string

	// READ-ONLY; Resource power state
	PowerState *AzureBareMetalInstancePowerStateEnum

	// READ-ONLY; State of provisioning of the AzureBareMetalInstance
	ProvisioningState *AzureBareMetalProvisioningStatesEnum

	// READ-ONLY; Resource proximity placement group
	ProximityPlacementGroup *string
}

// AzureBareMetalInstancesClientDeleteOptions contains the optional parameters for the AzureBareMetalInstancesClient.Delete
// method.
type AzureBareMetalInstancesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// AzureBareMetalInstancesClientGetOptions contains the optional parameters for the AzureBareMetalInstancesClient.Get method.
type AzureBareMetalInstancesClientGetOptions struct {
	// placeholder for future optional parameters
}

// AzureBareMetalInstancesClientListByResourceGroupOptions contains the optional parameters for the AzureBareMetalInstancesClient.NewListByResourceGroupPager
// method.
type AzureBareMetalInstancesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// AzureBareMetalInstancesClientListBySubscriptionOptions contains the optional parameters for the AzureBareMetalInstancesClient.NewListBySubscriptionPager
// method.
type AzureBareMetalInstancesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// AzureBareMetalInstancesClientPutOptions contains the optional parameters for the AzureBareMetalInstancesClient.Put method.
type AzureBareMetalInstancesClientPutOptions struct {
	// request body for put call
	RequestBodyParameters *AzureBareMetalInstance
}

// AzureBareMetalInstancesClientUpdateOptions contains the optional parameters for the AzureBareMetalInstancesClient.Update
// method.
type AzureBareMetalInstancesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// AzureBareMetalInstancesListResult - The response from the List AzureBareMetal Instances operation.
type AzureBareMetalInstancesListResult struct {
	// The URL to get the next set of AzureBareMetal instances.
	NextLink *string

	// The list of Azure BareMetal instances.
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

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system metadata relating to this resource.
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

// AzureBareMetalStorageInstancesClientCreateOptions contains the optional parameters for the AzureBareMetalStorageInstancesClient.Create
// method.
type AzureBareMetalStorageInstancesClientCreateOptions struct {
	// placeholder for future optional parameters
}

// AzureBareMetalStorageInstancesClientDeleteOptions contains the optional parameters for the AzureBareMetalStorageInstancesClient.Delete
// method.
type AzureBareMetalStorageInstancesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// AzureBareMetalStorageInstancesClientGetOptions contains the optional parameters for the AzureBareMetalStorageInstancesClient.Get
// method.
type AzureBareMetalStorageInstancesClientGetOptions struct {
	// placeholder for future optional parameters
}

// AzureBareMetalStorageInstancesClientListByResourceGroupOptions contains the optional parameters for the AzureBareMetalStorageInstancesClient.NewListByResourceGroupPager
// method.
type AzureBareMetalStorageInstancesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// AzureBareMetalStorageInstancesClientListBySubscriptionOptions contains the optional parameters for the AzureBareMetalStorageInstancesClient.NewListBySubscriptionPager
// method.
type AzureBareMetalStorageInstancesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// AzureBareMetalStorageInstancesClientUpdateOptions contains the optional parameters for the AzureBareMetalStorageInstancesClient.Update
// method.
type AzureBareMetalStorageInstancesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// AzureBareMetalStorageInstancesListResult - The response from the Get AzureBareMetalStorageInstances operation.
type AzureBareMetalStorageInstancesListResult struct {
	// The URL to get the next set of AzureBareMetalStorage instances.
	NextLink *string

	// The list of AzureBareMetalStorage instances.
	Value []*AzureBareMetalStorageInstance
}

// Disk - Specifies the disk information fo the AzureBareMetal instance
type Disk struct {
	// Specifies the size of an empty data disk in gigabytes.
	DiskSizeGB *int32

	// The disk name.
	Name *string

	// READ-ONLY; Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM
	// and therefore must be unique for each data disk attached to a VM.
	Lun *int32
}

// Display - Detailed BareMetal operation information
type Display struct {
	// READ-ONLY; The localized friendly description for the operation as shown to the user.
	Description *string

	// READ-ONLY; The localized friendly name for the operation as shown to the user.
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name.
	Provider *string

	// READ-ONLY; The localized friendly form of the resource type related to this action/operation.
	Resource *string
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

// HardwareProfile - Specifies the hardware settings for the AzureBareMetal instance.
type HardwareProfile struct {
	// READ-ONLY; Specifies the AzureBareMetal instance SKU.
	AzureBareMetalInstanceSize *AzureBareMetalInstanceSizeNamesEnum

	// READ-ONLY; Name of the hardware type (vendor and/or their product name)
	HardwareType *AzureBareMetalHardwareTypeNamesEnum
}

// IPAddress - Specifies the IP address of the network interface.
type IPAddress struct {
	// Specifies the IP address of the network interface.
	IPAddress *string
}

// NetworkProfile - Specifies the network settings for the AzureBareMetal instance disks.
type NetworkProfile struct {
	// Specifies the network interfaces for the AzureBareMetal instance.
	NetworkInterfaces []*IPAddress

	// READ-ONLY; Specifies the circuit id for connecting to express route.
	CircuitID *string
}

// OSProfile - Specifies the operating system settings for the AzureBareMetal instance.
type OSProfile struct {
	// Specifies the host OS name of the AzureBareMetal instance.
	ComputerName *string

	// Specifies the SSH public key used to access the operating system.
	SSHPublicKey *string

	// READ-ONLY; This property allows you to specify the type of the OS.
	OSType *string

	// READ-ONLY; Specifies version of operating system.
	Version *string
}

// Operation - AzureBareMetal operation information
type Operation struct {
	// Displayed AzureBareMetal operation information
	Display *Display

	// READ-ONLY; indicates whether an operation is a data action or not.
	IsDataAction *bool

	// READ-ONLY; The name of the operation being performed on this particular object. This name should match the action name
	// that appears in RBAC / the event service.
	Name *string
}

// OperationList - List of AzureBareMetal operations
type OperationList struct {
	// List of AzureBareMetal operations
	Value []*Operation
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// Result - Sample result definition
type Result struct {
	// Sample property of type string
	SampleProperty *string
}

// StorageBillingProperties - Describes the billing related details of the AzureBareMetalStorageInstance.
type StorageBillingProperties struct {
	// the SKU type that is provisioned
	AzureBareMetalStorageInstanceSize *string

	// the billing mode for the storage instance
	BillingMode *string
}

// StorageProfile - Specifies the storage settings for the AzureBareMetal instance disks.
type StorageProfile struct {
	// Specifies information about the operating system disk used by baremetal instance.
	OSDisks []*Disk

	// READ-ONLY; IP Address to connect to storage.
	NfsIPAddress *string
}

// StorageProperties - described the storage properties of the azure baremetalstorage instance
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

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}
