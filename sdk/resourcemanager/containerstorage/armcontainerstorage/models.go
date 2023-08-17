//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerstorage

import "time"

// DiskPoolProperties - Disk Pool Properties
type DiskPoolProperties struct {
	// List of KV pairs to set in StorageClass to configure CSI driver.
	CsiParams map[string]*string

	// Only required if individual disk selection is desired. Path to disk, e.g. :/dev/sda or WWN. Supports specifying multiple
	// disks (same syntax as tags).
	Disks []*string

	// Maximum capacity of the volumes in GiB the user intends to create. Default 512.
	MaxVolumeCapacityGiB *int64
}

// ElasticSanPoolProperties - Elastic San Pool Properties
type ElasticSanPoolProperties struct {
	// REQUIRED; Resource group of an existing SAN.
	ResourceGroup *string

	// REQUIRED; Name of an existing SAN.
	SanName *string

	// REQUIRED; Volume group of an existing SAN.
	VolumeGroup *string
}

// ElasticSanPoolPropertiesUpdate - Elastic San Pool Properties
type ElasticSanPoolPropertiesUpdate struct {
	// Resource group of an existing SAN.
	ResourceGroup *string

	// Name of an existing SAN.
	SanName *string

	// Volume group of an existing SAN.
	VolumeGroup *string
}

// EphemeralPoolProperties - Ephemeral Pool Properties
type EphemeralPoolProperties struct {
	// REQUIRED; Template name or KV pairs containing disk selection criteria, e.g. model="Microsoft NVMe Direct Disk" to match
	// all Lsv2 NVMe disks.
	DiskSelector []*string

	// REQUIRED; Only required if individual disk selection is desired. Path to disk, e.g. :/dev/sda or WWN. Supports specifying
	// multiple disks (same syntax as tags).
	Disks []*string

	// Consent to format the local disks.
	DiskFormat *bool
}

// EphemeralPoolPropertiesUpdate - Ephemeral Pool Properties
type EphemeralPoolPropertiesUpdate struct {
	// Consent to format the local disks.
	DiskFormat *bool

	// Template name or KV pairs containing disk selection criteria, e.g. model="Microsoft NVMe Direct Disk" to match all Lsv2
	// NVMe disks.
	DiskSelector []*string

	// Only required if individual disk selection is desired. Path to disk, e.g. :/dev/sda or WWN. Supports specifying multiple
	// disks (same syntax as tags).
	Disks []*string
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
	// REQUIRED; List of resources that should have access to the pool. Typically ARM references to AKS clusters or ACI Container
	// Groups. For local and standard this must be a single reference. For portable there can
	// be many.
	Assignments []*string

	// REQUIRED; Elastic San Pool Properties
	ElasticSanPoolProperties *ElasticSanPoolProperties

	// REQUIRED; Initial capacity of the pool in GiB.
	PoolCapacityGiB *int64

	// REQUIRED; Type of the Pool: ephemeral, disk, managed, or elasticsan.
	PoolType *PoolType

	// REQUIRED; List of availability zones that resources can be created in.
	Zones []*string

	// Disk Pool Properties
	DiskPoolProperties *DiskPoolProperties

	// Ephemeral Pool Properties
	EphemeralPoolProperties *EphemeralPoolProperties

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState
}

// PoolUpdate - The type used for update operations of the Pool.
type PoolUpdate struct {
	// The updatable properties of the Pool.
	Properties *PoolUpdateProperties

	// Resource tags.
	Tags map[string]*string
}

// PoolUpdateProperties - The updatable properties of the Pool.
type PoolUpdateProperties struct {
	// List of resources that should have access to the pool. Typically ARM references to AKS clusters or ACI Container Groups.
	// For local and standard this must be a single reference. For portable there can
	// be many.
	Assignments []*string

	// Disk Pool Properties
	DiskPoolProperties *DiskPoolProperties

	// Elastic San Pool Properties
	ElasticSanPoolProperties *ElasticSanPoolPropertiesUpdate

	// Ephemeral Pool Properties
	EphemeralPoolProperties *EphemeralPoolPropertiesUpdate

	// Initial capacity of the pool in GiB.
	PoolCapacityGiB *int64

	// Type of the Pool: ephemeral, disk, managed, or elasticsan.
	PoolType *PoolType

	// List of availability zones that resources can be created in.
	Zones []*string
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

	// REQUIRED; List of string mount options
	MountOptions []*string

	// REQUIRED; Reclaim Policy, Delete or Retain
	ReclaimPolicy *ReclaimPolicy

	// REQUIRED; Indicates how the volume should be attached
	VolumeMode *VolumeMode

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState
}

// VolumeSnapshot - Concrete proxy resource types can be created by aliasing this type using a specific property type.
type VolumeSnapshot struct {
	// The resource-specific properties for this resource.
	Properties *VolumeSnapshotProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// VolumeSnapshotListResult - The response of a VolumeSnapshot list operation.
type VolumeSnapshotListResult struct {
	// REQUIRED; The VolumeSnapshot items on this page
	Value []*VolumeSnapshot

	// The link to the next page of items
	NextLink *string
}

// VolumeSnapshotProperties - Volume Snapshot Properties
type VolumeSnapshotProperties struct {
	// REQUIRED; List of string mount options
	MountOptions []*string

	// REQUIRED; Reclaim Policy, Delete or Retain
	ReclaimPolicy *ReclaimPolicy

	// REQUIRED; Reference to the source volume
	Source *string

	// REQUIRED; Indicates how the volumes created from the snapshot should be attached
	VolumeMode *VolumeMode

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState
}

// VolumeSnapshotUpdate - The type used for update operations of the VolumeSnapshot.
type VolumeSnapshotUpdate struct {
	// The updatable properties of the VolumeSnapshot.
	Properties *VolumeSnapshotUpdateProperties
}

// VolumeSnapshotUpdateProperties - The updatable properties of the VolumeSnapshot.
type VolumeSnapshotUpdateProperties struct {
	// List of string mount options
	MountOptions []*string

	// Reclaim Policy, Delete or Retain
	ReclaimPolicy *ReclaimPolicy

	// Reference to the source volume
	Source *string

	// Indicates how the volumes created from the snapshot should be attached
	VolumeMode *VolumeMode
}

// VolumeUpdate - The type used for update operations of the Volume.
type VolumeUpdate struct {
	// The updatable properties of the Volume.
	Properties *VolumeUpdateProperties
}

// VolumeUpdateProperties - The updatable properties of the Volume.
type VolumeUpdateProperties struct {
	// Requested capacity in GiB
	CapacityGiB *int64

	// String KV pairs indicating labels
	Labels map[string]*string

	// List of string mount options
	MountOptions []*string

	// Reclaim Policy, Delete or Retain
	ReclaimPolicy *ReclaimPolicy

	// Indicates how the volume should be attached
	VolumeMode *VolumeMode
}
