//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armelasticsan

import "time"

// ElasticSan - Response for ElasticSan request.
type ElasticSan struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; Properties of ElasticSan.
	Properties *Properties

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

// EncryptionIdentity - Encryption identity for the volume group.
type EncryptionIdentity struct {
	// Resource identifier of the UserAssigned identity to be associated with server-side encryption on the volume group.
	EncryptionUserAssignedIdentity *string
}

// EncryptionProperties - The encryption settings on the volume group.
type EncryptionProperties struct {
	// The identity to be used with service-side encryption at rest.
	EncryptionIdentity *EncryptionIdentity

	// Properties provided by key vault.
	KeyVaultProperties *KeyVaultProperties
}

// Identity for the resource.
type Identity struct {
	// REQUIRED; The identity type.
	Type *IdentityType

	// Gets or sets a list of key value pairs that describe the set of User Assigned identities that will be used with this volume
	// group. The key is the ARM resource identifier of the identity.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The principal ID of resource identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of resource.
	TenantID *string
}

// IscsiTargetInfo - Iscsi target information
type IscsiTargetInfo struct {
	// Operational status of the iSCSI Target.
	Status *OperationalStatus

	// READ-ONLY; State of the operation on the resource.
	ProvisioningState *ProvisioningStates

	// READ-ONLY; iSCSI Target IQN (iSCSI Qualified Name); example: "iqn.2005-03.org.iscsi:server".
	TargetIqn *string

	// READ-ONLY; iSCSI Target Portal Host Name
	TargetPortalHostname *string

	// READ-ONLY; iSCSI Target Portal Port
	TargetPortalPort *int32
}

// KeyVaultProperties - Properties of key vault.
type KeyVaultProperties struct {
	// The name of KeyVault key.
	KeyName *string

	// The Uri of KeyVault.
	KeyVaultURI *string

	// The version of KeyVault key.
	KeyVersion *string

	// READ-ONLY; This is a read only property that represents the expiration time of the current version of the customer managed
	// key used for encryption.
	CurrentVersionedKeyExpirationTimestamp *time.Time

	// READ-ONLY; The object identifier of the current versioned Key Vault Key in use.
	CurrentVersionedKeyIdentifier *string

	// READ-ONLY; Timestamp of last rotation of the Key Vault Key.
	LastKeyRotationTimestamp *time.Time
}

// List of Elastic Sans
type List struct {
	// An array of Elastic San objects.
	Value []*ElasticSan

	// READ-ONLY; URI to fetch the next section of the paginated response.
	NextLink *string
}

// ManagedByInfo - Parent resource information.
type ManagedByInfo struct {
	// Resource ID of the resource managing the volume, this is a restricted field and can only be set for internal use.
	ResourceID *string
}

// NetworkRuleSet - A set of rules governing the network accessibility.
type NetworkRuleSet struct {
	// The list of virtual network rules.
	VirtualNetworkRules []*VirtualNetworkRule
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

// PrivateEndpoint - Response for PrivateEndpoint
type PrivateEndpoint struct {
	// READ-ONLY; The ARM identifier for Private Endpoint
	ID *string
}

// PrivateEndpointConnection - Response for PrivateEndpoint Connection object
type PrivateEndpointConnection struct {
	// REQUIRED; Private Endpoint Connection Properties.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateEndpointConnectionListResult - List of private endpoint connections associated with SAN
type PrivateEndpointConnectionListResult struct {
	// Array of private endpoint connections
	Value []*PrivateEndpointConnection

	// READ-ONLY; URI to fetch the next section of the paginated response.
	NextLink *string
}

// PrivateEndpointConnectionProperties - Response for PrivateEndpoint connection properties
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; Private Link Service Connection State.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState

	// List of resources private endpoint is mapped
	GroupIDs []*string

	// Private Endpoint resource
	PrivateEndpoint *PrivateEndpoint

	// READ-ONLY; Provisioning State of Private Endpoint connection resource
	ProvisioningState *ProvisioningStates
}

// PrivateLinkResource - A private link resource
type PrivateLinkResource struct {
	// Resource properties.
	Properties *PrivateLinkResourceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateLinkResourceListResult - A list of private link resources
type PrivateLinkResourceListResult struct {
	// Array of private link resources
	Value []*PrivateLinkResource

	// READ-ONLY; URI to fetch the next section of the paginated response.
	NextLink *string
}

// PrivateLinkResourceProperties - Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// The private link resource Private link DNS zone name.
	RequiredZoneNames []*string

	// READ-ONLY; The private link resource group id.
	GroupID *string

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string
}

// PrivateLinkServiceConnectionState - Response for Private Link Service Connection state
type PrivateLinkServiceConnectionState struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string

	// The reason for approval/rejection of the connection.
	Description *string

	// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus
}

// Properties - Elastic San response properties.
type Properties struct {
	// REQUIRED; Base size of the Elastic San appliance in TiB.
	BaseSizeTiB *int64

	// REQUIRED; Extended size of the Elastic San appliance in TiB.
	ExtendedCapacitySizeTiB *int64

	// REQUIRED; resource sku
	SKU *SKU

	// Logical zone for Elastic San resource; example: ["1"].
	AvailabilityZones []*string

	// Allow or disallow public network access to ElasticSan. Value is optional but if passed in, must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess *PublicNetworkAccess

	// READ-ONLY; The list of Private Endpoint Connections.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// READ-ONLY; State of the operation on the resource.
	ProvisioningState *ProvisioningStates

	// READ-ONLY; Total Provisioned IOPS of the Elastic San appliance.
	TotalIops *int64

	// READ-ONLY; Total Provisioned MBps Elastic San appliance.
	TotalMBps *int64

	// READ-ONLY; Total size of the Elastic San appliance in TB.
	TotalSizeTiB *int64

	// READ-ONLY; Total size of the provisioned Volumes in GiB.
	TotalVolumeSizeGiB *int64

	// READ-ONLY; Total number of volume groups in this Elastic San appliance.
	VolumeGroupCount *int64
}

// SKU - The SKU name. Required for account creation; optional for update.
type SKU struct {
	// REQUIRED; The sku name.
	Name *SKUName

	// The sku tier.
	Tier *SKUTier
}

// SKUCapability - The capability information in the specified SKU.
type SKUCapability struct {
	// READ-ONLY; The name of capability.
	Name *string

	// READ-ONLY; A string value to indicate states of given capability.
	Value *string
}

// SKUInformation - ElasticSAN SKU and its properties
type SKUInformation struct {
	// REQUIRED; Sku Name
	Name *SKUName

	// Sku Tier
	Tier *SKUTier

	// READ-ONLY; The capability information in the specified SKU.
	Capabilities []*SKUCapability

	// READ-ONLY; Availability of the SKU for the location/zone
	LocationInfo []*SKULocationInfo

	// READ-ONLY; The set of locations that the SKU is available. This will be supported and registered Azure Geo Regions (e.g.
	// West US, East US, Southeast Asia, etc.).
	Locations []*string

	// READ-ONLY; The type of the resource.
	ResourceType *string
}

// SKUInformationList - List of SKU Information objects
type SKUInformationList struct {
	// READ-ONLY; URI to fetch the next section of the paginated response.
	NextLink *string

	// READ-ONLY; List of ResourceType Sku
	Value []*SKUInformation
}

// SKULocationInfo - The location info.
type SKULocationInfo struct {
	// READ-ONLY; The location.
	Location *string

	// READ-ONLY; The zones.
	Zones []*string
}

// Snapshot - Response for Volume Snapshot request.
type Snapshot struct {
	// REQUIRED; Properties of Volume Snapshot.
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

// SnapshotCreationData - Data used when creating a volume snapshot.
type SnapshotCreationData struct {
	// REQUIRED; Fully qualified resource ID of the volume. E.g.
	// "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"
	SourceID *string
}

// SnapshotList - List of Snapshots
type SnapshotList struct {
	// An array of Snapshot objects.
	Value []*Snapshot

	// READ-ONLY; URI to fetch the next section of the paginated response.
	NextLink *string
}

// SnapshotProperties - Properties for Snapshot.
type SnapshotProperties struct {
	// REQUIRED; Data used when creating a volume snapshot.
	CreationData *SnapshotCreationData

	// READ-ONLY; State of the operation on the resource.
	ProvisioningState *ProvisioningStates

	// READ-ONLY; Size of Source Volume
	SourceVolumeSizeGiB *int64

	// READ-ONLY; Source Volume Name of a snapshot
	VolumeName *string
}

// SourceCreationData - Data source used when creating the volume.
type SourceCreationData struct {
	// This enumerates the possible sources of a volume creation.
	CreateSource *VolumeCreateOption

	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	SourceID *string
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

// Update - Response for ElasticSan update request.
type Update struct {
	// Properties of ElasticSan.
	Properties *UpdateProperties

	// Update tags
	Tags map[string]*string
}

// UpdateProperties - Elastic San update properties.
type UpdateProperties struct {
	// Base size of the Elastic San appliance in TiB.
	BaseSizeTiB *int64

	// Extended size of the Elastic San appliance in TiB.
	ExtendedCapacitySizeTiB *int64

	// Allow or disallow public network access to ElasticSan Account. Value is optional but if passed in, must be 'Enabled' or
	// 'Disabled'.
	PublicNetworkAccess *PublicNetworkAccess
}

// UserAssignedIdentity for the resource.
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the identity.
	ClientID *string

	// READ-ONLY; The principal ID of the identity.
	PrincipalID *string
}

// VirtualNetworkRule - Virtual Network rule.
type VirtualNetworkRule struct {
	// REQUIRED; Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.
	VirtualNetworkResourceID *string

	// The action of virtual network rule.
	Action *Action
}

// Volume - Response for Volume request.
type Volume struct {
	// REQUIRED; Properties of Volume.
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

// VolumeGroup - Response for Volume Group request.
type VolumeGroup struct {
	// The identity of the resource.
	Identity *Identity

	// Properties of VolumeGroup.
	Properties *VolumeGroupProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// VolumeGroupList - List of Volume Groups
type VolumeGroupList struct {
	// An array of Volume Groups objects.
	Value []*VolumeGroup

	// READ-ONLY; URI to fetch the next section of the paginated response.
	NextLink *string
}

// VolumeGroupProperties - VolumeGroup response properties.
type VolumeGroupProperties struct {
	// Type of encryption
	Encryption *EncryptionType

	// Encryption Properties describing Key Vault and Identity information
	EncryptionProperties *EncryptionProperties

	// A collection of rules governing the accessibility from specific network locations.
	NetworkACLs *NetworkRuleSet

	// Type of storage target
	ProtocolType *StorageTargetType

	// READ-ONLY; The list of Private Endpoint Connections.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// READ-ONLY; State of the operation on the resource.
	ProvisioningState *ProvisioningStates
}

// VolumeGroupUpdate - Volume Group request.
type VolumeGroupUpdate struct {
	// The identity of the resource.
	Identity *Identity

	// Properties of VolumeGroup.
	Properties *VolumeGroupUpdateProperties
}

// VolumeGroupUpdateProperties - VolumeGroup response properties.
type VolumeGroupUpdateProperties struct {
	// Type of encryption
	Encryption *EncryptionType

	// Encryption Properties describing Key Vault and Identity information
	EncryptionProperties *EncryptionProperties

	// A collection of rules governing the accessibility from specific network locations.
	NetworkACLs *NetworkRuleSet

	// Type of storage target
	ProtocolType *StorageTargetType
}

// VolumeList - List of Volumes
type VolumeList struct {
	// An array of Volume objects.
	Value []*Volume

	// READ-ONLY; URI to fetch the next section of the paginated response.
	NextLink *string
}

// VolumeProperties - Volume response properties.
type VolumeProperties struct {
	// REQUIRED; Volume size.
	SizeGiB *int64

	// State of the operation on the resource.
	CreationData *SourceCreationData

	// Parent resource information.
	ManagedBy *ManagedByInfo

	// READ-ONLY; State of the operation on the resource.
	ProvisioningState *ProvisioningStates

	// READ-ONLY; Storage target information
	StorageTarget *IscsiTargetInfo

	// READ-ONLY; Unique Id of the volume in GUID format
	VolumeID *string
}

// VolumeUpdate - Response for Volume request.
type VolumeUpdate struct {
	// Properties of Volume.
	Properties *VolumeUpdateProperties
}

// VolumeUpdateProperties - Volume response properties.
type VolumeUpdateProperties struct {
	// Parent resource information.
	ManagedBy *ManagedByInfo

	// Volume size.
	SizeGiB *int64
}
