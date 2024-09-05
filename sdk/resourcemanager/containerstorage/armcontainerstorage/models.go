//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerstorage

import "time"

// Assignment Properties
type Assignment struct {
	// REQUIRED; Resource id for the assigned resource
	ID *string

	// READ-ONLY; Indicates if the assignment is in a usable state
	Status *AssignmentStatus
}

// AssignmentStatus - Status of the assignment resource
type AssignmentStatus struct {
	// REQUIRED; State of the assignment resource
	State *AssignmentStatusState

	// Reason for the status
	Message *string
}

// AzureDisk - Azure Disk Pool Properties
type AzureDisk struct {
	// Only required if individual disk selection is desired. Path to disk, e.g. :/dev/sda or WWN. Supports specifying multiple
	// disks (same syntax as tags).
	Disks []*Disk

	// Encryption specifies the encryption configuration for the Azure Disk pool
	Encryption *Encryption

	// Sku name
	SKUName *AzureDiskSKUName

	// READ-ONLY; Managed resource group for the pool.
	ResourceGroup *string
}

// AzureDiskUpdate - Azure Disk Pool Properties
type AzureDiskUpdate struct {
	// Only required if individual disk selection is desired. Path to disk, e.g. :/dev/sda or WWN. Supports specifying multiple
	// disks (same syntax as tags).
	Disks []*Disk
}

// Disk - Model for disk for that pool is using
type Disk struct {
	// REQUIRED; ID is the disk identifier visible to the OS. It is typically the WWN or disk ID in formats such as eui.e8238fa6bf530001001b448b45263379
	// or 0x5002cf6cbc5dd460
	ID *string

	// REQUIRED; Reference is the location of the disk in an external system.
	Reference *string
}

// ElasticSan - Elastic San Pool Properties
type ElasticSan struct {
	// Encryption specifies the encryption configuration for the Azure Disk pool
	Encryption *Encryption

	// Sku name
	SKUName *ElasticSanSKUName

	// READ-ONLY; Managed resource group for the pool.
	ResourceGroup *string
}

// ElasticSanVolumeProperties - Properties of the ElasticSAN iSCSI target
type ElasticSanVolumeProperties struct {
	// REQUIRED; iSCSI Target IQN (iSCSI Qualified Name); example: "iqn.2005-03.org.iscsi:server"
	TargetIqn *string

	// REQUIRED; iSCSI Target Portal Host Name
	TargetPortalHostname *string

	// REQUIRED; iSCSI Target Portal Port
	TargetPortalPort *int32
}

// Encryption key properties for the pool.
type Encryption struct {
	// REQUIRED; The name of the key vault key.
	KeyName *string

	// REQUIRED; The URI of the key vault.
	KeyVaultURI *string

	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity
}

// EphemeralDisk - Ephemeral Disk Pool Properties
type EphemeralDisk struct {
	// Only required if individual disk selection is desired. Path to disk, e.g. :/dev/sda or WWN. Supports specifying multiple
	// disks (same syntax as tags).
	Disks []*Disk

	// The number of data copies. Default 3.
	Replicas *int64
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
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

// Pool resource
type Pool struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *PoolProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PoolListResult - The response of a Pool list operation.
type PoolListResult struct {
	// REQUIRED; The Pool items on this page
	Value []*Pool

	// The link to the next page of items
	NextLink *string
}

// PoolProperties - Pool Properties
type PoolProperties struct {
	// REQUIRED; Type of the Pool: ephemeralDisk, azureDisk, or elasticsan.
	PoolType *PoolType

	// List of resources that should have access to the pool. Typically ARM references to AKS clusters or ACI Container Groups.
	// For local and standard this must be a single reference. For ElasticSAN there
	// can be many.
	Assignments []*Assignment

	// ReclaimPolicy defines what happens to the backend storage when StoragePool is deleted
	ReclaimPolicy *ReclaimPolicy

	// Resources represent the resources the pool should have.
	Resources *Resources

	// List of availability zones that resources can be created in.
	Zones []*Zone

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The operational status of the resource
	Status *ResourceOperationalStatus
}

// PoolType - Type of the Pool: ephemeralDisk, azureDisk, or elasticsan
type PoolType struct {
	// Disk Pool Properties
	AzureDisk *AzureDisk

	// Elastic San Pool Properties
	ElasticSan *ElasticSan

	// Ephemeral Pool Properties
	EphemeralDisk *EphemeralDisk
}

// PoolTypeUpdate - Type of the Pool: ephemeralDisk, azureDisk, or elasticsan
type PoolTypeUpdate struct {
	// Disk Pool Properties
	AzureDisk *AzureDiskUpdate

	// Elastic San Pool Properties
	ElasticSan any

	// Ephemeral Pool Properties
	EphemeralDisk *EphemeralDisk
}

// PoolUpdate - The type used for update operations of the Pool.
type PoolUpdate struct {
	// The resource-specific properties for this resource.
	Properties *PoolUpdateProperties

	// Resource tags.
	Tags map[string]*string
}

// PoolUpdateProperties - The updatable properties of the Pool.
type PoolUpdateProperties struct {
	// List of resources that should have access to the pool. Typically ARM references to AKS clusters or ACI Container Groups.
	// For local and standard this must be a single reference. For ElasticSAN there
	// can be many.
	Assignments []*Assignment

	// Type of the Pool: ephemeralDisk, azureDisk, or elasticsan.
	PoolType *PoolTypeUpdate

	// Resources represent the resources the pool should have.
	Resources *Resources
}

// Requests for capacity for the pool.
type Requests struct {
	// Requested capacity of the pool in GiB.
	Storage *int64
}

// ResourceOperationalStatus - Status of the resource
type ResourceOperationalStatus struct {
	// REQUIRED; state of the resource
	State *ResourceOperationStatusState

	// Reason for state.
	Message *string
}

// Resources - Resource Requests for the pool.
type Resources struct {
	// Requests for capacity for the pool.
	Requests *Requests
}

// Snapshot - Concrete proxy resource types can be created by aliasing this type using a specific property type.
type Snapshot struct {
	// The resource-specific properties for this resource.
	Properties *SnapshotProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SnapshotListResult - The response of a Snapshot list operation.
type SnapshotListResult struct {
	// REQUIRED; The Snapshot items on this page
	Value []*Snapshot

	// The link to the next page of items
	NextLink *string
}

// SnapshotProperties - Volume Snapshot Properties
type SnapshotProperties struct {
	// REQUIRED; Reference to the source volume
	Source *string

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The status of the resource.
	Status *ResourceOperationalStatus
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

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}

// Volume - Concrete proxy resource types can be created by aliasing this type using a specific property type.
type Volume struct {
	// The resource-specific properties for this resource.
	Properties *VolumeProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// VolumeListResult - The response of a Volume list operation.
type VolumeListResult struct {
	// REQUIRED; The Volume items on this page
	Value []*Volume

	// The link to the next page of items
	NextLink *string
}

// VolumeProperties - Volume Properties
type VolumeProperties struct {
	// REQUIRED; Requested capacity in GiB
	CapacityGiB *int64

	// REQUIRED; String KV pairs indicating labels
	Labels map[string]*string

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The status of the resource.
	Status *ResourceOperationalStatus

	// READ-ONLY; Properties of the volume
	VolumeType *VolumeType
}

// VolumeType - Properties of the volume
type VolumeType struct {
	// READ-ONLY; Properties of the ElasticSAN iSCSI target
	ElasticSan *ElasticSanVolumeProperties
}

// VolumeUpdate - The type used for update operations of the Volume.
type VolumeUpdate struct {
	// The resource-specific properties for this resource.
	Properties *VolumeUpdateProperties
}

// VolumeUpdateProperties - The updatable properties of the Volume.
type VolumeUpdateProperties struct {
	// Requested capacity in GiB
	CapacityGiB *int64

	// String KV pairs indicating labels
	Labels map[string]*string
}
