//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsphere

import "time"

// Catalog - An Azure Sphere catalog
type Catalog struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *CatalogProperties

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

// CatalogListResult - The response of a Catalog list operation.
type CatalogListResult struct {
	// REQUIRED; The Catalog items on this page
	Value []*Catalog

	// The link to the next page of items
	NextLink *string
}

// CatalogProperties - Catalog properties
type CatalogProperties struct {
	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState
}

// CatalogUpdate - The type used for update operations of the Catalog.
type CatalogUpdate struct {
	// Resource tags.
	Tags map[string]*string
}

// Certificate - An certificate resource belonging to a catalog resource.
type Certificate struct {
	// The resource-specific properties for this resource.
	Properties *CertificateProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// CertificateChainResponse - The certificate chain response.
type CertificateChainResponse struct {
	// READ-ONLY; The certificate chain.
	CertificateChain *string
}

// CertificateListResult - The response of a Certificate list operation.
type CertificateListResult struct {
	// REQUIRED; The Certificate items on this page
	Value []*Certificate

	// The link to the next page of items
	NextLink *string
}

// CertificateProperties - The properties of certificate
type CertificateProperties struct {
	// READ-ONLY; The certificate as a UTF-8 encoded base 64 string.
	Certificate *string

	// READ-ONLY; The certificate expiry date.
	ExpiryUTC *time.Time

	// READ-ONLY; The certificate not before date.
	NotBeforeUTC *time.Time

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The certificate status.
	Status *CertificateStatus

	// READ-ONLY; The certificate subject.
	Subject *string

	// READ-ONLY; The certificate thumbprint.
	Thumbprint *string
}

// ClaimDevicesRequest - Request to the action call to bulk claim devices.
type ClaimDevicesRequest struct {
	// REQUIRED; Device identifiers of the devices to be claimed.
	DeviceIdentifiers []*string
}

// CountDeviceResponse - Response to the action call for count devices in a catalog (preview API).
type CountDeviceResponse struct {
	// REQUIRED; Number of children resources in parent resource.
	Value *int32
}

// Deployment - An deployment resource belonging to a device group resource.
type Deployment struct {
	// The resource-specific properties for this resource.
	Properties *DeploymentProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// DeploymentListResult - The response of a Deployment list operation.
type DeploymentListResult struct {
	// REQUIRED; The Deployment items on this page
	Value []*Deployment

	// The link to the next page of items
	NextLink *string
}

// DeploymentProperties - The properties of deployment
type DeploymentProperties struct {
	// Images deployed
	DeployedImages []*Image

	// Deployment ID
	DeploymentID *string

	// READ-ONLY; Deployment date UTC
	DeploymentDateUTC *time.Time

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState
}

// Device - An device resource belonging to a device group resource.
type Device struct {
	// The resource-specific properties for this resource.
	Properties *DeviceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// DeviceGroup - An device group resource belonging to a product resource.
type DeviceGroup struct {
	// The resource-specific properties for this resource.
	Properties *DeviceGroupProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// DeviceGroupListResult - The response of a DeviceGroup list operation.
type DeviceGroupListResult struct {
	// REQUIRED; The DeviceGroup items on this page
	Value []*DeviceGroup

	// The link to the next page of items
	NextLink *string
}

// DeviceGroupProperties - The properties of deviceGroup
type DeviceGroupProperties struct {
	// Flag to define if the user allows for crash dump collection.
	AllowCrashDumpsCollection *AllowCrashDumpCollection

	// Description of the device group.
	Description *string

	// Operating system feed type of the device group.
	OSFeedType *OSFeedType

	// Regional data boundary for the device group.
	RegionalDataBoundary *RegionalDataBoundary

	// Update policy of the device group.
	UpdatePolicy *UpdatePolicy

	// READ-ONLY; Deployment status for the device group.
	HasDeployment *bool

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState
}

// DeviceGroupUpdate - The type used for update operations of the DeviceGroup.
type DeviceGroupUpdate struct {
	// The updatable properties of the DeviceGroup.
	Properties *DeviceGroupUpdateProperties
}

// DeviceGroupUpdateProperties - The updatable properties of the DeviceGroup.
type DeviceGroupUpdateProperties struct {
	// Flag to define if the user allows for crash dump collection.
	AllowCrashDumpsCollection *AllowCrashDumpCollection

	// Description of the device group.
	Description *string

	// Operating system feed type of the device group.
	OSFeedType *OSFeedType

	// Regional data boundary for the device group.
	RegionalDataBoundary *RegionalDataBoundary

	// Update policy of the device group.
	UpdatePolicy *UpdatePolicy
}

// DeviceInsight - Device insight report.
type DeviceInsight struct {
	// REQUIRED; Event description
	Description *string

	// REQUIRED; Device ID
	DeviceID *string

	// REQUIRED; Event end timestamp
	EndTimestampUTC *time.Time

	// REQUIRED; Event category
	EventCategory *string

	// REQUIRED; Event class
	EventClass *string

	// REQUIRED; Event count
	EventCount *int32

	// REQUIRED; Event type
	EventType *string

	// REQUIRED; Event start timestamp
	StartTimestampUTC *time.Time
}

// DeviceListResult - The response of a Device list operation.
type DeviceListResult struct {
	// REQUIRED; The Device items on this page
	Value []*Device

	// The link to the next page of items
	NextLink *string
}

// DeviceProperties - The properties of device
type DeviceProperties struct {
	// Device ID
	DeviceID *string

	// READ-ONLY; SKU of the chip
	ChipSKU *string

	// READ-ONLY; OS version available for installation when update requested
	LastAvailableOsVersion *string

	// READ-ONLY; OS version running on device when update requested
	LastInstalledOsVersion *string

	// READ-ONLY; Time when update requested and new OS version available
	LastOsUpdateUTC *time.Time

	// READ-ONLY; Time when update was last requested
	LastUpdateRequestUTC *time.Time

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState
}

// DeviceUpdate - The type used for update operations of the Device.
type DeviceUpdate struct {
	// The updatable properties of the Device.
	Properties *DeviceUpdateProperties
}

// DeviceUpdateProperties - The updatable properties of the Device.
type DeviceUpdateProperties struct {
	// Device group id
	DeviceGroupID *string
}

// GenerateCapabilityImageRequest - Request of the action to create a signed device capability image
type GenerateCapabilityImageRequest struct {
	// REQUIRED; List of capabilities to create
	Capabilities []*CapabilityType
}

// Image - An image resource belonging to a catalog resource.
type Image struct {
	// The resource-specific properties for this resource.
	Properties *ImageProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ImageListResult - The response of a Image list operation.
type ImageListResult struct {
	// REQUIRED; The Image items on this page
	Value []*Image

	// The link to the next page of items
	NextLink *string
}

// ImageProperties - The properties of image
type ImageProperties struct {
	// Image as a UTF-8 encoded base 64 string on image create. This field contains the image URI on image reads.
	Image *string

	// Image ID
	ImageID *string

	// Regional data boundary for an image
	RegionalDataBoundary *RegionalDataBoundary

	// READ-ONLY; The image component id.
	ComponentID *string

	// READ-ONLY; The image description.
	Description *string

	// READ-ONLY; Image name
	ImageName *string

	// READ-ONLY; The image type.
	ImageType *ImageType

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState

	// READ-ONLY; Location the image
	URI *string
}

// ListDeviceGroupsRequest - Request of the action to list device groups for a catalog.
type ListDeviceGroupsRequest struct {
	// Device Group name.
	DeviceGroupName *string
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

// PagedDeviceInsight - Paged collection of DeviceInsight items
type PagedDeviceInsight struct {
	// REQUIRED; The DeviceInsight items on this page
	Value []*DeviceInsight

	// The link to the next page of items
	NextLink *string
}

// Product - An product resource belonging to a catalog resource.
type Product struct {
	// The resource-specific properties for this resource.
	Properties *ProductProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ProductListResult - The response of a Product list operation.
type ProductListResult struct {
	// REQUIRED; The Product items on this page
	Value []*Product

	// The link to the next page of items
	NextLink *string
}

// ProductProperties - The properties of product
type ProductProperties struct {
	// REQUIRED; Description of the product
	Description *string

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState
}

// ProductUpdate - The type used for update operations of the Product.
type ProductUpdate struct {
	// The updatable properties of the Product.
	Properties *ProductUpdateProperties
}

// ProductUpdateProperties - The updatable properties of the Product.
type ProductUpdateProperties struct {
	// Description of the product
	Description *string
}

// ProofOfPossessionNonceRequest - Request for the proof of possession nonce
type ProofOfPossessionNonceRequest struct {
	// REQUIRED; The proof of possession nonce
	ProofOfPossessionNonce *string
}

// ProofOfPossessionNonceResponse - Result of the action to generate a proof of possession nonce
type ProofOfPossessionNonceResponse struct {
	// READ-ONLY; The certificate as a UTF-8 encoded base 64 string.
	Certificate *string

	// READ-ONLY; The certificate expiry date.
	ExpiryUTC *time.Time

	// READ-ONLY; The certificate not before date.
	NotBeforeUTC *time.Time

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The certificate status.
	Status *CertificateStatus

	// READ-ONLY; The certificate subject.
	Subject *string

	// READ-ONLY; The certificate thumbprint.
	Thumbprint *string
}

// SignedCapabilityImageResponse - Signed device capability image response
type SignedCapabilityImageResponse struct {
	// READ-ONLY; The signed device capability image as a UTF-8 encoded base 64 string.
	Image *string
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
